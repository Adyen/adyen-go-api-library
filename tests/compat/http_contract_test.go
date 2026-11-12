package compat

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/adyen/adyen-go-api-library/v21/src/adyen"
	"github.com/adyen/adyen-go-api-library/v21/src/checkout"
	"github.com/adyen/adyen-go-api-library/v21/src/common"
	"github.com/adyen/adyen-go-api-library/v21/src/transfers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func fixture(t *testing.T, name string) []byte {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("testdata", name))
	require.NoError(t, err)
	return b
}

type recorded struct {
	method, path string
	header       http.Header
	body         []byte
}

// stubServer answers every request with the given status and body, and records the last request.
func stubServer(t *testing.T, status int, body []byte) (*httptest.Server, *recorded) {
	t.Helper()
	rec := &recorded{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec.method, rec.path, rec.header = r.Method, r.URL.Path, r.Header.Clone()
		rec.body, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write(body)
	}))
	t.Cleanup(srv.Close)
	return srv, rec
}

func cardPaymentRequest() checkout.PaymentRequest {
	card := checkout.NewCardDetails()
	card.SetEncryptedCardNumber("test_4111111111111111")
	card.SetEncryptedExpiryMonth("test_03")
	card.SetEncryptedExpiryYear("test_2030")
	card.SetEncryptedSecurityCode("test_737")
	req := checkout.NewPaymentRequest(
		*checkout.NewAmount("EUR", 1000),
		"TestMerchant",
		checkout.CardDetailsAsCheckoutPaymentMethod(card),
		"compat-reference",
		"https://example.org/return",
	)
	req.SetShopperInteraction("Ecommerce")
	return *req
}

func TestCheckoutPaymentsRequestContract(t *testing.T) {
	srv, rec := stubServer(t, http.StatusOK, fixture(t, "checkout_payment_response.json"))
	client := adyen.NewClient(&common.Config{ApiKey: "test_api_key", Environment: common.TestEnv, ApplicationName: "compat-suite"})
	service := client.Checkout()
	service.PaymentsApi.BasePath = func() string { return srv.URL }

	res, httpRes, err := service.PaymentsApi.Payments(context.Background(),
		service.PaymentsApi.PaymentsInput().IdempotencyKey("idem-123").PaymentRequest(cardPaymentRequest()))

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, httpRes.StatusCode)

	assert.Equal(t, http.MethodPost, rec.method)
	assert.Equal(t, "/payments", rec.path)
	assert.Equal(t, "test_api_key", rec.header.Get("X-API-Key"))
	assert.Equal(t, "idem-123", rec.header.Get("Idempotency-Key"))
	assert.Equal(t, "application/json", rec.header.Get("Content-Type"))
	assert.Equal(t, "application/json", rec.header.Get("Accept"))
	assert.Equal(t, "no-cache", rec.header.Get("Cache-Control"))
	assert.Equal(t, "compat-suite adyen-go-api-library/"+common.LibVersion, rec.header.Get("User-Agent"))
	assert.Equal(t, "adyen-go-api-library", rec.header.Get("adyen-library-name"))
	assert.Equal(t, common.LibVersion, rec.header.Get("adyen-library-version"))
	assert.JSONEq(t, string(fixture(t, "checkout_payment_request.json")), string(rec.body))

	// Responses carrying fields or enum values unknown to this SDK version must still decode.
	assert.Equal(t, "8515131751004933", res.GetPspReference())
	assert.Equal(t, "SomeFutureResultCode", res.GetResultCode())
	assert.Equal(t, map[string]string{"cardSummary": "1111", "paymentMethod": "visa"}, res.GetAdditionalData())

	body, err := io.ReadAll(httpRes.Body)
	require.NoError(t, err)
	assert.JSONEq(t, string(fixture(t, "checkout_payment_response.json")), string(body), "the raw response body stays readable")
}

