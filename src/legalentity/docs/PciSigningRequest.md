# PciSigningRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PciTemplateReferences** | **[]string** | The array of Adyen-generated unique identifiers for the questionnaires. | 
**SignedBy** | **string** | The [legal entity ID](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/legalEntities__resParam_id) of the individual who signs the PCI questionnaire. | 

## Methods

### NewPciSigningRequest

`func NewPciSigningRequest(pciTemplateReferences []string, signedBy string, ) *PciSigningRequest`

NewPciSigningRequest instantiates a new PciSigningRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPciSigningRequestWithDefaults

`func NewPciSigningRequestWithDefaults() *PciSigningRequest`

NewPciSigningRequestWithDefaults instantiates a new PciSigningRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPciTemplateReferences

`func (o *PciSigningRequest) GetPciTemplateReferences() []string`

GetPciTemplateReferences returns the PciTemplateReferences field if non-nil, zero value otherwise.

### GetPciTemplateReferencesOk

`func (o *PciSigningRequest) GetPciTemplateReferencesOk() (*[]string, bool)`

GetPciTemplateReferencesOk returns a tuple with the PciTemplateReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciTemplateReferences

`func (o *PciSigningRequest) SetPciTemplateReferences(v []string)`

SetPciTemplateReferences sets PciTemplateReferences field to given value.


### GetSignedBy

`func (o *PciSigningRequest) GetSignedBy() string`

GetSignedBy returns the SignedBy field if non-nil, zero value otherwise.

### GetSignedByOk

`func (o *PciSigningRequest) GetSignedByOk() (*string, bool)`

GetSignedByOk returns a tuple with the SignedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedBy

`func (o *PciSigningRequest) SetSignedBy(v string)`

SetSignedBy sets SignedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


