/*
Disputes API

API version: 30
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package disputes

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the SupplyDefenseDocumentRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SupplyDefenseDocumentRequest{}

// SupplyDefenseDocumentRequest struct for SupplyDefenseDocumentRequest
type SupplyDefenseDocumentRequest struct {
	// An array containing a list of the defense documents.
	DefenseDocuments []DefenseDocument `json:"defenseDocuments"`
	// The PSP reference assigned to the dispute.
	DisputePspReference string `json:"disputePspReference"`
	// The merchant account identifier, for which you want to process the dispute transaction.
	MerchantAccountCode string `json:"merchantAccountCode"`
}

// NewSupplyDefenseDocumentRequest instantiates a new SupplyDefenseDocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupplyDefenseDocumentRequest(defenseDocuments []DefenseDocument, disputePspReference string, merchantAccountCode string) *SupplyDefenseDocumentRequest {
	this := SupplyDefenseDocumentRequest{}
	this.DefenseDocuments = defenseDocuments
	this.DisputePspReference = disputePspReference
	this.MerchantAccountCode = merchantAccountCode
	return &this
}

// NewSupplyDefenseDocumentRequestWithDefaults instantiates a new SupplyDefenseDocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupplyDefenseDocumentRequestWithDefaults() *SupplyDefenseDocumentRequest {
	this := SupplyDefenseDocumentRequest{}
	return &this
}

// GetDefenseDocuments returns the DefenseDocuments field value
func (o *SupplyDefenseDocumentRequest) GetDefenseDocuments() []DefenseDocument {
	if o == nil {
		var ret []DefenseDocument
		return ret
	}

	return o.DefenseDocuments
}

// GetDefenseDocumentsOk returns a tuple with the DefenseDocuments field value
// and a boolean to check if the value has been set.
func (o *SupplyDefenseDocumentRequest) GetDefenseDocumentsOk() ([]DefenseDocument, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefenseDocuments, true
}

// SetDefenseDocuments sets field value
func (o *SupplyDefenseDocumentRequest) SetDefenseDocuments(v []DefenseDocument) {
	o.DefenseDocuments = v
}

// GetDisputePspReference returns the DisputePspReference field value
func (o *SupplyDefenseDocumentRequest) GetDisputePspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisputePspReference
}

// GetDisputePspReferenceOk returns a tuple with the DisputePspReference field value
// and a boolean to check if the value has been set.
func (o *SupplyDefenseDocumentRequest) GetDisputePspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisputePspReference, true
}

// SetDisputePspReference sets field value
func (o *SupplyDefenseDocumentRequest) SetDisputePspReference(v string) {
	o.DisputePspReference = v
}

// GetMerchantAccountCode returns the MerchantAccountCode field value
func (o *SupplyDefenseDocumentRequest) GetMerchantAccountCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccountCode
}

// GetMerchantAccountCodeOk returns a tuple with the MerchantAccountCode field value
// and a boolean to check if the value has been set.
func (o *SupplyDefenseDocumentRequest) GetMerchantAccountCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccountCode, true
}

// SetMerchantAccountCode sets field value
func (o *SupplyDefenseDocumentRequest) SetMerchantAccountCode(v string) {
	o.MerchantAccountCode = v
}

func (o SupplyDefenseDocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupplyDefenseDocumentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["defenseDocuments"] = o.DefenseDocuments
	toSerialize["disputePspReference"] = o.DisputePspReference
	toSerialize["merchantAccountCode"] = o.MerchantAccountCode
	return toSerialize, nil
}

type NullableSupplyDefenseDocumentRequest struct {
	value *SupplyDefenseDocumentRequest
	isSet bool
}

func (v NullableSupplyDefenseDocumentRequest) Get() *SupplyDefenseDocumentRequest {
	return v.value
}

func (v *NullableSupplyDefenseDocumentRequest) Set(val *SupplyDefenseDocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSupplyDefenseDocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSupplyDefenseDocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupplyDefenseDocumentRequest(val *SupplyDefenseDocumentRequest) *NullableSupplyDefenseDocumentRequest {
	return &NullableSupplyDefenseDocumentRequest{value: val, isSet: true}
}

func (v NullableSupplyDefenseDocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupplyDefenseDocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
