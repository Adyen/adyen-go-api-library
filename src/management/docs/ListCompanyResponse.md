# ListCompanyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] 
**Data** | Pointer to [**[]Company**](Company.md) | The list of companies. | [optional] 
**ItemsTotal** | **int32** | Total number of items. | 
**PagesTotal** | **int32** | Total number of pages. | 

## Methods

### NewListCompanyResponse

`func NewListCompanyResponse(itemsTotal int32, pagesTotal int32, ) *ListCompanyResponse`

NewListCompanyResponse instantiates a new ListCompanyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCompanyResponseWithDefaults

`func NewListCompanyResponseWithDefaults() *ListCompanyResponse`

NewListCompanyResponseWithDefaults instantiates a new ListCompanyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ListCompanyResponse) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListCompanyResponse) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListCompanyResponse) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListCompanyResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetData

`func (o *ListCompanyResponse) GetData() []Company`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListCompanyResponse) GetDataOk() (*[]Company, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListCompanyResponse) SetData(v []Company)`

SetData sets Data field to given value.

### HasData

`func (o *ListCompanyResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetItemsTotal

`func (o *ListCompanyResponse) GetItemsTotal() int32`

GetItemsTotal returns the ItemsTotal field if non-nil, zero value otherwise.

### GetItemsTotalOk

`func (o *ListCompanyResponse) GetItemsTotalOk() (*int32, bool)`

GetItemsTotalOk returns a tuple with the ItemsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsTotal

`func (o *ListCompanyResponse) SetItemsTotal(v int32)`

SetItemsTotal sets ItemsTotal field to given value.


### GetPagesTotal

`func (o *ListCompanyResponse) GetPagesTotal() int32`

GetPagesTotal returns the PagesTotal field if non-nil, zero value otherwise.

### GetPagesTotalOk

`func (o *ListCompanyResponse) GetPagesTotalOk() (*int32, bool)`

GetPagesTotalOk returns a tuple with the PagesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagesTotal

`func (o *ListCompanyResponse) SetPagesTotal(v int32)`

SetPagesTotal sets PagesTotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


