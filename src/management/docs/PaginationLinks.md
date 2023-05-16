# PaginationLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**First** | [**LinksElement**](LinksElement.md) |  | 
**Last** | [**LinksElement**](LinksElement.md) |  | 
**Next** | Pointer to [**LinksElement**](LinksElement.md) |  | [optional] 
**Prev** | Pointer to [**LinksElement**](LinksElement.md) |  | [optional] 
**Self** | [**LinksElement**](LinksElement.md) |  | 

## Methods

### NewPaginationLinks

`func NewPaginationLinks(first LinksElement, last LinksElement, self LinksElement, ) *PaginationLinks`

NewPaginationLinks instantiates a new PaginationLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationLinksWithDefaults

`func NewPaginationLinksWithDefaults() *PaginationLinks`

NewPaginationLinksWithDefaults instantiates a new PaginationLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirst

`func (o *PaginationLinks) GetFirst() LinksElement`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *PaginationLinks) GetFirstOk() (*LinksElement, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *PaginationLinks) SetFirst(v LinksElement)`

SetFirst sets First field to given value.


### GetLast

`func (o *PaginationLinks) GetLast() LinksElement`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *PaginationLinks) GetLastOk() (*LinksElement, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *PaginationLinks) SetLast(v LinksElement)`

SetLast sets Last field to given value.


### GetNext

`func (o *PaginationLinks) GetNext() LinksElement`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginationLinks) GetNextOk() (*LinksElement, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginationLinks) SetNext(v LinksElement)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginationLinks) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrev

`func (o *PaginationLinks) GetPrev() LinksElement`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *PaginationLinks) GetPrevOk() (*LinksElement, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *PaginationLinks) SetPrev(v LinksElement)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *PaginationLinks) HasPrev() bool`

HasPrev returns a boolean if a field has been set.

### GetSelf

`func (o *PaginationLinks) GetSelf() LinksElement`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *PaginationLinks) GetSelfOk() (*LinksElement, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *PaginationLinks) SetSelf(v LinksElement)`

SetSelf sets Self field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


