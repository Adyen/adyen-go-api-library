# CheckoutBalanceCheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountInfo** | Pointer to [**AccountInfo**](AccountInfo.md) |  | [optional] 
**AdditionalAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**AdditionalData** | Pointer to **map[string]string** | This field contains additional data, which may be required for a particular payment request.  The &#x60;additionalData&#x60; object consists of entries, each of which includes the key and value. | [optional] 
**Amount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**ApplicationInfo** | Pointer to [**ApplicationInfo**](ApplicationInfo.md) |  | [optional] 
**BillingAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**BrowserInfo** | Pointer to [**BrowserInfo**](BrowserInfo.md) |  | [optional] 
**CaptureDelayHours** | Pointer to **int32** | The delay between the authorisation and scheduled auto-capture, specified in hours. | [optional] 
**DateOfBirth** | Pointer to **string** | The shopper&#39;s date of birth.  Format [ISO-8601](https://www.w3.org/TR/NOTE-datetime): YYYY-MM-DD | [optional] 
**DccQuote** | Pointer to [**ForexQuote**](ForexQuote.md) |  | [optional] 
**DeliveryAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**DeliveryDate** | Pointer to **time.Time** | The date and time the purchased goods should be delivered.  Format [ISO 8601](https://www.w3.org/TR/NOTE-datetime): YYYY-MM-DDThh:mm:ss.sssTZD  Example: 2017-07-17T13:42:40.428+01:00 | [optional] 
**DeviceFingerprint** | Pointer to **string** | A string containing the shopper&#39;s device fingerprint. For more information, refer to [Device fingerprinting](https://docs.adyen.com/risk-management/device-fingerprinting). | [optional] 
**FraudOffset** | Pointer to **int32** | An integer value that is added to the normal fraud score. The value can be either positive or negative. | [optional] 
**Installments** | Pointer to [**Installments**](Installments.md) |  | [optional] 
**LocalizedShopperStatement** | Pointer to **map[string]string** | This field allows merchants to use dynamic shopper statement in local character sets. The local shopper statement field can be supplied in markets where localized merchant descriptors are used. Currently, Adyen only supports this in the Japanese market .The available character sets at the moment are: * Processing in Japan: **ja-Kana** The character set **ja-Kana** supports UTF-8 based Katakana and alphanumeric and special characters. Merchants should send the Katakana shopperStatement in full-width characters.  An example request would be: &gt; {   \&quot;shopperStatement\&quot; : \&quot;ADYEN - SELLER-A\&quot;,   \&quot;localizedShopperStatement\&quot; : {     \&quot;ja-Kana\&quot; : \&quot;ADYEN - セラーA\&quot;   } } We recommend merchants to always supply the field localizedShopperStatement in addition to the field shopperStatement.It is issuer dependent whether the localized shopper statement field is supported. In the case of non-domestic transactions (e.g. US-issued cards processed in JP) the field &#x60;shopperStatement&#x60; is used to modify the statement of the shopper. Adyen handles the complexity of ensuring the correct descriptors are assigned. | [optional] 
**Mcc** | Pointer to **string** | The [merchant category code](https://en.wikipedia.org/wiki/Merchant_category_code) (MCC) is a four-digit number, which relates to a particular market segment. This code reflects the predominant activity that is conducted by the merchant. | [optional] 
**MerchantAccount** | **string** | The merchant account identifier, with which you want to process the transaction. | 
**MerchantOrderReference** | Pointer to **string** | This reference allows linking multiple transactions to each other for reporting purposes (i.e. order auth-rate). The reference should be unique per billing cycle. The same merchant order reference should never be reused after the first authorised attempt. If used, this field should be supplied for all incoming authorisations. &gt; We strongly recommend you send the &#x60;merchantOrderReference&#x60; value to benefit from linking payment requests when authorisation retries take place. In addition, we recommend you provide &#x60;retry.orderAttemptNumber&#x60;, &#x60;retry.chainAttemptNumber&#x60;, and &#x60;retry.skipRetry&#x60; values in &#x60;PaymentRequest.additionalData&#x60;. | [optional] 
**MerchantRiskIndicator** | Pointer to [**MerchantRiskIndicator**](MerchantRiskIndicator.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** | Metadata consists of entries, each of which includes a key and a value. Limits: * Maximum 20 key-value pairs per request. When exceeding, the \&quot;177\&quot; error occurs: \&quot;Metadata size exceeds limit\&quot;. * Maximum 20 characters per key. * Maximum 80 characters per value.  | [optional] 
**OrderReference** | Pointer to **string** | When you are doing multiple partial (gift card) payments, this is the &#x60;pspReference&#x60; of the first payment. We use this to link the multiple payments to each other. As your own reference for linking multiple payments, use the &#x60;merchantOrderReference&#x60;instead. | [optional] 
**PaymentMethod** | **map[string]string** | The collection that contains the type of the payment method and its specific information. | 
**Recurring** | Pointer to [**Recurring**](Recurring.md) |  | [optional] 
**RecurringProcessingModel** | Pointer to **string** | Defines a recurring payment type. Allowed values: * &#x60;Subscription&#x60; – A transaction for a fixed or variable amount, which follows a fixed schedule. * &#x60;CardOnFile&#x60; – With a card-on-file (CoF) transaction, card details are stored to enable one-click or omnichannel journeys, or simply to streamline the checkout process. Any subscription not following a fixed schedule is also considered a card-on-file transaction. * &#x60;UnscheduledCardOnFile&#x60; – An unscheduled card-on-file (UCoF) transaction is a transaction that occurs on a non-fixed schedule and/or have variable amounts. For example, automatic top-ups when a cardholder&#39;s balance drops below a certain amount.  | [optional] 
**Reference** | Pointer to **string** | The reference to uniquely identify a payment. This reference is used in all communication with you about the payment status. We recommend using a unique value per payment; however, it is not a requirement. If you need to provide multiple references for a transaction, separate them with hyphens (\&quot;-\&quot;). Maximum length: 80 characters. | [optional] 
**SelectedBrand** | Pointer to **string** | Some payment methods require defining a value for this field to specify how to process the transaction.  For the Bancontact payment method, it can be set to: * &#x60;maestro&#x60; (default), to be processed like a Maestro card, or * &#x60;bcmc&#x60;, to be processed like a Bancontact card. | [optional] 
**SelectedRecurringDetailReference** | Pointer to **string** | The &#x60;recurringDetailReference&#x60; you want to use for this payment. The value &#x60;LATEST&#x60; can be used to select the most recently stored recurring detail. | [optional] 
**SessionId** | Pointer to **string** | A session ID used to identify a payment session. | [optional] 
**ShopperEmail** | Pointer to **string** | The shopper&#39;s email address. We recommend that you provide this data, as it is used in velocity fraud checks. &gt; For 3D Secure 2 transactions, schemes require &#x60;shopperEmail&#x60; for all browser-based and mobile implementations. | [optional] 
**ShopperIP** | Pointer to **string** | The shopper&#39;s IP address. In general, we recommend that you provide this data, as it is used in a number of risk checks (for instance, number of payment attempts or location-based checks). &gt; For 3D Secure 2 transactions, schemes require &#x60;shopperIP&#x60; for all browser-based implementations. This field is also mandatory for some merchants depending on your business model. For more information, [contact Support](https://www.adyen.help/hc/en-us/requests/new). | [optional] 
**ShopperInteraction** | Pointer to **string** | Specifies the sales channel, through which the shopper gives their card details, and whether the shopper is a returning customer. For the web service API, Adyen assumes Ecommerce shopper interaction by default.  This field has the following possible values: * &#x60;Ecommerce&#x60; - Online transactions where the cardholder is present (online). For better authorisation rates, we recommend sending the card security code (CSC) along with the request. * &#x60;ContAuth&#x60; - Card on file and/or subscription transactions, where the cardholder is known to the merchant (returning customer). If the shopper is present (online), you can supply also the CSC to improve authorisation (one-click payment). * &#x60;Moto&#x60; - Mail-order and telephone-order transactions where the shopper is in contact with the merchant via email or telephone. * &#x60;POS&#x60; - Point-of-sale transactions where the shopper is physically present to make a payment using a secure payment terminal. | [optional] 
**ShopperLocale** | Pointer to **string** | The combination of a language code and a country code to specify the language to be used in the payment. | [optional] 
**ShopperName** | Pointer to [**Name**](Name.md) |  | [optional] 
**ShopperReference** | Pointer to **string** | Required for recurring payments.  Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. &gt; Your reference must not include personally identifiable information (PII), for example name or email address. | [optional] 
**ShopperStatement** | Pointer to **string** | The text to be shown on the shopper&#39;s bank statement.  We recommend sending a maximum of 22 characters, otherwise banks might truncate the string.  Allowed characters: **a-z**, **A-Z**, **0-9**, spaces, and special characters **. , &#39; _ - ? + * /_**. | [optional] 
**SocialSecurityNumber** | Pointer to **string** | The shopper&#39;s social security number. | [optional] 
**Splits** | Pointer to [**[]Split**](Split.md) | An array of objects specifying how the payment should be split when using [Adyen for Platforms](https://docs.adyen.com/marketplaces-and-platforms/processing-payments#providing-split-information) or [Issuing](https://docs.adyen.com/issuing/add-manage-funds#split). | [optional] 
**Store** | Pointer to **string** | The ecommerce or point-of-sale store that is processing the payment. Used in [partner model integrations](https://docs.adyen.com/marketplaces-and-platforms/classic/platforms-for-partners#route-payments) for Adyen for Platforms. | [optional] 
**TelephoneNumber** | Pointer to **string** | The shopper&#39;s telephone number. | [optional] 
**ThreeDS2RequestData** | Pointer to [**ThreeDS2RequestData**](ThreeDS2RequestData.md) |  | [optional] 
**ThreeDSAuthenticationOnly** | Pointer to **bool** | If set to true, you will only perform the [3D Secure 2 authentication](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only), and not the payment authorisation. | [optional] [default to false]
**TotalsGroup** | Pointer to **string** | The reference value to aggregate sales totals in reporting. When not specified, the store field is used (if available). | [optional] 
**TrustedShopper** | Pointer to **bool** | Set to true if the payment should be routed to a trusted MID. | [optional] 

## Methods

### NewCheckoutBalanceCheckRequest

`func NewCheckoutBalanceCheckRequest(merchantAccount string, paymentMethod map[string]string, ) *CheckoutBalanceCheckRequest`

NewCheckoutBalanceCheckRequest instantiates a new CheckoutBalanceCheckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutBalanceCheckRequestWithDefaults

`func NewCheckoutBalanceCheckRequestWithDefaults() *CheckoutBalanceCheckRequest`

NewCheckoutBalanceCheckRequestWithDefaults instantiates a new CheckoutBalanceCheckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountInfo

`func (o *CheckoutBalanceCheckRequest) GetAccountInfo() AccountInfo`

GetAccountInfo returns the AccountInfo field if non-nil, zero value otherwise.

### GetAccountInfoOk

`func (o *CheckoutBalanceCheckRequest) GetAccountInfoOk() (*AccountInfo, bool)`

GetAccountInfoOk returns a tuple with the AccountInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountInfo

`func (o *CheckoutBalanceCheckRequest) SetAccountInfo(v AccountInfo)`

SetAccountInfo sets AccountInfo field to given value.

### HasAccountInfo

`func (o *CheckoutBalanceCheckRequest) HasAccountInfo() bool`

HasAccountInfo returns a boolean if a field has been set.

### GetAdditionalAmount

`func (o *CheckoutBalanceCheckRequest) GetAdditionalAmount() Amount`

GetAdditionalAmount returns the AdditionalAmount field if non-nil, zero value otherwise.

### GetAdditionalAmountOk

`func (o *CheckoutBalanceCheckRequest) GetAdditionalAmountOk() (*Amount, bool)`

GetAdditionalAmountOk returns a tuple with the AdditionalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAmount

`func (o *CheckoutBalanceCheckRequest) SetAdditionalAmount(v Amount)`

SetAdditionalAmount sets AdditionalAmount field to given value.

### HasAdditionalAmount

`func (o *CheckoutBalanceCheckRequest) HasAdditionalAmount() bool`

HasAdditionalAmount returns a boolean if a field has been set.

### GetAdditionalData

`func (o *CheckoutBalanceCheckRequest) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *CheckoutBalanceCheckRequest) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *CheckoutBalanceCheckRequest) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *CheckoutBalanceCheckRequest) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetAmount

`func (o *CheckoutBalanceCheckRequest) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CheckoutBalanceCheckRequest) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CheckoutBalanceCheckRequest) SetAmount(v Amount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CheckoutBalanceCheckRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetApplicationInfo

`func (o *CheckoutBalanceCheckRequest) GetApplicationInfo() ApplicationInfo`

GetApplicationInfo returns the ApplicationInfo field if non-nil, zero value otherwise.

### GetApplicationInfoOk

`func (o *CheckoutBalanceCheckRequest) GetApplicationInfoOk() (*ApplicationInfo, bool)`

GetApplicationInfoOk returns a tuple with the ApplicationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationInfo

`func (o *CheckoutBalanceCheckRequest) SetApplicationInfo(v ApplicationInfo)`

SetApplicationInfo sets ApplicationInfo field to given value.

### HasApplicationInfo

`func (o *CheckoutBalanceCheckRequest) HasApplicationInfo() bool`

HasApplicationInfo returns a boolean if a field has been set.

### GetBillingAddress

`func (o *CheckoutBalanceCheckRequest) GetBillingAddress() Address`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *CheckoutBalanceCheckRequest) GetBillingAddressOk() (*Address, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *CheckoutBalanceCheckRequest) SetBillingAddress(v Address)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *CheckoutBalanceCheckRequest) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetBrowserInfo

`func (o *CheckoutBalanceCheckRequest) GetBrowserInfo() BrowserInfo`

GetBrowserInfo returns the BrowserInfo field if non-nil, zero value otherwise.

### GetBrowserInfoOk

`func (o *CheckoutBalanceCheckRequest) GetBrowserInfoOk() (*BrowserInfo, bool)`

GetBrowserInfoOk returns a tuple with the BrowserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserInfo

`func (o *CheckoutBalanceCheckRequest) SetBrowserInfo(v BrowserInfo)`

SetBrowserInfo sets BrowserInfo field to given value.

### HasBrowserInfo

`func (o *CheckoutBalanceCheckRequest) HasBrowserInfo() bool`

HasBrowserInfo returns a boolean if a field has been set.

### GetCaptureDelayHours

`func (o *CheckoutBalanceCheckRequest) GetCaptureDelayHours() int32`

GetCaptureDelayHours returns the CaptureDelayHours field if non-nil, zero value otherwise.

### GetCaptureDelayHoursOk

`func (o *CheckoutBalanceCheckRequest) GetCaptureDelayHoursOk() (*int32, bool)`

GetCaptureDelayHoursOk returns a tuple with the CaptureDelayHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureDelayHours

`func (o *CheckoutBalanceCheckRequest) SetCaptureDelayHours(v int32)`

SetCaptureDelayHours sets CaptureDelayHours field to given value.

### HasCaptureDelayHours

`func (o *CheckoutBalanceCheckRequest) HasCaptureDelayHours() bool`

HasCaptureDelayHours returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *CheckoutBalanceCheckRequest) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *CheckoutBalanceCheckRequest) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *CheckoutBalanceCheckRequest) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *CheckoutBalanceCheckRequest) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetDccQuote

