/*
Adyen Checkout API

API version: 70
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the UtilityRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UtilityRequest{}

// UtilityRequest struct for UtilityRequest
type UtilityRequest struct {
	// The list of origin domains, for which origin keys are requested.
	OriginDomains []string `json:"originDomains"`
}

// NewUtilityRequest instantiates a new UtilityRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUtilityRequest(originDomains []string) *UtilityRequest {
	this := UtilityRequest{}
	this.OriginDomains = originDomains
	return &this
}

// NewUtilityRequestWithDefaults instantiates a new UtilityRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUtilityRequestWithDefaults() *UtilityRequest {
	this := UtilityRequest{}
	return &this
}

// GetOriginDomains returns the OriginDomains field value
func (o *UtilityRequest) GetOriginDomains() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OriginDomains
}

// GetOriginDomainsOk returns a tuple with the OriginDomains field value
// and a boolean to check if the value has been set.
func (o *UtilityRequest) GetOriginDomainsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginDomains, true
}

// SetOriginDomains sets field value
func (o *UtilityRequest) SetOriginDomains(v []string) {
	o.OriginDomains = v
}

func (o UtilityRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UtilityRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["originDomains"] = o.OriginDomains
	return toSerialize, nil
}

type NullableUtilityRequest struct {
	value *UtilityRequest
	isSet bool
}

func (v NullableUtilityRequest) Get() *UtilityRequest {
	return v.value
}

func (v *NullableUtilityRequest) Set(val *UtilityRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUtilityRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUtilityRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUtilityRequest(val *UtilityRequest) *NullableUtilityRequest {
	return &NullableUtilityRequest{value: val, isSet: true}
}

func (v NullableUtilityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUtilityRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}