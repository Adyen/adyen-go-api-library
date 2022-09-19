# CreateAllowedOriginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Domain** | **string** | Domain of the allowed origin. | 
**Id** | Pointer to **string** | Unique identifier of the allowed origin. | [optional] 

## Methods

### NewCreateAllowedOriginRequest

`func NewCreateAllowedOriginRequest(domain string, ) *CreateAllowedOriginRequest`

NewCreateAllowedOriginRequest instantiates a new CreateAllowedOriginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAllowedOriginRequestWithDefaults

`func NewCreateAllowedOriginRequestWithDefaults() *CreateAllowedOriginRequest`

NewCreateAllowedOriginRequestWithDefaults instantiates a new CreateAllowedOriginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *CreateAllowedOriginRequest) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateAllowedOriginRequest) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateAllowedOriginRequest) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateAllowedOriginRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetDomain

`func (o *CreateAllowedOriginRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateAllowedOriginRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateAllowedOriginRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetId

`func (o *CreateAllowedOriginRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateAllowedOriginRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateAllowedOriginRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateAllowedOriginRequest) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


