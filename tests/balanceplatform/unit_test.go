package balanceplatform

import (
	"context"
	"errors"
	"github.com/adyen/adyen-go-api-library/v8/src/adyen"
	"github.com/adyen/adyen-go-api-library/v8/src/balanceplatform"
	"github.com/adyen/adyen-go-api-library/v8/src/common"
	"github.com/adyen/adyen-go-api-library/v8/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestBalancePlatform(t *testing.T) {
	client := adyen.NewClient(&common.Config{
		ApiKey:      "YOUR_ADYEN_API_KEY",
		Environment: "TEST",
		Debug:       false,
	})
	service := client.BalancePlatform()

	mux := http.NewServeMux()
	mockResponse := tests.MockResponse(t, mux)

	// Success cases
	mockResponse(http.StatusOK, "GET PATCH", "/accountHolders/123", "account_holder.json")
	mux.HandleFunc("/accountHolders/123/balanceAccounts", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "GET", r.Method)
		assert.Equal(t, "5", r.URL.Query().Get("limit"))
		assert.Equal(t, "42", r.URL.Query().Get("offset"))
		w.Header().Set("Content-Type", "application/json")
		file, _ := os.Open("../fixtures/paginated_balance_accounts_response.json")
		io.Copy(w, file)
	})
	mux.HandleFunc("/balanceAccounts/balanceAccountId/sweeps/sweepId", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "DELETE", r.Method)
		w.WriteHeader(http.StatusNoContent)
		// poof!
	})
	mux.HandleFunc("/validateBankAccountIdentification", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		// no response
	})
	mockResponse(http.StatusOK, "GET PATCH", "/balanceAccounts/BA123/sweeps/SWPC123", "sweep.json")
	mockResponse(http.StatusOK, "DELETE", "/transactionRules/transactionRuleId", "transaction_rule.json")

	// Error case
	mockResponse(http.StatusUnprocessableEntity, "PATCH", "/paymentInstruments/666", "not_found.json")

	mockServer := httptest.NewServer(mux)
	defer mockServer.Close()
	// base path is shared between all endpoints
	client.BalancePlatform().AccountHoldersApi.BasePath = func() string {
		return mockServer.URL
	}

	// Network error
	mux.HandleFunc("/balanceAccounts", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		mockServer.CloseClientConnections()
	})

	t.Run("Configuration", func(t *testing.T) {
		testClient := adyen.NewClient(&common.Config{
			Environment: common.TestEnv,
		})
		assert.Equal(t, "https://balanceplatform-api-test.adyen.com/bcl/v2", testClient.BalancePlatform().AccountHoldersApi.BasePath())
		liveClient := adyen.NewClient(&common.Config{
			Environment: common.LiveEnv,
		})
		assert.Equal(t, "https://balanceplatform-api-live.adyen.com/bcl/v2", liveClient.BalancePlatform().BalanceAccountsApi.BasePath())
	})

	t.Run("Get an account holder", func(t *testing.T) {
		req := service.AccountHoldersApi.GetAccountHolderInput("123")
		res, httpRes, err := service.AccountHoldersApi.GetAccountHolder(context.Background(), req)

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

	t.Run("Update account holder", func(t *testing.T) {
		req := service.AccountHoldersApi.UpdateAccountHolderInput("123")

		req = req.AccountHolderUpdateRequest(balanceplatform.AccountHolderUpdateRequest{
			Status: common.PtrString("closed"),
		})

		res, httpRes, err := service.AccountHoldersApi.UpdateAccountHolder(context.Background(), req)

		assert.Equal(t, 200, httpRes.StatusCode)
		require.NoError(t, err)
		assert.Equal(t, "YOUR_BALANCE_PLATFORM", res.GetBalancePlatform())
	})

	t.Run("Get all balance accounts of an account holder", func(t *testing.T) {
		req := service.AccountHoldersApi.GetAllBalanceAccountsOfAccountHolderInput("123")
		req = req.Offset(42).Limit(5)
		res, httpRes, err := service.AccountHoldersApi.GetAllBalanceAccountsOfAccountHolder(context.Background(), req)

		require.NotNil(t, res)
		require.NotNil(t, httpRes)
		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		require.True(t, res.HasNext)
		assert.Equal(t, "AH32272223222B5CTBMZT6W2V", res.GetBalanceAccounts()[0].GetAccountHolderId())
	})

	t.Run("Error response", func(t *testing.T) {
		req := service.PaymentInstrumentsApi.UpdatePaymentInstrumentInput("666")
		body := balanceplatform.NewPaymentInstrumentUpdateRequest()
		body.SetBalanceAccountId("BA32272223222B5CM82WL892M")
		req.PaymentInstrumentUpdateRequest(*body)

		_, _, err := service.PaymentInstrumentsApi.UpdatePaymentInstrument(context.Background(), req)

		var serviceError common.RestServiceError
		errors.As(err, &serviceError)
		assert.Equal(t, int32(422), serviceError.Status)
		assert.Equal(t, "30_112", serviceError.GetErrorCode())
		assert.Equal(t, "Entity was not found", serviceError.GetTitle())
		assert.Equal(t, "Payment instrument not found", serviceError.GetDetail())
	})

	t.Run("Gateway Timeout error", func(t *testing.T) {
		req := service.BalanceAccountsApi.CreateBalanceAccountInput().BalanceAccountInfo(balanceplatform.BalanceAccountInfo{
			AccountHolderId: "AH123ABC",
			Description:     common.PtrString("S.Hopper - Main balance account"),
		})

		_, httpRes, err := service.BalanceAccountsApi.CreateBalanceAccount(context.Background(), req)

		assert.Error(t, err)
		assert.Nil(t, httpRes)
	})

	t.Run("Get a sweep", func(t *testing.T) {
		req := service.BalanceAccountsApi.GetSweepInput("BA123", "SWPC123")

		res, httpRes, err := service.BalanceAccountsApi.GetSweep(context.Background(), req)

		assert.Equal(t, 200, httpRes.StatusCode)
		assert.NoError(t, err)
		assert.Equal(t, "active", res.GetStatus())
	})

	t.Run("Update a sweep", func(t *testing.T) {
		req := service.BalanceAccountsApi.UpdateSweepInput("BA123", "SWPC123")

		cron := balanceplatform.NewSweepSchedule("cron")
		cron.SetCronExpression("6 6 6 6 6")
		req = req.UpdateSweepConfigurationV2(balanceplatform.UpdateSweepConfigurationV2{
			Schedule:    cron,
			Description: common.PtrString("Let's Go!"),
		})

		res, httpRes, err := service.BalanceAccountsApi.UpdateSweep(context.Background(), req)

		assert.Equal(t, 200, httpRes.StatusCode)
		assert.NoError(t, err)
		assert.Equal(t, "cron", res.Schedule.Type)
		assert.Equal(t, "6 6 6 6 6", *res.Schedule.CronExpression)
	})

	t.Run("Delete a sweep", func(t *testing.T) {
		req := service.BalanceAccountsApi.DeleteSweepInput("balanceAccountId", "sweepId")
		httpRes, err := service.BalanceAccountsApi.DeleteSweep(context.Background(), req)

		assert.Equal(t, 204, httpRes.StatusCode)
		require.Nil(t, err)
	})

	t.Run("Delete a transaction rule", func(t *testing.T) {
		req := service.TransactionRulesApi.DeleteTransactionRuleInput("transactionRuleId")
		res, httpRes, err := service.TransactionRulesApi.DeleteTransactionRule(context.Background(), req)

		require.NotNil(t, res)
		require.NotNil(t, httpRes)
		require.Nil(t, err)
		assert.Equal(t, "myRule12345", res.GetReference())
	})

	t.Run("Validate a bank account", func(t *testing.T) {
		req := service.BankAccountValidationApi.ValidateBankAccountIdentificationInput()
		httpRes, err := service.BankAccountValidationApi.ValidateBankAccountIdentification(context.Background(), req)

		assert.Equal(t, 200, httpRes.StatusCode)
		assert.NoError(t, err)
	})
}
