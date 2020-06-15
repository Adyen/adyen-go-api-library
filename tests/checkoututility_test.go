/*
 * Adyen API Client
 *
 * Contact: support@adyen.com
 */

package tests

import (
	"os"
	"testing"

	"github.com/adyen/adyen-go-api-library/src/adyen"
	"github.com/adyen/adyen-go-api-library/src/checkoututility"
	"github.com/adyen/adyen-go-api-library/src/common"
	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_CheckoutUtility(t *testing.T) {
	godotenv.Load("./../.env")

	var (
		APIKey = os.Getenv("ADYEN_API_KEY")
	)

	client := adyen.NewClient(&common.Config{
		ApiKey:      APIKey,
		Environment: "TEST",
	})
	// client.GetConfig().Debug = true

	t.Run("OriginKeys", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			res, httpRes, err := client.CheckoutUtility.OriginKeys(&checkoututility.CheckoutUtilityRequest{})
			require.NotNil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, "500 Internal Server Error: Required field 'originDomains' is null (validation: 702)", err.Error())
			assert.Equal(t, 500, httpRes.StatusCode)
			assert.Equal(t, "500 Internal Server Error", httpRes.Status)
			require.NotNil(t, res)
		})
		t.Run("Create an API request that should pass", func(t *testing.T) {
			domain := "http://example.com"
			res, httpRes, err := client.CheckoutUtility.OriginKeys(&checkoututility.CheckoutUtilityRequest{
				OriginDomains: []string{domain},
			})
			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, "200 OK", httpRes.Status)
			require.NotNil(t, res)
			originKeys := *res.OriginKeys
			assert.NotEmpty(t, originKeys[domain])
		})
	})
}
