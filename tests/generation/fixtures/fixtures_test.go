package fixtures

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// These tests pin how the generator and templates render each schema construct in
// GenerationFixtures.json. Regenerate the models with `./gradlew :go:fixtures` in
// adyen-sdk-automation, then run `make fmt`.

func roundTrip(t *testing.T, input string, target interface{}) string {
	t.Helper()
	require.NoError(t, json.Unmarshal([]byte(input), target))
	out, err := json.Marshal(target)
	require.NoError(t, err)
	return string(out)
}

func TestNullableFieldsAbsentNullAndValue(t *testing.T) {
	absent := roundTrip(t, `{}`, &NullableFields{})
	assert.JSONEq(t, `{}`, absent)

	allNull := `{"legacyBool":null,"legacyString":null,"legacyRef":null,
		"unionBool":null,"unionString":null,"unionDate":null,"unionInt32":null,
		"unionInt64":null,"unionEnum":null,"unionRef":null}`
	var nulls NullableFields
	assert.JSONEq(t, allNull, roundTrip(t, allNull, &nulls))
	assert.True(t, nulls.HasLegacyBool())
	assert.True(t, nulls.LegacyRef.IsSet())
	assert.Nil(t, nulls.LegacyRef.Get())
	assert.True(t, nulls.UnionRef.IsSet())
	assert.Nil(t, nulls.UnionRef.Get())

	values := `{"legacyBool":true,"legacyString":"a","legacyArray":["x"],
		"legacyRef":{"currency":"EUR","value":1},"unionBool":false,"unionString":"b",
		"unionDate":"2024-01-31","unionInt32":7,"unionInt64":9007199254740993,
		"unionArray":[{"currency":"USD","value":2}],"unionEnum":"active",
		"unionRef":{"currency":"GBP","value":3}}`
	var full NullableFields
	assert.JSONEq(t, values, roundTrip(t, values, &full))
	assert.Equal(t, int64(9007199254740993), full.GetUnionInt64())
	assert.Equal(t, "GBP", full.GetUnionRef().Currency)
	assert.Equal(t, "EUR", full.GetLegacyRef().Currency)
}

func TestNullableFieldsSetters(t *testing.T) {
	m := NewNullableFields()
	m.SetUnionBool(false)
	m.SetLegacyStringNil()
	m.SetUnionRef(Amount{Currency: "EUR", Value: 5})
	out, err := json.Marshal(m)
	require.NoError(t, err)
	assert.JSONEq(t, `{"unionBool":false,"legacyString":null,"unionRef":{"currency":"EUR","value":5}}`, string(out))

	m.UnsetLegacyString()
	m.SetUnionRefNil()
	out, err = json.Marshal(m)
	require.NoError(t, err)
	assert.JSONEq(t, `{"unionBool":false,"unionRef":null}`, string(out))
}

func TestPlainFieldsTypes(t *testing.T) {
	// Optional typed maps stay pointers; free-form objects and required maps do not.
	m := PlainFields{}
	var _ *map[string]string = m.Metadata
	var _ *map[string]Amount = m.Amounts
	var _ map[string]string = m.RequiredMap
	var _ map[string]interface{} = m.FreeForm
	var _ *Amount = m.Amount
	var _ []string = m.Tags
}

func TestPlainFieldsRoundTrip(t *testing.T) {
	input := `{"name":"n","enabled":true,"count":3,"amount":{"currency":"EUR","value":1},
		"metadata":{"k":"v"},"amounts":{"a":{"currency":"USD","value":2}},
		"requiredMap":{},"tags":["t"],"freeForm":{"nested":{"x":1}}}`
	var m PlainFields
	assert.JSONEq(t, input, roundTrip(t, input, &m))
	assert.Equal(t, map[string]string{"k": "v"}, m.GetMetadata())
	assert.Equal(t, int64(2), m.GetAmounts()["a"].Value)

	empty := NewPlainFields(map[string]string{})
	out, err := json.Marshal(empty)
	require.NoError(t, err)
	assert.JSONEq(t, `{"requiredMap":{},"enabled":false}`, string(out))
}

func TestDefaults(t *testing.T) {
	assert.False(t, *NewPlainFieldsWithDefaults().Enabled)
	assert.False(t, *NewPaymentRequestWithDefaults().ThreeDSAuthenticationOnly)
	assert.Equal(t, "scheme", NewCardDetails().GetType())

	// Unmarshalling does not apply defaults: an absent field stays unset.
	var req PaymentRequest
	require.NoError(t, json.Unmarshal([]byte(`{"amount":{"currency":"EUR","value":1},"paymentMethod":{"type":"applepay"}}`), &req))
	assert.False(t, req.HasThreeDSAuthenticationOnly())
}

