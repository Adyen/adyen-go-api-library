package tests

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/adyen/adyen-go-api-library/v21/src/adyen"
	"github.com/adyen/adyen-go-api-library/v21/src/checkout"
	"github.com/adyen/adyen-go-api-library/v21/src/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func checkoutAgainst(t *testing.T, status int, body string, cfg *common.Config) *checkout.APIClient {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		fmt.Fprint(w, body)
	}))
	t.Cleanup(server.Close)
	cfg.Environment = common.TestEnv
	service := adyen.NewClient(cfg).Checkout()
	service.PaymentsApi.BasePath = func() string { return server.URL }
	return service
}

func callPaymentMethods(service *checkout.APIClient) error {
	_, _, err := service.PaymentsApi.PaymentMethods(context.Background(), service.PaymentsApi.PaymentMethodsInput().
		PaymentMethodsRequest(checkout.PaymentMethodsRequest{MerchantAccount: "M"}))
	return err
}

// captureStdout returns what f writes to os.Stdout.
func captureStdout(t *testing.T, f func()) string {
	r, w, err := os.Pipe()
	require.NoError(t, err)
	orig := os.Stdout
	os.Stdout = w
	f()
	os.Stdout = orig
	require.NoError(t, w.Close())
	out, err := io.ReadAll(r)
	require.NoError(t, err)
	return string(out)
}

func TestLoggerReceivesDebugDumpsWithMaskedApiKey(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug}))
	service := checkoutAgainst(t, http.StatusOK, `{}`, &common.Config{ApiKey: "AQEsecretapikey12345", Debug: true, Logger: logger})

	var std bytes.Buffer
	log.SetOutput(&std)
	defer log.SetOutput(os.Stderr)
	require.NoError(t, callPaymentMethods(service))

	out := buf.String()
	assert.Contains(t, out, `"msg":"adyen request"`)
	assert.Contains(t, out, `"msg":"adyen response"`)
	assert.Contains(t, out, `"method":"POST"`)
	assert.NotContains(t, out, "AQEsecretapikey12345")
	assert.Contains(t, out, "12345")
	assert.Empty(t, std.String(), "nothing goes to the standard logger when Logger is set")
}

func TestLoggerReceivesErrorResponses(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, nil))
	body := `{"status":422,"errorCode":"14_030","message":"Return URL is missing."}`

	service := checkoutAgainst(t, http.StatusUnprocessableEntity, body, &common.Config{Log4XXError: true, Logger: logger})
	require.Error(t, callPaymentMethods(service))
	assert.Contains(t, buf.String(), "level=WARN")
	assert.Contains(t, buf.String(), "status=422")
	assert.Contains(t, buf.String(), "Return URL is missing.")

	buf.Reset()
	service = checkoutAgainst(t, http.StatusInternalServerError, `{"status":500}`, &common.Config{Log5XXError: true, Logger: logger})
	require.Error(t, callPaymentMethods(service))
	assert.Contains(t, buf.String(), "level=ERROR")

	buf.Reset()
	service = checkoutAgainst(t, http.StatusUnprocessableEntity, body, &common.Config{Log5XXError: true, Logger: logger})
	require.Error(t, callPaymentMethods(service))
	assert.Empty(t, buf.String(), "4XX responses are not logged when only Log5XXError is set")
}

func TestErrorResponsesGoToStdoutWithoutLogger(t *testing.T) {
	body := `{"status":422,"errorCode":"14_030","message":"Return URL is missing."}`
	service := checkoutAgainst(t, http.StatusUnprocessableEntity, body, &common.Config{Log4XXError: true})
	out := captureStdout(t, func() { _ = callPaymentMethods(service) })
	assert.Equal(t, body+"\n", out)
}

func TestUnparsableErrorBodyIsLoggedToLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, nil))
	service := checkoutAgainst(t, http.StatusBadGateway, `<html>bad gateway</html>`, &common.Config{Logger: logger})

	var std bytes.Buffer
	log.SetOutput(&std)
	defer log.SetOutput(os.Stderr)
	err := callPaymentMethods(service)
	require.Error(t, err)
	assert.Contains(t, buf.String(), "adyen error body is not JSON")
	assert.Empty(t, std.String())
	assert.True(t, strings.HasPrefix(err.Error(), "502"))
}

func TestPtr(t *testing.T) {
	assert.Equal(t, "EUR", *common.Ptr("EUR"))
	assert.Equal(t, int64(1000), *common.Ptr(int64(1000)))
	req := checkout.PaymentRequest{ShopperInteraction: common.Ptr(checkout.PaymentRequestShopperInteractionEcommerce)}
	assert.Equal(t, "Ecommerce", req.GetShopperInteraction())
}
