# CreatePermitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantAccount** | **string** | The merchant account identifier, with which you want to process the transaction. | 
**Permits** | [**[]Permit**](Permit.md) | The permits to create for this recurring contract. | 
**RecurringDetailReference** | **string** | The recurring contract the new permits will use. | 
**ShopperReference** | **string** | The shopper&#39;s reference to uniquely identify this shopper (e.g. user ID or account ID). | 

## Methods

### NewCreatePermitRequest

`func NewCreatePermitRequest(merchantAccount string, permits []Permit, recurringDetailReference string, shopperReference string, ) *CreatePermitRequest`

NewCreatePermitRequest instantiates a new CreatePermitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePermitRequestWithDefaults

`func NewCreatePermitRequestWithDefaults() *CreatePermitRequest`

NewCreatePermitRequestWithDefaults instantiates a new CreatePermitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantAccount

`func (o *CreatePermitRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *CreatePermitRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *CreatePermitRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetPermits

`func (o *CreatePermitRequest) GetPermits() []Permit`

GetPermits returns the Permits field if non-nil, zero value otherwise.

### GetPermitsOk

`func (o *CreatePermitRequest) GetPermitsOk() (*[]Permit, bool)`

GetPermitsOk returns a tuple with the Permits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermits

`func (o *CreatePermitRequest) SetPermits(v []Permit)`

SetPermits sets Permits field to given value.


### GetRecurringDetailReference

`func (o *CreatePermitRequest) GetRecurringDetailReference() string`

GetRecurringDetailReference returns the RecurringDetailReference field if non-nil, zero value otherwise.

### GetRecurringDetailReferenceOk

`func (o *CreatePermitRequest) GetRecurringDetailReferenceOk() (*string, bool)`

GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDetailReference

`func (o *CreatePermitRequest) SetRecurringDetailReference(v string)`

SetRecurringDetailReference sets RecurringDetailReference field to given value.


### GetShopperReference

`func (o *CreatePermitRequest) GetShopperReference() string`

GetShopperReference returns the ShopperReference field if non-nil, zero value otherwise.

### GetShopperReferenceOk

`func (o *CreatePermitRequest) GetShopperReferenceOk() (*string, bool)`

GetShopperReferenceOk returns a tuple with the ShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperReference

`func (o *CreatePermitRequest) SetShopperReference(v string)`

SetShopperReference sets ShopperReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


