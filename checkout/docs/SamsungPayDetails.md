# SamsungPayDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckoutAttemptId** | Pointer to **string** | The checkout attempt identifier. | [optional] 
**FundingSource** | Pointer to **string** | The funding source that should be used when multiple sources are available. For Brazilian combo cards, by default the funding source is credit. To use debit, set this value to **debit**. | [optional] 
**RecurringDetailReference** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 
**SamsungPayToken** | **string** | The payload you received from the Samsung Pay SDK response. | 
**StoredPaymentMethodId** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 
**Type** | Pointer to **string** | **samsungpay** | [optional] [default to "samsungpay"]

## Methods

### NewSamsungPayDetails

`func NewSamsungPayDetails(samsungPayToken string, ) *SamsungPayDetails`

NewSamsungPayDetails instantiates a new SamsungPayDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamsungPayDetailsWithDefaults

`func NewSamsungPayDetailsWithDefaults() *SamsungPayDetails`

NewSamsungPayDetailsWithDefaults instantiates a new SamsungPayDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckoutAttemptId

`func (o *SamsungPayDetails) GetCheckoutAttemptId() string`

GetCheckoutAttemptId returns the CheckoutAttemptId field if non-nil, zero value otherwise.

### GetCheckoutAttemptIdOk

`func (o *SamsungPayDetails) GetCheckoutAttemptIdOk() (*string, bool)`

GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutAttemptId

`func (o *SamsungPayDetails) SetCheckoutAttemptId(v string)`

SetCheckoutAttemptId sets CheckoutAttemptId field to given value.

### HasCheckoutAttemptId

`func (o *SamsungPayDetails) HasCheckoutAttemptId() bool`

HasCheckoutAttemptId returns a boolean if a field has been set.

### GetFundingSource

`func (o *SamsungPayDetails) GetFundingSource() string`

GetFundingSource returns the FundingSource field if non-nil, zero value otherwise.

### GetFundingSourceOk

`func (o *SamsungPayDetails) GetFundingSourceOk() (*string, bool)`

GetFundingSourceOk returns a tuple with the FundingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSource

`func (o *SamsungPayDetails) SetFundingSource(v string)`

SetFundingSource sets FundingSource field to given value.

### HasFundingSource

`func (o *SamsungPayDetails) HasFundingSource() bool`

HasFundingSource returns a boolean if a field has been set.

### GetRecurringDetailReference

`func (o *SamsungPayDetails) GetRecurringDetailReference() string`

GetRecurringDetailReference returns the RecurringDetailReference field if non-nil, zero value otherwise.

### GetRecurringDetailReferenceOk

`func (o *SamsungPayDetails) GetRecurringDetailReferenceOk() (*string, bool)`

GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDetailReference

`func (o *SamsungPayDetails) SetRecurringDetailReference(v string)`

SetRecurringDetailReference sets RecurringDetailReference field to given value.

### HasRecurringDetailReference

`func (o *SamsungPayDetails) HasRecurringDetailReference() bool`

HasRecurringDetailReference returns a boolean if a field has been set.

### GetSamsungPayToken

`func (o *SamsungPayDetails) GetSamsungPayToken() string`

GetSamsungPayToken returns the SamsungPayToken field if non-nil, zero value otherwise.

### GetSamsungPayTokenOk

`func (o *SamsungPayDetails) GetSamsungPayTokenOk() (*string, bool)`

GetSamsungPayTokenOk returns a tuple with the SamsungPayToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamsungPayToken

`func (o *SamsungPayDetails) SetSamsungPayToken(v string)`

SetSamsungPayToken sets SamsungPayToken field to given value.


### GetStoredPaymentMethodId

`func (o *SamsungPayDetails) GetStoredPaymentMethodId() string`

GetStoredPaymentMethodId returns the StoredPaymentMethodId field if non-nil, zero value otherwise.

### GetStoredPaymentMethodIdOk

`func (o *SamsungPayDetails) GetStoredPaymentMethodIdOk() (*string, bool)`

GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredPaymentMethodId

`func (o *SamsungPayDetails) SetStoredPaymentMethodId(v string)`

SetStoredPaymentMethodId sets StoredPaymentMethodId field to given value.

### HasStoredPaymentMethodId

`func (o *SamsungPayDetails) HasStoredPaymentMethodId() bool`

HasStoredPaymentMethodId returns a boolean if a field has been set.

### GetType

`func (o *SamsungPayDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SamsungPayDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SamsungPayDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SamsungPayDetails) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


