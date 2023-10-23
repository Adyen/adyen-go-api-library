package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/adyen/adyen-go-api-library/v8/src/webhook"
)

func TestWebhook_HandleRequest(t *testing.T) {
	tests := []struct {
		name    string
		req     string
		want    func(got *webhook.Webhook, t *testing.T)
		wantErr bool
	}{
		{
			"should return authorisation success",
			`{
                "live": "false",
                "notificationItems": [
                  {
                    "NotificationRequestItem": {
                      "additionalData": {
                        "expiryDate": "12\/2012",
                        " NAME1 ": "VALUE1",
                        "authCode": "1234",
                        "cardSummary": "7777",
                        "totalFraudScore": "10",
                        "hmacSignature": "OzDjCMZIsdtDqrZ+cl\/FWC+WdESrorctXTzAzW33dXI=",
                        "NAME2": "  VALUE2  ",
                        "fraudCheck-6-ShopperIpUsage": "10"
                      },
                      "amount": {
                        "currency": "EUR",
                        "value": 10100
                      },
                      "eventCode": "AUTHORISATION",
                      "eventDate": "2017-01-19T16:42:03+01:00",
                      "merchantAccountCode": "MagentoMerchantTest2",
                      "merchantReference": "8313842560770001",
                      "operations": [
                        "CANCEL",
                        "CAPTURE",
                        "REFUND"
                      ],
                      "paymentMethod": "visa",
                      "pspReference": "123456789",
                      "reason": "1234:7777:12\/2012",
                      "success": "true"
                    }
                  }
                ]
              }
            `,
			func(got *webhook.Webhook, t *testing.T) {
				require.NotNil(t, got)
				assert.Equal(t, 1, len(got.GetNotificationItems()))
				ni := got.GetNotificationItems()[0]
				assert.Equal(t, webhook.EventCodeAuthorisation, ni.EventCode)
				assert.Equal(t, "true", ni.Success)
				assert.Equal(t, "123456789", ni.PspReference)
				assert.NotEmpty(t, ni.AdditionalData)
				ad := *ni.AdditionalData
				assert.Equal(t, "1234", ad["authCode"])
				assert.NotEmpty(t, ni.Amount)
				assert.NotEmpty(t, ni.EventDate)
				assert.NotEmpty(t, ni.Operations)
			},
			false,
		},
		{
			"should return capture success",
			`{
                "live": "false",
                "notificationItems": [
                  {
                    "NotificationRequestItem": {
                      "additionalData": {
                        "hmacSignature": "qvS6I3Gdi1jx+jSh7IopAgcHtMoxvXlNL7DYQ+j1hd0="
                      },
                      "amount": {
                        "currency": "USD",
                        "value": 23623
                      },
                      "eventCode": "CAPTURE",
                      "eventDate": "2017-01-25T18:08:19+01:00",
                      "merchantAccountCode": "MagentoMerchantTest2",
                      "merchantReference": "00000001",
                      "originalReference": "ORIGINAL_PSP",
                      "paymentMethod": "visa",
                      "pspReference": "PSP_REFERENCE",
                      "reason": "",
                      "success": "true"
                    }
                  }
                ]
              }
            `,
			func(got *webhook.Webhook, t *testing.T) {
				require.NotNil(t, got)
				assert.Equal(t, 1, len(got.GetNotificationItems()))
				ni := got.GetNotificationItems()[0]
				assert.Equal(t, webhook.EventCodeCapture, ni.EventCode)
				assert.Equal(t, "true", ni.Success)
				assert.Equal(t, "PSP_REFERENCE", ni.PspReference)
				assert.Equal(t, "ORIGINAL_PSP", ni.OriginalReference)
			},
			false,
		},
		{
			"should return refund fail",
			`{
                "live": "false",
                "notificationItems": [
                  {
                    "NotificationRequestItem": {
                      "additionalData": {
                        "hmacSignature": "HZXziBYopfDIzDhk49iC\/\/yCfxmy\/z0xWuvvTxFNUSA="
                      },
                      "amount": {
                        "currency": "EUR",
                        "value": 1500
                      },
                      "eventCode": "REFUND",
                      "eventDate": "2017-02-17T11:04:07+01:00",
                      "merchantAccountCode": "MagentoMerchantTest2",
                      "merchantReference": "payment-2017-1-17-11-refund",
                      "originalReference": "ORIGINAL_PSP",
                      "paymentMethod": "visa",
                      "pspReference": "PSP_REFERENCE",
                      "reason": "Insufficient balance on payment",
                      "success": "false"
                    }
                  }
                ]
              }
            `,
			func(got *webhook.Webhook, t *testing.T) {
				require.NotNil(t, got)
				assert.Equal(t, 1, len(got.GetNotificationItems()))
				ni := got.GetNotificationItems()[0]
				assert.Equal(t, webhook.EventCodeRefund, ni.EventCode)
				assert.Equal(t, "false", ni.Success)
				assert.Equal(t, "PSP_REFERENCE", ni.PspReference)
				assert.Equal(t, "ORIGINAL_PSP", ni.OriginalReference)
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := webhook.HandleRequest(tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("webhook.HandleRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Nil(t, err)
			tt.want(got, t)
		})
	}
}
