# SubmitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **map[string]string** | This field contains additional data, which may be required for a particular request. | [optional] 
**Amount** | [**Amount**](Amount.md) |  | 
**DateOfBirth** | Pointer to **time.Time** | The date of birth. Format: ISO-8601; example: YYYY-MM-DD  For Paysafecard it must be the same as used when registering the Paysafecard account.  &gt; This field is mandatory for natural persons.  &gt; This field is required to update the existing &#x60;dateOfBirth&#x60; that is associated with this recurring contract. | [optional] 
**EntityType** | Pointer to **string** | The type of the entity the payout is processed for.  Allowed values: * NaturalPerson * Company &gt; This field is required to update the existing &#x60;entityType&#x60; that is associated with this recurring contract. | [optional] 
**FraudOffset** | Pointer to **int32** | An integer value that is added to the normal fraud score. The value can be either positive or negative. | [optional] 
**MerchantAccount** | **string** | The merchant account identifier you want to process the transaction request with. | 
**Nationality** | Pointer to **string** | The shopper&#39;s nationality.  A valid value is an ISO 2-character country code (e.g. &#39;NL&#39;).  &gt; This field is required to update the existing nationality that is associated with this recurring contract. | [optional] 
**Recurring** | [**Recurring**](Recurring.md) |  | 
**Reference** | **string** | The merchant reference for this payout. This reference will be used in all communication to the merchant about the status of the payout. Although it is a good idea to make sure it is unique, this is not a requirement. | 
**SelectedRecurringDetailReference** | **string** | This is the &#x60;recurringDetailReference&#x60; you want to use for this payout.  You can use the value LATEST to select the most recently used recurring detail. | 
**ShopperEmail** | **string** | The shopper&#39;s email address. | 
**ShopperName** | Pointer to [**Name**](Name.md) |  | [optional] 
**ShopperReference** | **string** | The shopper&#39;s reference for the payout transaction. | 
**ShopperStatement** | Pointer to **string** | The description of this payout. This description is shown on the bank statement of the shopper (if this is supported by the chosen payment method). | [optional] 
**SocialSecurityNumber** | Pointer to **string** | The shopper&#39;s social security number. | [optional] 

## Methods

### NewSubmitRequest

`func NewSubmitRequest(amount Amount, merchantAccount string, recurring Recurring, reference string, selectedRecurringDetailReference string, shopperEmail string, shopperReference string, ) *SubmitRequest`

NewSubmitRequest instantiates a new SubmitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitRequestWithDefaults

`func NewSubmitRequestWithDefaults() *SubmitRequest`

NewSubmitRequestWithDefaults instantiates a new SubmitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *SubmitRequest) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *SubmitRequest) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *SubmitRequest) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *SubmitRequest) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetAmount

`func (o *SubmitRequest) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SubmitRequest) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SubmitRequest) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetDateOfBirth

`func (o *SubmitRequest) GetDateOfBirth() time.Time`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *SubmitRequest) GetDateOfBirthOk() (*time.Time, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *SubmitRequest) SetDateOfBirth(v time.Time)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *SubmitRequest) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetEntityType

`func (o *SubmitRequest) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *SubmitRequest) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *SubmitRequest) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *SubmitRequest) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetFraudOffset

`func (o *SubmitRequest) GetFraudOffset() int32`

GetFraudOffset returns the FraudOffset field if non-nil, zero value otherwise.

### GetFraudOffsetOk

`func (o *SubmitRequest) GetFraudOffsetOk() (*int32, bool)`

GetFraudOffsetOk returns a tuple with the FraudOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudOffset

`func (o *SubmitRequest) SetFraudOffset(v int32)`

SetFraudOffset sets FraudOffset field to given value.

### HasFraudOffset

`func (o *SubmitRequest) HasFraudOffset() bool`

HasFraudOffset returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *SubmitRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *SubmitRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *SubmitRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetNationality

