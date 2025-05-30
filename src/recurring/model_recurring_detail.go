/*
Adyen Recurring API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package recurring

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the RecurringDetail type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RecurringDetail{}

// RecurringDetail struct for RecurringDetail
type RecurringDetail struct {
	// This field contains additional data, which may be returned in a particular response.  The additionalData object consists of entries, each of which includes the key and value.
	AdditionalData *map[string]string `json:"additionalData,omitempty"`
	// The alias of the credit card number.  Applies only to recurring contracts storing credit card details
	Alias *string `json:"alias,omitempty"`
	// The alias type of the credit card number.  Applies only to recurring contracts storing credit card details.
	AliasType      *string      `json:"aliasType,omitempty"`
	Bank           *BankAccount `json:"bank,omitempty"`
	BillingAddress *Address     `json:"billingAddress,omitempty"`
	Card           *Card        `json:"card,omitempty"`
	// Types of recurring contracts.
	ContractTypes []string `json:"contractTypes,omitempty"`
	// The date when the recurring details were created.
	CreationDate *time.Time `json:"creationDate,omitempty"`
	// The `pspReference` of the first recurring payment that created the recurring detail.
	FirstPspReference *string `json:"firstPspReference,omitempty"`
	// An optional descriptive name for this recurring detail.
	Name *string `json:"name,omitempty"`
	// Returned in the response if you are not tokenizing with Adyen and are using the Merchant-initiated transactions (MIT) framework from Mastercard or Visa.  This contains either the Mastercard Trace ID or the Visa Transaction ID.
	NetworkTxReference *string `json:"networkTxReference,omitempty"`
	// The  type or sub-brand of a payment method used, e.g. Visa Debit, Visa Corporate, etc. For more information, refer to [PaymentMethodVariant](https://docs.adyen.com/development-resources/paymentmethodvariant).
	PaymentMethodVariant *string `json:"paymentMethodVariant,omitempty"`
	// The reference that uniquely identifies the recurring detail.
	RecurringDetailReference string `json:"recurringDetailReference"`
	ShopperName              *Name  `json:"shopperName,omitempty"`
	// A shopper's social security number (only in countries where it is legal to collect).
	SocialSecurityNumber *string       `json:"socialSecurityNumber,omitempty"`
	TokenDetails         *TokenDetails `json:"tokenDetails,omitempty"`
	// The payment method, such as “mc\", \"visa\", \"ideal\", \"paypal\".
	Variant string `json:"variant"`
}

// NewRecurringDetail instantiates a new RecurringDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecurringDetail(recurringDetailReference string, variant string) *RecurringDetail {
	this := RecurringDetail{}
	this.RecurringDetailReference = recurringDetailReference
	this.Variant = variant
	return &this
}

// NewRecurringDetailWithDefaults instantiates a new RecurringDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecurringDetailWithDefaults() *RecurringDetail {
	this := RecurringDetail{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *RecurringDetail) GetAdditionalData() map[string]string {
	if o == nil || common.IsNil(o.AdditionalData) {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringDetail) GetAdditionalDataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *RecurringDetail) HasAdditionalData() bool {
	if o != nil && !common.IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]string and assigns it to the AdditionalData field.
func (o *RecurringDetail) SetAdditionalData(v map[string]string) {
	o.AdditionalData = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *RecurringDetail) GetAlias() string {
	if o == nil || common.IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringDetail) GetAliasOk() (*string, bool) {
	if o == nil || common.IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *RecurringDetail) HasAlias() bool {
	if o != nil && !common.IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *RecurringDetail) SetAlias(v string) {
	o.Alias = &v
}

// GetAliasType returns the AliasType field value if set, zero value otherwise.
func (o *RecurringDetail) GetAliasType() string {
	if o == nil || common.IsNil(o.AliasType) {
		var ret string
		return ret
	}
	return *o.AliasType
}

// GetAliasTypeOk returns a tuple with the AliasType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringDetail) GetAliasTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AliasType) {
		return nil, false
	}
	return o.AliasType, true
}

// HasAliasType returns a boolean if a field has been set.
func (o *RecurringDetail) HasAliasType() bool {
	if o != nil && !common.IsNil(o.AliasType) {
		return true
	}

	return false
}

// SetAliasType gets a reference to the given string and assigns it to the AliasType field.
func (o *RecurringDetail) SetAliasType(v string) {
	o.AliasType = &v
}

// GetBank returns the Bank field value if set, zero value otherwise.
func (o *RecurringDetail) GetBank() BankAccount {
	if o == nil || common.IsNil(o.Bank) {
		var ret BankAccount
		return ret
	}
	return *o.Bank
}

// GetBankOk returns a tuple with the Bank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringDetail) GetBankOk() (*BankAccount, bool) {
	if o == nil || common.IsNil(o.Bank) {
		return nil, false
	}
	return o.Bank, true
}

// HasBank returns a boolean if a field has been set.
func (o *RecurringDetail) HasBank() bool {
	if o != nil && !common.IsNil(o.Bank) {
		return true
	}

	return false
}

// SetBank gets a reference to the given BankAccount and assigns it to the Bank field.
func (o *RecurringDetail) SetBank(v BankAccount) {
	o.Bank = &v
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *RecurringDetail) GetBillingAddress() Address {
	if o == nil || common.IsNil(o.BillingAddress) {
		var ret Address
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringDetail) GetBillingAddressOk() (*Address, bool) {
	if o == nil || common.IsNil(o.BillingAddress) {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *RecurringDetail) HasBillingAddress() bool {
	if o != nil && !common.IsNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given Address and assigns it to the BillingAddress field.
func (o *RecurringDetail) SetBillingAddress(v Address) {
	o.BillingAddress = &v
}

// GetCard returns the Card field value if set, zero value otherwise.
func (o *RecurringDetail) GetCard() Card {
	if o == nil || common.IsNil(o.Card) {
		var ret Card
		return ret
	}
	return *o.Card
}

// GetCardOk returns a tuple with the Card field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringDetail) GetCardOk() (*Card, bool) {
	if o == nil || common.IsNil(o.Card) {
		return nil, false
	}
	return o.Card, true
}

// HasCard returns a boolean if a field has been set.
func (o *RecurringDetail) HasCard() bool {
	if o != nil && !common.IsNil(o.Card) {
		return true
	}

	return false
}

// SetCard gets a reference to the given Card and assigns it to the Card field.
func (o *RecurringDetail) SetCard(v Card) {
	o.Card = &v
}

// GetContractTypes returns the ContractTypes field value if set, zero value otherwise.
func (o *RecurringDetail) GetContractTypes() []string {
	if o == nil || common.IsNil(o.ContractTypes) {
		var ret []string
		return ret
	}
	return o.ContractTypes
}

// GetContractTypesOk returns a tuple with the ContractTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringDetail) GetContractTypesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.ContractTypes) {
		return nil, false
	}
	return o.ContractTypes, true
}

// HasContractTypes returns a boolean if a field has been set.
func (o *RecurringDetail) HasContractTypes() bool {
	if o != nil && !common.IsNil(o.ContractTypes) {
		return true
	}

	return false
}

// SetContractTypes gets a reference to the given []string and assigns it to the ContractTypes field.
func (o *RecurringDetail) SetContractTypes(v []string) {
	o.ContractTypes = v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *RecurringDetail) GetCreationDate() time.Time {
	if o == nil || common.IsNil(o.CreationDate) {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringDetail) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *RecurringDetail) HasCreationDate() bool {
	if o != nil && !common.IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *RecurringDetail) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetFirstPspReference returns the FirstPspReference field value if set, zero value otherwise.
func (o *RecurringDetail) GetFirstPspReference() string {
	if o == nil || common.IsNil(o.FirstPspReference) {
		var ret string
		return ret
	}
	return *o.FirstPspReference
}

// GetFirstPspReferenceOk returns a tuple with the FirstPspReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringDetail) GetFirstPspReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.FirstPspReference) {
		return nil, false
	}
	return o.FirstPspReference, true
}

// HasFirstPspReference returns a boolean if a field has been set.
func (o *RecurringDetail) HasFirstPspReference() bool {
	if o != nil && !common.IsNil(o.FirstPspReference) {
		return true
	}

	return false
}

// SetFirstPspReference gets a reference to the given string and assigns it to the FirstPspReference field.
func (o *RecurringDetail) SetFirstPspReference(v string) {
	o.FirstPspReference = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RecurringDetail) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringDetail) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RecurringDetail) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RecurringDetail) SetName(v string) {
	o.Name = &v
}

// GetNetworkTxReference returns the NetworkTxReference field value if set, zero value otherwise.
func (o *RecurringDetail) GetNetworkTxReference() string {
	if o == nil || common.IsNil(o.NetworkTxReference) {
		var ret string
		return ret
	}
	return *o.NetworkTxReference
}

// GetNetworkTxReferenceOk returns a tuple with the NetworkTxReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringDetail) GetNetworkTxReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.NetworkTxReference) {
		return nil, false
	}
	return o.NetworkTxReference, true
}

// HasNetworkTxReference returns a boolean if a field has been set.
func (o *RecurringDetail) HasNetworkTxReference() bool {
	if o != nil && !common.IsNil(o.NetworkTxReference) {
		return true
	}

	return false
}

// SetNetworkTxReference gets a reference to the given string and assigns it to the NetworkTxReference field.
func (o *RecurringDetail) SetNetworkTxReference(v string) {
	o.NetworkTxReference = &v
}

// GetPaymentMethodVariant returns the PaymentMethodVariant field value if set, zero value otherwise.
func (o *RecurringDetail) GetPaymentMethodVariant() string {
	if o == nil || common.IsNil(o.PaymentMethodVariant) {
		var ret string
		return ret
	}
	return *o.PaymentMethodVariant
}

// GetPaymentMethodVariantOk returns a tuple with the PaymentMethodVariant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringDetail) GetPaymentMethodVariantOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentMethodVariant) {
		return nil, false
	}
	return o.PaymentMethodVariant, true
}

// HasPaymentMethodVariant returns a boolean if a field has been set.
func (o *RecurringDetail) HasPaymentMethodVariant() bool {
	if o != nil && !common.IsNil(o.PaymentMethodVariant) {
		return true
	}

	return false
}

// SetPaymentMethodVariant gets a reference to the given string and assigns it to the PaymentMethodVariant field.
func (o *RecurringDetail) SetPaymentMethodVariant(v string) {
	o.PaymentMethodVariant = &v
}

// GetRecurringDetailReference returns the RecurringDetailReference field value
func (o *RecurringDetail) GetRecurringDetailReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecurringDetailReference
}

// GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field value
// and a boolean to check if the value has been set.
func (o *RecurringDetail) GetRecurringDetailReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecurringDetailReference, true
}

// SetRecurringDetailReference sets field value
func (o *RecurringDetail) SetRecurringDetailReference(v string) {
	o.RecurringDetailReference = v
}

// GetShopperName returns the ShopperName field value if set, zero value otherwise.
func (o *RecurringDetail) GetShopperName() Name {
	if o == nil || common.IsNil(o.ShopperName) {
		var ret Name
		return ret
	}
	return *o.ShopperName
}

// GetShopperNameOk returns a tuple with the ShopperName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringDetail) GetShopperNameOk() (*Name, bool) {
	if o == nil || common.IsNil(o.ShopperName) {
		return nil, false
	}
	return o.ShopperName, true
}

// HasShopperName returns a boolean if a field has been set.
func (o *RecurringDetail) HasShopperName() bool {
	if o != nil && !common.IsNil(o.ShopperName) {
		return true
	}

	return false
}

// SetShopperName gets a reference to the given Name and assigns it to the ShopperName field.
func (o *RecurringDetail) SetShopperName(v Name) {
	o.ShopperName = &v
}

// GetSocialSecurityNumber returns the SocialSecurityNumber field value if set, zero value otherwise.
func (o *RecurringDetail) GetSocialSecurityNumber() string {
	if o == nil || common.IsNil(o.SocialSecurityNumber) {
		var ret string
		return ret
	}
	return *o.SocialSecurityNumber
}

// GetSocialSecurityNumberOk returns a tuple with the SocialSecurityNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringDetail) GetSocialSecurityNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.SocialSecurityNumber) {
		return nil, false
	}
	return o.SocialSecurityNumber, true
}

// HasSocialSecurityNumber returns a boolean if a field has been set.
func (o *RecurringDetail) HasSocialSecurityNumber() bool {
	if o != nil && !common.IsNil(o.SocialSecurityNumber) {
		return true
	}

	return false
}

// SetSocialSecurityNumber gets a reference to the given string and assigns it to the SocialSecurityNumber field.
func (o *RecurringDetail) SetSocialSecurityNumber(v string) {
	o.SocialSecurityNumber = &v
}

// GetTokenDetails returns the TokenDetails field value if set, zero value otherwise.
func (o *RecurringDetail) GetTokenDetails() TokenDetails {
	if o == nil || common.IsNil(o.TokenDetails) {
		var ret TokenDetails
		return ret
	}
	return *o.TokenDetails
}

// GetTokenDetailsOk returns a tuple with the TokenDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringDetail) GetTokenDetailsOk() (*TokenDetails, bool) {
	if o == nil || common.IsNil(o.TokenDetails) {
		return nil, false
	}
	return o.TokenDetails, true
}

// HasTokenDetails returns a boolean if a field has been set.
func (o *RecurringDetail) HasTokenDetails() bool {
	if o != nil && !common.IsNil(o.TokenDetails) {
		return true
	}

	return false
}

// SetTokenDetails gets a reference to the given TokenDetails and assigns it to the TokenDetails field.
func (o *RecurringDetail) SetTokenDetails(v TokenDetails) {
	o.TokenDetails = &v
}

// GetVariant returns the Variant field value
func (o *RecurringDetail) GetVariant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Variant
}

// GetVariantOk returns a tuple with the Variant field value
// and a boolean to check if the value has been set.
func (o *RecurringDetail) GetVariantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variant, true
}

// SetVariant sets field value
func (o *RecurringDetail) SetVariant(v string) {
	o.Variant = v
}

func (o RecurringDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecurringDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AdditionalData) {
		toSerialize["additionalData"] = o.AdditionalData
	}
	if !common.IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if !common.IsNil(o.AliasType) {
		toSerialize["aliasType"] = o.AliasType
	}
	if !common.IsNil(o.Bank) {
		toSerialize["bank"] = o.Bank
	}
	if !common.IsNil(o.BillingAddress) {
		toSerialize["billingAddress"] = o.BillingAddress
	}
	if !common.IsNil(o.Card) {
		toSerialize["card"] = o.Card
	}
	if !common.IsNil(o.ContractTypes) {
		toSerialize["contractTypes"] = o.ContractTypes
	}
	if !common.IsNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	if !common.IsNil(o.FirstPspReference) {
		toSerialize["firstPspReference"] = o.FirstPspReference
	}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.NetworkTxReference) {
		toSerialize["networkTxReference"] = o.NetworkTxReference
	}
	if !common.IsNil(o.PaymentMethodVariant) {
		toSerialize["paymentMethodVariant"] = o.PaymentMethodVariant
	}
	toSerialize["recurringDetailReference"] = o.RecurringDetailReference
	if !common.IsNil(o.ShopperName) {
		toSerialize["shopperName"] = o.ShopperName
	}
	if !common.IsNil(o.SocialSecurityNumber) {
		toSerialize["socialSecurityNumber"] = o.SocialSecurityNumber
	}
	if !common.IsNil(o.TokenDetails) {
		toSerialize["tokenDetails"] = o.TokenDetails
	}
	toSerialize["variant"] = o.Variant
	return toSerialize, nil
}

type NullableRecurringDetail struct {
	value *RecurringDetail
	isSet bool
}

func (v NullableRecurringDetail) Get() *RecurringDetail {
	return v.value
}

func (v *NullableRecurringDetail) Set(val *RecurringDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableRecurringDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableRecurringDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecurringDetail(val *RecurringDetail) *NullableRecurringDetail {
	return &NullableRecurringDetail{value: val, isSet: true}
}

func (v NullableRecurringDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecurringDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
