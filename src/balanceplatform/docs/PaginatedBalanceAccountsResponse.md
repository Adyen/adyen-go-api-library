# PaginatedBalanceAccountsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceAccounts** | [**[]BalanceAccountBase**](BalanceAccountBase.md) | List of balance accounts. | 
**HasNext** | **bool** | Indicates whether there are more items on the next page. | 
**HasPrevious** | **bool** | Indicates whether there are more items on the previous page. | 

## Methods

### NewPaginatedBalanceAccountsResponse

`func NewPaginatedBalanceAccountsResponse(balanceAccounts []BalanceAccountBase, hasNext bool, hasPrevious bool, ) *PaginatedBalanceAccountsResponse`

NewPaginatedBalanceAccountsResponse instantiates a new PaginatedBalanceAccountsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedBalanceAccountsResponseWithDefaults

`func NewPaginatedBalanceAccountsResponseWithDefaults() *PaginatedBalanceAccountsResponse`

NewPaginatedBalanceAccountsResponseWithDefaults instantiates a new PaginatedBalanceAccountsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceAccounts

`func (o *PaginatedBalanceAccountsResponse) GetBalanceAccounts() []BalanceAccountBase`

GetBalanceAccounts returns the BalanceAccounts field if non-nil, zero value otherwise.

### GetBalanceAccountsOk

`func (o *PaginatedBalanceAccountsResponse) GetBalanceAccountsOk() (*[]BalanceAccountBase, bool)`

GetBalanceAccountsOk returns a tuple with the BalanceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAccounts

`func (o *PaginatedBalanceAccountsResponse) SetBalanceAccounts(v []BalanceAccountBase)`

SetBalanceAccounts sets BalanceAccounts field to given value.


### GetHasNext

`func (o *PaginatedBalanceAccountsResponse) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *PaginatedBalanceAccountsResponse) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *PaginatedBalanceAccountsResponse) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.


### GetHasPrevious

`func (o *PaginatedBalanceAccountsResponse) GetHasPrevious() bool`

GetHasPrevious returns the HasPrevious field if non-nil, zero value otherwise.

### GetHasPreviousOk

`func (o *PaginatedBalanceAccountsResponse) GetHasPreviousOk() (*bool, bool)`

GetHasPreviousOk returns a tuple with the HasPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrevious

`func (o *PaginatedBalanceAccountsResponse) SetHasPrevious(v bool)`

SetHasPrevious sets HasPrevious field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


