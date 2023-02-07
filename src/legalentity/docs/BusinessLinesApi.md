# \BusinessLinesApi

All URIs are relative to *https://kyc-test.adyen.com/lem/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBusinessLine**](BusinessLinesApi.md#DeleteBusinessLine) | **Delete** /businessLines/{id} | Delete a business line
[**GetBusinessLine**](BusinessLinesApi.md#GetBusinessLine) | **Get** /businessLines/{id} | Get a business line
[**UpdateBusinessLine**](BusinessLinesApi.md#UpdateBusinessLine) | **Patch** /businessLines/{id} | Update a business line
[**CreateBusinessLine**](BusinessLinesApi.md#CreateBusinessLine) | **Post** /businessLines | Create a business line



## DeleteBusinessLine

> DeleteBusinessLine(ctx, id).Execute()

Delete a business line



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
    id := "id_example" // string | The unique identifier of the business line.

	username := "YOUR_USERNAME"
    password := "YOUR_PASSWORD"
    env := openapiclient.TestEnv

    configuration, err := openapiclient.NewConfiguration(username, password, env)
    apiClient := openapiclient.NewAPIClient(configuration)

    r, err := apiClient.BusinessLinesApi.DeleteBusinessLine(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusinessLinesApi.DeleteBusinessLine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the business line. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBusinessLinesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBusinessLine

> BusinessLine GetBusinessLine(ctx, id).Execute()

Get a business line



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
    id := "id_example" // string | The unique identifier of the business line.

	username := "YOUR_USERNAME"
    password := "YOUR_PASSWORD"
    env := openapiclient.TestEnv

    configuration, err := openapiclient.NewConfiguration(username, password, env)
    apiClient := openapiclient.NewAPIClient(configuration)

    resp, r, err := apiClient.BusinessLinesApi.GetBusinessLine(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusinessLinesApi.GetBusinessLine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBusinessLine`: BusinessLine
    fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the business line. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBusinessLinesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BusinessLine**](BusinessLine.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBusinessLine

> BusinessLine UpdateBusinessLine(ctx, id).BusinessLineInfoUpdate(businessLineInfoUpdate).Execute()

Update a business line



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
    id := "id_example" // string | The unique identifier of the business line.
    businessLineInfoUpdate := *openapiclient.NewBusinessLineInfoUpdate() // BusinessLineInfoUpdate |  (optional)

	username := "YOUR_USERNAME"
    password := "YOUR_PASSWORD"
    env := openapiclient.TestEnv

    configuration, err := openapiclient.NewConfiguration(username, password, env)
    apiClient := openapiclient.NewAPIClient(configuration)

    resp, r, err := apiClient.BusinessLinesApi.UpdateBusinessLine(context.Background(), id).BusinessLineInfoUpdate(businessLineInfoUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusinessLinesApi.UpdateBusinessLine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBusinessLine`: BusinessLine
    fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the business line. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchBusinessLinesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **businessLineInfoUpdate** | [**BusinessLineInfoUpdate**](BusinessLineInfoUpdate.md) |  | 

### Return type

[**BusinessLine**](BusinessLine.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBusinessLine

> BusinessLine CreateBusinessLine(ctx).BusinessLineInfo(businessLineInfo).Execute()

Create a business line



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
    businessLineInfo := *openapiclient.NewBusinessLineInfo("Capability_example", "IndustryCode_example", "LegalEntityId_example") // BusinessLineInfo |  (optional)

	username := "YOUR_USERNAME"
    password := "YOUR_PASSWORD"
    env := openapiclient.TestEnv

    configuration, err := openapiclient.NewConfiguration(username, password, env)
    apiClient := openapiclient.NewAPIClient(configuration)

    resp, r, err := apiClient.BusinessLinesApi.CreateBusinessLine(context.Background()).BusinessLineInfo(businessLineInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusinessLinesApi.CreateBusinessLine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBusinessLine`: BusinessLine
    fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBusinessLinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **businessLineInfo** | [**BusinessLineInfo**](BusinessLineInfo.md) |  | 

### Return type

[**BusinessLine**](BusinessLine.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

