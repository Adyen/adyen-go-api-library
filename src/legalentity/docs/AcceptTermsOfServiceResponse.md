# AcceptTermsOfServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedBy** | Pointer to **string** | The unique identifier of the user that accepted the Terms of Service. | [optional] 
**Id** | Pointer to **string** | The unique identifier of the Terms of Service acceptance. | [optional] 
**IpAddress** | Pointer to **string** | The IP address of the user that accepted the Terms of Service. | [optional] 
**Language** | Pointer to **string** | The language used for the Terms of Service document, specified by the two letter [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language code. For example, **nl** for Dutch. | [optional] 
**TermsOfServiceDocumentId** | Pointer to **string** | The unique identifier of the Terms of Service document. | [optional] 
**Type** | Pointer to **string** | The type of Terms of Service. | [optional] 

## Methods

### NewAcceptTermsOfServiceResponse

`func NewAcceptTermsOfServiceResponse() *AcceptTermsOfServiceResponse`

NewAcceptTermsOfServiceResponse instantiates a new AcceptTermsOfServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceptTermsOfServiceResponseWithDefaults

`func NewAcceptTermsOfServiceResponseWithDefaults() *AcceptTermsOfServiceResponse`

NewAcceptTermsOfServiceResponseWithDefaults instantiates a new AcceptTermsOfServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedBy

`func (o *AcceptTermsOfServiceResponse) GetAcceptedBy() string`

GetAcceptedBy returns the AcceptedBy field if non-nil, zero value otherwise.

### GetAcceptedByOk

`func (o *AcceptTermsOfServiceResponse) GetAcceptedByOk() (*string, bool)`

GetAcceptedByOk returns a tuple with the AcceptedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedBy

`func (o *AcceptTermsOfServiceResponse) SetAcceptedBy(v string)`

SetAcceptedBy sets AcceptedBy field to given value.

### HasAcceptedBy

`func (o *AcceptTermsOfServiceResponse) HasAcceptedBy() bool`

HasAcceptedBy returns a boolean if a field has been set.

### GetId

`func (o *AcceptTermsOfServiceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AcceptTermsOfServiceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AcceptTermsOfServiceResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AcceptTermsOfServiceResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpAddress

`func (o *AcceptTermsOfServiceResponse) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *AcceptTermsOfServiceResponse) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *AcceptTermsOfServiceResponse) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *AcceptTermsOfServiceResponse) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLanguage

`func (o *AcceptTermsOfServiceResponse) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *AcceptTermsOfServiceResponse) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *AcceptTermsOfServiceResponse) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *AcceptTermsOfServiceResponse) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetTermsOfServiceDocumentId

`func (o *AcceptTermsOfServiceResponse) GetTermsOfServiceDocumentId() string`

GetTermsOfServiceDocumentId returns the TermsOfServiceDocumentId field if non-nil, zero value otherwise.

### GetTermsOfServiceDocumentIdOk

`func (o *AcceptTermsOfServiceResponse) GetTermsOfServiceDocumentIdOk() (*string, bool)`

GetTermsOfServiceDocumentIdOk returns a tuple with the TermsOfServiceDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfServiceDocumentId

`func (o *AcceptTermsOfServiceResponse) SetTermsOfServiceDocumentId(v string)`

SetTermsOfServiceDocumentId sets TermsOfServiceDocumentId field to given value.

### HasTermsOfServiceDocumentId

`func (o *AcceptTermsOfServiceResponse) HasTermsOfServiceDocumentId() bool`

HasTermsOfServiceDocumentId returns a boolean if a field has been set.

### GetType

`func (o *AcceptTermsOfServiceResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AcceptTermsOfServiceResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AcceptTermsOfServiceResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AcceptTermsOfServiceResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


