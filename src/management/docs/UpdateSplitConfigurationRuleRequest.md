# UpdateSplitConfigurationRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | The currency condition that defines whether the split logic applies. Its value must be a three-character [ISO currency code](https://en.wikipedia.org/wiki/ISO_4217). | 
**FundingSource** | Pointer to **string** | The funding source condition of the payment method (only for cards).  Possible values: **credit**, **debit**, or **ANY**. | [optional] 
**PaymentMethod** | **string** | The payment method condition that defines whether the split logic applies.  Possible values: * [Payment method variant](https://docs.adyen.com/development-resources/paymentmethodvariant): Apply the split logic for a specific payment method. * **ANY**: Apply the split logic for all available payment methods. | 
**ShopperInteraction** | **string** | The sales channel condition that defines whether the split logic applies.  Possible values: * **Ecommerce**: Online transactions where the cardholder is present. * **ContAuth**: Card on file and/or subscription transactions, where the cardholder is known to the merchant (returning customer). * **Moto**: Mail-order and telephone-order transactions where the customer is in contact with the merchant via email or telephone. * **POS**: Point-of-sale transactions where the customer is physically present to make a payment using a secure payment terminal. * **ANY**: All sales channels. | 

## Methods

### NewUpdateSplitConfigurationRuleRequest

`func NewUpdateSplitConfigurationRuleRequest(currency string, paymentMethod string, shopperInteraction string, ) *UpdateSplitConfigurationRuleRequest`

NewUpdateSplitConfigurationRuleRequest instantiates a new UpdateSplitConfigurationRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSplitConfigurationRuleRequestWithDefaults

`func NewUpdateSplitConfigurationRuleRequestWithDefaults() *UpdateSplitConfigurationRuleRequest`

NewUpdateSplitConfigurationRuleRequestWithDefaults instantiates a new UpdateSplitConfigurationRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *UpdateSplitConfigurationRuleRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UpdateSplitConfigurationRuleRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UpdateSplitConfigurationRuleRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetFundingSource

`func (o *UpdateSplitConfigurationRuleRequest) GetFundingSource() string`

GetFundingSource returns the FundingSource field if non-nil, zero value otherwise.

### GetFundingSourceOk

`func (o *UpdateSplitConfigurationRuleRequest) GetFundingSourceOk() (*string, bool)`

GetFundingSourceOk returns a tuple with the FundingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSource

`func (o *UpdateSplitConfigurationRuleRequest) SetFundingSource(v string)`

SetFundingSource sets FundingSource field to given value.

### HasFundingSource

`func (o *UpdateSplitConfigurationRuleRequest) HasFundingSource() bool`

HasFundingSource returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *UpdateSplitConfigurationRuleRequest) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *UpdateSplitConfigurationRuleRequest) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *UpdateSplitConfigurationRuleRequest) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetShopperInteraction

`func (o *UpdateSplitConfigurationRuleRequest) GetShopperInteraction() string`

GetShopperInteraction returns the ShopperInteraction field if non-nil, zero value otherwise.

### GetShopperInteractionOk

`func (o *UpdateSplitConfigurationRuleRequest) GetShopperInteractionOk() (*string, bool)`

GetShopperInteractionOk returns a tuple with the ShopperInteraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperInteraction

`func (o *UpdateSplitConfigurationRuleRequest) SetShopperInteraction(v string)`

SetShopperInteraction sets ShopperInteraction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


