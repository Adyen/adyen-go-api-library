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

// checks if the AndroidCertificate type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AndroidCertificate{}

// AndroidCertificate struct for AndroidCertificate
type AndroidCertificate struct {
	// The description that was provided when uploading the certificate.
	Description *string `json:"description,omitempty"`
	// The file format of the certificate, as indicated by the file extension. For example, **.cert** or **.pem**.
	Extension *string `json:"extension,omitempty"`
	// The unique identifier of the certificate.
	Id string `json:"id"`
	// The file name of the certificate. For example, **mycert**.
	Name *string `json:"name,omitempty"`
	// The date when the certificate stops to be valid.
	NotAfter *time.Time `json:"notAfter,omitempty"`
	// The date when the certificate starts to be valid.
	NotBefore *time.Time `json:"notBefore,omitempty"`
	// The status of the certificate.
	Status *string `json:"status,omitempty"`
}

// NewAndroidCertificate instantiates a new AndroidCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAndroidCertificate(id string) *AndroidCertificate {
	this := AndroidCertificate{}
	this.Id = id
	return &this
}

// NewAndroidCertificateWithDefaults instantiates a new AndroidCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAndroidCertificateWithDefaults() *AndroidCertificate {
	this := AndroidCertificate{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AndroidCertificate) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidCertificate) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AndroidCertificate) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AndroidCertificate) SetDescription(v string) {
	o.Description = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *AndroidCertificate) GetExtension() string {
	if o == nil || common.IsNil(o.Extension) {
		var ret string
		return ret
	}
	return *o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidCertificate) GetExtensionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *AndroidCertificate) HasExtension() bool {
	if o != nil && !common.IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given string and assigns it to the Extension field.
func (o *AndroidCertificate) SetExtension(v string) {
	o.Extension = &v
}

// GetId returns the Id field value
func (o *AndroidCertificate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AndroidCertificate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AndroidCertificate) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AndroidCertificate) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidCertificate) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AndroidCertificate) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AndroidCertificate) SetName(v string) {
	o.Name = &v
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise.
func (o *AndroidCertificate) GetNotAfter() time.Time {
	if o == nil || common.IsNil(o.NotAfter) {
		var ret time.Time
		return ret
	}
	return *o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidCertificate) GetNotAfterOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.NotAfter) {
		return nil, false
	}
	return o.NotAfter, true
}

// HasNotAfter returns a boolean if a field has been set.
func (o *AndroidCertificate) HasNotAfter() bool {
	if o != nil && !common.IsNil(o.NotAfter) {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given time.Time and assigns it to the NotAfter field.
func (o *AndroidCertificate) SetNotAfter(v time.Time) {
	o.NotAfter = &v
}

// GetNotBefore returns the NotBefore field value if set, zero value otherwise.
func (o *AndroidCertificate) GetNotBefore() time.Time {
	if o == nil || common.IsNil(o.NotBefore) {
		var ret time.Time
		return ret
	}
	return *o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidCertificate) GetNotBeforeOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.NotBefore) {
		return nil, false
	}
	return o.NotBefore, true
}

// HasNotBefore returns a boolean if a field has been set.
func (o *AndroidCertificate) HasNotBefore() bool {
	if o != nil && !common.IsNil(o.NotBefore) {
		return true
	}

	return false
}

// SetNotBefore gets a reference to the given time.Time and assigns it to the NotBefore field.
func (o *AndroidCertificate) SetNotBefore(v time.Time) {
	o.NotBefore = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AndroidCertificate) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidCertificate) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AndroidCertificate) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AndroidCertificate) SetStatus(v string) {
	o.Status = &v
}

func (o AndroidCertificate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AndroidCertificate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.Extension) {
		toSerialize["extension"] = o.Extension
	}
	toSerialize["id"] = o.Id
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.NotAfter) {
		toSerialize["notAfter"] = o.NotAfter
	}
	if !common.IsNil(o.NotBefore) {
		toSerialize["notBefore"] = o.NotBefore
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableAndroidCertificate struct {
	value *AndroidCertificate
	isSet bool
}

func (v NullableAndroidCertificate) Get() *AndroidCertificate {
	return v.value
}

func (v *NullableAndroidCertificate) Set(val *AndroidCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableAndroidCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableAndroidCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAndroidCertificate(val *AndroidCertificate) *NullableAndroidCertificate {
	return &NullableAndroidCertificate{value: val, isSet: true}
}

func (v NullableAndroidCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAndroidCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
