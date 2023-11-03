package tests

import (
	"context"
	"errors"
	"github.com/adyen/adyen-go-api-library/v8/src/transfers"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/adyen/adyen-go-api-library/v8/src/adyen"
	"github.com/adyen/adyen-go-api-library/v8/src/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Transfers(t *testing.T) {
	client := adyen.NewClient(&common.Config{
		ApiKey:      "YOUR_ADYEN_API_KEY",
		Environment: "TEST",
	})

	mux := http.NewServeMux()

	// Success cases
	mux.HandleFunc("/transfers", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{
			"id": "1W1UG35U8A9J5ZLG",
			"accountHolder": {
			  "description": "Your account holder description",
			  "id": "AH3227C223222C5GXQXF658WB",
			  "reference": "Your account holder reference"
			},
			"amount": {
			  "currency": "EUR",
			  "value": 10000
			},
			"balanceAccount": {
			  "description": "Your balance account description",
			  "id": "BAB8B2C3D4E5F6G7H8D9J6GD4",
			  "reference": "Your balance account reference"
			},
			"balanceAccountId": "BAB8B2C3D4E5F6G7H8D9J6GD4",
			"category": "internal",
			"counterparty": {
			  "balanceAccountId": "BA32272223222B5LPRFDW7J9G"
			},
			"referenceForBeneficiary": "Your-reference-sent-to-the-beneficiary",
			"reference": "Your internal reference for the transfer",
			"description": "Your description for the transfer",
			"direction": "outgoing",
			"reason": "approved",
			"status": "authorised"
		  }`)
	})

	mux.HandleFunc("/grants", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{
			"id": "GR00000000000000000000001",
			"grantAccountId": "CG00000000000000000000001",
			"grantOfferId": "0000000000000001",
			"counterparty": {
				"accountHolderId": "AH00000000000000000000001",
				"balanceAccountId": "BA00000000000000000000001"
			},
			"amount": {
				"currency": "EUR",
				"value": 1000000
			},
			"fee": {
				"amount": {
					"value": 120000,
					"currency": "EUR"
				}
			},
			"balances": {
				"currency": "EUR",
				"fee": 120000,
				"principal": 1000000,
				"total": 1120009
			},
			"repayment": {
				"basisPoints": 1400
			},
			"status": "Pending"
		}`)
	})

	mockResponse := MockResponse(t, mux)
	mockResponse(http.StatusOK, "GET", "/transactions", "all_transactions.json")

	// Error case
	mux.HandleFunc("/transactions/ERRForbidden403", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "GET", r.Method)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		io.WriteString(w, `{
			"type": "https://docs.adyen.com/errors/forbidden",
			"title": "Forbidden",
			"status": 403,
			"detail": "Not allowed",
			"errorCode": "00_403"
		}`)
	})

	mockServer := httptest.NewServer(mux)
	defer mockServer.Close()
	client.Transfers().TransfersApi.BasePath = func() string { return mockServer.URL }
	service := client.Transfers()

	t.Run("Configuration", func(t *testing.T) {
		testClient := adyen.NewClient(&common.Config{
			Environment: common.TestEnv,
		})
		assert.Equal(t, "https://balanceplatform-api-test.adyen.com/btl/v4", testClient.Transfers().CapitalApi.BasePath())
		liveClient := adyen.NewClient(&common.Config{
			Environment: common.LiveEnv,
		})
		assert.Equal(t, "https://balanceplatform-api-live.adyen.com/btl/v4", liveClient.Transfers().TransfersApi.BasePath())
	})

	t.Run("make successful transfer", func(t *testing.T) {
		request := service.TransfersApi.TransferFundsInput()

		_, httpRes, err := service.TransfersApi.TransferFunds(context.Background(), request)

		require.NoError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		require.Nil(t, err)
	})

	t.Run("Get a transaction", func(t *testing.T) {
		req := service.TransactionsApi.GetTransactionInput("ERRForbidden403")

		_, httpRes, err := service.TransactionsApi.GetTransaction(context.Background(), req)

		assert.Equal(t, 403, httpRes.StatusCode)
		require.NotNil(t, err)
		var serviceError common.RestServiceError
		errors.As(err, &serviceError)
		assert.Equal(t, int32(403), serviceError.Status)
		assert.Equal(t, "00_403", serviceError.GetErrorCode())
		assert.Equal(t, "Forbidden", serviceError.GetTitle())
		assert.Equal(t, "Not allowed", serviceError.GetDetail())
	})

	t.Run("Get all transactions", func(t *testing.T) {
		req := service.TransactionsApi.GetAllTransactionsInput()
		since := time.Date(2022, 1, 1, 1, 2, 3, 0, time.UTC)
		until := time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC)
		req = req.BalancePlatform("Your balance platform").CreatedSince(since).CreatedUntil(until)

		res, httpRes, err := service.TransactionsApi.GetAllTransactions(context.Background(), req)

		assert.Equal(t, 200, httpRes.StatusCode)
		require.NoError(t, err)
		require.Len(t, res.Data, 4)
		assert.Equal(t, 2023, res.GetData()[0].GetCreationDate().Year())
	})

	t.Run("Request a grant payout", func(t *testing.T) {
		request := service.CapitalApi.RequestGrantPayoutInput().CapitalGrantInfo(transfers.CapitalGrantInfo{
			GrantAccountId: "CG00000000000000000000001",
			GrantOfferId:   "0000000000000001",
			Counterparty: &transfers.Counterparty{
				AccountHolderId:  common.PtrString("AH00000000000000000000001"),
				BalanceAccountId: common.PtrString("BA00000000000000000000001"),
			},
		})

		grant, httpRes, err := service.CapitalApi.RequestGrantPayout(context.Background(), request)

		require.NoError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, "GR00000000000000000000001", grant.GetId())
		assert.Equal(t, "Pending", grant.GetStatus())
		balances := grant.GetBalances()
		assert.Equal(t, int64(1120009), balances.GetTotal())
	})
}
