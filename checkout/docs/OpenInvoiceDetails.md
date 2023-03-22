# OpenInvoiceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingAddress** | Pointer to **string** | The address where to send the invoice. | [optional] 
**CheckoutAttemptId** | Pointer to **string** | The checkout attempt identifier. | [optional] 
**DeliveryAddress** | Pointer to **string** | The address where the goods should be delivered. | [optional] 
**PersonalDetails** | Pointer to **string** | Shopper name, date of birth, phone number, and email address. | [optional] 
**RecurringDetailReference** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 
**StoredPaymentMethodId** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 
**Type** | Pointer to **string** | **openinvoice** | [optional] [default to "openinvoice"]

## Methods

### NewOpenInvoiceDetails

`func NewOpenInvoiceDetails() *OpenInvoiceDetails`

NewOpenInvoiceDetails instantiates a new OpenInvoiceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenInvoiceDetailsWithDefaults

`func NewOpenInvoiceDetailsWithDefaults() *OpenInvoiceDetails`

NewOpenInvoiceDetailsWithDefaults instantiates a new OpenInvoiceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingAddress

`func (o *OpenInvoiceDetails) GetBillingAddress() string`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *OpenInvoiceDetails) GetBillingAddressOk() (*string, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *OpenInvoiceDetails) SetBillingAddress(v string)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *OpenInvoiceDetails) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetCheckoutAttemptId

`func (o *OpenInvoiceDetails) GetCheckoutAttemptId() string`

GetCheckoutAttemptId returns the CheckoutAttemptId field if non-nil, zero value otherwise.

### GetCheckoutAttemptIdOk

`func (o *OpenInvoiceDetails) GetCheckoutAttemptIdOk() (*string, bool)`

GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutAttemptId

`func (o *OpenInvoiceDetails) SetCheckoutAttemptId(v string)`

SetCheckoutAttemptId sets CheckoutAttemptId field to given value.

### HasCheckoutAttemptId

`func (o *OpenInvoiceDetails) HasCheckoutAttemptId() bool`

HasCheckoutAttemptId returns a boolean if a field has been set.

### GetDeliveryAddress

`func (o *OpenInvoiceDetails) GetDeliveryAddress() string`

GetDeliveryAddress returns the DeliveryAddress field if non-nil, zero value otherwise.

### GetDeliveryAddressOk

`func (o *OpenInvoiceDetails) GetDeliveryAddressOk() (*string, bool)`

GetDeliveryAddressOk returns a tuple with the DeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddress

`func (o *OpenInvoiceDetails) SetDeliveryAddress(v string)`

SetDeliveryAddress sets DeliveryAddress field to given value.

### HasDeliveryAddress

`func (o *OpenInvoiceDetails) HasDeliveryAddress() bool`

HasDeliveryAddress returns a boolean if a field has been set.

### GetPersonalDetails

`func (o *OpenInvoiceDetails) GetPersonalDetails() string`

GetPersonalDetails returns the PersonalDetails field if non-nil, zero value otherwise.

### GetPersonalDetailsOk

`func (o *OpenInvoiceDetails) GetPersonalDetailsOk() (*string, bool)`

GetPersonalDetailsOk returns a tuple with the PersonalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalDetails

`func (o *OpenInvoiceDetails) SetPersonalDetails(v string)`

SetPersonalDetails sets PersonalDetails field to given value.

### HasPersonalDetails

`func (o *OpenInvoiceDetails) HasPersonalDetails() bool`

HasPersonalDetails returns a boolean if a field has been set.

### GetRecurringDetailReference

`func (o *OpenInvoiceDetails) GetRecurringDetailReference() string`

GetRecurringDetailReference returns the RecurringDetailReference field if non-nil, zero value otherwise.

### GetRecurringDetailReferenceOk

`func (o *OpenInvoiceDetails) GetRecurringDetailReferenceOk() (*string, bool)`

GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDetailReference

`func (o *OpenInvoiceDetails) SetRecurringDetailReference(v string)`

SetRecurringDetailReference sets RecurringDetailReference field to given value.

### HasRecurringDetailReference

`func (o *OpenInvoiceDetails) HasRecurringDetailReference() bool`

HasRecurringDetailReference returns a boolean if a field has been set.

### GetStoredPaymentMethodId

`func (o *OpenInvoiceDetails) GetStoredPaymentMethodId() string`

GetStoredPaymentMethodId returns the StoredPaymentMethodId field if non-nil, zero value otherwise.

### GetStoredPaymentMethodIdOk

`func (o *OpenInvoiceDetails) GetStoredPaymentMethodIdOk() (*string, bool)`

GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredPaymentMethodId

`func (o *OpenInvoiceDetails) SetStoredPaymentMethodId(v string)`

SetStoredPaymentMethodId sets StoredPaymentMethodId field to given value.

### HasStoredPaymentMethodId

`func (o *OpenInvoiceDetails) HasStoredPaymentMethodId() bool`

HasStoredPaymentMethodId returns a boolean if a field has been set.

### GetType

`func (o *OpenInvoiceDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpenInvoiceDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpenInvoiceDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OpenInvoiceDetails) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