func TestDeprecatedContextIdempotencyKey(t *testing.T) {
	srv, rec := stubServer(t, http.StatusOK, fixture(t, "checkout_payment_response.json"))
	client := adyen.NewClient(&common.Config{ApiKey: "test_api_key", Environment: common.TestEnv})
	service := client.Checkout()
	service.PaymentsApi.BasePath = func() string { return srv.URL }

	ctx := common.WithIdempotencyKey(context.Background(), "ctx-key")
	_, _, err := service.PaymentsApi.Payments(ctx, service.PaymentsApi.PaymentsInput().PaymentRequest(cardPaymentRequest()))

	require.NoError(t, err)
	assert.Equal(t, "ctx-key", rec.header.Get("Idempotency-Key"))
	assert.Equal(t, "adyen-go-api-library/"+common.LibVersion, rec.header.Get("User-Agent"))
}

func TestBasicAuthWithoutAPIKey(t *testing.T) {
	srv, rec := stubServer(t, http.StatusOK, fixture(t, "checkout_payment_response.json"))
	client := adyen.NewClient(&common.Config{Username: "ws@Company.Test", Password: "secret", Environment: common.TestEnv})
	service := client.Checkout()
	service.PaymentsApi.BasePath = func() string { return srv.URL }

	_, _, err := service.PaymentsApi.Payments(context.Background(), service.PaymentsApi.PaymentsInput().PaymentRequest(cardPaymentRequest()))

	require.NoError(t, err)
	user, pass, ok := (&http.Request{Header: rec.header}).BasicAuth()
	require.True(t, ok)
	assert.Equal(t, "ws@Company.Test", user)
	assert.Equal(t, "secret", pass)
	assert.Empty(t, rec.header.Get("X-API-Key"))
}

func TestCheckoutErrorContract(t *testing.T) {
	errBody := fixture(t, "checkout_validation_error.json")
	srv, _ := stubServer(t, http.StatusUnprocessableEntity, errBody)
	client := adyen.NewClient(&common.Config{ApiKey: "test_api_key", Environment: common.TestEnv})
	service := client.Checkout()
	service.PaymentsApi.BasePath = func() string { return srv.URL }

	_, httpRes, err := service.PaymentsApi.Payments(context.Background(), service.PaymentsApi.PaymentsInput().PaymentRequest(cardPaymentRequest()))

	require.Error(t, err)
	require.NotNil(t, httpRes)
	assert.Equal(t, http.StatusUnprocessableEntity, httpRes.StatusCode)

	var apiErr common.APIError
	require.True(t, errors.As(err, &apiErr), "error type is %T", err)
	assert.Equal(t, "422 Unprocessable Entity", apiErr.Err)
	assert.Equal(t, float64(422), apiErr.Status)
	assert.Equal(t, "14_030", apiErr.Code)
	assert.Equal(t, "Return URL is missing.", apiErr.Message)
	assert.Equal(t, "validation", apiErr.Type)
	assert.JSONEq(t, string(errBody), string(apiErr.RawBody))
	assert.Equal(t, "422 Unprocessable Entity: Return URL is missing. (validation: 14_030)", err.Error())
}

func TestTransfersRestServiceErrorContract(t *testing.T) {
	srv, rec := stubServer(t, http.StatusUnprocessableEntity, fixture(t, "transfers_rest_service_error.json"))
	client := adyen.NewClient(&common.Config{ApiKey: "test_api_key", Environment: common.TestEnv})
	service := client.Transfers()
	service.TransfersApi.BasePath = func() string { return srv.URL }

	_, httpRes, err := service.TransfersApi.TransferFunds(context.Background(),
		service.TransfersApi.TransferFundsInput().IdempotencyKey("transfer-1").TransferInfo(transfers.TransferInfo{}))

	require.Error(t, err)
	require.NotNil(t, httpRes)
	assert.Equal(t, "/transfers", rec.path)
	assert.Equal(t, "transfer-1", rec.header.Get("Idempotency-Key"))

	var serviceErr common.RestServiceError
	require.True(t, errors.As(err, &serviceErr), "error type is %T", err)
	assert.Equal(t, int32(422), serviceErr.Status)
	assert.Equal(t, "30_112", serviceErr.ErrorCode)
	assert.Equal(t, "Invalid transfer information provided", serviceErr.Title)
	assert.Equal(t, "https://docs.adyen.com/errors/validation", serviceErr.Type)
	assert.Equal(t, "req-123", serviceErr.GetRequestId())
	require.Len(t, serviceErr.InvalidFields, 1)
	assert.Equal(t, common.InvalidField{Name: "amount.value", Value: "0", Message: "Amount must be greater than zero"}, serviceErr.InvalidFields[0])
}
