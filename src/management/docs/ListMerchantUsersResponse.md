# ListMerchantUsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] 
**Data** | Pointer to [**[]User**](User.md) | The list of users. | [optional] 
**ItemsTotal** | **int32** | Total number of items. | 
**PagesTotal** | **int32** | Total number of pages. | 

## Methods

### NewListMerchantUsersResponse

`func NewListMerchantUsersResponse(itemsTotal int32, pagesTotal int32, ) *ListMerchantUsersResponse`

NewListMerchantUsersResponse instantiates a new ListMerchantUsersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMerchantUsersResponseWithDefaults

`func NewListMerchantUsersResponseWithDefaults() *ListMerchantUsersResponse`

NewListMerchantUsersResponseWithDefaults instantiates a new ListMerchantUsersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ListMerchantUsersResponse) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListMerchantUsersResponse) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListMerchantUsersResponse) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListMerchantUsersResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetData

`func (o *ListMerchantUsersResponse) GetData() []User`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListMerchantUsersResponse) GetDataOk() (*[]User, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListMerchantUsersResponse) SetData(v []User)`

SetData sets Data field to given value.

### HasData

`func (o *ListMerchantUsersResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetItemsTotal

`func (o *ListMerchantUsersResponse) GetItemsTotal() int32`

GetItemsTotal returns the ItemsTotal field if non-nil, zero value otherwise.

### GetItemsTotalOk

`func (o *ListMerchantUsersResponse) GetItemsTotalOk() (*int32, bool)`

GetItemsTotalOk returns a tuple with the ItemsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsTotal

`func (o *ListMerchantUsersResponse) SetItemsTotal(v int32)`

SetItemsTotal sets ItemsTotal field to given value.


### GetPagesTotal

`func (o *ListMerchantUsersResponse) GetPagesTotal() int32`

GetPagesTotal returns the PagesTotal field if non-nil, zero value otherwise.

### GetPagesTotalOk

`func (o *ListMerchantUsersResponse) GetPagesTotalOk() (*int32, bool)`

GetPagesTotalOk returns a tuple with the PagesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagesTotal

`func (o *ListMerchantUsersResponse) SetPagesTotal(v int32)`

SetPagesTotal sets PagesTotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


