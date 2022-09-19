# \AllowedOriginsMerchantLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMerchantsIdApiCredentialsApiCredentialIdAllowedOriginsOriginId**](AllowedOriginsMerchantLevelApi.md#DeleteMerchantsIdApiCredentialsApiCredentialIdAllowedOriginsOriginId) | **Delete** /merchants/{id}/apiCredentials/{apiCredentialId}/allowedOrigins/{originId} | Delete an allowed origin
[**GetMerchantsIdApiCredentialsApiCredentialIdAllowedOrigins**](AllowedOriginsMerchantLevelApi.md#GetMerchantsIdApiCredentialsApiCredentialIdAllowedOrigins) | **Get** /merchants/{id}/apiCredentials/{apiCredentialId}/allowedOrigins | Get a list of allowed origins
[**GetMerchantsIdApiCredentialsApiCredentialIdAllowedOriginsOriginId**](AllowedOriginsMerchantLevelApi.md#GetMerchantsIdApiCredentialsApiCredentialIdAllowedOriginsOriginId) | **Get** /merchants/{id}/apiCredentials/{apiCredentialId}/allowedOrigins/{originId} | Get an allowed origin
[**PostMerchantsIdApiCredentialsApiCredentialIdAllowedOrigins**](AllowedOriginsMerchantLevelApi.md#PostMerchantsIdApiCredentialsApiCredentialIdAllowedOrigins) | **Post** /merchants/{id}/apiCredentials/{apiCredentialId}/allowedOrigins | Create an allowed origin



## DeleteMerchantsIdApiCredentialsApiCredentialIdAllowedOriginsOriginId

> DeleteMerchantsIdApiCredentialsApiCredentialIdAllowedOriginsOriginId(ctx, id, apiCredentialId, originId).Execute()

Delete an allowed origin



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
    apiCredentialId := "apiCredentialId_example" // string | Unique identifier of the API credential.
    originId := "originId_example" // string | Unique identifier of the allowed origin.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllowedOriginsMerchantLevelApi.DeleteMerchantsIdApiCredentialsApiCredentialIdAllowedOriginsOriginId(context.Background(), id, apiCredentialId, originId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllowedOriginsMerchantLevelApi.DeleteMerchantsIdApiCredentialsApiCredentialIdAllowedOriginsOriginId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the merchant account. | 
**apiCredentialId** | **string** | Unique identifier of the API credential. | 
**originId** | **string** | Unique identifier of the allowed origin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMerchantsIdApiCredentialsApiCredentialIdAllowedOriginsOriginIdRequest struct via the builder pattern


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


## GetMerchantsIdApiCredentialsApiCredentialIdAllowedOrigins

> AllowedOriginsResponse GetMerchantsIdApiCredentialsApiCredentialIdAllowedOrigins(ctx, id, apiCredentialId).Execute()

Get a list of allowed origins



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
    apiCredentialId := "apiCredentialId_example" // string | Unique identifier of the API credential.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllowedOriginsMerchantLevelApi.GetMerchantsIdApiCredentialsApiCredentialIdAllowedOrigins(context.Background(), id, apiCredentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllowedOriginsMerchantLevelApi.GetMerchantsIdApiCredentialsApiCredentialIdAllowedOrigins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsIdApiCredentialsApiCredentialIdAllowedOrigins`: AllowedOriginsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllowedOriginsMerchantLevelApi.GetMerchantsIdApiCredentialsApiCredentialIdAllowedOrigins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the merchant account. | 
**apiCredentialId** | **string** | Unique identifier of the API credential. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsIdApiCredentialsApiCredentialIdAllowedOriginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AllowedOriginsResponse**](AllowedOriginsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMerchantsIdApiCredentialsApiCredentialIdAllowedOriginsOriginId

> AllowedOrigin GetMerchantsIdApiCredentialsApiCredentialIdAllowedOriginsOriginId(ctx, id, apiCredentialId, originId).Execute()

Get an allowed origin



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
    apiCredentialId := "apiCredentialId_example" // string | Unique identifier of the API credential.
    originId := "originId_example" // string | Unique identifier of the allowed origin.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllowedOriginsMerchantLevelApi.GetMerchantsIdApiCredentialsApiCredentialIdAllowedOriginsOriginId(context.Background(), id, apiCredentialId, originId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllowedOriginsMerchantLevelApi.GetMerchantsIdApiCredentialsApiCredentialIdAllowedOriginsOriginId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsIdApiCredentialsApiCredentialIdAllowedOriginsOriginId`: AllowedOrigin
    fmt.Fprintf(os.Stdout, "Response from `AllowedOriginsMerchantLevelApi.GetMerchantsIdApiCredentialsApiCredentialIdAllowedOriginsOriginId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the merchant account. | 
**apiCredentialId** | **string** | Unique identifier of the API credential. | 
**originId** | **string** | Unique identifier of the allowed origin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsIdApiCredentialsApiCredentialIdAllowedOriginsOriginIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AllowedOrigin**](AllowedOrigin.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMerchantsIdApiCredentialsApiCredentialIdAllowedOrigins

> AllowedOriginsResponse PostMerchantsIdApiCredentialsApiCredentialIdAllowedOrigins(ctx, id, apiCredentialId).AllowedOrigin(allowedOrigin).Execute()

Create an allowed origin



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
    apiCredentialId := "apiCredentialId_example" // string | Unique identifier of the API credential.
    allowedOrigin := *openapiclient.NewAllowedOrigin("https://adyen.com") // AllowedOrigin |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllowedOriginsMerchantLevelApi.PostMerchantsIdApiCredentialsApiCredentialIdAllowedOrigins(context.Background(), id, apiCredentialId).AllowedOrigin(allowedOrigin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllowedOriginsMerchantLevelApi.PostMerchantsIdApiCredentialsApiCredentialIdAllowedOrigins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMerchantsIdApiCredentialsApiCredentialIdAllowedOrigins`: AllowedOriginsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllowedOriginsMerchantLevelApi.PostMerchantsIdApiCredentialsApiCredentialIdAllowedOrigins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the merchant account. | 
**apiCredentialId** | **string** | Unique identifier of the API credential. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMerchantsIdApiCredentialsApiCredentialIdAllowedOriginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **allowedOrigin** | [**AllowedOrigin**](AllowedOrigin.md) |  | 

### Return type

[**AllowedOriginsResponse**](AllowedOriginsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

