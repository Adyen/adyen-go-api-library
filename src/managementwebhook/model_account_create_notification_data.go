/*
Management Webhooks

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package managementwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the AccountCreateNotificationData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountCreateNotificationData{}

// AccountCreateNotificationData struct for AccountCreateNotificationData
type AccountCreateNotificationData struct {
	// Key-value pairs that specify the actions that the merchant account can do and its settings. The key is a capability. For example, the **sendToTransferInstrument** is the capability required before you can pay out funds to the bank account. The value is an object containing the settings for the capability.
	Capabilities map[string]AccountCapabilityData `json:"capabilities"`
	// The unique identifier of the company account.
	CompanyId string `json:"companyId"`
	// The unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/legalentity/latest/post/legalEntities#responses-200-id).
	LegalEntityId *string `json:"legalEntityId,omitempty"`
	// The unique identifier of the merchant account.
	MerchantId string `json:"merchantId"`
	// The status of the merchant account.  Possible values:  * **PreActive**: The merchant account has been created. Users cannot access the merchant account in the Customer Area. The account cannot process payments. * **Active**: Users can access the merchant account in the Customer Area. If the company account is also **Active**, then payment processing and payouts are enabled. * **InactiveWithModifications**: Users can access the merchant account in the Customer Area. The account cannot process new payments but can still modify payments, for example issue refunds. The account can still receive payouts. * **Inactive**: Users can access the merchant account in the Customer Area. Payment processing and payouts are disabled. * **Closed**: The account is closed and this cannot be reversed. Users cannot log in. Payment processing and payouts are disabled.
	Status string `json:"status"`
}

// NewAccountCreateNotificationData instantiates a new AccountCreateNotificationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCreateNotificationData(capabilities map[string]AccountCapabilityData, companyId string, merchantId string, status string) *AccountCreateNotificationData {
	this := AccountCreateNotificationData{}
	this.Capabilities = capabilities
	this.CompanyId = companyId
	this.MerchantId = merchantId
	this.Status = status
	return &this
}

// NewAccountCreateNotificationDataWithDefaults instantiates a new AccountCreateNotificationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCreateNotificationDataWithDefaults() *AccountCreateNotificationData {
	this := AccountCreateNotificationData{}
	return &this
}

// GetCapabilities returns the Capabilities field value
func (o *AccountCreateNotificationData) GetCapabilities() map[string]AccountCapabilityData {
	if o == nil {
		var ret map[string]AccountCapabilityData
		return ret
	}

	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *AccountCreateNotificationData) GetCapabilitiesOk() (*map[string]AccountCapabilityData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capabilities, true
}

// SetCapabilities sets field value
func (o *AccountCreateNotificationData) SetCapabilities(v map[string]AccountCapabilityData) {
	o.Capabilities = v
}

// GetCompanyId returns the CompanyId field value
func (o *AccountCreateNotificationData) GetCompanyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *AccountCreateNotificationData) GetCompanyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *AccountCreateNotificationData) SetCompanyId(v string) {
	o.CompanyId = v
}

// GetLegalEntityId returns the LegalEntityId field value if set, zero value otherwise.
func (o *AccountCreateNotificationData) GetLegalEntityId() string {
	if o == nil || common.IsNil(o.LegalEntityId) {
		var ret string
		return ret
	}
	return *o.LegalEntityId
}

// GetLegalEntityIdOk returns a tuple with the LegalEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateNotificationData) GetLegalEntityIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.LegalEntityId) {
		return nil, false
	}
	return o.LegalEntityId, true
}

// HasLegalEntityId returns a boolean if a field has been set.
func (o *AccountCreateNotificationData) HasLegalEntityId() bool {
	if o != nil && !common.IsNil(o.LegalEntityId) {
		return true
	}

	return false
}

// SetLegalEntityId gets a reference to the given string and assigns it to the LegalEntityId field.
func (o *AccountCreateNotificationData) SetLegalEntityId(v string) {
	o.LegalEntityId = &v
}

// GetMerchantId returns the MerchantId field value
func (o *AccountCreateNotificationData) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *AccountCreateNotificationData) GetMerchantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *AccountCreateNotificationData) SetMerchantId(v string) {
	o.MerchantId = v
}

// GetStatus returns the Status field value
func (o *AccountCreateNotificationData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AccountCreateNotificationData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AccountCreateNotificationData) SetStatus(v string) {
	o.Status = v
}

func (o AccountCreateNotificationData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountCreateNotificationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["capabilities"] = o.Capabilities
	toSerialize["companyId"] = o.CompanyId
	if !common.IsNil(o.LegalEntityId) {
		toSerialize["legalEntityId"] = o.LegalEntityId
	}
	toSerialize["merchantId"] = o.MerchantId
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableAccountCreateNotificationData struct {
	value *AccountCreateNotificationData
	isSet bool
}

func (v NullableAccountCreateNotificationData) Get() *AccountCreateNotificationData {
	return v.value
}

func (v *NullableAccountCreateNotificationData) Set(val *AccountCreateNotificationData) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCreateNotificationData) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCreateNotificationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCreateNotificationData(val *AccountCreateNotificationData) *NullableAccountCreateNotificationData {
	return &NullableAccountCreateNotificationData{value: val, isSet: true}
}

func (v NullableAccountCreateNotificationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCreateNotificationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
