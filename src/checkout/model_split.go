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

// checks if the Split type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Split{}

// Split struct for Split
type Split struct {
	// The unique identifier of the account to which the split amount is booked. Required if `type` is **MarketPlace** or **BalanceAccount**.  * [Classic Platforms integration](https://docs.adyen.com/classic-platforms): The [`accountCode`](https://docs.adyen.com/api-explorer/Account/latest/post/updateAccount#request-accountCode) of the account to which the split amount is booked. * [Balance Platform](https://docs.adyen.com/adyen-for-platforms-model): The [`balanceAccountId`](https://docs.adyen.com/api-explorer/balanceplatform/latest/get/balanceAccounts/_id_#path-id) of the account to which the split amount is booked.
	Account *string      `json:"account,omitempty"`
	Amount  *SplitAmount `json:"amount,omitempty"`
	// Your description for the split item.
	Description *string `json:"description,omitempty"`
	// Your unique reference for the part of the payment booked to the specified `account`.  This is required if `type` is **MarketPlace** ([Classic Platforms integration](https://docs.adyen.com/classic-platforms)) or **BalanceAccount** ([Balance Platform](https://docs.adyen.com/adyen-for-platforms-model)).  For the other types, we also recommend providing a **unique** reference so you can reconcile the split and the associated payment in the transaction overview and in the reports.
	Reference *string `json:"reference,omitempty"`
	// The part of the payment you want to book to the specified `account`.  Possible values for the [Balance Platform](https://docs.adyen.com/adyen-for-platforms-model): * **BalanceAccount**: books part of the payment (specified in `amount`) to the specified `account`. * Transaction fees types that you can book to the specified `account`:    * **AcquiringFees**: the aggregated amount of the interchange and scheme fees.    * **PaymentFee**: the aggregated amount of all transaction fees.    * **AdyenFees**: the aggregated amount of Adyen's commission and markup fees.    * **AdyenCommission**: the transaction fees due to Adyen under [blended rates](https://www.adyen.com/knowledge-hub/interchange-fees-explained).    * **AdyenMarkup**: the transaction fees due to Adyen under [Interchange ++ pricing](https://www.adyen.com/knowledge-hub/interchange-fees-explained).    * **Interchange**: the fees paid to the issuer for each payment made with the card network.    * **SchemeFee**: the fees paid to the card scheme for using their network.  * **Commission**: your platform's commission on the payment (specified in `amount`), booked to your liable balance account. * **Remainder**: the amount left over after a currency conversion, booked to the specified `account`. * **TopUp**: allows you and your users to top up balance accounts using direct debit, card payments, or other payment methods. * **VAT**: the value-added tax charged on the payment, booked to your platforms liable balance account. * **Commission**: your platform's commission (specified in `amount`) on the payment, booked to your liable balance account. * **Default**: in very specific use cases, allows you to book the specified `amount` to the specified `account`. For more information, contact Adyen support.  Possible values for the [Classic Platforms integration](https://docs.adyen.com/classic-platforms): **Commission**, **Default**, **MarketPlace**, **PaymentFee**, **VAT**.
	Type string `json:"type"`
}

// NewSplit instantiates a new Split object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSplit(type_ string) *Split {
	this := Split{}
	this.Type = type_
	return &this
}

// NewSplitWithDefaults instantiates a new Split object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSplitWithDefaults() *Split {
	this := Split{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *Split) GetAccount() string {
	if o == nil || common.IsNil(o.Account) {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Split) GetAccountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *Split) HasAccount() bool {
	if o != nil && !common.IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *Split) SetAccount(v string) {
	o.Account = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Split) GetAmount() SplitAmount {
	if o == nil || common.IsNil(o.Amount) {
		var ret SplitAmount
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Split) GetAmountOk() (*SplitAmount, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Split) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given SplitAmount and assigns it to the Amount field.
func (o *Split) SetAmount(v SplitAmount) {
	o.Amount = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Split) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Split) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Split) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Split) SetDescription(v string) {
	o.Description = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *Split) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Split) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *Split) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *Split) SetReference(v string) {
	o.Reference = &v
}

// GetType returns the Type field value
func (o *Split) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Split) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Split) SetType(v string) {
	o.Type = v
}

func (o Split) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Split) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableSplit struct {
	value *Split
	isSet bool
}

func (v NullableSplit) Get() *Split {
	return v.value
}

func (v *NullableSplit) Set(val *Split) {
	v.value = val
	v.isSet = true
}

func (v NullableSplit) IsSet() bool {
	return v.isSet
}

func (v *NullableSplit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSplit(val *Split) *NullableSplit {
	return &NullableSplit{value: val, isSet: true}
}

func (v NullableSplit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSplit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *Split) isValidType() bool {
	var allowedEnumValues = []string{"AcquiringFees", "AdyenCommission", "AdyenFees", "AdyenMarkup", "BalanceAccount", "Commission", "Default", "Interchange", "MarketPlace", "PaymentFee", "Remainder", "SchemeFee", "Surcharge", "Tip", "TopUp", "VAT"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
