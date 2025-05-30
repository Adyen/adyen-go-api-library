/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the TerminalInstructions type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TerminalInstructions{}

// TerminalInstructions struct for TerminalInstructions
type TerminalInstructions struct {
	// Indicates whether the Adyen app on the payment terminal restarts automatically when the configuration is updated.
	AdyenAppRestart *bool `json:"adyenAppRestart,omitempty"`
}

// NewTerminalInstructions instantiates a new TerminalInstructions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalInstructions() *TerminalInstructions {
	this := TerminalInstructions{}
	return &this
}

// NewTerminalInstructionsWithDefaults instantiates a new TerminalInstructions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalInstructionsWithDefaults() *TerminalInstructions {
	this := TerminalInstructions{}
	return &this
}

// GetAdyenAppRestart returns the AdyenAppRestart field value if set, zero value otherwise.
func (o *TerminalInstructions) GetAdyenAppRestart() bool {
	if o == nil || common.IsNil(o.AdyenAppRestart) {
		var ret bool
		return ret
	}
	return *o.AdyenAppRestart
}

// GetAdyenAppRestartOk returns a tuple with the AdyenAppRestart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalInstructions) GetAdyenAppRestartOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AdyenAppRestart) {
		return nil, false
	}
	return o.AdyenAppRestart, true
}

// HasAdyenAppRestart returns a boolean if a field has been set.
func (o *TerminalInstructions) HasAdyenAppRestart() bool {
	if o != nil && !common.IsNil(o.AdyenAppRestart) {
		return true
	}

	return false
}

// SetAdyenAppRestart gets a reference to the given bool and assigns it to the AdyenAppRestart field.
func (o *TerminalInstructions) SetAdyenAppRestart(v bool) {
	o.AdyenAppRestart = &v
}

func (o TerminalInstructions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalInstructions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AdyenAppRestart) {
		toSerialize["adyenAppRestart"] = o.AdyenAppRestart
	}
	return toSerialize, nil
}

type NullableTerminalInstructions struct {
	value *TerminalInstructions
	isSet bool
}

func (v NullableTerminalInstructions) Get() *TerminalInstructions {
	return v.value
}

func (v *NullableTerminalInstructions) Set(val *TerminalInstructions) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalInstructions) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalInstructions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalInstructions(val *TerminalInstructions) *NullableTerminalInstructions {
	return &NullableTerminalInstructions{value: val, isSet: true}
}

func (v NullableTerminalInstructions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalInstructions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
