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

// checks if the AdditionalDataTemporaryServices type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AdditionalDataTemporaryServices{}

// AdditionalDataTemporaryServices struct for AdditionalDataTemporaryServices
type AdditionalDataTemporaryServices struct {
	// The customer code, if supplied by a customer. * Encoding: ASCII * maxLength: 25
	EnhancedSchemeDataCustomerReference *string `json:"enhancedSchemeData.customerReference,omitempty"`
	// The name or ID of the person working in a temporary capacity. * maxLength: 40.   * Must not be all spaces.  *Must not be all zeros.
	EnhancedSchemeDataEmployeeName *string `json:"enhancedSchemeData.employeeName,omitempty"`
	// The job description of the person working in a temporary capacity. * maxLength: 40  * Must not be all spaces.  *Must not be all zeros.
	EnhancedSchemeDataJobDescription *string `json:"enhancedSchemeData.jobDescription,omitempty"`
	// The amount paid for regular hours worked, [minor units](https://docs.adyen.com/development-resources/currency-codes). * maxLength: 7 * Must not be empty * Can be all zeros
	EnhancedSchemeDataRegularHoursRate *string `json:"enhancedSchemeData.regularHoursRate,omitempty"`
	// The hours worked. * maxLength: 7 * Must not be empty * Can be all zeros
	EnhancedSchemeDataRegularHoursWorked *string `json:"enhancedSchemeData.regularHoursWorked,omitempty"`
	// The name of the person requesting temporary services. * maxLength: 40 * Must not be all zeros * Must not be all spaces
	EnhancedSchemeDataRequestName *string `json:"enhancedSchemeData.requestName,omitempty"`
	// The billing period start date. * Format: ddMMyy * maxLength: 6
	EnhancedSchemeDataTempStartDate *string `json:"enhancedSchemeData.tempStartDate,omitempty"`
	// The billing period end date. * Format: ddMMyy * maxLength: 6
	EnhancedSchemeDataTempWeekEnding *string `json:"enhancedSchemeData.tempWeekEnding,omitempty"`
	// The total tax amount, in [minor units](https://docs.adyen.com/development-resources/currency-codes). For example, 2000 means USD 20.00 * maxLength: 12
	EnhancedSchemeDataTotalTaxAmount *string `json:"enhancedSchemeData.totalTaxAmount,omitempty"`
}

// NewAdditionalDataTemporaryServices instantiates a new AdditionalDataTemporaryServices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalDataTemporaryServices() *AdditionalDataTemporaryServices {
	this := AdditionalDataTemporaryServices{}
	return &this
}

// NewAdditionalDataTemporaryServicesWithDefaults instantiates a new AdditionalDataTemporaryServices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalDataTemporaryServicesWithDefaults() *AdditionalDataTemporaryServices {
	this := AdditionalDataTemporaryServices{}
	return &this
}

// GetEnhancedSchemeDataCustomerReference returns the EnhancedSchemeDataCustomerReference field value if set, zero value otherwise.
func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataCustomerReference() string {
	if o == nil || common.IsNil(o.EnhancedSchemeDataCustomerReference) {
		var ret string
		return ret
	}
	return *o.EnhancedSchemeDataCustomerReference
}

// GetEnhancedSchemeDataCustomerReferenceOk returns a tuple with the EnhancedSchemeDataCustomerReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataCustomerReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnhancedSchemeDataCustomerReference) {
		return nil, false
	}
	return o.EnhancedSchemeDataCustomerReference, true
}

// HasEnhancedSchemeDataCustomerReference returns a boolean if a field has been set.
func (o *AdditionalDataTemporaryServices) HasEnhancedSchemeDataCustomerReference() bool {
	if o != nil && !common.IsNil(o.EnhancedSchemeDataCustomerReference) {
		return true
	}

	return false
}

// SetEnhancedSchemeDataCustomerReference gets a reference to the given string and assigns it to the EnhancedSchemeDataCustomerReference field.
func (o *AdditionalDataTemporaryServices) SetEnhancedSchemeDataCustomerReference(v string) {
	o.EnhancedSchemeDataCustomerReference = &v
}

// GetEnhancedSchemeDataEmployeeName returns the EnhancedSchemeDataEmployeeName field value if set, zero value otherwise.
func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataEmployeeName() string {
	if o == nil || common.IsNil(o.EnhancedSchemeDataEmployeeName) {
		var ret string
		return ret
	}
	return *o.EnhancedSchemeDataEmployeeName
}

