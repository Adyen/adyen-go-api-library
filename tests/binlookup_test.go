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
	"github.com/adyen/adyen-go-api-library/src/binlookup"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Binlookup(t *testing.T) {
	godotenv.Load("./../.env")

	var (
		MerchantAccount = os.Getenv("ADYEN_MERCHANT")
		APIKey          = os.Getenv("ADYEN_API_KEY")
	)

	client := adyen.NewAPIClientWithAPIKey(APIKey, "TEST")
	client.GetConfig().Debug = true

	t.Run("Get 3DS Availability", func(t *testing.T) {
		t.Run("Create an API request that should support only 3DS1", func(t *testing.T) {
			res, httpRes, err := client.BinLookup.Get3dsAvailabilityPost(&binlookup.ThreeDSAvailabilityRequest{
				CardNumber:      "4111111111111111",
				MerchantAccount: MerchantAccount,
			})

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, "200 OK", httpRes.Status)
			assert.Equal(t, false, res.ThreeDS2supported)
			assert.Equal(t, true, res.ThreeDS1Supported)
		})
		t.Run("Create an API request that should support 3DS1 and 3DS2", func(t *testing.T) {
			res, httpRes, err := client.BinLookup.Get3dsAvailabilityPost(&binlookup.ThreeDSAvailabilityRequest{
				CardNumber:      "5454545454545454",
				MerchantAccount: MerchantAccount,
			})

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, "200 OK", httpRes.Status)
			assert.Equal(t, true, res.ThreeDS2supported)
			assert.Equal(t, true, res.ThreeDS1Supported)
		})
		t.Run("Create an API request that should not support 3DS1 and 3DS2", func(t *testing.T) {
			res, httpRes, err := client.BinLookup.Get3dsAvailabilityPost(&binlookup.ThreeDSAvailabilityRequest{
				CardNumber:      "4199350000000002",
				MerchantAccount: MerchantAccount,
			})

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, "200 OK", httpRes.Status)
			assert.Equal(t, false, res.ThreeDS2supported)
			assert.Equal(t, false, res.ThreeDS1Supported)
		})
	})
	t.Run("Get Cost Estimate", func(t *testing.T) {
		t.Run("Create an API request that should get an unsupported cost estimate", func(t *testing.T) {
			res, httpRes, err := client.BinLookup.GetCostEstimatePost(&binlookup.CostEstimateRequest{
				Amount: binlookup.Amount{
					Currency: "EUR",
					Value:    1250,
				},
				CardNumber:      "4111111111111111",
				MerchantAccount: MerchantAccount,
			})
			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, "200 OK", httpRes.Status)
			assert.Equal(t, "Unsupported", res.ResultCode)
		})
		t.Run("Create an API request that should get success cost estimate", func(t *testing.T) {
			res, httpRes, err := client.BinLookup.GetCostEstimatePost(&binlookup.CostEstimateRequest{
				Amount: binlookup.Amount{
					Currency: "EUR",
					Value:    1250,
				},
				CardNumber:      "5101180000000007",
				MerchantAccount: MerchantAccount,
			})
			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, "200 OK", httpRes.Status)
			assert.Equal(t, "Success", res.ResultCode)
		})
	})
}
