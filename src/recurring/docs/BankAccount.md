# BankAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankAccountNumber** | Pointer to **string** | The bank account number (without separators). | [optional] 
**BankCity** | Pointer to **string** | The bank city. | [optional] 
**BankLocationId** | Pointer to **string** | The location id of the bank. The field value is &#x60;nil&#x60; in most cases. | [optional] 
**BankName** | Pointer to **string** | The name of the bank. | [optional] 
**Bic** | Pointer to **string** | The [Business Identifier Code](https://en.wikipedia.org/wiki/ISO_9362) (BIC) is the SWIFT address assigned to a bank. The field value is &#x60;nil&#x60; in most cases. | [optional] 
**CountryCode** | Pointer to **string** | Country code where the bank is located.  A valid value is an ISO two-character country code (e.g. &#39;NL&#39;). | [optional] 
**Iban** | Pointer to **string** | The [International Bank Account Number](https://en.wikipedia.org/wiki/International_Bank_Account_Number) (IBAN). | [optional] 
**OwnerName** | Pointer to **string** | The name of the bank account holder. If you submit a name with non-Latin characters, we automatically replace some of them with corresponding Latin characters to meet the FATF recommendations. For example: * χ12 is converted to ch12. * üA is converted to euA. * Peter Møller is converted to Peter Mller, because banks don&#39;t accept &#39;ø&#39;. After replacement, the ownerName must have at least three alphanumeric characters (A-Z, a-z, 0-9), and at least one of them must be a valid Latin character (A-Z, a-z). For example: * John17 - allowed. * J17 - allowed. * 171 - not allowed. * John-7 - allowed. &gt; If provided details don&#39;t match the required format, the response returns the error message: 203 &#39;Invalid bank account holder name&#39;. | [optional] 
**TaxId** | Pointer to **string** | The bank account holder&#39;s tax ID. | [optional] 

## Methods

### NewBankAccount

`func NewBankAccount() *BankAccount`

NewBankAccount instantiates a new BankAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankAccountWithDefaults

`func NewBankAccountWithDefaults() *BankAccount`

NewBankAccountWithDefaults instantiates a new BankAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankAccountNumber

`func (o *BankAccount) GetBankAccountNumber() string`

GetBankAccountNumber returns the BankAccountNumber field if non-nil, zero value otherwise.

### GetBankAccountNumberOk

`func (o *BankAccount) GetBankAccountNumberOk() (*string, bool)`

GetBankAccountNumberOk returns a tuple with the BankAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountNumber

`func (o *BankAccount) SetBankAccountNumber(v string)`

SetBankAccountNumber sets BankAccountNumber field to given value.

### HasBankAccountNumber

`func (o *BankAccount) HasBankAccountNumber() bool`

HasBankAccountNumber returns a boolean if a field has been set.

### GetBankCity

`func (o *BankAccount) GetBankCity() string`

GetBankCity returns the BankCity field if non-nil, zero value otherwise.

### GetBankCityOk

`func (o *BankAccount) GetBankCityOk() (*string, bool)`

GetBankCityOk returns a tuple with the BankCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCity

`func (o *BankAccount) SetBankCity(v string)`

SetBankCity sets BankCity field to given value.

### HasBankCity

`func (o *BankAccount) HasBankCity() bool`

HasBankCity returns a boolean if a field has been set.

### GetBankLocationId

`func (o *BankAccount) GetBankLocationId() string`

GetBankLocationId returns the BankLocationId field if non-nil, zero value otherwise.

### GetBankLocationIdOk

`func (o *BankAccount) GetBankLocationIdOk() (*string, bool)`

GetBankLocationIdOk returns a tuple with the BankLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankLocationId

`func (o *BankAccount) SetBankLocationId(v string)`

SetBankLocationId sets BankLocationId field to given value.

### HasBankLocationId

`func (o *BankAccount) HasBankLocationId() bool`

HasBankLocationId returns a boolean if a field has been set.

### GetBankName

`func (o *BankAccount) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *BankAccount) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *BankAccount) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *BankAccount) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### GetBic

`func (o *BankAccount) GetBic() string`

GetBic returns the Bic field if non-nil, zero value otherwise.

### GetBicOk

`func (o *BankAccount) GetBicOk() (*string, bool)`

GetBicOk returns a tuple with the Bic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBic

`func (o *BankAccount) SetBic(v string)`

SetBic sets Bic field to given value.

### HasBic

`func (o *BankAccount) HasBic() bool`

HasBic returns a boolean if a field has been set.

### GetCountryCode

`func (o *BankAccount) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *BankAccount) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *BankAccount) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *BankAccount) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetIban

`func (o *BankAccount) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *BankAccount) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *BankAccount) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *BankAccount) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetOwnerName

`func (o *BankAccount) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *BankAccount) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *BankAccount) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *BankAccount) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### GetTaxId

`func (o *BankAccount) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *BankAccount) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *BankAccount) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *BankAccount) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


