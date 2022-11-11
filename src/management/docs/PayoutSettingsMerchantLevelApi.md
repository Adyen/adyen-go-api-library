# \PayoutSettingsMerchantLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMerchantsMerchantIdPayoutSettingsPayoutSettingsId**](PayoutSettingsMerchantLevelApi.md#DeleteMerchantsMerchantIdPayoutSettingsPayoutSettingsId) | **Delete** /merchants/{merchantId}/payoutSettings/{payoutSettingsId} | Delete a payout setting
[**GetMerchantsMerchantIdPayoutSettings**](PayoutSettingsMerchantLevelApi.md#GetMerchantsMerchantIdPayoutSettings) | **Get** /merchants/{merchantId}/payoutSettings | Get a list of payout settings
[**GetMerchantsMerchantIdPayoutSettingsPayoutSettingsId**](PayoutSettingsMerchantLevelApi.md#GetMerchantsMerchantIdPayoutSettingsPayoutSettingsId) | **Get** /merchants/{merchantId}/payoutSettings/{payoutSettingsId} | Get a payout setting
[**PatchMerchantsMerchantIdPayoutSettingsPayoutSettingsId**](PayoutSettingsMerchantLevelApi.md#PatchMerchantsMerchantIdPayoutSettingsPayoutSettingsId) | **Patch** /merchants/{merchantId}/payoutSettings/{payoutSettingsId} | Update a payout setting
[**PostMerchantsMerchantIdPayoutSettings**](PayoutSettingsMerchantLevelApi.md#PostMerchantsMerchantIdPayoutSettings) | **Post** /merchants/{merchantId}/payoutSettings | Add a payout setting



## DeleteMerchantsMerchantIdPayoutSettingsPayoutSettingsId

> interface{} DeleteMerchantsMerchantIdPayoutSettingsPayoutSettingsId(ctx, merchantId, payoutSettingsId).Execute()

Delete a payout setting



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
    payoutSettingsId := "payoutSettingsId_example" // string | The unique identifier of the payout setting.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayoutSettingsMerchantLevelApi.DeleteMerchantsMerchantIdPayoutSettingsPayoutSettingsId(context.Background(), merchantId, payoutSettingsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutSettingsMerchantLevelApi.DeleteMerchantsMerchantIdPayoutSettingsPayoutSettingsId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMerchantsMerchantIdPayoutSettingsPayoutSettingsId`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `PayoutSettingsMerchantLevelApi.DeleteMerchantsMerchantIdPayoutSettingsPayoutSettingsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**payoutSettingsId** | **string** | The unique identifier of the payout setting. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMerchantsMerchantIdPayoutSettingsPayoutSettingsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMerchantsMerchantIdPayoutSettings

> PayoutSettingsResponse GetMerchantsMerchantIdPayoutSettings(ctx, merchantId).Execute()

Get a list of payout settings



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayoutSettingsMerchantLevelApi.GetMerchantsMerchantIdPayoutSettings(context.Background(), merchantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutSettingsMerchantLevelApi.GetMerchantsMerchantIdPayoutSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsMerchantIdPayoutSettings`: PayoutSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `PayoutSettingsMerchantLevelApi.GetMerchantsMerchantIdPayoutSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsMerchantIdPayoutSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PayoutSettingsResponse**](PayoutSettingsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMerchantsMerchantIdPayoutSettingsPayoutSettingsId

> PayoutSettings GetMerchantsMerchantIdPayoutSettingsPayoutSettingsId(ctx, merchantId, payoutSettingsId).Execute()

Get a payout setting



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
    payoutSettingsId := "payoutSettingsId_example" // string | The unique identifier of the payout setting.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayoutSettingsMerchantLevelApi.GetMerchantsMerchantIdPayoutSettingsPayoutSettingsId(context.Background(), merchantId, payoutSettingsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutSettingsMerchantLevelApi.GetMerchantsMerchantIdPayoutSettingsPayoutSettingsId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsMerchantIdPayoutSettingsPayoutSettingsId`: PayoutSettings
    fmt.Fprintf(os.Stdout, "Response from `PayoutSettingsMerchantLevelApi.GetMerchantsMerchantIdPayoutSettingsPayoutSettingsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**payoutSettingsId** | **string** | The unique identifier of the payout setting. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsMerchantIdPayoutSettingsPayoutSettingsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PayoutSettings**](PayoutSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMerchantsMerchantIdPayoutSettingsPayoutSettingsId

> PayoutSettings PatchMerchantsMerchantIdPayoutSettingsPayoutSettingsId(ctx, merchantId, payoutSettingsId).UpdatePayoutSettingsRequest(updatePayoutSettingsRequest).Execute()

Update a payout setting



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
    payoutSettingsId := "payoutSettingsId_example" // string | The unique identifier of the payout setting.
    updatePayoutSettingsRequest := *openapiclient.NewUpdatePayoutSettingsRequest() // UpdatePayoutSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayoutSettingsMerchantLevelApi.PatchMerchantsMerchantIdPayoutSettingsPayoutSettingsId(context.Background(), merchantId, payoutSettingsId).UpdatePayoutSettingsRequest(updatePayoutSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutSettingsMerchantLevelApi.PatchMerchantsMerchantIdPayoutSettingsPayoutSettingsId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMerchantsMerchantIdPayoutSettingsPayoutSettingsId`: PayoutSettings
    fmt.Fprintf(os.Stdout, "Response from `PayoutSettingsMerchantLevelApi.PatchMerchantsMerchantIdPayoutSettingsPayoutSettingsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 
**payoutSettingsId** | **string** | The unique identifier of the payout setting. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMerchantsMerchantIdPayoutSettingsPayoutSettingsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updatePayoutSettingsRequest** | [**UpdatePayoutSettingsRequest**](UpdatePayoutSettingsRequest.md) |  | 

### Return type

[**PayoutSettings**](PayoutSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMerchantsMerchantIdPayoutSettings

> PayoutSettings PostMerchantsMerchantIdPayoutSettings(ctx, merchantId).PayoutSettingsRequest(payoutSettingsRequest).Execute()

Add a payout setting



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
    payoutSettingsRequest := *openapiclient.NewPayoutSettingsRequest("TransferInstrumentId_example") // PayoutSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayoutSettingsMerchantLevelApi.PostMerchantsMerchantIdPayoutSettings(context.Background(), merchantId).PayoutSettingsRequest(payoutSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutSettingsMerchantLevelApi.PostMerchantsMerchantIdPayoutSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMerchantsMerchantIdPayoutSettings`: PayoutSettings
    fmt.Fprintf(os.Stdout, "Response from `PayoutSettingsMerchantLevelApi.PostMerchantsMerchantIdPayoutSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | The unique identifier of the merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMerchantsMerchantIdPayoutSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payoutSettingsRequest** | [**PayoutSettingsRequest**](PayoutSettingsRequest.md) |  | 

### Return type

[**PayoutSettings**](PayoutSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

