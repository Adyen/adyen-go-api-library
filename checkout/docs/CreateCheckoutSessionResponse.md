# CreateCheckoutSessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountInfo** | Pointer to [**AccountInfo**](AccountInfo.md) |  | [optional] 
**AdditionalAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**AdditionalData** | Pointer to **map[string]string** | This field contains additional data, which may be required for a particular payment request.  The &#x60;additionalData&#x60; object consists of entries, each of which includes the key and value. | [optional] 
**AllowedPaymentMethods** | Pointer to **[]string** | List of payment methods to be presented to the shopper. To refer to payment methods, use their &#x60;paymentMethod.type&#x60;from [Payment methods overview](https://docs.adyen.com/payment-methods).  Example: &#x60;\&quot;allowedPaymentMethods\&quot;:[\&quot;ideal\&quot;,\&quot;giropay\&quot;]&#x60; | [optional] 
**Amount** | [**Amount**](Amount.md) |  | 
**ApplicationInfo** | Pointer to [**ApplicationInfo**](ApplicationInfo.md) |  | [optional] 
**AuthenticationData** | Pointer to [**AuthenticationData**](AuthenticationData.md) |  | [optional] 
**BillingAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**BlockedPaymentMethods** | Pointer to **[]string** | List of payment methods to be hidden from the shopper. To refer to payment methods, use their &#x60;paymentMethod.type&#x60;from [Payment methods overview](https://docs.adyen.com/payment-methods).  Example: &#x60;\&quot;blockedPaymentMethods\&quot;:[\&quot;ideal\&quot;,\&quot;giropay\&quot;]&#x60; | [optional] 
**CaptureDelayHours** | Pointer to **int32** | The delay between the authorisation and scheduled auto-capture, specified in hours. | [optional] 
**Channel** | Pointer to **string** | The platform where a payment transaction takes place. This field is optional for filtering out payment methods that are only available on specific platforms. If this value is not set, then we will try to infer it from the &#x60;sdkVersion&#x60; or &#x60;token&#x60;.  Possible values: * **iOS** * **Android** * **Web** | [optional] 
**Company** | Pointer to [**Company**](Company.md) |  | [optional] 
**CountryCode** | Pointer to **string** | The shopper&#39;s two-letter country code. | [optional] 
**DateOfBirth** | Pointer to **string** | The shopper&#39;s date of birth.  Format [ISO-8601](https://www.w3.org/TR/NOTE-datetime): YYYY-MM-DD | [optional] 
**DeliverAt** | Pointer to **time.Time** | The date and time when the purchased goods should be delivered.  [ISO 8601](https://www.w3.org/TR/NOTE-datetime) format: YYYY-MM-DDThh:mm:ss+TZD, for example, **2020-12-18T10:15:30+01:00**. | [optional] 
**DeliveryAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**EnableOneClick** | Pointer to **bool** | When true and &#x60;shopperReference&#x60; is provided, the shopper will be asked if the payment details should be stored for future one-click payments. | [optional] 
**EnablePayOut** | Pointer to **bool** | When true and &#x60;shopperReference&#x60; is provided, the payment details will be tokenized for payouts. | [optional] 
**EnableRecurring** | Pointer to **bool** | When true and &#x60;shopperReference&#x60; is provided, the payment details will be tokenized for recurring payments. | [optional] 
**ExpiresAt** | **time.Time** | The date the session expires in [ISO8601](https://www.iso.org/iso-8601-date-and-time-format.html) format. When not specified, the expiry date is set to 1 hour after session creation. You cannot set the session expiry to more than 24 hours after session creation. | 
**FundOrigin** | Pointer to [**FundOrigin**](FundOrigin.md) |  | [optional] 
**FundRecipient** | Pointer to [**FundRecipient**](FundRecipient.md) |  | [optional] 
**Id** | **string** | A unique identifier of the session. | [readonly] 
**InstallmentOptions** | Pointer to [**map[string]CheckoutSessionInstallmentOption**](CheckoutSessionInstallmentOption.md) | A set of key-value pairs that specifies the installment options available per payment method. The key must be a payment method name in lowercase. For example, **card** to specify installment options for all cards, or **visa** or **mc**. The value must be an object containing the installment options. | [optional] 
**LineItems** | Pointer to [**[]LineItem**](LineItem.md) | Price and product information about the purchased items, to be included on the invoice sent to the shopper. &gt; This field is required for 3x 4x Oney, Affirm, Afterpay, Clearpay, Klarna, Ratepay, and Zip. | [optional] 
**Mandate** | Pointer to [**Mandate**](Mandate.md) |  | [optional] 
**Mcc** | Pointer to **string** | The [merchant category code](https://en.wikipedia.org/wiki/Merchant_category_code) (MCC) is a four-digit number, which relates to a particular market segment. This code reflects the predominant activity that is conducted by the merchant. | [optional] 
**MerchantAccount** | **string** | The merchant account identifier, with which you want to process the transaction. | 
**MerchantOrderReference** | Pointer to **string** | This reference allows linking multiple transactions to each other for reporting purposes (i.e. order auth-rate). The reference should be unique per billing cycle. The same merchant order reference should never be reused after the first authorised attempt. If used, this field should be supplied for all incoming authorisations. &gt; We strongly recommend you send the &#x60;merchantOrderReference&#x60; value to benefit from linking payment requests when authorisation retries take place. In addition, we recommend you provide &#x60;retry.orderAttemptNumber&#x60;, &#x60;retry.chainAttemptNumber&#x60;, and &#x60;retry.skipRetry&#x60; values in &#x60;PaymentRequest.additionalData&#x60;. | [optional] 
**Metadata** | Pointer to **map[string]string** | Metadata consists of entries, each of which includes a key and a value. Limits: * Maximum 20 key-value pairs per request. * Maximum 20 characters per key. * Maximum 80 characters per value.  | [optional] 
**Mode** | Pointer to **string** | Indicates the type of front end integration. Possible values: * **embedded** (default): Drop-in or Components integration * **hosted**: Hosted Checkout integration | [optional] [default to "embedded"]
**MpiData** | Pointer to [**ThreeDSecureData**](ThreeDSecureData.md) |  | [optional] 
**RecurringExpiry** | Pointer to **string** | Date after which no further authorisations shall be performed. Only for 3D Secure 2. | [optional] 
**RecurringFrequency** | Pointer to **string** | Minimum number of days between authorisations. Only for 3D Secure 2. | [optional] 
**RecurringProcessingModel** | Pointer to **string** | Defines a recurring payment type. Required when creating a token to store payment details. Allowed values: * &#x60;Subscription&#x60; – A transaction for a fixed or variable amount, which follows a fixed schedule. * &#x60;CardOnFile&#x60; – With a card-on-file (CoF) transaction, card details are stored to enable one-click or omnichannel journeys, or simply to streamline the checkout process. Any subscription not following a fixed schedule is also considered a card-on-file transaction. * &#x60;UnscheduledCardOnFile&#x60; – An unscheduled card-on-file (UCoF) transaction is a transaction that occurs on a non-fixed schedule and/or have variable amounts. For example, automatic top-ups when a cardholder&#39;s balance drops below a certain amount.  | [optional] 
**RedirectFromIssuerMethod** | Pointer to **string** | Specifies the redirect method (GET or POST) when redirecting back from the issuer. | [optional] 
**RedirectToIssuerMethod** | Pointer to **string** | Specifies the redirect method (GET or POST) when redirecting to the issuer. | [optional] 
**Reference** | **string** | The reference to uniquely identify a payment. | 
**ReturnUrl** | **string** | The URL to return to when a redirect payment is completed. | 
**RiskData** | Pointer to [**RiskData**](RiskData.md) |  | [optional] 
**SessionData** | Pointer to **string** | The payment session data you need to pass to your front end. | [optional] 
**ShopperEmail** | Pointer to **string** | The shopper&#39;s email address. | [optional] 
**ShopperIP** | Pointer to **string** | The shopper&#39;s IP address. In general, we recommend that you provide this data, as it is used in a number of risk checks (for instance, number of payment attempts or location-based checks). &gt; For 3D Secure 2 transactions, schemes require &#x60;shopperIP&#x60; for all browser-based implementations. This field is also mandatory for some merchants depending on your business model. For more information, [contact Support](https://www.adyen.help/hc/en-us/requests/new). | [optional] 
**ShopperInteraction** | Pointer to **string** | Specifies the sales channel, through which the shopper gives their card details, and whether the shopper is a returning customer. For the web service API, Adyen assumes Ecommerce shopper interaction by default.  This field has the following possible values: * &#x60;Ecommerce&#x60; - Online transactions where the cardholder is present (online). For better authorisation rates, we recommend sending the card security code (CSC) along with the request. * &#x60;ContAuth&#x60; - Card on file and/or subscription transactions, where the cardholder is known to the merchant (returning customer). If the shopper is present (online), you can supply also the CSC to improve authorisation (one-click payment). * &#x60;Moto&#x60; - Mail-order and telephone-order transactions where the shopper is in contact with the merchant via email or telephone. * &#x60;POS&#x60; - Point-of-sale transactions where the shopper is physically present to make a payment using a secure payment terminal. | [optional] 
**ShopperLocale** | Pointer to **string** | The combination of a language code and a country code to specify the language to be used in the payment. | [optional] 
**ShopperName** | Pointer to [**Name**](Name.md) |  | [optional] 
**ShopperReference** | Pointer to **string** | Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. &gt; Your reference must not include personally identifiable information (PII), for example name or email address. | [optional] 
**ShopperStatement** | Pointer to **string** | The text to be shown on the shopper&#39;s bank statement.  We recommend sending a maximum of 22 characters, otherwise banks might truncate the string.  Allowed characters: **a-z**, **A-Z**, **0-9**, spaces, and special characters **. , &#39; _ - ? + * /_**. | [optional] 
**SocialSecurityNumber** | Pointer to **string** | The shopper&#39;s social security number. | [optional] 
**SplitCardFundingSources** | Pointer to **bool** | Boolean value indicating whether the card payment method should be split into separate debit and credit options. | [optional] [default to false]
**Splits** | Pointer to [**[]Split**](Split.md) | An array of objects specifying how the payment should be split when using [Adyen for Platforms](https://docs.adyen.com/platforms/processing-payments#providing-split-information) or [Issuing](https://docs.adyen.com/issuing/manage-funds#split). | [optional] 
**Store** | Pointer to **string** | The ecommerce or point-of-sale store that is processing the payment. | [optional] 
**StorePaymentMethod** | Pointer to **bool** | When this is set to **true** and the &#x60;shopperReference&#x60; is provided, the payment details will be stored. | [optional] 
**StorePaymentMethodMode** | Pointer to **string** | Indicates if the details of the payment method will be stored for the shopper. Possible values: * **disabled** – No details will be stored (default). * **askForConsent** – If the &#x60;shopperReference&#x60; is provided, the UI lets the shopper choose if they want their payment details to be stored. * **enabled** – If the &#x60;shopperReference&#x60; is provided, the details will be stored without asking the shopper for consent. | [optional] 
**TelephoneNumber** | Pointer to **string** | The shopper&#39;s telephone number. | [optional] 
**ThreeDSAuthenticationOnly** | Pointer to **bool** | If set to true, you will only perform the [3D Secure 2 authentication](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only), and not the payment authorisation. | [optional] [default to false]
**TrustedShopper** | Pointer to **bool** | Set to true if the payment should be routed to a trusted MID. | [optional] 

## Methods

### NewCreateCheckoutSessionResponse

`func NewCreateCheckoutSessionResponse(amount Amount, expiresAt time.Time, id string, merchantAccount string, reference string, returnUrl string, ) *CreateCheckoutSessionResponse`

NewCreateCheckoutSessionResponse instantiates a new CreateCheckoutSessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCheckoutSessionResponseWithDefaults

`func NewCreateCheckoutSessionResponseWithDefaults() *CreateCheckoutSessionResponse`

NewCreateCheckoutSessionResponseWithDefaults instantiates a new CreateCheckoutSessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountInfo

`func (o *CreateCheckoutSessionResponse) GetAccountInfo() AccountInfo`

GetAccountInfo returns the AccountInfo field if non-nil, zero value otherwise.

### GetAccountInfoOk

`func (o *CreateCheckoutSessionResponse) GetAccountInfoOk() (*AccountInfo, bool)`

GetAccountInfoOk returns a tuple with the AccountInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountInfo

`func (o *CreateCheckoutSessionResponse) SetAccountInfo(v AccountInfo)`

SetAccountInfo sets AccountInfo field to given value.

### HasAccountInfo

`func (o *CreateCheckoutSessionResponse) HasAccountInfo() bool`

HasAccountInfo returns a boolean if a field has been set.

### GetAdditionalAmount

`func (o *CreateCheckoutSessionResponse) GetAdditionalAmount() Amount`

GetAdditionalAmount returns the AdditionalAmount field if non-nil, zero value otherwise.

### GetAdditionalAmountOk

`func (o *CreateCheckoutSessionResponse) GetAdditionalAmountOk() (*Amount, bool)`

GetAdditionalAmountOk returns a tuple with the AdditionalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAmount

`func (o *CreateCheckoutSessionResponse) SetAdditionalAmount(v Amount)`

SetAdditionalAmount sets AdditionalAmount field to given value.

### HasAdditionalAmount

`func (o *CreateCheckoutSessionResponse) HasAdditionalAmount() bool`

HasAdditionalAmount returns a boolean if a field has been set.

### GetAdditionalData

`func (o *CreateCheckoutSessionResponse) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *CreateCheckoutSessionResponse) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *CreateCheckoutSessionResponse) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *CreateCheckoutSessionResponse) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetAllowedPaymentMethods

`func (o *CreateCheckoutSessionResponse) GetAllowedPaymentMethods() []string`

GetAllowedPaymentMethods returns the AllowedPaymentMethods field if non-nil, zero value otherwise.

### GetAllowedPaymentMethodsOk

`func (o *CreateCheckoutSessionResponse) GetAllowedPaymentMethodsOk() (*[]string, bool)`

GetAllowedPaymentMethodsOk returns a tuple with the AllowedPaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPaymentMethods

`func (o *CreateCheckoutSessionResponse) SetAllowedPaymentMethods(v []string)`

SetAllowedPaymentMethods sets AllowedPaymentMethods field to given value.

### HasAllowedPaymentMethods

`func (o *CreateCheckoutSessionResponse) HasAllowedPaymentMethods() bool`

HasAllowedPaymentMethods returns a boolean if a field has been set.

### GetAmount

`func (o *CreateCheckoutSessionResponse) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateCheckoutSessionResponse) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateCheckoutSessionResponse) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetApplicationInfo

`func (o *CreateCheckoutSessionResponse) GetApplicationInfo() ApplicationInfo`

GetApplicationInfo returns the ApplicationInfo field if non-nil, zero value otherwise.

### GetApplicationInfoOk

`func (o *CreateCheckoutSessionResponse) GetApplicationInfoOk() (*ApplicationInfo, bool)`

GetApplicationInfoOk returns a tuple with the ApplicationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationInfo

`func (o *CreateCheckoutSessionResponse) SetApplicationInfo(v ApplicationInfo)`

SetApplicationInfo sets ApplicationInfo field to given value.

### HasApplicationInfo

`func (o *CreateCheckoutSessionResponse) HasApplicationInfo() bool`

HasApplicationInfo returns a boolean if a field has been set.

### GetAuthenticationData

`func (o *CreateCheckoutSessionResponse) GetAuthenticationData() AuthenticationData`

GetAuthenticationData returns the AuthenticationData field if non-nil, zero value otherwise.

### GetAuthenticationDataOk

`func (o *CreateCheckoutSessionResponse) GetAuthenticationDataOk() (*AuthenticationData, bool)`

GetAuthenticationDataOk returns a tuple with the AuthenticationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationData

`func (o *CreateCheckoutSessionResponse) SetAuthenticationData(v AuthenticationData)`

SetAuthenticationData sets AuthenticationData field to given value.

### HasAuthenticationData

`func (o *CreateCheckoutSessionResponse) HasAuthenticationData() bool`

HasAuthenticationData returns a boolean if a field has been set.

### GetBillingAddress

`func (o *CreateCheckoutSessionResponse) GetBillingAddress() Address`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *CreateCheckoutSessionResponse) GetBillingAddressOk() (*Address, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *CreateCheckoutSessionResponse) SetBillingAddress(v Address)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *CreateCheckoutSessionResponse) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetBlockedPaymentMethods

