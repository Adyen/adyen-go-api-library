/*
Accounting notifications

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatformaccountingnotification

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// checks if the CALocalAccountIdentification type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CALocalAccountIdentification{}

// CALocalAccountIdentification struct for CALocalAccountIdentification
type CALocalAccountIdentification struct {
	// The 5- to 12-digit bank account number, without separators or whitespace.
	AccountNumber string `json:"accountNumber"`
	// The 3-digit institution number, without separators or whitespace.
	InstitutionNumber string `json:"institutionNumber"`
	// The 5-digit transit number, without separators or whitespace.
	TransitNumber string `json:"transitNumber"`
	// **caLocal**
	Type string `json:"type"`
}

// NewCALocalAccountIdentification instantiates a new CALocalAccountIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCALocalAccountIdentification(accountNumber string, institutionNumber string, transitNumber string, type_ string) *CALocalAccountIdentification {
	this := CALocalAccountIdentification{}
	this.AccountNumber = accountNumber
	this.InstitutionNumber = institutionNumber
	this.TransitNumber = transitNumber
	this.Type = type_
	return &this
}

// NewCALocalAccountIdentificationWithDefaults instantiates a new CALocalAccountIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCALocalAccountIdentificationWithDefaults() *CALocalAccountIdentification {
	this := CALocalAccountIdentification{}
	var type_ string = "caLocal"
	this.Type = type_
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *CALocalAccountIdentification) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *CALocalAccountIdentification) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *CALocalAccountIdentification) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetInstitutionNumber returns the InstitutionNumber field value
func (o *CALocalAccountIdentification) GetInstitutionNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstitutionNumber
}

// GetInstitutionNumberOk returns a tuple with the InstitutionNumber field value
// and a boolean to check if the value has been set.
func (o *CALocalAccountIdentification) GetInstitutionNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstitutionNumber, true
}

// SetInstitutionNumber sets field value
func (o *CALocalAccountIdentification) SetInstitutionNumber(v string) {
	o.InstitutionNumber = v
}

// GetTransitNumber returns the TransitNumber field value
func (o *CALocalAccountIdentification) GetTransitNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransitNumber
}

// GetTransitNumberOk returns a tuple with the TransitNumber field value
// and a boolean to check if the value has been set.
func (o *CALocalAccountIdentification) GetTransitNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransitNumber, true
}

// SetTransitNumber sets field value
func (o *CALocalAccountIdentification) SetTransitNumber(v string) {
	o.TransitNumber = v
}

// GetType returns the Type field value
func (o *CALocalAccountIdentification) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CALocalAccountIdentification) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CALocalAccountIdentification) SetType(v string) {
	o.Type = v
}

func (o CALocalAccountIdentification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CALocalAccountIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountNumber"] = o.AccountNumber
	toSerialize["institutionNumber"] = o.InstitutionNumber
	toSerialize["transitNumber"] = o.TransitNumber
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableCALocalAccountIdentification struct {
	value *CALocalAccountIdentification
	isSet bool
}

func (v NullableCALocalAccountIdentification) Get() *CALocalAccountIdentification {
	return v.value
}

func (v *NullableCALocalAccountIdentification) Set(val *CALocalAccountIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableCALocalAccountIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableCALocalAccountIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCALocalAccountIdentification(val *CALocalAccountIdentification) *NullableCALocalAccountIdentification {
	return &NullableCALocalAccountIdentification{value: val, isSet: true}
}

func (v NullableCALocalAccountIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCALocalAccountIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CALocalAccountIdentification) isValidType() bool {
	var allowedEnumValues = []string{"caLocal"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}