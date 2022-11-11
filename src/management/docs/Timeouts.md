# Timeouts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromActiveToSleep** | Pointer to **int32** | Indicates the number of seconds of inactivity after which the terminal display goes into sleep mode. | [optional] 

## Methods

### NewTimeouts

`func NewTimeouts() *Timeouts`

NewTimeouts instantiates a new Timeouts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeoutsWithDefaults

`func NewTimeoutsWithDefaults() *Timeouts`

NewTimeoutsWithDefaults instantiates a new Timeouts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromActiveToSleep

`func (o *Timeouts) GetFromActiveToSleep() int32`

GetFromActiveToSleep returns the FromActiveToSleep field if non-nil, zero value otherwise.

### GetFromActiveToSleepOk

`func (o *Timeouts) GetFromActiveToSleepOk() (*int32, bool)`

GetFromActiveToSleepOk returns a tuple with the FromActiveToSleep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromActiveToSleep

`func (o *Timeouts) SetFromActiveToSleep(v int32)`

SetFromActiveToSleep sets FromActiveToSleep field to given value.

### HasFromActiveToSleep

`func (o *Timeouts) HasFromActiveToSleep() bool`

HasFromActiveToSleep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


