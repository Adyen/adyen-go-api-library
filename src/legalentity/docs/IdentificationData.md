# IdentificationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardNumber** | Pointer to **string** | The card number of the document that was issued (AU only). | [optional] 
**ExpiryDate** | Pointer to **string** | The expiry date of the document, in YYYY-MM-DD format. | [optional] 
**IssuerCountry** | Pointer to **string** | The two-character [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code where the document was issued. For example, **US**. | [optional] 
**IssuerState** | Pointer to **string** | The state or province where the document was issued (AU only). | [optional] 
**NationalIdExempt** | Pointer to **bool** | Applies only to individuals in the US. Set to **true** if the individual does not have an SSN. To verify their identity, Adyen will require them to upload an ID document. | [optional] 
**Number** | Pointer to **string** | The number in the document. | [optional] 
**Type** | **string** | Type of document, used when providing an ID number or uploading a document. The possible values depend on the legal entity type.  When providing ID numbers: * For **individual**, the &#x60;type&#x60; values can be **driversLicense**, **identityCard**, **nationalIdNumber**, or **passport**.  When uploading documents: * For **organization**, the &#x60;type&#x60; values can be **proofOfAddress**, **registrationDocument**, **vatDocument**, **proofOfOrganizationTaxInfo**, **proofOfOwnership**, or **proofOfIndustry**.   * For **individual**, the &#x60;type&#x60; values can be **identityCard**, **driversLicense**, **passport**, **proofOfNationalIdNumber**, **proofOfResidency**, **proofOfIndustry**, or **proofOfIndividualTaxId**.  * For **soleProprietorship**, the &#x60;type&#x60; values can be **constitutionalDocument**, **proofOfAddress**, or **proofOfIndustry**.  * Use **bankStatement** to upload documents for a [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments__resParam_id). | 

## Methods

### NewIdentificationData

`func NewIdentificationData(type_ string, ) *IdentificationData`

NewIdentificationData instantiates a new IdentificationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentificationDataWithDefaults

`func NewIdentificationDataWithDefaults() *IdentificationData`

NewIdentificationDataWithDefaults instantiates a new IdentificationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardNumber

`func (o *IdentificationData) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *IdentificationData) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *IdentificationData) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *IdentificationData) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetExpiryDate

`func (o *IdentificationData) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *IdentificationData) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *IdentificationData) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *IdentificationData) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetIssuerCountry

`func (o *IdentificationData) GetIssuerCountry() string`

GetIssuerCountry returns the IssuerCountry field if non-nil, zero value otherwise.

### GetIssuerCountryOk

`func (o *IdentificationData) GetIssuerCountryOk() (*string, bool)`

GetIssuerCountryOk returns a tuple with the IssuerCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCountry

`func (o *IdentificationData) SetIssuerCountry(v string)`

SetIssuerCountry sets IssuerCountry field to given value.

### HasIssuerCountry

`func (o *IdentificationData) HasIssuerCountry() bool`

HasIssuerCountry returns a boolean if a field has been set.

### GetIssuerState

`func (o *IdentificationData) GetIssuerState() string`

GetIssuerState returns the IssuerState field if non-nil, zero value otherwise.

### GetIssuerStateOk

`func (o *IdentificationData) GetIssuerStateOk() (*string, bool)`

GetIssuerStateOk returns a tuple with the IssuerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerState

`func (o *IdentificationData) SetIssuerState(v string)`

SetIssuerState sets IssuerState field to given value.

### HasIssuerState

`func (o *IdentificationData) HasIssuerState() bool`

HasIssuerState returns a boolean if a field has been set.

### GetNationalIdExempt

`func (o *IdentificationData) GetNationalIdExempt() bool`

GetNationalIdExempt returns the NationalIdExempt field if non-nil, zero value otherwise.

### GetNationalIdExemptOk

`func (o *IdentificationData) GetNationalIdExemptOk() (*bool, bool)`

GetNationalIdExemptOk returns a tuple with the NationalIdExempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdExempt

`func (o *IdentificationData) SetNationalIdExempt(v bool)`

SetNationalIdExempt sets NationalIdExempt field to given value.

### HasNationalIdExempt

`func (o *IdentificationData) HasNationalIdExempt() bool`

HasNationalIdExempt returns a boolean if a field has been set.

### GetNumber

`func (o *IdentificationData) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *IdentificationData) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *IdentificationData) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *IdentificationData) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetType

`func (o *IdentificationData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdentificationData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdentificationData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


