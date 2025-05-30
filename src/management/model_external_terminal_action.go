/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the ExternalTerminalAction type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExternalTerminalAction{}

// ExternalTerminalAction struct for ExternalTerminalAction
type ExternalTerminalAction struct {
	// The type of terminal action: **InstallAndroidApp**, **UninstallAndroidApp**, **InstallAndroidCertificate**, or **UninstallAndroidCertificate**.
	ActionType *string `json:"actionType,omitempty"`
	// Technical information about the terminal action.
	Config *string `json:"config,omitempty"`
	// The date and time when the action was carried out.
	ConfirmedAt *time.Time `json:"confirmedAt,omitempty"`
	// The unique ID of the terminal action.
	Id *string `json:"id,omitempty"`
	// The result message for the action.
	Result *string `json:"result,omitempty"`
	// The date and time when the action was scheduled to happen.
	ScheduledAt *time.Time `json:"scheduledAt,omitempty"`
	// The status of the terminal action: **pending**, **successful**, **failed**, **cancelled**, or **tryLater**.
	Status *string `json:"status,omitempty"`
	// The unique ID of the terminal that the action applies to.
	TerminalId *string `json:"terminalId,omitempty"`
}

// NewExternalTerminalAction instantiates a new ExternalTerminalAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalTerminalAction() *ExternalTerminalAction {
	this := ExternalTerminalAction{}
	return &this
}

// NewExternalTerminalActionWithDefaults instantiates a new ExternalTerminalAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalTerminalActionWithDefaults() *ExternalTerminalAction {
	this := ExternalTerminalAction{}
	return &this
}

// GetActionType returns the ActionType field value if set, zero value otherwise.
func (o *ExternalTerminalAction) GetActionType() string {
	if o == nil || common.IsNil(o.ActionType) {
		var ret string
		return ret
	}
	return *o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalTerminalAction) GetActionTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ActionType) {
		return nil, false
	}
	return o.ActionType, true
}

// HasActionType returns a boolean if a field has been set.
func (o *ExternalTerminalAction) HasActionType() bool {
	if o != nil && !common.IsNil(o.ActionType) {
		return true
	}

	return false
}

// SetActionType gets a reference to the given string and assigns it to the ActionType field.
func (o *ExternalTerminalAction) SetActionType(v string) {
	o.ActionType = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *ExternalTerminalAction) GetConfig() string {
	if o == nil || common.IsNil(o.Config) {
		var ret string
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalTerminalAction) GetConfigOk() (*string, bool) {
	if o == nil || common.IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *ExternalTerminalAction) HasConfig() bool {
	if o != nil && !common.IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given string and assigns it to the Config field.
func (o *ExternalTerminalAction) SetConfig(v string) {
	o.Config = &v
}

// GetConfirmedAt returns the ConfirmedAt field value if set, zero value otherwise.
func (o *ExternalTerminalAction) GetConfirmedAt() time.Time {
	if o == nil || common.IsNil(o.ConfirmedAt) {
		var ret time.Time
		return ret
	}
	return *o.ConfirmedAt
}

// GetConfirmedAtOk returns a tuple with the ConfirmedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalTerminalAction) GetConfirmedAtOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.ConfirmedAt) {
		return nil, false
	}
	return o.ConfirmedAt, true
}

// HasConfirmedAt returns a boolean if a field has been set.
func (o *ExternalTerminalAction) HasConfirmedAt() bool {
	if o != nil && !common.IsNil(o.ConfirmedAt) {
		return true
	}

	return false
}

// SetConfirmedAt gets a reference to the given time.Time and assigns it to the ConfirmedAt field.
func (o *ExternalTerminalAction) SetConfirmedAt(v time.Time) {
	o.ConfirmedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExternalTerminalAction) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalTerminalAction) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExternalTerminalAction) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ExternalTerminalAction) SetId(v string) {
	o.Id = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ExternalTerminalAction) GetResult() string {
	if o == nil || common.IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalTerminalAction) GetResultOk() (*string, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ExternalTerminalAction) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *ExternalTerminalAction) SetResult(v string) {
	o.Result = &v
}

// GetScheduledAt returns the ScheduledAt field value if set, zero value otherwise.
func (o *ExternalTerminalAction) GetScheduledAt() time.Time {
	if o == nil || common.IsNil(o.ScheduledAt) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledAt
}

// GetScheduledAtOk returns a tuple with the ScheduledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalTerminalAction) GetScheduledAtOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.ScheduledAt) {
		return nil, false
	}
	return o.ScheduledAt, true
}

// HasScheduledAt returns a boolean if a field has been set.
func (o *ExternalTerminalAction) HasScheduledAt() bool {
	if o != nil && !common.IsNil(o.ScheduledAt) {
		return true
	}

	return false
}

// SetScheduledAt gets a reference to the given time.Time and assigns it to the ScheduledAt field.
func (o *ExternalTerminalAction) SetScheduledAt(v time.Time) {
	o.ScheduledAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ExternalTerminalAction) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalTerminalAction) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ExternalTerminalAction) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ExternalTerminalAction) SetStatus(v string) {
	o.Status = &v
}

// GetTerminalId returns the TerminalId field value if set, zero value otherwise.
func (o *ExternalTerminalAction) GetTerminalId() string {
	if o == nil || common.IsNil(o.TerminalId) {
		var ret string
		return ret
	}
	return *o.TerminalId
}

// GetTerminalIdOk returns a tuple with the TerminalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalTerminalAction) GetTerminalIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TerminalId) {
		return nil, false
	}
	return o.TerminalId, true
}

// HasTerminalId returns a boolean if a field has been set.
func (o *ExternalTerminalAction) HasTerminalId() bool {
	if o != nil && !common.IsNil(o.TerminalId) {
		return true
	}

	return false
}

// SetTerminalId gets a reference to the given string and assigns it to the TerminalId field.
func (o *ExternalTerminalAction) SetTerminalId(v string) {
	o.TerminalId = &v
}

func (o ExternalTerminalAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalTerminalAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ActionType) {
		toSerialize["actionType"] = o.ActionType
	}
	if !common.IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !common.IsNil(o.ConfirmedAt) {
		toSerialize["confirmedAt"] = o.ConfirmedAt
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !common.IsNil(o.ScheduledAt) {
		toSerialize["scheduledAt"] = o.ScheduledAt
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.TerminalId) {
		toSerialize["terminalId"] = o.TerminalId
	}
	return toSerialize, nil
}

type NullableExternalTerminalAction struct {
	value *ExternalTerminalAction
	isSet bool
}

func (v NullableExternalTerminalAction) Get() *ExternalTerminalAction {
	return v.value
}

func (v *NullableExternalTerminalAction) Set(val *ExternalTerminalAction) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalTerminalAction) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalTerminalAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalTerminalAction(val *ExternalTerminalAction) *NullableExternalTerminalAction {
	return &NullableExternalTerminalAction{value: val, isSet: true}
}

func (v NullableExternalTerminalAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalTerminalAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
