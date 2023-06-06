# BankAccountV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountHolder** | [**PartyIdentification**](PartyIdentification.md) |  | 
**AccountIdentification** | [**BankAccountV3AccountIdentification**](BankAccountV3AccountIdentification.md) |  | 

## Methods

### NewBankAccountV3

`func NewBankAccountV3(accountHolder PartyIdentification, accountIdentification BankAccountV3AccountIdentification, ) *BankAccountV3`

NewBankAccountV3 instantiates a new BankAccountV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankAccountV3WithDefaults

`func NewBankAccountV3WithDefaults() *BankAccountV3`

NewBankAccountV3WithDefaults instantiates a new BankAccountV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountHolder

`func (o *BankAccountV3) GetAccountHolder() PartyIdentification`

GetAccountHolder returns the AccountHolder field if non-nil, zero value otherwise.

### GetAccountHolderOk

`func (o *BankAccountV3) GetAccountHolderOk() (*PartyIdentification, bool)`

GetAccountHolderOk returns a tuple with the AccountHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolder

`func (o *BankAccountV3) SetAccountHolder(v PartyIdentification)`

SetAccountHolder sets AccountHolder field to given value.


### GetAccountIdentification

`func (o *BankAccountV3) GetAccountIdentification() BankAccountV3AccountIdentification`

GetAccountIdentification returns the AccountIdentification field if non-nil, zero value otherwise.

### GetAccountIdentificationOk

`func (o *BankAccountV3) GetAccountIdentificationOk() (*BankAccountV3AccountIdentification, bool)`

GetAccountIdentificationOk returns a tuple with the AccountIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdentification

`func (o *BankAccountV3) SetAccountIdentification(v BankAccountV3AccountIdentification)`

SetAccountIdentification sets AccountIdentification field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


