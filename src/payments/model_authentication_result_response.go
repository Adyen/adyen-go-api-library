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

// checks if the AuthenticationResultResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AuthenticationResultResponse{}

// AuthenticationResultResponse struct for AuthenticationResultResponse
type AuthenticationResultResponse struct {
	ThreeDS1Result *ThreeDS1Result `json:"threeDS1Result,omitempty"`
	ThreeDS2Result *ThreeDS2Result `json:"threeDS2Result,omitempty"`
}

// NewAuthenticationResultResponse instantiates a new AuthenticationResultResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationResultResponse() *AuthenticationResultResponse {
	this := AuthenticationResultResponse{}
	return &this
}

// NewAuthenticationResultResponseWithDefaults instantiates a new AuthenticationResultResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationResultResponseWithDefaults() *AuthenticationResultResponse {
	this := AuthenticationResultResponse{}
	return &this
}

// GetThreeDS1Result returns the ThreeDS1Result field value if set, zero value otherwise.
func (o *AuthenticationResultResponse) GetThreeDS1Result() ThreeDS1Result {
	if o == nil || common.IsNil(o.ThreeDS1Result) {
		var ret ThreeDS1Result
		return ret
	}
	return *o.ThreeDS1Result
}

// GetThreeDS1ResultOk returns a tuple with the ThreeDS1Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationResultResponse) GetThreeDS1ResultOk() (*ThreeDS1Result, bool) {
	if o == nil || common.IsNil(o.ThreeDS1Result) {
		return nil, false
	}
	return o.ThreeDS1Result, true
}

// HasThreeDS1Result returns a boolean if a field has been set.
func (o *AuthenticationResultResponse) HasThreeDS1Result() bool {
	if o != nil && !common.IsNil(o.ThreeDS1Result) {
		return true
	}

	return false
}

// SetThreeDS1Result gets a reference to the given ThreeDS1Result and assigns it to the ThreeDS1Result field.
func (o *AuthenticationResultResponse) SetThreeDS1Result(v ThreeDS1Result) {
	o.ThreeDS1Result = &v
}

// GetThreeDS2Result returns the ThreeDS2Result field value if set, zero value otherwise.
func (o *AuthenticationResultResponse) GetThreeDS2Result() ThreeDS2Result {
	if o == nil || common.IsNil(o.ThreeDS2Result) {
		var ret ThreeDS2Result
		return ret
	}
	return *o.ThreeDS2Result
}

// GetThreeDS2ResultOk returns a tuple with the ThreeDS2Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationResultResponse) GetThreeDS2ResultOk() (*ThreeDS2Result, bool) {
	if o == nil || common.IsNil(o.ThreeDS2Result) {
		return nil, false
	}
	return o.ThreeDS2Result, true
}

// HasThreeDS2Result returns a boolean if a field has been set.
func (o *AuthenticationResultResponse) HasThreeDS2Result() bool {
	if o != nil && !common.IsNil(o.ThreeDS2Result) {
		return true
	}

	return false
}

// SetThreeDS2Result gets a reference to the given ThreeDS2Result and assigns it to the ThreeDS2Result field.
func (o *AuthenticationResultResponse) SetThreeDS2Result(v ThreeDS2Result) {
	o.ThreeDS2Result = &v
}

func (o AuthenticationResultResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationResultResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ThreeDS1Result) {
		toSerialize["threeDS1Result"] = o.ThreeDS1Result
	}
	if !common.IsNil(o.ThreeDS2Result) {
		toSerialize["threeDS2Result"] = o.ThreeDS2Result
	}
	return toSerialize, nil
}

type NullableAuthenticationResultResponse struct {
	value *AuthenticationResultResponse
	isSet bool
}

func (v NullableAuthenticationResultResponse) Get() *AuthenticationResultResponse {
	return v.value
}

func (v *NullableAuthenticationResultResponse) Set(val *AuthenticationResultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationResultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationResultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationResultResponse(val *AuthenticationResultResponse) *NullableAuthenticationResultResponse {
	return &NullableAuthenticationResultResponse{value: val, isSet: true}
}

func (v NullableAuthenticationResultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationResultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
