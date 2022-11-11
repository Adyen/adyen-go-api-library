# \WebhooksCompanyLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompaniesCompanyIdWebhooksWebhookId**](WebhooksCompanyLevelApi.md#DeleteCompaniesCompanyIdWebhooksWebhookId) | **Delete** /companies/{companyId}/webhooks/{webhookId} | Remove a webhook
[**GetCompaniesCompanyIdWebhooks**](WebhooksCompanyLevelApi.md#GetCompaniesCompanyIdWebhooks) | **Get** /companies/{companyId}/webhooks | List all webhooks
[**GetCompaniesCompanyIdWebhooksWebhookId**](WebhooksCompanyLevelApi.md#GetCompaniesCompanyIdWebhooksWebhookId) | **Get** /companies/{companyId}/webhooks/{webhookId} | Get a webhook
[**PatchCompaniesCompanyIdWebhooksWebhookId**](WebhooksCompanyLevelApi.md#PatchCompaniesCompanyIdWebhooksWebhookId) | **Patch** /companies/{companyId}/webhooks/{webhookId} | Update a webhook
[**PostCompaniesCompanyIdWebhooks**](WebhooksCompanyLevelApi.md#PostCompaniesCompanyIdWebhooks) | **Post** /companies/{companyId}/webhooks | Set up a webhook
[**PostCompaniesCompanyIdWebhooksWebhookIdGenerateHmac**](WebhooksCompanyLevelApi.md#PostCompaniesCompanyIdWebhooksWebhookIdGenerateHmac) | **Post** /companies/{companyId}/webhooks/{webhookId}/generateHmac | Generate an HMAC key
[**PostCompaniesCompanyIdWebhooksWebhookIdTest**](WebhooksCompanyLevelApi.md#PostCompaniesCompanyIdWebhooksWebhookIdTest) | **Post** /companies/{companyId}/webhooks/{webhookId}/test | Test a webhook



## DeleteCompaniesCompanyIdWebhooksWebhookId

> DeleteCompaniesCompanyIdWebhooksWebhookId(ctx, companyId, webhookId).Execute()

Remove a webhook



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := "companyId_example" // string | The unique identifier of the company account.
    webhookId := "webhookId_example" // string | Unique identifier of the webhook configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksCompanyLevelApi.DeleteCompaniesCompanyIdWebhooksWebhookId(context.Background(), companyId, webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksCompanyLevelApi.DeleteCompaniesCompanyIdWebhooksWebhookId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 
**webhookId** | **string** | Unique identifier of the webhook configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompaniesCompanyIdWebhooksWebhookIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompaniesCompanyIdWebhooks

> ListWebhooksResponse GetCompaniesCompanyIdWebhooks(ctx, companyId).PageNumber(pageNumber).PageSize(pageSize).Execute()

List all webhooks



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := "companyId_example" // string | Unique identifier of the [company account](https://docs.adyen.com/account/account-structure#company-account).
    pageNumber := int32(56) // int32 | The number of the page to fetch. (optional)
    pageSize := int32(56) // int32 | The number of items to have on a page, maximum 100. The default is 10 items on a page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksCompanyLevelApi.GetCompaniesCompanyIdWebhooks(context.Background(), companyId).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksCompanyLevelApi.GetCompaniesCompanyIdWebhooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompaniesCompanyIdWebhooks`: ListWebhooksResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksCompanyLevelApi.GetCompaniesCompanyIdWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier of the [company account](https://docs.adyen.com/account/account-structure#company-account). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesCompanyIdWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | The number of the page to fetch. | 
 **pageSize** | **int32** | The number of items to have on a page, maximum 100. The default is 10 items on a page. | 

### Return type

[**ListWebhooksResponse**](ListWebhooksResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompaniesCompanyIdWebhooksWebhookId

> Webhook GetCompaniesCompanyIdWebhooksWebhookId(ctx, companyId, webhookId).Execute()

Get a webhook



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := "companyId_example" // string | Unique identifier of the [company account](https://docs.adyen.com/account/account-structure#company-account).
    webhookId := "webhookId_example" // string | Unique identifier of the webhook configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksCompanyLevelApi.GetCompaniesCompanyIdWebhooksWebhookId(context.Background(), companyId, webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksCompanyLevelApi.GetCompaniesCompanyIdWebhooksWebhookId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompaniesCompanyIdWebhooksWebhookId`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhooksCompanyLevelApi.GetCompaniesCompanyIdWebhooksWebhookId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier of the [company account](https://docs.adyen.com/account/account-structure#company-account). | 
**webhookId** | **string** | Unique identifier of the webhook configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesCompanyIdWebhooksWebhookIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Webhook**](Webhook.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCompaniesCompanyIdWebhooksWebhookId

> Webhook PatchCompaniesCompanyIdWebhooksWebhookId(ctx, companyId, webhookId).UpdateCompanyWebhookRequest(updateCompanyWebhookRequest).Execute()

Update a webhook



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := "companyId_example" // string | The unique identifier of the company account.
    webhookId := "webhookId_example" // string | Unique identifier of the webhook configuration.
    updateCompanyWebhookRequest := *openapiclient.NewUpdateCompanyWebhookRequest("http://www.adyen.com") // UpdateCompanyWebhookRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksCompanyLevelApi.PatchCompaniesCompanyIdWebhooksWebhookId(context.Background(), companyId, webhookId).UpdateCompanyWebhookRequest(updateCompanyWebhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksCompanyLevelApi.PatchCompaniesCompanyIdWebhooksWebhookId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCompaniesCompanyIdWebhooksWebhookId`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhooksCompanyLevelApi.PatchCompaniesCompanyIdWebhooksWebhookId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 
**webhookId** | **string** | Unique identifier of the webhook configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompaniesCompanyIdWebhooksWebhookIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateCompanyWebhookRequest** | [**UpdateCompanyWebhookRequest**](UpdateCompanyWebhookRequest.md) |  | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompaniesCompanyIdWebhooks

> Webhook PostCompaniesCompanyIdWebhooks(ctx, companyId).CreateCompanyWebhookRequest(createCompanyWebhookRequest).Execute()

Set up a webhook



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := "companyId_example" // string | Unique identifier of the [company account](https://docs.adyen.com/account/account-structure#company-account).
    createCompanyWebhookRequest := *openapiclient.NewCreateCompanyWebhookRequest(false, "SOAP", "FilterMerchantAccountType_example", []string{"FilterMerchantAccounts_example"}, "Type_example", "http://www.adyen.com") // CreateCompanyWebhookRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksCompanyLevelApi.PostCompaniesCompanyIdWebhooks(context.Background(), companyId).CreateCompanyWebhookRequest(createCompanyWebhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksCompanyLevelApi.PostCompaniesCompanyIdWebhooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCompaniesCompanyIdWebhooks`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhooksCompanyLevelApi.PostCompaniesCompanyIdWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier of the [company account](https://docs.adyen.com/account/account-structure#company-account). | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompaniesCompanyIdWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createCompanyWebhookRequest** | [**CreateCompanyWebhookRequest**](CreateCompanyWebhookRequest.md) |  | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompaniesCompanyIdWebhooksWebhookIdGenerateHmac

> GenerateHmacKeyResponse PostCompaniesCompanyIdWebhooksWebhookIdGenerateHmac(ctx, companyId, webhookId).Execute()

Generate an HMAC key



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := "companyId_example" // string | The unique identifier of the company account.
    webhookId := "webhookId_example" // string | Unique identifier of the webhook configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksCompanyLevelApi.PostCompaniesCompanyIdWebhooksWebhookIdGenerateHmac(context.Background(), companyId, webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksCompanyLevelApi.PostCompaniesCompanyIdWebhooksWebhookIdGenerateHmac``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCompaniesCompanyIdWebhooksWebhookIdGenerateHmac`: GenerateHmacKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksCompanyLevelApi.PostCompaniesCompanyIdWebhooksWebhookIdGenerateHmac`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 
**webhookId** | **string** | Unique identifier of the webhook configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompaniesCompanyIdWebhooksWebhookIdGenerateHmacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GenerateHmacKeyResponse**](GenerateHmacKeyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompaniesCompanyIdWebhooksWebhookIdTest

> TestWebhookResponse PostCompaniesCompanyIdWebhooksWebhookIdTest(ctx, companyId, webhookId).TestCompanyWebhookRequest(testCompanyWebhookRequest).Execute()

Test a webhook



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := "companyId_example" // string | The unique identifier of the company account.
    webhookId := "webhookId_example" // string | Unique identifier of the webhook configuration.
    testCompanyWebhookRequest := *openapiclient.NewTestCompanyWebhookRequest() // TestCompanyWebhookRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksCompanyLevelApi.PostCompaniesCompanyIdWebhooksWebhookIdTest(context.Background(), companyId, webhookId).TestCompanyWebhookRequest(testCompanyWebhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksCompanyLevelApi.PostCompaniesCompanyIdWebhooksWebhookIdTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCompaniesCompanyIdWebhooksWebhookIdTest`: TestWebhookResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksCompanyLevelApi.PostCompaniesCompanyIdWebhooksWebhookIdTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 
**webhookId** | **string** | Unique identifier of the webhook configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompaniesCompanyIdWebhooksWebhookIdTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **testCompanyWebhookRequest** | [**TestCompanyWebhookRequest**](TestCompanyWebhookRequest.md) |  | 

### Return type

[**TestWebhookResponse**](TestWebhookResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

