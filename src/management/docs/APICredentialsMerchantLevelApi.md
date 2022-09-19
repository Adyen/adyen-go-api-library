# \APICredentialsMerchantLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMerchantsIdApiCredentials**](APICredentialsMerchantLevelApi.md#GetMerchantsIdApiCredentials) | **Get** /merchants/{id}/apiCredentials | Get a list of API credentials
[**GetMerchantsIdApiCredentialsApiCredentialId**](APICredentialsMerchantLevelApi.md#GetMerchantsIdApiCredentialsApiCredentialId) | **Get** /merchants/{id}/apiCredentials/{apiCredentialId} | Get an API credential
[**PatchMerchantsIdApiCredentialsApiCredentialId**](APICredentialsMerchantLevelApi.md#PatchMerchantsIdApiCredentialsApiCredentialId) | **Patch** /merchants/{id}/apiCredentials/{apiCredentialId} | Update an API credential
[**PostMerchantsIdApiCredentials**](APICredentialsMerchantLevelApi.md#PostMerchantsIdApiCredentials) | **Post** /merchants/{id}/apiCredentials | Create an API credential



## GetMerchantsIdApiCredentials

> ListMerchantApiCredentialsResponse GetMerchantsIdApiCredentials(ctx, id).PageNumber(pageNumber).PageSize(pageSize).Execute()

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
    id := "id_example" // string | The unique identifier of the merchant account.
    pageNumber := int32(56) // int32 | The number of the page to fetch. (optional)
    pageSize := int32(56) // int32 | The number of items to have on a page, maximum 100. The default is 10 items on a page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APICredentialsMerchantLevelApi.GetMerchantsIdApiCredentials(context.Background(), id).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APICredentialsMerchantLevelApi.GetMerchantsIdApiCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsIdApiCredentials`: ListMerchantApiCredentialsResponse
    fmt.Fprintf(os.Stdout, "Response from `APICredentialsMerchantLevelApi.GetMerchantsIdApiCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsIdApiCredentialsRequest struct via the builder pattern


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


## GetMerchantsIdApiCredentialsApiCredentialId

> ApiCredential GetMerchantsIdApiCredentialsApiCredentialId(ctx, id, apiCredentialId).Execute()

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
    id := "id_example" // string | The unique identifier of the merchant account.
    apiCredentialId := "apiCredentialId_example" // string | Unique identifier of the API credential.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APICredentialsMerchantLevelApi.GetMerchantsIdApiCredentialsApiCredentialId(context.Background(), id, apiCredentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APICredentialsMerchantLevelApi.GetMerchantsIdApiCredentialsApiCredentialId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsIdApiCredentialsApiCredentialId`: ApiCredential
    fmt.Fprintf(os.Stdout, "Response from `APICredentialsMerchantLevelApi.GetMerchantsIdApiCredentialsApiCredentialId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the merchant account. | 
**apiCredentialId** | **string** | Unique identifier of the API credential. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsIdApiCredentialsApiCredentialIdRequest struct via the builder pattern


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


## PatchMerchantsIdApiCredentialsApiCredentialId

> ApiCredential PatchMerchantsIdApiCredentialsApiCredentialId(ctx, id, apiCredentialId).UpdateMerchantApiCredentialRequest(updateMerchantApiCredentialRequest).Execute()

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
    id := "id_example" // string | The unique identifier of the merchant account.
    apiCredentialId := "apiCredentialId_example" // string | Unique identifier of the API credential.
    updateMerchantApiCredentialRequest := *openapiclient.NewUpdateMerchantApiCredentialRequest() // UpdateMerchantApiCredentialRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APICredentialsMerchantLevelApi.PatchMerchantsIdApiCredentialsApiCredentialId(context.Background(), id, apiCredentialId).UpdateMerchantApiCredentialRequest(updateMerchantApiCredentialRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APICredentialsMerchantLevelApi.PatchMerchantsIdApiCredentialsApiCredentialId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMerchantsIdApiCredentialsApiCredentialId`: ApiCredential
    fmt.Fprintf(os.Stdout, "Response from `APICredentialsMerchantLevelApi.PatchMerchantsIdApiCredentialsApiCredentialId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the merchant account. | 
**apiCredentialId** | **string** | Unique identifier of the API credential. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMerchantsIdApiCredentialsApiCredentialIdRequest struct via the builder pattern


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


## PostMerchantsIdApiCredentials

> CreateApiCredentialResponse PostMerchantsIdApiCredentials(ctx, id).CreateMerchantApiCredentialRequest(createMerchantApiCredentialRequest).Execute()

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
    id := "id_example" // string | The unique identifier of the merchant account.
    createMerchantApiCredentialRequest := *openapiclient.NewCreateMerchantApiCredentialRequest() // CreateMerchantApiCredentialRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APICredentialsMerchantLevelApi.PostMerchantsIdApiCredentials(context.Background(), id).CreateMerchantApiCredentialRequest(createMerchantApiCredentialRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APICredentialsMerchantLevelApi.PostMerchantsIdApiCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMerchantsIdApiCredentials`: CreateApiCredentialResponse
    fmt.Fprintf(os.Stdout, "Response from `APICredentialsMerchantLevelApi.PostMerchantsIdApiCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMerchantsIdApiCredentialsRequest struct via the builder pattern


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

