## Client generation using OpenAPI generator

The [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) is used to generate the SDK for the API.

The following artifacts are generated using the OpenAPI Generator:
* services
* models
* docs
* README.md

The following artifacts are NOT generated (created manually):
* configuration.go
* tests


We use custom Mustache templates to customise the code generation. The templates are located under the `./templates/go-v6.4.0` folder.

The dafault [Go templates](https://github.com/OpenAPITools/openapi-generator/tree/master/modules/openapi-generator/src/main/resources/go) are 
modified to include configure/customize the following:
* set Username and Password for Basic Authentication
* use the Adyen extension `x-methodName` in the OpenAPI file (to name the Go functions)
* lowercase package name when necessary

**Step 1**: Use the OpenAPI Generator cli (or Docker) setting the required parameters: 

-   `-i` path to OpenAPI file (yaml or json)
-   `-t` path to the custom Mustache templates in `./templates/go`
-   `-g` generator type `go`
-   `-p packageName= <Api namespace in StartCase>`
-   `-o` output path `./src/<Api namespace in lowercase>`
-   `--git-repo-id` Lib name `adyen-go-api-library/v6`
-   `--git-user-id` Github user `adyen`
-   `--global-property apiTests=false` to skip tests generation

```
openapi-generator-cli generate \
    -i /schema/LegalEntityService-v2.yaml
    -t /templates/go-v6.4.0 \
    -g go \
    -p packageName=LegalEntity \
    -o /src/legalentity \
    --git-repo-id adyen-go-api-library/v6 --git-user-id adyen \
    --global-property apiTests=false
```

**Step 2**: Perform health check
``````
go run main.go
``````

**Step 3**: Run LegalEntity unit tests only
``````
go clean -testcache
go test -v ./tests/legalentity/...
``````

**Step 4**: Run LegalEntity unit and integration tests

**Note:** integration tests require `username` and `password` for Basic Authentication as well as
the IDs of entities to test

In the `.env` file configure the following
``````
ADYEN_LEM_USERNAME=your_lem_username
ADYEN_LEM_PASSWORD=your_lem_password

LEM_LEGAL_ENTITY_ID=your_legal_entity_id
LEM_BUSINESS_LINE_ID=your_business_line_id
LEM_DOCUMENT_ID=your_document_id
LEM_THEME_ID=your_theme_id
LEM_TRANSFER_INSTRUMENT_ID=your_transfer_instrument_id

``````
Run all tests
``````
go clean -testcache
go test -v ./tests/legalentity/... -tags=integration
``````

