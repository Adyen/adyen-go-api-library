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

// checks if the Avs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Avs{}

// Avs struct for Avs
type Avs struct {
	// Indicates whether the shopper is allowed to modify the billing address for the current payment request.
	AddressEditable *bool `json:"addressEditable,omitempty"`
	// Specifies whether the shopper should enter their billing address during checkout.  Allowed values: * yes — Perform AVS checks for every card payment. * automatic — Perform AVS checks only when required to optimize the conversion rate. * no — Do not perform AVS checks.
	Enabled *string `json:"enabled,omitempty"`
}

// NewAvs instantiates a new Avs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvs() *Avs {
	this := Avs{}
	return &this
}

// NewAvsWithDefaults instantiates a new Avs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvsWithDefaults() *Avs {
	this := Avs{}
	return &this
}

// GetAddressEditable returns the AddressEditable field value if set, zero value otherwise.
func (o *Avs) GetAddressEditable() bool {
	if o == nil || IsNil(o.AddressEditable) {
		var ret bool
		return ret
	}
	return *o.AddressEditable
}

// GetAddressEditableOk returns a tuple with the AddressEditable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Avs) GetAddressEditableOk() (*bool, bool) {
	if o == nil || IsNil(o.AddressEditable) {
		return nil, false
	}
	return o.AddressEditable, true
}

// HasAddressEditable returns a boolean if a field has been set.
func (o *Avs) HasAddressEditable() bool {
	if o != nil && !IsNil(o.AddressEditable) {
		return true
	}

	return false
}

// SetAddressEditable gets a reference to the given bool and assigns it to the AddressEditable field.
func (o *Avs) SetAddressEditable(v bool) {
	o.AddressEditable = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Avs) GetEnabled() string {
	if o == nil || IsNil(o.Enabled) {
		var ret string
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Avs) GetEnabledOk() (*string, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Avs) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given string and assigns it to the Enabled field.
func (o *Avs) SetEnabled(v string) {
	o.Enabled = &v
}

func (o Avs) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Avs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressEditable) {
		toSerialize["addressEditable"] = o.AddressEditable
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableAvs struct {
	value *Avs
	isSet bool
}

func (v NullableAvs) Get() *Avs {
	return v.value
}

func (v *NullableAvs) Set(val *Avs) {
	v.value = val
	v.isSet = true
}

func (v NullableAvs) IsSet() bool {
	return v.isSet
}

func (v *NullableAvs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvs(val *Avs) *NullableAvs {
	return &NullableAvs{value: val, isSet: true}
}

func (v NullableAvs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *Avs) isValidEnabled() bool {
	var allowedEnumValues = []string{"yes", "no", "automatic"}
	for _, allowed := range allowedEnumValues {
		if o.GetEnabled() == allowed {
			return true
		}
	}
	return false
}