# StoredValueLoadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | [**Amount**](Amount.md) |  | 
**LoadType** | Pointer to **string** | The type of load you are trying to do, when absent we default to &#39;Load&#39; | [optional] 
**MerchantAccount** | **string** | The merchant account identifier, with which you want to process the transaction. | 
**PaymentMethod** | **map[string]string** | The collection that contains the type of the payment method and its specific information if available | 
**RecurringDetailReference** | Pointer to **string** |  | [optional] 
**Reference** | **string** | The reference to uniquely identify a payment. This reference is used in all communication with you about the payment status. We recommend using a unique value per payment; however, it is not a requirement. If you need to provide multiple references for a transaction, separate them with hyphens (\&quot;-\&quot;). Maximum length: 80 characters. | 
**ShopperInteraction** | Pointer to **string** | Specifies the sales channel, through which the shopper gives their card details, and whether the shopper is a returning customer. For the web service API, Adyen assumes Ecommerce shopper interaction by default.  This field has the following possible values: * &#x60;Ecommerce&#x60; - Online transactions where the cardholder is present (online). For better authorisation rates, we recommend sending the card security code (CSC) along with the request. * &#x60;ContAuth&#x60; - Card on file and/or subscription transactions, where the cardholder is known to the merchant (returning customer). If the shopper is present (online), you can supply also the CSC to improve authorisation (one-click payment). * &#x60;Moto&#x60; - Mail-order and telephone-order transactions where the shopper is in contact with the merchant via email or telephone. * &#x60;POS&#x60; - Point-of-sale transactions where the shopper is physically present to make a payment using a secure payment terminal. | [optional] 
**ShopperReference** | Pointer to **string** |  | [optional] 
**Store** | Pointer to **string** | The physical store, for which this payment is processed. | [optional] 

## Methods

### NewStoredValueLoadRequest

`func NewStoredValueLoadRequest(amount Amount, merchantAccount string, paymentMethod map[string]string, reference string, ) *StoredValueLoadRequest`

NewStoredValueLoadRequest instantiates a new StoredValueLoadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoredValueLoadRequestWithDefaults

`func NewStoredValueLoadRequestWithDefaults() *StoredValueLoadRequest`

NewStoredValueLoadRequestWithDefaults instantiates a new StoredValueLoadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *StoredValueLoadRequest) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *StoredValueLoadRequest) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *StoredValueLoadRequest) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetLoadType

`func (o *StoredValueLoadRequest) GetLoadType() string`

GetLoadType returns the LoadType field if non-nil, zero value otherwise.

### GetLoadTypeOk

`func (o *StoredValueLoadRequest) GetLoadTypeOk() (*string, bool)`

GetLoadTypeOk returns a tuple with the LoadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadType

`func (o *StoredValueLoadRequest) SetLoadType(v string)`

SetLoadType sets LoadType field to given value.

### HasLoadType

`func (o *StoredValueLoadRequest) HasLoadType() bool`

HasLoadType returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *StoredValueLoadRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *StoredValueLoadRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *StoredValueLoadRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetPaymentMethod

`func (o *StoredValueLoadRequest) GetPaymentMethod() map[string]string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *StoredValueLoadRequest) GetPaymentMethodOk() (*map[string]string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *StoredValueLoadRequest) SetPaymentMethod(v map[string]string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetRecurringDetailReference

`func (o *StoredValueLoadRequest) GetRecurringDetailReference() string`

GetRecurringDetailReference returns the RecurringDetailReference field if non-nil, zero value otherwise.

### GetRecurringDetailReferenceOk

`func (o *StoredValueLoadRequest) GetRecurringDetailReferenceOk() (*string, bool)`

GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDetailReference

`func (o *StoredValueLoadRequest) SetRecurringDetailReference(v string)`

SetRecurringDetailReference sets RecurringDetailReference field to given value.

### HasRecurringDetailReference

`func (o *StoredValueLoadRequest) HasRecurringDetailReference() bool`

HasRecurringDetailReference returns a boolean if a field has been set.

### GetReference

`func (o *StoredValueLoadRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *StoredValueLoadRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *StoredValueLoadRequest) SetReference(v string)`

SetReference sets Reference field to given value.


### GetShopperInteraction

`func (o *StoredValueLoadRequest) GetShopperInteraction() string`

GetShopperInteraction returns the ShopperInteraction field if non-nil, zero value otherwise.

### GetShopperInteractionOk

`func (o *StoredValueLoadRequest) GetShopperInteractionOk() (*string, bool)`

GetShopperInteractionOk returns a tuple with the ShopperInteraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperInteraction

`func (o *StoredValueLoadRequest) SetShopperInteraction(v string)`

SetShopperInteraction sets ShopperInteraction field to given value.

### HasShopperInteraction

`func (o *StoredValueLoadRequest) HasShopperInteraction() bool`

HasShopperInteraction returns a boolean if a field has been set.

### GetShopperReference

`func (o *StoredValueLoadRequest) GetShopperReference() string`

GetShopperReference returns the ShopperReference field if non-nil, zero value otherwise.

### GetShopperReferenceOk

`func (o *StoredValueLoadRequest) GetShopperReferenceOk() (*string, bool)`

GetShopperReferenceOk returns a tuple with the ShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperReference

`func (o *StoredValueLoadRequest) SetShopperReference(v string)`

SetShopperReference sets ShopperReference field to given value.

### HasShopperReference

`func (o *StoredValueLoadRequest) HasShopperReference() bool`

HasShopperReference returns a boolean if a field has been set.

### GetStore

`func (o *StoredValueLoadRequest) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *StoredValueLoadRequest) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *StoredValueLoadRequest) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *StoredValueLoadRequest) HasStore() bool`

HasStore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


