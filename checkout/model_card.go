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

// checks if the Card type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Card{}

// Card struct for Card
type Card struct {
	// The [card verification code](https://docs.adyen.com/payments-fundamentals/payment-glossary#card-security-code-cvc-cvv-cid) (1-20 characters). Depending on the card brand, it is known also as: * CVV2/CVC2 – length: 3 digits * CID – length: 4 digits > If you are using [Client-Side Encryption](https://docs.adyen.com/classic-integration/cse-integration-ecommerce), the CVC code is present in the encrypted data. You must never post the card details to the server. > This field must be always present in a [one-click payment request](https://docs.adyen.com/classic-integration/recurring-payments). > When this value is returned in a response, it is always empty because it is not stored.
	Cvc *string `json:"cvc,omitempty"`
	// The card expiry month. Format: 2 digits, zero-padded for single digits. For example: * 03 = March * 11 = November
	ExpiryMonth *string `json:"expiryMonth,omitempty"`
	// The card expiry year. Format: 4 digits. For example: 2020
	ExpiryYear string `json:"expiryYear"`
	// The name of the cardholder, as printed on the card.
	HolderName string `json:"holderName"`
	// The issue number of the card (for some UK debit cards only).
	IssueNumber *string `json:"issueNumber,omitempty"`
	// The card number (4-19 characters). Do not use any separators. When this value is returned in a response, only the last 4 digits of the card number are returned.
	Number *string `json:"number,omitempty"`
	// The month component of the start date (for some UK debit cards only).
	StartMonth *string `json:"startMonth,omitempty"`
	// The year component of the start date (for some UK debit cards only).
	StartYear *string `json:"startYear,omitempty"`
}

// NewCard instantiates a new Card object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCard(expiryYear string, holderName string) *Card {
	this := Card{}
	this.ExpiryYear = expiryYear
	this.HolderName = holderName
	return &this
}

// NewCardWithDefaults instantiates a new Card object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardWithDefaults() *Card {
	this := Card{}
	return &this
}

// GetCvc returns the Cvc field value if set, zero value otherwise.
func (o *Card) GetCvc() string {
	if o == nil || IsNil(o.Cvc) {
		var ret string
		return ret
	}
	return *o.Cvc
}

// GetCvcOk returns a tuple with the Cvc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetCvcOk() (*string, bool) {
	if o == nil || IsNil(o.Cvc) {
		return nil, false
	}
	return o.Cvc, true
}

// HasCvc returns a boolean if a field has been set.
func (o *Card) HasCvc() bool {
	if o != nil && !IsNil(o.Cvc) {
		return true
	}

	return false
}

// SetCvc gets a reference to the given string and assigns it to the Cvc field.
func (o *Card) SetCvc(v string) {
	o.Cvc = &v
}

// GetExpiryMonth returns the ExpiryMonth field value if set, zero value otherwise.
func (o *Card) GetExpiryMonth() string {
	if o == nil || IsNil(o.ExpiryMonth) {
		var ret string
		return ret
	}
	return *o.ExpiryMonth
}

// GetExpiryMonthOk returns a tuple with the ExpiryMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetExpiryMonthOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiryMonth) {
		return nil, false
	}
	return o.ExpiryMonth, true
}

// HasExpiryMonth returns a boolean if a field has been set.
func (o *Card) HasExpiryMonth() bool {
	if o != nil && !IsNil(o.ExpiryMonth) {
		return true
	}

	return false
}

// SetExpiryMonth gets a reference to the given string and assigns it to the ExpiryMonth field.
func (o *Card) SetExpiryMonth(v string) {
	o.ExpiryMonth = &v
}

// GetExpiryYear returns the ExpiryYear field value
func (o *Card) GetExpiryYear() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpiryYear
}

// GetExpiryYearOk returns a tuple with the ExpiryYear field value
// and a boolean to check if the value has been set.
func (o *Card) GetExpiryYearOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiryYear, true
}

// SetExpiryYear sets field value
func (o *Card) SetExpiryYear(v string) {
	o.ExpiryYear = v
}

