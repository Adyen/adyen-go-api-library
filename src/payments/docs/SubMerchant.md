# SubMerchant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** | The city of the sub-merchant&#39;s address. * Format: Alphanumeric * Maximum length: 13 characters | [optional] 
**Country** | Pointer to **string** | The three-letter country code of the sub-merchant&#39;s address. For example, **BRA** for Brazil.  * Format: [ISO 3166-1 alpha-3](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3) * Fixed length: 3 characters | [optional] 
**Mcc** | Pointer to **string** | The sub-merchant&#39;s 4-digit Merchant Category Code (MCC).  * Format: Numeric * Fixed length: 4 digits | [optional] 
**Name** | Pointer to **string** | The name of the sub-merchant. Based on scheme specifications, this value will overwrite the shopper statement  that will appear in the card statement. * Format: Alphanumeric * Maximum length: 22 characters | [optional] 
**TaxId** | Pointer to **string** | The tax ID of the sub-merchant. * Format: Numeric * Fixed length: 11 digits for the CPF or 14 digits for the CNPJ | [optional] 

## Methods

### NewSubMerchant

`func NewSubMerchant() *SubMerchant`

NewSubMerchant instantiates a new SubMerchant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubMerchantWithDefaults

`func NewSubMerchantWithDefaults() *SubMerchant`

NewSubMerchantWithDefaults instantiates a new SubMerchant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *SubMerchant) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *SubMerchant) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *SubMerchant) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *SubMerchant) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *SubMerchant) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SubMerchant) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SubMerchant) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SubMerchant) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetMcc

`func (o *SubMerchant) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *SubMerchant) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *SubMerchant) SetMcc(v string)`

SetMcc sets Mcc field to given value.

### HasMcc

`func (o *SubMerchant) HasMcc() bool`

HasMcc returns a boolean if a field has been set.

### GetName

`func (o *SubMerchant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubMerchant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubMerchant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubMerchant) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTaxId

`func (o *SubMerchant) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *SubMerchant) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *SubMerchant) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *SubMerchant) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


