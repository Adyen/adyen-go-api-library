# \PaymentsApi

All URIs are relative to *https://checkout-test.adyen.com/v70*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCardDetails**](PaymentsApi.md#PostCardDetails) | **Post** /cardDetails | Get the list of brands on the card
[**PostDonations**](PaymentsApi.md#PostDonations) | **Post** /donations | Start a transaction for donations
[**PostPaymentMethods**](PaymentsApi.md#PostPaymentMethods) | **Post** /paymentMethods | Get a list of available payment methods
[**PostPayments**](PaymentsApi.md#PostPayments) | **Post** /payments | Start a transaction
[**PostPaymentsDetails**](PaymentsApi.md#PostPaymentsDetails) | **Post** /payments/details | Submit details for a payment
[**PostSessions**](PaymentsApi.md#PostSessions) | **Post** /sessions | Create a payment session



## PostCardDetails

> CardDetailsResponse PostCardDetails(ctx).IdempotencyKey(idempotencyKey).CardDetailsRequest(cardDetailsRequest).Execute()

Get the list of brands on the card



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
    cardDetailsRequest := *openapiclient.NewCardDetailsRequest("CardNumber_example", "MerchantAccount_example") // CardDetailsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsApi.PostCardDetails(context.Background()).IdempotencyKey(idempotencyKey).CardDetailsRequest(cardDetailsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.PostCardDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCardDetails`: CardDetailsResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.PostCardDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCardDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | A unique identifier for the message with a maximum of 64 characters (we recommend a UUID). | 
 **cardDetailsRequest** | [**CardDetailsRequest**](CardDetailsRequest.md) |  | 

### Return type

[**CardDetailsResponse**](CardDetailsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDonations

> DonationResponse PostDonations(ctx).IdempotencyKey(idempotencyKey).PaymentDonationRequest(paymentDonationRequest).Execute()

Start a transaction for donations



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
    paymentDonationRequest := *openapiclient.NewPaymentDonationRequest(*openapiclient.NewAmount("Currency_example", int64(123)), "DonationAccount_example", "MerchantAccount_example", openapiclient.PaymentDonationRequest_paymentMethod{AchDetails: openapiclient.NewAchDetails("BankAccountNumber_example")}, "Reference_example", "ReturnUrl_example") // PaymentDonationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsApi.PostDonations(context.Background()).IdempotencyKey(idempotencyKey).PaymentDonationRequest(paymentDonationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.PostDonations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostDonations`: DonationResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.PostDonations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDonationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | A unique identifier for the message with a maximum of 64 characters (we recommend a UUID). | 
 **paymentDonationRequest** | [**PaymentDonationRequest**](PaymentDonationRequest.md) |  | 

### Return type

[**DonationResponse**](DonationResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPaymentMethods

> PaymentMethodsResponse PostPaymentMethods(ctx).IdempotencyKey(idempotencyKey).PaymentMethodsRequest(paymentMethodsRequest).Execute()

Get a list of available payment methods



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
    paymentMethodsRequest := *openapiclient.NewPaymentMethodsRequest("MerchantAccount_example") // PaymentMethodsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsApi.PostPaymentMethods(context.Background()).IdempotencyKey(idempotencyKey).PaymentMethodsRequest(paymentMethodsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.PostPaymentMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPaymentMethods`: PaymentMethodsResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.PostPaymentMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPaymentMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | A unique identifier for the message with a maximum of 64 characters (we recommend a UUID). | 
 **paymentMethodsRequest** | [**PaymentMethodsRequest**](PaymentMethodsRequest.md) |  | 

### Return type

[**PaymentMethodsResponse**](PaymentMethodsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPayments

> PaymentResponse PostPayments(ctx).IdempotencyKey(idempotencyKey).PaymentRequest(paymentRequest).Execute()

Start a transaction



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
    paymentRequest := *openapiclient.NewPaymentRequest(*openapiclient.NewAmount("Currency_example", int64(123)), "MerchantAccount_example", openapiclient.PaymentDonationRequest_paymentMethod{AchDetails: openapiclient.NewAchDetails("BankAccountNumber_example")}, "Reference_example", "ReturnUrl_example") // PaymentRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsApi.PostPayments(context.Background()).IdempotencyKey(idempotencyKey).PaymentRequest(paymentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.PostPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPayments`: PaymentResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.PostPayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | A unique identifier for the message with a maximum of 64 characters (we recommend a UUID). | 
 **paymentRequest** | [**PaymentRequest**](PaymentRequest.md) |  | 

### Return type

[**PaymentResponse**](PaymentResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPaymentsDetails

> PaymentDetailsResponse PostPaymentsDetails(ctx).IdempotencyKey(idempotencyKey).DetailsRequest(detailsRequest).Execute()

Submit details for a payment



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
    detailsRequest := *openapiclient.NewDetailsRequest(*openapiclient.NewPaymentCompletionDetails()) // DetailsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsApi.PostPaymentsDetails(context.Background()).IdempotencyKey(idempotencyKey).DetailsRequest(detailsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.PostPaymentsDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPaymentsDetails`: PaymentDetailsResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.PostPaymentsDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPaymentsDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | A unique identifier for the message with a maximum of 64 characters (we recommend a UUID). | 
 **detailsRequest** | [**DetailsRequest**](DetailsRequest.md) |  | 

### Return type

[**PaymentDetailsResponse**](PaymentDetailsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSessions

> CreateCheckoutSessionResponse PostSessions(ctx).IdempotencyKey(idempotencyKey).CreateCheckoutSessionRequest(createCheckoutSessionRequest).Execute()

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
    createCheckoutSessionRequest := *openapiclient.NewCreateCheckoutSessionRequest(*openapiclient.NewAmount("Currency_example", int64(123)), "MerchantAccount_example", "Reference_example", "ReturnUrl_example") // CreateCheckoutSessionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsApi.PostSessions(context.Background()).IdempotencyKey(idempotencyKey).CreateCheckoutSessionRequest(createCheckoutSessionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.PostSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSessions`: CreateCheckoutSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.PostSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | A unique identifier for the message with a maximum of 64 characters (we recommend a UUID). | 
 **createCheckoutSessionRequest** | [**CreateCheckoutSessionRequest**](CreateCheckoutSessionRequest.md) |  | 

### Return type

[**CreateCheckoutSessionResponse**](CreateCheckoutSessionResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

