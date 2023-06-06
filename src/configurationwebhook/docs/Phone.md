# Phone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** | The full phone number provided as a single string.  For example, **\&quot;0031 6 11 22 33 44\&quot;**, **\&quot;+316/1122-3344\&quot;**,    or **\&quot;(0031) 611223344\&quot;**. | 
**Type** | **string** | Type of phone number. Possible values:  **Landline**, **Mobile**.  | 

## Methods

### NewPhone

`func NewPhone(number string, type_ string, ) *Phone`

NewPhone instantiates a new Phone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneWithDefaults

`func NewPhoneWithDefaults() *Phone`

NewPhoneWithDefaults instantiates a new Phone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *Phone) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Phone) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Phone) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetType

`func (o *Phone) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Phone) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Phone) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


