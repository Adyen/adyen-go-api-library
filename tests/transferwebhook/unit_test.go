package transferwebhook

import (
	"encoding/json"
	"testing"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
	"github.com/adyen/adyen-go-api-library/v21/src/transferwebhook"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestIssuedCardNetworkVariant verifies that all supported network variants
// can be set, read, serialized to JSON, and deserialized back into an IssuedCard.
func TestIssuedCardNetworkVariant(t *testing.T) {
	for _, networkVariant := range []string{"maestro_us", "mastercard", "visa"} {
		t.Run(networkVariant, func(t *testing.T) {
			// Create an IssuedCard and set the new NetworkVariant field.
			issuedCard := transferwebhook.NewIssuedCard()
			issuedCard.SetNetworkVariant(networkVariant)

			// Verify the generated getter methods return the value that was set.
			assert.Equal(t, networkVariant, issuedCard.GetNetworkVariant())
			value, ok := issuedCard.GetNetworkVariantOk()
			require.True(t, ok)
			require.NotNil(t, value)
			assert.Equal(t, networkVariant, *value)
			assert.True(t, issuedCard.HasNetworkVariant())

			// Verify the new field is included when the model is converted to JSON.
			data, err := json.Marshal(issuedCard)
			require.NoError(t, err)

			var payload map[string]interface{}
			require.NoError(t, json.Unmarshal(data, &payload))
			assert.Equal(t, networkVariant, payload["networkVariant"])

			// Verify the field can be read back when JSON is unmarshaled.
			var decoded transferwebhook.IssuedCard
			require.NoError(t, json.Unmarshal(data, &decoded))
			assert.Equal(t, networkVariant, decoded.GetNetworkVariant())
		})
	}
}

// TestNewTransferWebhookEnumValuesSerialize verifies that the newly supported
// transfer webhook enum values are included correctly when models are serialized to JSON.
func TestNewTransferWebhookEnumValuesSerialize(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
		field string
		want  string
	}{
		{
			name:  "modification reversal received status",
			value: transferwebhook.Modification{Status: common.PtrString("reversalReceived")},
			field: "status",
			want:  "reversalReceived",
		},
		{
			name:  "transfer data reversal received status",
			value: transferwebhook.TransferData{Status: "reversalReceived"},
			field: "status",
			want:  "reversalReceived",
		},
		{
			name:  "transfer data fx sell type",
			value: transferwebhook.TransferData{Type: common.PtrString("fxSell")},
			field: "type",
			want:  "fxSell",
		},
		{
			name:  "transfer data fx buy type",
			value: transferwebhook.TransferData{Type: common.PtrString("fxBuy")},
			field: "type",
			want:  "fxBuy",
		},
		{
			name:  "transfer event reversal received status",
			value: transferwebhook.TransferEvent{Status: common.PtrString("reversalReceived")},
			field: "status",
			want:  "reversalReceived",
		},
		{
			name:  "network reason us ach correction namespace",
			value: transferwebhook.NetworkReason{Namespace: common.PtrString("usAchCorrectionReasonCode")},
			field: "namespace",
			want:  "usAchCorrectionReasonCode",
		},
	}

	// Each test case checks one newly supported enum value.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.value)
			require.NoError(t, err)

			var payload map[string]interface{}
			require.NoError(t, json.Unmarshal(data, &payload))

			assert.Equal(t, tt.want, payload[tt.field])
		})
	}
}
