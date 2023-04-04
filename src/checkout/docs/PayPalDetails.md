# PayPalDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckoutAttemptId** | Pointer to **string** | The checkout attempt identifier. | [optional] 
**OrderID** | Pointer to **string** | The unique ID associated with the order. | [optional] 
**PayerID** | Pointer to **string** | The unique ID associated with the payer. | [optional] 
**RecurringDetailReference** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 
**StoredPaymentMethodId** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 
**Subtype** | Pointer to **string** | The type of flow to initiate. | [optional] 
**Type** | **string** | **paypal** | [default to "paypal"]

## Methods

### NewPayPalDetails

`func NewPayPalDetails(type_ string, ) *PayPalDetails`

NewPayPalDetails instantiates a new PayPalDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayPalDetailsWithDefaults

`func NewPayPalDetailsWithDefaults() *PayPalDetails`

NewPayPalDetailsWithDefaults instantiates a new PayPalDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckoutAttemptId

`func (o *PayPalDetails) GetCheckoutAttemptId() string`

GetCheckoutAttemptId returns the CheckoutAttemptId field if non-nil, zero value otherwise.

### GetCheckoutAttemptIdOk

`func (o *PayPalDetails) GetCheckoutAttemptIdOk() (*string, bool)`

GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutAttemptId

`func (o *PayPalDetails) SetCheckoutAttemptId(v string)`

SetCheckoutAttemptId sets CheckoutAttemptId field to given value.

### HasCheckoutAttemptId

`func (o *PayPalDetails) HasCheckoutAttemptId() bool`

HasCheckoutAttemptId returns a boolean if a field has been set.

### GetOrderID

`func (o *PayPalDetails) GetOrderID() string`

GetOrderID returns the OrderID field if non-nil, zero value otherwise.

### GetOrderIDOk

`func (o *PayPalDetails) GetOrderIDOk() (*string, bool)`

GetOrderIDOk returns a tuple with the OrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderID

`func (o *PayPalDetails) SetOrderID(v string)`

SetOrderID sets OrderID field to given value.

### HasOrderID

`func (o *PayPalDetails) HasOrderID() bool`

HasOrderID returns a boolean if a field has been set.

### GetPayerID

`func (o *PayPalDetails) GetPayerID() string`

GetPayerID returns the PayerID field if non-nil, zero value otherwise.

### GetPayerIDOk

`func (o *PayPalDetails) GetPayerIDOk() (*string, bool)`

GetPayerIDOk returns a tuple with the PayerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerID

`func (o *PayPalDetails) SetPayerID(v string)`

SetPayerID sets PayerID field to given value.

### HasPayerID

`func (o *PayPalDetails) HasPayerID() bool`

HasPayerID returns a boolean if a field has been set.

### GetRecurringDetailReference

`func (o *PayPalDetails) GetRecurringDetailReference() string`

GetRecurringDetailReference returns the RecurringDetailReference field if non-nil, zero value otherwise.

### GetRecurringDetailReferenceOk

`func (o *PayPalDetails) GetRecurringDetailReferenceOk() (*string, bool)`

GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDetailReference

`func (o *PayPalDetails) SetRecurringDetailReference(v string)`

SetRecurringDetailReference sets RecurringDetailReference field to given value.

### HasRecurringDetailReference

`func (o *PayPalDetails) HasRecurringDetailReference() bool`

HasRecurringDetailReference returns a boolean if a field has been set.

### GetStoredPaymentMethodId

`func (o *PayPalDetails) GetStoredPaymentMethodId() string`

GetStoredPaymentMethodId returns the StoredPaymentMethodId field if non-nil, zero value otherwise.

### GetStoredPaymentMethodIdOk

`func (o *PayPalDetails) GetStoredPaymentMethodIdOk() (*string, bool)`

GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredPaymentMethodId

`func (o *PayPalDetails) SetStoredPaymentMethodId(v string)`

SetStoredPaymentMethodId sets StoredPaymentMethodId field to given value.

### HasStoredPaymentMethodId

`func (o *PayPalDetails) HasStoredPaymentMethodId() bool`

HasStoredPaymentMethodId returns a boolean if a field has been set.

### GetSubtype

`func (o *PayPalDetails) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *PayPalDetails) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *PayPalDetails) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *PayPalDetails) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetType

`func (o *PayPalDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PayPalDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PayPalDetails) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


