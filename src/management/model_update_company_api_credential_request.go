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

// checks if the UpdateCompanyApiCredentialRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UpdateCompanyApiCredentialRequest{}

// UpdateCompanyApiCredentialRequest struct for UpdateCompanyApiCredentialRequest
type UpdateCompanyApiCredentialRequest struct {
	// Indicates if the API credential is enabled.
	Active *bool `json:"active,omitempty"`
	// The new list of [allowed origins](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) for the API credential.
	AllowedOrigins []string `json:"allowedOrigins,omitempty"`
	// List of merchant accounts that the API credential has access to.
	AssociatedMerchantAccounts []string `json:"associatedMerchantAccounts,omitempty"`
	// Description of the API credential.
	Description *string `json:"description,omitempty"`
	// List of [roles](https://docs.adyen.com/development-resources/api-credentials#roles-1) for the API credential. Only roles assigned to 'ws@Company.<CompanyName>' can be assigned to other API credentials.
	Roles []string `json:"roles,omitempty"`
}

// NewUpdateCompanyApiCredentialRequest instantiates a new UpdateCompanyApiCredentialRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCompanyApiCredentialRequest() *UpdateCompanyApiCredentialRequest {
	this := UpdateCompanyApiCredentialRequest{}
	return &this
}

// NewUpdateCompanyApiCredentialRequestWithDefaults instantiates a new UpdateCompanyApiCredentialRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCompanyApiCredentialRequestWithDefaults() *UpdateCompanyApiCredentialRequest {
	this := UpdateCompanyApiCredentialRequest{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *UpdateCompanyApiCredentialRequest) GetActive() bool {
	if o == nil || common.IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyApiCredentialRequest) GetActiveOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *UpdateCompanyApiCredentialRequest) HasActive() bool {
	if o != nil && !common.IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *UpdateCompanyApiCredentialRequest) SetActive(v bool) {
	o.Active = &v
}

// GetAllowedOrigins returns the AllowedOrigins field value if set, zero value otherwise.
func (o *UpdateCompanyApiCredentialRequest) GetAllowedOrigins() []string {
	if o == nil || common.IsNil(o.AllowedOrigins) {
		var ret []string
		return ret
	}
	return o.AllowedOrigins
}

// GetAllowedOriginsOk returns a tuple with the AllowedOrigins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyApiCredentialRequest) GetAllowedOriginsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.AllowedOrigins) {
		return nil, false
	}
	return o.AllowedOrigins, true
}

// HasAllowedOrigins returns a boolean if a field has been set.
func (o *UpdateCompanyApiCredentialRequest) HasAllowedOrigins() bool {
	if o != nil && !common.IsNil(o.AllowedOrigins) {
		return true
	}

	return false
}

// SetAllowedOrigins gets a reference to the given []string and assigns it to the AllowedOrigins field.
func (o *UpdateCompanyApiCredentialRequest) SetAllowedOrigins(v []string) {
	o.AllowedOrigins = v
}

// GetAssociatedMerchantAccounts returns the AssociatedMerchantAccounts field value if set, zero value otherwise.
func (o *UpdateCompanyApiCredentialRequest) GetAssociatedMerchantAccounts() []string {
	if o == nil || common.IsNil(o.AssociatedMerchantAccounts) {
		var ret []string
		return ret
	}
	return o.AssociatedMerchantAccounts
}

// GetAssociatedMerchantAccountsOk returns a tuple with the AssociatedMerchantAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyApiCredentialRequest) GetAssociatedMerchantAccountsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.AssociatedMerchantAccounts) {
		return nil, false
	}
	return o.AssociatedMerchantAccounts, true
}

// HasAssociatedMerchantAccounts returns a boolean if a field has been set.
func (o *UpdateCompanyApiCredentialRequest) HasAssociatedMerchantAccounts() bool {
	if o != nil && !common.IsNil(o.AssociatedMerchantAccounts) {
		return true
	}

	return false
}

// SetAssociatedMerchantAccounts gets a reference to the given []string and assigns it to the AssociatedMerchantAccounts field.
func (o *UpdateCompanyApiCredentialRequest) SetAssociatedMerchantAccounts(v []string) {
	o.AssociatedMerchantAccounts = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateCompanyApiCredentialRequest) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyApiCredentialRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateCompanyApiCredentialRequest) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateCompanyApiCredentialRequest) SetDescription(v string) {
	o.Description = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *UpdateCompanyApiCredentialRequest) GetRoles() []string {
	if o == nil || common.IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyApiCredentialRequest) GetRolesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UpdateCompanyApiCredentialRequest) HasRoles() bool {
	if o != nil && !common.IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *UpdateCompanyApiCredentialRequest) SetRoles(v []string) {
	o.Roles = v
}

func (o UpdateCompanyApiCredentialRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCompanyApiCredentialRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !common.IsNil(o.AllowedOrigins) {
		toSerialize["allowedOrigins"] = o.AllowedOrigins
	}
	if !common.IsNil(o.AssociatedMerchantAccounts) {
		toSerialize["associatedMerchantAccounts"] = o.AssociatedMerchantAccounts
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableUpdateCompanyApiCredentialRequest struct {
	value *UpdateCompanyApiCredentialRequest
	isSet bool
}

func (v NullableUpdateCompanyApiCredentialRequest) Get() *UpdateCompanyApiCredentialRequest {
	return v.value
}

func (v *NullableUpdateCompanyApiCredentialRequest) Set(val *UpdateCompanyApiCredentialRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCompanyApiCredentialRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCompanyApiCredentialRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCompanyApiCredentialRequest(val *UpdateCompanyApiCredentialRequest) *NullableUpdateCompanyApiCredentialRequest {
	return &NullableUpdateCompanyApiCredentialRequest{value: val, isSet: true}
}

func (v NullableUpdateCompanyApiCredentialRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCompanyApiCredentialRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
