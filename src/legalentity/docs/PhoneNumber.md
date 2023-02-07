# PhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** | The full phone number, including the country code. For example, **+3112345678**. | 
**Type** | Pointer to **string** | The type of phone number.  Possible values: **mobile**, **landline**, **sip**, **fax.**  | [optional] 

## Methods

### NewPhoneNumber

`func NewPhoneNumber(number string, ) *PhoneNumber`

NewPhoneNumber instantiates a new PhoneNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneNumberWithDefaults

`func NewPhoneNumberWithDefaults() *PhoneNumber`

NewPhoneNumberWithDefaults instantiates a new PhoneNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *PhoneNumber) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PhoneNumber) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PhoneNumber) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetType

`func (o *PhoneNumber) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PhoneNumber) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PhoneNumber) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PhoneNumber) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


