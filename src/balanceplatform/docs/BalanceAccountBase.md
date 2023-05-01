# BalanceAccountBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountHolderId** | **string** | The unique identifier of the [account holder](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/accountHolders__resParam_id) associated with the balance account. | 
**DefaultCurrencyCode** | Pointer to **string** | The default three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes) of the balance account. The default value is **EUR**. | [optional] 
**Description** | Pointer to **string** | A human-readable description of the balance account, maximum 300 characters. You can use this parameter to distinguish between multiple balance accounts under an account holder. | [optional] 
**Id** | **string** | The unique identifier of the balance account. | 
**Reference** | Pointer to **string** | Your reference for the balance account, maximum 150 characters. | [optional] 
**Status** | Pointer to **string** | The status of the balance account, set to **active** by default.   | [optional] 
**TimeZone** | Pointer to **string** | The [time zone](https://www.iana.org/time-zones) of the balance account. For example, **Europe/Amsterdam**. Defaults to the time zone of the account holder if no time zone is set. For possible values, see the [list of time zone codes](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). | [optional] 

## Methods

### NewBalanceAccountBase

`func NewBalanceAccountBase(accountHolderId string, id string, ) *BalanceAccountBase`

NewBalanceAccountBase instantiates a new BalanceAccountBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceAccountBaseWithDefaults

`func NewBalanceAccountBaseWithDefaults() *BalanceAccountBase`

NewBalanceAccountBaseWithDefaults instantiates a new BalanceAccountBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountHolderId

`func (o *BalanceAccountBase) GetAccountHolderId() string`

GetAccountHolderId returns the AccountHolderId field if non-nil, zero value otherwise.

### GetAccountHolderIdOk

`func (o *BalanceAccountBase) GetAccountHolderIdOk() (*string, bool)`

GetAccountHolderIdOk returns a tuple with the AccountHolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderId

`func (o *BalanceAccountBase) SetAccountHolderId(v string)`

SetAccountHolderId sets AccountHolderId field to given value.


### GetDefaultCurrencyCode

`func (o *BalanceAccountBase) GetDefaultCurrencyCode() string`

GetDefaultCurrencyCode returns the DefaultCurrencyCode field if non-nil, zero value otherwise.

### GetDefaultCurrencyCodeOk

`func (o *BalanceAccountBase) GetDefaultCurrencyCodeOk() (*string, bool)`

GetDefaultCurrencyCodeOk returns a tuple with the DefaultCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCurrencyCode

`func (o *BalanceAccountBase) SetDefaultCurrencyCode(v string)`

SetDefaultCurrencyCode sets DefaultCurrencyCode field to given value.

### HasDefaultCurrencyCode

`func (o *BalanceAccountBase) HasDefaultCurrencyCode() bool`

HasDefaultCurrencyCode returns a boolean if a field has been set.

### GetDescription

`func (o *BalanceAccountBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BalanceAccountBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BalanceAccountBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BalanceAccountBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *BalanceAccountBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BalanceAccountBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BalanceAccountBase) SetId(v string)`

SetId sets Id field to given value.


### GetReference

`func (o *BalanceAccountBase) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *BalanceAccountBase) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *BalanceAccountBase) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *BalanceAccountBase) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetStatus

`func (o *BalanceAccountBase) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BalanceAccountBase) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BalanceAccountBase) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BalanceAccountBase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeZone

`func (o *BalanceAccountBase) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *BalanceAccountBase) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *BalanceAccountBase) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *BalanceAccountBase) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


