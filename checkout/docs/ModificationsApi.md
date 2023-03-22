# \ModificationsApi

All URIs are relative to *https://checkout-test.adyen.com/v70*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCancels**](ModificationsApi.md#PostCancels) | **Post** /cancels | Cancel an authorised payment
[**PostPaymentsPaymentPspReferenceAmountUpdates**](ModificationsApi.md#PostPaymentsPaymentPspReferenceAmountUpdates) | **Post** /payments/{paymentPspReference}/amountUpdates | Update an authorised amount
[**PostPaymentsPaymentPspReferenceCancels**](ModificationsApi.md#PostPaymentsPaymentPspReferenceCancels) | **Post** /payments/{paymentPspReference}/cancels | Cancel an authorised payment
[**PostPaymentsPaymentPspReferenceCaptures**](ModificationsApi.md#PostPaymentsPaymentPspReferenceCaptures) | **Post** /payments/{paymentPspReference}/captures | Capture an authorised payment
[**PostPaymentsPaymentPspReferenceRefunds**](ModificationsApi.md#PostPaymentsPaymentPspReferenceRefunds) | **Post** /payments/{paymentPspReference}/refunds | Refund a captured payment
[**PostPaymentsPaymentPspReferenceReversals**](ModificationsApi.md#PostPaymentsPaymentPspReferenceReversals) | **Post** /payments/{paymentPspReference}/reversals | Refund or cancel a payment



## PostCancels

> StandalonePaymentCancelResource PostCancels(ctx).IdempotencyKey(idempotencyKey).CreateStandalonePaymentCancelRequest(createStandalonePaymentCancelRequest).Execute()

Cancel an authorised payment



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
    createStandalonePaymentCancelRequest := *openapiclient.NewCreateStandalonePaymentCancelRequest("MerchantAccount_example", "PaymentReference_example") // CreateStandalonePaymentCancelRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModificationsApi.PostCancels(context.Background()).IdempotencyKey(idempotencyKey).CreateStandalonePaymentCancelRequest(createStandalonePaymentCancelRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModificationsApi.PostCancels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCancels`: StandalonePaymentCancelResource
    fmt.Fprintf(os.Stdout, "Response from `ModificationsApi.PostCancels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCancelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | A unique identifier for the message with a maximum of 64 characters (we recommend a UUID). | 
 **createStandalonePaymentCancelRequest** | [**CreateStandalonePaymentCancelRequest**](CreateStandalonePaymentCancelRequest.md) |  | 

### Return type

[**StandalonePaymentCancelResource**](StandalonePaymentCancelResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPaymentsPaymentPspReferenceAmountUpdates

> PaymentAmountUpdateResource PostPaymentsPaymentPspReferenceAmountUpdates(ctx, paymentPspReference).IdempotencyKey(idempotencyKey).CreatePaymentAmountUpdateRequest(createPaymentAmountUpdateRequest).Execute()

Update an authorised amount



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
    paymentPspReference := "paymentPspReference_example" // string | The [`pspReference`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment.
    idempotencyKey := "37ca9c97-d1d1-4c62-89e8-706891a563ed" // string | A unique identifier for the message with a maximum of 64 characters (we recommend a UUID). (optional)
    createPaymentAmountUpdateRequest := *openapiclient.NewCreatePaymentAmountUpdateRequest(*openapiclient.NewAmount("Currency_example", int64(123)), "MerchantAccount_example") // CreatePaymentAmountUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModificationsApi.PostPaymentsPaymentPspReferenceAmountUpdates(context.Background(), paymentPspReference).IdempotencyKey(idempotencyKey).CreatePaymentAmountUpdateRequest(createPaymentAmountUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModificationsApi.PostPaymentsPaymentPspReferenceAmountUpdates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPaymentsPaymentPspReferenceAmountUpdates`: PaymentAmountUpdateResource
    fmt.Fprintf(os.Stdout, "Response from `ModificationsApi.PostPaymentsPaymentPspReferenceAmountUpdates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentPspReference** | **string** | The [&#x60;pspReference&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPaymentsPaymentPspReferenceAmountUpdatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | A unique identifier for the message with a maximum of 64 characters (we recommend a UUID). | 
 **createPaymentAmountUpdateRequest** | [**CreatePaymentAmountUpdateRequest**](CreatePaymentAmountUpdateRequest.md) |  | 

### Return type

[**PaymentAmountUpdateResource**](PaymentAmountUpdateResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPaymentsPaymentPspReferenceCancels

> PaymentCancelResource PostPaymentsPaymentPspReferenceCancels(ctx, paymentPspReference).IdempotencyKey(idempotencyKey).CreatePaymentCancelRequest(createPaymentCancelRequest).Execute()

Cancel an authorised payment



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
    paymentPspReference := "paymentPspReference_example" // string | The [`pspReference`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment that you want to cancel. 
    idempotencyKey := "37ca9c97-d1d1-4c62-89e8-706891a563ed" // string | A unique identifier for the message with a maximum of 64 characters (we recommend a UUID). (optional)
    createPaymentCancelRequest := *openapiclient.NewCreatePaymentCancelRequest("MerchantAccount_example") // CreatePaymentCancelRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModificationsApi.PostPaymentsPaymentPspReferenceCancels(context.Background(), paymentPspReference).IdempotencyKey(idempotencyKey).CreatePaymentCancelRequest(createPaymentCancelRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModificationsApi.PostPaymentsPaymentPspReferenceCancels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPaymentsPaymentPspReferenceCancels`: PaymentCancelResource
    fmt.Fprintf(os.Stdout, "Response from `ModificationsApi.PostPaymentsPaymentPspReferenceCancels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentPspReference** | **string** | The [&#x60;pspReference&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment that you want to cancel.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPaymentsPaymentPspReferenceCancelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | A unique identifier for the message with a maximum of 64 characters (we recommend a UUID). | 
 **createPaymentCancelRequest** | [**CreatePaymentCancelRequest**](CreatePaymentCancelRequest.md) |  | 

### Return type

[**PaymentCancelResource**](PaymentCancelResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPaymentsPaymentPspReferenceCaptures

> PaymentCaptureResource PostPaymentsPaymentPspReferenceCaptures(ctx, paymentPspReference).IdempotencyKey(idempotencyKey).CreatePaymentCaptureRequest(createPaymentCaptureRequest).Execute()

Capture an authorised payment



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
    paymentPspReference := "paymentPspReference_example" // string | The [`pspReference`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment that you want to capture.
    idempotencyKey := "37ca9c97-d1d1-4c62-89e8-706891a563ed" // string | A unique identifier for the message with a maximum of 64 characters (we recommend a UUID). (optional)
    createPaymentCaptureRequest := *openapiclient.NewCreatePaymentCaptureRequest(*openapiclient.NewAmount("Currency_example", int64(123)), "MerchantAccount_example") // CreatePaymentCaptureRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModificationsApi.PostPaymentsPaymentPspReferenceCaptures(context.Background(), paymentPspReference).IdempotencyKey(idempotencyKey).CreatePaymentCaptureRequest(createPaymentCaptureRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModificationsApi.PostPaymentsPaymentPspReferenceCaptures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPaymentsPaymentPspReferenceCaptures`: PaymentCaptureResource
    fmt.Fprintf(os.Stdout, "Response from `ModificationsApi.PostPaymentsPaymentPspReferenceCaptures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentPspReference** | **string** | The [&#x60;pspReference&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment that you want to capture. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPaymentsPaymentPspReferenceCapturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | A unique identifier for the message with a maximum of 64 characters (we recommend a UUID). | 
 **createPaymentCaptureRequest** | [**CreatePaymentCaptureRequest**](CreatePaymentCaptureRequest.md) |  | 

### Return type

[**PaymentCaptureResource**](PaymentCaptureResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPaymentsPaymentPspReferenceRefunds

> PaymentRefundResource PostPaymentsPaymentPspReferenceRefunds(ctx, paymentPspReference).IdempotencyKey(idempotencyKey).CreatePaymentRefundRequest(createPaymentRefundRequest).Execute()

Refund a captured payment



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
    paymentPspReference := "paymentPspReference_example" // string | The [`pspReference`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment that you want to refund.
    idempotencyKey := "37ca9c97-d1d1-4c62-89e8-706891a563ed" // string | A unique identifier for the message with a maximum of 64 characters (we recommend a UUID). (optional)
    createPaymentRefundRequest := *openapiclient.NewCreatePaymentRefundRequest(*openapiclient.NewAmount("Currency_example", int64(123)), "MerchantAccount_example") // CreatePaymentRefundRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModificationsApi.PostPaymentsPaymentPspReferenceRefunds(context.Background(), paymentPspReference).IdempotencyKey(idempotencyKey).CreatePaymentRefundRequest(createPaymentRefundRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModificationsApi.PostPaymentsPaymentPspReferenceRefunds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPaymentsPaymentPspReferenceRefunds`: PaymentRefundResource
    fmt.Fprintf(os.Stdout, "Response from `ModificationsApi.PostPaymentsPaymentPspReferenceRefunds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentPspReference** | **string** | The [&#x60;pspReference&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment that you want to refund. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPaymentsPaymentPspReferenceRefundsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | A unique identifier for the message with a maximum of 64 characters (we recommend a UUID). | 
 **createPaymentRefundRequest** | [**CreatePaymentRefundRequest**](CreatePaymentRefundRequest.md) |  | 

### Return type

[**PaymentRefundResource**](PaymentRefundResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPaymentsPaymentPspReferenceReversals

> PaymentReversalResource PostPaymentsPaymentPspReferenceReversals(ctx, paymentPspReference).IdempotencyKey(idempotencyKey).CreatePaymentReversalRequest(createPaymentReversalRequest).Execute()

Refund or cancel a payment



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
    paymentPspReference := "paymentPspReference_example" // string | The [`pspReference`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment that you want to reverse. 
    idempotencyKey := "37ca9c97-d1d1-4c62-89e8-706891a563ed" // string | A unique identifier for the message with a maximum of 64 characters (we recommend a UUID). (optional)
    createPaymentReversalRequest := *openapiclient.NewCreatePaymentReversalRequest("MerchantAccount_example") // CreatePaymentReversalRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModificationsApi.PostPaymentsPaymentPspReferenceReversals(context.Background(), paymentPspReference).IdempotencyKey(idempotencyKey).CreatePaymentReversalRequest(createPaymentReversalRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModificationsApi.PostPaymentsPaymentPspReferenceReversals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPaymentsPaymentPspReferenceReversals`: PaymentReversalResource
    fmt.Fprintf(os.Stdout, "Response from `ModificationsApi.PostPaymentsPaymentPspReferenceReversals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentPspReference** | **string** | The [&#x60;pspReference&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment that you want to reverse.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPaymentsPaymentPspReferenceReversalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | A unique identifier for the message with a maximum of 64 characters (we recommend a UUID). | 
 **createPaymentReversalRequest** | [**CreatePaymentReversalRequest**](CreatePaymentReversalRequest.md) |  | 

### Return type

[**PaymentReversalResource**](PaymentReversalResource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

