# LegalEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to [**map[string]LegalEntityCapability**](LegalEntityCapability.md) | Contains key-value pairs that specify the actions that the legal entity can do in your platform.The key is a capability required for your integration. For example, **issueCard** for Issuing.The value is an object containing the settings for the capability. | [optional] [readonly] 
**DocumentDetails** | Pointer to [**[]DocumentReference**](DocumentReference.md) | List of documents uploaded for the legal entity. | [optional] 
**Documents** | Pointer to [**[]EntityReference**](EntityReference.md) | List of documents uploaded for the legal entity. | [optional] 
**EntityAssociations** | Pointer to [**[]LegalEntityAssociation**](LegalEntityAssociation.md) | List of legal entities associated with the current legal entity. For example, ultimate beneficial owners associated with an organization through ownership or control, or as signatories. | [optional] 
**Id** | **string** | The unique identifier of the legal entity. | [readonly] 
**Individual** | Pointer to [**Individual**](Individual.md) |  | [optional] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 
**Reference** | Pointer to **string** | Your reference for the legal entity, maximum 150 characters. | [optional] 
**SoleProprietorship** | Pointer to [**SoleProprietorship**](SoleProprietorship.md) |  | [optional] 
**TransferInstruments** | Pointer to [**[]TransferInstrumentReference**](TransferInstrumentReference.md) | List of transfer instruments that the legal entity owns. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of legal entity.   Possible values: **individual**, **organization**, or **soleProprietorship**. | [optional] 

## Methods

### NewLegalEntity

`func NewLegalEntity(id string, ) *LegalEntity`

NewLegalEntity instantiates a new LegalEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalEntityWithDefaults

`func NewLegalEntityWithDefaults() *LegalEntity`

NewLegalEntityWithDefaults instantiates a new LegalEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *LegalEntity) GetCapabilities() map[string]LegalEntityCapability`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *LegalEntity) GetCapabilitiesOk() (*map[string]LegalEntityCapability, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *LegalEntity) SetCapabilities(v map[string]LegalEntityCapability)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *LegalEntity) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetDocumentDetails

`func (o *LegalEntity) GetDocumentDetails() []DocumentReference`

GetDocumentDetails returns the DocumentDetails field if non-nil, zero value otherwise.

### GetDocumentDetailsOk

`func (o *LegalEntity) GetDocumentDetailsOk() (*[]DocumentReference, bool)`

GetDocumentDetailsOk returns a tuple with the DocumentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentDetails

`func (o *LegalEntity) SetDocumentDetails(v []DocumentReference)`

SetDocumentDetails sets DocumentDetails field to given value.

### HasDocumentDetails

`func (o *LegalEntity) HasDocumentDetails() bool`

HasDocumentDetails returns a boolean if a field has been set.

### GetDocuments

`func (o *LegalEntity) GetDocuments() []EntityReference`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *LegalEntity) GetDocumentsOk() (*[]EntityReference, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *LegalEntity) SetDocuments(v []EntityReference)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *LegalEntity) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetEntityAssociations

`func (o *LegalEntity) GetEntityAssociations() []LegalEntityAssociation`

GetEntityAssociations returns the EntityAssociations field if non-nil, zero value otherwise.

### GetEntityAssociationsOk

`func (o *LegalEntity) GetEntityAssociationsOk() (*[]LegalEntityAssociation, bool)`

GetEntityAssociationsOk returns a tuple with the EntityAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityAssociations

`func (o *LegalEntity) SetEntityAssociations(v []LegalEntityAssociation)`

SetEntityAssociations sets EntityAssociations field to given value.

### HasEntityAssociations

`func (o *LegalEntity) HasEntityAssociations() bool`

HasEntityAssociations returns a boolean if a field has been set.

### GetId

`func (o *LegalEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LegalEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LegalEntity) SetId(v string)`

SetId sets Id field to given value.


### GetIndividual

`func (o *LegalEntity) GetIndividual() Individual`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *LegalEntity) GetIndividualOk() (*Individual, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *LegalEntity) SetIndividual(v Individual)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *LegalEntity) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.

### GetOrganization

`func (o *LegalEntity) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *LegalEntity) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *LegalEntity) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *LegalEntity) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetReference

`func (o *LegalEntity) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *LegalEntity) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *LegalEntity) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *LegalEntity) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSoleProprietorship

`func (o *LegalEntity) GetSoleProprietorship() SoleProprietorship`

GetSoleProprietorship returns the SoleProprietorship field if non-nil, zero value otherwise.

### GetSoleProprietorshipOk

`func (o *LegalEntity) GetSoleProprietorshipOk() (*SoleProprietorship, bool)`

GetSoleProprietorshipOk returns a tuple with the SoleProprietorship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoleProprietorship

`func (o *LegalEntity) SetSoleProprietorship(v SoleProprietorship)`

SetSoleProprietorship sets SoleProprietorship field to given value.

### HasSoleProprietorship

`func (o *LegalEntity) HasSoleProprietorship() bool`

HasSoleProprietorship returns a boolean if a field has been set.

### GetTransferInstruments

`func (o *LegalEntity) GetTransferInstruments() []TransferInstrumentReference`

GetTransferInstruments returns the TransferInstruments field if non-nil, zero value otherwise.

### GetTransferInstrumentsOk

`func (o *LegalEntity) GetTransferInstrumentsOk() (*[]TransferInstrumentReference, bool)`

GetTransferInstrumentsOk returns a tuple with the TransferInstruments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferInstruments

`func (o *LegalEntity) SetTransferInstruments(v []TransferInstrumentReference)`

SetTransferInstruments sets TransferInstruments field to given value.

### HasTransferInstruments

`func (o *LegalEntity) HasTransferInstruments() bool`

HasTransferInstruments returns a boolean if a field has been set.

### GetType

`func (o *LegalEntity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LegalEntity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LegalEntity) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LegalEntity) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


