# \APIKeyCompanyLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCompaniesCompanyIdApiCredentialsApiCredentialIdGenerateApiKey**](APIKeyCompanyLevelApi.md#PostCompaniesCompanyIdApiCredentialsApiCredentialIdGenerateApiKey) | **Post** /companies/{companyId}/apiCredentials/{apiCredentialId}/generateApiKey | Generate new API key



## PostCompaniesCompanyIdApiCredentialsApiCredentialIdGenerateApiKey

> GenerateApiKeyResponse PostCompaniesCompanyIdApiCredentialsApiCredentialIdGenerateApiKey(ctx, companyId, apiCredentialId).Execute()

Generate new API key



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
    resp, r, err := apiClient.APIKeyCompanyLevelApi.PostCompaniesCompanyIdApiCredentialsApiCredentialIdGenerateApiKey(context.Background(), companyId, apiCredentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeyCompanyLevelApi.PostCompaniesCompanyIdApiCredentialsApiCredentialIdGenerateApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCompaniesCompanyIdApiCredentialsApiCredentialIdGenerateApiKey`: GenerateApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `APIKeyCompanyLevelApi.PostCompaniesCompanyIdApiCredentialsApiCredentialIdGenerateApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 
**apiCredentialId** | **string** | Unique identifier of the API credential. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompaniesCompanyIdApiCredentialsApiCredentialIdGenerateApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GenerateApiKeyResponse**](GenerateApiKeyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

