# \ClientKeyMerchantLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostMerchantsIdApiCredentialsApiCredentialIdGenerateClientKey**](ClientKeyMerchantLevelApi.md#PostMerchantsIdApiCredentialsApiCredentialIdGenerateClientKey) | **Post** /merchants/{id}/apiCredentials/{apiCredentialId}/generateClientKey | Generate new client key



## PostMerchantsIdApiCredentialsApiCredentialIdGenerateClientKey

> GenerateClientKeyResponse PostMerchantsIdApiCredentialsApiCredentialIdGenerateClientKey(ctx, id, apiCredentialId).Execute()

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
    id := "id_example" // string | The unique identifier of the merchant account.
    apiCredentialId := "apiCredentialId_example" // string | Unique identifier of the API credential.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientKeyMerchantLevelApi.PostMerchantsIdApiCredentialsApiCredentialIdGenerateClientKey(context.Background(), id, apiCredentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientKeyMerchantLevelApi.PostMerchantsIdApiCredentialsApiCredentialIdGenerateClientKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMerchantsIdApiCredentialsApiCredentialIdGenerateClientKey`: GenerateClientKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientKeyMerchantLevelApi.PostMerchantsIdApiCredentialsApiCredentialIdGenerateClientKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the merchant account. | 
**apiCredentialId** | **string** | Unique identifier of the API credential. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMerchantsIdApiCredentialsApiCredentialIdGenerateClientKeyRequest struct via the builder pattern


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

