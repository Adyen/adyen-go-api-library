# UpdateMerchantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessLineId** | Pointer to **string** | Business line ID | [optional] 
**CaptureDelay** | Pointer to **string** | New time interval between authorising and capturing a payment. Specified in days. Possible values are: - **1** - **2** - **3** - **4** - **5** - **6** - **7** - **immediate** - **manual** | [optional] 
**MerchantCity** | Pointer to **string** | New city merchant operates from. | [optional] 
**Name** | Pointer to **string** | New name for the merchant. | [optional] 
**PricingPlan** | Pointer to **string** | Only applies to merchant accounts managed by Adyen&#39;s partners. The name of the pricing plan assigned to the merchant account. | [optional] 
**PrimarySettlementCurrency** | Pointer to **string** | New primary currency to settle payments. | [optional] 
**ShopWebAddress** | Pointer to **string** | New URL pointing to online shop. | [optional] 

## Methods

### NewUpdateMerchantRequest

`func NewUpdateMerchantRequest() *UpdateMerchantRequest`

NewUpdateMerchantRequest instantiates a new UpdateMerchantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMerchantRequestWithDefaults

`func NewUpdateMerchantRequestWithDefaults() *UpdateMerchantRequest`

NewUpdateMerchantRequestWithDefaults instantiates a new UpdateMerchantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessLineId

`func (o *UpdateMerchantRequest) GetBusinessLineId() string`

GetBusinessLineId returns the BusinessLineId field if non-nil, zero value otherwise.

### GetBusinessLineIdOk

`func (o *UpdateMerchantRequest) GetBusinessLineIdOk() (*string, bool)`

GetBusinessLineIdOk returns a tuple with the BusinessLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessLineId

`func (o *UpdateMerchantRequest) SetBusinessLineId(v string)`

SetBusinessLineId sets BusinessLineId field to given value.

### HasBusinessLineId

`func (o *UpdateMerchantRequest) HasBusinessLineId() bool`

HasBusinessLineId returns a boolean if a field has been set.

### GetCaptureDelay

`func (o *UpdateMerchantRequest) GetCaptureDelay() string`

GetCaptureDelay returns the CaptureDelay field if non-nil, zero value otherwise.

### GetCaptureDelayOk

`func (o *UpdateMerchantRequest) GetCaptureDelayOk() (*string, bool)`

GetCaptureDelayOk returns a tuple with the CaptureDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureDelay

`func (o *UpdateMerchantRequest) SetCaptureDelay(v string)`

SetCaptureDelay sets CaptureDelay field to given value.

### HasCaptureDelay

`func (o *UpdateMerchantRequest) HasCaptureDelay() bool`

HasCaptureDelay returns a boolean if a field has been set.

### GetMerchantCity

`func (o *UpdateMerchantRequest) GetMerchantCity() string`

GetMerchantCity returns the MerchantCity field if non-nil, zero value otherwise.

### GetMerchantCityOk

`func (o *UpdateMerchantRequest) GetMerchantCityOk() (*string, bool)`

GetMerchantCityOk returns a tuple with the MerchantCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCity

`func (o *UpdateMerchantRequest) SetMerchantCity(v string)`

SetMerchantCity sets MerchantCity field to given value.

### HasMerchantCity

`func (o *UpdateMerchantRequest) HasMerchantCity() bool`

HasMerchantCity returns a boolean if a field has been set.

### GetName

`func (o *UpdateMerchantRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMerchantRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMerchantRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateMerchantRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPricingPlan

`func (o *UpdateMerchantRequest) GetPricingPlan() string`

GetPricingPlan returns the PricingPlan field if non-nil, zero value otherwise.

### GetPricingPlanOk

`func (o *UpdateMerchantRequest) GetPricingPlanOk() (*string, bool)`

GetPricingPlanOk returns a tuple with the PricingPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingPlan

`func (o *UpdateMerchantRequest) SetPricingPlan(v string)`

SetPricingPlan sets PricingPlan field to given value.

### HasPricingPlan

`func (o *UpdateMerchantRequest) HasPricingPlan() bool`

HasPricingPlan returns a boolean if a field has been set.

### GetPrimarySettlementCurrency

`func (o *UpdateMerchantRequest) GetPrimarySettlementCurrency() string`

GetPrimarySettlementCurrency returns the PrimarySettlementCurrency field if non-nil, zero value otherwise.

### GetPrimarySettlementCurrencyOk

`func (o *UpdateMerchantRequest) GetPrimarySettlementCurrencyOk() (*string, bool)`

GetPrimarySettlementCurrencyOk returns a tuple with the PrimarySettlementCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySettlementCurrency

`func (o *UpdateMerchantRequest) SetPrimarySettlementCurrency(v string)`

SetPrimarySettlementCurrency sets PrimarySettlementCurrency field to given value.

### HasPrimarySettlementCurrency

`func (o *UpdateMerchantRequest) HasPrimarySettlementCurrency() bool`

HasPrimarySettlementCurrency returns a boolean if a field has been set.

### GetShopWebAddress

`func (o *UpdateMerchantRequest) GetShopWebAddress() string`

GetShopWebAddress returns the ShopWebAddress field if non-nil, zero value otherwise.

### GetShopWebAddressOk

`func (o *UpdateMerchantRequest) GetShopWebAddressOk() (*string, bool)`

GetShopWebAddressOk returns a tuple with the ShopWebAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopWebAddress

`func (o *UpdateMerchantRequest) SetShopWebAddress(v string)`

SetShopWebAddress sets ShopWebAddress field to given value.

### HasShopWebAddress

`func (o *UpdateMerchantRequest) HasShopWebAddress() bool`

HasShopWebAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


