# \TerminalOrdersMerchantLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMerchantsMerchantIdBillingEntities**](TerminalOrdersMerchantLevelApi.md#GetMerchantsMerchantIdBillingEntities) | **Get** /merchants/{merchantId}/billingEntities | Get a list of billing entities
[**GetMerchantsMerchantIdShippingLocations**](TerminalOrdersMerchantLevelApi.md#GetMerchantsMerchantIdShippingLocations) | **Get** /merchants/{merchantId}/shippingLocations | Get a list of shipping locations
[**GetMerchantsMerchantIdTerminalModels**](TerminalOrdersMerchantLevelApi.md#GetMerchantsMerchantIdTerminalModels) | **Get** /merchants/{merchantId}/terminalModels | Get a list of terminal models
[**GetMerchantsMerchantIdTerminalOrders**](TerminalOrdersMerchantLevelApi.md#GetMerchantsMerchantIdTerminalOrders) | **Get** /merchants/{merchantId}/terminalOrders | Get a list of orders
[**GetMerchantsMerchantIdTerminalOrdersOrderId**](TerminalOrdersMerchantLevelApi.md#GetMerchantsMerchantIdTerminalOrdersOrderId) | **Get** /merchants/{merchantId}/terminalOrders/{orderId} | Get an order
[**GetMerchantsMerchantIdTerminalProducts**](TerminalOrdersMerchantLevelApi.md#GetMerchantsMerchantIdTerminalProducts) | **Get** /merchants/{merchantId}/terminalProducts | Get a list of terminal products
[**PatchMerchantsMerchantIdTerminalOrdersOrderId**](TerminalOrdersMerchantLevelApi.md#PatchMerchantsMerchantIdTerminalOrdersOrderId) | **Patch** /merchants/{merchantId}/terminalOrders/{orderId} | Update an order
[**PostMerchantsMerchantIdShippingLocations**](TerminalOrdersMerchantLevelApi.md#PostMerchantsMerchantIdShippingLocations) | **Post** /merchants/{merchantId}/shippingLocations | Create a shipping location
[**PostMerchantsMerchantIdTerminalOrders**](TerminalOrdersMerchantLevelApi.md#PostMerchantsMerchantIdTerminalOrders) | **Post** /merchants/{merchantId}/terminalOrders | Create an order
[**PostMerchantsMerchantIdTerminalOrdersOrderIdCancel**](TerminalOrdersMerchantLevelApi.md#PostMerchantsMerchantIdTerminalOrdersOrderIdCancel) | **Post** /merchants/{merchantId}/terminalOrders/{orderId}/cancel | Cancel an order



## GetMerchantsMerchantIdBillingEntities

> BillingEntitiesResponse GetMerchantsMerchantIdBillingEntities(ctx, merchantId).Name(name).Execute()

Get a list of billing entities



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
    name := "name_example" // string | The name of the billing entity. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalOrdersMerchantLevelApi.GetMerchantsMerchantIdBillingEntities(context.Background(), merchantId).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalOrdersMerchantLevelApi.GetMerchantsMerchantIdBillingEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsMerchantIdBillingEntities`: BillingEntitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `TerminalOrdersMerchantLevelApi.GetMerchantsMerchantIdBillingEntities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsMerchantIdBillingEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | The name of the billing entity. | 

### Return type

[**BillingEntitiesResponse**](BillingEntitiesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMerchantsMerchantIdShippingLocations

> ShippingLocationsResponse GetMerchantsMerchantIdShippingLocations(ctx, merchantId).Name(name).Offset(offset).Limit(limit).Execute()

Get a list of shipping locations



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
    name := "name_example" // string | The name of the shipping location. (optional)
    offset := int32(56) // int32 | The number of locations to skip. (optional)
    limit := int32(56) // int32 | The number of locations to return. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalOrdersMerchantLevelApi.GetMerchantsMerchantIdShippingLocations(context.Background(), merchantId).Name(name).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalOrdersMerchantLevelApi.GetMerchantsMerchantIdShippingLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsMerchantIdShippingLocations`: ShippingLocationsResponse
    fmt.Fprintf(os.Stdout, "Response from `TerminalOrdersMerchantLevelApi.GetMerchantsMerchantIdShippingLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsMerchantIdShippingLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | The name of the shipping location. | 
 **offset** | **int32** | The number of locations to skip. | 
 **limit** | **int32** | The number of locations to return. | 

### Return type

[**ShippingLocationsResponse**](ShippingLocationsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMerchantsMerchantIdTerminalModels

> TerminalModelsResponse GetMerchantsMerchantIdTerminalModels(ctx, merchantId).Execute()

Get a list of terminal models



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalOrdersMerchantLevelApi.GetMerchantsMerchantIdTerminalModels(context.Background(), merchantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalOrdersMerchantLevelApi.GetMerchantsMerchantIdTerminalModels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsMerchantIdTerminalModels`: TerminalModelsResponse
    fmt.Fprintf(os.Stdout, "Response from `TerminalOrdersMerchantLevelApi.GetMerchantsMerchantIdTerminalModels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsMerchantIdTerminalModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TerminalModelsResponse**](TerminalModelsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMerchantsMerchantIdTerminalOrders

> TerminalOrdersResponse GetMerchantsMerchantIdTerminalOrders(ctx, merchantId).CustomerOrderReference(customerOrderReference).Status(status).Offset(offset).Limit(limit).Execute()

Get a list of orders



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
    merchantId := "merchantId_example" // string | 
    customerOrderReference := "customerOrderReference_example" // string | Your purchase order number. (optional)
    status := "status_example" // string | The order status. Possible values (not case-sensitive): Placed, Confirmed, Cancelled, Shipped, Delivered. (optional)
    offset := int32(56) // int32 | The number of orders to skip. (optional)
    limit := int32(56) // int32 | The number of orders to return. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalOrdersMerchantLevelApi.GetMerchantsMerchantIdTerminalOrders(context.Background(), merchantId).CustomerOrderReference(customerOrderReference).Status(status).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalOrdersMerchantLevelApi.GetMerchantsMerchantIdTerminalOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsMerchantIdTerminalOrders`: TerminalOrdersResponse
    fmt.Fprintf(os.Stdout, "Response from `TerminalOrdersMerchantLevelApi.GetMerchantsMerchantIdTerminalOrders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsMerchantIdTerminalOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerOrderReference** | **string** | Your purchase order number. | 
 **status** | **string** | The order status. Possible values (not case-sensitive): Placed, Confirmed, Cancelled, Shipped, Delivered. | 
 **offset** | **int32** | The number of orders to skip. | 
 **limit** | **int32** | The number of orders to return. | 

### Return type

[**TerminalOrdersResponse**](TerminalOrdersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMerchantsMerchantIdTerminalOrdersOrderId

> TerminalOrder GetMerchantsMerchantIdTerminalOrdersOrderId(ctx, merchantId, orderId).Execute()

Get an order



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
    orderId := "orderId_example" // string | The unique identifier of the order.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalOrdersMerchantLevelApi.GetMerchantsMerchantIdTerminalOrdersOrderId(context.Background(), merchantId, orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalOrdersMerchantLevelApi.GetMerchantsMerchantIdTerminalOrdersOrderId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsMerchantIdTerminalOrdersOrderId`: TerminalOrder
    fmt.Fprintf(os.Stdout, "Response from `TerminalOrdersMerchantLevelApi.GetMerchantsMerchantIdTerminalOrdersOrderId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**orderId** | **string** | The unique identifier of the order. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsMerchantIdTerminalOrdersOrderIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TerminalOrder**](TerminalOrder.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMerchantsMerchantIdTerminalProducts

> TerminalProductsResponse GetMerchantsMerchantIdTerminalProducts(ctx, merchantId).Country(country).TerminalModelId(terminalModelId).Offset(offset).Limit(limit).Execute()

Get a list of terminal products



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
    country := "country_example" // string | The country to return products for, in [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) format. For example, **US** (optional)
    terminalModelId := "terminalModelId_example" // string | The terminal model to return products for. Use the ID returned in the [GET `/terminalModels`](https://docs.adyen.com/api-explorer/#/ManagementService/latest/get/terminalModels) response. For example, **Verifone.M400** (optional)
    offset := int32(56) // int32 | The number of products to skip. (optional)
    limit := int32(56) // int32 | The number of products to return. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalOrdersMerchantLevelApi.GetMerchantsMerchantIdTerminalProducts(context.Background(), merchantId).Country(country).TerminalModelId(terminalModelId).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalOrdersMerchantLevelApi.GetMerchantsMerchantIdTerminalProducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsMerchantIdTerminalProducts`: TerminalProductsResponse
    fmt.Fprintf(os.Stdout, "Response from `TerminalOrdersMerchantLevelApi.GetMerchantsMerchantIdTerminalProducts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsMerchantIdTerminalProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **country** | **string** | The country to return products for, in [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) format. For example, **US** | 
 **terminalModelId** | **string** | The terminal model to return products for. Use the ID returned in the [GET &#x60;/terminalModels&#x60;](https://docs.adyen.com/api-explorer/#/ManagementService/latest/get/terminalModels) response. For example, **Verifone.M400** | 
 **offset** | **int32** | The number of products to skip. | 
 **limit** | **int32** | The number of products to return. | 

### Return type

[**TerminalProductsResponse**](TerminalProductsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMerchantsMerchantIdTerminalOrdersOrderId

> TerminalOrder PatchMerchantsMerchantIdTerminalOrdersOrderId(ctx, merchantId, orderId).TerminalOrderRequest(terminalOrderRequest).Execute()

Update an order



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
    orderId := "orderId_example" // string | The unique identifier of the order.
    terminalOrderRequest := *openapiclient.NewTerminalOrderRequest() // TerminalOrderRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalOrdersMerchantLevelApi.PatchMerchantsMerchantIdTerminalOrdersOrderId(context.Background(), merchantId, orderId).TerminalOrderRequest(terminalOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalOrdersMerchantLevelApi.PatchMerchantsMerchantIdTerminalOrdersOrderId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMerchantsMerchantIdTerminalOrdersOrderId`: TerminalOrder
    fmt.Fprintf(os.Stdout, "Response from `TerminalOrdersMerchantLevelApi.PatchMerchantsMerchantIdTerminalOrdersOrderId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**orderId** | **string** | The unique identifier of the order. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMerchantsMerchantIdTerminalOrdersOrderIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **terminalOrderRequest** | [**TerminalOrderRequest**](TerminalOrderRequest.md) |  | 

### Return type

[**TerminalOrder**](TerminalOrder.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMerchantsMerchantIdShippingLocations

> ShippingLocation PostMerchantsMerchantIdShippingLocations(ctx, merchantId).ShippingLocation(shippingLocation).Execute()

Create a shipping location



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
    shippingLocation := *openapiclient.NewShippingLocation() // ShippingLocation |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalOrdersMerchantLevelApi.PostMerchantsMerchantIdShippingLocations(context.Background(), merchantId).ShippingLocation(shippingLocation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalOrdersMerchantLevelApi.PostMerchantsMerchantIdShippingLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMerchantsMerchantIdShippingLocations`: ShippingLocation
    fmt.Fprintf(os.Stdout, "Response from `TerminalOrdersMerchantLevelApi.PostMerchantsMerchantIdShippingLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMerchantsMerchantIdShippingLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shippingLocation** | [**ShippingLocation**](ShippingLocation.md) |  | 

### Return type

[**ShippingLocation**](ShippingLocation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMerchantsMerchantIdTerminalOrders

> TerminalOrder PostMerchantsMerchantIdTerminalOrders(ctx, merchantId).TerminalOrderRequest(terminalOrderRequest).Execute()

Create an order



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
    terminalOrderRequest := *openapiclient.NewTerminalOrderRequest() // TerminalOrderRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalOrdersMerchantLevelApi.PostMerchantsMerchantIdTerminalOrders(context.Background(), merchantId).TerminalOrderRequest(terminalOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalOrdersMerchantLevelApi.PostMerchantsMerchantIdTerminalOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMerchantsMerchantIdTerminalOrders`: TerminalOrder
    fmt.Fprintf(os.Stdout, "Response from `TerminalOrdersMerchantLevelApi.PostMerchantsMerchantIdTerminalOrders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMerchantsMerchantIdTerminalOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **terminalOrderRequest** | [**TerminalOrderRequest**](TerminalOrderRequest.md) |  | 

### Return type

[**TerminalOrder**](TerminalOrder.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMerchantsMerchantIdTerminalOrdersOrderIdCancel

> TerminalOrder PostMerchantsMerchantIdTerminalOrdersOrderIdCancel(ctx, merchantId, orderId).Execute()

Cancel an order



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
    orderId := "orderId_example" // string | The unique identifier of the order.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalOrdersMerchantLevelApi.PostMerchantsMerchantIdTerminalOrdersOrderIdCancel(context.Background(), merchantId, orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalOrdersMerchantLevelApi.PostMerchantsMerchantIdTerminalOrdersOrderIdCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMerchantsMerchantIdTerminalOrdersOrderIdCancel`: TerminalOrder
    fmt.Fprintf(os.Stdout, "Response from `TerminalOrdersMerchantLevelApi.PostMerchantsMerchantIdTerminalOrdersOrderIdCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**orderId** | **string** | The unique identifier of the order. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMerchantsMerchantIdTerminalOrdersOrderIdCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TerminalOrder**](TerminalOrder.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

