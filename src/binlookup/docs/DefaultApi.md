# \DefaultApi

All URIs are relative to *https://pal-test.adyen.com/pal/servlet/BinLookup/v50*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get3dsAvailabilityPost**](DefaultApi.md#Get3dsAvailabilityPost) | **Post** /get3dsAvailability | Checks 3D Secure availability.
[**GetCostEstimatePost**](DefaultApi.md#GetCostEstimatePost) | **Post** /getCostEstimate | Gets a cost estimate.



## Get3dsAvailabilityPost

> ThreeDSAvailabilityResponse Get3dsAvailabilityPost(ctx, optional)

Checks 3D Secure availability.

Verifies whether 3D Secure is available for the specified BIN or card brand. For 3D Secure 2, this endpoint also returns device fingerprinting keys.  For more information, refer to [3D Secure 2](https://docs.adyen.com/checkout/3d-secure/native-3ds2).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Get3dsAvailabilityPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Get3dsAvailabilityPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **threeDSAvailabilityRequest** | [**optional.Interface of ThreeDSAvailabilityRequest**](ThreeDSAvailabilityRequest.md)|  | 

### Return type

[**ThreeDSAvailabilityResponse**](ThreeDSAvailabilityResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCostEstimatePost

> CostEstimateResponse GetCostEstimatePost(ctx, optional)

Gets a cost estimate.

Use the Adyen Cost Estimation API to pre-calculate interchange and scheme fee costs. Knowing these costs prior actual payment authorisation gives you an opportunity to charge those costs to the cardholder, if necessary.  To retrieve this information, make the call to the `/getCostEstimate` endpoint. The response to this call contains the amount of the interchange and scheme fees charged by the network for this transaction, and also which surcharging policy is possible (based on current regulations).  > Since not all information is known in advance (for example, if the cardholder will successfully authenticate via 3D Secure or if you also plan to provide additional Level 2/3 data), the returned amounts are based on a set of assumption criteria you define in the `assumptions` parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetCostEstimatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCostEstimatePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **costEstimateRequest** | [**optional.Interface of CostEstimateRequest**](CostEstimateRequest.md)|  | 

### Return type

[**CostEstimateResponse**](CostEstimateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

