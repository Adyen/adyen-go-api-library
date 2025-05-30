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

// checks if the PixDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PixDetails{}

// PixDetails struct for PixDetails
type PixDetails struct {
	// The checkout attempt identifier.
	CheckoutAttemptId *string       `json:"checkoutAttemptId,omitempty"`
	PixRecurring      *PixRecurring `json:"pixRecurring,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	// Deprecated since Adyen Checkout API v49
	// Use `storedPaymentMethodId` instead.
	RecurringDetailReference *string `json:"recurringDetailReference,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	StoredPaymentMethodId *string `json:"storedPaymentMethodId,omitempty"`
	// The payment method type.
	Type *string `json:"type,omitempty"`
}

// NewPixDetails instantiates a new PixDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPixDetails() *PixDetails {
	this := PixDetails{}
	return &this
}

// NewPixDetailsWithDefaults instantiates a new PixDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPixDetailsWithDefaults() *PixDetails {
	this := PixDetails{}
	return &this
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *PixDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PixDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *PixDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *PixDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetPixRecurring returns the PixRecurring field value if set, zero value otherwise.
func (o *PixDetails) GetPixRecurring() PixRecurring {
	if o == nil || common.IsNil(o.PixRecurring) {
		var ret PixRecurring
		return ret
	}
	return *o.PixRecurring
}

// GetPixRecurringOk returns a tuple with the PixRecurring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PixDetails) GetPixRecurringOk() (*PixRecurring, bool) {
	if o == nil || common.IsNil(o.PixRecurring) {
		return nil, false
	}
	return o.PixRecurring, true
}

// HasPixRecurring returns a boolean if a field has been set.
func (o *PixDetails) HasPixRecurring() bool {
	if o != nil && !common.IsNil(o.PixRecurring) {
		return true
	}

	return false
}

// SetPixRecurring gets a reference to the given PixRecurring and assigns it to the PixRecurring field.
func (o *PixDetails) SetPixRecurring(v PixRecurring) {
	o.PixRecurring = &v
}

// GetRecurringDetailReference returns the RecurringDetailReference field value if set, zero value otherwise.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *PixDetails) GetRecurringDetailReference() string {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		var ret string
		return ret
	}
	return *o.RecurringDetailReference
}

// GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *PixDetails) GetRecurringDetailReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		return nil, false
	}
	return o.RecurringDetailReference, true
}

// HasRecurringDetailReference returns a boolean if a field has been set.
func (o *PixDetails) HasRecurringDetailReference() bool {
	if o != nil && !common.IsNil(o.RecurringDetailReference) {
		return true
	}

	return false
}

// SetRecurringDetailReference gets a reference to the given string and assigns it to the RecurringDetailReference field.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *PixDetails) SetRecurringDetailReference(v string) {
	o.RecurringDetailReference = &v
}

// GetStoredPaymentMethodId returns the StoredPaymentMethodId field value if set, zero value otherwise.
func (o *PixDetails) GetStoredPaymentMethodId() string {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		var ret string
		return ret
	}
	return *o.StoredPaymentMethodId
}

// GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PixDetails) GetStoredPaymentMethodIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		return nil, false
	}
	return o.StoredPaymentMethodId, true
}

// HasStoredPaymentMethodId returns a boolean if a field has been set.
func (o *PixDetails) HasStoredPaymentMethodId() bool {
	if o != nil && !common.IsNil(o.StoredPaymentMethodId) {
		return true
	}

	return false
}

// SetStoredPaymentMethodId gets a reference to the given string and assigns it to the StoredPaymentMethodId field.
func (o *PixDetails) SetStoredPaymentMethodId(v string) {
	o.StoredPaymentMethodId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PixDetails) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PixDetails) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PixDetails) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PixDetails) SetType(v string) {
	o.Type = &v
}

func (o PixDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PixDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	if !common.IsNil(o.PixRecurring) {
		toSerialize["pixRecurring"] = o.PixRecurring
	}
	if !common.IsNil(o.RecurringDetailReference) {
		toSerialize["recurringDetailReference"] = o.RecurringDetailReference
	}
	if !common.IsNil(o.StoredPaymentMethodId) {
		toSerialize["storedPaymentMethodId"] = o.StoredPaymentMethodId
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullablePixDetails struct {
	value *PixDetails
	isSet bool
}

func (v NullablePixDetails) Get() *PixDetails {
	return v.value
}

func (v *NullablePixDetails) Set(val *PixDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePixDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePixDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePixDetails(val *PixDetails) *NullablePixDetails {
	return &NullablePixDetails{value: val, isSet: true}
}

func (v NullablePixDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePixDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PixDetails) isValidType() bool {
	var allowedEnumValues = []string{"pix"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
