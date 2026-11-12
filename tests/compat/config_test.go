package compat

import (
	"encoding/json"
	"testing"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Every Config field must have its own JSON key: fields sharing a key are
// silently dropped by encoding/json in both directions.
func TestConfigJSONKeys(t *testing.T) {
	cfg := common.Config{ApiKey: "key", Debug: true, Log3XXError: true, Log4XXError: true, Log5XXError: true}
	out, err := json.Marshal(cfg)
	require.NoError(t, err)

	var keys map[string]interface{}
	require.NoError(t, json.Unmarshal(out, &keys))
	for _, k := range []string{"apiKey", "debug", "log3XXError", "log4XXError", "log5XXError"} {
		assert.Equal(t, true, keys[k] != nil, "missing key %s in %s", k, out)
	}

	var back common.Config
	require.NoError(t, json.Unmarshal(out, &back))
	assert.True(t, back.Debug)
	assert.True(t, back.Log3XXError)
	assert.True(t, back.Log4XXError)
	assert.True(t, back.Log5XXError)
}
