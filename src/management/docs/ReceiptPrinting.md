# ReceiptPrinting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantApproved** | Pointer to **bool** | Print a merchant receipt when the payment is approved. | [optional] 
**MerchantCancelled** | Pointer to **bool** | Print a merchant receipt when the transaction is cancelled. | [optional] 
**MerchantCaptureApproved** | Pointer to **bool** | Print a merchant receipt when capturing the payment is approved. | [optional] 
**MerchantCaptureRefused** | Pointer to **bool** | Print a merchant receipt when capturing the payment is refused. | [optional] 
**MerchantRefundApproved** | Pointer to **bool** | Print a merchant receipt when the refund is approved. | [optional] 
**MerchantRefundRefused** | Pointer to **bool** | Print a merchant receipt when the refund is refused. | [optional] 
**MerchantRefused** | Pointer to **bool** | Print a merchant receipt when the payment is refused. | [optional] 
**MerchantVoid** | Pointer to **bool** | Print a merchant receipt when a previous transaction is voided. | [optional] 
**ShopperApproved** | Pointer to **bool** | Print a shopper receipt when the payment is approved. | [optional] 
**ShopperCancelled** | Pointer to **bool** | Print a shopper receipt when the transaction is cancelled. | [optional] 
**ShopperCaptureApproved** | Pointer to **bool** | Print a shopper receipt when capturing the payment is approved. | [optional] 
**ShopperCaptureRefused** | Pointer to **bool** | Print a shopper receipt when capturing the payment is refused. | [optional] 
**ShopperRefundApproved** | Pointer to **bool** | Print a shopper receipt when the refund is approved. | [optional] 
**ShopperRefundRefused** | Pointer to **bool** | Print a shopper receipt when the refund is refused. | [optional] 
**ShopperRefused** | Pointer to **bool** | Print a shopper receipt when the payment is refused. | [optional] 
**ShopperVoid** | Pointer to **bool** | Print a shopper receipt when a previous transaction is voided. | [optional] 

## Methods

### NewReceiptPrinting

`func NewReceiptPrinting() *ReceiptPrinting`

NewReceiptPrinting instantiates a new ReceiptPrinting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiptPrintingWithDefaults

`func NewReceiptPrintingWithDefaults() *ReceiptPrinting`

NewReceiptPrintingWithDefaults instantiates a new ReceiptPrinting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantApproved

`func (o *ReceiptPrinting) GetMerchantApproved() bool`

GetMerchantApproved returns the MerchantApproved field if non-nil, zero value otherwise.

### GetMerchantApprovedOk

`func (o *ReceiptPrinting) GetMerchantApprovedOk() (*bool, bool)`

GetMerchantApprovedOk returns a tuple with the MerchantApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantApproved

`func (o *ReceiptPrinting) SetMerchantApproved(v bool)`

SetMerchantApproved sets MerchantApproved field to given value.

### HasMerchantApproved

`func (o *ReceiptPrinting) HasMerchantApproved() bool`

HasMerchantApproved returns a boolean if a field has been set.

### GetMerchantCancelled

`func (o *ReceiptPrinting) GetMerchantCancelled() bool`

GetMerchantCancelled returns the MerchantCancelled field if non-nil, zero value otherwise.

### GetMerchantCancelledOk

`func (o *ReceiptPrinting) GetMerchantCancelledOk() (*bool, bool)`

GetMerchantCancelledOk returns a tuple with the MerchantCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCancelled

`func (o *ReceiptPrinting) SetMerchantCancelled(v bool)`

SetMerchantCancelled sets MerchantCancelled field to given value.

### HasMerchantCancelled

`func (o *ReceiptPrinting) HasMerchantCancelled() bool`

HasMerchantCancelled returns a boolean if a field has been set.

### GetMerchantCaptureApproved

`func (o *ReceiptPrinting) GetMerchantCaptureApproved() bool`

GetMerchantCaptureApproved returns the MerchantCaptureApproved field if non-nil, zero value otherwise.

### GetMerchantCaptureApprovedOk

`func (o *ReceiptPrinting) GetMerchantCaptureApprovedOk() (*bool, bool)`

GetMerchantCaptureApprovedOk returns a tuple with the MerchantCaptureApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCaptureApproved

`func (o *ReceiptPrinting) SetMerchantCaptureApproved(v bool)`

SetMerchantCaptureApproved sets MerchantCaptureApproved field to given value.

### HasMerchantCaptureApproved

`func (o *ReceiptPrinting) HasMerchantCaptureApproved() bool`

HasMerchantCaptureApproved returns a boolean if a field has been set.

### GetMerchantCaptureRefused

`func (o *ReceiptPrinting) GetMerchantCaptureRefused() bool`

GetMerchantCaptureRefused returns the MerchantCaptureRefused field if non-nil, zero value otherwise.

