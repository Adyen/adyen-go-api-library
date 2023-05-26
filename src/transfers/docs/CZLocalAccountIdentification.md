# CZLocalAccountIdentification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | The 2- to 16-digit bank account number (Číslo účtu) in the following format:  - The optional prefix (předčíslí).  - The required second part (základní část) which must be at least two non-zero digits.  Examples:  - **19-123457** (with prefix)  - **123457** (without prefix)  - **000019-0000123457** (with prefix, normalized)  - **000000-0000123457** (without prefix, normalized) | 
**BankCode** | **string** | The 4-digit bank code (Kód banky), without separators or whitespace. | 
**Type** | **string** | **czLocal** | [default to "czLocal"]

## Methods

### NewCZLocalAccountIdentification

`func NewCZLocalAccountIdentification(accountNumber string, bankCode string, type_ string, ) *CZLocalAccountIdentification`

NewCZLocalAccountIdentification instantiates a new CZLocalAccountIdentification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCZLocalAccountIdentificationWithDefaults

`func NewCZLocalAccountIdentificationWithDefaults() *CZLocalAccountIdentification`

NewCZLocalAccountIdentificationWithDefaults instantiates a new CZLocalAccountIdentification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *CZLocalAccountIdentification) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *CZLocalAccountIdentification) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *CZLocalAccountIdentification) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetBankCode

`func (o *CZLocalAccountIdentification) GetBankCode() string`

GetBankCode returns the BankCode field if non-nil, zero value otherwise.

### GetBankCodeOk

`func (o *CZLocalAccountIdentification) GetBankCodeOk() (*string, bool)`

GetBankCodeOk returns a tuple with the BankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCode

`func (o *CZLocalAccountIdentification) SetBankCode(v string)`

SetBankCode sets BankCode field to given value.


### GetType

`func (o *CZLocalAccountIdentification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CZLocalAccountIdentification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CZLocalAccountIdentification) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


