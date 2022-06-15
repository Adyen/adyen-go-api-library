## Client generation steps using Open API generator

We use custom mustache templates to generate the open api client for the API specs. The templates are located under the `./templates/go` folder

**Step 1**: Generate the the client for a Schema using open api Docker image or cli. Provide the following parameters

-   `-i` path to open api spec yaml
-   `-t` path to custom templates in `./templates/go`
-   `-g` generator type `go`
-   `-p packageName= <Api namespace in StartCase>`
-   `-o` output path `./src/<Api namespace in lowercase>`

```
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
    -i /local/schema/CheckoutService-v67.yaml \
    -t /local/templates/go \
    -g go \
    -p packageName=Checkout \
    -o /local/src/checkout
```

or

```
openapi-generator-cli generate \
    -i /schema/CheckoutService-v67.yaml
    -t /templates/go \
    -g go \
    -p packageName=Checkout \
    -o /src/checkout
```

**Step 2**: Delete the following files/folders from `./src/<api package folder name in lowercase>`. If the foldername is not in lowercase, rename it

-   `configuration.go`
-   `client.go`
-   `utils.go`
-   `response.go`
-   `.travis.yml`
-   `git_push.sh`
-   `go.mod`
-   `go.sum`
-   `docs`

**Step 3**: Remove the HTTP method(Post, Get, Put, Patch) suffix on API endpoint methods (Regex to find them `([A-Z][a-zA-Z0-9]*)Post\(request`)

**Step 4**: Add the new service to `APIClient` struct in `./src/api/api.go` and add import for the same

```go
type APIClient struct {
	client *common.Client
	// API Services
    Checkout *checkout.Checkout
    <Api namespace in StartCase> *<Api namespace in lowercase>.<Api namespace in StartCase>
}
```

Init service in the `NewClient` method

```go
func NewClient(cfg *common.Config) *APIClient {
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

**Step 5**: Run `make run` or `go run main.go` and Fix any issues found

**Step 6**: Change any model fields that have `anyof` or `oneof` fields to `interface{}` (if we have typed structs) or `map[string]interface{}` (if its a key value pair). If `interface{}` is used then make sure to add a custom `UnmarshalJSON` method on the structs with test cases. See `src/checkout/model_payment_response.go` for example

**Step 7**: Add tests for the new APIs created under `./src/<Api namespace in lowercase>`

**Step 8**: Run `make test` or `go test ./...` and Fix any issues found
