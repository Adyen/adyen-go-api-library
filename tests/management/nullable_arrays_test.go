package management

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adyen/adyen-go-api-library/v21/src/adyen"
	"github.com/adyen/adyen-go-api-library/v21/src/common"
	"github.com/adyen/adyen-go-api-library/v21/src/management"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateTerminalSettingsSendsNullGratuities(t *testing.T) {
	var body string
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw, _ := io.ReadAll(r.Body)
		body = string(raw)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"gratuities":null}`))
	}))
	defer mockServer.Close()

	client := adyen.NewClient(&common.Config{ApiKey: "YOUR_ADYEN_API_KEY", Environment: common.TestEnv})
	service := client.Management()
	service.TerminalSettingsMerchantLevelApi.BasePath = func() string { return mockServer.URL }

	settings := management.TerminalSettings{}
	settings.SetGratuitiesNil()
	req := service.TerminalSettingsMerchantLevelApi.UpdateTerminalSettingsInput("MERCHANT").TerminalSettings(settings)
	res, httpRes, err := service.TerminalSettingsMerchantLevelApi.UpdateTerminalSettings(context.Background(), req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, httpRes.StatusCode)
	assert.JSONEq(t, `{"gratuities":null}`, body)
	assert.True(t, res.HasGratuities())
	assert.Nil(t, res.GetGratuities())
}

func TestTerminalSettingsGratuitiesAbsentEmptyAndValue(t *testing.T) {
	out, err := json.Marshal(management.TerminalSettings{})
	require.NoError(t, err)
	assert.JSONEq(t, `{}`, string(out))

	settings := management.TerminalSettings{}
	settings.SetGratuities([]management.Gratuity{})
	out, err = json.Marshal(settings)
	require.NoError(t, err)
	assert.JSONEq(t, `{"gratuities":[]}`, string(out))

	currency := "EUR"
	settings.SetGratuities([]management.Gratuity{{Currency: &currency}})
	out, err = json.Marshal(settings)
	require.NoError(t, err)
	assert.JSONEq(t, `{"gratuities":[{"currency":"EUR"}]}`, string(out))

	var decoded management.TerminalSettings
	require.NoError(t, json.Unmarshal(out, &decoded))
	assert.Equal(t, "EUR", decoded.GetGratuities()[0].GetCurrency())
}
