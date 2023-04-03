# CreatePaymentLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedPaymentMethods** | Pointer to **[]string** | List of payment methods to be presented to the shopper. To refer to payment methods, use their &#x60;paymentMethod.type&#x60; from [Payment methods overview](https://docs.adyen.com/payment-methods).  Example: &#x60;\&quot;allowedPaymentMethods\&quot;:[\&quot;ideal\&quot;,\&quot;giropay\&quot;]&#x60; | [optional] 
**Amount** | [**Amount**](Amount.md) |  | 
**BillingAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**BlockedPaymentMethods** | Pointer to **[]string** | List of payment methods to be hidden from the shopper. To refer to payment methods, use their &#x60;paymentMethod.type&#x60; from [Payment methods overview](https://docs.adyen.com/payment-methods).  Example: &#x60;\&quot;blockedPaymentMethods\&quot;:[\&quot;ideal\&quot;,\&quot;giropay\&quot;]&#x60; | [optional] 
**CaptureDelayHours** | Pointer to **int32** | The delay between the authorisation and scheduled auto-capture, specified in hours. | [optional] 
**CountryCode** | Pointer to **string** | The shopper&#39;s two-letter country code. | [optional] 
**DateOfBirth** | Pointer to **string** | The shopper&#39;s date of birth.  Format [ISO-8601](https://www.w3.org/TR/NOTE-datetime): YYYY-MM-DD | [optional] 
**DeliverAt** | Pointer to **time.Time** | The date and time when the purchased goods should be delivered.  [ISO 8601](https://www.w3.org/TR/NOTE-datetime) format: YYYY-MM-DDThh:mm:ss+TZD, for example, **2020-12-18T10:15:30+01:00**. | [optional] 
**DeliveryAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**Description** | Pointer to **string** | A short description visible on the payment page. Maximum length: 280 characters. | [optional] 
**ExpiresAt** | Pointer to **string** | The date when the payment link expires.  [ISO 8601](https://www.w3.org/TR/NOTE-datetime) format: YYYY-MM-DDThh:mm:ss+TZD, for example, **2020-12-18T10:15:30+01:00**.  The maximum expiry date is 70 days after the payment link is created.  If not provided, the payment link expires 24 hours after it was created. | [optional] 
**InstallmentOptions** | Pointer to [**map[string]InstallmentOption**](InstallmentOption.md) | A set of key-value pairs that specifies the installment options available per payment method. The key must be a payment method name in lowercase. For example, **card** to specify installment options for all cards, or **visa** or **mc**. The value must be an object containing the installment options. | [optional] 
**LineItems** | Pointer to [**[]LineItem**](LineItem.md) | Price and product information about the purchased items, to be included on the invoice sent to the shopper. This parameter is required for open invoice (_buy now, pay later_) payment methods such Afterpay, Clearpay, Klarna, RatePay, and Zip. | [optional] 
**ManualCapture** | Pointer to **bool** | Indicates if the payment must be [captured manually](https://docs.adyen.com/online-payments/capture). | [optional] 
**Mcc** | Pointer to **string** | The [merchant category code](https://en.wikipedia.org/wiki/Merchant_category_code) (MCC) is a four-digit number, which relates to a particular market segment. This code reflects the predominant activity that is conducted by the merchant. | [optional] 
**MerchantAccount** | **string** | The merchant account identifier for which the payment link is created. | 
**MerchantOrderReference** | Pointer to **string** | This reference allows linking multiple transactions to each other for reporting purposes (for example, order auth-rate). The reference should be unique per billing cycle. | [optional] 
**Metadata** | Pointer to **map[string]string** | Metadata consists of entries, each of which includes a key and a value. Limitations: * Maximum 20 key-value pairs per request. Otherwise, error \&quot;177\&quot; occurs: \&quot;Metadata size exceeds limit\&quot; * Maximum 20 characters per key. Otherwise, error \&quot;178\&quot; occurs: \&quot;Metadata key size exceeds limit\&quot; * A key cannot have the name &#x60;checkout.linkId&#x60;. Any value that you provide with this key is going to be replaced by the real payment link ID. | [optional] 
**RecurringProcessingModel** | Pointer to **string** | Defines a recurring payment type. Required when creating a token to store payment details. Possible values: * **Subscription** – A transaction for a fixed or variable amount, which follows a fixed schedule. * **CardOnFile** – With a card-on-file (CoF) transaction, card details are stored to enable one-click or omnichannel journeys, or simply to streamline the checkout process. Any subscription not following a fixed schedule is also considered a card-on-file transaction. * **UnscheduledCardOnFile** – An unscheduled card-on-file (UCoF) transaction is a transaction that occurs on a non-fixed schedule and/or has variable amounts. For example, automatic top-ups when a cardholder&#39;s balance drops below a certain amount.  | [optional] 
**Reference** | **string** | A reference that is used to uniquely identify the payment in future communications about the payment status. | 
**RequiredShopperFields** | Pointer to **[]string** | List of fields that the shopper has to provide on the payment page before completing the payment. For more information, refer to [Provide shopper information](https://docs.adyen.com/unified-commerce/pay-by-link/payment-links/api#shopper-information).  Possible values: * **billingAddress** – The address where to send the invoice. * **deliveryAddress** – The address where the purchased goods should be delivered. * **shopperEmail** – The shopper&#39;s email address. * **shopperName** – The shopper&#39;s full name. * **telephoneNumber** – The shopper&#39;s phone number.  | [optional] 
**ReturnUrl** | Pointer to **string** | Website URL used for redirection after payment is completed. If provided, a **Continue** button will be shown on the payment page. If shoppers select the button, they are redirected to the specified URL. | [optional] 
**Reusable** | Pointer to **bool** | Indicates whether the payment link can be reused for multiple payments. If not provided, this defaults to **false** which means the link can be used for one successful payment only. | [optional] 
**RiskData** | Pointer to [**RiskData**](RiskData.md) |  | [optional] 
**ShopperEmail** | Pointer to **string** | The shopper&#39;s email address. | [optional] 
**ShopperLocale** | Pointer to **string** | The language to be used in the payment page, specified by a combination of a language and country code. For example, &#x60;en-US&#x60;.  For a list of shopper locales that Pay by Link supports, refer to [Language and localization](https://docs.adyen.com/unified-commerce/pay-by-link/payment-links/api#language). | [optional] 
**ShopperName** | Pointer to [**Name**](Name.md) |  | [optional] 
**ShopperReference** | Pointer to **string** | Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. &gt; Your reference must not include personally identifiable information (PII), for example name or email address. | [optional] 
**ShopperStatement** | Pointer to **string** | The text to be shown on the shopper&#39;s bank statement.  We recommend sending a maximum of 22 characters, otherwise banks might truncate the string.  Allowed characters: **a-z**, **A-Z**, **0-9**, spaces, and special characters **. , &#39; _ - ? + * /_**. | [optional] 
**ShowRemovePaymentMethodButton** | Pointer to **bool** | Set to **false** to hide the button that lets the shopper remove a stored payment method. | [optional] [default to true]
**SocialSecurityNumber** | Pointer to **string** | The shopper&#39;s social security number. | [optional] 
**SplitCardFundingSources** | Pointer to **bool** | Boolean value indicating whether the card payment method should be split into separate debit and credit options. | [optional] [default to false]
**Splits** | Pointer to [**[]Split**](Split.md) | An array of objects specifying how the payment should be split between accounts when using Adyen for Platforms. For details, refer to [Providing split information](https://docs.adyen.com/marketplaces-and-platforms/processing-payments#providing-split-information). | [optional] 
**Store** | Pointer to **string** | The physical store, for which this payment is processed. | [optional] 
**StorePaymentMethodMode** | Pointer to **string** | Indicates if the details of the payment method will be stored for the shopper. Possible values: * **disabled** – No details will be stored (default). * **askForConsent** – If the &#x60;shopperReference&#x60; is provided, the UI lets the shopper choose if they want their payment details to be stored. * **enabled** – If the &#x60;shopperReference&#x60; is provided, the details will be stored without asking the shopper for consent. | [optional] 
**TelephoneNumber** | Pointer to **string** | The shopper&#39;s telephone number. | [optional] 
**ThemeId** | Pointer to **string** | A [theme](https://docs.adyen.com/unified-commerce/pay-by-link/payment-links/api#themes) to customize the appearance of the payment page. If not specified, the payment page is rendered according to the theme set as default in your Customer Area. | [optional] 

## Methods

### NewCreatePaymentLinkRequest

`func NewCreatePaymentLinkRequest(amount Amount, merchantAccount string, reference string, ) *CreatePaymentLinkRequest`

NewCreatePaymentLinkRequest instantiates a new CreatePaymentLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaymentLinkRequestWithDefaults

`func NewCreatePaymentLinkRequestWithDefaults() *CreatePaymentLinkRequest`

NewCreatePaymentLinkRequestWithDefaults instantiates a new CreatePaymentLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedPaymentMethods

`func (o *CreatePaymentLinkRequest) GetAllowedPaymentMethods() []string`

GetAllowedPaymentMethods returns the AllowedPaymentMethods field if non-nil, zero value otherwise.

### GetAllowedPaymentMethodsOk

`func (o *CreatePaymentLinkRequest) GetAllowedPaymentMethodsOk() (*[]string, bool)`

GetAllowedPaymentMethodsOk returns a tuple with the AllowedPaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPaymentMethods

`func (o *CreatePaymentLinkRequest) SetAllowedPaymentMethods(v []string)`

SetAllowedPaymentMethods sets AllowedPaymentMethods field to given value.

### HasAllowedPaymentMethods

`func (o *CreatePaymentLinkRequest) HasAllowedPaymentMethods() bool`

HasAllowedPaymentMethods returns a boolean if a field has been set.

### GetAmount

`func (o *CreatePaymentLinkRequest) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreatePaymentLinkRequest) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreatePaymentLinkRequest) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetBillingAddress

`func (o *CreatePaymentLinkRequest) GetBillingAddress() Address`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *CreatePaymentLinkRequest) GetBillingAddressOk() (*Address, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *CreatePaymentLinkRequest) SetBillingAddress(v Address)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *CreatePaymentLinkRequest) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetBlockedPaymentMethods

`func (o *CreatePaymentLinkRequest) GetBlockedPaymentMethods() []string`

GetBlockedPaymentMethods returns the BlockedPaymentMethods field if non-nil, zero value otherwise.

### GetBlockedPaymentMethodsOk

`func (o *CreatePaymentLinkRequest) GetBlockedPaymentMethodsOk() (*[]string, bool)`

GetBlockedPaymentMethodsOk returns a tuple with the BlockedPaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedPaymentMethods

`func (o *CreatePaymentLinkRequest) SetBlockedPaymentMethods(v []string)`

SetBlockedPaymentMethods sets BlockedPaymentMethods field to given value.

### HasBlockedPaymentMethods

`func (o *CreatePaymentLinkRequest) HasBlockedPaymentMethods() bool`

HasBlockedPaymentMethods returns a boolean if a field has been set.

### GetCaptureDelayHours

`func (o *CreatePaymentLinkRequest) GetCaptureDelayHours() int32`

GetCaptureDelayHours returns the CaptureDelayHours field if non-nil, zero value otherwise.

### GetCaptureDelayHoursOk

`func (o *CreatePaymentLinkRequest) GetCaptureDelayHoursOk() (*int32, bool)`

GetCaptureDelayHoursOk returns a tuple with the CaptureDelayHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureDelayHours

`func (o *CreatePaymentLinkRequest) SetCaptureDelayHours(v int32)`

SetCaptureDelayHours sets CaptureDelayHours field to given value.

### HasCaptureDelayHours

`func (o *CreatePaymentLinkRequest) HasCaptureDelayHours() bool`

HasCaptureDelayHours returns a boolean if a field has been set.

### GetCountryCode

`func (o *CreatePaymentLinkRequest) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CreatePaymentLinkRequest) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CreatePaymentLinkRequest) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *CreatePaymentLinkRequest) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *CreatePaymentLinkRequest) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *CreatePaymentLinkRequest) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *CreatePaymentLinkRequest) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *CreatePaymentLinkRequest) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetDeliverAt

