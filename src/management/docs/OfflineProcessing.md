# OfflineProcessing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChipFloorLimit** | Pointer to **int32** | The (inclusive) limit for accepting chip cards offline, in the processing currency, in minor units | [optional] 
**OfflineSwipeLimits** | Pointer to [**[]MinorUnitsMonetaryValue**](MinorUnitsMonetaryValue.md) | The maximum amount up to which swiped credit cards can be accepted offline, in the specified currency | [optional] 

## Methods

### NewOfflineProcessing

`func NewOfflineProcessing() *OfflineProcessing`

NewOfflineProcessing instantiates a new OfflineProcessing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfflineProcessingWithDefaults

`func NewOfflineProcessingWithDefaults() *OfflineProcessing`

NewOfflineProcessingWithDefaults instantiates a new OfflineProcessing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChipFloorLimit

`func (o *OfflineProcessing) GetChipFloorLimit() int32`

GetChipFloorLimit returns the ChipFloorLimit field if non-nil, zero value otherwise.

### GetChipFloorLimitOk

`func (o *OfflineProcessing) GetChipFloorLimitOk() (*int32, bool)`

GetChipFloorLimitOk returns a tuple with the ChipFloorLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChipFloorLimit

`func (o *OfflineProcessing) SetChipFloorLimit(v int32)`

SetChipFloorLimit sets ChipFloorLimit field to given value.

### HasChipFloorLimit

`func (o *OfflineProcessing) HasChipFloorLimit() bool`

HasChipFloorLimit returns a boolean if a field has been set.

### GetOfflineSwipeLimits

`func (o *OfflineProcessing) GetOfflineSwipeLimits() []MinorUnitsMonetaryValue`

GetOfflineSwipeLimits returns the OfflineSwipeLimits field if non-nil, zero value otherwise.

### GetOfflineSwipeLimitsOk

`func (o *OfflineProcessing) GetOfflineSwipeLimitsOk() (*[]MinorUnitsMonetaryValue, bool)`

GetOfflineSwipeLimitsOk returns a tuple with the OfflineSwipeLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineSwipeLimits

`func (o *OfflineProcessing) SetOfflineSwipeLimits(v []MinorUnitsMonetaryValue)`

SetOfflineSwipeLimits sets OfflineSwipeLimits field to given value.

### HasOfflineSwipeLimits

`func (o *OfflineProcessing) HasOfflineSwipeLimits() bool`

HasOfflineSwipeLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


