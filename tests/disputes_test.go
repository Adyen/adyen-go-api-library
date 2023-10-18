package tests

import (
	"context"
	"os"
	"testing"

	"github.com/adyen/adyen-go-api-library/v7/src/adyen"
	"github.com/adyen/adyen-go-api-library/v7/src/checkout"
	"github.com/adyen/adyen-go-api-library/v7/src/common"
	"github.com/adyen/adyen-go-api-library/v7/src/disputes"
	"github.com/adyen/adyen-go-api-library/v7/src/payments"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Disputes(t *testing.T) {
	godotenv.Load("./../.env")

	var (
		MerchantAccount = os.Getenv("ADYEN_MERCHANT")
		APIKey          = os.Getenv("ADYEN_API_KEY")
	)

	client := adyen.NewClient(&common.Config{
		ApiKey:      APIKey,
		Environment: "TEST",
		Debug:       true,
	})
	checkoutApi := client.Checkout()
	service := client.Disputes()

	createTestPayment := func() string {
		card := checkout.NewCardDetails()
		card.SetEncryptedCardNumber("test_4111111111111111")
		card.SetEncryptedExpiryMonth("test_03")
		card.SetEncryptedExpiryYear("test_2030")
		card.SetEncryptedSecurityCode("test_737")
		card.SetHolderName("chargeback:10.4")
		req := checkoutApi.PaymentsApi.PaymentsInput().PaymentRequest(checkout.PaymentRequest{
			Amount: checkout.Amount{
				Currency: "EUR",
				Value:    1000,
			},
			Reference:       "DISPUTES_CHARGEBACK",
			PaymentMethod:   checkout.CardDetailsAsCheckoutPaymentMethod(card),
			ReturnUrl:       "https://adyen.com",
			MerchantAccount: MerchantAccount,
		})
		res, _, _ := checkoutApi.PaymentsApi.Payments(context.Background(), req)

		reference := "MODIFICATION_REFERENCE"
		body := payments.CaptureRequest{
			OriginalReference: res.GetPspReference(),
			ModificationAmount: payments.Amount{
				Currency: "EUR",
				Value:    1000,
			},
			Reference:       &reference,
			MerchantAccount: MerchantAccount,
		}
		paymentsApi := client.Payments()
		captureReq := paymentsApi.ModificationsApi.CaptureInput().CaptureRequest(body)
		captureRes, _, _ := paymentsApi.ModificationsApi.Capture(context.Background(), captureReq)
		return captureRes.GetPspReference()
	}

	t.Run("Disputes", func(t *testing.T) {
		t.Run("Retrieve applicable defense reasons", func(t *testing.T) {
			pspReference := createTestPayment()
			req := service.RetrieveApplicableDefenseReasonsInput().DefenseReasonsRequest(disputes.DefenseReasonsRequest{
				DisputePspReference: pspReference,
				MerchantAccountCode: MerchantAccount,
			})
			res, httpRes, err := service.RetrieveApplicableDefenseReasons(context.Background(), req)

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.NotNil(t, res)
			assert.Equal(t, "Dispute not found.", res.DisputeServiceResult.GetErrorMessage())
		})

		t.Run("Supply defense document", func(t *testing.T) {
			pspReference := createTestPayment()
			req := service.SupplyDefenseDocumentInput().SupplyDefenseDocumentRequest(disputes.SupplyDefenseDocumentRequest{
				DefenseDocuments: []disputes.DefenseDocument{
					{
						Content:                 "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z8BQDwAEhQGAhKmMIQAAAABJRU5ErkJggg==",
						ContentType:             "image/jpg",
						DefenseDocumentTypeCode: "DefenseMaterial",
					},
				},
				DisputePspReference: pspReference,
				MerchantAccountCode: MerchantAccount,
			})
			res, httpRes, err := service.SupplyDefenseDocument(context.Background(), req)

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.NotNil(t, res)
			assert.Equal(t, "Unknown dispute", res.DisputeServiceResult.GetErrorMessage())
		})

		t.Run("Delete dispute defense document", func(t *testing.T) {
			pspReference := createTestPayment()
			req := service.DeleteDisputeDefenseDocumentInput().DeleteDefenseDocumentRequest(disputes.DeleteDefenseDocumentRequest{
				DefenseDocumentType: "DefenseMaterial",
				DisputePspReference: pspReference,
				MerchantAccountCode: MerchantAccount,
			})
			res, httpRes, err := service.DeleteDisputeDefenseDocument(context.Background(), req)

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.NotNil(t, res)
			assert.Equal(t, "Unknown dispute", res.DisputeServiceResult.GetErrorMessage())
		})

		t.Run("Defend dispute", func(t *testing.T) {
			pspReference := createTestPayment()
			req := service.DefendDisputeInput().DefendDisputeRequest(disputes.DefendDisputeRequest{
				DefenseReasonCode:   "DuplicateChargeback",
				DisputePspReference: pspReference,
				MerchantAccountCode: MerchantAccount,
			})
			res, httpRes, err := service.DefendDispute(context.Background(), req)

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.NotNil(t, res)
			assert.Equal(t, "Dispute not found.", res.DisputeServiceResult.GetErrorMessage())
		})

		t.Run("Download dispute defense document", func(t *testing.T) {
			pspReference := createTestPayment()
			req := service.DownloadDisputeDefenseDocumentInput().DownloadDefenseDocumentRequest(disputes.DownloadDefenseDocumentRequest{
				DefenseDocumentType: "DefenseMaterial",
				DisputePspReference: pspReference,
				MerchantAccountCode: MerchantAccount,
			})
			_, httpRes, err := service.DownloadDisputeDefenseDocument(context.Background(), req)

			require.NotNil(t, err)
			assert.Equal(t, 403, httpRes.StatusCode)
		})
	})
}
