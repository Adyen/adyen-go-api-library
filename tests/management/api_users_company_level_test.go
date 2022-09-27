/*
 * Testing UsersCompanyLevelApiService
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

func Test_ManagementAPI_UsersCompanyLevelApiService(t *testing.T) {
	godotenv.Load("./../../.env")

	var (
		APIKey = os.Getenv("ADYEN_API_KEY")
		env    = Management.TestEnv
	)

	configuration, err := Management.NewManagementAPIConfiguration(APIKey, env)
	require.Nil(t, err, "Error creating Config object")
	apiClient := Management.NewAPIClient(configuration)

	t.Run("Test UsersCompanyLevelApiService GetCompaniesCompanyIdUsers", func(t *testing.T) {

		t.Run("Create an API request that should pass", func(t *testing.T) {

			companyId := "TestCompany123"

			resp, httpRes, err := apiClient.UsersCompanyLevelApi.GetCompaniesCompanyIdUsers(context.Background(), companyId).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error : %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, resp)
		})
	})

	t.Run("Test UsersCompanyLevelApiService GetCompaniesCompanyIdUsers", func(t *testing.T) {

		t.Run("Create an API request that should pass", func(t *testing.T) {

			t.Skip("skipping test: add new User TODO")

			companyId := "TestCompany123"
			userId := "S2-3E50415B3237"

			resp, httpRes, err := apiClient.UsersCompanyLevelApi.GetCompaniesCompanyIdUsersUserId(context.Background(), companyId, userId).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error : %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, resp)
		})
	})

}
