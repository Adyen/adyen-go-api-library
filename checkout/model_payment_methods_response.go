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

// checks if the PaymentMethodsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodsResponse{}

// PaymentMethodsResponse struct for PaymentMethodsResponse
type PaymentMethodsResponse struct {
	// Detailed list of payment methods required to generate payment forms.
	PaymentMethods []PaymentMethod `json:"paymentMethods,omitempty"`
	// List of all stored payment methods.
	StoredPaymentMethods []StoredPaymentMethod `json:"storedPaymentMethods,omitempty"`
}

// NewPaymentMethodsResponse instantiates a new PaymentMethodsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodsResponse() *PaymentMethodsResponse {
	this := PaymentMethodsResponse{}
	return &this
}

// NewPaymentMethodsResponseWithDefaults instantiates a new PaymentMethodsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodsResponseWithDefaults() *PaymentMethodsResponse {
	this := PaymentMethodsResponse{}
	return &this
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *PaymentMethodsResponse) GetPaymentMethods() []PaymentMethod {
	if o == nil || IsNil(o.PaymentMethods) {
		var ret []PaymentMethod
		return ret
	}
	return o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodsResponse) GetPaymentMethodsOk() ([]PaymentMethod, bool) {
	if o == nil || IsNil(o.PaymentMethods) {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *PaymentMethodsResponse) HasPaymentMethods() bool {
	if o != nil && !IsNil(o.PaymentMethods) {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given []PaymentMethod and assigns it to the PaymentMethods field.
func (o *PaymentMethodsResponse) SetPaymentMethods(v []PaymentMethod) {
	o.PaymentMethods = v
}

// GetStoredPaymentMethods returns the StoredPaymentMethods field value if set, zero value otherwise.
func (o *PaymentMethodsResponse) GetStoredPaymentMethods() []StoredPaymentMethod {
	if o == nil || IsNil(o.StoredPaymentMethods) {
		var ret []StoredPaymentMethod
		return ret
	}
	return o.StoredPaymentMethods
}

// GetStoredPaymentMethodsOk returns a tuple with the StoredPaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodsResponse) GetStoredPaymentMethodsOk() ([]StoredPaymentMethod, bool) {
	if o == nil || IsNil(o.StoredPaymentMethods) {
		return nil, false
	}
	return o.StoredPaymentMethods, true
}

// HasStoredPaymentMethods returns a boolean if a field has been set.
func (o *PaymentMethodsResponse) HasStoredPaymentMethods() bool {
	if o != nil && !IsNil(o.StoredPaymentMethods) {
		return true
	}

	return false
}

// SetStoredPaymentMethods gets a reference to the given []StoredPaymentMethod and assigns it to the StoredPaymentMethods field.
func (o *PaymentMethodsResponse) SetStoredPaymentMethods(v []StoredPaymentMethod) {
	o.StoredPaymentMethods = v
}

func (o PaymentMethodsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentMethods) {
		toSerialize["paymentMethods"] = o.PaymentMethods
	}
	if !IsNil(o.StoredPaymentMethods) {
		toSerialize["storedPaymentMethods"] = o.StoredPaymentMethods
	}
	return toSerialize, nil
}

type NullablePaymentMethodsResponse struct {
	value *PaymentMethodsResponse
	isSet bool
}

func (v NullablePaymentMethodsResponse) Get() *PaymentMethodsResponse {
	return v.value
}

func (v *NullablePaymentMethodsResponse) Set(val *PaymentMethodsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodsResponse(val *PaymentMethodsResponse) *NullablePaymentMethodsResponse {
	return &NullablePaymentMethodsResponse{value: val, isSet: true}
}

func (v NullablePaymentMethodsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}