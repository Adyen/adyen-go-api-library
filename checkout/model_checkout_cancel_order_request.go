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

// checks if the CheckoutCancelOrderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckoutCancelOrderRequest{}

// CheckoutCancelOrderRequest struct for CheckoutCancelOrderRequest
type CheckoutCancelOrderRequest struct {
	// The merchant account identifier that orderData belongs to.
	MerchantAccount string        `json:"merchantAccount"`
	Order           CheckoutOrder `json:"order"`
}

// NewCheckoutCancelOrderRequest instantiates a new CheckoutCancelOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutCancelOrderRequest(merchantAccount string, order CheckoutOrder) *CheckoutCancelOrderRequest {
	this := CheckoutCancelOrderRequest{}
	this.MerchantAccount = merchantAccount
	this.Order = order
	return &this
}

// NewCheckoutCancelOrderRequestWithDefaults instantiates a new CheckoutCancelOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutCancelOrderRequestWithDefaults() *CheckoutCancelOrderRequest {
	this := CheckoutCancelOrderRequest{}
	return &this
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *CheckoutCancelOrderRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *CheckoutCancelOrderRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *CheckoutCancelOrderRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetOrder returns the Order field value
func (o *CheckoutCancelOrderRequest) GetOrder() CheckoutOrder {
	if o == nil {
		var ret CheckoutOrder
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *CheckoutCancelOrderRequest) GetOrderOk() (*CheckoutOrder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *CheckoutCancelOrderRequest) SetOrder(v CheckoutOrder) {
	o.Order = v
}

func (o CheckoutCancelOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutCancelOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merchantAccount"] = o.MerchantAccount
	toSerialize["order"] = o.Order
	return toSerialize, nil
}

type NullableCheckoutCancelOrderRequest struct {
	value *CheckoutCancelOrderRequest
	isSet bool
}

func (v NullableCheckoutCancelOrderRequest) Get() *CheckoutCancelOrderRequest {
	return v.value
}

func (v *NullableCheckoutCancelOrderRequest) Set(val *CheckoutCancelOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutCancelOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutCancelOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutCancelOrderRequest(val *CheckoutCancelOrderRequest) *NullableCheckoutCancelOrderRequest {
	return &NullableCheckoutCancelOrderRequest{value: val, isSet: true}
}

func (v NullableCheckoutCancelOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutCancelOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}