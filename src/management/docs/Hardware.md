# Hardware

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayMaximumBackLight** | Pointer to **int32** | The brightness of the display when the terminal is being used, expressed as a percentage. | [optional] 
**RestartHour** | Pointer to **int32** | The hour (0 - 23) in which the device will reboot, reboot will happen in the timezone of the device | [optional] 

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

### GetRestartHour

`func (o *Hardware) GetRestartHour() int32`

GetRestartHour returns the RestartHour field if non-nil, zero value otherwise.

### GetRestartHourOk

`func (o *Hardware) GetRestartHourOk() (*int32, bool)`

GetRestartHourOk returns a tuple with the RestartHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartHour

`func (o *Hardware) SetRestartHour(v int32)`

SetRestartHour sets RestartHour field to given value.

### HasRestartHour

`func (o *Hardware) HasRestartHour() bool`

HasRestartHour returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


