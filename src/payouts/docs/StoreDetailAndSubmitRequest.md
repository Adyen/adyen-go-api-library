# StoreDetailAndSubmitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **map[string]string** | This field contains additional data, which may be required for a particular request. | [optional] 
**Amount** | [**Amount**](Amount.md) |  | 
**Bank** | Pointer to [**BankAccount**](BankAccount.md) |  | [optional] 
**BillingAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**Card** | Pointer to [**Card**](Card.md) |  | [optional] 
**DateOfBirth** | **time.Time** | The date of birth. Format: [ISO-8601](https://www.w3.org/TR/NOTE-datetime); example: YYYY-MM-DD For Paysafecard it must be the same as used when registering the Paysafecard account. &gt; This field is mandatory for natural persons. | 
**EntityType** | **string** | The type of the entity the payout is processed for. | 
**FraudOffset** | Pointer to **int32** | An integer value that is added to the normal fraud score. The value can be either positive or negative. | [optional] 
**MerchantAccount** | **string** | The merchant account identifier, with which you want to process the transaction. | 
**Nationality** | **string** | The shopper&#39;s nationality.  A valid value is an ISO 2-character country code (e.g. &#39;NL&#39;). | 
**Recurring** | [**Recurring**](Recurring.md) |  | 
**Reference** | **string** | The merchant reference for this payment. This reference will be used in all communication to the merchant about the status of the payout. Although it is a good idea to make sure it is unique, this is not a requirement. | 
**SelectedBrand** | Pointer to **string** | The name of the brand to make a payout to.  For Paysafecard it must be set to &#x60;paysafecard&#x60;. | [optional] 
**ShopperEmail** | **string** | The shopper&#39;s email address. | 
**ShopperName** | Pointer to [**Name**](Name.md) |  | [optional] 
**ShopperReference** | **string** | The shopper&#39;s reference for the payment transaction. | 
**ShopperStatement** | Pointer to **string** | The description of this payout. This description is shown on the bank statement of the shopper (if this is supported by the chosen payment method). | [optional] 
**SocialSecurityNumber** | Pointer to **string** | The shopper&#39;s social security number. | [optional] 
**TelephoneNumber** | Pointer to **string** | The shopper&#39;s phone number. | [optional] 

## Methods

### NewStoreDetailAndSubmitRequest

`func NewStoreDetailAndSubmitRequest(amount Amount, dateOfBirth time.Time, entityType string, merchantAccount string, nationality string, recurring Recurring, reference string, shopperEmail string, shopperReference string, ) *StoreDetailAndSubmitRequest`

NewStoreDetailAndSubmitRequest instantiates a new StoreDetailAndSubmitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreDetailAndSubmitRequestWithDefaults

`func NewStoreDetailAndSubmitRequestWithDefaults() *StoreDetailAndSubmitRequest`

NewStoreDetailAndSubmitRequestWithDefaults instantiates a new StoreDetailAndSubmitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *StoreDetailAndSubmitRequest) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *StoreDetailAndSubmitRequest) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *StoreDetailAndSubmitRequest) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *StoreDetailAndSubmitRequest) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetAmount

`func (o *StoreDetailAndSubmitRequest) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *StoreDetailAndSubmitRequest) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *StoreDetailAndSubmitRequest) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetBank

`func (o *StoreDetailAndSubmitRequest) GetBank() BankAccount`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *StoreDetailAndSubmitRequest) GetBankOk() (*BankAccount, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *StoreDetailAndSubmitRequest) SetBank(v BankAccount)`

SetBank sets Bank field to given value.

### HasBank

`func (o *StoreDetailAndSubmitRequest) HasBank() bool`

HasBank returns a boolean if a field has been set.

### GetBillingAddress

`func (o *StoreDetailAndSubmitRequest) GetBillingAddress() Address`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *StoreDetailAndSubmitRequest) GetBillingAddressOk() (*Address, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *StoreDetailAndSubmitRequest) SetBillingAddress(v Address)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *StoreDetailAndSubmitRequest) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetCard

`func (o *StoreDetailAndSubmitRequest) GetCard() Card`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *StoreDetailAndSubmitRequest) GetCardOk() (*Card, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *StoreDetailAndSubmitRequest) SetCard(v Card)`

SetCard sets Card field to given value.

### HasCard

`func (o *StoreDetailAndSubmitRequest) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *StoreDetailAndSubmitRequest) GetDateOfBirth() time.Time`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *StoreDetailAndSubmitRequest) GetDateOfBirthOk() (*time.Time, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *StoreDetailAndSubmitRequest) SetDateOfBirth(v time.Time)`

SetDateOfBirth sets DateOfBirth field to given value.


### GetEntityType

`func (o *StoreDetailAndSubmitRequest) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *StoreDetailAndSubmitRequest) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *StoreDetailAndSubmitRequest) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetFraudOffset

`func (o *StoreDetailAndSubmitRequest) GetFraudOffset() int32`

GetFraudOffset returns the FraudOffset field if non-nil, zero value otherwise.

### GetFraudOffsetOk

`func (o *StoreDetailAndSubmitRequest) GetFraudOffsetOk() (*int32, bool)`

GetFraudOffsetOk returns a tuple with the FraudOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudOffset

`func (o *StoreDetailAndSubmitRequest) SetFraudOffset(v int32)`

SetFraudOffset sets FraudOffset field to given value.

### HasFraudOffset

`func (o *StoreDetailAndSubmitRequest) HasFraudOffset() bool`

HasFraudOffset returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *StoreDetailAndSubmitRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *StoreDetailAndSubmitRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *StoreDetailAndSubmitRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetNationality

`func (o *StoreDetailAndSubmitRequest) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *StoreDetailAndSubmitRequest) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *StoreDetailAndSubmitRequest) SetNationality(v string)`

SetNationality sets Nationality field to given value.


### GetRecurring

`func (o *StoreDetailAndSubmitRequest) GetRecurring() Recurring`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *StoreDetailAndSubmitRequest) GetRecurringOk() (*Recurring, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *StoreDetailAndSubmitRequest) SetRecurring(v Recurring)`