`func (o *CheckoutBalanceCheckRequest) GetDccQuote() ForexQuote`

GetDccQuote returns the DccQuote field if non-nil, zero value otherwise.

### GetDccQuoteOk

`func (o *CheckoutBalanceCheckRequest) GetDccQuoteOk() (*ForexQuote, bool)`

GetDccQuoteOk returns a tuple with the DccQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDccQuote

`func (o *CheckoutBalanceCheckRequest) SetDccQuote(v ForexQuote)`

SetDccQuote sets DccQuote field to given value.

### HasDccQuote

`func (o *CheckoutBalanceCheckRequest) HasDccQuote() bool`

HasDccQuote returns a boolean if a field has been set.

### GetDeliveryAddress

`func (o *CheckoutBalanceCheckRequest) GetDeliveryAddress() Address`

GetDeliveryAddress returns the DeliveryAddress field if non-nil, zero value otherwise.

### GetDeliveryAddressOk

`func (o *CheckoutBalanceCheckRequest) GetDeliveryAddressOk() (*Address, bool)`

GetDeliveryAddressOk returns a tuple with the DeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddress

`func (o *CheckoutBalanceCheckRequest) SetDeliveryAddress(v Address)`

SetDeliveryAddress sets DeliveryAddress field to given value.

### HasDeliveryAddress

