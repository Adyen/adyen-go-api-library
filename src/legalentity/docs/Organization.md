# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateOfIncorporation** | Pointer to **string** | The date when the organization was incorporated in YYYY-MM-DD format. | [optional] 
**Description** | Pointer to **string** | Your description for the organization. | [optional] 
**DoingBusinessAs** | Pointer to **string** | The organization&#39;s trading name, if different from the registered legal name. | [optional] 
**Email** | Pointer to **string** | The email address of the legal entity. | [optional] 
**LegalName** | **string** | The organization&#39;s legal name. | 
**Phone** | Pointer to [**PhoneNumber**](PhoneNumber.md) |  | [optional] 
**PrincipalPlaceOfBusiness** | Pointer to [**Address**](Address.md) |  | [optional] 
**RegisteredAddress** | [**Address**](Address.md) |  | 
**RegistrationNumber** | Pointer to **string** | The organization&#39;s registration number. | [optional] 
**StockData** | Pointer to [**StockData**](StockData.md) |  | [optional] 
**TaxInformation** | Pointer to [**[]TaxInformation**](TaxInformation.md) | The tax information of the organization. | [optional] 
**TaxReportingClassification** | Pointer to [**TaxReportingClassification**](TaxReportingClassification.md) |  | [optional] 
**Type** | Pointer to **string** | Type of organization.   Possible values: **associationIncorporated**, **governmentalOrganization**, **listedPublicCompany**, **nonProfit**, **partnershipIncorporated**, **privateCompany**. | [optional] 
**VatAbsenceReason** | Pointer to **string** | The reason the organization has not provided a VAT number.  Possible values: **industryExemption**, **belowTaxThreshold**. | [optional] 
**VatNumber** | Pointer to **string** | The organization&#39;s VAT number. | [optional] 
**WebData** | Pointer to [**WebData**](WebData.md) |  | [optional] 

## Methods

### NewOrganization

`func NewOrganization(legalName string, registeredAddress Address, ) *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateOfIncorporation

`func (o *Organization) GetDateOfIncorporation() string`

GetDateOfIncorporation returns the DateOfIncorporation field if non-nil, zero value otherwise.

### GetDateOfIncorporationOk

`func (o *Organization) GetDateOfIncorporationOk() (*string, bool)`

GetDateOfIncorporationOk returns a tuple with the DateOfIncorporation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfIncorporation

`func (o *Organization) SetDateOfIncorporation(v string)`

SetDateOfIncorporation sets DateOfIncorporation field to given value.

### HasDateOfIncorporation

`func (o *Organization) HasDateOfIncorporation() bool`

HasDateOfIncorporation returns a boolean if a field has been set.

### GetDescription

`func (o *Organization) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Organization) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Organization) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Organization) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDoingBusinessAs

`func (o *Organization) GetDoingBusinessAs() string`

GetDoingBusinessAs returns the DoingBusinessAs field if non-nil, zero value otherwise.

### GetDoingBusinessAsOk

`func (o *Organization) GetDoingBusinessAsOk() (*string, bool)`

GetDoingBusinessAsOk returns a tuple with the DoingBusinessAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoingBusinessAs

`func (o *Organization) SetDoingBusinessAs(v string)`

SetDoingBusinessAs sets DoingBusinessAs field to given value.

### HasDoingBusinessAs

`func (o *Organization) HasDoingBusinessAs() bool`

HasDoingBusinessAs returns a boolean if a field has been set.

### GetEmail

`func (o *Organization) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Organization) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Organization) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Organization) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLegalName

`func (o *Organization) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *Organization) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *Organization) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.


### GetPhone

`func (o *Organization) GetPhone() PhoneNumber`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Organization) GetPhoneOk() (*PhoneNumber, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Organization) SetPhone(v PhoneNumber)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Organization) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPrincipalPlaceOfBusiness

`func (o *Organization) GetPrincipalPlaceOfBusiness() Address`

GetPrincipalPlaceOfBusiness returns the PrincipalPlaceOfBusiness field if non-nil, zero value otherwise.

### GetPrincipalPlaceOfBusinessOk

`func (o *Organization) GetPrincipalPlaceOfBusinessOk() (*Address, bool)`

GetPrincipalPlaceOfBusinessOk returns a tuple with the PrincipalPlaceOfBusiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalPlaceOfBusiness

`func (o *Organization) SetPrincipalPlaceOfBusiness(v Address)`

SetPrincipalPlaceOfBusiness sets PrincipalPlaceOfBusiness field to given value.

### HasPrincipalPlaceOfBusiness

