package tests

import (
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
	})
	//client.GetConfig().Debug = true
	service := client.BalancePlatform

	mux := http.NewServeMux()
	// Success case
	mux.HandleFunc("/accountHolders/123", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "GET", r.Method)
		w.Header().Set("Content-Type", "application/json")
		file, _ := os.Open("fixtures/account_holder.json")
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

	t.Run("Get an account holder", func(t *testing.T) {
		res, httpRes, err := service.AccountHoldersApi.GetAccountHolder(common.PtrString("123"))

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

	t.Run("Error response", func(t *testing.T) {
		id := "666"
		req := balanceplatform.NewPaymentInstrumentUpdateRequest()
		_, _, err := service.PaymentInstrumentsApi.UpdatePaymentInstrument(&id, req)

		apiError := err.(common.APIError)
		assert.Equal(t, float64(422), apiError.Status)
		assert.Equal(t, "30_112", apiError.Code)

		var restError balanceplatform.RestServiceError
		_ = json.Unmarshal(apiError.RawBody, &restError)
		assert.Equal(t, "Entity was not found", restError.Title)
		assert.Equal(t, "Payment instrument not found", restError.Detail)
	})
}
