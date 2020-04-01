# \DefaultApi

All URIs are relative to *https://checkout-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OriginKeysPost**](DefaultApi.md#OriginKeysPost) | **Post** /originKeys | Create originKey values for one or more merchant domains.



## OriginKeysPost

> CheckoutUtilityResponse OriginKeysPost(ctx, optional)

Create originKey values for one or more merchant domains.

This operation takes the origin domains and returns a JSON object containing the corresponding origin keys for the domains.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OriginKeysPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OriginKeysPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkoutUtilityRequest** | [**optional.Interface of CheckoutUtilityRequest**](CheckoutUtilityRequest.md)|  | 

### Return type

[**CheckoutUtilityResponse**](CheckoutUtilityResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

