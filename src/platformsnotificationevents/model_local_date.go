/*
 * Adyen for Platforms: Notifications
 *
 * The Notification API sends notifications to the endpoints specified in a given subscription. Subscriptions are managed through the Notification Configuration API. The API specifications listed here detail the format of each notification.  For more information, refer to our [documentation](https://docs.adyen.com/platforms/notifications).
 *
 * API version: 6
 * Contact: support@adyen.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package platformsnotificationevents

import (
	"encoding/json"
)

// LocalDate struct for LocalDate
type LocalDate struct {
	Month *int32 `json:"month,omitempty"`
	Year  *int32 `json:"year,omitempty"`
}

// NewLocalDate instantiates a new LocalDate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalDate() *LocalDate {
	this := LocalDate{}
	return &this
}

// NewLocalDateWithDefaults instantiates a new LocalDate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalDateWithDefaults() *LocalDate {
	this := LocalDate{}
	return &this
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *LocalDate) GetMonth() int32 {
	if o == nil || o.Month == nil {
		var ret int32
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDate) GetMonthOk() (*int32, bool) {
	if o == nil || o.Month == nil {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *LocalDate) HasMonth() bool {
	if o != nil && o.Month != nil {
		return true
	}

	return false
}

// SetMonth gets a reference to the given int32 and assigns it to the Month field.
func (o *LocalDate) SetMonth(v int32) {
	o.Month = &v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *LocalDate) GetYear() int32 {
	if o == nil || o.Year == nil {
		var ret int32
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalDate) GetYearOk() (*int32, bool) {
	if o == nil || o.Year == nil {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *LocalDate) HasYear() bool {
	if o != nil && o.Year != nil {
		return true
	}

	return false
}

// SetYear gets a reference to the given int32 and assigns it to the Year field.
func (o *LocalDate) SetYear(v int32) {
	o.Year = &v
}

func (o LocalDate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Month != nil {
		toSerialize["month"] = o.Month
	}
	if o.Year != nil {
		toSerialize["year"] = o.Year
	}
	return json.Marshal(toSerialize)
}

type NullableLocalDate struct {
	value *LocalDate
	isSet bool
}

func (v NullableLocalDate) Get() *LocalDate {
	return v.value
}

func (v *NullableLocalDate) Set(val *LocalDate) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalDate) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalDate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalDate(val *LocalDate) *NullableLocalDate {
	return &NullableLocalDate{value: val, isSet: true}
}

func (v NullableLocalDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalDate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
