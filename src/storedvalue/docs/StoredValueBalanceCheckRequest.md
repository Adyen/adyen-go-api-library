# StoredValueBalanceCheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**MerchantAccount** | **string** | The merchant account identifier, with which you want to process the transaction. | 
**PaymentMethod** | **map[string]string** | The collection that contains the type of the payment method and its specific information if available | 
**RecurringDetailReference** | Pointer to **string** |  | [optional] 
**Reference** | **string** | The reference to uniquely identify a payment. This reference is used in all communication with you about the payment status. We recommend using a unique value per payment; however, it is not a requirement. If you need to provide multiple references for a transaction, separate them with hyphens (\&quot;-\&quot;). Maximum length: 80 characters. | 
**ShopperInteraction** | Pointer to **string** | Specifies the sales channel, through which the shopper gives their card details, and whether the shopper is a returning customer. For the web service API, Adyen assumes Ecommerce shopper interaction by default.  This field has the following possible values: * &#x60;Ecommerce&#x60; - Online transactions where the cardholder is present (online). For better authorisation rates, we recommend sending the card security code (CSC) along with the request. * &#x60;ContAuth&#x60; - Card on file and/or subscription transactions, where the cardholder is known to the merchant (returning customer). If the shopper is present (online), you can supply also the CSC to improve authorisation (one-click payment). * &#x60;Moto&#x60; - Mail-order and telephone-order transactions where the shopper is in contact with the merchant via email or telephone. * &#x60;POS&#x60; - Point-of-sale transactions where the shopper is physically present to make a payment using a secure payment terminal. | [optional] 
**ShopperReference** | Pointer to **string** |  | [optional] 
**Store** | Pointer to **string** | The physical store, for which this payment is processed. | [optional] 

## Methods

### NewStoredValueBalanceCheckRequest

`func NewStoredValueBalanceCheckRequest(merchantAccount string, paymentMethod map[string]string, reference string, ) *StoredValueBalanceCheckRequest`

NewStoredValueBalanceCheckRequest instantiates a new StoredValueBalanceCheckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoredValueBalanceCheckRequestWithDefaults

`func NewStoredValueBalanceCheckRequestWithDefaults() *StoredValueBalanceCheckRequest`

NewStoredValueBalanceCheckRequestWithDefaults instantiates a new StoredValueBalanceCheckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *StoredValueBalanceCheckRequest) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *StoredValueBalanceCheckRequest) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *StoredValueBalanceCheckRequest) SetAmount(v Amount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *StoredValueBalanceCheckRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *StoredValueBalanceCheckRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *StoredValueBalanceCheckRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *StoredValueBalanceCheckRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetPaymentMethod

`func (o *StoredValueBalanceCheckRequest) GetPaymentMethod() map[string]string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *StoredValueBalanceCheckRequest) GetPaymentMethodOk() (*map[string]string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *StoredValueBalanceCheckRequest) SetPaymentMethod(v map[string]string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetRecurringDetailReference

`func (o *StoredValueBalanceCheckRequest) GetRecurringDetailReference() string`

GetRecurringDetailReference returns the RecurringDetailReference field if non-nil, zero value otherwise.

### GetRecurringDetailReferenceOk

`func (o *StoredValueBalanceCheckRequest) GetRecurringDetailReferenceOk() (*string, bool)`

GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDetailReference

`func (o *StoredValueBalanceCheckRequest) SetRecurringDetailReference(v string)`

SetRecurringDetailReference sets RecurringDetailReference field to given value.

### HasRecurringDetailReference

`func (o *StoredValueBalanceCheckRequest) HasRecurringDetailReference() bool`

HasRecurringDetailReference returns a boolean if a field has been set.

### GetReference

`func (o *StoredValueBalanceCheckRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *StoredValueBalanceCheckRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *StoredValueBalanceCheckRequest) SetReference(v string)`

SetReference sets Reference field to given value.


### GetShopperInteraction

`func (o *StoredValueBalanceCheckRequest) GetShopperInteraction() string`

GetShopperInteraction returns the ShopperInteraction field if non-nil, zero value otherwise.

### GetShopperInteractionOk

`func (o *StoredValueBalanceCheckRequest) GetShopperInteractionOk() (*string, bool)`

GetShopperInteractionOk returns a tuple with the ShopperInteraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperInteraction

`func (o *StoredValueBalanceCheckRequest) SetShopperInteraction(v string)`

SetShopperInteraction sets ShopperInteraction field to given value.

### HasShopperInteraction

`func (o *StoredValueBalanceCheckRequest) HasShopperInteraction() bool`

HasShopperInteraction returns a boolean if a field has been set.

### GetShopperReference

`func (o *StoredValueBalanceCheckRequest) GetShopperReference() string`

GetShopperReference returns the ShopperReference field if non-nil, zero value otherwise.

### GetShopperReferenceOk

`func (o *StoredValueBalanceCheckRequest) GetShopperReferenceOk() (*string, bool)`

GetShopperReferenceOk returns a tuple with the ShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperReference

`func (o *StoredValueBalanceCheckRequest) SetShopperReference(v string)`

SetShopperReference sets ShopperReference field to given value.

### HasShopperReference

`func (o *StoredValueBalanceCheckRequest) HasShopperReference() bool`

HasShopperReference returns a boolean if a field has been set.

### GetStore

`func (o *StoredValueBalanceCheckRequest) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *StoredValueBalanceCheckRequest) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *StoredValueBalanceCheckRequest) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *StoredValueBalanceCheckRequest) HasStore() bool`

HasStore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


