# \DefaultApi

All URIs are relative to *https://pal-test.adyen.com/pal/servlet/Payment/v52*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdjustAuthorisationPost**](DefaultApi.md#AdjustAuthorisationPost) | **Post** /adjustAuthorisation | Increases or decreases the authorised amount.
[**Authorise3dPost**](DefaultApi.md#Authorise3dPost) | **Post** /authorise3d | Completes a 3D Secure payment authorisation.
[**Authorise3ds2Post**](DefaultApi.md#Authorise3ds2Post) | **Post** /authorise3ds2 | Completes a 3D Secure 2 payment authorisation.
[**AuthorisePost**](DefaultApi.md#AuthorisePost) | **Post** /authorise | Creates a payment authorisation.
[**CancelOrRefundPost**](DefaultApi.md#CancelOrRefundPost) | **Post** /cancelOrRefund | Cancels or refunds a payment.
[**CancelPost**](DefaultApi.md#CancelPost) | **Post** /cancel | Cancels an authorised payment.
[**CapturePost**](DefaultApi.md#CapturePost) | **Post** /capture | Captures an authorised payment.
[**GetAuthenticationResultPost**](DefaultApi.md#GetAuthenticationResultPost) | **Post** /getAuthenticationResult | Return the authentication result after doing a 3D Secure authentication only.
[**RefundPost**](DefaultApi.md#RefundPost) | **Post** /refund | Refunds a captured payment.
[**Retrieve3ds2ResultPost**](DefaultApi.md#Retrieve3ds2ResultPost) | **Post** /retrieve3ds2Result | Retrieves the &#x60;threeDS2Result&#x60; after doing a 3D Secure 2 authentication only.
[**TechnicalCancelPost**](DefaultApi.md#TechnicalCancelPost) | **Post** /technicalCancel | Cancels a payment using your custom reference.
[**VoidPendingRefundPost**](DefaultApi.md#VoidPendingRefundPost) | **Post** /voidPendingRefund | Cancels a POS refund request before it has been completed.



## AdjustAuthorisationPost

> ModificationResult AdjustAuthorisationPost(ctx, optional)

Increases or decreases the authorised amount.

Allows you to increase or decrease the authorised amount after the initial authorisation has taken place. This functionality enables tipping, improving the chances your authorisation will be valid, charging the shopper when they have already left the merchant premises, etc.  For more information, refer to [Adjust Authorisation](https://docs.adyen.com/checkout/adjust-authorisation).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AdjustAuthorisationPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AdjustAuthorisationPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modificationRequest** | [**optional.Interface of ModificationRequest**](ModificationRequest.md)|  | 

### Return type

[**ModificationResult**](ModificationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Authorise3dPost

> PaymentResult Authorise3dPost(ctx, optional)

Completes a 3D Secure payment authorisation.

For an authenticated 3D Secure session, completes the payment authorisation. This endpoint must receive the `md` and `paResponse` parameters that you get from the card issuer after a shopper pays via 3D Secure.  For more information, refer to [3D Secure](https://docs.adyen.com/classic-integration/3d-secure).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Authorise3dPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Authorise3dPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentRequest3d** | [**optional.Interface of PaymentRequest3d**](PaymentRequest3d.md)|  | 

### Return type

[**PaymentResult**](PaymentResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Authorise3ds2Post

> PaymentResult Authorise3ds2Post(ctx, optional)

Completes a 3D Secure 2 payment authorisation.

For an authenticated 3D Secure 2 session, completes the payment authorisation. This endpoint must receive the `threeDS2Token` and `threeDS2Result` parameters.  For more information, refer to [3D Secure 2](https://docs.adyen.com/checkout/3d-secure/native-3ds2).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Authorise3ds2PostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Authorise3ds2PostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentRequest3ds2** | [**optional.Interface of PaymentRequest3ds2**](PaymentRequest3ds2.md)|  | 

### Return type

[**PaymentResult**](PaymentResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorisePost

> PaymentResult AuthorisePost(ctx, optional)

Creates a payment authorisation.

Creates a payment with a unique reference (`pspReference`) and attempts to obtain an authorisation hold. For cards, this amount can be captured or cancelled later. Non-card payment methods typically don't support this and will automatically capture as part of the authorisation.  For more information, refer to [Classic integration](https://docs.adyen.com/classic-integration).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthorisePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuthorisePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentRequest** | [**optional.Interface of PaymentRequest**](PaymentRequest.md)|  | 

### Return type

[**PaymentResult**](PaymentResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelOrRefundPost

> ModificationResult CancelOrRefundPost(ctx, optional)

Cancels or refunds a payment.

Cancels a payment if it has not been captured yet, or refunds it if it has already been captured. This is useful when it is not certain if the payment has been captured or not (for example, when using auto-capture).  > Do not use this request for payments that involve (multiple) partial captures.  For more information, refer to [Cancel or refund](https://docs.adyen.com/checkout/cancel-or-refund).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CancelOrRefundPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CancelOrRefundPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modificationRequest** | [**optional.Interface of ModificationRequest**](ModificationRequest.md)|  | 

### Return type

[**ModificationResult**](ModificationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelPost

> ModificationResult CancelPost(ctx, optional)

Cancels an authorised payment.

Cancels the authorisation hold on a payment, returning a unique reference for this request. You can cancel payments after authorisation only for payment methods that support distinct authorisations and captures.  For more information, refer to [Cancel](https://docs.adyen.com/checkout/cancel).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CancelPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CancelPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modificationRequest** | [**optional.Interface of ModificationRequest**](ModificationRequest.md)|  | 

### Return type

[**ModificationResult**](ModificationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CapturePost

> ModificationResult CapturePost(ctx, optional)

Captures an authorised payment.

Captures the authorisation hold on a payment, returning a unique reference for this request. Usually the full authorisation amount is captured, however it's also possible to capture a smaller amount, which results in cancelling the remaining authorisation balance.  Payment methods, which automatically capture as part of authorisation, don't need to be captured, but submitting a capture request on these transactions will not result in double charges. If immediate or delayed auto-capture is enabled, calling the capture method is not neccessary.  For more information, refer to [Capture](https://docs.adyen.com/checkout/capture).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CapturePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CapturePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modificationRequest** | [**optional.Interface of ModificationRequest**](ModificationRequest.md)|  | 

### Return type

[**ModificationResult**](ModificationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthenticationResultPost

> AuthenticationResultResponse GetAuthenticationResultPost(ctx, optional)

Return the authentication result after doing a 3D Secure authentication only.

Return the authentication result after doing a 3D Secure authentication only.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAuthenticationResultPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAuthenticationResultPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authenticationResultRequest** | [**optional.Interface of AuthenticationResultRequest**](AuthenticationResultRequest.md)|  | 

### Return type

[**AuthenticationResultResponse**](AuthenticationResultResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefundPost

> ModificationResult RefundPost(ctx, optional)

Refunds a captured payment.

Refunds a payment that has previously been captured, returning a unique reference for this request. Refunding can be done on the full captured amount or a partial amount. Multiple (partial) refunds will be accepted as long as their sum doesn't exceed the captured amount. Payments which have been authorised, but not captured, cannot be refunded, use the /cancel method instead.  > Some payment methods/gateways do not support partial/multiple refunds. > A margin above the captured limit can be configured to cover shipping/handling costs.  For more information, refer to [Refund](https://docs.adyen.com/checkout/refund).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RefundPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RefundPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modificationRequest** | [**optional.Interface of ModificationRequest**](ModificationRequest.md)|  | 

### Return type

[**ModificationResult**](ModificationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Retrieve3ds2ResultPost

> ThreeDS2ResultResponse Retrieve3ds2ResultPost(ctx, optional)

Retrieves the `threeDS2Result` after doing a 3D Secure 2 authentication only.

Retrieves the `threeDS2Result` after doing a 3D Secure 2 authentication only.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Retrieve3ds2ResultPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Retrieve3ds2ResultPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **threeDS2ResultRequest** | [**optional.Interface of ThreeDS2ResultRequest**](ThreeDS2ResultRequest.md)|  | 

### Return type

[**ThreeDS2ResultResponse**](ThreeDS2ResultResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TechnicalCancelPost

> ModificationResult TechnicalCancelPost(ctx, optional)

Cancels a payment using your custom reference.

This endpoint allows you to cancel a payment if you do not have the PSP reference of the original payment request available.  In your call, refer to the original payment by using the `reference` that you specified in your payment request.  For more information, see [Technical cancel](https://docs.adyen.com/checkout/cancel#technical-cancel).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TechnicalCancelPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TechnicalCancelPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modificationRequest** | [**optional.Interface of ModificationRequest**](ModificationRequest.md)|  | 

### Return type

[**ModificationResult**](ModificationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoidPendingRefundPost

> ModificationResult VoidPendingRefundPost(ctx, optional)

Cancels a POS refund request before it has been completed.

This endpoint allows you to cancel the refund request before it has been completed.  In your call, you can refer to the original refund request either by using the `tenderReference`, or the `pspReference`. We recommend implementing based on the `tenderReference`, as this is generated for both offline and online transactions.  For more information, refer to [Cancel a refund](https://docs.adyen.com/point-of-sale/refund-payment/cancel-a-pos-refund-request).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VoidPendingRefundPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VoidPendingRefundPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modificationRequest** | [**optional.Interface of ModificationRequest**](ModificationRequest.md)|  | 

### Return type

[**ModificationResult**](ModificationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