`func (o *CreatePaymentLinkRequest) GetDeliverAt() time.Time`

GetDeliverAt returns the DeliverAt field if non-nil, zero value otherwise.

### GetDeliverAtOk

`func (o *CreatePaymentLinkRequest) GetDeliverAtOk() (*time.Time, bool)`

GetDeliverAtOk returns a tuple with the DeliverAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverAt

`func (o *CreatePaymentLinkRequest) SetDeliverAt(v time.Time)`

SetDeliverAt sets DeliverAt field to given value.

### HasDeliverAt

`func (o *CreatePaymentLinkRequest) HasDeliverAt() bool`

HasDeliverAt returns a boolean if a field has been set.

### GetDeliveryAddress

`func (o *CreatePaymentLinkRequest) GetDeliveryAddress() Address`

GetDeliveryAddress returns the DeliveryAddress field if non-nil, zero value otherwise.

### GetDeliveryAddressOk

`func (o *CreatePaymentLinkRequest) GetDeliveryAddressOk() (*Address, bool)`

GetDeliveryAddressOk returns a tuple with the DeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddress

`func (o *CreatePaymentLinkRequest) SetDeliveryAddress(v Address)`

SetDeliveryAddress sets DeliveryAddress field to given value.

### HasDeliveryAddress

`func (o *CreatePaymentLinkRequest) HasDeliveryAddress() bool`

