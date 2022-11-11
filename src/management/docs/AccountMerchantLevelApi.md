# \AccountMerchantLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMerchants**](AccountMerchantLevelApi.md#GetMerchants) | **Get** /merchants | Get a list of merchant accounts
[**GetMerchantsMerchantId**](AccountMerchantLevelApi.md#GetMerchantsMerchantId) | **Get** /merchants/{merchantId} | Get a merchant account
[**PostMerchants**](AccountMerchantLevelApi.md#PostMerchants) | **Post** /merchants | Create a merchant account
[**PostMerchantsMerchantIdActivate**](AccountMerchantLevelApi.md#PostMerchantsMerchantIdActivate) | **Post** /merchants/{merchantId}/activate | Request to activate a merchant account



## GetMerchants

> ListMerchantResponse GetMerchants(ctx).PageNumber(pageNumber).PageSize(pageSize).Execute()

Get a list of merchant accounts



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
    pageNumber := int32(56) // int32 | The number of the page to fetch. (optional)
    pageSize := int32(56) // int32 | The number of items to have on a page, maximum 100. The default is 10 items on a page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountMerchantLevelApi.GetMerchants(context.Background()).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountMerchantLevelApi.GetMerchants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchants`: ListMerchantResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountMerchantLevelApi.GetMerchants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The number of the page to fetch. | 
 **pageSize** | **int32** | The number of items to have on a page, maximum 100. The default is 10 items on a page. | 

### Return type

[**ListMerchantResponse**](ListMerchantResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMerchantsMerchantId

> Merchant GetMerchantsMerchantId(ctx, merchantId).Execute()

Get a merchant account



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountMerchantLevelApi.GetMerchantsMerchantId(context.Background(), merchantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountMerchantLevelApi.GetMerchantsMerchantId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsMerchantId`: Merchant
    fmt.Fprintf(os.Stdout, "Response from `AccountMerchantLevelApi.GetMerchantsMerchantId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsMerchantIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Merchant**](Merchant.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMerchants

> CreateMerchantResponse PostMerchants(ctx).CreateMerchantRequest(createMerchantRequest).Execute()

Create a merchant account



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
    createMerchantRequest := *openapiclient.NewCreateMerchantRequest("CompanyId_example") // CreateMerchantRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountMerchantLevelApi.PostMerchants(context.Background()).CreateMerchantRequest(createMerchantRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountMerchantLevelApi.PostMerchants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMerchants`: CreateMerchantResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountMerchantLevelApi.PostMerchants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMerchantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMerchantRequest** | [**CreateMerchantRequest**](CreateMerchantRequest.md) |  | 

### Return type

[**CreateMerchantResponse**](CreateMerchantResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMerchantsMerchantIdActivate

> RequestActivationResponse PostMerchantsMerchantIdActivate(ctx, merchantId).Execute()

Request to activate a merchant account



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountMerchantLevelApi.PostMerchantsMerchantIdActivate(context.Background(), merchantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountMerchantLevelApi.PostMerchantsMerchantIdActivate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMerchantsMerchantIdActivate`: RequestActivationResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountMerchantLevelApi.PostMerchantsMerchantIdActivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMerchantsMerchantIdActivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestActivationResponse**](RequestActivationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

