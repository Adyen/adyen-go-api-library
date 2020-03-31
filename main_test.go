/*
 * Adyen API Client
 *
 * Contact: support@adyen.com
 */

package main

import (
	"os"
	"testing"

	"github.com/adyen/adyen-go-api-library/src/checkout"
	"github.com/joho/godotenv"

	common "github.com/adyen/adyen-go-api-library/src/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_main(t *testing.T) {
	t.Run("Create a new default config object", func(t *testing.T) {
		want := &common.Config{
			BasePath:      "https://checkout-test.adyen.com/v51",
			DefaultHeader: make(map[string]string),
			UserAgent:     "OpenAPI-Generator/1.0.0/go",
			Debug:         false,
			Servers: []common.ServerConfiguration{
				{
					Url:         "https://checkout-test.adyen.com/v51",
					Description: "No description provided",
				},
			},
		}
		config := NewConfig()
		require.NotNil(t, config)
		assert.Equal(t, want, config)
	})
	t.Run("Create a new APIClient", func(t *testing.T) {

		config := NewConfig()

		client := NewAPIClient(config)
		require.NotNil(t, client)
		require.NotNil(t, client.client)
		require.NotNil(t, client.client.Cfg)
		require.NotNil(t, client.client.Cfg.HTTPClient)
		require.NotNil(t, client.Checkout)

		t.Run("Create a API request that should fail", func(t *testing.T) {

			res, httpRes, err := client.Checkout.PaymentMethodsPost(&checkout.PaymentMethodsRequest{})
			require.NotNil(t, err)
			assert.Equal(t, "401 Unauthorized", err.Error())
			require.NotNil(t, res)
			assert.Equal(t, checkout.PaymentMethodsResponse{}, res)
			require.NotNil(t, httpRes)
			assert.Equal(t, 401, httpRes.StatusCode)
			assert.Equal(t, "401 Unauthorized", httpRes.Status)
		})

		godotenv.Load("./.env")

		var (
			MerchantAccount = os.Getenv("ADYEN_MERCHANT")
			APIKey          = os.Getenv("ADYEN_API_KEY")
		)

		config.AddDefaultHeader("x-API-key", APIKey)

		t.Run("Create a API request that should pass", func(t *testing.T) {

			res, httpRes, err := client.Checkout.PaymentMethodsPost(&checkout.PaymentMethodsRequest{
				MerchantAccount: MerchantAccount,
			})

			require.Nil(t, err)
			require.NotNil(t, res)
			assert.True(t, len(res.Groups) > 1)
			assert.True(t, len(res.PaymentMethods) > 1)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, "200 OK", httpRes.Status)
		})
	})
}
