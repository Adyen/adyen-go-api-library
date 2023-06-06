# DKLocalAccountIdentification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | The 4-10 digits bank account number (Kontonummer) (without separators or whitespace). | 
**BankCode** | **string** | The 4-digit bank code (Registreringsnummer) (without separators or whitespace). | 
**Type** | **string** | **dkLocal** | [default to "dkLocal"]

## Methods

### NewDKLocalAccountIdentification

`func NewDKLocalAccountIdentification(accountNumber string, bankCode string, type_ string, ) *DKLocalAccountIdentification`

NewDKLocalAccountIdentification instantiates a new DKLocalAccountIdentification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDKLocalAccountIdentificationWithDefaults

`func NewDKLocalAccountIdentificationWithDefaults() *DKLocalAccountIdentification`

NewDKLocalAccountIdentificationWithDefaults instantiates a new DKLocalAccountIdentification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *DKLocalAccountIdentification) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *DKLocalAccountIdentification) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *DKLocalAccountIdentification) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetBankCode

`func (o *DKLocalAccountIdentification) GetBankCode() string`

GetBankCode returns the BankCode field if non-nil, zero value otherwise.

### GetBankCodeOk

`func (o *DKLocalAccountIdentification) GetBankCodeOk() (*string, bool)`

GetBankCodeOk returns a tuple with the BankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCode

`func (o *DKLocalAccountIdentification) SetBankCode(v string)`

SetBankCode sets BankCode field to given value.


### GetType

`func (o *DKLocalAccountIdentification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DKLocalAccountIdentification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DKLocalAccountIdentification) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


