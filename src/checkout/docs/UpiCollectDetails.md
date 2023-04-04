# UpiCollectDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingSequenceNumber** | **string** | The sequence number for the debit. For example, send **2** if this is the second debit for the subscription. The sequence number is included in the notification sent to the shopper. | 
**CheckoutAttemptId** | Pointer to **string** | The checkout attempt identifier. | [optional] 
**RecurringDetailReference** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 
**ShopperNotificationReference** | Pointer to **string** | The &#x60;shopperNotificationReference&#x60; returned in the response when you requested to notify the shopper. Used for recurring payment only. | [optional] 
**StoredPaymentMethodId** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 
**Type** | **string** | **upi_collect** | [default to "upi_collect"]
**VirtualPaymentAddress** | Pointer to **string** | The virtual payment address for UPI. | [optional] 

## Methods

### NewUpiCollectDetails

`func NewUpiCollectDetails(billingSequenceNumber string, type_ string, ) *UpiCollectDetails`

NewUpiCollectDetails instantiates a new UpiCollectDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpiCollectDetailsWithDefaults

`func NewUpiCollectDetailsWithDefaults() *UpiCollectDetails`

NewUpiCollectDetailsWithDefaults instantiates a new UpiCollectDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingSequenceNumber

`func (o *UpiCollectDetails) GetBillingSequenceNumber() string`

GetBillingSequenceNumber returns the BillingSequenceNumber field if non-nil, zero value otherwise.

### GetBillingSequenceNumberOk

`func (o *UpiCollectDetails) GetBillingSequenceNumberOk() (*string, bool)`

GetBillingSequenceNumberOk returns a tuple with the BillingSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingSequenceNumber

`func (o *UpiCollectDetails) SetBillingSequenceNumber(v string)`

SetBillingSequenceNumber sets BillingSequenceNumber field to given value.


### GetCheckoutAttemptId

`func (o *UpiCollectDetails) GetCheckoutAttemptId() string`

GetCheckoutAttemptId returns the CheckoutAttemptId field if non-nil, zero value otherwise.

### GetCheckoutAttemptIdOk

`func (o *UpiCollectDetails) GetCheckoutAttemptIdOk() (*string, bool)`

GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutAttemptId

`func (o *UpiCollectDetails) SetCheckoutAttemptId(v string)`

SetCheckoutAttemptId sets CheckoutAttemptId field to given value.

### HasCheckoutAttemptId

`func (o *UpiCollectDetails) HasCheckoutAttemptId() bool`

HasCheckoutAttemptId returns a boolean if a field has been set.

### GetRecurringDetailReference

`func (o *UpiCollectDetails) GetRecurringDetailReference() string`

GetRecurringDetailReference returns the RecurringDetailReference field if non-nil, zero value otherwise.

### GetRecurringDetailReferenceOk

`func (o *UpiCollectDetails) GetRecurringDetailReferenceOk() (*string, bool)`

GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDetailReference

`func (o *UpiCollectDetails) SetRecurringDetailReference(v string)`

SetRecurringDetailReference sets RecurringDetailReference field to given value.

### HasRecurringDetailReference

`func (o *UpiCollectDetails) HasRecurringDetailReference() bool`

HasRecurringDetailReference returns a boolean if a field has been set.

### GetShopperNotificationReference

`func (o *UpiCollectDetails) GetShopperNotificationReference() string`

GetShopperNotificationReference returns the ShopperNotificationReference field if non-nil, zero value otherwise.

### GetShopperNotificationReferenceOk

`func (o *UpiCollectDetails) GetShopperNotificationReferenceOk() (*string, bool)`

GetShopperNotificationReferenceOk returns a tuple with the ShopperNotificationReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperNotificationReference

`func (o *UpiCollectDetails) SetShopperNotificationReference(v string)`

SetShopperNotificationReference sets ShopperNotificationReference field to given value.

### HasShopperNotificationReference

`func (o *UpiCollectDetails) HasShopperNotificationReference() bool`

HasShopperNotificationReference returns a boolean if a field has been set.

### GetStoredPaymentMethodId

`func (o *UpiCollectDetails) GetStoredPaymentMethodId() string`

GetStoredPaymentMethodId returns the StoredPaymentMethodId field if non-nil, zero value otherwise.

### GetStoredPaymentMethodIdOk

`func (o *UpiCollectDetails) GetStoredPaymentMethodIdOk() (*string, bool)`

GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredPaymentMethodId

`func (o *UpiCollectDetails) SetStoredPaymentMethodId(v string)`

SetStoredPaymentMethodId sets StoredPaymentMethodId field to given value.

### HasStoredPaymentMethodId

`func (o *UpiCollectDetails) HasStoredPaymentMethodId() bool`

HasStoredPaymentMethodId returns a boolean if a field has been set.

### GetType

`func (o *UpiCollectDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpiCollectDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpiCollectDetails) SetType(v string)`

SetType sets Type field to given value.


### GetVirtualPaymentAddress

`func (o *UpiCollectDetails) GetVirtualPaymentAddress() string`

GetVirtualPaymentAddress returns the VirtualPaymentAddress field if non-nil, zero value otherwise.

### GetVirtualPaymentAddressOk

`func (o *UpiCollectDetails) GetVirtualPaymentAddressOk() (*string, bool)`

GetVirtualPaymentAddressOk returns a tuple with the VirtualPaymentAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPaymentAddress

`func (o *UpiCollectDetails) SetVirtualPaymentAddress(v string)`

SetVirtualPaymentAddress sets VirtualPaymentAddress field to given value.

### HasVirtualPaymentAddress

`func (o *UpiCollectDetails) HasVirtualPaymentAddress() bool`

HasVirtualPaymentAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


