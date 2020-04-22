## Adyen Golang API Client Library

The Adyen API Library for golang enables you to work with Adyen APIs.

## Integration

The Library supports all APIs under the following services:

- [x] checkout
- [x] checkout utility
- [x] payments
    - [x] modifications
- [x] payouts
- [x] recurring
- [x] notifications
- [x] BIN lookup

## Requirements

- Go 1.13 or higher

## Installation

You can use go modules to add our library to your project

### Go module

```bash
go get github.com/adyen/adyen-go-api-library
```

## Documentation

- https://docs.adyen.com/developers/development-resources/libraries
- https://docs.adyen.com/developers/checkout

## Usage examples

```go
import (
	"github.com/adyen/adyen-go-api-library/src/checkout"
	"github.com/adyen/adyen-go-api-library/src/adyen"
)

client := adyen.NewClient(&common.Config{
    ApiKey:      "your api key",
    Environment: common.TestEnv,
})

res, httpRes, err := client.Checkout.PaymentMethods(&checkout.PaymentMethodsRequest{
    MerchantAccount: "your merchant account",
})
```

## Getting error details

```go
import (
	"github.com/adyen/adyen-go-api-library/src/common"
	"github.com/adyen/adyen-go-api-library/src/checkout"
	"github.com/adyen/adyen-go-api-library/src/adyen"
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

## HTTP Client Configuration

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

## Proxy configuration

You can configure a proxy connection by injecting your own `http.Client` with a custom Transport on your client instance.

Example:

```go
//creating the proxyURL
proxyURL, _ := url.Parse("http://localhost:7000")
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

If you have any problems, questions or suggestions, create an issue here or send your inquiry to support@adyen.com.

## Contributing

We strongly encourage you to join us in contributing to this repository so everyone can benefit from:

- New features and functionality
- Resolved bug fixes and issues
- Any general improvements

Read our [**contribution guidelines**](CONTRIBUTING.md) to find out how.

## Licence

MIT license. For more information, see the LICENSE file.
