/*
 * Adyen API Client
 *
 * Contact: support@adyen.com
 */

package adyen

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/adyen/adyen-go-api-library/v6/src/checkout"
	"github.com/adyen/adyen-go-api-library/v6/src/common"
	"github.com/adyen/adyen-go-api-library/v6/src/recurring"
	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_api(t *testing.T) {
	t.Run("Create a new APIClient", func(t *testing.T) {

		client := NewClient(&common.Config{
			Environment: "TEST",
		})
		require.NotNil(t, client)
		require.NotNil(t, client.client)
		require.NotNil(t, client.client.Cfg)
		require.NotNil(t, client.client.Cfg.HTTPClient)
		require.NotNil(t, client.Checkout)
		assert.Equal(t, "https://checkout-test.adyen.com/checkout/"+ CheckoutAPIVersion, client.Checkout.BasePath())

		t.Run("Create a API request that should fail", func(t *testing.T) {
			res, httpRes, err := client.Checkout.PaymentMethods(&checkout.PaymentMethodsRequest{})
			require.NotNil(t, err)
			assert.Equal(t, true, strings.Contains(err.Error(), "Unauthorized"))
			require.NotNil(t, res)
			assert.Equal(t, checkout.PaymentMethodsResponse{}, res)
			require.NotNil(t, httpRes)
			assert.Equal(t, 401, httpRes.StatusCode)
		})

		godotenv.Load("./../../.env")

		var (
			MerchantAccount = os.Getenv("ADYEN_MERCHANT")
			APIKey          = os.Getenv("ADYEN_API_KEY")
			USER            = os.Getenv("ADYEN_USER")
			PASS            = os.Getenv("ADYEN_PASSWORD")
		)

		client = NewClient(&common.Config{
			ApiKey:      APIKey,
			Environment: "TEST",
		})

		t.Run("Create a API request that uses API key auth and should pass", func(t *testing.T) {

			res, httpRes, err := client.Checkout.PaymentMethods(&checkout.PaymentMethodsRequest{
				MerchantAccount: MerchantAccount,
			})

			require.Nil(t, err)
			require.NotNil(t, res)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
		})

		client = NewClient(&common.Config{
			Username:    USER,
			Password:    PASS,
			Environment: "TEST",
		})

		t.Run("Create a API request that uses basic auth and should pass", func(t *testing.T) {
			res, httpRes, err := client.Recurring.ListRecurringDetails(&recurring.RecurringDetailsRequest{
				MerchantAccount: MerchantAccount,
				Recurring: &recurring.RecurringType{
					Contract: "RECURRING",
				},
				ShopperReference: time.Now().String(),
			})

			require.Nil(t, err)
			require.NotNil(t, res)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
		})
	})
}
