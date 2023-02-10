package tests

import (
	"github.com/adyen/adyen-go-api-library/v6/src/adyen"
	"github.com/adyen/adyen-go-api-library/v6/src/checkout"
	"github.com/adyen/adyen-go-api-library/v6/src/common"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"os"
	"testing"
)

type Tests struct {
	name          string
	server        *httptest.Server
	response      *checkout.PaymentCaptureResource
	expectedError error
}

func Test_API_Modifications(t *testing.T) {

	godotenv.Load("./../.env")

	var (
		MerchantAccount = os.Getenv("ADYEN_MERCHANT")
		APIKey          = os.Getenv("ADYEN_API_KEY")
	)

	client := adyen.NewClient(&common.Config{
		ApiKey:      APIKey,
		Environment: "TEST",
	})

	t.Run("API Modifications - Captures", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {

			pspReference := "psp0001"
			res, httpRes, err :=
				client.Checkout.CaptureAuthorisedPayment(&pspReference,
					&checkout.CreatePaymentCaptureRequest{
						MerchantAccount: MerchantAccount,
						Amount: checkout.Amount{
							Value:    1250,
							Currency: "EUR",
						},
					})

			require.NotNil(t, err)
			assert.Contains(t, err.Error(), "Original pspReference required for this operation")
			require.NotNil(t, httpRes)
			assert.Equal(t, 422, httpRes.StatusCode)
			require.NotNil(t, res)
		})
	})

	t.Run("API Modifications - Cancels", func(t *testing.T) {
		t.Run("Create an API request that should pass", func(t *testing.T) {

			res, httpRes, err :=
				client.Checkout.CancelAuthorisedPayment(
					&checkout.CreateStandalonePaymentCancelRequest{
						MerchantAccount:  MerchantAccount,
						PaymentReference: "paymentReference01",
						Reference:        "reference01",
					})

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 201, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, "received", res.Status)
		})
		t.Run("Create an API request that should fail", func(t *testing.T) {

			_, httpRes, err :=
				client.Checkout.CancelAuthorisedPayment(
					&checkout.CreateStandalonePaymentCancelRequest{
						MerchantAccount:  MerchantAccount,
						PaymentReference: "",
						Reference:        "reference01",
					})

			require.NotNil(t, err)
			assert.Contains(t, err.Error(), "Required field 'paymentReference' is not provided.")
			assert.Equal(t, 422, httpRes.StatusCode)
		})
	})

	t.Run("API Modifications - Refunds", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {

			pspReference := "psp0001"
			_, httpRes, err :=
				client.Checkout.RefundCapturedPayment(&pspReference,
					&checkout.CreatePaymentRefundRequest{
						MerchantAccount: MerchantAccount,
						Reference:       "reference01",
						Amount: checkout.Amount{
							Value:    1250,
							Currency: "EUR",
						},
					})

			require.NotNil(t, err)
			assert.Contains(t, err.Error(), "Original pspReference required")
			assert.Equal(t, 422, httpRes.StatusCode)
		})
	})

	t.Run("API Modifications - Reversals", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {

			pspReference := "psp0001"
			_, httpRes, err :=
				client.Checkout.RefundOrCancelPayment(&pspReference,
					&checkout.CreatePaymentReversalRequest{
						MerchantAccount: MerchantAccount,
						Reference:       "reference01",
					})

			require.NotNil(t, err)
			assert.Contains(t, err.Error(), "Original pspReference required")
			assert.Equal(t, 422, httpRes.StatusCode)
		})
	})

}
