# SoleProprietorship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryOfGoverningLaw** | **string** | The two-character [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the governing country. | 
**DateOfIncorporation** | Pointer to **string** | The date when the legal arrangement was incorporated in YYYY-MM-DD format. | [optional] 
**DoingBusinessAs** | Pointer to **string** | The registered name, if different from the &#x60;name&#x60;. | [optional] 
**Name** | **string** | The legal name. | 
**PrincipalPlaceOfBusiness** | Pointer to [**Address**](Address.md) |  | [optional] 
**RegisteredAddress** | [**Address**](Address.md) |  | 
**RegistrationNumber** | Pointer to **string** | The registration number. | [optional] 
**VatAbsenceReason** | Pointer to **string** | The reason for not providing a VAT number.  Possible values: **industryExemption**, **belowTaxThreshold**. | [optional] 
**VatNumber** | Pointer to **string** | The VAT number. | [optional] 

## Methods

### NewSoleProprietorship

`func NewSoleProprietorship(countryOfGoverningLaw string, name string, registeredAddress Address, ) *SoleProprietorship`

NewSoleProprietorship instantiates a new SoleProprietorship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoleProprietorshipWithDefaults

`func NewSoleProprietorshipWithDefaults() *SoleProprietorship`

NewSoleProprietorshipWithDefaults instantiates a new SoleProprietorship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryOfGoverningLaw

`func (o *SoleProprietorship) GetCountryOfGoverningLaw() string`

GetCountryOfGoverningLaw returns the CountryOfGoverningLaw field if non-nil, zero value otherwise.

### GetCountryOfGoverningLawOk

`func (o *SoleProprietorship) GetCountryOfGoverningLawOk() (*string, bool)`

GetCountryOfGoverningLawOk returns a tuple with the CountryOfGoverningLaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfGoverningLaw

`func (o *SoleProprietorship) SetCountryOfGoverningLaw(v string)`

SetCountryOfGoverningLaw sets CountryOfGoverningLaw field to given value.


### GetDateOfIncorporation

`func (o *SoleProprietorship) GetDateOfIncorporation() string`

GetDateOfIncorporation returns the DateOfIncorporation field if non-nil, zero value otherwise.

### GetDateOfIncorporationOk

`func (o *SoleProprietorship) GetDateOfIncorporationOk() (*string, bool)`

GetDateOfIncorporationOk returns a tuple with the DateOfIncorporation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfIncorporation

`func (o *SoleProprietorship) SetDateOfIncorporation(v string)`

SetDateOfIncorporation sets DateOfIncorporation field to given value.

### HasDateOfIncorporation

`func (o *SoleProprietorship) HasDateOfIncorporation() bool`

HasDateOfIncorporation returns a boolean if a field has been set.

### GetDoingBusinessAs

`func (o *SoleProprietorship) GetDoingBusinessAs() string`

GetDoingBusinessAs returns the DoingBusinessAs field if non-nil, zero value otherwise.

### GetDoingBusinessAsOk

`func (o *SoleProprietorship) GetDoingBusinessAsOk() (*string, bool)`

GetDoingBusinessAsOk returns a tuple with the DoingBusinessAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoingBusinessAs

`func (o *SoleProprietorship) SetDoingBusinessAs(v string)`

SetDoingBusinessAs sets DoingBusinessAs field to given value.

### HasDoingBusinessAs

`func (o *SoleProprietorship) HasDoingBusinessAs() bool`

HasDoingBusinessAs returns a boolean if a field has been set.

### GetName

`func (o *SoleProprietorship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SoleProprietorship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SoleProprietorship) SetName(v string)`

SetName sets Name field to given value.


### GetPrincipalPlaceOfBusiness

`func (o *SoleProprietorship) GetPrincipalPlaceOfBusiness() Address`

GetPrincipalPlaceOfBusiness returns the PrincipalPlaceOfBusiness field if non-nil, zero value otherwise.

### GetPrincipalPlaceOfBusinessOk

`func (o *SoleProprietorship) GetPrincipalPlaceOfBusinessOk() (*Address, bool)`

GetPrincipalPlaceOfBusinessOk returns a tuple with the PrincipalPlaceOfBusiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalPlaceOfBusiness

`func (o *SoleProprietorship) SetPrincipalPlaceOfBusiness(v Address)`

SetPrincipalPlaceOfBusiness sets PrincipalPlaceOfBusiness field to given value.

### HasPrincipalPlaceOfBusiness

`func (o *SoleProprietorship) HasPrincipalPlaceOfBusiness() bool`

HasPrincipalPlaceOfBusiness returns a boolean if a field has been set.

### GetRegisteredAddress

`func (o *SoleProprietorship) GetRegisteredAddress() Address`

GetRegisteredAddress returns the RegisteredAddress field if non-nil, zero value otherwise.

### GetRegisteredAddressOk

`func (o *SoleProprietorship) GetRegisteredAddressOk() (*Address, bool)`

GetRegisteredAddressOk returns a tuple with the RegisteredAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredAddress

`func (o *SoleProprietorship) SetRegisteredAddress(v Address)`

SetRegisteredAddress sets RegisteredAddress field to given value.


### GetRegistrationNumber

`func (o *SoleProprietorship) GetRegistrationNumber() string`

GetRegistrationNumber returns the RegistrationNumber field if non-nil, zero value otherwise.

### GetRegistrationNumberOk

`func (o *SoleProprietorship) GetRegistrationNumberOk() (*string, bool)`

GetRegistrationNumberOk returns a tuple with the RegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationNumber

`func (o *SoleProprietorship) SetRegistrationNumber(v string)`

SetRegistrationNumber sets RegistrationNumber field to given value.

### HasRegistrationNumber

`func (o *SoleProprietorship) HasRegistrationNumber() bool`

HasRegistrationNumber returns a boolean if a field has been set.

### GetVatAbsenceReason

`func (o *SoleProprietorship) GetVatAbsenceReason() string`

GetVatAbsenceReason returns the VatAbsenceReason field if non-nil, zero value otherwise.

### GetVatAbsenceReasonOk

`func (o *SoleProprietorship) GetVatAbsenceReasonOk() (*string, bool)`

GetVatAbsenceReasonOk returns a tuple with the VatAbsenceReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatAbsenceReason

`func (o *SoleProprietorship) SetVatAbsenceReason(v string)`

SetVatAbsenceReason sets VatAbsenceReason field to given value.

### HasVatAbsenceReason

`func (o *SoleProprietorship) HasVatAbsenceReason() bool`

HasVatAbsenceReason returns a boolean if a field has been set.

### GetVatNumber

`func (o *SoleProprietorship) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *SoleProprietorship) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *SoleProprietorship) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *SoleProprietorship) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


