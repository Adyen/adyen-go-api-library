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

// checks if the USInternationalAchAddressRequirement type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &USInternationalAchAddressRequirement{}

// USInternationalAchAddressRequirement struct for USInternationalAchAddressRequirement
type USInternationalAchAddressRequirement struct {
	// Specifies that you must provide a complete street address for International ACH (IAT) transactions.
	Description *string `json:"description,omitempty"`
	// **usInternationalAchAddressRequirement**
	Type string `json:"type"`
}

// NewUSInternationalAchAddressRequirement instantiates a new USInternationalAchAddressRequirement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUSInternationalAchAddressRequirement(type_ string) *USInternationalAchAddressRequirement {
	this := USInternationalAchAddressRequirement{}
	this.Type = type_
	return &this
}

// NewUSInternationalAchAddressRequirementWithDefaults instantiates a new USInternationalAchAddressRequirement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUSInternationalAchAddressRequirementWithDefaults() *USInternationalAchAddressRequirement {
	this := USInternationalAchAddressRequirement{}
	var type_ string = "usInternationalAchAddressRequirement"
	this.Type = type_
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *USInternationalAchAddressRequirement) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *USInternationalAchAddressRequirement) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *USInternationalAchAddressRequirement) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *USInternationalAchAddressRequirement) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *USInternationalAchAddressRequirement) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *USInternationalAchAddressRequirement) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *USInternationalAchAddressRequirement) SetType(v string) {
	o.Type = v
}

func (o USInternationalAchAddressRequirement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o USInternationalAchAddressRequirement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableUSInternationalAchAddressRequirement struct {
	value *USInternationalAchAddressRequirement
	isSet bool
}

func (v NullableUSInternationalAchAddressRequirement) Get() *USInternationalAchAddressRequirement {
	return v.value
}

func (v *NullableUSInternationalAchAddressRequirement) Set(val *USInternationalAchAddressRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableUSInternationalAchAddressRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableUSInternationalAchAddressRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUSInternationalAchAddressRequirement(val *USInternationalAchAddressRequirement) *NullableUSInternationalAchAddressRequirement {
	return &NullableUSInternationalAchAddressRequirement{value: val, isSet: true}
}

func (v NullableUSInternationalAchAddressRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUSInternationalAchAddressRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *USInternationalAchAddressRequirement) isValidType() bool {
	var allowedEnumValues = []string{"usInternationalAchAddressRequirement"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
