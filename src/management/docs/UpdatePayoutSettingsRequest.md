# UpdatePayoutSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Indicates if payouts to this bank account are enabled. Default: **true**.  To receive payouts into this bank account, both &#x60;enabled&#x60; and &#x60;allowed&#x60; must be **true**. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


