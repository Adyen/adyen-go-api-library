# \UtilityApi

All URIs are relative to *https://checkout-test.adyen.com/v70*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostApplePaySessions**](UtilityApi.md#PostApplePaySessions) | **Post** /applePay/sessions | Get an Apple Pay session
[**PostOriginKeys**](UtilityApi.md#PostOriginKeys) | **Post** /originKeys | Create originKey values for domains



## PostApplePaySessions

> ApplePaySessionResponse PostApplePaySessions(ctx).IdempotencyKey(idempotencyKey).CreateApplePaySessionRequest(createApplePaySessionRequest).Execute()

Get an Apple Pay session



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
    createApplePaySessionRequest := *openapiclient.NewCreateApplePaySessionRequest("DisplayName_example", "DomainName_example", "MerchantIdentifier_example") // CreateApplePaySessionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UtilityApi.PostApplePaySessions(context.Background()).IdempotencyKey(idempotencyKey).CreateApplePaySessionRequest(createApplePaySessionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilityApi.PostApplePaySessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostApplePaySessions`: ApplePaySessionResponse
    fmt.Fprintf(os.Stdout, "Response from `UtilityApi.PostApplePaySessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApplePaySessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | A unique identifier for the message with a maximum of 64 characters (we recommend a UUID). | 
 **createApplePaySessionRequest** | [**CreateApplePaySessionRequest**](CreateApplePaySessionRequest.md) |  | 

### Return type

[**ApplePaySessionResponse**](ApplePaySessionResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOriginKeys

> CheckoutUtilityResponse PostOriginKeys(ctx).IdempotencyKey(idempotencyKey).CheckoutUtilityRequest(checkoutUtilityRequest).Execute()

Create originKey values for domains



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
    checkoutUtilityRequest := *openapiclient.NewCheckoutUtilityRequest([]string{"OriginDomains_example"}) // CheckoutUtilityRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UtilityApi.PostOriginKeys(context.Background()).IdempotencyKey(idempotencyKey).CheckoutUtilityRequest(checkoutUtilityRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilityApi.PostOriginKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostOriginKeys`: CheckoutUtilityResponse
    fmt.Fprintf(os.Stdout, "Response from `UtilityApi.PostOriginKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostOriginKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | A unique identifier for the message with a maximum of 64 characters (we recommend a UUID). | 
 **checkoutUtilityRequest** | [**CheckoutUtilityRequest**](CheckoutUtilityRequest.md) |  | 

### Return type

[**CheckoutUtilityResponse**](CheckoutUtilityResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

