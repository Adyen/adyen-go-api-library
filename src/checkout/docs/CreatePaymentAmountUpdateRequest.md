# CreatePaymentAmountUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | [**Amount**](Amount.md) |  | 
**IndustryUsage** | Pointer to **string** | The reason for the amount update. Possible values:  * **delayedCharge**  * **noShow**  * **installment** | [optional] 
**MerchantAccount** | **string** | The merchant account that is used to process the payment. | 
**Reference** | Pointer to **string** | Your reference for the amount update request. Maximum length: 80 characters. | [optional] 
**Splits** | Pointer to [**[]Split**](Split.md) | An array of objects specifying how the amount should be split between accounts when using Adyen for Platforms. For details, refer to [Providing split information](https://docs.adyen.com/marketplaces-and-platforms/processing-payments#providing-split-information). | [optional] 

## Methods

### NewCreatePaymentAmountUpdateRequest

`func NewCreatePaymentAmountUpdateRequest(amount Amount, merchantAccount string, ) *CreatePaymentAmountUpdateRequest`

NewCreatePaymentAmountUpdateRequest instantiates a new CreatePaymentAmountUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaymentAmountUpdateRequestWithDefaults

`func NewCreatePaymentAmountUpdateRequestWithDefaults() *CreatePaymentAmountUpdateRequest`

NewCreatePaymentAmountUpdateRequestWithDefaults instantiates a new CreatePaymentAmountUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CreatePaymentAmountUpdateRequest) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreatePaymentAmountUpdateRequest) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreatePaymentAmountUpdateRequest) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetIndustryUsage

`func (o *CreatePaymentAmountUpdateRequest) GetIndustryUsage() string`

GetIndustryUsage returns the IndustryUsage field if non-nil, zero value otherwise.

### GetIndustryUsageOk

`func (o *CreatePaymentAmountUpdateRequest) GetIndustryUsageOk() (*string, bool)`

GetIndustryUsageOk returns a tuple with the IndustryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryUsage

`func (o *CreatePaymentAmountUpdateRequest) SetIndustryUsage(v string)`

SetIndustryUsage sets IndustryUsage field to given value.

### HasIndustryUsage

`func (o *CreatePaymentAmountUpdateRequest) HasIndustryUsage() bool`

HasIndustryUsage returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *CreatePaymentAmountUpdateRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *CreatePaymentAmountUpdateRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *CreatePaymentAmountUpdateRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetReference

`func (o *CreatePaymentAmountUpdateRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CreatePaymentAmountUpdateRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CreatePaymentAmountUpdateRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *CreatePaymentAmountUpdateRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSplits

`func (o *CreatePaymentAmountUpdateRequest) GetSplits() []Split`

GetSplits returns the Splits field if non-nil, zero value otherwise.

### GetSplitsOk

`func (o *CreatePaymentAmountUpdateRequest) GetSplitsOk() (*[]Split, bool)`

GetSplitsOk returns a tuple with the Splits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplits

`func (o *CreatePaymentAmountUpdateRequest) SetSplits(v []Split)`

SetSplits sets Splits field to given value.

### HasSplits

`func (o *CreatePaymentAmountUpdateRequest) HasSplits() bool`

HasSplits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


