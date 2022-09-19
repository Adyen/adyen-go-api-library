# \TerminalSettingsStoreLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMerchantsMerchantIdStoresReferenceTerminalLogos**](TerminalSettingsStoreLevelApi.md#GetMerchantsMerchantIdStoresReferenceTerminalLogos) | **Get** /merchants/{merchantId}/stores/{reference}/terminalLogos | Get the terminal logo
[**GetMerchantsMerchantIdStoresReferenceTerminalSettings**](TerminalSettingsStoreLevelApi.md#GetMerchantsMerchantIdStoresReferenceTerminalSettings) | **Get** /merchants/{merchantId}/stores/{reference}/terminalSettings | Get terminal settings
[**GetStoresStoreIdTerminalLogos**](TerminalSettingsStoreLevelApi.md#GetStoresStoreIdTerminalLogos) | **Get** /stores/{storeId}/terminalLogos | Get the terminal logo
[**GetStoresStoreIdTerminalSettings**](TerminalSettingsStoreLevelApi.md#GetStoresStoreIdTerminalSettings) | **Get** /stores/{storeId}/terminalSettings | Get terminal settings
[**PatchMerchantsMerchantIdStoresReferenceTerminalLogos**](TerminalSettingsStoreLevelApi.md#PatchMerchantsMerchantIdStoresReferenceTerminalLogos) | **Patch** /merchants/{merchantId}/stores/{reference}/terminalLogos | Update the terminal logo
[**PatchMerchantsMerchantIdStoresReferenceTerminalSettings**](TerminalSettingsStoreLevelApi.md#PatchMerchantsMerchantIdStoresReferenceTerminalSettings) | **Patch** /merchants/{merchantId}/stores/{reference}/terminalSettings | Update terminal settings
[**PatchStoresStoreIdTerminalLogos**](TerminalSettingsStoreLevelApi.md#PatchStoresStoreIdTerminalLogos) | **Patch** /stores/{storeId}/terminalLogos | Update the terminal logo
[**PatchStoresStoreIdTerminalSettings**](TerminalSettingsStoreLevelApi.md#PatchStoresStoreIdTerminalSettings) | **Patch** /stores/{storeId}/terminalSettings | Update terminal settings



## GetMerchantsMerchantIdStoresReferenceTerminalLogos

> Logo GetMerchantsMerchantIdStoresReferenceTerminalLogos(ctx, merchantId, reference).Model(model).Execute()

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
    reference := "reference_example" // string | The reference that identifies the store.
    model := "model_example" // string | The terminal model. Possible values: E355, VX675WIFIBT, VX680, VX690, VX700, VX820, M400, MX925, P400Plus, UX300, UX410, V200cPlus, V240mPlus, V400cPlus, V400m, e280, e285, e285p, S1E, S1EL, S1F2, S1L, S1U, S7T. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalSettingsStoreLevelApi.GetMerchantsMerchantIdStoresReferenceTerminalLogos(context.Background(), merchantId, reference).Model(model).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalSettingsStoreLevelApi.GetMerchantsMerchantIdStoresReferenceTerminalLogos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsMerchantIdStoresReferenceTerminalLogos`: Logo
    fmt.Fprintf(os.Stdout, "Response from `TerminalSettingsStoreLevelApi.GetMerchantsMerchantIdStoresReferenceTerminalLogos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**reference** | **string** | The reference that identifies the store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsMerchantIdStoresReferenceTerminalLogosRequest struct via the builder pattern


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


## GetMerchantsMerchantIdStoresReferenceTerminalSettings

> TerminalSettings GetMerchantsMerchantIdStoresReferenceTerminalSettings(ctx, merchantId, reference).Execute()

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
    reference := "reference_example" // string | The reference that identifies the store.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalSettingsStoreLevelApi.GetMerchantsMerchantIdStoresReferenceTerminalSettings(context.Background(), merchantId, reference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalSettingsStoreLevelApi.GetMerchantsMerchantIdStoresReferenceTerminalSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsMerchantIdStoresReferenceTerminalSettings`: TerminalSettings
    fmt.Fprintf(os.Stdout, "Response from `TerminalSettingsStoreLevelApi.GetMerchantsMerchantIdStoresReferenceTerminalSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**reference** | **string** | The reference that identifies the store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsMerchantIdStoresReferenceTerminalSettingsRequest struct via the builder pattern


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


## GetStoresStoreIdTerminalLogos

> Logo GetStoresStoreIdTerminalLogos(ctx, storeId).Model(model).Execute()

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
    storeId := "storeId_example" // string | The unique identifier of the store.
    model := "model_example" // string | The terminal model. Possible values: E355, VX675WIFIBT, VX680, VX690, VX700, VX820, M400, MX925, P400Plus, UX300, UX410, V200cPlus, V240mPlus, V400cPlus, V400m, e280, e285, e285p, S1E, S1EL, S1F2, S1L, S1U, S7T. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalSettingsStoreLevelApi.GetStoresStoreIdTerminalLogos(context.Background(), storeId).Model(model).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalSettingsStoreLevelApi.GetStoresStoreIdTerminalLogos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStoresStoreIdTerminalLogos`: Logo
    fmt.Fprintf(os.Stdout, "Response from `TerminalSettingsStoreLevelApi.GetStoresStoreIdTerminalLogos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The unique identifier of the store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoresStoreIdTerminalLogosRequest struct via the builder pattern


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


## GetStoresStoreIdTerminalSettings

> TerminalSettings GetStoresStoreIdTerminalSettings(ctx, storeId).Execute()

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
    storeId := "storeId_example" // string | The unique identifier of the store.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalSettingsStoreLevelApi.GetStoresStoreIdTerminalSettings(context.Background(), storeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalSettingsStoreLevelApi.GetStoresStoreIdTerminalSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStoresStoreIdTerminalSettings`: TerminalSettings
    fmt.Fprintf(os.Stdout, "Response from `TerminalSettingsStoreLevelApi.GetStoresStoreIdTerminalSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The unique identifier of the store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoresStoreIdTerminalSettingsRequest struct via the builder pattern


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


## PatchMerchantsMerchantIdStoresReferenceTerminalLogos

> Logo PatchMerchantsMerchantIdStoresReferenceTerminalLogos(ctx, merchantId, reference).Model(model).Logo(logo).Execute()

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
    reference := "reference_example" // string | The reference that identifies the store.
    model := "model_example" // string | The terminal model. Possible values: E355, VX675WIFIBT, VX680, VX690, VX700, VX820, M400, MX925, P400Plus, UX300, UX410, V200cPlus, V240mPlus, V400cPlus, V400m, e280, e285, e285p, S1E, S1EL, S1F2, S1L, S1U, S7T (optional)
    logo := *openapiclient.NewLogo() // Logo |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalSettingsStoreLevelApi.PatchMerchantsMerchantIdStoresReferenceTerminalLogos(context.Background(), merchantId, reference).Model(model).Logo(logo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalSettingsStoreLevelApi.PatchMerchantsMerchantIdStoresReferenceTerminalLogos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMerchantsMerchantIdStoresReferenceTerminalLogos`: Logo
    fmt.Fprintf(os.Stdout, "Response from `TerminalSettingsStoreLevelApi.PatchMerchantsMerchantIdStoresReferenceTerminalLogos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**reference** | **string** | The reference that identifies the store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMerchantsMerchantIdStoresReferenceTerminalLogosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | **string** | The terminal model. Possible values: E355, VX675WIFIBT, VX680, VX690, VX700, VX820, M400, MX925, P400Plus, UX300, UX410, V200cPlus, V240mPlus, V400cPlus, V400m, e280, e285, e285p, S1E, S1EL, S1F2, S1L, S1U, S7T | 
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


## PatchMerchantsMerchantIdStoresReferenceTerminalSettings

> TerminalSettings PatchMerchantsMerchantIdStoresReferenceTerminalSettings(ctx, merchantId, reference).TerminalSettings(terminalSettings).Execute()

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
    reference := "reference_example" // string | The reference that identifies the store.
    terminalSettings := *openapiclient.NewTerminalSettings() // TerminalSettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalSettingsStoreLevelApi.PatchMerchantsMerchantIdStoresReferenceTerminalSettings(context.Background(), merchantId, reference).TerminalSettings(terminalSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalSettingsStoreLevelApi.PatchMerchantsMerchantIdStoresReferenceTerminalSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMerchantsMerchantIdStoresReferenceTerminalSettings`: TerminalSettings
    fmt.Fprintf(os.Stdout, "Response from `TerminalSettingsStoreLevelApi.PatchMerchantsMerchantIdStoresReferenceTerminalSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**reference** | **string** | The reference that identifies the store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMerchantsMerchantIdStoresReferenceTerminalSettingsRequest struct via the builder pattern


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


## PatchStoresStoreIdTerminalLogos

> Logo PatchStoresStoreIdTerminalLogos(ctx, storeId).Model(model).Logo(logo).Execute()

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
    storeId := "storeId_example" // string | The unique identifier of the store.
    model := "model_example" // string | The terminal model. Possible values: E355, VX675WIFIBT, VX680, VX690, VX700, VX820, M400, MX925, P400Plus, UX300, UX410, V200cPlus, V240mPlus, V400cPlus, V400m, e280, e285, e285p, S1E, S1EL, S1F2, S1L, S1U, S7T. (optional)
    logo := *openapiclient.NewLogo() // Logo |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalSettingsStoreLevelApi.PatchStoresStoreIdTerminalLogos(context.Background(), storeId).Model(model).Logo(logo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalSettingsStoreLevelApi.PatchStoresStoreIdTerminalLogos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchStoresStoreIdTerminalLogos`: Logo
    fmt.Fprintf(os.Stdout, "Response from `TerminalSettingsStoreLevelApi.PatchStoresStoreIdTerminalLogos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The unique identifier of the store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchStoresStoreIdTerminalLogosRequest struct via the builder pattern


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


## PatchStoresStoreIdTerminalSettings

> TerminalSettings PatchStoresStoreIdTerminalSettings(ctx, storeId).TerminalSettings(terminalSettings).Execute()

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
    storeId := "storeId_example" // string | The unique identifier of the store.
    terminalSettings := *openapiclient.NewTerminalSettings() // TerminalSettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalSettingsStoreLevelApi.PatchStoresStoreIdTerminalSettings(context.Background(), storeId).TerminalSettings(terminalSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalSettingsStoreLevelApi.PatchStoresStoreIdTerminalSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchStoresStoreIdTerminalSettings`: TerminalSettings
    fmt.Fprintf(os.Stdout, "Response from `TerminalSettingsStoreLevelApi.PatchStoresStoreIdTerminalSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | The unique identifier of the store. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchStoresStoreIdTerminalSettingsRequest struct via the builder pattern


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

