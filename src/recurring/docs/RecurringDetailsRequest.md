# RecurringDetailsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantAccount** | **string** | The merchant account identifier you want to process the (transaction) request with. | 
**Recurring** | Pointer to [**Recurring**](Recurring.md) |  | [optional] 
**ShopperReference** | **string** | The reference you use to uniquely identify the shopper (e.g. user ID or account ID). | 

## Methods

### NewRecurringDetailsRequest

`func NewRecurringDetailsRequest(merchantAccount string, shopperReference string, ) *RecurringDetailsRequest`

NewRecurringDetailsRequest instantiates a new RecurringDetailsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurringDetailsRequestWithDefaults

`func NewRecurringDetailsRequestWithDefaults() *RecurringDetailsRequest`

NewRecurringDetailsRequestWithDefaults instantiates a new RecurringDetailsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantAccount

`func (o *RecurringDetailsRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *RecurringDetailsRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *RecurringDetailsRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetRecurring

`func (o *RecurringDetailsRequest) GetRecurring() Recurring`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *RecurringDetailsRequest) GetRecurringOk() (*Recurring, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *RecurringDetailsRequest) SetRecurring(v Recurring)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *RecurringDetailsRequest) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.

### GetShopperReference

`func (o *RecurringDetailsRequest) GetShopperReference() string`

GetShopperReference returns the ShopperReference field if non-nil, zero value otherwise.

### GetShopperReferenceOk

`func (o *RecurringDetailsRequest) GetShopperReferenceOk() (*string, bool)`

GetShopperReferenceOk returns a tuple with the ShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperReference

`func (o *RecurringDetailsRequest) SetShopperReference(v string)`

SetShopperReference sets ShopperReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


