package tests

import (
	"context"
	"encoding/json"
	"github.com/adyen/adyen-go-api-library/v6/src/adyen"
	"github.com/adyen/adyen-go-api-library/v6/src/balanceplatform"
	"github.com/adyen/adyen-go-api-library/v6/src/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func Test_BalancePlatform(t *testing.T) {
	client := adyen.NewClient(&common.Config{
		ApiKey:      "YOUR_ADYEN_API_KEY",
		Environment: "TEST",
		Debug:       false,
	})
	service := client.BalancePlatform

	mux := http.NewServeMux()
	// Success case
	mux.HandleFunc("/accountHolders/123", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "GET", r.Method)
		w.Header().Set("Content-Type", "application/json")
		file, _ := os.Open("fixtures/account_holder.json")
		io.Copy(w, file)
	})
	mux.HandleFunc("/accountHolders/123/balanceAccounts", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "GET", r.Method)
		assert.Equal(t, "5", r.URL.Query().Get("limit"))
		assert.Equal(t, "42", r.URL.Query().Get("offset"))
		w.Header().Set("Content-Type", "application/json")
		file, _ := os.Open("fixtures/paginated_balance_accounts_response.json")
		io.Copy(w, file)
	})
	mux.HandleFunc("/balanceAccounts/balanceAccountId/sweeps/sweepId", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "DELETE", r.Method)
		w.WriteHeader(http.StatusNoContent)
		// poof!
	})
	mux.HandleFunc("/transactionRules/transactionRuleId", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "DELETE", r.Method)
		w.Header().Set("Content-Type", "application/json")
		file, _ := os.Open("fixtures/transaction_rule.json")
		io.Copy(w, file)
	})
	// Error case
	mux.HandleFunc("/paymentInstruments/666", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "PATCH", r.Method)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		file, _ := os.Open("fixtures/not_found.json")
		io.Copy(w, file)
	})

	mockServer := httptest.NewServer(mux)
	defer mockServer.Close()
	// base path is shared between all endpoints
	client.BalancePlatform.AccountHoldersApi.BasePath = func() string {
		return mockServer.URL
	}

	t.Run("Configuration", func(t *testing.T) {
		testClient := adyen.NewClient(&common.Config{
			Environment: common.TestEnv,
		})
		assert.Equal(t, "https://balanceplatform-api-test.adyen.com/bcl/v2", testClient.BalancePlatform.AccountHoldersApi.BasePath())
		liveClient := adyen.NewClient(&common.Config{
			Environment: common.LiveEnv,
		})
		assert.Equal(t, "https://balanceplatform-api-live.adyen.com/bcl/v2", liveClient.BalancePlatform.BalanceAccountsApi.BasePath())
	})

	t.Run("Get an account holder", func(t *testing.T) {
		req := service.AccountHoldersApi.GetAccountHolderConfig(context.Background(), "123")
		res, httpRes, err := service.AccountHoldersApi.GetAccountHolder(req)

		require.NotNil(t, res)
		require.NotNil(t, httpRes)
		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, "YOUR_BALANCE_PLATFORM", res.GetBalancePlatform())
		accountHolderCapability := res.GetCapabilities()["receiveFromPlatformPayments"]
		require.True(t, accountHolderCapability.GetEnabled())
		assert.Equal(t, "pending", accountHolderCapability.GetVerificationStatus())
		assert.Equal(t, "active", *res.Status)
	})

	t.Run("Get all balance accounts of an account holder", func(t *testing.T) {
		req := service.AccountHoldersApi.GetAllBalanceAccountsOfAccountHolderConfig(context.Background(), "123")
		req = req.Offset(42).Limit(5)
		res, httpRes, err := service.AccountHoldersApi.GetAllBalanceAccountsOfAccountHolder(req)

		require.NotNil(t, res)
		require.NotNil(t, httpRes)
		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		require.True(t, res.HasNext)
		assert.Equal(t, "AH32272223222B5CTBMZT6W2V", res.GetBalanceAccounts()[0].GetAccountHolderId())
	})

	t.Run("Error response", func(t *testing.T) {
		req := service.PaymentInstrumentsApi.UpdatePaymentInstrumentConfig(context.Background(), "666")
		body := balanceplatform.NewPaymentInstrumentUpdateRequest()
		body.SetBalanceAccountId("BA32272223222B5CM82WL892M")
		req.PaymentInstrumentUpdateRequest(*body)
		_, _, err := service.PaymentInstrumentsApi.UpdatePaymentInstrument(req)

		apiError := err.(common.APIError)
		assert.Equal(t, float64(422), apiError.Status)
		assert.Equal(t, "30_112", apiError.Code)

		var restError balanceplatform.RestServiceError
		_ = json.Unmarshal(apiError.RawBody, &restError)
		assert.Equal(t, "Entity was not found", restError.Title)
		assert.Equal(t, "Payment instrument not found", restError.Detail)
	})

	t.Run("Delete a sweep", func(t *testing.T) {
		req := service.BalanceAccountsApi.DeleteSweepConfig(context.Background(), "balanceAccountId", "sweepId")
		httpRes, err := service.BalanceAccountsApi.DeleteSweep(req)

		assert.Equal(t, 204, httpRes.StatusCode)
		require.Nil(t, err)
	})

	t.Run("Delete a transaction rule", func(t *testing.T) {
		req := service.TransactionRulesApi.DeleteTransactionRuleConfig(context.Background(), "transactionRuleId")
		res, httpRes, err := service.TransactionRulesApi.DeleteTransactionRule(req)

		require.NotNil(t, res)
		require.NotNil(t, httpRes)
		require.Nil(t, err)
		assert.Equal(t, "myRule12345", res.GetReference())
	})
}