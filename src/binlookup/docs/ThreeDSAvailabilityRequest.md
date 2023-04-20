# ThreeDSAvailabilityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **map[string]string** | This field contains additional data, which may be required for a particular request.  The &#x60;additionalData&#x60; object consists of entries, each of which includes the key and value. | [optional] 
**Brands** | Pointer to **[]string** | List of brands. | [optional] 
**CardNumber** | Pointer to **string** | Card number or BIN. | [optional] 
**MerchantAccount** | **string** | The merchant account identifier. | 
**RecurringDetailReference** | Pointer to **string** | A recurring detail reference corresponding to a card. | [optional] 
**ShopperReference** | Pointer to **string** | The shopper&#39;s reference to uniquely identify this shopper (e.g. user ID or account ID). | [optional] 

## Methods

### NewThreeDSAvailabilityRequest

`func NewThreeDSAvailabilityRequest(merchantAccount string, ) *ThreeDSAvailabilityRequest`

NewThreeDSAvailabilityRequest instantiates a new ThreeDSAvailabilityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSAvailabilityRequestWithDefaults

`func NewThreeDSAvailabilityRequestWithDefaults() *ThreeDSAvailabilityRequest`

NewThreeDSAvailabilityRequestWithDefaults instantiates a new ThreeDSAvailabilityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *ThreeDSAvailabilityRequest) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *ThreeDSAvailabilityRequest) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *ThreeDSAvailabilityRequest) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *ThreeDSAvailabilityRequest) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetBrands

`func (o *ThreeDSAvailabilityRequest) GetBrands() []string`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *ThreeDSAvailabilityRequest) GetBrandsOk() (*[]string, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *ThreeDSAvailabilityRequest) SetBrands(v []string)`

SetBrands sets Brands field to given value.

### HasBrands

`func (o *ThreeDSAvailabilityRequest) HasBrands() bool`

HasBrands returns a boolean if a field has been set.

### GetCardNumber

`func (o *ThreeDSAvailabilityRequest) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *ThreeDSAvailabilityRequest) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *ThreeDSAvailabilityRequest) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *ThreeDSAvailabilityRequest) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *ThreeDSAvailabilityRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *ThreeDSAvailabilityRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *ThreeDSAvailabilityRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetRecurringDetailReference

`func (o *ThreeDSAvailabilityRequest) GetRecurringDetailReference() string`

GetRecurringDetailReference returns the RecurringDetailReference field if non-nil, zero value otherwise.

### GetRecurringDetailReferenceOk

`func (o *ThreeDSAvailabilityRequest) GetRecurringDetailReferenceOk() (*string, bool)`

GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDetailReference

`func (o *ThreeDSAvailabilityRequest) SetRecurringDetailReference(v string)`

SetRecurringDetailReference sets RecurringDetailReference field to given value.

### HasRecurringDetailReference

`func (o *ThreeDSAvailabilityRequest) HasRecurringDetailReference() bool`

HasRecurringDetailReference returns a boolean if a field has been set.

### GetShopperReference

`func (o *ThreeDSAvailabilityRequest) GetShopperReference() string`

GetShopperReference returns the ShopperReference field if non-nil, zero value otherwise.

### GetShopperReferenceOk

`func (o *ThreeDSAvailabilityRequest) GetShopperReferenceOk() (*string, bool)`

GetShopperReferenceOk returns a tuple with the ShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperReference

`func (o *ThreeDSAvailabilityRequest) SetShopperReference(v string)`

SetShopperReference sets ShopperReference field to given value.

### HasShopperReference

`func (o *ThreeDSAvailabilityRequest) HasShopperReference() bool`

HasShopperReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


