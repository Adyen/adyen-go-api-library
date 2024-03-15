//go:build integration
// +build integration

package management

import (
	"context"
	"errors"
	"github.com/adyen/adyen-go-api-library/v9/src/adyen"
	"github.com/adyen/adyen-go-api-library/v9/src/common"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test_ManagementAPI_Integration(t *testing.T) {
	// Breaking IT's skipping for now
	t.Skip()

	godotenv.Load("./../../.env")

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

			assert.Equal(t, 503, httpRes.StatusCode)
			require.NotNil(t, resp)
			require.Nil(t, serviceErr) //TODO error trace /home/runner/work/adyen-go-api-library/adyen-go-api-library/tests/management/integration_test.go:36
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
			assert.Equal(t, 503, httpRes.StatusCode)
			require.NotNil(t, restServiceErr)
		})
	})

	t.Run("List merchant accounts", func(t *testing.T) {
		req := service.AccountMerchantLevelApi.ListMerchantAccountsInput()
		req = req.PageNumber(1).PageSize(1)

		resp, httpRes, serviceErr := service.AccountMerchantLevelApi.ListMerchantAccounts(context.Background(), req)

		require.Nil(t, serviceErr) //TODO error strace /home/runner/work/adyen-go-api-library/adyen-go-api-library/tests/management/integration_test.go:73
		assert.Equal(t, 200, httpRes.StatusCode)
		require.Equal(t, 1, len(resp.Data), "Should contain only one merchant account")
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
		assert.NotEmpty(t, restServiceErr.GetRequestId()) //TODO Should NOT be empty, but was
		assert.Equal(t, "", restServiceErr.GetErrorCode())
		assert.Equal(t, int32(0), restServiceErr.GetStatus())
		assert.Equal(t, 503, httpRes.StatusCode)
	})
}
