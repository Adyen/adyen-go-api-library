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

// checks if the PaymentRefundRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentRefundRequest{}

// PaymentRefundRequest struct for PaymentRefundRequest
type PaymentRefundRequest struct {
	Amount          Amount           `json:"amount"`
	ApplicationInfo *ApplicationInfo `json:"applicationInfo,omitempty"`
	// Price and product information of the refunded items, required for [partial refunds](https://docs.adyen.com/online-payments/refund#refund-a-payment). > This field is required for partial refunds with 3x 4x Oney, Affirm, Afterpay, Atome, Clearpay, Klarna, Ratepay, Walley, and Zip.
	LineItems []LineItem `json:"lineItems,omitempty"`
	// The merchant account that is used to process the payment.
	MerchantAccount string `json:"merchantAccount"`
	// The reason for the refund request.  Possible values:  * **FRAUD**  * **CUSTOMER REQUEST**  * **RETURN**  * **DUPLICATE**  * **OTHER**
	MerchantRefundReason common.NullableString `json:"merchantRefundReason,omitempty"`
	// Your reference for the refund request. Maximum length: 80 characters.
	Reference *string `json:"reference,omitempty"`
	// An array of objects specifying how the amount should be split between accounts when using Adyen for Platforms. For more information, see how to process payments for [marketplaces](https://docs.adyen.com/marketplaces/split-payments) or [platforms](https://docs.adyen.com/platforms/online-payments/split-payments/).
	Splits []Split `json:"splits,omitempty"`
	// The online store or [physical store](https://docs.adyen.com/point-of-sale/design-your-integration/determine-account-structure/#create-stores) that is processing the refund. This must be the same as the store name configured in your Customer Area.  Otherwise, you get an error and the refund fails.
	Store *string `json:"store,omitempty"`
}

// NewPaymentRefundRequest instantiates a new PaymentRefundRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentRefundRequest(amount Amount, merchantAccount string) *PaymentRefundRequest {
	this := PaymentRefundRequest{}
	this.Amount = amount
	this.MerchantAccount = merchantAccount
	return &this
}

// NewPaymentRefundRequestWithDefaults instantiates a new PaymentRefundRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentRefundRequestWithDefaults() *PaymentRefundRequest {
	this := PaymentRefundRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *PaymentRefundRequest) GetAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PaymentRefundRequest) GetAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PaymentRefundRequest) SetAmount(v Amount) {
	o.Amount = v
}

// GetApplicationInfo returns the ApplicationInfo field value if set, zero value otherwise.
func (o *PaymentRefundRequest) GetApplicationInfo() ApplicationInfo {
	if o == nil || common.IsNil(o.ApplicationInfo) {
		var ret ApplicationInfo
		return ret
	}
	return *o.ApplicationInfo
}

// GetApplicationInfoOk returns a tuple with the ApplicationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRefundRequest) GetApplicationInfoOk() (*ApplicationInfo, bool) {
	if o == nil || common.IsNil(o.ApplicationInfo) {
		return nil, false
	}
	return o.ApplicationInfo, true
}

// HasApplicationInfo returns a boolean if a field has been set.
func (o *PaymentRefundRequest) HasApplicationInfo() bool {
	if o != nil && !common.IsNil(o.ApplicationInfo) {
		return true
	}

	return false
}

