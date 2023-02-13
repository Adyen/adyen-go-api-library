# GetTermsOfServiceDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | Pointer to **string** | The language to be used for the Terms of Service document, specified by the two letter [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language code. For example, **nl** for Dutch. | [optional] 
**Type** | Pointer to **string** | The type of Terms of Service. | [optional] 

## Methods

### NewGetTermsOfServiceDocumentRequest

`func NewGetTermsOfServiceDocumentRequest() *GetTermsOfServiceDocumentRequest`

NewGetTermsOfServiceDocumentRequest instantiates a new GetTermsOfServiceDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTermsOfServiceDocumentRequestWithDefaults

`func NewGetTermsOfServiceDocumentRequestWithDefaults() *GetTermsOfServiceDocumentRequest`

NewGetTermsOfServiceDocumentRequestWithDefaults instantiates a new GetTermsOfServiceDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *GetTermsOfServiceDocumentRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *GetTermsOfServiceDocumentRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *GetTermsOfServiceDocumentRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *GetTermsOfServiceDocumentRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetType

`func (o *GetTermsOfServiceDocumentRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetTermsOfServiceDocumentRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetTermsOfServiceDocumentRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetTermsOfServiceDocumentRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


