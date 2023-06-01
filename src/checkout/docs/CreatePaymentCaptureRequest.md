# CreatePaymentCaptureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | [**Amount**](Amount.md) |  | 
**LineItems** | Pointer to [**[]LineItem**](LineItem.md) | Price and product information of the refunded items, required for [partial refunds](https://docs.adyen.com/online-payments/refund#refund-a-payment). &gt; This field is required for partial refunds with 3x 4x Oney, Affirm, Afterpay, Atome, Clearpay, Klarna, Ratepay, Walley, and Zip. | [optional] 
**MerchantAccount** | **string** | The merchant account that is used to process the payment. | 
**Reference** | Pointer to **string** | Your reference for the capture request. Maximum length: 80 characters. | [optional] 
**Splits** | Pointer to [**[]Split**](Split.md) | An array of objects specifying how the amount should be split between accounts when using Adyen for Platforms. For details, refer to [Providing split information](https://docs.adyen.com/marketplaces-and-platforms/processing-payments#providing-split-information). | [optional] 

## Methods

### NewCreatePaymentCaptureRequest

`func NewCreatePaymentCaptureRequest(amount Amount, merchantAccount string, ) *CreatePaymentCaptureRequest`

NewCreatePaymentCaptureRequest instantiates a new CreatePaymentCaptureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaymentCaptureRequestWithDefaults

`func NewCreatePaymentCaptureRequestWithDefaults() *CreatePaymentCaptureRequest`

NewCreatePaymentCaptureRequestWithDefaults instantiates a new CreatePaymentCaptureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CreatePaymentCaptureRequest) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreatePaymentCaptureRequest) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreatePaymentCaptureRequest) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetLineItems

`func (o *CreatePaymentCaptureRequest) GetLineItems() []LineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *CreatePaymentCaptureRequest) GetLineItemsOk() (*[]LineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *CreatePaymentCaptureRequest) SetLineItems(v []LineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *CreatePaymentCaptureRequest) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *CreatePaymentCaptureRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *CreatePaymentCaptureRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *CreatePaymentCaptureRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetReference

`func (o *CreatePaymentCaptureRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CreatePaymentCaptureRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CreatePaymentCaptureRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *CreatePaymentCaptureRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSplits

`func (o *CreatePaymentCaptureRequest) GetSplits() []Split`

GetSplits returns the Splits field if non-nil, zero value otherwise.

### GetSplitsOk

`func (o *CreatePaymentCaptureRequest) GetSplitsOk() (*[]Split, bool)`

GetSplitsOk returns a tuple with the Splits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplits

`func (o *CreatePaymentCaptureRequest) SetSplits(v []Split)`

SetSplits sets Splits field to given value.

### HasSplits

`func (o *CreatePaymentCaptureRequest) HasSplits() bool`

HasSplits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


