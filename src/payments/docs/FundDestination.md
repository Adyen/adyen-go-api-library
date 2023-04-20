# FundDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **map[string]string** | a map of name/value pairs for passing in additional/industry-specific data | [optional] 
**BillingAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**Card** | Pointer to [**Card**](Card.md) |  | [optional] 
**SelectedRecurringDetailReference** | Pointer to **string** | The &#x60;recurringDetailReference&#x60; you want to use for this payment. The value &#x60;LATEST&#x60; can be used to select the most recently stored recurring detail. | [optional] 
**ShopperEmail** | Pointer to **string** | the email address of the person | [optional] 
**ShopperName** | Pointer to [**Name**](Name.md) |  | [optional] 
**ShopperReference** | Pointer to **string** | Required for recurring payments.  Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. &gt; Your reference must not include personally identifiable information (PII), for example name or email address. | [optional] 
**SubMerchant** | Pointer to [**SubMerchant**](SubMerchant.md) |  | [optional] 
**TelephoneNumber** | Pointer to **string** | the telephone number of the person | [optional] 

## Methods

### NewFundDestination

`func NewFundDestination() *FundDestination`

NewFundDestination instantiates a new FundDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundDestinationWithDefaults

`func NewFundDestinationWithDefaults() *FundDestination`

NewFundDestinationWithDefaults instantiates a new FundDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *FundDestination) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *FundDestination) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *FundDestination) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *FundDestination) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetBillingAddress

`func (o *FundDestination) GetBillingAddress() Address`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *FundDestination) GetBillingAddressOk() (*Address, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *FundDestination) SetBillingAddress(v Address)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *FundDestination) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetCard

`func (o *FundDestination) GetCard() Card`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *FundDestination) GetCardOk() (*Card, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *FundDestination) SetCard(v Card)`

SetCard sets Card field to given value.

### HasCard

`func (o *FundDestination) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetSelectedRecurringDetailReference

`func (o *FundDestination) GetSelectedRecurringDetailReference() string`

GetSelectedRecurringDetailReference returns the SelectedRecurringDetailReference field if non-nil, zero value otherwise.

### GetSelectedRecurringDetailReferenceOk

`func (o *FundDestination) GetSelectedRecurringDetailReferenceOk() (*string, bool)`

GetSelectedRecurringDetailReferenceOk returns a tuple with the SelectedRecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRecurringDetailReference

`func (o *FundDestination) SetSelectedRecurringDetailReference(v string)`

SetSelectedRecurringDetailReference sets SelectedRecurringDetailReference field to given value.

### HasSelectedRecurringDetailReference

`func (o *FundDestination) HasSelectedRecurringDetailReference() bool`

HasSelectedRecurringDetailReference returns a boolean if a field has been set.

### GetShopperEmail

`func (o *FundDestination) GetShopperEmail() string`

GetShopperEmail returns the ShopperEmail field if non-nil, zero value otherwise.

### GetShopperEmailOk

`func (o *FundDestination) GetShopperEmailOk() (*string, bool)`

GetShopperEmailOk returns a tuple with the ShopperEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperEmail

`func (o *FundDestination) SetShopperEmail(v string)`

SetShopperEmail sets ShopperEmail field to given value.

### HasShopperEmail

`func (o *FundDestination) HasShopperEmail() bool`

HasShopperEmail returns a boolean if a field has been set.

### GetShopperName

`func (o *FundDestination) GetShopperName() Name`

GetShopperName returns the ShopperName field if non-nil, zero value otherwise.

### GetShopperNameOk

`func (o *FundDestination) GetShopperNameOk() (*Name, bool)`

GetShopperNameOk returns a tuple with the ShopperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperName

`func (o *FundDestination) SetShopperName(v Name)`

SetShopperName sets ShopperName field to given value.

### HasShopperName

`func (o *FundDestination) HasShopperName() bool`

HasShopperName returns a boolean if a field has been set.

### GetShopperReference

`func (o *FundDestination) GetShopperReference() string`

GetShopperReference returns the ShopperReference field if non-nil, zero value otherwise.

### GetShopperReferenceOk

`func (o *FundDestination) GetShopperReferenceOk() (*string, bool)`

GetShopperReferenceOk returns a tuple with the ShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperReference

`func (o *FundDestination) SetShopperReference(v string)`

SetShopperReference sets ShopperReference field to given value.

### HasShopperReference

`func (o *FundDestination) HasShopperReference() bool`

HasShopperReference returns a boolean if a field has been set.

### GetSubMerchant

`func (o *FundDestination) GetSubMerchant() SubMerchant`

GetSubMerchant returns the SubMerchant field if non-nil, zero value otherwise.

### GetSubMerchantOk

`func (o *FundDestination) GetSubMerchantOk() (*SubMerchant, bool)`

GetSubMerchantOk returns a tuple with the SubMerchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchant

`func (o *FundDestination) SetSubMerchant(v SubMerchant)`

SetSubMerchant sets SubMerchant field to given value.

### HasSubMerchant

`func (o *FundDestination) HasSubMerchant() bool`

HasSubMerchant returns a boolean if a field has been set.

### GetTelephoneNumber

`func (o *FundDestination) GetTelephoneNumber() string`

GetTelephoneNumber returns the TelephoneNumber field if non-nil, zero value otherwise.

### GetTelephoneNumberOk

`func (o *FundDestination) GetTelephoneNumberOk() (*string, bool)`

GetTelephoneNumberOk returns a tuple with the TelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumber

`func (o *FundDestination) SetTelephoneNumber(v string)`

SetTelephoneNumber sets TelephoneNumber field to given value.

### HasTelephoneNumber

`func (o *FundDestination) HasTelephoneNumber() bool`

HasTelephoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


