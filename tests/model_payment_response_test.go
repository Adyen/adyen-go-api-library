package tests

import (
	"encoding/json"
	"testing"

	"github.com/adyen/adyen-go-api-library/v8/src/checkout"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPaymentResponse_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		req     checkout.PaymentResponse
		json    string
		wantErr bool
		wantFn  func(pm checkout.PaymentResponse, t *testing.T)
	}{
		{
			"unmarshalls an empty payment response correctly",
			checkout.PaymentResponse{},
			"{}",
			false,
			func(got checkout.PaymentResponse, t *testing.T) {
				require.NotNil(t, got)
				require.Nil(t, got.Action)
			},
		},
		{
			"unmarshalls a payment response without an action correctly",
			checkout.PaymentResponse{},
			`{
                "resultCode":"Error",
                "PspReference":"12345"
            }`,
			false,
			func(got checkout.PaymentResponse, t *testing.T) {
				require.NotNil(t, got)
				require.Nil(t, got.Action)
				assert.Equal(t, "Error", got.GetResultCode())
				assert.Equal(t, "12345", got.GetPspReference())
			},
		},
		{
			"unmarshalls a payment response with redirect action correctly",
			checkout.PaymentResponse{},
			`{
                "action":{
                    "method":"GET",
                    "paymentMethodType":"ideal",
                    "type":"redirect",
                    "url":"https://checkoutshopper-test.adyen.com/checkoutshopper/checkoutPaymentRedirect?redirectData=X6XtfGC3..."
                },
                "resultCode":"RedirectShopper"
            }`,
			false,
			func(got checkout.PaymentResponse, t *testing.T) {
				require.NotNil(t, got)
				require.NotNil(t, got.Action)
				require.NotNil(t, got.ResultCode)
				assert.Equal(t, "redirect", got.GetAction().CheckoutRedirectAction.GetType())
				assert.Equal(t, "ideal", got.GetAction().CheckoutRedirectAction.GetPaymentMethodType())
				assert.Equal(t, "https://checkoutshopper-test.adyen.com/checkoutshopper/checkoutPaymentRedirect?redirectData=X6XtfGC3...", got.GetAction().CheckoutRedirectAction.GetUrl())
			},
		},
		{
			"unmarshalls a payment response with 3ds2 action correctly",
			checkout.PaymentResponse{},
			`{
                "resultCode": "IdentifyShopper",
                "action": {
                  "paymentData": "Ab02b4c0!BQABAgCCFm6bRbO...",
                  "paymentMethodType": "scheme",
                  "authorisationToken": "BQABAQCOY3Jh1O4zAJC7AEESc...",
                  "subtype": "fingerprint",
                  "token": "eyJ0aHJlZURTTWVzc2FnZVZ...",
                  "type": "threeDS2"
                }
            }`,
			false,
			func(got checkout.PaymentResponse, t *testing.T) {
				require.NotNil(t, got)
				require.NotNil(t, got.Action)
				require.NotNil(t, got.ResultCode)
				assert.Equal(t, "threeDS2", got.GetAction().CheckoutThreeDS2Action.GetType())
				assert.Equal(t, "fingerprint", got.GetAction().CheckoutThreeDS2Action.GetSubtype())
				assert.Equal(t, "eyJ0aHJlZURTTWVzc2FnZVZ...", got.GetAction().CheckoutThreeDS2Action.GetToken())

				jsonString, err := json.Marshal(got)
				assert.Nil(t, err)
				assert.Equal(t, `{"action":{"authorisationToken":"BQABAQCOY3Jh1O4zAJC7AEESc...","paymentData":"Ab02b4c0!BQABAgCCFm6bRbO...","paymentMethodType":"scheme","subtype":"fingerprint","token":"eyJ0aHJlZURTTWVzc2FnZVZ...","type":"threeDS2"},"resultCode":"IdentifyShopper"}`, string(jsonString))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pm := checkout.PaymentResponse{}
			err := json.Unmarshal([]byte(tt.json), &pm)
			if (err != nil) != tt.wantErr {
				t.Errorf("PaymentResponse.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Nil(t, err)
			tt.wantFn(pm, t)
		})
	}

	t.Run("Cannot unmarshal type missing from the specification", func(t *testing.T) {
		pm := checkout.PaymentResponse{}
		inputJson := `{
            "resultCode": "IdentifyShopper",
            "action": {
              "paymentMethodType": "scheme",
              "token": "eyJ0aHJlZURTTWVzc2FnZVZ...",
              "type": "unknown"
            }
        }`
		err := json.Unmarshal([]byte(inputJson), &pm)
		require.ErrorContains(t, err, "data failed to match schemas in oneOf(PaymentResponseAction)")
	})
}
