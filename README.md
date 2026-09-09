![Go](https://github.com/Adyen/adyen-go-api-library/assets/156021/23cd84f0-f918-456d-8f3f-ee6d0fc8ef2e)

# Adyen Golang API Client Library

[![Go Version](https://img.shields.io/github/v/tag/adyen/adyen-go-api-library?label=version)](https://pkg.go.dev/github.com/adyen/adyen-go-api-library/v21?tab=versions)

This is the officially supported golang library for using Adyen's APIs.

## Supported API versions

The Library supports all APIs under the following services:

| API                                                                                                        | Description                                                                                                                                                                                                                                                                                                                              | Service constructor                         | Supported version |
|------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------|-------------------|
| [Checkout API](https://docs.adyen.com/api-explorer/Checkout/71/overview)                                   | Our latest integration for accepting online payments.                                                                                                                                                                                                                                                                                    | client.Checkout()                           | **v71**           |
| [Payouts API](https://docs.adyen.com/api-explorer/Payout/68/overview)                                      | Endpoints for sending funds to your customers.                                                                                                                                                                                                                                                                                           | client.Payout()                             | **v68**           |
| [Recurring API](https://docs.adyen.com/api-explorer/Recurring/68/overview)                                 | Endpoints for managing saved payment details.                                                                                                                                                                                                                                                                                            | client.Recurring()                          | **v68**           |
| [BIN lookup API](https://docs.adyen.com/api-explorer/BinLookup/54/overview)                                | The BIN Lookup API provides endpoints for retrieving information based on a given BIN.                                                                                                                                                                                                                                                   | client.BinLookup()                          | **v54**           |
| [Disputes API](https://docs.adyen.com/api-explorer/Disputes/30/overview)                                   | You can use the [Disputes API](https://docs.adyen.com/risk-management/disputes-api/) to automate the dispute handling process so that you can respond to disputes and chargebacks as soon as they are initiated. The Disputes API lets you retrieve defense reasons, supply and delete defense documents, and accept or defend disputes. | client.Disputes()                           | **v30**           |
| [POS Terminal Management API](https://docs.adyen.com/api-explorer/postfmapi/1/overview)                    | ~~Endpoints for managing your point-of-sale payment~~ ‼️ **Deprecated**: use instead the [Management API](https://docs.adyen.com/api-explorer/Management/latest/overview) for the management of your terminal fleet. terminals.                                                                                                                                                                                                                                                                             | ~~client.PosTerminalManagement()~~              | ~~**v1**~~            |
| [Management API](https://docs.adyen.com/api-explorer/Management/3/overview)                                | Configure and manage your Adyen company and merchant accounts, stores, and payment terminals.                                                                                                                                                                                                                                            | client.Management()                         | **v3**            |
| [Data Protection API](https://docs.adyen.com/development-resources/data-protection-api)                    | Adyen Data Protection API provides a way for you to process [Subject Erasure Requests](https://gdpr-info.eu/art-17-gdpr/) as mandated in GDPR.                                                                                                                                                                                           | client.DataProtection()                     | **v1**            |
| [Balance Control API](https://docs.adyen.com/api-explorer/BalanceControl/1/overview)                       | The Balance Control API lets you transfer funds between merchant accounts that belong to the same legal entity and are under the same company account.                                                                                                                                                                                   | client.BalanceControl()                     | **v1**            |
| [Legal Entity Management API](https://docs.adyen.com/api-explorer/legalentity/3/overview)                  | Manage legal entities that contain information required for verification.                                                                                                                                                                                                                                                                | client.LegalEntity()                        | **v3**            |
| [Configuration API](https://docs.adyen.com/api-explorer/balanceplatform/2/overview)                        | The Configuration API enables you to create a platform where you can onboard your users as account holders and create balance accounts, cards, and business accounts.                                                                                                                                                                    | client.BalancePlatform()                    | **v2**            |
| [Transfers API](https://docs.adyen.com/api-explorer/transfers/4/overview)                                  | The Transfers API provides endpoints that can be used to get information about all your transactions, move funds within your balance platform or send funds from your balance platform to a transfer instrument.                                                                                                                         | client.Transfers()                          | **v4**            |
| [Stored Value API](https://docs.adyen.com/payment-methods/gift-cards/stored-value-api)                     | Manage both online and point-of-sale gift cards and other stored-value cards.                                                                                                                                                                                                                                                            | client.StoredValue()                        | **v46**           |
| [POS Mobile API](https://docs.adyen.com/api-explorer/possdk/68/overview)                                    | The POS Mobile API is used in the mutual authentication flow between an Adyen Android or iOS [POS Mobile SDK](https://docs.adyen.com/point-of-sale/ipp-mobile/) and the Adyen payments platform. The POS Mobile SDK for Android or iOS devices enables businesses to accept in-person payments using a commercial off-the-shelf (COTS) device like a phone. For example, Tap to Pay transactions, or transactions on a mobile device in combination with a card reader |  client.PosMobile()                        | **v68**                               |
| [Payments App API](https://docs.adyen.com/api-explorer/payments-app/1/overview)                                    | The Payments App API is used to Board and manage the Adyen Payments App on your Android mobile devices | client.PaymentsApp()     | **v1**              | 
| [Payments API](https://docs.adyen.com/api-explorer/Payment/68/overview)                                    | Our classic integration for online payments.                                                                                                                                                                                                                                                                                             | client.Payments()                           | **v68**           |


For more information, refer to our [documentation](https://docs.adyen.com/) or the [API Explorer](https://docs.adyen.com/api-explorer/).

## Supported Webhook versions

The library supports all webhooks under the following model directories:

| Webhooks                                                                                          | Description                                                                                                                                                                                                                                                                         | Package                                                        | Supported Version |
|---------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------|-------------------|
| [Payment Webhooks](https://docs.adyen.com/api-explorer/Webhooks/1/overview)                       | Adyen uses webhooks to send notifications about payment status updates, newly available reports, and other events that can be subscribed to. For more information, refer to our [documentation](https://docs.adyen.com/development-resources/webhooks).                             | [webhook](src/webhook)                                         | **v1**            |
| [Authentication Webhooks](https://docs.adyen.com/api-explorer/acs-webhook/1/overview)             | Adyen sends this webhook when the process of cardholder authentication is finalized, whether it is completed successfully, fails, or expires.                                                                                                                                       | [acswebhook](src/acswebhook)                                   | **v1**            |
| [Configuration Webhooks](https://docs.adyen.com/api-explorer/balanceplatform-webhooks/2/overview) | You can use these webhooks to build your implementation. For example, you can use this information to update internal statuses when the status of a capability is changed.                                                                                                          | [configurationwebhook](src/configurationwebhook)               | **v2**            |
| [Transfer Webhooks](https://docs.adyen.com/api-explorer/transfer-webhooks/4/overview)             | You can use these webhooks to build your implementation. For example, you can use this information to update balances in your own dashboards or to keep track of incoming funds.                                                                                                    | [transferwebhook](src/transferwebhook)                         | **v4**            |
| [Report Webhooks](https://docs.adyen.com/api-explorer/report-webhooks/1/overview)                 | You can download reports programmatically by making an HTTP GET request, or manually from your Balance Platform Customer Area                                                                                                                                                       | [reportwebhook](src/reportwebhook)                             | **v1**            |
| [Management Webhooks](https://docs.adyen.com/api-explorer/ManagementNotification/3/overview)      | Adyen uses webhooks to inform your system about events that happen with your Adyen company and merchant accounts, stores, payment terminals, and payment methods when using Management API.                                                                                         | [managementwebhook](src/managementwebhook)                     | **v3**            |
| [Transaction Webhooks](https://docs.adyen.com/api-explorer/transaction-webhooks/4/overview)       | Adyen sends webhooks to inform your system about incoming and outgoing transfers in your platform. You can use these webhooks to build your implementation. For example, you can use this information to update balances in your own dashboards or to keep track of incoming funds. | [transactionwebhook](src/transactionwebhook)                   | **v4**            |

## Prerequisites

-   Go 1.23 or higher
-   [Adyen test account](https://docs.adyen.com/get-started-with-adyen)
-   [API key](https://docs.adyen.com/development-resources/api-credentials#generate-api-key). For testing, your API credential needs to have the [API PCI Payments role](https://docs.adyen.com/development-resources/api-credentials#roles).
-   [Basic Authentication](https://docs.adyen.com/development-resources/api-credentials#basic-authentication) (for Legal Entity Management API only). 


## Documentation

-   https://docs.adyen.com/developers/development-resources/libraries
-   https://docs.adyen.com/developers/checkout

## Installation

You can use go modules to add our library to your project

```bash
go get github.com/adyen/adyen-go-api-library/v21@v21.2.2
```

## Usage examples

### Using APIs with APIKey

```go
import (
	"context"
	"github.com/adyen/adyen-go-api-library/v21/src/checkout"
	"github.com/adyen/adyen-go-api-library/v21/src/common"
	"github.com/adyen/adyen-go-api-library/v21/src/adyen"
)

client := adyen.NewClient(&common.Config{
    ApiKey:      "your api key",
    Environment: common.TestEnv,
    Log4XXError: true,
})
service := client.Checkout()

req := service.PaymentsApi.PaymentMethodsInput()
req = req.PaymentMethodsRequest(checkout.PaymentMethodsRequest{
    MerchantAccount: "your merchant account",
})
res, httpRes, err := service.PaymentsApi.PaymentMethods(context.Background(), req)
```

### Using APIs with APIKey for Live env

```go
import (
    "github.com/adyen/adyen-go-api-library/v21/src/checkout"
    "github.com/adyen/adyen-go-api-library/v21/src/common"
    "github.com/adyen/adyen-go-api-library/v21/src/adyen"
)

client := adyen.NewClient(&common.Config{
    ApiKey:                "your api key",
    Environment:           common.LiveEnv,
    Log4XXError: true,
    LiveEndpointURLPrefix: "1797a841fbb37ca7-AdyenDemo", // Refer to https://docs.adyen.com/development-resources/live-endpoints#live-url-prefix
})
service := client.Checkout()

req := service.PaymentsApi.PaymentMethodsInput()
req = req.PaymentMethodsRequest(checkout.PaymentMethodsRequest{
    MerchantAccount: "your merchant account",
})
res, httpRes, err := service.PaymentsApi.PaymentMethods(context.Background(), req)
```

### Using API with Basic Auth

```go
import (
    "github.com/adyen/adyen-go-api-library/v21/src/recurring"
    "github.com/adyen/adyen-go-api-library/v21/src/common"
    "github.com/adyen/adyen-go-api-library/v21/src/adyen"
)

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
```

## Unmarshalling JSON Strings
In some setups you might need to unmarshal JSON strings to request objects. For example, when using the libraries in combination with Dropin/Components. Please use following methodology when unmarshalling JSON strings:

```go
import (
	"encoding/json"
	"github.com/adyen/adyen-go-api-library/v21/src/checkout"
)
paymentRequest := checkout.PaymentRequest{}
error := json.Unmarshal([]byte("YOUR_JSON_STRING"), &paymentRequest)

```

### Using the webhook parser

```go
import (
	"github.com/adyen/adyen-go-api-library/v21/src/webhook"
)

msg, err := webhook.HandleRequest(`{"live": "false", "notificationItems": []}`)
```

### Getting error details

```go
import (
	"github.com/adyen/adyen-go-api-library/v21/src/common"
	"github.com/adyen/adyen-go-api-library/v21/src/checkout"
	"github.com/adyen/adyen-go-api-library/v21/src/adyen"
)

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
```

Consider additional logging when an error occurs:
```go

client := adyen.NewClient(&common.Config{
ApiKey:      "your api key",
Environment: common.TestEnv,
Log3XXError: true,
Log4XXError: true,
Log5XXError: true,
})

service := client.Checkout()
```

### Configure the HTTP client

By default, the library uses Go's [`http.DefaultClient`](https://pkg.go.dev/net/http#DefaultClient). You can customize HTTP behavior by injecting a configured `*http.Client` through `common.Config.HTTPClient`. Create the client and its transport once, then share the resulting Adyen client across your application. Go HTTP clients and transports are safe for concurrent use.

```go
transport := http.DefaultTransport.(*http.Transport).Clone()
transport.MaxIdleConns = 100
transport.MaxIdleConnsPerHost = 20
transport.IdleConnTimeout = 30 * time.Second

httpClient := &http.Client{
    Transport: transport,
    Timeout:   60 * time.Second,
}

client := adyen.NewClient(&common.Config{
    HTTPClient:  httpClient,
    Environment: common.TestEnv,
    ApiKey:      "your api key",
})
```

Clone `http.DefaultTransport` before customizing it. This preserves Go's default proxy handling, connection and TLS timeouts, TCP keepalive, and HTTP/2 support. Do not modify `http.DefaultTransport` directly because it is shared by the process. Creating an empty `http.Transport` does not inherit these defaults.

You can configure the following settings on the client and transport:

- `http.Client.Timeout` sets the maximum duration for the complete request and response operation. It does not only control TCP connection establishment.
- `MaxIdleConns` sets the maximum number of idle connections retained across all destinations.
- `MaxIdleConnsPerHost` sets the maximum number of idle connections retained for one destination. Increasing it can reduce HTTP/1.1 connection churn under concurrent traffic.
- `IdleConnTimeout` controls how long an unused pooled connection remains available. Consider setting it below relevant proxy, firewall, load balancer, or server idle limits.
- `MaxConnsPerHost` limits dialing, active, and idle connections for one destination. Set it only when you intend to limit concurrency, because it can throttle requests.
- `DisableKeepAlives` disables HTTP connection reuse. It increases TCP and TLS setup overhead, so use it for diagnosis rather than as the default solution.
- `DialContext` and `net.Dialer.KeepAlive` configure low-level connection behavior. TCP keepalive probes are different from HTTP connection reuse.
- `Proxy` and `TLSClientConfig` allow proxy and TLS configuration specific to your network.

For high-volume traffic, keep HTTP connection reuse enabled and size the connection pool using observed concurrency and infrastructure limits. 
Do not create a new client or transport per request, because doing so prevents effective connection reuse. 
Use stable idempotency keys on supported POST operations when retries are possible. 

#### Proxy configuration

You can configure a proxy by setting `Proxy` on a cloned default transport.

```go
proxyURL, err := url.Parse("http://myproxy:7000")
if err != nil {
    // Handle the configuration error.
}

transport := http.DefaultTransport.(*http.Transport).Clone()
transport.Proxy = http.ProxyURL(proxyURL)

client := adyen.NewClient(&common.Config{
    HTTPClient: &http.Client{
        Transport: transport,
    },
    Environment: common.TestEnv,
    ApiKey:      "your api key",
})
```

## Support

If you have a feature request, or spotted a bug or a technical problem, [create an issue here](https://github.com/Adyen/adyen-go-api-library/issues/new/choose).

For other questions, [contact our Support Team](https://www.adyen.help/hc/en-us/requests/new?ticket_form_id=360000705420).

## Contributing

We strongly encourage you to join us in contributing to this repository so everyone can benefit from:

-   New features and functionality
-   Resolved bug fixes and issues
-   Any general improvements

Read our [**contribution guidelines**](CONTRIBUTING.md) to find out how.

## Feedback

We value your input! Help us enhance our API Libraries and improve the integration experience by providing your feedback. Please take a moment to fill out [our feedback form](https://forms.gle/A4EERrR6CWgKWe5r9) to share your thoughts, suggestions or ideas.

## License

MIT license. For more information, see the LICENSE file.

