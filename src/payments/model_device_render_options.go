/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the DeviceRenderOptions type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DeviceRenderOptions{}

// DeviceRenderOptions struct for DeviceRenderOptions
type DeviceRenderOptions struct {
	// Supported SDK interface types. Allowed values: * native * html * both
	SdkInterface *string `json:"sdkInterface,omitempty"`
	// UI types supported for displaying specific challenges. Allowed values: * text * singleSelect * outOfBand * otherHtml * multiSelect
	SdkUiType []string `json:"sdkUiType,omitempty"`
}

// NewDeviceRenderOptions instantiates a new DeviceRenderOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceRenderOptions() *DeviceRenderOptions {
	this := DeviceRenderOptions{}
	var sdkInterface string = "both"
	this.SdkInterface = &sdkInterface
	return &this
}

// NewDeviceRenderOptionsWithDefaults instantiates a new DeviceRenderOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceRenderOptionsWithDefaults() *DeviceRenderOptions {
	this := DeviceRenderOptions{}
	var sdkInterface string = "both"
	this.SdkInterface = &sdkInterface
	return &this
}

// GetSdkInterface returns the SdkInterface field value if set, zero value otherwise.
func (o *DeviceRenderOptions) GetSdkInterface() string {
	if o == nil || common.IsNil(o.SdkInterface) {
		var ret string
		return ret
	}
	return *o.SdkInterface
}

// GetSdkInterfaceOk returns a tuple with the SdkInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRenderOptions) GetSdkInterfaceOk() (*string, bool) {
	if o == nil || common.IsNil(o.SdkInterface) {
		return nil, false
	}
	return o.SdkInterface, true
}

// HasSdkInterface returns a boolean if a field has been set.
func (o *DeviceRenderOptions) HasSdkInterface() bool {
	if o != nil && !common.IsNil(o.SdkInterface) {
		return true
	}

	return false
}

// SetSdkInterface gets a reference to the given string and assigns it to the SdkInterface field.
func (o *DeviceRenderOptions) SetSdkInterface(v string) {
	o.SdkInterface = &v
}

// GetSdkUiType returns the SdkUiType field value if set, zero value otherwise.
func (o *DeviceRenderOptions) GetSdkUiType() []string {
	if o == nil || common.IsNil(o.SdkUiType) {
		var ret []string
		return ret
	}
	return o.SdkUiType
}

// GetSdkUiTypeOk returns a tuple with the SdkUiType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRenderOptions) GetSdkUiTypeOk() ([]string, bool) {
	if o == nil || common.IsNil(o.SdkUiType) {
		return nil, false
	}
	return o.SdkUiType, true
}

// HasSdkUiType returns a boolean if a field has been set.
func (o *DeviceRenderOptions) HasSdkUiType() bool {
	if o != nil && !common.IsNil(o.SdkUiType) {
		return true
	}

	return false
}

// SetSdkUiType gets a reference to the given []string and assigns it to the SdkUiType field.
func (o *DeviceRenderOptions) SetSdkUiType(v []string) {
	o.SdkUiType = v
}

func (o DeviceRenderOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceRenderOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.SdkInterface) {
		toSerialize["sdkInterface"] = o.SdkInterface
	}
	if !common.IsNil(o.SdkUiType) {
		toSerialize["sdkUiType"] = o.SdkUiType
	}
	return toSerialize, nil
}

type NullableDeviceRenderOptions struct {
	value *DeviceRenderOptions
	isSet bool
}

func (v NullableDeviceRenderOptions) Get() *DeviceRenderOptions {
	return v.value
}

func (v *NullableDeviceRenderOptions) Set(val *DeviceRenderOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceRenderOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceRenderOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceRenderOptions(val *DeviceRenderOptions) *NullableDeviceRenderOptions {
	return &NullableDeviceRenderOptions{value: val, isSet: true}
}

func (v NullableDeviceRenderOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceRenderOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *DeviceRenderOptions) isValidSdkInterface() bool {
	var allowedEnumValues = []string{"native", "html", "both"}
	for _, allowed := range allowedEnumValues {
		if o.GetSdkInterface() == allowed {
			return true
		}
	}
	return false
}
