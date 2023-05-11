# ShippingLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**Contact** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**Id** | Pointer to **string** | The unique identifier of the shipping location, for use as &#x60;shippingLocationId&#x60; when creating an order. | [optional] 
**Name** | Pointer to **string** | The unique name of the shipping location. | [optional] 

## Methods

### NewShippingLocation

`func NewShippingLocation() *ShippingLocation`

NewShippingLocation instantiates a new ShippingLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingLocationWithDefaults

`func NewShippingLocationWithDefaults() *ShippingLocation`

NewShippingLocationWithDefaults instantiates a new ShippingLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ShippingLocation) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ShippingLocation) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ShippingLocation) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ShippingLocation) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetContact

`func (o *ShippingLocation) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ShippingLocation) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ShippingLocation) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ShippingLocation) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetId

`func (o *ShippingLocation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShippingLocation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShippingLocation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShippingLocation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ShippingLocation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShippingLocation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShippingLocation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShippingLocation) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


