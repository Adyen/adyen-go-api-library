/*
 * Testing AccountCompanyLevelApiService
 *
 * Contact: support@adyen.com
 */

package management

import (
	"context"
	"fmt"
	Management "github.com/adyen/adyen-go-api-library/v6/src/management"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func Test_ManagementAPI_AccountCompanyLevelApiService(t *testing.T) {
	godotenv.Load("./../../.env")

	var (
		APIKey = "n/a"
		env    = Management.TestEnv
		server *httptest.Server
	)

	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// mock here
		switch strings.TrimSpace(r.Method) {
		case "GET":
			switch strings.TrimSpace(r.URL.Path) {
			case "/companies":
				companyId := "TestCompany123"
				model := Management.ListCompanyResponse{Data: []Management.Company{{Id: &companyId}}}
				mockResponse(http.StatusOK, w, model)
			case "/companies/TestCompany123":
				companyId := "TestCompany123"
				model := Management.Company{Id: &companyId}
				mockResponse(http.StatusOK, w, model)
			case "/companies/notExisting":
				model := Management.NewRestServiceErrorWithDefaults()
				mockResponse(http.StatusForbidden, w, model)
			case "/companies/TestCompany123/merchants":
				companyId := "TestCompany123"
				model := Management.Company{Id: &companyId}
				mockResponse(http.StatusOK, w, model)
			case "/companies/notExisting/merchants":
				model := Management.NewRestServiceErrorWithDefaults()
				mockResponse(http.StatusForbidden, w, model)
			default:
				t.Errorf("Mock not found")
				http.NotFoundHandler().ServeHTTP(w, r)
			}
		case "PATCH":
			switch strings.TrimSpace(r.URL.Path) {
			case "/companies/TestCompany123":
				legalEntityId := "updateCompanyLegalEntityId"
				model := Management.SetLegalEntityToAccountResponse{LegalEntityId: &legalEntityId}
				mockResponse(http.StatusOK, w, model)
			case "/companies/notExisting":
				model := Management.NewRestServiceErrorWithDefaults()
				mockResponse(http.StatusForbidden, w, model)
			}
		case "POST":
			switch strings.TrimSpace(r.URL.Path) {
			case "/companies":
				companyId := "newCompany"
				legalEntityId := "newCompanyLegalEntityId"
				model := Management.CreateCompanyResponse{CompanyId: &companyId, LegalEntityId: &legalEntityId}
				mockResponse(http.StatusOK, w, model)
			}
		}

	}))

	configuration, err := Management.NewManagementAPIConfiguration(APIKey, env)
	configuration.Servers = Management.ServerConfigurations{
		{
			URL:         server.URL,
			Description: "Mock Server",
		},
	}

	require.Nil(t, err, "Error creating Config object")
	apiClient := Management.NewAPIClient(configuration)

	t.Run("Test ListCompanyAccounts", func(t *testing.T) {

		t.Run("Create an API request that should pass", func(t *testing.T) {

			resp, httpRes, err := apiClient.AccountCompanyLevelApi.ListCompanyAccounts(context.Background()).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 200, httpRes.StatusCode)
			require.Nil(t, err)
		})
	})

	t.Run("Test GetCompanyAccount", func(t *testing.T) {

		t.Run("Create an API request that should pass", func(t *testing.T) {

			companyId := "TestCompany123"

			resp, httpRes, err := apiClient.AccountCompanyLevelApi.GetCompanyAccount(context.Background(), companyId).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 200, httpRes.StatusCode)
			require.Nil(t, err)
		})

		t.Run("Create an API request that should fail", func(t *testing.T) {

			companyId := "notExisting"
			resp, httpRes, err := apiClient.AccountCompanyLevelApi.GetCompanyAccount(context.Background(), companyId).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 403, httpRes.StatusCode)
			require.NotNil(t, err)
		})

	})

	t.Run("Test ListMerchantAccounts", func(t *testing.T) {

		t.Run("Create an API request that should pass", func(t *testing.T) {

			companyId := "TestCompany123"

			resp, httpRes, err := apiClient.AccountCompanyLevelApi.ListMerchantAccounts(context.Background(), companyId).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 200, httpRes.StatusCode)
			require.Nil(t, err)
		})

		t.Run("Create an API request that should fail", func(t *testing.T) {

			companyId := "notExisting"
			resp, httpRes, err := apiClient.AccountCompanyLevelApi.ListMerchantAccounts(context.Background(), companyId).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 403, httpRes.StatusCode)
			require.NotNil(t, err)
		})

	})

	t.Run("Test UpdateLegalEntityOfSpecificCompany", func(t *testing.T) {

		t.Run("Create an API request that should pass", func(t *testing.T) {

			companyId := "TestCompany123"

			resp, httpRes, err := apiClient.AccountCompanyLevelApi.UpdateLegalEntityOfSpecificCompany(context.Background(), companyId).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 200, httpRes.StatusCode)
			require.Nil(t, err)
		})

		t.Run("Create an API request that should fail", func(t *testing.T) {

			companyId := "notExisting"
			resp, httpRes, err := apiClient.AccountCompanyLevelApi.UpdateLegalEntityOfSpecificCompany(context.Background(), companyId).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 403, httpRes.StatusCode)
			require.NotNil(t, err)
		})
	})

	t.Run("Test CreateCompany", func(t *testing.T) {

		t.Run("Create an API request that should pass", func(t *testing.T) {

			//createCompanyRequest := *Management.NewCreateCompanyRequest()
			createCompanyRequest := Management.CreateCompanyRequest{Description: nil}

			resp, httpRes, err := apiClient.AccountCompanyLevelApi.CreateCompany(context.Background()).
				CreateCompanyRequest(createCompanyRequest).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 200, httpRes.StatusCode)
			require.Nil(t, err)
		})

	})

}
