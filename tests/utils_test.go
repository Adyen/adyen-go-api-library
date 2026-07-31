package tests

import (
	"testing"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewStrictDecoder(t *testing.T) {
	type variant struct {
		Type string `json:"type"`
	}

	t.Run("decodes a payload with known fields only", func(t *testing.T) {
		var got variant
		err := common.NewStrictDecoder([]byte(`{"type":"scheme"}`)).Decode(&got)

		require.NoError(t, err)
		assert.Equal(t, "scheme", got.Type)
	})

	t.Run("rejects a payload with unknown fields", func(t *testing.T) {
		var got variant
		err := common.NewStrictDecoder([]byte(`{"type":"scheme","brand":"visa"}`)).Decode(&got)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "unknown field")
	})

	t.Run("accepts any key when decoding into a map", func(t *testing.T) {
		var got map[string]interface{}
		err := common.NewStrictDecoder([]byte(`{"type":"scheme","brand":"visa"}`)).Decode(&got)

		require.NoError(t, err)
		assert.Equal(t, "scheme", got["type"])
	})
}
