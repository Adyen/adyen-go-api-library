/*
 * Testing AccountMerchantLevelApiService
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

func Test_ManagementAPI_AccountMerchantLevelApiService(t *testing.T) {
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
			case "/merchants":
				merchantId := "TestMerchantAccount"
				model := Management.ListMerchantResponse{Data: []Management.Merchant{{Id: &merchantId}}}
				mockResponse(http.StatusOK, w, model)
			case "/merchants/TestMerchantAccount":
				merchantId := "TestMerchantAccount"
				model := Management.Merchant{Id: &merchantId}
				mockResponse(http.StatusOK, w, model)
			case "/merchants/notExisting":
				model := Management.NewRestServiceErrorWithDefaults()
				mockResponse(http.StatusForbidden, w, model)
			default:
				t.Errorf("Mock not found")
				http.NotFoundHandler().ServeHTTP(w, r)
			}
		case "PATCH":
			switch strings.TrimSpace(r.URL.Path) {
			case "/merchants/TestMerchantAccount/account":
				legalEntityId := "updateMerchantLegalEntityId"
				model := Management.SetLegalEntityToAccountResponse{LegalEntityId: &legalEntityId}
				mockResponse(http.StatusOK, w, model)
			case "/merchants/notExisting/account":
				model := Management.NewRestServiceErrorWithDefaults()
				mockResponse(http.StatusForbidden, w, model)
			case "/merchants/TestMerchantAccount":
				merchantId := "TestMerchantAccount"
				model := Management.Merchant{Id: &merchantId}
				mockResponse(http.StatusOK, w, model)
			case "/merchants/notExisting":
				model := Management.NewRestServiceErrorWithDefaults()
				mockResponse(http.StatusForbidden, w, model)
			}
		case "POST":
			switch strings.TrimSpace(r.URL.Path) {
			case "/merchants":
				merchantId := "newMerchantAccount"
				legalEntityId := "newCompanyLegalEntityId"
				model := Management.CreateMerchantResponse{Id: &merchantId, LegalEntityId: &legalEntityId}
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

	t.Run("Test ListMerchantAccounts", func(t *testing.T) {

		t.Run("Create an API request that should pass", func(t *testing.T) {

			resp, httpRes, err := apiClient.AccountMerchantLevelApi.ListMerchantAccounts(context.Background()).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 200, httpRes.StatusCode)
			require.Nil(t, err)
		})
	})

	t.Run("Test GetMerchantAccount", func(t *testing.T) {

		t.Run("Create an API request that should pass", func(t *testing.T) {

			merchantId := "TestMerchantAccount"
			resp, httpRes, err := apiClient.AccountMerchantLevelApi.GetMerchantAccount(context.Background(), merchantId).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 200, httpRes.StatusCode)
			require.Nil(t, err)
		})

		t.Run("Create an API request that should fail", func(t *testing.T) {

			merchantId := "notExisting"
			resp, httpRes, err := apiClient.AccountMerchantLevelApi.GetMerchantAccount(context.Background(), merchantId).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 403, httpRes.StatusCode)
			require.NotNil(t, err)
		})

	})

	t.Run("Test UpdateLegalEntityOfSpecificMerchant", func(t *testing.T) {

		t.Run("Create an API request that should pass", func(t *testing.T) {

			merchantId := "TestMerchantAccount"
			resp, httpRes, err := apiClient.AccountMerchantLevelApi.UpdateLegalEntityOfSpecificMerchant(context.Background(), merchantId).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 200, httpRes.StatusCode)
			require.Nil(t, err)
		})

		t.Run("Create an API request that should fail", func(t *testing.T) {

			merchantId := "notExisting"
			resp, httpRes, err := apiClient.AccountMerchantLevelApi.UpdateLegalEntityOfSpecificMerchant(context.Background(), merchantId).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 403, httpRes.StatusCode)
			require.NotNil(t, err)
		})
	})

	t.Run("Test UpdateSpecificMerchant", func(t *testing.T) {

		t.Run("Create an API request that should pass", func(t *testing.T) {

			merchantId := "TestMerchantAccount"
			resp, httpRes, err := apiClient.AccountMerchantLevelApi.UpdateSpecificMerchant(context.Background(), merchantId).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 200, httpRes.StatusCode)
			require.Nil(t, err)
		})

		t.Run("Create an API request that should fail", func(t *testing.T) {

			merchantId := "notExisting"
			resp, httpRes, err := apiClient.AccountMerchantLevelApi.UpdateSpecificMerchant(context.Background(), merchantId).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 403, httpRes.StatusCode)
			require.NotNil(t, err)
		})
	})

	t.Run("Test CreateMerchantAccount", func(t *testing.T) {

		t.Run("Create an API request that should pass", func(t *testing.T) {

			createMerchantAccountRequest := *Management.NewCreateMerchantRequestWithDefaults()

			resp, httpRes, err := apiClient.AccountMerchantLevelApi.CreateMerchantAccount(context.Background()).
				CreateMerchantRequest(createMerchantAccountRequest).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 200, httpRes.StatusCode)
			require.Nil(t, err)
		})

	})

}
