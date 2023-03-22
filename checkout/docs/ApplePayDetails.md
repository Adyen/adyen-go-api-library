# ApplePayDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplePayToken** | **string** | The stringified and base64 encoded &#x60;paymentData&#x60; you retrieved from the Apple framework. | 
**CheckoutAttemptId** | Pointer to **string** | The checkout attempt identifier. | [optional] 
**FundingSource** | Pointer to **string** | The funding source that should be used when multiple sources are available. For Brazilian combo cards, by default the funding source is credit. To use debit, set this value to **debit**. | [optional] 
**RecurringDetailReference** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 
**StoredPaymentMethodId** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 
**Type** | Pointer to **string** | **applepay** | [optional] [default to "applepay"]

## Methods

### NewApplePayDetails

`func NewApplePayDetails(applePayToken string, ) *ApplePayDetails`

NewApplePayDetails instantiates a new ApplePayDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplePayDetailsWithDefaults

`func NewApplePayDetailsWithDefaults() *ApplePayDetails`

NewApplePayDetailsWithDefaults instantiates a new ApplePayDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplePayToken

`func (o *ApplePayDetails) GetApplePayToken() string`

GetApplePayToken returns the ApplePayToken field if non-nil, zero value otherwise.

### GetApplePayTokenOk

`func (o *ApplePayDetails) GetApplePayTokenOk() (*string, bool)`

GetApplePayTokenOk returns a tuple with the ApplePayToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplePayToken

`func (o *ApplePayDetails) SetApplePayToken(v string)`

SetApplePayToken sets ApplePayToken field to given value.


### GetCheckoutAttemptId

`func (o *ApplePayDetails) GetCheckoutAttemptId() string`

GetCheckoutAttemptId returns the CheckoutAttemptId field if non-nil, zero value otherwise.

### GetCheckoutAttemptIdOk

`func (o *ApplePayDetails) GetCheckoutAttemptIdOk() (*string, bool)`

GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutAttemptId

`func (o *ApplePayDetails) SetCheckoutAttemptId(v string)`

SetCheckoutAttemptId sets CheckoutAttemptId field to given value.

### HasCheckoutAttemptId

`func (o *ApplePayDetails) HasCheckoutAttemptId() bool`

HasCheckoutAttemptId returns a boolean if a field has been set.

### GetFundingSource

`func (o *ApplePayDetails) GetFundingSource() string`

GetFundingSource returns the FundingSource field if non-nil, zero value otherwise.

### GetFundingSourceOk

`func (o *ApplePayDetails) GetFundingSourceOk() (*string, bool)`

GetFundingSourceOk returns a tuple with the FundingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSource

`func (o *ApplePayDetails) SetFundingSource(v string)`

SetFundingSource sets FundingSource field to given value.

### HasFundingSource

`func (o *ApplePayDetails) HasFundingSource() bool`

HasFundingSource returns a boolean if a field has been set.

### GetRecurringDetailReference

`func (o *ApplePayDetails) GetRecurringDetailReference() string`

GetRecurringDetailReference returns the RecurringDetailReference field if non-nil, zero value otherwise.

### GetRecurringDetailReferenceOk

`func (o *ApplePayDetails) GetRecurringDetailReferenceOk() (*string, bool)`

GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDetailReference

`func (o *ApplePayDetails) SetRecurringDetailReference(v string)`

SetRecurringDetailReference sets RecurringDetailReference field to given value.

### HasRecurringDetailReference

`func (o *ApplePayDetails) HasRecurringDetailReference() bool`

HasRecurringDetailReference returns a boolean if a field has been set.

### GetStoredPaymentMethodId

`func (o *ApplePayDetails) GetStoredPaymentMethodId() string`

GetStoredPaymentMethodId returns the StoredPaymentMethodId field if non-nil, zero value otherwise.

### GetStoredPaymentMethodIdOk

`func (o *ApplePayDetails) GetStoredPaymentMethodIdOk() (*string, bool)`

GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredPaymentMethodId

`func (o *ApplePayDetails) SetStoredPaymentMethodId(v string)`

SetStoredPaymentMethodId sets StoredPaymentMethodId field to given value.

### HasStoredPaymentMethodId

`func (o *ApplePayDetails) HasStoredPaymentMethodId() bool`

HasStoredPaymentMethodId returns a boolean if a field has been set.

### GetType

`func (o *ApplePayDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplePayDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplePayDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApplePayDetails) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


