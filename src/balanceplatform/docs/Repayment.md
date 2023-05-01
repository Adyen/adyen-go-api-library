# Repayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasisPoints** | **int32** | The repayment that is deducted daily from incoming net volume, in [basis points](https://www.investopedia.com/terms/b/basispoint.asp). | 
**Term** | Pointer to [**RepaymentTerm**](RepaymentTerm.md) |  | [optional] 
**Threshold** | Pointer to [**ThresholdRepayment**](ThresholdRepayment.md) |  | [optional] 

## Methods

### NewRepayment

`func NewRepayment(basisPoints int32, ) *Repayment`

NewRepayment instantiates a new Repayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepaymentWithDefaults

`func NewRepaymentWithDefaults() *Repayment`

NewRepaymentWithDefaults instantiates a new Repayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasisPoints

`func (o *Repayment) GetBasisPoints() int32`

GetBasisPoints returns the BasisPoints field if non-nil, zero value otherwise.

### GetBasisPointsOk

`func (o *Repayment) GetBasisPointsOk() (*int32, bool)`

GetBasisPointsOk returns a tuple with the BasisPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasisPoints

`func (o *Repayment) SetBasisPoints(v int32)`

SetBasisPoints sets BasisPoints field to given value.


### GetTerm

`func (o *Repayment) GetTerm() RepaymentTerm`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *Repayment) GetTermOk() (*RepaymentTerm, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *Repayment) SetTerm(v RepaymentTerm)`

SetTerm sets Term field to given value.

### HasTerm

`func (o *Repayment) HasTerm() bool`

HasTerm returns a boolean if a field has been set.

### GetThreshold

`func (o *Repayment) GetThreshold() ThresholdRepayment`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *Repayment) GetThresholdOk() (*ThresholdRepayment, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *Repayment) SetThreshold(v ThresholdRepayment)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *Repayment) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