`func (o *Organization) HasPrincipalPlaceOfBusiness() bool`

HasPrincipalPlaceOfBusiness returns a boolean if a field has been set.

### GetRegisteredAddress

`func (o *Organization) GetRegisteredAddress() Address`

GetRegisteredAddress returns the RegisteredAddress field if non-nil, zero value otherwise.

### GetRegisteredAddressOk

`func (o *Organization) GetRegisteredAddressOk() (*Address, bool)`

GetRegisteredAddressOk returns a tuple with the RegisteredAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredAddress

`func (o *Organization) SetRegisteredAddress(v Address)`

SetRegisteredAddress sets RegisteredAddress field to given value.


### GetRegistrationNumber

`func (o *Organization) GetRegistrationNumber() string`

GetRegistrationNumber returns the RegistrationNumber field if non-nil, zero value otherwise.

### GetRegistrationNumberOk

`func (o *Organization) GetRegistrationNumberOk() (*string, bool)`

GetRegistrationNumberOk returns a tuple with the RegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationNumber

`func (o *Organization) SetRegistrationNumber(v string)`

SetRegistrationNumber sets RegistrationNumber field to given value.

### HasRegistrationNumber

`func (o *Organization) HasRegistrationNumber() bool`

HasRegistrationNumber returns a boolean if a field has been set.

### GetStockData

`func (o *Organization) GetStockData() StockData`

GetStockData returns the StockData field if non-nil, zero value otherwise.

### GetStockDataOk

`func (o *Organization) GetStockDataOk() (*StockData, bool)`

GetStockDataOk returns a tuple with the StockData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockData

`func (o *Organization) SetStockData(v StockData)`

SetStockData sets StockData field to given value.

### HasStockData

`func (o *Organization) HasStockData() bool`

HasStockData returns a boolean if a field has been set.

### GetTaxInformation

`func (o *Organization) GetTaxInformation() []TaxInformation`

GetTaxInformation returns the TaxInformation field if non-nil, zero value otherwise.

### GetTaxInformationOk

`func (o *Organization) GetTaxInformationOk() (*[]TaxInformation, bool)`

GetTaxInformationOk returns a tuple with the TaxInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxInformation

`func (o *Organization) SetTaxInformation(v []TaxInformation)`

SetTaxInformation sets TaxInformation field to given value.

### HasTaxInformation

`func (o *Organization) HasTaxInformation() bool`

HasTaxInformation returns a boolean if a field has been set.

### GetTaxReportingClassification

`func (o *Organization) GetTaxReportingClassification() TaxReportingClassification`

GetTaxReportingClassification returns the TaxReportingClassification field if non-nil, zero value otherwise.

### GetTaxReportingClassificationOk

`func (o *Organization) GetTaxReportingClassificationOk() (*TaxReportingClassification, bool)`

GetTaxReportingClassificationOk returns a tuple with the TaxReportingClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxReportingClassification

`func (o *Organization) SetTaxReportingClassification(v TaxReportingClassification)`

SetTaxReportingClassification sets TaxReportingClassification field to given value.

### HasTaxReportingClassification

`func (o *Organization) HasTaxReportingClassification() bool`

HasTaxReportingClassification returns a boolean if a field has been set.

### GetType

`func (o *Organization) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Organization) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Organization) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Organization) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVatAbsenceReason

`func (o *Organization) GetVatAbsenceReason() string`

GetVatAbsenceReason returns the VatAbsenceReason field if non-nil, zero value otherwise.

### GetVatAbsenceReasonOk

`func (o *Organization) GetVatAbsenceReasonOk() (*string, bool)`

GetVatAbsenceReasonOk returns a tuple with the VatAbsenceReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatAbsenceReason

`func (o *Organization) SetVatAbsenceReason(v string)`

SetVatAbsenceReason sets VatAbsenceReason field to given value.

### HasVatAbsenceReason

`func (o *Organization) HasVatAbsenceReason() bool`

HasVatAbsenceReason returns a boolean if a field has been set.

### GetVatNumber

`func (o *Organization) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *Organization) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *Organization) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *Organization) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### GetWebData

`func (o *Organization) GetWebData() WebData`

GetWebData returns the WebData field if non-nil, zero value otherwise.

### GetWebDataOk

`func (o *Organization) GetWebDataOk() (*WebData, bool)`

GetWebDataOk returns a tuple with the WebData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebData

`func (o *Organization) SetWebData(v WebData)`

SetWebData sets WebData field to given value.

### HasWebData

`func (o *Organization) HasWebData() bool`

HasWebData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