### GetMerchantCaptureRefusedOk

`func (o *ReceiptPrinting) GetMerchantCaptureRefusedOk() (*bool, bool)`

GetMerchantCaptureRefusedOk returns a tuple with the MerchantCaptureRefused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCaptureRefused

`func (o *ReceiptPrinting) SetMerchantCaptureRefused(v bool)`

SetMerchantCaptureRefused sets MerchantCaptureRefused field to given value.

### HasMerchantCaptureRefused

`func (o *ReceiptPrinting) HasMerchantCaptureRefused() bool`

HasMerchantCaptureRefused returns a boolean if a field has been set.

### GetMerchantRefundApproved

`func (o *ReceiptPrinting) GetMerchantRefundApproved() bool`

GetMerchantRefundApproved returns the MerchantRefundApproved field if non-nil, zero value otherwise.

### GetMerchantRefundApprovedOk

`func (o *ReceiptPrinting) GetMerchantRefundApprovedOk() (*bool, bool)`

GetMerchantRefundApprovedOk returns a tuple with the MerchantRefundApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantRefundApproved

`func (o *ReceiptPrinting) SetMerchantRefundApproved(v bool)`

SetMerchantRefundApproved sets MerchantRefundApproved field to given value.

### HasMerchantRefundApproved

`func (o *ReceiptPrinting) HasMerchantRefundApproved() bool`

HasMerchantRefundApproved returns a boolean if a field has been set.

### GetMerchantRefundRefused

`func (o *ReceiptPrinting) GetMerchantRefundRefused() bool`

GetMerchantRefundRefused returns the MerchantRefundRefused field if non-nil, zero value otherwise.

### GetMerchantRefundRefusedOk

`func (o *ReceiptPrinting) GetMerchantRefundRefusedOk() (*bool, bool)`

GetMerchantRefundRefusedOk returns a tuple with the MerchantRefundRefused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantRefundRefused

`func (o *ReceiptPrinting) SetMerchantRefundRefused(v bool)`

SetMerchantRefundRefused sets MerchantRefundRefused field to given value.

### HasMerchantRefundRefused

`func (o *ReceiptPrinting) HasMerchantRefundRefused() bool`

HasMerchantRefundRefused returns a boolean if a field has been set.

### GetMerchantRefused

`func (o *ReceiptPrinting) GetMerchantRefused() bool`

GetMerchantRefused returns the MerchantRefused field if non-nil, zero value otherwise.

### GetMerchantRefusedOk

`func (o *ReceiptPrinting) GetMerchantRefusedOk() (*bool, bool)`

GetMerchantRefusedOk returns a tuple with the MerchantRefused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantRefused

`func (o *ReceiptPrinting) SetMerchantRefused(v bool)`

SetMerchantRefused sets MerchantRefused field to given value.

### HasMerchantRefused

`func (o *ReceiptPrinting) HasMerchantRefused() bool`

HasMerchantRefused returns a boolean if a field has been set.

### GetMerchantVoid

`func (o *ReceiptPrinting) GetMerchantVoid() bool`

GetMerchantVoid returns the MerchantVoid field if non-nil, zero value otherwise.

### GetMerchantVoidOk

`func (o *ReceiptPrinting) GetMerchantVoidOk() (*bool, bool)`

GetMerchantVoidOk returns a tuple with the MerchantVoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantVoid

`func (o *ReceiptPrinting) SetMerchantVoid(v bool)`

SetMerchantVoid sets MerchantVoid field to given value.

### HasMerchantVoid

`func (o *ReceiptPrinting) HasMerchantVoid() bool`

HasMerchantVoid returns a boolean if a field has been set.

### GetShopperApproved

`func (o *ReceiptPrinting) GetShopperApproved() bool`

GetShopperApproved returns the ShopperApproved field if non-nil, zero value otherwise.

### GetShopperApprovedOk

`func (o *ReceiptPrinting) GetShopperApprovedOk() (*bool, bool)`

GetShopperApprovedOk returns a tuple with the ShopperApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperApproved

`func (o *ReceiptPrinting) SetShopperApproved(v bool)`

SetShopperApproved sets ShopperApproved field to given value.

### HasShopperApproved

`func (o *ReceiptPrinting) HasShopperApproved() bool`

HasShopperApproved returns a boolean if a field has been set.

### GetShopperCancelled

`func (o *ReceiptPrinting) GetShopperCancelled() bool`

GetShopperCancelled returns the ShopperCancelled field if non-nil, zero value otherwise.

### GetShopperCancelledOk

`func (o *ReceiptPrinting) GetShopperCancelledOk() (*bool, bool)`

GetShopperCancelledOk returns a tuple with the ShopperCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperCancelled

