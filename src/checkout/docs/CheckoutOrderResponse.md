# CheckoutOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**ExpiresAt** | Pointer to **string** | The expiry date for the order. | [optional] 
**OrderData** | Pointer to **string** | The encrypted order data. | [optional] 
**PspReference** | **string** | The &#x60;pspReference&#x60; that belongs to the order. | 
**Reference** | Pointer to **string** | The merchant reference for the order. | [optional] 
**RemainingAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 

## Methods

### NewCheckoutOrderResponse

`func NewCheckoutOrderResponse(pspReference string, ) *CheckoutOrderResponse`

NewCheckoutOrderResponse instantiates a new CheckoutOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutOrderResponseWithDefaults

`func NewCheckoutOrderResponseWithDefaults() *CheckoutOrderResponse`

NewCheckoutOrderResponseWithDefaults instantiates a new CheckoutOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CheckoutOrderResponse) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CheckoutOrderResponse) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CheckoutOrderResponse) SetAmount(v Amount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CheckoutOrderResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CheckoutOrderResponse) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CheckoutOrderResponse) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CheckoutOrderResponse) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CheckoutOrderResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetOrderData

`func (o *CheckoutOrderResponse) GetOrderData() string`

GetOrderData returns the OrderData field if non-nil, zero value otherwise.

### GetOrderDataOk

`func (o *CheckoutOrderResponse) GetOrderDataOk() (*string, bool)`

GetOrderDataOk returns a tuple with the OrderData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderData

`func (o *CheckoutOrderResponse) SetOrderData(v string)`

SetOrderData sets OrderData field to given value.

### HasOrderData

`func (o *CheckoutOrderResponse) HasOrderData() bool`

HasOrderData returns a boolean if a field has been set.

### GetPspReference

`func (o *CheckoutOrderResponse) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *CheckoutOrderResponse) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *CheckoutOrderResponse) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.


### GetReference

`func (o *CheckoutOrderResponse) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CheckoutOrderResponse) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CheckoutOrderResponse) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *CheckoutOrderResponse) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetRemainingAmount

`func (o *CheckoutOrderResponse) GetRemainingAmount() Amount`

GetRemainingAmount returns the RemainingAmount field if non-nil, zero value otherwise.

### GetRemainingAmountOk

`func (o *CheckoutOrderResponse) GetRemainingAmountOk() (*Amount, bool)`

GetRemainingAmountOk returns a tuple with the RemainingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingAmount

`func (o *CheckoutOrderResponse) SetRemainingAmount(v Amount)`

SetRemainingAmount sets RemainingAmount field to given value.

### HasRemainingAmount

`func (o *CheckoutOrderResponse) HasRemainingAmount() bool`

HasRemainingAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


