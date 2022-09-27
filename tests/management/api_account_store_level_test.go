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

func Test_ManagementAPI_AccountStoreLevelApiService(t *testing.T) {
	godotenv.Load("./../../.env")

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

			resp, httpRes, err := apiClient.AccountStoreLevelApi.GetMerchantsMerchantIdStores(context.Background(), merchantId).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error : %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 200, httpRes.StatusCode)
			require.Nil(t, err)
		})
	})

	t.Run("Test AccountStoreLevelApiService GetMerchantsMerchantIdStoresStoreId", func(t *testing.T) {

		t.Run("Create an API request that should pass", func(t *testing.T) {

			t.Skip("skipping test: add new Store TODO")

			merchantId := "TestMerchantAccount"
			storeId := "ST3226522322285GPN5DP32DM"

			resp, httpRes, err := apiClient.AccountStoreLevelApi.GetMerchantsMerchantIdStoresStoreId(context.Background(), merchantId, storeId).Execute()
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

			t.Skip("skipping test: add new Store TODO")

			resp, httpRes, err := apiClient.AccountStoreLevelApi.GetStores(context.Background()).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error : %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 200, httpRes.StatusCode)
			require.Nil(t, err)
			assert.True(t, len(resp.Data) > 0)
		})
	})

	t.Run("Test AccountStoreLevelApiService GetStoresStoreId", func(t *testing.T) {

		t.Run("Create an API request that should pass", func(t *testing.T) {

			t.Skip("skipping test: add new Store TODO")

			storeId := "ST3226522322285GPN5DP32DM"

			resp, httpRes, err := apiClient.AccountStoreLevelApi.GetStoresStoreId(context.Background(), storeId).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error : %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 200, httpRes.StatusCode)
			require.Nil(t, err)
		})
	})

}
