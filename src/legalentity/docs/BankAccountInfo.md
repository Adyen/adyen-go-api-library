# BankAccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIdentification** | Pointer to [**BankAccountInfoAccountIdentification**](BankAccountInfoAccountIdentification.md) |  | [optional] 
**AccountType** | Pointer to **string** | The type of bank account. | [optional] 
**CountryCode** | Pointer to **string** | The two-character [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code where the bank account is registered. For example, **NL**. | [optional] 

## Methods

### NewBankAccountInfo

`func NewBankAccountInfo() *BankAccountInfo`

NewBankAccountInfo instantiates a new BankAccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankAccountInfoWithDefaults

`func NewBankAccountInfoWithDefaults() *BankAccountInfo`

NewBankAccountInfoWithDefaults instantiates a new BankAccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountIdentification

`func (o *BankAccountInfo) GetAccountIdentification() BankAccountInfoAccountIdentification`

GetAccountIdentification returns the AccountIdentification field if non-nil, zero value otherwise.

### GetAccountIdentificationOk

`func (o *BankAccountInfo) GetAccountIdentificationOk() (*BankAccountInfoAccountIdentification, bool)`

GetAccountIdentificationOk returns a tuple with the AccountIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdentification

`func (o *BankAccountInfo) SetAccountIdentification(v BankAccountInfoAccountIdentification)`

SetAccountIdentification sets AccountIdentification field to given value.

### HasAccountIdentification

`func (o *BankAccountInfo) HasAccountIdentification() bool`

HasAccountIdentification returns a boolean if a field has been set.

### GetAccountType

`func (o *BankAccountInfo) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *BankAccountInfo) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *BankAccountInfo) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *BankAccountInfo) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetCountryCode

`func (o *BankAccountInfo) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *BankAccountInfo) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *BankAccountInfo) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *BankAccountInfo) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


