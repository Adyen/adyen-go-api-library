# DifferentCurrenciesRestriction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | **string** | Defines how the condition must be evaluated. | 
**Value** | Pointer to **bool** | Checks the currency of the payment against the currency of the payment instrument.  Possible values:  - **true**: The currency of the payment is different from the currency of the payment instrument.  - **false**: The currencies are the same.   | [optional] 

## Methods

### NewDifferentCurrenciesRestriction

`func NewDifferentCurrenciesRestriction(operation string, ) *DifferentCurrenciesRestriction`

NewDifferentCurrenciesRestriction instantiates a new DifferentCurrenciesRestriction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDifferentCurrenciesRestrictionWithDefaults

`func NewDifferentCurrenciesRestrictionWithDefaults() *DifferentCurrenciesRestriction`

NewDifferentCurrenciesRestrictionWithDefaults instantiates a new DifferentCurrenciesRestriction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *DifferentCurrenciesRestriction) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *DifferentCurrenciesRestriction) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *DifferentCurrenciesRestriction) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetValue

`func (o *DifferentCurrenciesRestriction) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DifferentCurrenciesRestriction) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DifferentCurrenciesRestriction) SetValue(v bool)`

SetValue sets Value field to given value.

### HasValue

`func (o *DifferentCurrenciesRestriction) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


