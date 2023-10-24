package main

import (
	"context"
	"strings"
	"testing"

	"github.com/adyen/adyen-go-api-library/v8/src/adyen"
	"github.com/adyen/adyen-go-api-library/v8/src/checkout"
	"github.com/adyen/adyen-go-api-library/v8/src/common"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_main(t *testing.T) {
	t.Run("Create a new APIClient", func(t *testing.T) {
		client := adyen.NewClient(&common.Config{
			Environment: "TEST",
		})

		require.NotNil(t, client)
		checkoutClient := client.Checkout()
		require.NotNil(t, checkoutClient)
		require.NotNil(t, checkoutClient.PaymentsApi)
		require.NotNil(t, checkoutClient.PaymentsApi.Client.Cfg)
		require.Equal(t, common.TestEnv, checkoutClient.PaymentLinksApi.Client.Cfg.Environment)
		assert.Equal(t, "https://checkout-test.adyen.com/checkout/"+adyen.CheckoutAPIVersion, checkoutClient.ModificationsApi.BasePath())

		t.Run("Create a API request that should fail", func(t *testing.T) {
			req := checkoutClient.PaymentsApi.PaymentMethodsInput().PaymentMethodsRequest(checkout.PaymentMethodsRequest{})

			res, httpRes, err := checkoutClient.PaymentsApi.PaymentMethods(context.Background(), req)

			require.NotNil(t, err)
			assert.Equal(t, true, strings.Contains(err.Error(), "Unauthorized"))
			require.NotNil(t, res)
			assert.Equal(t, checkout.PaymentMethodsResponse{}, res)
			require.NotNil(t, httpRes)
			assert.Equal(t, 401, httpRes.StatusCode)
		})
	})
}