HasDeliveryAddress returns a boolean if a field has been set.

### GetDescription

`func (o *CreatePaymentLinkRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePaymentLinkRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePaymentLinkRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatePaymentLinkRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CreatePaymentLinkRequest) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreatePaymentLinkRequest) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreatePaymentLinkRequest) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreatePaymentLinkRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetInstallmentOptions

`func (o *CreatePaymentLinkRequest) GetInstallmentOptions() map[string]InstallmentOption`

GetInstallmentOptions returns the InstallmentOptions field if non-nil, zero value otherwise.

### GetInstallmentOptionsOk

`func (o *CreatePaymentLinkRequest) GetInstallmentOptionsOk() (*map[string]InstallmentOption, bool)`

GetInstallmentOptionsOk returns a tuple with the InstallmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallmentOptions

`func (o *CreatePaymentLinkRequest) SetInstallmentOptions(v map[string]InstallmentOption)`

SetInstallmentOptions sets InstallmentOptions field to given value.

### HasInstallmentOptions

`func (o *CreatePaymentLinkRequest) HasInstallmentOptions() bool`

HasInstallmentOptions returns a boolean if a field has been set.

### GetLineItems

`func (o *CreatePaymentLinkRequest) GetLineItems() []LineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *CreatePaymentLinkRequest) GetLineItemsOk() (*[]LineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *CreatePaymentLinkRequest) SetLineItems(v []LineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *CreatePaymentLinkRequest) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetManualCapture

`func (o *CreatePaymentLinkRequest) GetManualCapture() bool`

GetManualCapture returns the ManualCapture field if non-nil, zero value otherwise.

### GetManualCaptureOk

`func (o *CreatePaymentLinkRequest) GetManualCaptureOk() (*bool, bool)`

GetManualCaptureOk returns a tuple with the ManualCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualCapture

`func (o *CreatePaymentLinkRequest) SetManualCapture(v bool)`

SetManualCapture sets ManualCapture field to given value.

### HasManualCapture

`func (o *CreatePaymentLinkRequest) HasManualCapture() bool`

HasManualCapture returns a boolean if a field has been set.

### GetMcc

`func (o *CreatePaymentLinkRequest) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *CreatePaymentLinkRequest) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *CreatePaymentLinkRequest) SetMcc(v string)`

SetMcc sets Mcc field to given value.

### HasMcc

`func (o *CreatePaymentLinkRequest) HasMcc() bool`

HasMcc returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *CreatePaymentLinkRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *CreatePaymentLinkRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *CreatePaymentLinkRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetMerchantOrderReference

`func (o *CreatePaymentLinkRequest) GetMerchantOrderReference() string`

GetMerchantOrderReference returns the MerchantOrderReference field if non-nil, zero value otherwise.

### GetMerchantOrderReferenceOk

`func (o *CreatePaymentLinkRequest) GetMerchantOrderReferenceOk() (*string, bool)`

GetMerchantOrderReferenceOk returns a tuple with the MerchantOrderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderReference

`func (o *CreatePaymentLinkRequest) SetMerchantOrderReference(v string)`

SetMerchantOrderReference sets MerchantOrderReference field to given value.

