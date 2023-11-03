package balanceplatform

import (
	"github.com/adyen/adyen-go-api-library/v8/src/acswebhook"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleAuthenticationNotificationRequest(t *testing.T) {
	t.Run("on balancePlatform.authentication.created", func(t *testing.T) {
		notificationJson := `{
          "data": {
            "authentication": {
              "acsTransID": "6a4c1709-a42e-4c7f-96c7-1043adacfc97",
              "challengeIndicator": "01",
              "createdAt": "2022-12-22T15:45:03+01:00",
              "deviceChannel": "app",
              "dsTransID": "a3b86754-444d-46ca-95a2-ada351d3f42c",
              "exemptionIndicator": "lowValue",
              "inPSD2scope": true,
              "messageCategory": "payment",
              "messageVersion": "2.2.0",
              "riskScore": 0,
              "threeDSServerTransID": "6edcc246-23ee-4e94-ac5d-8ae620bea7d9",
              "transStatus": "Y",
              "type": "frictionless"
            },
            "balancePlatform": "YOUR_BALANCE_PLATFORM",
            "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
            "paymentInstrumentId": "PI3227C223222B5BPCMFXD2XG",
            "purchase": {
              "date": "2022-12-22T15:49:03+01:00",
              "merchantName": "TeaShop_NL",
              "originalAmount": {
                "currency": "EUR",
                "value": 1000
              }
            },
            "status": "authenticated"
          },
          "environment": "test",
          "type": "balancePlatform.authentication.created"
        }`

		webhook, _ := acswebhook.HandleAuthenticationNotificationRequest(notificationJson)

		assert.Equal(t, "balancePlatform.authentication.created", webhook.Type)
		assert.Equal(t, "authenticated", webhook.Data.Status)
		assert.Equal(t, 2022, webhook.Data.Authentication.CreatedAt.Year())
		assert.True(t, webhook.Data.Authentication.InPSD2Scope)
	})
}
