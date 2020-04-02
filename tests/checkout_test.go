/*
 * Adyen API Client
 *
 * Contact: support@adyen.com
 */

package tests

import (
	"os"
	"testing"

	"github.com/adyen/adyen-go-api-library/src/checkout"

	adyen "github.com/adyen/adyen-go-api-library/src/api"
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

	client := adyen.NewAPIClientWithAPIKey(APIKey, "TEST")

	t.Run("PaymentLinks", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {

			res, httpRes, err := client.Checkout.PaymentLinksPost(&checkout.CreatePaymentLinkRequest{
				Amount: checkout.Amount{
					Value:    1250,
					Currency: "EUR",
				},
				MerchantAccount: MerchantAccount,
			})

			require.NotNil(t, err)
			assert.Equal(t, "422 Unprocessable Entity", err.Error())
			require.NotNil(t, httpRes)
			assert.Equal(t, 422, httpRes.StatusCode)
			assert.Equal(t, "422 Unprocessable Entity", httpRes.Status)
			require.NotNil(t, res)
		})
		t.Run("Create an API request that should pass", func(t *testing.T) {

			res, httpRes, err := client.Checkout.PaymentLinksPost(&checkout.CreatePaymentLinkRequest{
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
			assert.Equal(t, "200 OK", httpRes.Status)
			require.NotNil(t, res)
			assert.Equal(t, &checkout.Amount{Currency: "EUR", Value: 1250}, res.Amount)
			assert.NotNil(t, res.Url)
		})
	})

	t.Run("PaymentMethods", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {

			res, httpRes, err := client.Checkout.PaymentMethodsPost(&checkout.PaymentMethodsRequest{})

			require.NotNil(t, err)
			assert.Equal(t, "403 Forbidden", err.Error())
			require.NotNil(t, httpRes)
			assert.Equal(t, 403, httpRes.StatusCode)
			assert.Equal(t, "403 Forbidden", httpRes.Status)
			require.NotNil(t, res)
		})
		t.Run("Create an API request that should pass", func(t *testing.T) {

			res, httpRes, err := client.Checkout.PaymentMethodsPost(&checkout.PaymentMethodsRequest{
				MerchantAccount: MerchantAccount,
			})

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, "200 OK", httpRes.Status)
			require.NotNil(t, res)
			assert.True(t, len(*res.Groups) > 1)
			assert.True(t, len(*res.PaymentMethods) > 1)
		})
	})

	t.Run("Payments", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {

			res, httpRes, err := client.Checkout.PaymentsPost(&checkout.PaymentRequest{
				MerchantAccount: MerchantAccount,
			})

			require.NotNil(t, err)
			assert.Equal(t, "422 Unprocessable Entity", err.Error())
			require.NotNil(t, httpRes)
			assert.Equal(t, 422, httpRes.StatusCode)
			assert.Equal(t, "422 Unprocessable Entity", httpRes.Status)
			require.NotNil(t, res)
		})
		t.Run("Create an API request that should pass", func(t *testing.T) {

			res, httpRes, err := client.Checkout.PaymentsPost(&checkout.PaymentRequest{
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
			})

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, "200 OK", httpRes.Status)
			require.NotNil(t, res)
			assert.Equal(t, "RedirectShopper", res.ResultCode)
			require.NotNil(t, res.Action)
			assert.Equal(t, "ideal", res.Action.PaymentMethodType)
			require.NotNil(t, res.PaymentData)
		})
	})

	t.Run("PaymentDetails", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			res, httpRes, err := client.Checkout.PaymentsDetailsPost(&checkout.DetailsRequest{
				PaymentData: "1234",
				Details: map[string]interface{}{
					"MD":    "Ab02b4c0!BQABAgCW5sxB4e/==",
					"PaRes": "eNrNV0mTo7gS...",
				},
			})

			require.NotNil(t, err)
			assert.Equal(t, "422 Unprocessable Entity", httpRes.Status)
			require.NotNil(t, httpRes)
			assert.Equal(t, 422, httpRes.StatusCode)
			assert.Equal(t, "422 Unprocessable Entity", httpRes.Status)
			require.NotNil(t, res)
		})
		t.Run("Create an API request that should pass", func(t *testing.T) {
			payRes, _, _ := client.Checkout.PaymentsPost(&checkout.PaymentRequest{
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
			})
			res, httpRes, err := client.Checkout.PaymentsDetailsPost(&checkout.DetailsRequest{
				PaymentData: payRes.PaymentData,
				Details: map[string]interface{}{
					"MD":    "Ab02b4c0!BQABAgCW5sxB4e/==",
					"PaRes": "eNrNV0mTo7gS...",
				},
			})

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, "200 OK", httpRes.Status)
			require.NotNil(t, res)
			assert.Equal(t, "Authorised", res.ResultCode)
		})
	})
}
