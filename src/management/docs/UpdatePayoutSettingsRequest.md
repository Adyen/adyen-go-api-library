# UpdatePayoutSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Indicates if payouts to this bank account are enabled. Default: **true**.  To receive payouts into this bank account, both &#x60;enabled&#x60; and &#x60;allowed&#x60; must be **true**. | [optional] 
**EnabledFromDate** | Pointer to **string** | The date when Adyen starts paying out to this bank account.  Format: [ISO 8601](https://www.w3.org/TR/NOTE-datetime), for example, **2019-11-23T12:25:28Z** or **2020-05-27T20:25:28+08:00**.  If not specified, the &#x60;enabled&#x60; field indicates if payouts are enabled for this bank account.  If a date is specified and:  * &#x60;enabled&#x60;: **true**, payouts are enabled starting the specified date. * &#x60;enabled&#x60;: **false**, payouts are disabled until the specified date. On the specified date, &#x60;enabled&#x60; changes to **true** and this field is reset to **null**. | [optional] 
**Priority** | Pointer to **string** | Determines how long it takes for the funds to reach the bank account. Adyen pays out based on the [payout frequency](https://docs.adyen.com/account/getting-paid#payout-frequency). Depending on the currencies and banks involved in transferring the money, it may take up to three days for the payout funds to arrive in the bank account.   Possible values: * **first**: same day. * **urgent**: the next day. * **normal**: between 1 and 3 days. | [optional] 

## Methods

### NewUpdatePayoutSettingsRequest

`func NewUpdatePayoutSettingsRequest() *UpdatePayoutSettingsRequest`

NewUpdatePayoutSettingsRequest instantiates a new UpdatePayoutSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePayoutSettingsRequestWithDefaults

`func NewUpdatePayoutSettingsRequestWithDefaults() *UpdatePayoutSettingsRequest`

NewUpdatePayoutSettingsRequestWithDefaults instantiates a new UpdatePayoutSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdatePayoutSettingsRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdatePayoutSettingsRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdatePayoutSettingsRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdatePayoutSettingsRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEnabledFromDate

`func (o *UpdatePayoutSettingsRequest) GetEnabledFromDate() string`

GetEnabledFromDate returns the EnabledFromDate field if non-nil, zero value otherwise.

### GetEnabledFromDateOk

`func (o *UpdatePayoutSettingsRequest) GetEnabledFromDateOk() (*string, bool)`

GetEnabledFromDateOk returns a tuple with the EnabledFromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledFromDate

`func (o *UpdatePayoutSettingsRequest) SetEnabledFromDate(v string)`

SetEnabledFromDate sets EnabledFromDate field to given value.

### HasEnabledFromDate

`func (o *UpdatePayoutSettingsRequest) HasEnabledFromDate() bool`

HasEnabledFromDate returns a boolean if a field has been set.

### GetPriority

`func (o *UpdatePayoutSettingsRequest) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *UpdatePayoutSettingsRequest) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *UpdatePayoutSettingsRequest) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *UpdatePayoutSettingsRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


