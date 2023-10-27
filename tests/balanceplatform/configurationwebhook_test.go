package balanceplatform

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com///testify/require"

	"github.com/adyen/adyen-go-api-library/v8/src/configurationwebhook"
)

func TestHandleAccountHolderNotificationRequest(t *testing.T) {
	t.Run("should return accountHolder created success", func(t *testing.T) {
		notificationJson := `{
			"data": {
			  "balancePlatform": "YOUR_BALANCE_PLATFORM",
			  "accountHolder": {
				"contactDetails": {
				  "email": "test@adyen.com",
				  "phone": {
					"number": "0612345678",
					"type": "Mobile"
				  },
				  "address": {
					"houseNumberOrName": "23",
					"city": "Amsterdam",
					"country": "NL",
					"postalCode": "12345",
					"street": "Teststreet 1"
				  }
				},
				"description": "Shelly Eller",
				"legalEntityId": "LE43319330319C8AYX89L2V59",
				"reference": "YOUR_REFERENCE-2412C",
				"capabilities": {
				  "issueCard": {
					"enabled": true,
					"requested": true,
					"allowed": false,
					"verificationStatus": "pending"
				  },
				  "receiveFromTransferInstrument": {
					"enabled": true,
					"requested": true,
					"allowed": false,
					"verificationStatus": "pending"
				  },
				  "sendToTransferInstrument": {
					"enabled": true,
					"requested": true,
					"allowed": false,
					"verificationStatus": "pending"
				  },
				  "sendToBalanceAccount": {
					"enabled": true,
					"requested": true,
					"allowed": false,
					"verificationStatus": "pending"
				  },
				  "receiveFromBalanceAccount": {
					"enabled": true,
					"requested": true,
					"allowed": false,
					"verificationStatus": "pending"
				  }
				},
				"id": "AH32272223222B5CZW6QZ2V34",
				"status": "Active"
			  }
			},
			"environment": "test",
			"type": "balancePlatform.accountHolder.created"
		  }
		`
		balancePlatformNotification, _ := configurationwebhook.HandleAccountHolderNotificationRequest(notificationJson)
		balancePlatformValue := "YOUR_BALANCE_PLATFORM"
		assert.Equal(t, &balancePlatformValue, balancePlatformNotification.Data.BalancePlatform)
		assert.Equal(t, "test@adyen.com", balancePlatformNotification.Data.AccountHolder.ContactDetails.Email)
	})
}