`func (o *CheckoutBalanceCheckRequest) HasDeliveryAddress() bool`

HasDeliveryAddress returns a boolean if a field has been set.

### GetDeliveryDate

`func (o *CheckoutBalanceCheckRequest) GetDeliveryDate() time.Time`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *CheckoutBalanceCheckRequest) GetDeliveryDateOk() (*time.Time, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *CheckoutBalanceCheckRequest) SetDeliveryDate(v time.Time)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *CheckoutBalanceCheckRequest) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### GetDeviceFingerprint

`func (o *CheckoutBalanceCheckRequest) GetDeviceFingerprint() string`

GetDeviceFingerprint returns the DeviceFingerprint field if non-nil, zero value otherwise.

### GetDeviceFingerprintOk

`func (o *CheckoutBalanceCheckRequest) GetDeviceFingerprintOk() (*string, bool)`

GetDeviceFingerprintOk returns a tuple with the DeviceFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFingerprint

`func (o *CheckoutBalanceCheckRequest) SetDeviceFingerprint(v string)`

SetDeviceFingerprint sets DeviceFingerprint field to given value.

### HasDeviceFingerprint

`func (o *CheckoutBalanceCheckRequest) HasDeviceFingerprint() bool`

HasDeviceFingerprint returns a boolean if a field has been set.

