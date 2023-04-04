# PaymentDonationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountInfo** | Pointer to [**AccountInfo**](AccountInfo.md) |  | [optional] 
**AdditionalData** | Pointer to **map[string]string** | This field contains additional data, which may be required for a particular payment request.  The &#x60;additionalData&#x60; object consists of entries, each of which includes the key and value. | [optional] 
**Amount** | [**Amount**](Amount.md) |  | 
**ApplicationInfo** | Pointer to [**ApplicationInfo**](ApplicationInfo.md) |  | [optional] 
**AuthenticationData** | Pointer to [**AuthenticationData**](AuthenticationData.md) |  | [optional] 
**BillingAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**BrowserInfo** | Pointer to [**BrowserInfo**](BrowserInfo.md) |  | [optional] 
**CaptureDelayHours** | Pointer to **int32** | The delay between the authorisation and scheduled auto-capture, specified in hours. | [optional] 
**Channel** | Pointer to **string** | The platform where a payment transaction takes place. This field is optional for filtering out payment methods that are only available on specific platforms. If this value is not set, then we will try to infer it from the &#x60;sdkVersion&#x60; or &#x60;token&#x60;.  Possible values: * iOS * Android * Web | [optional] 
**CheckoutAttemptId** | Pointer to **string** | Checkout attempt ID that corresponds to the Id generated for tracking user payment journey. | [optional] 
**Company** | Pointer to [**Company**](Company.md) |  | [optional] 
**ConversionId** | Pointer to **string** | Conversion ID that corresponds to the Id generated for tracking user payment journey. | [optional] 
**CountryCode** | Pointer to **string** | The shopper country.  Format: [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) Example: NL or DE | [optional] 
**DateOfBirth** | Pointer to **string** | The shopper&#39;s date of birth.  Format [ISO-8601](https://www.w3.org/TR/NOTE-datetime): YYYY-MM-DD | [optional] 
**DccQuote** | Pointer to [**ForexQuote**](ForexQuote.md) |  | [optional] 
**DeliveryAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**DeliveryDate** | Pointer to **time.Time** | The date and time the purchased goods should be delivered.  Format [ISO 8601](https://www.w3.org/TR/NOTE-datetime): YYYY-MM-DDThh:mm:ss.sssTZD  Example: 2017-07-17T13:42:40.428+01:00 | [optional] 
**DeviceFingerprint** | Pointer to **string** | A string containing the shopper&#39;s device fingerprint. For more information, refer to [Device fingerprinting](https://docs.adyen.com/risk-management/device-fingerprinting). | [optional] 
**DonationAccount** | **string** | Donation account to which the transaction is credited. | 
**DonationOriginalPspReference** | Pointer to **string** | PSP reference of the transaction from which the donation token is generated. Required when &#x60;donationToken&#x60; is provided. | [optional] 
**DonationToken** | Pointer to **string** | Donation token received in the &#x60;/payments&#x60; call. | [optional] 
**EnableOneClick** | Pointer to **bool** | When true and &#x60;shopperReference&#x60; is provided, the shopper will be asked if the payment details should be stored for future one-click payments. | [optional] 
**EnablePayOut** | Pointer to **bool** | When true and &#x60;shopperReference&#x60; is provided, the payment details will be tokenized for payouts. | [optional] 
**EnableRecurring** | Pointer to **bool** | When true and &#x60;shopperReference&#x60; is provided, the payment details will be tokenized for recurring payments. | [optional] 
**EntityType** | Pointer to **string** | The type of the entity the payment is processed for. | [optional] 
**FraudOffset** | Pointer to **int32** | An integer value that is added to the normal fraud score. The value can be either positive or negative. | [optional] 
**IndustryUsage** | Pointer to **string** | The reason for the amount update. Possible values:  * **delayedCharge**  * **noShow**  * **installment** | [optional] 
**Installments** | Pointer to [**Installments**](Installments.md) |  | [optional] 
**LineItems** | Pointer to [**[]LineItem**](LineItem.md) | Price and product information about the purchased items, to be included on the invoice sent to the shopper. &gt; This field is required for 3x 4x Oney, Affirm, Afterpay, Clearpay, Klarna, Ratepay, Zip and Atome. | [optional] 
**LocalizedShopperStatement** | Pointer to **map[string]string** | This field allows merchants to use dynamic shopper statement in local character sets. The local shopper statement field can be supplied in markets where localized merchant descriptors are used. Currently, Adyen only supports this in the Japanese market .The available character sets at the moment are: * Processing in Japan: **ja-Kana** The character set **ja-Kana** supports UTF-8 based Katakana and alphanumeric and special characters. Merchants should send the Katakana shopperStatement in full-width characters.  An example request would be: &gt; {   \&quot;shopperStatement\&quot; : \&quot;ADYEN - SELLER-A\&quot;,   \&quot;localizedShopperStatement\&quot; : {     \&quot;ja-Kana\&quot; : \&quot;ADYEN - セラーA\&quot;   } } We recommend merchants to always supply the field localizedShopperStatement in addition to the field shopperStatement.It is issuer dependent whether the localized shopper statement field is supported. In the case of non-domestic transactions (e.g. US-issued cards processed in JP) the field &#x60;shopperStatement&#x60; is used to modify the statement of the shopper. Adyen handles the complexity of ensuring the correct descriptors are assigned. | [optional] 
**Mandate** | Pointer to [**Mandate**](Mandate.md) |  | [optional] 
**Mcc** | Pointer to **string** | The [merchant category code](https://en.wikipedia.org/wiki/Merchant_category_code) (MCC) is a four-digit number, which relates to a particular market segment. This code reflects the predominant activity that is conducted by the merchant. | [optional] 
**MerchantAccount** | **string** | The merchant account identifier, with which you want to process the transaction. | 
**MerchantOrderReference** | Pointer to **string** | This reference allows linking multiple transactions to each other for reporting purposes (i.e. order auth-rate). The reference should be unique per billing cycle. The same merchant order reference should never be reused after the first authorised attempt. If used, this field should be supplied for all incoming authorisations. &gt; We strongly recommend you send the &#x60;merchantOrderReference&#x60; value to benefit from linking payment requests when authorisation retries take place. In addition, we recommend you provide &#x60;retry.orderAttemptNumber&#x60;, &#x60;retry.chainAttemptNumber&#x60;, and &#x60;retry.skipRetry&#x60; values in &#x60;PaymentRequest.additionalData&#x60;. | [optional] 
**MerchantRiskIndicator** | Pointer to [**MerchantRiskIndicator**](MerchantRiskIndicator.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** | Metadata consists of entries, each of which includes a key and a value. Limits: * Maximum 20 key-value pairs per request. When exceeding, the \&quot;177\&quot; error occurs: \&quot;Metadata size exceeds limit\&quot;. * Maximum 20 characters per key. * Maximum 80 characters per value.  | [optional] 
**MpiData** | Pointer to [**ThreeDSecureData**](ThreeDSecureData.md) |  | [optional] 
**Order** | Pointer to [**CheckoutOrder**](CheckoutOrder.md) |  | [optional] 
**OrderReference** | Pointer to **string** | When you are doing multiple partial (gift card) payments, this is the &#x60;pspReference&#x60; of the first payment. We use this to link the multiple payments to each other. As your own reference for linking multiple payments, use the &#x60;merchantOrderReference&#x60;instead. | [optional] 
**Origin** | Pointer to **string** | Required for the 3D Secure 2 &#x60;channel&#x60; **Web** integration.  Set this parameter to the origin URL of the page that you are loading the 3D Secure Component from. | [optional] 
**PaymentMethod** | [**CheckoutPaymentMethod**](CheckoutPaymentMethod.md) |  | 
**PlatformChargebackLogic** | Pointer to [**PlatformChargebackLogic**](PlatformChargebackLogic.md) |  | [optional] 
**RecurringExpiry** | Pointer to **string** | Date after which no further authorisations shall be performed. Only for 3D Secure 2. | [optional] 
**RecurringFrequency** | Pointer to **string** | Minimum number of days between authorisations. Only for 3D Secure 2. | [optional] 
**RecurringProcessingModel** | Pointer to **string** | Defines a recurring payment type. Allowed values: * &#x60;Subscription&#x60; – A transaction for a fixed or variable amount, which follows a fixed schedule. * &#x60;CardOnFile&#x60; – With a card-on-file (CoF) transaction, card details are stored to enable one-click or omnichannel journeys, or simply to streamline the checkout process. Any subscription not following a fixed schedule is also considered a card-on-file transaction. * &#x60;UnscheduledCardOnFile&#x60; – An unscheduled card-on-file (UCoF) transaction is a transaction that occurs on a non-fixed schedule and/or have variable amounts. For example, automatic top-ups when a cardholder&#39;s balance drops below a certain amount.  | [optional] 
**RedirectFromIssuerMethod** | Pointer to **string** | Specifies the redirect method (GET or POST) when redirecting back from the issuer. | [optional] 
**RedirectToIssuerMethod** | Pointer to **string** | Specifies the redirect method (GET or POST) when redirecting to the issuer. | [optional] 
**Reference** | **string** | The reference to uniquely identify a payment. This reference is used in all communication with you about the payment status. We recommend using a unique value per payment; however, it is not a requirement. If you need to provide multiple references for a transaction, separate them with hyphens (\&quot;-\&quot;). Maximum length: 80 characters. | 
**ReturnUrl** | **string** | The URL to return to in case of a redirection. The format depends on the channel. This URL can have a maximum of 1024 characters. * For web, include the protocol &#x60;http://&#x60; or &#x60;https://&#x60;. You can also include your own additional query parameters, for example, shopper ID or order reference number. Example: &#x60;https://your-company.com/checkout?shopperOrder&#x3D;12xy&#x60; * For iOS, use the custom URL for your app. To know more about setting custom URL schemes, refer to the [Apple Developer documentation](https://developer.apple.com/documentation/uikit/inter-process_communication/allowing_apps_and_websites_to_link_to_your_content/defining_a_custom_url_scheme_for_your_app). Example: &#x60;my-app://&#x60; * For Android, use a custom URL handled by an Activity on your app. You can configure it with an [intent filter](https://developer.android.com/guide/components/intents-filters). Example: &#x60;my-app://your.package.name&#x60; | 
**RiskData** | Pointer to [**RiskData**](RiskData.md) |  | [optional] 
**SessionValidity** | Pointer to **string** | The date and time until when the session remains valid, in [ISO 8601](https://www.w3.org/TR/NOTE-datetime) format.  For example: 2020-07-18T15:42:40.428+01:00 | [optional] 
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
**StorePaymentMethod** | Pointer to **bool** | When true and &#x60;shopperReference&#x60; is provided, the payment details will be stored. | [optional] 
**TelephoneNumber** | Pointer to **string** | The shopper&#39;s telephone number. | [optional] 
**ThreeDS2RequestData** | Pointer to [**ThreeDS2RequestData**](ThreeDS2RequestData.md) |  | [optional] 
**ThreeDSAuthenticationOnly** | Pointer to **bool** | If set to true, you will only perform the [3D Secure 2 authentication](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only), and not the payment authorisation. | [optional] [default to false]
**TrustedShopper** | Pointer to **bool** | Set to true if the payment should be routed to a trusted MID. | [optional] 

## Methods

### NewPaymentDonationRequest

`func NewPaymentDonationRequest(amount Amount, donationAccount string, merchantAccount string, paymentMethod CheckoutPaymentMethod, reference string, returnUrl string, ) *PaymentDonationRequest`

NewPaymentDonationRequest instantiates a new PaymentDonationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentDonationRequestWithDefaults

`func NewPaymentDonationRequestWithDefaults() *PaymentDonationRequest`

NewPaymentDonationRequestWithDefaults instantiates a new PaymentDonationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountInfo

`func (o *PaymentDonationRequest) GetAccountInfo() AccountInfo`

GetAccountInfo returns the AccountInfo field if non-nil, zero value otherwise.

### GetAccountInfoOk

`func (o *PaymentDonationRequest) GetAccountInfoOk() (*AccountInfo, bool)`

GetAccountInfoOk returns a tuple with the AccountInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountInfo

`func (o *PaymentDonationRequest) SetAccountInfo(v AccountInfo)`

SetAccountInfo sets AccountInfo field to given value.

### HasAccountInfo

`func (o *PaymentDonationRequest) HasAccountInfo() bool`

HasAccountInfo returns a boolean if a field has been set.

### GetAdditionalData

`func (o *PaymentDonationRequest) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *PaymentDonationRequest) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *PaymentDonationRequest) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *PaymentDonationRequest) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetAmount

`func (o *PaymentDonationRequest) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentDonationRequest) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentDonationRequest) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetApplicationInfo

`func (o *PaymentDonationRequest) GetApplicationInfo() ApplicationInfo`

GetApplicationInfo returns the ApplicationInfo field if non-nil, zero value otherwise.

### GetApplicationInfoOk

`func (o *PaymentDonationRequest) GetApplicationInfoOk() (*ApplicationInfo, bool)`

GetApplicationInfoOk returns a tuple with the ApplicationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationInfo

`func (o *PaymentDonationRequest) SetApplicationInfo(v ApplicationInfo)`

SetApplicationInfo sets ApplicationInfo field to given value.

### HasApplicationInfo

`func (o *PaymentDonationRequest) HasApplicationInfo() bool`

HasApplicationInfo returns a boolean if a field has been set.

### GetAuthenticationData

`func (o *PaymentDonationRequest) GetAuthenticationData() AuthenticationData`

GetAuthenticationData returns the AuthenticationData field if non-nil, zero value otherwise.

### GetAuthenticationDataOk

`func (o *PaymentDonationRequest) GetAuthenticationDataOk() (*AuthenticationData, bool)`

GetAuthenticationDataOk returns a tuple with the AuthenticationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationData

`func (o *PaymentDonationRequest) SetAuthenticationData(v AuthenticationData)`

SetAuthenticationData sets AuthenticationData field to given value.

### HasAuthenticationData

`func (o *PaymentDonationRequest) HasAuthenticationData() bool`

HasAuthenticationData returns a boolean if a field has been set.

### GetBillingAddress

`func (o *PaymentDonationRequest) GetBillingAddress() Address`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *PaymentDonationRequest) GetBillingAddressOk() (*Address, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *PaymentDonationRequest) SetBillingAddress(v Address)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *PaymentDonationRequest) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetBrowserInfo

`func (o *PaymentDonationRequest) GetBrowserInfo() BrowserInfo`

GetBrowserInfo returns the BrowserInfo field if non-nil, zero value otherwise.

### GetBrowserInfoOk

`func (o *PaymentDonationRequest) GetBrowserInfoOk() (*BrowserInfo, bool)`

GetBrowserInfoOk returns a tuple with the BrowserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserInfo

`func (o *PaymentDonationRequest) SetBrowserInfo(v BrowserInfo)`

SetBrowserInfo sets BrowserInfo field to given value.

### HasBrowserInfo

`func (o *PaymentDonationRequest) HasBrowserInfo() bool`

HasBrowserInfo returns a boolean if a field has been set.

### GetCaptureDelayHours

`func (o *PaymentDonationRequest) GetCaptureDelayHours() int32`

GetCaptureDelayHours returns the CaptureDelayHours field if non-nil, zero value otherwise.

### GetCaptureDelayHoursOk

`func (o *PaymentDonationRequest) GetCaptureDelayHoursOk() (*int32, bool)`

GetCaptureDelayHoursOk returns a tuple with the CaptureDelayHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureDelayHours

`func (o *PaymentDonationRequest) SetCaptureDelayHours(v int32)`

SetCaptureDelayHours sets CaptureDelayHours field to given value.

