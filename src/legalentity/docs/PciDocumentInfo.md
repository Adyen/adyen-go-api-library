# PciDocumentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | The date the questionnaire was created, in ISO 8601 extended format. For example, 2022-12-18T10:15:30+01:00 | [optional] 
**Id** | Pointer to **string** | The unique identifier of the signed questionnaire. | [optional] 
**ValidUntil** | Pointer to **time.Time** | The expiration date of the questionnaire, in ISO 8601 extended format. For example, 2022-12-18T10:15:30+01:00 | [optional] 

## Methods

### NewPciDocumentInfo

`func NewPciDocumentInfo() *PciDocumentInfo`

NewPciDocumentInfo instantiates a new PciDocumentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPciDocumentInfoWithDefaults

`func NewPciDocumentInfoWithDefaults() *PciDocumentInfo`

NewPciDocumentInfoWithDefaults instantiates a new PciDocumentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PciDocumentInfo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PciDocumentInfo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PciDocumentInfo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PciDocumentInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *PciDocumentInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PciDocumentInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PciDocumentInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PciDocumentInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValidUntil

`func (o *PciDocumentInfo) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *PciDocumentInfo) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *PciDocumentInfo) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *PciDocumentInfo) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


