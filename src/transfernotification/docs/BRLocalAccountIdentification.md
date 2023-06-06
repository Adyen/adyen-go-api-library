# BRLocalAccountIdentification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | The bank account number (without separators or whitespace). | 
**BankCode** | **string** | The 3-digit Brazilian bank code (with leading zeros). | 
**BranchNumber** | **string** | The bank account branch number (without separators or whitespace). | 
**Type** | **string** | **brLocal** | [default to "brLocal"]

## Methods

### NewBRLocalAccountIdentification

`func NewBRLocalAccountIdentification(accountNumber string, bankCode string, branchNumber string, type_ string, ) *BRLocalAccountIdentification`

NewBRLocalAccountIdentification instantiates a new BRLocalAccountIdentification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBRLocalAccountIdentificationWithDefaults

`func NewBRLocalAccountIdentificationWithDefaults() *BRLocalAccountIdentification`

NewBRLocalAccountIdentificationWithDefaults instantiates a new BRLocalAccountIdentification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *BRLocalAccountIdentification) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *BRLocalAccountIdentification) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *BRLocalAccountIdentification) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetBankCode

`func (o *BRLocalAccountIdentification) GetBankCode() string`

GetBankCode returns the BankCode field if non-nil, zero value otherwise.

### GetBankCodeOk

`func (o *BRLocalAccountIdentification) GetBankCodeOk() (*string, bool)`

GetBankCodeOk returns a tuple with the BankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCode

`func (o *BRLocalAccountIdentification) SetBankCode(v string)`

SetBankCode sets BankCode field to given value.


### GetBranchNumber

`func (o *BRLocalAccountIdentification) GetBranchNumber() string`

GetBranchNumber returns the BranchNumber field if non-nil, zero value otherwise.

### GetBranchNumberOk

`func (o *BRLocalAccountIdentification) GetBranchNumberOk() (*string, bool)`

GetBranchNumberOk returns a tuple with the BranchNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchNumber

`func (o *BRLocalAccountIdentification) SetBranchNumber(v string)`

SetBranchNumber sets BranchNumber field to given value.


### GetType

`func (o *BRLocalAccountIdentification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BRLocalAccountIdentification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BRLocalAccountIdentification) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


