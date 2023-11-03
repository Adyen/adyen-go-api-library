//go:build integration
// +build integration

package tests

import (
	"context"
	"github.com/adyen/adyen-go-api-library/v8/src/adyen"
	"github.com/adyen/adyen-go-api-library/v8/src/checkout"
	"github.com/adyen/adyen-go-api-library/v8/src/common"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

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
	service := client.Checkout()

	t.Run("API Modifications - Captures", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			pspReference := "psp0001"
			req := service.ModificationsApi.CaptureAuthorisedPaymentInput(pspReference)
			req = req.PaymentCaptureRequest(checkout.PaymentCaptureRequest{
				MerchantAccount: MerchantAccount,
				Amount: checkout.Amount{
					Value:    1250,
					Currency: "EUR",
				},
			})
			res, httpRes, err := service.ModificationsApi.CaptureAuthorisedPayment(context.Background(), req)

			require.NotNil(t, err)
			assert.Contains(t, err.Error(), "Original pspReference required for this operation")
			require.NotNil(t, httpRes)
			assert.Equal(t, 422, httpRes.StatusCode)
			require.NotNil(t, res)
		})
	})

	t.Run("API Modifications - Cancels", func(t *testing.T) {
		t.Run("Create an API request that should pass", func(t *testing.T) {
			req := service.ModificationsApi.CancelAuthorisedPaymentInput()
			req = req.StandalonePaymentCancelRequest(checkout.StandalonePaymentCancelRequest{
				MerchantAccount:  MerchantAccount,
				PaymentReference: "paymentReference01",
				Reference:        common.PtrString("reference01"),
			})
			res, httpRes, err := service.ModificationsApi.CancelAuthorisedPayment(context.Background(), req)

			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 201, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, "received", res.Status)
		})

		t.Run("Create an API request that should fail", func(t *testing.T) {
			req := service.ModificationsApi.CancelAuthorisedPaymentInput()
			req = req.StandalonePaymentCancelRequest(checkout.StandalonePaymentCancelRequest{
				MerchantAccount:  MerchantAccount,
				PaymentReference: "",
				Reference:        common.PtrString("reference01"),
			})
			_, httpRes, err := service.ModificationsApi.CancelAuthorisedPayment(context.Background(), req)

			require.NotNil(t, err)
			assert.Contains(t, err.Error(), "Required field 'paymentReference' is not provided.")
			assert.Equal(t, 422, httpRes.StatusCode)
		})
	})

	t.Run("API Modifications - Refunds", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			pspReference := "psp0001"
			req := service.ModificationsApi.RefundCapturedPaymentInput(pspReference)
			req = req.PaymentRefundRequest(checkout.PaymentRefundRequest{
				MerchantAccount: MerchantAccount,
				Reference:       common.PtrString("reference01"),
				Amount: checkout.Amount{
					Value:    1250,
					Currency: "EUR",
				},
			})
			_, httpRes, err := service.ModificationsApi.RefundCapturedPayment(context.Background(), req)

			require.NotNil(t, err)
			assert.Contains(t, err.Error(), "Original pspReference required")
			assert.Equal(t, 422, httpRes.StatusCode)
		})
	})

	t.Run("API Modifications - Reversals", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			pspReference := "psp0001"
			req := service.ModificationsApi.RefundOrCancelPaymentInput(pspReference)
			req = req.PaymentReversalRequest(checkout.PaymentReversalRequest{
				MerchantAccount: MerchantAccount,
				Reference:       common.PtrString("reference01"),
			})
			_, httpRes, err := service.ModificationsApi.RefundOrCancelPayment(context.Background(), req)

			require.NotNil(t, err)
			assert.Contains(t, err.Error(), "Original pspReference required")
			assert.Equal(t, 422, httpRes.StatusCode)
		})
	})
}
