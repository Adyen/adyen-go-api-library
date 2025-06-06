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

// checks if the CheckoutDelegatedAuthenticationAction type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CheckoutDelegatedAuthenticationAction{}

// CheckoutDelegatedAuthenticationAction struct for CheckoutDelegatedAuthenticationAction
type CheckoutDelegatedAuthenticationAction struct {
	// A token needed to authorise a payment.
	AuthorisationToken *string `json:"authorisationToken,omitempty"`
	// Encoded payment data.
	PaymentData *string `json:"paymentData,omitempty"`
	// Specifies the payment method.
	PaymentMethodType *string `json:"paymentMethodType,omitempty"`
	// A token to pass to the delegatedAuthentication component.
	Token *string `json:"token,omitempty"`
	// **delegatedAuthentication**
	Type string `json:"type"`
	// Specifies the URL to redirect to.
	Url *string `json:"url,omitempty"`
}

// NewCheckoutDelegatedAuthenticationAction instantiates a new CheckoutDelegatedAuthenticationAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutDelegatedAuthenticationAction(type_ string) *CheckoutDelegatedAuthenticationAction {
	this := CheckoutDelegatedAuthenticationAction{}
	this.Type = type_
	return &this
}

// NewCheckoutDelegatedAuthenticationActionWithDefaults instantiates a new CheckoutDelegatedAuthenticationAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutDelegatedAuthenticationActionWithDefaults() *CheckoutDelegatedAuthenticationAction {
	this := CheckoutDelegatedAuthenticationAction{}
	return &this
}

// GetAuthorisationToken returns the AuthorisationToken field value if set, zero value otherwise.
func (o *CheckoutDelegatedAuthenticationAction) GetAuthorisationToken() string {
	if o == nil || common.IsNil(o.AuthorisationToken) {
		var ret string
		return ret
	}
	return *o.AuthorisationToken
}

// GetAuthorisationTokenOk returns a tuple with the AuthorisationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutDelegatedAuthenticationAction) GetAuthorisationTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.AuthorisationToken) {
		return nil, false
	}
	return o.AuthorisationToken, true
}

// HasAuthorisationToken returns a boolean if a field has been set.
func (o *CheckoutDelegatedAuthenticationAction) HasAuthorisationToken() bool {
	if o != nil && !common.IsNil(o.AuthorisationToken) {
		return true
	}

	return false
}

// SetAuthorisationToken gets a reference to the given string and assigns it to the AuthorisationToken field.
func (o *CheckoutDelegatedAuthenticationAction) SetAuthorisationToken(v string) {
	o.AuthorisationToken = &v
}

// GetPaymentData returns the PaymentData field value if set, zero value otherwise.
func (o *CheckoutDelegatedAuthenticationAction) GetPaymentData() string {
	if o == nil || common.IsNil(o.PaymentData) {
		var ret string
		return ret
	}
	return *o.PaymentData
}

// GetPaymentDataOk returns a tuple with the PaymentData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutDelegatedAuthenticationAction) GetPaymentDataOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentData) {
		return nil, false
	}
	return o.PaymentData, true
}

// HasPaymentData returns a boolean if a field has been set.
func (o *CheckoutDelegatedAuthenticationAction) HasPaymentData() bool {
	if o != nil && !common.IsNil(o.PaymentData) {
		return true
	}

	return false
}

// SetPaymentData gets a reference to the given string and assigns it to the PaymentData field.
func (o *CheckoutDelegatedAuthenticationAction) SetPaymentData(v string) {
	o.PaymentData = &v
}

// GetPaymentMethodType returns the PaymentMethodType field value if set, zero value otherwise.
func (o *CheckoutDelegatedAuthenticationAction) GetPaymentMethodType() string {
	if o == nil || common.IsNil(o.PaymentMethodType) {
		var ret string
		return ret
	}
	return *o.PaymentMethodType
}

// GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutDelegatedAuthenticationAction) GetPaymentMethodTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentMethodType) {
		return nil, false
	}
	return o.PaymentMethodType, true
}

// HasPaymentMethodType returns a boolean if a field has been set.
func (o *CheckoutDelegatedAuthenticationAction) HasPaymentMethodType() bool {
	if o != nil && !common.IsNil(o.PaymentMethodType) {
		return true
	}

	return false
}

// SetPaymentMethodType gets a reference to the given string and assigns it to the PaymentMethodType field.
func (o *CheckoutDelegatedAuthenticationAction) SetPaymentMethodType(v string) {
	o.PaymentMethodType = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CheckoutDelegatedAuthenticationAction) GetToken() string {
	if o == nil || common.IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutDelegatedAuthenticationAction) GetTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CheckoutDelegatedAuthenticationAction) HasToken() bool {
	if o != nil && !common.IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CheckoutDelegatedAuthenticationAction) SetToken(v string) {
	o.Token = &v
}

// GetType returns the Type field value
func (o *CheckoutDelegatedAuthenticationAction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CheckoutDelegatedAuthenticationAction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CheckoutDelegatedAuthenticationAction) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CheckoutDelegatedAuthenticationAction) GetUrl() string {
	if o == nil || common.IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutDelegatedAuthenticationAction) GetUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CheckoutDelegatedAuthenticationAction) HasUrl() bool {
	if o != nil && !common.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CheckoutDelegatedAuthenticationAction) SetUrl(v string) {
	o.Url = &v
}

func (o CheckoutDelegatedAuthenticationAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutDelegatedAuthenticationAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AuthorisationToken) {
		toSerialize["authorisationToken"] = o.AuthorisationToken
	}
	if !common.IsNil(o.PaymentData) {
		toSerialize["paymentData"] = o.PaymentData
	}
	if !common.IsNil(o.PaymentMethodType) {
		toSerialize["paymentMethodType"] = o.PaymentMethodType
	}
	if !common.IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	toSerialize["type"] = o.Type
	if !common.IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableCheckoutDelegatedAuthenticationAction struct {
	value *CheckoutDelegatedAuthenticationAction
	isSet bool
}

func (v NullableCheckoutDelegatedAuthenticationAction) Get() *CheckoutDelegatedAuthenticationAction {
	return v.value
}

func (v *NullableCheckoutDelegatedAuthenticationAction) Set(val *CheckoutDelegatedAuthenticationAction) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutDelegatedAuthenticationAction) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutDelegatedAuthenticationAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutDelegatedAuthenticationAction(val *CheckoutDelegatedAuthenticationAction) *NullableCheckoutDelegatedAuthenticationAction {
	return &NullableCheckoutDelegatedAuthenticationAction{value: val, isSet: true}
}

func (v NullableCheckoutDelegatedAuthenticationAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutDelegatedAuthenticationAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CheckoutDelegatedAuthenticationAction) isValidType() bool {
	var allowedEnumValues = []string{"delegatedAuthentication"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
