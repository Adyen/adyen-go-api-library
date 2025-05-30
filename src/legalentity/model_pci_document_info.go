/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the PciDocumentInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PciDocumentInfo{}

// PciDocumentInfo struct for PciDocumentInfo
type PciDocumentInfo struct {
	// The date the questionnaire was created, in ISO 8601 extended format. For example, 2022-12-18T10:15:30+01:00
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The unique identifier of the signed questionnaire.
	Id *string `json:"id,omitempty"`
	// The expiration date of the questionnaire, in ISO 8601 extended format. For example, 2022-12-18T10:15:30+01:00
	ValidUntil *time.Time `json:"validUntil,omitempty"`
}

// NewPciDocumentInfo instantiates a new PciDocumentInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPciDocumentInfo() *PciDocumentInfo {
	this := PciDocumentInfo{}
	return &this
}

// NewPciDocumentInfoWithDefaults instantiates a new PciDocumentInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPciDocumentInfoWithDefaults() *PciDocumentInfo {
	this := PciDocumentInfo{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PciDocumentInfo) GetCreatedAt() time.Time {
	if o == nil || common.IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciDocumentInfo) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PciDocumentInfo) HasCreatedAt() bool {
	if o != nil && !common.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PciDocumentInfo) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PciDocumentInfo) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciDocumentInfo) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PciDocumentInfo) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PciDocumentInfo) SetId(v string) {
	o.Id = &v
}

// GetValidUntil returns the ValidUntil field value if set, zero value otherwise.
func (o *PciDocumentInfo) GetValidUntil() time.Time {
	if o == nil || common.IsNil(o.ValidUntil) {
		var ret time.Time
		return ret
	}
	return *o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciDocumentInfo) GetValidUntilOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.ValidUntil) {
		return nil, false
	}
	return o.ValidUntil, true
}

// HasValidUntil returns a boolean if a field has been set.
func (o *PciDocumentInfo) HasValidUntil() bool {
	if o != nil && !common.IsNil(o.ValidUntil) {
		return true
	}

	return false
}

// SetValidUntil gets a reference to the given time.Time and assigns it to the ValidUntil field.
func (o *PciDocumentInfo) SetValidUntil(v time.Time) {
	o.ValidUntil = &v
}

func (o PciDocumentInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PciDocumentInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.ValidUntil) {
		toSerialize["validUntil"] = o.ValidUntil
	}
	return toSerialize, nil
}

type NullablePciDocumentInfo struct {
	value *PciDocumentInfo
	isSet bool
}

func (v NullablePciDocumentInfo) Get() *PciDocumentInfo {
	return v.value
}

func (v *NullablePciDocumentInfo) Set(val *PciDocumentInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePciDocumentInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePciDocumentInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePciDocumentInfo(val *PciDocumentInfo) *NullablePciDocumentInfo {
	return &NullablePciDocumentInfo{value: val, isSet: true}
}

func (v NullablePciDocumentInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePciDocumentInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
