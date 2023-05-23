# AdditionalDataTemporaryServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnhancedSchemeDataCustomerReference** | Pointer to **string** | The customer code, if supplied by a customer. * Encoding: ASCII * maxLength: 25 | [optional] 
**EnhancedSchemeDataEmployeeName** | Pointer to **string** | The name or ID of the person working in a temporary capacity. * maxLength: 40 * Must not be all zeros * Must not be all spaces | [optional] 
**EnhancedSchemeDataJobDescription** | Pointer to **string** | The job description of the person working in a temporary capacity. * maxLength: 40 * Must not be all zeros * Must not be all spaces | [optional] 
**EnhancedSchemeDataRegularHoursRate** | Pointer to **string** | The amount paid for regular hours worked, [minor units](https://docs.adyen.com/development-resources/currency-codes). * maxLength: 7 * Must not be empty * Can be all zeros | [optional] 
**EnhancedSchemeDataRegularHoursWorked** | Pointer to **string** | The hours worked. * maxLength: 7 * Must not be empty * Can be all zeros | [optional] 
**EnhancedSchemeDataRequestName** | Pointer to **string** | The name of the person requesting temporary services. * maxLength: 40 * Must not be all zeros * Must not be all spaces | [optional] 
**EnhancedSchemeDataTempStartDate** | Pointer to **string** | The billing period start date. * Format: ddMMyy * maxLength: 6 | [optional] 
**EnhancedSchemeDataTempWeekEnding** | Pointer to **string** | The billing period end date. * Format: ddMMyy * maxLength: 6 | [optional] 
**EnhancedSchemeDataTotalTaxAmount** | Pointer to **string** | The total tax amount, in [minor units](https://docs.adyen.com/development-resources/currency-codes). For example, 2000 means USD 20.00 * maxLength: 12 | [optional] 

## Methods

### NewAdditionalDataTemporaryServices

`func NewAdditionalDataTemporaryServices() *AdditionalDataTemporaryServices`

NewAdditionalDataTemporaryServices instantiates a new AdditionalDataTemporaryServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalDataTemporaryServicesWithDefaults

`func NewAdditionalDataTemporaryServicesWithDefaults() *AdditionalDataTemporaryServices`

NewAdditionalDataTemporaryServicesWithDefaults instantiates a new AdditionalDataTemporaryServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnhancedSchemeDataCustomerReference

`func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataCustomerReference() string`

GetEnhancedSchemeDataCustomerReference returns the EnhancedSchemeDataCustomerReference field if non-nil, zero value otherwise.

### GetEnhancedSchemeDataCustomerReferenceOk

`func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataCustomerReferenceOk() (*string, bool)`

GetEnhancedSchemeDataCustomerReferenceOk returns a tuple with the EnhancedSchemeDataCustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSchemeDataCustomerReference

`func (o *AdditionalDataTemporaryServices) SetEnhancedSchemeDataCustomerReference(v string)`

SetEnhancedSchemeDataCustomerReference sets EnhancedSchemeDataCustomerReference field to given value.

### HasEnhancedSchemeDataCustomerReference

`func (o *AdditionalDataTemporaryServices) HasEnhancedSchemeDataCustomerReference() bool`

HasEnhancedSchemeDataCustomerReference returns a boolean if a field has been set.

### GetEnhancedSchemeDataEmployeeName

`func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataEmployeeName() string`

GetEnhancedSchemeDataEmployeeName returns the EnhancedSchemeDataEmployeeName field if non-nil, zero value otherwise.

### GetEnhancedSchemeDataEmployeeNameOk

`func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataEmployeeNameOk() (*string, bool)`

GetEnhancedSchemeDataEmployeeNameOk returns a tuple with the EnhancedSchemeDataEmployeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSchemeDataEmployeeName

`func (o *AdditionalDataTemporaryServices) SetEnhancedSchemeDataEmployeeName(v string)`

SetEnhancedSchemeDataEmployeeName sets EnhancedSchemeDataEmployeeName field to given value.

### HasEnhancedSchemeDataEmployeeName

`func (o *AdditionalDataTemporaryServices) HasEnhancedSchemeDataEmployeeName() bool`

HasEnhancedSchemeDataEmployeeName returns a boolean if a field has been set.

### GetEnhancedSchemeDataJobDescription

`func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataJobDescription() string`

GetEnhancedSchemeDataJobDescription returns the EnhancedSchemeDataJobDescription field if non-nil, zero value otherwise.

### GetEnhancedSchemeDataJobDescriptionOk

`func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataJobDescriptionOk() (*string, bool)`

GetEnhancedSchemeDataJobDescriptionOk returns a tuple with the EnhancedSchemeDataJobDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSchemeDataJobDescription

`func (o *AdditionalDataTemporaryServices) SetEnhancedSchemeDataJobDescription(v string)`

SetEnhancedSchemeDataJobDescription sets EnhancedSchemeDataJobDescription field to given value.

### HasEnhancedSchemeDataJobDescription

`func (o *AdditionalDataTemporaryServices) HasEnhancedSchemeDataJobDescription() bool`

HasEnhancedSchemeDataJobDescription returns a boolean if a field has been set.

### GetEnhancedSchemeDataRegularHoursRate

`func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataRegularHoursRate() string`

GetEnhancedSchemeDataRegularHoursRate returns the EnhancedSchemeDataRegularHoursRate field if non-nil, zero value otherwise.

### GetEnhancedSchemeDataRegularHoursRateOk

`func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataRegularHoursRateOk() (*string, bool)`

GetEnhancedSchemeDataRegularHoursRateOk returns a tuple with the EnhancedSchemeDataRegularHoursRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSchemeDataRegularHoursRate

`func (o *AdditionalDataTemporaryServices) SetEnhancedSchemeDataRegularHoursRate(v string)`

SetEnhancedSchemeDataRegularHoursRate sets EnhancedSchemeDataRegularHoursRate field to given value.

### HasEnhancedSchemeDataRegularHoursRate

`func (o *AdditionalDataTemporaryServices) HasEnhancedSchemeDataRegularHoursRate() bool`

HasEnhancedSchemeDataRegularHoursRate returns a boolean if a field has been set.

### GetEnhancedSchemeDataRegularHoursWorked

`func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataRegularHoursWorked() string`

GetEnhancedSchemeDataRegularHoursWorked returns the EnhancedSchemeDataRegularHoursWorked field if non-nil, zero value otherwise.

### GetEnhancedSchemeDataRegularHoursWorkedOk

`func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataRegularHoursWorkedOk() (*string, bool)`

GetEnhancedSchemeDataRegularHoursWorkedOk returns a tuple with the EnhancedSchemeDataRegularHoursWorked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSchemeDataRegularHoursWorked

`func (o *AdditionalDataTemporaryServices) SetEnhancedSchemeDataRegularHoursWorked(v string)`

SetEnhancedSchemeDataRegularHoursWorked sets EnhancedSchemeDataRegularHoursWorked field to given value.

### HasEnhancedSchemeDataRegularHoursWorked

`func (o *AdditionalDataTemporaryServices) HasEnhancedSchemeDataRegularHoursWorked() bool`

HasEnhancedSchemeDataRegularHoursWorked returns a boolean if a field has been set.

### GetEnhancedSchemeDataRequestName

`func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataRequestName() string`

GetEnhancedSchemeDataRequestName returns the EnhancedSchemeDataRequestName field if non-nil, zero value otherwise.

### GetEnhancedSchemeDataRequestNameOk

`func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataRequestNameOk() (*string, bool)`

GetEnhancedSchemeDataRequestNameOk returns a tuple with the EnhancedSchemeDataRequestName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSchemeDataRequestName

`func (o *AdditionalDataTemporaryServices) SetEnhancedSchemeDataRequestName(v string)`

SetEnhancedSchemeDataRequestName sets EnhancedSchemeDataRequestName field to given value.

### HasEnhancedSchemeDataRequestName

`func (o *AdditionalDataTemporaryServices) HasEnhancedSchemeDataRequestName() bool`

HasEnhancedSchemeDataRequestName returns a boolean if a field has been set.

### GetEnhancedSchemeDataTempStartDate

`func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataTempStartDate() string`

GetEnhancedSchemeDataTempStartDate returns the EnhancedSchemeDataTempStartDate field if non-nil, zero value otherwise.

### GetEnhancedSchemeDataTempStartDateOk

`func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataTempStartDateOk() (*string, bool)`

GetEnhancedSchemeDataTempStartDateOk returns a tuple with the EnhancedSchemeDataTempStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSchemeDataTempStartDate

`func (o *AdditionalDataTemporaryServices) SetEnhancedSchemeDataTempStartDate(v string)`

SetEnhancedSchemeDataTempStartDate sets EnhancedSchemeDataTempStartDate field to given value.

### HasEnhancedSchemeDataTempStartDate

`func (o *AdditionalDataTemporaryServices) HasEnhancedSchemeDataTempStartDate() bool`

HasEnhancedSchemeDataTempStartDate returns a boolean if a field has been set.

### GetEnhancedSchemeDataTempWeekEnding

`func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataTempWeekEnding() string`

GetEnhancedSchemeDataTempWeekEnding returns the EnhancedSchemeDataTempWeekEnding field if non-nil, zero value otherwise.

### GetEnhancedSchemeDataTempWeekEndingOk

`func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataTempWeekEndingOk() (*string, bool)`

GetEnhancedSchemeDataTempWeekEndingOk returns a tuple with the EnhancedSchemeDataTempWeekEnding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSchemeDataTempWeekEnding

`func (o *AdditionalDataTemporaryServices) SetEnhancedSchemeDataTempWeekEnding(v string)`

SetEnhancedSchemeDataTempWeekEnding sets EnhancedSchemeDataTempWeekEnding field to given value.

### HasEnhancedSchemeDataTempWeekEnding

`func (o *AdditionalDataTemporaryServices) HasEnhancedSchemeDataTempWeekEnding() bool`

HasEnhancedSchemeDataTempWeekEnding returns a boolean if a field has been set.

### GetEnhancedSchemeDataTotalTaxAmount

`func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataTotalTaxAmount() string`

GetEnhancedSchemeDataTotalTaxAmount returns the EnhancedSchemeDataTotalTaxAmount field if non-nil, zero value otherwise.

### GetEnhancedSchemeDataTotalTaxAmountOk

`func (o *AdditionalDataTemporaryServices) GetEnhancedSchemeDataTotalTaxAmountOk() (*string, bool)`

GetEnhancedSchemeDataTotalTaxAmountOk returns a tuple with the EnhancedSchemeDataTotalTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSchemeDataTotalTaxAmount

`func (o *AdditionalDataTemporaryServices) SetEnhancedSchemeDataTotalTaxAmount(v string)`

SetEnhancedSchemeDataTotalTaxAmount sets EnhancedSchemeDataTotalTaxAmount field to given value.

### HasEnhancedSchemeDataTotalTaxAmount

`func (o *AdditionalDataTemporaryServices) HasEnhancedSchemeDataTotalTaxAmount() bool`

HasEnhancedSchemeDataTotalTaxAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


