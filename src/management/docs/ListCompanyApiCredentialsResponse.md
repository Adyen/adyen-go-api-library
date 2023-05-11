# ListCompanyApiCredentialsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] 
**Data** | Pointer to [**[]CompanyApiCredential**](CompanyApiCredential.md) | The list of API credentials. | [optional] 
**ItemsTotal** | **int32** | Total number of items. | 
**PagesTotal** | **int32** | Total number of pages. | 

## Methods

### NewListCompanyApiCredentialsResponse

`func NewListCompanyApiCredentialsResponse(itemsTotal int32, pagesTotal int32, ) *ListCompanyApiCredentialsResponse`

NewListCompanyApiCredentialsResponse instantiates a new ListCompanyApiCredentialsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCompanyApiCredentialsResponseWithDefaults

`func NewListCompanyApiCredentialsResponseWithDefaults() *ListCompanyApiCredentialsResponse`

NewListCompanyApiCredentialsResponseWithDefaults instantiates a new ListCompanyApiCredentialsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ListCompanyApiCredentialsResponse) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListCompanyApiCredentialsResponse) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListCompanyApiCredentialsResponse) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListCompanyApiCredentialsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetData

`func (o *ListCompanyApiCredentialsResponse) GetData() []CompanyApiCredential`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListCompanyApiCredentialsResponse) GetDataOk() (*[]CompanyApiCredential, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListCompanyApiCredentialsResponse) SetData(v []CompanyApiCredential)`

SetData sets Data field to given value.

### HasData

`func (o *ListCompanyApiCredentialsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetItemsTotal

`func (o *ListCompanyApiCredentialsResponse) GetItemsTotal() int32`

GetItemsTotal returns the ItemsTotal field if non-nil, zero value otherwise.

### GetItemsTotalOk

`func (o *ListCompanyApiCredentialsResponse) GetItemsTotalOk() (*int32, bool)`

GetItemsTotalOk returns a tuple with the ItemsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsTotal

`func (o *ListCompanyApiCredentialsResponse) SetItemsTotal(v int32)`

SetItemsTotal sets ItemsTotal field to given value.


### GetPagesTotal

`func (o *ListCompanyApiCredentialsResponse) GetPagesTotal() int32`

GetPagesTotal returns the PagesTotal field if non-nil, zero value otherwise.

### GetPagesTotalOk

`func (o *ListCompanyApiCredentialsResponse) GetPagesTotalOk() (*int32, bool)`

GetPagesTotalOk returns a tuple with the PagesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagesTotal

`func (o *ListCompanyApiCredentialsResponse) SetPagesTotal(v int32)`

SetPagesTotal sets PagesTotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


