# \UsersCompanyLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompaniesCompanyIdUsers**](UsersCompanyLevelApi.md#GetCompaniesCompanyIdUsers) | **Get** /companies/{companyId}/users | Get a list of users
[**GetCompaniesCompanyIdUsersUserId**](UsersCompanyLevelApi.md#GetCompaniesCompanyIdUsersUserId) | **Get** /companies/{companyId}/users/{userId} | Get user details
[**PatchCompaniesCompanyIdUsersUserId**](UsersCompanyLevelApi.md#PatchCompaniesCompanyIdUsersUserId) | **Patch** /companies/{companyId}/users/{userId} | Update user details
[**PostCompaniesCompanyIdUsers**](UsersCompanyLevelApi.md#PostCompaniesCompanyIdUsers) | **Post** /companies/{companyId}/users | Create a new user



## GetCompaniesCompanyIdUsers

> ListCompanyUsersResponse GetCompaniesCompanyIdUsers(ctx, companyId).PageNumber(pageNumber).PageSize(pageSize).Execute()

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
    companyId := "companyId_example" // string | The unique identifier of the company account.
    pageNumber := int32(56) // int32 | The number of the page to return. (optional)
    pageSize := int32(56) // int32 | The number of items to have on a page. Maximum value is **100**. The default is **10** items on a page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersCompanyLevelApi.GetCompaniesCompanyIdUsers(context.Background(), companyId).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersCompanyLevelApi.GetCompaniesCompanyIdUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompaniesCompanyIdUsers`: ListCompanyUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersCompanyLevelApi.GetCompaniesCompanyIdUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesCompanyIdUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | The number of the page to return. | 
 **pageSize** | **int32** | The number of items to have on a page. Maximum value is **100**. The default is **10** items on a page. | 

### Return type

[**ListCompanyUsersResponse**](ListCompanyUsersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompaniesCompanyIdUsersUserId

> CompanyUser GetCompaniesCompanyIdUsersUserId(ctx, companyId, userId).Execute()

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
    companyId := "companyId_example" // string | The unique identifier of the company account.
    userId := "userId_example" // string | The unique identifier of the user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersCompanyLevelApi.GetCompaniesCompanyIdUsersUserId(context.Background(), companyId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersCompanyLevelApi.GetCompaniesCompanyIdUsersUserId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompaniesCompanyIdUsersUserId`: CompanyUser
    fmt.Fprintf(os.Stdout, "Response from `UsersCompanyLevelApi.GetCompaniesCompanyIdUsersUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesCompanyIdUsersUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CompanyUser**](CompanyUser.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCompaniesCompanyIdUsersUserId

> CompanyUser PatchCompaniesCompanyIdUsersUserId(ctx, companyId, userId).UpdateCompanyUserRequest(updateCompanyUserRequest).Execute()

Update user details



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
    userId := "userId_example" // string | The unique identifier of the user.
    updateCompanyUserRequest := *openapiclient.NewUpdateCompanyUserRequest() // UpdateCompanyUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersCompanyLevelApi.PatchCompaniesCompanyIdUsersUserId(context.Background(), companyId, userId).UpdateCompanyUserRequest(updateCompanyUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersCompanyLevelApi.PatchCompaniesCompanyIdUsersUserId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCompaniesCompanyIdUsersUserId`: CompanyUser
    fmt.Fprintf(os.Stdout, "Response from `UsersCompanyLevelApi.PatchCompaniesCompanyIdUsersUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompaniesCompanyIdUsersUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateCompanyUserRequest** | [**UpdateCompanyUserRequest**](UpdateCompanyUserRequest.md) |  | 

### Return type

[**CompanyUser**](CompanyUser.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompaniesCompanyIdUsers

> CreateCompanyUserResponse PostCompaniesCompanyIdUsers(ctx, companyId).CreateCompanyUserRequest(createCompanyUserRequest).Execute()

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
    companyId := "companyId_example" // string | The unique identifier of the company account.
    createCompanyUserRequest := *openapiclient.NewCreateCompanyUserRequest("Email_example", *openapiclient.NewName("FirstName_example", "LastName_example"), "Username_example") // CreateCompanyUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersCompanyLevelApi.PostCompaniesCompanyIdUsers(context.Background(), companyId).CreateCompanyUserRequest(createCompanyUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersCompanyLevelApi.PostCompaniesCompanyIdUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCompaniesCompanyIdUsers`: CreateCompanyUserResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersCompanyLevelApi.PostCompaniesCompanyIdUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompaniesCompanyIdUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createCompanyUserRequest** | [**CreateCompanyUserRequest**](CreateCompanyUserRequest.md) |  | 

### Return type

[**CreateCompanyUserResponse**](CreateCompanyUserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

