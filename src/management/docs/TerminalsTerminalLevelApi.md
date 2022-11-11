# \TerminalsTerminalLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTerminals**](TerminalsTerminalLevelApi.md#GetTerminals) | **Get** /terminals | Get a list of terminals



## GetTerminals

> ListTerminalsResponse GetTerminals(ctx).SearchQuery(searchQuery).Countries(countries).MerchantIds(merchantIds).StoreIds(storeIds).BrandModels(brandModels).PageNumber(pageNumber).PageSize(pageSize).Execute()

Get a list of terminals



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
    searchQuery := "searchQuery_example" // string | Returns terminals with an ID that contains the specified string. If present, other query parameters are ignored. (optional)
    countries := "countries_example" // string | Returns terminals located in the countries specified by their [two-letter country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). (optional)
    merchantIds := "merchantIds_example" // string | Returns terminals that belong to the merchant accounts specified by their unique merchant account ID. (optional)
    storeIds := "storeIds_example" // string | Returns terminals that are assigned to the [stores](https://docs.adyen.com/api-explorer/#/ManagementService/latest/get/stores) specified by their unique store ID. (optional)
    brandModels := "brandModels_example" // string | Returns terminals of the [models](https://docs.adyen.com/api-explorer/#/ManagementService/latest/get/companies/{companyId}/terminalModels) specified in the format *brand.model*. (optional)
    pageNumber := int32(56) // int32 | The number of the page to fetch. (optional)
    pageSize := int32(56) // int32 | The number of items to have on a page, maximum 100. The default is 20 items on a page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalsTerminalLevelApi.GetTerminals(context.Background()).SearchQuery(searchQuery).Countries(countries).MerchantIds(merchantIds).StoreIds(storeIds).BrandModels(brandModels).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalsTerminalLevelApi.GetTerminals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTerminals`: ListTerminalsResponse
    fmt.Fprintf(os.Stdout, "Response from `TerminalsTerminalLevelApi.GetTerminals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTerminalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchQuery** | **string** | Returns terminals with an ID that contains the specified string. If present, other query parameters are ignored. | 
 **countries** | **string** | Returns terminals located in the countries specified by their [two-letter country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). | 
 **merchantIds** | **string** | Returns terminals that belong to the merchant accounts specified by their unique merchant account ID. | 
 **storeIds** | **string** | Returns terminals that are assigned to the [stores](https://docs.adyen.com/api-explorer/#/ManagementService/latest/get/stores) specified by their unique store ID. | 
 **brandModels** | **string** | Returns terminals of the [models](https://docs.adyen.com/api-explorer/#/ManagementService/latest/get/companies/{companyId}/terminalModels) specified in the format *brand.model*. | 
 **pageNumber** | **int32** | The number of the page to fetch. | 
 **pageSize** | **int32** | The number of items to have on a page, maximum 100. The default is 20 items on a page. | 

### Return type

[**ListTerminalsResponse**](ListTerminalsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

