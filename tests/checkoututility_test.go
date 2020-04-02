/*
 * Adyen API Client
 *
 * Contact: support@adyen.com
 */

package tests

import (
	"os"
	"testing"

	adyen "github.com/adyen/adyen-go-api-library/src/api"
	"github.com/adyen/adyen-go-api-library/src/checkoututility"
	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_CheckoutUtility(t *testing.T) {
	godotenv.Load("./../env")

	var (
		APIKey = os.Getenv("ADYEN_API_KEY")
	)

	client := adyen.NewAPIClientWithAPIKey(APIKey, "TEST")

	t.Run("OriginKeys", func(t *testing.T) {
		t.Run("Create an API request that should pass", func(t *testing.T) {
			domain := "http.example.com"
			res, httpRes, err := client.CheckoutUtility.OriginKeysPost(&checkoututility.CheckoutUtilityRequest{
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
