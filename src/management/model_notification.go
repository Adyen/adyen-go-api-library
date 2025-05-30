/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the Notification type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Notification{}

// Notification struct for Notification
type Notification struct {
	// The type of event notification sent when you select the notification button.
	Category *string `json:"category,omitempty"`
	// The text shown in the prompt which opens when you select the notification button. For example, the description of the input box for pay-at-table.
	Details *string `json:"details,omitempty"`
	// Enables sending event notifications either by pressing the Confirm key on terminals with a keypad or by tapping the event notification button on the terminal screen.
	Enabled *bool `json:"enabled,omitempty"`
	// Shows or hides the event notification button on the screen of terminal models that have a keypad.
	ShowButton *bool `json:"showButton,omitempty"`
	// The name of the notification button on the terminal screen.
	Title *string `json:"title,omitempty"`
}

// NewNotification instantiates a new Notification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotification() *Notification {
	this := Notification{}
	return &this
}

// NewNotificationWithDefaults instantiates a new Notification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationWithDefaults() *Notification {
	this := Notification{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *Notification) GetCategory() string {
	if o == nil || common.IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notification) GetCategoryOk() (*string, bool) {
	if o == nil || common.IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *Notification) HasCategory() bool {
	if o != nil && !common.IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *Notification) SetCategory(v string) {
	o.Category = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *Notification) GetDetails() string {
	if o == nil || common.IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notification) GetDetailsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *Notification) HasDetails() bool {
	if o != nil && !common.IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *Notification) SetDetails(v string) {
	o.Details = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Notification) GetEnabled() bool {
	if o == nil || common.IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notification) GetEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Notification) HasEnabled() bool {
	if o != nil && !common.IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Notification) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetShowButton returns the ShowButton field value if set, zero value otherwise.
func (o *Notification) GetShowButton() bool {
	if o == nil || common.IsNil(o.ShowButton) {
		var ret bool
		return ret
	}
	return *o.ShowButton
}

// GetShowButtonOk returns a tuple with the ShowButton field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notification) GetShowButtonOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ShowButton) {
		return nil, false
	}
	return o.ShowButton, true
}

// HasShowButton returns a boolean if a field has been set.
func (o *Notification) HasShowButton() bool {
	if o != nil && !common.IsNil(o.ShowButton) {
		return true
	}

	return false
}

// SetShowButton gets a reference to the given bool and assigns it to the ShowButton field.
func (o *Notification) SetShowButton(v bool) {
	o.ShowButton = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Notification) GetTitle() string {
	if o == nil || common.IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notification) GetTitleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Notification) HasTitle() bool {
	if o != nil && !common.IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Notification) SetTitle(v string) {
	o.Title = &v
}

func (o Notification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Notification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !common.IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !common.IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !common.IsNil(o.ShowButton) {
		toSerialize["showButton"] = o.ShowButton
	}
	if !common.IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableNotification struct {
	value *Notification
	isSet bool
}

func (v NullableNotification) Get() *Notification {
	return v.value
}

func (v *NullableNotification) Set(val *Notification) {
	v.value = val
	v.isSet = true
}

func (v NullableNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotification(val *Notification) *NullableNotification {
	return &NullableNotification{value: val, isSet: true}
}

func (v NullableNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *Notification) isValidCategory() bool {
	var allowedEnumValues = []string{"SaleWakeUp", "KeyPressed", ""}
	for _, allowed := range allowedEnumValues {
		if o.GetCategory() == allowed {
			return true
		}
	}
	return false
}