### GetFraudOffset

`func (o *CheckoutBalanceCheckRequest) GetFraudOffset() int32`

GetFraudOffset returns the FraudOffset field if non-nil, zero value otherwise.

### GetFraudOffsetOk

`func (o *CheckoutBalanceCheckRequest) GetFraudOffsetOk() (*int32, bool)`

GetFraudOffsetOk returns a tuple with the FraudOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudOffset

`func (o *CheckoutBalanceCheckRequest) SetFraudOffset(v int32)`

SetFraudOffset sets FraudOffset field to given value.

### HasFraudOffset

`func (o *CheckoutBalanceCheckRequest) HasFraudOffset() bool`

HasFraudOffset returns a boolean if a field has been set.

### GetInstallments

`func (o *CheckoutBalanceCheckRequest) GetInstallments() Installments`

GetInstallments returns the Installments field if non-nil, zero value otherwise.

### GetInstallmentsOk

`func (o *CheckoutBalanceCheckRequest) GetInstallmentsOk() (*Installments, bool)`

GetInstallmentsOk returns a tuple with the Installments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallments

`func (o *CheckoutBalanceCheckRequest) SetInstallments(v Installments)`

SetInstallments sets Installments field to given value.

### HasInstallments

`func (o *CheckoutBalanceCheckRequest) HasInstallments() bool`

