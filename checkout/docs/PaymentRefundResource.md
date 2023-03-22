# PaymentRefundResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | [**Amount**](Amount.md) |  | 
**LineItems** | Pointer to [**[]LineItem**](LineItem.md) | Price and product information of the refunded items, required for [partial refunds](https://docs.adyen.com/online-payments/refund#refund-a-payment). &gt; This field is required for partial refunds with 3x 4x Oney, Affirm, Afterpay, Clearpay, Klarna, Ratepay, Zip and Atome. | [optional] 
**MerchantAccount** | **string** | The merchant account that is used to process the payment. | 
**MerchantRefundReason** | Pointer to **string** | Your reason for the refund request. | [optional] 
**PaymentPspReference** | **string** | The [&#x60;pspReference&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment to refund.  | 
**PspReference** | **string** | Adyen&#39;s 16-character reference associated with the refund request. | 
**Reference** | Pointer to **string** | Your reference for the refund request. | [optional] 
**Splits** | Pointer to [**[]Split**](Split.md) | An array of objects specifying how the amount should be split between accounts when using Adyen for Platforms. For details, refer to [Providing split information](https://docs.adyen.com/marketplaces-and-platforms/processing-payments#providing-split-information). | [optional] 
**Status** | **string** | The status of your request. This will always have the value **received**. | 

## Methods

### NewPaymentRefundResource

`func NewPaymentRefundResource(amount Amount, merchantAccount string, paymentPspReference string, pspReference string, status string, ) *PaymentRefundResource`

NewPaymentRefundResource instantiates a new PaymentRefundResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRefundResourceWithDefaults

`func NewPaymentRefundResourceWithDefaults() *PaymentRefundResource`

NewPaymentRefundResourceWithDefaults instantiates a new PaymentRefundResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PaymentRefundResource) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentRefundResource) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentRefundResource) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetLineItems

`func (o *PaymentRefundResource) GetLineItems() []LineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *PaymentRefundResource) GetLineItemsOk() (*[]LineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *PaymentRefundResource) SetLineItems(v []LineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *PaymentRefundResource) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *PaymentRefundResource) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *PaymentRefundResource) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *PaymentRefundResource) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetMerchantRefundReason

`func (o *PaymentRefundResource) GetMerchantRefundReason() string`

GetMerchantRefundReason returns the MerchantRefundReason field if non-nil, zero value otherwise.

### GetMerchantRefundReasonOk

`func (o *PaymentRefundResource) GetMerchantRefundReasonOk() (*string, bool)`

GetMerchantRefundReasonOk returns a tuple with the MerchantRefundReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantRefundReason

`func (o *PaymentRefundResource) SetMerchantRefundReason(v string)`

SetMerchantRefundReason sets MerchantRefundReason field to given value.

### HasMerchantRefundReason

`func (o *PaymentRefundResource) HasMerchantRefundReason() bool`

HasMerchantRefundReason returns a boolean if a field has been set.

### GetPaymentPspReference

`func (o *PaymentRefundResource) GetPaymentPspReference() string`

GetPaymentPspReference returns the PaymentPspReference field if non-nil, zero value otherwise.

### GetPaymentPspReferenceOk

`func (o *PaymentRefundResource) GetPaymentPspReferenceOk() (*string, bool)`

GetPaymentPspReferenceOk returns a tuple with the PaymentPspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentPspReference

`func (o *PaymentRefundResource) SetPaymentPspReference(v string)`

SetPaymentPspReference sets PaymentPspReference field to given value.


### GetPspReference

`func (o *PaymentRefundResource) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *PaymentRefundResource) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *PaymentRefundResource) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.


### GetReference

`func (o *PaymentRefundResource) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaymentRefundResource) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaymentRefundResource) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PaymentRefundResource) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSplits

`func (o *PaymentRefundResource) GetSplits() []Split`

GetSplits returns the Splits field if non-nil, zero value otherwise.

### GetSplitsOk

`func (o *PaymentRefundResource) GetSplitsOk() (*[]Split, bool)`

GetSplitsOk returns a tuple with the Splits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplits

`func (o *PaymentRefundResource) SetSplits(v []Split)`

SetSplits sets Splits field to given value.

### HasSplits

`func (o *PaymentRefundResource) HasSplits() bool`

HasSplits returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentRefundResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentRefundResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentRefundResource) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


