# \APICredentialsMerchantLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMerchantsMerchantIdApiCredentials**](APICredentialsMerchantLevelApi.md#GetMerchantsMerchantIdApiCredentials) | **Get** /merchants/{merchantId}/apiCredentials | Get a list of API credentials
[**GetMerchantsMerchantIdApiCredentialsApiCredentialId**](APICredentialsMerchantLevelApi.md#GetMerchantsMerchantIdApiCredentialsApiCredentialId) | **Get** /merchants/{merchantId}/apiCredentials/{apiCredentialId} | Get an API credential
[**PatchMerchantsMerchantIdApiCredentialsApiCredentialId**](APICredentialsMerchantLevelApi.md#PatchMerchantsMerchantIdApiCredentialsApiCredentialId) | **Patch** /merchants/{merchantId}/apiCredentials/{apiCredentialId} | Update an API credential
[**PostMerchantsMerchantIdApiCredentials**](APICredentialsMerchantLevelApi.md#PostMerchantsMerchantIdApiCredentials) | **Post** /merchants/{merchantId}/apiCredentials | Create an API credential



## GetMerchantsMerchantIdApiCredentials

> ListMerchantApiCredentialsResponse GetMerchantsMerchantIdApiCredentials(ctx, merchantId).PageNumber(pageNumber).PageSize(pageSize).Execute()

Get a list of API credentials



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
    resp, r, err := apiClient.APICredentialsMerchantLevelApi.GetMerchantsMerchantIdApiCredentials(context.Background(), merchantId).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APICredentialsMerchantLevelApi.GetMerchantsMerchantIdApiCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsMerchantIdApiCredentials`: ListMerchantApiCredentialsResponse
    fmt.Fprintf(os.Stdout, "Response from `APICredentialsMerchantLevelApi.GetMerchantsMerchantIdApiCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsMerchantIdApiCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | The number of the page to fetch. | 
 **pageSize** | **int32** | The number of items to have on a page, maximum 100. The default is 10 items on a page. | 

### Return type

[**ListMerchantApiCredentialsResponse**](ListMerchantApiCredentialsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMerchantsMerchantIdApiCredentialsApiCredentialId

> ApiCredential GetMerchantsMerchantIdApiCredentialsApiCredentialId(ctx, merchantId, apiCredentialId).Execute()

Get an API credential



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
    resp, r, err := apiClient.APICredentialsMerchantLevelApi.GetMerchantsMerchantIdApiCredentialsApiCredentialId(context.Background(), merchantId, apiCredentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APICredentialsMerchantLevelApi.GetMerchantsMerchantIdApiCredentialsApiCredentialId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsMerchantIdApiCredentialsApiCredentialId`: ApiCredential
    fmt.Fprintf(os.Stdout, "Response from `APICredentialsMerchantLevelApi.GetMerchantsMerchantIdApiCredentialsApiCredentialId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**apiCredentialId** | **string** | Unique identifier of the API credential. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsMerchantIdApiCredentialsApiCredentialIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiCredential**](ApiCredential.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMerchantsMerchantIdApiCredentialsApiCredentialId

> ApiCredential PatchMerchantsMerchantIdApiCredentialsApiCredentialId(ctx, merchantId, apiCredentialId).UpdateMerchantApiCredentialRequest(updateMerchantApiCredentialRequest).Execute()

Update an API credential



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
    updateMerchantApiCredentialRequest := *openapiclient.NewUpdateMerchantApiCredentialRequest() // UpdateMerchantApiCredentialRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APICredentialsMerchantLevelApi.PatchMerchantsMerchantIdApiCredentialsApiCredentialId(context.Background(), merchantId, apiCredentialId).UpdateMerchantApiCredentialRequest(updateMerchantApiCredentialRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APICredentialsMerchantLevelApi.PatchMerchantsMerchantIdApiCredentialsApiCredentialId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMerchantsMerchantIdApiCredentialsApiCredentialId`: ApiCredential
    fmt.Fprintf(os.Stdout, "Response from `APICredentialsMerchantLevelApi.PatchMerchantsMerchantIdApiCredentialsApiCredentialId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**apiCredentialId** | **string** | Unique identifier of the API credential. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMerchantsMerchantIdApiCredentialsApiCredentialIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateMerchantApiCredentialRequest** | [**UpdateMerchantApiCredentialRequest**](UpdateMerchantApiCredentialRequest.md) |  | 

### Return type

[**ApiCredential**](ApiCredential.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMerchantsMerchantIdApiCredentials

> CreateApiCredentialResponse PostMerchantsMerchantIdApiCredentials(ctx, merchantId).CreateMerchantApiCredentialRequest(createMerchantApiCredentialRequest).Execute()

Create an API credential



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
    createMerchantApiCredentialRequest := *openapiclient.NewCreateMerchantApiCredentialRequest() // CreateMerchantApiCredentialRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APICredentialsMerchantLevelApi.PostMerchantsMerchantIdApiCredentials(context.Background(), merchantId).CreateMerchantApiCredentialRequest(createMerchantApiCredentialRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APICredentialsMerchantLevelApi.PostMerchantsMerchantIdApiCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMerchantsMerchantIdApiCredentials`: CreateApiCredentialResponse
    fmt.Fprintf(os.Stdout, "Response from `APICredentialsMerchantLevelApi.PostMerchantsMerchantIdApiCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMerchantsMerchantIdApiCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createMerchantApiCredentialRequest** | [**CreateMerchantApiCredentialRequest**](CreateMerchantApiCredentialRequest.md) |  | 

### Return type

[**CreateApiCredentialResponse**](CreateApiCredentialResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

