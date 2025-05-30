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

// checks if the FundOrigin type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FundOrigin{}

// FundOrigin struct for FundOrigin
type FundOrigin struct {
	BillingAddress *Address `json:"billingAddress,omitempty"`
	// The email address of the person funding the money.
	ShopperEmail *string `json:"shopperEmail,omitempty"`
	ShopperName  *Name   `json:"shopperName,omitempty"`
	// The phone number of the person funding the money.
	TelephoneNumber *string `json:"telephoneNumber,omitempty"`
	// The unique identifier of the wallet where the funds are coming from.
	WalletIdentifier *string `json:"walletIdentifier,omitempty"`
}

// NewFundOrigin instantiates a new FundOrigin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundOrigin() *FundOrigin {
	this := FundOrigin{}
	return &this
}

// NewFundOriginWithDefaults instantiates a new FundOrigin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundOriginWithDefaults() *FundOrigin {
	this := FundOrigin{}
	return &this
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *FundOrigin) GetBillingAddress() Address {
	if o == nil || common.IsNil(o.BillingAddress) {
		var ret Address
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundOrigin) GetBillingAddressOk() (*Address, bool) {
	if o == nil || common.IsNil(o.BillingAddress) {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *FundOrigin) HasBillingAddress() bool {
	if o != nil && !common.IsNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given Address and assigns it to the BillingAddress field.
func (o *FundOrigin) SetBillingAddress(v Address) {
	o.BillingAddress = &v
}

// GetShopperEmail returns the ShopperEmail field value if set, zero value otherwise.
func (o *FundOrigin) GetShopperEmail() string {
	if o == nil || common.IsNil(o.ShopperEmail) {
		var ret string
		return ret
	}
	return *o.ShopperEmail
}

// GetShopperEmailOk returns a tuple with the ShopperEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundOrigin) GetShopperEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperEmail) {
		return nil, false
	}
	return o.ShopperEmail, true
}

// HasShopperEmail returns a boolean if a field has been set.
func (o *FundOrigin) HasShopperEmail() bool {
	if o != nil && !common.IsNil(o.ShopperEmail) {
		return true
	}

	return false
}

// SetShopperEmail gets a reference to the given string and assigns it to the ShopperEmail field.
func (o *FundOrigin) SetShopperEmail(v string) {
	o.ShopperEmail = &v
}

// GetShopperName returns the ShopperName field value if set, zero value otherwise.
func (o *FundOrigin) GetShopperName() Name {
	if o == nil || common.IsNil(o.ShopperName) {
		var ret Name
		return ret
	}
	return *o.ShopperName
}

// GetShopperNameOk returns a tuple with the ShopperName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundOrigin) GetShopperNameOk() (*Name, bool) {
	if o == nil || common.IsNil(o.ShopperName) {
		return nil, false
	}
	return o.ShopperName, true
}

// HasShopperName returns a boolean if a field has been set.
func (o *FundOrigin) HasShopperName() bool {
	if o != nil && !common.IsNil(o.ShopperName) {
		return true
	}

	return false
}

// SetShopperName gets a reference to the given Name and assigns it to the ShopperName field.
func (o *FundOrigin) SetShopperName(v Name) {
	o.ShopperName = &v
}

// GetTelephoneNumber returns the TelephoneNumber field value if set, zero value otherwise.
func (o *FundOrigin) GetTelephoneNumber() string {
	if o == nil || common.IsNil(o.TelephoneNumber) {
		var ret string
		return ret
	}
	return *o.TelephoneNumber
}

// GetTelephoneNumberOk returns a tuple with the TelephoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundOrigin) GetTelephoneNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.TelephoneNumber) {
		return nil, false
	}
	return o.TelephoneNumber, true
}

// HasTelephoneNumber returns a boolean if a field has been set.
func (o *FundOrigin) HasTelephoneNumber() bool {
	if o != nil && !common.IsNil(o.TelephoneNumber) {
		return true
	}

	return false
}

// SetTelephoneNumber gets a reference to the given string and assigns it to the TelephoneNumber field.
func (o *FundOrigin) SetTelephoneNumber(v string) {
	o.TelephoneNumber = &v
}

// GetWalletIdentifier returns the WalletIdentifier field value if set, zero value otherwise.
func (o *FundOrigin) GetWalletIdentifier() string {
	if o == nil || common.IsNil(o.WalletIdentifier) {
		var ret string
		return ret
	}
	return *o.WalletIdentifier
}

// GetWalletIdentifierOk returns a tuple with the WalletIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundOrigin) GetWalletIdentifierOk() (*string, bool) {
	if o == nil || common.IsNil(o.WalletIdentifier) {
		return nil, false
	}
	return o.WalletIdentifier, true
}

// HasWalletIdentifier returns a boolean if a field has been set.
func (o *FundOrigin) HasWalletIdentifier() bool {
	if o != nil && !common.IsNil(o.WalletIdentifier) {
		return true
	}

	return false
}

// SetWalletIdentifier gets a reference to the given string and assigns it to the WalletIdentifier field.
func (o *FundOrigin) SetWalletIdentifier(v string) {
	o.WalletIdentifier = &v
}

func (o FundOrigin) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FundOrigin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BillingAddress) {
		toSerialize["billingAddress"] = o.BillingAddress
	}
	if !common.IsNil(o.ShopperEmail) {
		toSerialize["shopperEmail"] = o.ShopperEmail
	}
	if !common.IsNil(o.ShopperName) {
		toSerialize["shopperName"] = o.ShopperName
	}
	if !common.IsNil(o.TelephoneNumber) {
		toSerialize["telephoneNumber"] = o.TelephoneNumber
	}
	if !common.IsNil(o.WalletIdentifier) {
		toSerialize["walletIdentifier"] = o.WalletIdentifier
	}
	return toSerialize, nil
}

type NullableFundOrigin struct {
	value *FundOrigin
	isSet bool
}

func (v NullableFundOrigin) Get() *FundOrigin {
	return v.value
}

func (v *NullableFundOrigin) Set(val *FundOrigin) {
	v.value = val
	v.isSet = true
}

func (v NullableFundOrigin) IsSet() bool {
	return v.isSet
}

func (v *NullableFundOrigin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundOrigin(val *FundOrigin) *NullableFundOrigin {
	return &NullableFundOrigin{value: val, isSet: true}
}

func (v NullableFundOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundOrigin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
