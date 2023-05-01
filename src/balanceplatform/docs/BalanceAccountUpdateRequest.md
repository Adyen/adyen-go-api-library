# BalanceAccountUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountHolderId** | Pointer to **string** | The unique identifier of the [account holder](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/accountHolders__resParam_id) associated with the balance account. | [optional] 
**DefaultCurrencyCode** | Pointer to **string** | The default currency code of this balance account, in three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes) format.  The default value is **EUR**. | [optional] 
**Description** | Pointer to **string** | A human-readable description of the balance account, maximum 300 characters. You can use this parameter to distinguish between multiple balance accounts under an account holder. | [optional] 
**Reference** | Pointer to **string** | Your reference to the balance account, maximum 150 characters. | [optional] 
**Status** | Pointer to **string** | The status of the balance account. Payment instruments linked to the balance account can only be used if the balance account status is **active**.  Possible values: **active**, **inactive**, **closed**, **suspended**. | [optional] 
**TimeZone** | Pointer to **string** | The [time zone](https://www.iana.org/time-zones) of the balance account. For example, **Europe/Amsterdam**. Defaults to the time zone of the account holder if no time zone is set. For possible values, see the [list of time zone codes](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). | [optional] 

## Methods

### NewBalanceAccountUpdateRequest

`func NewBalanceAccountUpdateRequest() *BalanceAccountUpdateRequest`

NewBalanceAccountUpdateRequest instantiates a new BalanceAccountUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceAccountUpdateRequestWithDefaults

`func NewBalanceAccountUpdateRequestWithDefaults() *BalanceAccountUpdateRequest`

NewBalanceAccountUpdateRequestWithDefaults instantiates a new BalanceAccountUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountHolderId

`func (o *BalanceAccountUpdateRequest) GetAccountHolderId() string`

GetAccountHolderId returns the AccountHolderId field if non-nil, zero value otherwise.

### GetAccountHolderIdOk

`func (o *BalanceAccountUpdateRequest) GetAccountHolderIdOk() (*string, bool)`

GetAccountHolderIdOk returns a tuple with the AccountHolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderId

`func (o *BalanceAccountUpdateRequest) SetAccountHolderId(v string)`

SetAccountHolderId sets AccountHolderId field to given value.

### HasAccountHolderId

`func (o *BalanceAccountUpdateRequest) HasAccountHolderId() bool`

HasAccountHolderId returns a boolean if a field has been set.

### GetDefaultCurrencyCode

`func (o *BalanceAccountUpdateRequest) GetDefaultCurrencyCode() string`

GetDefaultCurrencyCode returns the DefaultCurrencyCode field if non-nil, zero value otherwise.

### GetDefaultCurrencyCodeOk

`func (o *BalanceAccountUpdateRequest) GetDefaultCurrencyCodeOk() (*string, bool)`

GetDefaultCurrencyCodeOk returns a tuple with the DefaultCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCurrencyCode

`func (o *BalanceAccountUpdateRequest) SetDefaultCurrencyCode(v string)`

SetDefaultCurrencyCode sets DefaultCurrencyCode field to given value.

### HasDefaultCurrencyCode

`func (o *BalanceAccountUpdateRequest) HasDefaultCurrencyCode() bool`

HasDefaultCurrencyCode returns a boolean if a field has been set.

### GetDescription

`func (o *BalanceAccountUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BalanceAccountUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BalanceAccountUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BalanceAccountUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetReference

`func (o *BalanceAccountUpdateRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *BalanceAccountUpdateRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *BalanceAccountUpdateRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *BalanceAccountUpdateRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetStatus

`func (o *BalanceAccountUpdateRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BalanceAccountUpdateRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BalanceAccountUpdateRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BalanceAccountUpdateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeZone

`func (o *BalanceAccountUpdateRequest) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *BalanceAccountUpdateRequest) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *BalanceAccountUpdateRequest) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *BalanceAccountUpdateRequest) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


