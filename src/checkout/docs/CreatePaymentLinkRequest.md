# CreatePaymentLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedPaymentMethods** | **[]string** | List of payments methods to be presented to the shopper. To refer to payment methods, use their &#x60;paymentMethod.type&#x60; from [Payment methods overview](https://docs.adyen.com/payment-methods).  Example: &#x60;\&quot;allowedPaymentMethods\&quot;:[\&quot;ideal\&quot;,\&quot;giropay\&quot;]&#x60; | [optional] 
**Amount** |  Pointer to [**Amount**](Amount.md) |  | 
**BillingAddress** |  Pointer to [**Address**](Address.md) |  | [optional] 
**BlockedPaymentMethods** | **[]string** | List of payments methods to be hidden from the shopper. To refer to payment methods, use their &#x60;paymentMethod.type&#x60; from [Payment methods overview](https://docs.adyen.com/payment-methods).  Example: &#x60;\&quot;blockedPaymentMethods\&quot;:[\&quot;ideal\&quot;,\&quot;giropay\&quot;]&#x60; | [optional] 
**CountryCode** | **string** | The shopper&#39;s country code. | 
**DeliveryAddress** |  Pointer to [**Address**](Address.md) |  | [optional] 
**Description** | **string** | A short description visible on the Pay By Link page. Maximum length: 280 characters. | [optional] 
**ExpiresAt** | **string** | The date that the Pay By Link expires, in ISO 8601 format. For example, 2019-11-23T12:25:28Z. Maximum expiry date should be 30 days from when the payment link is created. If not provided, the default expiry duration is 24 hours. | [optional] 
**MerchantAccount** | **string** | The merchant account identifier, with which you want to process the transaction. | 
**Reference** | **string** | The reference to uniquely identify a payment. This reference is used in all communication with you about the payment status. We recommend using a unique value per payment; however, it is not a requirement. If you need to provide multiple references for a transaction, separate them with hyphens (\&quot;-\&quot;). Maximum length: 80 characters. | 
**ReturnUrl** | **string** | Website URL used for redirection after payment is completed. If provided, a **Continue** button will be shown on the page. If shoppers select the button, they are redirected to the specified URL. | [optional] 
**ShopperEmail** | **string** | The shopper&#39;s email address. We recommend that you provide this data, as it is used in velocity fraud checks. &gt; For 3D Secure 2 transactions, schemes require the &#x60;shopperEmail&#x60; for both &#x60;deviceChannel&#x60; **browser** and **app**. | [optional] 
**ShopperLocale** | **string** | The combination of a language code and a country code to specify the language to be used in the payment. | [optional] 
**ShopperReference** | **string** | The shopper&#39;s reference to uniquely identify this shopper (e.g. user ID or account ID). &gt; This field is required for recurring payments. | [optional] 
**StorePaymentMethod** | **bool** | When true and &#x60;shopperReference&#x60; is provided, the payment details will be stored. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


