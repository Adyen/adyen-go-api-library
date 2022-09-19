# \PaymentMethodsMerchantLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMerchantsIdPaymentMethodSettings**](PaymentMethodsMerchantLevelApi.md#GetMerchantsIdPaymentMethodSettings) | **Get** /merchants/{id}/paymentMethodSettings | Get all payment methods
[**GetMerchantsIdPaymentMethodSettingsPaymentMethodId**](PaymentMethodsMerchantLevelApi.md#GetMerchantsIdPaymentMethodSettingsPaymentMethodId) | **Get** /merchants/{id}/paymentMethodSettings/{paymentMethodId} | Get payment method details
[**PatchMerchantsIdPaymentMethodSettingsPaymentMethodId**](PaymentMethodsMerchantLevelApi.md#PatchMerchantsIdPaymentMethodSettingsPaymentMethodId) | **Patch** /merchants/{id}/paymentMethodSettings/{paymentMethodId} | Update a payment method
[**PostMerchantsIdPaymentMethodSettings**](PaymentMethodsMerchantLevelApi.md#PostMerchantsIdPaymentMethodSettings) | **Post** /merchants/{id}/paymentMethodSettings | Request a payment method



## GetMerchantsIdPaymentMethodSettings

> PaymentMethodResponse GetMerchantsIdPaymentMethodSettings(ctx, id).StoreId(storeId).BusinessLineId(businessLineId).PageSize(pageSize).PageNumber(pageNumber).Execute()

Get all payment methods



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
    storeId := "storeId_example" // string | The unique identifier of the store for which to return the payment methods. (optional)
    businessLineId := "businessLineId_example" // string | The unique identifier of the Business Line for which to return the payment methods. (optional)
    pageSize := int32(56) // int32 | The number of items to have on a page, maximum 100. The default is 10 items on a page. (optional)
    pageNumber := int32(56) // int32 | The number of the page to fetch. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentMethodsMerchantLevelApi.GetMerchantsIdPaymentMethodSettings(context.Background(), id).StoreId(storeId).BusinessLineId(businessLineId).PageSize(pageSize).PageNumber(pageNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsMerchantLevelApi.GetMerchantsIdPaymentMethodSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsIdPaymentMethodSettings`: PaymentMethodResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsMerchantLevelApi.GetMerchantsIdPaymentMethodSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsIdPaymentMethodSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storeId** | **string** | The unique identifier of the store for which to return the payment methods. | 
 **businessLineId** | **string** | The unique identifier of the Business Line for which to return the payment methods. | 
 **pageSize** | **int32** | The number of items to have on a page, maximum 100. The default is 10 items on a page. | 
 **pageNumber** | **int32** | The number of the page to fetch. | 

### Return type

[**PaymentMethodResponse**](PaymentMethodResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMerchantsIdPaymentMethodSettingsPaymentMethodId

> PaymentMethod GetMerchantsIdPaymentMethodSettingsPaymentMethodId(ctx, id, paymentMethodId).Execute()

Get payment method details



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
    paymentMethodId := "paymentMethodId_example" // string | The unique identifier of the payment method.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentMethodsMerchantLevelApi.GetMerchantsIdPaymentMethodSettingsPaymentMethodId(context.Background(), id, paymentMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsMerchantLevelApi.GetMerchantsIdPaymentMethodSettingsPaymentMethodId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsIdPaymentMethodSettingsPaymentMethodId`: PaymentMethod
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsMerchantLevelApi.GetMerchantsIdPaymentMethodSettingsPaymentMethodId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the merchant account. | 
**paymentMethodId** | **string** | The unique identifier of the payment method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsIdPaymentMethodSettingsPaymentMethodIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaymentMethod**](PaymentMethod.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMerchantsIdPaymentMethodSettingsPaymentMethodId

> PaymentMethod PatchMerchantsIdPaymentMethodSettingsPaymentMethodId(ctx, id, paymentMethodId).UpdatePaymentMethodInfo(updatePaymentMethodInfo).Execute()

Update a payment method



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
    paymentMethodId := "paymentMethodId_example" // string | The unique identifier of the payment method.
    updatePaymentMethodInfo := *openapiclient.NewUpdatePaymentMethodInfo() // UpdatePaymentMethodInfo |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentMethodsMerchantLevelApi.PatchMerchantsIdPaymentMethodSettingsPaymentMethodId(context.Background(), id, paymentMethodId).UpdatePaymentMethodInfo(updatePaymentMethodInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsMerchantLevelApi.PatchMerchantsIdPaymentMethodSettingsPaymentMethodId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMerchantsIdPaymentMethodSettingsPaymentMethodId`: PaymentMethod
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsMerchantLevelApi.PatchMerchantsIdPaymentMethodSettingsPaymentMethodId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the merchant account. | 
**paymentMethodId** | **string** | The unique identifier of the payment method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMerchantsIdPaymentMethodSettingsPaymentMethodIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updatePaymentMethodInfo** | [**UpdatePaymentMethodInfo**](UpdatePaymentMethodInfo.md) |  | 

### Return type

[**PaymentMethod**](PaymentMethod.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMerchantsIdPaymentMethodSettings

> PaymentMethod PostMerchantsIdPaymentMethodSettings(ctx, id).PaymentMethodSetupInfo(paymentMethodSetupInfo).Execute()

Request a payment method



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
    paymentMethodSetupInfo := *openapiclient.NewPaymentMethodSetupInfo("Type_example") // PaymentMethodSetupInfo |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentMethodsMerchantLevelApi.PostMerchantsIdPaymentMethodSettings(context.Background(), id).PaymentMethodSetupInfo(paymentMethodSetupInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsMerchantLevelApi.PostMerchantsIdPaymentMethodSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMerchantsIdPaymentMethodSettings`: PaymentMethod
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsMerchantLevelApi.PostMerchantsIdPaymentMethodSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMerchantsIdPaymentMethodSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentMethodSetupInfo** | [**PaymentMethodSetupInfo**](PaymentMethodSetupInfo.md) |  | 

### Return type

[**PaymentMethod**](PaymentMethod.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

