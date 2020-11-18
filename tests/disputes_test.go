package tests

import (
	"os"
	"testing"

	"github.com/adyen/adyen-go-api-library/v4/src/adyen"
	"github.com/adyen/adyen-go-api-library/v4/src/checkout"
	"github.com/adyen/adyen-go-api-library/v4/src/common"
	"github.com/adyen/adyen-go-api-library/v4/src/disputes"
	"github.com/adyen/adyen-go-api-library/v4/src/payments"
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
	})

	createTestPayment := func() string {
		res, _, _ := client.Checkout.Payments(&checkout.PaymentRequest{
			Amount: checkout.Amount{
				Currency: "EUR",
				Value:    1000,
			},
			Reference: "DISPUTES_CHARGEBACK",
			PaymentMethod: map[string]interface{}{
				"type":        "scheme",
				"number":      "4111111111111111",
				"expiryMonth": "03",
				"expiryYear":  "2030",
				"holderName":  "chargeback:10.4",
				"cvc":         "737",
			},
			ReturnUrl:       "https://adyen.com",
			MerchantAccount: MerchantAccount,
		})
		captureRes, _, _ := client.Payments.Capture(&payments.ModificationRequest{
			OriginalReference: res.PspReference,
			ModificationAmount: &payments.Amount{
				Currency: "EUR",
				Value:    1000,
			},
			Reference:       "MODIFICATION_REFERENCE",
			MerchantAccount: MerchantAccount,
		})
		return captureRes.PspReference
	}

	t.Run("Disputes", func(t *testing.T) {
		t.Run("Retrieve applicable defense reasons", func(t *testing.T) {
			pspReference := createTestPayment()
			res, httpRes, err := client.Disputes.RetrieveApplicableDefenseReasons(&disputes.DefenseReasonsRequest{
				DisputePspReference: pspReference,
				MerchantAccountCode: MerchantAccount,
			})

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.NotNil(t, res)
			assert.Equal(t, "Dispute not found.", res.DisputeServiceResult.ErrorMessage)
		})

		t.Run("Supply defense document", func(t *testing.T) {
			pspReference := createTestPayment()
			res, httpRes, err := client.Disputes.SupplyDefenseDocument(&disputes.SupplyDefenseDocumentRequest{
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

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.NotNil(t, res)
			assert.Equal(t, "Unknown dispute", res.DisputeServiceResult.ErrorMessage)
		})

		t.Run("Delete dispute defense document", func(t *testing.T) {
			pspReference := createTestPayment()
			res, httpRes, err := client.Disputes.DeleteDisputeDefenseDocument(&disputes.DeleteDefenseDocumentRequest{
				DefenseDocumentType: "DefenseMaterial",
				DisputePspReference: pspReference,
				MerchantAccountCode: MerchantAccount,
			})

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.NotNil(t, res)
			assert.Equal(t, "Unknown dispute", res.DisputeServiceResult.ErrorMessage)
		})

		t.Run("Defend dispute", func(t *testing.T) {
			pspReference := createTestPayment()
			res, httpRes, err := client.Disputes.DefendDispute(&disputes.DefendDisputeRequest{
				DefenseReasonCode:   "DuplicateChargeback",
				DisputePspReference: pspReference,
				MerchantAccountCode: MerchantAccount,
			})

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.NotNil(t, res)
			assert.Equal(t, "Dispute not found.", res.DisputeServiceResult.ErrorMessage)
		})

		t.Run("Download dispute defense document", func(t *testing.T) {
			pspReference := createTestPayment()
			_, httpRes, err := client.Disputes.DownloadDisputeDefenseDocument(&disputes.DownloadDefenseDocumentRequest{
				DefenseDocumentType: "DefenseMaterial",
				DisputePspReference: pspReference,
				MerchantAccountCode: MerchantAccount,
			})

			require.NotNil(t, err)
			assert.Equal(t, 403, httpRes.StatusCode)
		})
	})
}
