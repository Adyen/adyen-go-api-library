# PaymentAmountUpdateResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | [**Amount**](Amount.md) |  | 
**IndustryUsage** | Pointer to **string** | The reason for the amount update. Possible values:  * **delayedCharge**  * **noShow**  * **installment** | [optional] 
**MerchantAccount** | **string** | The merchant account that is used to process the payment. | 
**PaymentPspReference** | **string** | The [&#x60;pspReference&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment to update.  | 
**PspReference** | **string** | Adyen&#39;s 16-character reference associated with the amount update request. | 
**Reference** | **string** | Your reference for the amount update request. Maximum length: 80 characters. | 
**Splits** | Pointer to [**[]Split**](Split.md) | An array of objects specifying how the amount should be split between accounts when using Adyen for Platforms. For details, refer to [Providing split information](https://docs.adyen.com/marketplaces-and-platforms/processing-payments#providing-split-information). | [optional] 
**Status** | **string** | The status of your request. This will always have the value **received**. | 

## Methods

### NewPaymentAmountUpdateResource

`func NewPaymentAmountUpdateResource(amount Amount, merchantAccount string, paymentPspReference string, pspReference string, reference string, status string, ) *PaymentAmountUpdateResource`

NewPaymentAmountUpdateResource instantiates a new PaymentAmountUpdateResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentAmountUpdateResourceWithDefaults

`func NewPaymentAmountUpdateResourceWithDefaults() *PaymentAmountUpdateResource`

NewPaymentAmountUpdateResourceWithDefaults instantiates a new PaymentAmountUpdateResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PaymentAmountUpdateResource) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentAmountUpdateResource) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentAmountUpdateResource) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetIndustryUsage

`func (o *PaymentAmountUpdateResource) GetIndustryUsage() string`

GetIndustryUsage returns the IndustryUsage field if non-nil, zero value otherwise.

### GetIndustryUsageOk

`func (o *PaymentAmountUpdateResource) GetIndustryUsageOk() (*string, bool)`

GetIndustryUsageOk returns a tuple with the IndustryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryUsage

`func (o *PaymentAmountUpdateResource) SetIndustryUsage(v string)`

SetIndustryUsage sets IndustryUsage field to given value.

### HasIndustryUsage

`func (o *PaymentAmountUpdateResource) HasIndustryUsage() bool`

HasIndustryUsage returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *PaymentAmountUpdateResource) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *PaymentAmountUpdateResource) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *PaymentAmountUpdateResource) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetPaymentPspReference

`func (o *PaymentAmountUpdateResource) GetPaymentPspReference() string`

GetPaymentPspReference returns the PaymentPspReference field if non-nil, zero value otherwise.

### GetPaymentPspReferenceOk

`func (o *PaymentAmountUpdateResource) GetPaymentPspReferenceOk() (*string, bool)`

GetPaymentPspReferenceOk returns a tuple with the PaymentPspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentPspReference

`func (o *PaymentAmountUpdateResource) SetPaymentPspReference(v string)`

SetPaymentPspReference sets PaymentPspReference field to given value.


### GetPspReference

`func (o *PaymentAmountUpdateResource) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *PaymentAmountUpdateResource) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *PaymentAmountUpdateResource) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.


### GetReference

`func (o *PaymentAmountUpdateResource) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaymentAmountUpdateResource) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaymentAmountUpdateResource) SetReference(v string)`

SetReference sets Reference field to given value.


### GetSplits

`func (o *PaymentAmountUpdateResource) GetSplits() []Split`

GetSplits returns the Splits field if non-nil, zero value otherwise.

### GetSplitsOk

`func (o *PaymentAmountUpdateResource) GetSplitsOk() (*[]Split, bool)`

GetSplitsOk returns a tuple with the Splits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplits

`func (o *PaymentAmountUpdateResource) SetSplits(v []Split)`

SetSplits sets Splits field to given value.

### HasSplits

`func (o *PaymentAmountUpdateResource) HasSplits() bool`

HasSplits returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentAmountUpdateResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentAmountUpdateResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentAmountUpdateResource) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


