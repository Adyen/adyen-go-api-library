# SweepConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceAccountId** | Pointer to **string** | The unique identifier of the destination or source [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts__resParam_id).   You can only use this for periodic sweep schedules such as &#x60;schedule.type&#x60; **daily** or **monthly**. | [optional] 
**Id** | **string** | The unique identifier of the sweep. | 
**MerchantAccount** | Pointer to **string** | The merchant account that will be the source of funds. You can only use this if you are processing payments with Adyen. This can only be used for sweeps of &#x60;type&#x60; **pull** and &#x60;schedule.type&#x60; **balance**. | [optional] 
**Schedule** | [**SweepConfigurationSchedule**](SweepConfigurationSchedule.md) |  | 
**Status** | Pointer to **string** | The status of the sweep. If not provided, by default, this is set to **active**.  Possible values:    * **active**:  the sweep is enabled and funds will be pulled in or pushed out based on the defined configuration.    * **inactive**: the sweep is disabled and cannot be triggered.    | [optional] 
**SweepAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**TargetAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**TransferInstrumentId** | Pointer to **string** | The unique identifier of the destination or source [transfer instrument](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/transferInstruments__resParam_id).  You can also use this in combination with a &#x60;merchantAccount&#x60; and a &#x60;type&#x60; **pull** to start a direct debit request from the source transfer instrument. To use this feature, reach out to your Adyen contact. | [optional] 
**TriggerAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**Type** | Pointer to **string** | The direction of sweep, whether pushing out or pulling in funds to the balance account. If not provided, by default, this is set to **push**.  Possible values:   * **push**: _push out funds_ to a destination balance account or transfer instrument.   * **pull**: _pull in funds_ from a source merchant account, transfer instrument, or balance account. | [optional] [default to "push"]

## Methods

### NewSweepConfiguration

`func NewSweepConfiguration(id string, schedule SweepConfigurationSchedule, ) *SweepConfiguration`

NewSweepConfiguration instantiates a new SweepConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSweepConfigurationWithDefaults

`func NewSweepConfigurationWithDefaults() *SweepConfiguration`

NewSweepConfigurationWithDefaults instantiates a new SweepConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceAccountId

`func (o *SweepConfiguration) GetBalanceAccountId() string`

GetBalanceAccountId returns the BalanceAccountId field if non-nil, zero value otherwise.

### GetBalanceAccountIdOk

`func (o *SweepConfiguration) GetBalanceAccountIdOk() (*string, bool)`

GetBalanceAccountIdOk returns a tuple with the BalanceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAccountId

`func (o *SweepConfiguration) SetBalanceAccountId(v string)`

SetBalanceAccountId sets BalanceAccountId field to given value.

### HasBalanceAccountId

`func (o *SweepConfiguration) HasBalanceAccountId() bool`

HasBalanceAccountId returns a boolean if a field has been set.

### GetId

`func (o *SweepConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SweepConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SweepConfiguration) SetId(v string)`

SetId sets Id field to given value.


### GetMerchantAccount

`func (o *SweepConfiguration) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *SweepConfiguration) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *SweepConfiguration) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.

### HasMerchantAccount

`func (o *SweepConfiguration) HasMerchantAccount() bool`

HasMerchantAccount returns a boolean if a field has been set.

### GetSchedule

`func (o *SweepConfiguration) GetSchedule() SweepConfigurationSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *SweepConfiguration) GetScheduleOk() (*SweepConfigurationSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *SweepConfiguration) SetSchedule(v SweepConfigurationSchedule)`

SetSchedule sets Schedule field to given value.


### GetStatus

`func (o *SweepConfiguration) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SweepConfiguration) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SweepConfiguration) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SweepConfiguration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSweepAmount

`func (o *SweepConfiguration) GetSweepAmount() Amount`

GetSweepAmount returns the SweepAmount field if non-nil, zero value otherwise.

### GetSweepAmountOk

`func (o *SweepConfiguration) GetSweepAmountOk() (*Amount, bool)`

GetSweepAmountOk returns a tuple with the SweepAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSweepAmount

`func (o *SweepConfiguration) SetSweepAmount(v Amount)`

SetSweepAmount sets SweepAmount field to given value.

### HasSweepAmount

`func (o *SweepConfiguration) HasSweepAmount() bool`

HasSweepAmount returns a boolean if a field has been set.

### GetTargetAmount

`func (o *SweepConfiguration) GetTargetAmount() Amount`

GetTargetAmount returns the TargetAmount field if non-nil, zero value otherwise.

### GetTargetAmountOk

`func (o *SweepConfiguration) GetTargetAmountOk() (*Amount, bool)`

GetTargetAmountOk returns a tuple with the TargetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAmount

`func (o *SweepConfiguration) SetTargetAmount(v Amount)`

SetTargetAmount sets TargetAmount field to given value.

### HasTargetAmount

`func (o *SweepConfiguration) HasTargetAmount() bool`

HasTargetAmount returns a boolean if a field has been set.

### GetTransferInstrumentId

`func (o *SweepConfiguration) GetTransferInstrumentId() string`

GetTransferInstrumentId returns the TransferInstrumentId field if non-nil, zero value otherwise.

### GetTransferInstrumentIdOk

`func (o *SweepConfiguration) GetTransferInstrumentIdOk() (*string, bool)`

GetTransferInstrumentIdOk returns a tuple with the TransferInstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferInstrumentId

`func (o *SweepConfiguration) SetTransferInstrumentId(v string)`

SetTransferInstrumentId sets TransferInstrumentId field to given value.

### HasTransferInstrumentId

`func (o *SweepConfiguration) HasTransferInstrumentId() bool`

HasTransferInstrumentId returns a boolean if a field has been set.

### GetTriggerAmount

`func (o *SweepConfiguration) GetTriggerAmount() Amount`

GetTriggerAmount returns the TriggerAmount field if non-nil, zero value otherwise.

### GetTriggerAmountOk

`func (o *SweepConfiguration) GetTriggerAmountOk() (*Amount, bool)`

GetTriggerAmountOk returns a tuple with the TriggerAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerAmount

`func (o *SweepConfiguration) SetTriggerAmount(v Amount)`

SetTriggerAmount sets TriggerAmount field to given value.

### HasTriggerAmount

`func (o *SweepConfiguration) HasTriggerAmount() bool`

HasTriggerAmount returns a boolean if a field has been set.

### GetType

`func (o *SweepConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SweepConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SweepConfiguration) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SweepConfiguration) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


