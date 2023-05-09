# InternationalTransactionRestriction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | **string** | Defines how the condition must be evaluated. | 
**Value** | Pointer to **bool** | Boolean indicating whether transaction is an international transaction.  Possible values:  - **true**: The transaction is an international transaction.  - **false**: The transaction is a domestic transaction.   | [optional] 

## Methods

### NewInternationalTransactionRestriction

`func NewInternationalTransactionRestriction(operation string, ) *InternationalTransactionRestriction`

NewInternationalTransactionRestriction instantiates a new InternationalTransactionRestriction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternationalTransactionRestrictionWithDefaults

`func NewInternationalTransactionRestrictionWithDefaults() *InternationalTransactionRestriction`

NewInternationalTransactionRestrictionWithDefaults instantiates a new InternationalTransactionRestriction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *InternationalTransactionRestriction) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *InternationalTransactionRestriction) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *InternationalTransactionRestriction) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetValue

`func (o *InternationalTransactionRestriction) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InternationalTransactionRestriction) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InternationalTransactionRestriction) SetValue(v bool)`

SetValue sets Value field to given value.

### HasValue

`func (o *InternationalTransactionRestriction) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


