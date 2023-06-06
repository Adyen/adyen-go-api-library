package management

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/adyen/adyen-go-api-library/v7/src/adyen"
	"github.com/adyen/adyen-go-api-library/v7/src/common"
	"github.com/adyen/adyen-go-api-library/v7/src/management"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	service := client.Management()

	t.Run("Test MyAPICredentialApiService GetMe", func(t *testing.T) {
		t.Run("Create an API request that should pass", func(t *testing.T) {
			req := service.MyAPICredentialApi.GetApiCredentialDetailsConfig(context.Background())

			resp, httpRes, serviceErr, err := service.MyAPICredentialApi.GetApiCredentialDetails(req)

			require.NotNil(t, serviceErr)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, resp)
			require.Nil(t, err)
		})

		t.Run("Create an API request that should fail", func(t *testing.T) {
			// create misconfigured client to test invalid apiKey
			invalidKeyClient := adyen.NewClient(&common.Config{
				ApiKey:      "xxx",
				Environment: common.TestEnv,
			})
			req := invalidKeyClient.Management().MyAPICredentialApi.GetApiCredentialDetailsConfig(context.Background())

			_, httpRes, serviceErr, _ := invalidKeyClient.Management().MyAPICredentialApi.GetApiCredentialDetails(req)

			assert.Equal(t, 401, httpRes.StatusCode)
			require.NotNil(t, serviceErr)
		})
	})

	t.Run("List terminals", func(t *testing.T) {
		req := service.TerminalsTerminalLevelApi.ListTerminalsConfig(context.Background())
		req = req.Countries("NL").PageSize(1)

		resp, httpRes, serviceErr, err := service.TerminalsTerminalLevelApi.ListTerminals(req)

		require.NotNil(t, serviceErr)
		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		require.NotNil(t, resp)
	})

	t.Run("Create an API request that should fail for company", func(t *testing.T) {
		// Creates a test that should fail because of the wrong Id
		req := service.AccountCompanyLevelApi.GetCompanyAccountConfig(context.Background(), "99999")
		_, httpRes, serviceErr, err := service.AccountCompanyLevelApi.GetCompanyAccount(req)
		assert.NotEmpty(t, serviceErr.GetRequestId())
		assert.Equal(t, "010", serviceErr.GetErrorCode())
		assert.Equal(t, int32(403), serviceErr.GetStatus())
		assert.Equal(t, 403, httpRes.StatusCode)
		assert.Equal(t, reflect.TypeOf(serviceErr), reflect.TypeOf(management.RestServiceError{}))
		assert.NotNil(t, err)
	})
}