func TestAllOfIncludesInheritedProperties(t *testing.T) {
	payTo := NewPayToPaymentMethod("payTo")
	payTo.SetShopperReference("ref")
	out, err := json.Marshal(payTo)
	require.NoError(t, err)
	assert.JSONEq(t, `{"type":"payTo","shopperReference":"ref"}`, string(out))

	var upi UPIPaymentMethod
	assert.JSONEq(t, `{"type":"upi_collect","virtualPaymentAddress":"v@upi"}`,
		roundTrip(t, `{"type":"upi_collect","virtualPaymentAddress":"v@upi"}`, &upi))
	assert.Equal(t, "upi_collect", upi.GetType())
}

func TestDiscriminatedOneOf(t *testing.T) {
	var m DiscriminatedMethod
	assert.JSONEq(t, `{"type":"upi_collect","virtualPaymentAddress":"v@upi"}`,
		roundTrip(t, `{"type":"upi_collect","virtualPaymentAddress":"v@upi"}`, &m))
	require.NotNil(t, m.UPIPaymentMethod)
	assert.Nil(t, m.PayToPaymentMethod)

	m = DiscriminatedMethod{}
	require.NoError(t, json.Unmarshal([]byte(`{"type":"payTo","shopperReference":"r"}`), &m))
	require.NotNil(t, m.PayToPaymentMethod)
	assert.Equal(t, "r", m.PayToPaymentMethod.GetShopperReference())

	require.NoError(t, json.Unmarshal([]byte(`{"type":"upi_collect"}`), &m))
	assert.Nil(t, m.PayToPaymentMethod, "previous variant must be cleared")
	require.NotNil(t, m.UPIPaymentMethod)

	assert.Error(t, json.Unmarshal([]byte(`{"type":"unknown"}`), &m))
	assert.Nil(t, m.GetActualInstance())

	wrapped := PayToPaymentMethodAsDiscriminatedMethod(NewPayToPaymentMethod("payTo"))
	out, err := json.Marshal(wrapped)
	require.NoError(t, err)
	assert.JSONEq(t, `{"type":"payTo"}`, string(out))
}

func TestOneOfWithoutDiscriminator(t *testing.T) {
	var card PaymentMethod
	assert.JSONEq(t, `{"type":"scheme","number":"4111"}`, roundTrip(t, `{"type":"scheme","number":"4111"}`, &card))
	require.NotNil(t, card.CardDetails)
	assert.Nil(t, card.WalletDetails)

	var wallet PaymentMethod
	assert.JSONEq(t, `{"type":"applepay","token":"tok"}`, roundTrip(t, `{"type":"applepay","token":"tok"}`, &wallet))
	require.NotNil(t, wallet.WalletDetails)
	assert.Nil(t, wallet.CardDetails)

	var bad PaymentMethod
	assert.Error(t, json.Unmarshal([]byte(`{"type":"unknown"}`), &bad))
}

func TestOneOfMemberWithoutTypeEnum(t *testing.T) {
	var amount AmountOrWallet
	assert.JSONEq(t, `{"currency":"EUR","value":1}`, roundTrip(t, `{"currency":"EUR","value":1}`, &amount))
	require.NotNil(t, amount.Amount)
	assert.Nil(t, amount.WalletDetails)

	// Decoding ignores unknown fields, so a member without a "type" enum matches any object:
	// a wallet payload matches both members and is rejected as ambiguous.
	var wallet AmountOrWallet
	assert.Error(t, json.Unmarshal([]byte(`{"type":"applepay","token":"tok"}`), &wallet))
	assert.Nil(t, wallet.GetActualInstance())

	out, err := json.Marshal(WalletDetailsAsAmountOrWallet(NewWalletDetails("applepay")))
	require.NoError(t, err)
	assert.JSONEq(t, `{"type":"applepay"}`, string(out))
}

func TestAnyOf(t *testing.T) {
	var amount AnyOfValue
	assert.JSONEq(t, `{"currency":"EUR","value":1}`, roundTrip(t, `{"currency":"EUR","value":1}`, &amount))
	require.NotNil(t, amount.Amount)

	req := PaymentRequest{
		Amount:        Amount{Currency: "EUR", Value: 10},
		PaymentMethod: CardDetailsAsPaymentMethod(NewCardDetails()),
	}
	out, err := json.Marshal(req)
	require.NoError(t, err)
	assert.JSONEq(t, `{"amount":{"currency":"EUR","value":10},"paymentMethod":{"type":"scheme"}}`, string(out))
}

func TestTopLevelEnumsArePrefixed(t *testing.T) {
	assert.Equal(t, CALocalBankAccountType("checking"), CALOCALBANKACCOUNTTYPE_CHECKING)
	assert.Equal(t, USLocalBankAccountType("savings"), USLOCALBANKACCOUNTTYPE_SAVINGS)

	var ca CALocalAccountIdentification
	assert.JSONEq(t, `{"accountType":"savings"}`, roundTrip(t, `{"accountType":"savings"}`, &ca))
	assert.Equal(t, CALOCALBANKACCOUNTTYPE_SAVINGS, ca.GetAccountType())

	var us USLocalAccountIdentification
	assert.Error(t, json.Unmarshal([]byte(`{"accountType":"other"}`), &us))
}
