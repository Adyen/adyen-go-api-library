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

// checks if the CheckoutUtilityResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckoutUtilityResponse{}

// CheckoutUtilityResponse struct for CheckoutUtilityResponse
type CheckoutUtilityResponse struct {
	// The list of origin keys for all requested domains. For each list item, the key is the domain and the value is the origin key.
	OriginKeys *map[string]string `json:"originKeys,omitempty"`
}

// NewCheckoutUtilityResponse instantiates a new CheckoutUtilityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutUtilityResponse() *CheckoutUtilityResponse {
	this := CheckoutUtilityResponse{}
	return &this
}

// NewCheckoutUtilityResponseWithDefaults instantiates a new CheckoutUtilityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutUtilityResponseWithDefaults() *CheckoutUtilityResponse {
	this := CheckoutUtilityResponse{}
	return &this
}

// GetOriginKeys returns the OriginKeys field value if set, zero value otherwise.
func (o *CheckoutUtilityResponse) GetOriginKeys() map[string]string {
	if o == nil || IsNil(o.OriginKeys) {
		var ret map[string]string
		return ret
	}
	return *o.OriginKeys
}

// GetOriginKeysOk returns a tuple with the OriginKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutUtilityResponse) GetOriginKeysOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.OriginKeys) {
		return nil, false
	}
	return o.OriginKeys, true
}

// HasOriginKeys returns a boolean if a field has been set.
func (o *CheckoutUtilityResponse) HasOriginKeys() bool {
	if o != nil && !IsNil(o.OriginKeys) {
		return true
	}

	return false
}

// SetOriginKeys gets a reference to the given map[string]string and assigns it to the OriginKeys field.
func (o *CheckoutUtilityResponse) SetOriginKeys(v map[string]string) {
	o.OriginKeys = &v
}

func (o CheckoutUtilityResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutUtilityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OriginKeys) {
		toSerialize["originKeys"] = o.OriginKeys
	}
	return toSerialize, nil
}

type NullableCheckoutUtilityResponse struct {
	value *CheckoutUtilityResponse
	isSet bool
}

func (v NullableCheckoutUtilityResponse) Get() *CheckoutUtilityResponse {
	return v.value
}

func (v *NullableCheckoutUtilityResponse) Set(val *CheckoutUtilityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutUtilityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutUtilityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutUtilityResponse(val *CheckoutUtilityResponse) *NullableCheckoutUtilityResponse {
	return &NullableCheckoutUtilityResponse{value: val, isSet: true}
}

func (v NullableCheckoutUtilityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutUtilityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}