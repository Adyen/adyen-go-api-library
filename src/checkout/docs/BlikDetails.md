# BlikDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlikCode** | Pointer to **string** | BLIK code consisting of 6 digits. | [optional] 
**CheckoutAttemptId** | Pointer to **string** | The checkout attempt identifier. | [optional] 
**RecurringDetailReference** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 
**StoredPaymentMethodId** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 
**Type** | Pointer to **string** | **blik** | [optional] 

## Methods

### NewBlikDetails

`func NewBlikDetails() *BlikDetails`

NewBlikDetails instantiates a new BlikDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlikDetailsWithDefaults

`func NewBlikDetailsWithDefaults() *BlikDetails`

NewBlikDetailsWithDefaults instantiates a new BlikDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlikCode

`func (o *BlikDetails) GetBlikCode() string`

GetBlikCode returns the BlikCode field if non-nil, zero value otherwise.

### GetBlikCodeOk

`func (o *BlikDetails) GetBlikCodeOk() (*string, bool)`

GetBlikCodeOk returns a tuple with the BlikCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlikCode

`func (o *BlikDetails) SetBlikCode(v string)`

SetBlikCode sets BlikCode field to given value.

### HasBlikCode

`func (o *BlikDetails) HasBlikCode() bool`

HasBlikCode returns a boolean if a field has been set.

### GetCheckoutAttemptId

`func (o *BlikDetails) GetCheckoutAttemptId() string`

GetCheckoutAttemptId returns the CheckoutAttemptId field if non-nil, zero value otherwise.

### GetCheckoutAttemptIdOk

`func (o *BlikDetails) GetCheckoutAttemptIdOk() (*string, bool)`

GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutAttemptId

`func (o *BlikDetails) SetCheckoutAttemptId(v string)`

SetCheckoutAttemptId sets CheckoutAttemptId field to given value.

### HasCheckoutAttemptId

`func (o *BlikDetails) HasCheckoutAttemptId() bool`

HasCheckoutAttemptId returns a boolean if a field has been set.

### GetRecurringDetailReference

`func (o *BlikDetails) GetRecurringDetailReference() string`

GetRecurringDetailReference returns the RecurringDetailReference field if non-nil, zero value otherwise.

### GetRecurringDetailReferenceOk

`func (o *BlikDetails) GetRecurringDetailReferenceOk() (*string, bool)`

GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDetailReference

`func (o *BlikDetails) SetRecurringDetailReference(v string)`

SetRecurringDetailReference sets RecurringDetailReference field to given value.

### HasRecurringDetailReference

`func (o *BlikDetails) HasRecurringDetailReference() bool`

HasRecurringDetailReference returns a boolean if a field has been set.

### GetStoredPaymentMethodId

`func (o *BlikDetails) GetStoredPaymentMethodId() string`

GetStoredPaymentMethodId returns the StoredPaymentMethodId field if non-nil, zero value otherwise.

### GetStoredPaymentMethodIdOk

`func (o *BlikDetails) GetStoredPaymentMethodIdOk() (*string, bool)`

GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredPaymentMethodId

`func (o *BlikDetails) SetStoredPaymentMethodId(v string)`

SetStoredPaymentMethodId sets StoredPaymentMethodId field to given value.

### HasStoredPaymentMethodId

`func (o *BlikDetails) HasStoredPaymentMethodId() bool`

HasStoredPaymentMethodId returns a boolean if a field has been set.

### GetType

`func (o *BlikDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlikDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlikDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BlikDetails) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


