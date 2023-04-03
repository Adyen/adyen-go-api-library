# PaymentCancelResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantAccount** | **string** | The merchant account that is used to process the payment. | 
**PaymentPspReference** | **string** | The [&#x60;pspReference&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment to cancel.  | 
**PspReference** | **string** | Adyen&#39;s 16-character reference associated with the cancel request. | 
**Reference** | Pointer to **string** | Your reference for the cancel request. | [optional] 
**Status** | **string** | The status of your request. This will always have the value **received**. | 

## Methods

### NewPaymentCancelResource

`func NewPaymentCancelResource(merchantAccount string, paymentPspReference string, pspReference string, status string, ) *PaymentCancelResource`

NewPaymentCancelResource instantiates a new PaymentCancelResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentCancelResourceWithDefaults

`func NewPaymentCancelResourceWithDefaults() *PaymentCancelResource`

NewPaymentCancelResourceWithDefaults instantiates a new PaymentCancelResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantAccount

`func (o *PaymentCancelResource) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *PaymentCancelResource) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *PaymentCancelResource) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetPaymentPspReference

`func (o *PaymentCancelResource) GetPaymentPspReference() string`

GetPaymentPspReference returns the PaymentPspReference field if non-nil, zero value otherwise.

### GetPaymentPspReferenceOk

`func (o *PaymentCancelResource) GetPaymentPspReferenceOk() (*string, bool)`

GetPaymentPspReferenceOk returns a tuple with the PaymentPspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentPspReference

`func (o *PaymentCancelResource) SetPaymentPspReference(v string)`

SetPaymentPspReference sets PaymentPspReference field to given value.


### GetPspReference

`func (o *PaymentCancelResource) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *PaymentCancelResource) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *PaymentCancelResource) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.


### GetReference

`func (o *PaymentCancelResource) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaymentCancelResource) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaymentCancelResource) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PaymentCancelResource) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentCancelResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentCancelResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentCancelResource) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


