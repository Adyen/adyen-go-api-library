/*
Transfers API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the CapitalGrantInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CapitalGrantInfo{}

// CapitalGrantInfo struct for CapitalGrantInfo
type CapitalGrantInfo struct {
	Counterparty *Counterparty2 `json:"counterparty,omitempty"`
	// The identifier of the grant account used for the grant.
	GrantAccountId string `json:"grantAccountId"`
	// The identifier of the grant offer that has been selected and from which the grant details will be used.
	GrantOfferId string `json:"grantOfferId"`
}

// NewCapitalGrantInfo instantiates a new CapitalGrantInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapitalGrantInfo(grantAccountId string, grantOfferId string) *CapitalGrantInfo {
	this := CapitalGrantInfo{}
	this.GrantAccountId = grantAccountId
	this.GrantOfferId = grantOfferId
	return &this
}

// NewCapitalGrantInfoWithDefaults instantiates a new CapitalGrantInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapitalGrantInfoWithDefaults() *CapitalGrantInfo {
	this := CapitalGrantInfo{}
	return &this
}

// GetCounterparty returns the Counterparty field value if set, zero value otherwise.
func (o *CapitalGrantInfo) GetCounterparty() Counterparty2 {
	if o == nil || common.IsNil(o.Counterparty) {
		var ret Counterparty2
		return ret
	}
	return *o.Counterparty
}

// GetCounterpartyOk returns a tuple with the Counterparty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapitalGrantInfo) GetCounterpartyOk() (*Counterparty2, bool) {
	if o == nil || common.IsNil(o.Counterparty) {
		return nil, false
	}
	return o.Counterparty, true
}

// HasCounterparty returns a boolean if a field has been set.
func (o *CapitalGrantInfo) HasCounterparty() bool {
	if o != nil && !common.IsNil(o.Counterparty) {
		return true
	}

	return false
}

// SetCounterparty gets a reference to the given Counterparty2 and assigns it to the Counterparty field.
func (o *CapitalGrantInfo) SetCounterparty(v Counterparty2) {
	o.Counterparty = &v
}

// GetGrantAccountId returns the GrantAccountId field value
func (o *CapitalGrantInfo) GetGrantAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GrantAccountId
}

// GetGrantAccountIdOk returns a tuple with the GrantAccountId field value
// and a boolean to check if the value has been set.
func (o *CapitalGrantInfo) GetGrantAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantAccountId, true
}

// SetGrantAccountId sets field value
func (o *CapitalGrantInfo) SetGrantAccountId(v string) {
	o.GrantAccountId = v
}

// GetGrantOfferId returns the GrantOfferId field value
func (o *CapitalGrantInfo) GetGrantOfferId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GrantOfferId
}

// GetGrantOfferIdOk returns a tuple with the GrantOfferId field value
// and a boolean to check if the value has been set.
func (o *CapitalGrantInfo) GetGrantOfferIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantOfferId, true
}

// SetGrantOfferId sets field value
func (o *CapitalGrantInfo) SetGrantOfferId(v string) {
	o.GrantOfferId = v
}

func (o CapitalGrantInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapitalGrantInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Counterparty) {
		toSerialize["counterparty"] = o.Counterparty
	}
	toSerialize["grantAccountId"] = o.GrantAccountId
	toSerialize["grantOfferId"] = o.GrantOfferId
	return toSerialize, nil
}

type NullableCapitalGrantInfo struct {
	value *CapitalGrantInfo
	isSet bool
}

func (v NullableCapitalGrantInfo) Get() *CapitalGrantInfo {
	return v.value
}

func (v *NullableCapitalGrantInfo) Set(val *CapitalGrantInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCapitalGrantInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCapitalGrantInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapitalGrantInfo(val *CapitalGrantInfo) *NullableCapitalGrantInfo {
	return &NullableCapitalGrantInfo{value: val, isSet: true}
}

func (v NullableCapitalGrantInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapitalGrantInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}