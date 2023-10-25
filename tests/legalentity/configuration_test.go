package legalentity

import (
	"context"
	"encoding/json"
	"github.com/adyen/adyen-go-api-library/v8/src/adyen"
	"github.com/adyen/adyen-go-api-library/v8/src/common"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_LegalEntity_Configuration(t *testing.T) {
	t.Run("Configuration", func(t *testing.T) {
		testClient := adyen.NewClient(&common.Config{
			Environment: common.TestEnv,
		})
		assert.Equal(t, "https://kyc-test.adyen.com/lem/v3", testClient.LegalEntity().BusinessLinesApi.BasePath())
		liveClient := adyen.NewClient(&common.Config{
			Environment: common.LiveEnv,
		})
		assert.Equal(t, "https://kyc-live.adyen.com/lem/v3", liveClient.LegalEntity().DocumentsApi.BasePath())
	})

	basicAuthClient := adyen.NewClient(&common.Config{
		Username:    "lem",
		Password:    "lem",
		Environment: "TEST",
	})
	service := basicAuthClient.LegalEntity()
	var (
		mockServer  *httptest.Server
		mockRequest *http.Request
	)
	mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mockRequest = r
	}))
	basicAuthClient.LegalEntity().TermsOfServiceApi.BasePath = func() string {
		return mockServer.URL
	}
	defer mockServer.Close()

	t.Run("Basic auth", func(t *testing.T) {
		req := service.BusinessLinesApi.UpdateBusinessLineInput("123")

		service.BusinessLinesApi.UpdateBusinessLine(context.Background(), req)

		assert.NotEmpty(t, mockRequest.Header.Get("Authorization"))
	})
}

// mock Response given the model
func mockResponse(status int, w http.ResponseWriter, model interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if model != nil {
		json.NewEncoder(w).Encode(model)
	}
}
