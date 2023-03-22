# PayUUpiDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckoutAttemptId** | Pointer to **string** | The checkout attempt identifier. | [optional] 
**RecurringDetailReference** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 
**ShopperNotificationReference** | Pointer to **string** | The &#x60;shopperNotificationReference&#x60; returned in the response when you requested to notify the shopper. Used for recurring payment only. | [optional] 
**StoredPaymentMethodId** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 
**Type** | **string** | **payu_IN_upi** | [default to "payu_IN_upi"]
**VirtualPaymentAddress** | Pointer to **string** | The virtual payment address for UPI. | [optional] 

## Methods

### NewPayUUpiDetails

`func NewPayUUpiDetails(type_ string, ) *PayUUpiDetails`

NewPayUUpiDetails instantiates a new PayUUpiDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayUUpiDetailsWithDefaults

`func NewPayUUpiDetailsWithDefaults() *PayUUpiDetails`

NewPayUUpiDetailsWithDefaults instantiates a new PayUUpiDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckoutAttemptId

`func (o *PayUUpiDetails) GetCheckoutAttemptId() string`

GetCheckoutAttemptId returns the CheckoutAttemptId field if non-nil, zero value otherwise.

### GetCheckoutAttemptIdOk

`func (o *PayUUpiDetails) GetCheckoutAttemptIdOk() (*string, bool)`

GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutAttemptId

`func (o *PayUUpiDetails) SetCheckoutAttemptId(v string)`

SetCheckoutAttemptId sets CheckoutAttemptId field to given value.

### HasCheckoutAttemptId

`func (o *PayUUpiDetails) HasCheckoutAttemptId() bool`

HasCheckoutAttemptId returns a boolean if a field has been set.

### GetRecurringDetailReference

`func (o *PayUUpiDetails) GetRecurringDetailReference() string`

GetRecurringDetailReference returns the RecurringDetailReference field if non-nil, zero value otherwise.

### GetRecurringDetailReferenceOk

`func (o *PayUUpiDetails) GetRecurringDetailReferenceOk() (*string, bool)`

GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDetailReference

`func (o *PayUUpiDetails) SetRecurringDetailReference(v string)`

SetRecurringDetailReference sets RecurringDetailReference field to given value.

### HasRecurringDetailReference

`func (o *PayUUpiDetails) HasRecurringDetailReference() bool`

HasRecurringDetailReference returns a boolean if a field has been set.

### GetShopperNotificationReference

`func (o *PayUUpiDetails) GetShopperNotificationReference() string`

GetShopperNotificationReference returns the ShopperNotificationReference field if non-nil, zero value otherwise.

### GetShopperNotificationReferenceOk

`func (o *PayUUpiDetails) GetShopperNotificationReferenceOk() (*string, bool)`

GetShopperNotificationReferenceOk returns a tuple with the ShopperNotificationReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperNotificationReference

`func (o *PayUUpiDetails) SetShopperNotificationReference(v string)`

SetShopperNotificationReference sets ShopperNotificationReference field to given value.

### HasShopperNotificationReference

`func (o *PayUUpiDetails) HasShopperNotificationReference() bool`

HasShopperNotificationReference returns a boolean if a field has been set.

### GetStoredPaymentMethodId

`func (o *PayUUpiDetails) GetStoredPaymentMethodId() string`

GetStoredPaymentMethodId returns the StoredPaymentMethodId field if non-nil, zero value otherwise.

### GetStoredPaymentMethodIdOk

`func (o *PayUUpiDetails) GetStoredPaymentMethodIdOk() (*string, bool)`

GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredPaymentMethodId

`func (o *PayUUpiDetails) SetStoredPaymentMethodId(v string)`

SetStoredPaymentMethodId sets StoredPaymentMethodId field to given value.

### HasStoredPaymentMethodId

`func (o *PayUUpiDetails) HasStoredPaymentMethodId() bool`

HasStoredPaymentMethodId returns a boolean if a field has been set.

### GetType

`func (o *PayUUpiDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PayUUpiDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PayUUpiDetails) SetType(v string)`

SetType sets Type field to given value.


### GetVirtualPaymentAddress

`func (o *PayUUpiDetails) GetVirtualPaymentAddress() string`

GetVirtualPaymentAddress returns the VirtualPaymentAddress field if non-nil, zero value otherwise.

### GetVirtualPaymentAddressOk

`func (o *PayUUpiDetails) GetVirtualPaymentAddressOk() (*string, bool)`

GetVirtualPaymentAddressOk returns a tuple with the VirtualPaymentAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPaymentAddress

`func (o *PayUUpiDetails) SetVirtualPaymentAddress(v string)`

SetVirtualPaymentAddress sets VirtualPaymentAddress field to given value.

### HasVirtualPaymentAddress

`func (o *PayUUpiDetails) HasVirtualPaymentAddress() bool`

HasVirtualPaymentAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


