# TerminalOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingEntity** | Pointer to [**BillingEntity**](BillingEntity.md) |  | [optional] 
**CustomerOrderReference** | Pointer to **string** | The merchant-defined purchase order number. This will be printed on the packing list. | [optional] 
**Id** | Pointer to **string** | The unique identifier of the order. | [optional] 
**Items** | Pointer to [**[]OrderItem**](OrderItem.md) | The products included in the order. | [optional] 
**OrderDate** | Pointer to **string** | The date and time that the order was placed, in UTC ISO 8601 format. For example, \&quot;2011-12-03T10:15:30Z\&quot;. | [optional] 
**ShippingLocation** | Pointer to [**ShippingLocation**](ShippingLocation.md) |  | [optional] 
**Status** | Pointer to **string** | The processing status of the order. | [optional] 
**TrackingUrl** | Pointer to **string** | The URL, provided by the carrier company, where the shipment can be tracked. | [optional] 

## Methods

### NewTerminalOrder

`func NewTerminalOrder() *TerminalOrder`

NewTerminalOrder instantiates a new TerminalOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminalOrderWithDefaults

`func NewTerminalOrderWithDefaults() *TerminalOrder`

NewTerminalOrderWithDefaults instantiates a new TerminalOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingEntity

`func (o *TerminalOrder) GetBillingEntity() BillingEntity`

GetBillingEntity returns the BillingEntity field if non-nil, zero value otherwise.

### GetBillingEntityOk

`func (o *TerminalOrder) GetBillingEntityOk() (*BillingEntity, bool)`

GetBillingEntityOk returns a tuple with the BillingEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEntity

`func (o *TerminalOrder) SetBillingEntity(v BillingEntity)`

SetBillingEntity sets BillingEntity field to given value.

### HasBillingEntity

`func (o *TerminalOrder) HasBillingEntity() bool`

HasBillingEntity returns a boolean if a field has been set.

### GetCustomerOrderReference

`func (o *TerminalOrder) GetCustomerOrderReference() string`

GetCustomerOrderReference returns the CustomerOrderReference field if non-nil, zero value otherwise.

### GetCustomerOrderReferenceOk

`func (o *TerminalOrder) GetCustomerOrderReferenceOk() (*string, bool)`

GetCustomerOrderReferenceOk returns a tuple with the CustomerOrderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerOrderReference

`func (o *TerminalOrder) SetCustomerOrderReference(v string)`

SetCustomerOrderReference sets CustomerOrderReference field to given value.

### HasCustomerOrderReference

`func (o *TerminalOrder) HasCustomerOrderReference() bool`

HasCustomerOrderReference returns a boolean if a field has been set.

### GetId

`func (o *TerminalOrder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TerminalOrder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TerminalOrder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TerminalOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *TerminalOrder) GetItems() []OrderItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TerminalOrder) GetItemsOk() (*[]OrderItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TerminalOrder) SetItems(v []OrderItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *TerminalOrder) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOrderDate

`func (o *TerminalOrder) GetOrderDate() string`

GetOrderDate returns the OrderDate field if non-nil, zero value otherwise.

### GetOrderDateOk

`func (o *TerminalOrder) GetOrderDateOk() (*string, bool)`

GetOrderDateOk returns a tuple with the OrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDate

`func (o *TerminalOrder) SetOrderDate(v string)`

SetOrderDate sets OrderDate field to given value.

### HasOrderDate

`func (o *TerminalOrder) HasOrderDate() bool`

HasOrderDate returns a boolean if a field has been set.

### GetShippingLocation

`func (o *TerminalOrder) GetShippingLocation() ShippingLocation`

GetShippingLocation returns the ShippingLocation field if non-nil, zero value otherwise.

### GetShippingLocationOk

`func (o *TerminalOrder) GetShippingLocationOk() (*ShippingLocation, bool)`

GetShippingLocationOk returns a tuple with the ShippingLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLocation

`func (o *TerminalOrder) SetShippingLocation(v ShippingLocation)`

SetShippingLocation sets ShippingLocation field to given value.

### HasShippingLocation

`func (o *TerminalOrder) HasShippingLocation() bool`

HasShippingLocation returns a boolean if a field has been set.

### GetStatus

`func (o *TerminalOrder) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TerminalOrder) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TerminalOrder) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TerminalOrder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTrackingUrl

`func (o *TerminalOrder) GetTrackingUrl() string`

GetTrackingUrl returns the TrackingUrl field if non-nil, zero value otherwise.

### GetTrackingUrlOk

`func (o *TerminalOrder) GetTrackingUrlOk() (*string, bool)`

GetTrackingUrlOk returns a tuple with the TrackingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingUrl

`func (o *TerminalOrder) SetTrackingUrl(v string)`

SetTrackingUrl sets TrackingUrl field to given value.

### HasTrackingUrl

`func (o *TerminalOrder) HasTrackingUrl() bool`

HasTrackingUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


