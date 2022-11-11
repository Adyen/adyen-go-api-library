# \APIKeyMerchantLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostMerchantsMerchantIdApiCredentialsApiCredentialIdGenerateApiKey**](APIKeyMerchantLevelApi.md#PostMerchantsMerchantIdApiCredentialsApiCredentialIdGenerateApiKey) | **Post** /merchants/{merchantId}/apiCredentials/{apiCredentialId}/generateApiKey | Generate new API key



## PostMerchantsMerchantIdApiCredentialsApiCredentialIdGenerateApiKey

> GenerateApiKeyResponse PostMerchantsMerchantIdApiCredentialsApiCredentialIdGenerateApiKey(ctx, merchantId, apiCredentialId).Execute()

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
    merchantId := "merchantId_example" // string | The unique identifier of the merchant account.
    apiCredentialId := "apiCredentialId_example" // string | Unique identifier of the API credential.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIKeyMerchantLevelApi.PostMerchantsMerchantIdApiCredentialsApiCredentialIdGenerateApiKey(context.Background(), merchantId, apiCredentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeyMerchantLevelApi.PostMerchantsMerchantIdApiCredentialsApiCredentialIdGenerateApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMerchantsMerchantIdApiCredentialsApiCredentialIdGenerateApiKey`: GenerateApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `APIKeyMerchantLevelApi.PostMerchantsMerchantIdApiCredentialsApiCredentialIdGenerateApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**apiCredentialId** | **string** | Unique identifier of the API credential. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMerchantsMerchantIdApiCredentialsApiCredentialIdGenerateApiKeyRequest struct via the builder pattern


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

