# TimeOfDay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **string** | The end time in a time-only ISO-8601 extended offset format. For example: **08:00:00+02:00**, **22:30:00-03:00**.   | [optional] 
**StartTime** | Pointer to **string** | The start time in a time-only ISO-8601 extended offset format. For example: **08:00:00+02:00**, **22:30:00-03:00**.   | [optional] 

## Methods

### NewTimeOfDay

`func NewTimeOfDay() *TimeOfDay`

NewTimeOfDay instantiates a new TimeOfDay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeOfDayWithDefaults

`func NewTimeOfDayWithDefaults() *TimeOfDay`

NewTimeOfDayWithDefaults instantiates a new TimeOfDay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *TimeOfDay) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TimeOfDay) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TimeOfDay) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TimeOfDay) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetStartTime

`func (o *TimeOfDay) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TimeOfDay) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TimeOfDay) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TimeOfDay) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


