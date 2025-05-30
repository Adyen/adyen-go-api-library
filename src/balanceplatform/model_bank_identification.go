/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the BankIdentification type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BankIdentification{}

// BankIdentification struct for BankIdentification
type BankIdentification struct {
	// Two-character [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code.
	Country *string `json:"country,omitempty"`
	// The bank identification code.
	Identification *string `json:"identification,omitempty"`
	// The type of the identification.  Possible values: **iban**, **routingNumber**, **sortCode**, **bic**.
	IdentificationType *string `json:"identificationType,omitempty"`
}

// NewBankIdentification instantiates a new BankIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankIdentification() *BankIdentification {
	this := BankIdentification{}
	return &this
}

// NewBankIdentificationWithDefaults instantiates a new BankIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankIdentificationWithDefaults() *BankIdentification {
	this := BankIdentification{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *BankIdentification) GetCountry() string {
	if o == nil || common.IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankIdentification) GetCountryOk() (*string, bool) {
	if o == nil || common.IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *BankIdentification) HasCountry() bool {
	if o != nil && !common.IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *BankIdentification) SetCountry(v string) {
	o.Country = &v
}

// GetIdentification returns the Identification field value if set, zero value otherwise.
func (o *BankIdentification) GetIdentification() string {
	if o == nil || common.IsNil(o.Identification) {
		var ret string
		return ret
	}
	return *o.Identification
}

// GetIdentificationOk returns a tuple with the Identification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankIdentification) GetIdentificationOk() (*string, bool) {
	if o == nil || common.IsNil(o.Identification) {
		return nil, false
	}
	return o.Identification, true
}

// HasIdentification returns a boolean if a field has been set.
func (o *BankIdentification) HasIdentification() bool {
	if o != nil && !common.IsNil(o.Identification) {
		return true
	}

	return false
}

// SetIdentification gets a reference to the given string and assigns it to the Identification field.
func (o *BankIdentification) SetIdentification(v string) {
	o.Identification = &v
}

// GetIdentificationType returns the IdentificationType field value if set, zero value otherwise.
func (o *BankIdentification) GetIdentificationType() string {
	if o == nil || common.IsNil(o.IdentificationType) {
		var ret string
		return ret
	}
	return *o.IdentificationType
}

// GetIdentificationTypeOk returns a tuple with the IdentificationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankIdentification) GetIdentificationTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.IdentificationType) {
		return nil, false
	}
	return o.IdentificationType, true
}

// HasIdentificationType returns a boolean if a field has been set.
func (o *BankIdentification) HasIdentificationType() bool {
	if o != nil && !common.IsNil(o.IdentificationType) {
		return true
	}

	return false
}

// SetIdentificationType gets a reference to the given string and assigns it to the IdentificationType field.
func (o *BankIdentification) SetIdentificationType(v string) {
	o.IdentificationType = &v
}

func (o BankIdentification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BankIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !common.IsNil(o.Identification) {
		toSerialize["identification"] = o.Identification
	}
	if !common.IsNil(o.IdentificationType) {
		toSerialize["identificationType"] = o.IdentificationType
	}
	return toSerialize, nil
}

type NullableBankIdentification struct {
	value *BankIdentification
	isSet bool
}

func (v NullableBankIdentification) Get() *BankIdentification {
	return v.value
}

func (v *NullableBankIdentification) Set(val *BankIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableBankIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableBankIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankIdentification(val *BankIdentification) *NullableBankIdentification {
	return &NullableBankIdentification{value: val, isSet: true}
}

func (v NullableBankIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *BankIdentification) isValidIdentificationType() bool {
	var allowedEnumValues = []string{"bic", "iban", "routingNumber", "sortCode"}
	for _, allowed := range allowedEnumValues {
		if o.GetIdentificationType() == allowed {
			return true
		}
	}
	return false
}