`func (o *ReceiptPrinting) SetShopperCancelled(v bool)`

SetShopperCancelled sets ShopperCancelled field to given value.

### HasShopperCancelled

`func (o *ReceiptPrinting) HasShopperCancelled() bool`

HasShopperCancelled returns a boolean if a field has been set.

### GetShopperCaptureApproved

`func (o *ReceiptPrinting) GetShopperCaptureApproved() bool`

GetShopperCaptureApproved returns the ShopperCaptureApproved field if non-nil, zero value otherwise.

### GetShopperCaptureApprovedOk

`func (o *ReceiptPrinting) GetShopperCaptureApprovedOk() (*bool, bool)`

GetShopperCaptureApprovedOk returns a tuple with the ShopperCaptureApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperCaptureApproved

`func (o *ReceiptPrinting) SetShopperCaptureApproved(v bool)`

SetShopperCaptureApproved sets ShopperCaptureApproved field to given value.

### HasShopperCaptureApproved

`func (o *ReceiptPrinting) HasShopperCaptureApproved() bool`

HasShopperCaptureApproved returns a boolean if a field has been set.

### GetShopperCaptureRefused

`func (o *ReceiptPrinting) GetShopperCaptureRefused() bool`

GetShopperCaptureRefused returns the ShopperCaptureRefused field if non-nil, zero value otherwise.

### GetShopperCaptureRefusedOk

`func (o *ReceiptPrinting) GetShopperCaptureRefusedOk() (*bool, bool)`

GetShopperCaptureRefusedOk returns a tuple with the ShopperCaptureRefused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperCaptureRefused

`func (o *ReceiptPrinting) SetShopperCaptureRefused(v bool)`

SetShopperCaptureRefused sets ShopperCaptureRefused field to given value.

### HasShopperCaptureRefused

`func (o *ReceiptPrinting) HasShopperCaptureRefused() bool`

HasShopperCaptureRefused returns a boolean if a field has been set.

### GetShopperRefundApproved

`func (o *ReceiptPrinting) GetShopperRefundApproved() bool`

GetShopperRefundApproved returns the ShopperRefundApproved field if non-nil, zero value otherwise.

### GetShopperRefundApprovedOk

`func (o *ReceiptPrinting) GetShopperRefundApprovedOk() (*bool, bool)`

GetShopperRefundApprovedOk returns a tuple with the ShopperRefundApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperRefundApproved

`func (o *ReceiptPrinting) SetShopperRefundApproved(v bool)`

SetShopperRefundApproved sets ShopperRefundApproved field to given value.

### HasShopperRefundApproved

`func (o *ReceiptPrinting) HasShopperRefundApproved() bool`

HasShopperRefundApproved returns a boolean if a field has been set.

### GetShopperRefundRefused

`func (o *ReceiptPrinting) GetShopperRefundRefused() bool`

GetShopperRefundRefused returns the ShopperRefundRefused field if non-nil, zero value otherwise.

### GetShopperRefundRefusedOk

`func (o *ReceiptPrinting) GetShopperRefundRefusedOk() (*bool, bool)`

GetShopperRefundRefusedOk returns a tuple with the ShopperRefundRefused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperRefundRefused

`func (o *ReceiptPrinting) SetShopperRefundRefused(v bool)`

SetShopperRefundRefused sets ShopperRefundRefused field to given value.

### HasShopperRefundRefused

`func (o *ReceiptPrinting) HasShopperRefundRefused() bool`

HasShopperRefundRefused returns a boolean if a field has been set.

### GetShopperRefused

`func (o *ReceiptPrinting) GetShopperRefused() bool`

GetShopperRefused returns the ShopperRefused field if non-nil, zero value otherwise.

### GetShopperRefusedOk

`func (o *ReceiptPrinting) GetShopperRefusedOk() (*bool, bool)`

GetShopperRefusedOk returns a tuple with the ShopperRefused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperRefused

`func (o *ReceiptPrinting) SetShopperRefused(v bool)`

SetShopperRefused sets ShopperRefused field to given value.

### HasShopperRefused

`func (o *ReceiptPrinting) HasShopperRefused() bool`

HasShopperRefused returns a boolean if a field has been set.

### GetShopperVoid

`func (o *ReceiptPrinting) GetShopperVoid() bool`

GetShopperVoid returns the ShopperVoid field if non-nil, zero value otherwise.

### GetShopperVoidOk

`func (o *ReceiptPrinting) GetShopperVoidOk() (*bool, bool)`

GetShopperVoidOk returns a tuple with the ShopperVoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperVoid

`func (o *ReceiptPrinting) SetShopperVoid(v bool)`

SetShopperVoid sets ShopperVoid field to given value.

### HasShopperVoid

`func (o *ReceiptPrinting) HasShopperVoid() bool`

HasShopperVoid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


