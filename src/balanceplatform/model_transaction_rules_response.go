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

// checks if the TransactionRulesResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransactionRulesResponse{}

// TransactionRulesResponse struct for TransactionRulesResponse
type TransactionRulesResponse struct {
	// List of transaction rules.
	TransactionRules []TransactionRule `json:"transactionRules,omitempty"`
}

// NewTransactionRulesResponse instantiates a new TransactionRulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRulesResponse() *TransactionRulesResponse {
	this := TransactionRulesResponse{}
	return &this
}

// NewTransactionRulesResponseWithDefaults instantiates a new TransactionRulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRulesResponseWithDefaults() *TransactionRulesResponse {
	this := TransactionRulesResponse{}
	return &this
}

// GetTransactionRules returns the TransactionRules field value if set, zero value otherwise.
func (o *TransactionRulesResponse) GetTransactionRules() []TransactionRule {
	if o == nil || common.IsNil(o.TransactionRules) {
		var ret []TransactionRule
		return ret
	}
	return o.TransactionRules
}

// GetTransactionRulesOk returns a tuple with the TransactionRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRulesResponse) GetTransactionRulesOk() ([]TransactionRule, bool) {
	if o == nil || common.IsNil(o.TransactionRules) {
		return nil, false
	}
	return o.TransactionRules, true
}

// HasTransactionRules returns a boolean if a field has been set.
func (o *TransactionRulesResponse) HasTransactionRules() bool {
	if o != nil && !common.IsNil(o.TransactionRules) {
		return true
	}

	return false
}

// SetTransactionRules gets a reference to the given []TransactionRule and assigns it to the TransactionRules field.
func (o *TransactionRulesResponse) SetTransactionRules(v []TransactionRule) {
	o.TransactionRules = v
}

func (o TransactionRulesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionRulesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TransactionRules) {
		toSerialize["transactionRules"] = o.TransactionRules
	}
	return toSerialize, nil
}

type NullableTransactionRulesResponse struct {
	value *TransactionRulesResponse
	isSet bool
}

func (v NullableTransactionRulesResponse) Get() *TransactionRulesResponse {
	return v.value
}

func (v *NullableTransactionRulesResponse) Set(val *TransactionRulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRulesResponse(val *TransactionRulesResponse) *NullableTransactionRulesResponse {
	return &NullableTransactionRulesResponse{value: val, isSet: true}
}

func (v NullableTransactionRulesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}