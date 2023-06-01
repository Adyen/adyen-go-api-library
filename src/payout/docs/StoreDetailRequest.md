# StoreDetailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **map[string]string** | This field contains additional data, which may be required for a particular request. | [optional] 
**Bank** | Pointer to [**BankAccount**](BankAccount.md) |  | [optional] 
**BillingAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**Card** | Pointer to [**Card**](Card.md) |  | [optional] 
**DateOfBirth** | **string** | The date of birth. Format: [ISO-8601](https://www.w3.org/TR/NOTE-datetime); example: YYYY-MM-DD For Paysafecard it must be the same as used when registering the Paysafecard account. &gt; This field is mandatory for natural persons. | 
**EntityType** | **string** | The type of the entity the payout is processed for. | 
**FraudOffset** | Pointer to **int32** | An integer value that is added to the normal fraud score. The value can be either positive or negative. | [optional] 
**MerchantAccount** | **string** | The merchant account identifier, with which you want to process the transaction. | 
**Nationality** | **string** | The shopper&#39;s nationality.  A valid value is an ISO 2-character country code (e.g. &#39;NL&#39;). | 
**Recurring** | [**Recurring**](Recurring.md) |  | 
**SelectedBrand** | Pointer to **string** | The name of the brand to make a payout to.  For Paysafecard it must be set to &#x60;paysafecard&#x60;. | [optional] 
**ShopperEmail** | **string** | The shopper&#39;s email address. | 
**ShopperName** | Pointer to [**Name**](Name.md) |  | [optional] 
**ShopperReference** | **string** | The shopper&#39;s reference for the payment transaction. | 
**SocialSecurityNumber** | Pointer to **string** | The shopper&#39;s social security number. | [optional] 
**TelephoneNumber** | Pointer to **string** | The shopper&#39;s phone number. | [optional] 

## Methods

### NewStoreDetailRequest

`func NewStoreDetailRequest(dateOfBirth string, entityType string, merchantAccount string, nationality string, recurring Recurring, shopperEmail string, shopperReference string, ) *StoreDetailRequest`

NewStoreDetailRequest instantiates a new StoreDetailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreDetailRequestWithDefaults

`func NewStoreDetailRequestWithDefaults() *StoreDetailRequest`

NewStoreDetailRequestWithDefaults instantiates a new StoreDetailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *StoreDetailRequest) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *StoreDetailRequest) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *StoreDetailRequest) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *StoreDetailRequest) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetBank

`func (o *StoreDetailRequest) GetBank() BankAccount`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *StoreDetailRequest) GetBankOk() (*BankAccount, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *StoreDetailRequest) SetBank(v BankAccount)`

SetBank sets Bank field to given value.

### HasBank

`func (o *StoreDetailRequest) HasBank() bool`

HasBank returns a boolean if a field has been set.

### GetBillingAddress

`func (o *StoreDetailRequest) GetBillingAddress() Address`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *StoreDetailRequest) GetBillingAddressOk() (*Address, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *StoreDetailRequest) SetBillingAddress(v Address)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *StoreDetailRequest) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetCard

`func (o *StoreDetailRequest) GetCard() Card`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *StoreDetailRequest) GetCardOk() (*Card, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *StoreDetailRequest) SetCard(v Card)`

SetCard sets Card field to given value.

### HasCard

`func (o *StoreDetailRequest) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *StoreDetailRequest) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *StoreDetailRequest) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *StoreDetailRequest) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.


### GetEntityType

`func (o *StoreDetailRequest) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *StoreDetailRequest) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *StoreDetailRequest) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetFraudOffset

`func (o *StoreDetailRequest) GetFraudOffset() int32`

GetFraudOffset returns the FraudOffset field if non-nil, zero value otherwise.

### GetFraudOffsetOk

`func (o *StoreDetailRequest) GetFraudOffsetOk() (*int32, bool)`

GetFraudOffsetOk returns a tuple with the FraudOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudOffset

`func (o *StoreDetailRequest) SetFraudOffset(v int32)`

SetFraudOffset sets FraudOffset field to given value.

### HasFraudOffset

`func (o *StoreDetailRequest) HasFraudOffset() bool`

HasFraudOffset returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *StoreDetailRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *StoreDetailRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *StoreDetailRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetNationality

