/*
Adyen Checkout API

API version: 70
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the PaymentCancelResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentCancelResponse{}

// PaymentCancelResponse struct for PaymentCancelResponse
type PaymentCancelResponse struct {
	// The merchant account that is used to process the payment.
	MerchantAccount string `json:"merchantAccount"`
	// The [`pspReference`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment to cancel.
	PaymentPspReference string `json:"paymentPspReference"`
	// Adyen's 16-character reference associated with the cancel request.
	PspReference string `json:"pspReference"`
	// Your reference for the cancel request.
	Reference *string `json:"reference,omitempty"`
	// The status of your request. This will always have the value **received**.
	Status string `json:"status"`
}

// NewPaymentCancelResponse instantiates a new PaymentCancelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentCancelResponse(merchantAccount string, paymentPspReference string, pspReference string, status string) *PaymentCancelResponse {
	this := PaymentCancelResponse{}
	this.MerchantAccount = merchantAccount
	this.PaymentPspReference = paymentPspReference
	this.PspReference = pspReference
	this.Status = status
	return &this
}

// NewPaymentCancelResponseWithDefaults instantiates a new PaymentCancelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentCancelResponseWithDefaults() *PaymentCancelResponse {
	this := PaymentCancelResponse{}
	return &this
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *PaymentCancelResponse) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *PaymentCancelResponse) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *PaymentCancelResponse) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetPaymentPspReference returns the PaymentPspReference field value
func (o *PaymentCancelResponse) GetPaymentPspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentPspReference
}

// GetPaymentPspReferenceOk returns a tuple with the PaymentPspReference field value
// and a boolean to check if the value has been set.
func (o *PaymentCancelResponse) GetPaymentPspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentPspReference, true
}

// SetPaymentPspReference sets field value
func (o *PaymentCancelResponse) SetPaymentPspReference(v string) {
	o.PaymentPspReference = v
}

// GetPspReference returns the PspReference field value
func (o *PaymentCancelResponse) GetPspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value
// and a boolean to check if the value has been set.
func (o *PaymentCancelResponse) GetPspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PspReference, true
}

// SetPspReference sets field value
func (o *PaymentCancelResponse) SetPspReference(v string) {
	o.PspReference = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PaymentCancelResponse) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCancelResponse) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PaymentCancelResponse) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PaymentCancelResponse) SetReference(v string) {
	o.Reference = &v
}

// GetStatus returns the Status field value
func (o *PaymentCancelResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentCancelResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentCancelResponse) SetStatus(v string) {
	o.Status = v
}

func (o PaymentCancelResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentCancelResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merchantAccount"] = o.MerchantAccount
	toSerialize["paymentPspReference"] = o.PaymentPspReference
	toSerialize["pspReference"] = o.PspReference
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullablePaymentCancelResponse struct {
	value *PaymentCancelResponse
	isSet bool
}

func (v NullablePaymentCancelResponse) Get() *PaymentCancelResponse {
	return v.value
}

func (v *NullablePaymentCancelResponse) Set(val *PaymentCancelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentCancelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentCancelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentCancelResponse(val *PaymentCancelResponse) *NullablePaymentCancelResponse {
	return &NullablePaymentCancelResponse{value: val, isSet: true}
}

func (v NullablePaymentCancelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentCancelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PaymentCancelResponse) isValidStatus() bool {
	var allowedEnumValues = []string{"received"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}