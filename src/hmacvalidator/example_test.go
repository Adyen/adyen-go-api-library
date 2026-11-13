package hmacvalidator_test

import (
	"fmt"

	"github.com/adyen/adyen-go-api-library/v21/src/hmacvalidator"
	"github.com/adyen/adyen-go-api-library/v21/src/transferwebhook"
	"github.com/adyen/adyen-go-api-library/v21/src/webhook"
)

const hmacKey = "0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF" // placeholder; use your HMAC key

// Validate the HMAC signature of each item in a standard webhook before acting on it. The
// signature is in the additionalData of the item.
func ExampleValidateHmac() {
	body := `{
  "live": "false",
  "notificationItems": [{
    "NotificationRequestItem": {
      "additionalData": {"hmacSignature": "M1eDlZlWw7YgVxTJHz3qdKsxzIHRRHDIiNKvMho3pI4="},
      "amount": {"currency": "EUR", "value": 1000},
      "eventCode": "AUTHORISATION",
      "eventDate": "2026-01-01T01:00:00+01:00",
      "merchantAccountCode": "YOUR_MERCHANT_ACCOUNT",
      "merchantReference": "order-12345",
      "paymentMethod": "visa",
      "pspReference": "PSP0001",
      "reason": "",
      "success": "true"
    }
  }]
}`
	notification, err := webhook.HandleRequest(body)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, item := range *notification.NotificationItems {
		request := item.NotificationRequestItem
		if !hmacvalidator.ValidateHmac(request, hmacKey) {
			fmt.Println("invalid signature, ignoring", request.PspReference)
			continue
		}
		if request.EventCode == webhook.EventCodeAuthorisation {
			fmt.Println("authorised:", request.PspReference, request.Success)
		}
	}
	// Output: authorised: PSP0001 true
}

// Validate a webhook that carries its HMAC signature in the HmacSignature HTTP header, such as
// the Balance Platform and Management webhooks, before parsing it with the handler of its package.
func ExampleValidateHmacPayload() {
	body := `{"type":"balancePlatform.transfer.created","environment":"test","data":{"id":"TRANSFER0001","status":"received","amount":{"currency":"EUR","value":1000}}}`
	signatureHeader := "Gllcu6qz73yiWnaipdT4qDeDR5xIWi4Myl7fLKA9Jaw="

	if !hmacvalidator.ValidateHmacPayload(signatureHeader, hmacKey, body) {
		fmt.Println("invalid signature")
		return
	}
	notification, err := transferwebhook.HandleTransferNotificationRequest(body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if notification.GetType() == transferwebhook.TransferNotificationRequestTypeBalancePlatformTransferCreated {
		fmt.Println("transfer created:", notification.Data.GetId(), notification.Data.Status)
	}
	// Output: transfer created: TRANSFER0001 received
}
