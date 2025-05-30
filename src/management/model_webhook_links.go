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

// checks if the WebhookLinks type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &WebhookLinks{}

// WebhookLinks struct for WebhookLinks
type WebhookLinks struct {
	Company      *LinksElement `json:"company,omitempty"`
	GenerateHmac LinksElement  `json:"generateHmac"`
	Merchant     *LinksElement `json:"merchant,omitempty"`
	Self         LinksElement  `json:"self"`
	TestWebhook  LinksElement  `json:"testWebhook"`
}

// NewWebhookLinks instantiates a new WebhookLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookLinks(generateHmac LinksElement, self LinksElement, testWebhook LinksElement) *WebhookLinks {
	this := WebhookLinks{}
	this.GenerateHmac = generateHmac
	this.Self = self
	this.TestWebhook = testWebhook
	return &this
}

// NewWebhookLinksWithDefaults instantiates a new WebhookLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookLinksWithDefaults() *WebhookLinks {
	this := WebhookLinks{}
	return &this
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *WebhookLinks) GetCompany() LinksElement {
	if o == nil || common.IsNil(o.Company) {
		var ret LinksElement
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookLinks) GetCompanyOk() (*LinksElement, bool) {
	if o == nil || common.IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *WebhookLinks) HasCompany() bool {
	if o != nil && !common.IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given LinksElement and assigns it to the Company field.
func (o *WebhookLinks) SetCompany(v LinksElement) {
	o.Company = &v
}

// GetGenerateHmac returns the GenerateHmac field value
func (o *WebhookLinks) GetGenerateHmac() LinksElement {
	if o == nil {
		var ret LinksElement
		return ret
	}

	return o.GenerateHmac
}

// GetGenerateHmacOk returns a tuple with the GenerateHmac field value
// and a boolean to check if the value has been set.
func (o *WebhookLinks) GetGenerateHmacOk() (*LinksElement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GenerateHmac, true
}

// SetGenerateHmac sets field value
func (o *WebhookLinks) SetGenerateHmac(v LinksElement) {
	o.GenerateHmac = v
}

// GetMerchant returns the Merchant field value if set, zero value otherwise.
func (o *WebhookLinks) GetMerchant() LinksElement {
	if o == nil || common.IsNil(o.Merchant) {
		var ret LinksElement
		return ret
	}
	return *o.Merchant
}

// GetMerchantOk returns a tuple with the Merchant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookLinks) GetMerchantOk() (*LinksElement, bool) {
	if o == nil || common.IsNil(o.Merchant) {
		return nil, false
	}
	return o.Merchant, true
}

// HasMerchant returns a boolean if a field has been set.
func (o *WebhookLinks) HasMerchant() bool {
	if o != nil && !common.IsNil(o.Merchant) {
		return true
	}

	return false
}

// SetMerchant gets a reference to the given LinksElement and assigns it to the Merchant field.
func (o *WebhookLinks) SetMerchant(v LinksElement) {
	o.Merchant = &v
}

// GetSelf returns the Self field value
func (o *WebhookLinks) GetSelf() LinksElement {
	if o == nil {
		var ret LinksElement
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *WebhookLinks) GetSelfOk() (*LinksElement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *WebhookLinks) SetSelf(v LinksElement) {
	o.Self = v
}

// GetTestWebhook returns the TestWebhook field value
func (o *WebhookLinks) GetTestWebhook() LinksElement {
	if o == nil {
		var ret LinksElement
		return ret
	}

	return o.TestWebhook
}

// GetTestWebhookOk returns a tuple with the TestWebhook field value
// and a boolean to check if the value has been set.
func (o *WebhookLinks) GetTestWebhookOk() (*LinksElement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestWebhook, true
}

// SetTestWebhook sets field value
func (o *WebhookLinks) SetTestWebhook(v LinksElement) {
	o.TestWebhook = v
}

func (o WebhookLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	toSerialize["generateHmac"] = o.GenerateHmac
	if !common.IsNil(o.Merchant) {
		toSerialize["merchant"] = o.Merchant
	}
	toSerialize["self"] = o.Self
	toSerialize["testWebhook"] = o.TestWebhook
	return toSerialize, nil
}

type NullableWebhookLinks struct {
	value *WebhookLinks
	isSet bool
}

func (v NullableWebhookLinks) Get() *WebhookLinks {
	return v.value
}

func (v *NullableWebhookLinks) Set(val *WebhookLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookLinks(val *WebhookLinks) *NullableWebhookLinks {
	return &NullableWebhookLinks{value: val, isSet: true}
}

func (v NullableWebhookLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
