# \ClientKeyCompanyLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCompaniesCompanyIdApiCredentialsApiCredentialIdGenerateClientKey**](ClientKeyCompanyLevelApi.md#PostCompaniesCompanyIdApiCredentialsApiCredentialIdGenerateClientKey) | **Post** /companies/{companyId}/apiCredentials/{apiCredentialId}/generateClientKey | Generate new client key



## PostCompaniesCompanyIdApiCredentialsApiCredentialIdGenerateClientKey

> GenerateClientKeyResponse PostCompaniesCompanyIdApiCredentialsApiCredentialIdGenerateClientKey(ctx, companyId, apiCredentialId).Execute()

Generate new client key



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
    resp, r, err := apiClient.ClientKeyCompanyLevelApi.PostCompaniesCompanyIdApiCredentialsApiCredentialIdGenerateClientKey(context.Background(), companyId, apiCredentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientKeyCompanyLevelApi.PostCompaniesCompanyIdApiCredentialsApiCredentialIdGenerateClientKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCompaniesCompanyIdApiCredentialsApiCredentialIdGenerateClientKey`: GenerateClientKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientKeyCompanyLevelApi.PostCompaniesCompanyIdApiCredentialsApiCredentialIdGenerateClientKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 
**apiCredentialId** | **string** | Unique identifier of the API credential. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompaniesCompanyIdApiCredentialsApiCredentialIdGenerateClientKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GenerateClientKeyResponse**](GenerateClientKeyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

