package tests

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/adyen/adyen-go-api-library/v8/src/adyen"
	"github.com/adyen/adyen-go-api-library/v8/src/common"
	"github.com/adyen/adyen-go-api-library/v8/src/disputes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Disputes(t *testing.T) {
	client := adyen.NewClient(&common.Config{
		ApiKey:      "YOUR_ADYEN_API_KEY",
		Environment: "TEST",
		Debug:       "true" == os.Getenv("DEBUG"),
	})

	mux := http.NewServeMux()

	// Success cases
	mux.HandleFunc("/retrieveApplicableDefenseReasons", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"disputeServiceResult":{"errorMessage":"Dispute not found.","success":false}}`)
	})

	mux.HandleFunc("/supplyDefenseDocument", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"disputeServiceResult":{"errorMessage":"Unknown dispute","success":false}}`)
	})

	mux.HandleFunc("/deleteDisputeDefenseDocument", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"disputeServiceResult":{"errorMessage":"Unknown dispute","success":false}}`)
	})

	mux.HandleFunc("/defendDispute", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"disputeServiceResult":{"errorMessage":"Dispute not found.","success":false}}`)
	})

	mockServer := httptest.NewServer(mux)
	defer mockServer.Close()

	service := client.Disputes()
	service.BasePath = func() string { return mockServer.URL }
	pspReference := "FOO12345BAR"
	merchantAccount := "YOUR_MERCHANT_ACCOUNT"

	t.Run("Configuration", func(t *testing.T) {
		testClient := adyen.NewClient(&common.Config{
			Environment: common.TestEnv,
		})
		assert.Equal(t, "https://ca-test.adyen.com/ca/services/DisputeService/v30", testClient.Disputes().BasePath())
		liveClient := adyen.NewClient(&common.Config{
			Environment: common.LiveEnv,
		})
		assert.Equal(t, "https://ca-live.adyen.com/ca/services/DisputeService/v30", liveClient.Disputes().BasePath())
	})

	t.Run("Retrieve applicable defense reasons", func(t *testing.T) {
		req := service.RetrieveApplicableDefenseReasonsInput().DefenseReasonsRequest(disputes.DefenseReasonsRequest{
			DisputePspReference: pspReference,
			MerchantAccountCode: merchantAccount,
		})
		res, httpRes, err := service.RetrieveApplicableDefenseReasons(context.Background(), req)

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.NotNil(t, res)
		assert.Equal(t, "Dispute not found.", res.DisputeServiceResult.GetErrorMessage())
	})

	t.Run("Supply defense document", func(t *testing.T) {
		req := service.SupplyDefenseDocumentInput().SupplyDefenseDocumentRequest(disputes.SupplyDefenseDocumentRequest{
			DefenseDocuments: []disputes.DefenseDocument{
				{
					Content:                 "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z8BQDwAEhQGAhKmMIQAAAABJRU5ErkJggg==",
					ContentType:             "image/jpg",
					DefenseDocumentTypeCode: "DefenseMaterial",
				},
			},
			DisputePspReference: pspReference,
			MerchantAccountCode: merchantAccount,
		})
		res, httpRes, err := service.SupplyDefenseDocument(context.Background(), req)

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.NotNil(t, res)
		assert.Equal(t, "Unknown dispute", res.DisputeServiceResult.GetErrorMessage())
	})

	t.Run("Delete dispute defense document", func(t *testing.T) {
		req := service.DeleteDisputeDefenseDocumentInput().DeleteDefenseDocumentRequest(disputes.DeleteDefenseDocumentRequest{
			DefenseDocumentType: "DefenseMaterial",
			DisputePspReference: pspReference,
			MerchantAccountCode: merchantAccount,
		})
		res, httpRes, err := service.DeleteDisputeDefenseDocument(context.Background(), req)

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.NotNil(t, res)
		assert.Equal(t, "Unknown dispute", res.DisputeServiceResult.GetErrorMessage())
	})

	t.Run("Defend dispute", func(t *testing.T) {
		req := service.DefendDisputeInput().DefendDisputeRequest(disputes.DefendDisputeRequest{
			DefenseReasonCode:   "DuplicateChargeback",
			DisputePspReference: pspReference,
			MerchantAccountCode: merchantAccount,
		})
		res, httpRes, err := service.DefendDispute(context.Background(), req)

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.NotNil(t, res)
		assert.Equal(t, "Dispute not found.", res.DisputeServiceResult.GetErrorMessage())
	})
}
