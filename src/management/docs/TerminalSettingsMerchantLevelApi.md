# \TerminalSettingsMerchantLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMerchantsMerchantIdTerminalLogos**](TerminalSettingsMerchantLevelApi.md#GetMerchantsMerchantIdTerminalLogos) | **Get** /merchants/{merchantId}/terminalLogos | Get the terminal logo
[**GetMerchantsMerchantIdTerminalSettings**](TerminalSettingsMerchantLevelApi.md#GetMerchantsMerchantIdTerminalSettings) | **Get** /merchants/{merchantId}/terminalSettings | Get terminal settings
[**PatchMerchantsMerchantIdTerminalLogos**](TerminalSettingsMerchantLevelApi.md#PatchMerchantsMerchantIdTerminalLogos) | **Patch** /merchants/{merchantId}/terminalLogos | Update the terminal logo
[**PatchMerchantsMerchantIdTerminalSettings**](TerminalSettingsMerchantLevelApi.md#PatchMerchantsMerchantIdTerminalSettings) | **Patch** /merchants/{merchantId}/terminalSettings | Update terminal settings



## GetMerchantsMerchantIdTerminalLogos

> Logo GetMerchantsMerchantIdTerminalLogos(ctx, merchantId).Model(model).Execute()

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
    merchantId := "merchantId_example" // string | The unique identifier of the merchant account.
    model := "model_example" // string | The terminal model. Possible values: E355, VX675WIFIBT, VX680, VX690, VX700, VX820, M400, MX925, P400Plus, UX300, UX410, V200cPlus, V240mPlus, V400cPlus, V400m, e280, e285, e285p, S1E, S1EL, S1F2, S1L, S1U, S7T. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalSettingsMerchantLevelApi.GetMerchantsMerchantIdTerminalLogos(context.Background(), merchantId).Model(model).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalSettingsMerchantLevelApi.GetMerchantsMerchantIdTerminalLogos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsMerchantIdTerminalLogos`: Logo
    fmt.Fprintf(os.Stdout, "Response from `TerminalSettingsMerchantLevelApi.GetMerchantsMerchantIdTerminalLogos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsMerchantIdTerminalLogosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | **string** | The terminal model. Possible values: E355, VX675WIFIBT, VX680, VX690, VX700, VX820, M400, MX925, P400Plus, UX300, UX410, V200cPlus, V240mPlus, V400cPlus, V400m, e280, e285, e285p, S1E, S1EL, S1F2, S1L, S1U, S7T. | 

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


## GetMerchantsMerchantIdTerminalSettings

> TerminalSettings GetMerchantsMerchantIdTerminalSettings(ctx, merchantId).Execute()

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
    merchantId := "merchantId_example" // string | The unique identifier of the merchant account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalSettingsMerchantLevelApi.GetMerchantsMerchantIdTerminalSettings(context.Background(), merchantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalSettingsMerchantLevelApi.GetMerchantsMerchantIdTerminalSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsMerchantIdTerminalSettings`: TerminalSettings
    fmt.Fprintf(os.Stdout, "Response from `TerminalSettingsMerchantLevelApi.GetMerchantsMerchantIdTerminalSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsMerchantIdTerminalSettingsRequest struct via the builder pattern


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


## PatchMerchantsMerchantIdTerminalLogos

> Logo PatchMerchantsMerchantIdTerminalLogos(ctx, merchantId).Model(model).Logo(logo).Execute()

Update the terminal logo



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
    model := "model_example" // string | The terminal model. Allowed values: E355, VX675WIFIBT, VX680, VX690, VX700, VX820, M400, MX925, P400Plus, UX300, UX410, V200cPlus, V240mPlus, V400cPlus, V400m, e280, e285, e285p, S1E, S1EL, S1F2, S1L, S1U, S7T. (optional)
    logo := *openapiclient.NewLogo() // Logo |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalSettingsMerchantLevelApi.PatchMerchantsMerchantIdTerminalLogos(context.Background(), merchantId).Model(model).Logo(logo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalSettingsMerchantLevelApi.PatchMerchantsMerchantIdTerminalLogos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMerchantsMerchantIdTerminalLogos`: Logo
    fmt.Fprintf(os.Stdout, "Response from `TerminalSettingsMerchantLevelApi.PatchMerchantsMerchantIdTerminalLogos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMerchantsMerchantIdTerminalLogosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | **string** | The terminal model. Allowed values: E355, VX675WIFIBT, VX680, VX690, VX700, VX820, M400, MX925, P400Plus, UX300, UX410, V200cPlus, V240mPlus, V400cPlus, V400m, e280, e285, e285p, S1E, S1EL, S1F2, S1L, S1U, S7T. | 
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


## PatchMerchantsMerchantIdTerminalSettings

> TerminalSettings PatchMerchantsMerchantIdTerminalSettings(ctx, merchantId).TerminalSettings(terminalSettings).Execute()

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
    merchantId := "merchantId_example" // string | The unique identifier of the merchant account.
    terminalSettings := *openapiclient.NewTerminalSettings() // TerminalSettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalSettingsMerchantLevelApi.PatchMerchantsMerchantIdTerminalSettings(context.Background(), merchantId).TerminalSettings(terminalSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalSettingsMerchantLevelApi.PatchMerchantsMerchantIdTerminalSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMerchantsMerchantIdTerminalSettings`: TerminalSettings
    fmt.Fprintf(os.Stdout, "Response from `TerminalSettingsMerchantLevelApi.PatchMerchantsMerchantIdTerminalSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMerchantsMerchantIdTerminalSettingsRequest struct via the builder pattern


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