`func (o *CreateCheckoutSessionResponse) GetBlockedPaymentMethods() []string`

GetBlockedPaymentMethods returns the BlockedPaymentMethods field if non-nil, zero value otherwise.

### GetBlockedPaymentMethodsOk

`func (o *CreateCheckoutSessionResponse) GetBlockedPaymentMethodsOk() (*[]string, bool)`

GetBlockedPaymentMethodsOk returns a tuple with the BlockedPaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedPaymentMethods

`func (o *CreateCheckoutSessionResponse) SetBlockedPaymentMethods(v []string)`

SetBlockedPaymentMethods sets BlockedPaymentMethods field to given value.

### HasBlockedPaymentMethods

`func (o *CreateCheckoutSessionResponse) HasBlockedPaymentMethods() bool`

HasBlockedPaymentMethods returns a boolean if a field has been set.

### GetCaptureDelayHours

`func (o *CreateCheckoutSessionResponse) GetCaptureDelayHours() int32`

GetCaptureDelayHours returns the CaptureDelayHours field if non-nil, zero value otherwise.

### GetCaptureDelayHoursOk

`func (o *CreateCheckoutSessionResponse) GetCaptureDelayHoursOk() (*int32, bool)`

GetCaptureDelayHoursOk returns a tuple with the CaptureDelayHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureDelayHours

`func (o *CreateCheckoutSessionResponse) SetCaptureDelayHours(v int32)`

SetCaptureDelayHours sets CaptureDelayHours field to given value.

### HasCaptureDelayHours

`func (o *CreateCheckoutSessionResponse) HasCaptureDelayHours() bool`

HasCaptureDelayHours returns a boolean if a field has been set.

### GetChannel

`func (o *CreateCheckoutSessionResponse) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CreateCheckoutSessionResponse) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CreateCheckoutSessionResponse) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *CreateCheckoutSessionResponse) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCompany

`func (o *CreateCheckoutSessionResponse) GetCompany() Company`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CreateCheckoutSessionResponse) GetCompanyOk() (*Company, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CreateCheckoutSessionResponse) SetCompany(v Company)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CreateCheckoutSessionResponse) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCountryCode

`func (o *CreateCheckoutSessionResponse) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CreateCheckoutSessionResponse) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CreateCheckoutSessionResponse) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *CreateCheckoutSessionResponse) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *CreateCheckoutSessionResponse) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *CreateCheckoutSessionResponse) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *CreateCheckoutSessionResponse) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *CreateCheckoutSessionResponse) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetDeliverAt

`func (o *CreateCheckoutSessionResponse) GetDeliverAt() time.Time`

GetDeliverAt returns the DeliverAt field if non-nil, zero value otherwise.

