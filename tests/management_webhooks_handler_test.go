package tests

import (
	"github.com/adyen/adyen-go-api-library/v8/src/managementwebhook"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Management_Webhooks_HandleRequest(t *testing.T) {
	t.Run("on merchant.created", func(t *testing.T) {
		notificationJson := `{
          "type": "merchant.created",
          "environment": "test",
          "createdAt": "2022-08-12T10:50:01+02:00",
          "data": {
            "capabilities": {
              "sendToTransferInstrument": {
                "requested": true,
                "requestedLevel": "notApplicable"
              }
            },
            "companyId": "YOUR_COMPANY_ID",
            "merchantId": "MC3224X22322535GH8D537TJR",
            "status": "PreActive"
          }
        }`

		webhook, _ := managementwebhook.HandleMerchantCreatedNotificationRequest(notificationJson)

		assert.Equal(t, "merchant.created", webhook.Type)
		assert.Equal(t, "test", webhook.Environment)
		assert.Equal(t, 2022, webhook.CreatedAt.Year())
		assert.Equal(t, "MC3224X22322535GH8D537TJR", webhook.Data.MerchantId)
		assert.True(t, webhook.Data.Capabilities["sendToTransferInstrument"].Requested)
	})
}