// GetEnhancedSchemeDataEmployeeNameOk returns a tuple with the EnhancedSchemeDataEmployeeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataEmployeeNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnhancedSchemeDataEmployeeName) {
		return nil, false
	}
	return o.EnhancedSchemeDataEmployeeName, true
}

// HasEnhancedSchemeDataEmployeeName returns a boolean if a field has been set.
func (o *AdditionalDataTemporaryServices) HasEnhancedSchemeDataEmployeeName() bool {
	if o != nil && !common.IsNil(o.EnhancedSchemeDataEmployeeName) {
		return true
	}

	return false
}

// SetEnhancedSchemeDataEmployeeName gets a reference to the given string and assigns it to the EnhancedSchemeDataEmployeeName field.
func (o *AdditionalDataTemporaryServices) SetEnhancedSchemeDataEmployeeName(v string) {
	o.EnhancedSchemeDataEmployeeName = &v
}

// GetEnhancedSchemeDataJobDescription returns the EnhancedSchemeDataJobDescription field value if set, zero value otherwise.
func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataJobDescription() string {
	if o == nil || common.IsNil(o.EnhancedSchemeDataJobDescription) {
		var ret string
		return ret
	}
	return *o.EnhancedSchemeDataJobDescription
}

// GetEnhancedSchemeDataJobDescriptionOk returns a tuple with the EnhancedSchemeDataJobDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataJobDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnhancedSchemeDataJobDescription) {
		return nil, false
	}
	return o.EnhancedSchemeDataJobDescription, true
}

// HasEnhancedSchemeDataJobDescription returns a boolean if a field has been set.
func (o *AdditionalDataTemporaryServices) HasEnhancedSchemeDataJobDescription() bool {
	if o != nil && !common.IsNil(o.EnhancedSchemeDataJobDescription) {
		return true
	}

	return false
}

// SetEnhancedSchemeDataJobDescription gets a reference to the given string and assigns it to the EnhancedSchemeDataJobDescription field.
func (o *AdditionalDataTemporaryServices) SetEnhancedSchemeDataJobDescription(v string) {
	o.EnhancedSchemeDataJobDescription = &v
}

// GetEnhancedSchemeDataRegularHoursRate returns the EnhancedSchemeDataRegularHoursRate field value if set, zero value otherwise.
func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataRegularHoursRate() string {
	if o == nil || common.IsNil(o.EnhancedSchemeDataRegularHoursRate) {
		var ret string
		return ret
	}
	return *o.EnhancedSchemeDataRegularHoursRate
}

// GetEnhancedSchemeDataRegularHoursRateOk returns a tuple with the EnhancedSchemeDataRegularHoursRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataRegularHoursRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnhancedSchemeDataRegularHoursRate) {
		return nil, false
	}
	return o.EnhancedSchemeDataRegularHoursRate, true
}

// HasEnhancedSchemeDataRegularHoursRate returns a boolean if a field has been set.
func (o *AdditionalDataTemporaryServices) HasEnhancedSchemeDataRegularHoursRate() bool {
	if o != nil && !common.IsNil(o.EnhancedSchemeDataRegularHoursRate) {
		return true
	}

	return false
}

// SetEnhancedSchemeDataRegularHoursRate gets a reference to the given string and assigns it to the EnhancedSchemeDataRegularHoursRate field.
func (o *AdditionalDataTemporaryServices) SetEnhancedSchemeDataRegularHoursRate(v string) {
	o.EnhancedSchemeDataRegularHoursRate = &v
}

// GetEnhancedSchemeDataRegularHoursWorked returns the EnhancedSchemeDataRegularHoursWorked field value if set, zero value otherwise.
func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataRegularHoursWorked() string {
	if o == nil || common.IsNil(o.EnhancedSchemeDataRegularHoursWorked) {
		var ret string
		return ret
	}
	return *o.EnhancedSchemeDataRegularHoursWorked
}

// GetEnhancedSchemeDataRegularHoursWorkedOk returns a tuple with the EnhancedSchemeDataRegularHoursWorked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataRegularHoursWorkedOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnhancedSchemeDataRegularHoursWorked) {
		return nil, false
	}
	return o.EnhancedSchemeDataRegularHoursWorked, true
}

