//go:build integration
// +build integration

package tests

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/adyen/adyen-go-api-library/v8/src/adyen"
	"github.com/adyen/adyen-go-api-library/v8/src/common"

	"github.com/adyen/adyen-go-api-library/v8/src/recurring"
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
		Debug:       "true" == os.Getenv("DEBUG"),
	})
	service := client.Recurring()

	t.Run("ListRecurringDetails", func(t *testing.T) {
		t.Run("Create an API request that should fail", func(t *testing.T) {
			req := service.ListRecurringDetailsInput().
				RecurringDetailsRequest(*recurring.NewRecurringDetailsRequestWithDefaults())

			res, httpRes, err := service.ListRecurringDetails(context.Background(), req)

			require.NotNil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, true, strings.Contains(err.Error(), "Invalid Merchant Account"))
			assert.Equal(t, 422, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, "Invalid Merchant Account", err.(common.APIError).Message)
		})

		t.Run("Create an API request that should pass", func(t *testing.T) {
			req := service.ListRecurringDetailsInput()
			req = req.RecurringDetailsRequest(recurring.RecurringDetailsRequest{
				MerchantAccount:  MerchantAccount,
				ShopperReference: "4343553GFGFYFY4654654675765",
				Recurring: &recurring.Recurring{
					Contract: common.PtrString("RECURRING"),
				},
			})

			res, httpRes, err := service.ListRecurringDetails(context.Background(), req)

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
			body := recurring.DisableRequest{
				MerchantAccount:  MerchantAccount,
				ShopperReference: "4343553GFGFYFY4654654675765",
			}
			body.SetRecurringDetailReference("8314442372419167")
			req := service.DisableInput().DisableRequest(body)

			res, httpRes, err := service.Disable(context.Background(), req)

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
			body := recurring.NotifyShopperRequest{
				Amount: recurring.Amount{
					Currency: "INR",
					Value:    1234,
				},
				BillingDate:              common.PtrString("2030-12-31"),
				BillingSequenceNumber:    common.PtrString("adhoc"),
				MerchantAccount:          MerchantAccount,
				Reference:                "4343553GFGFYFY4654654675765",
				ShopperReference:         "4343553GFGFYFY4654654675765",
				RecurringDetailReference: common.PtrString("8314442372419167"),
				StoredPaymentMethodId:    common.PtrString("8314442372419167"),
			}
			req := service.NotifyShopperInput().NotifyShopperRequest(body)

			res, httpRes, err := service.NotifyShopper(context.Background(), req)

			require.NotNil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, true, strings.Contains(err.Error(), "PaymentDetail not found"))
			assert.Equal(t, 500, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, "PaymentDetail not found", err.(common.APIError).Message)
		})

		t.Run("Create an API request that should fail because of invalid date", func(t *testing.T) {
			body := recurring.NotifyShopperRequest{
				Amount: recurring.Amount{
					Currency: "INR",
					Value:    1234,
				},
				MerchantAccount:  MerchantAccount,
				Reference:        "4343553GFGFYFY4654654675765",
				ShopperReference: "4343553GFGFYFY4654654675765",
			}
			body.SetBillingDate("2")
			body.SetBillingSequenceNumber("adhoc")
			body.SetRecurringDetailReference("8314442372419167")
			body.SetStoredPaymentMethodId("8314442372419167")
			req := service.NotifyShopperInput().NotifyShopperRequest(body)

			res, httpRes, err := service.NotifyShopper(context.Background(), req)

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
			body := recurring.ScheduleAccountUpdaterRequest{
				MerchantAccount: MerchantAccount,
				Reference:       "4343553GFGFYFY4654654675765",
			}
			req := service.ScheduleAccountUpdaterInput().ScheduleAccountUpdaterRequest(body)

			res, httpRes, err := service.ScheduleAccountUpdater(context.Background(), req)

			require.NotNil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, true, strings.Contains(err.Error(), "Missing card or shopperReference & selectedRecurringDetailReference"))
			assert.Equal(t, 422, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, "Missing card or shopperReference & selectedRecurringDetailReference", err.(common.APIError).Message)
		})

		t.Run("Create an API request that should fail with card", func(t *testing.T) {
			card := &recurring.Card{}
			card.SetExpiryMonth("99")
			card.SetExpiryYear("2030")
			card.SetHolderName("Adyen test")
			card.SetNumber("4111111111111111")
			body := recurring.ScheduleAccountUpdaterRequest{
				MerchantAccount: MerchantAccount,
				Reference:       "4343553GFGFYFY4654654675765",
				Card:            card,
			}
			req := service.ScheduleAccountUpdaterInput().ScheduleAccountUpdaterRequest(body)

			res, httpRes, err := service.ScheduleAccountUpdater(context.Background(), req)

			require.NotNil(t, err)
			require.NotNil(t, httpRes)
			assert.Equal(t, true, strings.Contains(err.Error(), "Expiry month should be between 1 and 12 inclusive: 99"))
			assert.Equal(t, 422, httpRes.StatusCode)
			require.NotNil(t, res)
			assert.Equal(t, "Expiry month should be between 1 and 12 inclusive: 99", err.(common.APIError).Message)
		})
	})
}