### HasCaptureDelayHours

`func (o *PaymentDonationRequest) HasCaptureDelayHours() bool`

HasCaptureDelayHours returns a boolean if a field has been set.

### GetChannel

`func (o *PaymentDonationRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PaymentDonationRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PaymentDonationRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *PaymentDonationRequest) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCheckoutAttemptId

`func (o *PaymentDonationRequest) GetCheckoutAttemptId() string`

GetCheckoutAttemptId returns the CheckoutAttemptId field if non-nil, zero value otherwise.

### GetCheckoutAttemptIdOk

`func (o *PaymentDonationRequest) GetCheckoutAttemptIdOk() (*string, bool)`

GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutAttemptId

`func (o *PaymentDonationRequest) SetCheckoutAttemptId(v string)`

SetCheckoutAttemptId sets CheckoutAttemptId field to given value.

### HasCheckoutAttemptId

`func (o *PaymentDonationRequest) HasCheckoutAttemptId() bool`

HasCheckoutAttemptId returns a boolean if a field has been set.

### GetCompany

`func (o *PaymentDonationRequest) GetCompany() Company`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PaymentDonationRequest) GetCompanyOk() (*Company, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PaymentDonationRequest) SetCompany(v Company)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *PaymentDonationRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetConversionId

`func (o *PaymentDonationRequest) GetConversionId() string`

GetConversionId returns the ConversionId field if non-nil, zero value otherwise.

### GetConversionIdOk

`func (o *PaymentDonationRequest) GetConversionIdOk() (*string, bool)`

GetConversionIdOk returns a tuple with the ConversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionId

`func (o *PaymentDonationRequest) SetConversionId(v string)`

SetConversionId sets ConversionId field to given value.

### HasConversionId

`func (o *PaymentDonationRequest) HasConversionId() bool`

HasConversionId returns a boolean if a field has been set.

### GetCountryCode

`func (o *PaymentDonationRequest) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PaymentDonationRequest) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PaymentDonationRequest) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *PaymentDonationRequest) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *PaymentDonationRequest) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *PaymentDonationRequest) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *PaymentDonationRequest) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *PaymentDonationRequest) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetDccQuote

`func (o *PaymentDonationRequest) GetDccQuote() ForexQuote`

GetDccQuote returns the DccQuote field if non-nil, zero value otherwise.

### GetDccQuoteOk

`func (o *PaymentDonationRequest) GetDccQuoteOk() (*ForexQuote, bool)`

