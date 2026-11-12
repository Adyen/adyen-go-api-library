package tests

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/adyen/adyen-go-api-library/v21/src/adyen"
	"github.com/adyen/adyen-go-api-library/v21/src/common"
	"github.com/adyen/adyen-go-api-library/v21/src/documentcollector"
	"github.com/adyen/adyen-go-api-library/v21/src/foreignexchange"
	"github.com/adyen/adyen-go-api-library/v21/src/posmobile"
	"github.com/adyen/adyen-go-api-library/v21/src/sessionauthentication"
	"github.com/adyen/adyen-go-api-library/v21/src/softposconfiguration"
)

// recordedRequest is what the mock server received.
type recordedRequest struct {
	method, path, contentType, apiKey string
	body                              []byte
	request                           *http.Request
}

// newAPIServer starts a server that records the request and answers with status and response.
func newAPIServer(t *testing.T, status int, response string) (*httptest.Server, *recordedRequest) {
	t.Helper()
	rec := &recordedRequest{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec.method, rec.path = r.Method, r.URL.Path
		rec.contentType, rec.apiKey = r.Header.Get("Content-Type"), r.Header.Get("X-API-Key")
		if r.MultipartForm == nil && r.Header.Get("Content-Type") != "" && r.Method != http.MethodGet {
			if err := r.ParseMultipartForm(1 << 20); err != nil {
				rec.body, _ = io.ReadAll(r.Body)
			}
		}
		rec.request = r
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write([]byte(response))
	}))
	t.Cleanup(server.Close)
	return server, rec
}

func newTestClient() *adyen.APIClient {
	return adyen.NewClient(&common.Config{ApiKey: "test-key", Environment: common.TestEnv})
}

func TestCapitalGetGrant(t *testing.T) {
	server, rec := newAPIServer(t, http.StatusOK, `{
		"id": "GR00000000000000000000001",
		"grantAccountId": "CG00000000000000000000001",
		"grantOfferId": "GO00000000000000000000001",
		"status": {"code": "Active"},
		"balances": {"currency": "EUR", "principal": 10000, "fee": 1000, "total": 11000}
	}`)
	client := newTestClient()
	client.GetConfig().CapitalEndpoint = server.URL

	service := client.Capital().GrantsApi
	grant, httpRes, err := service.GetGrant(context.Background(), service.GetGrantInput("GR00000000000000000000001"))

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, httpRes.StatusCode)
	assert.Equal(t, http.MethodGet, rec.method)
	assert.Equal(t, "/grants/GR00000000000000000000001", rec.path)
	assert.Equal(t, "test-key", rec.apiKey)
	assert.Equal(t, "CG00000000000000000000001", grant.GrantAccountId)
	assert.Equal(t, "Active", grant.Status.Code)
	assert.Equal(t, int64(11000), grant.Balances.GetTotal())
}

func TestForeignExchangeCalculate(t *testing.T) {
	server, rec := newAPIServer(t, http.StatusOK, `{
		"exchangeCalculations": [{
			"appliedExchangeRate": 1.1,
			"exchangeSide": "sell",
			"sourceAmount": {"currency": "EUR", "value": 1000},
			"targetAmount": {"currency": "USD", "value": 1100},
			"type": "transfer"
		}]
	}`)
	client := newTestClient()
	client.GetConfig().ForeignExchangeEndpoint = server.URL

	service := client.ForeignExchange().RatesApi
	item := foreignexchange.NewCalculateRateRequestItem(foreignexchange.EXCHANGESIDE_SELL, *foreignexchange.NewAmount("EUR", 1000), "USD", foreignexchange.RATETYPE_TRANSFER)
	req := service.CalculateInput().CalculateRateRequest(*foreignexchange.NewCalculateRateRequest([]foreignexchange.CalculateRateRequestItem{*item}))
	res, _, err := service.Calculate(context.Background(), req)

	require.NoError(t, err)
	assert.Equal(t, http.MethodPost, rec.method)
	assert.Equal(t, "/rates/calculate", rec.path)
	assert.JSONEq(t, `{"exchangeCalculations":[{"exchangeSide":"sell","sourceAmount":{"currency":"EUR","value":1000},"targetCurrency":"USD","type":"transfer"}]}`, string(rec.body))
	require.Len(t, res.ExchangeCalculations, 1)
	assert.Equal(t, int64(1100), res.ExchangeCalculations[0].GetTargetAmount().Value)
}

func TestSessionAuthenticationCreateSession(t *testing.T) {
	server, rec := newAPIServer(t, http.StatusOK, `{"id": "session-id", "token": "session-token"}`)
	client := newTestClient()
	client.GetConfig().SessionAuthenticationEndpoint = server.URL

	service := client.SessionAuthentication().SessionAuthenticationApi
	policy := sessionauthentication.NewPolicy()
	policy.SetRoles([]string{"Transactions Overview Component: View"})
	req := service.CreateAuthenticationSessionInput().AuthenticationSessionRequest(
		*sessionauthentication.NewAuthenticationSessionRequest(*policy, sessionauthentication.PRODUCTTYPE_PLATFORM))
	res, _, err := service.CreateAuthenticationSession(context.Background(), req)

	require.NoError(t, err)
	assert.Equal(t, "/sessions", rec.path)
	var sent map[string]interface{}
	require.NoError(t, json.Unmarshal(rec.body, &sent))
	assert.Equal(t, "platform", sent["product"])
	assert.Equal(t, "session-token", res.GetToken())
}

