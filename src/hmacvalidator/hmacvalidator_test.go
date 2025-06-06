package hmacvalidator

import (
	"testing"
	"time"

	"github.com/adyen/adyen-go-api-library/v21/src/webhook"
	"github.com/stretchr/testify/assert"
)

const key = "DFB1EB5485895CFA84146406857104ABB4CBCABDC8AAF103A624C8F6A3EAAB00"
const expectedSign = "ipnxGCaUZ4l8TUW75a71/ghd2Fe5ffvX0pV4TLTntIc="

var eventDate = time.Date(1970, time.January, 01, 0, 0, 0, 0, time.UTC)
var notificationRequestItem = webhook.NotificationRequestItem{
	AdditionalData: &map[string]interface{}{"hmacSignature": expectedSign},
	Amount: webhook.Amount{
		Currency: "EUR",
		Value:    1000,
	},
	EventCode:           "EVENT",
	EventDate:           &eventDate,
	MerchantAccountCode: "merchantAccount",
	MerchantReference:   "reference",
	OriginalReference:   "originalReference",
	PaymentMethod:       "VISA",
	PspReference:        "pspReference",
	Reason:              "reason",
	Success:             "true",
}

// webhook payload as string
const payloadAsString = `{
    "type": "merchant.created",
    "environment": "test",
    "createdAt": "01-01-2024",
    "data": {
        "capabilities": {
            "sendToTransferInstrument": {
                "requested": true,
                "requestedLevel": "notApplicable"
            }
        },
        "companyId": "YOUR_COMPANY_ID",
        "merchantId": YOUR_MERCHANT_ACCOUNT",
        "status": "PreActive"
    }
}`

// signature for the webhook payload above
const expectedSignFromPayloadAsString = "yitBTPnDGtbAMG9rTxUa1e6lFGpa4PnrD6rErOmEhJ8="

func Test_Hmacvalidator(t *testing.T) {
	t.Run("GetDataToSign", func(t *testing.T) {
		t.Run("Get correct data", func(t *testing.T) {
			data := map[string]string{"MerchantAccount": "ACC", "CurrencyCode": "EUR"}
			dataToSign := GetDataToSign(data)
			assert.Equal(t, "CurrencyCode:MerchantAccount:EUR:ACC", dataToSign)
		})
		t.Run("Get correct data with escaped characters", func(t *testing.T) {
			data := map[string]string{"CurrencyCode": "EUR", "MerchantAccount": "ACC:\\", "SessionValidity": "2019-09-21T11:45:24.637Z"}
			dataToSign := GetDataToSign(data)
			assert.Equal(t, "CurrencyCode:MerchantAccount:SessionValidity:EUR:ACC\\:\\\\:2019-09-21T11\\:45\\:24.637Z", dataToSign)
		})
		t.Run("Get correct data to sign", func(t *testing.T) {
			data := GetDataToSign(notificationRequestItem)
			assert.Equal(t, "pspReference:originalReference:merchantAccount:reference:1000:EUR:EVENT:true", data)
		})
	})
	t.Run("CalculateHmac", func(t *testing.T) {
		t.Run("Encrypt correctly", func(t *testing.T) {
			data := "countryCode:currencyCode:merchantAccount:merchantReference:paymentAmount:sessionValidity:skinCode:NL:EUR:MagentoMerchantTest2:TEST-PAYMENT-2017-02-01-14\\:02\\:05:199:2017-02-02T14\\:02\\:05+01\\:00:PKz2KML1"
			encrypted, err := CalculateHmac(data, key)
			assert.Nil(t, err)
			assert.Equal(t, "34oR8T1whkQWTv9P+SzKyp8zhusf9n0dpqrm9nsqSJs=", encrypted)
		})
		t.Run("Get Valid HMAC", func(t *testing.T) {
			enc, err := CalculateHmac(notificationRequestItem, key)
			assert.Nil(t, err)
			assert.Equal(t, expectedSign, enc)
		})
	})
	t.Run("ValidateHmac", func(t *testing.T) {
		t.Run("Validate HMAC", func(t *testing.T) {
			assert.True(t, ValidateHmac(notificationRequestItem, key))
		})
		t.Run("Get Invalid HMAC", func(t *testing.T) {
			notificationRequestItem.AdditionalData = &map[string]interface{}{"hmacSignature": "InvalidSignature"}
			assert.False(t, ValidateHmac(notificationRequestItem, key))
		})
		t.Run("Nil AdditionalData", func(t *testing.T) {
			// Create a copy and set AdditionalData to nil.
			testItem := notificationRequestItem
			testItem.AdditionalData = nil
			assert.False(t, ValidateHmac(testItem, key))
		})
	})
	t.Run("ValidateHmacPayload", func(t *testing.T) {
		t.Run("Validate HMAC from string payload", func(t *testing.T) {
			assert.True(t, ValidateHmacPayload(expectedSignFromPayloadAsString, key, payloadAsString))
		})
		t.Run("Test invalid HMAC from string payload", func(t *testing.T) {
			assert.False(t, ValidateHmacPayload("MismatchingHmacKey=", key, payloadAsString))
		})
	})
}
