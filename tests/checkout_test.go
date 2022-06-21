/*
 * Adyen API Client
 *
 * Contact: support@adyen.com
 */

package tests

import (
	"context"
	"encoding/json"
	_nethttp "net/http"
	"os"
	"strings"
	"testing"

	"github.com/google/uuid"

	"github.com/adyen/adyen-go-api-library/v5/src/adyen"
	"github.com/adyen/adyen-go-api-library/v5/src/checkout"
	"github.com/adyen/adyen-go-api-library/v5/src/common"

	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Checkout(t *testing.T) {
	godotenv.Load("./../.env")

	var (
		MerchantAccount = os.Getenv("ADYEN_MERCHANT")
		APIKey          = os.Getenv("ADYEN_API_KEY")
	)

	client := adyen.NewClient(&common.Config{
		ApiKey:      APIKey,
		Environment: "TEST",
	})
	// client.GetConfig().Debug = true

	t.Run("PaymentLinks", func(t *testing.T) {
		createPaymentLink := func() (checkout.PaymentLinkResponse, *_nethttp.Response, error) {
			return client.Checkout.PaymentLinks(&checkout.CreatePaymentLinkRequest{
				Reference: "123456781235",
				Amount: checkout.Amount{
					Value:    1250,
					Currency: "EUR",
				},
				CountryCode:      "NL",
				ShopperReference: "12345678",
				ShopperEmail:     "test@email.com",
				ShopperLocale:    "nl_NL",
				BillingAddress: &checkout.Address{
					Street:            "Roque Petroni Jr",
					PostalCode:        "59000060",
					City:              "São Paulo",
					HouseNumberOrName: "999",
					Country:           "BR",
					StateOrProvince:   "SP",
				},
				MerchantAccount: MerchantAccount,
			})
		}
		t.Run("Create an API request that should fail", func(t *testing.T) {
			res, httpRes, err := client.Checkout.PaymentLinks(&checkout.CreatePaymentLinkRequest{
				Amount: checkout.Amount{
					Value:    1250,
					Currency: "EUR",
				},
				MerchantAccount: MerchantAccount,
			})

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
			res, httpRes, err := client.Checkout.GetPaymentLink(paymentLink.Id)

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
			res, httpRes, err := client.Checkout.UpdatePaymentLink(paymentLink.Id, &checkout.UpdatePaymentLinkRequest{
				Status: "expired",
			})

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, paymentLink.Reference, res.Reference)
			assert.NotEqual(t, paymentLink.Status, res.Status)
			assert.NotNil(t, res.Url)
		})
	})

	t.Run("PaymentMethods", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			res, httpRes, err :=
				client.Checkout.PaymentMethods(&checkout.PaymentMethodsRequest{})

			require.NotNil(t, err)
			assert.Equal(t, true, strings.Contains(err.Error(), "Invalid Merchant Account (security: 901)"))
			require.NotNil(t, httpRes)
			assert.Equal(t, 403, httpRes.StatusCode)
			require.NotNil(t, res)
		})
		t.Run("Create an API request that should pass", func(t *testing.T) {
			res, httpRes, err := client.Checkout.PaymentMethods(&checkout.PaymentMethodsRequest{
				MerchantAccount: MerchantAccount,
			})

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
		})
	})

	t.Run("Payments", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			res, httpRes, err := client.Checkout.Payments(&checkout.PaymentRequest{
				MerchantAccount: MerchantAccount,
			})

			require.NotNil(t, err)
			assert.Contains(t, err.Error(), "'paymentMethod' is not provided")
			require.NotNil(t, httpRes)
			assert.Equal(t, 422, httpRes.StatusCode)
			require.NotNil(t, res)
		})
		t.Run("Create an API request that should pass", func(t *testing.T) {
			req := &checkout.PaymentRequest{
				Reference: "123456781235",
				Amount: checkout.Amount{
					Value:    1250,
					Currency: "EUR",
				},
				CountryCode:     "NL",
				MerchantAccount: MerchantAccount,
				Channel:         "Web",
				ReturnUrl:       "http://localhost:3000/redirect",
				PaymentMethod: checkout.IdealDetails{
					Type:   "ideal",
					Issuer: "1121",
				},
			}

			require.Nil(t, req.ApplicationInfo)

			res, httpRes, err := client.Checkout.Payments(req)

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, common.RedirectShopper, *res.ResultCode)
			require.NotNil(t, res.Action)

			// Make sure the actions is there
			redirectAction := res.Action.(*checkout.CheckoutRedirectAction)
			require.NotNil(t, redirectAction)
			require.NotNil(t, redirectAction.Url)
			require.Equal(t, "GET", redirectAction.Method)

			// check if req has ApplicationInfo added to it
			require.NotNil(t, req.ApplicationInfo)
			require.NotNil(t, req.ApplicationInfo.AdyenLibrary)
			require.Equal(t, common.LibName, req.ApplicationInfo.AdyenLibrary.Name)
			require.Equal(t, common.LibVersion, req.ApplicationInfo.AdyenLibrary.Version)
		})

		t.Run("Create two APIs requests that should be identical when using the same Idempotency Key", func(t *testing.T) {
			req := &checkout.PaymentRequest{
				Reference: "123456781235",
				Amount: checkout.Amount{
					Value:    1250,
					Currency: "EUR",
				},
				CountryCode:     "NL",
				MerchantAccount: MerchantAccount,
				Channel:         "Web",
				ReturnUrl:       "http://localhost:3000/redirect",
				PaymentMethod: map[string]interface{}{
					"type":        "scheme",
					"number":      "4111111111111111",
					"cvc":         "737",
					"expiryMonth": "03",
					"expiryYear":  "2030",
					"holderName":  "John Smith",
				},
			}

			require.Nil(t, req.ApplicationInfo)

			iKey := uuid.New().String()
			ctx := context.Background()
			ctx = common.WithIdempotencyKey(ctx, iKey)
			res, _, _ := client.Checkout.Payments(req, ctx)
			pspRef := res.PspReference

			ctx = common.WithIdempotencyKey(ctx, iKey)
			res, _, _ = client.Checkout.Payments(req, ctx)
			require.Equal(t, pspRef, res.PspReference)

			// Idempotency Key is not set for this request. Should have a new PspReference.
			res, _, _ = client.Checkout.Payments(req)
			require.NotEqual(t, pspRef, res.PspReference)
		})

		t.Run("Create an API request that should merge ApplicationInfo", func(t *testing.T) {
			req := &checkout.PaymentRequest{
				Reference: "123456781235",
				Amount: checkout.Amount{
					Value:    1250,
					Currency: "EUR",
				},
				CountryCode:     "NL",
				MerchantAccount: MerchantAccount,
				Channel:         "Web",
				ReturnUrl:       "http://localhost:3000/redirect",
				PaymentMethod: map[string]interface{}{
					"type":   "ideal",
					"issuer": "1121",
				},
				ApplicationInfo: &checkout.ApplicationInfo{
					AdyenPaymentSource: &checkout.CommonField{
						Name:    "test",
						Version: "v67",
					},
					AdyenLibrary: &checkout.CommonField{
						Name:    "test",
						Version: "v67",
					},
				},
			}

			res, httpRes, err := client.Checkout.Payments(req)

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, common.RedirectShopper, *res.ResultCode)
			require.NotNil(t, res.Action)

			// check if req has ApplicationInfo added to it
			require.NotNil(t, req.ApplicationInfo)
			require.NotNil(t, req.ApplicationInfo.AdyenLibrary)
			require.Equal(t, common.LibName, req.ApplicationInfo.AdyenLibrary.Name)
			require.Equal(t, common.LibVersion, req.ApplicationInfo.AdyenLibrary.Version)
			require.Equal(t, "test", req.ApplicationInfo.AdyenPaymentSource.Name)
			require.Equal(t, "v67", req.ApplicationInfo.AdyenPaymentSource.Version)
		})
	})

	t.Run("PaymentDetails", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			res, httpRes, err := client.Checkout.PaymentsDetails(&checkout.DetailsRequest{
				PaymentData: "1234",
				Details: checkout.PaymentCompletionDetails{
					MD:    "Ab02b4c0!BQABAgCW5sxB4e/==",
					PaRes: "eNrNV0mTo7gS...",
				},
			})

			require.NotNil(t, err)
			assert.Contains(t, err.Error(), "'paymentData' is not valid")
			require.NotNil(t, httpRes)
			assert.Equal(t, 422, httpRes.StatusCode)
			require.NotNil(t, res)
			require.Nil(t, res.ResultCode)
		})
	})

	t.Run("PaymentSession", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			res, httpRes, err := client.Checkout.PaymentSession(&checkout.PaymentSetupRequest{
				Reference: "123456781235",
				Amount: checkout.Amount{
					Value:    1250,
					Currency: "EUR",
				},
				CountryCode:     "NL",
				MerchantAccount: MerchantAccount,
				Channel:         "iOS",
				ReturnUrl:       "http://localhost:3000/redirect",
			})

			require.NotNil(t, err)
			assert.Contains(t, err.Error(), "'token' is not provided.")
			require.NotNil(t, httpRes)
			assert.Equal(t, 422, httpRes.StatusCode)
			require.NotNil(t, res)
		})
		t.Run("Create an API request that should pass", func(t *testing.T) {
			res, httpRes, err := client.Checkout.PaymentSession(&checkout.PaymentSetupRequest{
				Reference: "123456781235",
				Amount: checkout.Amount{
					Value:    1250,
					Currency: "EUR",
				},
				CountryCode:     "NL",
				MerchantAccount: MerchantAccount,
				Channel:         "Web",
				ReturnUrl:       "http://localhost:3000/redirect",
				SdkVersion:      "1.9.5",
			})

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.NotEmpty(t, res.PaymentSession)
		})
	})

	t.Run("PaymentsResult", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			res, httpRes, err := client.Checkout.PaymentsResult(&checkout.PaymentVerificationRequest{Payload: "dummyPayload"})
			
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
			res, httpRes, err := client.Checkout.OriginKeys(&checkout.CheckoutUtilityRequest{
				OriginDomains: []string{
					domain,
				},
			})
			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
			originKeys := res.OriginKeys
			assert.NotEmpty(t, originKeys[domain])
		})
	})

	t.Run("Orders", func(t *testing.T) {
		t.Run("Get balance", func(t *testing.T) {
			t.Skip("Payment method not correctly configured in the backoffice")
			res, httpRes, err := client.Checkout.PaymentMethodsBalance(&checkout.CheckoutBalanceCheckRequest{
				MerchantAccount: MerchantAccount,
				PaymentMethod: map[string]interface{}{
					"type":       "giftcard",
					"brand":      "givex",
					"number":     "603628672882001915092",
					"holderName": "balance EUR 100",
					"cvc":        "5754",
					"additionalAmount": map[string]interface{}{
						"currency": "EUR",
						"value":    0,
					},
				},
			})

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, int64(100), res.Balance.Value)
		})
		t.Run("Create order", func(t *testing.T) {
			res, httpRes, err := client.Checkout.Orders(&checkout.CheckoutCreateOrderRequest{
				Amount: checkout.Amount{
					Currency: "EUR",
					Value:    1000,
				},
				MerchantAccount: MerchantAccount,
				Reference:       "CREATE_ORDER_REF",
			})

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, int64(1000), res.RemainingAmount.Value)
		})
		t.Run("Cancel order", func(t *testing.T) {
			order, _, _ := client.Checkout.Orders(&checkout.CheckoutCreateOrderRequest{
				Amount: checkout.Amount{
					Currency: "EUR",
					Value:    1000,
				},
				MerchantAccount: MerchantAccount,
				Reference:       "CREATE_ORDER_REF",
			})

			res, httpRes, err := client.Checkout.OrdersCancel(&checkout.CheckoutCancelOrderRequest{
				MerchantAccount: MerchantAccount,
				Order: checkout.CheckoutOrder{
					OrderData:    order.OrderData,
					PspReference: order.PspReference,
				},
			})

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, "Received", res.ResultCode)
		})
		t.Run("Create an API request that should fail", func(t *testing.T) {
			res, httpRes, err := client.Checkout.Orders(&checkout.CheckoutCreateOrderRequest{
				Amount: checkout.Amount{
					Currency: "EUR",
					Value:    1000,
				},
				MerchantAccount: MerchantAccount,
				//Reference:       "CREATE_ORDER_REF", // create order without Reference
			})

			require.NotNil(t, err)
			assert.Equal(t, "validation", err.(common.APIError).Type)
			require.NotNil(t, httpRes)
			assert.Equal(t, 422, httpRes.StatusCode)
			require.NotNil(t, res)
		})
	})

	t.Run("Sessions", func(t *testing.T) {
		t.Run("Create an API request that should pass", func(t *testing.T) {

			req := &checkout.CreateCheckoutSessionRequest{
				Reference: "123456781235",
				Amount: checkout.Amount{
					Value:    1250,
					Currency: "EUR",
				},
				CountryCode:     "NL",
				MerchantAccount: MerchantAccount,
				Channel:         "Web",
				ReturnUrl:       "http://localhost:3000/redirect",
			}

			require.Nil(t, req.ApplicationInfo)

			res, httpRes, err :=
				client.Checkout.Sessions(req)

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 201, httpRes.StatusCode)
			require.NotNil(t, res)

			// check if req has ApplicationInfo added to it
			require.NotNil(t, req.ApplicationInfo)
			require.NotNil(t, req.ApplicationInfo.AdyenLibrary)
			require.Equal(t, common.LibName, req.ApplicationInfo.AdyenLibrary.Name)
			require.Equal(t, common.LibVersion, req.ApplicationInfo.AdyenLibrary.Version)

		})
		t.Run("Create an API request that should fail", func(t *testing.T) {
			res, httpRes, err :=
				client.Checkout.Sessions(&checkout.CreateCheckoutSessionRequest{})

			require.NotNil(t, err)
			assert.Equal(t, true, strings.Contains(err.Error(), "Invalid Merchant Account (validation: 901)"))
			require.NotNil(t, httpRes)
			assert.Equal(t, 422, httpRes.StatusCode)
			require.NotNil(t, res)
		})
	})

	t.Run("Card Details", func(t *testing.T) {
		t.Run("Create an API request that should pass", func(t *testing.T) {

			req := &checkout.CardDetailsRequest{
				CardNumber:     "37000000",				
				CountryCode:     "NL",
				MerchantAccount: MerchantAccount,
			}

			res, httpRes, err :=
				client.Checkout.CardDetails(req)

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, "amex", (*res.Brands)[0].Type)

		})
		t.Run("Create an API request that should fail", func(t *testing.T) {
			req := &checkout.CardDetailsRequest{
				CardNumber:     "3700",				
				CountryCode:     "NL",
				MerchantAccount: MerchantAccount,
			}

			res, httpRes, err :=
				client.Checkout.CardDetails(req)

			require.NotNil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 422, httpRes.StatusCode)
			require.NotNil(t, res)
		})
	})

}