GetDccQuoteOk returns a tuple with the DccQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDccQuote

`func (o *PaymentDonationRequest) SetDccQuote(v ForexQuote)`

SetDccQuote sets DccQuote field to given value.

### HasDccQuote

`func (o *PaymentDonationRequest) HasDccQuote() bool`

HasDccQuote returns a boolean if a field has been set.

### GetDeliveryAddress

`func (o *PaymentDonationRequest) GetDeliveryAddress() Address`

GetDeliveryAddress returns the DeliveryAddress field if non-nil, zero value otherwise.

### GetDeliveryAddressOk

`func (o *PaymentDonationRequest) GetDeliveryAddressOk() (*Address, bool)`

GetDeliveryAddressOk returns a tuple with the DeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddress

`func (o *PaymentDonationRequest) SetDeliveryAddress(v Address)`

SetDeliveryAddress sets DeliveryAddress field to given value.

### HasDeliveryAddress

`func (o *PaymentDonationRequest) HasDeliveryAddress() bool`

HasDeliveryAddress returns a boolean if a field has been set.

### GetDeliveryDate

`func (o *PaymentDonationRequest) GetDeliveryDate() time.Time`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *PaymentDonationRequest) GetDeliveryDateOk() (*time.Time, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *PaymentDonationRequest) SetDeliveryDate(v time.Time)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *PaymentDonationRequest) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### GetDeviceFingerprint

`func (o *PaymentDonationRequest) GetDeviceFingerprint() string`

GetDeviceFingerprint returns the DeviceFingerprint field if non-nil, zero value otherwise.

### GetDeviceFingerprintOk

`func (o *PaymentDonationRequest) GetDeviceFingerprintOk() (*string, bool)`

GetDeviceFingerprintOk returns a tuple with the DeviceFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFingerprint

`func (o *PaymentDonationRequest) SetDeviceFingerprint(v string)`

SetDeviceFingerprint sets DeviceFingerprint field to given value.

### HasDeviceFingerprint

`func (o *PaymentDonationRequest) HasDeviceFingerprint() bool`

HasDeviceFingerprint returns a boolean if a field has been set.

### GetDonationAccount

`func (o *PaymentDonationRequest) GetDonationAccount() string`

GetDonationAccount returns the DonationAccount field if non-nil, zero value otherwise.

### GetDonationAccountOk

`func (o *PaymentDonationRequest) GetDonationAccountOk() (*string, bool)`

GetDonationAccountOk returns a tuple with the DonationAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDonationAccount

`func (o *PaymentDonationRequest) SetDonationAccount(v string)`

SetDonationAccount sets DonationAccount field to given value.


### GetDonationOriginalPspReference

`func (o *PaymentDonationRequest) GetDonationOriginalPspReference() string`

GetDonationOriginalPspReference returns the DonationOriginalPspReference field if non-nil, zero value otherwise.

### GetDonationOriginalPspReferenceOk

`func (o *PaymentDonationRequest) GetDonationOriginalPspReferenceOk() (*string, bool)`

GetDonationOriginalPspReferenceOk returns a tuple with the DonationOriginalPspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDonationOriginalPspReference

`func (o *PaymentDonationRequest) SetDonationOriginalPspReference(v string)`

SetDonationOriginalPspReference sets DonationOriginalPspReference field to given value.

### HasDonationOriginalPspReference

`func (o *PaymentDonationRequest) HasDonationOriginalPspReference() bool`

HasDonationOriginalPspReference returns a boolean if a field has been set.

### GetDonationToken

`func (o *PaymentDonationRequest) GetDonationToken() string`

GetDonationToken returns the DonationToken field if non-nil, zero value otherwise.

### GetDonationTokenOk

`func (o *PaymentDonationRequest) GetDonationTokenOk() (*string, bool)`

GetDonationTokenOk returns a tuple with the DonationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDonationToken

`func (o *PaymentDonationRequest) SetDonationToken(v string)`

SetDonationToken sets DonationToken field to given value.

### HasDonationToken

`func (o *PaymentDonationRequest) HasDonationToken() bool`

HasDonationToken returns a boolean if a field has been set.

### GetEnableOneClick

`func (o *PaymentDonationRequest) GetEnableOneClick() bool`

GetEnableOneClick returns the EnableOneClick field if non-nil, zero value otherwise.

### GetEnableOneClickOk

`func (o *PaymentDonationRequest) GetEnableOneClickOk() (*bool, bool)`

GetEnableOneClickOk returns a tuple with the EnableOneClick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOneClick

`func (o *PaymentDonationRequest) SetEnableOneClick(v bool)`

SetEnableOneClick sets EnableOneClick field to given value.

### HasEnableOneClick

`func (o *PaymentDonationRequest) HasEnableOneClick() bool`

HasEnableOneClick returns a boolean if a field has been set.

### GetEnablePayOut

`func (o *PaymentDonationRequest) GetEnablePayOut() bool`

GetEnablePayOut returns the EnablePayOut field if non-nil, zero value otherwise.

### GetEnablePayOutOk

`func (o *PaymentDonationRequest) GetEnablePayOutOk() (*bool, bool)`

GetEnablePayOutOk returns a tuple with the EnablePayOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePayOut

`func (o *PaymentDonationRequest) SetEnablePayOut(v bool)`

SetEnablePayOut sets EnablePayOut field to given value.

### HasEnablePayOut

`func (o *PaymentDonationRequest) HasEnablePayOut() bool`

HasEnablePayOut returns a boolean if a field has been set.

### GetEnableRecurring

`func (o *PaymentDonationRequest) GetEnableRecurring() bool`

GetEnableRecurring returns the EnableRecurring field if non-nil, zero value otherwise.

### GetEnableRecurringOk

`func (o *PaymentDonationRequest) GetEnableRecurringOk() (*bool, bool)`

GetEnableRecurringOk returns a tuple with the EnableRecurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRecurring

`func (o *PaymentDonationRequest) SetEnableRecurring(v bool)`

SetEnableRecurring sets EnableRecurring field to given value.

### HasEnableRecurring

`func (o *PaymentDonationRequest) HasEnableRecurring() bool`

