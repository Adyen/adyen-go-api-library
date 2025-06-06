/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the BusinessLineInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BusinessLineInfo{}

// BusinessLineInfo struct for BusinessLineInfo
type BusinessLineInfo struct {
	// The capability for which you are creating the business line.  Possible values: **receivePayments**, **receiveFromPlatformPayments**, **issueBankAccount**
	// Deprecated since Legal Entity Management API v3
	// Use `service` instead.
	Capability *string `json:"capability,omitempty"`
	// A code that represents the industry of the legal entity for [marketplaces](https://docs.adyen.com/marketplaces/verification-requirements/reference-additional-products/#list-industry-codes) or [platforms](https://docs.adyen.com/platforms/verification-requirements/reference-additional-products/#list-industry-codes). For example, **4431A** for computer software stores.
	IndustryCode string `json:"industryCode"`
	// Unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/legalEntities__resParam_id) that owns the business line.
	LegalEntityId string `json:"legalEntityId"`
	// A list of channels where goods or services are sold.  Possible values: **pos**, **posMoto**, **eCommerce**, **ecomMoto**, **payByLink**.  Required only in combination with the `service` **paymentProcessing**.
	SalesChannels []string `json:"salesChannels,omitempty"`
	// The service for which you are creating the business line.    Possible values: *  **paymentProcessing** *  **banking**
	Service       string         `json:"service"`
	SourceOfFunds *SourceOfFunds `json:"sourceOfFunds,omitempty"`
	// List of website URLs where your user's goods or services are sold. When this is required for a service but your user does not have an online presence, provide the reason in the `webDataExemption` object.
	WebData          []WebData         `json:"webData,omitempty"`
	WebDataExemption *WebDataExemption `json:"webDataExemption,omitempty"`
}

// NewBusinessLineInfo instantiates a new BusinessLineInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessLineInfo(industryCode string, legalEntityId string, service string) *BusinessLineInfo {
	this := BusinessLineInfo{}
	this.IndustryCode = industryCode
	this.LegalEntityId = legalEntityId
	this.Service = service
	return &this
}

// NewBusinessLineInfoWithDefaults instantiates a new BusinessLineInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessLineInfoWithDefaults() *BusinessLineInfo {
	this := BusinessLineInfo{}
	return &this
}

// GetCapability returns the Capability field value if set, zero value otherwise.
// Deprecated since Legal Entity Management API v3
// Use `service` instead.
func (o *BusinessLineInfo) GetCapability() string {
	if o == nil || common.IsNil(o.Capability) {
		var ret string
		return ret
	}
	return *o.Capability
}

// GetCapabilityOk returns a tuple with the Capability field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated since Legal Entity Management API v3
// Use `service` instead.
func (o *BusinessLineInfo) GetCapabilityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Capability) {
		return nil, false
	}
	return o.Capability, true
}

// HasCapability returns a boolean if a field has been set.
func (o *BusinessLineInfo) HasCapability() bool {
	if o != nil && !common.IsNil(o.Capability) {
		return true
	}

	return false
}

// SetCapability gets a reference to the given string and assigns it to the Capability field.
// Deprecated since Legal Entity Management API v3
// Use `service` instead.
func (o *BusinessLineInfo) SetCapability(v string) {
	o.Capability = &v
}

// GetIndustryCode returns the IndustryCode field value
func (o *BusinessLineInfo) GetIndustryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IndustryCode
}

// GetIndustryCodeOk returns a tuple with the IndustryCode field value
// and a boolean to check if the value has been set.
func (o *BusinessLineInfo) GetIndustryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndustryCode, true
}

// SetIndustryCode sets field value
func (o *BusinessLineInfo) SetIndustryCode(v string) {
	o.IndustryCode = v
}

// GetLegalEntityId returns the LegalEntityId field value
func (o *BusinessLineInfo) GetLegalEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LegalEntityId
}

// GetLegalEntityIdOk returns a tuple with the LegalEntityId field value
// and a boolean to check if the value has been set.
func (o *BusinessLineInfo) GetLegalEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LegalEntityId, true
}

// SetLegalEntityId sets field value
func (o *BusinessLineInfo) SetLegalEntityId(v string) {
	o.LegalEntityId = v
}

// GetSalesChannels returns the SalesChannels field value if set, zero value otherwise.
func (o *BusinessLineInfo) GetSalesChannels() []string {
	if o == nil || common.IsNil(o.SalesChannels) {
		var ret []string
		return ret
	}
	return o.SalesChannels
}

// GetSalesChannelsOk returns a tuple with the SalesChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessLineInfo) GetSalesChannelsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.SalesChannels) {
		return nil, false
	}
	return o.SalesChannels, true
}

// HasSalesChannels returns a boolean if a field has been set.
func (o *BusinessLineInfo) HasSalesChannels() bool {
	if o != nil && !common.IsNil(o.SalesChannels) {
		return true
	}

	return false
}

// SetSalesChannels gets a reference to the given []string and assigns it to the SalesChannels field.
func (o *BusinessLineInfo) SetSalesChannels(v []string) {
	o.SalesChannels = v
}