// GetHolderName returns the HolderName field value
func (o *Card) GetHolderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HolderName
}

// GetHolderNameOk returns a tuple with the HolderName field value
// and a boolean to check if the value has been set.
func (o *Card) GetHolderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HolderName, true
}

// SetHolderName sets field value
func (o *Card) SetHolderName(v string) {
	o.HolderName = v
}

// GetIssueNumber returns the IssueNumber field value if set, zero value otherwise.
func (o *Card) GetIssueNumber() string {
	if o == nil || IsNil(o.IssueNumber) {
		var ret string
		return ret
	}
	return *o.IssueNumber
}

// GetIssueNumberOk returns a tuple with the IssueNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetIssueNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IssueNumber) {
		return nil, false
	}
	return o.IssueNumber, true
}

// HasIssueNumber returns a boolean if a field has been set.
func (o *Card) HasIssueNumber() bool {
	if o != nil && !IsNil(o.IssueNumber) {
		return true
	}

	return false
}

// SetIssueNumber gets a reference to the given string and assigns it to the IssueNumber field.
func (o *Card) SetIssueNumber(v string) {
	o.IssueNumber = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *Card) GetNumber() string {
	if o == nil || IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetNumberOk() (*string, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *Card) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *Card) SetNumber(v string) {
	o.Number = &v
}

// GetStartMonth returns the StartMonth field value if set, zero value otherwise.
func (o *Card) GetStartMonth() string {
	if o == nil || IsNil(o.StartMonth) {
		var ret string
		return ret
	}
	return *o.StartMonth
}

// GetStartMonthOk returns a tuple with the StartMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetStartMonthOk() (*string, bool) {
	if o == nil || IsNil(o.StartMonth) {
		return nil, false
	}
	return o.StartMonth, true
}

// HasStartMonth returns a boolean if a field has been set.
func (o *Card) HasStartMonth() bool {
	if o != nil && !IsNil(o.StartMonth) {
		return true
	}

	return false
}

// SetStartMonth gets a reference to the given string and assigns it to the StartMonth field.
func (o *Card) SetStartMonth(v string) {
	o.StartMonth = &v
}

// GetStartYear returns the StartYear field value if set, zero value otherwise.
func (o *Card) GetStartYear() string {
	if o == nil || IsNil(o.StartYear) {
		var ret string
		return ret
	}
	return *o.StartYear
}

// GetStartYearOk returns a tuple with the StartYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetStartYearOk() (*string, bool) {
	if o == nil || IsNil(o.StartYear) {
		return nil, false
	}
	return o.StartYear, true
}

// HasStartYear returns a boolean if a field has been set.
func (o *Card) HasStartYear() bool {
	if o != nil && !IsNil(o.StartYear) {
		return true
	}

	return false
}

// SetStartYear gets a reference to the given string and assigns it to the StartYear field.
func (o *Card) SetStartYear(v string) {
	o.StartYear = &v
}

func (o Card) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Card) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cvc) {
		toSerialize["cvc"] = o.Cvc
	}
	if !IsNil(o.ExpiryMonth) {
		toSerialize["expiryMonth"] = o.ExpiryMonth
	}
	toSerialize["expiryYear"] = o.ExpiryYear
	toSerialize["holderName"] = o.HolderName
	if !IsNil(o.IssueNumber) {
		toSerialize["issueNumber"] = o.IssueNumber
	}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.StartMonth) {
		toSerialize["startMonth"] = o.StartMonth
	}
	if !IsNil(o.StartYear) {
		toSerialize["startYear"] = o.StartYear
	}
	return toSerialize, nil
}

type NullableCard struct {
	value *Card
	isSet bool
}

func (v NullableCard) Get() *Card {
	return v.value
}

func (v *NullableCard) Set(val *Card) {
	v.value = val
	v.isSet = true
}

func (v NullableCard) IsSet() bool {
	return v.isSet
}

func (v *NullableCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCard(val *Card) *NullableCard {
	return &NullableCard{value: val, isSet: true}
}

func (v NullableCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}