HasInstallments returns a boolean if a field has been set.

### GetLocalizedShopperStatement

`func (o *CheckoutBalanceCheckRequest) GetLocalizedShopperStatement() map[string]string`

GetLocalizedShopperStatement returns the LocalizedShopperStatement field if non-nil, zero value otherwise.

### GetLocalizedShopperStatementOk

`func (o *CheckoutBalanceCheckRequest) GetLocalizedShopperStatementOk() (*map[string]string, bool)`

GetLocalizedShopperStatementOk returns a tuple with the LocalizedShopperStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedShopperStatement

`func (o *CheckoutBalanceCheckRequest) SetLocalizedShopperStatement(v map[string]string)`

SetLocalizedShopperStatement sets LocalizedShopperStatement field to given value.

### HasLocalizedShopperStatement

`func (o *CheckoutBalanceCheckRequest) HasLocalizedShopperStatement() bool`

HasLocalizedShopperStatement returns a boolean if a field has been set.

### GetMcc

`func (o *CheckoutBalanceCheckRequest) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *CheckoutBalanceCheckRequest) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *CheckoutBalanceCheckRequest) SetMcc(v string)`

SetMcc sets Mcc field to given value.

### HasMcc

`func (o *CheckoutBalanceCheckRequest) HasMcc() bool`

HasMcc returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *CheckoutBalanceCheckRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *CheckoutBalanceCheckRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *CheckoutBalanceCheckRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetMerchantOrderReference

`func (o *CheckoutBalanceCheckRequest) GetMerchantOrderReference() string`

GetMerchantOrderReference returns the MerchantOrderReference field if non-nil, zero value otherwise.

### GetMerchantOrderReferenceOk

`func (o *CheckoutBalanceCheckRequest) GetMerchantOrderReferenceOk() (*string, bool)`

GetMerchantOrderReferenceOk returns a tuple with the MerchantOrderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderReference

`func (o *CheckoutBalanceCheckRequest) SetMerchantOrderReference(v string)`

SetMerchantOrderReference sets MerchantOrderReference field to given value.

### HasMerchantOrderReference

`func (o *CheckoutBalanceCheckRequest) HasMerchantOrderReference() bool`

HasMerchantOrderReference returns a boolean if a field has been set.

### GetMerchantRiskIndicator

`func (o *CheckoutBalanceCheckRequest) GetMerchantRiskIndicator() MerchantRiskIndicator`

GetMerchantRiskIndicator returns the MerchantRiskIndicator field if non-nil, zero value otherwise.

### GetMerchantRiskIndicatorOk

`func (o *CheckoutBalanceCheckRequest) GetMerchantRiskIndicatorOk() (*MerchantRiskIndicator, bool)`

GetMerchantRiskIndicatorOk returns a tuple with the MerchantRiskIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantRiskIndicator

`func (o *CheckoutBalanceCheckRequest) SetMerchantRiskIndicator(v MerchantRiskIndicator)`

SetMerchantRiskIndicator sets MerchantRiskIndicator field to given value.

### HasMerchantRiskIndicator

`func (o *CheckoutBalanceCheckRequest) HasMerchantRiskIndicator() bool`

HasMerchantRiskIndicator returns a boolean if a field has been set.

### GetMetadata

`func (o *CheckoutBalanceCheckRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CheckoutBalanceCheckRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CheckoutBalanceCheckRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CheckoutBalanceCheckRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOrderReference

`func (o *CheckoutBalanceCheckRequest) GetOrderReference() string`

GetOrderReference returns the OrderReference field if non-nil, zero value otherwise.

### GetOrderReferenceOk

`func (o *CheckoutBalanceCheckRequest) GetOrderReferenceOk() (*string, bool)`

GetOrderReferenceOk returns a tuple with the OrderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderReference

`func (o *CheckoutBalanceCheckRequest) SetOrderReference(v string)`

SetOrderReference sets OrderReference field to given value.

### HasOrderReference

`func (o *CheckoutBalanceCheckRequest) HasOrderReference() bool`

HasOrderReference returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *CheckoutBalanceCheckRequest) GetPaymentMethod() map[string]string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *CheckoutBalanceCheckRequest) GetPaymentMethodOk() (*map[string]string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *CheckoutBalanceCheckRequest) SetPaymentMethod(v map[string]string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetRecurring

`func (o *CheckoutBalanceCheckRequest) GetRecurring() Recurring`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *CheckoutBalanceCheckRequest) GetRecurringOk() (*Recurring, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *CheckoutBalanceCheckRequest) SetRecurring(v Recurring)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *CheckoutBalanceCheckRequest) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.

### GetRecurringProcessingModel

`func (o *CheckoutBalanceCheckRequest) GetRecurringProcessingModel() string`

GetRecurringProcessingModel returns the RecurringProcessingModel field if non-nil, zero value otherwise.

### GetRecurringProcessingModelOk

`func (o *CheckoutBalanceCheckRequest) GetRecurringProcessingModelOk() (*string, bool)`

GetRecurringProcessingModelOk returns a tuple with the RecurringProcessingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringProcessingModel

`func (o *CheckoutBalanceCheckRequest) SetRecurringProcessingModel(v string)`

SetRecurringProcessingModel sets RecurringProcessingModel field to given value.

### HasRecurringProcessingModel

`func (o *CheckoutBalanceCheckRequest) HasRecurringProcessingModel() bool`

HasRecurringProcessingModel returns a boolean if a field has been set.

### GetReference

`func (o *CheckoutBalanceCheckRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CheckoutBalanceCheckRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CheckoutBalanceCheckRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *CheckoutBalanceCheckRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSelectedBrand

`func (o *CheckoutBalanceCheckRequest) GetSelectedBrand() string`

GetSelectedBrand returns the SelectedBrand field if non-nil, zero value otherwise.

### GetSelectedBrandOk

`func (o *CheckoutBalanceCheckRequest) GetSelectedBrandOk() (*string, bool)`

GetSelectedBrandOk returns a tuple with the SelectedBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedBrand

`func (o *CheckoutBalanceCheckRequest) SetSelectedBrand(v string)`

SetSelectedBrand sets SelectedBrand field to given value.

### HasSelectedBrand

`func (o *CheckoutBalanceCheckRequest) HasSelectedBrand() bool`

HasSelectedBrand returns a boolean if a field has been set.

### GetSelectedRecurringDetailReference

`func (o *CheckoutBalanceCheckRequest) GetSelectedRecurringDetailReference() string`

GetSelectedRecurringDetailReference returns the SelectedRecurringDetailReference field if non-nil, zero value otherwise.

### GetSelectedRecurringDetailReferenceOk

`func (o *CheckoutBalanceCheckRequest) GetSelectedRecurringDetailReferenceOk() (*string, bool)`

GetSelectedRecurringDetailReferenceOk returns a tuple with the SelectedRecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRecurringDetailReference

`func (o *CheckoutBalanceCheckRequest) SetSelectedRecurringDetailReference(v string)`

SetSelectedRecurringDetailReference sets SelectedRecurringDetailReference field to given value.

### HasSelectedRecurringDetailReference

`func (o *CheckoutBalanceCheckRequest) HasSelectedRecurringDetailReference() bool`

HasSelectedRecurringDetailReference returns a boolean if a field has been set.

### GetSessionId

`func (o *CheckoutBalanceCheckRequest) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CheckoutBalanceCheckRequest) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *CheckoutBalanceCheckRequest) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *CheckoutBalanceCheckRequest) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetShopperEmail

`func (o *CheckoutBalanceCheckRequest) GetShopperEmail() string`

GetShopperEmail returns the ShopperEmail field if non-nil, zero value otherwise.

### GetShopperEmailOk

`func (o *CheckoutBalanceCheckRequest) GetShopperEmailOk() (*string, bool)`

GetShopperEmailOk returns a tuple with the ShopperEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperEmail

`func (o *CheckoutBalanceCheckRequest) SetShopperEmail(v string)`

SetShopperEmail sets ShopperEmail field to given value.

### HasShopperEmail

`func (o *CheckoutBalanceCheckRequest) HasShopperEmail() bool`

HasShopperEmail returns a boolean if a field has been set.

### GetShopperIP

`func (o *CheckoutBalanceCheckRequest) GetShopperIP() string`

GetShopperIP returns the ShopperIP field if non-nil, zero value otherwise.

### GetShopperIPOk

`func (o *CheckoutBalanceCheckRequest) GetShopperIPOk() (*string, bool)`

GetShopperIPOk returns a tuple with the ShopperIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperIP

`func (o *CheckoutBalanceCheckRequest) SetShopperIP(v string)`

SetShopperIP sets ShopperIP field to given value.

### HasShopperIP

`func (o *CheckoutBalanceCheckRequest) HasShopperIP() bool`

HasShopperIP returns a boolean if a field has been set.

### GetShopperInteraction

`func (o *CheckoutBalanceCheckRequest) GetShopperInteraction() string`

GetShopperInteraction returns the ShopperInteraction field if non-nil, zero value otherwise.

### GetShopperInteractionOk

`func (o *CheckoutBalanceCheckRequest) GetShopperInteractionOk() (*string, bool)`

GetShopperInteractionOk returns a tuple with the ShopperInteraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperInteraction

`func (o *CheckoutBalanceCheckRequest) SetShopperInteraction(v string)`

SetShopperInteraction sets ShopperInteraction field to given value.

### HasShopperInteraction

`func (o *CheckoutBalanceCheckRequest) HasShopperInteraction() bool`

HasShopperInteraction returns a boolean if a field has been set.

### GetShopperLocale

`func (o *CheckoutBalanceCheckRequest) GetShopperLocale() string`

GetShopperLocale returns the ShopperLocale field if non-nil, zero value otherwise.

### GetShopperLocaleOk

`func (o *CheckoutBalanceCheckRequest) GetShopperLocaleOk() (*string, bool)`

GetShopperLocaleOk returns a tuple with the ShopperLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperLocale

`func (o *CheckoutBalanceCheckRequest) SetShopperLocale(v string)`

SetShopperLocale sets ShopperLocale field to given value.

### HasShopperLocale

`func (o *CheckoutBalanceCheckRequest) HasShopperLocale() bool`

HasShopperLocale returns a boolean if a field has been set.

### GetShopperName

`func (o *CheckoutBalanceCheckRequest) GetShopperName() Name`

GetShopperName returns the ShopperName field if non-nil, zero value otherwise.

### GetShopperNameOk

`func (o *CheckoutBalanceCheckRequest) GetShopperNameOk() (*Name, bool)`

GetShopperNameOk returns a tuple with the ShopperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperName

`func (o *CheckoutBalanceCheckRequest) SetShopperName(v Name)`

SetShopperName sets ShopperName field to given value.

### HasShopperName

`func (o *CheckoutBalanceCheckRequest) HasShopperName() bool`

HasShopperName returns a boolean if a field has been set.

### GetShopperReference

`func (o *CheckoutBalanceCheckRequest) GetShopperReference() string`

GetShopperReference returns the ShopperReference field if non-nil, zero value otherwise.

### GetShopperReferenceOk

`func (o *CheckoutBalanceCheckRequest) GetShopperReferenceOk() (*string, bool)`

GetShopperReferenceOk returns a tuple with the ShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperReference

`func (o *CheckoutBalanceCheckRequest) SetShopperReference(v string)`

SetShopperReference sets ShopperReference field to given value.

### HasShopperReference

`func (o *CheckoutBalanceCheckRequest) HasShopperReference() bool`

HasShopperReference returns a boolean if a field has been set.

### GetShopperStatement

`func (o *CheckoutBalanceCheckRequest) GetShopperStatement() string`

GetShopperStatement returns the ShopperStatement field if non-nil, zero value otherwise.

### GetShopperStatementOk

`func (o *CheckoutBalanceCheckRequest) GetShopperStatementOk() (*string, bool)`

GetShopperStatementOk returns a tuple with the ShopperStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperStatement

`func (o *CheckoutBalanceCheckRequest) SetShopperStatement(v string)`

SetShopperStatement sets ShopperStatement field to given value.

### HasShopperStatement

`func (o *CheckoutBalanceCheckRequest) HasShopperStatement() bool`

HasShopperStatement returns a boolean if a field has been set.

### GetSocialSecurityNumber

`func (o *CheckoutBalanceCheckRequest) GetSocialSecurityNumber() string`

GetSocialSecurityNumber returns the SocialSecurityNumber field if non-nil, zero value otherwise.

### GetSocialSecurityNumberOk

`func (o *CheckoutBalanceCheckRequest) GetSocialSecurityNumberOk() (*string, bool)`

GetSocialSecurityNumberOk returns a tuple with the SocialSecurityNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialSecurityNumber

`func (o *CheckoutBalanceCheckRequest) SetSocialSecurityNumber(v string)`

SetSocialSecurityNumber sets SocialSecurityNumber field to given value.

### HasSocialSecurityNumber

`func (o *CheckoutBalanceCheckRequest) HasSocialSecurityNumber() bool`

HasSocialSecurityNumber returns a boolean if a field has been set.

### GetSplits

`func (o *CheckoutBalanceCheckRequest) GetSplits() []Split`

GetSplits returns the Splits field if non-nil, zero value otherwise.

### GetSplitsOk

`func (o *CheckoutBalanceCheckRequest) GetSplitsOk() (*[]Split, bool)`

GetSplitsOk returns a tuple with the Splits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplits

`func (o *CheckoutBalanceCheckRequest) SetSplits(v []Split)`

SetSplits sets Splits field to given value.

### HasSplits

`func (o *CheckoutBalanceCheckRequest) HasSplits() bool`

HasSplits returns a boolean if a field has been set.

### GetStore

`func (o *CheckoutBalanceCheckRequest) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *CheckoutBalanceCheckRequest) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *CheckoutBalanceCheckRequest) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *CheckoutBalanceCheckRequest) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetTelephoneNumber

