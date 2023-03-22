# \PaymentLinksApi

All URIs are relative to *https://checkout-test.adyen.com/v70*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPaymentLinksLinkId**](PaymentLinksApi.md#GetPaymentLinksLinkId) | **Get** /paymentLinks/{linkId} | Get a payment link
[**PatchPaymentLinksLinkId**](PaymentLinksApi.md#PatchPaymentLinksLinkId) | **Patch** /paymentLinks/{linkId} | Update the status of a payment link
[**PostPaymentLinks**](PaymentLinksApi.md#PostPaymentLinks) | **Post** /paymentLinks | Create a payment link



## GetPaymentLinksLinkId

> PaymentLinkResponse GetPaymentLinksLinkId(ctx, linkId).Execute()

Get a payment link



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
    linkId := "linkId_example" // string | Unique identifier of the payment link.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentLinksApi.GetPaymentLinksLinkId(context.Background(), linkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksApi.GetPaymentLinksLinkId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentLinksLinkId`: PaymentLinkResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentLinksApi.GetPaymentLinksLinkId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkId** | **string** | Unique identifier of the payment link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentLinksLinkIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentLinkResponse**](PaymentLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPaymentLinksLinkId

> PaymentLinkResponse PatchPaymentLinksLinkId(ctx, linkId).UpdatePaymentLinkRequest(updatePaymentLinkRequest).Execute()

Update the status of a payment link



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
    linkId := "linkId_example" // string | Unique identifier of the payment link.
    updatePaymentLinkRequest := *openapiclient.NewUpdatePaymentLinkRequest("Status_example") // UpdatePaymentLinkRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentLinksApi.PatchPaymentLinksLinkId(context.Background(), linkId).UpdatePaymentLinkRequest(updatePaymentLinkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksApi.PatchPaymentLinksLinkId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPaymentLinksLinkId`: PaymentLinkResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentLinksApi.PatchPaymentLinksLinkId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkId** | **string** | Unique identifier of the payment link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPaymentLinksLinkIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePaymentLinkRequest** | [**UpdatePaymentLinkRequest**](UpdatePaymentLinkRequest.md) |  | 

### Return type

[**PaymentLinkResponse**](PaymentLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPaymentLinks

> PaymentLinkResponse PostPaymentLinks(ctx).IdempotencyKey(idempotencyKey).CreatePaymentLinkRequest(createPaymentLinkRequest).Execute()

Create a payment link



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
    idempotencyKey := "37ca9c97-d1d1-4c62-89e8-706891a563ed" // string | A unique identifier for the message with a maximum of 64 characters (we recommend a UUID). (optional)
    createPaymentLinkRequest := *openapiclient.NewCreatePaymentLinkRequest(*openapiclient.NewAmount("Currency_example", int64(123)), "MerchantAccount_example", "Reference_example") // CreatePaymentLinkRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentLinksApi.PostPaymentLinks(context.Background()).IdempotencyKey(idempotencyKey).CreatePaymentLinkRequest(createPaymentLinkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksApi.PostPaymentLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPaymentLinks`: PaymentLinkResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentLinksApi.PostPaymentLinks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPaymentLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | A unique identifier for the message with a maximum of 64 characters (we recommend a UUID). | 
 **createPaymentLinkRequest** | [**CreatePaymentLinkRequest**](CreatePaymentLinkRequest.md) |  | 

### Return type

[**PaymentLinkResponse**](PaymentLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

