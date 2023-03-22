# \ClassicCheckoutSDKApi

All URIs are relative to *https://checkout-test.adyen.com/v70*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostPaymentSession**](ClassicCheckoutSDKApi.md#PostPaymentSession) | **Post** /paymentSession | Create a payment session
[**PostPaymentsResult**](ClassicCheckoutSDKApi.md#PostPaymentsResult) | **Post** /payments/result | Verify a payment result



## PostPaymentSession

> PaymentSetupResponse PostPaymentSession(ctx).IdempotencyKey(idempotencyKey).PaymentSetupRequest(paymentSetupRequest).Execute()

Create a payment session



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
    paymentSetupRequest := *openapiclient.NewPaymentSetupRequest(*openapiclient.NewAmount("Currency_example", int64(123)), "CountryCode_example", "MerchantAccount_example", "Reference_example", "ReturnUrl_example") // PaymentSetupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClassicCheckoutSDKApi.PostPaymentSession(context.Background()).IdempotencyKey(idempotencyKey).PaymentSetupRequest(paymentSetupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClassicCheckoutSDKApi.PostPaymentSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPaymentSession`: PaymentSetupResponse
    fmt.Fprintf(os.Stdout, "Response from `ClassicCheckoutSDKApi.PostPaymentSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPaymentSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | A unique identifier for the message with a maximum of 64 characters (we recommend a UUID). | 
 **paymentSetupRequest** | [**PaymentSetupRequest**](PaymentSetupRequest.md) |  | 

### Return type

[**PaymentSetupResponse**](PaymentSetupResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPaymentsResult

> PaymentVerificationResponse PostPaymentsResult(ctx).IdempotencyKey(idempotencyKey).PaymentVerificationRequest(paymentVerificationRequest).Execute()

Verify a payment result



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
    paymentVerificationRequest := *openapiclient.NewPaymentVerificationRequest("Payload_example") // PaymentVerificationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClassicCheckoutSDKApi.PostPaymentsResult(context.Background()).IdempotencyKey(idempotencyKey).PaymentVerificationRequest(paymentVerificationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClassicCheckoutSDKApi.PostPaymentsResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPaymentsResult`: PaymentVerificationResponse
    fmt.Fprintf(os.Stdout, "Response from `ClassicCheckoutSDKApi.PostPaymentsResult`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPaymentsResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | A unique identifier for the message with a maximum of 64 characters (we recommend a UUID). | 
 **paymentVerificationRequest** | [**PaymentVerificationRequest**](PaymentVerificationRequest.md) |  | 

### Return type

[**PaymentVerificationResponse**](PaymentVerificationResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

