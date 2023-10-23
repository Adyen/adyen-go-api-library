package tests

import (
	"encoding/json"
	"testing"

	"github.com/adyen/adyen-go-api-library/v8/src/checkout"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPaymentRequest_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		req     checkout.PaymentRequest
		json    string
		wantErr bool
		wantFn  func(pm checkout.PaymentRequest, t *testing.T)
	}{
		{
			"unmarshalls an empty payment request correctly",
			checkout.PaymentRequest{},
			"{}",
			false,
			func(got checkout.PaymentRequest, t *testing.T) {
				require.NotNil(t, got)
				require.Nil(t, got.PaymentMethod.GetActualInstance())
			},
		},
		{
			"unmarshalls a payment request without payment method correctly",
			checkout.PaymentRequest{},
			`{
                "riskData":{"clientData":"eyJ2ZXJzaW9uIjoiMS4w"},
                "browserInfo":{"acceptHeader":"*/*","colorDepth":24,"language":"en-US","javaEnabled":false,"screenHeight":1080,"screenWidth":1920,"userAgent":"Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:85.0) Gecko/20100101 Firefox/85.0","timeZoneOffset":-60}
            }`,
			false,
			func(got checkout.PaymentRequest, t *testing.T) {
				require.NotNil(t, got)
				require.Nil(t, got.PaymentMethod.GetActualInstance())
				require.NotNil(t, got.RiskData)
				require.NotNil(t, got.BrowserInfo)
			},
		},
		{
			"unmarshalls a payment request with ideal correctly",
			checkout.PaymentRequest{},
			`{
                "riskData":{"clientData":"eyJ2ZXJzaW9uIjoiMS4"},
                "paymentMethod":{"type":"ideal","issuer":"1121"}
            }`,
			false,
			func(got checkout.PaymentRequest, t *testing.T) {
				require.NotNil(t, got)
				require.NotNil(t, got.PaymentMethod)
				require.NotNil(t, got.RiskData)
				assert.Equal(t, "ideal", got.GetPaymentMethod().IdealDetails.GetType())
				assert.Equal(t, "1121", got.GetPaymentMethod().IdealDetails.GetIssuer())
			},
		},
		{
			"unmarshalls a payment request with scheme correctly",
			checkout.PaymentRequest{},
			`{
                "riskData":{"clientData":"eyJ2ZXJzaW9uIjoiMS4w"},
                "paymentMethod":{"type":"scheme","holderName":"d","encryptedCardNumber":"adyenjs_0_1_25$J8/5xp5l6DjYVPokO6FwAQj","encryptedExpiryMonth":"adyenjs_0_1_25$bLCWe/ZHR37Okz0d28bzrDBYXw","encryptedExpiryYear":"adyenjs_0_1_25$nqasksbOSfn0grzrmna2vpWkQMhOHT6Cd","encryptedSecurityCode":"adyenjs_0_1_25$TbomjrfaGwHFfxpPuf","brand":"amex"},
                "browserInfo":{"acceptHeader":"*/*","colorDepth":24,"language":"en-US","javaEnabled":false,"screenHeight":1080,"screenWidth":1920,"userAgent":"Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:85.0) Gecko/20100101 Firefox/85.0","timeZoneOffset":-60}
            }`,
			false,
			func(got checkout.PaymentRequest, t *testing.T) {
				require.NotNil(t, got)
				require.NotNil(t, got.PaymentMethod)
				require.NotNil(t, got.RiskData)
				require.NotNil(t, got.BrowserInfo)
				assert.Equal(t, "scheme", got.GetPaymentMethod().CardDetails.GetType())
				assert.Equal(t, "adyenjs_0_1_25$J8/5xp5l6DjYVPokO6FwAQj", got.GetPaymentMethod().CardDetails.GetEncryptedCardNumber())

				jsonString, err := json.Marshal(got)
				assert.Nil(t, err)
				assert.Equal(t, `{"amount":{"currency":"","value":0},"browserInfo":{"acceptHeader":"*/*","colorDepth":24,"javaEnabled":false,"language":"en-US","screenHeight":1080,"screenWidth":1920,"timeZoneOffset":-60,"userAgent":"Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:85.0) Gecko/20100101 Firefox/85.0"},"merchantAccount":"","paymentMethod":{"brand":"amex","encryptedCardNumber":"adyenjs_0_1_25$J8/5xp5l6DjYVPokO6FwAQj","encryptedExpiryMonth":"adyenjs_0_1_25$bLCWe/ZHR37Okz0d28bzrDBYXw","encryptedExpiryYear":"adyenjs_0_1_25$nqasksbOSfn0grzrmna2vpWkQMhOHT6Cd","encryptedSecurityCode":"adyenjs_0_1_25$TbomjrfaGwHFfxpPuf","holderName":"d","type":"scheme"},"reference":"","returnUrl":"","riskData":{"clientData":"eyJ2ZXJzaW9uIjoiMS4w"}}`, string(jsonString))
			},
		},
		{
			"unmarshalls a payment request with ach correctly",
			checkout.PaymentRequest{},
			`{
                "riskData":{"clientData":"eyJ2ZX"},
                "paymentMethod":{"type":"ach","encryptedBankAccountNumber":"adyenjs_0_1_25","encryptedBankLocationId":"adyenjs_0_1_25","ownerName":"test"},
                "billingAddress":{"street":"test","houseNumberOrName":"2","postalCode":"123456","city":"porto rico","stateOrProvince":"N/A","country":"PR"}
            }`,
			false,
			func(got checkout.PaymentRequest, t *testing.T) {
				require.NotNil(t, got)
				require.NotNil(t, got.PaymentMethod)
				require.NotNil(t, got.RiskData)
				require.NotNil(t, got.BillingAddress)
				assert.Equal(t, "test", got.BillingAddress.Street)
				assert.Equal(t, "ach", got.GetPaymentMethod().AchDetails.GetType())
				assert.Equal(t, "adyenjs_0_1_25", got.GetPaymentMethod().AchDetails.GetEncryptedBankAccountNumber())
				assert.Equal(t, "test", got.GetPaymentMethod().AchDetails.GetOwnerName())
			},
		},
		{
			"unmarshalls a payment request with klarna correctly",
			checkout.PaymentRequest{},
			`{
                "riskData":{"clientData":"eyJ2ZX"},
                "paymentMethod":{"type":"klarna","billingAddress":"adyenjs_0_1_25"},
                "billingAddress":{"street":"test","houseNumberOrName":"2","postalCode":"123456","city":"porto rico","stateOrProvince":"N/A","country":"PR"}
            }`,
			false,
			func(got checkout.PaymentRequest, t *testing.T) {
				require.NotNil(t, got)
				require.NotNil(t, got.PaymentMethod)
				require.NotNil(t, got.RiskData)
				require.NotNil(t, got.BillingAddress)
				assert.Equal(t, "test", got.BillingAddress.Street)
				assert.Equal(t, "klarna", got.GetPaymentMethod().KlarnaDetails.GetType())
				assert.Equal(t, "adyenjs_0_1_25", got.GetPaymentMethod().KlarnaDetails.GetBillingAddress())
			},
		},
		{
			"unmarshalls a payment request with klarna_paynow correctly",
			checkout.PaymentRequest{},
			`{
                "riskData":{"clientData":"eyJ2ZX"},
                "paymentMethod":{"type":"klarna_paynow","billingAddress":"adyenjs_0_1_25"},
                "billingAddress":{"street":"test","houseNumberOrName":"2","postalCode":"123456","city":"porto rico","stateOrProvince":"N/A","country":"PR"}
            }`,
			false,
			func(got checkout.PaymentRequest, t *testing.T) {
				require.NotNil(t, got)
				require.NotNil(t, got.PaymentMethod)
				require.NotNil(t, got.RiskData)
				require.NotNil(t, got.BillingAddress)
				assert.Equal(t, "test", got.BillingAddress.Street)
				assert.Equal(t, "klarna_paynow", got.GetPaymentMethod().KlarnaDetails.GetType())
				assert.Equal(t, "adyenjs_0_1_25", got.GetPaymentMethod().KlarnaDetails.GetBillingAddress())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pm := checkout.PaymentRequest{}
			err := json.Unmarshal([]byte(tt.json), &pm)
			if (err != nil) != tt.wantErr {
				t.Errorf("PaymentRequest.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Nil(t, err)
			tt.wantFn(pm, t)
		})
	}

	t.Run("Cannot unmarshal type missing from the specification", func(t *testing.T) {
		pm := checkout.PaymentRequest{}
		inputJson := `{
            "riskData":{"clientData":"eyJ2ZXJzaW9uIjoiMS4w"},
            "paymentMethod":{"type":"mypay","holderName":"d","number":"1234", "brand":"amex"},
            "browserInfo":{"acceptHeader":"*/*","colorDepth":24,"language":"en-US","javaEnabled":false,"screenHeight":1080,"screenWidth":1920,"userAgent":"Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:85.0) Gecko/20100101 Firefox/85.0","timeZoneOffset":-60}
        }`
		err := json.Unmarshal([]byte(inputJson), &pm)
		require.ErrorContains(t, err, "data failed to match schemas in oneOf(CheckoutPaymentMethod)")
	})
}
