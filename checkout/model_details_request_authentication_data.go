/*
Adyen Checkout API

Adyen Checkout API provides a simple and flexible way to initiate and authorise online payments. You can use the same integration for payments made with cards (including 3D Secure), mobile wallets, and local payment methods (for example, iDEAL and Sofort).  This API reference provides information on available endpoints and how to interact with them. To learn more about the API, visit [online payments documentation](https://docs.adyen.com/online-payments).  ## Authentication Each request to Checkout API must be signed with an API key. For this, [get your API key](https://docs.adyen.com/development-resources/api-credentials#generate-api-key) from your Customer Area, and set this key to the `X-API-Key` header value, for example:  ``` curl -H \"Content-Type: application/json\" \\ -H \"X-API-Key: YOUR_API_KEY\" \\ ... ``` ## Versioning Checkout API supports [versioning](https://docs.adyen.com/development-resources/versioning) using a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://checkout-test.adyen.com/v70/payments ```  ## Going live  To access the live endpoints, you need an API key from your live Customer Area.  The live endpoint URLs contain a prefix which is unique to your company account, for example: ``` https://{PREFIX}-checkout-live.adyenpayments.com/checkout/v70/payments ```  Get your `{PREFIX}` from your live Customer Area under **Developers** > **API URLs** > **Prefix**.  When preparing to do live transactions with Checkout API, follow the [go-live checklist](https://docs.adyen.com/online-payments/go-live-checklist) to make sure you've got all the required configuration in place.  ## Release notes Have a look at the [release notes](https://docs.adyen.com/online-payments/release-notes?integration_type=api&version=70) to find out what changed in this version!

API version: 70
Contact: developer-experience@adyen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"
)

// checks if the DetailsRequestAuthenticationData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetailsRequestAuthenticationData{}

// DetailsRequestAuthenticationData struct for DetailsRequestAuthenticationData
type DetailsRequestAuthenticationData struct {
	// If set to true, you will only perform the [3D Secure 2 authentication](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only), and not the payment authorisation. Default: *false**.
	AuthenticationOnly *bool `json:"authenticationOnly,omitempty"`
}

// NewDetailsRequestAuthenticationData instantiates a new DetailsRequestAuthenticationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetailsRequestAuthenticationData() *DetailsRequestAuthenticationData {
	this := DetailsRequestAuthenticationData{}
	var authenticationOnly bool = false
	this.AuthenticationOnly = &authenticationOnly
	return &this
}

// NewDetailsRequestAuthenticationDataWithDefaults instantiates a new DetailsRequestAuthenticationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailsRequestAuthenticationDataWithDefaults() *DetailsRequestAuthenticationData {
	this := DetailsRequestAuthenticationData{}
	var authenticationOnly bool = false
	this.AuthenticationOnly = &authenticationOnly
	return &this
}

// GetAuthenticationOnly returns the AuthenticationOnly field value if set, zero value otherwise.
func (o *DetailsRequestAuthenticationData) GetAuthenticationOnly() bool {
	if o == nil || IsNil(o.AuthenticationOnly) {
		var ret bool
		return ret
	}
	return *o.AuthenticationOnly
}

// GetAuthenticationOnlyOk returns a tuple with the AuthenticationOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailsRequestAuthenticationData) GetAuthenticationOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.AuthenticationOnly) {
		return nil, false
	}
	return o.AuthenticationOnly, true
}

// HasAuthenticationOnly returns a boolean if a field has been set.
func (o *DetailsRequestAuthenticationData) HasAuthenticationOnly() bool {
	if o != nil && !IsNil(o.AuthenticationOnly) {
		return true
	}

	return false
}

// SetAuthenticationOnly gets a reference to the given bool and assigns it to the AuthenticationOnly field.
func (o *DetailsRequestAuthenticationData) SetAuthenticationOnly(v bool) {
	o.AuthenticationOnly = &v
}

func (o DetailsRequestAuthenticationData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetailsRequestAuthenticationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthenticationOnly) {
		toSerialize["authenticationOnly"] = o.AuthenticationOnly
	}
	return toSerialize, nil
}

type NullableDetailsRequestAuthenticationData struct {
	value *DetailsRequestAuthenticationData
	isSet bool
}

func (v NullableDetailsRequestAuthenticationData) Get() *DetailsRequestAuthenticationData {
	return v.value
}

func (v *NullableDetailsRequestAuthenticationData) Set(val *DetailsRequestAuthenticationData) {
	v.value = val
	v.isSet = true
}

func (v NullableDetailsRequestAuthenticationData) IsSet() bool {
	return v.isSet
}

func (v *NullableDetailsRequestAuthenticationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetailsRequestAuthenticationData(val *DetailsRequestAuthenticationData) *NullableDetailsRequestAuthenticationData {
	return &NullableDetailsRequestAuthenticationData{value: val, isSet: true}
}

func (v NullableDetailsRequestAuthenticationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetailsRequestAuthenticationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}