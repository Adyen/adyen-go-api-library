# \TermsOfServiceApi

All URIs are relative to *https://kyc-test.adyen.com/lem/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTermsOfServiceInformationForLegalEntity**](TermsOfServiceApi.md#GetTermsOfServiceInformationForLegalEntity) | **Get** /legalEntities/{id}/termsOfServiceAcceptanceInfos | Get Terms of Service information for a legal entity
[**GetTermsOfServiceStatus**](TermsOfServiceApi.md#GetTermsOfServiceStatus) | **Get** /legalEntities/{id}/termsOfServiceStatus | Get Terms of Service status
[**AcceptTermsOfService**](TermsOfServiceApi.md#AcceptTermsOfService) | **Patch** /legalEntities/{id}/termsOfService/{termsofservicedocumentid} | Accept Terms of Service
[**GetTermsOfServiceDocument**](TermsOfServiceApi.md#GetTermsOfServiceDocument) | **Post** /legalEntities/{id}/termsOfService | Get Terms of Service document



## GetTermsOfServiceInformationForLegalEntity

> GetTermsOfServiceAcceptanceInfosResponse GetTermsOfServiceInformationForLegalEntity(ctx, id).Execute()

Get Terms of Service information for a legal entity



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

    resp, r, err := apiClient.TermsOfServiceApi.GetTermsOfServiceInformationForLegalEntity(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TermsOfServiceApi.GetTermsOfServiceInformationForLegalEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTermsOfServiceInformationForLegalEntity`: GetTermsOfServiceAcceptanceInfosResponse
    fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the legal entity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLegalEntitiesIdTermsOfServiceAcceptanceInfosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTermsOfServiceAcceptanceInfosResponse**](GetTermsOfServiceAcceptanceInfosResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTermsOfServiceStatus

> CalculateTermsOfServiceStatusResponse GetTermsOfServiceStatus(ctx, id).Execute()

Get Terms of Service status



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

    resp, r, err := apiClient.TermsOfServiceApi.GetTermsOfServiceStatus(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TermsOfServiceApi.GetTermsOfServiceStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTermsOfServiceStatus`: CalculateTermsOfServiceStatusResponse
    fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the legal entity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLegalEntitiesIdTermsOfServiceStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CalculateTermsOfServiceStatusResponse**](CalculateTermsOfServiceStatusResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcceptTermsOfService

> AcceptTermsOfServiceResponse AcceptTermsOfService(ctx, id, termsofservicedocumentid).AcceptTermsOfServiceRequest(acceptTermsOfServiceRequest).Execute()

Accept Terms of Service



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
    termsofservicedocumentid := "termsofservicedocumentid_example" // string | The unique identifier of the Terms of Service document.
    acceptTermsOfServiceRequest := *openapiclient.NewAcceptTermsOfServiceRequest() // AcceptTermsOfServiceRequest |  (optional)

    username := "YOUR_USERNAME"
    password := "YOUR_PASSWORD"
    env := openapiclient.TestEnv

    configuration, err := openapiclient.NewConfiguration(username, password, env)
    apiClient := openapiclient.NewAPIClient(configuration)

    resp, r, err := apiClient.TermsOfServiceApi.AcceptTermsOfService(context.Background(), id, termsofservicedocumentid).AcceptTermsOfServiceRequest(acceptTermsOfServiceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TermsOfServiceApi.AcceptTermsOfService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcceptTermsOfService`: AcceptTermsOfServiceResponse
    fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the legal entity. | 
**termsofservicedocumentid** | **string** | The unique identifier of the Terms of Service document. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchLegalEntitiesIdTermsOfServiceTermsofservicedocumentidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptTermsOfServiceRequest** | [**AcceptTermsOfServiceRequest**](AcceptTermsOfServiceRequest.md) |  | 

### Return type

[**AcceptTermsOfServiceResponse**](AcceptTermsOfServiceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTermsOfServiceDocument

> GetTermsOfServiceDocumentResponse GetTermsOfServiceDocument(ctx, id).GetTermsOfServiceDocumentRequest(getTermsOfServiceDocumentRequest).Execute()

Get Terms of Service document



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
    getTermsOfServiceDocumentRequest := *openapiclient.NewGetTermsOfServiceDocumentRequest() // GetTermsOfServiceDocumentRequest |  (optional)

    username := "YOUR_USERNAME"
    password := "YOUR_PASSWORD"
    env := openapiclient.TestEnv

    configuration, err := openapiclient.NewConfiguration(username, password, env)
    apiClient := openapiclient.NewAPIClient(configuration)

    resp, r, err := apiClient.TermsOfServiceApi.GetTermsOfServiceDocument(context.Background(), id).GetTermsOfServiceDocumentRequest(getTermsOfServiceDocumentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TermsOfServiceApi.GetTermsOfServiceDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTermsOfServiceDocument`: GetTermsOfServiceDocumentResponse
    fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the legal entity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLegalEntitiesIdTermsOfServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getTermsOfServiceDocumentRequest** | [**GetTermsOfServiceDocumentRequest**](GetTermsOfServiceDocumentRequest.md) |  | 

### Return type

[**GetTermsOfServiceDocumentResponse**](GetTermsOfServiceDocumentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

