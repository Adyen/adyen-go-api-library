/*
Adyen Recurring API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package recurring

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the DisableResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DisableResult{}

// DisableResult struct for DisableResult
type DisableResult struct {
	// Depending on whether a specific recurring detail was in the request, result is either [detail-successfully-disabled] or [all-details-successfully-disabled].
	Response *string `json:"response,omitempty"`
}

// NewDisableResult instantiates a new DisableResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisableResult() *DisableResult {
	this := DisableResult{}
	return &this
}

// NewDisableResultWithDefaults instantiates a new DisableResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisableResultWithDefaults() *DisableResult {
	this := DisableResult{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DisableResult) GetResponse() string {
	if o == nil || common.IsNil(o.Response) {
		var ret string
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisableResult) GetResponseOk() (*string, bool) {
	if o == nil || common.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DisableResult) HasResponse() bool {
	if o != nil && !common.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given string and assigns it to the Response field.
func (o *DisableResult) SetResponse(v string) {
	o.Response = &v
}

func (o DisableResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisableResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Response) {
		toSerialize["response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDisableResult struct {
	value *DisableResult
	isSet bool
}

func (v NullableDisableResult) Get() *DisableResult {
	return v.value
}

func (v *NullableDisableResult) Set(val *DisableResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDisableResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDisableResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisableResult(val *DisableResult) *NullableDisableResult {
	return &NullableDisableResult{value: val, isSet: true}
}

func (v NullableDisableResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisableResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
