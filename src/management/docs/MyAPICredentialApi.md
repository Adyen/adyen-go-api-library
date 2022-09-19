# \MyAPICredentialApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMeAllowedOriginsOriginId**](MyAPICredentialApi.md#DeleteMeAllowedOriginsOriginId) | **Delete** /me/allowedOrigins/{originId} | Remove allowed origin
[**GetMe**](MyAPICredentialApi.md#GetMe) | **Get** /me | Get API credential details
[**GetMeAllowedOrigins**](MyAPICredentialApi.md#GetMeAllowedOrigins) | **Get** /me/allowedOrigins | Get allowed origins
[**GetMeAllowedOriginsOriginId**](MyAPICredentialApi.md#GetMeAllowedOriginsOriginId) | **Get** /me/allowedOrigins/{originId} | Get allowed origin details
[**PostMeAllowedOrigins**](MyAPICredentialApi.md#PostMeAllowedOrigins) | **Post** /me/allowedOrigins | Add allowed origin



## DeleteMeAllowedOriginsOriginId

> DeleteMeAllowedOriginsOriginId(ctx, originId).Execute()

Remove allowed origin



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
    originId := "originId_example" // string | Unique identifier of the allowed origin.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MyAPICredentialApi.DeleteMeAllowedOriginsOriginId(context.Background(), originId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MyAPICredentialApi.DeleteMeAllowedOriginsOriginId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**originId** | **string** | Unique identifier of the allowed origin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMeAllowedOriginsOriginIdRequest struct via the builder pattern


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


## GetMe

> MeApiCredential GetMe(ctx).Execute()

Get API credential details



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MyAPICredentialApi.GetMe(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MyAPICredentialApi.GetMe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMe`: MeApiCredential
    fmt.Fprintf(os.Stdout, "Response from `MyAPICredentialApi.GetMe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMeRequest struct via the builder pattern


### Return type

[**MeApiCredential**](MeApiCredential.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMeAllowedOrigins

> AllowedOriginsResponse GetMeAllowedOrigins(ctx).Execute()

Get allowed origins



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MyAPICredentialApi.GetMeAllowedOrigins(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MyAPICredentialApi.GetMeAllowedOrigins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMeAllowedOrigins`: AllowedOriginsResponse
    fmt.Fprintf(os.Stdout, "Response from `MyAPICredentialApi.GetMeAllowedOrigins`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMeAllowedOriginsRequest struct via the builder pattern


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


## GetMeAllowedOriginsOriginId

> AllowedOrigin GetMeAllowedOriginsOriginId(ctx, originId).Execute()

Get allowed origin details



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
    originId := "originId_example" // string | Unique identifier of the allowed origin.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MyAPICredentialApi.GetMeAllowedOriginsOriginId(context.Background(), originId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MyAPICredentialApi.GetMeAllowedOriginsOriginId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMeAllowedOriginsOriginId`: AllowedOrigin
    fmt.Fprintf(os.Stdout, "Response from `MyAPICredentialApi.GetMeAllowedOriginsOriginId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**originId** | **string** | Unique identifier of the allowed origin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMeAllowedOriginsOriginIdRequest struct via the builder pattern


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


## PostMeAllowedOrigins

> AllowedOriginsResponse PostMeAllowedOrigins(ctx).CreateAllowedOriginRequest(createAllowedOriginRequest).Execute()

Add allowed origin



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
    createAllowedOriginRequest := *openapiclient.NewCreateAllowedOriginRequest("https://adyen.com") // CreateAllowedOriginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MyAPICredentialApi.PostMeAllowedOrigins(context.Background()).CreateAllowedOriginRequest(createAllowedOriginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MyAPICredentialApi.PostMeAllowedOrigins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMeAllowedOrigins`: AllowedOriginsResponse
    fmt.Fprintf(os.Stdout, "Response from `MyAPICredentialApi.PostMeAllowedOrigins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMeAllowedOriginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAllowedOriginRequest** | [**CreateAllowedOriginRequest**](CreateAllowedOriginRequest.md) |  | 

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

