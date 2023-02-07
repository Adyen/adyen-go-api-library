package legalentity

import (
	openapiclient "github.com/adyen/adyen-go-api-library/v6/src/legalentity"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_LegalEntity_Configuration(t *testing.T) {
	godotenv.Load("./../../.env")

	var (
		username = "usr"
		password = "pwd"
		env      = openapiclient.TestEnv
	)

	t.Run("Test Configuration", func(t *testing.T) {
		t.Run("Create a Client that should pass", func(t *testing.T) {

			configuration, err := openapiclient.NewConfiguration(username, password, env)
			require.Nil(t, err, "Error creating Config object")

			apiClient := openapiclient.NewAPIClient(configuration)

			// credentials
			require.NotNil(t, apiClient.GetConfig().Username)
			require.NotNil(t, apiClient.GetConfig().Password)
			// env
			require.NotNil(t, apiClient.GetConfig().Environment)
			assert.Equal(t, openapiclient.TestEnv, apiClient.GetConfig().Environment)
			// URL
			assert.NotNil(t, apiClient.GetConfig().Servers)
			assert.Equal(t, "https://kyc-test.adyen.com/lem/v2", apiClient.GetConfig().Servers[0].URL)

		})
		t.Run("Create a Client that should fail (missing apiKey)", func(t *testing.T) {

			_, err := openapiclient.NewConfiguration("", password, env)
			require.NotNil(t, err, "Error creating Config object")
			assert.Equal(t, "username is not provided", err.Error())

		})
		t.Run("Create a Client that should fail (invalid Environment)", func(t *testing.T) {

			_, err := openapiclient.NewConfiguration(username, "", env)
			require.NotNil(t, err, "Error creating Config object")
			assert.Equal(t, "password is not provided", err.Error())

		})
	})

}
