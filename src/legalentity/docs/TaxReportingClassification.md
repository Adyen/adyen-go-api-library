# TaxReportingClassification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessType** | Pointer to **string** | The organization&#39;s business type.  Possible values: **other**, **listedPublicCompany**, **subsidiaryOfListedPublicCompany**, **governmentalOrganization**, **internationalOrganization**, **financialInstitution**. | [optional] 
**FinancialInstitutionNumber** | Pointer to **string** | The Global Intermediary Identification Number (GIIN) required for FATCA. | [optional] 
**MainSourceOfIncome** | Pointer to **string** | The organization&#39;s main source of income.  Possible values: **businessOperation**, **realEstateSales**, **investmentInterestOrRoyalty**, **propertyRental**, **other**. | [optional] 
**Type** | Pointer to **string** | The tax reporting classification type.  Possible values: **nonFinancialNonReportable**, **financialNonReportable**, **nonFinancialActive**, **nonFinancialPassive**. | [optional] 

## Methods

### NewTaxReportingClassification

`func NewTaxReportingClassification() *TaxReportingClassification`

NewTaxReportingClassification instantiates a new TaxReportingClassification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxReportingClassificationWithDefaults

`func NewTaxReportingClassificationWithDefaults() *TaxReportingClassification`

NewTaxReportingClassificationWithDefaults instantiates a new TaxReportingClassification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessType

`func (o *TaxReportingClassification) GetBusinessType() string`

GetBusinessType returns the BusinessType field if non-nil, zero value otherwise.

### GetBusinessTypeOk

`func (o *TaxReportingClassification) GetBusinessTypeOk() (*string, bool)`

GetBusinessTypeOk returns a tuple with the BusinessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessType

`func (o *TaxReportingClassification) SetBusinessType(v string)`

SetBusinessType sets BusinessType field to given value.

### HasBusinessType

`func (o *TaxReportingClassification) HasBusinessType() bool`

HasBusinessType returns a boolean if a field has been set.

### GetFinancialInstitutionNumber

`func (o *TaxReportingClassification) GetFinancialInstitutionNumber() string`

GetFinancialInstitutionNumber returns the FinancialInstitutionNumber field if non-nil, zero value otherwise.

### GetFinancialInstitutionNumberOk

`func (o *TaxReportingClassification) GetFinancialInstitutionNumberOk() (*string, bool)`

GetFinancialInstitutionNumberOk returns a tuple with the FinancialInstitutionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialInstitutionNumber

`func (o *TaxReportingClassification) SetFinancialInstitutionNumber(v string)`

SetFinancialInstitutionNumber sets FinancialInstitutionNumber field to given value.

### HasFinancialInstitutionNumber

`func (o *TaxReportingClassification) HasFinancialInstitutionNumber() bool`

HasFinancialInstitutionNumber returns a boolean if a field has been set.

### GetMainSourceOfIncome

`func (o *TaxReportingClassification) GetMainSourceOfIncome() string`

GetMainSourceOfIncome returns the MainSourceOfIncome field if non-nil, zero value otherwise.

### GetMainSourceOfIncomeOk

`func (o *TaxReportingClassification) GetMainSourceOfIncomeOk() (*string, bool)`

GetMainSourceOfIncomeOk returns a tuple with the MainSourceOfIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainSourceOfIncome

`func (o *TaxReportingClassification) SetMainSourceOfIncome(v string)`

SetMainSourceOfIncome sets MainSourceOfIncome field to given value.

### HasMainSourceOfIncome

`func (o *TaxReportingClassification) HasMainSourceOfIncome() bool`

HasMainSourceOfIncome returns a boolean if a field has been set.

### GetType

`func (o *TaxReportingClassification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaxReportingClassification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaxReportingClassification) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TaxReportingClassification) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


