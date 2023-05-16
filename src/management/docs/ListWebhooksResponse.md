# ListWebhooksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] 
**AccountReference** | Pointer to **string** | Reference to the account. | [optional] 
**Data** | Pointer to [**[]Webhook**](Webhook.md) | The list of webhooks configured for this account. | [optional] 
**ItemsTotal** | **int32** | Total number of items. | 
**PagesTotal** | **int32** | Total number of pages. | 

## Methods

### NewListWebhooksResponse

`func NewListWebhooksResponse(itemsTotal int32, pagesTotal int32, ) *ListWebhooksResponse`

NewListWebhooksResponse instantiates a new ListWebhooksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWebhooksResponseWithDefaults

`func NewListWebhooksResponseWithDefaults() *ListWebhooksResponse`

NewListWebhooksResponseWithDefaults instantiates a new ListWebhooksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ListWebhooksResponse) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListWebhooksResponse) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListWebhooksResponse) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListWebhooksResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAccountReference

`func (o *ListWebhooksResponse) GetAccountReference() string`

GetAccountReference returns the AccountReference field if non-nil, zero value otherwise.

### GetAccountReferenceOk

`func (o *ListWebhooksResponse) GetAccountReferenceOk() (*string, bool)`

GetAccountReferenceOk returns a tuple with the AccountReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountReference

`func (o *ListWebhooksResponse) SetAccountReference(v string)`

SetAccountReference sets AccountReference field to given value.

### HasAccountReference

`func (o *ListWebhooksResponse) HasAccountReference() bool`

HasAccountReference returns a boolean if a field has been set.

### GetData

`func (o *ListWebhooksResponse) GetData() []Webhook`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListWebhooksResponse) GetDataOk() (*[]Webhook, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListWebhooksResponse) SetData(v []Webhook)`

SetData sets Data field to given value.

### HasData

`func (o *ListWebhooksResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetItemsTotal

`func (o *ListWebhooksResponse) GetItemsTotal() int32`

GetItemsTotal returns the ItemsTotal field if non-nil, zero value otherwise.

### GetItemsTotalOk

`func (o *ListWebhooksResponse) GetItemsTotalOk() (*int32, bool)`

GetItemsTotalOk returns a tuple with the ItemsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsTotal

`func (o *ListWebhooksResponse) SetItemsTotal(v int32)`

SetItemsTotal sets ItemsTotal field to given value.


### GetPagesTotal

`func (o *ListWebhooksResponse) GetPagesTotal() int32`

GetPagesTotal returns the PagesTotal field if non-nil, zero value otherwise.

### GetPagesTotalOk

`func (o *ListWebhooksResponse) GetPagesTotalOk() (*int32, bool)`

GetPagesTotalOk returns a tuple with the PagesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagesTotal

`func (o *ListWebhooksResponse) SetPagesTotal(v int32)`

SetPagesTotal sets PagesTotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


