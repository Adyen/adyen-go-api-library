# DocumentReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Identifies whether the document is active and used for checks. | [optional] 
**Description** | Pointer to **string** | Your description for the document. | [optional] 
**FileName** | Pointer to **string** | Document name. | [optional] 
**Id** | Pointer to **string** | The unique identifier of the resource. | [optional] 
**ModificationDate** | Pointer to **time.Time** | The modification date of the document. | [optional] 
**Type** | Pointer to **string** | Type of document, used when providing an ID number or uploading a document. | [optional] 

## Methods

### NewDocumentReference

`func NewDocumentReference() *DocumentReference`

NewDocumentReference instantiates a new DocumentReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentReferenceWithDefaults

`func NewDocumentReferenceWithDefaults() *DocumentReference`

NewDocumentReferenceWithDefaults instantiates a new DocumentReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *DocumentReference) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *DocumentReference) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *DocumentReference) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *DocumentReference) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDescription

`func (o *DocumentReference) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DocumentReference) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DocumentReference) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DocumentReference) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFileName

`func (o *DocumentReference) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *DocumentReference) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *DocumentReference) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *DocumentReference) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetId

`func (o *DocumentReference) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentReference) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentReference) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DocumentReference) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModificationDate

`func (o *DocumentReference) GetModificationDate() time.Time`

GetModificationDate returns the ModificationDate field if non-nil, zero value otherwise.

### GetModificationDateOk

`func (o *DocumentReference) GetModificationDateOk() (*time.Time, bool)`

GetModificationDateOk returns a tuple with the ModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationDate

`func (o *DocumentReference) SetModificationDate(v time.Time)`

SetModificationDate sets ModificationDate field to given value.

### HasModificationDate

`func (o *DocumentReference) HasModificationDate() bool`

HasModificationDate returns a boolean if a field has been set.

### GetType

`func (o *DocumentReference) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DocumentReference) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DocumentReference) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DocumentReference) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


