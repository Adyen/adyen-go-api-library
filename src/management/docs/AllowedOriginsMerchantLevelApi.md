# \AllowedOriginsMerchantLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOriginsOriginId**](AllowedOriginsMerchantLevelApi.md#DeleteMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOriginsOriginId) | **Delete** /merchants/{merchantId}/apiCredentials/{apiCredentialId}/allowedOrigins/{originId} | Delete an allowed origin
[**GetMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOrigins**](AllowedOriginsMerchantLevelApi.md#GetMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOrigins) | **Get** /merchants/{merchantId}/apiCredentials/{apiCredentialId}/allowedOrigins | Get a list of allowed origins
[**GetMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOriginsOriginId**](AllowedOriginsMerchantLevelApi.md#GetMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOriginsOriginId) | **Get** /merchants/{merchantId}/apiCredentials/{apiCredentialId}/allowedOrigins/{originId} | Get an allowed origin
[**PostMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOrigins**](AllowedOriginsMerchantLevelApi.md#PostMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOrigins) | **Post** /merchants/{merchantId}/apiCredentials/{apiCredentialId}/allowedOrigins | Create an allowed origin



## DeleteMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOriginsOriginId

> DeleteMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOriginsOriginId(ctx, merchantId, apiCredentialId, originId).Execute()

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
    merchantId := "merchantId_example" // string | The unique identifier of the merchant account.
    apiCredentialId := "apiCredentialId_example" // string | Unique identifier of the API credential.
    originId := "originId_example" // string | Unique identifier of the allowed origin.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllowedOriginsMerchantLevelApi.DeleteMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOriginsOriginId(context.Background(), merchantId, apiCredentialId, originId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllowedOriginsMerchantLevelApi.DeleteMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOriginsOriginId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**apiCredentialId** | **string** | Unique identifier of the API credential. | 
**originId** | **string** | Unique identifier of the allowed origin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOriginsOriginIdRequest struct via the builder pattern


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


## GetMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOrigins

> AllowedOriginsResponse GetMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOrigins(ctx, merchantId, apiCredentialId).Execute()

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
    merchantId := "merchantId_example" // string | The unique identifier of the merchant account.
    apiCredentialId := "apiCredentialId_example" // string | Unique identifier of the API credential.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllowedOriginsMerchantLevelApi.GetMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOrigins(context.Background(), merchantId, apiCredentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllowedOriginsMerchantLevelApi.GetMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOrigins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOrigins`: AllowedOriginsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllowedOriginsMerchantLevelApi.GetMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOrigins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**apiCredentialId** | **string** | Unique identifier of the API credential. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOriginsRequest struct via the builder pattern


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


## GetMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOriginsOriginId

> AllowedOrigin GetMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOriginsOriginId(ctx, merchantId, apiCredentialId, originId).Execute()

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
    merchantId := "merchantId_example" // string | The unique identifier of the merchant account.
    apiCredentialId := "apiCredentialId_example" // string | Unique identifier of the API credential.
    originId := "originId_example" // string | Unique identifier of the allowed origin.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllowedOriginsMerchantLevelApi.GetMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOriginsOriginId(context.Background(), merchantId, apiCredentialId, originId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllowedOriginsMerchantLevelApi.GetMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOriginsOriginId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOriginsOriginId`: AllowedOrigin
    fmt.Fprintf(os.Stdout, "Response from `AllowedOriginsMerchantLevelApi.GetMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOriginsOriginId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**apiCredentialId** | **string** | Unique identifier of the API credential. | 
**originId** | **string** | Unique identifier of the allowed origin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOriginsOriginIdRequest struct via the builder pattern


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


## PostMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOrigins

> AllowedOriginsResponse PostMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOrigins(ctx, merchantId, apiCredentialId).AllowedOrigin(allowedOrigin).Execute()

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
    merchantId := "merchantId_example" // string | The unique identifier of the merchant account.
    apiCredentialId := "apiCredentialId_example" // string | Unique identifier of the API credential.
    allowedOrigin := *openapiclient.NewAllowedOrigin("https://adyen.com") // AllowedOrigin |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllowedOriginsMerchantLevelApi.PostMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOrigins(context.Background(), merchantId, apiCredentialId).AllowedOrigin(allowedOrigin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllowedOriginsMerchantLevelApi.PostMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOrigins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOrigins`: AllowedOriginsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllowedOriginsMerchantLevelApi.PostMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOrigins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**apiCredentialId** | **string** | Unique identifier of the API credential. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMerchantsMerchantIdApiCredentialsApiCredentialIdAllowedOriginsRequest struct via the builder pattern


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

