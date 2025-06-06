/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the AdditionalBankIdentification type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AdditionalBankIdentification{}

// AdditionalBankIdentification struct for AdditionalBankIdentification
type AdditionalBankIdentification struct {
	// The value of the additional bank identification.
	Code *string `json:"code,omitempty"`
	// The type of additional bank identification, depending on the country.  Possible values:   * **gbSortCode**: The 6-digit [UK sort code](https://en.wikipedia.org/wiki/Sort_code), without separators or spaces  * **usRoutingNumber**: The 9-digit [routing number](https://en.wikipedia.org/wiki/ABA_routing_transit_number), without separators or spaces.
	Type *string `json:"type,omitempty"`
}

// NewAdditionalBankIdentification instantiates a new AdditionalBankIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalBankIdentification() *AdditionalBankIdentification {
	this := AdditionalBankIdentification{}
	return &this
}

// NewAdditionalBankIdentificationWithDefaults instantiates a new AdditionalBankIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalBankIdentificationWithDefaults() *AdditionalBankIdentification {
	this := AdditionalBankIdentification{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AdditionalBankIdentification) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalBankIdentification) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AdditionalBankIdentification) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AdditionalBankIdentification) SetCode(v string) {
	o.Code = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AdditionalBankIdentification) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalBankIdentification) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AdditionalBankIdentification) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AdditionalBankIdentification) SetType(v string) {
	o.Type = &v
}

func (o AdditionalBankIdentification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalBankIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAdditionalBankIdentification struct {
	value *AdditionalBankIdentification
	isSet bool
}

func (v NullableAdditionalBankIdentification) Get() *AdditionalBankIdentification {
	return v.value
}

func (v *NullableAdditionalBankIdentification) Set(val *AdditionalBankIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalBankIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalBankIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalBankIdentification(val *AdditionalBankIdentification) *NullableAdditionalBankIdentification {
	return &NullableAdditionalBankIdentification{value: val, isSet: true}
}

func (v NullableAdditionalBankIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalBankIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *AdditionalBankIdentification) isValidType() bool {
	var allowedEnumValues = []string{"gbSortCode", "usRoutingNumber"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
