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

// checks if the PaymentDetailsRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentDetailsRequest{}

// PaymentDetailsRequest struct for PaymentDetailsRequest
type PaymentDetailsRequest struct {
	AuthenticationData *DetailsRequestAuthenticationData `json:"authenticationData,omitempty"`
	Details            PaymentCompletionDetails          `json:"details"`
	// Encoded payment data. For [authorizing a payment after using 3D Secure 2 Authentication-only](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only/#authorise-the-payment-with-adyen):  If you received `resultCode`: **AuthenticationNotRequired** in the `/payments` response, use the `threeDSPaymentData` from the same response.  If you received `resultCode`: **AuthenticationFinished** in the `/payments` response, use the `action.paymentData` from the same response.
	PaymentData *string `json:"paymentData,omitempty"`
	// Change the `authenticationOnly` indicator originally set in the `/payments` request. Only needs to be set if you want to modify the value set previously.
	// Deprecated since Adyen Checkout API v69
	// Use `authenticationData.authenticationOnly` instead.
	ThreeDSAuthenticationOnly *bool `json:"threeDSAuthenticationOnly,omitempty"`
}

// NewPaymentDetailsRequest instantiates a new PaymentDetailsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentDetailsRequest(details PaymentCompletionDetails) *PaymentDetailsRequest {
	this := PaymentDetailsRequest{}
	this.Details = details
	return &this
}

// NewPaymentDetailsRequestWithDefaults instantiates a new PaymentDetailsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentDetailsRequestWithDefaults() *PaymentDetailsRequest {
	this := PaymentDetailsRequest{}
	return &this
}

// GetAuthenticationData returns the AuthenticationData field value if set, zero value otherwise.
func (o *PaymentDetailsRequest) GetAuthenticationData() DetailsRequestAuthenticationData {
	if o == nil || common.IsNil(o.AuthenticationData) {
		var ret DetailsRequestAuthenticationData
		return ret
	}
	return *o.AuthenticationData
}

// GetAuthenticationDataOk returns a tuple with the AuthenticationData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetailsRequest) GetAuthenticationDataOk() (*DetailsRequestAuthenticationData, bool) {
	if o == nil || common.IsNil(o.AuthenticationData) {
		return nil, false
	}
	return o.AuthenticationData, true
}

// HasAuthenticationData returns a boolean if a field has been set.
func (o *PaymentDetailsRequest) HasAuthenticationData() bool {
	if o != nil && !common.IsNil(o.AuthenticationData) {
		return true
	}

	return false
}

// SetAuthenticationData gets a reference to the given DetailsRequestAuthenticationData and assigns it to the AuthenticationData field.
func (o *PaymentDetailsRequest) SetAuthenticationData(v DetailsRequestAuthenticationData) {
	o.AuthenticationData = &v
}

// GetDetails returns the Details field value
func (o *PaymentDetailsRequest) GetDetails() PaymentCompletionDetails {
	if o == nil {
		var ret PaymentCompletionDetails
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *PaymentDetailsRequest) GetDetailsOk() (*PaymentCompletionDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *PaymentDetailsRequest) SetDetails(v PaymentCompletionDetails) {
	o.Details = v
}

// GetPaymentData returns the PaymentData field value if set, zero value otherwise.
func (o *PaymentDetailsRequest) GetPaymentData() string {
	if o == nil || common.IsNil(o.PaymentData) {
		var ret string
		return ret
	}
	return *o.PaymentData
}

// GetPaymentDataOk returns a tuple with the PaymentData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetailsRequest) GetPaymentDataOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentData) {
		return nil, false
	}
	return o.PaymentData, true
}

// HasPaymentData returns a boolean if a field has been set.
func (o *PaymentDetailsRequest) HasPaymentData() bool {
	if o != nil && !common.IsNil(o.PaymentData) {
		return true
	}

	return false
}

// SetPaymentData gets a reference to the given string and assigns it to the PaymentData field.
func (o *PaymentDetailsRequest) SetPaymentData(v string) {
	o.PaymentData = &v
}

// GetThreeDSAuthenticationOnly returns the ThreeDSAuthenticationOnly field value if set, zero value otherwise.
// Deprecated since Adyen Checkout API v69
// Use `authenticationData.authenticationOnly` instead.
func (o *PaymentDetailsRequest) GetThreeDSAuthenticationOnly() bool {
	if o == nil || common.IsNil(o.ThreeDSAuthenticationOnly) {
		var ret bool
		return ret
	}
	return *o.ThreeDSAuthenticationOnly
}

// GetThreeDSAuthenticationOnlyOk returns a tuple with the ThreeDSAuthenticationOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated since Adyen Checkout API v69
// Use `authenticationData.authenticationOnly` instead.
func (o *PaymentDetailsRequest) GetThreeDSAuthenticationOnlyOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ThreeDSAuthenticationOnly) {
		return nil, false
	}
	return o.ThreeDSAuthenticationOnly, true
}

// HasThreeDSAuthenticationOnly returns a boolean if a field has been set.
func (o *PaymentDetailsRequest) HasThreeDSAuthenticationOnly() bool {
	if o != nil && !common.IsNil(o.ThreeDSAuthenticationOnly) {
		return true
	}

	return false
}

// SetThreeDSAuthenticationOnly gets a reference to the given bool and assigns it to the ThreeDSAuthenticationOnly field.
// Deprecated since Adyen Checkout API v69
// Use `authenticationData.authenticationOnly` instead.
func (o *PaymentDetailsRequest) SetThreeDSAuthenticationOnly(v bool) {
	o.ThreeDSAuthenticationOnly = &v
}

func (o PaymentDetailsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentDetailsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AuthenticationData) {
		toSerialize["authenticationData"] = o.AuthenticationData
	}
	toSerialize["details"] = o.Details
	if !common.IsNil(o.PaymentData) {
		toSerialize["paymentData"] = o.PaymentData
	}
	if !common.IsNil(o.ThreeDSAuthenticationOnly) {
		toSerialize["threeDSAuthenticationOnly"] = o.ThreeDSAuthenticationOnly
	}
	return toSerialize, nil
}

type NullablePaymentDetailsRequest struct {
	value *PaymentDetailsRequest
	isSet bool
}

func (v NullablePaymentDetailsRequest) Get() *PaymentDetailsRequest {
	return v.value
}

func (v *NullablePaymentDetailsRequest) Set(val *PaymentDetailsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentDetailsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentDetailsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentDetailsRequest(val *PaymentDetailsRequest) *NullablePaymentDetailsRequest {
	return &NullablePaymentDetailsRequest{value: val, isSet: true}
}

func (v NullablePaymentDetailsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentDetailsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
