package compat

import (
	"strings"
	"testing"
	"time"

	"github.com/adyen/adyen-go-api-library/v21/src/acswebhook"
	"github.com/adyen/adyen-go-api-library/v21/src/balancewebhook"
	"github.com/adyen/adyen-go-api-library/v21/src/configurationwebhook"
	"github.com/adyen/adyen-go-api-library/v21/src/hmacvalidator"
	"github.com/adyen/adyen-go-api-library/v21/src/managementwebhook"
	"github.com/adyen/adyen-go-api-library/v21/src/reportwebhook"
	"github.com/adyen/adyen-go-api-library/v21/src/transactionwebhook"
	"github.com/adyen/adyen-go-api-library/v21/src/transferwebhook"
	"github.com/adyen/adyen-go-api-library/v21/src/webhook"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Synthetic test key, built at runtime so it is never mistaken for a credential.
// The signature was computed independently of this library (Python hmac/sha256).
var hmacKey = strings.Repeat("0123456789ABCDEF", 4)

const hmacSignature = "jvaDtTfp2bNC25w587i21cMKu826o37u+v5ByRzuMrc="

func TestClassicNotificationContract(t *testing.T) {
	n, err := webhook.HandleRequest(string(fixture(t, "classic_notification.json")))
	require.NoError(t, err)

	assert.Equal(t, "false", n.Live)
	items := n.GetNotificationItems()
	require.Len(t, items, 1)
	item := items[0]

	assert.Equal(t, "AUTHORISATION", item.EventCode)
	assert.Equal(t, "true", item.Success)
	assert.Equal(t, webhook.Amount{Currency: "EUR", Value: 1130}, item.Amount)
	assert.Equal(t, []string{"CANCEL", "CAPTURE", "REFUND"}, item.Operations)
	require.NotNil(t, item.EventDate)
	assert.True(t, time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC).Equal(*item.EventDate))

	// additionalData keeps arbitrary JSON values, not only strings.
	require.NotNil(t, item.AdditionalData)
	ad := *item.AdditionalData
	assert.Equal(t, hmacSignature, ad["hmacSignature"])
	assert.Equal(t, true, ad["recurring.enabled"])
	assert.Equal(t, float64(3), ad["installments.value"])
	assert.Equal(t, map[string]interface{}{"street": "Main"}, ad["nested"])

	assert.True(t, hmacvalidator.ValidateHmac(*item, hmacKey))
}

func TestHmacValidation(t *testing.T) {
	item := webhook.NotificationRequestItem{
		AdditionalData:      &map[string]interface{}{"hmacSignature": hmacSignature},
		Amount:              webhook.Amount{Currency: "EUR", Value: 1130},
		EventCode:           "AUTHORISATION",
		MerchantAccountCode: "TestMerchant",
		MerchantReference:   "TestPayment-1407325143704",
		PspReference:        "7914073381342284",
		Success:             "true",
	}
	assert.Equal(t, "7914073381342284::TestMerchant:TestPayment-1407325143704:1130:EUR:AUTHORISATION:true", hmacvalidator.GetDataToSign(item))

	sig, err := hmacvalidator.CalculateHmac(item, hmacKey)
	require.NoError(t, err)
	assert.Equal(t, hmacSignature, sig)
	assert.True(t, hmacvalidator.ValidateHmac(item, hmacKey))

	tampered := item
	tampered.Amount.Value = 1131
	assert.False(t, hmacvalidator.ValidateHmac(tampered, hmacKey))

	missing := item
	missing.AdditionalData = nil
	assert.NotPanics(t, func() { assert.False(t, hmacvalidator.ValidateHmac(missing, hmacKey)) })

	nonString := item
	nonString.AdditionalData = &map[string]interface{}{"hmacSignature": 42}
	assert.False(t, hmacvalidator.ValidateHmac(nonString, hmacKey))

	_, err = hmacvalidator.CalculateHmac(item, "not-hex")
	assert.Error(t, err)
}

func TestTypedWebhookHandlers(t *testing.T) {
	t.Run("payload", func(t *testing.T) {
		n, err := configurationwebhook.HandleBalanceAccountNotificationRequest(string(fixture(t, "balance_account_notification.json")))
		require.NoError(t, err)
		assert.Equal(t, "balancePlatform.balanceAccount.created", n.Type)
		assert.Equal(t, "test", n.Environment)
		assert.Equal(t, "YOUR_BALANCE_PLATFORM", n.Data.GetBalancePlatform())
		assert.Equal(t, "BA00000000000000000000001", n.Data.BalanceAccount.GetId())
	})

	malformed := `{"type": "x", "data": `
	handlers := map[string]func(string) (interface{}, error){
		"acs": func(s string) (interface{}, error) { return acswebhook.HandleAuthenticationNotificationRequest(s) },
		"balance": func(s string) (interface{}, error) {
			return balancewebhook.HandleBalanceAccountBalanceNotificationRequest(s)
		},
		"configuration.accountHolder": func(s string) (interface{}, error) {
			return configurationwebhook.HandleAccountHolderNotificationRequest(s)
		},
		"configuration.balanceAccount": func(s string) (interface{}, error) {
			return configurationwebhook.HandleBalanceAccountNotificationRequest(s)
		},
		"configuration.cardOrder": func(s string) (interface{}, error) { return configurationwebhook.HandleCardOrderNotificationRequest(s) },
		"configuration.payment":   func(s string) (interface{}, error) { return configurationwebhook.HandlePaymentNotificationRequest(s) },
		"configuration.sweep": func(s string) (interface{}, error) {
			return configurationwebhook.HandleSweepConfigurationNotificationRequest(s)
		},
		"management.merchantCreated": func(s string) (interface{}, error) {
			return managementwebhook.HandleMerchantCreatedNotificationRequest(s)
		},
		"management.merchantUpdated": func(s string) (interface{}, error) {
			return managementwebhook.HandleMerchantUpdatedNotificationRequest(s)
		},
		"management.pmCreated": func(s string) (interface{}, error) {
			return managementwebhook.HandlePaymentMethodCreatedNotificationRequest(s)
		},
		"management.pmRemoved": func(s string) (interface{}, error) {
			return managementwebhook.HandlePaymentMethodRequestRemovedNotificationRequest(s)
		},
		"management.pmScheduled": func(s string) (interface{}, error) {
			return managementwebhook.HandlePaymentMethodScheduledForRemovalNotificationRequest(s)
		},
		"report": func(s string) (interface{}, error) { return reportwebhook.HandleReportNotificationRequest(s) },
		"transaction": func(s string) (interface{}, error) {
			return transactionwebhook.HandleTransactionNotificationRequestV4(s)
		},
		"transfer": func(s string) (interface{}, error) { return transferwebhook.HandleTransferNotificationRequest(s) },
		"classic":  func(s string) (interface{}, error) { return webhook.HandleRequest(s) },
	}
	for name, handle := range handlers {
		t.Run("malformed "+name, func(t *testing.T) {
			res, err := handle(malformed)
			assert.Error(t, err)
			assert.Nil(t, res)
		})
	}
}
