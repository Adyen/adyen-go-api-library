# CheckoutCreateOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | [**Amount**](Amount.md) |  | 
**ExpiresAt** | Pointer to **string** | The date that order expires; e.g. 2019-03-23T12:25:28Z. If not provided, the default expiry duration is 1 day. | [optional] 
**MerchantAccount** | **string** | The merchant account identifier, with which you want to process the order. | 
**Reference** | **string** | A custom reference identifying the order. | 

## Methods

### NewCheckoutCreateOrderRequest

`func NewCheckoutCreateOrderRequest(amount Amount, merchantAccount string, reference string, ) *CheckoutCreateOrderRequest`

NewCheckoutCreateOrderRequest instantiates a new CheckoutCreateOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutCreateOrderRequestWithDefaults

`func NewCheckoutCreateOrderRequestWithDefaults() *CheckoutCreateOrderRequest`

NewCheckoutCreateOrderRequestWithDefaults instantiates a new CheckoutCreateOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CheckoutCreateOrderRequest) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CheckoutCreateOrderRequest) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CheckoutCreateOrderRequest) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetExpiresAt

`func (o *CheckoutCreateOrderRequest) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CheckoutCreateOrderRequest) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CheckoutCreateOrderRequest) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CheckoutCreateOrderRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *CheckoutCreateOrderRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *CheckoutCreateOrderRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *CheckoutCreateOrderRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetReference

`func (o *CheckoutCreateOrderRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CheckoutCreateOrderRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CheckoutCreateOrderRequest) SetReference(v string)`

SetReference sets Reference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


