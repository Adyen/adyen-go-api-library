/*
 * Adyen API Client
 *
 * Contact: support@adyen.com
 */

package tests

import (
	"os"
	"testing"

	"github.com/adyen/adyen-go-api-library/src/common"

	adyen "github.com/adyen/adyen-go-api-library/src/api"
	"github.com/adyen/adyen-go-api-library/src/recurring"
	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Recurring(t *testing.T) {
	godotenv.Load("./../.env")

	var (
		// APIKey = os.Getenv("ADYEN_API_KEY")
		USER = os.Getenv("ADYEN_USER")
		PASS = os.Getenv("ADYEN_PASSWORD")
	)

	client := adyen.NewAPIClient(&common.Config{
		Environment: "TEST",
		Username:    USER,
		Password:    PASS,
	})
	t.Run("ListRecurringDetails", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			res, httpRes, err := client.Recurring.ListRecurringDetailsPost(&recurring.RecurringDetailsRequest{})
			require.NotNil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, "500 Internal Server Error", err.Error())
			assert.Equal(t, 500, httpRes.StatusCode)
			assert.Equal(t, "500 Internal Server Error", httpRes.Status)
			require.NotNil(t, res)
		})
		t.Run("Create an API request that should pass", func(t *testing.T) {
			res, httpRes, err := client.Recurring.ListRecurringDetailsPost(&recurring.RecurringDetailsRequest{})
			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, "200 OK", httpRes.Status)
			require.NotNil(t, res)
		})
	})
}
