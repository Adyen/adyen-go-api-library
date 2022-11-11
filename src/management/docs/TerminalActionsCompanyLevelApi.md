# \TerminalActionsCompanyLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompaniesCompanyIdAndroidApps**](TerminalActionsCompanyLevelApi.md#GetCompaniesCompanyIdAndroidApps) | **Get** /companies/{companyId}/androidApps | Get a list of Android apps
[**GetCompaniesCompanyIdAndroidCertificates**](TerminalActionsCompanyLevelApi.md#GetCompaniesCompanyIdAndroidCertificates) | **Get** /companies/{companyId}/androidCertificates | Get a list of Android certificates
[**GetCompaniesCompanyIdTerminalActions**](TerminalActionsCompanyLevelApi.md#GetCompaniesCompanyIdTerminalActions) | **Get** /companies/{companyId}/terminalActions | Get a list of terminal actions
[**GetCompaniesCompanyIdTerminalActionsActionId**](TerminalActionsCompanyLevelApi.md#GetCompaniesCompanyIdTerminalActionsActionId) | **Get** /companies/{companyId}/terminalActions/{actionId} | Get terminal action



## GetCompaniesCompanyIdAndroidApps

> AndroidAppsResponse GetCompaniesCompanyIdAndroidApps(ctx, companyId).PageNumber(pageNumber).PageSize(pageSize).Execute()

Get a list of Android apps



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
    pageNumber := int32(56) // int32 | The number of the page to fetch. (optional)
    pageSize := int32(56) // int32 | The number of items to have on a page, maximum 100. The default is 20 items on a page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalActionsCompanyLevelApi.GetCompaniesCompanyIdAndroidApps(context.Background(), companyId).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalActionsCompanyLevelApi.GetCompaniesCompanyIdAndroidApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompaniesCompanyIdAndroidApps`: AndroidAppsResponse
    fmt.Fprintf(os.Stdout, "Response from `TerminalActionsCompanyLevelApi.GetCompaniesCompanyIdAndroidApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesCompanyIdAndroidAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | The number of the page to fetch. | 
 **pageSize** | **int32** | The number of items to have on a page, maximum 100. The default is 20 items on a page. | 

### Return type

[**AndroidAppsResponse**](AndroidAppsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompaniesCompanyIdAndroidCertificates

> AndroidCertificatesResponse GetCompaniesCompanyIdAndroidCertificates(ctx, companyId).PageNumber(pageNumber).PageSize(pageSize).Execute()

Get a list of Android certificates



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
    pageNumber := int32(56) // int32 | The number of the page to fetch. (optional)
    pageSize := int32(56) // int32 | The number of items to have on a page, maximum 100. The default is 20 items on a page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalActionsCompanyLevelApi.GetCompaniesCompanyIdAndroidCertificates(context.Background(), companyId).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalActionsCompanyLevelApi.GetCompaniesCompanyIdAndroidCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompaniesCompanyIdAndroidCertificates`: AndroidCertificatesResponse
    fmt.Fprintf(os.Stdout, "Response from `TerminalActionsCompanyLevelApi.GetCompaniesCompanyIdAndroidCertificates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesCompanyIdAndroidCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | The number of the page to fetch. | 
 **pageSize** | **int32** | The number of items to have on a page, maximum 100. The default is 20 items on a page. | 

### Return type

[**AndroidCertificatesResponse**](AndroidCertificatesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompaniesCompanyIdTerminalActions

> ListExternalTerminalActionsResponse GetCompaniesCompanyIdTerminalActions(ctx, companyId).PageNumber(pageNumber).PageSize(pageSize).Status(status).Type_(type_).Execute()

Get a list of terminal actions



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
    pageNumber := int32(56) // int32 | The number of the page to fetch. (optional)
    pageSize := int32(56) // int32 | The number of items to have on a page, maximum 100. The default is 20 items on a page. (optional)
    status := "status_example" // string | Returns terminal actions with the specified status.  Allowed values: **pending**, **successful**, **failed**, **cancelled**, **tryLater**. (optional)
    type_ := "type__example" // string | Returns terminal actions of the specified type.  Allowed values: **InstallAndroidApp**, **UninstallAndroidApp**, **InstallAndroidCertificate**, **UninstallAndroidCertificate**. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalActionsCompanyLevelApi.GetCompaniesCompanyIdTerminalActions(context.Background(), companyId).PageNumber(pageNumber).PageSize(pageSize).Status(status).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalActionsCompanyLevelApi.GetCompaniesCompanyIdTerminalActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompaniesCompanyIdTerminalActions`: ListExternalTerminalActionsResponse
    fmt.Fprintf(os.Stdout, "Response from `TerminalActionsCompanyLevelApi.GetCompaniesCompanyIdTerminalActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesCompanyIdTerminalActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | The number of the page to fetch. | 
 **pageSize** | **int32** | The number of items to have on a page, maximum 100. The default is 20 items on a page. | 
 **status** | **string** | Returns terminal actions with the specified status.  Allowed values: **pending**, **successful**, **failed**, **cancelled**, **tryLater**. | 
 **type_** | **string** | Returns terminal actions of the specified type.  Allowed values: **InstallAndroidApp**, **UninstallAndroidApp**, **InstallAndroidCertificate**, **UninstallAndroidCertificate**. | 

### Return type

[**ListExternalTerminalActionsResponse**](ListExternalTerminalActionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompaniesCompanyIdTerminalActionsActionId

> ExternalTerminalAction GetCompaniesCompanyIdTerminalActionsActionId(ctx, companyId, actionId).Execute()

Get terminal action



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
    actionId := "actionId_example" // string | The unique identifier of the terminal action.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalActionsCompanyLevelApi.GetCompaniesCompanyIdTerminalActionsActionId(context.Background(), companyId, actionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalActionsCompanyLevelApi.GetCompaniesCompanyIdTerminalActionsActionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompaniesCompanyIdTerminalActionsActionId`: ExternalTerminalAction
    fmt.Fprintf(os.Stdout, "Response from `TerminalActionsCompanyLevelApi.GetCompaniesCompanyIdTerminalActionsActionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 
**actionId** | **string** | The unique identifier of the terminal action. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesCompanyIdTerminalActionsActionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ExternalTerminalAction**](ExternalTerminalAction.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

