# GetStoresUnderAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stores** | Pointer to [**[]Store**](Store.md) | Array that returns a list of all stores for the specified merchant account, or for all merchant accounts under the company account. | [optional] 

## Methods

### NewGetStoresUnderAccountResponse

`func NewGetStoresUnderAccountResponse() *GetStoresUnderAccountResponse`

NewGetStoresUnderAccountResponse instantiates a new GetStoresUnderAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStoresUnderAccountResponseWithDefaults

`func NewGetStoresUnderAccountResponseWithDefaults() *GetStoresUnderAccountResponse`

NewGetStoresUnderAccountResponseWithDefaults instantiates a new GetStoresUnderAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStores

`func (o *GetStoresUnderAccountResponse) GetStores() []Store`

GetStores returns the Stores field if non-nil, zero value otherwise.

### GetStoresOk

`func (o *GetStoresUnderAccountResponse) GetStoresOk() (*[]Store, bool)`

GetStoresOk returns a tuple with the Stores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStores

`func (o *GetStoresUnderAccountResponse) SetStores(v []Store)`

SetStores sets Stores field to given value.

### HasStores

`func (o *GetStoresUnderAccountResponse) HasStores() bool`

HasStores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


