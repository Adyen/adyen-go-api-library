/*
 * Testing APICredentialsMerchantLevelApiService
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

func Test_ManagementAPI_APICredentialsMerchantLevelApiService(t *testing.T) {
	godotenv.Load("./../../.env")

	var (
		APIKey = "n/a"
		env    = Management.TestEnv
		server *httptest.Server
	)

	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// mock here
		switch strings.TrimSpace(r.URL.Path) {
		case "/merchants/TestMerchantAccount/apiCredentials":
			model := Management.ListMerchantApiCredentialsResponse{Data: []Management.ApiCredential{{Id: "TestMerchantAccount"}}}
			mockResponse(http.StatusOK, w, model)
		case "/merchants/notExisting/apiCredentials":
			model := Management.ListMerchantApiCredentialsResponse{Data: []Management.ApiCredential{{Id: "TestMerchantAccount"}}}
			mockResponse(http.StatusForbidden, w, model)
		default:
			t.Errorf("Mock not found")
			http.NotFoundHandler().ServeHTTP(w, r)
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

	t.Run("Test APICredentialsMerchantLevelApi GetMerchantsIdApiCredentials", func(t *testing.T) {

		t.Run("Create an API request that should pass", func(t *testing.T) {

			id := "TestMerchantAccount"
			resp, httpRes, err := apiClient.APICredentialsMerchantLevelApi.GetMerchantsIdApiCredentials(context.Background(), id).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error when calling `APICredentialsMerchantLevelApi.GetMerchantsIdApiCredentials``: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 200, httpRes.StatusCode)
			require.Nil(t, err)
		})
		t.Run("Create an API request that should fail", func(t *testing.T) {

			id := "notExisting"
			resp, httpRes, err := apiClient.APICredentialsMerchantLevelApi.GetMerchantsIdApiCredentials(context.Background(), id).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error when calling `APICredentialsMerchantLevelApi.GetMerchantsIdApiCredentials``: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 403, httpRes.StatusCode)
			require.NotNil(t, err)
			assert.Equal(t, "403 Forbidden", err.Error())
		})
	})

}
