# \WebhooksMerchantLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMerchantsMerchantIdWebhooksWebhookId**](WebhooksMerchantLevelApi.md#DeleteMerchantsMerchantIdWebhooksWebhookId) | **Delete** /merchants/{merchantId}/webhooks/{webhookId} | Remove a webhook
[**GetMerchantsMerchantIdWebhooks**](WebhooksMerchantLevelApi.md#GetMerchantsMerchantIdWebhooks) | **Get** /merchants/{merchantId}/webhooks | List all webhooks
[**GetMerchantsMerchantIdWebhooksWebhookId**](WebhooksMerchantLevelApi.md#GetMerchantsMerchantIdWebhooksWebhookId) | **Get** /merchants/{merchantId}/webhooks/{webhookId} | Get a webhook
[**PatchMerchantsMerchantIdWebhooksWebhookId**](WebhooksMerchantLevelApi.md#PatchMerchantsMerchantIdWebhooksWebhookId) | **Patch** /merchants/{merchantId}/webhooks/{webhookId} | Update a webhook
[**PostMerchantsMerchantIdWebhooks**](WebhooksMerchantLevelApi.md#PostMerchantsMerchantIdWebhooks) | **Post** /merchants/{merchantId}/webhooks | Set up a webhook
[**PostMerchantsMerchantIdWebhooksWebhookIdGenerateHmac**](WebhooksMerchantLevelApi.md#PostMerchantsMerchantIdWebhooksWebhookIdGenerateHmac) | **Post** /merchants/{merchantId}/webhooks/{webhookId}/generateHmac | Generate an HMAC key
[**PostMerchantsMerchantIdWebhooksWebhookIdTest**](WebhooksMerchantLevelApi.md#PostMerchantsMerchantIdWebhooksWebhookIdTest) | **Post** /merchants/{merchantId}/webhooks/{webhookId}/test | Test a webhook



## DeleteMerchantsMerchantIdWebhooksWebhookId

> DeleteMerchantsMerchantIdWebhooksWebhookId(ctx, merchantId, webhookId).Execute()

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
    merchantId := "merchantId_example" // string | The unique identifier of the merchant account.
    webhookId := "webhookId_example" // string | Unique identifier of the webhook configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksMerchantLevelApi.DeleteMerchantsMerchantIdWebhooksWebhookId(context.Background(), merchantId, webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksMerchantLevelApi.DeleteMerchantsMerchantIdWebhooksWebhookId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**webhookId** | **string** | Unique identifier of the webhook configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMerchantsMerchantIdWebhooksWebhookIdRequest struct via the builder pattern


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


## GetMerchantsMerchantIdWebhooks

> ListWebhooksResponse GetMerchantsMerchantIdWebhooks(ctx, merchantId).PageNumber(pageNumber).PageSize(pageSize).Execute()

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
    merchantId := "merchantId_example" // string | The unique identifier of the merchant account.
    pageNumber := int32(56) // int32 | The number of the page to fetch. (optional)
    pageSize := int32(56) // int32 | The number of items to have on a page, maximum 100. The default is 10 items on a page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksMerchantLevelApi.GetMerchantsMerchantIdWebhooks(context.Background(), merchantId).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksMerchantLevelApi.GetMerchantsMerchantIdWebhooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsMerchantIdWebhooks`: ListWebhooksResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksMerchantLevelApi.GetMerchantsMerchantIdWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsMerchantIdWebhooksRequest struct via the builder pattern


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


## GetMerchantsMerchantIdWebhooksWebhookId

> Webhook GetMerchantsMerchantIdWebhooksWebhookId(ctx, merchantId, webhookId).Execute()

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
    merchantId := "merchantId_example" // string | The unique identifier of the merchant account.
    webhookId := "webhookId_example" // string | Unique identifier of the webhook configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksMerchantLevelApi.GetMerchantsMerchantIdWebhooksWebhookId(context.Background(), merchantId, webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksMerchantLevelApi.GetMerchantsMerchantIdWebhooksWebhookId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsMerchantIdWebhooksWebhookId`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhooksMerchantLevelApi.GetMerchantsMerchantIdWebhooksWebhookId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**webhookId** | **string** | Unique identifier of the webhook configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsMerchantIdWebhooksWebhookIdRequest struct via the builder pattern


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


## PatchMerchantsMerchantIdWebhooksWebhookId

> Webhook PatchMerchantsMerchantIdWebhooksWebhookId(ctx, merchantId, webhookId).UpdateMerchantWebhookRequest(updateMerchantWebhookRequest).Execute()

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
    merchantId := "merchantId_example" // string | The unique identifier of the merchant account.
    webhookId := "webhookId_example" // string | Unique identifier of the webhook configuration.
    updateMerchantWebhookRequest := *openapiclient.NewUpdateMerchantWebhookRequest(false, "SOAP", "http://www.adyen.com") // UpdateMerchantWebhookRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksMerchantLevelApi.PatchMerchantsMerchantIdWebhooksWebhookId(context.Background(), merchantId, webhookId).UpdateMerchantWebhookRequest(updateMerchantWebhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksMerchantLevelApi.PatchMerchantsMerchantIdWebhooksWebhookId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMerchantsMerchantIdWebhooksWebhookId`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhooksMerchantLevelApi.PatchMerchantsMerchantIdWebhooksWebhookId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**webhookId** | **string** | Unique identifier of the webhook configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMerchantsMerchantIdWebhooksWebhookIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateMerchantWebhookRequest** | [**UpdateMerchantWebhookRequest**](UpdateMerchantWebhookRequest.md) |  | 

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


## PostMerchantsMerchantIdWebhooks

> Webhook PostMerchantsMerchantIdWebhooks(ctx, merchantId).CreateMerchantWebhookRequest(createMerchantWebhookRequest).Execute()

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
    merchantId := "merchantId_example" // string | The unique identifier of the merchant account.
    createMerchantWebhookRequest := *openapiclient.NewCreateMerchantWebhookRequest(false, "SOAP", "Type_example", "http://www.adyen.com") // CreateMerchantWebhookRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksMerchantLevelApi.PostMerchantsMerchantIdWebhooks(context.Background(), merchantId).CreateMerchantWebhookRequest(createMerchantWebhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksMerchantLevelApi.PostMerchantsMerchantIdWebhooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMerchantsMerchantIdWebhooks`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhooksMerchantLevelApi.PostMerchantsMerchantIdWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMerchantsMerchantIdWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createMerchantWebhookRequest** | [**CreateMerchantWebhookRequest**](CreateMerchantWebhookRequest.md) |  | 

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


## PostMerchantsMerchantIdWebhooksWebhookIdGenerateHmac

> GenerateHmacKeyResponse PostMerchantsMerchantIdWebhooksWebhookIdGenerateHmac(ctx, merchantId, webhookId).Execute()

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
    merchantId := "merchantId_example" // string | The unique identifier of the merchant account.
    webhookId := "webhookId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksMerchantLevelApi.PostMerchantsMerchantIdWebhooksWebhookIdGenerateHmac(context.Background(), merchantId, webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksMerchantLevelApi.PostMerchantsMerchantIdWebhooksWebhookIdGenerateHmac``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMerchantsMerchantIdWebhooksWebhookIdGenerateHmac`: GenerateHmacKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksMerchantLevelApi.PostMerchantsMerchantIdWebhooksWebhookIdGenerateHmac`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMerchantsMerchantIdWebhooksWebhookIdGenerateHmacRequest struct via the builder pattern


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


## PostMerchantsMerchantIdWebhooksWebhookIdTest

> TestWebhookResponse PostMerchantsMerchantIdWebhooksWebhookIdTest(ctx, merchantId, webhookId).TestWebhookRequest(testWebhookRequest).Execute()

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
    merchantId := "merchantId_example" // string | The unique identifier of the merchant account.
    webhookId := "webhookId_example" // string | Unique identifier of the webhook configuration.
    testWebhookRequest := *openapiclient.NewTestWebhookRequest() // TestWebhookRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksMerchantLevelApi.PostMerchantsMerchantIdWebhooksWebhookIdTest(context.Background(), merchantId, webhookId).TestWebhookRequest(testWebhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksMerchantLevelApi.PostMerchantsMerchantIdWebhooksWebhookIdTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMerchantsMerchantIdWebhooksWebhookIdTest`: TestWebhookResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksMerchantLevelApi.PostMerchantsMerchantIdWebhooksWebhookIdTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**webhookId** | **string** | Unique identifier of the webhook configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMerchantsMerchantIdWebhooksWebhookIdTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **testWebhookRequest** | [**TestWebhookRequest**](TestWebhookRequest.md) |  | 

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

