/*
 * Adyen API Client
 *
 * Contact: support@adyen.com
 */

package management

import (
	"os"
	"testing"
	Management "github.com/adyen/adyen-go-api-library/v6/src/management"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ManagementAPI_Configuration(t *testing.T) {
	godotenv.Load("./../../.env")

	var (
		APIKey          = os.Getenv("ADYEN_API_KEY")
		env 			= Management.TestEnv
	)

	t.Run("Test Configuration", func(t *testing.T) {
		t.Run("Create a Client that should pass", func(t *testing.T) {

			configuration, err := Management.NewManagementAPIConfiguration(APIKey, env)
			require.Nil(t, err, "Error creating Config object")

			apiClient := Management.NewAPIClient(configuration)
		
			// API key
			require.NotNil(t, apiClient.GetConfig().ApiKey)
			// env
			require.NotNil(t, apiClient.GetConfig().Environment)
			assert.Equal(t, Management.TestEnv, apiClient.GetConfig().Environment)
			// URL
			assert.NotNil(t, apiClient.GetConfig().Servers)
			assert.Equal(t, "https://management-test.adyen.com/v1", apiClient.GetConfig().Servers[0].URL)

		})
		t.Run("Create a Client that should fail (missing apiKey)", func(t *testing.T) {

			_, err := Management.NewManagementAPIConfiguration("", env)
			require.NotNil(t, err, "Error creating Config object")
			assert.Equal(t, "apiKey is not provided", err.Error())

		})
		t.Run("Create a Client that should fail (invalid Environment)", func(t *testing.T) {

			_, err := Management.NewManagementAPIConfiguration(APIKey, "not_existing")
			require.NotNil(t, err, "Error creating Config object")
			assert.Equal(t, "unsupported environment not_existing", err.Error())

		})
	})
	
}
