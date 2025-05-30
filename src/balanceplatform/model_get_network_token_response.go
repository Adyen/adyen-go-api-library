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

// checks if the GetNetworkTokenResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetNetworkTokenResponse{}

// GetNetworkTokenResponse struct for GetNetworkTokenResponse
type GetNetworkTokenResponse struct {
	Token NetworkToken `json:"token"`
}

// NewGetNetworkTokenResponse instantiates a new GetNetworkTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkTokenResponse(token NetworkToken) *GetNetworkTokenResponse {
	this := GetNetworkTokenResponse{}
	this.Token = token
	return &this
}

// NewGetNetworkTokenResponseWithDefaults instantiates a new GetNetworkTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkTokenResponseWithDefaults() *GetNetworkTokenResponse {
	this := GetNetworkTokenResponse{}
	return &this
}

// GetToken returns the Token field value
func (o *GetNetworkTokenResponse) GetToken() NetworkToken {
	if o == nil {
		var ret NetworkToken
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *GetNetworkTokenResponse) GetTokenOk() (*NetworkToken, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *GetNetworkTokenResponse) SetToken(v NetworkToken) {
	o.Token = v
}

func (o GetNetworkTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

type NullableGetNetworkTokenResponse struct {
	value *GetNetworkTokenResponse
	isSet bool
}

func (v NullableGetNetworkTokenResponse) Get() *GetNetworkTokenResponse {
	return v.value
}

func (v *NullableGetNetworkTokenResponse) Set(val *GetNetworkTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkTokenResponse(val *GetNetworkTokenResponse) *NullableGetNetworkTokenResponse {
	return &NullableGetNetworkTokenResponse{value: val, isSet: true}
}

func (v NullableGetNetworkTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