### GetDeliverAtOk

`func (o *CreateCheckoutSessionResponse) GetDeliverAtOk() (*time.Time, bool)`

GetDeliverAtOk returns a tuple with the DeliverAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverAt

`func (o *CreateCheckoutSessionResponse) SetDeliverAt(v time.Time)`

SetDeliverAt sets DeliverAt field to given value.

### HasDeliverAt

`func (o *CreateCheckoutSessionResponse) HasDeliverAt() bool`

HasDeliverAt returns a boolean if a field has been set.

### GetDeliveryAddress

`func (o *CreateCheckoutSessionResponse) GetDeliveryAddress() Address`

GetDeliveryAddress returns the DeliveryAddress field if non-nil, zero value otherwise.

### GetDeliveryAddressOk

`func (o *CreateCheckoutSessionResponse) GetDeliveryAddressOk() (*Address, bool)`

GetDeliveryAddressOk returns a tuple with the DeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddress

`func (o *CreateCheckoutSessionResponse) SetDeliveryAddress(v Address)`

SetDeliveryAddress sets DeliveryAddress field to given value.

### HasDeliveryAddress

`func (o *CreateCheckoutSessionResponse) HasDeliveryAddress() bool`

HasDeliveryAddress returns a boolean if a field has been set.

### GetEnableOneClick

`func (o *CreateCheckoutSessionResponse) GetEnableOneClick() bool`

GetEnableOneClick returns the EnableOneClick field if non-nil, zero value otherwise.

### GetEnableOneClickOk

`func (o *CreateCheckoutSessionResponse) GetEnableOneClickOk() (*bool, bool)`

GetEnableOneClickOk returns a tuple with the EnableOneClick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOneClick

`func (o *CreateCheckoutSessionResponse) SetEnableOneClick(v bool)`

SetEnableOneClick sets EnableOneClick field to given value.

### HasEnableOneClick

`func (o *CreateCheckoutSessionResponse) HasEnableOneClick() bool`

HasEnableOneClick returns a boolean if a field has been set.

### GetEnablePayOut

`func (o *CreateCheckoutSessionResponse) GetEnablePayOut() bool`

GetEnablePayOut returns the EnablePayOut field if non-nil, zero value otherwise.

### GetEnablePayOutOk

`func (o *CreateCheckoutSessionResponse) GetEnablePayOutOk() (*bool, bool)`

GetEnablePayOutOk returns a tuple with the EnablePayOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePayOut

`func (o *CreateCheckoutSessionResponse) SetEnablePayOut(v bool)`

SetEnablePayOut sets EnablePayOut field to given value.

### HasEnablePayOut

`func (o *CreateCheckoutSessionResponse) HasEnablePayOut() bool`

HasEnablePayOut returns a boolean if a field has been set.

### GetEnableRecurring

`func (o *CreateCheckoutSessionResponse) GetEnableRecurring() bool`

GetEnableRecurring returns the EnableRecurring field if non-nil, zero value otherwise.

### GetEnableRecurringOk

`func (o *CreateCheckoutSessionResponse) GetEnableRecurringOk() (*bool, bool)`

GetEnableRecurringOk returns a tuple with the EnableRecurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRecurring

`func (o *CreateCheckoutSessionResponse) SetEnableRecurring(v bool)`

SetEnableRecurring sets EnableRecurring field to given value.

### HasEnableRecurring

`func (o *CreateCheckoutSessionResponse) HasEnableRecurring() bool`

HasEnableRecurring returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CreateCheckoutSessionResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateCheckoutSessionResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateCheckoutSessionResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetFundOrigin

