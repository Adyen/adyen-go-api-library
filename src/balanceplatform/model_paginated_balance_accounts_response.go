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

// checks if the PaginatedBalanceAccountsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaginatedBalanceAccountsResponse{}

// PaginatedBalanceAccountsResponse struct for PaginatedBalanceAccountsResponse
type PaginatedBalanceAccountsResponse struct {
	// List of balance accounts.
	BalanceAccounts []BalanceAccountBase `json:"balanceAccounts"`
	// Indicates whether there are more items on the next page.
	HasNext bool `json:"hasNext"`
	// Indicates whether there are more items on the previous page.
	HasPrevious bool `json:"hasPrevious"`
}

// NewPaginatedBalanceAccountsResponse instantiates a new PaginatedBalanceAccountsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedBalanceAccountsResponse(balanceAccounts []BalanceAccountBase, hasNext bool, hasPrevious bool) *PaginatedBalanceAccountsResponse {
	this := PaginatedBalanceAccountsResponse{}
	this.BalanceAccounts = balanceAccounts
	this.HasNext = hasNext
	this.HasPrevious = hasPrevious
	return &this
}

// NewPaginatedBalanceAccountsResponseWithDefaults instantiates a new PaginatedBalanceAccountsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedBalanceAccountsResponseWithDefaults() *PaginatedBalanceAccountsResponse {
	this := PaginatedBalanceAccountsResponse{}
	return &this
}

// GetBalanceAccounts returns the BalanceAccounts field value
func (o *PaginatedBalanceAccountsResponse) GetBalanceAccounts() []BalanceAccountBase {
	if o == nil {
		var ret []BalanceAccountBase
		return ret
	}

	return o.BalanceAccounts
}

// GetBalanceAccountsOk returns a tuple with the BalanceAccounts field value
// and a boolean to check if the value has been set.
func (o *PaginatedBalanceAccountsResponse) GetBalanceAccountsOk() ([]BalanceAccountBase, bool) {
	if o == nil {
		return nil, false
	}
	return o.BalanceAccounts, true
}

// SetBalanceAccounts sets field value
func (o *PaginatedBalanceAccountsResponse) SetBalanceAccounts(v []BalanceAccountBase) {
	o.BalanceAccounts = v
}

// GetHasNext returns the HasNext field value
func (o *PaginatedBalanceAccountsResponse) GetHasNext() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasNext
}

// GetHasNextOk returns a tuple with the HasNext field value
// and a boolean to check if the value has been set.
func (o *PaginatedBalanceAccountsResponse) GetHasNextOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasNext, true
}

// SetHasNext sets field value
func (o *PaginatedBalanceAccountsResponse) SetHasNext(v bool) {
	o.HasNext = v
}

// GetHasPrevious returns the HasPrevious field value
func (o *PaginatedBalanceAccountsResponse) GetHasPrevious() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasPrevious
}

// GetHasPreviousOk returns a tuple with the HasPrevious field value
// and a boolean to check if the value has been set.
func (o *PaginatedBalanceAccountsResponse) GetHasPreviousOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasPrevious, true
}

// SetHasPrevious sets field value
func (o *PaginatedBalanceAccountsResponse) SetHasPrevious(v bool) {
	o.HasPrevious = v
}

func (o PaginatedBalanceAccountsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedBalanceAccountsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["balanceAccounts"] = o.BalanceAccounts
	toSerialize["hasNext"] = o.HasNext
	toSerialize["hasPrevious"] = o.HasPrevious
	return toSerialize, nil
}

type NullablePaginatedBalanceAccountsResponse struct {
	value *PaginatedBalanceAccountsResponse
	isSet bool
}

func (v NullablePaginatedBalanceAccountsResponse) Get() *PaginatedBalanceAccountsResponse {
	return v.value
}

func (v *NullablePaginatedBalanceAccountsResponse) Set(val *PaginatedBalanceAccountsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedBalanceAccountsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedBalanceAccountsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedBalanceAccountsResponse(val *PaginatedBalanceAccountsResponse) *NullablePaginatedBalanceAccountsResponse {
	return &NullablePaginatedBalanceAccountsResponse{value: val, isSet: true}
}

func (v NullablePaginatedBalanceAccountsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedBalanceAccountsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