`func (o *CheckoutBalanceCheckRequest) GetTelephoneNumber() string`

GetTelephoneNumber returns the TelephoneNumber field if non-nil, zero value otherwise.

### GetTelephoneNumberOk

`func (o *CheckoutBalanceCheckRequest) GetTelephoneNumberOk() (*string, bool)`

GetTelephoneNumberOk returns a tuple with the TelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumber

`func (o *CheckoutBalanceCheckRequest) SetTelephoneNumber(v string)`

SetTelephoneNumber sets TelephoneNumber field to given value.

### HasTelephoneNumber

`func (o *CheckoutBalanceCheckRequest) HasTelephoneNumber() bool`

HasTelephoneNumber returns a boolean if a field has been set.

### GetThreeDS2RequestData

`func (o *CheckoutBalanceCheckRequest) GetThreeDS2RequestData() ThreeDS2RequestData`

GetThreeDS2RequestData returns the ThreeDS2RequestData field if non-nil, zero value otherwise.

### GetThreeDS2RequestDataOk

`func (o *CheckoutBalanceCheckRequest) GetThreeDS2RequestDataOk() (*ThreeDS2RequestData, bool)`

GetThreeDS2RequestDataOk returns a tuple with the ThreeDS2RequestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDS2RequestData

`func (o *CheckoutBalanceCheckRequest) SetThreeDS2RequestData(v ThreeDS2RequestData)`

