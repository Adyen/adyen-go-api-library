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

// checks if the Configuration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Configuration{}

// Configuration struct for Configuration
type Configuration struct {
	Avs *Avs `json:"avs,omitempty"`
	// Determines whether the cardholder name should be provided or not.  Permitted values: * NONE * OPTIONAL * REQUIRED
	CardHolderName *string             `json:"cardHolderName,omitempty"`
	Installments   *InstallmentsNumber `json:"installments,omitempty"`
	ShopperInput   *ShopperInput       `json:"shopperInput,omitempty"`
}

// NewConfiguration instantiates a new Configuration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfiguration() *Configuration {
	this := Configuration{}
	return &this
}

// NewConfigurationWithDefaults instantiates a new Configuration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationWithDefaults() *Configuration {
	this := Configuration{}
	return &this
}

// GetAvs returns the Avs field value if set, zero value otherwise.
func (o *Configuration) GetAvs() Avs {
	if o == nil || IsNil(o.Avs) {
		var ret Avs
		return ret
	}
	return *o.Avs
}

// GetAvsOk returns a tuple with the Avs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Configuration) GetAvsOk() (*Avs, bool) {
	if o == nil || IsNil(o.Avs) {
		return nil, false
	}
	return o.Avs, true
}

// HasAvs returns a boolean if a field has been set.
func (o *Configuration) HasAvs() bool {
	if o != nil && !IsNil(o.Avs) {
		return true
	}

	return false
}

// SetAvs gets a reference to the given Avs and assigns it to the Avs field.
func (o *Configuration) SetAvs(v Avs) {
	o.Avs = &v
}

// GetCardHolderName returns the CardHolderName field value if set, zero value otherwise.
func (o *Configuration) GetCardHolderName() string {
	if o == nil || IsNil(o.CardHolderName) {
		var ret string
		return ret
	}
	return *o.CardHolderName
}

// GetCardHolderNameOk returns a tuple with the CardHolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Configuration) GetCardHolderNameOk() (*string, bool) {
	if o == nil || IsNil(o.CardHolderName) {
		return nil, false
	}
	return o.CardHolderName, true
}

// HasCardHolderName returns a boolean if a field has been set.
func (o *Configuration) HasCardHolderName() bool {
	if o != nil && !IsNil(o.CardHolderName) {
		return true
	}

	return false
}

// SetCardHolderName gets a reference to the given string and assigns it to the CardHolderName field.
func (o *Configuration) SetCardHolderName(v string) {
	o.CardHolderName = &v
}

// GetInstallments returns the Installments field value if set, zero value otherwise.
func (o *Configuration) GetInstallments() InstallmentsNumber {
	if o == nil || IsNil(o.Installments) {
		var ret InstallmentsNumber
		return ret
	}
	return *o.Installments
}

// GetInstallmentsOk returns a tuple with the Installments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Configuration) GetInstallmentsOk() (*InstallmentsNumber, bool) {
	if o == nil || IsNil(o.Installments) {
		return nil, false
	}
	return o.Installments, true
}

// HasInstallments returns a boolean if a field has been set.
func (o *Configuration) HasInstallments() bool {
	if o != nil && !IsNil(o.Installments) {
		return true
	}

	return false
}

// SetInstallments gets a reference to the given InstallmentsNumber and assigns it to the Installments field.
func (o *Configuration) SetInstallments(v InstallmentsNumber) {
	o.Installments = &v
}

// GetShopperInput returns the ShopperInput field value if set, zero value otherwise.
func (o *Configuration) GetShopperInput() ShopperInput {
	if o == nil || IsNil(o.ShopperInput) {
		var ret ShopperInput
		return ret
	}
	return *o.ShopperInput
}

// GetShopperInputOk returns a tuple with the ShopperInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Configuration) GetShopperInputOk() (*ShopperInput, bool) {
	if o == nil || IsNil(o.ShopperInput) {
		return nil, false
	}
	return o.ShopperInput, true
}

// HasShopperInput returns a boolean if a field has been set.
func (o *Configuration) HasShopperInput() bool {
	if o != nil && !IsNil(o.ShopperInput) {
		return true
	}

	return false
}

// SetShopperInput gets a reference to the given ShopperInput and assigns it to the ShopperInput field.
func (o *Configuration) SetShopperInput(v ShopperInput) {
	o.ShopperInput = &v
}

func (o Configuration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Configuration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Avs) {
		toSerialize["avs"] = o.Avs
	}
	if !IsNil(o.CardHolderName) {
		toSerialize["cardHolderName"] = o.CardHolderName
	}
	if !IsNil(o.Installments) {
		toSerialize["installments"] = o.Installments
	}
	if !IsNil(o.ShopperInput) {
		toSerialize["shopperInput"] = o.ShopperInput
	}
	return toSerialize, nil
}

type NullableConfiguration struct {
	value *Configuration
	isSet bool
}

func (v NullableConfiguration) Get() *Configuration {
	return v.value
}

func (v *NullableConfiguration) Set(val *Configuration) {
	v.value = val
	v.isSet = true
}

func (v NullableConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfiguration(val *Configuration) *NullableConfiguration {
	return &NullableConfiguration{value: val, isSet: true}
}

func (v NullableConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *Configuration) isValidCardHolderName() bool {
	var allowedEnumValues = []string{"NONE", "OPTIONAL", "REQUIRED"}
	for _, allowed := range allowedEnumValues {
		if o.GetCardHolderName() == allowed {
			return true
		}
	}
	return false
}