`func (o *SubmitRequest) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *SubmitRequest) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *SubmitRequest) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *SubmitRequest) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetRecurring

`func (o *SubmitRequest) GetRecurring() Recurring`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *SubmitRequest) GetRecurringOk() (*Recurring, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *SubmitRequest) SetRecurring(v Recurring)`

SetRecurring sets Recurring field to given value.


### GetReference

`func (o *SubmitRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *SubmitRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *SubmitRequest) SetReference(v string)`

SetReference sets Reference field to given value.


### GetSelectedRecurringDetailReference

`func (o *SubmitRequest) GetSelectedRecurringDetailReference() string`

GetSelectedRecurringDetailReference returns the SelectedRecurringDetailReference field if non-nil, zero value otherwise.

### GetSelectedRecurringDetailReferenceOk

`func (o *SubmitRequest) GetSelectedRecurringDetailReferenceOk() (*string, bool)`

GetSelectedRecurringDetailReferenceOk returns a tuple with the SelectedRecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRecurringDetailReference

`func (o *SubmitRequest) SetSelectedRecurringDetailReference(v string)`

SetSelectedRecurringDetailReference sets SelectedRecurringDetailReference field to given value.


### GetShopperEmail

`func (o *SubmitRequest) GetShopperEmail() string`

GetShopperEmail returns the ShopperEmail field if non-nil, zero value otherwise.

### GetShopperEmailOk

`func (o *SubmitRequest) GetShopperEmailOk() (*string, bool)`

GetShopperEmailOk returns a tuple with the ShopperEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperEmail

`func (o *SubmitRequest) SetShopperEmail(v string)`

SetShopperEmail sets ShopperEmail field to given value.


### GetShopperName

`func (o *SubmitRequest) GetShopperName() Name`

GetShopperName returns the ShopperName field if non-nil, zero value otherwise.

### GetShopperNameOk

`func (o *SubmitRequest) GetShopperNameOk() (*Name, bool)`

GetShopperNameOk returns a tuple with the ShopperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperName

`func (o *SubmitRequest) SetShopperName(v Name)`

SetShopperName sets ShopperName field to given value.

### HasShopperName

`func (o *SubmitRequest) HasShopperName() bool`

HasShopperName returns a boolean if a field has been set.

### GetShopperReference

`func (o *SubmitRequest) GetShopperReference() string`

GetShopperReference returns the ShopperReference field if non-nil, zero value otherwise.

### GetShopperReferenceOk

`func (o *SubmitRequest) GetShopperReferenceOk() (*string, bool)`

GetShopperReferenceOk returns a tuple with the ShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperReference

`func (o *SubmitRequest) SetShopperReference(v string)`

SetShopperReference sets ShopperReference field to given value.


### GetShopperStatement

`func (o *SubmitRequest) GetShopperStatement() string`

GetShopperStatement returns the ShopperStatement field if non-nil, zero value otherwise.

### GetShopperStatementOk

`func (o *SubmitRequest) GetShopperStatementOk() (*string, bool)`

GetShopperStatementOk returns a tuple with the ShopperStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperStatement

`func (o *SubmitRequest) SetShopperStatement(v string)`

SetShopperStatement sets ShopperStatement field to given value.

### HasShopperStatement

`func (o *SubmitRequest) HasShopperStatement() bool`

HasShopperStatement returns a boolean if a field has been set.

### GetSocialSecurityNumber

`func (o *SubmitRequest) GetSocialSecurityNumber() string`

GetSocialSecurityNumber returns the SocialSecurityNumber field if non-nil, zero value otherwise.

### GetSocialSecurityNumberOk

`func (o *SubmitRequest) GetSocialSecurityNumberOk() (*string, bool)`

GetSocialSecurityNumberOk returns a tuple with the SocialSecurityNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialSecurityNumber

`func (o *SubmitRequest) SetSocialSecurityNumber(v string)`

SetSocialSecurityNumber sets SocialSecurityNumber field to given value.

### HasSocialSecurityNumber

`func (o *SubmitRequest) HasSocialSecurityNumber() bool`

HasSocialSecurityNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


