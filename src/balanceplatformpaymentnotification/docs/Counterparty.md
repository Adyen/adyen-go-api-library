# Counterparty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceAccountId** | Pointer to **string** | Unique identifier of the [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts__resParam_id). | [optional] 
**BankAccount** | Pointer to [**BankAccountInfo**](BankAccountInfo.md) |  | [optional] 
**Merchant** | Pointer to [**MerchantData**](MerchantData.md) |  | [optional] 
**TransferInstrumentId** | Pointer to **string** | Unique identifier of the [transfer instrument](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/transferInstruments__resParam_id). | [optional] 

## Methods

### NewCounterparty

`func NewCounterparty() *Counterparty`

NewCounterparty instantiates a new Counterparty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCounterpartyWithDefaults

`func NewCounterpartyWithDefaults() *Counterparty`

NewCounterpartyWithDefaults instantiates a new Counterparty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceAccountId

`func (o *Counterparty) GetBalanceAccountId() string`

GetBalanceAccountId returns the BalanceAccountId field if non-nil, zero value otherwise.

### GetBalanceAccountIdOk

`func (o *Counterparty) GetBalanceAccountIdOk() (*string, bool)`

GetBalanceAccountIdOk returns a tuple with the BalanceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAccountId

`func (o *Counterparty) SetBalanceAccountId(v string)`

SetBalanceAccountId sets BalanceAccountId field to given value.

### HasBalanceAccountId

`func (o *Counterparty) HasBalanceAccountId() bool`

HasBalanceAccountId returns a boolean if a field has been set.

### GetBankAccount

`func (o *Counterparty) GetBankAccount() BankAccountInfo`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *Counterparty) GetBankAccountOk() (*BankAccountInfo, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *Counterparty) SetBankAccount(v BankAccountInfo)`

SetBankAccount sets BankAccount field to given value.

### HasBankAccount

`func (o *Counterparty) HasBankAccount() bool`

HasBankAccount returns a boolean if a field has been set.

### GetMerchant

`func (o *Counterparty) GetMerchant() MerchantData`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *Counterparty) GetMerchantOk() (*MerchantData, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *Counterparty) SetMerchant(v MerchantData)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *Counterparty) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetTransferInstrumentId

`func (o *Counterparty) GetTransferInstrumentId() string`

GetTransferInstrumentId returns the TransferInstrumentId field if non-nil, zero value otherwise.

### GetTransferInstrumentIdOk

`func (o *Counterparty) GetTransferInstrumentIdOk() (*string, bool)`

GetTransferInstrumentIdOk returns a tuple with the TransferInstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferInstrumentId

`func (o *Counterparty) SetTransferInstrumentId(v string)`

SetTransferInstrumentId sets TransferInstrumentId field to given value.

### HasTransferInstrumentId

`func (o *Counterparty) HasTransferInstrumentId() bool`

HasTransferInstrumentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


