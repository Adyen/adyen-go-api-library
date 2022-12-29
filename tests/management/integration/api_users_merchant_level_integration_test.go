//go:build integration
// +build integration

/*
 * Testing UsersMerchantLevelApiService
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

func Test_Integration_ManagementAPI_UsersMerchantLevelApiService(t *testing.T) {
	godotenv.Load("./../../../.env")

	var (
		APIKey = os.Getenv("ADYEN_API_KEY")
		env    = Management.TestEnv
	)

	configuration, err := Management.NewManagementAPIConfiguration(APIKey, env)
	require.Nil(t, err, "Error creating Config object")
	apiClient := Management.NewAPIClient(configuration)

	t.Run("Test UsersMerchantLevelApiService GetMerchantsIdUsers", func(t *testing.T) {

		t.Run("Create an API request that should pass", func(t *testing.T) {

			merchantId := "TestMerchantAccount"

			resp, httpRes, err := apiClient.UsersMerchantLevelApi.ListUsers(context.Background(), merchantId).Execute()
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
