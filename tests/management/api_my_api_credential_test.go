/*
 * Adyen API Client
 *
 * Contact: support@adyen.com
 */

package management

import (
	"context"
	"fmt"
	"os"
	"testing"
	Management "github.com/adyen/adyen-go-api-library/v6/src/management"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ManagementAPI_MyAPICredentialApiService(t *testing.T) {
	godotenv.Load("./../../.env")

	var (
		APIKey          = os.Getenv("ADYEN_API_KEY")
		env 			= Management.TestEnv
	)

	configuration, err := Management.NewManagementAPIConfiguration(APIKey, env)
	require.Nil(t, err, "Error creating Config object")
	apiClient := Management.NewAPIClient(configuration)

	t.Run("Test MyAPICredentialApiService GetMe", func(t *testing.T) {

		t.Run("Create an API request that should pass", func(t *testing.T) {

			resp, httpRes, err := apiClient.MyAPICredentialApi.GetMe(context.Background()).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error when calling `MyAPICredentialApi.GetMe``: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, resp)
		})
		t.Run("Create an API request that should fail", func(t *testing.T) {

			// create misconfiged client to test invalid apiKey
			misconfig, err := Management.NewManagementAPIConfiguration("xxx", env)
			require.Nil(t, err, "Error creating Config object")
			apiInvalidKeyClient := Management.NewAPIClient(misconfig)

			resp, httpRes, err := apiInvalidKeyClient.MyAPICredentialApi.GetMe(context.Background()).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error when calling `MyAPICredentialApi.GetMe``: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			assert.Equal(t, 401, httpRes.StatusCode)
			require.NotNil(t, err)
		})		
	})
	
	t.Run("Test MyAPICredentialApiService GetMeAllowedOrigins", func(t *testing.T) {

		t.Run("Create an API request that should pass", func(t *testing.T) {

			resp, httpRes, err := apiClient.MyAPICredentialApi.GetMeAllowedOrigins(context.Background()).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error when calling `MyAPICredentialApi.GetMeAllowedOrigins``: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, resp)
		})
	})

	t.Run("Test MyAPICredentialApiService PostMeAllowedOrigins", func(t *testing.T) {

		t.Run("Create an API request that should pass", func(t *testing.T) {

			apiPostMeAllowedOriginsRequest := apiClient.MyAPICredentialApi.PostMeAllowedOrigins(context.Background())
			apiPostMeAllowedOriginsRequest.CreateAllowedOriginRequest().Domain = "a"
			resp, httpRes, err := apiPostMeAllowedOriginsRequest.Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error when calling `PostMeAllowedOrigins.GetMeAllowedOrigins``: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
			}

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, resp)
		})
	})
}
