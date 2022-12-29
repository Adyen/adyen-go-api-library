# \GeneralApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostOnboardingLinks**](GeneralApi.md#PostOnboardingLinks) | **Post** /onboardingLinks | Creates link to account onboarding page



## PostOnboardingLinks

> OnboardingLink PostOnboardingLinks(ctx).CreateOnboardingLinkRequest(createOnboardingLinkRequest).Execute()

Creates link to account onboarding page



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
    createOnboardingLinkRequest := *openapiclient.NewCreateOnboardingLinkRequest(*openapiclient.NewUser2("Email_example", *openapiclient.NewName2()), "Country_example") // CreateOnboardingLinkRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GeneralApi.PostOnboardingLinks(context.Background()).CreateOnboardingLinkRequest(createOnboardingLinkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneralApi.PostOnboardingLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostOnboardingLinks`: OnboardingLink
    fmt.Fprintf(os.Stdout, "Response from `GeneralApi.PostOnboardingLinks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostOnboardingLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOnboardingLinkRequest** | [**CreateOnboardingLinkRequest**](CreateOnboardingLinkRequest.md) |  | 

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

