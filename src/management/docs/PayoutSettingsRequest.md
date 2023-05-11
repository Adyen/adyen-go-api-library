# PayoutSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Indicates if payouts to this bank account are enabled. Default: **true**.  To receive payouts into this bank account, both &#x60;enabled&#x60; and &#x60;allowed&#x60; must be **true**. | [optional] 
**EnabledFromDate** | Pointer to **string** | The date when Adyen starts paying out to this bank account.  Format: [ISO 8601](https://www.w3.org/TR/NOTE-datetime), for example, **2019-11-23T12:25:28Z** or **2020-05-27T20:25:28+08:00**.  If not specified, the &#x60;enabled&#x60; field indicates if payouts are enabled for this bank account.  If a date is specified and:  * &#x60;enabled&#x60;: **true**, payouts are enabled starting the specified date. * &#x60;enabled&#x60;: **false**, payouts are disabled until the specified date. On the specified date, &#x60;enabled&#x60; changes to **true** and this field is reset to **null**. | [optional] 
**TransferInstrumentId** | **string** | The unique identifier of the [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments) that contains the details of the bank account. | 

## Methods

### NewPayoutSettingsRequest

`func NewPayoutSettingsRequest(transferInstrumentId string, ) *PayoutSettingsRequest`

NewPayoutSettingsRequest instantiates a new PayoutSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutSettingsRequestWithDefaults

`func NewPayoutSettingsRequestWithDefaults() *PayoutSettingsRequest`

NewPayoutSettingsRequestWithDefaults instantiates a new PayoutSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *PayoutSettingsRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PayoutSettingsRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PayoutSettingsRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PayoutSettingsRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEnabledFromDate

`func (o *PayoutSettingsRequest) GetEnabledFromDate() string`

GetEnabledFromDate returns the EnabledFromDate field if non-nil, zero value otherwise.

### GetEnabledFromDateOk

`func (o *PayoutSettingsRequest) GetEnabledFromDateOk() (*string, bool)`

GetEnabledFromDateOk returns a tuple with the EnabledFromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledFromDate

`func (o *PayoutSettingsRequest) SetEnabledFromDate(v string)`

SetEnabledFromDate sets EnabledFromDate field to given value.

### HasEnabledFromDate

`func (o *PayoutSettingsRequest) HasEnabledFromDate() bool`

HasEnabledFromDate returns a boolean if a field has been set.

### GetTransferInstrumentId

`func (o *PayoutSettingsRequest) GetTransferInstrumentId() string`

GetTransferInstrumentId returns the TransferInstrumentId field if non-nil, zero value otherwise.

### GetTransferInstrumentIdOk

`func (o *PayoutSettingsRequest) GetTransferInstrumentIdOk() (*string, bool)`

GetTransferInstrumentIdOk returns a tuple with the TransferInstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferInstrumentId

`func (o *PayoutSettingsRequest) SetTransferInstrumentId(v string)`

SetTransferInstrumentId sets TransferInstrumentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


