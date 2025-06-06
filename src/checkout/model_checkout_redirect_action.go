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

// checks if the CheckoutRedirectAction type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CheckoutRedirectAction{}

// CheckoutRedirectAction struct for CheckoutRedirectAction
type CheckoutRedirectAction struct {
	// When the redirect URL must be accessed via POST, use this data to post to the redirect URL.
	Data *map[string]string `json:"data,omitempty"`
	// Specifies the HTTP method, for example GET or POST.
	Method *string `json:"method,omitempty"`
	// Specifies the payment method.
	PaymentMethodType *string `json:"paymentMethodType,omitempty"`
	// **redirect**
	Type string `json:"type"`
	// Specifies the URL to redirect to.
	Url *string `json:"url,omitempty"`
}

// NewCheckoutRedirectAction instantiates a new CheckoutRedirectAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutRedirectAction(type_ string) *CheckoutRedirectAction {
	this := CheckoutRedirectAction{}
	this.Type = type_
	return &this
}

// NewCheckoutRedirectActionWithDefaults instantiates a new CheckoutRedirectAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutRedirectActionWithDefaults() *CheckoutRedirectAction {
	this := CheckoutRedirectAction{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CheckoutRedirectAction) GetData() map[string]string {
	if o == nil || common.IsNil(o.Data) {
		var ret map[string]string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutRedirectAction) GetDataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CheckoutRedirectAction) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]string and assigns it to the Data field.
func (o *CheckoutRedirectAction) SetData(v map[string]string) {
	o.Data = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *CheckoutRedirectAction) GetMethod() string {
	if o == nil || common.IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutRedirectAction) GetMethodOk() (*string, bool) {
	if o == nil || common.IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *CheckoutRedirectAction) HasMethod() bool {
	if o != nil && !common.IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *CheckoutRedirectAction) SetMethod(v string) {
	o.Method = &v
}

// GetPaymentMethodType returns the PaymentMethodType field value if set, zero value otherwise.
func (o *CheckoutRedirectAction) GetPaymentMethodType() string {
	if o == nil || common.IsNil(o.PaymentMethodType) {
		var ret string
		return ret
	}
	return *o.PaymentMethodType
}

// GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutRedirectAction) GetPaymentMethodTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentMethodType) {
		return nil, false
	}
	return o.PaymentMethodType, true
}

// HasPaymentMethodType returns a boolean if a field has been set.
func (o *CheckoutRedirectAction) HasPaymentMethodType() bool {
	if o != nil && !common.IsNil(o.PaymentMethodType) {
		return true
	}

	return false
}

// SetPaymentMethodType gets a reference to the given string and assigns it to the PaymentMethodType field.
func (o *CheckoutRedirectAction) SetPaymentMethodType(v string) {
	o.PaymentMethodType = &v
}

// GetType returns the Type field value
func (o *CheckoutRedirectAction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CheckoutRedirectAction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CheckoutRedirectAction) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CheckoutRedirectAction) GetUrl() string {
	if o == nil || common.IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutRedirectAction) GetUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CheckoutRedirectAction) HasUrl() bool {
	if o != nil && !common.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CheckoutRedirectAction) SetUrl(v string) {
	o.Url = &v
}

func (o CheckoutRedirectAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutRedirectAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !common.IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !common.IsNil(o.PaymentMethodType) {
		toSerialize["paymentMethodType"] = o.PaymentMethodType
	}
	toSerialize["type"] = o.Type
	if !common.IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableCheckoutRedirectAction struct {
	value *CheckoutRedirectAction
	isSet bool
}

func (v NullableCheckoutRedirectAction) Get() *CheckoutRedirectAction {
	return v.value
}

func (v *NullableCheckoutRedirectAction) Set(val *CheckoutRedirectAction) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutRedirectAction) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutRedirectAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutRedirectAction(val *CheckoutRedirectAction) *NullableCheckoutRedirectAction {
	return &NullableCheckoutRedirectAction{value: val, isSet: true}
}

func (v NullableCheckoutRedirectAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutRedirectAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CheckoutRedirectAction) isValidType() bool {
	var allowedEnumValues = []string{"redirect"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