HasEnableRecurring returns a boolean if a field has been set.

### GetEntityType

`func (o *PaymentDonationRequest) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *PaymentDonationRequest) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *PaymentDonationRequest) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *PaymentDonationRequest) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetFraudOffset

`func (o *PaymentDonationRequest) GetFraudOffset() int32`

GetFraudOffset returns the FraudOffset field if non-nil, zero value otherwise.

### GetFraudOffsetOk

`func (o *PaymentDonationRequest) GetFraudOffsetOk() (*int32, bool)`

GetFraudOffsetOk returns a tuple with the FraudOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudOffset

`func (o *PaymentDonationRequest) SetFraudOffset(v int32)`

SetFraudOffset sets FraudOffset field to given value.

### HasFraudOffset

`func (o *PaymentDonationRequest) HasFraudOffset() bool`

HasFraudOffset returns a boolean if a field has been set.

### GetIndustryUsage

`func (o *PaymentDonationRequest) GetIndustryUsage() string`

GetIndustryUsage returns the IndustryUsage field if non-nil, zero value otherwise.

### GetIndustryUsageOk

`func (o *PaymentDonationRequest) GetIndustryUsageOk() (*string, bool)`

GetIndustryUsageOk returns a tuple with the IndustryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryUsage

`func (o *PaymentDonationRequest) SetIndustryUsage(v string)`

SetIndustryUsage sets IndustryUsage field to given value.

### HasIndustryUsage

`func (o *PaymentDonationRequest) HasIndustryUsage() bool`

HasIndustryUsage returns a boolean if a field has been set.

### GetInstallments

`func (o *PaymentDonationRequest) GetInstallments() Installments`

GetInstallments returns the Installments field if non-nil, zero value otherwise.

### GetInstallmentsOk

`func (o *PaymentDonationRequest) GetInstallmentsOk() (*Installments, bool)`

GetInstallmentsOk returns a tuple with the Installments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallments

`func (o *PaymentDonationRequest) SetInstallments(v Installments)`

SetInstallments sets Installments field to given value.

### HasInstallments

`func (o *PaymentDonationRequest) HasInstallments() bool`

HasInstallments returns a boolean if a field has been set.

### GetLineItems

`func (o *PaymentDonationRequest) GetLineItems() []LineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *PaymentDonationRequest) GetLineItemsOk() (*[]LineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *PaymentDonationRequest) SetLineItems(v []LineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *PaymentDonationRequest) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetLocalizedShopperStatement

`func (o *PaymentDonationRequest) GetLocalizedShopperStatement() map[string]string`

GetLocalizedShopperStatement returns the LocalizedShopperStatement field if non-nil, zero value otherwise.

### GetLocalizedShopperStatementOk

`func (o *PaymentDonationRequest) GetLocalizedShopperStatementOk() (*map[string]string, bool)`

GetLocalizedShopperStatementOk returns a tuple with the LocalizedShopperStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedShopperStatement

`func (o *PaymentDonationRequest) SetLocalizedShopperStatement(v map[string]string)`

SetLocalizedShopperStatement sets LocalizedShopperStatement field to given value.

### HasLocalizedShopperStatement

`func (o *PaymentDonationRequest) HasLocalizedShopperStatement() bool`

HasLocalizedShopperStatement returns a boolean if a field has been set.

### GetMandate

`func (o *PaymentDonationRequest) GetMandate() Mandate`

GetMandate returns the Mandate field if non-nil, zero value otherwise.

### GetMandateOk

`func (o *PaymentDonationRequest) GetMandateOk() (*Mandate, bool)`

GetMandateOk returns a tuple with the Mandate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandate

`func (o *PaymentDonationRequest) SetMandate(v Mandate)`

SetMandate sets Mandate field to given value.

### HasMandate

`func (o *PaymentDonationRequest) HasMandate() bool`

HasMandate returns a boolean if a field has been set.

### GetMcc

`func (o *PaymentDonationRequest) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *PaymentDonationRequest) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *PaymentDonationRequest) SetMcc(v string)`

SetMcc sets Mcc field to given value.

### HasMcc

`func (o *PaymentDonationRequest) HasMcc() bool`

HasMcc returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *PaymentDonationRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *PaymentDonationRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *PaymentDonationRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetMerchantOrderReference

`func (o *PaymentDonationRequest) GetMerchantOrderReference() string`

GetMerchantOrderReference returns the MerchantOrderReference field if non-nil, zero value otherwise.

### GetMerchantOrderReferenceOk

`func (o *PaymentDonationRequest) GetMerchantOrderReferenceOk() (*string, bool)`

GetMerchantOrderReferenceOk returns a tuple with the MerchantOrderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderReference

`func (o *PaymentDonationRequest) SetMerchantOrderReference(v string)`

SetMerchantOrderReference sets MerchantOrderReference field to given value.

### HasMerchantOrderReference

`func (o *PaymentDonationRequest) HasMerchantOrderReference() bool`

HasMerchantOrderReference returns a boolean if a field has been set.

### GetMerchantRiskIndicator

`func (o *PaymentDonationRequest) GetMerchantRiskIndicator() MerchantRiskIndicator`

GetMerchantRiskIndicator returns the MerchantRiskIndicator field if non-nil, zero value otherwise.

### GetMerchantRiskIndicatorOk

`func (o *PaymentDonationRequest) GetMerchantRiskIndicatorOk() (*MerchantRiskIndicator, bool)`

GetMerchantRiskIndicatorOk returns a tuple with the MerchantRiskIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantRiskIndicator

`func (o *PaymentDonationRequest) SetMerchantRiskIndicator(v MerchantRiskIndicator)`

SetMerchantRiskIndicator sets MerchantRiskIndicator field to given value.

### HasMerchantRiskIndicator

`func (o *PaymentDonationRequest) HasMerchantRiskIndicator() bool`

HasMerchantRiskIndicator returns a boolean if a field has been set.

### GetMetadata

