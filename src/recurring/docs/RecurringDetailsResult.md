# RecurringDetailsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDate** | Pointer to **time.Time** | The date when the recurring details were created. | [optional] 
**Details** | Pointer to [**[]RecurringDetailWrapper**](RecurringDetailWrapper.md) | Payment details stored for recurring payments. | [optional] 
**LastKnownShopperEmail** | Pointer to **string** | The most recent email for this shopper (if available). | [optional] 
**ShopperReference** | Pointer to **string** | The reference you use to uniquely identify the shopper (e.g. user ID or account ID). | [optional] 

## Methods

### NewRecurringDetailsResult

`func NewRecurringDetailsResult() *RecurringDetailsResult`

NewRecurringDetailsResult instantiates a new RecurringDetailsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurringDetailsResultWithDefaults

`func NewRecurringDetailsResultWithDefaults() *RecurringDetailsResult`

NewRecurringDetailsResultWithDefaults instantiates a new RecurringDetailsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationDate

`func (o *RecurringDetailsResult) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *RecurringDetailsResult) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *RecurringDetailsResult) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *RecurringDetailsResult) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDetails

`func (o *RecurringDetailsResult) GetDetails() []RecurringDetailWrapper`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RecurringDetailsResult) GetDetailsOk() (*[]RecurringDetailWrapper, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *RecurringDetailsResult) SetDetails(v []RecurringDetailWrapper)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *RecurringDetailsResult) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetLastKnownShopperEmail

`func (o *RecurringDetailsResult) GetLastKnownShopperEmail() string`

GetLastKnownShopperEmail returns the LastKnownShopperEmail field if non-nil, zero value otherwise.

### GetLastKnownShopperEmailOk

`func (o *RecurringDetailsResult) GetLastKnownShopperEmailOk() (*string, bool)`

GetLastKnownShopperEmailOk returns a tuple with the LastKnownShopperEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastKnownShopperEmail

`func (o *RecurringDetailsResult) SetLastKnownShopperEmail(v string)`

SetLastKnownShopperEmail sets LastKnownShopperEmail field to given value.

### HasLastKnownShopperEmail

`func (o *RecurringDetailsResult) HasLastKnownShopperEmail() bool`

HasLastKnownShopperEmail returns a boolean if a field has been set.

### GetShopperReference

`func (o *RecurringDetailsResult) GetShopperReference() string`

GetShopperReference returns the ShopperReference field if non-nil, zero value otherwise.

### GetShopperReferenceOk

`func (o *RecurringDetailsResult) GetShopperReferenceOk() (*string, bool)`

GetShopperReferenceOk returns a tuple with the ShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperReference

`func (o *RecurringDetailsResult) SetShopperReference(v string)`

SetShopperReference sets ShopperReference field to given value.

### HasShopperReference

`func (o *RecurringDetailsResult) HasShopperReference() bool`

HasShopperReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


