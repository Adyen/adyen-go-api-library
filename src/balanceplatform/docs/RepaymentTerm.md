# RepaymentTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EstimatedDays** | **int32** | The estimated term for repaying the grant, in days. | 
**MaximumDays** | Pointer to **int32** | The maximum term for repaying the grant, in days. Only applies when &#x60;contractType&#x60; is **loan**. | [optional] 

## Methods

### NewRepaymentTerm

`func NewRepaymentTerm(estimatedDays int32, ) *RepaymentTerm`

NewRepaymentTerm instantiates a new RepaymentTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepaymentTermWithDefaults

`func NewRepaymentTermWithDefaults() *RepaymentTerm`

NewRepaymentTermWithDefaults instantiates a new RepaymentTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstimatedDays

`func (o *RepaymentTerm) GetEstimatedDays() int32`

GetEstimatedDays returns the EstimatedDays field if non-nil, zero value otherwise.

### GetEstimatedDaysOk

`func (o *RepaymentTerm) GetEstimatedDaysOk() (*int32, bool)`

GetEstimatedDaysOk returns a tuple with the EstimatedDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDays

`func (o *RepaymentTerm) SetEstimatedDays(v int32)`

SetEstimatedDays sets EstimatedDays field to given value.


### GetMaximumDays

`func (o *RepaymentTerm) GetMaximumDays() int32`

GetMaximumDays returns the MaximumDays field if non-nil, zero value otherwise.

### GetMaximumDaysOk

`func (o *RepaymentTerm) GetMaximumDaysOk() (*int32, bool)`

GetMaximumDaysOk returns a tuple with the MaximumDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDays

`func (o *RepaymentTerm) SetMaximumDays(v int32)`

SetMaximumDays sets MaximumDays field to given value.

### HasMaximumDays

`func (o *RepaymentTerm) HasMaximumDays() bool`

HasMaximumDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


