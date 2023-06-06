package tests

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/adyen/adyen-go-api-library/v7/src/adyen"
	"github.com/adyen/adyen-go-api-library/v7/src/common"
	"github.com/adyen/adyen-go-api-library/v7/src/payments"
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
		req := service.GeneralApi.AuthoriseConfig(context.Background()).PaymentRequest(body)
		return service.GeneralApi.Authorise(req)
	}

	t.Run("General", func(t *testing.T) {
		t.Run("Authorise3d", func(t *testing.T) {
			t.Skip("skipping since 3d requires manual user authentication")
			req := service.GeneralApi.Authorise3dConfig(context.Background()).PaymentRequest3d(payments.PaymentRequest3d{})
			res, httpRes, err := service.GeneralApi.Authorise3d(req)
			assertForSuccessResponse(res, httpRes, err)
		})

		t.Run("Authorise3ds2", func(t *testing.T) {
			t.Skip("skipping since 3d requires manual user authentication")
			req := service.GeneralApi.Authorise3ds2Config(context.Background()).PaymentRequest3ds2(payments.PaymentRequest3ds2{})
			res, httpRes, err := service.GeneralApi.Authorise3ds2(req)
			assertForSuccessResponse(res, httpRes, err)
		})

		t.Run("Authorise", func(t *testing.T) {
			resultCode := common.Authorised.String()
			res, httpRes, err := authorisePost()

			assertForSuccessResponse(res, httpRes, err)
			assert.NotNil(t, res.PspReference)
			assert.Equal(t, res.ResultCode, &resultCode)
		})

		t.Run("GetAuthenticationResult", func(t *testing.T) {
			t.Skip("skipping since this returns auth result after a 3d auth")
			req := service.GeneralApi.GetAuthenticationResultConfig(context.Background())
			res, httpRes, err := service.GeneralApi.GetAuthenticationResult(req)
			assertForSuccessResponse(res, httpRes, err)
		})

		t.Run("Retrieve3ds2Result", func(t *testing.T) {
			t.Skip("skipping since this returns auth result after a 3d auth")
			req := service.GeneralApi.Retrieve3ds2ResultConfig(context.Background())
			res, httpRes, err := service.GeneralApi.Retrieve3ds2Result(req)
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
			req := service.ModificationsApi.AdjustAuthorisationConfig(context.Background()).AdjustAuthorisationRequest(body)
			res, httpRes, err := service.ModificationsApi.AdjustAuthorisation(req)
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
			req := service.ModificationsApi.CancelOrRefundConfig(context.Background()).CancelOrRefundRequest(body)
			res, httpRes, err := service.ModificationsApi.CancelOrRefund(req)
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
			req := service.ModificationsApi.CancelConfig(context.Background()).CancelRequest(body)
			res, httpRes, err := service.ModificationsApi.Cancel(req)
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
			req := service.ModificationsApi.CaptureConfig(context.Background()).CaptureRequest(body)
			res, httpRes, err := service.ModificationsApi.Capture(req)
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
			req := service.ModificationsApi.RefundConfig(context.Background()).RefundRequest(body)
			res, httpRes, err := service.ModificationsApi.Refund(req)
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
			req := service.ModificationsApi.TechnicalCancelConfig(context.Background()).TechnicalCancelRequest(body)
			res, httpRes, err := service.ModificationsApi.TechnicalCancel(req)
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
			req := service.ModificationsApi.VoidPendingRefundConfig(context.Background()).VoidPendingRefundRequest(body)
			res, httpRes, err := service.ModificationsApi.VoidPendingRefund(req)
			assertForSuccessResponse(res, httpRes, err)
			assert.Equal(t, "[voidPendingRefund-received]", res.Response)
		})
	})
}
