/*
Legal Entity Management API

Testing LegalEntitiesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package legalentity

import (
	"context"
	"encoding/json"
	openapiclient "github.com/adyen/adyen-go-api-library/v6/src/legalentity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_LegalEntity_LegalEntitiesApiService(t *testing.T) {

	var (
		username    = "n/a"
		password    = "n/a"
		env         = openapiclient.TestEnv
		mockServer  *httptest.Server
		mockRequest *http.Request
	)

	mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		mockRequest = r

		switch strings.TrimSpace(r.Method) {
		case "GET":
			switch strings.TrimSpace(r.URL.Path) {
			case "/legalEntities/123":
				legalEntityId := "123"
				model := openapiclient.LegalEntity{Id: legalEntityId}
				mockResponse(http.StatusOK, w, model)
			case "/legalEntities/123/businessLines":
				model := openapiclient.NewBusinessLinesWithDefaults()
				mockResponse(http.StatusOK, w, model)
			}
		case "PATCH":
			switch strings.TrimSpace(r.URL.Path) {
			case "/legalEntities/123":
				legalEntityId := "123"
				model := openapiclient.LegalEntity{Id: legalEntityId}
				mockResponse(http.StatusOK, w, model)
			}
		case "POST":
			switch strings.TrimSpace(r.URL.Path) {
			case "/legalEntities":
				legalEntityId := "123"
				model := openapiclient.LegalEntity{Id: legalEntityId}
				mockResponse(http.StatusOK, w, model)
			}
		}

	}))
	defer mockServer.Close()

	configuration, err := openapiclient.NewConfiguration(username, password, env)
	configuration.Servers = openapiclient.ServerConfigurations{
		{
			URL:         mockServer.URL,
			Description: "Mock Server",
		},
	}

	require.Nil(t, err, "Error creating Config object")
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LegalEntitiesApiService GetLegalEntity", func(t *testing.T) {

		apiClient.LegalEntitiesApi.GetLegalEntity(context.Background(), "123").Execute()

		assert.Equal(t, "GET", mockRequest.Method)
		assert.Equal(t, "/legalEntities/123", mockRequest.URL.Path)
		assert.NotEmpty(t, mockRequest.Header.Get("Authorization"))

	})

	t.Run("Test LegalEntitiesApiService GetAllBusinessLinesUnderLegalEntity", func(t *testing.T) {

		apiClient.LegalEntitiesApi.GetAllBusinessLinesUnderLegalEntity(context.Background(), "123").Execute()

		assert.Equal(t, "GET", mockRequest.Method)
		assert.Equal(t, "/legalEntities/123/businessLines", mockRequest.URL.Path)
		assert.NotEmpty(t, mockRequest.Header.Get("Authorization"))

	})

	t.Run("Test LegalEntitiesApiService UpdateLegalEntity", func(t *testing.T) {

		legalEntityInfo := *openapiclient.NewLegalEntityInfoWithDefaults()

		apiClient.LegalEntitiesApi.UpdateLegalEntity(context.Background(), "123").LegalEntityInfo(legalEntityInfo).Execute()

		assert.Equal(t, "PATCH", mockRequest.Method)
		assert.Equal(t, "/legalEntities/123", mockRequest.URL.Path)
		assert.NotEmpty(t, mockRequest.Header.Get("Authorization"))

	})

	t.Run("Test LegalEntitiesApiService CreateLegalEntity", func(t *testing.T) {

		apiClient.LegalEntitiesApi.CreateLegalEntity(context.Background()).Execute()

		assert.Equal(t, "POST", mockRequest.Method)
		assert.Equal(t, "/legalEntities", mockRequest.URL.Path)
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