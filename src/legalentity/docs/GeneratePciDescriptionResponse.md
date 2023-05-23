# GeneratePciDescriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | The generated questionnaires in a base64 encoded format. | [optional] 
**Language** | Pointer to **string** | The two-letter [ISO-639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language code for the questionnaire. For example, **en**. | [optional] 
**PciTemplateReferences** | Pointer to **[]string** | The array of Adyen-generated unique identifiers for the questionnaires. | [optional] 

## Methods

### NewGeneratePciDescriptionResponse

`func NewGeneratePciDescriptionResponse() *GeneratePciDescriptionResponse`

NewGeneratePciDescriptionResponse instantiates a new GeneratePciDescriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneratePciDescriptionResponseWithDefaults

`func NewGeneratePciDescriptionResponseWithDefaults() *GeneratePciDescriptionResponse`

NewGeneratePciDescriptionResponseWithDefaults instantiates a new GeneratePciDescriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *GeneratePciDescriptionResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GeneratePciDescriptionResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GeneratePciDescriptionResponse) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *GeneratePciDescriptionResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetLanguage

`func (o *GeneratePciDescriptionResponse) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *GeneratePciDescriptionResponse) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *GeneratePciDescriptionResponse) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *GeneratePciDescriptionResponse) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPciTemplateReferences

`func (o *GeneratePciDescriptionResponse) GetPciTemplateReferences() []string`

GetPciTemplateReferences returns the PciTemplateReferences field if non-nil, zero value otherwise.

### GetPciTemplateReferencesOk

`func (o *GeneratePciDescriptionResponse) GetPciTemplateReferencesOk() (*[]string, bool)`

GetPciTemplateReferencesOk returns a tuple with the PciTemplateReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciTemplateReferences

`func (o *GeneratePciDescriptionResponse) SetPciTemplateReferences(v []string)`

SetPciTemplateReferences sets PciTemplateReferences field to given value.

### HasPciTemplateReferences

`func (o *GeneratePciDescriptionResponse) HasPciTemplateReferences() bool`

HasPciTemplateReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