`func (o *PaymentDonationRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PaymentDonationRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PaymentDonationRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PaymentDonationRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMpiData

`func (o *PaymentDonationRequest) GetMpiData() ThreeDSecureData`

GetMpiData returns the MpiData field if non-nil, zero value otherwise.

### GetMpiDataOk

`func (o *PaymentDonationRequest) GetMpiDataOk() (*ThreeDSecureData, bool)`

GetMpiDataOk returns a tuple with the MpiData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpiData

`func (o *PaymentDonationRequest) SetMpiData(v ThreeDSecureData)`

SetMpiData sets MpiData field to given value.

### HasMpiData

`func (o *PaymentDonationRequest) HasMpiData() bool`

HasMpiData returns a boolean if a field has been set.

### GetOrder

`func (o *PaymentDonationRequest) GetOrder() CheckoutOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PaymentDonationRequest) GetOrderOk() (*CheckoutOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PaymentDonationRequest) SetOrder(v CheckoutOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PaymentDonationRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOrderReference

`func (o *PaymentDonationRequest) GetOrderReference() string`

GetOrderReference returns the OrderReference field if non-nil, zero value otherwise.

### GetOrderReferenceOk

`func (o *PaymentDonationRequest) GetOrderReferenceOk() (*string, bool)`

GetOrderReferenceOk returns a tuple with the OrderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderReference

`func (o *PaymentDonationRequest) SetOrderReference(v string)`

SetOrderReference sets OrderReference field to given value.

### HasOrderReference

`func (o *PaymentDonationRequest) HasOrderReference() bool`

HasOrderReference returns a boolean if a field has been set.

### GetOrigin

`func (o *PaymentDonationRequest) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *PaymentDonationRequest) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *PaymentDonationRequest) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *PaymentDonationRequest) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *PaymentDonationRequest) GetPaymentMethod() CheckoutPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PaymentDonationRequest) GetPaymentMethodOk() (*CheckoutPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PaymentDonationRequest) SetPaymentMethod(v CheckoutPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetPlatformChargebackLogic

`func (o *PaymentDonationRequest) GetPlatformChargebackLogic() PlatformChargebackLogic`

GetPlatformChargebackLogic returns the PlatformChargebackLogic field if non-nil, zero value otherwise.

### GetPlatformChargebackLogicOk

`func (o *PaymentDonationRequest) GetPlatformChargebackLogicOk() (*PlatformChargebackLogic, bool)`

GetPlatformChargebackLogicOk returns a tuple with the PlatformChargebackLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformChargebackLogic

`func (o *PaymentDonationRequest) SetPlatformChargebackLogic(v PlatformChargebackLogic)`

SetPlatformChargebackLogic sets PlatformChargebackLogic field to given value.

### HasPlatformChargebackLogic

`func (o *PaymentDonationRequest) HasPlatformChargebackLogic() bool`

HasPlatformChargebackLogic returns a boolean if a field has been set.

### GetRecurringExpiry

`func (o *PaymentDonationRequest) GetRecurringExpiry() string`

GetRecurringExpiry returns the RecurringExpiry field if non-nil, zero value otherwise.

### GetRecurringExpiryOk

`func (o *PaymentDonationRequest) GetRecurringExpiryOk() (*string, bool)`

GetRecurringExpiryOk returns a tuple with the RecurringExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringExpiry

`func (o *PaymentDonationRequest) SetRecurringExpiry(v string)`

SetRecurringExpiry sets RecurringExpiry field to given value.

### HasRecurringExpiry

`func (o *PaymentDonationRequest) HasRecurringExpiry() bool`

HasRecurringExpiry returns a boolean if a field has been set.

### GetRecurringFrequency

`func (o *PaymentDonationRequest) GetRecurringFrequency() string`

GetRecurringFrequency returns the RecurringFrequency field if non-nil, zero value otherwise.

### GetRecurringFrequencyOk

`func (o *PaymentDonationRequest) GetRecurringFrequencyOk() (*string, bool)`

GetRecurringFrequencyOk returns a tuple with the RecurringFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringFrequency

`func (o *PaymentDonationRequest) SetRecurringFrequency(v string)`

SetRecurringFrequency sets RecurringFrequency field to given value.

### HasRecurringFrequency

`func (o *PaymentDonationRequest) HasRecurringFrequency() bool`

HasRecurringFrequency returns a boolean if a field has been set.

### GetRecurringProcessingModel

`func (o *PaymentDonationRequest) GetRecurringProcessingModel() string`

GetRecurringProcessingModel returns the RecurringProcessingModel field if non-nil, zero value otherwise.

### GetRecurringProcessingModelOk

`func (o *PaymentDonationRequest) GetRecurringProcessingModelOk() (*string, bool)`

GetRecurringProcessingModelOk returns a tuple with the RecurringProcessingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringProcessingModel

`func (o *PaymentDonationRequest) SetRecurringProcessingModel(v string)`

SetRecurringProcessingModel sets RecurringProcessingModel field to given value.

### HasRecurringProcessingModel

`func (o *PaymentDonationRequest) HasRecurringProcessingModel() bool`

HasRecurringProcessingModel returns a boolean if a field has been set.

### GetRedirectFromIssuerMethod

`func (o *PaymentDonationRequest) GetRedirectFromIssuerMethod() string`

GetRedirectFromIssuerMethod returns the RedirectFromIssuerMethod field if non-nil, zero value otherwise.

### GetRedirectFromIssuerMethodOk

`func (o *PaymentDonationRequest) GetRedirectFromIssuerMethodOk() (*string, bool)`

GetRedirectFromIssuerMethodOk returns a tuple with the RedirectFromIssuerMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectFromIssuerMethod

`func (o *PaymentDonationRequest) SetRedirectFromIssuerMethod(v string)`

SetRedirectFromIssuerMethod sets RedirectFromIssuerMethod field to given value.

### HasRedirectFromIssuerMethod

`func (o *PaymentDonationRequest) HasRedirectFromIssuerMethod() bool`

HasRedirectFromIssuerMethod returns a boolean if a field has been set.

### GetRedirectToIssuerMethod

`func (o *PaymentDonationRequest) GetRedirectToIssuerMethod() string`

GetRedirectToIssuerMethod returns the RedirectToIssuerMethod field if non-nil, zero value otherwise.

### GetRedirectToIssuerMethodOk

`func (o *PaymentDonationRequest) GetRedirectToIssuerMethodOk() (*string, bool)`

GetRedirectToIssuerMethodOk returns a tuple with the RedirectToIssuerMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectToIssuerMethod

`func (o *PaymentDonationRequest) SetRedirectToIssuerMethod(v string)`

SetRedirectToIssuerMethod sets RedirectToIssuerMethod field to given value.

### HasRedirectToIssuerMethod

`func (o *PaymentDonationRequest) HasRedirectToIssuerMethod() bool`

HasRedirectToIssuerMethod returns a boolean if a field has been set.

### GetReference

`func (o *PaymentDonationRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaymentDonationRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaymentDonationRequest) SetReference(v string)`

SetReference sets Reference field to given value.


### GetReturnUrl

`func (o *PaymentDonationRequest) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *PaymentDonationRequest) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *PaymentDonationRequest) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.


