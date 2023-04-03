# CommonField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the field. For example, Name of External Platform. | [optional] 
**Version** | Pointer to **string** | Version of the field. For example, Version of External Platform. | [optional] 

## Methods

### NewCommonField

`func NewCommonField() *CommonField`

NewCommonField instantiates a new CommonField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonFieldWithDefaults

`func NewCommonFieldWithDefaults() *CommonField`

NewCommonFieldWithDefaults instantiates a new CommonField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CommonField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonField) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CommonField) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *CommonField) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CommonField) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CommonField) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CommonField) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


