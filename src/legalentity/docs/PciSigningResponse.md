# PciSigningResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PciQuestionnaireIds** | Pointer to **[]string** | The unique identifiers of the signed PCI documents. | [optional] 
**SignedBy** | Pointer to **string** | The [legal entity ID](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/legalEntities__resParam_id) of the individual who signed the PCI questionnaire. | [optional] 

## Methods

### NewPciSigningResponse

`func NewPciSigningResponse() *PciSigningResponse`

NewPciSigningResponse instantiates a new PciSigningResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPciSigningResponseWithDefaults

`func NewPciSigningResponseWithDefaults() *PciSigningResponse`

NewPciSigningResponseWithDefaults instantiates a new PciSigningResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPciQuestionnaireIds

`func (o *PciSigningResponse) GetPciQuestionnaireIds() []string`

GetPciQuestionnaireIds returns the PciQuestionnaireIds field if non-nil, zero value otherwise.

### GetPciQuestionnaireIdsOk

`func (o *PciSigningResponse) GetPciQuestionnaireIdsOk() (*[]string, bool)`

GetPciQuestionnaireIdsOk returns a tuple with the PciQuestionnaireIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciQuestionnaireIds

`func (o *PciSigningResponse) SetPciQuestionnaireIds(v []string)`

SetPciQuestionnaireIds sets PciQuestionnaireIds field to given value.

### HasPciQuestionnaireIds

`func (o *PciSigningResponse) HasPciQuestionnaireIds() bool`

HasPciQuestionnaireIds returns a boolean if a field has been set.

### GetSignedBy

`func (o *PciSigningResponse) GetSignedBy() string`

GetSignedBy returns the SignedBy field if non-nil, zero value otherwise.

### GetSignedByOk

`func (o *PciSigningResponse) GetSignedByOk() (*string, bool)`

GetSignedByOk returns a tuple with the SignedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedBy

`func (o *PciSigningResponse) SetSignedBy(v string)`

SetSignedBy sets SignedBy field to given value.

### HasSignedBy

`func (o *PciSigningResponse) HasSignedBy() bool`

HasSignedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


