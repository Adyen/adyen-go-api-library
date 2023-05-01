# BalanceAccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountHolderId** | **string** | The unique identifier of the [account holder](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/accountHolders__resParam_id) associated with the balance account. | 
**DefaultCurrencyCode** | Pointer to **string** | The default three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes) of the balance account. The default value is **EUR**. | [optional] 
**Description** | Pointer to **string** | A human-readable description of the balance account, maximum 300 characters. You can use this parameter to distinguish between multiple balance accounts under an account holder. | [optional] 
**Reference** | Pointer to **string** | Your reference for the balance account, maximum 150 characters. | [optional] 
**TimeZone** | Pointer to **string** | The [time zone](https://www.iana.org/time-zones) of the balance account. For example, **Europe/Amsterdam**. Defaults to the time zone of the account holder if no time zone is set. For possible values, see the [list of time zone codes](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). | [optional] 

## Methods

### NewBalanceAccountInfo

`func NewBalanceAccountInfo(accountHolderId string, ) *BalanceAccountInfo`

NewBalanceAccountInfo instantiates a new BalanceAccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceAccountInfoWithDefaults

`func NewBalanceAccountInfoWithDefaults() *BalanceAccountInfo`

NewBalanceAccountInfoWithDefaults instantiates a new BalanceAccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountHolderId

`func (o *BalanceAccountInfo) GetAccountHolderId() string`

GetAccountHolderId returns the AccountHolderId field if non-nil, zero value otherwise.

### GetAccountHolderIdOk

`func (o *BalanceAccountInfo) GetAccountHolderIdOk() (*string, bool)`

GetAccountHolderIdOk returns a tuple with the AccountHolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderId

`func (o *BalanceAccountInfo) SetAccountHolderId(v string)`

SetAccountHolderId sets AccountHolderId field to given value.


### GetDefaultCurrencyCode

`func (o *BalanceAccountInfo) GetDefaultCurrencyCode() string`

GetDefaultCurrencyCode returns the DefaultCurrencyCode field if non-nil, zero value otherwise.

### GetDefaultCurrencyCodeOk

`func (o *BalanceAccountInfo) GetDefaultCurrencyCodeOk() (*string, bool)`

GetDefaultCurrencyCodeOk returns a tuple with the DefaultCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCurrencyCode

`func (o *BalanceAccountInfo) SetDefaultCurrencyCode(v string)`

SetDefaultCurrencyCode sets DefaultCurrencyCode field to given value.

### HasDefaultCurrencyCode

`func (o *BalanceAccountInfo) HasDefaultCurrencyCode() bool`

HasDefaultCurrencyCode returns a boolean if a field has been set.

### GetDescription

`func (o *BalanceAccountInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BalanceAccountInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BalanceAccountInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BalanceAccountInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetReference

`func (o *BalanceAccountInfo) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *BalanceAccountInfo) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *BalanceAccountInfo) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *BalanceAccountInfo) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetTimeZone

`func (o *BalanceAccountInfo) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *BalanceAccountInfo) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *BalanceAccountInfo) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *BalanceAccountInfo) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


