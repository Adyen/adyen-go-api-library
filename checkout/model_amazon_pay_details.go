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

// checks if the AmazonPayDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmazonPayDetails{}

// AmazonPayDetails struct for AmazonPayDetails
type AmazonPayDetails struct {
	// This is the `amazonPayToken` that you obtained from the [Get Checkout Session](https://amazon-pay-acquirer-guide.s3-eu-west-1.amazonaws.com/v1/amazon-pay-api-v2/checkout-session.html#get-checkout-session) response.
	AmazonPayToken *string `json:"amazonPayToken,omitempty"`
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// **amazonpay**
	Type *string `json:"type,omitempty"`
}

// NewAmazonPayDetails instantiates a new AmazonPayDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonPayDetails() *AmazonPayDetails {
	this := AmazonPayDetails{}
	var type_ string = "amazonpay"
	this.Type = &type_
	return &this
}

// NewAmazonPayDetailsWithDefaults instantiates a new AmazonPayDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonPayDetailsWithDefaults() *AmazonPayDetails {
	this := AmazonPayDetails{}
	var type_ string = "amazonpay"
	this.Type = &type_
	return &this
}

// GetAmazonPayToken returns the AmazonPayToken field value if set, zero value otherwise.
func (o *AmazonPayDetails) GetAmazonPayToken() string {
	if o == nil || IsNil(o.AmazonPayToken) {
		var ret string
		return ret
	}
	return *o.AmazonPayToken
}

// GetAmazonPayTokenOk returns a tuple with the AmazonPayToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonPayDetails) GetAmazonPayTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AmazonPayToken) {
		return nil, false
	}
	return o.AmazonPayToken, true
}

// HasAmazonPayToken returns a boolean if a field has been set.
func (o *AmazonPayDetails) HasAmazonPayToken() bool {
	if o != nil && !IsNil(o.AmazonPayToken) {
		return true
	}

	return false
}

// SetAmazonPayToken gets a reference to the given string and assigns it to the AmazonPayToken field.
func (o *AmazonPayDetails) SetAmazonPayToken(v string) {
	o.AmazonPayToken = &v
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *AmazonPayDetails) GetCheckoutAttemptId() string {
	if o == nil || IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonPayDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *AmazonPayDetails) HasCheckoutAttemptId() bool {
	if o != nil && !IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *AmazonPayDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AmazonPayDetails) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonPayDetails) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AmazonPayDetails) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AmazonPayDetails) SetType(v string) {
	o.Type = &v
}

func (o AmazonPayDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmazonPayDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AmazonPayToken) {
		toSerialize["amazonPayToken"] = o.AmazonPayToken
	}
	if !IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAmazonPayDetails struct {
	value *AmazonPayDetails
	isSet bool
}

func (v NullableAmazonPayDetails) Get() *AmazonPayDetails {
	return v.value
}

func (v *NullableAmazonPayDetails) Set(val *AmazonPayDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonPayDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonPayDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonPayDetails(val *AmazonPayDetails) *NullableAmazonPayDetails {
	return &NullableAmazonPayDetails{value: val, isSet: true}
}

func (v NullableAmazonPayDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonPayDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *AmazonPayDetails) isValidType() bool {
	var allowedEnumValues = []string{"amazonpay"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}