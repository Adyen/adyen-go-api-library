# PaymentMethodsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **map[string]string** | This field contains additional data, which may be required for a particular payment request.  The &#x60;additionalData&#x60; object consists of entries, each of which includes the key and value. | [optional] 
**AllowedPaymentMethods** | Pointer to **[]string** | List of payment methods to be presented to the shopper. To refer to payment methods, use their [payment method type](https://docs.adyen.com/payment-methods/payment-method-types).  Example: &#x60;\&quot;allowedPaymentMethods\&quot;:[\&quot;ideal\&quot;,\&quot;giropay\&quot;]&#x60; | [optional] 
**Amount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**BlockedPaymentMethods** | Pointer to **[]string** | List of payment methods to be hidden from the shopper. To refer to payment methods, use their [payment method type](https://docs.adyen.com/payment-methods/payment-method-types).  Example: &#x60;\&quot;blockedPaymentMethods\&quot;:[\&quot;ideal\&quot;,\&quot;giropay\&quot;]&#x60; | [optional] 
**Channel** | Pointer to **string** | The platform where a payment transaction takes place. This field can be used for filtering out payment methods that are only available on specific platforms. Possible values: * iOS * Android * Web | [optional] 
**CountryCode** | Pointer to **string** | The shopper&#39;s country code. | [optional] 
**MerchantAccount** | **string** | The merchant account identifier, with which you want to process the transaction. | 
**Order** | Pointer to [**EncryptedOrderData**](EncryptedOrderData.md) |  | [optional] 
**ShopperLocale** | Pointer to **string** | The combination of a language code and a country code to specify the language to be used in the payment. | [optional] 
**ShopperReference** | Pointer to **string** | Required for recurring payments.  Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. &gt; Your reference must not include personally identifiable information (PII), for example name or email address. | [optional] 
**SplitCardFundingSources** | Pointer to **bool** | Boolean value indicating whether the card payment method should be split into separate debit and credit options. | [optional] [default to false]
**Store** | Pointer to **string** | The ecommerce or point-of-sale store that is processing the payment. Used in [partner model integrations](https://docs.adyen.com/marketplaces-and-platforms/classic/platforms-for-partners#route-payments) for Adyen for Platforms. | [optional] 

## Methods

### NewPaymentMethodsRequest

`func NewPaymentMethodsRequest(merchantAccount string, ) *PaymentMethodsRequest`

NewPaymentMethodsRequest instantiates a new PaymentMethodsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodsRequestWithDefaults

`func NewPaymentMethodsRequestWithDefaults() *PaymentMethodsRequest`

NewPaymentMethodsRequestWithDefaults instantiates a new PaymentMethodsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *PaymentMethodsRequest) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *PaymentMethodsRequest) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *PaymentMethodsRequest) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *PaymentMethodsRequest) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetAllowedPaymentMethods

`func (o *PaymentMethodsRequest) GetAllowedPaymentMethods() []string`

GetAllowedPaymentMethods returns the AllowedPaymentMethods field if non-nil, zero value otherwise.

### GetAllowedPaymentMethodsOk

`func (o *PaymentMethodsRequest) GetAllowedPaymentMethodsOk() (*[]string, bool)`

GetAllowedPaymentMethodsOk returns a tuple with the AllowedPaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPaymentMethods

`func (o *PaymentMethodsRequest) SetAllowedPaymentMethods(v []string)`

SetAllowedPaymentMethods sets AllowedPaymentMethods field to given value.

### HasAllowedPaymentMethods

`func (o *PaymentMethodsRequest) HasAllowedPaymentMethods() bool`

HasAllowedPaymentMethods returns a boolean if a field has been set.

### GetAmount

`func (o *PaymentMethodsRequest) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentMethodsRequest) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentMethodsRequest) SetAmount(v Amount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentMethodsRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBlockedPaymentMethods

`func (o *PaymentMethodsRequest) GetBlockedPaymentMethods() []string`

GetBlockedPaymentMethods returns the BlockedPaymentMethods field if non-nil, zero value otherwise.

### GetBlockedPaymentMethodsOk

`func (o *PaymentMethodsRequest) GetBlockedPaymentMethodsOk() (*[]string, bool)`

GetBlockedPaymentMethodsOk returns a tuple with the BlockedPaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedPaymentMethods

`func (o *PaymentMethodsRequest) SetBlockedPaymentMethods(v []string)`

SetBlockedPaymentMethods sets BlockedPaymentMethods field to given value.

### HasBlockedPaymentMethods

`func (o *PaymentMethodsRequest) HasBlockedPaymentMethods() bool`

HasBlockedPaymentMethods returns a boolean if a field has been set.

### GetChannel

`func (o *PaymentMethodsRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PaymentMethodsRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PaymentMethodsRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *PaymentMethodsRequest) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCountryCode

`func (o *PaymentMethodsRequest) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PaymentMethodsRequest) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PaymentMethodsRequest) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *PaymentMethodsRequest) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *PaymentMethodsRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *PaymentMethodsRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *PaymentMethodsRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetOrder

`func (o *PaymentMethodsRequest) GetOrder() EncryptedOrderData`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PaymentMethodsRequest) GetOrderOk() (*EncryptedOrderData, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PaymentMethodsRequest) SetOrder(v EncryptedOrderData)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PaymentMethodsRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetShopperLocale

`func (o *PaymentMethodsRequest) GetShopperLocale() string`

GetShopperLocale returns the ShopperLocale field if non-nil, zero value otherwise.

### GetShopperLocaleOk

`func (o *PaymentMethodsRequest) GetShopperLocaleOk() (*string, bool)`

GetShopperLocaleOk returns a tuple with the ShopperLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperLocale

`func (o *PaymentMethodsRequest) SetShopperLocale(v string)`

SetShopperLocale sets ShopperLocale field to given value.

### HasShopperLocale

`func (o *PaymentMethodsRequest) HasShopperLocale() bool`

HasShopperLocale returns a boolean if a field has been set.

### GetShopperReference

`func (o *PaymentMethodsRequest) GetShopperReference() string`

GetShopperReference returns the ShopperReference field if non-nil, zero value otherwise.

### GetShopperReferenceOk

`func (o *PaymentMethodsRequest) GetShopperReferenceOk() (*string, bool)`

GetShopperReferenceOk returns a tuple with the ShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperReference

`func (o *PaymentMethodsRequest) SetShopperReference(v string)`

SetShopperReference sets ShopperReference field to given value.

### HasShopperReference

`func (o *PaymentMethodsRequest) HasShopperReference() bool`

HasShopperReference returns a boolean if a field has been set.

### GetSplitCardFundingSources

`func (o *PaymentMethodsRequest) GetSplitCardFundingSources() bool`

GetSplitCardFundingSources returns the SplitCardFundingSources field if non-nil, zero value otherwise.

### GetSplitCardFundingSourcesOk

`func (o *PaymentMethodsRequest) GetSplitCardFundingSourcesOk() (*bool, bool)`

GetSplitCardFundingSourcesOk returns a tuple with the SplitCardFundingSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitCardFundingSources

`func (o *PaymentMethodsRequest) SetSplitCardFundingSources(v bool)`

SetSplitCardFundingSources sets SplitCardFundingSources field to given value.

### HasSplitCardFundingSources

`func (o *PaymentMethodsRequest) HasSplitCardFundingSources() bool`

HasSplitCardFundingSources returns a boolean if a field has been set.

### GetStore

`func (o *PaymentMethodsRequest) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *PaymentMethodsRequest) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *PaymentMethodsRequest) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *PaymentMethodsRequest) HasStore() bool`

HasStore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