// GetService returns the Service field value
func (o *BusinessLineInfo) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *BusinessLineInfo) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *BusinessLineInfo) SetService(v string) {
	o.Service = v
}

// GetSourceOfFunds returns the SourceOfFunds field value if set, zero value otherwise.
func (o *BusinessLineInfo) GetSourceOfFunds() SourceOfFunds {
	if o == nil || common.IsNil(o.SourceOfFunds) {
		var ret SourceOfFunds
		return ret
	}
	return *o.SourceOfFunds
}

// GetSourceOfFundsOk returns a tuple with the SourceOfFunds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessLineInfo) GetSourceOfFundsOk() (*SourceOfFunds, bool) {
	if o == nil || common.IsNil(o.SourceOfFunds) {
		return nil, false
	}
	return o.SourceOfFunds, true
}

// HasSourceOfFunds returns a boolean if a field has been set.
func (o *BusinessLineInfo) HasSourceOfFunds() bool {
	if o != nil && !common.IsNil(o.SourceOfFunds) {
		return true
	}

	return false
}

// SetSourceOfFunds gets a reference to the given SourceOfFunds and assigns it to the SourceOfFunds field.
func (o *BusinessLineInfo) SetSourceOfFunds(v SourceOfFunds) {
	o.SourceOfFunds = &v
}

// GetWebData returns the WebData field value if set, zero value otherwise.
func (o *BusinessLineInfo) GetWebData() []WebData {
	if o == nil || common.IsNil(o.WebData) {
		var ret []WebData
		return ret
	}
	return o.WebData
}

// GetWebDataOk returns a tuple with the WebData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessLineInfo) GetWebDataOk() ([]WebData, bool) {
	if o == nil || common.IsNil(o.WebData) {
		return nil, false
	}
	return o.WebData, true
}

// HasWebData returns a boolean if a field has been set.
func (o *BusinessLineInfo) HasWebData() bool {
	if o != nil && !common.IsNil(o.WebData) {
		return true
	}

	return false
}

// SetWebData gets a reference to the given []WebData and assigns it to the WebData field.
func (o *BusinessLineInfo) SetWebData(v []WebData) {
	o.WebData = v
}

// GetWebDataExemption returns the WebDataExemption field value if set, zero value otherwise.
func (o *BusinessLineInfo) GetWebDataExemption() WebDataExemption {
	if o == nil || common.IsNil(o.WebDataExemption) {
		var ret WebDataExemption
		return ret
	}
	return *o.WebDataExemption
}

// GetWebDataExemptionOk returns a tuple with the WebDataExemption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessLineInfo) GetWebDataExemptionOk() (*WebDataExemption, bool) {
	if o == nil || common.IsNil(o.WebDataExemption) {
		return nil, false
	}
	return o.WebDataExemption, true
}

// HasWebDataExemption returns a boolean if a field has been set.
func (o *BusinessLineInfo) HasWebDataExemption() bool {
	if o != nil && !common.IsNil(o.WebDataExemption) {
		return true
	}

	return false
}

// SetWebDataExemption gets a reference to the given WebDataExemption and assigns it to the WebDataExemption field.
func (o *BusinessLineInfo) SetWebDataExemption(v WebDataExemption) {
	o.WebDataExemption = &v
}

func (o BusinessLineInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BusinessLineInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Capability) {
		toSerialize["capability"] = o.Capability
	}
	toSerialize["industryCode"] = o.IndustryCode
	toSerialize["legalEntityId"] = o.LegalEntityId
	if !common.IsNil(o.SalesChannels) {
		toSerialize["salesChannels"] = o.SalesChannels
	}
	toSerialize["service"] = o.Service
	if !common.IsNil(o.SourceOfFunds) {
		toSerialize["sourceOfFunds"] = o.SourceOfFunds
	}
	if !common.IsNil(o.WebData) {
		toSerialize["webData"] = o.WebData
	}
	if !common.IsNil(o.WebDataExemption) {
		toSerialize["webDataExemption"] = o.WebDataExemption
	}
	return toSerialize, nil
}

type NullableBusinessLineInfo struct {
	value *BusinessLineInfo
	isSet bool
}

func (v NullableBusinessLineInfo) Get() *BusinessLineInfo {
	return v.value
}

func (v *NullableBusinessLineInfo) Set(val *BusinessLineInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessLineInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessLineInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessLineInfo(val *BusinessLineInfo) *NullableBusinessLineInfo {
	return &NullableBusinessLineInfo{value: val, isSet: true}
}

func (v NullableBusinessLineInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusinessLineInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *BusinessLineInfo) isValidCapability() bool {
	var allowedEnumValues = []string{"receivePayments", "receiveFromPlatformPayments", "issueBankAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetCapability() == allowed {
			return true
		}
	}
	return false
}
func (o *BusinessLineInfo) isValidService() bool {
	var allowedEnumValues = []string{"paymentProcessing", "banking"}
	for _, allowed := range allowedEnumValues {
		if o.GetService() == allowed {
			return true
		}
	}
	return false
}