// HasEnhancedSchemeDataRegularHoursWorked returns a boolean if a field has been set.
func (o *AdditionalDataTemporaryServices) HasEnhancedSchemeDataRegularHoursWorked() bool {
	if o != nil && !common.IsNil(o.EnhancedSchemeDataRegularHoursWorked) {
		return true
	}

	return false
}

// SetEnhancedSchemeDataRegularHoursWorked gets a reference to the given string and assigns it to the EnhancedSchemeDataRegularHoursWorked field.
func (o *AdditionalDataTemporaryServices) SetEnhancedSchemeDataRegularHoursWorked(v string) {
	o.EnhancedSchemeDataRegularHoursWorked = &v
}

// GetEnhancedSchemeDataRequestName returns the EnhancedSchemeDataRequestName field value if set, zero value otherwise.
func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataRequestName() string {
	if o == nil || common.IsNil(o.EnhancedSchemeDataRequestName) {
		var ret string
		return ret
	}
	return *o.EnhancedSchemeDataRequestName
}

// GetEnhancedSchemeDataRequestNameOk returns a tuple with the EnhancedSchemeDataRequestName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataRequestNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnhancedSchemeDataRequestName) {
		return nil, false
	}
	return o.EnhancedSchemeDataRequestName, true
}

// HasEnhancedSchemeDataRequestName returns a boolean if a field has been set.
func (o *AdditionalDataTemporaryServices) HasEnhancedSchemeDataRequestName() bool {
	if o != nil && !common.IsNil(o.EnhancedSchemeDataRequestName) {
		return true
	}

	return false
}

// SetEnhancedSchemeDataRequestName gets a reference to the given string and assigns it to the EnhancedSchemeDataRequestName field.
func (o *AdditionalDataTemporaryServices) SetEnhancedSchemeDataRequestName(v string) {
	o.EnhancedSchemeDataRequestName = &v
}

// GetEnhancedSchemeDataTempStartDate returns the EnhancedSchemeDataTempStartDate field value if set, zero value otherwise.
func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataTempStartDate() string {
	if o == nil || common.IsNil(o.EnhancedSchemeDataTempStartDate) {
		var ret string
		return ret
	}
	return *o.EnhancedSchemeDataTempStartDate
}

// GetEnhancedSchemeDataTempStartDateOk returns a tuple with the EnhancedSchemeDataTempStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataTempStartDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnhancedSchemeDataTempStartDate) {
		return nil, false
	}
	return o.EnhancedSchemeDataTempStartDate, true
}

// HasEnhancedSchemeDataTempStartDate returns a boolean if a field has been set.
func (o *AdditionalDataTemporaryServices) HasEnhancedSchemeDataTempStartDate() bool {
	if o != nil && !common.IsNil(o.EnhancedSchemeDataTempStartDate) {
		return true
	}

	return false
}

// SetEnhancedSchemeDataTempStartDate gets a reference to the given string and assigns it to the EnhancedSchemeDataTempStartDate field.
func (o *AdditionalDataTemporaryServices) SetEnhancedSchemeDataTempStartDate(v string) {
	o.EnhancedSchemeDataTempStartDate = &v
}

// GetEnhancedSchemeDataTempWeekEnding returns the EnhancedSchemeDataTempWeekEnding field value if set, zero value otherwise.
func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataTempWeekEnding() string {
	if o == nil || common.IsNil(o.EnhancedSchemeDataTempWeekEnding) {
		var ret string
		return ret
	}
	return *o.EnhancedSchemeDataTempWeekEnding
}

// GetEnhancedSchemeDataTempWeekEndingOk returns a tuple with the EnhancedSchemeDataTempWeekEnding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataTempWeekEndingOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnhancedSchemeDataTempWeekEnding) {
		return nil, false
	}
	return o.EnhancedSchemeDataTempWeekEnding, true
}

// HasEnhancedSchemeDataTempWeekEnding returns a boolean if a field has been set.
func (o *AdditionalDataTemporaryServices) HasEnhancedSchemeDataTempWeekEnding() bool {
	if o != nil && !common.IsNil(o.EnhancedSchemeDataTempWeekEnding) {
		return true
	}

	return false
}

