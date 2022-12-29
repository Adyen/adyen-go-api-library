//go:build integration
// +build integration

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
	"os"
	"testing"
)

func Test_Integration_ManagementAPI_AccountCompanyLevelApiService(t *testing.T) {
	godotenv.Load("./../../../.env")

	var (
		APIKey = os.Getenv("ADYEN_API_KEY")
		env    = Management.TestEnv
	)

	configuration, err := Management.NewManagementAPIConfiguration(APIKey, env)
	require.Nil(t, err, "Error creating Config object")
	apiClient := Management.NewAPIClient(configuration)

	t.Run("Test AccountCompanyLevelApiService GetCompanies", func(t *testing.T) {

		t.Run("Create an API request that should pass", func(t *testing.T) {

			resp, httpRes, err := apiClient.AccountCompanyLevelApi.ListCompanyAccounts(context.Background()).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error : %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 200, httpRes.StatusCode)
			require.Nil(t, err)
		})
		t.Run("Create an API request that should fail", func(t *testing.T) {

			var invalidPageNumber int32 = 0

			resp, httpRes, err := apiClient.AccountCompanyLevelApi.ListCompanyAccounts(context.Background()).PageNumber(invalidPageNumber).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error : %v\n", err)
				fmt.Fprintf(os.Stderr, "Full response: %v\n", resp)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
			}

			require.NotNil(t, err)
			assert.Equal(t, 422, httpRes.StatusCode)
		})
	})

	t.Run("Test AccountCompanyLevelApiService GetCompaniesCompanyId", func(t *testing.T) {

		t.Run("Create an API request that should pass", func(t *testing.T) {

			companyId := "TestCompany123"

			resp, httpRes, err := apiClient.AccountCompanyLevelApi.GetCompanyAccount(context.Background(), companyId).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error : %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 200, httpRes.StatusCode)
			require.Nil(t, err)
		})
	})

	t.Run("Test AccountCompanyLevelApiService GetCompaniesCompanyIdMerchants", func(t *testing.T) {

		t.Run("Create an API request that should pass", func(t *testing.T) {

			companyId := "TestCompany123"

			resp, httpRes, err := apiClient.AccountCompanyLevelApi.ListMerchantAccounts(context.Background(), companyId).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error : %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 200, httpRes.StatusCode)
			require.Nil(t, err)
		})
	})
}
