package balanceplatform

import (
	"github.com/adyen/adyen-go-api-library/v21/src/balancewebhook"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleBalanceAccountBalanceNotificationRequest(t *testing.T) {
	t.Run("on balanceAccount.balance.updated", func(t *testing.T) {
		notificationJson := `{
          "data": {
            "balanceAccountId": "BWHS00000000000000000000000001",
            "balancePlatform": "YOUR_BALANCE_PLATFORM",
            "balances": {
              "available": 499900,
              "pending": 350000,
              "reserved": 120000,
              "balance": 470000
            },
            "creationDate": "2025-01-19T13:37:38+02:00",
            "currency": "USD",
            "settingIds": ["WK1", "WK2"]
          },
          "environment": "test",
          "type": "balancePlatform.balanceAccount.balance.updated"
        }`

		webhook, _ := balancewebhook.HandleBalanceAccountBalanceNotificationRequest(notificationJson)

		assert.Equal(t, "balancePlatform.balanceAccount.balance.updated", webhook.Type)
		assert.Equal(t, "BWHS00000000000000000000000001", webhook.Data.BalanceAccountId)
		assert.True(t, len(webhook.Data.SettingIds) > 0)
	})
}
