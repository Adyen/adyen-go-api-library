# \DefaultApi

All URIs are relative to *https://checkout-test.adyen.com/v51*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaymentLinksPost**](DefaultApi.md#PaymentLinksPost) | **Post** /paymentLinks | Creates a payment link.
[**PaymentMethodsPost**](DefaultApi.md#PaymentMethodsPost) | **Post** /paymentMethods | Returns available payment methods.
[**PaymentSessionPost**](DefaultApi.md#PaymentSessionPost) | **Post** /paymentSession | Creates a payment session.
[**PaymentsDetailsPost**](DefaultApi.md#PaymentsDetailsPost) | **Post** /payments/details | Submits details for a payment.
[**PaymentsPost**](DefaultApi.md#PaymentsPost) | **Post** /payments | Starts a transaction.
[**PaymentsResultPost**](DefaultApi.md#PaymentsResultPost) | **Post** /payments/result | Verifies payment result.



## PaymentLinksPost

> CreatePaymentLinkResponse PaymentLinksPost(ctx, optional)

Creates a payment link.

Creates a payment link to our hosted payment form where shoppers can pay. The list of payment methods presented to the shopper depends on the `currency` and `country` parameters sent in the request.  For more information, refer to [Pay by Link documentation](https://docs.adyen.com/checkout/pay-by-link#create-payment-links-through-api).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PaymentLinksPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PaymentLinksPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPaymentLinkRequest** | [**optional.Interface of CreatePaymentLinkRequest**](CreatePaymentLinkRequest.md)|  | 

### Return type

[**CreatePaymentLinkResponse**](CreatePaymentLinkResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentMethodsPost

> PaymentMethodsResponse PaymentMethodsPost(ctx, optional)

Returns available payment methods.

Queries the available payment methods for a transaction based on the transaction context (like amount, country, and currency). Besides giving back a list of the available payment methods, the response also returns which input details you need to collect from the shopper (to be submitted to `/payments`).  Although we highly recommend using this endpoint to ensure you are always offering the most up-to-date list of payment methods, its usage is optional. You can, for example, also cache the `/paymentMethods` response and update it once a week.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PaymentMethodsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PaymentMethodsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentMethodsRequest** | [**optional.Interface of PaymentMethodsRequest**](PaymentMethodsRequest.md)|  | 

### Return type

[**PaymentMethodsResponse**](PaymentMethodsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentSessionPost

> PaymentSetupResponse PaymentSessionPost(ctx, optional)

Creates a payment session.

Provides the data object that can be used to start the Checkout SDK. To set up the payment, pass its amount, currency, and other required parameters. We use this to optimise the payment flow and perform better risk assessment of the transaction.  For more information, refer to [How it works](https://docs.adyen.com/checkout#howitworks).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PaymentSessionPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PaymentSessionPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentSetupRequest** | [**optional.Interface of PaymentSetupRequest**](PaymentSetupRequest.md)|  | 

### Return type

[**PaymentSetupResponse**](PaymentSetupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentsDetailsPost

> PaymentResponse PaymentsDetailsPost(ctx, optional)

Submits details for a payment.

Submits details for a payment created using `/payments`. This step is only needed when no final state has been reached on the `/payments` request (for example for 3D Secure, or when getting redirected back directly from a payment method using an app switch).  The exact details, which need to be sent to this endpoint, are always specified in the response of the associated `/payments` request.  In addition, the endpoint can be used to verify a `payload`, which is returned after coming back from the Checkout SDK or any of the redirect based methods on the Checkout API.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PaymentsDetailsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PaymentsDetailsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **detailsRequest** | [**optional.Interface of DetailsRequest**](DetailsRequest.md)|  | 

### Return type

[**PaymentResponse**](PaymentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentsPost

> PaymentResponse PaymentsPost(ctx, optional)

Starts a transaction.

Sends payment parameters (like amount, country, and currency) together with the input details collected from the shopper. The response returns the result of the payment request: * For some payment methods (e.g. Visa, Mastercard, and SEPA Direct Debits) you'll get a final state in the `resultCode` (e.g. `authorised` or `refused`). * For other payment methods, you'll receive `redirectShopper` as `resultCode` together with a `redirectUrl`. In this case, the shopper must finalize the payment on the page behind the `redirectUrl`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PaymentsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PaymentsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentRequest** | [**optional.Interface of PaymentRequest**](PaymentRequest.md)|  | 

### Return type

[**PaymentResponse**](PaymentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentsResultPost

> PaymentVerificationResponse PaymentsResultPost(ctx, optional)

Verifies payment result.

Verifies the payment result using the payload returned from the Checkout SDK.  For more information, refer to [How it works](https://docs.adyen.com/checkout#howitworks).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PaymentsResultPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PaymentsResultPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentVerificationRequest** | [**optional.Interface of PaymentVerificationRequest**](PaymentVerificationRequest.md)|  | 

### Return type

[**PaymentVerificationResponse**](PaymentVerificationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

