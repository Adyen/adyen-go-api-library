# PayoutSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | Pointer to **bool** | Indicates if payouts to the bank account are allowed. This value is set automatically based on the status of the verification process. The value is:  * **true** if &#x60;verificationStatus&#x60; is **valid**. * **false** for all other values. | [optional] 
**Enabled** | Pointer to **bool** | Indicates if payouts to this bank account are enabled. Default: **true**.  To receive payouts into this bank account, both &#x60;enabled&#x60; and &#x60;allowed&#x60; must be **true**. | [optional] 
**EnabledFromDate** | Pointer to **string** | The date when Adyen starts paying out to this bank account.  Format: [ISO 8601](https://www.w3.org/TR/NOTE-datetime), for example, **2019-11-23T12:25:28Z** or **2020-05-27T20:25:28+08:00**.  If not specified, the &#x60;enabled&#x60; field indicates if payouts are enabled for this bank account.  If a date is specified and:  * &#x60;enabled&#x60;: **true**, payouts are enabled starting the specified date. * &#x60;enabled&#x60;: **false**, payouts are disabled until the specified date. On the specified date, &#x60;enabled&#x60; changes to **true** and this field is reset to **null**. | [optional] 
**Id** | **string** | The unique identifier of the payout setting. | 
**Priority** | Pointer to **string** | Determines how long it takes for the funds to reach the bank account. Adyen pays out based on the [payout frequency](https://docs.adyen.com/account/getting-paid#payout-frequency). Depending on the currencies and banks involved in transferring the money, it may take up to three days for the payout funds to arrive in the bank account.   Possible values: * **first**: same day. * **urgent**: the next day. * **normal**: between 1 and 3 days. | [optional] 
**TransferInstrumentId** | **string** | The unique identifier of the [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments) that contains the details of the bank account. | 
**VerificationStatus** | Pointer to **string** | The status of the verification process for the bank account.  Possible values: * **valid**: the verification was successful. * **pending**: the verification is in progress. * **invalid**: the information provided is not complete. * **rejected**:  there are reasons to refuse working with this entity. | [optional] 

## Methods

### NewPayoutSettings

`func NewPayoutSettings(id string, transferInstrumentId string, ) *PayoutSettings`

NewPayoutSettings instantiates a new PayoutSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutSettingsWithDefaults

`func NewPayoutSettingsWithDefaults() *PayoutSettings`

NewPayoutSettingsWithDefaults instantiates a new PayoutSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *PayoutSettings) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *PayoutSettings) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *PayoutSettings) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.

### HasAllowed

`func (o *PayoutSettings) HasAllowed() bool`

HasAllowed returns a boolean if a field has been set.

### GetEnabled

`func (o *PayoutSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PayoutSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PayoutSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PayoutSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEnabledFromDate

`func (o *PayoutSettings) GetEnabledFromDate() string`

GetEnabledFromDate returns the EnabledFromDate field if non-nil, zero value otherwise.

### GetEnabledFromDateOk

`func (o *PayoutSettings) GetEnabledFromDateOk() (*string, bool)`

GetEnabledFromDateOk returns a tuple with the EnabledFromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledFromDate

`func (o *PayoutSettings) SetEnabledFromDate(v string)`

SetEnabledFromDate sets EnabledFromDate field to given value.

### HasEnabledFromDate

`func (o *PayoutSettings) HasEnabledFromDate() bool`

HasEnabledFromDate returns a boolean if a field has been set.

### GetId

`func (o *PayoutSettings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PayoutSettings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PayoutSettings) SetId(v string)`

SetId sets Id field to given value.


### GetPriority

`func (o *PayoutSettings) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PayoutSettings) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PayoutSettings) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PayoutSettings) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTransferInstrumentId

`func (o *PayoutSettings) GetTransferInstrumentId() string`

GetTransferInstrumentId returns the TransferInstrumentId field if non-nil, zero value otherwise.

### GetTransferInstrumentIdOk

`func (o *PayoutSettings) GetTransferInstrumentIdOk() (*string, bool)`

GetTransferInstrumentIdOk returns a tuple with the TransferInstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferInstrumentId

`func (o *PayoutSettings) SetTransferInstrumentId(v string)`

SetTransferInstrumentId sets TransferInstrumentId field to given value.


### GetVerificationStatus

`func (o *PayoutSettings) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *PayoutSettings) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *PayoutSettings) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *PayoutSettings) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


