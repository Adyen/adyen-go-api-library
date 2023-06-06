# \LegalEntitiesApi

All URIs are relative to *https://kyc-test.adyen.com/lem/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLegalEntity**](LegalEntitiesApi.md#GetLegalEntity) | **Get** /legalEntities/{id} | Get a legal entity
[**GetAllBusinessLinesUnderLegalEntity**](LegalEntitiesApi.md#GetAllBusinessLinesUnderLegalEntity) | **Get** /legalEntities/{id}/businessLines | Get all business lines under a legal entity
[**UpdateLegalEntity**](LegalEntitiesApi.md#UpdateLegalEntity) | **Patch** /legalEntities/{id} | Update a legal entity
[**CreateLegalEntity**](LegalEntitiesApi.md#CreateLegalEntity) | **Post** /legalEntities | Create a legal entity



## GetLegalEntity

> LegalEntity GetLegalEntity(ctx, id).Execute()

Get a legal entity



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/adyen/adyen-go-api-library/v7/src/legalentity"
)

func main() {
    id := "id_example" // string | The unique identifier of the legal entity.

    username := "YOUR_USERNAME"
    password := "YOUR_PASSWORD"
    env := openapiclient.TestEnv

    configuration, err := openapiclient.NewConfiguration(username, password, env)
    apiClient := openapiclient.NewAPIClient(configuration)

    resp, r, err := apiClient.LegalEntitiesApi.GetLegalEntity(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegalEntitiesApi.GetLegalEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLegalEntity`: LegalEntity
    fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the legal entity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLegalEntitiesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LegalEntity**](LegalEntity.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllBusinessLinesUnderLegalEntity

> BusinessLines GetAllBusinessLinesUnderLegalEntity(ctx, id).Execute()

Get all business lines under a legal entity



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/adyen/adyen-go-api-library/v7/src/legalentity"
)

func main() {
    id := "id_example" // string | The unique identifier of the legal entity.

    username := "YOUR_USERNAME"
    password := "YOUR_PASSWORD"
    env := openapiclient.TestEnv

    configuration, err := openapiclient.NewConfiguration(username, password, env)
    apiClient := openapiclient.NewAPIClient(configuration)

    resp, r, err := apiClient.LegalEntitiesApi.GetAllBusinessLinesUnderLegalEntity(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegalEntitiesApi.GetAllBusinessLinesUnderLegalEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllBusinessLinesUnderLegalEntity`: BusinessLines
    fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the legal entity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLegalEntitiesIdBusinessLinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BusinessLines**](BusinessLines.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLegalEntity

> LegalEntity UpdateLegalEntity(ctx, id).LegalEntityInfo(legalEntityInfo).Execute()

Update a legal entity



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/adyen/adyen-go-api-library/v7/src/legalentity"
)

func main() {
    id := "id_example" // string | The unique identifier of the legal entity.
    legalEntityInfo := *openapiclient.NewLegalEntityInfo() // LegalEntityInfo |  (optional)

    username := "YOUR_USERNAME"
    password := "YOUR_PASSWORD"
    env := openapiclient.TestEnv

    configuration, err := openapiclient.NewConfiguration(username, password, env)
    apiClient := openapiclient.NewAPIClient(configuration)

    resp, r, err := apiClient.LegalEntitiesApi.UpdateLegalEntity(context.Background(), id).LegalEntityInfo(legalEntityInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegalEntitiesApi.UpdateLegalEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLegalEntity`: LegalEntity
    fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the legal entity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchLegalEntitiesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **legalEntityInfo** | [**LegalEntityInfo**](LegalEntityInfo.md) |  | 

### Return type

[**LegalEntity**](LegalEntity.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLegalEntity

> LegalEntity CreateLegalEntity(ctx).LegalEntityInfoRequiredType(legalEntityInfoRequiredType).Execute()

Create a legal entity



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/adyen/adyen-go-api-library/v7/src/legalentity"
)

func main() {
    legalEntityInfoRequiredType := *openapiclient.NewLegalEntityInfoRequiredType("Type_example") // LegalEntityInfoRequiredType |  (optional)

    username := "YOUR_USERNAME"
    password := "YOUR_PASSWORD"
    env := openapiclient.TestEnv

    configuration, err := openapiclient.NewConfiguration(username, password, env)
    apiClient := openapiclient.NewAPIClient(configuration)

    resp, r, err := apiClient.LegalEntitiesApi.CreateLegalEntity(context.Background()).LegalEntityInfoRequiredType(legalEntityInfoRequiredType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegalEntitiesApi.CreateLegalEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLegalEntity`: LegalEntity
    fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLegalEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **legalEntityInfoRequiredType** | [**LegalEntityInfoRequiredType**](LegalEntityInfoRequiredType.md) |  | 

### Return type

[**LegalEntity**](LegalEntity.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

