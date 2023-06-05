# SweepSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The schedule type.  Possible values:  * **cron**: push out funds based on a cron expression.  * **daily**: push out funds daily at 07:00 AM CET.  * **weekly**: push out funds every Monday at 07:00 AM CET.  * **monthly**: push out funds every first of the month at 07:00 AM CET.  * **balance**: pull in funds instantly if the balance is less than or equal to the &#x60;triggerAmount&#x60;. You can only use this for sweeps of &#x60;type&#x60; **pull** and when the source is a &#x60;merchantAccount&#x60; or &#x60;transferInstrument&#x60;.If the source is transferInstrument, merchant account identifier is still required, with which you want to process the transaction. | [optional] 

## Methods

### NewSweepSchedule

`func NewSweepSchedule() *SweepSchedule`

NewSweepSchedule instantiates a new SweepSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSweepScheduleWithDefaults

`func NewSweepScheduleWithDefaults() *SweepSchedule`

NewSweepScheduleWithDefaults instantiates a new SweepSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SweepSchedule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SweepSchedule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SweepSchedule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SweepSchedule) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


