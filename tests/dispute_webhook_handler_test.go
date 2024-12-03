package tests

import (
	"github.com/adyen/adyen-go-api-library/v15/src/disputewebhook"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Dispute_Webhook_HandleRequest(t *testing.T) {
	t.Run("on balancePlatform.dispute.created", func(t *testing.T) {
		notificationJson := `{
          "type": "balancePlatform.dispute.created",
          "data": {
            "balancePlatform": "YOUR_BALANCE_PLATFORM",
            "arn": "1 123456 2 123 4567890123456 1 1 1 1",
            "description": "Dispute request has been created",
            "disputedAmount": {
                "currency": "USD",
                "value": 1000
            },
            "id": "DS0001",
            "status": "Open",
            "statusDetail": "n/a",
            "transactionId": "n/a"
          }
        }`

		webhook, _ := disputewebhook.HandleDisputeNotificationRequest(notificationJson)

		assert.Equal(t, "balancePlatform.dispute.created", webhook.Type)
		// Environment missing!?
		// assert.Equal(t, "test", webhook.Environment)
		// CreatedAt missing!?
		// assert.Equal(t, 2022, webhook.CreatedAt.Year())
		assert.Equal(t, "DS0001", webhook.Data.GetId())
		assert.Equal(t, int64(1000), webhook.Data.DisputedAmount.GetValue())
	})
}
