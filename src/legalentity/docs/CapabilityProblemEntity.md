# CapabilityProblemEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Documents** | Pointer to **[]string** | List of document IDs corresponding to the verification errors from capabilities. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**CapabilityProblemEntityRecursive**](CapabilityProblemEntityRecursive.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewCapabilityProblemEntity

`func NewCapabilityProblemEntity() *CapabilityProblemEntity`

NewCapabilityProblemEntity instantiates a new CapabilityProblemEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityProblemEntityWithDefaults

`func NewCapabilityProblemEntityWithDefaults() *CapabilityProblemEntity`

NewCapabilityProblemEntityWithDefaults instantiates a new CapabilityProblemEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocuments

`func (o *CapabilityProblemEntity) GetDocuments() []string`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *CapabilityProblemEntity) GetDocumentsOk() (*[]string, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *CapabilityProblemEntity) SetDocuments(v []string)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *CapabilityProblemEntity) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetId

`func (o *CapabilityProblemEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CapabilityProblemEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CapabilityProblemEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CapabilityProblemEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOwner

`func (o *CapabilityProblemEntity) GetOwner() CapabilityProblemEntityRecursive`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CapabilityProblemEntity) GetOwnerOk() (*CapabilityProblemEntityRecursive, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CapabilityProblemEntity) SetOwner(v CapabilityProblemEntityRecursive)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CapabilityProblemEntity) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetType

`func (o *CapabilityProblemEntity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CapabilityProblemEntity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CapabilityProblemEntity) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CapabilityProblemEntity) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


