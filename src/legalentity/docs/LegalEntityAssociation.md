# LegalEntityAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatorId** | Pointer to **string** | The unique identifier of another legal entity with which the &#x60;legalEntityId&#x60; is associated. When the &#x60;legalEntityId&#x60; is associated to legal entities other than the current one, the response returns all the associations. | [optional] [readonly] 
**EntityType** | Pointer to **string** | The legal entity type of associated legal entity.   For example, **organization**, **soleProprietorship** or **individual**.  | [optional] [readonly] 
**JobTitle** | Pointer to **string** | The individual&#39;s job title if the &#x60;type&#x60; is **uboThroughControl** or **signatory**. | [optional] 
**LegalEntityId** | **string** | The unique identifier of the associated [legal entity](https://docs.adyen.com/api-explorer/legalentity/latest/post/legalEntities#responses-200-id). | 
**Name** | Pointer to **string** | The name of the associated [legal entity](https://docs.adyen.com/api-explorer/legalentity/latest/post/legalEntities#responses-200-id).  - For **individual**, &#x60;name.firstName&#x60; and &#x60;name.lastName&#x60;. - For **organization**, &#x60;legalName&#x60;. - For **soleProprietorship**, &#x60;name&#x60;. | [optional] [readonly] 
**Type** | **string** | Defines the relationship of the legal entity to the current legal entity.   Possible values for organizations: **uboThroughOwnership**, **uboThroughControl**, **signatory**, or **ultimateParentCompany**.   Possible values for sole proprietorships: **soleProprietorship**.  | 

## Methods

### NewLegalEntityAssociation

`func NewLegalEntityAssociation(legalEntityId string, type_ string, ) *LegalEntityAssociation`

NewLegalEntityAssociation instantiates a new LegalEntityAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalEntityAssociationWithDefaults

`func NewLegalEntityAssociationWithDefaults() *LegalEntityAssociation`

NewLegalEntityAssociationWithDefaults instantiates a new LegalEntityAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociatorId

`func (o *LegalEntityAssociation) GetAssociatorId() string`

GetAssociatorId returns the AssociatorId field if non-nil, zero value otherwise.

### GetAssociatorIdOk

`func (o *LegalEntityAssociation) GetAssociatorIdOk() (*string, bool)`

GetAssociatorIdOk returns a tuple with the AssociatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatorId

`func (o *LegalEntityAssociation) SetAssociatorId(v string)`

SetAssociatorId sets AssociatorId field to given value.

### HasAssociatorId

`func (o *LegalEntityAssociation) HasAssociatorId() bool`

HasAssociatorId returns a boolean if a field has been set.

### GetEntityType

`func (o *LegalEntityAssociation) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *LegalEntityAssociation) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *LegalEntityAssociation) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *LegalEntityAssociation) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetJobTitle

`func (o *LegalEntityAssociation) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *LegalEntityAssociation) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *LegalEntityAssociation) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *LegalEntityAssociation) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetLegalEntityId

`func (o *LegalEntityAssociation) GetLegalEntityId() string`

GetLegalEntityId returns the LegalEntityId field if non-nil, zero value otherwise.

### GetLegalEntityIdOk

`func (o *LegalEntityAssociation) GetLegalEntityIdOk() (*string, bool)`

GetLegalEntityIdOk returns a tuple with the LegalEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalEntityId

`func (o *LegalEntityAssociation) SetLegalEntityId(v string)`

SetLegalEntityId sets LegalEntityId field to given value.


### GetName

`func (o *LegalEntityAssociation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LegalEntityAssociation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LegalEntityAssociation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LegalEntityAssociation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *LegalEntityAssociation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LegalEntityAssociation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LegalEntityAssociation) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