func TestSoftPOSConfigurationIssueCertificates(t *testing.T) {
	server, rec := newAPIServer(t, http.StatusOK, `{"id": "cert-id", "installationId": "install-id", "sdkData": "sdk"}`)
	client := newTestClient()
	client.GetConfig().SoftPOSConfigurationEndpoint = server.URL

	service := client.SoftPOSConfiguration().SessionAuthenticationApi
	req := service.IssueCertificatesInput().CertificateLoadingRequest(*softposconfiguration.NewCertificateLoadingRequest("YOUR_MERCHANT_ACCOUNT", "setup-token"))
	res, _, err := service.IssueCertificates(context.Background(), req)

	require.NoError(t, err)
	assert.Equal(t, "/auth/certificate", rec.path)
	assert.JSONEq(t, `{"merchantAccount":"YOUR_MERCHANT_ACCOUNT","setupToken":"setup-token"}`, string(rec.body))
	assert.Equal(t, "install-id", res.GetInstallationId())
}

func TestPosMobileCreateCommunicationSession(t *testing.T) {
	server, rec := newAPIServer(t, http.StatusOK, `{"id": "session-id", "sdkData": "sdk"}`)
	client := newTestClient()
	client.GetConfig().PosMobileEndpoint = server.URL

	service := client.PosMobile()
	req := service.CreateCommunicationSessionInput().CreateSessionRequest(*posmobile.NewCreateSessionRequest("YOUR_MERCHANT_ACCOUNT", "setup-token"))
	res, _, err := service.CreateCommunicationSession(context.Background(), req)

	require.NoError(t, err)
	assert.Equal(t, "/sessions", rec.path)
	assert.Equal(t, "session-id", res.GetId())
}

func TestDocumentCollectorUploadIsMultipart(t *testing.T) {
	server, rec := newAPIServer(t, http.StatusCreated, `{
		"pspReference": "QRMJC25GDZRKDM92",
		"merchantAccount": "YourMerchantAccount",
		"context": "paCbInvoice",
		"fileName": "QRMJC25GDZRKDM92.pdf"
	}`)
	client := newTestClient()
	client.GetConfig().DocumentCollectorEndpoint = server.URL

	path := filepath.Join(t.TempDir(), "QRMJC25GDZRKDM92.pdf")
	require.NoError(t, os.WriteFile(path, []byte("%PDF-1.4 test"), 0o600))
	file, err := os.Open(path)
	require.NoError(t, err)
	defer file.Close()

	service := client.DocumentCollector().DocumentsApi
	req := service.UploadCrossBorderInvoiceInput().
		Context(documentcollector.DOCUMENTCONTEXT_PA_CB_INVOICE).
		File(file).
		MerchantAccount("YourMerchantAccount")
	res, httpRes, err := service.UploadCrossBorderInvoice(context.Background(), req)

	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode)
	assert.Equal(t, "/crossBorderInvoices", rec.path)
	assert.Regexp(t, `^multipart/form-data; boundary=`, rec.contentType)
	form := rec.request.MultipartForm
	require.NotNil(t, form)
	assert.Equal(t, []string{"paCbInvoice"}, form.Value["context"])
	assert.Equal(t, []string{"YourMerchantAccount"}, form.Value["merchantAccount"])
	require.Len(t, form.File["file"], 1)
	assert.Equal(t, "QRMJC25GDZRKDM92.pdf", form.File["file"][0].Filename)
	part, err := form.File["file"][0].Open()
	require.NoError(t, err)
	content, _ := io.ReadAll(part)
	assert.Equal(t, "%PDF-1.4 test", string(content))
	assert.Equal(t, "QRMJC25GDZRKDM92.pdf", res.GetFileName())
}

func TestSoftPOSConfigurationLiveRegions(t *testing.T) {
	client := adyen.NewClient(&common.Config{ApiKey: "key", Environment: common.LiveEnv})
	assert.Equal(t, "https://softposconfig-live.adyen.com/softposconfig/v3", client.SoftPOSConfiguration().SessionAuthenticationApi.BasePath())

	client.GetConfig().SoftPOSConfigurationEndpoint = adyen.SoftPOSConfigurationEndpointLiveUS + "/" + adyen.SoftPOSConfigurationAPIVersion
	assert.Equal(t, "https://softposconfig-live-us.adyen.com/softposconfig/v3", client.SoftPOSConfiguration().SessionAuthenticationApi.BasePath())
}

func TestPosMobileLiveNeedsPrefix(t *testing.T) {
	client := adyen.NewClient(&common.Config{ApiKey: "key", Environment: common.LiveEnv})
	assert.Equal(t, "", client.PosMobile().BasePath())
}
