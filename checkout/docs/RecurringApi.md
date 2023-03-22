# \RecurringApi

All URIs are relative to *https://checkout-test.adyen.com/v70*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteStoredPaymentMethodsRecurringId**](RecurringApi.md#DeleteStoredPaymentMethodsRecurringId) | **Delete** /storedPaymentMethods/{recurringId} | Delete a token for stored payment details
[**GetStoredPaymentMethods**](RecurringApi.md#GetStoredPaymentMethods) | **Get** /storedPaymentMethods | Get tokens for stored payment details



## DeleteStoredPaymentMethodsRecurringId

> StoredPaymentMethodResource DeleteStoredPaymentMethodsRecurringId(ctx, recurringId).ShopperReference(shopperReference).MerchantAccount(merchantAccount).Execute()

Delete a token for stored payment details



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/adyen/adyen-go-api-library/v6"
)

func main() {
    recurringId := "recurringId_example" // string | The unique identifier of the token.
    shopperReference := "shopperReference_example" // string | Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. > Your reference must not include personally identifiable information (PII), for example name or email address. (optional)
    merchantAccount := "merchantAccount_example" // string | Your merchant account. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecurringApi.DeleteStoredPaymentMethodsRecurringId(context.Background(), recurringId).ShopperReference(shopperReference).MerchantAccount(merchantAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringApi.DeleteStoredPaymentMethodsRecurringId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteStoredPaymentMethodsRecurringId`: StoredPaymentMethodResource
    fmt.Fprintf(os.Stdout, "Response from `RecurringApi.DeleteStoredPaymentMethodsRecurringId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recurringId** | **string** | The unique identifier of the token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStoredPaymentMethodsRecurringIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shopperReference** | **string** | Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. &gt; Your reference must not include personally identifiable information (PII), for example name or email address. | 
 **merchantAccount** | **string** | Your merchant account. | 

### Return type

[**StoredPaymentMethodResource**](StoredPaymentMethodResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStoredPaymentMethods

> ListStoredPaymentMethodsResponse GetStoredPaymentMethods(ctx).ShopperReference(shopperReference).MerchantAccount(merchantAccount).Execute()

Get tokens for stored payment details



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/adyen/adyen-go-api-library/v6"
)

func main() {
    shopperReference := "shopperReference_example" // string | Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. > Your reference must not include personally identifiable information (PII), for example name or email address. (optional)
    merchantAccount := "merchantAccount_example" // string | Your merchant account. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecurringApi.GetStoredPaymentMethods(context.Background()).ShopperReference(shopperReference).MerchantAccount(merchantAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringApi.GetStoredPaymentMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStoredPaymentMethods`: ListStoredPaymentMethodsResponse
    fmt.Fprintf(os.Stdout, "Response from `RecurringApi.GetStoredPaymentMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStoredPaymentMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shopperReference** | **string** | Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. &gt; Your reference must not include personally identifiable information (PII), for example name or email address. | 
 **merchantAccount** | **string** | Your merchant account. | 

### Return type

[**ListStoredPaymentMethodsResponse**](ListStoredPaymentMethodsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