SetRecurring sets Recurring field to given value.


### GetReference

`func (o *StoreDetailAndSubmitRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *StoreDetailAndSubmitRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *StoreDetailAndSubmitRequest) SetReference(v string)`

SetReference sets Reference field to given value.


### GetSelectedBrand

`func (o *StoreDetailAndSubmitRequest) GetSelectedBrand() string`

GetSelectedBrand returns the SelectedBrand field if non-nil, zero value otherwise.

### GetSelectedBrandOk

`func (o *StoreDetailAndSubmitRequest) GetSelectedBrandOk() (*string, bool)`

GetSelectedBrandOk returns a tuple with the SelectedBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedBrand

`func (o *StoreDetailAndSubmitRequest) SetSelectedBrand(v string)`

SetSelectedBrand sets SelectedBrand field to given value.

### HasSelectedBrand

`func (o *StoreDetailAndSubmitRequest) HasSelectedBrand() bool`

HasSelectedBrand returns a boolean if a field has been set.

### GetShopperEmail

`func (o *StoreDetailAndSubmitRequest) GetShopperEmail() string`

GetShopperEmail returns the ShopperEmail field if non-nil, zero value otherwise.

### GetShopperEmailOk

`func (o *StoreDetailAndSubmitRequest) GetShopperEmailOk() (*string, bool)`

GetShopperEmailOk returns a tuple with the ShopperEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperEmail

`func (o *StoreDetailAndSubmitRequest) SetShopperEmail(v string)`

SetShopperEmail sets ShopperEmail field to given value.


### GetShopperName

`func (o *StoreDetailAndSubmitRequest) GetShopperName() Name`

GetShopperName returns the ShopperName field if non-nil, zero value otherwise.

### GetShopperNameOk

`func (o *StoreDetailAndSubmitRequest) GetShopperNameOk() (*Name, bool)`

GetShopperNameOk returns a tuple with the ShopperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperName

`func (o *StoreDetailAndSubmitRequest) SetShopperName(v Name)`

SetShopperName sets ShopperName field to given value.

### HasShopperName

`func (o *StoreDetailAndSubmitRequest) HasShopperName() bool`

HasShopperName returns a boolean if a field has been set.

### GetShopperReference

`func (o *StoreDetailAndSubmitRequest) GetShopperReference() string`

GetShopperReference returns the ShopperReference field if non-nil, zero value otherwise.

### GetShopperReferenceOk

`func (o *StoreDetailAndSubmitRequest) GetShopperReferenceOk() (*string, bool)`

GetShopperReferenceOk returns a tuple with the ShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperReference

`func (o *StoreDetailAndSubmitRequest) SetShopperReference(v string)`

SetShopperReference sets ShopperReference field to given value.


### GetShopperStatement

`func (o *StoreDetailAndSubmitRequest) GetShopperStatement() string`

GetShopperStatement returns the ShopperStatement field if non-nil, zero value otherwise.

### GetShopperStatementOk

`func (o *StoreDetailAndSubmitRequest) GetShopperStatementOk() (*string, bool)`

GetShopperStatementOk returns a tuple with the ShopperStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperStatement

`func (o *StoreDetailAndSubmitRequest) SetShopperStatement(v string)`

SetShopperStatement sets ShopperStatement field to given value.

### HasShopperStatement

`func (o *StoreDetailAndSubmitRequest) HasShopperStatement() bool`

HasShopperStatement returns a boolean if a field has been set.

### GetSocialSecurityNumber

`func (o *StoreDetailAndSubmitRequest) GetSocialSecurityNumber() string`

GetSocialSecurityNumber returns the SocialSecurityNumber field if non-nil, zero value otherwise.

### GetSocialSecurityNumberOk

`func (o *StoreDetailAndSubmitRequest) GetSocialSecurityNumberOk() (*string, bool)`

GetSocialSecurityNumberOk returns a tuple with the SocialSecurityNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialSecurityNumber

`func (o *StoreDetailAndSubmitRequest) SetSocialSecurityNumber(v string)`

SetSocialSecurityNumber sets SocialSecurityNumber field to given value.

### HasSocialSecurityNumber

`func (o *StoreDetailAndSubmitRequest) HasSocialSecurityNumber() bool`

HasSocialSecurityNumber returns a boolean if a field has been set.

### GetTelephoneNumber

`func (o *StoreDetailAndSubmitRequest) GetTelephoneNumber() string`

GetTelephoneNumber returns the TelephoneNumber field if non-nil, zero value otherwise.

### GetTelephoneNumberOk

`func (o *StoreDetailAndSubmitRequest) GetTelephoneNumberOk() (*string, bool)`

GetTelephoneNumberOk returns a tuple with the TelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumber

`func (o *StoreDetailAndSubmitRequest) SetTelephoneNumber(v string)`

SetTelephoneNumber sets TelephoneNumber field to given value.

### HasTelephoneNumber

`func (o *StoreDetailAndSubmitRequest) HasTelephoneNumber() bool`

HasTelephoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


