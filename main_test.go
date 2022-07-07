/*
 * Adyen API Client
 *
 * Contact: support@adyen.com
 */

package main

import (
	"strings"
	"testing"

	"github.com/adyen/adyen-go-api-library/v6/src/adyen"
	"github.com/adyen/adyen-go-api-library/v6/src/checkout"
	"github.com/adyen/adyen-go-api-library/v6/src/common"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_main(t *testing.T) {
	t.Run("Create a new APIClient", func(t *testing.T) {

		client := adyen.NewClient(&common.Config{
			Environment: "TEST",
		})
		require.NotNil(t, client)
		require.NotNil(t, client.Checkout)
		require.NotNil(t, client.Checkout.Client)
		require.NotNil(t, client.Checkout.Client.Cfg)
		require.Equal(t, common.TestEnv, client.Checkout.Client.Cfg.Environment)
		assert.Equal(t, "https://checkout-test.adyen.com/checkout/" + adyen.CheckoutAPIVersion, client.Checkout.BasePath())

		t.Run("Create a API request that should fail", func(t *testing.T) {
			res, httpRes, err := client.Checkout.PaymentMethods(&checkout.PaymentMethodsRequest{})
			require.NotNil(t, err)
			assert.Equal(t, true, strings.Contains(err.Error(), "Unauthorized"))
			require.NotNil(t, res)
			assert.Equal(t, checkout.PaymentMethodsResponse{}, res)
			require.NotNil(t, httpRes)
			assert.Equal(t, 401, httpRes.StatusCode)
		})
	})
}
