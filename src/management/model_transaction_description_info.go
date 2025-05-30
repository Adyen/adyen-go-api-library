/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the TransactionDescriptionInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransactionDescriptionInfo{}

// TransactionDescriptionInfo struct for TransactionDescriptionInfo
type TransactionDescriptionInfo struct {
	// The text to be shown on the shopper's bank statement.  We recommend sending a maximum of 22 characters, otherwise banks might truncate the string.  Allowed characters: **a-z**, **A-Z**, **0-9**, spaces, and special characters **. , ' _ - ? + * /_**.
	DoingBusinessAsName *string `json:"doingBusinessAsName,omitempty"`
	// The type of transaction description you want to use: - **fixed**: The transaction description set in this request is used for all payments with this payment method. - **append**: The transaction description set in this request is used as a base for all payments with this payment method. The [transaction description set in the request to process the payment](https://docs.adyen.com/api-explorer/Checkout/70/post/sessions#request-shopperStatement) is appended to this base description. Note that if the combined length exceeds 22 characters, banks may truncate the string. - **dynamic**: Only the [transaction description set in the request to process the payment](https://docs.adyen.com/api-explorer/Checkout/70/post/sessions#request-shopperStatement) is used for payments with this payment method.
	Type *string `json:"type,omitempty"`
}

// NewTransactionDescriptionInfo instantiates a new TransactionDescriptionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionDescriptionInfo() *TransactionDescriptionInfo {
	this := TransactionDescriptionInfo{}
	var type_ string = "dynamic"
	this.Type = &type_
	return &this
}

// NewTransactionDescriptionInfoWithDefaults instantiates a new TransactionDescriptionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionDescriptionInfoWithDefaults() *TransactionDescriptionInfo {
	this := TransactionDescriptionInfo{}
	var type_ string = "dynamic"
	this.Type = &type_
	return &this
}

// GetDoingBusinessAsName returns the DoingBusinessAsName field value if set, zero value otherwise.
func (o *TransactionDescriptionInfo) GetDoingBusinessAsName() string {
	if o == nil || common.IsNil(o.DoingBusinessAsName) {
		var ret string
		return ret
	}
	return *o.DoingBusinessAsName
}

// GetDoingBusinessAsNameOk returns a tuple with the DoingBusinessAsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDescriptionInfo) GetDoingBusinessAsNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.DoingBusinessAsName) {
		return nil, false
	}
	return o.DoingBusinessAsName, true
}

// HasDoingBusinessAsName returns a boolean if a field has been set.
func (o *TransactionDescriptionInfo) HasDoingBusinessAsName() bool {
	if o != nil && !common.IsNil(o.DoingBusinessAsName) {
		return true
	}

	return false
}

// SetDoingBusinessAsName gets a reference to the given string and assigns it to the DoingBusinessAsName field.
func (o *TransactionDescriptionInfo) SetDoingBusinessAsName(v string) {
	o.DoingBusinessAsName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TransactionDescriptionInfo) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDescriptionInfo) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TransactionDescriptionInfo) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TransactionDescriptionInfo) SetType(v string) {
	o.Type = &v
}

func (o TransactionDescriptionInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionDescriptionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.DoingBusinessAsName) {
		toSerialize["doingBusinessAsName"] = o.DoingBusinessAsName
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableTransactionDescriptionInfo struct {
	value *TransactionDescriptionInfo
	isSet bool
}

func (v NullableTransactionDescriptionInfo) Get() *TransactionDescriptionInfo {
	return v.value
}

func (v *NullableTransactionDescriptionInfo) Set(val *TransactionDescriptionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionDescriptionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionDescriptionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionDescriptionInfo(val *TransactionDescriptionInfo) *NullableTransactionDescriptionInfo {
	return &NullableTransactionDescriptionInfo{value: val, isSet: true}
}

func (v NullableTransactionDescriptionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionDescriptionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *TransactionDescriptionInfo) isValidType() bool {
	var allowedEnumValues = []string{"append", "dynamic", "fixed"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
