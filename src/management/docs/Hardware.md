# Hardware

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayMaximumBackLight** | Pointer to **int32** | The brightness of the display when the terminal is being used, expressed as a percentage. | [optional] 

## Methods

### NewHardware

`func NewHardware() *Hardware`

NewHardware instantiates a new Hardware object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHardwareWithDefaults

`func NewHardwareWithDefaults() *Hardware`

NewHardwareWithDefaults instantiates a new Hardware object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayMaximumBackLight

`func (o *Hardware) GetDisplayMaximumBackLight() int32`

GetDisplayMaximumBackLight returns the DisplayMaximumBackLight field if non-nil, zero value otherwise.

### GetDisplayMaximumBackLightOk

`func (o *Hardware) GetDisplayMaximumBackLightOk() (*int32, bool)`

GetDisplayMaximumBackLightOk returns a tuple with the DisplayMaximumBackLight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMaximumBackLight

`func (o *Hardware) SetDisplayMaximumBackLight(v int32)`

SetDisplayMaximumBackLight sets DisplayMaximumBackLight field to given value.

### HasDisplayMaximumBackLight

`func (o *Hardware) HasDisplayMaximumBackLight() bool`

HasDisplayMaximumBackLight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


