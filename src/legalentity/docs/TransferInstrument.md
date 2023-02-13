# TransferInstrument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankAccount** | [**BankAccountInfo**](BankAccountInfo.md) |  | 
**DocumentDetails** | Pointer to [**[]DocumentReference**](DocumentReference.md) | List of documents uploaded for the transfer instrument. | [optional] 
**Id** | **string** | The unique identifier of the transfer instrument. | [readonly] 
**LegalEntityId** | **string** | The unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/legalEntities__resParam_id) that owns the transfer instrument. | 
**Type** | **string** | The type of transfer instrument.  Possible value: **bankAccount**. | 

## Methods

### NewTransferInstrument

`func NewTransferInstrument(bankAccount BankAccountInfo, id string, legalEntityId string, type_ string, ) *TransferInstrument`

NewTransferInstrument instantiates a new TransferInstrument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferInstrumentWithDefaults

`func NewTransferInstrumentWithDefaults() *TransferInstrument`

NewTransferInstrumentWithDefaults instantiates a new TransferInstrument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankAccount

`func (o *TransferInstrument) GetBankAccount() BankAccountInfo`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *TransferInstrument) GetBankAccountOk() (*BankAccountInfo, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *TransferInstrument) SetBankAccount(v BankAccountInfo)`

SetBankAccount sets BankAccount field to given value.


### GetDocumentDetails

`func (o *TransferInstrument) GetDocumentDetails() []DocumentReference`

GetDocumentDetails returns the DocumentDetails field if non-nil, zero value otherwise.

### GetDocumentDetailsOk

`func (o *TransferInstrument) GetDocumentDetailsOk() (*[]DocumentReference, bool)`

GetDocumentDetailsOk returns a tuple with the DocumentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentDetails

`func (o *TransferInstrument) SetDocumentDetails(v []DocumentReference)`

SetDocumentDetails sets DocumentDetails field to given value.

### HasDocumentDetails

`func (o *TransferInstrument) HasDocumentDetails() bool`

HasDocumentDetails returns a boolean if a field has been set.

### GetId

`func (o *TransferInstrument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransferInstrument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransferInstrument) SetId(v string)`

SetId sets Id field to given value.


### GetLegalEntityId

`func (o *TransferInstrument) GetLegalEntityId() string`

GetLegalEntityId returns the LegalEntityId field if non-nil, zero value otherwise.

### GetLegalEntityIdOk

`func (o *TransferInstrument) GetLegalEntityIdOk() (*string, bool)`

GetLegalEntityIdOk returns a tuple with the LegalEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalEntityId

`func (o *TransferInstrument) SetLegalEntityId(v string)`

SetLegalEntityId sets LegalEntityId field to given value.


### GetType

`func (o *TransferInstrument) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransferInstrument) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransferInstrument) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


