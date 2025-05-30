/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the TransferRouteRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransferRouteRequest{}

// TransferRouteRequest struct for TransferRouteRequest
type TransferRouteRequest struct {
	// The unique identifier of the source [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts__resParam_id). Required if `counterparty` is **transferInstrumentId**.
	BalanceAccountId *string `json:"balanceAccountId,omitempty"`
	// The unique identifier assigned to the balance platform associated with the account holder.
	BalancePlatform string `json:"balancePlatform"`
	//  The type of transfer. Possible values:    - **bank**: Transfer to a [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments__resParam_id) or a bank account.
	Category     string        `json:"category"`
	Counterparty *Counterparty `json:"counterparty,omitempty"`
	// The two-character ISO-3166-1 alpha-2 country code of the counterparty. For example, **US** or **NL**.  > Either `counterparty` or `country` field must be provided in a transfer route request.
	Country *string `json:"country,omitempty"`
	// The three-character ISO currency code of transfer. For example, **USD** or **EUR**.
	Currency string `json:"currency"`
	// The list of priorities for the bank transfer. Priorities set the speed at which the transfer is sent and the fees that you have to pay. Multiple values can be provided. Possible values:  * **regular**: for normal, low-value transactions.  * **fast**: a faster way to transfer funds, but the fees are higher. Recommended for high-priority, low-value transactions.  * **wire**: the fastest way to transfer funds, but this has the highest fees. Recommended for high-priority, high-value transactions.  * **instant**: for instant funds transfers in [SEPA countries](https://www.ecb.europa.eu/paym/integration/retail/sepa/html/index.en.html).  * **crossBorder**: for high-value transfers to a recipient in a different country.  * **internal**: for transfers to an Adyen-issued business bank account (by bank account number/IBAN).
	Priorities []string `json:"priorities,omitempty"`
}

// NewTransferRouteRequest instantiates a new TransferRouteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRouteRequest(balancePlatform string, category string, currency string) *TransferRouteRequest {
	this := TransferRouteRequest{}
	this.BalancePlatform = balancePlatform
	this.Category = category
	this.Currency = currency
	return &this
}

// NewTransferRouteRequestWithDefaults instantiates a new TransferRouteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRouteRequestWithDefaults() *TransferRouteRequest {
	this := TransferRouteRequest{}
	return &this
}

// GetBalanceAccountId returns the BalanceAccountId field value if set, zero value otherwise.
func (o *TransferRouteRequest) GetBalanceAccountId() string {
	if o == nil || common.IsNil(o.BalanceAccountId) {
		var ret string
		return ret
	}
	return *o.BalanceAccountId
}

// GetBalanceAccountIdOk returns a tuple with the BalanceAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRouteRequest) GetBalanceAccountIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.BalanceAccountId) {
		return nil, false
	}
	return o.BalanceAccountId, true
}

// HasBalanceAccountId returns a boolean if a field has been set.
func (o *TransferRouteRequest) HasBalanceAccountId() bool {
	if o != nil && !common.IsNil(o.BalanceAccountId) {
		return true
	}

	return false
}

// SetBalanceAccountId gets a reference to the given string and assigns it to the BalanceAccountId field.
func (o *TransferRouteRequest) SetBalanceAccountId(v string) {
	o.BalanceAccountId = &v
}

// GetBalancePlatform returns the BalancePlatform field value
func (o *TransferRouteRequest) GetBalancePlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BalancePlatform
}

// GetBalancePlatformOk returns a tuple with the BalancePlatform field value
// and a boolean to check if the value has been set.
func (o *TransferRouteRequest) GetBalancePlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BalancePlatform, true
}

// SetBalancePlatform sets field value
func (o *TransferRouteRequest) SetBalancePlatform(v string) {
	o.BalancePlatform = v
}

// GetCategory returns the Category field value
func (o *TransferRouteRequest) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *TransferRouteRequest) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *TransferRouteRequest) SetCategory(v string) {
	o.Category = v
}

// GetCounterparty returns the Counterparty field value if set, zero value otherwise.
func (o *TransferRouteRequest) GetCounterparty() Counterparty {
	if o == nil || common.IsNil(o.Counterparty) {
		var ret Counterparty
		return ret
	}
	return *o.Counterparty
}

// GetCounterpartyOk returns a tuple with the Counterparty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRouteRequest) GetCounterpartyOk() (*Counterparty, bool) {
	if o == nil || common.IsNil(o.Counterparty) {
		return nil, false
	}
	return o.Counterparty, true
}

// HasCounterparty returns a boolean if a field has been set.
func (o *TransferRouteRequest) HasCounterparty() bool {
	if o != nil && !common.IsNil(o.Counterparty) {
		return true
	}

	return false
}

// SetCounterparty gets a reference to the given Counterparty and assigns it to the Counterparty field.
func (o *TransferRouteRequest) SetCounterparty(v Counterparty) {
	o.Counterparty = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *TransferRouteRequest) GetCountry() string {
	if o == nil || common.IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRouteRequest) GetCountryOk() (*string, bool) {
	if o == nil || common.IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *TransferRouteRequest) HasCountry() bool {
	if o != nil && !common.IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *TransferRouteRequest) SetCountry(v string) {
	o.Country = &v
}

// GetCurrency returns the Currency field value
func (o *TransferRouteRequest) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *TransferRouteRequest) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *TransferRouteRequest) SetCurrency(v string) {
	o.Currency = v
}

// GetPriorities returns the Priorities field value if set, zero value otherwise.
func (o *TransferRouteRequest) GetPriorities() []string {
	if o == nil || common.IsNil(o.Priorities) {
		var ret []string
		return ret
	}
	return o.Priorities
}

// GetPrioritiesOk returns a tuple with the Priorities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRouteRequest) GetPrioritiesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Priorities) {
		return nil, false
	}
	return o.Priorities, true
}

// HasPriorities returns a boolean if a field has been set.
func (o *TransferRouteRequest) HasPriorities() bool {
	if o != nil && !common.IsNil(o.Priorities) {
		return true
	}

	return false
}

// SetPriorities gets a reference to the given []string and assigns it to the Priorities field.
func (o *TransferRouteRequest) SetPriorities(v []string) {
	o.Priorities = v
}

func (o TransferRouteRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferRouteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BalanceAccountId) {
		toSerialize["balanceAccountId"] = o.BalanceAccountId
	}
	toSerialize["balancePlatform"] = o.BalancePlatform
	toSerialize["category"] = o.Category
	if !common.IsNil(o.Counterparty) {
		toSerialize["counterparty"] = o.Counterparty
	}
	if !common.IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	toSerialize["currency"] = o.Currency
	if !common.IsNil(o.Priorities) {
		toSerialize["priorities"] = o.Priorities
	}
	return toSerialize, nil
}

type NullableTransferRouteRequest struct {
	value *TransferRouteRequest
	isSet bool
}

func (v NullableTransferRouteRequest) Get() *TransferRouteRequest {
	return v.value
}

func (v *NullableTransferRouteRequest) Set(val *TransferRouteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRouteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRouteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRouteRequest(val *TransferRouteRequest) *NullableTransferRouteRequest {
	return &NullableTransferRouteRequest{value: val, isSet: true}
}

func (v NullableTransferRouteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRouteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *TransferRouteRequest) isValidCategory() bool {
	var allowedEnumValues = []string{"bank"}
	for _, allowed := range allowedEnumValues {
		if o.GetCategory() == allowed {
			return true
		}
	}
	return false
}
