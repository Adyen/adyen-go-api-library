package transactionwebhook

import (
	"encoding/json"
	"testing"

	"github.com/adyen/adyen-go-api-library/v21/src/transactionwebhook"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestIssuedCardNetworkVariant verifies that NetworkVariant is available
// through the generated accessors and survives JSON serialization.
func TestIssuedCardNetworkVariant(t *testing.T) {
	for _, networkVariant := range []string{"maestro_us", "mastercard", "visa"} {
		t.Run(networkVariant, func(t *testing.T) {
			issuedCard := transactionwebhook.NewIssuedCard()
			issuedCard.SetNetworkVariant(networkVariant)

			assert.Equal(t, networkVariant, issuedCard.GetNetworkVariant())
			value, ok := issuedCard.GetNetworkVariantOk()
			require.True(t, ok)
			require.NotNil(t, value)
			assert.Equal(t, networkVariant, *value)
			assert.True(t, issuedCard.HasNetworkVariant())

			data, err := json.Marshal(issuedCard)
			require.NoError(t, err)

			var payload map[string]interface{}
			require.NoError(t, json.Unmarshal(data, &payload))
			assert.Equal(t, networkVariant, payload["networkVariant"])

			var decoded transactionwebhook.IssuedCard
			require.NoError(t, json.Unmarshal(data, &decoded))
			assert.Equal(t, networkVariant, decoded.GetNetworkVariant())
		})
	}
}
