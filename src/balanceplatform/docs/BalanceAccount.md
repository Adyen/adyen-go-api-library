# BalanceAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountHolderId** | **string** | The unique identifier of the [account holder](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/accountHolders__resParam_id) associated with the balance account. | 
**Balances** | Pointer to [**[]Balance**](Balance.md) | List of balances with the amount and currency. | [optional] 
**DefaultCurrencyCode** | Pointer to **string** | The default three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes) of the balance account. The default value is **EUR**. | [optional] 
**Description** | Pointer to **string** | A human-readable description of the balance account, maximum 300 characters. You can use this parameter to distinguish between multiple balance accounts under an account holder. | [optional] 
**Id** | **string** | The unique identifier of the balance account. | 
**Reference** | Pointer to **string** | Your reference for the balance account, maximum 150 characters. | [optional] 
**Status** | Pointer to **string** | The status of the balance account, set to **active** by default.   | [optional] 
**TimeZone** | Pointer to **string** | The time zone of the balance account. For example, **Europe/Amsterdam**. Defaults to the time zone of the account holder if no time zone is set. For possible values, see the [list of time zone codes](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). | [optional] 

## Methods

### NewBalanceAccount

`func NewBalanceAccount(accountHolderId string, id string, ) *BalanceAccount`

NewBalanceAccount instantiates a new BalanceAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceAccountWithDefaults

`func NewBalanceAccountWithDefaults() *BalanceAccount`

NewBalanceAccountWithDefaults instantiates a new BalanceAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountHolderId

`func (o *BalanceAccount) GetAccountHolderId() string`

GetAccountHolderId returns the AccountHolderId field if non-nil, zero value otherwise.

### GetAccountHolderIdOk

`func (o *BalanceAccount) GetAccountHolderIdOk() (*string, bool)`

GetAccountHolderIdOk returns a tuple with the AccountHolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderId

`func (o *BalanceAccount) SetAccountHolderId(v string)`

SetAccountHolderId sets AccountHolderId field to given value.


### GetBalances

`func (o *BalanceAccount) GetBalances() []Balance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *BalanceAccount) GetBalancesOk() (*[]Balance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *BalanceAccount) SetBalances(v []Balance)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *BalanceAccount) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetDefaultCurrencyCode

`func (o *BalanceAccount) GetDefaultCurrencyCode() string`

GetDefaultCurrencyCode returns the DefaultCurrencyCode field if non-nil, zero value otherwise.

### GetDefaultCurrencyCodeOk

`func (o *BalanceAccount) GetDefaultCurrencyCodeOk() (*string, bool)`

GetDefaultCurrencyCodeOk returns a tuple with the DefaultCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCurrencyCode

`func (o *BalanceAccount) SetDefaultCurrencyCode(v string)`

SetDefaultCurrencyCode sets DefaultCurrencyCode field to given value.

### HasDefaultCurrencyCode

`func (o *BalanceAccount) HasDefaultCurrencyCode() bool`

HasDefaultCurrencyCode returns a boolean if a field has been set.

### GetDescription

`func (o *BalanceAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BalanceAccount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BalanceAccount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BalanceAccount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *BalanceAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BalanceAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BalanceAccount) SetId(v string)`

SetId sets Id field to given value.


### GetReference

`func (o *BalanceAccount) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *BalanceAccount) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *BalanceAccount) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *BalanceAccount) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetStatus

`func (o *BalanceAccount) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BalanceAccount) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BalanceAccount) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BalanceAccount) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeZone

`func (o *BalanceAccount) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *BalanceAccount) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *BalanceAccount) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *BalanceAccount) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


