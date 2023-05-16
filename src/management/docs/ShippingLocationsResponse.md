# ShippingLocationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ShippingLocation**](ShippingLocation.md) | Physical locations where orders can be shipped to. | [optional] 

## Methods

### NewShippingLocationsResponse

`func NewShippingLocationsResponse() *ShippingLocationsResponse`

NewShippingLocationsResponse instantiates a new ShippingLocationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingLocationsResponseWithDefaults

`func NewShippingLocationsResponseWithDefaults() *ShippingLocationsResponse`

NewShippingLocationsResponseWithDefaults instantiates a new ShippingLocationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ShippingLocationsResponse) GetData() []ShippingLocation`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ShippingLocationsResponse) GetDataOk() (*[]ShippingLocation, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ShippingLocationsResponse) SetData(v []ShippingLocation)`

SetData sets Data field to given value.

### HasData

`func (o *ShippingLocationsResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


