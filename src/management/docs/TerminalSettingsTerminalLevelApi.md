# \TerminalSettingsTerminalLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTerminalsTerminalIdTerminalLogos**](TerminalSettingsTerminalLevelApi.md#GetTerminalsTerminalIdTerminalLogos) | **Get** /terminals/{terminalId}/terminalLogos | Get the terminal logo
[**GetTerminalsTerminalIdTerminalSettings**](TerminalSettingsTerminalLevelApi.md#GetTerminalsTerminalIdTerminalSettings) | **Get** /terminals/{terminalId}/terminalSettings | Get terminal settings
[**PatchTerminalsTerminalIdTerminalLogos**](TerminalSettingsTerminalLevelApi.md#PatchTerminalsTerminalIdTerminalLogos) | **Patch** /terminals/{terminalId}/terminalLogos | Update the logo
[**PatchTerminalsTerminalIdTerminalSettings**](TerminalSettingsTerminalLevelApi.md#PatchTerminalsTerminalIdTerminalSettings) | **Patch** /terminals/{terminalId}/terminalSettings | Update terminal settings



## GetTerminalsTerminalIdTerminalLogos

> Logo GetTerminalsTerminalIdTerminalLogos(ctx, terminalId).Execute()

Get the terminal logo



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
    terminalId := "terminalId_example" // string | The unique identifier of the payment terminal.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalSettingsTerminalLevelApi.GetTerminalsTerminalIdTerminalLogos(context.Background(), terminalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalSettingsTerminalLevelApi.GetTerminalsTerminalIdTerminalLogos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTerminalsTerminalIdTerminalLogos`: Logo
    fmt.Fprintf(os.Stdout, "Response from `TerminalSettingsTerminalLevelApi.GetTerminalsTerminalIdTerminalLogos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**terminalId** | **string** | The unique identifier of the payment terminal. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTerminalsTerminalIdTerminalLogosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Logo**](Logo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTerminalsTerminalIdTerminalSettings

> TerminalSettings GetTerminalsTerminalIdTerminalSettings(ctx, terminalId).Execute()

Get terminal settings



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
    terminalId := "terminalId_example" // string | The unique identifier of the payment terminal.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalSettingsTerminalLevelApi.GetTerminalsTerminalIdTerminalSettings(context.Background(), terminalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalSettingsTerminalLevelApi.GetTerminalsTerminalIdTerminalSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTerminalsTerminalIdTerminalSettings`: TerminalSettings
    fmt.Fprintf(os.Stdout, "Response from `TerminalSettingsTerminalLevelApi.GetTerminalsTerminalIdTerminalSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**terminalId** | **string** | The unique identifier of the payment terminal. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTerminalsTerminalIdTerminalSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TerminalSettings**](TerminalSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchTerminalsTerminalIdTerminalLogos

> Logo PatchTerminalsTerminalIdTerminalLogos(ctx, terminalId).Logo(logo).Execute()

Update the logo



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
    terminalId := "terminalId_example" // string | The unique identifier of the payment terminal.
    logo := *openapiclient.NewLogo() // Logo |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalSettingsTerminalLevelApi.PatchTerminalsTerminalIdTerminalLogos(context.Background(), terminalId).Logo(logo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalSettingsTerminalLevelApi.PatchTerminalsTerminalIdTerminalLogos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchTerminalsTerminalIdTerminalLogos`: Logo
    fmt.Fprintf(os.Stdout, "Response from `TerminalSettingsTerminalLevelApi.PatchTerminalsTerminalIdTerminalLogos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**terminalId** | **string** | The unique identifier of the payment terminal. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTerminalsTerminalIdTerminalLogosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logo** | [**Logo**](Logo.md) |  | 

### Return type

[**Logo**](Logo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchTerminalsTerminalIdTerminalSettings

> TerminalSettings PatchTerminalsTerminalIdTerminalSettings(ctx, terminalId).TerminalSettings(terminalSettings).Execute()

Update terminal settings



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
    terminalId := "terminalId_example" // string | The unique identifier of the payment terminal.
    terminalSettings := *openapiclient.NewTerminalSettings() // TerminalSettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalSettingsTerminalLevelApi.PatchTerminalsTerminalIdTerminalSettings(context.Background(), terminalId).TerminalSettings(terminalSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalSettingsTerminalLevelApi.PatchTerminalsTerminalIdTerminalSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchTerminalsTerminalIdTerminalSettings`: TerminalSettings
    fmt.Fprintf(os.Stdout, "Response from `TerminalSettingsTerminalLevelApi.PatchTerminalsTerminalIdTerminalSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**terminalId** | **string** | The unique identifier of the payment terminal. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTerminalsTerminalIdTerminalSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **terminalSettings** | [**TerminalSettings**](TerminalSettings.md) |  | 

### Return type

[**TerminalSettings**](TerminalSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

