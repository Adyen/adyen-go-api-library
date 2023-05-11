# TerminalOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingEntityId** | Pointer to **string** | The identification of the billing entity to use for the order. | [optional] 
**CustomerOrderReference** | Pointer to **string** | The merchant-defined purchase order reference. | [optional] 
**Items** | Pointer to [**[]OrderItem**](OrderItem.md) | The products included in the order. | [optional] 
**ShippingLocationId** | Pointer to **string** | The identification of the shipping location to use for the order. | [optional] 
**TaxId** | Pointer to **string** | The tax number of the billing entity. | [optional] 

## Methods

### NewTerminalOrderRequest

`func NewTerminalOrderRequest() *TerminalOrderRequest`

NewTerminalOrderRequest instantiates a new TerminalOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminalOrderRequestWithDefaults

`func NewTerminalOrderRequestWithDefaults() *TerminalOrderRequest`

NewTerminalOrderRequestWithDefaults instantiates a new TerminalOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingEntityId

`func (o *TerminalOrderRequest) GetBillingEntityId() string`

GetBillingEntityId returns the BillingEntityId field if non-nil, zero value otherwise.

### GetBillingEntityIdOk

`func (o *TerminalOrderRequest) GetBillingEntityIdOk() (*string, bool)`

GetBillingEntityIdOk returns a tuple with the BillingEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEntityId

`func (o *TerminalOrderRequest) SetBillingEntityId(v string)`

SetBillingEntityId sets BillingEntityId field to given value.

### HasBillingEntityId

`func (o *TerminalOrderRequest) HasBillingEntityId() bool`

HasBillingEntityId returns a boolean if a field has been set.

### GetCustomerOrderReference

`func (o *TerminalOrderRequest) GetCustomerOrderReference() string`

GetCustomerOrderReference returns the CustomerOrderReference field if non-nil, zero value otherwise.

### GetCustomerOrderReferenceOk

`func (o *TerminalOrderRequest) GetCustomerOrderReferenceOk() (*string, bool)`

GetCustomerOrderReferenceOk returns a tuple with the CustomerOrderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerOrderReference

`func (o *TerminalOrderRequest) SetCustomerOrderReference(v string)`

SetCustomerOrderReference sets CustomerOrderReference field to given value.

### HasCustomerOrderReference

`func (o *TerminalOrderRequest) HasCustomerOrderReference() bool`

HasCustomerOrderReference returns a boolean if a field has been set.

### GetItems

`func (o *TerminalOrderRequest) GetItems() []OrderItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TerminalOrderRequest) GetItemsOk() (*[]OrderItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TerminalOrderRequest) SetItems(v []OrderItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *TerminalOrderRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetShippingLocationId

`func (o *TerminalOrderRequest) GetShippingLocationId() string`

GetShippingLocationId returns the ShippingLocationId field if non-nil, zero value otherwise.

### GetShippingLocationIdOk

`func (o *TerminalOrderRequest) GetShippingLocationIdOk() (*string, bool)`

GetShippingLocationIdOk returns a tuple with the ShippingLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLocationId

`func (o *TerminalOrderRequest) SetShippingLocationId(v string)`

SetShippingLocationId sets ShippingLocationId field to given value.

### HasShippingLocationId

`func (o *TerminalOrderRequest) HasShippingLocationId() bool`

HasShippingLocationId returns a boolean if a field has been set.

### GetTaxId

`func (o *TerminalOrderRequest) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *TerminalOrderRequest) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *TerminalOrderRequest) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *TerminalOrderRequest) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


