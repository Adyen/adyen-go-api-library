# StandalonePaymentCancelResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantAccount** | **string** | The merchant account that is used to process the payment. | 
**PaymentReference** | **string** | The [&#x60;reference&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__reqParam_reference) of the payment to cancel. | 
**PspReference** | **string** | Adyen&#39;s 16-character reference associated with the cancel request. | 
**Reference** | Pointer to **string** | Your reference for the cancel request. | [optional] 
**Status** | **string** | The status of your request. This will always have the value **received**. | 

## Methods

### NewStandalonePaymentCancelResource

`func NewStandalonePaymentCancelResource(merchantAccount string, paymentReference string, pspReference string, status string, ) *StandalonePaymentCancelResource`

NewStandalonePaymentCancelResource instantiates a new StandalonePaymentCancelResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandalonePaymentCancelResourceWithDefaults

`func NewStandalonePaymentCancelResourceWithDefaults() *StandalonePaymentCancelResource`

NewStandalonePaymentCancelResourceWithDefaults instantiates a new StandalonePaymentCancelResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantAccount

`func (o *StandalonePaymentCancelResource) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *StandalonePaymentCancelResource) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *StandalonePaymentCancelResource) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetPaymentReference

`func (o *StandalonePaymentCancelResource) GetPaymentReference() string`

GetPaymentReference returns the PaymentReference field if non-nil, zero value otherwise.

### GetPaymentReferenceOk

`func (o *StandalonePaymentCancelResource) GetPaymentReferenceOk() (*string, bool)`

GetPaymentReferenceOk returns a tuple with the PaymentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentReference

`func (o *StandalonePaymentCancelResource) SetPaymentReference(v string)`

SetPaymentReference sets PaymentReference field to given value.


### GetPspReference

`func (o *StandalonePaymentCancelResource) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *StandalonePaymentCancelResource) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *StandalonePaymentCancelResource) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.


### GetReference

`func (o *StandalonePaymentCancelResource) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *StandalonePaymentCancelResource) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *StandalonePaymentCancelResource) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *StandalonePaymentCancelResource) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetStatus

`func (o *StandalonePaymentCancelResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StandalonePaymentCancelResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StandalonePaymentCancelResource) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


