//go:build integration
// +build integration

package checkout

import (
	"context"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/adyen/adyen-go-api-library/v21/src/adyen"
	"github.com/adyen/adyen-go-api-library/v21/src/checkout"
	"github.com/adyen/adyen-go-api-library/v21/src/common"
	"github.com/google/uuid"

	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCheckoutIntegration(t *testing.T) {
	godotenv.Load("./../../.env")

	merchantAccount := os.Getenv("ADYEN_MERCHANT")

	client := adyen.NewClient(&common.Config{
		ApiKey:      os.Getenv("ADYEN_API_KEY"),
		Environment: "TEST",
		Debug:       "true" == os.Getenv("DEBUG"),
	})
	service := client.Checkout()

	t.Run("PaymentMethods", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			req := service.PaymentsApi.PaymentMethodsInput().
				PaymentMethodsRequest(checkout.PaymentMethodsRequest{})
			res, httpRes, err := service.PaymentsApi.PaymentMethods(context.Background(), req)

			require.NotNil(t, res)
			require.NotNil(t, httpRes)
			require.NotNil(t, err)
			assert.Equal(t, 403, httpRes.StatusCode)
			assert.Equal(t, "403 Forbidden: Invalid Merchant Account (security: 901)", err.Error())
		})

		t.Run("Create an API request that should pass", func(t *testing.T) {
			paymentMethodsRequest := *checkout.NewPaymentMethodsRequest(merchantAccount)
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
		t.Run("Credit card payment", func(t *testing.T) {
			card := checkout.NewCardDetails()
			card.SetEncryptedCardNumber("test_4111111111111111")
			card.SetEncryptedExpiryMonth("test_03")
			card.SetEncryptedExpiryYear("test_2030")
			card.SetEncryptedSecurityCode("test_737")
			paymentRequest := *checkout.NewPaymentRequest(
				*checkout.NewAmount("EUR", int64(1234)),
				merchantAccount,
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
			idempotencyKey := uuid.New().String()
			ctx := common.WithIdempotencyKey(context.Background(), idempotencyKey)
			ideal := checkout.NewIdealDetails()
			paymentRequest := *checkout.NewPaymentRequest(
				*checkout.NewAmount("EUR", int64(1234)),
				merchantAccount,
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
	})

	t.Run("PaymentDetails", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			req := service.PaymentsApi.PaymentsDetailsInput()
			req = req.PaymentDetailsRequest(checkout.PaymentDetailsRequest{
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
		t.Run("Create an API request that should fail", func(t *testing.T) {
			req := service.PaymentLinksApi.PaymentLinksInput()
			req = req.PaymentLinkRequest(checkout.PaymentLinkRequest{
				Amount: checkout.Amount{
					Value:    1250,
					Currency: "EUR",
				},
				MerchantAccount: merchantAccount,
			})
			res, httpRes, err := service.PaymentLinksApi.PaymentLinks(context.Background(), req)

			require.NotNil(t, err)
			assert.Contains(t, err.Error(), "'reference' is not provided")
			require.NotNil(t, httpRes)
			require.NotNil(t, res)
		})

		t.Run("Create an API request that should pass", func(t *testing.T) {
			req := service.PaymentLinksApi.PaymentLinksInput()
			req = req.PaymentLinkRequest(checkout.PaymentLinkRequest{
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
				MerchantAccount: merchantAccount,
			})
			res, httpRes, err := service.PaymentLinksApi.PaymentLinks(context.Background(), req)

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 201, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, checkout.Amount{Currency: "EUR", Value: 1250}, res.Amount)
			assert.NotNil(t, res.Url)
		})
	})

	t.Run("Utility", func(t *testing.T) {
		t.Run("Get origin keys", func(t *testing.T) {
			domain := "https://adyen.com"
			req := service.UtilityApi.OriginKeysInput()
			req = req.UtilityRequest(checkout.UtilityRequest{
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
		t.Run("Create an API request that should fail", func(t *testing.T) {
			req := service.OrdersApi.OrdersInput()
			// missing required "reference" field
			req = req.CreateOrderRequest(checkout.CreateOrderRequest{
				Amount: checkout.Amount{
					Currency: "EUR",
					Value:    1000,
				},
				MerchantAccount: merchantAccount,
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
				MerchantAccount: merchantAccount,
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
			t.Skip("Disabling test as the endpoint behavior has changed and needs re-evaluation.")
			// Body without Reference
			body := checkout.CreateCheckoutSessionRequest{
				Amount: checkout.Amount{
					Value:    1250,
					Currency: "EUR",
				},
				CountryCode:     common.PtrString("NL"),
				MerchantAccount: merchantAccount,
				Channel:         common.PtrString("Web"),
				ReturnUrl:       "http://localhost:3000/redirect",
			}
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
				MerchantAccount: merchantAccount,
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
				MerchantAccount: merchantAccount,
			}
			req := service.PaymentsApi.CardDetailsInput().CardDetailsRequest(body)

			res, httpRes, err := service.PaymentsApi.CardDetails(context.Background(), req)

			require.NotNil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 422, httpRes.StatusCode)
			require.NotNil(t, res)

			// cast to API Error
			e := err.(common.APIError)

			require.NotNil(t, e.RawBody)

			assert.Equal(t, float64(422), e.Status)
			assert.Equal(t, "validation", e.Type)
			assert.Equal(t, "208", e.Code)
		})
	})
}
