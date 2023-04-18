# ScheduleAccountUpdaterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **map[string]string** | This field contains additional data, which may be required for a particular request. | [optional] 
**Card** | Pointer to [**Card**](Card.md) |  | [optional] 
**MerchantAccount** | **string** | Account of the merchant. | 
**Reference** | **string** | A reference that merchants can apply for the call. | 
**SelectedRecurringDetailReference** | Pointer to **string** | The selected detail recurring reference.  Optional if &#x60;card&#x60; is provided. | [optional] 
**ShopperReference** | Pointer to **string** | The reference of the shopper that owns the recurring contract.  Optional if &#x60;card&#x60; is provided. | [optional] 

## Methods

### NewScheduleAccountUpdaterRequest

`func NewScheduleAccountUpdaterRequest(merchantAccount string, reference string, ) *ScheduleAccountUpdaterRequest`

NewScheduleAccountUpdaterRequest instantiates a new ScheduleAccountUpdaterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleAccountUpdaterRequestWithDefaults

`func NewScheduleAccountUpdaterRequestWithDefaults() *ScheduleAccountUpdaterRequest`

NewScheduleAccountUpdaterRequestWithDefaults instantiates a new ScheduleAccountUpdaterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *ScheduleAccountUpdaterRequest) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *ScheduleAccountUpdaterRequest) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *ScheduleAccountUpdaterRequest) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *ScheduleAccountUpdaterRequest) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetCard

`func (o *ScheduleAccountUpdaterRequest) GetCard() Card`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *ScheduleAccountUpdaterRequest) GetCardOk() (*Card, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *ScheduleAccountUpdaterRequest) SetCard(v Card)`

SetCard sets Card field to given value.

### HasCard

`func (o *ScheduleAccountUpdaterRequest) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *ScheduleAccountUpdaterRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *ScheduleAccountUpdaterRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *ScheduleAccountUpdaterRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetReference

`func (o *ScheduleAccountUpdaterRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ScheduleAccountUpdaterRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ScheduleAccountUpdaterRequest) SetReference(v string)`

SetReference sets Reference field to given value.


### GetSelectedRecurringDetailReference

`func (o *ScheduleAccountUpdaterRequest) GetSelectedRecurringDetailReference() string`

GetSelectedRecurringDetailReference returns the SelectedRecurringDetailReference field if non-nil, zero value otherwise.

### GetSelectedRecurringDetailReferenceOk

`func (o *ScheduleAccountUpdaterRequest) GetSelectedRecurringDetailReferenceOk() (*string, bool)`

GetSelectedRecurringDetailReferenceOk returns a tuple with the SelectedRecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRecurringDetailReference

`func (o *ScheduleAccountUpdaterRequest) SetSelectedRecurringDetailReference(v string)`

SetSelectedRecurringDetailReference sets SelectedRecurringDetailReference field to given value.

### HasSelectedRecurringDetailReference

`func (o *ScheduleAccountUpdaterRequest) HasSelectedRecurringDetailReference() bool`

HasSelectedRecurringDetailReference returns a boolean if a field has been set.

### GetShopperReference

`func (o *ScheduleAccountUpdaterRequest) GetShopperReference() string`

GetShopperReference returns the ShopperReference field if non-nil, zero value otherwise.

### GetShopperReferenceOk

`func (o *ScheduleAccountUpdaterRequest) GetShopperReferenceOk() (*string, bool)`

GetShopperReferenceOk returns a tuple with the ShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperReference

`func (o *ScheduleAccountUpdaterRequest) SetShopperReference(v string)`

SetShopperReference sets ShopperReference field to given value.

### HasShopperReference

`func (o *ScheduleAccountUpdaterRequest) HasShopperReference() bool`

HasShopperReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


