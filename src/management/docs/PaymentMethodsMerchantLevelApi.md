# \PaymentMethodsMerchantLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMerchantsMerchantIdPaymentMethodSettings**](PaymentMethodsMerchantLevelApi.md#GetMerchantsMerchantIdPaymentMethodSettings) | **Get** /merchants/{merchantId}/paymentMethodSettings | Get all payment methods
[**GetMerchantsMerchantIdPaymentMethodSettingsPaymentMethodId**](PaymentMethodsMerchantLevelApi.md#GetMerchantsMerchantIdPaymentMethodSettingsPaymentMethodId) | **Get** /merchants/{merchantId}/paymentMethodSettings/{paymentMethodId} | Get payment method details
[**PatchMerchantsMerchantIdPaymentMethodSettingsPaymentMethodId**](PaymentMethodsMerchantLevelApi.md#PatchMerchantsMerchantIdPaymentMethodSettingsPaymentMethodId) | **Patch** /merchants/{merchantId}/paymentMethodSettings/{paymentMethodId} | Update a payment method
[**PostMerchantsMerchantIdPaymentMethodSettings**](PaymentMethodsMerchantLevelApi.md#PostMerchantsMerchantIdPaymentMethodSettings) | **Post** /merchants/{merchantId}/paymentMethodSettings | Request a payment method



## GetMerchantsMerchantIdPaymentMethodSettings

> PaymentMethodResponse GetMerchantsMerchantIdPaymentMethodSettings(ctx, merchantId).StoreId(storeId).BusinessLineId(businessLineId).PageSize(pageSize).PageNumber(pageNumber).Execute()

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
    merchantId := "merchantId_example" // string | The unique identifier of the merchant account.
    storeId := "storeId_example" // string | The unique identifier of the store for which to return the payment methods. (optional)
    businessLineId := "businessLineId_example" // string | The unique identifier of the Business Line for which to return the payment methods. (optional)
    pageSize := int32(56) // int32 | The number of items to have on a page, maximum 100. The default is 10 items on a page. (optional)
    pageNumber := int32(56) // int32 | The number of the page to fetch. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentMethodsMerchantLevelApi.GetMerchantsMerchantIdPaymentMethodSettings(context.Background(), merchantId).StoreId(storeId).BusinessLineId(businessLineId).PageSize(pageSize).PageNumber(pageNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsMerchantLevelApi.GetMerchantsMerchantIdPaymentMethodSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsMerchantIdPaymentMethodSettings`: PaymentMethodResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsMerchantLevelApi.GetMerchantsMerchantIdPaymentMethodSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsMerchantIdPaymentMethodSettingsRequest struct via the builder pattern


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


## GetMerchantsMerchantIdPaymentMethodSettingsPaymentMethodId

> PaymentMethod GetMerchantsMerchantIdPaymentMethodSettingsPaymentMethodId(ctx, merchantId, paymentMethodId).Execute()

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
    merchantId := "merchantId_example" // string | The unique identifier of the merchant account.
    paymentMethodId := "paymentMethodId_example" // string | The unique identifier of the payment method.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentMethodsMerchantLevelApi.GetMerchantsMerchantIdPaymentMethodSettingsPaymentMethodId(context.Background(), merchantId, paymentMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsMerchantLevelApi.GetMerchantsMerchantIdPaymentMethodSettingsPaymentMethodId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsMerchantIdPaymentMethodSettingsPaymentMethodId`: PaymentMethod
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsMerchantLevelApi.GetMerchantsMerchantIdPaymentMethodSettingsPaymentMethodId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**paymentMethodId** | **string** | The unique identifier of the payment method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsMerchantIdPaymentMethodSettingsPaymentMethodIdRequest struct via the builder pattern


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


## PatchMerchantsMerchantIdPaymentMethodSettingsPaymentMethodId

> PaymentMethod PatchMerchantsMerchantIdPaymentMethodSettingsPaymentMethodId(ctx, merchantId, paymentMethodId).UpdatePaymentMethodInfo(updatePaymentMethodInfo).Execute()

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
    merchantId := "merchantId_example" // string | The unique identifier of the merchant account.
    paymentMethodId := "paymentMethodId_example" // string | The unique identifier of the payment method.
    updatePaymentMethodInfo := *openapiclient.NewUpdatePaymentMethodInfo() // UpdatePaymentMethodInfo |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentMethodsMerchantLevelApi.PatchMerchantsMerchantIdPaymentMethodSettingsPaymentMethodId(context.Background(), merchantId, paymentMethodId).UpdatePaymentMethodInfo(updatePaymentMethodInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsMerchantLevelApi.PatchMerchantsMerchantIdPaymentMethodSettingsPaymentMethodId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMerchantsMerchantIdPaymentMethodSettingsPaymentMethodId`: PaymentMethod
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsMerchantLevelApi.PatchMerchantsMerchantIdPaymentMethodSettingsPaymentMethodId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**paymentMethodId** | **string** | The unique identifier of the payment method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMerchantsMerchantIdPaymentMethodSettingsPaymentMethodIdRequest struct via the builder pattern


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


## PostMerchantsMerchantIdPaymentMethodSettings

> PaymentMethod PostMerchantsMerchantIdPaymentMethodSettings(ctx, merchantId).PaymentMethodSetupInfo(paymentMethodSetupInfo).Execute()

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
    merchantId := "merchantId_example" // string | The unique identifier of the merchant account.
    paymentMethodSetupInfo := *openapiclient.NewPaymentMethodSetupInfo("Type_example") // PaymentMethodSetupInfo |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentMethodsMerchantLevelApi.PostMerchantsMerchantIdPaymentMethodSettings(context.Background(), merchantId).PaymentMethodSetupInfo(paymentMethodSetupInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsMerchantLevelApi.PostMerchantsMerchantIdPaymentMethodSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMerchantsMerchantIdPaymentMethodSettings`: PaymentMethod
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsMerchantLevelApi.PostMerchantsMerchantIdPaymentMethodSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMerchantsMerchantIdPaymentMethodSettingsRequest struct via the builder pattern


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