// SetEnhancedSchemeDataTempWeekEnding gets a reference to the given string and assigns it to the EnhancedSchemeDataTempWeekEnding field.
func (o *AdditionalDataTemporaryServices) SetEnhancedSchemeDataTempWeekEnding(v string) {
	o.EnhancedSchemeDataTempWeekEnding = &v
}

// GetEnhancedSchemeDataTotalTaxAmount returns the EnhancedSchemeDataTotalTaxAmount field value if set, zero value otherwise.
func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataTotalTaxAmount() string {
	if o == nil || common.IsNil(o.EnhancedSchemeDataTotalTaxAmount) {
		var ret string
		return ret
	}
	return *o.EnhancedSchemeDataTotalTaxAmount
}

// GetEnhancedSchemeDataTotalTaxAmountOk returns a tuple with the EnhancedSchemeDataTotalTaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataTotalTaxAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.EnhancedSchemeDataTotalTaxAmount) {
		return nil, false
	}
	return o.EnhancedSchemeDataTotalTaxAmount, true
}

// HasEnhancedSchemeDataTotalTaxAmount returns a boolean if a field has been set.
func (o *AdditionalDataTemporaryServices) HasEnhancedSchemeDataTotalTaxAmount() bool {
	if o != nil && !common.IsNil(o.EnhancedSchemeDataTotalTaxAmount) {
		return true
	}

	return false
}

// SetEnhancedSchemeDataTotalTaxAmount gets a reference to the given string and assigns it to the EnhancedSchemeDataTotalTaxAmount field.
func (o *AdditionalDataTemporaryServices) SetEnhancedSchemeDataTotalTaxAmount(v string) {
	o.EnhancedSchemeDataTotalTaxAmount = &v
}

func (o AdditionalDataTemporaryServices) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalDataTemporaryServices) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.EnhancedSchemeDataCustomerReference) {
		toSerialize["enhancedSchemeData.customerReference"] = o.EnhancedSchemeDataCustomerReference
	}
	if !common.IsNil(o.EnhancedSchemeDataEmployeeName) {
		toSerialize["enhancedSchemeData.employeeName"] = o.EnhancedSchemeDataEmployeeName
	}
	if !common.IsNil(o.EnhancedSchemeDataJobDescription) {
		toSerialize["enhancedSchemeData.jobDescription"] = o.EnhancedSchemeDataJobDescription
	}
	if !common.IsNil(o.EnhancedSchemeDataRegularHoursRate) {
		toSerialize["enhancedSchemeData.regularHoursRate"] = o.EnhancedSchemeDataRegularHoursRate
	}
	if !common.IsNil(o.EnhancedSchemeDataRegularHoursWorked) {
		toSerialize["enhancedSchemeData.regularHoursWorked"] = o.EnhancedSchemeDataRegularHoursWorked
	}
	if !common.IsNil(o.EnhancedSchemeDataRequestName) {
		toSerialize["enhancedSchemeData.requestName"] = o.EnhancedSchemeDataRequestName
	}
	if !common.IsNil(o.EnhancedSchemeDataTempStartDate) {
		toSerialize["enhancedSchemeData.tempStartDate"] = o.EnhancedSchemeDataTempStartDate
	}
	if !common.IsNil(o.EnhancedSchemeDataTempWeekEnding) {
		toSerialize["enhancedSchemeData.tempWeekEnding"] = o.EnhancedSchemeDataTempWeekEnding
	}
	if !common.IsNil(o.EnhancedSchemeDataTotalTaxAmount) {
		toSerialize["enhancedSchemeData.totalTaxAmount"] = o.EnhancedSchemeDataTotalTaxAmount
	}
	return toSerialize, nil
}

type NullableAdditionalDataTemporaryServices struct {
	value *AdditionalDataTemporaryServices
	isSet bool
}

func (v NullableAdditionalDataTemporaryServices) Get() *AdditionalDataTemporaryServices {
	return v.value
}

func (v *NullableAdditionalDataTemporaryServices) Set(val *AdditionalDataTemporaryServices) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalDataTemporaryServices) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalDataTemporaryServices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalDataTemporaryServices(val *AdditionalDataTemporaryServices) *NullableAdditionalDataTemporaryServices {
	return &NullableAdditionalDataTemporaryServices{value: val, isSet: true}
}

func (v NullableAdditionalDataTemporaryServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalDataTemporaryServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
