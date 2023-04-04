# CreateStandalonePaymentCancelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantAccount** | **string** | The merchant account that is used to process the payment. | 
**PaymentReference** | **string** | The [&#x60;reference&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__reqParam_reference) of the payment that you want to cancel. | 
**Reference** | Pointer to **string** | Your reference for the cancel request. Maximum length: 80 characters. | [optional] 

## Methods

### NewCreateStandalonePaymentCancelRequest

`func NewCreateStandalonePaymentCancelRequest(merchantAccount string, paymentReference string, ) *CreateStandalonePaymentCancelRequest`

NewCreateStandalonePaymentCancelRequest instantiates a new CreateStandalonePaymentCancelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStandalonePaymentCancelRequestWithDefaults

`func NewCreateStandalonePaymentCancelRequestWithDefaults() *CreateStandalonePaymentCancelRequest`

NewCreateStandalonePaymentCancelRequestWithDefaults instantiates a new CreateStandalonePaymentCancelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantAccount

`func (o *CreateStandalonePaymentCancelRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *CreateStandalonePaymentCancelRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *CreateStandalonePaymentCancelRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetPaymentReference

`func (o *CreateStandalonePaymentCancelRequest) GetPaymentReference() string`

GetPaymentReference returns the PaymentReference field if non-nil, zero value otherwise.

### GetPaymentReferenceOk

`func (o *CreateStandalonePaymentCancelRequest) GetPaymentReferenceOk() (*string, bool)`

GetPaymentReferenceOk returns a tuple with the PaymentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentReference

`func (o *CreateStandalonePaymentCancelRequest) SetPaymentReference(v string)`

SetPaymentReference sets PaymentReference field to given value.


### GetReference

`func (o *CreateStandalonePaymentCancelRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CreateStandalonePaymentCancelRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CreateStandalonePaymentCancelRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *CreateStandalonePaymentCancelRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


