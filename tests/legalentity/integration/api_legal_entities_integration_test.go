//go:build integration
// +build integration

package legalentity

import (
	"context"
	openapiclient "github.com/adyen/adyen-go-api-library/v6/src/legalentity"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test_Integration_LegalEntity_LegalEntitiesApiService(t *testing.T) {

	godotenv.Load("./../../../.env")

	var (
		username = os.Getenv("ADYEN_LEM_USERNAME")
		password = os.Getenv("ADYEN_LEM_PASSWORD")
		env      = openapiclient.TestEnv
	)

	configuration, err := openapiclient.NewConfiguration(username, password, env)
	require.Nil(t, err, "Error creating Config object")
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LegalEntitiesApiService GetLegalEntity", func(t *testing.T) {
		t.Run("Create an API request that should pass", func(t *testing.T) {

			var id = os.Getenv("LEM_LEGAL_ENTITY_ID")

			resp, httpRes, err := apiClient.LegalEntitiesApi.GetLegalEntity(context.Background(), id).Execute()

			require.Nil(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, httpRes.StatusCode)
		})
		t.Run("Create an API request that should fail", func(t *testing.T) {

			var id = "xxx"

			resp, httpRes, err := apiClient.LegalEntitiesApi.GetLegalEntity(context.Background(), id).Execute()

			require.NotNil(t, err)
			assert.Equal(t, 422, httpRes.StatusCode)
			require.NotNil(t, httpRes)
			require.Nil(t, resp)
		})
		t.Run("Create an API request that should not authorize", func(t *testing.T) {

			misconfig, err := openapiclient.NewConfiguration(username, "xxx", env)
			require.Nil(t, err, "Error creating Config object")
			invalidPasswordClient := openapiclient.NewAPIClient(misconfig)

			var id = "xxx"

			_, httpRes, err := invalidPasswordClient.LegalEntitiesApi.GetLegalEntity(context.Background(), id).Execute()

			require.NotNil(t, err)
			assert.Equal(t, 401, httpRes.StatusCode)
		})
	})

	t.Run("Test LegalEntitiesApiService GetAllBusinessLinesUnderLegalEntity", func(t *testing.T) {

		var id = os.Getenv("LEM_LEGAL_ENTITY_ID")

		resp, httpRes, err := apiClient.LegalEntitiesApi.GetAllBusinessLinesUnderLegalEntity(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
