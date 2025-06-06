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

// checks if the UpdateMerchantApiCredentialRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UpdateMerchantApiCredentialRequest{}

// UpdateMerchantApiCredentialRequest struct for UpdateMerchantApiCredentialRequest
type UpdateMerchantApiCredentialRequest struct {
	// Indicates if the API credential is enabled.
	Active *bool `json:"active,omitempty"`
	// The new list of [allowed origins](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) for the API credential.
	AllowedOrigins []string `json:"allowedOrigins,omitempty"`
	// Description of the API credential.
	Description *string `json:"description,omitempty"`
	// List of [roles](https://docs.adyen.com/development-resources/api-credentials#roles-1) for the API credential. Only roles assigned to 'ws@Company.<CompanyName>' can be assigned to other API credentials.
	Roles []string `json:"roles,omitempty"`
}

// NewUpdateMerchantApiCredentialRequest instantiates a new UpdateMerchantApiCredentialRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMerchantApiCredentialRequest() *UpdateMerchantApiCredentialRequest {
	this := UpdateMerchantApiCredentialRequest{}
	return &this
}

// NewUpdateMerchantApiCredentialRequestWithDefaults instantiates a new UpdateMerchantApiCredentialRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMerchantApiCredentialRequestWithDefaults() *UpdateMerchantApiCredentialRequest {
	this := UpdateMerchantApiCredentialRequest{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *UpdateMerchantApiCredentialRequest) GetActive() bool {
	if o == nil || common.IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMerchantApiCredentialRequest) GetActiveOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *UpdateMerchantApiCredentialRequest) HasActive() bool {
	if o != nil && !common.IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *UpdateMerchantApiCredentialRequest) SetActive(v bool) {
	o.Active = &v
}

// GetAllowedOrigins returns the AllowedOrigins field value if set, zero value otherwise.
func (o *UpdateMerchantApiCredentialRequest) GetAllowedOrigins() []string {
	if o == nil || common.IsNil(o.AllowedOrigins) {
		var ret []string
		return ret
	}
	return o.AllowedOrigins
}

// GetAllowedOriginsOk returns a tuple with the AllowedOrigins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMerchantApiCredentialRequest) GetAllowedOriginsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.AllowedOrigins) {
		return nil, false
	}
	return o.AllowedOrigins, true
}

// HasAllowedOrigins returns a boolean if a field has been set.
func (o *UpdateMerchantApiCredentialRequest) HasAllowedOrigins() bool {
	if o != nil && !common.IsNil(o.AllowedOrigins) {
		return true
	}

	return false
}

// SetAllowedOrigins gets a reference to the given []string and assigns it to the AllowedOrigins field.
func (o *UpdateMerchantApiCredentialRequest) SetAllowedOrigins(v []string) {
	o.AllowedOrigins = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateMerchantApiCredentialRequest) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMerchantApiCredentialRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateMerchantApiCredentialRequest) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateMerchantApiCredentialRequest) SetDescription(v string) {
	o.Description = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *UpdateMerchantApiCredentialRequest) GetRoles() []string {
	if o == nil || common.IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMerchantApiCredentialRequest) GetRolesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UpdateMerchantApiCredentialRequest) HasRoles() bool {
	if o != nil && !common.IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *UpdateMerchantApiCredentialRequest) SetRoles(v []string) {
	o.Roles = v
}

func (o UpdateMerchantApiCredentialRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMerchantApiCredentialRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !common.IsNil(o.AllowedOrigins) {
		toSerialize["allowedOrigins"] = o.AllowedOrigins
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableUpdateMerchantApiCredentialRequest struct {
	value *UpdateMerchantApiCredentialRequest
	isSet bool
}

func (v NullableUpdateMerchantApiCredentialRequest) Get() *UpdateMerchantApiCredentialRequest {
	return v.value
}

func (v *NullableUpdateMerchantApiCredentialRequest) Set(val *UpdateMerchantApiCredentialRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMerchantApiCredentialRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMerchantApiCredentialRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMerchantApiCredentialRequest(val *UpdateMerchantApiCredentialRequest) *NullableUpdateMerchantApiCredentialRequest {
	return &NullableUpdateMerchantApiCredentialRequest{value: val, isSet: true}
}

func (v NullableUpdateMerchantApiCredentialRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMerchantApiCredentialRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
