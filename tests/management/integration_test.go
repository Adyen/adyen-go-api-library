package management

import (
	"context"
	"encoding/json"
	"github.com/adyen/adyen-go-api-library/v6/src/adyen"
	"github.com/adyen/adyen-go-api-library/v6/src/common"
	"github.com/adyen/adyen-go-api-library/v6/src/management"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test_ManagementAPI_Integration(t *testing.T) {
	godotenv.Load("./../../.env")

	if os.Getenv("ADYEN_API_KEY") == "" {
		t.Skip("Integration tests require credentials")
	}

	client := adyen.NewClient(&common.Config{
		ApiKey:      os.Getenv("ADYEN_API_KEY"),
		Environment: common.TestEnv,
		Debug:       "true" == os.Getenv("DEBUG"),
	})
	service := client.Management

	t.Run("Test MyAPICredentialApiService GetMe", func(t *testing.T) {
		t.Run("Create an API request that should pass", func(t *testing.T) {
			req := service.MyAPICredentialApi.GetApiCredentialDetailsConfig(context.Background())

			resp, httpRes, err := service.MyAPICredentialApi.GetApiCredentialDetails(req)

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, resp)
		})

		t.Run("Create an API request that should fail", func(t *testing.T) {
			// create misconfigured client to test invalid apiKey
			invalidKeyClient := adyen.NewClient(&common.Config{
				ApiKey:      "xxx",
				Environment: common.TestEnv,
			})
			req := invalidKeyClient.Management.MyAPICredentialApi.GetApiCredentialDetailsConfig(context.Background())

			_, httpRes, err := invalidKeyClient.Management.MyAPICredentialApi.GetApiCredentialDetails(req)

			assert.Equal(t, 401, httpRes.StatusCode)
			require.NotNil(t, err)
		})
	})

	t.Run("List terminals", func(t *testing.T) {
		req := service.TerminalsTerminalLevelApi.ListTerminalsConfig(context.Background())
		req = req.Countries("NL").PageSize(1)

		resp, httpRes, err := service.TerminalsTerminalLevelApi.ListTerminals(req)

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		require.NotNil(t, resp)
	})

	t.Run("Create an API request that should fail for company", func(t *testing.T) {
		// create misconfigured client to test invalid apiKey
		req := service.AccountCompanyLevelApi.GetCompanyAccountConfig(context.Background(), "99999")

		_, httpRes, err := service.AccountCompanyLevelApi.GetCompanyAccount(req)

		apiError := err.(common.APIError)
		var restError management.RestServiceError
		_ = json.Unmarshal(apiError.RawBody, &restError)
		assert.Equal(t, 403, httpRes.StatusCode)
		require.NotNil(t, string(apiError.RawBody))
	})
}