`func (o *CreateCheckoutSessionResponse) GetFundOrigin() FundOrigin`

GetFundOrigin returns the FundOrigin field if non-nil, zero value otherwise.

### GetFundOriginOk

`func (o *CreateCheckoutSessionResponse) GetFundOriginOk() (*FundOrigin, bool)`

GetFundOriginOk returns a tuple with the FundOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundOrigin

`func (o *CreateCheckoutSessionResponse) SetFundOrigin(v FundOrigin)`

SetFundOrigin sets FundOrigin field to given value.

### HasFundOrigin

`func (o *CreateCheckoutSessionResponse) HasFundOrigin() bool`

HasFundOrigin returns a boolean if a field has been set.

### GetFundRecipient

`func (o *CreateCheckoutSessionResponse) GetFundRecipient() FundRecipient`

GetFundRecipient returns the FundRecipient field if non-nil, zero value otherwise.

### GetFundRecipientOk

`func (o *CreateCheckoutSessionResponse) GetFundRecipientOk() (*FundRecipient, bool)`

GetFundRecipientOk returns a tuple with the FundRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundRecipient

`func (o *CreateCheckoutSessionResponse) SetFundRecipient(v FundRecipient)`

SetFundRecipient sets FundRecipient field to given value.

### HasFundRecipient

`func (o *CreateCheckoutSessionResponse) HasFundRecipient() bool`

HasFundRecipient returns a boolean if a field has been set.

### GetId

`func (o *CreateCheckoutSessionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateCheckoutSessionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateCheckoutSessionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetInstallmentOptions

`func (o *CreateCheckoutSessionResponse) GetInstallmentOptions() map[string]CheckoutSessionInstallmentOption`

GetInstallmentOptions returns the InstallmentOptions field if non-nil, zero value otherwise.

### GetInstallmentOptionsOk

`func (o *CreateCheckoutSessionResponse) GetInstallmentOptionsOk() (*map[string]CheckoutSessionInstallmentOption, bool)`

GetInstallmentOptionsOk returns a tuple with the InstallmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallmentOptions

`func (o *CreateCheckoutSessionResponse) SetInstallmentOptions(v map[string]CheckoutSessionInstallmentOption)`

SetInstallmentOptions sets InstallmentOptions field to given value.

### HasInstallmentOptions

`func (o *CreateCheckoutSessionResponse) HasInstallmentOptions() bool`

HasInstallmentOptions returns a boolean if a field has been set.

### GetLineItems

`func (o *CreateCheckoutSessionResponse) GetLineItems() []LineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *CreateCheckoutSessionResponse) GetLineItemsOk() (*[]LineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *CreateCheckoutSessionResponse) SetLineItems(v []LineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *CreateCheckoutSessionResponse) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetMandate

`func (o *CreateCheckoutSessionResponse) GetMandate() Mandate`

GetMandate returns the Mandate field if non-nil, zero value otherwise.

### GetMandateOk

`func (o *CreateCheckoutSessionResponse) GetMandateOk() (*Mandate, bool)`

GetMandateOk returns a tuple with the Mandate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandate

`func (o *CreateCheckoutSessionResponse) SetMandate(v Mandate)`

SetMandate sets Mandate field to given value.

### HasMandate

`func (o *CreateCheckoutSessionResponse) HasMandate() bool`

HasMandate returns a boolean if a field has been set.

### GetMcc

`func (o *CreateCheckoutSessionResponse) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *CreateCheckoutSessionResponse) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *CreateCheckoutSessionResponse) SetMcc(v string)`

SetMcc sets Mcc field to given value.

### HasMcc

`func (o *CreateCheckoutSessionResponse) HasMcc() bool`

HasMcc returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *CreateCheckoutSessionResponse) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *CreateCheckoutSessionResponse) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *CreateCheckoutSessionResponse) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetMerchantOrderReference

`func (o *CreateCheckoutSessionResponse) GetMerchantOrderReference() string`

GetMerchantOrderReference returns the MerchantOrderReference field if non-nil, zero value otherwise.

### GetMerchantOrderReferenceOk

`func (o *CreateCheckoutSessionResponse) GetMerchantOrderReferenceOk() (*string, bool)`

GetMerchantOrderReferenceOk returns a tuple with the MerchantOrderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderReference

`func (o *CreateCheckoutSessionResponse) SetMerchantOrderReference(v string)`

SetMerchantOrderReference sets MerchantOrderReference field to given value.

### HasMerchantOrderReference

`func (o *CreateCheckoutSessionResponse) HasMerchantOrderReference() bool`

HasMerchantOrderReference returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateCheckoutSessionResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateCheckoutSessionResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateCheckoutSessionResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateCheckoutSessionResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMode

