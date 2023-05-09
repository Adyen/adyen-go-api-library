# PaginatedAccountHoldersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountHolders** | [**[]AccountHolder**](AccountHolder.md) | List of account holders. | 
**HasNext** | **bool** | Indicates whether there are more items on the next page. | 
**HasPrevious** | **bool** | Indicates whether there are more items on the previous page. | 

## Methods

### NewPaginatedAccountHoldersResponse

`func NewPaginatedAccountHoldersResponse(accountHolders []AccountHolder, hasNext bool, hasPrevious bool, ) *PaginatedAccountHoldersResponse`

NewPaginatedAccountHoldersResponse instantiates a new PaginatedAccountHoldersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedAccountHoldersResponseWithDefaults

`func NewPaginatedAccountHoldersResponseWithDefaults() *PaginatedAccountHoldersResponse`

NewPaginatedAccountHoldersResponseWithDefaults instantiates a new PaginatedAccountHoldersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountHolders

`func (o *PaginatedAccountHoldersResponse) GetAccountHolders() []AccountHolder`

GetAccountHolders returns the AccountHolders field if non-nil, zero value otherwise.

### GetAccountHoldersOk

`func (o *PaginatedAccountHoldersResponse) GetAccountHoldersOk() (*[]AccountHolder, bool)`

GetAccountHoldersOk returns a tuple with the AccountHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolders

`func (o *PaginatedAccountHoldersResponse) SetAccountHolders(v []AccountHolder)`

SetAccountHolders sets AccountHolders field to given value.


### GetHasNext

`func (o *PaginatedAccountHoldersResponse) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *PaginatedAccountHoldersResponse) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *PaginatedAccountHoldersResponse) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.


### GetHasPrevious

`func (o *PaginatedAccountHoldersResponse) GetHasPrevious() bool`

GetHasPrevious returns the HasPrevious field if non-nil, zero value otherwise.

### GetHasPreviousOk

`func (o *PaginatedAccountHoldersResponse) GetHasPreviousOk() (*bool, bool)`

GetHasPreviousOk returns a tuple with the HasPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrevious

`func (o *PaginatedAccountHoldersResponse) SetHasPrevious(v bool)`

SetHasPrevious sets HasPrevious field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


