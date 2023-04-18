# RecurringDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **map[string]string** | This field contains additional data, which may be returned in a particular response.  The additionalData object consists of entries, each of which includes the key and value. | [optional] 
**Alias** | Pointer to **string** | The alias of the credit card number.  Applies only to recurring contracts storing credit card details | [optional] 
**AliasType** | Pointer to **string** | The alias type of the credit card number.  Applies only to recurring contracts storing credit card details. | [optional] 
**Bank** | Pointer to [**BankAccount**](BankAccount.md) |  | [optional] 
**BillingAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**Card** | Pointer to [**Card**](Card.md) |  | [optional] 
**ContractTypes** | Pointer to **[]string** | Types of recurring contracts. | [optional] 
**CreationDate** | Pointer to **time.Time** | The date when the recurring details were created. | [optional] 
**FirstPspReference** | Pointer to **string** | The &#x60;pspReference&#x60; of the first recurring payment that created the recurring detail. | [optional] 
**Name** | Pointer to **string** | An optional descriptive name for this recurring detail. | [optional] 
**NetworkTxReference** | Pointer to **string** | Returned in the response if you are not tokenizing with Adyen and are using the Merchant-initiated transactions (MIT) framework from Mastercard or Visa.  This contains either the Mastercard Trace ID or the Visa Transaction ID. | [optional] 
**PaymentMethodVariant** | Pointer to **string** | The  type or sub-brand of a payment method used, e.g. Visa Debit, Visa Corporate, etc. For more information, refer to [PaymentMethodVariant](https://docs.adyen.com/development-resources/paymentmethodvariant). | [optional] 
**RecurringDetailReference** | **string** | The reference that uniquely identifies the recurring detail. | 
**ShopperName** | Pointer to [**Name**](Name.md) |  | [optional] 
**SocialSecurityNumber** | Pointer to **string** | A shopper&#39;s social security number (only in countries where it is legal to collect). | [optional] 
**TokenDetails** | Pointer to [**TokenDetails**](TokenDetails.md) |  | [optional] 
**Variant** | **string** | The payment method, such as “mc\&quot;, \&quot;visa\&quot;, \&quot;ideal\&quot;, \&quot;paypal\&quot;. | 

## Methods

### NewRecurringDetail

`func NewRecurringDetail(recurringDetailReference string, variant string, ) *RecurringDetail`

NewRecurringDetail instantiates a new RecurringDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurringDetailWithDefaults

`func NewRecurringDetailWithDefaults() *RecurringDetail`

NewRecurringDetailWithDefaults instantiates a new RecurringDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *RecurringDetail) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *RecurringDetail) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *RecurringDetail) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *RecurringDetail) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetAlias

`func (o *RecurringDetail) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *RecurringDetail) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *RecurringDetail) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *RecurringDetail) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetAliasType

`func (o *RecurringDetail) GetAliasType() string`

GetAliasType returns the AliasType field if non-nil, zero value otherwise.

### GetAliasTypeOk

`func (o *RecurringDetail) GetAliasTypeOk() (*string, bool)`

GetAliasTypeOk returns a tuple with the AliasType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasType

`func (o *RecurringDetail) SetAliasType(v string)`

SetAliasType sets AliasType field to given value.

### HasAliasType

`func (o *RecurringDetail) HasAliasType() bool`

HasAliasType returns a boolean if a field has been set.

### GetBank

`func (o *RecurringDetail) GetBank() BankAccount`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *RecurringDetail) GetBankOk() (*BankAccount, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *RecurringDetail) SetBank(v BankAccount)`

SetBank sets Bank field to given value.

### HasBank

`func (o *RecurringDetail) HasBank() bool`

HasBank returns a boolean if a field has been set.

### GetBillingAddress

`func (o *RecurringDetail) GetBillingAddress() Address`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *RecurringDetail) GetBillingAddressOk() (*Address, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *RecurringDetail) SetBillingAddress(v Address)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *RecurringDetail) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetCard

`func (o *RecurringDetail) GetCard() Card`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *RecurringDetail) GetCardOk() (*Card, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *RecurringDetail) SetCard(v Card)`

SetCard sets Card field to given value.

### HasCard

`func (o *RecurringDetail) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetContractTypes

`func (o *RecurringDetail) GetContractTypes() []string`

GetContractTypes returns the ContractTypes field if non-nil, zero value otherwise.

### GetContractTypesOk

`func (o *RecurringDetail) GetContractTypesOk() (*[]string, bool)`

GetContractTypesOk returns a tuple with the ContractTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTypes

`func (o *RecurringDetail) SetContractTypes(v []string)`

SetContractTypes sets ContractTypes field to given value.

### HasContractTypes

`func (o *RecurringDetail) HasContractTypes() bool`

HasContractTypes returns a boolean if a field has been set.

### GetCreationDate

