package tests

import (
	"context"
	"fmt"
	"github.com/adyen/adyen-go-api-library/v8/src/adyen"
	"github.com/adyen/adyen-go-api-library/v8/src/checkout"
	"github.com/adyen/adyen-go-api-library/v8/src/common"
	"github.com/adyen/adyen-go-api-library/v8/src/recurring"
	"github.com/adyen/adyen-go-api-library/v8/src/webhook"
	"net/http"
	"net/url"
	"time"
)

func ExamplePaymentsApi_PaymentMethods() {
	client := adyen.NewClient(&common.Config{
		ApiKey:      "your api key",
		Environment: common.TestEnv,
	})
	service := client.Checkout()

	req := service.PaymentsApi.PaymentMethodsInput()
	req = req.PaymentMethodsRequest(checkout.PaymentMethodsRequest{
		MerchantAccount: "your merchant account",
	})
	res, httpRes, err := service.PaymentsApi.PaymentMethods(context.Background(), req)

	fmt.Println(res.GetPaymentMethods(), httpRes.StatusCode, err)
	// Output: [] 401 401 : HTTP Status Response - Unauthorized (security: 000)
}

func ExampleLiveEnv() {
	client := adyen.NewClient(&common.Config{
		ApiKey:                "your api key",
		Environment:           common.LiveEnv,
		LiveEndpointURLPrefix: "1797a841fbb37ca7-AdyenDemo", // Refer to https://docs.adyen.com/development-resources/live-endpoints#live-url-prefix
	})
	service := client.Checkout()

	req := service.PaymentsApi.PaymentMethodsInput()
	req = req.PaymentMethodsRequest(checkout.PaymentMethodsRequest{
		MerchantAccount: "your merchant account",
	})
	res, httpRes, err := service.PaymentsApi.PaymentMethods(context.Background(), req)

	fmt.Println(res.GetPaymentMethods(), httpRes, err.(*url.Error).URL)
	// Output: [] <nil> https://1797a841fbb37ca7-AdyenDemo-checkout-live.adyenpayments.com/checkout/v71/paymentMethods
}

func ExampleBasicAuth() {
	client := adyen.NewClient(&common.Config{
		Username:    "your ws user",
		Password:    "your secret password",
		Environment: common.TestEnv,
		UserAgent:   "Custom Application",
	})
	service := client.Recurring()

	req := service.ListRecurringDetailsInput()
	req = req.RecurringDetailsRequest(recurring.RecurringDetailsRequest{
		MerchantAccount: "your merchant account",
		Recurring: &recurring.Recurring{
			Contract: common.PtrString("RECURRING"),
		},
		ShopperReference: "ref",
	})
	res, httpRes, err := service.ListRecurringDetails(context.Background(), req)

	fmt.Println(res, httpRes, err)
}

func ExampleWebhook() {
	msg, err := webhook.HandleRequest(`{"live": "false", "notificationItems": []}`)

	fmt.Println(msg.Live, len(msg.GetNotificationItems()), err)
	// Output: false 0 <nil>
}

func ExampleError() {
	client := adyen.NewClient(&common.Config{
		ApiKey:      "your api key",
		Environment: common.TestEnv,
	})
	service := client.Checkout()

	req := service.PaymentsApi.PaymentsInput()
	paymentMethod := checkout.IdealDetailsAsCheckoutPaymentMethod(checkout.NewIdealDetails("1121"))
	_, httpRes, err := service.PaymentsApi.Payments(context.Background(), req.PaymentRequest(checkout.PaymentRequest{
		Reference: "123456781235",
		Amount: checkout.Amount{
			Value:    1250,
			Currency: "EUR",
		},
		CountryCode:     common.PtrString("NL"),
		MerchantAccount: "your merchant account",
		Channel:         common.PtrString("Web"),
		ReturnUrl:       "http://localhost:3000/redirect",
		PaymentMethod:   paymentMethod,
	}))

	httpStatusCode := httpRes.StatusCode
	errorMessage := err.(common.APIError).Message
	errorCode := err.(common.APIError).Code
	errorType := err.(common.APIError).Type

	fmt.Println(httpStatusCode, errorMessage, errorCode, errorType)
	fmt.Println(err.Error())
	// Output:
	// 401 HTTP Status Response - Unauthorized 000 security
	// 401 : HTTP Status Response - Unauthorized (security: 000)
}

func ExampleCustomHTTPClientConfiguration() {
	client := adyen.NewClient(&common.Config{
		HTTPClient: &http.Client{
			Timeout: 512 * time.Millisecond,
		},
		Environment: common.TestEnv,
		ApiKey:      "your api key",
	})

	fmt.Println(client.BinLookup().Client.Cfg.HTTPClient.Timeout)
	// Output: 512ms
}

func ExampleProxy() {
	// creating the proxyURL
	proxyURL, _ := url.Parse("http://myproxy:7000")
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	client := adyen.NewClient(&common.Config{
		HTTPClient: &http.Client{
			Transport: transport,
		},
		Environment: common.TestEnv,
		ApiKey:      "your api key",
	})

	fmt.Println(client)
}
