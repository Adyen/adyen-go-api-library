# GetTermsOfServiceDocumentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Document** | Pointer to **string** | The Terms of Service document in Base64-encoded format. | [optional] 
**Id** | Pointer to **string** | The unique identifier of the legal entity. | [optional] 
**Language** | Pointer to **string** | The language used for the Terms of Service document, specified by the two letter [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language code. For example, **nl** for Dutch. | [optional] 
**TermsOfServiceDocumentId** | Pointer to **string** | The unique identifier of the Terms of Service document. | [optional] 
**Type** | Pointer to **string** | The type of Terms of Service. | [optional] 

## Methods

### NewGetTermsOfServiceDocumentResponse

`func NewGetTermsOfServiceDocumentResponse() *GetTermsOfServiceDocumentResponse`

NewGetTermsOfServiceDocumentResponse instantiates a new GetTermsOfServiceDocumentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTermsOfServiceDocumentResponseWithDefaults

`func NewGetTermsOfServiceDocumentResponseWithDefaults() *GetTermsOfServiceDocumentResponse`

NewGetTermsOfServiceDocumentResponseWithDefaults instantiates a new GetTermsOfServiceDocumentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocument

`func (o *GetTermsOfServiceDocumentResponse) GetDocument() string`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *GetTermsOfServiceDocumentResponse) GetDocumentOk() (*string, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *GetTermsOfServiceDocumentResponse) SetDocument(v string)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *GetTermsOfServiceDocumentResponse) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetId

`func (o *GetTermsOfServiceDocumentResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTermsOfServiceDocumentResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTermsOfServiceDocumentResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetTermsOfServiceDocumentResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLanguage

`func (o *GetTermsOfServiceDocumentResponse) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *GetTermsOfServiceDocumentResponse) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *GetTermsOfServiceDocumentResponse) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *GetTermsOfServiceDocumentResponse) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetTermsOfServiceDocumentId

`func (o *GetTermsOfServiceDocumentResponse) GetTermsOfServiceDocumentId() string`

GetTermsOfServiceDocumentId returns the TermsOfServiceDocumentId field if non-nil, zero value otherwise.

### GetTermsOfServiceDocumentIdOk

`func (o *GetTermsOfServiceDocumentResponse) GetTermsOfServiceDocumentIdOk() (*string, bool)`

GetTermsOfServiceDocumentIdOk returns a tuple with the TermsOfServiceDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfServiceDocumentId

`func (o *GetTermsOfServiceDocumentResponse) SetTermsOfServiceDocumentId(v string)`

SetTermsOfServiceDocumentId sets TermsOfServiceDocumentId field to given value.

### HasTermsOfServiceDocumentId

`func (o *GetTermsOfServiceDocumentResponse) HasTermsOfServiceDocumentId() bool`

HasTermsOfServiceDocumentId returns a boolean if a field has been set.

### GetType

`func (o *GetTermsOfServiceDocumentResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetTermsOfServiceDocumentResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetTermsOfServiceDocumentResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetTermsOfServiceDocumentResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


