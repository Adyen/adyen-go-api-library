/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the NetworkToken type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NetworkToken{}

// NetworkToken struct for NetworkToken
type NetworkToken struct {
	// The card brand variant of the payment instrument associated with the network token. For example, **mc_prepaid_mrw**.
	BrandVariant *string `json:"brandVariant,omitempty"`
	// Date and time when the network token was created, in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) extended format. For example, **2020-12-18T10:15:30+01:00**..
	CreationDate *time.Time  `json:"creationDate,omitempty"`
	Device       *DeviceInfo `json:"device,omitempty"`
	// The unique identifier of the network token.
	Id *string `json:"id,omitempty"`
	// The unique identifier of the payment instrument to which this network token belongs to.
	PaymentInstrumentId *string `json:"paymentInstrumentId,omitempty"`
	// The status of the network token. Possible values: **active**, **inactive**, **suspended**, **closed**.
	Status *string `json:"status,omitempty"`
	// The last four digits of the network token `id`.
	TokenLastFour *string `json:"tokenLastFour,omitempty"`
	// The type of network token. For example, **wallet**, **cof**.
	Type *string `json:"type,omitempty"`
}

// NewNetworkToken instantiates a new NetworkToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkToken() *NetworkToken {
	this := NetworkToken{}
	return &this
}

// NewNetworkTokenWithDefaults instantiates a new NetworkToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkTokenWithDefaults() *NetworkToken {
	this := NetworkToken{}
	return &this
}

// GetBrandVariant returns the BrandVariant field value if set, zero value otherwise.
func (o *NetworkToken) GetBrandVariant() string {
	if o == nil || common.IsNil(o.BrandVariant) {
		var ret string
		return ret
	}
	return *o.BrandVariant
}

// GetBrandVariantOk returns a tuple with the BrandVariant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkToken) GetBrandVariantOk() (*string, bool) {
	if o == nil || common.IsNil(o.BrandVariant) {
		return nil, false
	}
	return o.BrandVariant, true
}

// HasBrandVariant returns a boolean if a field has been set.
func (o *NetworkToken) HasBrandVariant() bool {
	if o != nil && !common.IsNil(o.BrandVariant) {
		return true
	}

	return false
}

// SetBrandVariant gets a reference to the given string and assigns it to the BrandVariant field.
func (o *NetworkToken) SetBrandVariant(v string) {
	o.BrandVariant = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *NetworkToken) GetCreationDate() time.Time {
	if o == nil || common.IsNil(o.CreationDate) {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkToken) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *NetworkToken) HasCreationDate() bool {
	if o != nil && !common.IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *NetworkToken) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *NetworkToken) GetDevice() DeviceInfo {
	if o == nil || common.IsNil(o.Device) {
		var ret DeviceInfo
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkToken) GetDeviceOk() (*DeviceInfo, bool) {
	if o == nil || common.IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *NetworkToken) HasDevice() bool {
	if o != nil && !common.IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given DeviceInfo and assigns it to the Device field.
func (o *NetworkToken) SetDevice(v DeviceInfo) {
	o.Device = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworkToken) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkToken) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworkToken) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NetworkToken) SetId(v string) {
	o.Id = &v
}

// GetPaymentInstrumentId returns the PaymentInstrumentId field value if set, zero value otherwise.
func (o *NetworkToken) GetPaymentInstrumentId() string {
	if o == nil || common.IsNil(o.PaymentInstrumentId) {
		var ret string
		return ret
	}
	return *o.PaymentInstrumentId
}

// GetPaymentInstrumentIdOk returns a tuple with the PaymentInstrumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkToken) GetPaymentInstrumentIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentInstrumentId) {
		return nil, false
	}
	return o.PaymentInstrumentId, true
}

// HasPaymentInstrumentId returns a boolean if a field has been set.
func (o *NetworkToken) HasPaymentInstrumentId() bool {
	if o != nil && !common.IsNil(o.PaymentInstrumentId) {
		return true
	}

	return false
}

// SetPaymentInstrumentId gets a reference to the given string and assigns it to the PaymentInstrumentId field.
func (o *NetworkToken) SetPaymentInstrumentId(v string) {
	o.PaymentInstrumentId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NetworkToken) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkToken) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NetworkToken) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NetworkToken) SetStatus(v string) {
	o.Status = &v
}

// GetTokenLastFour returns the TokenLastFour field value if set, zero value otherwise.
func (o *NetworkToken) GetTokenLastFour() string {
	if o == nil || common.IsNil(o.TokenLastFour) {
		var ret string
		return ret
	}
	return *o.TokenLastFour
}

// GetTokenLastFourOk returns a tuple with the TokenLastFour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkToken) GetTokenLastFourOk() (*string, bool) {
	if o == nil || common.IsNil(o.TokenLastFour) {
		return nil, false
	}
	return o.TokenLastFour, true
}

// HasTokenLastFour returns a boolean if a field has been set.
func (o *NetworkToken) HasTokenLastFour() bool {
	if o != nil && !common.IsNil(o.TokenLastFour) {
		return true
	}

	return false
}

// SetTokenLastFour gets a reference to the given string and assigns it to the TokenLastFour field.
func (o *NetworkToken) SetTokenLastFour(v string) {
	o.TokenLastFour = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworkToken) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkToken) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworkToken) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NetworkToken) SetType(v string) {
	o.Type = &v
}

func (o NetworkToken) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BrandVariant) {
		toSerialize["brandVariant"] = o.BrandVariant
	}
	if !common.IsNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	if !common.IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.PaymentInstrumentId) {
		toSerialize["paymentInstrumentId"] = o.PaymentInstrumentId
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.TokenLastFour) {
		toSerialize["tokenLastFour"] = o.TokenLastFour
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableNetworkToken struct {
	value *NetworkToken
	isSet bool
}

func (v NullableNetworkToken) Get() *NetworkToken {
	return v.value
}

func (v *NullableNetworkToken) Set(val *NetworkToken) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkToken) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkToken(val *NetworkToken) *NullableNetworkToken {
	return &NullableNetworkToken{value: val, isSet: true}
}

func (v NullableNetworkToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *NetworkToken) isValidStatus() bool {
	var allowedEnumValues = []string{"active", "inactive", "suspended", "closed"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
