package main

import (
	"os"
	"testing"

	"github.com/adyen/adyen-go-api-library/v3/src/adyen"
	"github.com/adyen/adyen-go-api-library/v3/src/common"
	"github.com/adyen/adyen-go-api-library/v3/src/disputes"
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

	t.Run("Disputes", func(t *testing.T) {
		t.Run("Retrieve applicable defense reasons", func(t *testing.T) {
			res, httpRes, err := client.Disputes.RetrieveApplicableDefenseReasons(&disputes.DefenseReasonsRequest{
				DisputePspReference: string,
				MerchantAccountCode: string,
			})

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
		})
		t.Run("Supply defense document", func(t *testing.T) {
			res, httpRes, err := client.Disputes.SupplyDefenseDocument(&disputes.SupplyDefenseDocumentRequest{
				DefenseDocuments: []disputes.DefenseDocument{
					{
						Content:                 string,
						ContentType:             string,
						DefenseDocumentTypeCode: string,
					},
				},
				DisputePspReference: string,
				MerchantAccountCode: string,
			})

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
		})
		t.Run("Delete dispute defense document", func(t *testing.T) {
			res, httpRes, err := client.Disputes.DeleteDisputeDefenseDocument(&disputes.DeleteDefenseDocumentRequest{
				DefenseDocumentType: string,
				DisputePspReference: string,
				MerchantAccountCode: string,
			})

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
		})

		t.Run("Defend dispute", func(t *testing.T) {
			res, httpRes, err := client.Disputes.DefendDispute(&disputes.DefendDisputeRequest{
				DefenseReasonCode:   string,
				DisputePspReference: string,
				MerchantAccountCode: string,
			})

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
		})

		t.Run("Defend dispute", func(t *testing.T) {
			res, httpRes, err := client.Disputes.DownloadDisputeDefenseDocument(&disputes.DownloadDefenseDocumentRequest{
				DefenseDocumentType: string,
				DisputePspReference: string,
				MerchantAccountCode: string,
			})

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
		})
	})
}
