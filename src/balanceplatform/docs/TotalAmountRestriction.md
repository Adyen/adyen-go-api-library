# TotalAmountRestriction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | **string** | Defines how the condition must be evaluated. | 
**Value** | Pointer to [**Amount**](Amount.md) |  | [optional] 

## Methods

### NewTotalAmountRestriction

`func NewTotalAmountRestriction(operation string, ) *TotalAmountRestriction`

NewTotalAmountRestriction instantiates a new TotalAmountRestriction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTotalAmountRestrictionWithDefaults

`func NewTotalAmountRestrictionWithDefaults() *TotalAmountRestriction`

NewTotalAmountRestrictionWithDefaults instantiates a new TotalAmountRestriction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *TotalAmountRestriction) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *TotalAmountRestriction) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *TotalAmountRestriction) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetValue

`func (o *TotalAmountRestriction) GetValue() Amount`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TotalAmountRestriction) GetValueOk() (*Amount, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TotalAmountRestriction) SetValue(v Amount)`

SetValue sets Value field to given value.

### HasValue

`func (o *TotalAmountRestriction) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

