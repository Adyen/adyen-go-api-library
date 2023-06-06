# \HostedOnboardingApi

All URIs are relative to *https://kyc-test.adyen.com/lem/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListHostedOnboardingPageThemes**](HostedOnboardingApi.md#ListHostedOnboardingPageThemes) | **Get** /themes | Get a list of hosted onboarding page themes
[**GetOnboardingLinkTheme**](HostedOnboardingApi.md#GetOnboardingLinkTheme) | **Get** /themes/{id} | Get an onboarding link theme
[**GetLinkToAdyenhostedOnboardingPage**](HostedOnboardingApi.md#GetLinkToAdyenhostedOnboardingPage) | **Post** /legalEntities/{id}/onboardingLinks | Get a link to an Adyen-hosted onboarding page



## ListHostedOnboardingPageThemes

> OnboardingThemes ListHostedOnboardingPageThemes(ctx).Execute()

Get a list of hosted onboarding page themes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/adyen/adyen-go-api-library/v7/src/legalentity"
)

func main() {

    username := "YOUR_USERNAME"
    password := "YOUR_PASSWORD"
    env := openapiclient.TestEnv

    configuration, err := openapiclient.NewConfiguration(username, password, env)
    apiClient := openapiclient.NewAPIClient(configuration)

    resp, r, err := apiClient.HostedOnboardingApi.ListHostedOnboardingPageThemes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostedOnboardingApi.ListHostedOnboardingPageThemes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHostedOnboardingPageThemes`: OnboardingThemes
    fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetThemesRequest struct via the builder pattern


### Return type

[**OnboardingThemes**](OnboardingThemes.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnboardingLinkTheme

> OnboardingTheme GetOnboardingLinkTheme(ctx, id).Execute()

Get an onboarding link theme



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/adyen/adyen-go-api-library/v7/src/legalentity"
)

func main() {
    id := "id_example" // string | The unique identifier of the theme

    username := "YOUR_USERNAME"
    password := "YOUR_PASSWORD"
    env := openapiclient.TestEnv

    configuration, err := openapiclient.NewConfiguration(username, password, env)
    apiClient := openapiclient.NewAPIClient(configuration)

    resp, r, err := apiClient.HostedOnboardingApi.GetOnboardingLinkTheme(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostedOnboardingApi.GetOnboardingLinkTheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOnboardingLinkTheme`: OnboardingTheme
    fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the theme | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThemesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OnboardingTheme**](OnboardingTheme.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkToAdyenhostedOnboardingPage

> OnboardingLink GetLinkToAdyenhostedOnboardingPage(ctx, id).OnboardingLinkInfo(onboardingLinkInfo).Execute()

Get a link to an Adyen-hosted onboarding page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/adyen/adyen-go-api-library/v7/src/legalentity"
)

func main() {
    id := "id_example" // string | The unique identifier of the legal entity
    onboardingLinkInfo := *openapiclient.NewOnboardingLinkInfo() // OnboardingLinkInfo |  (optional)

    username := "YOUR_USERNAME"
    password := "YOUR_PASSWORD"
    env := openapiclient.TestEnv

    configuration, err := openapiclient.NewConfiguration(username, password, env)
    apiClient := openapiclient.NewAPIClient(configuration)

    resp, r, err := apiClient.HostedOnboardingApi.GetLinkToAdyenhostedOnboardingPage(context.Background(), id).OnboardingLinkInfo(onboardingLinkInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostedOnboardingApi.GetLinkToAdyenhostedOnboardingPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLinkToAdyenhostedOnboardingPage`: OnboardingLink
    fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the legal entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLegalEntitiesIdOnboardingLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **onboardingLinkInfo** | [**OnboardingLinkInfo**](OnboardingLinkInfo.md) |  | 

### Return type

[**OnboardingLink**](OnboardingLink.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

