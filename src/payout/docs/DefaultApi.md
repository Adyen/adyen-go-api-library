# \DefaultApi

All URIs are relative to *https://pal-test.adyen.com/pal/servlet/Payout/v52*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmThirdPartyPost**](DefaultApi.md#ConfirmThirdPartyPost) | **Post** /confirmThirdParty | Confirms a payout.
[**DeclineThirdPartyPost**](DefaultApi.md#DeclineThirdPartyPost) | **Post** /declineThirdParty | Cancels a payout.
[**PayoutPost**](DefaultApi.md#PayoutPost) | **Post** /payout | Pay out directly.
[**StoreDetailAndSubmitThirdPartyPost**](DefaultApi.md#StoreDetailAndSubmitThirdPartyPost) | **Post** /storeDetailAndSubmitThirdParty | Stores details and submits a payout.
[**StoreDetailPost**](DefaultApi.md#StoreDetailPost) | **Post** /storeDetail | Stores payout details.
[**SubmitThirdPartyPost**](DefaultApi.md#SubmitThirdPartyPost) | **Post** /submitThirdParty | Submits a payout.



## ConfirmThirdPartyPost

> ModifyResponse ConfirmThirdPartyPost(ctx, optional)

Confirms a payout.

Confirms a previously submitted payout.  To cancel a payout, use the `/declineThirdParty` endpoint.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfirmThirdPartyPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfirmThirdPartyPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modifyRequest** | [**optional.Interface of ModifyRequest**](ModifyRequest.md)|  | 

### Return type

[**ModifyResponse**](ModifyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeclineThirdPartyPost

> ModifyResponse DeclineThirdPartyPost(ctx, optional)

Cancels a payout.

Cancels a previously submitted payout.  To confirm and send a payout, use the `/confirmThirdParty` endpoint.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeclineThirdPartyPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeclineThirdPartyPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modifyRequest** | [**optional.Interface of ModifyRequest**](ModifyRequest.md)|  | 

### Return type

[**ModifyResponse**](ModifyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PayoutPost

> PayoutResponse PayoutPost(ctx, optional)

Pay out directly.

With this call, you can pay out to your customers, and funds will be made available within 30 minutes on the cardholder's bank account (this is dependent on whether the issuer supports this functionality). Instant card payouts are only supported for Visa and Mastercard cards.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PayoutPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PayoutPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payoutRequest** | [**optional.Interface of PayoutRequest**](PayoutRequest.md)|  | 

### Return type

[**PayoutResponse**](PayoutResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreDetailAndSubmitThirdPartyPost

> StoreDetailAndSubmitResponse StoreDetailAndSubmitThirdPartyPost(ctx, optional)

Stores details and submits a payout.

Submits a payout and stores its details for subsequent payouts.  The submitted payout must be confirmed or declined either by a reviewer or via `/confirmThirdParty` or `/declineThirdParty` calls.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StoreDetailAndSubmitThirdPartyPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StoreDetailAndSubmitThirdPartyPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeDetailAndSubmitRequest** | [**optional.Interface of StoreDetailAndSubmitRequest**](StoreDetailAndSubmitRequest.md)|  | 

### Return type

[**StoreDetailAndSubmitResponse**](StoreDetailAndSubmitResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreDetailPost

> StoreDetailResponse StoreDetailPost(ctx, optional)

Stores payout details.

Stores payment details under the `PAYOUT` recurring contract. These payment details can be used later to submit a payout via the `/submitThirdParty` call.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StoreDetailPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StoreDetailPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeDetailRequest** | [**optional.Interface of StoreDetailRequest**](StoreDetailRequest.md)|  | 

### Return type

[**StoreDetailResponse**](StoreDetailResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitThirdPartyPost

> SubmitResponse SubmitThirdPartyPost(ctx, optional)

Submits a payout.

Submits a payout using the previously stored payment details. To store payment details, use the `/storeDetail` API call.  The submitted payout must be confirmed or declined either by a reviewer or via `/confirmThirdParty` or `/declineThirdParty` calls.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SubmitThirdPartyPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SubmitThirdPartyPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitRequest** | [**optional.Interface of SubmitRequest**](SubmitRequest.md)|  | 

### Return type

[**SubmitResponse**](SubmitResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

