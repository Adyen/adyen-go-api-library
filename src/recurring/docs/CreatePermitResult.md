# CreatePermitResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermitResultList** | Pointer to [**[]PermitResult**](PermitResult.md) | List of new permits. | [optional] 
**PspReference** | Pointer to **string** | A unique reference associated with the request. This value is globally unique; quote it when communicating with us about this request. | [optional] 

## Methods

### NewCreatePermitResult

`func NewCreatePermitResult() *CreatePermitResult`

NewCreatePermitResult instantiates a new CreatePermitResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePermitResultWithDefaults

`func NewCreatePermitResultWithDefaults() *CreatePermitResult`

NewCreatePermitResultWithDefaults instantiates a new CreatePermitResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermitResultList

`func (o *CreatePermitResult) GetPermitResultList() []PermitResult`

GetPermitResultList returns the PermitResultList field if non-nil, zero value otherwise.

### GetPermitResultListOk

`func (o *CreatePermitResult) GetPermitResultListOk() (*[]PermitResult, bool)`

GetPermitResultListOk returns a tuple with the PermitResultList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermitResultList

`func (o *CreatePermitResult) SetPermitResultList(v []PermitResult)`

SetPermitResultList sets PermitResultList field to given value.

### HasPermitResultList

`func (o *CreatePermitResult) HasPermitResultList() bool`

HasPermitResultList returns a boolean if a field has been set.

### GetPspReference

`func (o *CreatePermitResult) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *CreatePermitResult) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *CreatePermitResult) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.

### HasPspReference

`func (o *CreatePermitResult) HasPspReference() bool`

HasPspReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


