# AllowedOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Domain** | **string** | Domain of the allowed origin. | 
**Id** | Pointer to **string** | Unique identifier of the allowed origin. | [optional] 

## Methods

### NewAllowedOrigin

`func NewAllowedOrigin(domain string, ) *AllowedOrigin`

NewAllowedOrigin instantiates a new AllowedOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowedOriginWithDefaults

`func NewAllowedOriginWithDefaults() *AllowedOrigin`

NewAllowedOriginWithDefaults instantiates a new AllowedOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AllowedOrigin) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AllowedOrigin) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AllowedOrigin) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AllowedOrigin) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetDomain

`func (o *AllowedOrigin) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *AllowedOrigin) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *AllowedOrigin) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetId

`func (o *AllowedOrigin) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AllowedOrigin) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AllowedOrigin) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AllowedOrigin) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


