# CounterpartyInfoV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceAccountId** | Pointer to **string** | Unique identifier of the [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts__resParam_id). | [optional] 
**BankAccount** | Pointer to [**BankAccountV3**](BankAccountV3.md) |  | [optional] 
**TransferInstrumentId** | Pointer to **string** | Unique identifier of the [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments__resParam_id). | [optional] 

## Methods

### NewCounterpartyInfoV3

`func NewCounterpartyInfoV3() *CounterpartyInfoV3`

NewCounterpartyInfoV3 instantiates a new CounterpartyInfoV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCounterpartyInfoV3WithDefaults

`func NewCounterpartyInfoV3WithDefaults() *CounterpartyInfoV3`

NewCounterpartyInfoV3WithDefaults instantiates a new CounterpartyInfoV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceAccountId

`func (o *CounterpartyInfoV3) GetBalanceAccountId() string`

GetBalanceAccountId returns the BalanceAccountId field if non-nil, zero value otherwise.

### GetBalanceAccountIdOk

`func (o *CounterpartyInfoV3) GetBalanceAccountIdOk() (*string, bool)`

GetBalanceAccountIdOk returns a tuple with the BalanceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAccountId

`func (o *CounterpartyInfoV3) SetBalanceAccountId(v string)`

SetBalanceAccountId sets BalanceAccountId field to given value.

### HasBalanceAccountId

`func (o *CounterpartyInfoV3) HasBalanceAccountId() bool`

HasBalanceAccountId returns a boolean if a field has been set.

### GetBankAccount

`func (o *CounterpartyInfoV3) GetBankAccount() BankAccountV3`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *CounterpartyInfoV3) GetBankAccountOk() (*BankAccountV3, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *CounterpartyInfoV3) SetBankAccount(v BankAccountV3)`

SetBankAccount sets BankAccount field to given value.

### HasBankAccount

`func (o *CounterpartyInfoV3) HasBankAccount() bool`

HasBankAccount returns a boolean if a field has been set.

### GetTransferInstrumentId

`func (o *CounterpartyInfoV3) GetTransferInstrumentId() string`

GetTransferInstrumentId returns the TransferInstrumentId field if non-nil, zero value otherwise.

### GetTransferInstrumentIdOk

`func (o *CounterpartyInfoV3) GetTransferInstrumentIdOk() (*string, bool)`

GetTransferInstrumentIdOk returns a tuple with the TransferInstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferInstrumentId

`func (o *CounterpartyInfoV3) SetTransferInstrumentId(v string)`

SetTransferInstrumentId sets TransferInstrumentId field to given value.

### HasTransferInstrumentId

`func (o *CounterpartyInfoV3) HasTransferInstrumentId() bool`

HasTransferInstrumentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

