# \DefaultApi

All URIs are relative to *https://pal-test.adyen.com/pal/servlet/Recurring/v49*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisablePost**](DefaultApi.md#DisablePost) | **Post** /disable | Disables stored payment details.
[**ListRecurringDetailsPost**](DefaultApi.md#ListRecurringDetailsPost) | **Post** /listRecurringDetails | Retrieves stored payment details for a shopper.
[**ScheduleAccountUpdaterPost**](DefaultApi.md#ScheduleAccountUpdaterPost) | **Post** /scheduleAccountUpdater | Schedules running of the Account Updater.



## DisablePost

> DisableResult DisablePost(ctx, optional)

Disables stored payment details.

Disables stored payment details to stop charging a shopper with this particular recurring detail ID.  For more information, refer to [Disable stored details](https://docs.adyen.com/classic-integration/recurring-payments/disable-stored-details/).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DisablePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DisablePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableRequest** | [**optional.Interface of DisableRequest**](DisableRequest.md)|  | 

### Return type

[**DisableResult**](DisableResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecurringDetailsPost

> RecurringDetailsResult ListRecurringDetailsPost(ctx, optional)

Retrieves stored payment details for a shopper.

Lists the stored payment details for a shopper, if there are any available. The recurring detail ID can be used with a regular authorisation request to charge the shopper. A summary of the payment detail is returned for presentation to the shopper.  For more information, refer to [Retrieve stored details](https://docs.adyen.com/classic-integration/recurring-payments/retrieve-stored-details/).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListRecurringDetailsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRecurringDetailsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recurringDetailsRequest** | [**optional.Interface of RecurringDetailsRequest**](RecurringDetailsRequest.md)|  | 

### Return type

[**RecurringDetailsResult**](RecurringDetailsResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleAccountUpdaterPost

> ScheduleAccountUpdaterResult ScheduleAccountUpdaterPost(ctx, optional)

Schedules running of the Account Updater.

When making the API call, you can submit either the credit card information, or the recurring detail reference and the shopper reference: * If the card information is provided, all the sub-fields for `card` are mandatory. * If the recurring detail reference is provided, the fields for `shopperReference` and `selectedRecurringDetailReference` are mandatory.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ScheduleAccountUpdaterPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ScheduleAccountUpdaterPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scheduleAccountUpdaterRequest** | [**optional.Interface of ScheduleAccountUpdaterRequest**](ScheduleAccountUpdaterRequest.md)|  | 

### Return type

[**ScheduleAccountUpdaterResult**](ScheduleAccountUpdaterResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

