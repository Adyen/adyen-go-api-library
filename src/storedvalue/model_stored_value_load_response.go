/*
Adyen Stored Value API

API version: 46
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storedvalue

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the StoredValueLoadResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StoredValueLoadResponse{}

// StoredValueLoadResponse struct for StoredValueLoadResponse
type StoredValueLoadResponse struct {
	// Authorisation code: * When the payment is authorised, this field holds the authorisation code for the payment. * When the payment is not authorised, this field is empty.
	AuthCode       *string `json:"authCode,omitempty"`
	CurrentBalance *Amount `json:"currentBalance,omitempty"`
	// Adyen's 16-character string reference associated with the transaction/request. This value is globally unique; quote it when communicating with us about this request.
	PspReference *string `json:"pspReference,omitempty"`
	// If the transaction is refused or an error occurs, this field holds Adyen's mapped reason for the refusal or a description of the error.  When a transaction fails, the authorisation response includes `resultCode` and `refusalReason` values.
	RefusalReason *string `json:"refusalReason,omitempty"`
	// The result of the payment. Possible values:  * **Success** – The operation has been completed successfully.  * **Refused** – The operation was refused. The reason is given in the `refusalReason` field.  * **Error** – There was an error when the operation was processed. The reason is given in the `refusalReason` field.  * **NotEnoughBalance** – The amount on the payment method is lower than the amount given in the request. Only applicable to balance checks.
	ResultCode *string `json:"resultCode,omitempty"`
	// Raw refusal reason received from the third party, where available
	ThirdPartyRefusalReason *string `json:"thirdPartyRefusalReason,omitempty"`
}

// NewStoredValueLoadResponse instantiates a new StoredValueLoadResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoredValueLoadResponse() *StoredValueLoadResponse {
	this := StoredValueLoadResponse{}
	return &this
}

// NewStoredValueLoadResponseWithDefaults instantiates a new StoredValueLoadResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoredValueLoadResponseWithDefaults() *StoredValueLoadResponse {
	this := StoredValueLoadResponse{}
	return &this
}

// GetAuthCode returns the AuthCode field value if set, zero value otherwise.
func (o *StoredValueLoadResponse) GetAuthCode() string {
	if o == nil || common.IsNil(o.AuthCode) {
		var ret string
		return ret
	}
	return *o.AuthCode
}

// GetAuthCodeOk returns a tuple with the AuthCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredValueLoadResponse) GetAuthCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AuthCode) {
		return nil, false
	}
	return o.AuthCode, true
}

// HasAuthCode returns a boolean if a field has been set.
func (o *StoredValueLoadResponse) HasAuthCode() bool {
	if o != nil && !common.IsNil(o.AuthCode) {
		return true
	}

	return false
}

// SetAuthCode gets a reference to the given string and assigns it to the AuthCode field.
func (o *StoredValueLoadResponse) SetAuthCode(v string) {
	o.AuthCode = &v
}

// GetCurrentBalance returns the CurrentBalance field value if set, zero value otherwise.
func (o *StoredValueLoadResponse) GetCurrentBalance() Amount {
	if o == nil || common.IsNil(o.CurrentBalance) {
		var ret Amount
		return ret
	}
	return *o.CurrentBalance
}

// GetCurrentBalanceOk returns a tuple with the CurrentBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredValueLoadResponse) GetCurrentBalanceOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.CurrentBalance) {
		return nil, false
	}
	return o.CurrentBalance, true
}

// HasCurrentBalance returns a boolean if a field has been set.
func (o *StoredValueLoadResponse) HasCurrentBalance() bool {
	if o != nil && !common.IsNil(o.CurrentBalance) {
		return true
	}

	return false
}

// SetCurrentBalance gets a reference to the given Amount and assigns it to the CurrentBalance field.
func (o *StoredValueLoadResponse) SetCurrentBalance(v Amount) {
	o.CurrentBalance = &v
}

// GetPspReference returns the PspReference field value if set, zero value otherwise.
func (o *StoredValueLoadResponse) GetPspReference() string {
	if o == nil || common.IsNil(o.PspReference) {
		var ret string
		return ret
	}
	return *o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredValueLoadResponse) GetPspReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.PspReference) {
		return nil, false
	}
	return o.PspReference, true
}

// HasPspReference returns a boolean if a field has been set.
func (o *StoredValueLoadResponse) HasPspReference() bool {
	if o != nil && !common.IsNil(o.PspReference) {
		return true
	}

	return false
}

// SetPspReference gets a reference to the given string and assigns it to the PspReference field.
func (o *StoredValueLoadResponse) SetPspReference(v string) {
	o.PspReference = &v
}

// GetRefusalReason returns the RefusalReason field value if set, zero value otherwise.
func (o *StoredValueLoadResponse) GetRefusalReason() string {
	if o == nil || common.IsNil(o.RefusalReason) {
		var ret string
		return ret
	}
	return *o.RefusalReason
}

// GetRefusalReasonOk returns a tuple with the RefusalReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredValueLoadResponse) GetRefusalReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.RefusalReason) {
		return nil, false
	}
	return o.RefusalReason, true
}

// HasRefusalReason returns a boolean if a field has been set.
func (o *StoredValueLoadResponse) HasRefusalReason() bool {
	if o != nil && !common.IsNil(o.RefusalReason) {
		return true
	}

	return false
}

// SetRefusalReason gets a reference to the given string and assigns it to the RefusalReason field.
func (o *StoredValueLoadResponse) SetRefusalReason(v string) {
	o.RefusalReason = &v
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *StoredValueLoadResponse) GetResultCode() string {
	if o == nil || common.IsNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredValueLoadResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *StoredValueLoadResponse) HasResultCode() bool {
	if o != nil && !common.IsNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *StoredValueLoadResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetThirdPartyRefusalReason returns the ThirdPartyRefusalReason field value if set, zero value otherwise.
func (o *StoredValueLoadResponse) GetThirdPartyRefusalReason() string {
	if o == nil || common.IsNil(o.ThirdPartyRefusalReason) {
		var ret string
		return ret
	}
	return *o.ThirdPartyRefusalReason
}

// GetThirdPartyRefusalReasonOk returns a tuple with the ThirdPartyRefusalReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoredValueLoadResponse) GetThirdPartyRefusalReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThirdPartyRefusalReason) {
		return nil, false
	}
	return o.ThirdPartyRefusalReason, true
}

// HasThirdPartyRefusalReason returns a boolean if a field has been set.
func (o *StoredValueLoadResponse) HasThirdPartyRefusalReason() bool {
	if o != nil && !common.IsNil(o.ThirdPartyRefusalReason) {
		return true
	}

	return false
}

// SetThirdPartyRefusalReason gets a reference to the given string and assigns it to the ThirdPartyRefusalReason field.
func (o *StoredValueLoadResponse) SetThirdPartyRefusalReason(v string) {
	o.ThirdPartyRefusalReason = &v
}

func (o StoredValueLoadResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoredValueLoadResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AuthCode) {
		toSerialize["authCode"] = o.AuthCode
	}
	if !common.IsNil(o.CurrentBalance) {
		toSerialize["currentBalance"] = o.CurrentBalance
	}
	if !common.IsNil(o.PspReference) {
		toSerialize["pspReference"] = o.PspReference
	}
	if !common.IsNil(o.RefusalReason) {
		toSerialize["refusalReason"] = o.RefusalReason
	}
	if !common.IsNil(o.ResultCode) {
		toSerialize["resultCode"] = o.ResultCode
	}
	if !common.IsNil(o.ThirdPartyRefusalReason) {
		toSerialize["thirdPartyRefusalReason"] = o.ThirdPartyRefusalReason
	}
	return toSerialize, nil
}

type NullableStoredValueLoadResponse struct {
	value *StoredValueLoadResponse
	isSet bool
}

func (v NullableStoredValueLoadResponse) Get() *StoredValueLoadResponse {
	return v.value
}

func (v *NullableStoredValueLoadResponse) Set(val *StoredValueLoadResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStoredValueLoadResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStoredValueLoadResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoredValueLoadResponse(val *StoredValueLoadResponse) *NullableStoredValueLoadResponse {
	return &NullableStoredValueLoadResponse{value: val, isSet: true}
}

func (v NullableStoredValueLoadResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoredValueLoadResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *StoredValueLoadResponse) isValidResultCode() bool {
	var allowedEnumValues = []string{"Success", "Refused", "Error", "NotEnoughBalance"}
	for _, allowed := range allowedEnumValues {
		if o.GetResultCode() == allowed {
			return true
		}
	}
	return false
}