### GetRiskData

`func (o *PaymentDonationRequest) GetRiskData() RiskData`

GetRiskData returns the RiskData field if non-nil, zero value otherwise.

### GetRiskDataOk

`func (o *PaymentDonationRequest) GetRiskDataOk() (*RiskData, bool)`

GetRiskDataOk returns a tuple with the RiskData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskData

`func (o *PaymentDonationRequest) SetRiskData(v RiskData)`

SetRiskData sets RiskData field to given value.

### HasRiskData

`func (o *PaymentDonationRequest) HasRiskData() bool`

HasRiskData returns a boolean if a field has been set.

### GetSessionValidity

`func (o *PaymentDonationRequest) GetSessionValidity() string`

GetSessionValidity returns the SessionValidity field if non-nil, zero value otherwise.

### GetSessionValidityOk

`func (o *PaymentDonationRequest) GetSessionValidityOk() (*string, bool)`

GetSessionValidityOk returns a tuple with the SessionValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionValidity

`func (o *PaymentDonationRequest) SetSessionValidity(v string)`

SetSessionValidity sets SessionValidity field to given value.

### HasSessionValidity

`func (o *PaymentDonationRequest) HasSessionValidity() bool`

HasSessionValidity returns a boolean if a field has been set.

### GetShopperEmail

`func (o *PaymentDonationRequest) GetShopperEmail() string`

GetShopperEmail returns the ShopperEmail field if non-nil, zero value otherwise.

### GetShopperEmailOk

`func (o *PaymentDonationRequest) GetShopperEmailOk() (*string, bool)`

GetShopperEmailOk returns a tuple with the ShopperEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperEmail

`func (o *PaymentDonationRequest) SetShopperEmail(v string)`

SetShopperEmail sets ShopperEmail field to given value.

### HasShopperEmail

`func (o *PaymentDonationRequest) HasShopperEmail() bool`

HasShopperEmail returns a boolean if a field has been set.

### GetShopperIP

`func (o *PaymentDonationRequest) GetShopperIP() string`

GetShopperIP returns the ShopperIP field if non-nil, zero value otherwise.

### GetShopperIPOk

`func (o *PaymentDonationRequest) GetShopperIPOk() (*string, bool)`

GetShopperIPOk returns a tuple with the ShopperIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperIP

`func (o *PaymentDonationRequest) SetShopperIP(v string)`

SetShopperIP sets ShopperIP field to given value.

### HasShopperIP

`func (o *PaymentDonationRequest) HasShopperIP() bool`

HasShopperIP returns a boolean if a field has been set.

### GetShopperInteraction

`func (o *PaymentDonationRequest) GetShopperInteraction() string`

GetShopperInteraction returns the ShopperInteraction field if non-nil, zero value otherwise.

### GetShopperInteractionOk

`func (o *PaymentDonationRequest) GetShopperInteractionOk() (*string, bool)`

GetShopperInteractionOk returns a tuple with the ShopperInteraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperInteraction

`func (o *PaymentDonationRequest) SetShopperInteraction(v string)`

SetShopperInteraction sets ShopperInteraction field to given value.

### HasShopperInteraction

`func (o *PaymentDonationRequest) HasShopperInteraction() bool`

HasShopperInteraction returns a boolean if a field has been set.

### GetShopperLocale

`func (o *PaymentDonationRequest) GetShopperLocale() string`

GetShopperLocale returns the ShopperLocale field if non-nil, zero value otherwise.

### GetShopperLocaleOk

`func (o *PaymentDonationRequest) GetShopperLocaleOk() (*string, bool)`

GetShopperLocaleOk returns a tuple with the ShopperLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperLocale

`func (o *PaymentDonationRequest) SetShopperLocale(v string)`

SetShopperLocale sets ShopperLocale field to given value.

### HasShopperLocale

`func (o *PaymentDonationRequest) HasShopperLocale() bool`

HasShopperLocale returns a boolean if a field has been set.

### GetShopperName

`func (o *PaymentDonationRequest) GetShopperName() Name`

GetShopperName returns the ShopperName field if non-nil, zero value otherwise.

### GetShopperNameOk

`func (o *PaymentDonationRequest) GetShopperNameOk() (*Name, bool)`

GetShopperNameOk returns a tuple with the ShopperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperName

`func (o *PaymentDonationRequest) SetShopperName(v Name)`

SetShopperName sets ShopperName field to given value.

### HasShopperName

`func (o *PaymentDonationRequest) HasShopperName() bool`

HasShopperName returns a boolean if a field has been set.

### GetShopperReference

`func (o *PaymentDonationRequest) GetShopperReference() string`

GetShopperReference returns the ShopperReference field if non-nil, zero value otherwise.

### GetShopperReferenceOk

`func (o *PaymentDonationRequest) GetShopperReferenceOk() (*string, bool)`

GetShopperReferenceOk returns a tuple with the ShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperReference

`func (o *PaymentDonationRequest) SetShopperReference(v string)`

SetShopperReference sets ShopperReference field to given value.

### HasShopperReference

`func (o *PaymentDonationRequest) HasShopperReference() bool`

HasShopperReference returns a boolean if a field has been set.

### GetShopperStatement

`func (o *PaymentDonationRequest) GetShopperStatement() string`

