# PaymentReversalResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantAccount** | **string** | The merchant account that is used to process the payment. | 
**PaymentPspReference** | **string** | The [&#x60;pspReference&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment to reverse.  | 
**PspReference** | **string** | Adyen&#39;s 16-character reference associated with the reversal request. | 
**Reference** | Pointer to **string** | Your reference for the reversal request. | [optional] 
**Status** | **string** | The status of your request. This will always have the value **received**. | 

## Methods

### NewPaymentReversalResource

`func NewPaymentReversalResource(merchantAccount string, paymentPspReference string, pspReference string, status string, ) *PaymentReversalResource`

NewPaymentReversalResource instantiates a new PaymentReversalResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentReversalResourceWithDefaults

`func NewPaymentReversalResourceWithDefaults() *PaymentReversalResource`

NewPaymentReversalResourceWithDefaults instantiates a new PaymentReversalResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantAccount

`func (o *PaymentReversalResource) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *PaymentReversalResource) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *PaymentReversalResource) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetPaymentPspReference

`func (o *PaymentReversalResource) GetPaymentPspReference() string`

GetPaymentPspReference returns the PaymentPspReference field if non-nil, zero value otherwise.

### GetPaymentPspReferenceOk

`func (o *PaymentReversalResource) GetPaymentPspReferenceOk() (*string, bool)`

GetPaymentPspReferenceOk returns a tuple with the PaymentPspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentPspReference

`func (o *PaymentReversalResource) SetPaymentPspReference(v string)`

SetPaymentPspReference sets PaymentPspReference field to given value.


### GetPspReference

`func (o *PaymentReversalResource) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *PaymentReversalResource) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *PaymentReversalResource) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.


### GetReference

`func (o *PaymentReversalResource) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaymentReversalResource) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaymentReversalResource) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PaymentReversalResource) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentReversalResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentReversalResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentReversalResource) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