### HasMerchantOrderReference

`func (o *CreatePaymentLinkRequest) HasMerchantOrderReference() bool`

HasMerchantOrderReference returns a boolean if a field has been set.

### GetMetadata

`func (o *CreatePaymentLinkRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreatePaymentLinkRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreatePaymentLinkRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreatePaymentLinkRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRecurringProcessingModel

`func (o *CreatePaymentLinkRequest) GetRecurringProcessingModel() string`

GetRecurringProcessingModel returns the RecurringProcessingModel field if non-nil, zero value otherwise.

### GetRecurringProcessingModelOk

`func (o *CreatePaymentLinkRequest) GetRecurringProcessingModelOk() (*string, bool)`

GetRecurringProcessingModelOk returns a tuple with the RecurringProcessingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringProcessingModel

`func (o *CreatePaymentLinkRequest) SetRecurringProcessingModel(v string)`

SetRecurringProcessingModel sets RecurringProcessingModel field to given value.

### HasRecurringProcessingModel

`func (o *CreatePaymentLinkRequest) HasRecurringProcessingModel() bool`

HasRecurringProcessingModel returns a boolean if a field has been set.

### GetReference

`func (o *CreatePaymentLinkRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CreatePaymentLinkRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CreatePaymentLinkRequest) SetReference(v string)`

SetReference sets Reference field to given value.


### GetRequiredShopperFields

`func (o *CreatePaymentLinkRequest) GetRequiredShopperFields() []string`

GetRequiredShopperFields returns the RequiredShopperFields field if non-nil, zero value otherwise.

### GetRequiredShopperFieldsOk

`func (o *CreatePaymentLinkRequest) GetRequiredShopperFieldsOk() (*[]string, bool)`

GetRequiredShopperFieldsOk returns a tuple with the RequiredShopperFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredShopperFields

`func (o *CreatePaymentLinkRequest) SetRequiredShopperFields(v []string)`

SetRequiredShopperFields sets RequiredShopperFields field to given value.

### HasRequiredShopperFields

`func (o *CreatePaymentLinkRequest) HasRequiredShopperFields() bool`

HasRequiredShopperFields returns a boolean if a field has been set.

### GetReturnUrl

`func (o *CreatePaymentLinkRequest) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *CreatePaymentLinkRequest) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *CreatePaymentLinkRequest) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *CreatePaymentLinkRequest) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetReusable

`func (o *CreatePaymentLinkRequest) GetReusable() bool`

GetReusable returns the Reusable field if non-nil, zero value otherwise.

### GetReusableOk

`func (o *CreatePaymentLinkRequest) GetReusableOk() (*bool, bool)`

GetReusableOk returns a tuple with the Reusable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusable

`func (o *CreatePaymentLinkRequest) SetReusable(v bool)`

SetReusable sets Reusable field to given value.

### HasReusable

`func (o *CreatePaymentLinkRequest) HasReusable() bool`

HasReusable returns a boolean if a field has been set.

### GetRiskData

`func (o *CreatePaymentLinkRequest) GetRiskData() RiskData`

GetRiskData returns the RiskData field if non-nil, zero value otherwise.

### GetRiskDataOk

`func (o *CreatePaymentLinkRequest) GetRiskDataOk() (*RiskData, bool)`

GetRiskDataOk returns a tuple with the RiskData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskData

`func (o *CreatePaymentLinkRequest) SetRiskData(v RiskData)`

SetRiskData sets RiskData field to given value.

### HasRiskData

`func (o *CreatePaymentLinkRequest) HasRiskData() bool`

HasRiskData returns a boolean if a field has been set.

### GetShopperEmail

`func (o *CreatePaymentLinkRequest) GetShopperEmail() string`

GetShopperEmail returns the ShopperEmail field if non-nil, zero value otherwise.

### GetShopperEmailOk

`func (o *CreatePaymentLinkRequest) GetShopperEmailOk() (*string, bool)`

GetShopperEmailOk returns a tuple with the ShopperEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperEmail

`func (o *CreatePaymentLinkRequest) SetShopperEmail(v string)`

SetShopperEmail sets ShopperEmail field to given value.

### HasShopperEmail

`func (o *CreatePaymentLinkRequest) HasShopperEmail() bool`

