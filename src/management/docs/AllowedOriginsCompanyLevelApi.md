# \AllowedOriginsCompanyLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOriginsOriginId**](AllowedOriginsCompanyLevelApi.md#DeleteCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOriginsOriginId) | **Delete** /companies/{companyId}/apiCredentials/{apiCredentialId}/allowedOrigins/{originId} | Delete an allowed origin
[**GetCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOrigins**](AllowedOriginsCompanyLevelApi.md#GetCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOrigins) | **Get** /companies/{companyId}/apiCredentials/{apiCredentialId}/allowedOrigins | Get a list of allowed origins
[**GetCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOriginsOriginId**](AllowedOriginsCompanyLevelApi.md#GetCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOriginsOriginId) | **Get** /companies/{companyId}/apiCredentials/{apiCredentialId}/allowedOrigins/{originId} | Get an allowed origin
[**PostCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOrigins**](AllowedOriginsCompanyLevelApi.md#PostCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOrigins) | **Post** /companies/{companyId}/apiCredentials/{apiCredentialId}/allowedOrigins | Create an allowed origin



## DeleteCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOriginsOriginId

> DeleteCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOriginsOriginId(ctx, companyId, apiCredentialId, originId).Execute()

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
    companyId := "companyId_example" // string | The unique identifier of the company account.
    apiCredentialId := "apiCredentialId_example" // string | Unique identifier of the API credential.
    originId := "originId_example" // string | Unique identifier of the allowed origin.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllowedOriginsCompanyLevelApi.DeleteCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOriginsOriginId(context.Background(), companyId, apiCredentialId, originId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllowedOriginsCompanyLevelApi.DeleteCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOriginsOriginId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 
**apiCredentialId** | **string** | Unique identifier of the API credential. | 
**originId** | **string** | Unique identifier of the allowed origin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOriginsOriginIdRequest struct via the builder pattern


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


## GetCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOrigins

> AllowedOriginsResponse GetCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOrigins(ctx, companyId, apiCredentialId).Execute()

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
    companyId := "companyId_example" // string | The unique identifier of the company account.
    apiCredentialId := "apiCredentialId_example" // string | Unique identifier of the API credential.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllowedOriginsCompanyLevelApi.GetCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOrigins(context.Background(), companyId, apiCredentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllowedOriginsCompanyLevelApi.GetCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOrigins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOrigins`: AllowedOriginsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllowedOriginsCompanyLevelApi.GetCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOrigins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 
**apiCredentialId** | **string** | Unique identifier of the API credential. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOriginsRequest struct via the builder pattern


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


## GetCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOriginsOriginId

> AllowedOrigin GetCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOriginsOriginId(ctx, companyId, apiCredentialId, originId).Execute()

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
    companyId := "companyId_example" // string | The unique identifier of the company account.
    apiCredentialId := "apiCredentialId_example" // string | Unique identifier of the API credential.
    originId := "originId_example" // string | Unique identifier of the allowed origin.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllowedOriginsCompanyLevelApi.GetCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOriginsOriginId(context.Background(), companyId, apiCredentialId, originId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllowedOriginsCompanyLevelApi.GetCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOriginsOriginId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOriginsOriginId`: AllowedOrigin
    fmt.Fprintf(os.Stdout, "Response from `AllowedOriginsCompanyLevelApi.GetCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOriginsOriginId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 
**apiCredentialId** | **string** | Unique identifier of the API credential. | 
**originId** | **string** | Unique identifier of the allowed origin. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOriginsOriginIdRequest struct via the builder pattern


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


## PostCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOrigins

> AllowedOriginsResponse PostCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOrigins(ctx, companyId, apiCredentialId).AllowedOrigin(allowedOrigin).Execute()

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
    companyId := "companyId_example" // string | The unique identifier of the company account.
    apiCredentialId := "apiCredentialId_example" // string | Unique identifier of the API credential.
    allowedOrigin := *openapiclient.NewAllowedOrigin("https://adyen.com") // AllowedOrigin |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllowedOriginsCompanyLevelApi.PostCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOrigins(context.Background(), companyId, apiCredentialId).AllowedOrigin(allowedOrigin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllowedOriginsCompanyLevelApi.PostCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOrigins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOrigins`: AllowedOriginsResponse
    fmt.Fprintf(os.Stdout, "Response from `AllowedOriginsCompanyLevelApi.PostCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOrigins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 
**apiCredentialId** | **string** | Unique identifier of the API credential. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompaniesCompanyIdApiCredentialsApiCredentialIdAllowedOriginsRequest struct via the builder pattern


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

