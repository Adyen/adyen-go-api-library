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

// checks if the CellulantDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CellulantDetails{}

// CellulantDetails struct for CellulantDetails
type CellulantDetails struct {
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// The Cellulant issuer.
	Issuer *string `json:"issuer,omitempty"`
	// **Cellulant**
	Type *string `json:"type,omitempty"`
}

// NewCellulantDetails instantiates a new CellulantDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCellulantDetails() *CellulantDetails {
	this := CellulantDetails{}
	var type_ string = "cellulant"
	this.Type = &type_
	return &this
}

// NewCellulantDetailsWithDefaults instantiates a new CellulantDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCellulantDetailsWithDefaults() *CellulantDetails {
	this := CellulantDetails{}
	var type_ string = "cellulant"
	this.Type = &type_
	return &this
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *CellulantDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellulantDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *CellulantDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *CellulantDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *CellulantDetails) GetIssuer() string {
	if o == nil || common.IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellulantDetails) GetIssuerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *CellulantDetails) HasIssuer() bool {
	if o != nil && !common.IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *CellulantDetails) SetIssuer(v string) {
	o.Issuer = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CellulantDetails) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellulantDetails) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CellulantDetails) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CellulantDetails) SetType(v string) {
	o.Type = &v
}

func (o CellulantDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CellulantDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	if !common.IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableCellulantDetails struct {
	value *CellulantDetails
	isSet bool
}

func (v NullableCellulantDetails) Get() *CellulantDetails {
	return v.value
}

func (v *NullableCellulantDetails) Set(val *CellulantDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCellulantDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCellulantDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCellulantDetails(val *CellulantDetails) *NullableCellulantDetails {
	return &NullableCellulantDetails{value: val, isSet: true}
}

func (v NullableCellulantDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCellulantDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CellulantDetails) isValidType() bool {
	var allowedEnumValues = []string{"cellulant"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
