package tests

import (
	"context"
	"os"
	"testing"

	"github.com/adyen/adyen-go-api-library/v6/checkout"

	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Checkout_Next(t *testing.T) {
	godotenv.Load("./../.env")

	var (
		MerchantAccount = os.Getenv("ADYEN_MERCHANT")
		APIKey          = os.Getenv("ADYEN_API_KEY")
	)

	configuration := checkout.NewClientConfig()
	client := checkout.NewAPIClient(configuration)
	// client.GetConfig().Debug = true

	t.Run("PaymentMethods", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			paymentMethodsRequest := checkout.PaymentMethodsRequest{}

			res, httpRes, err := client.PaymentsApi.PaymentMethods(context.Background()).PaymentMethodsRequest(paymentMethodsRequest).Execute()

			require.Nil(t, res)
			require.NotNil(t, httpRes)
			require.NotNil(t, err)
			assert.Equal(t, 401, httpRes.StatusCode)
		})

		t.Run("Create an API request that should pass", func(t *testing.T) {
			paymentMethodsRequest := *checkout.NewPaymentMethodsRequest(MerchantAccount)

			contextKey := map[string]checkout.APIKey{"ApiKeyAuth": {Key: APIKey}}
			auth := context.WithValue(context.Background(), checkout.ContextAPIKeys, contextKey)

			res, httpRes, err := client.PaymentsApi.PaymentMethods(auth).PaymentMethodsRequest(paymentMethodsRequest).Execute()

			require.NotNil(t, res)
			require.NotNil(t, httpRes)
			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
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
				MerchantAccount,
				checkout.CardDetailsAsPaymentDonationRequestPaymentMethod(card),
				"Reference_example",
				"ReturnUrl_example",
			) // PaymentRequest |  (optional)
			paymentRequest.SetCaptureDelayHours(0) // assert int zero value is sent
			contextKey := map[string]checkout.APIKey{"ApiKeyAuth": {Key: APIKey}}
			auth := context.WithValue(context.Background(), checkout.ContextAPIKeys, contextKey)

			res, httpRes, err := client.PaymentsApi.Payments(auth).PaymentRequest(paymentRequest).Execute()

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
			ideal := checkout.NewIdealDetails("1121")
			paymentRequest := *checkout.NewPaymentRequest(
				*checkout.NewAmount("EUR", int64(1234)),
				MerchantAccount,
				checkout.IdealDetailsAsPaymentDonationRequestPaymentMethod(ideal),
				"Reference_example",
				"ReturnUrl_example",
			) // PaymentRequest |  (optional)
			contextKey := map[string]checkout.APIKey{"ApiKeyAuth": {Key: APIKey}}
			auth := context.WithValue(context.Background(), checkout.ContextAPIKeys, contextKey)

			res, httpRes, err := client.PaymentsApi.Payments(auth).PaymentRequest(paymentRequest).Execute()

			require.NotNil(t, res)
			require.NotNil(t, httpRes)
			require.Nil(t, err)

			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, "RedirectShopper", res.GetResultCode())
			require.True(t, res.HasAction())
			assert.Equal(t, "ideal", res.GetAction().CheckoutRedirectAction.PaymentMethodType)
			assert.NotEmpty(t, res.GetResultCode())
		})
	})
}
