# PaymentRequest3ds2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountInfo** |  Pointer to [**AccountInfo**](AccountInfo.md) |  | [optional] 
**AdditionalAmount** |  Pointer to [**Amount**](Amount.md) |  | [optional] 
**AdditionalData** | **map[string]string** | This field contains additional data, which may be required for a particular payment request.  The &#x60;additionalData&#x60; object consists of entries, each of which includes the key and value. | [optional] 
**Amount** |  Pointer to [**Amount**](Amount.md) |  | 
**ApplicationInfo** |  Pointer to [**ApplicationInfo**](ApplicationInfo.md) |  | [optional] 
**BillingAddress** |  Pointer to [**Address**](Address.md) |  | [optional] 
**BrowserInfo** |  Pointer to [**BrowserInfo**](BrowserInfo.md) |  | [optional] 
**CaptureDelayHours** | **int32** | The delay between the authorisation and scheduled auto-capture, specified in hours. | [optional] 
**DateOfBirth** |  Pointer to [**time.Time**](time.Time.md) | The shopper&#39;s date of birth.  Format [ISO-8601](https://www.w3.org/TR/NOTE-datetime): YYYY-MM-DD | [optional] 
**DccQuote** |  Pointer to [**ForexQuote**](ForexQuote.md) |  | [optional] 
**DeliveryAddress** |  Pointer to [**Address**](Address.md) |  | [optional] 
**DeliveryDate** |  Pointer to [**time.Time**](time.Time.md) | The date and time the purchased goods should be delivered.  Format [ISO 8601](https://www.w3.org/TR/NOTE-datetime): YYYY-MM-DDThh:mm:ss.sssTZD  Example: 2017-07-17T13:42:40.428+01:00 | [optional] 
**DeviceFingerprint** | **string** | A string containing the shopper&#39;s device fingerprint. For more information, refer to [Device fingerprinting](https://docs.adyen.com/risk-management/device-fingerprinting). | [optional] 
**EnableRealTimeUpdate** | **bool** | Choose if a specific transaction should use the Real-time Account Updater, regardless of other settings. | [optional] 
**FraudOffset** | **int32** | An integer value that is added to the normal fraud score. The value can be either positive or negative. | [optional] 
**Installments** |  Pointer to [**Installments**](Installments.md) |  | [optional] 
**Mcc** | **string** | The [merchant category code](https://en.wikipedia.org/wiki/Merchant_category_code) (MCC) is a four-digit number, which relates to a particular market segment. This code reflects the predominant activity that is conducted by the merchant. | [optional] 
**MerchantAccount** | **string** | The merchant account identifier, with which you want to process the transaction. | 
**MerchantOrderReference** | **string** | This reference allows linking multiple transactions to each other for reporting purposes (i.e. order auth-rate). The reference should be unique per billing cycle. The same merchant order reference should never be reused after the first authorised attempt. If used, this field should be supplied for all incoming authorisations. &gt; We strongly recommend you send the &#x60;merchantOrderReference&#x60; value to benefit from linking payment requests when authorisation retries take place. In addition, we recommend you provide &#x60;retry.orderAttemptNumber&#x60;, &#x60;retry.chainAttemptNumber&#x60;, and &#x60;retry.skipRetry&#x60; values in &#x60;PaymentRequest.additionalData&#x60;. | [optional] 
**MerchantRiskIndicator** |  Pointer to [**MerchantRiskIndicator**](MerchantRiskIndicator.md) |  | [optional] 
**Metadata** | **map[string]string** | Metadata consists of entries, each of which includes a key and a value. Limitations: Maximum 20 key-value pairs per request. When exceeding, the \&quot;177\&quot; error occurs: \&quot;Metadata size exceeds limit\&quot;. | [optional] 
**OrderReference** | **string** | When you are doing multiple partial (gift card) payments, this is the &#x60;pspReference&#x60; of the first payment. We use this to link the multiple payments to each other. As your own reference for linking multiple payments, use the &#x60;merchantOrderReference&#x60;instead. | [optional] 
**Recurring** |  Pointer to [**Recurring**](Recurring.md) |  | [optional] 
**RecurringProcessingModel** | **string** | Defines a recurring payment type. Allowed values: * &#x60;Subscription&#x60; – A transaction for a fixed or variable amount, which follows a fixed schedule. * &#x60;CardOnFile&#x60; – Card details are stored to enable one-click or omnichannel journeys, or simply to streamline the checkout process. Any subscription not following a fixed schedule is also considered a card-on-file transaction. * &#x60;UnscheduledCardOnFile&#x60; – A transaction that occurs on a non-fixed schedule and/or have variable amounts. For example, automatic top-ups when a cardholder&#39;s balance drops below a certain amount.  | [optional] 
**Reference** | **string** | The reference to uniquely identify a payment. This reference is used in all communication with you about the payment status. We recommend using a unique value per payment; however, it is not a requirement. If you need to provide multiple references for a transaction, separate them with hyphens (\&quot;-\&quot;). Maximum length: 80 characters. | 
**SelectedBrand** | **string** | Some payment methods require defining a value for this field to specify how to process the transaction.  For the Bancontact payment method, it can be set to: * &#x60;maestro&#x60; (default), to be processed like a Maestro card, or * &#x60;bcmc&#x60;, to be processed like a Bancontact card. | [optional] 
**SelectedRecurringDetailReference** | **string** | The &#x60;recurringDetailReference&#x60; you want to use for this payment. The value &#x60;LATEST&#x60; can be used to select the most recently stored recurring detail. | [optional] 
**SessionId** | **string** | A session ID used to identify a payment session. | [optional] 
**ShopperEmail** | **string** | The shopper&#39;s email address. We recommend that you provide this data, as it is used in velocity fraud checks. &gt; For 3D Secure 2 transactions, schemes require the &#x60;shopperEmail&#x60; for both &#x60;deviceChannel&#x60; **browser** and **app**. | [optional] 
**ShopperIP** | **string** | The shopper&#39;s IP address. In general, we recommend that you provide this data, as it is used in a number of risk checks (for instance, number of payment attempts or location-based checks). &gt; Required for 3D Secure 2 transactions. This field is also mandatory for some merchants depending on your business model. For more information, [contact Support](https://support.adyen.com/hc/en-us/requests/new). | [optional] 
**ShopperInteraction** | **string** | Specifies the sales channel, through which the shopper gives their card details, and whether the shopper is a returning customer. For the web service API, Adyen assumes Ecommerce shopper interaction by default.  This field has the following possible values: * &#x60;Ecommerce&#x60; - Online transactions where the cardholder is present (online). For better authorisation rates, we recommend sending the card security code (CSC) along with the request. * &#x60;ContAuth&#x60; - Card on file and/or subscription transactions, where the cardholder is known to the merchant (returning customer). If the shopper is present (online), you can supply also the CSC to improve authorisation (one-click payment). * &#x60;Moto&#x60; - Mail-order and telephone-order transactions where the shopper is in contact with the merchant via email or telephone. * &#x60;POS&#x60; - Point-of-sale transactions where the shopper is physically present to make a payment using a secure payment terminal. | [optional] 
**ShopperLocale** | **string** | The combination of a language code and a country code to specify the language to be used in the payment. | [optional] 
**ShopperName** |  Pointer to [**Name**](Name.md) |  | [optional] 
**ShopperReference** | **string** | The shopper&#39;s reference to uniquely identify this shopper (e.g. user ID or account ID). &gt; This field is required for recurring payments. | [optional] 
**ShopperStatement** | **string** | The text to appear on the shopper&#39;s bank statement. | [optional] 
**SocialSecurityNumber** | **string** | The shopper&#39;s social security number. | [optional] 
**Splits** |  Pointer to [**[]Split**](Split.md) | Information on how the payment should be split between accounts when using [Adyen for Platforms](https://docs.adyen.com/marketpay/processing-payments#providing-split-information). | [optional] 
**Store** | **string** | The physical store, for which this payment is processed. | [optional] 
**TelephoneNumber** | **string** | The shopper&#39;s telephone number. | [optional] 
**ThreeDS2RequestData** |  Pointer to [**ThreeDS2RequestData**](ThreeDS2RequestData.md) |  | [optional] 
**ThreeDS2Result** |  Pointer to [**ThreeDS2Result**](ThreeDS2Result.md) |  | [optional] 
**ThreeDS2Token** | **string** | The ThreeDS2Token that was returned in the /authorise call. | [optional] 
**ThreeDSAuthenticationOnly** | **bool** | If set to true, you will only perform the [3D Secure 2 authentication](https://docs.adyen.com/checkout/3d-secure/other-3ds-flows/authentication-only), and not the payment authorisation. | [optional] 
**TotalsGroup** | **string** | The reference value to aggregate sales totals in reporting. When not specified, the store field is used (if available). | [optional] 
**TrustedShopper** | **bool** | Set to true if the payment should be routed to a trusted MID. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


