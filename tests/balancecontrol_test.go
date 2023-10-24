package tests

import (
	"context"
	"github.com/adyen/adyen-go-api-library/v8/src/balancecontrol"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adyen/adyen-go-api-library/v8/src/adyen"
	"github.com/adyen/adyen-go-api-library/v8/src/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_BalanceControl(t *testing.T) {
	client := adyen.NewClient(&common.Config{
		ApiKey:      "YOUR_ADYEN_API_KEY",
		Environment: "TEST",
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/balanceTransfer", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{
          "amount": {
            "value": 50000,
            "currency": "EUR"
          },
          "createdAt": "2022-01-24T14:59:11+01:00",
          "description": "Your description for the transfer",
          "fromMerchant": "MerchantAccount_NL",
          "toMerchant": "MerchantAccount_DE",
          "type": "debit",
          "reference": "Unique reference for the transfer",
          "pspReference": "8816080397613514",
          "status": "transferred"
        }`)
	})

	mockServer := httptest.NewServer(mux)
	defer mockServer.Close()
	client.BalanceControl().BasePath = func() string { return mockServer.URL }
	service := client.BalanceControl()

	t.Run("Configuration", func(t *testing.T) {
		testClient := adyen.NewClient(&common.Config{
			Environment: common.TestEnv,
		})
		assert.Equal(t, "https://pal-test.adyen.com/pal/servlet/BalanceControl/v1", testClient.BalanceControl().BasePath())
		liveClient := adyen.NewClient(&common.Config{
			Environment: common.LiveEnv,
		})
		assert.Equal(t, "https://pal-live.adyen.com/pal/servlet/BalanceControl/v1", liveClient.BalanceControl().BasePath())
	})

	t.Run("Start a balance transfer", func(t *testing.T) {
		request := service.BalanceTransferInput().BalanceTransferRequest(balancecontrol.BalanceTransferRequest{
			Amount: balancecontrol.Amount{
				Currency: "EUR",
				Value:    50000,
			},
			Description:  common.PtrString("Testing"),
			FromMerchant: "MerchantAccount_NL",
			ToMerchant:   "MerchantAccount_DE",
			Type:         "debit",
		})

		balanceTransfer, httpRes, err := service.BalanceTransfer(context.Background(), request)

		require.NoError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, "8816080397613514", balanceTransfer.GetPspReference())
	})
}
