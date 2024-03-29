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
	"time"
)

// AccountHolderStatusChangeNotification struct for AccountHolderStatusChangeNotification
type AccountHolderStatusChangeNotification struct {
	Content *AccountHolderStatusChangeNotificationContent `json:"content,omitempty"`
	Error   *NotificationErrorContainer                   `json:"error,omitempty"`
	// The date and time when an event has been completed.
	EventDate time.Time `json:"eventDate"`
	// The event type of the notification.
	EventType string `json:"eventType"`
	// The user or process that has triggered the notification.
	ExecutingUserKey string `json:"executingUserKey"`
	// Indicates whether the notification originated from the live environment or the test environment. If true, the notification originated from the live environment. If false, the notification originated from the test environment.
	Live bool `json:"live"`
	// The PSP reference of the request from which the notification originates.
	PspReference string `json:"pspReference"`
}

// NewAccountHolderStatusChangeNotification instantiates a new AccountHolderStatusChangeNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountHolderStatusChangeNotification(eventDate time.Time, eventType string, executingUserKey string, live bool, pspReference string) *AccountHolderStatusChangeNotification {
	this := AccountHolderStatusChangeNotification{}
	this.EventDate = eventDate
	this.EventType = eventType
	this.ExecutingUserKey = executingUserKey
	this.Live = live
	this.PspReference = pspReference
	return &this
}

// NewAccountHolderStatusChangeNotificationWithDefaults instantiates a new AccountHolderStatusChangeNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountHolderStatusChangeNotificationWithDefaults() *AccountHolderStatusChangeNotification {
	this := AccountHolderStatusChangeNotification{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *AccountHolderStatusChangeNotification) GetContent() AccountHolderStatusChangeNotificationContent {
	if o == nil || o.Content == nil {
		var ret AccountHolderStatusChangeNotificationContent
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderStatusChangeNotification) GetContentOk() (*AccountHolderStatusChangeNotificationContent, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *AccountHolderStatusChangeNotification) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given AccountHolderStatusChangeNotificationContent and assigns it to the Content field.
func (o *AccountHolderStatusChangeNotification) SetContent(v AccountHolderStatusChangeNotificationContent) {
	o.Content = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *AccountHolderStatusChangeNotification) GetError() NotificationErrorContainer {
	if o == nil || o.Error == nil {
		var ret NotificationErrorContainer
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderStatusChangeNotification) GetErrorOk() (*NotificationErrorContainer, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *AccountHolderStatusChangeNotification) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given NotificationErrorContainer and assigns it to the Error field.
func (o *AccountHolderStatusChangeNotification) SetError(v NotificationErrorContainer) {
	o.Error = &v
}

// GetEventDate returns the EventDate field value
func (o *AccountHolderStatusChangeNotification) GetEventDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventDate
}

// GetEventDateOk returns a tuple with the EventDate field value
// and a boolean to check if the value has been set.
func (o *AccountHolderStatusChangeNotification) GetEventDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventDate, true
}

// SetEventDate sets field value
func (o *AccountHolderStatusChangeNotification) SetEventDate(v time.Time) {
	o.EventDate = v
}

// GetEventType returns the EventType field value
func (o *AccountHolderStatusChangeNotification) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *AccountHolderStatusChangeNotification) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *AccountHolderStatusChangeNotification) SetEventType(v string) {
	o.EventType = v
}

// GetExecutingUserKey returns the ExecutingUserKey field value
func (o *AccountHolderStatusChangeNotification) GetExecutingUserKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExecutingUserKey
}

// GetExecutingUserKeyOk returns a tuple with the ExecutingUserKey field value
// and a boolean to check if the value has been set.
func (o *AccountHolderStatusChangeNotification) GetExecutingUserKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutingUserKey, true
}

// SetExecutingUserKey sets field value
func (o *AccountHolderStatusChangeNotification) SetExecutingUserKey(v string) {
	o.ExecutingUserKey = v
}

// GetLive returns the Live field value
func (o *AccountHolderStatusChangeNotification) GetLive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Live
}

// GetLiveOk returns a tuple with the Live field value
// and a boolean to check if the value has been set.
func (o *AccountHolderStatusChangeNotification) GetLiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Live, true
}

// SetLive sets field value
func (o *AccountHolderStatusChangeNotification) SetLive(v bool) {
	o.Live = v
}

// GetPspReference returns the PspReference field value
func (o *AccountHolderStatusChangeNotification) GetPspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value
// and a boolean to check if the value has been set.
func (o *AccountHolderStatusChangeNotification) GetPspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PspReference, true
}

// SetPspReference sets field value
func (o *AccountHolderStatusChangeNotification) SetPspReference(v string) {
	o.PspReference = v
}

func (o AccountHolderStatusChangeNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if true {
		toSerialize["eventDate"] = o.EventDate
	}
	if true {
		toSerialize["eventType"] = o.EventType
	}
	if true {
		toSerialize["executingUserKey"] = o.ExecutingUserKey
	}
	if true {
		toSerialize["live"] = o.Live
	}
	if true {
		toSerialize["pspReference"] = o.PspReference
	}
	return json.Marshal(toSerialize)
}

type NullableAccountHolderStatusChangeNotification struct {
	value *AccountHolderStatusChangeNotification
	isSet bool
}

func (v NullableAccountHolderStatusChangeNotification) Get() *AccountHolderStatusChangeNotification {
	return v.value
}

func (v *NullableAccountHolderStatusChangeNotification) Set(val *AccountHolderStatusChangeNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountHolderStatusChangeNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountHolderStatusChangeNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountHolderStatusChangeNotification(val *AccountHolderStatusChangeNotification) *NullableAccountHolderStatusChangeNotification {
	return &NullableAccountHolderStatusChangeNotification{value: val, isSet: true}
}

func (v NullableAccountHolderStatusChangeNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountHolderStatusChangeNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
