# \UsersMerchantLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMerchantsIdUsers**](UsersMerchantLevelApi.md#GetMerchantsIdUsers) | **Get** /merchants/{id}/users | Get a list of users
[**GetMerchantsIdUsersUserId**](UsersMerchantLevelApi.md#GetMerchantsIdUsersUserId) | **Get** /merchants/{id}/users/{userId} | Get user details
[**PatchMerchantsIdUsersUserId**](UsersMerchantLevelApi.md#PatchMerchantsIdUsersUserId) | **Patch** /merchants/{id}/users/{userId} | Update a user
[**PostMerchantsIdUsers**](UsersMerchantLevelApi.md#PostMerchantsIdUsers) | **Post** /merchants/{id}/users | Create a new user



## GetMerchantsIdUsers

> ListMerchantUsersResponse GetMerchantsIdUsers(ctx, id).PageNumber(pageNumber).PageSize(pageSize).Execute()

Get a list of users



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
    id := "id_example" // string | Unique identifier of the merchant.
    pageNumber := int32(56) // int32 | The number of the page to fetch. (optional)
    pageSize := int32(56) // int32 | The number of items to have on a page. Maximum value is **100**. The default is **10** items on a page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersMerchantLevelApi.GetMerchantsIdUsers(context.Background(), id).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMerchantLevelApi.GetMerchantsIdUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsIdUsers`: ListMerchantUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersMerchantLevelApi.GetMerchantsIdUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the merchant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsIdUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | The number of the page to fetch. | 
 **pageSize** | **int32** | The number of items to have on a page. Maximum value is **100**. The default is **10** items on a page. | 

### Return type

[**ListMerchantUsersResponse**](ListMerchantUsersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMerchantsIdUsersUserId

> User GetMerchantsIdUsersUserId(ctx, id, userId).Execute()

Get user details



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
    id := "id_example" // string | Unique identifier of the merchant.
    userId := "userId_example" // string | Unique identifier of the user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersMerchantLevelApi.GetMerchantsIdUsersUserId(context.Background(), id, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMerchantLevelApi.GetMerchantsIdUsersUserId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsIdUsersUserId`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersMerchantLevelApi.GetMerchantsIdUsersUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the merchant. | 
**userId** | **string** | Unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsIdUsersUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**User**](User.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMerchantsIdUsersUserId

> User PatchMerchantsIdUsersUserId(ctx, id, userId).UpdateMerchantUserRequest(updateMerchantUserRequest).Execute()

Update a user



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
    id := "id_example" // string | Unique identifier of the merchant.
    userId := "userId_example" // string | Unique identifier of the user.
    updateMerchantUserRequest := *openapiclient.NewUpdateMerchantUserRequest() // UpdateMerchantUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersMerchantLevelApi.PatchMerchantsIdUsersUserId(context.Background(), id, userId).UpdateMerchantUserRequest(updateMerchantUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMerchantLevelApi.PatchMerchantsIdUsersUserId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMerchantsIdUsersUserId`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersMerchantLevelApi.PatchMerchantsIdUsersUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the merchant. | 
**userId** | **string** | Unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMerchantsIdUsersUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateMerchantUserRequest** | [**UpdateMerchantUserRequest**](UpdateMerchantUserRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMerchantsIdUsers

> CreateUserResponse PostMerchantsIdUsers(ctx, id).CreateMerchantUserRequest(createMerchantUserRequest).Execute()

Create a new user



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
    id := "id_example" // string | Unique identifier of the merchant.
    createMerchantUserRequest := *openapiclient.NewCreateMerchantUserRequest("Email_example", *openapiclient.NewName("FirstName_example", "LastName_example"), "Username_example") // CreateMerchantUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersMerchantLevelApi.PostMerchantsIdUsers(context.Background(), id).CreateMerchantUserRequest(createMerchantUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMerchantLevelApi.PostMerchantsIdUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMerchantsIdUsers`: CreateUserResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersMerchantLevelApi.PostMerchantsIdUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the merchant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMerchantsIdUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createMerchantUserRequest** | [**CreateMerchantUserRequest**](CreateMerchantUserRequest.md) |  | 

### Return type

[**CreateUserResponse**](CreateUserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

