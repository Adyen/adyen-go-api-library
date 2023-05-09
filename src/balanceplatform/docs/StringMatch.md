# StringMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | Pointer to **string** | The type of string matching operation. Possible values:  **startsWith**, **endsWith**, **isEqualTo**, **contains**, | [optional] 
**Value** | Pointer to **string** | The string to be matched. | [optional] 

## Methods

### NewStringMatch

`func NewStringMatch() *StringMatch`

NewStringMatch instantiates a new StringMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStringMatchWithDefaults

`func NewStringMatchWithDefaults() *StringMatch`

NewStringMatchWithDefaults instantiates a new StringMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *StringMatch) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *StringMatch) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *StringMatch) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *StringMatch) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetValue

`func (o *StringMatch) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StringMatch) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StringMatch) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *StringMatch) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


