# Duration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | Pointer to **string** | The unit of time. You can only use **minutes** and **hours** if the &#x60;interval.type&#x60; is **sliding**.  Possible values: **minutes**, **hours**, **days**, **weeks**, or **months** | [optional] 
**Value** | Pointer to **int32** | The length of time by the unit. For example, 5 days.  The maximum duration is 90 days or an equivalent in other units. For example, 3 months. | [optional] 

## Methods

### NewDuration

`func NewDuration() *Duration`

NewDuration instantiates a new Duration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDurationWithDefaults

`func NewDurationWithDefaults() *Duration`

NewDurationWithDefaults instantiates a new Duration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *Duration) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Duration) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Duration) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Duration) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetValue

`func (o *Duration) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Duration) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Duration) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *Duration) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