HasShopperEmail returns a boolean if a field has been set.

### GetShopperLocale

`func (o *CreatePaymentLinkRequest) GetShopperLocale() string`

GetShopperLocale returns the ShopperLocale field if non-nil, zero value otherwise.

### GetShopperLocaleOk

`func (o *CreatePaymentLinkRequest) GetShopperLocaleOk() (*string, bool)`

GetShopperLocaleOk returns a tuple with the ShopperLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperLocale

`func (o *CreatePaymentLinkRequest) SetShopperLocale(v string)`

SetShopperLocale sets ShopperLocale field to given value.

### HasShopperLocale

`func (o *CreatePaymentLinkRequest) HasShopperLocale() bool`

HasShopperLocale returns a boolean if a field has been set.

### GetShopperName

`func (o *CreatePaymentLinkRequest) GetShopperName() Name`

GetShopperName returns the ShopperName field if non-nil, zero value otherwise.

### GetShopperNameOk

`func (o *CreatePaymentLinkRequest) GetShopperNameOk() (*Name, bool)`

GetShopperNameOk returns a tuple with the ShopperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperName

`func (o *CreatePaymentLinkRequest) SetShopperName(v Name)`

SetShopperName sets ShopperName field to given value.

### HasShopperName

`func (o *CreatePaymentLinkRequest) HasShopperName() bool`

HasShopperName returns a boolean if a field has been set.

### GetShopperReference

`func (o *CreatePaymentLinkRequest) GetShopperReference() string`

GetShopperReference returns the ShopperReference field if non-nil, zero value otherwise.

### GetShopperReferenceOk

`func (o *CreatePaymentLinkRequest) GetShopperReferenceOk() (*string, bool)`

GetShopperReferenceOk returns a tuple with the ShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperReference

`func (o *CreatePaymentLinkRequest) SetShopperReference(v string)`

SetShopperReference sets ShopperReference field to given value.

### HasShopperReference

`func (o *CreatePaymentLinkRequest) HasShopperReference() bool`

HasShopperReference returns a boolean if a field has been set.

### GetShopperStatement

`func (o *CreatePaymentLinkRequest) GetShopperStatement() string`

GetShopperStatement returns the ShopperStatement field if non-nil, zero value otherwise.

### GetShopperStatementOk

`func (o *CreatePaymentLinkRequest) GetShopperStatementOk() (*string, bool)`

GetShopperStatementOk returns a tuple with the ShopperStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperStatement

`func (o *CreatePaymentLinkRequest) SetShopperStatement(v string)`

SetShopperStatement sets ShopperStatement field to given value.

### HasShopperStatement

`func (o *CreatePaymentLinkRequest) HasShopperStatement() bool`

HasShopperStatement returns a boolean if a field has been set.

### GetShowRemovePaymentMethodButton

`func (o *CreatePaymentLinkRequest) GetShowRemovePaymentMethodButton() bool`

GetShowRemovePaymentMethodButton returns the ShowRemovePaymentMethodButton field if non-nil, zero value otherwise.

### GetShowRemovePaymentMethodButtonOk

`func (o *CreatePaymentLinkRequest) GetShowRemovePaymentMethodButtonOk() (*bool, bool)`

GetShowRemovePaymentMethodButtonOk returns a tuple with the ShowRemovePaymentMethodButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowRemovePaymentMethodButton

`func (o *CreatePaymentLinkRequest) SetShowRemovePaymentMethodButton(v bool)`

SetShowRemovePaymentMethodButton sets ShowRemovePaymentMethodButton field to given value.

### HasShowRemovePaymentMethodButton

`func (o *CreatePaymentLinkRequest) HasShowRemovePaymentMethodButton() bool`

HasShowRemovePaymentMethodButton returns a boolean if a field has been set.

### GetSocialSecurityNumber

`func (o *CreatePaymentLinkRequest) GetSocialSecurityNumber() string`

GetSocialSecurityNumber returns the SocialSecurityNumber field if non-nil, zero value otherwise.

### GetSocialSecurityNumberOk

`func (o *CreatePaymentLinkRequest) GetSocialSecurityNumberOk() (*string, bool)`

GetSocialSecurityNumberOk returns a tuple with the SocialSecurityNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialSecurityNumber

