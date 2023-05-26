# ResourceReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the resource. | [optional] 
**Id** | Pointer to **string** | The unique identifier of the resource. | [optional] 
**Reference** | Pointer to **string** | The reference for the resource. | [optional] 

## Methods

### NewResourceReference

`func NewResourceReference() *ResourceReference`

NewResourceReference instantiates a new ResourceReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceReferenceWithDefaults

`func NewResourceReferenceWithDefaults() *ResourceReference`

NewResourceReferenceWithDefaults instantiates a new ResourceReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ResourceReference) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceReference) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceReference) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResourceReference) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ResourceReference) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceReference) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceReference) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceReference) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReference

`func (o *ResourceReference) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ResourceReference) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ResourceReference) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ResourceReference) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


