package management

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/adyen/adyen-go-api-library/v7/src/adyen"
	"github.com/adyen/adyen-go-api-library/v7/src/common"
	"github.com/adyen/adyen-go-api-library/v7/src/management"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func Test_ManagementAPI(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch strings.TrimSpace(r.Method) {
		case "GET":
			switch strings.TrimSpace(r.URL.Path) {
			case "/companies":
				companyId := "TestCompany123"
				model := management.ListCompanyResponse{Data: []management.Company{{Id: &companyId}}}
				mockResponse(http.StatusOK, w, model)
			case "/companies/TestCompany123":
				companyId := "TestCompany123"
				model := management.Company{Id: &companyId}
				mockResponse(http.StatusOK, w, model)
			case "/companies/notExisting":
				model := management.NewRestServiceErrorWithDefaults()
				mockResponse(http.StatusForbidden, w, model)
			case "/companies/TestCompany123/merchants":
				companyId := "TestCompany123"
				model := management.Company{Id: &companyId}
				mockResponse(http.StatusOK, w, model)
			case "/companies/notExisting/merchants":
				model := management.NewRestServiceErrorWithDefaults()
				mockResponse(http.StatusForbidden, w, model)
			default:
				t.Errorf("Mock not found")
				http.NotFoundHandler().ServeHTTP(w, r)
			}
		case "PATCH":
			switch strings.TrimSpace(r.URL.Path) {
			case "/companies/notExisting":
				model := management.NewRestServiceErrorWithDefaults()
				mockResponse(http.StatusForbidden, w, model)
			}
		}
	}))
	defer mockServer.Close()

	client := adyen.NewClient(&common.Config{
		ApiKey:      "YOUR_ADYEN_API_KEY",
		Environment: common.TestEnv,
		Debug:       "true" == os.Getenv("DEBUG"),
	})
	service := client.Management()
	// base path is shared between all endpoints
	service.APICredentialsCompanyLevelApi.BasePath = func() string {
		return mockServer.URL
	}

	t.Run("Configuration", func(t *testing.T) {
		testClient := adyen.NewClient(&common.Config{
			Environment: common.TestEnv,
		})
		assert.Equal(t, "https://management-test.adyen.com/v1", testClient.Management().MyAPICredentialApi.BasePath())
		liveClient := adyen.NewClient(&common.Config{
			Environment: common.LiveEnv,
		})
		assert.Equal(t, "https://management-live.adyen.com/v1", liveClient.Management().WebhooksCompanyLevelApi.BasePath())
	})

	t.Run("Test ListCompanyAccounts", func(t *testing.T) {
		t.Run("Create an API request that should pass", func(t *testing.T) {
			req := service.AccountCompanyLevelApi.ListCompanyAccountsConfig(context.Background())

			_, httpRes, _, err := service.AccountCompanyLevelApi.ListCompanyAccounts(req)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.Nil(t, err)
		})
	})

	t.Run("Test GetCompanyAccount", func(t *testing.T) {
		t.Run("Create an API request that should pass", func(t *testing.T) {
			companyId := "TestCompany123"
			req := service.AccountCompanyLevelApi.GetCompanyAccountConfig(context.Background(), companyId)
			resp, httpRes, serviceError, err := service.AccountCompanyLevelApi.GetCompanyAccount(req)
			if serviceError.ErrorCode != "500" {
				fmt.Fprintf(os.Stderr, "Error: %v\n", serviceError)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 200, httpRes.StatusCode)
			require.Nil(t, err)
		})

		t.Run("Create an API request that should fail", func(t *testing.T) {
			companyId := "notExisting"
			req := service.AccountCompanyLevelApi.GetCompanyAccountConfig(context.Background(), companyId)
			resp, httpRes, serviceError, _ := service.AccountCompanyLevelApi.GetCompanyAccount(req)
			if serviceError.ErrorCode != "500" {
				fmt.Fprintf(os.Stderr, "Error: %v\n", serviceError)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 403, httpRes.StatusCode)
		})
	})

	t.Run("Test ListMerchantAccounts", func(t *testing.T) {
		t.Run("Create an API request that should pass", func(t *testing.T) {
			companyId := "TestCompany123"
			req := service.AccountCompanyLevelApi.ListMerchantAccountsConfig(context.Background(), companyId)
			resp, httpRes, serviceError, _ := service.AccountCompanyLevelApi.ListMerchantAccounts(req)
			if serviceError.ErrorCode != "500" {
				fmt.Fprintf(os.Stderr, "Error: %v\n", serviceError)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 200, httpRes.StatusCode)
		})

		t.Run("Create an API request that should fail", func(t *testing.T) {
			companyId := "notExisting"
			req := service.AccountCompanyLevelApi.ListMerchantAccountsConfig(context.Background(), companyId)
			resp, httpRes, serviceError, err := service.AccountCompanyLevelApi.ListMerchantAccounts(req)
			if serviceError.ErrorCode != "500" {
				fmt.Fprintf(os.Stderr, "Error: %v\n", serviceError)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 403, httpRes.StatusCode)
			require.NotNil(t, err)
		})
	})
}

// mock Response given the model
func mockResponse(status int, w http.ResponseWriter, model interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(model)
}
