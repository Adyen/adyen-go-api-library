/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the OnboardingLink type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OnboardingLink{}

// OnboardingLink struct for OnboardingLink
type OnboardingLink struct {
	// The URL of the hosted onboarding page where you need to redirect your user. This URL expires after 4 minutes and can only be used once.  If the link expires, you need to create a new link.
	Url *string `json:"url,omitempty"`
}

// NewOnboardingLink instantiates a new OnboardingLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnboardingLink() *OnboardingLink {
	this := OnboardingLink{}
	return &this
}

// NewOnboardingLinkWithDefaults instantiates a new OnboardingLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnboardingLinkWithDefaults() *OnboardingLink {
	this := OnboardingLink{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *OnboardingLink) GetUrl() string {
	if o == nil || common.IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingLink) GetUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *OnboardingLink) HasUrl() bool {
	if o != nil && !common.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *OnboardingLink) SetUrl(v string) {
	o.Url = &v
}

func (o OnboardingLink) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnboardingLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableOnboardingLink struct {
	value *OnboardingLink
	isSet bool
}

func (v NullableOnboardingLink) Get() *OnboardingLink {
	return v.value
}

func (v *NullableOnboardingLink) Set(val *OnboardingLink) {
	v.value = val
	v.isSet = true
}

func (v NullableOnboardingLink) IsSet() bool {
	return v.isSet
}

func (v *NullableOnboardingLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnboardingLink(val *OnboardingLink) *NullableOnboardingLink {
	return &NullableOnboardingLink{value: val, isSet: true}
}

func (v NullableOnboardingLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnboardingLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