GetShopperStatement returns the ShopperStatement field if non-nil, zero value otherwise.

### GetShopperStatementOk

`func (o *PaymentDonationRequest) GetShopperStatementOk() (*string, bool)`

GetShopperStatementOk returns a tuple with the ShopperStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperStatement

`func (o *PaymentDonationRequest) SetShopperStatement(v string)`

SetShopperStatement sets ShopperStatement field to given value.

### HasShopperStatement

`func (o *PaymentDonationRequest) HasShopperStatement() bool`

HasShopperStatement returns a boolean if a field has been set.

### GetSocialSecurityNumber

`func (o *PaymentDonationRequest) GetSocialSecurityNumber() string`

GetSocialSecurityNumber returns the SocialSecurityNumber field if non-nil, zero value otherwise.

### GetSocialSecurityNumberOk

`func (o *PaymentDonationRequest) GetSocialSecurityNumberOk() (*string, bool)`

GetSocialSecurityNumberOk returns a tuple with the SocialSecurityNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialSecurityNumber

`func (o *PaymentDonationRequest) SetSocialSecurityNumber(v string)`

SetSocialSecurityNumber sets SocialSecurityNumber field to given value.

### HasSocialSecurityNumber

`func (o *PaymentDonationRequest) HasSocialSecurityNumber() bool`

HasSocialSecurityNumber returns a boolean if a field has been set.

### GetSplits

`func (o *PaymentDonationRequest) GetSplits() []Split`

GetSplits returns the Splits field if non-nil, zero value otherwise.

### GetSplitsOk

`func (o *PaymentDonationRequest) GetSplitsOk() (*[]Split, bool)`

GetSplitsOk returns a tuple with the Splits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplits

`func (o *PaymentDonationRequest) SetSplits(v []Split)`

SetSplits sets Splits field to given value.

### HasSplits

`func (o *PaymentDonationRequest) HasSplits() bool`

HasSplits returns a boolean if a field has been set.

### GetStore

`func (o *PaymentDonationRequest) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *PaymentDonationRequest) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *PaymentDonationRequest) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *PaymentDonationRequest) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetStorePaymentMethod

`func (o *PaymentDonationRequest) GetStorePaymentMethod() bool`

GetStorePaymentMethod returns the StorePaymentMethod field if non-nil, zero value otherwise.

### GetStorePaymentMethodOk

`func (o *PaymentDonationRequest) GetStorePaymentMethodOk() (*bool, bool)`

GetStorePaymentMethodOk returns a tuple with the StorePaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePaymentMethod

`func (o *PaymentDonationRequest) SetStorePaymentMethod(v bool)`

SetStorePaymentMethod sets StorePaymentMethod field to given value.

### HasStorePaymentMethod

`func (o *PaymentDonationRequest) HasStorePaymentMethod() bool`

HasStorePaymentMethod returns a boolean if a field has been set.

### GetTelephoneNumber

`func (o *PaymentDonationRequest) GetTelephoneNumber() string`

GetTelephoneNumber returns the TelephoneNumber field if non-nil, zero value otherwise.

### GetTelephoneNumberOk

`func (o *PaymentDonationRequest) GetTelephoneNumberOk() (*string, bool)`

GetTelephoneNumberOk returns a tuple with the TelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumber

`func (o *PaymentDonationRequest) SetTelephoneNumber(v string)`

SetTelephoneNumber sets TelephoneNumber field to given value.

### HasTelephoneNumber

`func (o *PaymentDonationRequest) HasTelephoneNumber() bool`

HasTelephoneNumber returns a boolean if a field has been set.

### GetThreeDS2RequestData

`func (o *PaymentDonationRequest) GetThreeDS2RequestData() ThreeDS2RequestData`

GetThreeDS2RequestData returns the ThreeDS2RequestData field if non-nil, zero value otherwise.

### GetThreeDS2RequestDataOk

`func (o *PaymentDonationRequest) GetThreeDS2RequestDataOk() (*ThreeDS2RequestData, bool)`

GetThreeDS2RequestDataOk returns a tuple with the ThreeDS2RequestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDS2RequestData

`func (o *PaymentDonationRequest) SetThreeDS2RequestData(v ThreeDS2RequestData)`

SetThreeDS2RequestData sets ThreeDS2RequestData field to given value.

### HasThreeDS2RequestData

`func (o *PaymentDonationRequest) HasThreeDS2RequestData() bool`

HasThreeDS2RequestData returns a boolean if a field has been set.

### GetThreeDSAuthenticationOnly

`func (o *PaymentDonationRequest) GetThreeDSAuthenticationOnly() bool`

GetThreeDSAuthenticationOnly returns the ThreeDSAuthenticationOnly field if non-nil, zero value otherwise.

### GetThreeDSAuthenticationOnlyOk

`func (o *PaymentDonationRequest) GetThreeDSAuthenticationOnlyOk() (*bool, bool)`

GetThreeDSAuthenticationOnlyOk returns a tuple with the ThreeDSAuthenticationOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSAuthenticationOnly

`func (o *PaymentDonationRequest) SetThreeDSAuthenticationOnly(v bool)`

SetThreeDSAuthenticationOnly sets ThreeDSAuthenticationOnly field to given value.

### HasThreeDSAuthenticationOnly

`func (o *PaymentDonationRequest) HasThreeDSAuthenticationOnly() bool`

HasThreeDSAuthenticationOnly returns a boolean if a field has been set.

### GetTrustedShopper

`func (o *PaymentDonationRequest) GetTrustedShopper() bool`

GetTrustedShopper returns the TrustedShopper field if non-nil, zero value otherwise.

### GetTrustedShopperOk

`func (o *PaymentDonationRequest) GetTrustedShopperOk() (*bool, bool)`

GetTrustedShopperOk returns a tuple with the TrustedShopper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedShopper

`func (o *PaymentDonationRequest) SetTrustedShopper(v bool)`

SetTrustedShopper sets TrustedShopper field to given value.

### HasTrustedShopper

`func (o *PaymentDonationRequest) HasTrustedShopper() bool`

HasTrustedShopper returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


