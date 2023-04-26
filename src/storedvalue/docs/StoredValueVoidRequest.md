# StoredValueVoidRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantAccount** | **string** | The merchant account identifier, with which you want to process the transaction. | 
**OriginalReference** | **string** | The original pspReference of the payment to modify. | 
**Reference** | Pointer to **string** | Your reference for the payment modification. This reference is visible in Customer Area and in reports. Maximum length: 80 characters. | [optional] 
**Store** | Pointer to **string** | The physical store, for which this payment is processed. | [optional] 
**TenderReference** | Pointer to **string** | The reference of the tender. | [optional] 
**UniqueTerminalId** | Pointer to **string** | The unique ID of a POS terminal. | [optional] 

## Methods

### NewStoredValueVoidRequest

`func NewStoredValueVoidRequest(merchantAccount string, originalReference string, ) *StoredValueVoidRequest`

NewStoredValueVoidRequest instantiates a new StoredValueVoidRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoredValueVoidRequestWithDefaults

`func NewStoredValueVoidRequestWithDefaults() *StoredValueVoidRequest`

NewStoredValueVoidRequestWithDefaults instantiates a new StoredValueVoidRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantAccount

`func (o *StoredValueVoidRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *StoredValueVoidRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *StoredValueVoidRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetOriginalReference

`func (o *StoredValueVoidRequest) GetOriginalReference() string`

GetOriginalReference returns the OriginalReference field if non-nil, zero value otherwise.

### GetOriginalReferenceOk

`func (o *StoredValueVoidRequest) GetOriginalReferenceOk() (*string, bool)`

GetOriginalReferenceOk returns a tuple with the OriginalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalReference

`func (o *StoredValueVoidRequest) SetOriginalReference(v string)`

SetOriginalReference sets OriginalReference field to given value.


### GetReference

`func (o *StoredValueVoidRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *StoredValueVoidRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *StoredValueVoidRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *StoredValueVoidRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetStore

`func (o *StoredValueVoidRequest) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *StoredValueVoidRequest) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *StoredValueVoidRequest) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *StoredValueVoidRequest) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetTenderReference

`func (o *StoredValueVoidRequest) GetTenderReference() string`

GetTenderReference returns the TenderReference field if non-nil, zero value otherwise.

### GetTenderReferenceOk

`func (o *StoredValueVoidRequest) GetTenderReferenceOk() (*string, bool)`

GetTenderReferenceOk returns a tuple with the TenderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenderReference

`func (o *StoredValueVoidRequest) SetTenderReference(v string)`

SetTenderReference sets TenderReference field to given value.

### HasTenderReference

`func (o *StoredValueVoidRequest) HasTenderReference() bool`

HasTenderReference returns a boolean if a field has been set.

### GetUniqueTerminalId

`func (o *StoredValueVoidRequest) GetUniqueTerminalId() string`

GetUniqueTerminalId returns the UniqueTerminalId field if non-nil, zero value otherwise.

### GetUniqueTerminalIdOk

`func (o *StoredValueVoidRequest) GetUniqueTerminalIdOk() (*string, bool)`

GetUniqueTerminalIdOk returns a tuple with the UniqueTerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueTerminalId

`func (o *StoredValueVoidRequest) SetUniqueTerminalId(v string)`

SetUniqueTerminalId sets UniqueTerminalId field to given value.

### HasUniqueTerminalId

`func (o *StoredValueVoidRequest) HasUniqueTerminalId() bool`

HasUniqueTerminalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


