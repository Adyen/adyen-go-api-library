/*
Legal Entity Management API

The Legal Entity Management API enables you to manage legal entities that contain information required for verification.  ## Authentication To connect to the Legal Entity Management API, you must use the basic authentication credentials of your web service user. If you don't have one, contact the [Adyen Support Team](https://www.adyen.help/hc/en-us/requests/new). Use the web service user credentials to authenticate your request, for example:  ``` curl -U \"ws12345@Scope.BalancePlatform_YourBalancePlatform\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning The Legal Entity Management API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://kyc-test.adyen.com/lem/v2/legalEntities ``` ## Going live When going live, your Adyen contact will provide your API credential for the live environment. You can then use the username and password to send requests to `https://kyc-live.adyen.com/lem/v2`.  

API version: 2
Contact: developer-experience@adyen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"
)

// checks if the SoleProprietorship type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SoleProprietorship{}

// SoleProprietorship struct for SoleProprietorship
type SoleProprietorship struct {
	// The two-character [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the governing country.
	CountryOfGoverningLaw string `json:"countryOfGoverningLaw"`
	// The date when the legal arrangement was incorporated in YYYY-MM-DD format.
	DateOfIncorporation *string `json:"dateOfIncorporation,omitempty"`
	// The registered name, if different from the `name`.
	DoingBusinessAs *string `json:"doingBusinessAs,omitempty"`
	// The legal name.
	Name string `json:"name"`
	PrincipalPlaceOfBusiness *Address `json:"principalPlaceOfBusiness,omitempty"`
	RegisteredAddress Address `json:"registeredAddress"`
	// The registration number.
	RegistrationNumber *string `json:"registrationNumber,omitempty"`
	// The reason for not providing a VAT number.  Possible values: **industryExemption**, **belowTaxThreshold**.
	VatAbsenceReason *string `json:"vatAbsenceReason,omitempty"`
	// The VAT number.
	VatNumber *string `json:"vatNumber,omitempty"`
}

// NewSoleProprietorship instantiates a new SoleProprietorship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoleProprietorship(countryOfGoverningLaw string, name string, registeredAddress Address) *SoleProprietorship {
	this := SoleProprietorship{}
	this.CountryOfGoverningLaw = countryOfGoverningLaw
	this.Name = name
	this.RegisteredAddress = registeredAddress
	return &this
}

// NewSoleProprietorshipWithDefaults instantiates a new SoleProprietorship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoleProprietorshipWithDefaults() *SoleProprietorship {
	this := SoleProprietorship{}
	return &this
}

// GetCountryOfGoverningLaw returns the CountryOfGoverningLaw field value
func (o *SoleProprietorship) GetCountryOfGoverningLaw() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryOfGoverningLaw
}

// GetCountryOfGoverningLawOk returns a tuple with the CountryOfGoverningLaw field value
// and a boolean to check if the value has been set.
func (o *SoleProprietorship) GetCountryOfGoverningLawOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryOfGoverningLaw, true
}

// SetCountryOfGoverningLaw sets field value
func (o *SoleProprietorship) SetCountryOfGoverningLaw(v string) {
	o.CountryOfGoverningLaw = v
}

// GetDateOfIncorporation returns the DateOfIncorporation field value if set, zero value otherwise.
func (o *SoleProprietorship) GetDateOfIncorporation() string {
	if o == nil || IsNil(o.DateOfIncorporation) {
		var ret string
		return ret
	}
	return *o.DateOfIncorporation
}

// GetDateOfIncorporationOk returns a tuple with the DateOfIncorporation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoleProprietorship) GetDateOfIncorporationOk() (*string, bool) {
	if o == nil || IsNil(o.DateOfIncorporation) {
		return nil, false
	}
	return o.DateOfIncorporation, true
}

// HasDateOfIncorporation returns a boolean if a field has been set.
func (o *SoleProprietorship) HasDateOfIncorporation() bool {
	if o != nil && !IsNil(o.DateOfIncorporation) {
		return true
	}

	return false
}

// SetDateOfIncorporation gets a reference to the given string and assigns it to the DateOfIncorporation field.
func (o *SoleProprietorship) SetDateOfIncorporation(v string) {
	o.DateOfIncorporation = &v
}

// GetDoingBusinessAs returns the DoingBusinessAs field value if set, zero value otherwise.
func (o *SoleProprietorship) GetDoingBusinessAs() string {
	if o == nil || IsNil(o.DoingBusinessAs) {
		var ret string
		return ret
	}
	return *o.DoingBusinessAs
}

// GetDoingBusinessAsOk returns a tuple with the DoingBusinessAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoleProprietorship) GetDoingBusinessAsOk() (*string, bool) {
	if o == nil || IsNil(o.DoingBusinessAs) {
		return nil, false
	}
	return o.DoingBusinessAs, true
}

// HasDoingBusinessAs returns a boolean if a field has been set.
func (o *SoleProprietorship) HasDoingBusinessAs() bool {
	if o != nil && !IsNil(o.DoingBusinessAs) {
		return true
	}

	return false
}

// SetDoingBusinessAs gets a reference to the given string and assigns it to the DoingBusinessAs field.
func (o *SoleProprietorship) SetDoingBusinessAs(v string) {
	o.DoingBusinessAs = &v
}

// GetName returns the Name field value
func (o *SoleProprietorship) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SoleProprietorship) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SoleProprietorship) SetName(v string) {
	o.Name = v
}

// GetPrincipalPlaceOfBusiness returns the PrincipalPlaceOfBusiness field value if set, zero value otherwise.
func (o *SoleProprietorship) GetPrincipalPlaceOfBusiness() Address {
	if o == nil || IsNil(o.PrincipalPlaceOfBusiness) {
		var ret Address
		return ret
	}
	return *o.PrincipalPlaceOfBusiness
}

// GetPrincipalPlaceOfBusinessOk returns a tuple with the PrincipalPlaceOfBusiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoleProprietorship) GetPrincipalPlaceOfBusinessOk() (*Address, bool) {
	if o == nil || IsNil(o.PrincipalPlaceOfBusiness) {
		return nil, false
	}
	return o.PrincipalPlaceOfBusiness, true
}

// HasPrincipalPlaceOfBusiness returns a boolean if a field has been set.
func (o *SoleProprietorship) HasPrincipalPlaceOfBusiness() bool {
	if o != nil && !IsNil(o.PrincipalPlaceOfBusiness) {
		return true
	}

	return false
}

// SetPrincipalPlaceOfBusiness gets a reference to the given Address and assigns it to the PrincipalPlaceOfBusiness field.
func (o *SoleProprietorship) SetPrincipalPlaceOfBusiness(v Address) {
	o.PrincipalPlaceOfBusiness = &v
}

// GetRegisteredAddress returns the RegisteredAddress field value
func (o *SoleProprietorship) GetRegisteredAddress() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.RegisteredAddress
}

// GetRegisteredAddressOk returns a tuple with the RegisteredAddress field value
// and a boolean to check if the value has been set.
func (o *SoleProprietorship) GetRegisteredAddressOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegisteredAddress, true
}

// SetRegisteredAddress sets field value
func (o *SoleProprietorship) SetRegisteredAddress(v Address) {
	o.RegisteredAddress = v
}

// GetRegistrationNumber returns the RegistrationNumber field value if set, zero value otherwise.
func (o *SoleProprietorship) GetRegistrationNumber() string {
	if o == nil || IsNil(o.RegistrationNumber) {
		var ret string
		return ret
	}
	return *o.RegistrationNumber
}

// GetRegistrationNumberOk returns a tuple with the RegistrationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoleProprietorship) GetRegistrationNumberOk() (*string, bool) {
	if o == nil || IsNil(o.RegistrationNumber) {
		return nil, false
	}
	return o.RegistrationNumber, true
}

// HasRegistrationNumber returns a boolean if a field has been set.
func (o *SoleProprietorship) HasRegistrationNumber() bool {
	if o != nil && !IsNil(o.RegistrationNumber) {
		return true
	}

	return false
}

// SetRegistrationNumber gets a reference to the given string and assigns it to the RegistrationNumber field.
func (o *SoleProprietorship) SetRegistrationNumber(v string) {
	o.RegistrationNumber = &v
}

// GetVatAbsenceReason returns the VatAbsenceReason field value if set, zero value otherwise.
func (o *SoleProprietorship) GetVatAbsenceReason() string {
	if o == nil || IsNil(o.VatAbsenceReason) {
		var ret string
		return ret
	}
	return *o.VatAbsenceReason
}

// GetVatAbsenceReasonOk returns a tuple with the VatAbsenceReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoleProprietorship) GetVatAbsenceReasonOk() (*string, bool) {
	if o == nil || IsNil(o.VatAbsenceReason) {
		return nil, false
	}
	return o.VatAbsenceReason, true
}

// HasVatAbsenceReason returns a boolean if a field has been set.
func (o *SoleProprietorship) HasVatAbsenceReason() bool {
	if o != nil && !IsNil(o.VatAbsenceReason) {
		return true
	}

	return false
}

// SetVatAbsenceReason gets a reference to the given string and assigns it to the VatAbsenceReason field.
func (o *SoleProprietorship) SetVatAbsenceReason(v string) {
	o.VatAbsenceReason = &v
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise.
func (o *SoleProprietorship) GetVatNumber() string {
	if o == nil || IsNil(o.VatNumber) {
		var ret string
		return ret
	}
	return *o.VatNumber
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoleProprietorship) GetVatNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VatNumber) {
		return nil, false
	}
	return o.VatNumber, true
}

// HasVatNumber returns a boolean if a field has been set.
func (o *SoleProprietorship) HasVatNumber() bool {
	if o != nil && !IsNil(o.VatNumber) {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given string and assigns it to the VatNumber field.
func (o *SoleProprietorship) SetVatNumber(v string) {
	o.VatNumber = &v
}

func (o SoleProprietorship) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SoleProprietorship) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["countryOfGoverningLaw"] = o.CountryOfGoverningLaw
	if !IsNil(o.DateOfIncorporation) {
		toSerialize["dateOfIncorporation"] = o.DateOfIncorporation
	}
	if !IsNil(o.DoingBusinessAs) {
		toSerialize["doingBusinessAs"] = o.DoingBusinessAs
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.PrincipalPlaceOfBusiness) {
		toSerialize["principalPlaceOfBusiness"] = o.PrincipalPlaceOfBusiness
	}
	toSerialize["registeredAddress"] = o.RegisteredAddress
	if !IsNil(o.RegistrationNumber) {
		toSerialize["registrationNumber"] = o.RegistrationNumber
	}
	if !IsNil(o.VatAbsenceReason) {
		toSerialize["vatAbsenceReason"] = o.VatAbsenceReason
	}
	if !IsNil(o.VatNumber) {
		toSerialize["vatNumber"] = o.VatNumber
	}
	return toSerialize, nil
}

type NullableSoleProprietorship struct {
	value *SoleProprietorship
	isSet bool
}

func (v NullableSoleProprietorship) Get() *SoleProprietorship {
	return v.value
}

func (v *NullableSoleProprietorship) Set(val *SoleProprietorship) {
	v.value = val
	v.isSet = true
}

func (v NullableSoleProprietorship) IsSet() bool {
	return v.isSet
}

func (v *NullableSoleProprietorship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoleProprietorship(val *SoleProprietorship) *NullableSoleProprietorship {
	return &NullableSoleProprietorship{value: val, isSet: true}
}

func (v NullableSoleProprietorship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoleProprietorship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

