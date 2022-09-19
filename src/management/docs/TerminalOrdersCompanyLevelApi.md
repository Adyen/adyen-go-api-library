# \TerminalOrdersCompanyLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompaniesCompanyIdBillingEntities**](TerminalOrdersCompanyLevelApi.md#GetCompaniesCompanyIdBillingEntities) | **Get** /companies/{companyId}/billingEntities | Get a list of billing entities
[**GetCompaniesCompanyIdShippingLocations**](TerminalOrdersCompanyLevelApi.md#GetCompaniesCompanyIdShippingLocations) | **Get** /companies/{companyId}/shippingLocations | Get a list of shipping locations
[**GetCompaniesCompanyIdTerminalModels**](TerminalOrdersCompanyLevelApi.md#GetCompaniesCompanyIdTerminalModels) | **Get** /companies/{companyId}/terminalModels | Get a list of terminal models
[**GetCompaniesCompanyIdTerminalOrders**](TerminalOrdersCompanyLevelApi.md#GetCompaniesCompanyIdTerminalOrders) | **Get** /companies/{companyId}/terminalOrders | Get a list of orders
[**GetCompaniesCompanyIdTerminalOrdersOrderId**](TerminalOrdersCompanyLevelApi.md#GetCompaniesCompanyIdTerminalOrdersOrderId) | **Get** /companies/{companyId}/terminalOrders/{orderId} | Get an order
[**GetCompaniesCompanyIdTerminalProducts**](TerminalOrdersCompanyLevelApi.md#GetCompaniesCompanyIdTerminalProducts) | **Get** /companies/{companyId}/terminalProducts | Get a list of terminal products
[**PatchCompaniesCompanyIdTerminalOrdersOrderId**](TerminalOrdersCompanyLevelApi.md#PatchCompaniesCompanyIdTerminalOrdersOrderId) | **Patch** /companies/{companyId}/terminalOrders/{orderId} | Update an order
[**PostCompaniesCompanyIdShippingLocations**](TerminalOrdersCompanyLevelApi.md#PostCompaniesCompanyIdShippingLocations) | **Post** /companies/{companyId}/shippingLocations | Create a shipping location
[**PostCompaniesCompanyIdTerminalOrders**](TerminalOrdersCompanyLevelApi.md#PostCompaniesCompanyIdTerminalOrders) | **Post** /companies/{companyId}/terminalOrders | Create an order
[**PostCompaniesCompanyIdTerminalOrdersOrderIdCancel**](TerminalOrdersCompanyLevelApi.md#PostCompaniesCompanyIdTerminalOrdersOrderIdCancel) | **Post** /companies/{companyId}/terminalOrders/{orderId}/cancel | Cancel an order



## GetCompaniesCompanyIdBillingEntities

> BillingEntitiesResponse GetCompaniesCompanyIdBillingEntities(ctx, companyId).Name(name).Execute()

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
    companyId := "companyId_example" // string | The unique identifier of the company account.
    name := "name_example" // string | The name of the billing entity. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalOrdersCompanyLevelApi.GetCompaniesCompanyIdBillingEntities(context.Background(), companyId).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalOrdersCompanyLevelApi.GetCompaniesCompanyIdBillingEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompaniesCompanyIdBillingEntities`: BillingEntitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `TerminalOrdersCompanyLevelApi.GetCompaniesCompanyIdBillingEntities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesCompanyIdBillingEntitiesRequest struct via the builder pattern


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


## GetCompaniesCompanyIdShippingLocations

> ShippingLocationsResponse GetCompaniesCompanyIdShippingLocations(ctx, companyId).Name(name).Offset(offset).Limit(limit).Execute()

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
    companyId := "companyId_example" // string | The unique identifier of the company account.
    name := "name_example" // string | The name of the shipping location. (optional)
    offset := int32(56) // int32 | The number of locations to skip. (optional)
    limit := int32(56) // int32 | The number of locations to return. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalOrdersCompanyLevelApi.GetCompaniesCompanyIdShippingLocations(context.Background(), companyId).Name(name).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalOrdersCompanyLevelApi.GetCompaniesCompanyIdShippingLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompaniesCompanyIdShippingLocations`: ShippingLocationsResponse
    fmt.Fprintf(os.Stdout, "Response from `TerminalOrdersCompanyLevelApi.GetCompaniesCompanyIdShippingLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesCompanyIdShippingLocationsRequest struct via the builder pattern


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


## GetCompaniesCompanyIdTerminalModels

> TerminalModelsResponse GetCompaniesCompanyIdTerminalModels(ctx, companyId).Execute()

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
    companyId := "companyId_example" // string | The unique identifier of the company account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalOrdersCompanyLevelApi.GetCompaniesCompanyIdTerminalModels(context.Background(), companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalOrdersCompanyLevelApi.GetCompaniesCompanyIdTerminalModels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompaniesCompanyIdTerminalModels`: TerminalModelsResponse
    fmt.Fprintf(os.Stdout, "Response from `TerminalOrdersCompanyLevelApi.GetCompaniesCompanyIdTerminalModels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesCompanyIdTerminalModelsRequest struct via the builder pattern


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


## GetCompaniesCompanyIdTerminalOrders

> TerminalOrdersResponse GetCompaniesCompanyIdTerminalOrders(ctx, companyId).CustomerOrderReference(customerOrderReference).Status(status).Offset(offset).Limit(limit).Execute()

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
    companyId := "companyId_example" // string | The unique identifier of the company account.
    customerOrderReference := "customerOrderReference_example" // string | Your purchase order number. (optional)
    status := "status_example" // string | The order status. Possible values (not case-sensitive): Placed, Confirmed, Cancelled, Shipped, Delivered. (optional)
    offset := int32(56) // int32 | The number of orders to skip. (optional)
    limit := int32(56) // int32 | The number of orders to return. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalOrdersCompanyLevelApi.GetCompaniesCompanyIdTerminalOrders(context.Background(), companyId).CustomerOrderReference(customerOrderReference).Status(status).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalOrdersCompanyLevelApi.GetCompaniesCompanyIdTerminalOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompaniesCompanyIdTerminalOrders`: TerminalOrdersResponse
    fmt.Fprintf(os.Stdout, "Response from `TerminalOrdersCompanyLevelApi.GetCompaniesCompanyIdTerminalOrders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesCompanyIdTerminalOrdersRequest struct via the builder pattern


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


## GetCompaniesCompanyIdTerminalOrdersOrderId

> TerminalOrder GetCompaniesCompanyIdTerminalOrdersOrderId(ctx, companyId, orderId).Execute()

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
    companyId := "companyId_example" // string | The unique identifier of the company account.
    orderId := "orderId_example" // string | The unique identifier of the order.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalOrdersCompanyLevelApi.GetCompaniesCompanyIdTerminalOrdersOrderId(context.Background(), companyId, orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalOrdersCompanyLevelApi.GetCompaniesCompanyIdTerminalOrdersOrderId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompaniesCompanyIdTerminalOrdersOrderId`: TerminalOrder
    fmt.Fprintf(os.Stdout, "Response from `TerminalOrdersCompanyLevelApi.GetCompaniesCompanyIdTerminalOrdersOrderId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 
**orderId** | **string** | The unique identifier of the order. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesCompanyIdTerminalOrdersOrderIdRequest struct via the builder pattern


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


## GetCompaniesCompanyIdTerminalProducts

> TerminalProductsResponse GetCompaniesCompanyIdTerminalProducts(ctx, companyId).Country(country).TerminalModelId(terminalModelId).Offset(offset).Limit(limit).Execute()

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
    companyId := "companyId_example" // string | The unique identifier of the company account.
    country := "country_example" // string | The country to return products for, in [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) format. For example, **US** (optional)
    terminalModelId := "terminalModelId_example" // string | The terminal model to return products for. Use the ID returned in the [GET `/terminalModels`](https://docs.adyen.com/api-explorer/#/ManagementService/latest/get/terminalModels) response. For example, **Verifone.M400** (optional)
    offset := int32(56) // int32 | The number of products to skip. (optional)
    limit := int32(56) // int32 | The number of products to return. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalOrdersCompanyLevelApi.GetCompaniesCompanyIdTerminalProducts(context.Background(), companyId).Country(country).TerminalModelId(terminalModelId).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalOrdersCompanyLevelApi.GetCompaniesCompanyIdTerminalProducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompaniesCompanyIdTerminalProducts`: TerminalProductsResponse
    fmt.Fprintf(os.Stdout, "Response from `TerminalOrdersCompanyLevelApi.GetCompaniesCompanyIdTerminalProducts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesCompanyIdTerminalProductsRequest struct via the builder pattern


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


## PatchCompaniesCompanyIdTerminalOrdersOrderId

> TerminalOrder PatchCompaniesCompanyIdTerminalOrdersOrderId(ctx, companyId, orderId).TerminalOrderRequest(terminalOrderRequest).Execute()

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
    companyId := "companyId_example" // string | The unique identifier of the company account.
    orderId := "orderId_example" // string | The unique identifier of the order.
    terminalOrderRequest := *openapiclient.NewTerminalOrderRequest() // TerminalOrderRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalOrdersCompanyLevelApi.PatchCompaniesCompanyIdTerminalOrdersOrderId(context.Background(), companyId, orderId).TerminalOrderRequest(terminalOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalOrdersCompanyLevelApi.PatchCompaniesCompanyIdTerminalOrdersOrderId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCompaniesCompanyIdTerminalOrdersOrderId`: TerminalOrder
    fmt.Fprintf(os.Stdout, "Response from `TerminalOrdersCompanyLevelApi.PatchCompaniesCompanyIdTerminalOrdersOrderId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 
**orderId** | **string** | The unique identifier of the order. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompaniesCompanyIdTerminalOrdersOrderIdRequest struct via the builder pattern


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


## PostCompaniesCompanyIdShippingLocations

> ShippingLocation PostCompaniesCompanyIdShippingLocations(ctx, companyId).ShippingLocation(shippingLocation).Execute()

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
    companyId := "companyId_example" // string | The unique identifier of the company account.
    shippingLocation := *openapiclient.NewShippingLocation() // ShippingLocation |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalOrdersCompanyLevelApi.PostCompaniesCompanyIdShippingLocations(context.Background(), companyId).ShippingLocation(shippingLocation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalOrdersCompanyLevelApi.PostCompaniesCompanyIdShippingLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCompaniesCompanyIdShippingLocations`: ShippingLocation
    fmt.Fprintf(os.Stdout, "Response from `TerminalOrdersCompanyLevelApi.PostCompaniesCompanyIdShippingLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompaniesCompanyIdShippingLocationsRequest struct via the builder pattern


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


## PostCompaniesCompanyIdTerminalOrders

> TerminalOrder PostCompaniesCompanyIdTerminalOrders(ctx, companyId).TerminalOrderRequest(terminalOrderRequest).Execute()

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
    companyId := "companyId_example" // string | The unique identifier of the company account.
    terminalOrderRequest := *openapiclient.NewTerminalOrderRequest() // TerminalOrderRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalOrdersCompanyLevelApi.PostCompaniesCompanyIdTerminalOrders(context.Background(), companyId).TerminalOrderRequest(terminalOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalOrdersCompanyLevelApi.PostCompaniesCompanyIdTerminalOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCompaniesCompanyIdTerminalOrders`: TerminalOrder
    fmt.Fprintf(os.Stdout, "Response from `TerminalOrdersCompanyLevelApi.PostCompaniesCompanyIdTerminalOrders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompaniesCompanyIdTerminalOrdersRequest struct via the builder pattern


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


## PostCompaniesCompanyIdTerminalOrdersOrderIdCancel

> TerminalOrder PostCompaniesCompanyIdTerminalOrdersOrderIdCancel(ctx, companyId, orderId).Execute()

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
    companyId := "companyId_example" // string | The unique identifier of the company account.
    orderId := "orderId_example" // string | The unique identifier of the order.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalOrdersCompanyLevelApi.PostCompaniesCompanyIdTerminalOrdersOrderIdCancel(context.Background(), companyId, orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalOrdersCompanyLevelApi.PostCompaniesCompanyIdTerminalOrdersOrderIdCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCompaniesCompanyIdTerminalOrdersOrderIdCancel`: TerminalOrder
    fmt.Fprintf(os.Stdout, "Response from `TerminalOrdersCompanyLevelApi.PostCompaniesCompanyIdTerminalOrdersOrderIdCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 
**orderId** | **string** | The unique identifier of the order. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompaniesCompanyIdTerminalOrdersOrderIdCancelRequest struct via the builder pattern


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

