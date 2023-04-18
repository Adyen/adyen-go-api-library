/*
 * Adyen API Client
 *
 * Contact: support@adyen.com
 */

package tests

import (
	"os"
	"strings"
	"testing"

	"github.com/adyen/adyen-go-api-library/v6/src/adyen"
	"github.com/adyen/adyen-go-api-library/v6/src/common"

	"github.com/adyen/adyen-go-api-library/v6/src/recurring"
	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Recurring(t *testing.T) {
	godotenv.Load("./../.env")

	var (
		APIKey          = os.Getenv("ADYEN_API_KEY")
		MerchantAccount = os.Getenv("ADYEN_MERCHANT")
	)

	client := adyen.NewClient(&common.Config{
		ApiKey:      APIKey,
		Environment: "TEST",
	})
	// client.GetConfig().Debug = true

	t.Run("ListRecurringDetails", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			res, httpRes, err := client.Recurring.ListRecurringDetails(&recurring.RecurringDetailsRequest{})
			require.NotNil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, true, strings.Contains(err.Error(), "Invalid Merchant Account"))
			assert.Equal(t, 422, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, "Invalid Merchant Account", err.(common.APIError).Message)
		})
		t.Run("Create an API request that should pass", func(t *testing.T) {
			res, httpRes, err := client.Recurring.ListRecurringDetails(&recurring.RecurringDetailsRequest{
				MerchantAccount:  MerchantAccount,
				ShopperReference: "4343553GFGFYFY4654654675765",
				Recurring: &recurring.Recurring{
					Contract: "RECURRING",
				},
			})
			require.Nil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, 200, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Empty(t, res.ShopperReference)
			assert.Empty(t, res.Details)
		})
	})

	t.Run("Disable", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			res, httpRes, err := client.Recurring.Disable(&recurring.DisableRequest{
				MerchantAccount:          MerchantAccount,
				ShopperReference:         "4343553GFGFYFY4654654675765",
				RecurringDetailReference: "8314442372419167",
			})
			require.NotNil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, true, strings.Contains(err.Error(), "Contract not found"))
			assert.Equal(t, 422, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, "Contract not found", err.(common.APIError).Message)
		})
	})

	t.Run("NotifyShopper", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			res, httpRes, err := client.Recurring.NotifyShopper(&recurring.NotifyShopperRequest{
				Amount: &recurring.Amount{
					Currency: "INR",
					Value:    1234,
				},
				BillingDate:              "2030-12-31",
				BillingSequenceNumber:    "adhoc",
				MerchantAccount:          MerchantAccount,
				Reference:                "4343553GFGFYFY4654654675765",
				ShopperReference:         "4343553GFGFYFY4654654675765",
				RecurringDetailReference: "8314442372419167",
				StoredPaymentMethodID:    "8314442372419167",
			})
			require.NotNil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, true, strings.Contains(err.Error(), "PaymentDetail not found"))
			assert.Equal(t, 500, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, "PaymentDetail not found", err.(common.APIError).Message)
		})

		t.Run("Create an API request that should fail because of invalid date", func(t *testing.T) {
			res, httpRes, err := client.Recurring.NotifyShopper(&recurring.NotifyShopperRequest{
				Amount: &recurring.Amount{
					Currency: "INR",
					Value:    1234,
				},
				BillingDate:              "2",
				BillingSequenceNumber:    "adhoc",
				MerchantAccount:          MerchantAccount,
				Reference:                "4343553GFGFYFY4654654675765",
				ShopperReference:         "4343553GFGFYFY4654654675765",
				RecurringDetailReference: "8314442372419167",
				StoredPaymentMethodID:    "8314442372419167",
			})
			require.NotNil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, true, strings.Contains(err.Error(), "Inner validation error(s) occurred: billingDate should be in yyyy-MM-dd format"))
			assert.Equal(t, 422, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, "Inner validation error(s) occurred: billingDate should be in yyyy-MM-dd format", err.(common.APIError).Message)
		})
	})

	t.Run("ScheduleAccountUpdater", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			res, httpRes, err := client.Recurring.ScheduleAccountUpdater(&recurring.ScheduleAccountUpdaterRequest{
				MerchantAccount: MerchantAccount,
				Reference:       "4343553GFGFYFY4654654675765",
				// RecurringDetailReference: "8314442372419167",
			})
			require.NotNil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, true, strings.Contains(err.Error(), "Missing card or shopperReference & selectedRecurringDetailReference"))
			assert.Equal(t, 422, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, "Missing card or shopperReference & selectedRecurringDetailReference", err.(common.APIError).Message)
		})
		t.Run("Create an API request that should fail with card", func(t *testing.T) {
			res, httpRes, err := client.Recurring.ScheduleAccountUpdater(&recurring.ScheduleAccountUpdaterRequest{
				MerchantAccount: MerchantAccount,
				Reference:       "4343553GFGFYFY4654654675765",
				Card: &recurring.Card{
					ExpiryMonth: "99",
					ExpiryYear:  "2030",
					HolderName:  "Adyen test",
					Number:      "4111111111111111",
				},
			})
			require.NotNil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, true, strings.Contains(err.Error(), "Expiry month should be between 1 and 12 inclusive: 99"))
			assert.Equal(t, 422, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, "Expiry month should be between 1 and 12 inclusive: 99", err.(common.APIError).Message)
		})
	})

	t.Run("CreatePermit", func(t *testing.T) {

	})

	t.Run("DisablePermit", func(t *testing.T) {

	})
}
