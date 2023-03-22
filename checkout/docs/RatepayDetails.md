# RatepayDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingAddress** | Pointer to **string** | The address where to send the invoice. | [optional] 
**CheckoutAttemptId** | Pointer to **string** | The checkout attempt identifier. | [optional] 
**DeliveryAddress** | Pointer to **string** | The address where the goods should be delivered. | [optional] 
**PersonalDetails** | Pointer to **string** | Shopper name, date of birth, phone number, and email address. | [optional] 
**RecurringDetailReference** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 
**StoredPaymentMethodId** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 
**Type** | **string** | **ratepay** | [default to "ratepay"]

## Methods

### NewRatepayDetails

`func NewRatepayDetails(type_ string, ) *RatepayDetails`

NewRatepayDetails instantiates a new RatepayDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatepayDetailsWithDefaults

`func NewRatepayDetailsWithDefaults() *RatepayDetails`

NewRatepayDetailsWithDefaults instantiates a new RatepayDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingAddress

`func (o *RatepayDetails) GetBillingAddress() string`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *RatepayDetails) GetBillingAddressOk() (*string, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *RatepayDetails) SetBillingAddress(v string)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *RatepayDetails) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetCheckoutAttemptId

`func (o *RatepayDetails) GetCheckoutAttemptId() string`

GetCheckoutAttemptId returns the CheckoutAttemptId field if non-nil, zero value otherwise.

### GetCheckoutAttemptIdOk

`func (o *RatepayDetails) GetCheckoutAttemptIdOk() (*string, bool)`

GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutAttemptId

`func (o *RatepayDetails) SetCheckoutAttemptId(v string)`

SetCheckoutAttemptId sets CheckoutAttemptId field to given value.

### HasCheckoutAttemptId

`func (o *RatepayDetails) HasCheckoutAttemptId() bool`

HasCheckoutAttemptId returns a boolean if a field has been set.

### GetDeliveryAddress

`func (o *RatepayDetails) GetDeliveryAddress() string`

GetDeliveryAddress returns the DeliveryAddress field if non-nil, zero value otherwise.

### GetDeliveryAddressOk

`func (o *RatepayDetails) GetDeliveryAddressOk() (*string, bool)`

GetDeliveryAddressOk returns a tuple with the DeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddress

`func (o *RatepayDetails) SetDeliveryAddress(v string)`

SetDeliveryAddress sets DeliveryAddress field to given value.

### HasDeliveryAddress

`func (o *RatepayDetails) HasDeliveryAddress() bool`

HasDeliveryAddress returns a boolean if a field has been set.

### GetPersonalDetails

`func (o *RatepayDetails) GetPersonalDetails() string`

GetPersonalDetails returns the PersonalDetails field if non-nil, zero value otherwise.

### GetPersonalDetailsOk

`func (o *RatepayDetails) GetPersonalDetailsOk() (*string, bool)`

GetPersonalDetailsOk returns a tuple with the PersonalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalDetails

`func (o *RatepayDetails) SetPersonalDetails(v string)`

SetPersonalDetails sets PersonalDetails field to given value.

### HasPersonalDetails

`func (o *RatepayDetails) HasPersonalDetails() bool`

HasPersonalDetails returns a boolean if a field has been set.

### GetRecurringDetailReference

`func (o *RatepayDetails) GetRecurringDetailReference() string`

GetRecurringDetailReference returns the RecurringDetailReference field if non-nil, zero value otherwise.

### GetRecurringDetailReferenceOk

`func (o *RatepayDetails) GetRecurringDetailReferenceOk() (*string, bool)`

GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDetailReference

`func (o *RatepayDetails) SetRecurringDetailReference(v string)`

SetRecurringDetailReference sets RecurringDetailReference field to given value.

### HasRecurringDetailReference

`func (o *RatepayDetails) HasRecurringDetailReference() bool`

HasRecurringDetailReference returns a boolean if a field has been set.

### GetStoredPaymentMethodId

`func (o *RatepayDetails) GetStoredPaymentMethodId() string`

GetStoredPaymentMethodId returns the StoredPaymentMethodId field if non-nil, zero value otherwise.

### GetStoredPaymentMethodIdOk

`func (o *RatepayDetails) GetStoredPaymentMethodIdOk() (*string, bool)`

GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredPaymentMethodId

`func (o *RatepayDetails) SetStoredPaymentMethodId(v string)`

SetStoredPaymentMethodId sets StoredPaymentMethodId field to given value.

### HasStoredPaymentMethodId

`func (o *RatepayDetails) HasStoredPaymentMethodId() bool`

HasStoredPaymentMethodId returns a boolean if a field has been set.

### GetType

`func (o *RatepayDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RatepayDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RatepayDetails) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


