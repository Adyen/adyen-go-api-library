/*
Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the SplitConfigurationLogic type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SplitConfigurationLogic{}

// SplitConfigurationLogic struct for SplitConfigurationLogic
type SplitConfigurationLogic struct {
	AdditionalCommission *AdditionalCommission `json:"additionalCommission,omitempty"`
	// Specifies the logic to apply when booking the chargeback amount.  Possible values: **deductFromLiableAccount**, **deductFromOneBalanceAccount**, **deductAccordingToSplitRatio**.
	Chargeback *string    `json:"chargeback,omitempty"`
	Commission Commission `json:"commission"`
	// Specifies the logic to apply when booking the transaction fees.  Possible values: **deductFromLiableAccount**, **deductFromOneBalanceAccount**.
	PaymentFee string `json:"paymentFee"`
	// Specifies the logic to apply when booking the amount left over after currency conversion.  Possible values: **addToLiableAccount**, **addToOneBalanceAccount**.
	Remainder *string `json:"remainder,omitempty"`
	// Unique identifier of the split logic that is applied when the split configuration conditions are met.
	SplitLogicId *string `json:"splitLogicId,omitempty"`
	// Specifies the logic to apply when booking the surcharge amount.  Possible values: **addToLiableAccount**, **addToOneBalanceAccount**
	Surcharge *string `json:"surcharge,omitempty"`
	// Specifies the logic to apply when booking tips (gratuity).  Possible values: **addToLiableAccount**, **addToOneBalanceAccount**.
	Tip *string `json:"tip,omitempty"`
}

// NewSplitConfigurationLogic instantiates a new SplitConfigurationLogic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSplitConfigurationLogic(commission Commission, paymentFee string) *SplitConfigurationLogic {
	this := SplitConfigurationLogic{}
	this.Commission = commission
	this.PaymentFee = paymentFee
	return &this
}

// NewSplitConfigurationLogicWithDefaults instantiates a new SplitConfigurationLogic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSplitConfigurationLogicWithDefaults() *SplitConfigurationLogic {
	this := SplitConfigurationLogic{}
	return &this
}

// GetAdditionalCommission returns the AdditionalCommission field value if set, zero value otherwise.
func (o *SplitConfigurationLogic) GetAdditionalCommission() AdditionalCommission {
	if o == nil || common.IsNil(o.AdditionalCommission) {
		var ret AdditionalCommission
		return ret
	}
	return *o.AdditionalCommission
}

// GetAdditionalCommissionOk returns a tuple with the AdditionalCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfigurationLogic) GetAdditionalCommissionOk() (*AdditionalCommission, bool) {
	if o == nil || common.IsNil(o.AdditionalCommission) {
		return nil, false
	}
	return o.AdditionalCommission, true
}

// HasAdditionalCommission returns a boolean if a field has been set.
func (o *SplitConfigurationLogic) HasAdditionalCommission() bool {
	if o != nil && !common.IsNil(o.AdditionalCommission) {
		return true
	}

	return false
}

// SetAdditionalCommission gets a reference to the given AdditionalCommission and assigns it to the AdditionalCommission field.
func (o *SplitConfigurationLogic) SetAdditionalCommission(v AdditionalCommission) {
	o.AdditionalCommission = &v
}

// GetChargeback returns the Chargeback field value if set, zero value otherwise.
func (o *SplitConfigurationLogic) GetChargeback() string {
	if o == nil || common.IsNil(o.Chargeback) {
		var ret string
		return ret
	}
	return *o.Chargeback
}

// GetChargebackOk returns a tuple with the Chargeback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfigurationLogic) GetChargebackOk() (*string, bool) {
	if o == nil || common.IsNil(o.Chargeback) {
		return nil, false
	}
	return o.Chargeback, true
}

// HasChargeback returns a boolean if a field has been set.
func (o *SplitConfigurationLogic) HasChargeback() bool {
	if o != nil && !common.IsNil(o.Chargeback) {
		return true
	}

	return false
}

// SetChargeback gets a reference to the given string and assigns it to the Chargeback field.
func (o *SplitConfigurationLogic) SetChargeback(v string) {
	o.Chargeback = &v
}

// GetCommission returns the Commission field value
func (o *SplitConfigurationLogic) GetCommission() Commission {
	if o == nil {
		var ret Commission
		return ret
	}

	return o.Commission
}

// GetCommissionOk returns a tuple with the Commission field value
// and a boolean to check if the value has been set.
func (o *SplitConfigurationLogic) GetCommissionOk() (*Commission, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Commission, true
}

// SetCommission sets field value
func (o *SplitConfigurationLogic) SetCommission(v Commission) {
	o.Commission = v
}

// GetPaymentFee returns the PaymentFee field value
func (o *SplitConfigurationLogic) GetPaymentFee() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentFee
}

// GetPaymentFeeOk returns a tuple with the PaymentFee field value
// and a boolean to check if the value has been set.
func (o *SplitConfigurationLogic) GetPaymentFeeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentFee, true
}

// SetPaymentFee sets field value
func (o *SplitConfigurationLogic) SetPaymentFee(v string) {
	o.PaymentFee = v
}

// GetRemainder returns the Remainder field value if set, zero value otherwise.
func (o *SplitConfigurationLogic) GetRemainder() string {
	if o == nil || common.IsNil(o.Remainder) {
		var ret string
		return ret
	}
	return *o.Remainder
}

// GetRemainderOk returns a tuple with the Remainder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfigurationLogic) GetRemainderOk() (*string, bool) {
	if o == nil || common.IsNil(o.Remainder) {
		return nil, false
	}
	return o.Remainder, true
}

// HasRemainder returns a boolean if a field has been set.
func (o *SplitConfigurationLogic) HasRemainder() bool {
	if o != nil && !common.IsNil(o.Remainder) {
		return true
	}

	return false
}

// SetRemainder gets a reference to the given string and assigns it to the Remainder field.
func (o *SplitConfigurationLogic) SetRemainder(v string) {
	o.Remainder = &v
}

// GetSplitLogicId returns the SplitLogicId field value if set, zero value otherwise.
func (o *SplitConfigurationLogic) GetSplitLogicId() string {
	if o == nil || common.IsNil(o.SplitLogicId) {
		var ret string
		return ret
	}
	return *o.SplitLogicId
}

// GetSplitLogicIdOk returns a tuple with the SplitLogicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfigurationLogic) GetSplitLogicIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.SplitLogicId) {
		return nil, false
	}
	return o.SplitLogicId, true
}

// HasSplitLogicId returns a boolean if a field has been set.
func (o *SplitConfigurationLogic) HasSplitLogicId() bool {
	if o != nil && !common.IsNil(o.SplitLogicId) {
		return true
	}

	return false
}

// SetSplitLogicId gets a reference to the given string and assigns it to the SplitLogicId field.
func (o *SplitConfigurationLogic) SetSplitLogicId(v string) {
	o.SplitLogicId = &v
}

// GetSurcharge returns the Surcharge field value if set, zero value otherwise.
func (o *SplitConfigurationLogic) GetSurcharge() string {
	if o == nil || common.IsNil(o.Surcharge) {
		var ret string
		return ret
	}
	return *o.Surcharge
}

// GetSurchargeOk returns a tuple with the Surcharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfigurationLogic) GetSurchargeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Surcharge) {
		return nil, false
	}
	return o.Surcharge, true
}

// HasSurcharge returns a boolean if a field has been set.
func (o *SplitConfigurationLogic) HasSurcharge() bool {
	if o != nil && !common.IsNil(o.Surcharge) {
		return true
	}

	return false
}

// SetSurcharge gets a reference to the given string and assigns it to the Surcharge field.
func (o *SplitConfigurationLogic) SetSurcharge(v string) {
	o.Surcharge = &v
}

// GetTip returns the Tip field value if set, zero value otherwise.
func (o *SplitConfigurationLogic) GetTip() string {
	if o == nil || common.IsNil(o.Tip) {
		var ret string
		return ret
	}
	return *o.Tip
}

// GetTipOk returns a tuple with the Tip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfigurationLogic) GetTipOk() (*string, bool) {
	if o == nil || common.IsNil(o.Tip) {
		return nil, false
	}
	return o.Tip, true
}

// HasTip returns a boolean if a field has been set.
func (o *SplitConfigurationLogic) HasTip() bool {
	if o != nil && !common.IsNil(o.Tip) {
		return true
	}

	return false
}

// SetTip gets a reference to the given string and assigns it to the Tip field.
func (o *SplitConfigurationLogic) SetTip(v string) {
	o.Tip = &v
}

func (o SplitConfigurationLogic) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SplitConfigurationLogic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AdditionalCommission) {
		toSerialize["additionalCommission"] = o.AdditionalCommission
	}
	if !common.IsNil(o.Chargeback) {
		toSerialize["chargeback"] = o.Chargeback
	}
	toSerialize["commission"] = o.Commission
	toSerialize["paymentFee"] = o.PaymentFee
	if !common.IsNil(o.Remainder) {
		toSerialize["remainder"] = o.Remainder
	}
	if !common.IsNil(o.SplitLogicId) {
		toSerialize["splitLogicId"] = o.SplitLogicId
	}
	if !common.IsNil(o.Surcharge) {
		toSerialize["surcharge"] = o.Surcharge
	}
	if !common.IsNil(o.Tip) {
		toSerialize["tip"] = o.Tip
	}
	return toSerialize, nil
}

type NullableSplitConfigurationLogic struct {
	value *SplitConfigurationLogic
	isSet bool
}

func (v NullableSplitConfigurationLogic) Get() *SplitConfigurationLogic {
	return v.value
}

func (v *NullableSplitConfigurationLogic) Set(val *SplitConfigurationLogic) {
	v.value = val
	v.isSet = true
}

func (v NullableSplitConfigurationLogic) IsSet() bool {
	return v.isSet
}

func (v *NullableSplitConfigurationLogic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSplitConfigurationLogic(val *SplitConfigurationLogic) *NullableSplitConfigurationLogic {
	return &NullableSplitConfigurationLogic{value: val, isSet: true}
}

func (v NullableSplitConfigurationLogic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSplitConfigurationLogic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *SplitConfigurationLogic) isValidChargeback() bool {
	var allowedEnumValues = []string{"deductFromLiableAccount", "deductFromOneBalanceAccount", "deductAccordingToSplitRatio"}
	for _, allowed := range allowedEnumValues {
		if o.GetChargeback() == allowed {
			return true
		}
	}
	return false
}
func (o *SplitConfigurationLogic) isValidPaymentFee() bool {
	var allowedEnumValues = []string{"deductFromLiableAccount", "deductFromOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetPaymentFee() == allowed {
			return true
		}
	}
	return false
}
func (o *SplitConfigurationLogic) isValidRemainder() bool {
	var allowedEnumValues = []string{"addToLiableAccount", "addToOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetRemainder() == allowed {
			return true
		}
	}
	return false
}
func (o *SplitConfigurationLogic) isValidSurcharge() bool {
	var allowedEnumValues = []string{"addToLiableAccount", "addToOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetSurcharge() == allowed {
			return true
		}
	}
	return false
}
func (o *SplitConfigurationLogic) isValidTip() bool {
	var allowedEnumValues = []string{"addToLiableAccount", "addToOneBalanceAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetTip() == allowed {
			return true
		}
	}
	return false
}