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
- [ ] notifications
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

## Sample Usage

```go
import (
	"github.com/adyen/adyen-go-api-library/src/checkout"
	adyen "github.com/adyen/adyen-go-api-library/src/api"
)

client := adyen.NewAPIClientWithAPIKey("your api key", "TEST")

res, httpRes, err := client.Checkout.PaymentMethodsPost(&checkout.PaymentMethodsRequest{
    MerchantAccount: "your merchant account",
})
```

## Getting error details

```go
import (
	"github.com/adyen/adyen-go-api-library/src/common"
	"github.com/adyen/adyen-go-api-library/src/checkout"
	adyen "github.com/adyen/adyen-go-api-library/src/api"
)

client := adyen.NewAPIClientWithAPIKey("your api key", "TEST")

res, httpRes, err := client.Checkout.PaymentsPost(&checkout.PaymentRequest{
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
client := NewAPIClient(&common.Config{
    HTTPClient: &http.Client{
        CheckRedirect: redirectPolicyFunc,
        Timeout: 10 * time.MilliSeconds,
    },
    Environment: "TEST",
    ApiKey: "your api key",
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
client = NewAPIClient(&common.Config{
    HTTPClient: &http.Client{
        Transport: transport,
    },
    Environment: "TEST",
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

## Client generation steps using Open API generator

We use custom mustache templates to generate the open api client for the API specs. The templates are located under the `./templates/go` folder

**Step 1**: Generate the the client for a Schema using open api Docker image or cli. Provide the following parameters

- `-i` path to open api spec yaml
- `-t` path to custom templates in `./templates/go`
- `-g` generator type `go-experimental`
- `-p packageName= <Api namespace in StartCase>`
- `-o` output path `./src/<Api namespace in lowercase>`

```
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
    -i /local/schema/CheckoutService-v51.yaml \
    -t /local/templates/go \
    -g go-experimental \
    -p packageName=Checkout \
    -o /local/src/checkout
```

or

```
openapi-generator-cli generate \
    -i /schema/CheckoutService-v51.yaml
    -t /templates/go \
    -g go-experimental \
    -p packageName=Checkout \
    -o /src/checkout
```

**Step 2**: Delete the following files from `./src/<api package folder name in lowercase>`. If the foldername is not in lowercase, rename it

- `configuration.go`
- `client.go`
- `utils.go`
- `response.go`
- `travis.yml`
- `git_push.sh`
- `go.mod`
- `go.sum`

**Step 3**: Add the new service to `APIClient` struct in `./src/api/api.go` and add import for the same

```go
type APIClient struct {
	client *common.Client
	// API Services
    Checkout *checkout.Checkout
    <Api namespace in StartCase> *<Api namespace in lowercase>.<Api namespace in StartCase>
}
```

Init service in the `NewAPIClient` method

```go
func NewAPIClient(cfg *common.Config) *APIClient {
    ...
    // API Services
	c.Checkout = &checkout.Checkout{
		Client: c.client,
		BasePath: func() string {
			return fmt.Sprintf("%s/%s", c.client.Cfg.CheckoutEndpoint, CheckoutAPIVersion)
		},
	}
    c.<Api namespace in StartCase> = &<Api namespace in lowercase>.<Api namespace in StartCase>{
        Client: c.client,
        BasePath: func() string {
            return fmt.Sprintf("%s/%s", c.client.Cfg.<API end point>, <API version constant>)
        },
    }
}
```

**Step 4**: Run `make run` or `go run main.go` and Fix any issues found

**Step 5**: Add tests for the new APIs created under `./src/<Api namespace in lowercase>`

**Step 6**: Run `make test` or `go test ./...` and Fix any issues found
