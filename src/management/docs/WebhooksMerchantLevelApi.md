# \WebhooksMerchantLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMerchantsIdWebhooksWebhookId**](WebhooksMerchantLevelApi.md#DeleteMerchantsIdWebhooksWebhookId) | **Delete** /merchants/{id}/webhooks/{webhookId} | Remove a webhook
[**GetMerchantsIdWebhooks**](WebhooksMerchantLevelApi.md#GetMerchantsIdWebhooks) | **Get** /merchants/{id}/webhooks | List all webhooks
[**GetMerchantsIdWebhooksWebhookId**](WebhooksMerchantLevelApi.md#GetMerchantsIdWebhooksWebhookId) | **Get** /merchants/{id}/webhooks/{webhookId} | Get a webhook
[**PatchMerchantsIdWebhooksWebhookId**](WebhooksMerchantLevelApi.md#PatchMerchantsIdWebhooksWebhookId) | **Patch** /merchants/{id}/webhooks/{webhookId} | Update a webhook
[**PostMerchantsIdWebhooks**](WebhooksMerchantLevelApi.md#PostMerchantsIdWebhooks) | **Post** /merchants/{id}/webhooks | Set up a webhook
[**PostMerchantsIdWebhooksWebhookIdGenerateHmac**](WebhooksMerchantLevelApi.md#PostMerchantsIdWebhooksWebhookIdGenerateHmac) | **Post** /merchants/{id}/webhooks/{webhookId}/generateHmac | Generate an HMAC key
[**PostMerchantsIdWebhooksWebhookIdTest**](WebhooksMerchantLevelApi.md#PostMerchantsIdWebhooksWebhookIdTest) | **Post** /merchants/{id}/webhooks/{webhookId}/test | Test a webhook



## DeleteMerchantsIdWebhooksWebhookId

> DeleteMerchantsIdWebhooksWebhookId(ctx, id, webhookId).Execute()

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
    id := "id_example" // string | The unique identifier of the merchant account.
    webhookId := "webhookId_example" // string | Unique identifier of the webhook configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksMerchantLevelApi.DeleteMerchantsIdWebhooksWebhookId(context.Background(), id, webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksMerchantLevelApi.DeleteMerchantsIdWebhooksWebhookId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the merchant account. | 
**webhookId** | **string** | Unique identifier of the webhook configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMerchantsIdWebhooksWebhookIdRequest struct via the builder pattern


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


## GetMerchantsIdWebhooks

> ListWebhooksResponse GetMerchantsIdWebhooks(ctx, id).PageNumber(pageNumber).PageSize(pageSize).Execute()

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
    id := "id_example" // string | The unique identifier of the merchant account.
    pageNumber := int32(56) // int32 | The number of the page to fetch. (optional)
    pageSize := int32(56) // int32 | The number of items to have on a page, maximum 100. The default is 10 items on a page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksMerchantLevelApi.GetMerchantsIdWebhooks(context.Background(), id).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksMerchantLevelApi.GetMerchantsIdWebhooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsIdWebhooks`: ListWebhooksResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksMerchantLevelApi.GetMerchantsIdWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsIdWebhooksRequest struct via the builder pattern


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


## GetMerchantsIdWebhooksWebhookId

> Webhook GetMerchantsIdWebhooksWebhookId(ctx, id, webhookId).Execute()

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
    id := "id_example" // string | The unique identifier of the merchant account.
    webhookId := "webhookId_example" // string | Unique identifier of the webhook configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksMerchantLevelApi.GetMerchantsIdWebhooksWebhookId(context.Background(), id, webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksMerchantLevelApi.GetMerchantsIdWebhooksWebhookId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsIdWebhooksWebhookId`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhooksMerchantLevelApi.GetMerchantsIdWebhooksWebhookId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the merchant account. | 
**webhookId** | **string** | Unique identifier of the webhook configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsIdWebhooksWebhookIdRequest struct via the builder pattern


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


## PatchMerchantsIdWebhooksWebhookId

> Webhook PatchMerchantsIdWebhooksWebhookId(ctx, id, webhookId).UpdateMerchantWebhookRequest(updateMerchantWebhookRequest).Execute()

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
    id := "id_example" // string | The unique identifier of the merchant account.
    webhookId := "webhookId_example" // string | Unique identifier of the webhook configuration.
    updateMerchantWebhookRequest := *openapiclient.NewUpdateMerchantWebhookRequest(false, "SOAP", "http://www.adyen.com") // UpdateMerchantWebhookRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksMerchantLevelApi.PatchMerchantsIdWebhooksWebhookId(context.Background(), id, webhookId).UpdateMerchantWebhookRequest(updateMerchantWebhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksMerchantLevelApi.PatchMerchantsIdWebhooksWebhookId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMerchantsIdWebhooksWebhookId`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhooksMerchantLevelApi.PatchMerchantsIdWebhooksWebhookId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the merchant account. | 
**webhookId** | **string** | Unique identifier of the webhook configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMerchantsIdWebhooksWebhookIdRequest struct via the builder pattern


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


## PostMerchantsIdWebhooks

> Webhook PostMerchantsIdWebhooks(ctx, id).CreateMerchantWebhookRequest(createMerchantWebhookRequest).Execute()

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
    id := "id_example" // string | The unique identifier of the merchant account.
    createMerchantWebhookRequest := *openapiclient.NewCreateMerchantWebhookRequest(false, "SOAP", "Type_example", "http://www.adyen.com") // CreateMerchantWebhookRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksMerchantLevelApi.PostMerchantsIdWebhooks(context.Background(), id).CreateMerchantWebhookRequest(createMerchantWebhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksMerchantLevelApi.PostMerchantsIdWebhooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMerchantsIdWebhooks`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhooksMerchantLevelApi.PostMerchantsIdWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMerchantsIdWebhooksRequest struct via the builder pattern


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


## PostMerchantsIdWebhooksWebhookIdGenerateHmac

> GenerateHmacKeyResponse PostMerchantsIdWebhooksWebhookIdGenerateHmac(ctx, id, webhookId).Execute()

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
    id := "id_example" // string | The unique identifier of the merchant account.
    webhookId := "webhookId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksMerchantLevelApi.PostMerchantsIdWebhooksWebhookIdGenerateHmac(context.Background(), id, webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksMerchantLevelApi.PostMerchantsIdWebhooksWebhookIdGenerateHmac``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMerchantsIdWebhooksWebhookIdGenerateHmac`: GenerateHmacKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksMerchantLevelApi.PostMerchantsIdWebhooksWebhookIdGenerateHmac`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the merchant account. | 
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMerchantsIdWebhooksWebhookIdGenerateHmacRequest struct via the builder pattern


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


## PostMerchantsIdWebhooksWebhookIdTest

> TestWebhookResponse PostMerchantsIdWebhooksWebhookIdTest(ctx, id, webhookId).TestWebhookRequest(testWebhookRequest).Execute()

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
    id := "id_example" // string | The unique identifier of the merchant account.
    webhookId := "webhookId_example" // string | Unique identifier of the webhook configuration.
    testWebhookRequest := *openapiclient.NewTestWebhookRequest() // TestWebhookRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksMerchantLevelApi.PostMerchantsIdWebhooksWebhookIdTest(context.Background(), id, webhookId).TestWebhookRequest(testWebhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksMerchantLevelApi.PostMerchantsIdWebhooksWebhookIdTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMerchantsIdWebhooksWebhookIdTest`: TestWebhookResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksMerchantLevelApi.PostMerchantsIdWebhooksWebhookIdTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the merchant account. | 
**webhookId** | **string** | Unique identifier of the webhook configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMerchantsIdWebhooksWebhookIdTestRequest struct via the builder pattern


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

