# TransferInstrumentReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIdentifier** | **string** | The masked IBAN or bank account number. | 
**Id** | **string** | The unique identifier of the resource. | 
**RealLastFour** | Pointer to **string** | Four last digits of the bank account number. | [optional] 

## Methods

### NewTransferInstrumentReference

`func NewTransferInstrumentReference(accountIdentifier string, id string, ) *TransferInstrumentReference`

NewTransferInstrumentReference instantiates a new TransferInstrumentReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferInstrumentReferenceWithDefaults

`func NewTransferInstrumentReferenceWithDefaults() *TransferInstrumentReference`

NewTransferInstrumentReferenceWithDefaults instantiates a new TransferInstrumentReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountIdentifier

`func (o *TransferInstrumentReference) GetAccountIdentifier() string`

GetAccountIdentifier returns the AccountIdentifier field if non-nil, zero value otherwise.

### GetAccountIdentifierOk

`func (o *TransferInstrumentReference) GetAccountIdentifierOk() (*string, bool)`

GetAccountIdentifierOk returns a tuple with the AccountIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdentifier

`func (o *TransferInstrumentReference) SetAccountIdentifier(v string)`

SetAccountIdentifier sets AccountIdentifier field to given value.


### GetId

`func (o *TransferInstrumentReference) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransferInstrumentReference) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransferInstrumentReference) SetId(v string)`

SetId sets Id field to given value.


### GetRealLastFour

`func (o *TransferInstrumentReference) GetRealLastFour() string`

GetRealLastFour returns the RealLastFour field if non-nil, zero value otherwise.

### GetRealLastFourOk

`func (o *TransferInstrumentReference) GetRealLastFourOk() (*string, bool)`

GetRealLastFourOk returns a tuple with the RealLastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealLastFour

`func (o *TransferInstrumentReference) SetRealLastFour(v string)`

SetRealLastFour sets RealLastFour field to given value.

### HasRealLastFour

`func (o *TransferInstrumentReference) HasRealLastFour() bool`

HasRealLastFour returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


