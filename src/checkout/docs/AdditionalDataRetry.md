# AdditionalDataRetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetryChainAttemptNumber** | Pointer to **string** | The number of times the transaction (not order) has been retried between different payment service providers. For instance, the &#x60;chainAttemptNumber&#x60; set to 2 means that this transaction has been recently tried on another provider before being sent to Adyen.  &gt; If you submit &#x60;retry.chainAttemptNumber&#x60;, &#x60;retry.orderAttemptNumber&#x60;, and &#x60;retry.skipRetry&#x60; values, we also recommend you provide the &#x60;merchantOrderReference&#x60; to facilitate linking payment attempts together. | [optional] 
**RetryOrderAttemptNumber** | Pointer to **string** | The index of the attempt to bill a particular order, which is identified by the &#x60;merchantOrderReference&#x60; field. For example, if a recurring transaction fails and is retried one day later, then the order number for these attempts would be 1 and 2, respectively.  &gt; If you submit &#x60;retry.chainAttemptNumber&#x60;, &#x60;retry.orderAttemptNumber&#x60;, and &#x60;retry.skipRetry&#x60; values, we also recommend you provide the &#x60;merchantOrderReference&#x60; to facilitate linking payment attempts together. | [optional] 
**RetrySkipRetry** | Pointer to **string** | The Boolean value indicating whether Adyen should skip or retry this transaction, if possible.  &gt; If you submit &#x60;retry.chainAttemptNumber&#x60;, &#x60;retry.orderAttemptNumber&#x60;, and &#x60;retry.skipRetry&#x60; values, we also recommend you provide the &#x60;merchantOrderReference&#x60; to facilitate linking payment attempts together. | [optional] 

## Methods

### NewAdditionalDataRetry

`func NewAdditionalDataRetry() *AdditionalDataRetry`

NewAdditionalDataRetry instantiates a new AdditionalDataRetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalDataRetryWithDefaults

`func NewAdditionalDataRetryWithDefaults() *AdditionalDataRetry`

NewAdditionalDataRetryWithDefaults instantiates a new AdditionalDataRetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetryChainAttemptNumber

`func (o *AdditionalDataRetry) GetRetryChainAttemptNumber() string`

GetRetryChainAttemptNumber returns the RetryChainAttemptNumber field if non-nil, zero value otherwise.

### GetRetryChainAttemptNumberOk

`func (o *AdditionalDataRetry) GetRetryChainAttemptNumberOk() (*string, bool)`

GetRetryChainAttemptNumberOk returns a tuple with the RetryChainAttemptNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryChainAttemptNumber

`func (o *AdditionalDataRetry) SetRetryChainAttemptNumber(v string)`

SetRetryChainAttemptNumber sets RetryChainAttemptNumber field to given value.

### HasRetryChainAttemptNumber

`func (o *AdditionalDataRetry) HasRetryChainAttemptNumber() bool`

HasRetryChainAttemptNumber returns a boolean if a field has been set.

### GetRetryOrderAttemptNumber

`func (o *AdditionalDataRetry) GetRetryOrderAttemptNumber() string`

GetRetryOrderAttemptNumber returns the RetryOrderAttemptNumber field if non-nil, zero value otherwise.

### GetRetryOrderAttemptNumberOk

`func (o *AdditionalDataRetry) GetRetryOrderAttemptNumberOk() (*string, bool)`

GetRetryOrderAttemptNumberOk returns a tuple with the RetryOrderAttemptNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryOrderAttemptNumber

`func (o *AdditionalDataRetry) SetRetryOrderAttemptNumber(v string)`

SetRetryOrderAttemptNumber sets RetryOrderAttemptNumber field to given value.

### HasRetryOrderAttemptNumber

`func (o *AdditionalDataRetry) HasRetryOrderAttemptNumber() bool`

HasRetryOrderAttemptNumber returns a boolean if a field has been set.

### GetRetrySkipRetry

`func (o *AdditionalDataRetry) GetRetrySkipRetry() string`

GetRetrySkipRetry returns the RetrySkipRetry field if non-nil, zero value otherwise.

### GetRetrySkipRetryOk

`func (o *AdditionalDataRetry) GetRetrySkipRetryOk() (*string, bool)`

GetRetrySkipRetryOk returns a tuple with the RetrySkipRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrySkipRetry

`func (o *AdditionalDataRetry) SetRetrySkipRetry(v string)`

SetRetrySkipRetry sets RetrySkipRetry field to given value.

### HasRetrySkipRetry

`func (o *AdditionalDataRetry) HasRetrySkipRetry() bool`

HasRetrySkipRetry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


