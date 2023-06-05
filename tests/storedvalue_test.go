package tests

import (
	"github.com/adyen/adyen-go-api-library/v6/src/adyen"
	"github.com/adyen/adyen-go-api-library/v6/src/common"
	"github.com/adyen/adyen-go-api-library/v6/src/storedvalue"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_StoredValue(t *testing.T) {
	client := adyen.NewClient(&common.Config{
		ApiKey:      "YOUR_ADYEN_API_KEY",
		Environment: "TEST",
	})
	//client.GetConfig().Debug = true

	mux := http.NewServeMux()
	// Success case
	mux.HandleFunc("/checkBalance", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{
    "currentBalance": {
        "currency": "EUR",
        "value": 5600
    },
    "pspReference": "881564657480267D",
    "resultCode": "Success"
}`)
	})
	// Error case
	mux.HandleFunc("/voidTransaction", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status":500,"errorCode":"306","message":"No voidable transaction found","errorType":"internal"}`)
	})

	mockServer := httptest.NewServer(mux)
	defer mockServer.Close()
	client.StoredValue().BasePath = func() string { return mockServer.URL }

	t.Run("Configuration", func(t *testing.T) {
		testClient := adyen.NewClient(&common.Config{
			Environment: common.TestEnv,
		})
		assert.Equal(t, "https://pal-test.adyen.com/pal/servlet/StoredValue/v46", testClient.StoredValue().BasePath())
		liveClient := adyen.NewClient(&common.Config{
			Environment: common.LiveEnv,
		})
		assert.Equal(t, "https://pal-live.adyen.com/pal/servlet/StoredValue/v46", liveClient.StoredValue().BasePath())
	})

	t.Run("Check balance", func(t *testing.T) {
		request := storedvalue.NewStoredValueBalanceCheckRequest("YOUR_MERCHANT_ACCOUNT", map[string]string{
			"type":         "svs",
			"number":       "603628672882001915092",
			"securityCode": "5754",
		}, "YOUR_REFERENCE")
		request.SetStore("YOUR_STORE_ID")

		res, httpRes, err := client.StoredValue().CheckBalance(request)

		require.NotNil(t, res)
		require.NotNil(t, httpRes)
		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, "Success", res.GetResultCode())
		balance := res.GetCurrentBalance()
		assert.Equal(t, int64(5600), balance.GetValue())
	})

	t.Run("Error response", func(t *testing.T) {
		_, httpRes, err := client.StoredValue().VoidTransaction(&storedvalue.StoredValueVoidRequest{
			MerchantAccount:   "YOUR_MERCHANT_ACCOUNT",
			OriginalReference: "",
			Reference:         nil,
		})

		assert.Equal(t, 500, httpRes.StatusCode)
		apiError := err.(common.APIError)
		assert.Equal(t, "306", apiError.Code)
		assert.Equal(t, "No voidable transaction found", apiError.Message)
	})
}
