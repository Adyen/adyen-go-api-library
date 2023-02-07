# LegalEntityInfoRequiredType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to [**map[string]LegalEntityCapability**](LegalEntityCapability.md) | Contains key-value pairs that specify the actions that the legal entity can do in your platform.The key is a capability required for your integration. For example, **issueCard** for Issuing.The value is an object containing the settings for the capability. | [optional] [readonly] 
**EntityAssociations** | Pointer to [**[]LegalEntityAssociation**](LegalEntityAssociation.md) | List of legal entities associated with the current legal entity. For example, ultimate beneficial owners associated with an organization through ownership or control, or as signatories. | [optional] 
**Individual** | Pointer to [**Individual**](Individual.md) |  | [optional] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 
**Reference** | Pointer to **string** | Your reference for the legal entity, maximum 150 characters. | [optional] 
**SoleProprietorship** | Pointer to [**SoleProprietorship**](SoleProprietorship.md) |  | [optional] 
**Type** | **string** | The type of legal entity.   Possible values: **individual**, **organization**, or **soleProprietorship**. | 

## Methods

### NewLegalEntityInfoRequiredType

`func NewLegalEntityInfoRequiredType(type_ string, ) *LegalEntityInfoRequiredType`

NewLegalEntityInfoRequiredType instantiates a new LegalEntityInfoRequiredType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalEntityInfoRequiredTypeWithDefaults

`func NewLegalEntityInfoRequiredTypeWithDefaults() *LegalEntityInfoRequiredType`

NewLegalEntityInfoRequiredTypeWithDefaults instantiates a new LegalEntityInfoRequiredType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *LegalEntityInfoRequiredType) GetCapabilities() map[string]LegalEntityCapability`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *LegalEntityInfoRequiredType) GetCapabilitiesOk() (*map[string]LegalEntityCapability, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *LegalEntityInfoRequiredType) SetCapabilities(v map[string]LegalEntityCapability)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *LegalEntityInfoRequiredType) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetEntityAssociations

`func (o *LegalEntityInfoRequiredType) GetEntityAssociations() []LegalEntityAssociation`

GetEntityAssociations returns the EntityAssociations field if non-nil, zero value otherwise.

### GetEntityAssociationsOk

`func (o *LegalEntityInfoRequiredType) GetEntityAssociationsOk() (*[]LegalEntityAssociation, bool)`

GetEntityAssociationsOk returns a tuple with the EntityAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityAssociations

`func (o *LegalEntityInfoRequiredType) SetEntityAssociations(v []LegalEntityAssociation)`

SetEntityAssociations sets EntityAssociations field to given value.

### HasEntityAssociations

`func (o *LegalEntityInfoRequiredType) HasEntityAssociations() bool`

HasEntityAssociations returns a boolean if a field has been set.

### GetIndividual

`func (o *LegalEntityInfoRequiredType) GetIndividual() Individual`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *LegalEntityInfoRequiredType) GetIndividualOk() (*Individual, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *LegalEntityInfoRequiredType) SetIndividual(v Individual)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *LegalEntityInfoRequiredType) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.

### GetOrganization

`func (o *LegalEntityInfoRequiredType) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *LegalEntityInfoRequiredType) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *LegalEntityInfoRequiredType) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *LegalEntityInfoRequiredType) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetReference

`func (o *LegalEntityInfoRequiredType) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *LegalEntityInfoRequiredType) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *LegalEntityInfoRequiredType) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *LegalEntityInfoRequiredType) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSoleProprietorship

`func (o *LegalEntityInfoRequiredType) GetSoleProprietorship() SoleProprietorship`

GetSoleProprietorship returns the SoleProprietorship field if non-nil, zero value otherwise.

### GetSoleProprietorshipOk

`func (o *LegalEntityInfoRequiredType) GetSoleProprietorshipOk() (*SoleProprietorship, bool)`

GetSoleProprietorshipOk returns a tuple with the SoleProprietorship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoleProprietorship

`func (o *LegalEntityInfoRequiredType) SetSoleProprietorship(v SoleProprietorship)`

SetSoleProprietorship sets SoleProprietorship field to given value.

### HasSoleProprietorship

`func (o *LegalEntityInfoRequiredType) HasSoleProprietorship() bool`

HasSoleProprietorship returns a boolean if a field has been set.

### GetType

`func (o *LegalEntityInfoRequiredType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LegalEntityInfoRequiredType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LegalEntityInfoRequiredType) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


