# CounterpartyV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceAccountId** | Pointer to **string** | Unique identifier of the [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts__resParam_id). | [optional] 
**BankAccount** | Pointer to [**BankAccountV3**](BankAccountV3.md) |  | [optional] 
**Merchant** | Pointer to [**MerchantData**](MerchantData.md) |  | [optional] 
**TransferInstrumentId** | Pointer to **string** | Unique identifier of the [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments__resParam_id). | [optional] 

## Methods

### NewCounterpartyV3

`func NewCounterpartyV3() *CounterpartyV3`

NewCounterpartyV3 instantiates a new CounterpartyV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCounterpartyV3WithDefaults

`func NewCounterpartyV3WithDefaults() *CounterpartyV3`

NewCounterpartyV3WithDefaults instantiates a new CounterpartyV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceAccountId

`func (o *CounterpartyV3) GetBalanceAccountId() string`

GetBalanceAccountId returns the BalanceAccountId field if non-nil, zero value otherwise.

### GetBalanceAccountIdOk

`func (o *CounterpartyV3) GetBalanceAccountIdOk() (*string, bool)`

GetBalanceAccountIdOk returns a tuple with the BalanceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAccountId

`func (o *CounterpartyV3) SetBalanceAccountId(v string)`

SetBalanceAccountId sets BalanceAccountId field to given value.

### HasBalanceAccountId

`func (o *CounterpartyV3) HasBalanceAccountId() bool`

HasBalanceAccountId returns a boolean if a field has been set.

### GetBankAccount

`func (o *CounterpartyV3) GetBankAccount() BankAccountV3`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *CounterpartyV3) GetBankAccountOk() (*BankAccountV3, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *CounterpartyV3) SetBankAccount(v BankAccountV3)`

SetBankAccount sets BankAccount field to given value.

### HasBankAccount

`func (o *CounterpartyV3) HasBankAccount() bool`

HasBankAccount returns a boolean if a field has been set.

### GetMerchant

`func (o *CounterpartyV3) GetMerchant() MerchantData`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *CounterpartyV3) GetMerchantOk() (*MerchantData, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *CounterpartyV3) SetMerchant(v MerchantData)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *CounterpartyV3) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetTransferInstrumentId

`func (o *CounterpartyV3) GetTransferInstrumentId() string`

GetTransferInstrumentId returns the TransferInstrumentId field if non-nil, zero value otherwise.

### GetTransferInstrumentIdOk

`func (o *CounterpartyV3) GetTransferInstrumentIdOk() (*string, bool)`

GetTransferInstrumentIdOk returns a tuple with the TransferInstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferInstrumentId

`func (o *CounterpartyV3) SetTransferInstrumentId(v string)`

SetTransferInstrumentId sets TransferInstrumentId field to given value.

### HasTransferInstrumentId

`func (o *CounterpartyV3) HasTransferInstrumentId() bool`

HasTransferInstrumentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


