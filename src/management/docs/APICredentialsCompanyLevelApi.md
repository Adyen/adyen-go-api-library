# \APICredentialsCompanyLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompaniesCompanyIdApiCredentials**](APICredentialsCompanyLevelApi.md#GetCompaniesCompanyIdApiCredentials) | **Get** /companies/{companyId}/apiCredentials | Get a list of API credentials
[**GetCompaniesCompanyIdApiCredentialsApiCredentialId**](APICredentialsCompanyLevelApi.md#GetCompaniesCompanyIdApiCredentialsApiCredentialId) | **Get** /companies/{companyId}/apiCredentials/{apiCredentialId} | Get an API credential
[**PatchCompaniesCompanyIdApiCredentialsApiCredentialId**](APICredentialsCompanyLevelApi.md#PatchCompaniesCompanyIdApiCredentialsApiCredentialId) | **Patch** /companies/{companyId}/apiCredentials/{apiCredentialId} | Update an API credential.
[**PostCompaniesCompanyIdApiCredentials**](APICredentialsCompanyLevelApi.md#PostCompaniesCompanyIdApiCredentials) | **Post** /companies/{companyId}/apiCredentials | Create an API credential.



## GetCompaniesCompanyIdApiCredentials

> ListCompanyApiCredentialsResponse GetCompaniesCompanyIdApiCredentials(ctx, companyId).PageNumber(pageNumber).PageSize(pageSize).Execute()

Get a list of API credentials



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
    pageSize := int32(56) // int32 | The number of items to have on a page, maximum 100. The default is 10 items on a page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APICredentialsCompanyLevelApi.GetCompaniesCompanyIdApiCredentials(context.Background(), companyId).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APICredentialsCompanyLevelApi.GetCompaniesCompanyIdApiCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompaniesCompanyIdApiCredentials`: ListCompanyApiCredentialsResponse
    fmt.Fprintf(os.Stdout, "Response from `APICredentialsCompanyLevelApi.GetCompaniesCompanyIdApiCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesCompanyIdApiCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | The number of the page to fetch. | 
 **pageSize** | **int32** | The number of items to have on a page, maximum 100. The default is 10 items on a page. | 

### Return type

[**ListCompanyApiCredentialsResponse**](ListCompanyApiCredentialsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompaniesCompanyIdApiCredentialsApiCredentialId

> CompanyApiCredential GetCompaniesCompanyIdApiCredentialsApiCredentialId(ctx, companyId, apiCredentialId).Execute()

Get an API credential



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
    apiCredentialId := "apiCredentialId_example" // string | Unique identifier of the API credential.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APICredentialsCompanyLevelApi.GetCompaniesCompanyIdApiCredentialsApiCredentialId(context.Background(), companyId, apiCredentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APICredentialsCompanyLevelApi.GetCompaniesCompanyIdApiCredentialsApiCredentialId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompaniesCompanyIdApiCredentialsApiCredentialId`: CompanyApiCredential
    fmt.Fprintf(os.Stdout, "Response from `APICredentialsCompanyLevelApi.GetCompaniesCompanyIdApiCredentialsApiCredentialId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 
**apiCredentialId** | **string** | Unique identifier of the API credential. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesCompanyIdApiCredentialsApiCredentialIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CompanyApiCredential**](CompanyApiCredential.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCompaniesCompanyIdApiCredentialsApiCredentialId

> CompanyApiCredential PatchCompaniesCompanyIdApiCredentialsApiCredentialId(ctx, companyId, apiCredentialId).UpdateCompanyApiCredentialRequest(updateCompanyApiCredentialRequest).Execute()

Update an API credential.



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
    apiCredentialId := "apiCredentialId_example" // string | Unique identifier of the API credential.
    updateCompanyApiCredentialRequest := *openapiclient.NewUpdateCompanyApiCredentialRequest() // UpdateCompanyApiCredentialRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APICredentialsCompanyLevelApi.PatchCompaniesCompanyIdApiCredentialsApiCredentialId(context.Background(), companyId, apiCredentialId).UpdateCompanyApiCredentialRequest(updateCompanyApiCredentialRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APICredentialsCompanyLevelApi.PatchCompaniesCompanyIdApiCredentialsApiCredentialId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCompaniesCompanyIdApiCredentialsApiCredentialId`: CompanyApiCredential
    fmt.Fprintf(os.Stdout, "Response from `APICredentialsCompanyLevelApi.PatchCompaniesCompanyIdApiCredentialsApiCredentialId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 
**apiCredentialId** | **string** | Unique identifier of the API credential. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompaniesCompanyIdApiCredentialsApiCredentialIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateCompanyApiCredentialRequest** | [**UpdateCompanyApiCredentialRequest**](UpdateCompanyApiCredentialRequest.md) |  | 

### Return type

[**CompanyApiCredential**](CompanyApiCredential.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompaniesCompanyIdApiCredentials

> CreateCompanyApiCredentialResponse PostCompaniesCompanyIdApiCredentials(ctx, companyId).CreateCompanyApiCredentialRequest(createCompanyApiCredentialRequest).Execute()

Create an API credential.



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
    createCompanyApiCredentialRequest := *openapiclient.NewCreateCompanyApiCredentialRequest() // CreateCompanyApiCredentialRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APICredentialsCompanyLevelApi.PostCompaniesCompanyIdApiCredentials(context.Background(), companyId).CreateCompanyApiCredentialRequest(createCompanyApiCredentialRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APICredentialsCompanyLevelApi.PostCompaniesCompanyIdApiCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCompaniesCompanyIdApiCredentials`: CreateCompanyApiCredentialResponse
    fmt.Fprintf(os.Stdout, "Response from `APICredentialsCompanyLevelApi.PostCompaniesCompanyIdApiCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompaniesCompanyIdApiCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createCompanyApiCredentialRequest** | [**CreateCompanyApiCredentialRequest**](CreateCompanyApiCredentialRequest.md) |  | 

### Return type

[**CreateCompanyApiCredentialResponse**](CreateCompanyApiCredentialResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

