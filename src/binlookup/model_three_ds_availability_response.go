/*
Adyen BinLookup API

API version: 54
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package binlookup

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the ThreeDSAvailabilityResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ThreeDSAvailabilityResponse{}

// ThreeDSAvailabilityResponse struct for ThreeDSAvailabilityResponse
type ThreeDSAvailabilityResponse struct {
	BinDetails *BinDetail `json:"binDetails,omitempty"`
	// List of Directory Server (DS) public keys.
	DsPublicKeys []DSPublicKeyDetail `json:"dsPublicKeys,omitempty"`
	// Indicator if 3D Secure 1 is supported.
	ThreeDS1Supported *bool `json:"threeDS1Supported,omitempty"`
	// List of brand and card range pairs.
	ThreeDS2CardRangeDetails []ThreeDS2CardRangeDetail `json:"threeDS2CardRangeDetails,omitempty"`
	// Indicator if 3D Secure 2 is supported.
	ThreeDS2supported *bool `json:"threeDS2supported,omitempty"`
}

// NewThreeDSAvailabilityResponse instantiates a new ThreeDSAvailabilityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreeDSAvailabilityResponse() *ThreeDSAvailabilityResponse {
	this := ThreeDSAvailabilityResponse{}
	return &this
}

// NewThreeDSAvailabilityResponseWithDefaults instantiates a new ThreeDSAvailabilityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreeDSAvailabilityResponseWithDefaults() *ThreeDSAvailabilityResponse {
	this := ThreeDSAvailabilityResponse{}
	return &this
}

// GetBinDetails returns the BinDetails field value if set, zero value otherwise.
func (o *ThreeDSAvailabilityResponse) GetBinDetails() BinDetail {
	if o == nil || common.IsNil(o.BinDetails) {
		var ret BinDetail
		return ret
	}
	return *o.BinDetails
}

// GetBinDetailsOk returns a tuple with the BinDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSAvailabilityResponse) GetBinDetailsOk() (*BinDetail, bool) {
	if o == nil || common.IsNil(o.BinDetails) {
		return nil, false
	}
	return o.BinDetails, true
}

// HasBinDetails returns a boolean if a field has been set.
func (o *ThreeDSAvailabilityResponse) HasBinDetails() bool {
	if o != nil && !common.IsNil(o.BinDetails) {
		return true
	}

	return false
}

// SetBinDetails gets a reference to the given BinDetail and assigns it to the BinDetails field.
func (o *ThreeDSAvailabilityResponse) SetBinDetails(v BinDetail) {
	o.BinDetails = &v
}

// GetDsPublicKeys returns the DsPublicKeys field value if set, zero value otherwise.
func (o *ThreeDSAvailabilityResponse) GetDsPublicKeys() []DSPublicKeyDetail {
	if o == nil || common.IsNil(o.DsPublicKeys) {
		var ret []DSPublicKeyDetail
		return ret
	}
	return o.DsPublicKeys
}

// GetDsPublicKeysOk returns a tuple with the DsPublicKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSAvailabilityResponse) GetDsPublicKeysOk() ([]DSPublicKeyDetail, bool) {
	if o == nil || common.IsNil(o.DsPublicKeys) {
		return nil, false
	}
	return o.DsPublicKeys, true
}

// HasDsPublicKeys returns a boolean if a field has been set.
func (o *ThreeDSAvailabilityResponse) HasDsPublicKeys() bool {
	if o != nil && !common.IsNil(o.DsPublicKeys) {
		return true
	}

	return false
}

// SetDsPublicKeys gets a reference to the given []DSPublicKeyDetail and assigns it to the DsPublicKeys field.
func (o *ThreeDSAvailabilityResponse) SetDsPublicKeys(v []DSPublicKeyDetail) {
	o.DsPublicKeys = v
}

// GetThreeDS1Supported returns the ThreeDS1Supported field value if set, zero value otherwise.
func (o *ThreeDSAvailabilityResponse) GetThreeDS1Supported() bool {
	if o == nil || common.IsNil(o.ThreeDS1Supported) {
		var ret bool
		return ret
	}
	return *o.ThreeDS1Supported
}

// GetThreeDS1SupportedOk returns a tuple with the ThreeDS1Supported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSAvailabilityResponse) GetThreeDS1SupportedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ThreeDS1Supported) {
		return nil, false
	}
	return o.ThreeDS1Supported, true
}

// HasThreeDS1Supported returns a boolean if a field has been set.
func (o *ThreeDSAvailabilityResponse) HasThreeDS1Supported() bool {
	if o != nil && !common.IsNil(o.ThreeDS1Supported) {
		return true
	}

	return false
}

// SetThreeDS1Supported gets a reference to the given bool and assigns it to the ThreeDS1Supported field.
func (o *ThreeDSAvailabilityResponse) SetThreeDS1Supported(v bool) {
	o.ThreeDS1Supported = &v
}

// GetThreeDS2CardRangeDetails returns the ThreeDS2CardRangeDetails field value if set, zero value otherwise.
func (o *ThreeDSAvailabilityResponse) GetThreeDS2CardRangeDetails() []ThreeDS2CardRangeDetail {
	if o == nil || common.IsNil(o.ThreeDS2CardRangeDetails) {
		var ret []ThreeDS2CardRangeDetail
		return ret
	}
	return o.ThreeDS2CardRangeDetails
}

// GetThreeDS2CardRangeDetailsOk returns a tuple with the ThreeDS2CardRangeDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSAvailabilityResponse) GetThreeDS2CardRangeDetailsOk() ([]ThreeDS2CardRangeDetail, bool) {
	if o == nil || common.IsNil(o.ThreeDS2CardRangeDetails) {
		return nil, false
	}
	return o.ThreeDS2CardRangeDetails, true
}

// HasThreeDS2CardRangeDetails returns a boolean if a field has been set.
func (o *ThreeDSAvailabilityResponse) HasThreeDS2CardRangeDetails() bool {
	if o != nil && !common.IsNil(o.ThreeDS2CardRangeDetails) {
		return true
	}

	return false
}

// SetThreeDS2CardRangeDetails gets a reference to the given []ThreeDS2CardRangeDetail and assigns it to the ThreeDS2CardRangeDetails field.
func (o *ThreeDSAvailabilityResponse) SetThreeDS2CardRangeDetails(v []ThreeDS2CardRangeDetail) {
	o.ThreeDS2CardRangeDetails = v
}

// GetThreeDS2supported returns the ThreeDS2supported field value if set, zero value otherwise.
func (o *ThreeDSAvailabilityResponse) GetThreeDS2supported() bool {
	if o == nil || common.IsNil(o.ThreeDS2supported) {
		var ret bool
		return ret
	}
	return *o.ThreeDS2supported
}

// GetThreeDS2supportedOk returns a tuple with the ThreeDS2supported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSAvailabilityResponse) GetThreeDS2supportedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ThreeDS2supported) {
		return nil, false
	}
	return o.ThreeDS2supported, true
}

// HasThreeDS2supported returns a boolean if a field has been set.
func (o *ThreeDSAvailabilityResponse) HasThreeDS2supported() bool {
	if o != nil && !common.IsNil(o.ThreeDS2supported) {
		return true
	}

	return false
}

// SetThreeDS2supported gets a reference to the given bool and assigns it to the ThreeDS2supported field.
func (o *ThreeDSAvailabilityResponse) SetThreeDS2supported(v bool) {
	o.ThreeDS2supported = &v
}

func (o ThreeDSAvailabilityResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreeDSAvailabilityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BinDetails) {
		toSerialize["binDetails"] = o.BinDetails
	}
	if !common.IsNil(o.DsPublicKeys) {
		toSerialize["dsPublicKeys"] = o.DsPublicKeys
	}
	if !common.IsNil(o.ThreeDS1Supported) {
		toSerialize["threeDS1Supported"] = o.ThreeDS1Supported
	}
	if !common.IsNil(o.ThreeDS2CardRangeDetails) {
		toSerialize["threeDS2CardRangeDetails"] = o.ThreeDS2CardRangeDetails
	}
	if !common.IsNil(o.ThreeDS2supported) {
		toSerialize["threeDS2supported"] = o.ThreeDS2supported
	}
	return toSerialize, nil
}

type NullableThreeDSAvailabilityResponse struct {
	value *ThreeDSAvailabilityResponse
	isSet bool
}

func (v NullableThreeDSAvailabilityResponse) Get() *ThreeDSAvailabilityResponse {
	return v.value
}

func (v *NullableThreeDSAvailabilityResponse) Set(val *ThreeDSAvailabilityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThreeDSAvailabilityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThreeDSAvailabilityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreeDSAvailabilityResponse(val *ThreeDSAvailabilityResponse) *NullableThreeDSAvailabilityResponse {
	return &NullableThreeDSAvailabilityResponse{value: val, isSet: true}
}

func (v NullableThreeDSAvailabilityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreeDSAvailabilityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
