# Merchant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**MerchantLinks**](MerchantLinks.md) |  | [optional] 
**CaptureDelay** | Pointer to **string** | The [capture delay](https://docs.adyen.com/online-payments/capture#capture-delay) set for the merchant account.  Possible values: * **Immediate** * **Manual** * Number of days from **1** to **29** | [optional] 
**DefaultShopperInteraction** | Pointer to **string** | The default [&#x60;shopperInteraction&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/v68/post/payments__reqParam_shopperInteraction) value used when processing payments through this merchant account. | [optional] 
**Id** | Pointer to **string** | The unique identifier of the merchant account. | [optional] 
**MerchantCity** | Pointer to **string** | The city where the legal entity of this merchant account is registered. | [optional] 
**Name** | Pointer to **string** | The name of the legal entity associated with the merchant account. | [optional] 
**PricingPlan** | Pointer to **string** | Only applies to merchant accounts managed by Adyen&#39;s partners. The name of the pricing plan assigned to the merchant account. | [optional] 
**PrimarySettlementCurrency** | Pointer to **string** | The currency of the country where the legal entity of this merchant account is registered. Format: [ISO currency code](https://docs.adyen.com/development-resources/currency-codes). For example, a legal entity based in the United States has USD as the primary settlement currency. | [optional] 
**ShopWebAddress** | Pointer to **string** | The URL for the ecommerce website used with this merchant account. | [optional] 
**Status** | Pointer to **string** | The status of the merchant account.  Possible values:  * **PreActive**: The merchant account has been created. Users cannot access the merchant account in the Customer Area. The account cannot process payments. * **Active**: Users can access the merchant account in the Customer Area. If the company account is also **Active**, then payment processing and payouts are enabled. * **InactiveWithModifications**: Users can access the merchant account in the Customer Area. You cannot process new payments but you can still modify payments, for example issue refunds. You can still receive payouts. * **Inactive**: Users can access the merchant account in the Customer Area. Payment processing and payouts are disabled. * **Closed**: The account is closed and this cannot be reversed. Users cannot log in. Payment processing and payouts are disabled. | [optional] 

## Methods

### NewMerchant

`func NewMerchant() *Merchant`

NewMerchant instantiates a new Merchant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantWithDefaults

`func NewMerchantWithDefaults() *Merchant`

NewMerchantWithDefaults instantiates a new Merchant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Merchant) GetLinks() MerchantLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Merchant) GetLinksOk() (*MerchantLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Merchant) SetLinks(v MerchantLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Merchant) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCaptureDelay

`func (o *Merchant) GetCaptureDelay() string`

GetCaptureDelay returns the CaptureDelay field if non-nil, zero value otherwise.

### GetCaptureDelayOk

`func (o *Merchant) GetCaptureDelayOk() (*string, bool)`

GetCaptureDelayOk returns a tuple with the CaptureDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureDelay

`func (o *Merchant) SetCaptureDelay(v string)`

SetCaptureDelay sets CaptureDelay field to given value.

### HasCaptureDelay

`func (o *Merchant) HasCaptureDelay() bool`

HasCaptureDelay returns a boolean if a field has been set.

### GetDefaultShopperInteraction

`func (o *Merchant) GetDefaultShopperInteraction() string`

GetDefaultShopperInteraction returns the DefaultShopperInteraction field if non-nil, zero value otherwise.

### GetDefaultShopperInteractionOk

`func (o *Merchant) GetDefaultShopperInteractionOk() (*string, bool)`

GetDefaultShopperInteractionOk returns a tuple with the DefaultShopperInteraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultShopperInteraction

`func (o *Merchant) SetDefaultShopperInteraction(v string)`

SetDefaultShopperInteraction sets DefaultShopperInteraction field to given value.

### HasDefaultShopperInteraction

`func (o *Merchant) HasDefaultShopperInteraction() bool`

HasDefaultShopperInteraction returns a boolean if a field has been set.

### GetId

`func (o *Merchant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Merchant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Merchant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Merchant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantCity

`func (o *Merchant) GetMerchantCity() string`

GetMerchantCity returns the MerchantCity field if non-nil, zero value otherwise.

### GetMerchantCityOk

`func (o *Merchant) GetMerchantCityOk() (*string, bool)`

GetMerchantCityOk returns a tuple with the MerchantCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCity

`func (o *Merchant) SetMerchantCity(v string)`

SetMerchantCity sets MerchantCity field to given value.

### HasMerchantCity

`func (o *Merchant) HasMerchantCity() bool`

HasMerchantCity returns a boolean if a field has been set.

### GetName

`func (o *Merchant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Merchant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Merchant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Merchant) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPricingPlan

`func (o *Merchant) GetPricingPlan() string`

GetPricingPlan returns the PricingPlan field if non-nil, zero value otherwise.

### GetPricingPlanOk

`func (o *Merchant) GetPricingPlanOk() (*string, bool)`

GetPricingPlanOk returns a tuple with the PricingPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingPlan

`func (o *Merchant) SetPricingPlan(v string)`

SetPricingPlan sets PricingPlan field to given value.

### HasPricingPlan

`func (o *Merchant) HasPricingPlan() bool`

HasPricingPlan returns a boolean if a field has been set.

### GetPrimarySettlementCurrency

`func (o *Merchant) GetPrimarySettlementCurrency() string`

GetPrimarySettlementCurrency returns the PrimarySettlementCurrency field if non-nil, zero value otherwise.

### GetPrimarySettlementCurrencyOk

`func (o *Merchant) GetPrimarySettlementCurrencyOk() (*string, bool)`

GetPrimarySettlementCurrencyOk returns a tuple with the PrimarySettlementCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySettlementCurrency

`func (o *Merchant) SetPrimarySettlementCurrency(v string)`

SetPrimarySettlementCurrency sets PrimarySettlementCurrency field to given value.

### HasPrimarySettlementCurrency

`func (o *Merchant) HasPrimarySettlementCurrency() bool`

HasPrimarySettlementCurrency returns a boolean if a field has been set.

### GetShopWebAddress

`func (o *Merchant) GetShopWebAddress() string`

GetShopWebAddress returns the ShopWebAddress field if non-nil, zero value otherwise.

### GetShopWebAddressOk

`func (o *Merchant) GetShopWebAddressOk() (*string, bool)`

GetShopWebAddressOk returns a tuple with the ShopWebAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopWebAddress

`func (o *Merchant) SetShopWebAddress(v string)`

SetShopWebAddress sets ShopWebAddress field to given value.

### HasShopWebAddress

`func (o *Merchant) HasShopWebAddress() bool`

HasShopWebAddress returns a boolean if a field has been set.

### GetStatus

`func (o *Merchant) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Merchant) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Merchant) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Merchant) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


