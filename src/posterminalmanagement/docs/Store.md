# Store

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**Description** | Pointer to **string** | The description of the store. | [optional] 
**InStoreTerminals** | Pointer to **[]string** | The list of terminals assigned to the store. | [optional] 
**MerchantAccountCode** | Pointer to **string** | The code of the merchant account. | [optional] 
**Status** | Pointer to **string** | The status of the store:  - &#x60;PreActive&#x60;: the store has been created, but not yet activated.   - &#x60;Active&#x60;: the store has been activated. This means you can process payments for this store.   - &#x60;Inactive&#x60;: the store is currently not active.   - &#x60;InactiveWithModifications&#x60;: the store is currently not active, but payment modifications such as refunds are possible.   - &#x60;Closed&#x60;: the store has been closed.  | [optional] 
**Store** | **string** | The code of the store. | 

## Methods

### NewStore

`func NewStore(store string, ) *Store`

NewStore instantiates a new Store object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreWithDefaults

`func NewStoreWithDefaults() *Store`

NewStoreWithDefaults instantiates a new Store object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Store) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Store) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Store) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Store) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDescription

`func (o *Store) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Store) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Store) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Store) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInStoreTerminals

`func (o *Store) GetInStoreTerminals() []string`

GetInStoreTerminals returns the InStoreTerminals field if non-nil, zero value otherwise.

### GetInStoreTerminalsOk

`func (o *Store) GetInStoreTerminalsOk() (*[]string, bool)`

GetInStoreTerminalsOk returns a tuple with the InStoreTerminals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInStoreTerminals

`func (o *Store) SetInStoreTerminals(v []string)`

SetInStoreTerminals sets InStoreTerminals field to given value.

### HasInStoreTerminals

`func (o *Store) HasInStoreTerminals() bool`

HasInStoreTerminals returns a boolean if a field has been set.

### GetMerchantAccountCode

`func (o *Store) GetMerchantAccountCode() string`

GetMerchantAccountCode returns the MerchantAccountCode field if non-nil, zero value otherwise.

### GetMerchantAccountCodeOk

`func (o *Store) GetMerchantAccountCodeOk() (*string, bool)`

GetMerchantAccountCodeOk returns a tuple with the MerchantAccountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountCode

`func (o *Store) SetMerchantAccountCode(v string)`

SetMerchantAccountCode sets MerchantAccountCode field to given value.

### HasMerchantAccountCode

`func (o *Store) HasMerchantAccountCode() bool`

HasMerchantAccountCode returns a boolean if a field has been set.

### GetStatus

`func (o *Store) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Store) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Store) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Store) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStore

`func (o *Store) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *Store) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *Store) SetStore(v string)`

SetStore sets Store field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


