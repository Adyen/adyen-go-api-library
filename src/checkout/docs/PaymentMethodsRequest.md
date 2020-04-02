# PaymentMethodsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | **interface{}** [**AdditionalData | CommonAdditionalData | 3DSecureAdditionalData | AirlineAdditionalData | CarRentalAdditionalData | Level23AdditionalData | LodgingAdditionalData | OpenInvoiceAdditionalData | RatepayAdditionalData | RetryAdditionalData | RiskAdditionalData | RiskStandaloneAdditionalData | TemporaryServices**]() | This field contains additional data, which may be required for a particular payment request.  The &#x60;additionalData&#x60; object consists of entries, each of which includes the key and value. | [optional] 
**AllowedPaymentMethods** | **[]string** | List of payments methods to be presented to the shopper. To refer to payment methods, use their &#x60;paymentMethod.type&#x60; from [Payment methods overview](https://docs.adyen.com/payment-methods).  Example: &#x60;\&quot;allowedPaymentMethods\&quot;:[\&quot;ideal\&quot;,\&quot;giropay\&quot;]&#x60; | [optional] 
**Amount** |  Pointer to [**Amount**](Amount.md) |  | [optional] 
**BlockedPaymentMethods** | **[]string** | List of payments methods to be hidden from the shopper. To refer to payment methods, use their &#x60;paymentMethod.type&#x60; from [Payment methods overview](https://docs.adyen.com/payment-methods).  Example: &#x60;\&quot;blockedPaymentMethods\&quot;:[\&quot;ideal\&quot;,\&quot;giropay\&quot;]&#x60; | [optional] 
**Channel** | **string** | The platform where a payment transaction takes place. This field can be used for filtering out payment methods that are only available on specific platforms. Possible values: * iOS * Android * Web | [optional] 
**CountryCode** | **string** | The shopper&#39;s country code. | [optional] 
**EnableRealTimeUpdate** | **bool** | Choose if a specific transaction should use the Real-time Account Updater, regardless of other settings. | [optional] 
**MerchantAccount** | **string** | The merchant account identifier, with which you want to process the transaction. | 
**ShopperLocale** | **string** | The combination of a language code and a country code to specify the language to be used in the payment. | [optional] 
**ShopperReference** | **string** | The shopper&#39;s reference to uniquely identify this shopper (e.g. user ID or account ID). &gt; This field is required for recurring payments. | [optional] 
**ThreeDSAuthenticationOnly** | **bool** | If set to true, you will only perform the [3D Secure 2 authentication](https://docs.adyen.com/checkout/3d-secure/other-3ds-flows/authentication-only), and not the payment authorisation. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