`func (o *CreateCheckoutSessionResponse) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CreateCheckoutSessionResponse) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CreateCheckoutSessionResponse) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *CreateCheckoutSessionResponse) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetMpiData

`func (o *CreateCheckoutSessionResponse) GetMpiData() ThreeDSecureData`

GetMpiData returns the MpiData field if non-nil, zero value otherwise.

### GetMpiDataOk

`func (o *CreateCheckoutSessionResponse) GetMpiDataOk() (*ThreeDSecureData, bool)`

GetMpiDataOk returns a tuple with the MpiData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpiData

`func (o *CreateCheckoutSessionResponse) SetMpiData(v ThreeDSecureData)`

SetMpiData sets MpiData field to given value.

### HasMpiData

`func (o *CreateCheckoutSessionResponse) HasMpiData() bool`

HasMpiData returns a boolean if a field has been set.

### GetRecurringExpiry

`func (o *CreateCheckoutSessionResponse) GetRecurringExpiry() string`

GetRecurringExpiry returns the RecurringExpiry field if non-nil, zero value otherwise.

### GetRecurringExpiryOk

`func (o *CreateCheckoutSessionResponse) GetRecurringExpiryOk() (*string, bool)`

GetRecurringExpiryOk returns a tuple with the RecurringExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringExpiry

`func (o *CreateCheckoutSessionResponse) SetRecurringExpiry(v string)`

SetRecurringExpiry sets RecurringExpiry field to given value.

### HasRecurringExpiry

`func (o *CreateCheckoutSessionResponse) HasRecurringExpiry() bool`

HasRecurringExpiry returns a boolean if a field has been set.

### GetRecurringFrequency

`func (o *CreateCheckoutSessionResponse) GetRecurringFrequency() string`

GetRecurringFrequency returns the RecurringFrequency field if non-nil, zero value otherwise.

### GetRecurringFrequencyOk

`func (o *CreateCheckoutSessionResponse) GetRecurringFrequencyOk() (*string, bool)`

GetRecurringFrequencyOk returns a tuple with the RecurringFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringFrequency

`func (o *CreateCheckoutSessionResponse) SetRecurringFrequency(v string)`

SetRecurringFrequency sets RecurringFrequency field to given value.

### HasRecurringFrequency

`func (o *CreateCheckoutSessionResponse) HasRecurringFrequency() bool`

HasRecurringFrequency returns a boolean if a field has been set.

### GetRecurringProcessingModel

`func (o *CreateCheckoutSessionResponse) GetRecurringProcessingModel() string`

GetRecurringProcessingModel returns the RecurringProcessingModel field if non-nil, zero value otherwise.

### GetRecurringProcessingModelOk

`func (o *CreateCheckoutSessionResponse) GetRecurringProcessingModelOk() (*string, bool)`

GetRecurringProcessingModelOk returns a tuple with the RecurringProcessingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringProcessingModel

`func (o *CreateCheckoutSessionResponse) SetRecurringProcessingModel(v string)`

SetRecurringProcessingModel sets RecurringProcessingModel field to given value.

### HasRecurringProcessingModel

`func (o *CreateCheckoutSessionResponse) HasRecurringProcessingModel() bool`

HasRecurringProcessingModel returns a boolean if a field has been set.

### GetRedirectFromIssuerMethod

`func (o *CreateCheckoutSessionResponse) GetRedirectFromIssuerMethod() string`

GetRedirectFromIssuerMethod returns the RedirectFromIssuerMethod field if non-nil, zero value otherwise.

### GetRedirectFromIssuerMethodOk

`func (o *CreateCheckoutSessionResponse) GetRedirectFromIssuerMethodOk() (*string, bool)`

GetRedirectFromIssuerMethodOk returns a tuple with the RedirectFromIssuerMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectFromIssuerMethod

`func (o *CreateCheckoutSessionResponse) SetRedirectFromIssuerMethod(v string)`

SetRedirectFromIssuerMethod sets RedirectFromIssuerMethod field to given value.

### HasRedirectFromIssuerMethod

`func (o *CreateCheckoutSessionResponse) HasRedirectFromIssuerMethod() bool`

HasRedirectFromIssuerMethod returns a boolean if a field has been set.

### GetRedirectToIssuerMethod

`func (o *CreateCheckoutSessionResponse) GetRedirectToIssuerMethod() string`

GetRedirectToIssuerMethod returns the RedirectToIssuerMethod field if non-nil, zero value otherwise.

### GetRedirectToIssuerMethodOk

`func (o *CreateCheckoutSessionResponse) GetRedirectToIssuerMethodOk() (*string, bool)`

GetRedirectToIssuerMethodOk returns a tuple with the RedirectToIssuerMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectToIssuerMethod

`func (o *CreateCheckoutSessionResponse) SetRedirectToIssuerMethod(v string)`

SetRedirectToIssuerMethod sets RedirectToIssuerMethod field to given value.

### HasRedirectToIssuerMethod

`func (o *CreateCheckoutSessionResponse) HasRedirectToIssuerMethod() bool`

HasRedirectToIssuerMethod returns a boolean if a field has been set.

### GetReference

`func (o *CreateCheckoutSessionResponse) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CreateCheckoutSessionResponse) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CreateCheckoutSessionResponse) SetReference(v string)`

SetReference sets Reference field to given value.


### GetReturnUrl

`func (o *CreateCheckoutSessionResponse) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *CreateCheckoutSessionResponse) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *CreateCheckoutSessionResponse) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.


### GetRiskData

`func (o *CreateCheckoutSessionResponse) GetRiskData() RiskData`

GetRiskData returns the RiskData field if non-nil, zero value otherwise.

### GetRiskDataOk

`func (o *CreateCheckoutSessionResponse) GetRiskDataOk() (*RiskData, bool)`

GetRiskDataOk returns a tuple with the RiskData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskData

`func (o *CreateCheckoutSessionResponse) SetRiskData(v RiskData)`

SetRiskData sets RiskData field to given value.

### HasRiskData

`func (o *CreateCheckoutSessionResponse) HasRiskData() bool`

HasRiskData returns a boolean if a field has been set.

### GetSessionData

`func (o *CreateCheckoutSessionResponse) GetSessionData() string`

GetSessionData returns the SessionData field if non-nil, zero value otherwise.

### GetSessionDataOk

`func (o *CreateCheckoutSessionResponse) GetSessionDataOk() (*string, bool)`

GetSessionDataOk returns a tuple with the SessionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionData

`func (o *CreateCheckoutSessionResponse) SetSessionData(v string)`

SetSessionData sets SessionData field to given value.

### HasSessionData

`func (o *CreateCheckoutSessionResponse) HasSessionData() bool`

HasSessionData returns a boolean if a field has been set.

