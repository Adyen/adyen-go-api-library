# FundSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **map[string]string** | A map of name-value pairs for passing additional or industry-specific data. | [optional] 
**BillingAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**Card** | Pointer to [**Card**](Card.md) |  | [optional] 
**ShopperEmail** | Pointer to **string** | Email address of the person. | [optional] 
**ShopperName** | Pointer to [**Name**](Name.md) |  | [optional] 
**TelephoneNumber** | Pointer to **string** | Phone number of the person | [optional] 

## Methods

### NewFundSource

`func NewFundSource() *FundSource`

NewFundSource instantiates a new FundSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundSourceWithDefaults

`func NewFundSourceWithDefaults() *FundSource`

NewFundSourceWithDefaults instantiates a new FundSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *FundSource) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *FundSource) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *FundSource) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *FundSource) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetBillingAddress

`func (o *FundSource) GetBillingAddress() Address`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *FundSource) GetBillingAddressOk() (*Address, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *FundSource) SetBillingAddress(v Address)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *FundSource) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetCard

`func (o *FundSource) GetCard() Card`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *FundSource) GetCardOk() (*Card, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *FundSource) SetCard(v Card)`

SetCard sets Card field to given value.

### HasCard

`func (o *FundSource) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetShopperEmail

`func (o *FundSource) GetShopperEmail() string`

GetShopperEmail returns the ShopperEmail field if non-nil, zero value otherwise.

### GetShopperEmailOk

`func (o *FundSource) GetShopperEmailOk() (*string, bool)`

GetShopperEmailOk returns a tuple with the ShopperEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperEmail

`func (o *FundSource) SetShopperEmail(v string)`

SetShopperEmail sets ShopperEmail field to given value.

### HasShopperEmail

`func (o *FundSource) HasShopperEmail() bool`

HasShopperEmail returns a boolean if a field has been set.

### GetShopperName

`func (o *FundSource) GetShopperName() Name`

GetShopperName returns the ShopperName field if non-nil, zero value otherwise.

### GetShopperNameOk

`func (o *FundSource) GetShopperNameOk() (*Name, bool)`

GetShopperNameOk returns a tuple with the ShopperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperName

`func (o *FundSource) SetShopperName(v Name)`

SetShopperName sets ShopperName field to given value.

### HasShopperName

`func (o *FundSource) HasShopperName() bool`

HasShopperName returns a boolean if a field has been set.

### GetTelephoneNumber

`func (o *FundSource) GetTelephoneNumber() string`

GetTelephoneNumber returns the TelephoneNumber field if non-nil, zero value otherwise.

### GetTelephoneNumberOk

`func (o *FundSource) GetTelephoneNumberOk() (*string, bool)`

GetTelephoneNumberOk returns a tuple with the TelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumber

`func (o *FundSource) SetTelephoneNumber(v string)`

SetTelephoneNumber sets TelephoneNumber field to given value.

### HasTelephoneNumber

`func (o *FundSource) HasTelephoneNumber() bool`

HasTelephoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


