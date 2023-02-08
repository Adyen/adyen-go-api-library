## Adyen Golang API Client Library

This is the officially supported golang library for using Adyen's APIs.

## Integration

The Library supports all APIs under the following services:

* [Checkout API](https://docs.adyen.com/api-explorer/#/CheckoutService/v69/overview): Our latest integration for accepting online payments. Current supported version: **v69**
* [Payments API](https://docs.adyen.com/api-explorer/#/Payment/v64/overview): Our classic integration for online payments. Current supported version: **v64**
* [Recurring API](https://docs.adyen.com/api-explorer/#/Recurring/v49/overview): Endpoints for managing saved payment details. Current supported version: **v49**
* [Payouts API](https://docs.adyen.com/api-explorer/#/Payout/v64/overview): Endpoints for sending funds to your customers. Current supported version: **v64**
* [Platforms APIs](https://docs.adyen.com/platforms/api): Set of APIs when using Adyen for Platforms.
    * [Account API](https://docs.adyen.com/api-explorer/#/Account/v6/overview) Current supported version: **v6**
    * [Fund API](https://docs.adyen.com/api-explorer/#/Fund/v6/overview) Current supported version: **v6**
    * [Notification Configuration API](https://docs.adyen.com/api-explorer/#/NotificationConfigurationService/v6/overview) Current supported version: **v6**
* [Cloud-based Terminal API](https://docs.adyen.com/point-of-sale/terminal-api-reference): Our point-of-sale integration.
* [Legal Entity Management API](https://docs.adyen.com/api-explorer/#/legalentity/v2/overview): Manage legal entities that contain information required for verification. Current supported version: **v2**  

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
go get github.com/adyen/adyen-go-api-library/v6
```

## Usage examples

### Using APIs with APIKey

```go
import (
	"github.com/adyen/adyen-go-api-library/v6/src/checkout"
	"github.com/adyen/adyen-go-api-library/v6/src/common"
	"github.com/adyen/adyen-go-api-library/v6/src/adyen"
)

client := adyen.NewClient(&common.Config{
    ApiKey:      "your api key",
    Environment: common.TestEnv,
})

res, httpRes, err := client.Checkout.PaymentMethods(&checkout.PaymentMethodsRequest{
    MerchantAccount: "your merchant account",
})
```

### Using APIs with APIKey for Live env

```go
import (
    "github.com/adyen/adyen-go-api-library/v6/src/checkout"
    "github.com/adyen/adyen-go-api-library/v6/src/common"
    "github.com/adyen/adyen-go-api-library/v6/src/adyen"
)

client := adyen.NewClient(&common.Config{
    ApiKey:      "your api key",
    Environment: common.LiveEnv,
    LiveEndpointURLPrefix: "1797a841fbb37ca7-AdyenDemo", // Refer to https://docs.adyen.com/development-resources/live-endpoints#live-url-prefix
})

res, httpRes, err := client.Checkout.PaymentMethods(&checkout.PaymentMethodsRequest{
    MerchantAccount: "your merchant account",
})
```

### Using API with Basic Auth

```go
import (
    "github.com/adyen/adyen-go-api-library/v6/src/recurring"
    "github.com/adyen/adyen-go-api-library/v6/src/common"
    "github.com/adyen/adyen-go-api-library/v6/src/adyen"
)

client := adyen.NewClient(&common.Config{
    Username:        USER,
    Password:        PASS,
    Environment:     common.TestEnv,
    ApplicationName: "adyen-api-go-library",
})

res, httpRes, err := client.Recurring.ListRecurringDetails(&recurring.RecurringDetailsRequest{
    MerchantAccount: MerchantAccount,
    Recurring: &recurring.RecurringType{
        Contract: "RECURRING",
    },
    ShopperReference: "ref",
})
```

### Using Notifications parser

```go
import (
    "github.com/adyen/adyen-go-api-library/v6/src/adyen"
    "github.com/adyen/adyen-go-api-library/v6/src/common"
)

client := adyen.NewClient(&common.Config{
    ApiKey:      "your api key",
    Environment: common.TestEnv,
})

notification, err := client.Notification.HandleNotificationRequest(jsonRequestString)
```

### Getting error details

```go
import (
	"github.com/adyen/adyen-go-api-library/v6/src/common"
	"github.com/adyen/adyen-go-api-library/v6/src/checkout"
	"github.com/adyen/adyen-go-api-library/v6/src/adyen"
)

client := adyen.NewClient(&common.Config{
    ApiKey:      "your api key",
    Environment: common.TestEnv,
})

res, httpRes, err := client.Checkout.Payments(&checkout.PaymentRequest{
    Reference: "123456781235",
    Amount: checkout.Amount{
        Value:    1250,
        Currency: "EUR",
    },
    CountryCode:     "NL",
    MerchantAccount: MerchantAccount,
    Channel:         "Web",
    ReturnUrl:       "http://localhost:3000/redirect",
    PaymentMethod: map[string]interface{}{
        "type":   "ideal",
        "issuer": "1121",
    },
})

errorText := err.Error()
errorMessage := err.(common.APIError).Message
errorCode := err.(common.APIError).Code
errorType := err.(common.APIError).Type

httpStatusCode := httpRes.StatusCode
httpStatus := httpRes.Status
```

### Custom HTTP Client Configuration

By default, Go [`http.DefaultClient`](https://golang.org/pkg/net/http/) will be used to submit requests to the API. But you can change that by injecting your own HttpClient on your client instance.

```go
client := adyen.NewClient(&common.Config{
    HTTPClient:  &http.Client{
        CheckRedirect: redirectPolicyFunc,
        Timeout: 10 * time.MilliSeconds,
    },
    Environment: common.TestEnv,
    ApiKey:      "your api key",
})
```

### Proxy configuration

You can configure a proxy connection by injecting your own `http.Client` with a custom Transport on your client instance.

Example:

```go
//creating the proxyURL
proxyURL, _ := url.Parse("http://myproxy:7000")
transport := &http.Transport{
    Proxy: http.ProxyURL(proxyURL),
}
client = adyen.NewClient(&common.Config{
    HTTPClient:  &http.Client{
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
