/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// checks if the StringMatch type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StringMatch{}

// StringMatch struct for StringMatch
type StringMatch struct {
	// The type of string matching operation. Possible values:  **startsWith**, **endsWith**, **isEqualTo**, **contains**,
	Operation *string `json:"operation,omitempty"`
	// The string to be matched.
	Value *string `json:"value,omitempty"`
}

// NewStringMatch instantiates a new StringMatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringMatch() *StringMatch {
	this := StringMatch{}
	return &this
}

// NewStringMatchWithDefaults instantiates a new StringMatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringMatchWithDefaults() *StringMatch {
	this := StringMatch{}
	return &this
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *StringMatch) GetOperation() string {
	if o == nil || common.IsNil(o.Operation) {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMatch) GetOperationOk() (*string, bool) {
	if o == nil || common.IsNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *StringMatch) HasOperation() bool {
	if o != nil && !common.IsNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *StringMatch) SetOperation(v string) {
	o.Operation = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *StringMatch) GetValue() string {
	if o == nil || common.IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMatch) GetValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *StringMatch) HasValue() bool {
	if o != nil && !common.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *StringMatch) SetValue(v string) {
	o.Value = &v
}

func (o StringMatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StringMatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if !common.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableStringMatch struct {
	value *StringMatch
	isSet bool
}

func (v NullableStringMatch) Get() *StringMatch {
	return v.value
}

func (v *NullableStringMatch) Set(val *StringMatch) {
	v.value = val
	v.isSet = true
}

func (v NullableStringMatch) IsSet() bool {
	return v.isSet
}

func (v *NullableStringMatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringMatch(val *StringMatch) *NullableStringMatch {
	return &NullableStringMatch{value: val, isSet: true}
}

func (v NullableStringMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringMatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *StringMatch) isValidOperation() bool {
	var allowedEnumValues = []string{"contains", "endsWith", "isEqualTo", "startsWith"}
	for _, allowed := range allowedEnumValues {
		if o.GetOperation() == allowed {
			return true
		}
	}
	return false
}