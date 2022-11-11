# \UsersMerchantLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMerchantsMerchantIdUsers**](UsersMerchantLevelApi.md#GetMerchantsMerchantIdUsers) | **Get** /merchants/{merchantId}/users | Get a list of users
[**GetMerchantsMerchantIdUsersUserId**](UsersMerchantLevelApi.md#GetMerchantsMerchantIdUsersUserId) | **Get** /merchants/{merchantId}/users/{userId} | Get user details
[**PatchMerchantsMerchantIdUsersUserId**](UsersMerchantLevelApi.md#PatchMerchantsMerchantIdUsersUserId) | **Patch** /merchants/{merchantId}/users/{userId} | Update a user
[**PostMerchantsMerchantIdUsers**](UsersMerchantLevelApi.md#PostMerchantsMerchantIdUsers) | **Post** /merchants/{merchantId}/users | Create a new user



## GetMerchantsMerchantIdUsers

> ListMerchantUsersResponse GetMerchantsMerchantIdUsers(ctx, merchantId).PageNumber(pageNumber).PageSize(pageSize).Execute()

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
    merchantId := "merchantId_example" // string | Unique identifier of the merchant.
    pageNumber := int32(56) // int32 | The number of the page to fetch. (optional)
    pageSize := int32(56) // int32 | The number of items to have on a page. Maximum value is **100**. The default is **10** items on a page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersMerchantLevelApi.GetMerchantsMerchantIdUsers(context.Background(), merchantId).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMerchantLevelApi.GetMerchantsMerchantIdUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsMerchantIdUsers`: ListMerchantUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersMerchantLevelApi.GetMerchantsMerchantIdUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | Unique identifier of the merchant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsMerchantIdUsersRequest struct via the builder pattern


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


## GetMerchantsMerchantIdUsersUserId

> User GetMerchantsMerchantIdUsersUserId(ctx, merchantId, userId).Execute()

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
    merchantId := "merchantId_example" // string | Unique identifier of the merchant.
    userId := "userId_example" // string | Unique identifier of the user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersMerchantLevelApi.GetMerchantsMerchantIdUsersUserId(context.Background(), merchantId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMerchantLevelApi.GetMerchantsMerchantIdUsersUserId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantsMerchantIdUsersUserId`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersMerchantLevelApi.GetMerchantsMerchantIdUsersUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | Unique identifier of the merchant. | 
**userId** | **string** | Unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantsMerchantIdUsersUserIdRequest struct via the builder pattern


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


## PatchMerchantsMerchantIdUsersUserId

> User PatchMerchantsMerchantIdUsersUserId(ctx, merchantId, userId).UpdateMerchantUserRequest(updateMerchantUserRequest).Execute()

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
    merchantId := "merchantId_example" // string | Unique identifier of the merchant.
    userId := "userId_example" // string | Unique identifier of the user.
    updateMerchantUserRequest := *openapiclient.NewUpdateMerchantUserRequest() // UpdateMerchantUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersMerchantLevelApi.PatchMerchantsMerchantIdUsersUserId(context.Background(), merchantId, userId).UpdateMerchantUserRequest(updateMerchantUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMerchantLevelApi.PatchMerchantsMerchantIdUsersUserId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMerchantsMerchantIdUsersUserId`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersMerchantLevelApi.PatchMerchantsMerchantIdUsersUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | Unique identifier of the merchant. | 
**userId** | **string** | Unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMerchantsMerchantIdUsersUserIdRequest struct via the builder pattern


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


## PostMerchantsMerchantIdUsers

> CreateUserResponse PostMerchantsMerchantIdUsers(ctx, merchantId).CreateMerchantUserRequest(createMerchantUserRequest).Execute()

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
    merchantId := "merchantId_example" // string | Unique identifier of the merchant.
    createMerchantUserRequest := *openapiclient.NewCreateMerchantUserRequest("Email_example", *openapiclient.NewName("FirstName_example", "LastName_example"), "Username_example") // CreateMerchantUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersMerchantLevelApi.PostMerchantsMerchantIdUsers(context.Background(), merchantId).CreateMerchantUserRequest(createMerchantUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMerchantLevelApi.PostMerchantsMerchantIdUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMerchantsMerchantIdUsers`: CreateUserResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersMerchantLevelApi.PostMerchantsMerchantIdUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | **string** | Unique identifier of the merchant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMerchantsMerchantIdUsersRequest struct via the builder pattern


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

