package management

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/adyen/adyen-go-api-library/v7/src/adyen"
	"github.com/adyen/adyen-go-api-library/v7/src/common"
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
			req := service.MyAPICredentialApi.GetApiCredentialDetailsInput()

			resp, httpRes, serviceErr := service.MyAPICredentialApi.GetApiCredentialDetails(context.Background(), req)

			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, resp)
			require.Nil(t, serviceErr)
		})

		t.Run("Create an API request that should fail", func(t *testing.T) {
			// create misconfigured client to test invalid apiKey
			invalidKeyClient := adyen.NewClient(&common.Config{
				ApiKey:      "xxx",
				Environment: common.TestEnv,
			})
			req := invalidKeyClient.Management().MyAPICredentialApi.GetApiCredentialDetailsInput()

			_, httpRes, serviceErr := invalidKeyClient.Management().MyAPICredentialApi.GetApiCredentialDetails(context.Background(), req)

			var restServiceErr common.RestServiceError
			errors.As(serviceErr, &restServiceErr)
			assert.Equal(t, 401, httpRes.StatusCode)
			require.NotNil(t, restServiceErr)
		})
	})

	t.Run("List terminals", func(t *testing.T) {
		req := service.TerminalsTerminalLevelApi.ListTerminalsInput()
		req = req.Countries("NL").PageSize(1)

		resp, httpRes, serviceErr := service.TerminalsTerminalLevelApi.ListTerminals(context.Background(), req)

		require.Nil(t, serviceErr)
		assert.Equal(t, 200, httpRes.StatusCode)
		require.NotNil(t, resp)
	})

	t.Run("Create an API request that should fail for company", func(t *testing.T) {
		// Creates a test that should fail because of the wrong Id
		req := service.AccountCompanyLevelApi.GetCompanyAccountInput("99999")

		_, httpRes, serviceErr := service.AccountCompanyLevelApi.GetCompanyAccount(context.Background(), req)

		var restServiceErr common.RestServiceError
		errors.As(serviceErr, &restServiceErr)
		assert.NotEmpty(t, restServiceErr.GetRequestId())
		assert.Equal(t, "010", restServiceErr.GetErrorCode())
		assert.Equal(t, int32(403), restServiceErr.GetStatus())
		assert.Equal(t, 403, httpRes.StatusCode)
	})
}
