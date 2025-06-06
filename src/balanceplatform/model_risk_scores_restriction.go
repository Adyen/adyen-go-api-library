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

// checks if the RiskScoresRestriction type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RiskScoresRestriction{}

// RiskScoresRestriction struct for RiskScoresRestriction
type RiskScoresRestriction struct {
	// Defines how the condition must be evaluated.
	Operation string      `json:"operation"`
	Value     *RiskScores `json:"value,omitempty"`
}

// NewRiskScoresRestriction instantiates a new RiskScoresRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskScoresRestriction(operation string) *RiskScoresRestriction {
	this := RiskScoresRestriction{}
	this.Operation = operation
	return &this
}

// NewRiskScoresRestrictionWithDefaults instantiates a new RiskScoresRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskScoresRestrictionWithDefaults() *RiskScoresRestriction {
	this := RiskScoresRestriction{}
	return &this
}

// GetOperation returns the Operation field value
func (o *RiskScoresRestriction) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *RiskScoresRestriction) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *RiskScoresRestriction) SetOperation(v string) {
	o.Operation = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RiskScoresRestriction) GetValue() RiskScores {
	if o == nil || common.IsNil(o.Value) {
		var ret RiskScores
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskScoresRestriction) GetValueOk() (*RiskScores, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RiskScoresRestriction) HasValue() bool {
	if o != nil && !common.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given RiskScores and assigns it to the Value field.
func (o *RiskScoresRestriction) SetValue(v RiskScores) {
	o.Value = &v
}

func (o RiskScoresRestriction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskScoresRestriction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operation"] = o.Operation
	if !common.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableRiskScoresRestriction struct {
	value *RiskScoresRestriction
	isSet bool
}

func (v NullableRiskScoresRestriction) Get() *RiskScoresRestriction {
	return v.value
}

func (v *NullableRiskScoresRestriction) Set(val *RiskScoresRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskScoresRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskScoresRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskScoresRestriction(val *RiskScoresRestriction) *NullableRiskScoresRestriction {
	return &NullableRiskScoresRestriction{value: val, isSet: true}
}

func (v NullableRiskScoresRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskScoresRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
