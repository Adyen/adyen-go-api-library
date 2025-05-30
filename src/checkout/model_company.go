/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the Company type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Company{}

// Company struct for Company
type Company struct {
	// The company website's home page.
	Homepage *string `json:"homepage,omitempty"`
	// The company name.
	Name *string `json:"name,omitempty"`
	// Registration number of the company.
	RegistrationNumber *string `json:"registrationNumber,omitempty"`
	// Registry location of the company.
	RegistryLocation *string `json:"registryLocation,omitempty"`
	// Tax ID of the company.
	TaxId *string `json:"taxId,omitempty"`
	// The company type.
	Type *string `json:"type,omitempty"`
}

// NewCompany instantiates a new Company object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompany() *Company {
	this := Company{}
	return &this
}

// NewCompanyWithDefaults instantiates a new Company object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyWithDefaults() *Company {
	this := Company{}
	return &this
}

// GetHomepage returns the Homepage field value if set, zero value otherwise.
func (o *Company) GetHomepage() string {
	if o == nil || common.IsNil(o.Homepage) {
		var ret string
		return ret
	}
	return *o.Homepage
}

// GetHomepageOk returns a tuple with the Homepage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetHomepageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Homepage) {
		return nil, false
	}
	return o.Homepage, true
}

// HasHomepage returns a boolean if a field has been set.
func (o *Company) HasHomepage() bool {
	if o != nil && !common.IsNil(o.Homepage) {
		return true
	}

	return false
}

// SetHomepage gets a reference to the given string and assigns it to the Homepage field.
func (o *Company) SetHomepage(v string) {
	o.Homepage = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Company) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Company) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Company) SetName(v string) {
	o.Name = &v
}

// GetRegistrationNumber returns the RegistrationNumber field value if set, zero value otherwise.
func (o *Company) GetRegistrationNumber() string {
	if o == nil || common.IsNil(o.RegistrationNumber) {
		var ret string
		return ret
	}
	return *o.RegistrationNumber
}

// GetRegistrationNumberOk returns a tuple with the RegistrationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetRegistrationNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.RegistrationNumber) {
		return nil, false
	}
	return o.RegistrationNumber, true
}

// HasRegistrationNumber returns a boolean if a field has been set.
func (o *Company) HasRegistrationNumber() bool {
	if o != nil && !common.IsNil(o.RegistrationNumber) {
		return true
	}

	return false
}

// SetRegistrationNumber gets a reference to the given string and assigns it to the RegistrationNumber field.
func (o *Company) SetRegistrationNumber(v string) {
	o.RegistrationNumber = &v
}

// GetRegistryLocation returns the RegistryLocation field value if set, zero value otherwise.
func (o *Company) GetRegistryLocation() string {
	if o == nil || common.IsNil(o.RegistryLocation) {
		var ret string
		return ret
	}
	return *o.RegistryLocation
}

// GetRegistryLocationOk returns a tuple with the RegistryLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetRegistryLocationOk() (*string, bool) {
	if o == nil || common.IsNil(o.RegistryLocation) {
		return nil, false
	}
	return o.RegistryLocation, true
}

// HasRegistryLocation returns a boolean if a field has been set.
func (o *Company) HasRegistryLocation() bool {
	if o != nil && !common.IsNil(o.RegistryLocation) {
		return true
	}

	return false
}

// SetRegistryLocation gets a reference to the given string and assigns it to the RegistryLocation field.
func (o *Company) SetRegistryLocation(v string) {
	o.RegistryLocation = &v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *Company) GetTaxId() string {
	if o == nil || common.IsNil(o.TaxId) {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetTaxIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TaxId) {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *Company) HasTaxId() bool {
	if o != nil && !common.IsNil(o.TaxId) {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *Company) SetTaxId(v string) {
	o.TaxId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Company) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Company) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Company) SetType(v string) {
	o.Type = &v
}

func (o Company) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Company) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Homepage) {
		toSerialize["homepage"] = o.Homepage
	}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.RegistrationNumber) {
		toSerialize["registrationNumber"] = o.RegistrationNumber
	}
	if !common.IsNil(o.RegistryLocation) {
		toSerialize["registryLocation"] = o.RegistryLocation
	}
	if !common.IsNil(o.TaxId) {
		toSerialize["taxId"] = o.TaxId
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableCompany struct {
	value *Company
	isSet bool
}

func (v NullableCompany) Get() *Company {
	return v.value
}

func (v *NullableCompany) Set(val *Company) {
	v.value = val
	v.isSet = true
}

func (v NullableCompany) IsSet() bool {
	return v.isSet
}

func (v *NullableCompany) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompany(val *Company) *NullableCompany {
	return &NullableCompany{value: val, isSet: true}
}

func (v NullableCompany) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompany) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
