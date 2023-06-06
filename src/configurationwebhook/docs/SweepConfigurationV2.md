# SweepConfigurationV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Counterparty** | [**SweepCounterparty**](SweepCounterparty.md) |  | 
**Currency** | **string** | The three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes) in uppercase. For example, **EUR**.  The sweep currency must match any of the [balances currencies](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/get/balanceAccounts/{id}__resParam_balances). | 
**Description** | Pointer to **string** | The message that will be used in the sweep transfer&#39;s description body with a maximum length of 140 characters.  If the message is longer after replacing placeholders, the message will be cut off at 140 characters. | [optional] 
**Id** | **string** | The unique identifier of the sweep. | 
**Reason** | Pointer to **string** | The reason for disabling the sweep. | [optional] 
**Schedule** | [**SweepConfigurationSchedule**](SweepConfigurationSchedule.md) |  | 
**Status** | Pointer to **string** | The status of the sweep. If not provided, by default, this is set to **active**.  Possible values:    * **active**:  the sweep is enabled and funds will be pulled in or pushed out based on the defined configuration.    * **inactive**: the sweep is disabled and cannot be triggered.    | [optional] 
**SweepAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**TargetAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**TriggerAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**Type** | Pointer to **string** | The direction of sweep, whether pushing out or pulling in funds to the balance account. If not provided, by default, this is set to **push**.  Possible values:   * **push**: _push out funds_ to a destination balance account or transfer instrument.   * **pull**: _pull in funds_ from a source merchant account, transfer instrument, or balance account. | [optional] [default to "push"]

## Methods

### NewSweepConfigurationV2

`func NewSweepConfigurationV2(counterparty SweepCounterparty, currency string, id string, schedule SweepConfigurationSchedule, ) *SweepConfigurationV2`

NewSweepConfigurationV2 instantiates a new SweepConfigurationV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSweepConfigurationV2WithDefaults

`func NewSweepConfigurationV2WithDefaults() *SweepConfigurationV2`

NewSweepConfigurationV2WithDefaults instantiates a new SweepConfigurationV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounterparty

`func (o *SweepConfigurationV2) GetCounterparty() SweepCounterparty`

GetCounterparty returns the Counterparty field if non-nil, zero value otherwise.

### GetCounterpartyOk

`func (o *SweepConfigurationV2) GetCounterpartyOk() (*SweepCounterparty, bool)`

GetCounterpartyOk returns a tuple with the Counterparty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterparty

`func (o *SweepConfigurationV2) SetCounterparty(v SweepCounterparty)`

SetCounterparty sets Counterparty field to given value.


### GetCurrency

`func (o *SweepConfigurationV2) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SweepConfigurationV2) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SweepConfigurationV2) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *SweepConfigurationV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SweepConfigurationV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SweepConfigurationV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SweepConfigurationV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *SweepConfigurationV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SweepConfigurationV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SweepConfigurationV2) SetId(v string)`

SetId sets Id field to given value.


### GetReason

`func (o *SweepConfigurationV2) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SweepConfigurationV2) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SweepConfigurationV2) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SweepConfigurationV2) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSchedule

`func (o *SweepConfigurationV2) GetSchedule() SweepConfigurationSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *SweepConfigurationV2) GetScheduleOk() (*SweepConfigurationSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *SweepConfigurationV2) SetSchedule(v SweepConfigurationSchedule)`

SetSchedule sets Schedule field to given value.


### GetStatus

`func (o *SweepConfigurationV2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SweepConfigurationV2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SweepConfigurationV2) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SweepConfigurationV2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSweepAmount

`func (o *SweepConfigurationV2) GetSweepAmount() Amount`

GetSweepAmount returns the SweepAmount field if non-nil, zero value otherwise.

### GetSweepAmountOk

`func (o *SweepConfigurationV2) GetSweepAmountOk() (*Amount, bool)`

GetSweepAmountOk returns a tuple with the SweepAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSweepAmount

`func (o *SweepConfigurationV2) SetSweepAmount(v Amount)`

SetSweepAmount sets SweepAmount field to given value.

### HasSweepAmount

`func (o *SweepConfigurationV2) HasSweepAmount() bool`

HasSweepAmount returns a boolean if a field has been set.

### GetTargetAmount

`func (o *SweepConfigurationV2) GetTargetAmount() Amount`

GetTargetAmount returns the TargetAmount field if non-nil, zero value otherwise.

### GetTargetAmountOk

`func (o *SweepConfigurationV2) GetTargetAmountOk() (*Amount, bool)`

GetTargetAmountOk returns a tuple with the TargetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAmount

`func (o *SweepConfigurationV2) SetTargetAmount(v Amount)`

SetTargetAmount sets TargetAmount field to given value.

### HasTargetAmount

`func (o *SweepConfigurationV2) HasTargetAmount() bool`

HasTargetAmount returns a boolean if a field has been set.

### GetTriggerAmount

`func (o *SweepConfigurationV2) GetTriggerAmount() Amount`

GetTriggerAmount returns the TriggerAmount field if non-nil, zero value otherwise.

### GetTriggerAmountOk

`func (o *SweepConfigurationV2) GetTriggerAmountOk() (*Amount, bool)`

GetTriggerAmountOk returns a tuple with the TriggerAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerAmount

`func (o *SweepConfigurationV2) SetTriggerAmount(v Amount)`

SetTriggerAmount sets TriggerAmount field to given value.

### HasTriggerAmount

`func (o *SweepConfigurationV2) HasTriggerAmount() bool`

HasTriggerAmount returns a boolean if a field has been set.

### GetType

`func (o *SweepConfigurationV2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SweepConfigurationV2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SweepConfigurationV2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SweepConfigurationV2) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


