package tests

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adyen/adyen-go-api-library/v6/src/adyen"
	"github.com/adyen/adyen-go-api-library/v6/src/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Pos_Terminal_Management(t *testing.T) {
	client := adyen.NewClient(&common.Config{
		ApiKey:      "YOUR_ADYEN_API_KEY",
		Environment: "TEST",
	})

	mux := http.NewServeMux()

	// Success case
	mux.HandleFunc("/assignTerminals", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{
			"results": {
			  "P400Plus-275479597": "RemoveConfigScheduled"
			}
		}`)
	})

	// Error case
	mux.HandleFunc("/findTerminal", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
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
	client.PosTerminalManagement.BasePath = func() string { return mockServer.URL }

	t.Run("successfully assign terminal", func(t *testing.T) {
		service := client.PosTerminalManagement

		request := service.AssignTerminalsConfig(context.Background())
		_, httpRes, err := service.AssignTerminals(request)

		require.NoError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		require.Nil(t, err)
	})

	t.Run("make unsuccesful findTerminal call", func(t *testing.T) {
		service := client.PosTerminalManagement

		request := service.FindTerminalConfig(context.Background())
		_, httpRes, err := service.FindTerminal(request)

		assert.Equal(t, 403, httpRes.StatusCode)
		require.NotNil(t, err)

	})
}