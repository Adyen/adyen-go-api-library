package adyen_test

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/adyen/adyen-go-api-library/v21/src/adyen"
	"github.com/adyen/adyen-go-api-library/v21/src/checkout"
	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// fakeAdyen stands in for the Adyen API so that the examples run offline.
func fakeAdyen(status int, body string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		fmt.Fprint(w, body)
	}))
}

func cardPayment() checkout.PaymentRequest {
	return checkout.PaymentRequest{
		Amount:          checkout.Amount{Currency: "EUR", Value: 1000},
		MerchantAccount: "YOUR_MERCHANT_ACCOUNT",
		Reference:       "order-12345",
		ReturnUrl:       "https://your-company.example/checkout",
		PaymentMethod: checkout.CardDetailsAsCheckoutPaymentMethod(&checkout.CardDetails{
			Type:                  common.Ptr(checkout.CardDetailsTypeScheme),
			EncryptedCardNumber:   common.Ptr("test_4111111111111111"),
			EncryptedExpiryMonth:  common.Ptr("test_03"),
			EncryptedExpiryYear:   common.Ptr("test_2030"),
			EncryptedSecurityCode: common.Ptr("test_737"),
		}),
		ShopperInteraction:       common.Ptr(checkout.PaymentRequestShopperInteractionEcommerce),
		RecurringProcessingModel: common.Ptr(checkout.PaymentRequestRecurringProcessingModelCardOnFile),
	}
}

// Make a card payment with the Checkout API. Inline enum values such as the shopper
// interaction have constants named after the model and field.
func ExampleAPIClient_Checkout() {
	server := fakeAdyen(http.StatusOK, `{"pspReference":"PSP0001","resultCode":"Authorised"}`)
	defer server.Close()

	client := adyen.NewClient(&common.Config{ApiKey: "YOUR_API_KEY", Environment: common.TestEnv})
	service := client.Checkout()
	service.PaymentsApi.BasePath = func() string { return server.URL } // only for this example

	req := service.PaymentsApi.PaymentsInput().
		IdempotencyKey("order-12345-attempt-1").
		PaymentRequest(cardPayment())
	res, _, err := service.PaymentsApi.Payments(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}
	if res.GetResultCode() == checkout.PaymentResponseResultCodeAuthorised {
		fmt.Println("authorised:", res.GetPspReference())
	}
	// Output: authorised: PSP0001
}

// Read the error code and message of a request that Adyen rejected.
func ExampleAPIClient_Checkout_errorHandling() {
	server := fakeAdyen(http.StatusUnprocessableEntity,
		`{"status":422,"errorCode":"14_030","message":"Return URL is missing.","errorType":"validation"}`)
	defer server.Close()

	client := adyen.NewClient(&common.Config{ApiKey: "YOUR_API_KEY", Environment: common.TestEnv})
	service := client.Checkout()
	service.PaymentsApi.BasePath = func() string { return server.URL } // only for this example

	_, httpRes, err := service.PaymentsApi.Payments(context.Background(),
		service.PaymentsApi.PaymentsInput().PaymentRequest(cardPayment()))
	var apiErr common.APIError
	if errors.As(err, &apiErr) {
		fmt.Println(httpRes.StatusCode, apiErr.Code, apiErr.Message)
	}
	// Output: 422 14_030 Return URL is missing.
}

// Send the SDK's logs to a log/slog logger. Error responses are logged when the Log*XXError
// option for their status class is set; request and response dumps when Debug is set.
func ExampleNewClient_logger() {
	server := fakeAdyen(http.StatusUnprocessableEntity, `{"status":422,"errorCode":"14_030"}`)
	defer server.Close()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		// Drop the time and the URL (a random local port) to keep the example output stable.
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey || a.Key == "url" {
				return slog.Attr{}
			}
			return a
		},
	}))
	client := adyen.NewClient(&common.Config{
		ApiKey:      "YOUR_API_KEY",
		Environment: common.TestEnv,
		Log4XXError: true,
		Logger:      logger,
	})
	service := client.Checkout()
	service.PaymentsApi.BasePath = func() string { return server.URL } // only for this example

	_, _, _ = service.PaymentsApi.Payments(context.Background(),
		service.PaymentsApi.PaymentsInput().PaymentRequest(cardPayment()))
	// Output: level=WARN msg="adyen error response" status=422 body="{\"status\":422,\"errorCode\":\"14_030\"}" method=POST
}
