# \ClientKeyMerchantLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostMerchantsMerchantIdApiCredentialsApiCredentialIdGenerateClientKey**](ClientKeyMerchantLevelApi.md#PostMerchantsMerchantIdApiCredentialsApiCredentialIdGenerateClientKey) | **Post** /merchants/{merchantId}/apiCredentials/{apiCredentialId}/generateClientKey | Generate new client key



## PostMerchantsMerchantIdApiCredentialsApiCredentialIdGenerateClientKey

> GenerateClientKeyResponse PostMerchantsMerchantIdApiCredentialsApiCredentialIdGenerateClientKey(ctx, merchantId, apiCredentialId).Execute()

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
    merchantId := "merchantId_example" // string | The unique identifier of the merchant account.
    apiCredentialId := "apiCredentialId_example" // string | Unique identifier of the API credential.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientKeyMerchantLevelApi.PostMerchantsMerchantIdApiCredentialsApiCredentialIdGenerateClientKey(context.Background(), merchantId, apiCredentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientKeyMerchantLevelApi.PostMerchantsMerchantIdApiCredentialsApiCredentialIdGenerateClientKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMerchantsMerchantIdApiCredentialsApiCredentialIdGenerateClientKey`: GenerateClientKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientKeyMerchantLevelApi.PostMerchantsMerchantIdApiCredentialsApiCredentialIdGenerateClientKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**apiCredentialId** | **string** | Unique identifier of the API credential. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMerchantsMerchantIdApiCredentialsApiCredentialIdGenerateClientKeyRequest struct via the builder pattern


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

