# \TransferInstrumentsApi

All URIs are relative to *https://kyc-test.adyen.com/lem/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTransferInstrument**](TransferInstrumentsApi.md#DeleteTransferInstrument) | **Delete** /transferInstruments/{id} | Delete a transfer instrument
[**GetTransferInstrument**](TransferInstrumentsApi.md#GetTransferInstrument) | **Get** /transferInstruments/{id} | Get a transfer instrument
[**UpdateTransferInstrument**](TransferInstrumentsApi.md#UpdateTransferInstrument) | **Patch** /transferInstruments/{id} | Update a transfer instrument
[**CreateTransferInstrument**](TransferInstrumentsApi.md#CreateTransferInstrument) | **Post** /transferInstruments | Create a transfer instrument



## DeleteTransferInstrument

> interface{} DeleteTransferInstrument(ctx, id).Execute()

Delete a transfer instrument



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/adyen/adyen-go-api-library/v6/src/legalentity"
)

func main() {
    id := "id_example" // string | The unique identifier of the transfer instrument to be deleted.

    username := "YOUR_USERNAME"
    password := "YOUR_PASSWORD"
    env := openapiclient.TestEnv

    configuration, err := openapiclient.NewConfiguration(username, password, env)
    apiClient := openapiclient.NewAPIClient(configuration)

    resp, r, err := apiClient.TransferInstrumentsApi.DeleteTransferInstrument(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransferInstrumentsApi.DeleteTransferInstrument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTransferInstrument`: interface{}
    fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the transfer instrument to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTransferInstrumentsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransferInstrument

> TransferInstrument GetTransferInstrument(ctx, id).Execute()

Get a transfer instrument



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/adyen/adyen-go-api-library/v6/src/legalentity"
)

func main() {
    id := "id_example" // string | The unique identifier of the transfer instrument.

    username := "YOUR_USERNAME"
    password := "YOUR_PASSWORD"
    env := openapiclient.TestEnv

    configuration, err := openapiclient.NewConfiguration(username, password, env)
    apiClient := openapiclient.NewAPIClient(configuration)

    resp, r, err := apiClient.TransferInstrumentsApi.GetTransferInstrument(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransferInstrumentsApi.GetTransferInstrument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransferInstrument`: TransferInstrument
    fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the transfer instrument. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransferInstrumentsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransferInstrument**](TransferInstrument.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTransferInstrument

> TransferInstrument UpdateTransferInstrument(ctx, id).TransferInstrumentInfo(transferInstrumentInfo).Execute()

Update a transfer instrument



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/adyen/adyen-go-api-library/v6/src/legalentity"
)

func main() {
    id := "id_example" // string | The unique identifier of the transfer instrument.
    transferInstrumentInfo := *openapiclient.NewTransferInstrumentInfo(*openapiclient.NewBankAccountInfo("CountryCode_example", "CurrencyCode_example"), "LegalEntityId_example", "Type_example") // TransferInstrumentInfo |  (optional)

    username := "YOUR_USERNAME"
    password := "YOUR_PASSWORD"
    env := openapiclient.TestEnv

    configuration, err := openapiclient.NewConfiguration(username, password, env)
    apiClient := openapiclient.NewAPIClient(configuration)

    resp, r, err := apiClient.TransferInstrumentsApi.UpdateTransferInstrument(context.Background(), id).TransferInstrumentInfo(transferInstrumentInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransferInstrumentsApi.UpdateTransferInstrument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTransferInstrument`: TransferInstrument
    fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the transfer instrument. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTransferInstrumentsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transferInstrumentInfo** | [**TransferInstrumentInfo**](TransferInstrumentInfo.md) |  | 

### Return type

[**TransferInstrument**](TransferInstrument.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTransferInstrument

> TransferInstrument CreateTransferInstrument(ctx).TransferInstrumentInfo(transferInstrumentInfo).Execute()

Create a transfer instrument



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/adyen/adyen-go-api-library/v6/src/legalentity"
)

func main() {
    transferInstrumentInfo := *openapiclient.NewTransferInstrumentInfo(*openapiclient.NewBankAccountInfo("CountryCode_example", "CurrencyCode_example"), "LegalEntityId_example", "Type_example") // TransferInstrumentInfo |  (optional)

    username := "YOUR_USERNAME"
    password := "YOUR_PASSWORD"
    env := openapiclient.TestEnv

    configuration, err := openapiclient.NewConfiguration(username, password, env)
    apiClient := openapiclient.NewAPIClient(configuration)

    resp, r, err := apiClient.TransferInstrumentsApi.CreateTransferInstrument(context.Background()).TransferInstrumentInfo(transferInstrumentInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransferInstrumentsApi.CreateTransferInstrument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTransferInstrument`: TransferInstrument
    fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTransferInstrumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferInstrumentInfo** | [**TransferInstrumentInfo**](TransferInstrumentInfo.md) |  | 

### Return type

[**TransferInstrument**](TransferInstrument.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

