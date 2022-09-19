# \TerminalSettingsCompanyLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompaniesCompanyIdTerminalLogos**](TerminalSettingsCompanyLevelApi.md#GetCompaniesCompanyIdTerminalLogos) | **Get** /companies/{companyId}/terminalLogos | Get the terminal logo
[**GetCompaniesCompanyIdTerminalSettings**](TerminalSettingsCompanyLevelApi.md#GetCompaniesCompanyIdTerminalSettings) | **Get** /companies/{companyId}/terminalSettings | Get terminal settings
[**PatchCompaniesCompanyIdTerminalLogos**](TerminalSettingsCompanyLevelApi.md#PatchCompaniesCompanyIdTerminalLogos) | **Patch** /companies/{companyId}/terminalLogos | Update the terminal logo
[**PatchCompaniesCompanyIdTerminalSettings**](TerminalSettingsCompanyLevelApi.md#PatchCompaniesCompanyIdTerminalSettings) | **Patch** /companies/{companyId}/terminalSettings | Update terminal settings



## GetCompaniesCompanyIdTerminalLogos

> Logo GetCompaniesCompanyIdTerminalLogos(ctx, companyId).Model(model).Execute()

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
    companyId := "companyId_example" // string | The unique identifier of the company account.
    model := "model_example" // string | The terminal model. Possible values: E355, VX675WIFIBT, VX680, VX690, VX700, VX820, M400, MX925, P400Plus, UX300, UX410, V200cPlus, V240mPlus, V400cPlus, V400m, e280, e285, e285p, S1E, S1EL, S1F2, S1L, S1U, S7T. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalSettingsCompanyLevelApi.GetCompaniesCompanyIdTerminalLogos(context.Background(), companyId).Model(model).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalSettingsCompanyLevelApi.GetCompaniesCompanyIdTerminalLogos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompaniesCompanyIdTerminalLogos`: Logo
    fmt.Fprintf(os.Stdout, "Response from `TerminalSettingsCompanyLevelApi.GetCompaniesCompanyIdTerminalLogos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesCompanyIdTerminalLogosRequest struct via the builder pattern


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


## GetCompaniesCompanyIdTerminalSettings

> TerminalSettings GetCompaniesCompanyIdTerminalSettings(ctx, companyId).Execute()

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
    companyId := "companyId_example" // string | The unique identifier of the company account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalSettingsCompanyLevelApi.GetCompaniesCompanyIdTerminalSettings(context.Background(), companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalSettingsCompanyLevelApi.GetCompaniesCompanyIdTerminalSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompaniesCompanyIdTerminalSettings`: TerminalSettings
    fmt.Fprintf(os.Stdout, "Response from `TerminalSettingsCompanyLevelApi.GetCompaniesCompanyIdTerminalSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesCompanyIdTerminalSettingsRequest struct via the builder pattern


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


## PatchCompaniesCompanyIdTerminalLogos

> Logo PatchCompaniesCompanyIdTerminalLogos(ctx, companyId).Model(model).Logo(logo).Execute()

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
    companyId := "companyId_example" // string | The unique identifier of the company account.
    model := "model_example" // string | The terminal model. Possible values: E355, VX675WIFIBT, VX680, VX690, VX700, VX820, M400, MX925, P400Plus, UX300, UX410, V200cPlus, V240mPlus, V400cPlus, V400m, e280, e285, e285p, S1E, S1EL, S1F2, S1L, S1U, S7T. (optional)
    logo := *openapiclient.NewLogo() // Logo |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalSettingsCompanyLevelApi.PatchCompaniesCompanyIdTerminalLogos(context.Background(), companyId).Model(model).Logo(logo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalSettingsCompanyLevelApi.PatchCompaniesCompanyIdTerminalLogos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCompaniesCompanyIdTerminalLogos`: Logo
    fmt.Fprintf(os.Stdout, "Response from `TerminalSettingsCompanyLevelApi.PatchCompaniesCompanyIdTerminalLogos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompaniesCompanyIdTerminalLogosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | **string** | The terminal model. Possible values: E355, VX675WIFIBT, VX680, VX690, VX700, VX820, M400, MX925, P400Plus, UX300, UX410, V200cPlus, V240mPlus, V400cPlus, V400m, e280, e285, e285p, S1E, S1EL, S1F2, S1L, S1U, S7T. | 
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


## PatchCompaniesCompanyIdTerminalSettings

> TerminalSettings PatchCompaniesCompanyIdTerminalSettings(ctx, companyId).TerminalSettings(terminalSettings).Execute()

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
    companyId := "companyId_example" // string | The unique identifier of the company account.
    terminalSettings := *openapiclient.NewTerminalSettings() // TerminalSettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalSettingsCompanyLevelApi.PatchCompaniesCompanyIdTerminalSettings(context.Background(), companyId).TerminalSettings(terminalSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalSettingsCompanyLevelApi.PatchCompaniesCompanyIdTerminalSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCompaniesCompanyIdTerminalSettings`: TerminalSettings
    fmt.Fprintf(os.Stdout, "Response from `TerminalSettingsCompanyLevelApi.PatchCompaniesCompanyIdTerminalSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompaniesCompanyIdTerminalSettingsRequest struct via the builder pattern


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

