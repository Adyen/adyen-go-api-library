package compat

import (
	"encoding/json"
	"testing"

	"github.com/adyen/adyen-go-api-library/v21/src/checkout"
	"github.com/adyen/adyen-go-api-library/v21/src/management"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Nullable fields have three wire states (absent, null, value) that must survive a round trip.
func TestNullableFieldStates(t *testing.T) {
	cases := []struct {
		name      string
		in        string
		has       bool
		value     *string
		roundTrip string
	}{
		{"absent", `{}`, false, nil, `{}`},
		{"explicit null", `{"paymentInstrument":null}`, true, nil, `{"paymentInstrument":null}`},
		{"value", `{"paymentInstrument":"Cash"}`, true, ptr("Cash"), `{"paymentInstrument":"Cash"}`},
		{"empty string", `{"paymentInstrument":""}`, true, ptr(""), `{"paymentInstrument":""}`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var p management.PayAtTable
			require.NoError(t, json.Unmarshal([]byte(tc.in), &p))
			assert.Equal(t, tc.has, p.HasPaymentInstrument())
			got, _ := p.GetPaymentInstrumentOk()
			assert.Equal(t, tc.value, got)
			out, err := json.Marshal(p)
			require.NoError(t, err)
			assert.JSONEq(t, tc.roundTrip, string(out))
		})
	}

	t.Run("setters", func(t *testing.T) {
		var p management.PayAtTable
		p.SetPaymentInstrumentNil()
		out, _ := json.Marshal(p)
		assert.JSONEq(t, `{"paymentInstrument":null}`, string(out))
		p.UnsetPaymentInstrument()
		out, _ = json.Marshal(p)
		assert.JSONEq(t, `{}`, string(out))
	})
}

// Zero values of optional fields are sent when explicitly set, and omitted otherwise.
func TestOptionalZeroValues(t *testing.T) {
	var p management.PayAtTable
	out, _ := json.Marshal(p)
	assert.JSONEq(t, `{}`, string(out))

	p.SetEnablePayAtTable(false)
	out, _ = json.Marshal(p)
	assert.JSONEq(t, `{"enablePayAtTable":false}`, string(out))

	var back management.PayAtTable
	require.NoError(t, json.Unmarshal(out, &back))
	assert.True(t, back.HasEnablePayAtTable())
	assert.False(t, back.GetEnablePayAtTable())
}

func TestOneOfPaymentMethod(t *testing.T) {
	card := `{"type":"scheme","encryptedCardNumber":"test_4111111111111111"}`
	ideal := `{"type":"ideal","issuer":"1121"}`

	t.Run("discriminates by type", func(t *testing.T) {
		var pm checkout.CheckoutPaymentMethod
		require.NoError(t, json.Unmarshal([]byte(card), &pm))
		require.NotNil(t, pm.CardDetails)
		assert.IsType(t, &checkout.CardDetails{}, pm.GetActualInstance())
		assert.Equal(t, "test_4111111111111111", pm.CardDetails.GetEncryptedCardNumber())
		out, err := json.Marshal(pm)
		require.NoError(t, err)
		assert.JSONEq(t, card, string(out))
	})

	t.Run("repeated unmarshal into the same destination", func(t *testing.T) {
		var pm checkout.CheckoutPaymentMethod
		require.NoError(t, json.Unmarshal([]byte(card), &pm))
		require.NoError(t, json.Unmarshal([]byte(ideal), &pm))
		assert.Nil(t, pm.CardDetails, "previous variant must be cleared")
		require.NotNil(t, pm.IdealDetails)
		assert.IsType(t, &checkout.IdealDetails{}, pm.GetActualInstance())
	})

	t.Run("unknown variant is rejected", func(t *testing.T) {
		var pm checkout.CheckoutPaymentMethod
		err := json.Unmarshal([]byte(`{"type":"doesNotExist"}`), &pm)
		require.Error(t, err)
		assert.Nil(t, pm.GetActualInstance())
	})

	t.Run("embedded in a request", func(t *testing.T) {
		var req checkout.PaymentRequest
		require.NoError(t, json.Unmarshal([]byte(`{"amount":{"currency":"EUR","value":0},"merchantAccount":"M","reference":"R","returnUrl":"https://example.org","paymentMethod":`+ideal+`}`), &req))
		require.NotNil(t, req.PaymentMethod.IdealDetails)
		assert.Equal(t, int64(0), req.Amount.Value, "zero amount is a valid value")
	})
}

func TestResponseToleratesUnknownFields(t *testing.T) {
	var res checkout.PaymentResponse
	require.NoError(t, json.Unmarshal([]byte(`{"pspReference":"X","resultCode":"Authorised","brandNewField":{"nested":[1,2]}}`), &res))
	assert.Equal(t, "X", res.GetPspReference())
}

func ptr(s string) *string { return &s }
