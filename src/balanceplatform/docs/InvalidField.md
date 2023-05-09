# InvalidField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Description of the validation error. | 
**Name** | **string** | The field that has an invalid value. | 
**Value** | **string** | The invalid value. | 

## Methods

### NewInvalidField

`func NewInvalidField(message string, name string, value string, ) *InvalidField`

NewInvalidField instantiates a new InvalidField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvalidFieldWithDefaults

`func NewInvalidFieldWithDefaults() *InvalidField`

NewInvalidFieldWithDefaults instantiates a new InvalidField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *InvalidField) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InvalidField) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InvalidField) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetName

`func (o *InvalidField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvalidField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvalidField) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *InvalidField) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InvalidField) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InvalidField) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


