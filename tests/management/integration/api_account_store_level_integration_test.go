//go:build integration
// +build integration

/*
 * Testing AccountStoreLevelApiService
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

func Test_Integration_ManagementAPI_AccountStoreLevelApiService(t *testing.T) {
	godotenv.Load("./../../../.env")

	var (
		APIKey = os.Getenv("ADYEN_API_KEY")
		env    = Management.TestEnv
	)

	configuration, err := Management.NewManagementAPIConfiguration(APIKey, env)
	require.Nil(t, err, "Error creating Config object")
	apiClient := Management.NewAPIClient(configuration)

	t.Run("Test AccountStoreLevelApiService GetMerchantsMerchantIdStores", func(t *testing.T) {

		t.Run("Create an API request that should pass", func(t *testing.T) {

			merchantId := "TestMerchantAccount"

			resp, httpRes, err := apiClient.AccountStoreLevelApi.ListStoresByMerchantId(context.Background(), merchantId).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error : %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 200, httpRes.StatusCode)
			require.Nil(t, err)
		})
	})

	t.Run("Test AccountStoreLevelApiService GetStores", func(t *testing.T) {

		t.Run("Create an API request that should pass", func(t *testing.T) {

			resp, httpRes, err := apiClient.AccountStoreLevelApi.ListStores(context.Background()).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error : %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 200, httpRes.StatusCode)
			require.Nil(t, err)
		})
	})

	t.Run("Test AccountStoreLevelApiService GetStoresStoreId", func(t *testing.T) {

		t.Run("Create an API request that should fail", func(t *testing.T) {

			storeId := "notExisting"

			resp, httpRes, err := apiClient.AccountStoreLevelApi.GetStoreById(context.Background(), storeId).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error : %v\n", err.Error())
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 422, httpRes.StatusCode)
			require.NotNil(t, err)
		})
	})

}
