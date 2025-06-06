/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the AuthenticationData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AuthenticationData{}

// AuthenticationData struct for AuthenticationData
type AuthenticationData struct {
	// Indicates when 3D Secure authentication should be attempted. This overrides all other rules, including [Dynamic 3D Secure settings](https://docs.adyen.com/risk-management/dynamic-3d-secure).  Possible values:  * **always**: Perform 3D Secure authentication. * **never**: Don't perform 3D Secure authentication. If PSD2 SCA or other national regulations require authentication, the transaction gets declined.
	AttemptAuthentication *string `json:"attemptAuthentication,omitempty"`
	// If set to true, you will only perform the [3D Secure 2 authentication](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only), and not the payment authorisation. Default: **false**.
	AuthenticationOnly *bool               `json:"authenticationOnly,omitempty"`
	ThreeDSRequestData *ThreeDSRequestData `json:"threeDSRequestData,omitempty"`
}

// NewAuthenticationData instantiates a new AuthenticationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationData() *AuthenticationData {
	this := AuthenticationData{}
	var authenticationOnly bool = false
	this.AuthenticationOnly = &authenticationOnly
	return &this
}

// NewAuthenticationDataWithDefaults instantiates a new AuthenticationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationDataWithDefaults() *AuthenticationData {
	this := AuthenticationData{}
	var authenticationOnly bool = false
	this.AuthenticationOnly = &authenticationOnly
	return &this
}

// GetAttemptAuthentication returns the AttemptAuthentication field value if set, zero value otherwise.
func (o *AuthenticationData) GetAttemptAuthentication() string {
	if o == nil || common.IsNil(o.AttemptAuthentication) {
		var ret string
		return ret
	}
	return *o.AttemptAuthentication
}

// GetAttemptAuthenticationOk returns a tuple with the AttemptAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationData) GetAttemptAuthenticationOk() (*string, bool) {
	if o == nil || common.IsNil(o.AttemptAuthentication) {
		return nil, false
	}
	return o.AttemptAuthentication, true
}

// HasAttemptAuthentication returns a boolean if a field has been set.
func (o *AuthenticationData) HasAttemptAuthentication() bool {
	if o != nil && !common.IsNil(o.AttemptAuthentication) {
		return true
	}

	return false
}

// SetAttemptAuthentication gets a reference to the given string and assigns it to the AttemptAuthentication field.
func (o *AuthenticationData) SetAttemptAuthentication(v string) {
	o.AttemptAuthentication = &v
}

// GetAuthenticationOnly returns the AuthenticationOnly field value if set, zero value otherwise.
func (o *AuthenticationData) GetAuthenticationOnly() bool {
	if o == nil || common.IsNil(o.AuthenticationOnly) {
		var ret bool
		return ret
	}
	return *o.AuthenticationOnly
}

// GetAuthenticationOnlyOk returns a tuple with the AuthenticationOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationData) GetAuthenticationOnlyOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AuthenticationOnly) {
		return nil, false
	}
	return o.AuthenticationOnly, true
}

// HasAuthenticationOnly returns a boolean if a field has been set.
func (o *AuthenticationData) HasAuthenticationOnly() bool {
	if o != nil && !common.IsNil(o.AuthenticationOnly) {
		return true
	}

	return false
}

// SetAuthenticationOnly gets a reference to the given bool and assigns it to the AuthenticationOnly field.
func (o *AuthenticationData) SetAuthenticationOnly(v bool) {
	o.AuthenticationOnly = &v
}

// GetThreeDSRequestData returns the ThreeDSRequestData field value if set, zero value otherwise.
func (o *AuthenticationData) GetThreeDSRequestData() ThreeDSRequestData {
	if o == nil || common.IsNil(o.ThreeDSRequestData) {
		var ret ThreeDSRequestData
		return ret
	}
	return *o.ThreeDSRequestData
}

// GetThreeDSRequestDataOk returns a tuple with the ThreeDSRequestData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationData) GetThreeDSRequestDataOk() (*ThreeDSRequestData, bool) {
	if o == nil || common.IsNil(o.ThreeDSRequestData) {
		return nil, false
	}
	return o.ThreeDSRequestData, true
}

// HasThreeDSRequestData returns a boolean if a field has been set.
func (o *AuthenticationData) HasThreeDSRequestData() bool {
	if o != nil && !common.IsNil(o.ThreeDSRequestData) {
		return true
	}

	return false
}

// SetThreeDSRequestData gets a reference to the given ThreeDSRequestData and assigns it to the ThreeDSRequestData field.
func (o *AuthenticationData) SetThreeDSRequestData(v ThreeDSRequestData) {
	o.ThreeDSRequestData = &v
}

func (o AuthenticationData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AttemptAuthentication) {
		toSerialize["attemptAuthentication"] = o.AttemptAuthentication
	}
	if !common.IsNil(o.AuthenticationOnly) {
		toSerialize["authenticationOnly"] = o.AuthenticationOnly
	}
	if !common.IsNil(o.ThreeDSRequestData) {
		toSerialize["threeDSRequestData"] = o.ThreeDSRequestData
	}
	return toSerialize, nil
}

type NullableAuthenticationData struct {
	value *AuthenticationData
	isSet bool
}

func (v NullableAuthenticationData) Get() *AuthenticationData {
	return v.value
}

func (v *NullableAuthenticationData) Set(val *AuthenticationData) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationData) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationData(val *AuthenticationData) *NullableAuthenticationData {
	return &NullableAuthenticationData{value: val, isSet: true}
}

func (v NullableAuthenticationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *AuthenticationData) isValidAttemptAuthentication() bool {
	var allowedEnumValues = []string{"always", "never"}
	for _, allowed := range allowedEnumValues {
		if o.GetAttemptAuthentication() == allowed {
			return true
		}
	}
	return false
}
