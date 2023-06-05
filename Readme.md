## Adyen Golang API Client Library

This is the officially supported golang library for using Adyen's APIs.

## Supported API versions

The Library supports all APIs under the following services:

| API                                                                                                                  | Description                                                                                                                                                                                                                                             | Client accessor method                      | Supported version |
|----------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------|-------------------|
| [Checkout API](https://docs.adyen.com/api-explorer/Checkout/70/overview)                                             | Our latest integration for accepting online payments.                                                                                                                                                                                                   | client.Checkout()                           | **v70**           |
| [Payouts API](https://docs.adyen.com/api-explorer/#/Payout/v68/overview)                                             | Endpoints for sending funds to your customers.                                                                                                                                                                                                          | client.Payout()                             | **v68**           |
| [Recurring API](https://docs.adyen.com/api-explorer/#/Recurring/v68/overview)                                        | Endpoints for managing saved payment details.                                                                                                                                                                                                           | client.Recurring()                          | **v68**           |
| [BIN lookup API](https://docs.adyen.com/api-explorer/BinLookup/54/overview)                                          | The BIN Lookup API provides endpoints for retrieving information based on a given BIN.                                                                                                                                                                  | client.BinLookup()                          | **v54**           |
| [POS Terminal Management API](https://docs.adyen.com/api-explorer/#/postfmapi/v1/overview)                           | Endpoints for managing your point-of-sale payment terminals.                                                                                                                                                                                            | client.PosTerminalManagement()              | **v1**            |
| [Management API](https://docs.adyen.com/api-explorer/#/ManagementService/v1/overview)                                | Configure and manage your Adyen company and merchant accounts, stores, and payment terminals.                                                                                                                                                           | client.Management()                         | **v1**            |
| [Legal Entity Management API](https://docs.adyen.com/api-explorer/legalentity/3/overview)                            | Manage legal entities that contain information required for verification.                                                                                                                                                                               | client.LegalEntity()                        | **v3**            |
| [Configuration API](https://docs.adyen.com/api-explorer/#/balanceplatform/v2/overview)                               | The Configuration API enables you to create a platform where you can onboard your users as account holders and create balance accounts, cards, and business accounts.                                                                                   | client.BalancePlatform()                    | **v2**            |
| [Transfers API](https://docs.adyen.com/api-explorer/transfers/3/overview)                                            | The Transfers API provides endpoints that can be used to get information about all your transactions, move funds within your balance platform or send funds from your balance platform to a transfer instrument.                                        | client.Transfers()                          | **v3**            |
| [Payments API](https://docs.adyen.com/api-explorer/#/Payment/v68/overview)                                           | Our classic integration for online payments.                                                                                                                                                                                                            | client.Payments()                           | **v68**           |
| [Account API](https://docs.adyen.com/api-explorer/#/Account/v6/overview)                                             | This API is used for the classic integration. If you are just starting your implementation, refer to our [new integration guide](https://docs.adyen.com/marketplaces-and-platforms) instead.                                                            | client.PlatformsAccount()                   | **v6**            |
| [Fund API](https://docs.adyen.com/api-explorer/#/Fund/v6/overview)                                                   | This API is used for the classic integration. If you are just starting your implementation, refer to our [new integration guide](https://docs.adyen.com/marketplaces-and-platforms) instead.                                                            | client.PlatformsFund()                      | **v6**            |
| [Hosted onboarding API](https://docs.adyen.com/api-explorer/#/Hop/v6/overview)                                       | This API is used for the classic integration. If you are just starting your implementation, refer to our [new integration guide](https://docs.adyen.com/marketplaces-and-platforms) instead.                                                            | client.PlatformsHostedOnboardingPage()      | **v6**            |
| [Notification Configuration API](https://docs.adyen.com/api-explorer/#/NotificationConfigurationService/v6/overview) | This API is used for the classic integration. If you are just starting your implementation, refer to our [new integration guide](https://docs.adyen.com/marketplaces-and-platforms) instead.                                                            | client.PlatformsNotificationConfiguration() | **v6**            |
| [Platforms Notifications Webhooks](https://docs.adyen.com/api-explorer/#/NotificationService/v6/overview)            |                                                                                                                                                                                                                                                         | *Models only*                               | **v6**            |
| [Stored Value API](https://docs.adyen.com/payment-methods/gift-cards/stored-value-api)                               | Manage both online and point-of-sale gift cards and other stored-value cards.                                                                                                                                                                           | client.StoredValue()                        | **v46**           |
| [Webhooks](https://docs.adyen.com/api-explorer/Webhooks/1/overview)                                                  | Adyen uses webhooks to send notifications about payment status updates, newly available reports, and other events that can be subscribed to. For more information, refer to our [documentation](https://docs.adyen.com/development-resources/webhooks). | *Models only*                               | **v1**            |

For more information, refer to our [documentation](https://docs.adyen.com/) or the [API Explorer](https://docs.adyen.com/api-explorer/).

## Prerequisites

-   Go 1.13 or higher
-   [Adyen test account](https://docs.adyen.com/get-started-with-adyen)
-   [API key](https://docs.adyen.com/development-resources/api-credentials#generate-api-key). For testing, your API credential needs to have the [API PCI Payments role](https://docs.adyen.com/development-resources/api-credentials#roles).
-   [Basic Authentication](https://docs.adyen.com/development-resources/api-credentials#basic-authentication) (for Legal Entity Management API only). 


## Documentation

-   https://docs.adyen.com/developers/development-resources/libraries
-   https://docs.adyen.com/developers/checkout

## Installation

You can use go modules to add our library to your project

```bash
go get github.com/adyen/adyen-go-api-library/v7
```

## Usage examples

### Using APIs with APIKey

```go
import (
	"context"
	"github.com/adyen/adyen-go-api-library/v7/src/checkout"
	"github.com/adyen/adyen-go-api-library/v7/src/common"
	"github.com/adyen/adyen-go-api-library/v7/src/adyen"
)

client := adyen.NewClient(&common.Config{
    ApiKey:      "your api key",
    Environment: common.TestEnv,
})
service := client.Checkout()

req := service.PaymentsApi.PaymentMethodsConfig(context.Background())
req = req.PaymentMethodsRequest(checkout.PaymentMethodsRequest{
    MerchantAccount: "your merchant account",
})
res, httpRes, err := service.PaymentsApi.PaymentMethods(req)
```

### Using APIs with APIKey for Live env

```go
import (
    "github.com/adyen/adyen-go-api-library/v7/src/checkout"
    "github.com/adyen/adyen-go-api-library/v7/src/common"
    "github.com/adyen/adyen-go-api-library/v7/src/adyen"
)

client := adyen.NewClient(&common.Config{
    ApiKey:                "your api key",
    Environment:           common.LiveEnv,
    LiveEndpointURLPrefix: "1797a841fbb37ca7-AdyenDemo", // Refer to https://docs.adyen.com/development-resources/live-endpoints#live-url-prefix
})
service := client.Checkout()

req := service.PaymentsApi.PaymentMethodsConfig(context.Background())
req = req.PaymentMethodsRequest(checkout.PaymentMethodsRequest{
    MerchantAccount: "your merchant account",
})
res, httpRes, err := service.PaymentsApi.PaymentMethods(req)
```

### Using API with Basic Auth

```go
import (
    "github.com/adyen/adyen-go-api-library/v7/src/recurring"
    "github.com/adyen/adyen-go-api-library/v7/src/common"
    "github.com/adyen/adyen-go-api-library/v7/src/adyen"
)

client := adyen.NewClient(&common.Config{
    Username:    "your ws user",
    Password:    "your secret password",
    Environment: common.TestEnv,
    UserAgent:   "Custom Application",
})
service := client.Recurring()

req := service.ListRecurringDetailsConfig(context.Background())
req = req.RecurringDetailsRequest(recurring.RecurringDetailsRequest{
    MerchantAccount: "your merchant account",
    Recurring: &recurring.Recurring{
        Contract: common.PtrString("RECURRING"),
    },
    ShopperReference: "ref",
})
res, httpRes, err := service.ListRecurringDetails(req)
```

### Using Webhook parser

```go
import (
	"github.com/adyen/adyen-go-api-library/v7/src/webhook"
)

msg, err := webhook.HandleRequest(`{"live": "false", "notificationItems": []}`)
```

### Getting error details

```go
import (
	"github.com/adyen/adyen-go-api-library/v7/src/common"
	"github.com/adyen/adyen-go-api-library/v7/src/checkout"
	"github.com/adyen/adyen-go-api-library/v7/src/adyen"
)

client := adyen.NewClient(&common.Config{
    ApiKey:      "your api key",
    Environment: common.TestEnv,
})
service := client.Checkout()

req := service.PaymentsApi.PaymentsConfig(context.Background())
paymentMethod := checkout.IdealDetailsAsCheckoutPaymentMethod(checkout.NewIdealDetails("1121"))
_, httpRes, err := service.PaymentsApi.Payments(req.PaymentRequest(checkout.PaymentRequest{
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

### Custom HTTP Client Configuration

By default, Go [`http.DefaultClient`](https://golang.org/pkg/net/http/) will be used to submit requests to the API. But you can change that by injecting your own HttpClient on your client instance.

```go
client := adyen.NewClient(&common.Config{
    HTTPClient: &http.Client{
        Timeout: 512 * time.Millisecond,
    },
    Environment: common.TestEnv,
    ApiKey:      "your api key",
})
```

### Proxy configuration

You can configure a proxy connection by injecting your own `http.Client` with a custom Transport on your client instance.

Example:

```go
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
```

## Support

If you have a feature request, or spotted a bug or a technical problem, create a github issue. 
For other questions, contact our [support team](https://support.adyen.com/hc/en-us/requests/new?ticket_form_id=360000705420).

## Contributing

We strongly encourage you to join us in contributing to this repository so everyone can benefit from:

-   New features and functionality
-   Resolved bug fixes and issues
-   Any general improvements

Read our [**contribution guidelines**](CONTRIBUTING.md) to find out how.

## Licence

MIT license. For more information, see the LICENSE file.
