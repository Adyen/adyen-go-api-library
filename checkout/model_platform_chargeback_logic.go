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

// checks if the PlatformChargebackLogic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlatformChargebackLogic{}

// PlatformChargebackLogic struct for PlatformChargebackLogic
type PlatformChargebackLogic struct {
	Behavior      *string `json:"behavior,omitempty"`
	TargetAccount *string `json:"targetAccount,omitempty"`
}

// NewPlatformChargebackLogic instantiates a new PlatformChargebackLogic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformChargebackLogic() *PlatformChargebackLogic {
	this := PlatformChargebackLogic{}
	return &this
}

// NewPlatformChargebackLogicWithDefaults instantiates a new PlatformChargebackLogic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformChargebackLogicWithDefaults() *PlatformChargebackLogic {
	this := PlatformChargebackLogic{}
	return &this
}

// GetBehavior returns the Behavior field value if set, zero value otherwise.
func (o *PlatformChargebackLogic) GetBehavior() string {
	if o == nil || IsNil(o.Behavior) {
		var ret string
		return ret
	}
	return *o.Behavior
}

// GetBehaviorOk returns a tuple with the Behavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformChargebackLogic) GetBehaviorOk() (*string, bool) {
	if o == nil || IsNil(o.Behavior) {
		return nil, false
	}
	return o.Behavior, true
}

// HasBehavior returns a boolean if a field has been set.
func (o *PlatformChargebackLogic) HasBehavior() bool {
	if o != nil && !IsNil(o.Behavior) {
		return true
	}

	return false
}

// SetBehavior gets a reference to the given string and assigns it to the Behavior field.
func (o *PlatformChargebackLogic) SetBehavior(v string) {
	o.Behavior = &v
}

// GetTargetAccount returns the TargetAccount field value if set, zero value otherwise.
func (o *PlatformChargebackLogic) GetTargetAccount() string {
	if o == nil || IsNil(o.TargetAccount) {
		var ret string
		return ret
	}
	return *o.TargetAccount
}

// GetTargetAccountOk returns a tuple with the TargetAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformChargebackLogic) GetTargetAccountOk() (*string, bool) {
	if o == nil || IsNil(o.TargetAccount) {
		return nil, false
	}
	return o.TargetAccount, true
}

// HasTargetAccount returns a boolean if a field has been set.
func (o *PlatformChargebackLogic) HasTargetAccount() bool {
	if o != nil && !IsNil(o.TargetAccount) {
		return true
	}

	return false
}

// SetTargetAccount gets a reference to the given string and assigns it to the TargetAccount field.
func (o *PlatformChargebackLogic) SetTargetAccount(v string) {
	o.TargetAccount = &v
}

func (o PlatformChargebackLogic) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlatformChargebackLogic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Behavior) {
		toSerialize["behavior"] = o.Behavior
	}
	if !IsNil(o.TargetAccount) {
		toSerialize["targetAccount"] = o.TargetAccount
	}
	return toSerialize, nil
}

type NullablePlatformChargebackLogic struct {
	value *PlatformChargebackLogic
	isSet bool
}

func (v NullablePlatformChargebackLogic) Get() *PlatformChargebackLogic {
	return v.value
}

func (v *NullablePlatformChargebackLogic) Set(val *PlatformChargebackLogic) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformChargebackLogic) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformChargebackLogic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformChargebackLogic(val *PlatformChargebackLogic) *NullablePlatformChargebackLogic {
	return &NullablePlatformChargebackLogic{value: val, isSet: true}
}

func (v NullablePlatformChargebackLogic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformChargebackLogic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PlatformChargebackLogic) isValidBehavior() bool {
	var allowedEnumValues = []string{"deductAccordingToSplitRatio", "deductFromLiableAccount", "deductFromOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetBehavior() == allowed {
			return true
		}
	}
	return false
}