`func (o *RecurringDetail) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *RecurringDetail) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *RecurringDetail) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *RecurringDetail) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetFirstPspReference

`func (o *RecurringDetail) GetFirstPspReference() string`

GetFirstPspReference returns the FirstPspReference field if non-nil, zero value otherwise.

### GetFirstPspReferenceOk

`func (o *RecurringDetail) GetFirstPspReferenceOk() (*string, bool)`

GetFirstPspReferenceOk returns a tuple with the FirstPspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPspReference

`func (o *RecurringDetail) SetFirstPspReference(v string)`

SetFirstPspReference sets FirstPspReference field to given value.

### HasFirstPspReference

`func (o *RecurringDetail) HasFirstPspReference() bool`

HasFirstPspReference returns a boolean if a field has been set.

### GetName

`func (o *RecurringDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecurringDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecurringDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecurringDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkTxReference

`func (o *RecurringDetail) GetNetworkTxReference() string`

GetNetworkTxReference returns the NetworkTxReference field if non-nil, zero value otherwise.

### GetNetworkTxReferenceOk

`func (o *RecurringDetail) GetNetworkTxReferenceOk() (*string, bool)`

GetNetworkTxReferenceOk returns a tuple with the NetworkTxReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTxReference

`func (o *RecurringDetail) SetNetworkTxReference(v string)`

SetNetworkTxReference sets NetworkTxReference field to given value.

### HasNetworkTxReference

`func (o *RecurringDetail) HasNetworkTxReference() bool`

HasNetworkTxReference returns a boolean if a field has been set.

### GetPaymentMethodVariant

`func (o *RecurringDetail) GetPaymentMethodVariant() string`

GetPaymentMethodVariant returns the PaymentMethodVariant field if non-nil, zero value otherwise.

### GetPaymentMethodVariantOk

`func (o *RecurringDetail) GetPaymentMethodVariantOk() (*string, bool)`

GetPaymentMethodVariantOk returns a tuple with the PaymentMethodVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodVariant

`func (o *RecurringDetail) SetPaymentMethodVariant(v string)`

SetPaymentMethodVariant sets PaymentMethodVariant field to given value.

### HasPaymentMethodVariant

`func (o *RecurringDetail) HasPaymentMethodVariant() bool`

HasPaymentMethodVariant returns a boolean if a field has been set.

### GetRecurringDetailReference

`func (o *RecurringDetail) GetRecurringDetailReference() string`

GetRecurringDetailReference returns the RecurringDetailReference field if non-nil, zero value otherwise.

### GetRecurringDetailReferenceOk

`func (o *RecurringDetail) GetRecurringDetailReferenceOk() (*string, bool)`

GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDetailReference

`func (o *RecurringDetail) SetRecurringDetailReference(v string)`

SetRecurringDetailReference sets RecurringDetailReference field to given value.


### GetShopperName

`func (o *RecurringDetail) GetShopperName() Name`

GetShopperName returns the ShopperName field if non-nil, zero value otherwise.

### GetShopperNameOk

`func (o *RecurringDetail) GetShopperNameOk() (*Name, bool)`

GetShopperNameOk returns a tuple with the ShopperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperName

`func (o *RecurringDetail) SetShopperName(v Name)`

SetShopperName sets ShopperName field to given value.

### HasShopperName

`func (o *RecurringDetail) HasShopperName() bool`

HasShopperName returns a boolean if a field has been set.

### GetSocialSecurityNumber

`func (o *RecurringDetail) GetSocialSecurityNumber() string`

GetSocialSecurityNumber returns the SocialSecurityNumber field if non-nil, zero value otherwise.

### GetSocialSecurityNumberOk

`func (o *RecurringDetail) GetSocialSecurityNumberOk() (*string, bool)`

GetSocialSecurityNumberOk returns a tuple with the SocialSecurityNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialSecurityNumber

`func (o *RecurringDetail) SetSocialSecurityNumber(v string)`

SetSocialSecurityNumber sets SocialSecurityNumber field to given value.

### HasSocialSecurityNumber

`func (o *RecurringDetail) HasSocialSecurityNumber() bool`

HasSocialSecurityNumber returns a boolean if a field has been set.

### GetTokenDetails

`func (o *RecurringDetail) GetTokenDetails() TokenDetails`

GetTokenDetails returns the TokenDetails field if non-nil, zero value otherwise.

### GetTokenDetailsOk

`func (o *RecurringDetail) GetTokenDetailsOk() (*TokenDetails, bool)`

GetTokenDetailsOk returns a tuple with the TokenDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenDetails

`func (o *RecurringDetail) SetTokenDetails(v TokenDetails)`

SetTokenDetails sets TokenDetails field to given value.

### HasTokenDetails

`func (o *RecurringDetail) HasTokenDetails() bool`

HasTokenDetails returns a boolean if a field has been set.

### GetVariant

`func (o *RecurringDetail) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *RecurringDetail) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *RecurringDetail) SetVariant(v string)`

SetVariant sets Variant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


