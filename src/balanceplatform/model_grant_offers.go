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

// checks if the GrantOffers type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GrantOffers{}

// GrantOffers struct for GrantOffers
type GrantOffers struct {
	// A list of available grant offers.
	GrantOffers []GrantOffer `json:"grantOffers"`
}

// NewGrantOffers instantiates a new GrantOffers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantOffers(grantOffers []GrantOffer) *GrantOffers {
	this := GrantOffers{}
	this.GrantOffers = grantOffers
	return &this
}

// NewGrantOffersWithDefaults instantiates a new GrantOffers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantOffersWithDefaults() *GrantOffers {
	this := GrantOffers{}
	return &this
}

// GetGrantOffers returns the GrantOffers field value
func (o *GrantOffers) GetGrantOffers() []GrantOffer {
	if o == nil {
		var ret []GrantOffer
		return ret
	}

	return o.GrantOffers
}

// GetGrantOffersOk returns a tuple with the GrantOffers field value
// and a boolean to check if the value has been set.
func (o *GrantOffers) GetGrantOffersOk() ([]GrantOffer, bool) {
	if o == nil {
		return nil, false
	}
	return o.GrantOffers, true
}

// SetGrantOffers sets field value
func (o *GrantOffers) SetGrantOffers(v []GrantOffer) {
	o.GrantOffers = v
}

func (o GrantOffers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrantOffers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["grantOffers"] = o.GrantOffers
	return toSerialize, nil
}

type NullableGrantOffers struct {
	value *GrantOffers
	isSet bool
}

func (v NullableGrantOffers) Get() *GrantOffers {
	return v.value
}

func (v *NullableGrantOffers) Set(val *GrantOffers) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantOffers) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantOffers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantOffers(val *GrantOffers) *NullableGrantOffers {
	return &NullableGrantOffers{value: val, isSet: true}
}

func (v NullableGrantOffers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantOffers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
