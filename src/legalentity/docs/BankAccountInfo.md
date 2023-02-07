# BankAccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **string** | The bank account number (without separators).   When this is provided, the &#x60;branchCode&#x60; is also required. | [optional] 
**AccountType** | Pointer to **string** | The type of bank account. | [optional] 
**BankBicSwift** | Pointer to **string** | The bank&#39;s BIC or SWIFT code. | [optional] 
**BankCity** | Pointer to **string** | The city where the bank is located. | [optional] 
**BankCode** | Pointer to **string** | The bank code of the banking institution with which the bank account is registered. | [optional] 
**BankName** | Pointer to **string** | The name of the banking institution where the bank account is held. | [optional] 
**BranchCode** | Pointer to **string** | The branch code of the branch under which the bank account is registered.  Required when you provide an &#x60;accountNumber&#x60;.   In the following countries, this value corresponds to:   * United States: routing number * United Kingdom: sort code * Germany: Bankleitzahl | [optional] 
**CheckCode** | Pointer to **string** | The check code of the bank account. | [optional] 
**CountryCode** | **string** | The two-character [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code where the bank account is registered. For example, **NL**. | 
**CurrencyCode** | **string** | The account&#39;s three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes). For example, **EUR**. | 
**Iban** | Pointer to **string** | The international bank account number as defined in the [ISO-13616](https://www.iso.org/standard/81090.html) standard. | [optional] 

## Methods

### NewBankAccountInfo

`func NewBankAccountInfo(countryCode string, currencyCode string, ) *BankAccountInfo`

NewBankAccountInfo instantiates a new BankAccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankAccountInfoWithDefaults

`func NewBankAccountInfoWithDefaults() *BankAccountInfo`

NewBankAccountInfoWithDefaults instantiates a new BankAccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *BankAccountInfo) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *BankAccountInfo) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *BankAccountInfo) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *BankAccountInfo) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

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

### GetBankBicSwift

`func (o *BankAccountInfo) GetBankBicSwift() string`

GetBankBicSwift returns the BankBicSwift field if non-nil, zero value otherwise.

### GetBankBicSwiftOk

`func (o *BankAccountInfo) GetBankBicSwiftOk() (*string, bool)`

GetBankBicSwiftOk returns a tuple with the BankBicSwift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankBicSwift

`func (o *BankAccountInfo) SetBankBicSwift(v string)`

SetBankBicSwift sets BankBicSwift field to given value.

### HasBankBicSwift

`func (o *BankAccountInfo) HasBankBicSwift() bool`

HasBankBicSwift returns a boolean if a field has been set.

### GetBankCity

`func (o *BankAccountInfo) GetBankCity() string`

GetBankCity returns the BankCity field if non-nil, zero value otherwise.

### GetBankCityOk

`func (o *BankAccountInfo) GetBankCityOk() (*string, bool)`

GetBankCityOk returns a tuple with the BankCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCity

`func (o *BankAccountInfo) SetBankCity(v string)`

SetBankCity sets BankCity field to given value.

### HasBankCity

`func (o *BankAccountInfo) HasBankCity() bool`

HasBankCity returns a boolean if a field has been set.

### GetBankCode

`func (o *BankAccountInfo) GetBankCode() string`

GetBankCode returns the BankCode field if non-nil, zero value otherwise.

### GetBankCodeOk

`func (o *BankAccountInfo) GetBankCodeOk() (*string, bool)`

GetBankCodeOk returns a tuple with the BankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCode

`func (o *BankAccountInfo) SetBankCode(v string)`

SetBankCode sets BankCode field to given value.

### HasBankCode

`func (o *BankAccountInfo) HasBankCode() bool`

HasBankCode returns a boolean if a field has been set.

### GetBankName

`func (o *BankAccountInfo) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *BankAccountInfo) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *BankAccountInfo) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *BankAccountInfo) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### GetBranchCode

`func (o *BankAccountInfo) GetBranchCode() string`

GetBranchCode returns the BranchCode field if non-nil, zero value otherwise.

### GetBranchCodeOk

`func (o *BankAccountInfo) GetBranchCodeOk() (*string, bool)`

GetBranchCodeOk returns a tuple with the BranchCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchCode

`func (o *BankAccountInfo) SetBranchCode(v string)`

SetBranchCode sets BranchCode field to given value.

### HasBranchCode

`func (o *BankAccountInfo) HasBranchCode() bool`

HasBranchCode returns a boolean if a field has been set.

### GetCheckCode

`func (o *BankAccountInfo) GetCheckCode() string`

GetCheckCode returns the CheckCode field if non-nil, zero value otherwise.

### GetCheckCodeOk

`func (o *BankAccountInfo) GetCheckCodeOk() (*string, bool)`

GetCheckCodeOk returns a tuple with the CheckCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCode

`func (o *BankAccountInfo) SetCheckCode(v string)`

SetCheckCode sets CheckCode field to given value.

### HasCheckCode

`func (o *BankAccountInfo) HasCheckCode() bool`

HasCheckCode returns a boolean if a field has been set.

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


### GetCurrencyCode

`func (o *BankAccountInfo) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *BankAccountInfo) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *BankAccountInfo) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetIban

`func (o *BankAccountInfo) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *BankAccountInfo) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *BankAccountInfo) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *BankAccountInfo) HasIban() bool`

HasIban returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


