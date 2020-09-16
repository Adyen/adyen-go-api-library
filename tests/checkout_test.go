/*
 * Adyen API Client
 *
 * Contact: support@adyen.com
 */

package tests

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/google/uuid"

	"github.com/adyen/adyen-go-api-library/v3/src/adyen"
	"github.com/adyen/adyen-go-api-library/v3/src/checkout"
	"github.com/adyen/adyen-go-api-library/v3/src/common"

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
	//client.GetConfig().Debug = true

	t.Run("PaymentLinks", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {

			res, httpRes, err := client.Checkout.PaymentLinks(&checkout.CreatePaymentLinkRequest{
				Amount: checkout.Amount{
					Value:    1250,
					Currency: "EUR",
				},
				MerchantAccount: MerchantAccount,
			})

			require.NotNil(t, err)
			assert.Equal(t, true, strings.Contains(err.Error(), "Reference Missing"))
			require.NotNil(t, httpRes)
			require.NotNil(t, res)
		})
		t.Run("Create an API request that should pass", func(t *testing.T) {

			res, httpRes, err := client.Checkout.PaymentLinks(&checkout.CreatePaymentLinkRequest{
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

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, &checkout.Amount{Currency: "EUR", Value: 1250}, res.Amount)
			assert.NotNil(t, res.Url)
		})
	})

	t.Run("PaymentMethods", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {

			res, httpRes, err := client.Checkout.PaymentMethods(&checkout.PaymentMethodsRequest{})

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
			assert.Equal(t, true, strings.Contains(err.Error(), "Unsupported currency specified (validation: 138)"))
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
				PaymentMethod: map[string]interface{}{
					"type":   "ideal",
					"issuer": "1121",
				},
			}

			require.Nil(t, req.ApplicationInfo)

			res, httpRes, err := client.Checkout.Payments(req)

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, common.RedirectShopper, res.ResultCode)
			require.NotNil(t, res.Action)
			assert.Equal(t, "ideal", res.Action.PaymentMethodType)
			require.NotNil(t, res.PaymentData)

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
						Version: "v50",
					},
					AdyenLibrary: &checkout.CommonField{
						Name:    "test",
						Version: "v50",
					},
				},
			}

			res, httpRes, err := client.Checkout.Payments(req)

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, common.RedirectShopper, res.ResultCode)
			require.NotNil(t, res.Action)
			assert.Equal(t, "ideal", res.Action.PaymentMethodType)
			require.NotNil(t, res.PaymentData)

			// check if req has ApplicationInfo added to it
			require.NotNil(t, req.ApplicationInfo)
			require.NotNil(t, req.ApplicationInfo.AdyenLibrary)
			require.Equal(t, common.LibName, req.ApplicationInfo.AdyenLibrary.Name)
			require.Equal(t, common.LibVersion, req.ApplicationInfo.AdyenLibrary.Version)
			require.Equal(t, "test", req.ApplicationInfo.AdyenPaymentSource.Name)
			require.Equal(t, "v50", req.ApplicationInfo.AdyenPaymentSource.Version)
		})
	})

	t.Run("PaymentDetails", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			res, httpRes, err := client.Checkout.PaymentsDetails(&checkout.DetailsRequest{
				PaymentData: "1234",
				Details: map[string]interface{}{
					"MD":    "Ab02b4c0!BQABAgCW5sxB4e/==",
					"PaRes": "eNrNV0mTo7gS...",
				},
			})

			require.NotNil(t, err)
			assert.Equal(t, true, strings.Contains(err.Error(), "Invalid paymentData (validation: 14_003)"))
			require.NotNil(t, httpRes)
			assert.Equal(t, 422, httpRes.StatusCode)
			require.NotNil(t, res)
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
			assert.Equal(t, true, strings.Contains(err.Error(), "Token is missing"))
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
		})
	})
}
