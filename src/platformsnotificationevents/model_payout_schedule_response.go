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

// PayoutScheduleResponse struct for PayoutScheduleResponse
type PayoutScheduleResponse struct {
	// The date of the next scheduled payout.
	NextScheduledPayout time.Time `json:"nextScheduledPayout"`
	// The payout schedule of the account. >Permitted values: `DEFAULT`, `HOLD`, `DAILY`, `WEEKLY`, `MONTHLY`.
	Schedule string `json:"schedule"`
}

// NewPayoutScheduleResponse instantiates a new PayoutScheduleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayoutScheduleResponse(nextScheduledPayout time.Time, schedule string) *PayoutScheduleResponse {
	this := PayoutScheduleResponse{}
	this.NextScheduledPayout = nextScheduledPayout
	this.Schedule = schedule
	return &this
}

// NewPayoutScheduleResponseWithDefaults instantiates a new PayoutScheduleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayoutScheduleResponseWithDefaults() *PayoutScheduleResponse {
	this := PayoutScheduleResponse{}
	return &this
}

// GetNextScheduledPayout returns the NextScheduledPayout field value
func (o *PayoutScheduleResponse) GetNextScheduledPayout() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.NextScheduledPayout
}

// GetNextScheduledPayoutOk returns a tuple with the NextScheduledPayout field value
// and a boolean to check if the value has been set.
func (o *PayoutScheduleResponse) GetNextScheduledPayoutOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextScheduledPayout, true
}

// SetNextScheduledPayout sets field value
func (o *PayoutScheduleResponse) SetNextScheduledPayout(v time.Time) {
	o.NextScheduledPayout = v
}

// GetSchedule returns the Schedule field value
func (o *PayoutScheduleResponse) GetSchedule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *PayoutScheduleResponse) GetScheduleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *PayoutScheduleResponse) SetSchedule(v string) {
	o.Schedule = v
}

func (o PayoutScheduleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nextScheduledPayout"] = o.NextScheduledPayout
	}
	if true {
		toSerialize["schedule"] = o.Schedule
	}
	return json.Marshal(toSerialize)
}

type NullablePayoutScheduleResponse struct {
	value *PayoutScheduleResponse
	isSet bool
}

func (v NullablePayoutScheduleResponse) Get() *PayoutScheduleResponse {
	return v.value
}

func (v *NullablePayoutScheduleResponse) Set(val *PayoutScheduleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePayoutScheduleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePayoutScheduleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayoutScheduleResponse(val *PayoutScheduleResponse) *NullablePayoutScheduleResponse {
	return &NullablePayoutScheduleResponse{value: val, isSet: true}
}

func (v NullablePayoutScheduleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayoutScheduleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
