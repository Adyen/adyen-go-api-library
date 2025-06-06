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

// checks if the PixRecurring type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PixRecurring{}

// PixRecurring struct for PixRecurring
type PixRecurring struct {
	// The date on which the shopper's payment method will be charged, in YYYY-MM-DD format.
	BillingDate *string `json:"billingDate,omitempty"`
	// End date of the billing plan, in YYYY-MM-DD format. The end date must align with the frequency and the start date of the billing plan. If left blank, the subscription will continue indefinitely unless it is cancelled by the shopper.
	EndsAt *string `json:"endsAt,omitempty"`
	// The frequency at which the shopper will be charged.
	Frequency *string `json:"frequency,omitempty"`
	MinAmount *Amount `json:"minAmount,omitempty"`
	// The pspReference for the failed recurring payment. Find this in AUTHORISATION webhook you received after the billing date.
	OriginalPspReference *string `json:"originalPspReference,omitempty"`
	RecurringAmount      *Amount `json:"recurringAmount,omitempty"`
	// The text that that will be shown on the shopper's bank statement for the recurring payments. We recommend to add a descriptive text about the subscription to let your shoppers recognize your recurring payments. Maximum length: 35 characters.
	RecurringStatement *string `json:"recurringStatement,omitempty"`
	// When set to true, you can retry for failed recurring payments. The default value is true.
	RetryPolicy *bool `json:"retryPolicy,omitempty"`
	// Start date of the billing plan, in YYYY-MM-DD format. The default value is the transaction date.
	StartsAt *string `json:"startsAt,omitempty"`
}

// NewPixRecurring instantiates a new PixRecurring object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPixRecurring() *PixRecurring {
	this := PixRecurring{}
	return &this
}

// NewPixRecurringWithDefaults instantiates a new PixRecurring object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPixRecurringWithDefaults() *PixRecurring {
	this := PixRecurring{}
	return &this
}

// GetBillingDate returns the BillingDate field value if set, zero value otherwise.
func (o *PixRecurring) GetBillingDate() string {
	if o == nil || common.IsNil(o.BillingDate) {
		var ret string
		return ret
	}
	return *o.BillingDate
}

// GetBillingDateOk returns a tuple with the BillingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PixRecurring) GetBillingDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.BillingDate) {
		return nil, false
	}
	return o.BillingDate, true
}

// HasBillingDate returns a boolean if a field has been set.
func (o *PixRecurring) HasBillingDate() bool {
	if o != nil && !common.IsNil(o.BillingDate) {
		return true
	}

	return false
}

// SetBillingDate gets a reference to the given string and assigns it to the BillingDate field.
func (o *PixRecurring) SetBillingDate(v string) {
	o.BillingDate = &v
}

// GetEndsAt returns the EndsAt field value if set, zero value otherwise.
func (o *PixRecurring) GetEndsAt() string {
	if o == nil || common.IsNil(o.EndsAt) {
		var ret string
		return ret
	}
	return *o.EndsAt
}

// GetEndsAtOk returns a tuple with the EndsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PixRecurring) GetEndsAtOk() (*string, bool) {
	if o == nil || common.IsNil(o.EndsAt) {
		return nil, false
	}
	return o.EndsAt, true
}

// HasEndsAt returns a boolean if a field has been set.
func (o *PixRecurring) HasEndsAt() bool {
	if o != nil && !common.IsNil(o.EndsAt) {
		return true
	}

	return false
}

// SetEndsAt gets a reference to the given string and assigns it to the EndsAt field.
func (o *PixRecurring) SetEndsAt(v string) {
	o.EndsAt = &v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *PixRecurring) GetFrequency() string {
	if o == nil || common.IsNil(o.Frequency) {
		var ret string
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PixRecurring) GetFrequencyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Frequency) {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *PixRecurring) HasFrequency() bool {
	if o != nil && !common.IsNil(o.Frequency) {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given string and assigns it to the Frequency field.
func (o *PixRecurring) SetFrequency(v string) {
	o.Frequency = &v
}

// GetMinAmount returns the MinAmount field value if set, zero value otherwise.
func (o *PixRecurring) GetMinAmount() Amount {
	if o == nil || common.IsNil(o.MinAmount) {
		var ret Amount
		return ret
	}
	return *o.MinAmount
}

// GetMinAmountOk returns a tuple with the MinAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PixRecurring) GetMinAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.MinAmount) {
		return nil, false
	}
	return o.MinAmount, true
}

// HasMinAmount returns a boolean if a field has been set.
func (o *PixRecurring) HasMinAmount() bool {
	if o != nil && !common.IsNil(o.MinAmount) {
		return true
	}

	return false
}

// SetMinAmount gets a reference to the given Amount and assigns it to the MinAmount field.
func (o *PixRecurring) SetMinAmount(v Amount) {
	o.MinAmount = &v
}

// GetOriginalPspReference returns the OriginalPspReference field value if set, zero value otherwise.
func (o *PixRecurring) GetOriginalPspReference() string {
	if o == nil || common.IsNil(o.OriginalPspReference) {
		var ret string
		return ret
	}
	return *o.OriginalPspReference
}

// GetOriginalPspReferenceOk returns a tuple with the OriginalPspReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PixRecurring) GetOriginalPspReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.OriginalPspReference) {
		return nil, false
	}
	return o.OriginalPspReference, true
}

// HasOriginalPspReference returns a boolean if a field has been set.
func (o *PixRecurring) HasOriginalPspReference() bool {
	if o != nil && !common.IsNil(o.OriginalPspReference) {
		return true
	}

	return false
}

