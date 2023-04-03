# CreatePaymentCancelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantAccount** | **string** | The merchant account that is used to process the payment. | 
**Reference** | Pointer to **string** | Your reference for the cancel request. Maximum length: 80 characters. | [optional] 

## Methods

### NewCreatePaymentCancelRequest

`func NewCreatePaymentCancelRequest(merchantAccount string, ) *CreatePaymentCancelRequest`

NewCreatePaymentCancelRequest instantiates a new CreatePaymentCancelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaymentCancelRequestWithDefaults

`func NewCreatePaymentCancelRequestWithDefaults() *CreatePaymentCancelRequest`

NewCreatePaymentCancelRequestWithDefaults instantiates a new CreatePaymentCancelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantAccount

`func (o *CreatePaymentCancelRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *CreatePaymentCancelRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *CreatePaymentCancelRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetReference

`func (o *CreatePaymentCancelRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CreatePaymentCancelRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CreatePaymentCancelRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *CreatePaymentCancelRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


