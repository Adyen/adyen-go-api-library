package tests

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adyen/adyen-go-api-library/v21/src/adyen"
	"github.com/adyen/adyen-go-api-library/v21/src/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_BalancePlatform(t *testing.T) {
	client := adyen.NewClient(&common.Config{
		ApiKey:      "YOUR_ADYEN_API_KEY",
		Environment: "TEST",
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/accountHolders/1001", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "GET", r.Method)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{
              "balancePlatform": "YOUR_BALANCE_PLATFORM",
              "description": "Account holder used for international payments and payouts",
              "legalEntityId": "LE322JV223222D5GG42KN6869",
              "reference": "S.Eller-001",
              "capabilities": {
                "receiveFromPlatformPayments": {
                  "enabled": true,
                  "requested": true,
                  "allowed": false,
                  "verificationStatus": "pending"
                },
                "receiveFromBalanceAccount": {
                  "enabled": true,
                  "requested": true,
                  "allowed": false,
                  "verificationStatus": "pending"
                },
                "sendToBalanceAccount": {
                  "enabled": true,
                  "requested": true,
                  "allowed": false,
                  "verificationStatus": "pending"
                },
                "sendToTransferInstrument": {
                  "enabled": true,
                  "requested": true,
                  "allowed": false,
                  "transferInstruments": [
                    {
                      "enabled": true,
                      "requested": true,
                      "allowed": false,
                      "id": "SE322KH223222F5GXZFNM3BGP",
                      "verificationStatus": "pending"
                    }
                  ],
                  "verificationStatus": "pending"
                }
              },
              "id": "AH3227C223222C5GXQXF658WB",
              "status": "active"
        }`)
	})
	// with additional attributes
	mux.HandleFunc("/accountHolders/1002", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "GET", r.Method)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{
              "balancePlatform": "YOUR_BALANCE_PLATFORM",
              "description": "Account holder used for international payments and payouts",
              "legalEntityId": "LE322JV223222D5GG42KN6869",
              "reference": "S.Eller-001",
              "additionalAttribute": "something",
              "capabilities": {
                "receiveFromPlatformPayments": {
                  "enabled": true,
                  "requested": true,
                  "allowed": false,
                  "verificationStatus": "pending"
                },
                "receiveFromBalanceAccount": {
                  "enabled": true,
                  "requested": true,
                  "allowed": false,
                  "verificationStatus": "pending"
                }
              },
              "id": "AH3227C223222C5GXQXF658WB",
              "status": "active"
        }`)
	})
	// with unknown enums
	mux.HandleFunc("/accountHolders/1003", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "GET", r.Method)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{
              "balancePlatform": "YOUR_BALANCE_PLATFORM",
              "description": "Account holder used for international payments and payouts",
              "legalEntityId": "LE322JV223222D5GG42KN6869",
              "reference": "S.Eller-001",
              "capabilities": {
                "receiveFromPlatformPayments": {
                  "enabled": true,
                  "requested": true,
                  "allowed": false,
                  "verificationStatus": "this is unknown"
                },
                "receiveFromBalanceAccount": {
                  "enabled": true,
                  "requested": true,
                  "allowed": false,
                  "verificationStatus": "pending"
                }
              },
              "id": "AH3227C223222C5GXQXF658WB",
              "status": "active"
        }`)
	})

	mockServer := httptest.NewServer(mux)
	defer mockServer.Close()
	client.BalancePlatform().AccountHoldersApi.BasePath = func() string { return mockServer.URL }
	service := client.BalancePlatform()

	t.Run("Configuration", func(t *testing.T) {
		testClient := adyen.NewClient(&common.Config{
			Environment: common.TestEnv,
		})
		assert.Equal(t, "https://balanceplatform-api-test.adyen.com/bcl/v2", testClient.BalancePlatform().AccountHoldersApi.BasePath())
		liveClient := adyen.NewClient(&common.Config{
			Environment: common.LiveEnv,
		})
		assert.Equal(t, "https://balanceplatform-api-live.adyen.com/bcl/v2", liveClient.BalancePlatform().AccountHoldersApi.BasePath())
	})

	t.Run("Get an account holder", func(t *testing.T) {
		request := service.AccountHoldersApi.GetAccountHolderInput("1001")
		response, httpRes, err := service.AccountHoldersApi.GetAccountHolder(context.Background(), request)

		require.NoError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, "AH3227C223222C5GXQXF658WB", response.GetId())
	})

	t.Run("Get an account holder with additional attributes", func(t *testing.T) {
		request := service.AccountHoldersApi.GetAccountHolderInput("1002")
		response, httpRes, err := service.AccountHoldersApi.GetAccountHolder(context.Background(), request)

		require.NoError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, "AH3227C223222C5GXQXF658WB", response.GetId())
	})

	t.Run("Get an account holder with unknown enum", func(t *testing.T) {
		request := service.AccountHoldersApi.GetAccountHolderInput("1003")
		response, httpRes, err := service.AccountHoldersApi.GetAccountHolder(context.Background(), request)

		require.NoError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, "AH3227C223222C5GXQXF658WB", response.GetId())
		capability := response.GetCapabilities()["receiveFromPlatformPayments"]
		assert.Equal(t, "this is unknown", capability.GetVerificationStatus())
	})
}