// SetOriginalPspReference gets a reference to the given string and assigns it to the OriginalPspReference field.
func (o *PixRecurring) SetOriginalPspReference(v string) {
	o.OriginalPspReference = &v
}

// GetRecurringAmount returns the RecurringAmount field value if set, zero value otherwise.
func (o *PixRecurring) GetRecurringAmount() Amount {
	if o == nil || common.IsNil(o.RecurringAmount) {
		var ret Amount
		return ret
	}
	return *o.RecurringAmount
}

// GetRecurringAmountOk returns a tuple with the RecurringAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PixRecurring) GetRecurringAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.RecurringAmount) {
		return nil, false
	}
	return o.RecurringAmount, true
}

// HasRecurringAmount returns a boolean if a field has been set.
func (o *PixRecurring) HasRecurringAmount() bool {
	if o != nil && !common.IsNil(o.RecurringAmount) {
		return true
	}

	return false
}

// SetRecurringAmount gets a reference to the given Amount and assigns it to the RecurringAmount field.
func (o *PixRecurring) SetRecurringAmount(v Amount) {
	o.RecurringAmount = &v
}

// GetRecurringStatement returns the RecurringStatement field value if set, zero value otherwise.
func (o *PixRecurring) GetRecurringStatement() string {
	if o == nil || common.IsNil(o.RecurringStatement) {
		var ret string
		return ret
	}
	return *o.RecurringStatement
}

// GetRecurringStatementOk returns a tuple with the RecurringStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PixRecurring) GetRecurringStatementOk() (*string, bool) {
	if o == nil || common.IsNil(o.RecurringStatement) {
		return nil, false
	}
	return o.RecurringStatement, true
}

// HasRecurringStatement returns a boolean if a field has been set.
func (o *PixRecurring) HasRecurringStatement() bool {
	if o != nil && !common.IsNil(o.RecurringStatement) {
		return true
	}

	return false
}

// SetRecurringStatement gets a reference to the given string and assigns it to the RecurringStatement field.
func (o *PixRecurring) SetRecurringStatement(v string) {
	o.RecurringStatement = &v
}

// GetRetryPolicy returns the RetryPolicy field value if set, zero value otherwise.
func (o *PixRecurring) GetRetryPolicy() bool {
	if o == nil || common.IsNil(o.RetryPolicy) {
		var ret bool
		return ret
	}
	return *o.RetryPolicy
}

// GetRetryPolicyOk returns a tuple with the RetryPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PixRecurring) GetRetryPolicyOk() (*bool, bool) {
	if o == nil || common.IsNil(o.RetryPolicy) {
		return nil, false
	}
	return o.RetryPolicy, true
}

// HasRetryPolicy returns a boolean if a field has been set.
func (o *PixRecurring) HasRetryPolicy() bool {
	if o != nil && !common.IsNil(o.RetryPolicy) {
		return true
	}

	return false
}

// SetRetryPolicy gets a reference to the given bool and assigns it to the RetryPolicy field.
func (o *PixRecurring) SetRetryPolicy(v bool) {
	o.RetryPolicy = &v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise.
func (o *PixRecurring) GetStartsAt() string {
	if o == nil || common.IsNil(o.StartsAt) {
		var ret string
		return ret
	}
	return *o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PixRecurring) GetStartsAtOk() (*string, bool) {
	if o == nil || common.IsNil(o.StartsAt) {
		return nil, false
	}
	return o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *PixRecurring) HasStartsAt() bool {
	if o != nil && !common.IsNil(o.StartsAt) {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given string and assigns it to the StartsAt field.
func (o *PixRecurring) SetStartsAt(v string) {
	o.StartsAt = &v
}

func (o PixRecurring) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PixRecurring) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BillingDate) {
		toSerialize["billingDate"] = o.BillingDate
	}
	if !common.IsNil(o.EndsAt) {
		toSerialize["endsAt"] = o.EndsAt
	}
	if !common.IsNil(o.Frequency) {
		toSerialize["frequency"] = o.Frequency
	}
	if !common.IsNil(o.MinAmount) {
		toSerialize["minAmount"] = o.MinAmount
	}
	if !common.IsNil(o.OriginalPspReference) {
		toSerialize["originalPspReference"] = o.OriginalPspReference
	}
	if !common.IsNil(o.RecurringAmount) {
		toSerialize["recurringAmount"] = o.RecurringAmount
	}
	if !common.IsNil(o.RecurringStatement) {
		toSerialize["recurringStatement"] = o.RecurringStatement
	}
	if !common.IsNil(o.RetryPolicy) {
		toSerialize["retryPolicy"] = o.RetryPolicy
	}
	if !common.IsNil(o.StartsAt) {
		toSerialize["startsAt"] = o.StartsAt
	}
	return toSerialize, nil
}

type NullablePixRecurring struct {
	value *PixRecurring
	isSet bool
}

func (v NullablePixRecurring) Get() *PixRecurring {
	return v.value
}

func (v *NullablePixRecurring) Set(val *PixRecurring) {
	v.value = val
	v.isSet = true
}

func (v NullablePixRecurring) IsSet() bool {
	return v.isSet
}

func (v *NullablePixRecurring) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePixRecurring(val *PixRecurring) *NullablePixRecurring {
	return &NullablePixRecurring{value: val, isSet: true}
}

func (v NullablePixRecurring) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePixRecurring) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PixRecurring) isValidFrequency() bool {
	var allowedEnumValues = []string{"weekly", "monthly", "quarterly", "half-yearly", "yearly"}
	for _, allowed := range allowedEnumValues {
		if o.GetFrequency() == allowed {
			return true
		}
	}
	return false
}