### GetShopperEmail

`func (o *CreateCheckoutSessionResponse) GetShopperEmail() string`

GetShopperEmail returns the ShopperEmail field if non-nil, zero value otherwise.

### GetShopperEmailOk

`func (o *CreateCheckoutSessionResponse) GetShopperEmailOk() (*string, bool)`

GetShopperEmailOk returns a tuple with the ShopperEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperEmail

`func (o *CreateCheckoutSessionResponse) SetShopperEmail(v string)`

SetShopperEmail sets ShopperEmail field to given value.

### HasShopperEmail

`func (o *CreateCheckoutSessionResponse) HasShopperEmail() bool`

HasShopperEmail returns a boolean if a field has been set.

### GetShopperIP

`func (o *CreateCheckoutSessionResponse) GetShopperIP() string`

GetShopperIP returns the ShopperIP field if non-nil, zero value otherwise.

### GetShopperIPOk

`func (o *CreateCheckoutSessionResponse) GetShopperIPOk() (*string, bool)`

GetShopperIPOk returns a tuple with the ShopperIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperIP

`func (o *CreateCheckoutSessionResponse) SetShopperIP(v string)`

SetShopperIP sets ShopperIP field to given value.

### HasShopperIP

`func (o *CreateCheckoutSessionResponse) HasShopperIP() bool`

HasShopperIP returns a boolean if a field has been set.

### GetShopperInteraction

`func (o *CreateCheckoutSessionResponse) GetShopperInteraction() string`

GetShopperInteraction returns the ShopperInteraction field if non-nil, zero value otherwise.

### GetShopperInteractionOk

`func (o *CreateCheckoutSessionResponse) GetShopperInteractionOk() (*string, bool)`

GetShopperInteractionOk returns a tuple with the ShopperInteraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperInteraction

`func (o *CreateCheckoutSessionResponse) SetShopperInteraction(v string)`

SetShopperInteraction sets ShopperInteraction field to given value.

### HasShopperInteraction

`func (o *CreateCheckoutSessionResponse) HasShopperInteraction() bool`

HasShopperInteraction returns a boolean if a field has been set.

### GetShopperLocale

`func (o *CreateCheckoutSessionResponse) GetShopperLocale() string`

GetShopperLocale returns the ShopperLocale field if non-nil, zero value otherwise.

### GetShopperLocaleOk

`func (o *CreateCheckoutSessionResponse) GetShopperLocaleOk() (*string, bool)`

GetShopperLocaleOk returns a tuple with the ShopperLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperLocale

`func (o *CreateCheckoutSessionResponse) SetShopperLocale(v string)`

SetShopperLocale sets ShopperLocale field to given value.

### HasShopperLocale

`func (o *CreateCheckoutSessionResponse) HasShopperLocale() bool`

HasShopperLocale returns a boolean if a field has been set.

### GetShopperName

`func (o *CreateCheckoutSessionResponse) GetShopperName() Name`

GetShopperName returns the ShopperName field if non-nil, zero value otherwise.

### GetShopperNameOk

`func (o *CreateCheckoutSessionResponse) GetShopperNameOk() (*Name, bool)`

GetShopperNameOk returns a tuple with the ShopperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperName

`func (o *CreateCheckoutSessionResponse) SetShopperName(v Name)`

SetShopperName sets ShopperName field to given value.

### HasShopperName

`func (o *CreateCheckoutSessionResponse) HasShopperName() bool`

HasShopperName returns a boolean if a field has been set.

### GetShopperReference

`func (o *CreateCheckoutSessionResponse) GetShopperReference() string`

GetShopperReference returns the ShopperReference field if non-nil, zero value otherwise.

### GetShopperReferenceOk

`func (o *CreateCheckoutSessionResponse) GetShopperReferenceOk() (*string, bool)`

GetShopperReferenceOk returns a tuple with the ShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperReference

`func (o *CreateCheckoutSessionResponse) SetShopperReference(v string)`

SetShopperReference sets ShopperReference field to given value.

### HasShopperReference

`func (o *CreateCheckoutSessionResponse) HasShopperReference() bool`

HasShopperReference returns a boolean if a field has been set.

### GetShopperStatement

`func (o *CreateCheckoutSessionResponse) GetShopperStatement() string`

GetShopperStatement returns the ShopperStatement field if non-nil, zero value otherwise.

### GetShopperStatementOk

`func (o *CreateCheckoutSessionResponse) GetShopperStatementOk() (*string, bool)`

GetShopperStatementOk returns a tuple with the ShopperStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperStatement

`func (o *CreateCheckoutSessionResponse) SetShopperStatement(v string)`

SetShopperStatement sets ShopperStatement field to given value.

### HasShopperStatement

`func (o *CreateCheckoutSessionResponse) HasShopperStatement() bool`

HasShopperStatement returns a boolean if a field has been set.

### GetSocialSecurityNumber

`func (o *CreateCheckoutSessionResponse) GetSocialSecurityNumber() string`

GetSocialSecurityNumber returns the SocialSecurityNumber field if non-nil, zero value otherwise.

### GetSocialSecurityNumberOk

`func (o *CreateCheckoutSessionResponse) GetSocialSecurityNumberOk() (*string, bool)`

GetSocialSecurityNumberOk returns a tuple with the SocialSecurityNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialSecurityNumber

`func (o *CreateCheckoutSessionResponse) SetSocialSecurityNumber(v string)`

SetSocialSecurityNumber sets SocialSecurityNumber field to given value.

### HasSocialSecurityNumber

`func (o *CreateCheckoutSessionResponse) HasSocialSecurityNumber() bool`

HasSocialSecurityNumber returns a boolean if a field has been set.

### GetSplitCardFundingSources