// SetApplicationInfo gets a reference to the given ApplicationInfo and assigns it to the ApplicationInfo field.
func (o *PaymentRefundRequest) SetApplicationInfo(v ApplicationInfo) {
	o.ApplicationInfo = &v
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *PaymentRefundRequest) GetLineItems() []LineItem {
	if o == nil || common.IsNil(o.LineItems) {
		var ret []LineItem
		return ret
	}
	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRefundRequest) GetLineItemsOk() ([]LineItem, bool) {
	if o == nil || common.IsNil(o.LineItems) {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *PaymentRefundRequest) HasLineItems() bool {
	if o != nil && !common.IsNil(o.LineItems) {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given []LineItem and assigns it to the LineItems field.
func (o *PaymentRefundRequest) SetLineItems(v []LineItem) {
	o.LineItems = v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *PaymentRefundRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *PaymentRefundRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *PaymentRefundRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetMerchantRefundReason returns the MerchantRefundReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentRefundRequest) GetMerchantRefundReason() string {
	if o == nil || common.IsNil(o.MerchantRefundReason.Get()) {
		var ret string
		return ret
	}
	return *o.MerchantRefundReason.Get()
}

// GetMerchantRefundReasonOk returns a tuple with the MerchantRefundReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRefundRequest) GetMerchantRefundReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MerchantRefundReason.Get(), o.MerchantRefundReason.IsSet()
}

// HasMerchantRefundReason returns a boolean if a field has been set.
func (o *PaymentRefundRequest) HasMerchantRefundReason() bool {
	if o != nil && o.MerchantRefundReason.IsSet() {
		return true
	}

	return false
}

// SetMerchantRefundReason gets a reference to the given NullableString and assigns it to the MerchantRefundReason field.
func (o *PaymentRefundRequest) SetMerchantRefundReason(v string) {
	o.MerchantRefundReason.Set(&v)
}

// SetMerchantRefundReasonNil sets the value for MerchantRefundReason to be an explicit nil
func (o *PaymentRefundRequest) SetMerchantRefundReasonNil() {
	o.MerchantRefundReason.Set(nil)
}

// UnsetMerchantRefundReason ensures that no value is present for MerchantRefundReason, not even an explicit nil
func (o *PaymentRefundRequest) UnsetMerchantRefundReason() {
	o.MerchantRefundReason.Unset()
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PaymentRefundRequest) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRefundRequest) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PaymentRefundRequest) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PaymentRefundRequest) SetReference(v string) {
	o.Reference = &v
}

// GetSplits returns the Splits field value if set, zero value otherwise.
func (o *PaymentRefundRequest) GetSplits() []Split {
	if o == nil || common.IsNil(o.Splits) {
		var ret []Split
		return ret
	}
	return o.Splits
}

// GetSplitsOk returns a tuple with the Splits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRefundRequest) GetSplitsOk() ([]Split, bool) {
	if o == nil || common.IsNil(o.Splits) {
		return nil, false
	}
	return o.Splits, true
}

// HasSplits returns a boolean if a field has been set.
func (o *PaymentRefundRequest) HasSplits() bool {
	if o != nil && !common.IsNil(o.Splits) {
		return true
	}

	return false
}

// SetSplits gets a reference to the given []Split and assigns it to the Splits field.
func (o *PaymentRefundRequest) SetSplits(v []Split) {
	o.Splits = v
}

// GetStore returns the Store field value if set, zero value otherwise.
func (o *PaymentRefundRequest) GetStore() string {
	if o == nil || common.IsNil(o.Store) {
		var ret string
		return ret
	}
	return *o.Store
}

// GetStoreOk returns a tuple with the Store field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRefundRequest) GetStoreOk() (*string, bool) {
	if o == nil || common.IsNil(o.Store) {
		return nil, false
	}
	return o.Store, true
}

// HasStore returns a boolean if a field has been set.
func (o *PaymentRefundRequest) HasStore() bool {
	if o != nil && !common.IsNil(o.Store) {
		return true
	}

	return false
}

// SetStore gets a reference to the given string and assigns it to the Store field.
func (o *PaymentRefundRequest) SetStore(v string) {
	o.Store = &v
}

func (o PaymentRefundRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentRefundRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !common.IsNil(o.ApplicationInfo) {
		toSerialize["applicationInfo"] = o.ApplicationInfo
	}
	if !common.IsNil(o.LineItems) {
		toSerialize["lineItems"] = o.LineItems
	}
	toSerialize["merchantAccount"] = o.MerchantAccount
	if o.MerchantRefundReason.IsSet() {
		toSerialize["merchantRefundReason"] = o.MerchantRefundReason.Get()
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.Splits) {
		toSerialize["splits"] = o.Splits
	}
	if !common.IsNil(o.Store) {
		toSerialize["store"] = o.Store
	}
	return toSerialize, nil
}

type NullablePaymentRefundRequest struct {
	value *PaymentRefundRequest
	isSet bool
}

func (v NullablePaymentRefundRequest) Get() *PaymentRefundRequest {
	return v.value
}

func (v *NullablePaymentRefundRequest) Set(val *PaymentRefundRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentRefundRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentRefundRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentRefundRequest(val *PaymentRefundRequest) *NullablePaymentRefundRequest {
	return &NullablePaymentRefundRequest{value: val, isSet: true}
}

func (v NullablePaymentRefundRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentRefundRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PaymentRefundRequest) isValidMerchantRefundReason() bool {
	var allowedEnumValues = []string{"FRAUD", "CUSTOMER REQUEST", "RETURN", "DUPLICATE", "OTHER"}
	for _, allowed := range allowedEnumValues {
		if o.GetMerchantRefundReason() == allowed {
			return true
		}
	}
	return false
}
