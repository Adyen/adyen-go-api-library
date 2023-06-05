# CardOrderItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalancePlatform** | Pointer to **string** | The unique identifier of the balance platform. | [optional] 
**Card** | Pointer to [**CardOrderItemDeliveryStatus**](CardOrderItemDeliveryStatus.md) |  | [optional] 
**CardOrderItemId** | Pointer to **string** | The unique identifier of the card order item. | [optional] 
**CreationDate** | Pointer to **time.Time** | The date and time when the event was triggered, in ISO 8601 extended format. For example, **2020-12-18T10:15:30+01:00**. | [optional] 
**Id** | Pointer to **string** | The ID of the resource. | [optional] 
**PaymentInstrumentId** | Pointer to **string** | The unique identifier of the payment instrument related to the card order item. | [optional] 
**Pin** | Pointer to [**CardOrderItemDeliveryStatus**](CardOrderItemDeliveryStatus.md) |  | [optional] 
**ShippingMethod** | Pointer to **string** | Shipping method used to deliver the card or the PIN. | [optional] 

## Methods

### NewCardOrderItem

`func NewCardOrderItem() *CardOrderItem`

NewCardOrderItem instantiates a new CardOrderItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardOrderItemWithDefaults

`func NewCardOrderItemWithDefaults() *CardOrderItem`

NewCardOrderItemWithDefaults instantiates a new CardOrderItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalancePlatform

`func (o *CardOrderItem) GetBalancePlatform() string`

GetBalancePlatform returns the BalancePlatform field if non-nil, zero value otherwise.

### GetBalancePlatformOk

`func (o *CardOrderItem) GetBalancePlatformOk() (*string, bool)`

GetBalancePlatformOk returns a tuple with the BalancePlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancePlatform

`func (o *CardOrderItem) SetBalancePlatform(v string)`

SetBalancePlatform sets BalancePlatform field to given value.

### HasBalancePlatform

`func (o *CardOrderItem) HasBalancePlatform() bool`

HasBalancePlatform returns a boolean if a field has been set.

### GetCard

`func (o *CardOrderItem) GetCard() CardOrderItemDeliveryStatus`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *CardOrderItem) GetCardOk() (*CardOrderItemDeliveryStatus, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *CardOrderItem) SetCard(v CardOrderItemDeliveryStatus)`

SetCard sets Card field to given value.

### HasCard

`func (o *CardOrderItem) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetCardOrderItemId

`func (o *CardOrderItem) GetCardOrderItemId() string`

GetCardOrderItemId returns the CardOrderItemId field if non-nil, zero value otherwise.

### GetCardOrderItemIdOk

`func (o *CardOrderItem) GetCardOrderItemIdOk() (*string, bool)`

GetCardOrderItemIdOk returns a tuple with the CardOrderItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardOrderItemId

`func (o *CardOrderItem) SetCardOrderItemId(v string)`

SetCardOrderItemId sets CardOrderItemId field to given value.

### HasCardOrderItemId

`func (o *CardOrderItem) HasCardOrderItemId() bool`

HasCardOrderItemId returns a boolean if a field has been set.

### GetCreationDate

`func (o *CardOrderItem) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *CardOrderItem) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *CardOrderItem) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *CardOrderItem) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetId

`func (o *CardOrderItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardOrderItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardOrderItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CardOrderItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPaymentInstrumentId

`func (o *CardOrderItem) GetPaymentInstrumentId() string`

GetPaymentInstrumentId returns the PaymentInstrumentId field if non-nil, zero value otherwise.

### GetPaymentInstrumentIdOk

`func (o *CardOrderItem) GetPaymentInstrumentIdOk() (*string, bool)`

GetPaymentInstrumentIdOk returns a tuple with the PaymentInstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrumentId

`func (o *CardOrderItem) SetPaymentInstrumentId(v string)`

SetPaymentInstrumentId sets PaymentInstrumentId field to given value.

### HasPaymentInstrumentId

`func (o *CardOrderItem) HasPaymentInstrumentId() bool`

HasPaymentInstrumentId returns a boolean if a field has been set.

### GetPin

`func (o *CardOrderItem) GetPin() CardOrderItemDeliveryStatus`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *CardOrderItem) GetPinOk() (*CardOrderItemDeliveryStatus, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *CardOrderItem) SetPin(v CardOrderItemDeliveryStatus)`

SetPin sets Pin field to given value.

### HasPin

`func (o *CardOrderItem) HasPin() bool`

HasPin returns a boolean if a field has been set.

### GetShippingMethod

`func (o *CardOrderItem) GetShippingMethod() string`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *CardOrderItem) GetShippingMethodOk() (*string, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *CardOrderItem) SetShippingMethod(v string)`

SetShippingMethod sets ShippingMethod field to given value.

### HasShippingMethod

`func (o *CardOrderItem) HasShippingMethod() bool`

HasShippingMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


