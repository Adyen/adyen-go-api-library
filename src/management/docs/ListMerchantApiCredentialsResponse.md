# ListMerchantApiCredentialsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] 
**Data** | Pointer to [**[]ApiCredential**](ApiCredential.md) | The list of API credentials. | [optional] 
**ItemsTotal** | **int32** | Total number of items. | 
**PagesTotal** | **int32** | Total number of pages. | 

## Methods

### NewListMerchantApiCredentialsResponse

`func NewListMerchantApiCredentialsResponse(itemsTotal int32, pagesTotal int32, ) *ListMerchantApiCredentialsResponse`

NewListMerchantApiCredentialsResponse instantiates a new ListMerchantApiCredentialsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMerchantApiCredentialsResponseWithDefaults

`func NewListMerchantApiCredentialsResponseWithDefaults() *ListMerchantApiCredentialsResponse`

NewListMerchantApiCredentialsResponseWithDefaults instantiates a new ListMerchantApiCredentialsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ListMerchantApiCredentialsResponse) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListMerchantApiCredentialsResponse) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListMerchantApiCredentialsResponse) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListMerchantApiCredentialsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetData

`func (o *ListMerchantApiCredentialsResponse) GetData() []ApiCredential`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListMerchantApiCredentialsResponse) GetDataOk() (*[]ApiCredential, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListMerchantApiCredentialsResponse) SetData(v []ApiCredential)`

SetData sets Data field to given value.

### HasData

`func (o *ListMerchantApiCredentialsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetItemsTotal

`func (o *ListMerchantApiCredentialsResponse) GetItemsTotal() int32`

GetItemsTotal returns the ItemsTotal field if non-nil, zero value otherwise.

### GetItemsTotalOk

`func (o *ListMerchantApiCredentialsResponse) GetItemsTotalOk() (*int32, bool)`

GetItemsTotalOk returns a tuple with the ItemsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsTotal

`func (o *ListMerchantApiCredentialsResponse) SetItemsTotal(v int32)`

SetItemsTotal sets ItemsTotal field to given value.


### GetPagesTotal

`func (o *ListMerchantApiCredentialsResponse) GetPagesTotal() int32`

GetPagesTotal returns the PagesTotal field if non-nil, zero value otherwise.

### GetPagesTotalOk

`func (o *ListMerchantApiCredentialsResponse) GetPagesTotalOk() (*int32, bool)`

GetPagesTotalOk returns a tuple with the PagesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagesTotal

`func (o *ListMerchantApiCredentialsResponse) SetPagesTotal(v int32)`

SetPagesTotal sets PagesTotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


