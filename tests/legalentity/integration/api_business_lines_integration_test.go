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

func Test_Integration_LegalEntity_BusinessLinesApiService(t *testing.T) {

	godotenv.Load("./../../../.env")

	var (
		username = os.Getenv("ADYEN_LEM_USERNAME")
		password = os.Getenv("ADYEN_LEM_PASSWORD")
		env      = openapiclient.TestEnv
	)

	configuration, err := openapiclient.NewConfiguration(username, password, env)
	require.Nil(t, err, "Error creating Config object")
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BusinessLinesApiService GetBusinessLine", func(t *testing.T) {

		var id = os.Getenv("LEM_BUSINESS_LINE_ID")

		resp, httpRes, err := apiClient.BusinessLinesApi.GetBusinessLine(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
