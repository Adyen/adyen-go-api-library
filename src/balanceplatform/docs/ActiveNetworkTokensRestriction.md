# ActiveNetworkTokensRestriction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | **string** | Defines how the condition must be evaluated. | 
**Value** | Pointer to **int32** | The number of tokens. | [optional] 

## Methods

### NewActiveNetworkTokensRestriction

`func NewActiveNetworkTokensRestriction(operation string, ) *ActiveNetworkTokensRestriction`

NewActiveNetworkTokensRestriction instantiates a new ActiveNetworkTokensRestriction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveNetworkTokensRestrictionWithDefaults

`func NewActiveNetworkTokensRestrictionWithDefaults() *ActiveNetworkTokensRestriction`

NewActiveNetworkTokensRestrictionWithDefaults instantiates a new ActiveNetworkTokensRestriction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *ActiveNetworkTokensRestriction) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ActiveNetworkTokensRestriction) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ActiveNetworkTokensRestriction) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetValue

`func (o *ActiveNetworkTokensRestriction) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ActiveNetworkTokensRestriction) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ActiveNetworkTokensRestriction) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ActiveNetworkTokensRestriction) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


