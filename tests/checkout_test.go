package tests

import (
	"context"
	"encoding/json"
	"github.com/adyen/adyen-go-api-library/v7/src/adyen"
	"github.com/adyen/adyen-go-api-library/v7/src/checkout"
	"github.com/adyen/adyen-go-api-library/v7/src/common"
	"github.com/google/uuid"
	"io/ioutil"
	_nethttp "net/http"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Checkout(t *testing.T) {
	godotenv.Load("./../.env")

	MerchantAccount := os.Getenv("ADYEN_MERCHANT")

	client := adyen.NewClient(&common.Config{
		ApiKey:      os.Getenv("ADYEN_API_KEY"),
		Environment: "TEST",
		Debug:       "true" == os.Getenv("DEBUG"),
	})
	service := client.Checkout()

	t.Run("Configuration", func(t *testing.T) {
		liveClient := adyen.NewClient(&common.Config{
			ApiKey:                "LIVE_API_KEY",
			Environment:           common.LiveEnv,
			LiveEndpointURLPrefix: "abc123",
			Debug:                 false,
		})
		require.Equal(t, "https://abc123-checkout-live.adyenpayments.com/checkout/v70", liveClient.Checkout().PaymentsApi.BasePath())
	})

	t.Run("PaymentMethods", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			req := service.PaymentsApi.PaymentMethodsInput().
				PaymentMethodsRequest(checkout.PaymentMethodsRequest{})
			res, httpRes, err := service.PaymentsApi.PaymentMethods(context.Background(), req)

			require.NotNil(t, res)
			require.NotNil(t, httpRes)
			require.NotNil(t, err)
			assert.Equal(t, 403, httpRes.StatusCode)
			assert.Equal(t, "403 : Invalid Merchant Account (security: 901)", err.Error())
		})

		t.Run("Create an API request that should pass", func(t *testing.T) {
			paymentMethodsRequest := *checkout.NewPaymentMethodsRequest(MerchantAccount)
			req := service.PaymentsApi.PaymentMethodsInput().PaymentMethodsRequest(paymentMethodsRequest)
			res, httpRes, err := service.PaymentsApi.PaymentMethods(context.Background(), req)

			require.NotNil(t, res)
			require.NotNil(t, httpRes)
			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			resBody, readErr := ioutil.ReadAll(httpRes.Body)
			require.NoError(t, readErr)
			assert.Contains(t, string(resBody), "paymentMethods")
		})
	})

	t.Run("Payments", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			res, httpRes, err := service.PaymentsApi.Payments(
				context.Background(),
				// missing payment method
				service.PaymentsApi.PaymentsInput().PaymentRequest(checkout.PaymentRequest{
					MerchantAccount: MerchantAccount,
				}),
			)

			require.NotNil(t, err)
			assert.Contains(t, err.Error(), "error calling MarshalJSON for type checkout.CheckoutPaymentMethod")
			require.Nil(t, httpRes)
			require.NotNil(t, res)
		})

		t.Run("Credit card payment", func(t *testing.T) {
			card := checkout.NewCardDetails()
			card.SetEncryptedCardNumber("test_4111111111111111")
			card.SetEncryptedExpiryMonth("test_03")
			card.SetEncryptedExpiryYear("test_2030")
			card.SetEncryptedSecurityCode("test_737")
			paymentRequest := *checkout.NewPaymentRequest(
				*checkout.NewAmount("EUR", int64(1234)),
				MerchantAccount,
				checkout.CardDetailsAsCheckoutPaymentMethod(card),
				"Reference_example",
				"ReturnUrl_example",
			)
			paymentRequest.SetCaptureDelayHours(0) // assert int zero value is sent

			req := service.PaymentsApi.PaymentsInput()
			req = req.PaymentRequest(paymentRequest)
			res, httpRes, err := service.PaymentsApi.Payments(context.Background(), req)

			require.NotNil(t, res)
			require.NotNil(t, httpRes)
			require.Nil(t, err)

			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, "Authorised", res.GetResultCode())
			require.False(t, res.HasAction())
			assert.NotEmpty(t, res.GetPspReference())
			assert.NotEmpty(t, res.GetResultCode())
		})

		t.Run("iDEAL payment", func(t *testing.T) {
			idempotencyKey := "b9c3947f-b282-4059-a645-56ddbbd2fef3"
			ctx := common.WithIdempotencyKey(context.Background(), idempotencyKey)
			ideal := checkout.NewIdealDetails("1121")
			paymentRequest := *checkout.NewPaymentRequest(
				*checkout.NewAmount("EUR", int64(1234)),
				MerchantAccount,
				checkout.IdealDetailsAsCheckoutPaymentMethod(ideal),
				"Reference_example",
				"ReturnUrl_example",
			)

			req := service.PaymentsApi.PaymentsInput().PaymentRequest(paymentRequest)
			res, httpRes, err := service.PaymentsApi.Payments(ctx, req)

			require.NotNil(t, res)
			require.NotNil(t, httpRes)
			require.Nil(t, err)

			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, "RedirectShopper", res.GetResultCode())
			require.True(t, res.HasAction())
			assert.Equal(t, "ideal", res.GetAction().CheckoutRedirectAction.GetPaymentMethodType())
			assert.NotEmpty(t, res.GetResultCode())
		})

		t.Run("Create two APIs requests that should be identical when using the same Idempotency Key", func(t *testing.T) {
			card := checkout.NewCardDetails()
			card.SetEncryptedCardNumber("test_4111111111111111")
			card.SetEncryptedExpiryMonth("test_03")
			card.SetEncryptedExpiryYear("test_2030")
			card.SetEncryptedSecurityCode("test_737")
			card.SetHolderName("John Smith")
			body := checkout.PaymentRequest{
				Reference: "123456781235",
				Amount: checkout.Amount{
					Value:    1250,
					Currency: "EUR",
				},
				CountryCode:     common.PtrString("NL"),
				MerchantAccount: MerchantAccount,
				Channel:         common.PtrString("Web"),
				ReturnUrl:       "http://localhost:3000/redirect",
				PaymentMethod:   checkout.CardDetailsAsCheckoutPaymentMethod(card),
			}
			iKey := uuid.New().String()
			ctx := context.Background()

			req := service.PaymentsApi.PaymentsInput().IdempotencyKey(iKey).PaymentRequest(body)
			res, _, _ := service.PaymentsApi.Payments(ctx, req)
			pspRef := res.GetPspReference()

			res, _, _ = service.PaymentsApi.Payments(ctx, req)
			require.Equal(t, pspRef, res.GetPspReference())

			// Idempotency Key is not set for this request. Should have a new PspReference.
			req = service.PaymentsApi.PaymentsInput().PaymentRequest(body)
			res, _, _ = service.PaymentsApi.Payments(context.Background(), req)
			require.NotEqual(t, pspRef, res.GetPspReference())
		})
	})

	t.Run("PaymentDetails", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			req := service.PaymentsApi.PaymentsDetailsInput()
			req = req.DetailsRequest(checkout.DetailsRequest{
				PaymentData: common.PtrString("1234"),
				Details: checkout.PaymentCompletionDetails{
					MD:    common.PtrString("Ab02b4c0!BQABAgCW5sxB4e/=="),
					PaRes: common.PtrString("eNrNV0mTo7gS..."),
				},
			})
			res, httpRes, err := service.PaymentsApi.PaymentsDetails(context.Background(), req)

			require.NotNil(t, err)
			assert.Contains(t, err.Error(), "'paymentData' is not valid")
			require.NotNil(t, httpRes)
			assert.Equal(t, 422, httpRes.StatusCode)
			require.NotNil(t, res)
			require.Nil(t, res.ResultCode)
		})
	})

	t.Run("PaymentLinks", func(t *testing.T) {
		createPaymentLink := func() (checkout.PaymentLinkResponse, *_nethttp.Response, error) {
			req := service.PaymentLinksApi.PaymentLinksInput()
			req = req.CreatePaymentLinkRequest(checkout.CreatePaymentLinkRequest{
				Reference: "123456781235",
				Amount: checkout.Amount{
					Value:    1250,
					Currency: "EUR",
				},
				CountryCode:      common.PtrString("NL"),
				ShopperReference: common.PtrString("12345678"),
				ShopperEmail:     common.PtrString("test@email.com"),
				ShopperLocale:    common.PtrString("nl_NL"),
				BillingAddress: &checkout.Address{
					Street:            "Roque Petroni Jr",
					PostalCode:        "59000060",
					City:              "São Paulo",
					HouseNumberOrName: "999",
					Country:           "BR",
					StateOrProvince:   common.PtrString("SP"),
				},
				MerchantAccount: MerchantAccount,
			})
			return service.PaymentLinksApi.PaymentLinks(context.Background(), req)
		}

		t.Run("Create an API request that should fail", func(t *testing.T) {
			req := service.PaymentLinksApi.PaymentLinksInput()
			req = req.CreatePaymentLinkRequest(checkout.CreatePaymentLinkRequest{
				Amount: checkout.Amount{
					Value:    1250,
					Currency: "EUR",
				},
				MerchantAccount: MerchantAccount,
			})
			res, httpRes, err := service.PaymentLinksApi.PaymentLinks(context.Background(), req)

			require.NotNil(t, err)
			assert.Contains(t, err.Error(), "'reference' is not provided")
			require.NotNil(t, httpRes)
			require.NotNil(t, res)
		})
		t.Run("Create an API request that should pass", func(t *testing.T) {
			res, httpRes, err := createPaymentLink()

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 201, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, checkout.Amount{Currency: "EUR", Value: 1250}, res.Amount)
			assert.NotNil(t, res.Url)
		})

		t.Run("Get payment link", func(t *testing.T) {
			paymentLink, _, _ := createPaymentLink()
			req := service.PaymentLinksApi.GetPaymentLinkInput(paymentLink.Id)
			res, httpRes, err := service.PaymentLinksApi.GetPaymentLink(context.Background(), req)

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, paymentLink.Reference, res.Reference)
			assert.Equal(t, paymentLink.Status, res.Status)
			assert.NotNil(t, res.Url)
		})

		t.Run("Update payment link", func(t *testing.T) {
			paymentLink, _, _ := createPaymentLink()
			req := service.PaymentLinksApi.UpdatePaymentLinkInput(paymentLink.Id)
			req = req.UpdatePaymentLinkRequest(checkout.UpdatePaymentLinkRequest{
				Status: "expired",
			})
			res, httpRes, err := service.PaymentLinksApi.UpdatePaymentLink(context.Background(), req)

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, paymentLink.Reference, res.Reference)
			assert.NotEqual(t, paymentLink.Status, res.Status)
			assert.NotNil(t, res.Url)
		})
	})

	t.Run("PaymentSession", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			req := service.ClassicCheckoutSDKApi.PaymentSessionInput()
			req = req.PaymentSetupRequest(checkout.PaymentSetupRequest{
				Reference: "123456781235",
				Amount: checkout.Amount{
					Value:    1250,
					Currency: "EUR",
				},
				CountryCode:     "NL",
				MerchantAccount: MerchantAccount,
				Channel:         common.PtrString("iOS"),
				ReturnUrl:       "http://localhost:3000/redirect",
			})
			res, httpRes, err := service.ClassicCheckoutSDKApi.PaymentSession(context.Background(), req)

			require.NotNil(t, err)
			assert.Contains(t, err.Error(), "'token' is not provided.")
			require.NotNil(t, httpRes)
			assert.Equal(t, 422, httpRes.StatusCode)
			require.NotNil(t, res)
		})
		t.Run("Create an API request that should pass", func(t *testing.T) {
			req := service.ClassicCheckoutSDKApi.PaymentSessionInput()
			req = req.PaymentSetupRequest(checkout.PaymentSetupRequest{
				Reference: "123456781235",
				Amount: checkout.Amount{
					Value:    1250,
					Currency: "EUR",
				},
				CountryCode:     "NL",
				MerchantAccount: MerchantAccount,
				Channel:         common.PtrString("Web"),
				ReturnUrl:       "http://localhost:3000/redirect",
				SdkVersion:      common.PtrString("1.9.5"),
			})
			res, httpRes, err := service.ClassicCheckoutSDKApi.PaymentSession(context.Background(), req)

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.NotEmpty(t, res.PaymentSession)
		})
	})

	t.Run("PaymentsResult", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			req := service.ClassicCheckoutSDKApi.VerifyPaymentResultInput()
			req = req.PaymentVerificationRequest(checkout.PaymentVerificationRequest{Payload: "dummyPayload"})
			res, httpRes, err := service.ClassicCheckoutSDKApi.VerifyPaymentResult(context.Background(), req)

			require.NotNil(t, err)
			assert.Equal(t, true, strings.Contains(err.Error(), "Invalid payload provided"))
			require.NotNil(t, httpRes)
			assert.Equal(t, 422, httpRes.StatusCode)
			assert.Equal(t, "Invalid payload provided", err.(common.APIError).Message)
			require.NotNil(t, res)

			// verify ServiceError2 includes PspReference
			require.NotNil(t, err.(common.APIError).RawBody)

			var serviceError checkout.ServiceError2
			json.Unmarshal(err.(common.APIError).RawBody, &serviceError)
			require.NotNil(t, serviceError)
			require.NotNil(t, serviceError.PspReference)
		})
	})

	t.Run("Utility", func(t *testing.T) {
		t.Run("Get origin keys", func(t *testing.T) {
			domain := "https://adyen.com"
			req := service.UtilityApi.OriginKeysInput()
			req = req.CheckoutUtilityRequest(checkout.CheckoutUtilityRequest{
				OriginDomains: []string{
					domain,
				},
			})

			res, httpRes, err := service.UtilityApi.OriginKeys(context.Background(), req)

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
			originKeys := res.GetOriginKeys()
			assert.NotEmpty(t, originKeys[domain])
		})
	})

	t.Run("Orders", func(t *testing.T) {
		t.Run("Get balance", func(t *testing.T) {
			// @TODO review this test/config
			t.Skip("Payment method not correctly configured in the backoffice")
			req := service.OrdersApi.GetBalanceOfGiftCardInput()
			req = req.CheckoutBalanceCheckRequest(checkout.CheckoutBalanceCheckRequest{
				MerchantAccount: MerchantAccount,
				PaymentMethod: map[string]string{
					"type":       "giftcard",
					"brand":      "givex",
					"number":     "603628672882001915092",
					"holderName": "balance EUR 100",
					"cvc":        "5754",
				},
			})
			res, httpRes, err := service.OrdersApi.GetBalanceOfGiftCard(context.Background(), req)

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, int64(100), res.Balance.Value)
		})

		t.Run("Create order", func(t *testing.T) {
			req := service.OrdersApi.OrdersInput()
			req = req.CheckoutCreateOrderRequest(checkout.CheckoutCreateOrderRequest{
				Amount: checkout.Amount{
					Currency: "EUR",
					Value:    1000,
				},
				MerchantAccount: MerchantAccount,
				Reference:       "CREATE_ORDER_REF",
			})
			res, httpRes, err := service.OrdersApi.Orders(context.Background(), req)

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, int64(1000), res.RemainingAmount.Value)
		})

		t.Run("Cancel order", func(t *testing.T) {
			req := service.OrdersApi.OrdersInput()
			req = req.CheckoutCreateOrderRequest(checkout.CheckoutCreateOrderRequest{
				Amount: checkout.Amount{
					Currency: "EUR",
					Value:    1000,
				},
				MerchantAccount: MerchantAccount,
				Reference:       "CREATE_ORDER_REF",
			})
			order, _, _ := service.OrdersApi.Orders(context.Background(), req)

			cancelReq := service.OrdersApi.CancelOrderInput()
			cancelReq = cancelReq.CheckoutCancelOrderRequest(checkout.CheckoutCancelOrderRequest{
				MerchantAccount: MerchantAccount,
				Order: checkout.EncryptedOrderData{
					OrderData:    order.OrderData,
					PspReference: order.GetPspReference(),
				},
			})
			res, httpRes, err := service.OrdersApi.CancelOrder(context.Background(), cancelReq)

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, "Received", res.ResultCode)
		})

		t.Run("Create an API request that should fail", func(t *testing.T) {
			req := service.OrdersApi.OrdersInput()
			// missing required "reference" field
			req = req.CheckoutCreateOrderRequest(checkout.CheckoutCreateOrderRequest{
				Amount: checkout.Amount{
					Currency: "EUR",
					Value:    1000,
				},
				MerchantAccount: MerchantAccount,
			})
			res, httpRes, err := service.OrdersApi.Orders(context.Background(), req)

			require.NotNil(t, err)
			assert.Equal(t, "validation", err.(common.APIError).Type)
			require.NotNil(t, httpRes)
			assert.Equal(t, 422, httpRes.StatusCode)
			require.NotNil(t, res)
		})
	})

	t.Run("Sessions", func(t *testing.T) {
		t.Run("Create an API request that should pass", func(t *testing.T) {
			body := checkout.CreateCheckoutSessionRequest{
				Reference: "123456781235",
				Amount: checkout.Amount{
					Value:    1250,
					Currency: "EUR",
				},
				CountryCode:     common.PtrString("NL"),
				MerchantAccount: MerchantAccount,
				Channel:         common.PtrString("Web"),
				ReturnUrl:       "http://localhost:3000/redirect",
			}
			req := service.PaymentsApi.SessionsInput().CreateCheckoutSessionRequest(body)

			res, httpRes, err := service.PaymentsApi.Sessions(context.Background(), req)

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 201, httpRes.StatusCode)
			require.NotNil(t, res)
		})

		t.Run("Create an API request that should fail", func(t *testing.T) {
			body := checkout.CreateCheckoutSessionRequest{}
			req := service.PaymentsApi.SessionsInput().CreateCheckoutSessionRequest(body)

			res, httpRes, err := service.PaymentsApi.Sessions(context.Background(), req)

			require.NotNil(t, err)
			assert.Equal(t, true, strings.Contains(err.Error(), "Required field 'reference' is not provided. (validation: 130)"))
			require.NotNil(t, httpRes)
			assert.Equal(t, 422, httpRes.StatusCode)
			require.NotNil(t, res)
		})
	})

	t.Run("Card Details", func(t *testing.T) {
		t.Run("Create an API request that should pass", func(t *testing.T) {
			body := checkout.CardDetailsRequest{
				CardNumber:      "37000000",
				CountryCode:     common.PtrString("NL"),
				MerchantAccount: MerchantAccount,
			}
			req := service.PaymentsApi.CardDetailsInput().CardDetailsRequest(body)

			res, httpRes, err := service.PaymentsApi.CardDetails(context.Background(), req)

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, "amex", res.GetBrands()[0].GetType())
		})

		t.Run("Create an API request that should fail", func(t *testing.T) {
			body := checkout.CardDetailsRequest{
				CardNumber:      "3700",
				CountryCode:     common.PtrString("NL"),
				MerchantAccount: MerchantAccount,
			}
			req := service.PaymentsApi.CardDetailsInput().CardDetailsRequest(body)

			res, httpRes, err := service.PaymentsApi.CardDetails(context.Background(), req)

			require.NotNil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 422, httpRes.StatusCode)
			require.NotNil(t, res)
		})
	})
}
