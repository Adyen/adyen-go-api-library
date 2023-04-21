# TransferInstrumentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankAccount** | [**BankAccountInfo**](BankAccountInfo.md) |  | 
**LegalEntityId** | **string** | The unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/legalentity/latest/post/legalEntities#responses-200-id) that owns the transfer instrument. | 
**Type** | **string** | The type of transfer instrument.  Possible value: **bankAccount**. | 

## Methods

### NewTransferInstrumentInfo

`func NewTransferInstrumentInfo(bankAccount BankAccountInfo, legalEntityId string, type_ string, ) *TransferInstrumentInfo`

NewTransferInstrumentInfo instantiates a new TransferInstrumentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferInstrumentInfoWithDefaults

`func NewTransferInstrumentInfoWithDefaults() *TransferInstrumentInfo`

NewTransferInstrumentInfoWithDefaults instantiates a new TransferInstrumentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankAccount

`func (o *TransferInstrumentInfo) GetBankAccount() BankAccountInfo`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *TransferInstrumentInfo) GetBankAccountOk() (*BankAccountInfo, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *TransferInstrumentInfo) SetBankAccount(v BankAccountInfo)`

SetBankAccount sets BankAccount field to given value.


### GetLegalEntityId

`func (o *TransferInstrumentInfo) GetLegalEntityId() string`

GetLegalEntityId returns the LegalEntityId field if non-nil, zero value otherwise.

### GetLegalEntityIdOk

`func (o *TransferInstrumentInfo) GetLegalEntityIdOk() (*string, bool)`

GetLegalEntityIdOk returns a tuple with the LegalEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalEntityId

`func (o *TransferInstrumentInfo) SetLegalEntityId(v string)`

SetLegalEntityId sets LegalEntityId field to given value.


### GetType

`func (o *TransferInstrumentInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransferInstrumentInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransferInstrumentInfo) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


