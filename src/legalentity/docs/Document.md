# Document

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachment** | Pointer to [**Attachment**](Attachment.md) |  | [optional] 
**Attachments** | [**[]Attachment**](Attachment.md) | Array that contains the document. The array supports multiple attachments for uploading different sides or pages of a document. | 
**CreationDate** | Pointer to **time.Time** | The creation date of the document. | [optional] [readonly] 
**Description** | **string** | Your description for the document. | 
**ExpiryDate** | Pointer to **string** | The expiry date of the document, in YYYY-MM-DD format. | [optional] 
**FileName** | Pointer to **string** | The filename of the document. | [optional] 
**Id** | **string** | The unique identifier of the document. | [readonly] 
**IssuerCountry** | Pointer to **string** | The two-character [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code where the document was issued. For example, **US**. | [optional] 
**IssuerState** | Pointer to **string** | The state or province where the document was issued (AU only). | [optional] 
**ModificationDate** | Pointer to **time.Time** | The modification date of the document. | [optional] [readonly] 
**Number** | Pointer to **string** | The number in the document. | [optional] 
**Owner** | [**OwnerEntity**](OwnerEntity.md) |  | 
**Type** | **string** | Type of document, used when providing an ID number or uploading a document. The possible values depend on the legal entity type.  When providing ID numbers: * For **individual**, the &#x60;type&#x60; values can be **driversLicense**, **identityCard**, **nationalIdNumber**, or **passport**.  When uploading documents: * For **organization**, the &#x60;type&#x60; values can be **proofOfAddress**, **registrationDocument**, **vatDocument**, **proofOfOrganizationTaxInfo**, **proofOfOwnership**, or **proofOfIndustry**.   * For **individual**, the &#x60;type&#x60; values can be **identityCard**, **driversLicense**, **passport**, **proofOfNationalIdNumber**, **proofOfResidency**, **proofOfIndustry**, or **proofOfIndividualTaxId**.  * For **soleProprietorship**, the &#x60;type&#x60; values can be **constitutionalDocument**, **proofOfAddress**, or **proofOfIndustry**.  * Use **bankStatement** to upload documents for a [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments__resParam_id). | 

## Methods

### NewDocument

`func NewDocument(attachments []Attachment, description string, id string, owner OwnerEntity, type_ string, ) *Document`

NewDocument instantiates a new Document object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentWithDefaults

`func NewDocumentWithDefaults() *Document`

NewDocumentWithDefaults instantiates a new Document object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachment

`func (o *Document) GetAttachment() Attachment`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *Document) GetAttachmentOk() (*Attachment, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *Document) SetAttachment(v Attachment)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *Document) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.

### GetAttachments

`func (o *Document) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *Document) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *Document) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.


### GetCreationDate

`func (o *Document) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Document) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Document) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *Document) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDescription

`func (o *Document) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Document) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Document) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetExpiryDate

`func (o *Document) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *Document) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *Document) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *Document) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetFileName

`func (o *Document) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *Document) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *Document) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *Document) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetId

`func (o *Document) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Document) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Document) SetId(v string)`

SetId sets Id field to given value.


### GetIssuerCountry

`func (o *Document) GetIssuerCountry() string`

GetIssuerCountry returns the IssuerCountry field if non-nil, zero value otherwise.

### GetIssuerCountryOk

`func (o *Document) GetIssuerCountryOk() (*string, bool)`

GetIssuerCountryOk returns a tuple with the IssuerCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCountry

`func (o *Document) SetIssuerCountry(v string)`

SetIssuerCountry sets IssuerCountry field to given value.

### HasIssuerCountry

`func (o *Document) HasIssuerCountry() bool`

HasIssuerCountry returns a boolean if a field has been set.

### GetIssuerState

`func (o *Document) GetIssuerState() string`

GetIssuerState returns the IssuerState field if non-nil, zero value otherwise.

### GetIssuerStateOk

`func (o *Document) GetIssuerStateOk() (*string, bool)`

GetIssuerStateOk returns a tuple with the IssuerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerState

`func (o *Document) SetIssuerState(v string)`

SetIssuerState sets IssuerState field to given value.

### HasIssuerState

`func (o *Document) HasIssuerState() bool`

HasIssuerState returns a boolean if a field has been set.

### GetModificationDate

`func (o *Document) GetModificationDate() time.Time`

GetModificationDate returns the ModificationDate field if non-nil, zero value otherwise.

### GetModificationDateOk

`func (o *Document) GetModificationDateOk() (*time.Time, bool)`

GetModificationDateOk returns a tuple with the ModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationDate

`func (o *Document) SetModificationDate(v time.Time)`

SetModificationDate sets ModificationDate field to given value.

### HasModificationDate

`func (o *Document) HasModificationDate() bool`

HasModificationDate returns a boolean if a field has been set.

### GetNumber

`func (o *Document) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Document) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Document) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Document) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetOwner

`func (o *Document) GetOwner() OwnerEntity`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Document) GetOwnerOk() (*OwnerEntity, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Document) SetOwner(v OwnerEntity)`

SetOwner sets Owner field to given value.


### GetType

`func (o *Document) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Document) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Document) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


