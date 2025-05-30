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

// checks if the PublicKeyResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PublicKeyResponse{}

// PublicKeyResponse struct for PublicKeyResponse
type PublicKeyResponse struct {
	// The public key you need for encrypting a symmetric session key.
	PublicKey string `json:"publicKey"`
	// The expiry date of the public key.
	PublicKeyExpiryDate string `json:"publicKeyExpiryDate"`
}

// NewPublicKeyResponse instantiates a new PublicKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicKeyResponse(publicKey string, publicKeyExpiryDate string) *PublicKeyResponse {
	this := PublicKeyResponse{}
	this.PublicKey = publicKey
	this.PublicKeyExpiryDate = publicKeyExpiryDate
	return &this
}

// NewPublicKeyResponseWithDefaults instantiates a new PublicKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicKeyResponseWithDefaults() *PublicKeyResponse {
	this := PublicKeyResponse{}
	return &this
}

// GetPublicKey returns the PublicKey field value
func (o *PublicKeyResponse) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *PublicKeyResponse) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *PublicKeyResponse) SetPublicKey(v string) {
	o.PublicKey = v
}

// GetPublicKeyExpiryDate returns the PublicKeyExpiryDate field value
func (o *PublicKeyResponse) GetPublicKeyExpiryDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKeyExpiryDate
}

// GetPublicKeyExpiryDateOk returns a tuple with the PublicKeyExpiryDate field value
// and a boolean to check if the value has been set.
func (o *PublicKeyResponse) GetPublicKeyExpiryDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKeyExpiryDate, true
}

// SetPublicKeyExpiryDate sets field value
func (o *PublicKeyResponse) SetPublicKeyExpiryDate(v string) {
	o.PublicKeyExpiryDate = v
}

func (o PublicKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicKeyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["publicKey"] = o.PublicKey
	toSerialize["publicKeyExpiryDate"] = o.PublicKeyExpiryDate
	return toSerialize, nil
}

type NullablePublicKeyResponse struct {
	value *PublicKeyResponse
	isSet bool
}

func (v NullablePublicKeyResponse) Get() *PublicKeyResponse {
	return v.value
}

func (v *NullablePublicKeyResponse) Set(val *PublicKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicKeyResponse(val *PublicKeyResponse) *NullablePublicKeyResponse {
	return &NullablePublicKeyResponse{value: val, isSet: true}
}

func (v NullablePublicKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
