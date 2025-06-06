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

// checks if the AndroidAppsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AndroidAppsResponse{}

// AndroidAppsResponse struct for AndroidAppsResponse
type AndroidAppsResponse struct {
	// Apps uploaded for Android payment terminals.
	Data []AndroidApp `json:"data,omitempty"`
}

// NewAndroidAppsResponse instantiates a new AndroidAppsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAndroidAppsResponse() *AndroidAppsResponse {
	this := AndroidAppsResponse{}
	return &this
}

// NewAndroidAppsResponseWithDefaults instantiates a new AndroidAppsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAndroidAppsResponseWithDefaults() *AndroidAppsResponse {
	this := AndroidAppsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AndroidAppsResponse) GetData() []AndroidApp {
	if o == nil || common.IsNil(o.Data) {
		var ret []AndroidApp
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidAppsResponse) GetDataOk() ([]AndroidApp, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AndroidAppsResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AndroidApp and assigns it to the Data field.
func (o *AndroidAppsResponse) SetData(v []AndroidApp) {
	o.Data = v
}

func (o AndroidAppsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AndroidAppsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAndroidAppsResponse struct {
	value *AndroidAppsResponse
	isSet bool
}

func (v NullableAndroidAppsResponse) Get() *AndroidAppsResponse {
	return v.value
}

func (v *NullableAndroidAppsResponse) Set(val *AndroidAppsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAndroidAppsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAndroidAppsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAndroidAppsResponse(val *AndroidAppsResponse) *NullableAndroidAppsResponse {
	return &NullableAndroidAppsResponse{value: val, isSet: true}
}

func (v NullableAndroidAppsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAndroidAppsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
