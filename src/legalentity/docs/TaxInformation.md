# TaxInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | The two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code. | [optional] 
**Number** | Pointer to **string** | The tax ID number (TIN) of the organization or individual. | [optional] 
**Type** | Pointer to **string** | The TIN type depending on the country where it was issued. Provide only for countries that have multiple tax IDs, such as Sweden, the UK, or the US. For example, provide **SSN**, **EIN**, or **ITIN** for the US. | [optional] 

## Methods

### NewTaxInformation

`func NewTaxInformation() *TaxInformation`

NewTaxInformation instantiates a new TaxInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxInformationWithDefaults

`func NewTaxInformationWithDefaults() *TaxInformation`

NewTaxInformationWithDefaults instantiates a new TaxInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *TaxInformation) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *TaxInformation) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *TaxInformation) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *TaxInformation) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetNumber

`func (o *TaxInformation) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *TaxInformation) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *TaxInformation) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *TaxInformation) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetType

`func (o *TaxInformation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaxInformation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaxInformation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TaxInformation) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


