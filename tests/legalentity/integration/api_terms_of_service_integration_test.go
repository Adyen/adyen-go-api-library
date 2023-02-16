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

func Test_Integration_LegalEntity_TermsOfServiceApiService(t *testing.T) {

	godotenv.Load("./../../../.env")

	var (
		username = os.Getenv("ADYEN_LEM_USERNAME")
		password = os.Getenv("ADYEN_LEM_PASSWORD")
		env      = openapiclient.TestEnv
	)

	configuration, err := openapiclient.NewConfiguration(username, password, env)
	require.Nil(t, err, "Error creating Config object")
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TermsOfServiceApiService GetTermsOfServiceInformationForLegalEntity", func(t *testing.T) {

		var id = "LE322JV223222J5HFLRNL5VZ3"

		resp, httpRes, err := apiClient.TermsOfServiceApi.GetTermsOfServiceInformationForLegalEntity(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TermsOfServiceApiService GetTermsOfServiceStatus", func(t *testing.T) {

		// endpoint error - re-enable the test and verify server-side fix
		t.SkipNow()

		var id = "LE322JV223222J5HFLRNL5VZ3"

		resp, httpRes, err := apiClient.TermsOfServiceApi.GetTermsOfServiceStatus(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
