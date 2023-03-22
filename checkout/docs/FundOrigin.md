# FundOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**ShopperName** | Pointer to [**Name**](Name.md) |  | [optional] 

## Methods

### NewFundOrigin

`func NewFundOrigin() *FundOrigin`

NewFundOrigin instantiates a new FundOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundOriginWithDefaults

`func NewFundOriginWithDefaults() *FundOrigin`

NewFundOriginWithDefaults instantiates a new FundOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingAddress

`func (o *FundOrigin) GetBillingAddress() Address`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *FundOrigin) GetBillingAddressOk() (*Address, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *FundOrigin) SetBillingAddress(v Address)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *FundOrigin) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetShopperName

`func (o *FundOrigin) GetShopperName() Name`

GetShopperName returns the ShopperName field if non-nil, zero value otherwise.

### GetShopperNameOk

`func (o *FundOrigin) GetShopperNameOk() (*Name, bool)`

GetShopperNameOk returns a tuple with the ShopperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperName

`func (o *FundOrigin) SetShopperName(v Name)`

SetShopperName sets ShopperName field to given value.

### HasShopperName

`func (o *FundOrigin) HasShopperName() bool`

HasShopperName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


