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

// checks if the ReprocessAndroidAppResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ReprocessAndroidAppResponse{}

// ReprocessAndroidAppResponse struct for ReprocessAndroidAppResponse
type ReprocessAndroidAppResponse struct {
	// The result of the reprocess.
	Message *string `json:"Message,omitempty"`
}

// NewReprocessAndroidAppResponse instantiates a new ReprocessAndroidAppResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReprocessAndroidAppResponse() *ReprocessAndroidAppResponse {
	this := ReprocessAndroidAppResponse{}
	return &this
}

// NewReprocessAndroidAppResponseWithDefaults instantiates a new ReprocessAndroidAppResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReprocessAndroidAppResponseWithDefaults() *ReprocessAndroidAppResponse {
	this := ReprocessAndroidAppResponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ReprocessAndroidAppResponse) GetMessage() string {
	if o == nil || common.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReprocessAndroidAppResponse) GetMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ReprocessAndroidAppResponse) HasMessage() bool {
	if o != nil && !common.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ReprocessAndroidAppResponse) SetMessage(v string) {
	o.Message = &v
}

func (o ReprocessAndroidAppResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReprocessAndroidAppResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Message) {
		toSerialize["Message"] = o.Message
	}
	return toSerialize, nil
}

type NullableReprocessAndroidAppResponse struct {
	value *ReprocessAndroidAppResponse
	isSet bool
}

func (v NullableReprocessAndroidAppResponse) Get() *ReprocessAndroidAppResponse {
	return v.value
}

func (v *NullableReprocessAndroidAppResponse) Set(val *ReprocessAndroidAppResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReprocessAndroidAppResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReprocessAndroidAppResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReprocessAndroidAppResponse(val *ReprocessAndroidAppResponse) *NullableReprocessAndroidAppResponse {
	return &NullableReprocessAndroidAppResponse{value: val, isSet: true}
}

func (v NullableReprocessAndroidAppResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReprocessAndroidAppResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
