# Attachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | The document in Base64-encoded string format. | 
**ContentType** | Pointer to **string** | The file format.   Possible values: **application/pdf**, **image/jpg**, **image/jpeg**, **image/png**.  | [optional] 
**Filename** | Pointer to **string** | The name of the file including the file extension. | [optional] 
**PageName** | Pointer to **string** | The name of the file including the file extension. | [optional] 
**PageType** | Pointer to **string** | Specifies which side of the ID card is uploaded.  * When &#x60;type&#x60; is **driversLicense** or **identityCard**, set this to **front** or **back**.  * When omitted, we infer the page number based on the order of attachments. | [optional] 

## Methods

### NewAttachment

`func NewAttachment(content string, ) *Attachment`

NewAttachment instantiates a new Attachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentWithDefaults

`func NewAttachmentWithDefaults() *Attachment`

NewAttachmentWithDefaults instantiates a new Attachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *Attachment) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Attachment) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Attachment) SetContent(v string)`

SetContent sets Content field to given value.


### GetContentType

`func (o *Attachment) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *Attachment) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *Attachment) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *Attachment) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetFilename

`func (o *Attachment) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *Attachment) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *Attachment) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *Attachment) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetPageName

`func (o *Attachment) GetPageName() string`

GetPageName returns the PageName field if non-nil, zero value otherwise.

### GetPageNameOk

`func (o *Attachment) GetPageNameOk() (*string, bool)`

GetPageNameOk returns a tuple with the PageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageName

`func (o *Attachment) SetPageName(v string)`

SetPageName sets PageName field to given value.

### HasPageName

`func (o *Attachment) HasPageName() bool`

HasPageName returns a boolean if a field has been set.

### GetPageType

`func (o *Attachment) GetPageType() string`

GetPageType returns the PageType field if non-nil, zero value otherwise.

### GetPageTypeOk

`func (o *Attachment) GetPageTypeOk() (*string, bool)`

GetPageTypeOk returns a tuple with the PageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageType

`func (o *Attachment) SetPageType(v string)`

SetPageType sets PageType field to given value.

### HasPageType

`func (o *Attachment) HasPageType() bool`

HasPageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


