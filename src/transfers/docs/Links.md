# Links

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to [**Link**](Link.md) |  | [optional] 
**Prev** | Pointer to [**Link**](Link.md) |  | [optional] 

## Methods

### NewLinks

`func NewLinks() *Links`

NewLinks instantiates a new Links object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinksWithDefaults

`func NewLinksWithDefaults() *Links`

NewLinksWithDefaults instantiates a new Links object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *Links) GetNext() Link`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *Links) GetNextOk() (*Link, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *Links) SetNext(v Link)`

SetNext sets Next field to given value.

### HasNext

`func (o *Links) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrev

`func (o *Links) GetPrev() Link`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *Links) GetPrevOk() (*Link, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *Links) SetPrev(v Link)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *Links) HasPrev() bool`

HasPrev returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


