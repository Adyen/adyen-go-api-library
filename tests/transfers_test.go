package tests

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adyen/adyen-go-api-library/v7/src/adyen"
	"github.com/adyen/adyen-go-api-library/v7/src/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Transfers(t *testing.T) {
	client := adyen.NewClient(&common.Config{
		ApiKey:      "YOUR_ADYEN_API_KEY",
		Environment: "TEST",
	})

	mux := http.NewServeMux()

	// Success case
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

	// Error case
	mux.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
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
	client.Transfers.BasePath = func() string { return mockServer.URL }

	t.Run("make succesful transfer", func(t *testing.T) {
		service := client.Transfers

		request := service.TransferFundsConfig(context.Background())
		_, httpRes, err := service.TransferFunds(request)

		require.NoError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		require.Nil(t, err)
	})

	t.Run("make unsuccesful get transactions call", func(t *testing.T) {
		service := client.Transfers

		request := service.GetAllTransactionsConfig(context.Background())
		_, httpRes, err := service.GetAllTransactions(request)

		assert.Equal(t, 403, httpRes.StatusCode)
		require.NotNil(t, err)

	})
}
