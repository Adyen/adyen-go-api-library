# Commission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FixedAmount** | Pointer to **int64** | A fixed commission fee, in minor units. | [optional] 
**VariablePercentage** | Pointer to **int64** | A variable commission fee, in basis points. | [optional] 

## Methods

### NewCommission

`func NewCommission() *Commission`

NewCommission instantiates a new Commission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommissionWithDefaults

`func NewCommissionWithDefaults() *Commission`

NewCommissionWithDefaults instantiates a new Commission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFixedAmount

`func (o *Commission) GetFixedAmount() int64`

GetFixedAmount returns the FixedAmount field if non-nil, zero value otherwise.

### GetFixedAmountOk

`func (o *Commission) GetFixedAmountOk() (*int64, bool)`

GetFixedAmountOk returns a tuple with the FixedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAmount

`func (o *Commission) SetFixedAmount(v int64)`

SetFixedAmount sets FixedAmount field to given value.

### HasFixedAmount

`func (o *Commission) HasFixedAmount() bool`

HasFixedAmount returns a boolean if a field has been set.

### GetVariablePercentage

`func (o *Commission) GetVariablePercentage() int64`

GetVariablePercentage returns the VariablePercentage field if non-nil, zero value otherwise.

### GetVariablePercentageOk

`func (o *Commission) GetVariablePercentageOk() (*int64, bool)`

GetVariablePercentageOk returns a tuple with the VariablePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariablePercentage

`func (o *Commission) SetVariablePercentage(v int64)`

SetVariablePercentage sets VariablePercentage field to given value.

### HasVariablePercentage

`func (o *Commission) HasVariablePercentage() bool`

HasVariablePercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