`func (o *CreateCheckoutSessionResponse) GetSplitCardFundingSources() bool`

GetSplitCardFundingSources returns the SplitCardFundingSources field if non-nil, zero value otherwise.

### GetSplitCardFundingSourcesOk

`func (o *CreateCheckoutSessionResponse) GetSplitCardFundingSourcesOk() (*bool, bool)`

GetSplitCardFundingSourcesOk returns a tuple with the SplitCardFundingSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitCardFundingSources

`func (o *CreateCheckoutSessionResponse) SetSplitCardFundingSources(v bool)`

SetSplitCardFundingSources sets SplitCardFundingSources field to given value.

### HasSplitCardFundingSources

`func (o *CreateCheckoutSessionResponse) HasSplitCardFundingSources() bool`

HasSplitCardFundingSources returns a boolean if a field has been set.

### GetSplits

`func (o *CreateCheckoutSessionResponse) GetSplits() []Split`

GetSplits returns the Splits field if non-nil, zero value otherwise.

### GetSplitsOk

`func (o *CreateCheckoutSessionResponse) GetSplitsOk() (*[]Split, bool)`

GetSplitsOk returns a tuple with the Splits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplits

`func (o *CreateCheckoutSessionResponse) SetSplits(v []Split)`

SetSplits sets Splits field to given value.

### HasSplits

`func (o *CreateCheckoutSessionResponse) HasSplits() bool`

HasSplits returns a boolean if a field has been set.

### GetStore

`func (o *CreateCheckoutSessionResponse) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *CreateCheckoutSessionResponse) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *CreateCheckoutSessionResponse) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *CreateCheckoutSessionResponse) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetStorePaymentMethod

`func (o *CreateCheckoutSessionResponse) GetStorePaymentMethod() bool`

GetStorePaymentMethod returns the StorePaymentMethod field if non-nil, zero value otherwise.

### GetStorePaymentMethodOk

`func (o *CreateCheckoutSessionResponse) GetStorePaymentMethodOk() (*bool, bool)`

GetStorePaymentMethodOk returns a tuple with the StorePaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePaymentMethod

`func (o *CreateCheckoutSessionResponse) SetStorePaymentMethod(v bool)`

SetStorePaymentMethod sets StorePaymentMethod field to given value.

### HasStorePaymentMethod

`func (o *CreateCheckoutSessionResponse) HasStorePaymentMethod() bool`

HasStorePaymentMethod returns a boolean if a field has been set.

### GetStorePaymentMethodMode

`func (o *CreateCheckoutSessionResponse) GetStorePaymentMethodMode() string`

GetStorePaymentMethodMode returns the StorePaymentMethodMode field if non-nil, zero value otherwise.

### GetStorePaymentMethodModeOk

`func (o *CreateCheckoutSessionResponse) GetStorePaymentMethodModeOk() (*string, bool)`

GetStorePaymentMethodModeOk returns a tuple with the StorePaymentMethodMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePaymentMethodMode

`func (o *CreateCheckoutSessionResponse) SetStorePaymentMethodMode(v string)`

SetStorePaymentMethodMode sets StorePaymentMethodMode field to given value.

### HasStorePaymentMethodMode

`func (o *CreateCheckoutSessionResponse) HasStorePaymentMethodMode() bool`

HasStorePaymentMethodMode returns a boolean if a field has been set.

### GetTelephoneNumber

`func (o *CreateCheckoutSessionResponse) GetTelephoneNumber() string`

GetTelephoneNumber returns the TelephoneNumber field if non-nil, zero value otherwise.

### GetTelephoneNumberOk

`func (o *CreateCheckoutSessionResponse) GetTelephoneNumberOk() (*string, bool)`

GetTelephoneNumberOk returns a tuple with the TelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumber

`func (o *CreateCheckoutSessionResponse) SetTelephoneNumber(v string)`

SetTelephoneNumber sets TelephoneNumber field to given value.

### HasTelephoneNumber

`func (o *CreateCheckoutSessionResponse) HasTelephoneNumber() bool`

HasTelephoneNumber returns a boolean if a field has been set.

### GetThreeDSAuthenticationOnly

`func (o *CreateCheckoutSessionResponse) GetThreeDSAuthenticationOnly() bool`

GetThreeDSAuthenticationOnly returns the ThreeDSAuthenticationOnly field if non-nil, zero value otherwise.

### GetThreeDSAuthenticationOnlyOk

`func (o *CreateCheckoutSessionResponse) GetThreeDSAuthenticationOnlyOk() (*bool, bool)`

GetThreeDSAuthenticationOnlyOk returns a tuple with the ThreeDSAuthenticationOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSAuthenticationOnly

`func (o *CreateCheckoutSessionResponse) SetThreeDSAuthenticationOnly(v bool)`

SetThreeDSAuthenticationOnly sets ThreeDSAuthenticationOnly field to given value.

### HasThreeDSAuthenticationOnly

`func (o *CreateCheckoutSessionResponse) HasThreeDSAuthenticationOnly() bool`

HasThreeDSAuthenticationOnly returns a boolean if a field has been set.

### GetTrustedShopper

`func (o *CreateCheckoutSessionResponse) GetTrustedShopper() bool`

GetTrustedShopper returns the TrustedShopper field if non-nil, zero value otherwise.

### GetTrustedShopperOk

`func (o *CreateCheckoutSessionResponse) GetTrustedShopperOk() (*bool, bool)`

GetTrustedShopperOk returns a tuple with the TrustedShopper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedShopper

`func (o *CreateCheckoutSessionResponse) SetTrustedShopper(v bool)`

SetTrustedShopper sets TrustedShopper field to given value.

### HasTrustedShopper

`func (o *CreateCheckoutSessionResponse) HasTrustedShopper() bool`

HasTrustedShopper returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


