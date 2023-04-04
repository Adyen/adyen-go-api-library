# CreatePaymentRefundRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | [**Amount**](Amount.md) |  | 
**LineItems** | Pointer to [**[]LineItem**](LineItem.md) | Price and product information of the refunded items, required for [partial refunds](https://docs.adyen.com/online-payments/refund#refund-a-payment). &gt; This field is required for partial refunds with 3x 4x Oney, Affirm, Afterpay, Clearpay, Klarna, Ratepay, Zip and Atome. | [optional] 
**MerchantAccount** | **string** | The merchant account that is used to process the payment. | 
**MerchantRefundReason** | Pointer to **string** | Your reason for the refund request | [optional] 
**Reference** | Pointer to **string** | Your reference for the refund request. Maximum length: 80 characters. | [optional] 
**Splits** | Pointer to [**[]Split**](Split.md) | An array of objects specifying how the amount should be split between accounts when using Adyen for Platforms. For details, refer to [Providing split information](https://docs.adyen.com/marketplaces-and-platforms/processing-payments#providing-split-information). | [optional] 

## Methods

### NewCreatePaymentRefundRequest

`func NewCreatePaymentRefundRequest(amount Amount, merchantAccount string, ) *CreatePaymentRefundRequest`

NewCreatePaymentRefundRequest instantiates a new CreatePaymentRefundRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaymentRefundRequestWithDefaults

`func NewCreatePaymentRefundRequestWithDefaults() *CreatePaymentRefundRequest`

NewCreatePaymentRefundRequestWithDefaults instantiates a new CreatePaymentRefundRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CreatePaymentRefundRequest) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreatePaymentRefundRequest) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreatePaymentRefundRequest) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetLineItems

`func (o *CreatePaymentRefundRequest) GetLineItems() []LineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *CreatePaymentRefundRequest) GetLineItemsOk() (*[]LineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *CreatePaymentRefundRequest) SetLineItems(v []LineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *CreatePaymentRefundRequest) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *CreatePaymentRefundRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *CreatePaymentRefundRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *CreatePaymentRefundRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetMerchantRefundReason

`func (o *CreatePaymentRefundRequest) GetMerchantRefundReason() string`

GetMerchantRefundReason returns the MerchantRefundReason field if non-nil, zero value otherwise.

### GetMerchantRefundReasonOk

`func (o *CreatePaymentRefundRequest) GetMerchantRefundReasonOk() (*string, bool)`

GetMerchantRefundReasonOk returns a tuple with the MerchantRefundReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantRefundReason

`func (o *CreatePaymentRefundRequest) SetMerchantRefundReason(v string)`

SetMerchantRefundReason sets MerchantRefundReason field to given value.

### HasMerchantRefundReason

`func (o *CreatePaymentRefundRequest) HasMerchantRefundReason() bool`

HasMerchantRefundReason returns a boolean if a field has been set.

### GetReference

`func (o *CreatePaymentRefundRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CreatePaymentRefundRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CreatePaymentRefundRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *CreatePaymentRefundRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSplits

`func (o *CreatePaymentRefundRequest) GetSplits() []Split`

GetSplits returns the Splits field if non-nil, zero value otherwise.

### GetSplitsOk

`func (o *CreatePaymentRefundRequest) GetSplitsOk() (*[]Split, bool)`

GetSplitsOk returns a tuple with the Splits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplits

`func (o *CreatePaymentRefundRequest) SetSplits(v []Split)`

SetSplits sets Splits field to given value.

### HasSplits

`func (o *CreatePaymentRefundRequest) HasSplits() bool`

HasSplits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


