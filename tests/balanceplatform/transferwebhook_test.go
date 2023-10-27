package balanceplatform

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/adyen/adyen-go-api-library/v8/src/transferwebhook"
)

func TestHandleTransferNotificationRequest(t *testing.T) {
	t.Run("should return transfer success", func(t *testing.T) {
		notificationJson := `{
			"data": {
			  "balancePlatform": "YOUR_BALANCE_PLATFORM",
			  "accountHolder": {
				"id": "AH32272223222B59MTF7458DP"
			  },
			  "balanceAccount": {
				"id": "BA3227C223222B5B9SCR82TMV"
			  },
			  "creationDate": "2021-07-02T02:01:08+02:00",
			  "fileName": "balanceplatform_payments_accounting_report_2021_07_01.csv",
			  "type": "refund"
			},
			"environment": "test"
		  }
		`
		balancePlatformNotification, _ := transferwebhook.HandleTransferNotificationRequest(notificationJson)
		balancePlatformValue := "YOUR_BALANCE_PLATFORM"
		transferType := "refund"
		assert.Equal(t, &balancePlatformValue, balancePlatformNotification.Data.BalancePlatform)
		assert.Equal(t, &transferType, balancePlatformNotification.Data.Type)
	})
}
