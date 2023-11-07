//go:build integration
// +build integration

package tests

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/adyen/adyen-go-api-library/v8/src/adyen"
	"github.com/adyen/adyen-go-api-library/v8/src/common"
	"github.com/adyen/adyen-go-api-library/v8/src/payments"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Payment(t *testing.T) {
	godotenv.Load("./../.env")

	var (
		MerchantAccount = os.Getenv("ADYEN_MERCHANT")
		APIKey          = os.Getenv("ADYEN_API_KEY")
	)

	cvc := "737"
	expiryMonth := "03"
	number := "4111111111111111"
	card := payments.Card{
		Cvc:         &cvc,
		ExpiryMonth: &expiryMonth,
		ExpiryYear:  common.PtrString("2030"),
		HolderName:  common.PtrString("John Smith"),
		Number:      &number,
	}

	amount := payments.Amount{
		Currency: "EUR",
		Value:    1500,
	}

	client := adyen.NewClient(&common.Config{
		ApiKey:      APIKey,
		Environment: "TEST",
		Debug:       "true" == os.Getenv("DEBUG"),
	})
	service := client.Payments()

	assertForSuccessResponse := func(res interface{}, httpRes *http.Response, err error) {
		require.Nil(t, err)
		require.NotNil(t, httpRes)
		assert.Equal(t, 200, httpRes.StatusCode)
		require.NotNil(t, res)
	}

	authorisePost := func() (payments.PaymentResult, *http.Response, error) {
		body := payments.PaymentRequest{
			Card:            &card,
			Amount:          amount,
			Reference:       time.Now().String(),
			MerchantAccount: MerchantAccount,
		}
		req := service.PaymentsApi.AuthoriseInput().PaymentRequest(body)
		return service.PaymentsApi.Authorise(context.Background(), req)
	}

	t.Run("General", func(t *testing.T) {
		t.Run("Authorise3d", func(t *testing.T) {
			t.Skip("skipping since 3d requires manual user authentication")
			req := service.PaymentsApi.Authorise3dInput().PaymentRequest3d(payments.PaymentRequest3d{})
			res, httpRes, err := service.PaymentsApi.Authorise3d(context.Background(), req)
			assertForSuccessResponse(res, httpRes, err)
		})

		t.Run("Authorise3ds2", func(t *testing.T) {
			t.Skip("skipping since 3d requires manual user authentication")
			req := service.PaymentsApi.Authorise3ds2Input().PaymentRequest3ds2(payments.PaymentRequest3ds2{})
			res, httpRes, err := service.PaymentsApi.Authorise3ds2(context.Background(), req)
			assertForSuccessResponse(res, httpRes, err)
		})

		t.Run("Authorise", func(t *testing.T) {
			res, httpRes, err := authorisePost()

			assertForSuccessResponse(res, httpRes, err)
			assert.NotNil(t, res.PspReference)
			assert.Equal(t, res.GetResultCode(), "Authorised")
		})

		t.Run("GetAuthenticationResult", func(t *testing.T) {
			t.Skip("skipping since this returns auth result after a 3d auth")
			req := service.PaymentsApi.GetAuthenticationResultInput()
			res, httpRes, err := service.PaymentsApi.GetAuthenticationResult(context.Background(), req)
			assertForSuccessResponse(res, httpRes, err)
		})

		t.Run("Retrieve3ds2Result", func(t *testing.T) {
			t.Skip("skipping since this returns auth result after a 3d auth")
			req := service.PaymentsApi.Retrieve3ds2ResultInput()
			res, httpRes, err := service.PaymentsApi.Retrieve3ds2Result(context.Background(), req)
			assertForSuccessResponse(res, httpRes, err)
		})
	})

	t.Run("Modifications", func(t *testing.T) {
		t.Run("AdjustAuthorisation", func(t *testing.T) {
			authRes, _, _ := authorisePost()
			reference := time.Now().String()
			body := payments.AdjustAuthorisationRequest{
				OriginalReference: *authRes.PspReference,
				ModificationAmount: payments.Amount{
					Currency: "EUR",
					Value:    1234,
				},
				Reference:       &reference,
				MerchantAccount: MerchantAccount,
			}
			req := service.ModificationsApi.AdjustAuthorisationInput().AdjustAuthorisationRequest(body)
			res, httpRes, err := service.ModificationsApi.AdjustAuthorisation(context.Background(), req)
			assertForSuccessResponse(res, httpRes, err)
			assert.Equal(t, "[adjustAuthorisation-received]", res.Response)
		})

		t.Run("CancelOrRefund", func(t *testing.T) {
			authRes, _, _ := authorisePost()
			reference := time.Now().String()
			body := payments.CancelOrRefundRequest{
				OriginalReference: *authRes.PspReference,
				Reference:         &reference,
				MerchantAccount:   MerchantAccount,
			}
			req := service.ModificationsApi.CancelOrRefundInput().CancelOrRefundRequest(body)
			res, httpRes, err := service.ModificationsApi.CancelOrRefund(context.Background(), req)
			assertForSuccessResponse(res, httpRes, err)
			assert.Equal(t, "[cancelOrRefund-received]", res.Response)
		})

		t.Run("Cancel", func(t *testing.T) {
			authRes, _, _ := authorisePost()
			reference := time.Now().String()
			body := payments.CancelRequest{
				OriginalReference: *authRes.PspReference,
				Reference:         &reference,
				MerchantAccount:   MerchantAccount,
			}
			req := service.ModificationsApi.CancelInput().CancelRequest(body)
			res, httpRes, err := service.ModificationsApi.Cancel(context.Background(), req)
			assertForSuccessResponse(res, httpRes, err)
			assert.Equal(t, "[cancel-received]", res.Response)
		})

		t.Run("Capture", func(t *testing.T) {
			authRes, _, _ := authorisePost()
			reference := time.Now().String()
			body := payments.CaptureRequest{
				OriginalReference: *authRes.PspReference,
				ModificationAmount: payments.Amount{
					Currency: "EUR",
					Value:    1234,
				},
				Reference:       &reference,
				MerchantAccount: MerchantAccount,
			}
			req := service.ModificationsApi.CaptureInput().CaptureRequest(body)
			res, httpRes, err := service.ModificationsApi.Capture(context.Background(), req)
			assertForSuccessResponse(res, httpRes, err)
			assert.Equal(t, "[capture-received]", res.Response)
		})

		t.Run("Refund", func(t *testing.T) {
			authRes, _, _ := authorisePost()
			reference := time.Now().String()
			body := payments.RefundRequest{
				OriginalReference: *authRes.PspReference,
				ModificationAmount: payments.Amount{
					Currency: "EUR",
					Value:    1234,
				},
				Reference:       &reference,
				MerchantAccount: MerchantAccount,
			}
			req := service.ModificationsApi.RefundInput().RefundRequest(body)
			res, httpRes, err := service.ModificationsApi.Refund(context.Background(), req)
			assertForSuccessResponse(res, httpRes, err)
			assert.Equal(t, "[refund-received]", res.Response)
		})

		t.Run("TechnicalCancel", func(t *testing.T) {
			authRes, _, _ := authorisePost()
			reference := time.Now().String()
			body := payments.TechnicalCancelRequest{
				OriginalMerchantReference: *authRes.PspReference,
				Reference:                 &reference,
				MerchantAccount:           MerchantAccount,
			}
			req := service.ModificationsApi.TechnicalCancelInput().TechnicalCancelRequest(body)
			res, httpRes, err := service.ModificationsApi.TechnicalCancel(context.Background(), req)
			assertForSuccessResponse(res, httpRes, err)
			assert.Equal(t, "[technical-cancel-received]", res.Response)
		})

		t.Run("VoidPendingRefund", func(t *testing.T) {
			authRes, _, _ := authorisePost()
			reference := time.Now().String()
			body := payments.VoidPendingRefundRequest{
				OriginalReference: authRes.PspReference,
				Reference:         &reference,
				MerchantAccount:   MerchantAccount,
			}
			req := service.ModificationsApi.VoidPendingRefundInput().VoidPendingRefundRequest(body)
			res, httpRes, err := service.ModificationsApi.VoidPendingRefund(context.Background(), req)
			assertForSuccessResponse(res, httpRes, err)
			assert.Equal(t, "[voidPendingRefund-received]", res.Response)
		})
	})
}