SetThreeDS2RequestData sets ThreeDS2RequestData field to given value.

### HasThreeDS2RequestData

`func (o *CheckoutBalanceCheckRequest) HasThreeDS2RequestData() bool`

HasThreeDS2RequestData returns a boolean if a field has been set.

### GetThreeDSAuthenticationOnly

`func (o *CheckoutBalanceCheckRequest) GetThreeDSAuthenticationOnly() bool`

GetThreeDSAuthenticationOnly returns the ThreeDSAuthenticationOnly field if non-nil, zero value otherwise.

### GetThreeDSAuthenticationOnlyOk

`func (o *CheckoutBalanceCheckRequest) GetThreeDSAuthenticationOnlyOk() (*bool, bool)`

GetThreeDSAuthenticationOnlyOk returns a tuple with the ThreeDSAuthenticationOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSAuthenticationOnly

`func (o *CheckoutBalanceCheckRequest) SetThreeDSAuthenticationOnly(v bool)`

SetThreeDSAuthenticationOnly sets ThreeDSAuthenticationOnly field to given value.

### HasThreeDSAuthenticationOnly

`func (o *CheckoutBalanceCheckRequest) HasThreeDSAuthenticationOnly() bool`

HasThreeDSAuthenticationOnly returns a boolean if a field has been set.

### GetTotalsGroup

`func (o *CheckoutBalanceCheckRequest) GetTotalsGroup() string`

GetTotalsGroup returns the TotalsGroup field if non-nil, zero value otherwise.

### GetTotalsGroupOk

`func (o *CheckoutBalanceCheckRequest) GetTotalsGroupOk() (*string, bool)`

GetTotalsGroupOk returns a tuple with the TotalsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalsGroup

`func (o *CheckoutBalanceCheckRequest) SetTotalsGroup(v string)`

SetTotalsGroup sets TotalsGroup field to given value.

### HasTotalsGroup

`func (o *CheckoutBalanceCheckRequest) HasTotalsGroup() bool`

HasTotalsGroup returns a boolean if a field has been set.

### GetTrustedShopper

`func (o *CheckoutBalanceCheckRequest) GetTrustedShopper() bool`

GetTrustedShopper returns the TrustedShopper field if non-nil, zero value otherwise.

### GetTrustedShopperOk

`func (o *CheckoutBalanceCheckRequest) GetTrustedShopperOk() (*bool, bool)`

GetTrustedShopperOk returns a tuple with the TrustedShopper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedShopper

`func (o *CheckoutBalanceCheckRequest) SetTrustedShopper(v bool)`

SetTrustedShopper sets TrustedShopper field to given value.

### HasTrustedShopper

`func (o *CheckoutBalanceCheckRequest) HasTrustedShopper() bool`

HasTrustedShopper returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


