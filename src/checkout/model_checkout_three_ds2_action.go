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

// checks if the CheckoutThreeDS2Action type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CheckoutThreeDS2Action{}

// CheckoutThreeDS2Action struct for CheckoutThreeDS2Action
type CheckoutThreeDS2Action struct {
	// A token needed to authorise a payment.
	AuthorisationToken *string `json:"authorisationToken,omitempty"`
	// Encoded payment data.
	PaymentData *string `json:"paymentData,omitempty"`
	// Specifies the payment method.
	PaymentMethodType *string `json:"paymentMethodType,omitempty"`
	// A subtype of the token.
	Subtype *string `json:"subtype,omitempty"`
	// A token to pass to the 3DS2 Component to get the fingerprint.
	Token *string `json:"token,omitempty"`
	// **threeDS2**
	Type string `json:"type"`
	// Specifies the URL to redirect to.
	Url *string `json:"url,omitempty"`
}

// NewCheckoutThreeDS2Action instantiates a new CheckoutThreeDS2Action object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutThreeDS2Action(type_ string) *CheckoutThreeDS2Action {
	this := CheckoutThreeDS2Action{}
	this.Type = type_
	return &this
}

// NewCheckoutThreeDS2ActionWithDefaults instantiates a new CheckoutThreeDS2Action object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutThreeDS2ActionWithDefaults() *CheckoutThreeDS2Action {
	this := CheckoutThreeDS2Action{}
	return &this
}

// GetAuthorisationToken returns the AuthorisationToken field value if set, zero value otherwise.
func (o *CheckoutThreeDS2Action) GetAuthorisationToken() string {
	if o == nil || common.IsNil(o.AuthorisationToken) {
		var ret string
		return ret
	}
	return *o.AuthorisationToken
}

// GetAuthorisationTokenOk returns a tuple with the AuthorisationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutThreeDS2Action) GetAuthorisationTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.AuthorisationToken) {
		return nil, false
	}
	return o.AuthorisationToken, true
}

// HasAuthorisationToken returns a boolean if a field has been set.
func (o *CheckoutThreeDS2Action) HasAuthorisationToken() bool {
	if o != nil && !common.IsNil(o.AuthorisationToken) {
		return true
	}

	return false
}

// SetAuthorisationToken gets a reference to the given string and assigns it to the AuthorisationToken field.
func (o *CheckoutThreeDS2Action) SetAuthorisationToken(v string) {
	o.AuthorisationToken = &v
}

// GetPaymentData returns the PaymentData field value if set, zero value otherwise.
func (o *CheckoutThreeDS2Action) GetPaymentData() string {
	if o == nil || common.IsNil(o.PaymentData) {
		var ret string
		return ret
	}
	return *o.PaymentData
}

// GetPaymentDataOk returns a tuple with the PaymentData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutThreeDS2Action) GetPaymentDataOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentData) {
		return nil, false
	}
	return o.PaymentData, true
}

// HasPaymentData returns a boolean if a field has been set.
func (o *CheckoutThreeDS2Action) HasPaymentData() bool {
	if o != nil && !common.IsNil(o.PaymentData) {
		return true
	}

	return false
}

// SetPaymentData gets a reference to the given string and assigns it to the PaymentData field.
func (o *CheckoutThreeDS2Action) SetPaymentData(v string) {
	o.PaymentData = &v
}

// GetPaymentMethodType returns the PaymentMethodType field value if set, zero value otherwise.
func (o *CheckoutThreeDS2Action) GetPaymentMethodType() string {
	if o == nil || common.IsNil(o.PaymentMethodType) {
		var ret string
		return ret
	}
	return *o.PaymentMethodType
}

// GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutThreeDS2Action) GetPaymentMethodTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentMethodType) {
		return nil, false
	}
	return o.PaymentMethodType, true
}

// HasPaymentMethodType returns a boolean if a field has been set.
func (o *CheckoutThreeDS2Action) HasPaymentMethodType() bool {
	if o != nil && !common.IsNil(o.PaymentMethodType) {
		return true
	}

	return false
}

// SetPaymentMethodType gets a reference to the given string and assigns it to the PaymentMethodType field.
func (o *CheckoutThreeDS2Action) SetPaymentMethodType(v string) {
	o.PaymentMethodType = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *CheckoutThreeDS2Action) GetSubtype() string {
	if o == nil || common.IsNil(o.Subtype) {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutThreeDS2Action) GetSubtypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Subtype) {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *CheckoutThreeDS2Action) HasSubtype() bool {
	if o != nil && !common.IsNil(o.Subtype) {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *CheckoutThreeDS2Action) SetSubtype(v string) {
	o.Subtype = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CheckoutThreeDS2Action) GetToken() string {
	if o == nil || common.IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutThreeDS2Action) GetTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CheckoutThreeDS2Action) HasToken() bool {
	if o != nil && !common.IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CheckoutThreeDS2Action) SetToken(v string) {
	o.Token = &v
}

// GetType returns the Type field value
func (o *CheckoutThreeDS2Action) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CheckoutThreeDS2Action) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CheckoutThreeDS2Action) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CheckoutThreeDS2Action) GetUrl() string {
	if o == nil || common.IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutThreeDS2Action) GetUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CheckoutThreeDS2Action) HasUrl() bool {
	if o != nil && !common.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CheckoutThreeDS2Action) SetUrl(v string) {
	o.Url = &v
}

func (o CheckoutThreeDS2Action) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutThreeDS2Action) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Subtype) {
		toSerialize["subtype"] = o.Subtype
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

type NullableCheckoutThreeDS2Action struct {
	value *CheckoutThreeDS2Action
	isSet bool
}

func (v NullableCheckoutThreeDS2Action) Get() *CheckoutThreeDS2Action {
	return v.value
}

func (v *NullableCheckoutThreeDS2Action) Set(val *CheckoutThreeDS2Action) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutThreeDS2Action) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutThreeDS2Action) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutThreeDS2Action(val *CheckoutThreeDS2Action) *NullableCheckoutThreeDS2Action {
	return &NullableCheckoutThreeDS2Action{value: val, isSet: true}
}

func (v NullableCheckoutThreeDS2Action) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutThreeDS2Action) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CheckoutThreeDS2Action) isValidType() bool {
	var allowedEnumValues = []string{"threeDS2"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
