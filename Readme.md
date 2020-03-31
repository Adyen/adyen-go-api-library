## Adyen Golang API Client Library

The Adyen API Library for golang enables you to work with Adyen APIs.

## Integration

The Library supports all APIs under the following services:

- [x] checkout
- [] checkout utility
- [] payments
- [] modifications
- [] payouts
- [] recurring
- [] notifications
- [] BIN lookup

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

## HTTP Client Configuration

By default, Go [`http.DefaultClient`](https://golang.org/pkg/net/http/) will be used to submit requests to the API. But you can change that by injecting your own HttpClient on your client instance.

```go
config := NewConfig()

config.AddDefaultHeader("x-API-key", APIKey)
config.HTTPClient = &http.Client{
	CheckRedirect: redirectPolicyFunc,
}

client := NewAPIClient(config)
```

## Proxy configuration

You can configure a proxy connection by injecting your own `http.Client` on your client instance and changing the `ProxyFromEnvironment` or `ProxyURL` setter value in the Client Transport object.

Example:
// TODO

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

**Step 2**: Delete the following files from `./src/<api package folder name in lowercase>` - `configuration.go` - `client.go` - `utils.go` - `travis.yml` - `git_push.sh` - `go.mod` - `go.sum` - `go.sum`

**Step 3**: Add the new service to `APIClient` struct in `./main.go` and add import for the same

```go
type APIClient struct {
	client *common.Client
	// API Services
    Checkout *checkout.Checkout
    <Api namespace in StartCase> *<Api namespace in lowercase>.<Api namespace in StartCase>
}
```

**Step 4**: Run `go build main.go` and Fix any issues found

**Step 5**: Add tests for the new APIs created under `./src/<Api namespace in lowercase>`
