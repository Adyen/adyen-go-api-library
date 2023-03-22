# \OrdersApi

All URIs are relative to *https://checkout-test.adyen.com/v70*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostOrders**](OrdersApi.md#PostOrders) | **Post** /orders | Create an order
[**PostOrdersCancel**](OrdersApi.md#PostOrdersCancel) | **Post** /orders/cancel | Cancel an order
[**PostPaymentMethodsBalance**](OrdersApi.md#PostPaymentMethodsBalance) | **Post** /paymentMethods/balance | Get the balance of a gift card



## PostOrders

> CheckoutCreateOrderResponse PostOrders(ctx).IdempotencyKey(idempotencyKey).CheckoutCreateOrderRequest(checkoutCreateOrderRequest).Execute()

Create an order



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
    checkoutCreateOrderRequest := *openapiclient.NewCheckoutCreateOrderRequest(*openapiclient.NewAmount("Currency_example", int64(123)), "MerchantAccount_example", "Reference_example") // CheckoutCreateOrderRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.PostOrders(context.Background()).IdempotencyKey(idempotencyKey).CheckoutCreateOrderRequest(checkoutCreateOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.PostOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostOrders`: CheckoutCreateOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.PostOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | A unique identifier for the message with a maximum of 64 characters (we recommend a UUID). | 
 **checkoutCreateOrderRequest** | [**CheckoutCreateOrderRequest**](CheckoutCreateOrderRequest.md) |  | 

### Return type

[**CheckoutCreateOrderResponse**](CheckoutCreateOrderResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrdersCancel

> CheckoutCancelOrderResponse PostOrdersCancel(ctx).IdempotencyKey(idempotencyKey).CheckoutCancelOrderRequest(checkoutCancelOrderRequest).Execute()

Cancel an order



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
    checkoutCancelOrderRequest := *openapiclient.NewCheckoutCancelOrderRequest("MerchantAccount_example", *openapiclient.NewCheckoutOrder("OrderData_example", "PspReference_example")) // CheckoutCancelOrderRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.PostOrdersCancel(context.Background()).IdempotencyKey(idempotencyKey).CheckoutCancelOrderRequest(checkoutCancelOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.PostOrdersCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostOrdersCancel`: CheckoutCancelOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.PostOrdersCancel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostOrdersCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | A unique identifier for the message with a maximum of 64 characters (we recommend a UUID). | 
 **checkoutCancelOrderRequest** | [**CheckoutCancelOrderRequest**](CheckoutCancelOrderRequest.md) |  | 

### Return type

[**CheckoutCancelOrderResponse**](CheckoutCancelOrderResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPaymentMethodsBalance

> CheckoutBalanceCheckResponse PostPaymentMethodsBalance(ctx).IdempotencyKey(idempotencyKey).CheckoutBalanceCheckRequest(checkoutBalanceCheckRequest).Execute()

Get the balance of a gift card



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
    checkoutBalanceCheckRequest := *openapiclient.NewCheckoutBalanceCheckRequest("MerchantAccount_example", map[string]string{"key": "Inner_example"}) // CheckoutBalanceCheckRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.PostPaymentMethodsBalance(context.Background()).IdempotencyKey(idempotencyKey).CheckoutBalanceCheckRequest(checkoutBalanceCheckRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.PostPaymentMethodsBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPaymentMethodsBalance`: CheckoutBalanceCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.PostPaymentMethodsBalance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPaymentMethodsBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | A unique identifier for the message with a maximum of 64 characters (we recommend a UUID). | 
 **checkoutBalanceCheckRequest** | [**CheckoutBalanceCheckRequest**](CheckoutBalanceCheckRequest.md) |  | 

### Return type

[**CheckoutBalanceCheckResponse**](CheckoutBalanceCheckResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

