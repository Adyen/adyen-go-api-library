/*
 * Testing AllowedOriginsCompanyLevelApiService
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
	"os"
	"testing"
)

func Test_ManagementAPI_AllowedOriginsCompanyLevelApiService(t *testing.T) {
	godotenv.Load("./../../.env")

	var (
		APIKey = os.Getenv("ADYEN_API_KEY")
		env    = Management.TestEnv
	)

	configuration, err := Management.NewManagementAPIConfiguration(APIKey, env)
	require.Nil(t, err, "Error creating Config object")
	apiClient := Management.NewAPIClient(configuration)

	t.Run("Test AllowedOriginsCompanyLevelApiService GetCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOrigins", func(t *testing.T) {

		t.Run("Create an API request that should pass", func(t *testing.T) {

			t.Skip("skipping test: add new ApiCredential TODO")

			companyId := "TestCompany123"
			apiCredentialId := "na"

			resp, httpRes, err := apiClient.AllowedOriginsCompanyLevelApi.GetCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOrigins(context.Background(), companyId, apiCredentialId).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error : %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 200, httpRes.StatusCode)
			require.Nil(t, err)
		})
	})

	t.Run("Test AllowedOriginsCompanyLevelApiService GetCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOriginsOriginId", func(t *testing.T) {

		t.Run("Create an API request that should pass", func(t *testing.T) {

			t.Skip("skipping test: add new AllowOrigin TODO")

			companyId := "TestCompany123"
			apiCredentialId := "na"
			originId := "na"

			resp, httpRes, err := apiClient.AllowedOriginsCompanyLevelApi.GetCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOriginsOriginId(context.Background(), companyId, apiCredentialId, originId).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error : %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 200, httpRes.StatusCode)
			require.Nil(t, err)
		})
	})

}