`func (o *StoreDetailRequest) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *StoreDetailRequest) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *StoreDetailRequest) SetNationality(v string)`

SetNationality sets Nationality field to given value.


### GetRecurring

`func (o *StoreDetailRequest) GetRecurring() Recurring`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *StoreDetailRequest) GetRecurringOk() (*Recurring, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *StoreDetailRequest) SetRecurring(v Recurring)`

SetRecurring sets Recurring field to given value.


### GetSelectedBrand

`func (o *StoreDetailRequest) GetSelectedBrand() string`

GetSelectedBrand returns the SelectedBrand field if non-nil, zero value otherwise.

### GetSelectedBrandOk

`func (o *StoreDetailRequest) GetSelectedBrandOk() (*string, bool)`

GetSelectedBrandOk returns a tuple with the SelectedBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedBrand

`func (o *StoreDetailRequest) SetSelectedBrand(v string)`

SetSelectedBrand sets SelectedBrand field to given value.

### HasSelectedBrand

`func (o *StoreDetailRequest) HasSelectedBrand() bool`

HasSelectedBrand returns a boolean if a field has been set.

### GetShopperEmail

`func (o *StoreDetailRequest) GetShopperEmail() string`

GetShopperEmail returns the ShopperEmail field if non-nil, zero value otherwise.

### GetShopperEmailOk

`func (o *StoreDetailRequest) GetShopperEmailOk() (*string, bool)`

GetShopperEmailOk returns a tuple with the ShopperEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperEmail

`func (o *StoreDetailRequest) SetShopperEmail(v string)`

SetShopperEmail sets ShopperEmail field to given value.


### GetShopperName

`func (o *StoreDetailRequest) GetShopperName() Name`

GetShopperName returns the ShopperName field if non-nil, zero value otherwise.

### GetShopperNameOk

`func (o *StoreDetailRequest) GetShopperNameOk() (*Name, bool)`

GetShopperNameOk returns a tuple with the ShopperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperName

`func (o *StoreDetailRequest) SetShopperName(v Name)`

SetShopperName sets ShopperName field to given value.

### HasShopperName

`func (o *StoreDetailRequest) HasShopperName() bool`

HasShopperName returns a boolean if a field has been set.

### GetShopperReference

`func (o *StoreDetailRequest) GetShopperReference() string`

GetShopperReference returns the ShopperReference field if non-nil, zero value otherwise.

### GetShopperReferenceOk

`func (o *StoreDetailRequest) GetShopperReferenceOk() (*string, bool)`

GetShopperReferenceOk returns a tuple with the ShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperReference

`func (o *StoreDetailRequest) SetShopperReference(v string)`

SetShopperReference sets ShopperReference field to given value.


### GetSocialSecurityNumber

`func (o *StoreDetailRequest) GetSocialSecurityNumber() string`

GetSocialSecurityNumber returns the SocialSecurityNumber field if non-nil, zero value otherwise.

### GetSocialSecurityNumberOk

`func (o *StoreDetailRequest) GetSocialSecurityNumberOk() (*string, bool)`

GetSocialSecurityNumberOk returns a tuple with the SocialSecurityNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialSecurityNumber

`func (o *StoreDetailRequest) SetSocialSecurityNumber(v string)`

SetSocialSecurityNumber sets SocialSecurityNumber field to given value.

### HasSocialSecurityNumber

`func (o *StoreDetailRequest) HasSocialSecurityNumber() bool`

HasSocialSecurityNumber returns a boolean if a field has been set.

### GetTelephoneNumber

`func (o *StoreDetailRequest) GetTelephoneNumber() string`

GetTelephoneNumber returns the TelephoneNumber field if non-nil, zero value otherwise.

### GetTelephoneNumberOk

`func (o *StoreDetailRequest) GetTelephoneNumberOk() (*string, bool)`

GetTelephoneNumberOk returns a tuple with the TelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumber

`func (o *StoreDetailRequest) SetTelephoneNumber(v string)`

SetTelephoneNumber sets TelephoneNumber field to given value.

### HasTelephoneNumber

`func (o *StoreDetailRequest) HasTelephoneNumber() bool`

HasTelephoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


