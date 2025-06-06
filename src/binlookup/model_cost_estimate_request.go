/*
Adyen BinLookup API

API version: 54
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package binlookup

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the CostEstimateRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CostEstimateRequest{}

// CostEstimateRequest struct for CostEstimateRequest
type CostEstimateRequest struct {
	Amount      Amount                   `json:"amount"`
	Assumptions *CostEstimateAssumptions `json:"assumptions,omitempty"`
	// The card number (4-19 characters) for PCI compliant use cases. Do not use any separators.  > Either the `cardNumber` or `encryptedCardNumber` field must be provided in a payment request.
	CardNumber *string `json:"cardNumber,omitempty"`
	// Encrypted data that stores card information for non PCI-compliant use cases. The encrypted data must be created with the Checkout Card Component or Secured Fields Component, and must contain the `encryptedCardNumber` field.  > Either the `cardNumber` or `encryptedCardNumber` field must be provided in a payment request.
	EncryptedCardNumber *string `json:"encryptedCardNumber,omitempty"`
	// The merchant account identifier you want to process the (transaction) request with.
	MerchantAccount string           `json:"merchantAccount"`
	MerchantDetails *MerchantDetails `json:"merchantDetails,omitempty"`
	Recurring       *Recurring       `json:"recurring,omitempty"`
	// The `recurringDetailReference` you want to use for this cost estimate. The value `LATEST` can be used to select the most recently stored recurring detail.
	SelectedRecurringDetailReference *string `json:"selectedRecurringDetailReference,omitempty"`
	// Specifies the sales channel, through which the shopper gives their card details, and whether the shopper is a returning customer. For the web service API, Adyen assumes Ecommerce shopper interaction by default.  This field has the following possible values: * `Ecommerce` - Online transactions where the cardholder is present (online). For better authorisation rates, we recommend sending the card security code (CSC) along with the request. * `ContAuth` - Card on file and/or subscription transactions, where the card holder is known to the merchant (returning customer). If the shopper is present (online), you can supply also the CSC to improve authorisation (one-click payment). * `Moto` - Mail-order and telephone-order transactions where the shopper is in contact with the merchant via email or telephone. * `POS` - Point-of-sale transactions where the shopper is physically present to make a payment using a secure payment terminal.
	ShopperInteraction *string `json:"shopperInteraction,omitempty"`
	// Required for recurring payments.  Your reference to uniquely identify this shopper, for example user ID or account ID. The value is case-sensitive and must be at least three characters. > Your reference must not include personally identifiable information (PII) such as name or email address.
	ShopperReference *string `json:"shopperReference,omitempty"`
}

// NewCostEstimateRequest instantiates a new CostEstimateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCostEstimateRequest(amount Amount, merchantAccount string) *CostEstimateRequest {
	this := CostEstimateRequest{}
	this.Amount = amount
	this.MerchantAccount = merchantAccount
	return &this
}

// NewCostEstimateRequestWithDefaults instantiates a new CostEstimateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCostEstimateRequestWithDefaults() *CostEstimateRequest {
	this := CostEstimateRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *CostEstimateRequest) GetAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CostEstimateRequest) GetAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CostEstimateRequest) SetAmount(v Amount) {
	o.Amount = v
}

// GetAssumptions returns the Assumptions field value if set, zero value otherwise.
func (o *CostEstimateRequest) GetAssumptions() CostEstimateAssumptions {
	if o == nil || common.IsNil(o.Assumptions) {
		var ret CostEstimateAssumptions
		return ret
	}
	return *o.Assumptions
}

// GetAssumptionsOk returns a tuple with the Assumptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostEstimateRequest) GetAssumptionsOk() (*CostEstimateAssumptions, bool) {
	if o == nil || common.IsNil(o.Assumptions) {
		return nil, false
	}
	return o.Assumptions, true
}

// HasAssumptions returns a boolean if a field has been set.
func (o *CostEstimateRequest) HasAssumptions() bool {
	if o != nil && !common.IsNil(o.Assumptions) {
		return true
	}

	return false
}

// SetAssumptions gets a reference to the given CostEstimateAssumptions and assigns it to the Assumptions field.
func (o *CostEstimateRequest) SetAssumptions(v CostEstimateAssumptions) {
	o.Assumptions = &v
}

// GetCardNumber returns the CardNumber field value if set, zero value otherwise.
func (o *CostEstimateRequest) GetCardNumber() string {
	if o == nil || common.IsNil(o.CardNumber) {
		var ret string
		return ret
	}
	return *o.CardNumber
}

// GetCardNumberOk returns a tuple with the CardNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostEstimateRequest) GetCardNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.CardNumber) {
		return nil, false
	}
	return o.CardNumber, true
}

// HasCardNumber returns a boolean if a field has been set.
func (o *CostEstimateRequest) HasCardNumber() bool {
	if o != nil && !common.IsNil(o.CardNumber) {
		return true
	}

	return false
}

// SetCardNumber gets a reference to the given string and assigns it to the CardNumber field.
func (o *CostEstimateRequest) SetCardNumber(v string) {
	o.CardNumber = &v
}

// GetEncryptedCardNumber returns the EncryptedCardNumber field value if set, zero value otherwise.
func (o *CostEstimateRequest) GetEncryptedCardNumber() string {
	if o == nil || common.IsNil(o.EncryptedCardNumber) {
		var ret string
		return ret
	}
	return *o.EncryptedCardNumber
}

// GetEncryptedCardNumberOk returns a tuple with the EncryptedCardNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostEstimateRequest) GetEncryptedCardNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.EncryptedCardNumber) {
		return nil, false
	}
	return o.EncryptedCardNumber, true
}

// HasEncryptedCardNumber returns a boolean if a field has been set.
func (o *CostEstimateRequest) HasEncryptedCardNumber() bool {
	if o != nil && !common.IsNil(o.EncryptedCardNumber) {
		return true
	}

	return false
}

// SetEncryptedCardNumber gets a reference to the given string and assigns it to the EncryptedCardNumber field.
func (o *CostEstimateRequest) SetEncryptedCardNumber(v string) {
	o.EncryptedCardNumber = &v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *CostEstimateRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *CostEstimateRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *CostEstimateRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetMerchantDetails returns the MerchantDetails field value if set, zero value otherwise.
func (o *CostEstimateRequest) GetMerchantDetails() MerchantDetails {
	if o == nil || common.IsNil(o.MerchantDetails) {
		var ret MerchantDetails
		return ret
	}
	return *o.MerchantDetails
}

// GetMerchantDetailsOk returns a tuple with the MerchantDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostEstimateRequest) GetMerchantDetailsOk() (*MerchantDetails, bool) {
	if o == nil || common.IsNil(o.MerchantDetails) {
		return nil, false
	}
	return o.MerchantDetails, true
}

// HasMerchantDetails returns a boolean if a field has been set.
func (o *CostEstimateRequest) HasMerchantDetails() bool {
	if o != nil && !common.IsNil(o.MerchantDetails) {
		return true
	}

	return false
}

// SetMerchantDetails gets a reference to the given MerchantDetails and assigns it to the MerchantDetails field.
func (o *CostEstimateRequest) SetMerchantDetails(v MerchantDetails) {
	o.MerchantDetails = &v
}

// GetRecurring returns the Recurring field value if set, zero value otherwise.
func (o *CostEstimateRequest) GetRecurring() Recurring {
	if o == nil || common.IsNil(o.Recurring) {
		var ret Recurring
		return ret
	}
	return *o.Recurring
}

// GetRecurringOk returns a tuple with the Recurring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostEstimateRequest) GetRecurringOk() (*Recurring, bool) {
	if o == nil || common.IsNil(o.Recurring) {
		return nil, false
	}
	return o.Recurring, true
}

// HasRecurring returns a boolean if a field has been set.
func (o *CostEstimateRequest) HasRecurring() bool {
	if o != nil && !common.IsNil(o.Recurring) {
		return true
	}

	return false
}

// SetRecurring gets a reference to the given Recurring and assigns it to the Recurring field.
func (o *CostEstimateRequest) SetRecurring(v Recurring) {
	o.Recurring = &v
}

// GetSelectedRecurringDetailReference returns the SelectedRecurringDetailReference field value if set, zero value otherwise.
func (o *CostEstimateRequest) GetSelectedRecurringDetailReference() string {
	if o == nil || common.IsNil(o.SelectedRecurringDetailReference) {
		var ret string
		return ret
	}
	return *o.SelectedRecurringDetailReference
}

// GetSelectedRecurringDetailReferenceOk returns a tuple with the SelectedRecurringDetailReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostEstimateRequest) GetSelectedRecurringDetailReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.SelectedRecurringDetailReference) {
		return nil, false
	}
	return o.SelectedRecurringDetailReference, true
}

// HasSelectedRecurringDetailReference returns a boolean if a field has been set.
func (o *CostEstimateRequest) HasSelectedRecurringDetailReference() bool {
	if o != nil && !common.IsNil(o.SelectedRecurringDetailReference) {
		return true
	}

	return false
}

// SetSelectedRecurringDetailReference gets a reference to the given string and assigns it to the SelectedRecurringDetailReference field.
func (o *CostEstimateRequest) SetSelectedRecurringDetailReference(v string) {
	o.SelectedRecurringDetailReference = &v
}

// GetShopperInteraction returns the ShopperInteraction field value if set, zero value otherwise.
func (o *CostEstimateRequest) GetShopperInteraction() string {
	if o == nil || common.IsNil(o.ShopperInteraction) {
		var ret string
		return ret
	}
	return *o.ShopperInteraction
}

// GetShopperInteractionOk returns a tuple with the ShopperInteraction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostEstimateRequest) GetShopperInteractionOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperInteraction) {
		return nil, false
	}
	return o.ShopperInteraction, true
}

// HasShopperInteraction returns a boolean if a field has been set.
func (o *CostEstimateRequest) HasShopperInteraction() bool {
	if o != nil && !common.IsNil(o.ShopperInteraction) {
		return true
	}

	return false
}

// SetShopperInteraction gets a reference to the given string and assigns it to the ShopperInteraction field.
func (o *CostEstimateRequest) SetShopperInteraction(v string) {
	o.ShopperInteraction = &v
}

// GetShopperReference returns the ShopperReference field value if set, zero value otherwise.
func (o *CostEstimateRequest) GetShopperReference() string {
	if o == nil || common.IsNil(o.ShopperReference) {
		var ret string
		return ret
	}
	return *o.ShopperReference
}

// GetShopperReferenceOk returns a tuple with the ShopperReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostEstimateRequest) GetShopperReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperReference) {
		return nil, false
	}
	return o.ShopperReference, true
}

// HasShopperReference returns a boolean if a field has been set.
func (o *CostEstimateRequest) HasShopperReference() bool {
	if o != nil && !common.IsNil(o.ShopperReference) {
		return true
	}

	return false
}

// SetShopperReference gets a reference to the given string and assigns it to the ShopperReference field.
func (o *CostEstimateRequest) SetShopperReference(v string) {
	o.ShopperReference = &v
}

func (o CostEstimateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CostEstimateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !common.IsNil(o.Assumptions) {
		toSerialize["assumptions"] = o.Assumptions
	}
	if !common.IsNil(o.CardNumber) {
		toSerialize["cardNumber"] = o.CardNumber
	}
	if !common.IsNil(o.EncryptedCardNumber) {
		toSerialize["encryptedCardNumber"] = o.EncryptedCardNumber
	}
	toSerialize["merchantAccount"] = o.MerchantAccount
	if !common.IsNil(o.MerchantDetails) {
		toSerialize["merchantDetails"] = o.MerchantDetails
	}
	if !common.IsNil(o.Recurring) {
		toSerialize["recurring"] = o.Recurring
	}
	if !common.IsNil(o.SelectedRecurringDetailReference) {
		toSerialize["selectedRecurringDetailReference"] = o.SelectedRecurringDetailReference
	}
	if !common.IsNil(o.ShopperInteraction) {
		toSerialize["shopperInteraction"] = o.ShopperInteraction
	}
	if !common.IsNil(o.ShopperReference) {
		toSerialize["shopperReference"] = o.ShopperReference
	}
	return toSerialize, nil
}

type NullableCostEstimateRequest struct {
	value *CostEstimateRequest
	isSet bool
}

func (v NullableCostEstimateRequest) Get() *CostEstimateRequest {
	return v.value
}

func (v *NullableCostEstimateRequest) Set(val *CostEstimateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCostEstimateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCostEstimateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCostEstimateRequest(val *CostEstimateRequest) *NullableCostEstimateRequest {
	return &NullableCostEstimateRequest{value: val, isSet: true}
}

func (v NullableCostEstimateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCostEstimateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CostEstimateRequest) isValidShopperInteraction() bool {
	var allowedEnumValues = []string{"Ecommerce", "ContAuth", "Moto", "POS"}
	for _, allowed := range allowedEnumValues {
		if o.GetShopperInteraction() == allowed {
			return true
		}
	}
	return false
}
