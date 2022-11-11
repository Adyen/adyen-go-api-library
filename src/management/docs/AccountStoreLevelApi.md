# \AccountStoreLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMerchantsMerchantIdStores**](AccountStoreLevelApi.md#GetMerchantsMerchantIdStores) | **Get** /merchants/{merchantId}/stores | Get a list of stores
[**GetMerchantsMerchantIdStoresStoreId**](AccountStoreLevelApi.md#GetMerchantsMerchantIdStoresStoreId) | **Get** /merchants/{merchantId}/stores/{storeId} | Get a store
[**GetStores**](AccountStoreLevelApi.md#GetStores) | **Get** /stores | Get a list of stores
[**GetStoresStoreId**](AccountStoreLevelApi.md#GetStoresStoreId) | **Get** /stores/{storeId} | Get a store
[**PatchMerchantsMerchantIdStoresStoreId**](AccountStoreLevelApi.md#PatchMerchantsMerchantIdStoresStoreId) | **Patch** /merchants/{merchantId}/stores/{storeId} | Update a store
[**PatchStoresStoreId**](AccountStoreLevelApi.md#PatchStoresStoreId) | **Patch** /stores/{storeId} | Update a store
[**PostMerchantsMerchantIdStores**](AccountStoreLevelApi.md#PostMerchantsMerchantIdStores) | **Post** /merchants/{merchantId}/stores | Create a store
[**PostStores**](AccountStoreLevelApi.md#PostStores) | **Post** /stores | Create a store



## GetMerchantsMerchantIdStores

> ListStoresResponse GetMerchantsMerchantIdStores(ctx, merchantId).PageNumber(pageNumber).PageSize(pageSize).Reference(reference).Execute()

Get a list of stores



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
    reference := "reference_example" // string | The reference of the store. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountStoreLevelApi.GetMerchantsMerchantIdStores(context.Background(), merchantId).PageNumber(pageNumber).PageSize(pageSize).Reference(reference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountStoreLevelApi.GetMerchantsMerchantIdStores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsMerchantIdStores`: ListStoresResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountStoreLevelApi.GetMerchantsMerchantIdStores`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsMerchantIdStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | The number of the page to fetch. | 
 **pageSize** | **int32** | The number of items to have on a page, maximum 100. The default is 10 items on a page. | 
 **reference** | **string** | The reference of the store. | 

### Return type

[**ListStoresResponse**](ListStoresResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMerchantsMerchantIdStoresStoreId

> Store GetMerchantsMerchantIdStoresStoreId(ctx, merchantId, storeId).Execute()

Get a store



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
    storeId := "storeId_example" // string | The unique identifier of the store.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountStoreLevelApi.GetMerchantsMerchantIdStoresStoreId(context.Background(), merchantId, storeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountStoreLevelApi.GetMerchantsMerchantIdStoresStoreId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsMerchantIdStoresStoreId`: Store
    fmt.Fprintf(os.Stdout, "Response from `AccountStoreLevelApi.GetMerchantsMerchantIdStoresStoreId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**storeId** | **string** | The unique identifier of the store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsMerchantIdStoresStoreIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Store**](Store.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStores

> ListStoresResponse GetStores(ctx).PageNumber(pageNumber).PageSize(pageSize).Reference(reference).MerchantId(merchantId).Execute()

Get a list of stores



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
    reference := "reference_example" // string | The reference of the store. (optional)
    merchantId := "merchantId_example" // string | The unique identifier of the merchant account. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountStoreLevelApi.GetStores(context.Background()).PageNumber(pageNumber).PageSize(pageSize).Reference(reference).MerchantId(merchantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountStoreLevelApi.GetStores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStores`: ListStoresResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountStoreLevelApi.GetStores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The number of the page to fetch. | 
 **pageSize** | **int32** | The number of items to have on a page, maximum 100. The default is 10 items on a page. | 
 **reference** | **string** | The reference of the store. | 
 **merchantId** | **string** | The unique identifier of the merchant account. | 

### Return type

[**ListStoresResponse**](ListStoresResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStoresStoreId

> Store GetStoresStoreId(ctx, storeId).Execute()

Get a store



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
    storeId := "storeId_example" // string | The unique identifier of the store.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountStoreLevelApi.GetStoresStoreId(context.Background(), storeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountStoreLevelApi.GetStoresStoreId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStoresStoreId`: Store
    fmt.Fprintf(os.Stdout, "Response from `AccountStoreLevelApi.GetStoresStoreId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The unique identifier of the store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoresStoreIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Store**](Store.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMerchantsMerchantIdStoresStoreId

> Store PatchMerchantsMerchantIdStoresStoreId(ctx, merchantId, storeId).UpdateStoreRequest(updateStoreRequest).Execute()

Update a store



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
    storeId := "storeId_example" // string | The unique identifier of the store.
    updateStoreRequest := *openapiclient.NewUpdateStoreRequest() // UpdateStoreRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountStoreLevelApi.PatchMerchantsMerchantIdStoresStoreId(context.Background(), merchantId, storeId).UpdateStoreRequest(updateStoreRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountStoreLevelApi.PatchMerchantsMerchantIdStoresStoreId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMerchantsMerchantIdStoresStoreId`: Store
    fmt.Fprintf(os.Stdout, "Response from `AccountStoreLevelApi.PatchMerchantsMerchantIdStoresStoreId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**storeId** | **string** | The unique identifier of the store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMerchantsMerchantIdStoresStoreIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateStoreRequest** | [**UpdateStoreRequest**](UpdateStoreRequest.md) |  | 

### Return type

[**Store**](Store.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchStoresStoreId

> Store PatchStoresStoreId(ctx, storeId).UpdateStoreRequest(updateStoreRequest).Execute()

Update a store



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
    storeId := "storeId_example" // string | The unique identifier of the store.
    updateStoreRequest := *openapiclient.NewUpdateStoreRequest() // UpdateStoreRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountStoreLevelApi.PatchStoresStoreId(context.Background(), storeId).UpdateStoreRequest(updateStoreRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountStoreLevelApi.PatchStoresStoreId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchStoresStoreId`: Store
    fmt.Fprintf(os.Stdout, "Response from `AccountStoreLevelApi.PatchStoresStoreId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The unique identifier of the store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchStoresStoreIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateStoreRequest** | [**UpdateStoreRequest**](UpdateStoreRequest.md) |  | 

### Return type

[**Store**](Store.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMerchantsMerchantIdStores

> Store PostMerchantsMerchantIdStores(ctx, merchantId).StoreCreationRequest(storeCreationRequest).Execute()

Create a store



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
    storeCreationRequest := *openapiclient.NewStoreCreationRequest(*openapiclient.NewAddress2("Country_example"), "Description_example", "PhoneNumber_example", "ShopperStatement_example") // StoreCreationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountStoreLevelApi.PostMerchantsMerchantIdStores(context.Background(), merchantId).StoreCreationRequest(storeCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountStoreLevelApi.PostMerchantsMerchantIdStores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMerchantsMerchantIdStores`: Store
    fmt.Fprintf(os.Stdout, "Response from `AccountStoreLevelApi.PostMerchantsMerchantIdStores`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMerchantsMerchantIdStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storeCreationRequest** | [**StoreCreationRequest**](StoreCreationRequest.md) |  | 

### Return type

[**Store**](Store.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostStores

> Store PostStores(ctx).StoreCreationWithMerchantCodeRequest(storeCreationWithMerchantCodeRequest).Execute()

Create a store



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
    storeCreationWithMerchantCodeRequest := *openapiclient.NewStoreCreationWithMerchantCodeRequest(*openapiclient.NewAddress2("Country_example"), "Description_example", "MerchantId_example", "PhoneNumber_example", "ShopperStatement_example") // StoreCreationWithMerchantCodeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountStoreLevelApi.PostStores(context.Background()).StoreCreationWithMerchantCodeRequest(storeCreationWithMerchantCodeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountStoreLevelApi.PostStores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostStores`: Store
    fmt.Fprintf(os.Stdout, "Response from `AccountStoreLevelApi.PostStores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeCreationWithMerchantCodeRequest** | [**StoreCreationWithMerchantCodeRequest**](StoreCreationWithMerchantCodeRequest.md) |  | 

### Return type

[**Store**](Store.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

