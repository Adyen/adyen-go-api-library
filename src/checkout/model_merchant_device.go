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

// checks if the MerchantDevice type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MerchantDevice{}

// MerchantDevice struct for MerchantDevice
type MerchantDevice struct {
	// Operating system running on the merchant device.
	Os *string `json:"os,omitempty"`
	// Version of the operating system on the merchant device.
	OsVersion *string `json:"osVersion,omitempty"`
	// Merchant device reference.
	Reference *string `json:"reference,omitempty"`
}

// NewMerchantDevice instantiates a new MerchantDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantDevice() *MerchantDevice {
	this := MerchantDevice{}
	return &this
}

// NewMerchantDeviceWithDefaults instantiates a new MerchantDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantDeviceWithDefaults() *MerchantDevice {
	this := MerchantDevice{}
	return &this
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *MerchantDevice) GetOs() string {
	if o == nil || common.IsNil(o.Os) {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantDevice) GetOsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Os) {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *MerchantDevice) HasOs() bool {
	if o != nil && !common.IsNil(o.Os) {
		return true
	}

	return false
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *MerchantDevice) SetOs(v string) {
	o.Os = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *MerchantDevice) GetOsVersion() string {
	if o == nil || common.IsNil(o.OsVersion) {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantDevice) GetOsVersionOk() (*string, bool) {
	if o == nil || common.IsNil(o.OsVersion) {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *MerchantDevice) HasOsVersion() bool {
	if o != nil && !common.IsNil(o.OsVersion) {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *MerchantDevice) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *MerchantDevice) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantDevice) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *MerchantDevice) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *MerchantDevice) SetReference(v string) {
	o.Reference = &v
}

func (o MerchantDevice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Os) {
		toSerialize["os"] = o.Os
	}
	if !common.IsNil(o.OsVersion) {
		toSerialize["osVersion"] = o.OsVersion
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	return toSerialize, nil
}

type NullableMerchantDevice struct {
	value *MerchantDevice
	isSet bool
}

func (v NullableMerchantDevice) Get() *MerchantDevice {
	return v.value
}

func (v *NullableMerchantDevice) Set(val *MerchantDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantDevice(val *MerchantDevice) *NullableMerchantDevice {
	return &NullableMerchantDevice{value: val, isSet: true}
}

func (v NullableMerchantDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