`func (o *CreatePaymentLinkRequest) SetSocialSecurityNumber(v string)`

SetSocialSecurityNumber sets SocialSecurityNumber field to given value.

### HasSocialSecurityNumber

`func (o *CreatePaymentLinkRequest) HasSocialSecurityNumber() bool`

HasSocialSecurityNumber returns a boolean if a field has been set.

### GetSplitCardFundingSources

`func (o *CreatePaymentLinkRequest) GetSplitCardFundingSources() bool`

GetSplitCardFundingSources returns the SplitCardFundingSources field if non-nil, zero value otherwise.

### GetSplitCardFundingSourcesOk

`func (o *CreatePaymentLinkRequest) GetSplitCardFundingSourcesOk() (*bool, bool)`

GetSplitCardFundingSourcesOk returns a tuple with the SplitCardFundingSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitCardFundingSources

`func (o *CreatePaymentLinkRequest) SetSplitCardFundingSources(v bool)`

SetSplitCardFundingSources sets SplitCardFundingSources field to given value.

### HasSplitCardFundingSources

`func (o *CreatePaymentLinkRequest) HasSplitCardFundingSources() bool`

HasSplitCardFundingSources returns a boolean if a field has been set.

### GetSplits

`func (o *CreatePaymentLinkRequest) GetSplits() []Split`

GetSplits returns the Splits field if non-nil, zero value otherwise.

### GetSplitsOk

`func (o *CreatePaymentLinkRequest) GetSplitsOk() (*[]Split, bool)`

GetSplitsOk returns a tuple with the Splits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplits

`func (o *CreatePaymentLinkRequest) SetSplits(v []Split)`

SetSplits sets Splits field to given value.

### HasSplits

`func (o *CreatePaymentLinkRequest) HasSplits() bool`

HasSplits returns a boolean if a field has been set.

### GetStore

`func (o *CreatePaymentLinkRequest) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *CreatePaymentLinkRequest) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *CreatePaymentLinkRequest) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *CreatePaymentLinkRequest) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetStorePaymentMethodMode

`func (o *CreatePaymentLinkRequest) GetStorePaymentMethodMode() string`

GetStorePaymentMethodMode returns the StorePaymentMethodMode field if non-nil, zero value otherwise.

### GetStorePaymentMethodModeOk

`func (o *CreatePaymentLinkRequest) GetStorePaymentMethodModeOk() (*string, bool)`

GetStorePaymentMethodModeOk returns a tuple with the StorePaymentMethodMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePaymentMethodMode

`func (o *CreatePaymentLinkRequest) SetStorePaymentMethodMode(v string)`

SetStorePaymentMethodMode sets StorePaymentMethodMode field to given value.

### HasStorePaymentMethodMode

`func (o *CreatePaymentLinkRequest) HasStorePaymentMethodMode() bool`

HasStorePaymentMethodMode returns a boolean if a field has been set.

### GetTelephoneNumber

`func (o *CreatePaymentLinkRequest) GetTelephoneNumber() string`

GetTelephoneNumber returns the TelephoneNumber field if non-nil, zero value otherwise.

### GetTelephoneNumberOk

`func (o *CreatePaymentLinkRequest) GetTelephoneNumberOk() (*string, bool)`

GetTelephoneNumberOk returns a tuple with the TelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumber

`func (o *CreatePaymentLinkRequest) SetTelephoneNumber(v string)`

SetTelephoneNumber sets TelephoneNumber field to given value.

### HasTelephoneNumber

`func (o *CreatePaymentLinkRequest) HasTelephoneNumber() bool`

HasTelephoneNumber returns a boolean if a field has been set.

### GetThemeId

`func (o *CreatePaymentLinkRequest) GetThemeId() string`

GetThemeId returns the ThemeId field if non-nil, zero value otherwise.

### GetThemeIdOk

`func (o *CreatePaymentLinkRequest) GetThemeIdOk() (*string, bool)`

GetThemeIdOk returns a tuple with the ThemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeId

`func (o *CreatePaymentLinkRequest) SetThemeId(v string)`

SetThemeId sets ThemeId field to given value.

### HasThemeId

`func (o *CreatePaymentLinkRequest) HasThemeId() bool`

HasThemeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


