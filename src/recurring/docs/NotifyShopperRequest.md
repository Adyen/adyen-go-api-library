# NotifyShopperRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | [**Amount**](Amount.md) |  | 
**BillingDate** | Pointer to **string** | Date on which the subscription amount will be debited from the shopper. In YYYY-MM-DD format | [optional] 
**BillingSequenceNumber** | Pointer to **string** | Sequence of the debit. Depends on Frequency and Billing Attempts Rule. | [optional] 
**DisplayedReference** | Pointer to **string** | Reference of Pre-debit notification that is displayed to the shopper. Optional field. Maps to reference if missing | [optional] 
**MerchantAccount** | **string** | The merchant account identifier with which you want to process the transaction. | 
**RecurringDetailReference** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 
**Reference** | **string** | Pre-debit notification reference sent by the merchant. This is a mandatory field | 
**ShopperReference** | **string** | The ID that uniquely identifies the shopper.  This &#x60;shopperReference&#x60; must be the same as the &#x60;shopperReference&#x60; used in the initial payment. | 
**StoredPaymentMethodId** | Pointer to **string** | This is the &#x60;recurringDetailReference&#x60; returned in the response when you created the token. | [optional] 

## Methods

### NewNotifyShopperRequest

`func NewNotifyShopperRequest(amount Amount, merchantAccount string, reference string, shopperReference string, ) *NotifyShopperRequest`

NewNotifyShopperRequest instantiates a new NotifyShopperRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyShopperRequestWithDefaults

`func NewNotifyShopperRequestWithDefaults() *NotifyShopperRequest`

NewNotifyShopperRequestWithDefaults instantiates a new NotifyShopperRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *NotifyShopperRequest) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *NotifyShopperRequest) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *NotifyShopperRequest) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetBillingDate

`func (o *NotifyShopperRequest) GetBillingDate() string`

GetBillingDate returns the BillingDate field if non-nil, zero value otherwise.

### GetBillingDateOk

`func (o *NotifyShopperRequest) GetBillingDateOk() (*string, bool)`

GetBillingDateOk returns a tuple with the BillingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingDate

`func (o *NotifyShopperRequest) SetBillingDate(v string)`

SetBillingDate sets BillingDate field to given value.

### HasBillingDate

`func (o *NotifyShopperRequest) HasBillingDate() bool`

HasBillingDate returns a boolean if a field has been set.

### GetBillingSequenceNumber

`func (o *NotifyShopperRequest) GetBillingSequenceNumber() string`

GetBillingSequenceNumber returns the BillingSequenceNumber field if non-nil, zero value otherwise.

### GetBillingSequenceNumberOk

`func (o *NotifyShopperRequest) GetBillingSequenceNumberOk() (*string, bool)`

GetBillingSequenceNumberOk returns a tuple with the BillingSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingSequenceNumber

`func (o *NotifyShopperRequest) SetBillingSequenceNumber(v string)`

SetBillingSequenceNumber sets BillingSequenceNumber field to given value.

### HasBillingSequenceNumber

`func (o *NotifyShopperRequest) HasBillingSequenceNumber() bool`

HasBillingSequenceNumber returns a boolean if a field has been set.

### GetDisplayedReference

`func (o *NotifyShopperRequest) GetDisplayedReference() string`

GetDisplayedReference returns the DisplayedReference field if non-nil, zero value otherwise.

### GetDisplayedReferenceOk

`func (o *NotifyShopperRequest) GetDisplayedReferenceOk() (*string, bool)`

GetDisplayedReferenceOk returns a tuple with the DisplayedReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayedReference

`func (o *NotifyShopperRequest) SetDisplayedReference(v string)`

SetDisplayedReference sets DisplayedReference field to given value.

### HasDisplayedReference

`func (o *NotifyShopperRequest) HasDisplayedReference() bool`

HasDisplayedReference returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *NotifyShopperRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *NotifyShopperRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *NotifyShopperRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetRecurringDetailReference

`func (o *NotifyShopperRequest) GetRecurringDetailReference() string`

GetRecurringDetailReference returns the RecurringDetailReference field if non-nil, zero value otherwise.

### GetRecurringDetailReferenceOk

`func (o *NotifyShopperRequest) GetRecurringDetailReferenceOk() (*string, bool)`

GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDetailReference

`func (o *NotifyShopperRequest) SetRecurringDetailReference(v string)`

SetRecurringDetailReference sets RecurringDetailReference field to given value.

### HasRecurringDetailReference

`func (o *NotifyShopperRequest) HasRecurringDetailReference() bool`

HasRecurringDetailReference returns a boolean if a field has been set.

### GetReference

`func (o *NotifyShopperRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *NotifyShopperRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *NotifyShopperRequest) SetReference(v string)`

SetReference sets Reference field to given value.


### GetShopperReference

`func (o *NotifyShopperRequest) GetShopperReference() string`

GetShopperReference returns the ShopperReference field if non-nil, zero value otherwise.

### GetShopperReferenceOk

`func (o *NotifyShopperRequest) GetShopperReferenceOk() (*string, bool)`

GetShopperReferenceOk returns a tuple with the ShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperReference

`func (o *NotifyShopperRequest) SetShopperReference(v string)`

SetShopperReference sets ShopperReference field to given value.


### GetStoredPaymentMethodId

`func (o *NotifyShopperRequest) GetStoredPaymentMethodId() string`

GetStoredPaymentMethodId returns the StoredPaymentMethodId field if non-nil, zero value otherwise.

### GetStoredPaymentMethodIdOk

`func (o *NotifyShopperRequest) GetStoredPaymentMethodIdOk() (*string, bool)`

GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredPaymentMethodId

`func (o *NotifyShopperRequest) SetStoredPaymentMethodId(v string)`

SetStoredPaymentMethodId sets StoredPaymentMethodId field to given value.

### HasStoredPaymentMethodId

`func (o *NotifyShopperRequest) HasStoredPaymentMethodId() bool`

HasStoredPaymentMethodId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


