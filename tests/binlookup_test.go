//go:build integration
// +build integration

package tests

import (
	"context"
	"github.com/adyen/adyen-go-api-library/v8/src/binlookup"
	"os"
	"testing"

	"github.com/adyen/adyen-go-api-library/v8/src/adyen"
	"github.com/adyen/adyen-go-api-library/v8/src/common"
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

	client := adyen.NewClient(&common.Config{
		ApiKey:      APIKey,
		Environment: "TEST",
		Debug:       "true" == os.Getenv("DEBUG"),
	})

	t.Run("Get 3DS Availability", func(t *testing.T) {
		t.Run("Create an API request that should support only 3DS1", func(t *testing.T) {
			cardNumber := "4111111111111111"
			body := binlookup.ThreeDSAvailabilityRequest{
				CardNumber:      &cardNumber,
				MerchantAccount: MerchantAccount,
			}
			req := client.BinLookup().Get3dsAvailabilityInput().ThreeDSAvailabilityRequest(body)
			res, httpRes, err := client.BinLookup().Get3dsAvailability(context.Background(), req)

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, false, *res.ThreeDS2supported)
			assert.Equal(t, true, *res.ThreeDS1Supported)
		})
		t.Run("Create an API request that should support 3DS1 and 3DS2", func(t *testing.T) {
			cardNumber := "4917610000000000"
			body := binlookup.ThreeDSAvailabilityRequest{
				CardNumber:      &cardNumber,
				MerchantAccount: MerchantAccount,
			}
			req := client.BinLookup().Get3dsAvailabilityInput().ThreeDSAvailabilityRequest(body)

			res, httpRes, err := client.BinLookup().Get3dsAvailability(context.Background(), req)

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, true, *res.ThreeDS2supported)
			assert.Equal(t, true, *res.ThreeDS1Supported)
		})
		t.Run("Create an API request that should not support 3DS1 and 3DS2", func(t *testing.T) {
			cardNumber := "4199350000000002"
			body := binlookup.ThreeDSAvailabilityRequest{
				CardNumber:      &cardNumber,
				MerchantAccount: MerchantAccount,
			}
			req := client.BinLookup().Get3dsAvailabilityInput().ThreeDSAvailabilityRequest(body)

			res, httpRes, err := client.BinLookup().Get3dsAvailability(context.Background(), req)

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, false, *res.ThreeDS2supported)
			assert.Equal(t, false, *res.ThreeDS1Supported)
		})
	})
	t.Run("Get Cost Estimate", func(t *testing.T) {
		t.Run("Create an API request that should get an unsupported cost estimate", func(t *testing.T) {
			cardNumber := "4111111111111111"
			body := binlookup.CostEstimateRequest{
				Amount: binlookup.Amount{
					Currency: "EUR",
					Value:    1250,
				},
				CardNumber:      &cardNumber,
				MerchantAccount: MerchantAccount,
			}
			req := client.BinLookup().GetCostEstimateInput().CostEstimateRequest(body)

			res, httpRes, err := client.BinLookup().GetCostEstimate(context.Background(), req)

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, "Unsupported", *res.ResultCode)
		})
		t.Run("Create an API request that should get success cost estimate", func(t *testing.T) {
			cardNumber := "5101180000000007"
			body := binlookup.CostEstimateRequest{
				Amount: binlookup.Amount{
					Currency: "EUR",
					Value:    1250,
				},
				CardNumber:      &cardNumber,
				MerchantAccount: MerchantAccount,
			}
			req := client.BinLookup().GetCostEstimateInput().CostEstimateRequest(body)

			res, httpRes, err := client.BinLookup().GetCostEstimate(context.Background(), req)

			require.Nil(t, err)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, "Success", *res.ResultCode)
		})
	})
}
