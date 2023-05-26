# MerchantAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InStoreTerminals** | Pointer to **[]string** | List of terminals assigned to this merchant account as in-store terminals. This means that the terminal is ready to be boarded, or is already boarded. | [optional] 
**InventoryTerminals** | Pointer to **[]string** | List of terminals assigned to the inventory of this merchant account. | [optional] 
**MerchantAccount** | **string** | The merchant account. | 
**Stores** | Pointer to [**[]Store**](Store.md) | Array of stores under this merchant account. | [optional] 

## Methods

### NewMerchantAccount

`func NewMerchantAccount(merchantAccount string, ) *MerchantAccount`

NewMerchantAccount instantiates a new MerchantAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantAccountWithDefaults

`func NewMerchantAccountWithDefaults() *MerchantAccount`

NewMerchantAccountWithDefaults instantiates a new MerchantAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInStoreTerminals

`func (o *MerchantAccount) GetInStoreTerminals() []string`

GetInStoreTerminals returns the InStoreTerminals field if non-nil, zero value otherwise.

### GetInStoreTerminalsOk

`func (o *MerchantAccount) GetInStoreTerminalsOk() (*[]string, bool)`

GetInStoreTerminalsOk returns a tuple with the InStoreTerminals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInStoreTerminals

`func (o *MerchantAccount) SetInStoreTerminals(v []string)`

SetInStoreTerminals sets InStoreTerminals field to given value.

### HasInStoreTerminals

`func (o *MerchantAccount) HasInStoreTerminals() bool`

HasInStoreTerminals returns a boolean if a field has been set.

### GetInventoryTerminals

`func (o *MerchantAccount) GetInventoryTerminals() []string`

GetInventoryTerminals returns the InventoryTerminals field if non-nil, zero value otherwise.

### GetInventoryTerminalsOk

`func (o *MerchantAccount) GetInventoryTerminalsOk() (*[]string, bool)`

GetInventoryTerminalsOk returns a tuple with the InventoryTerminals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryTerminals

`func (o *MerchantAccount) SetInventoryTerminals(v []string)`

SetInventoryTerminals sets InventoryTerminals field to given value.

### HasInventoryTerminals

`func (o *MerchantAccount) HasInventoryTerminals() bool`

HasInventoryTerminals returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *MerchantAccount) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *MerchantAccount) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *MerchantAccount) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetStores

`func (o *MerchantAccount) GetStores() []Store`

GetStores returns the Stores field if non-nil, zero value otherwise.

### GetStoresOk

`func (o *MerchantAccount) GetStoresOk() (*[]Store, bool)`

GetStoresOk returns a tuple with the Stores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStores

`func (o *MerchantAccount) SetStores(v []Store)`

SetStores sets Stores field to given value.

### HasStores

`func (o *MerchantAccount) HasStores() bool`

HasStores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


