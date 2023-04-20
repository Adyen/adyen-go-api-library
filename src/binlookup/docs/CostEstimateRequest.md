# CostEstimateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | [**Amount**](Amount.md) |  | 
**Assumptions** | Pointer to [**CostEstimateAssumptions**](CostEstimateAssumptions.md) |  | [optional] 
**CardNumber** | Pointer to **string** | The card number (4-19 characters) for PCI compliant use cases. Do not use any separators.  &gt; Either the &#x60;cardNumber&#x60; or &#x60;encryptedCardNumber&#x60; field must be provided in a payment request. | [optional] 
**EncryptedCardNumber** | Pointer to **string** | Encrypted data that stores card information for non PCI-compliant use cases. The encrypted data must be created with the Checkout Card Component or Secured Fields Component, and must contain the &#x60;encryptedCardNumber&#x60; field.  &gt; Either the &#x60;cardNumber&#x60; or &#x60;encryptedCardNumber&#x60; field must be provided in a payment request. | [optional] 
**MerchantAccount** | **string** | The merchant account identifier you want to process the (transaction) request with. | 
**MerchantDetails** | Pointer to [**MerchantDetails**](MerchantDetails.md) |  | [optional] 
**Recurring** | Pointer to [**Recurring**](Recurring.md) |  | [optional] 
**SelectedRecurringDetailReference** | Pointer to **string** | The &#x60;recurringDetailReference&#x60; you want to use for this cost estimate. The value &#x60;LATEST&#x60; can be used to select the most recently stored recurring detail. | [optional] 
**ShopperInteraction** | Pointer to **string** | Specifies the sales channel, through which the shopper gives their card details, and whether the shopper is a returning customer. For the web service API, Adyen assumes Ecommerce shopper interaction by default.  This field has the following possible values: * &#x60;Ecommerce&#x60; - Online transactions where the cardholder is present (online). For better authorisation rates, we recommend sending the card security code (CSC) along with the request. * &#x60;ContAuth&#x60; - Card on file and/or subscription transactions, where the card holder is known to the merchant (returning customer). If the shopper is present (online), you can supply also the CSC to improve authorisation (one-click payment). * &#x60;Moto&#x60; - Mail-order and telephone-order transactions where the shopper is in contact with the merchant via email or telephone. * &#x60;POS&#x60; - Point-of-sale transactions where the shopper is physically present to make a payment using a secure payment terminal. | [optional] 
**ShopperReference** | Pointer to **string** | Required for recurring payments.  Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. &gt; Your reference must not include personally identifiable information (PII), for example name or email address. | [optional] 

## Methods

### NewCostEstimateRequest

`func NewCostEstimateRequest(amount Amount, merchantAccount string, ) *CostEstimateRequest`

NewCostEstimateRequest instantiates a new CostEstimateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostEstimateRequestWithDefaults

`func NewCostEstimateRequestWithDefaults() *CostEstimateRequest`

NewCostEstimateRequestWithDefaults instantiates a new CostEstimateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CostEstimateRequest) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CostEstimateRequest) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CostEstimateRequest) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetAssumptions

`func (o *CostEstimateRequest) GetAssumptions() CostEstimateAssumptions`

GetAssumptions returns the Assumptions field if non-nil, zero value otherwise.

### GetAssumptionsOk

`func (o *CostEstimateRequest) GetAssumptionsOk() (*CostEstimateAssumptions, bool)`

GetAssumptionsOk returns a tuple with the Assumptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumptions

`func (o *CostEstimateRequest) SetAssumptions(v CostEstimateAssumptions)`

SetAssumptions sets Assumptions field to given value.

### HasAssumptions

`func (o *CostEstimateRequest) HasAssumptions() bool`

HasAssumptions returns a boolean if a field has been set.

### GetCardNumber

`func (o *CostEstimateRequest) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *CostEstimateRequest) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *CostEstimateRequest) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *CostEstimateRequest) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetEncryptedCardNumber

`func (o *CostEstimateRequest) GetEncryptedCardNumber() string`

GetEncryptedCardNumber returns the EncryptedCardNumber field if non-nil, zero value otherwise.

### GetEncryptedCardNumberOk

`func (o *CostEstimateRequest) GetEncryptedCardNumberOk() (*string, bool)`

GetEncryptedCardNumberOk returns a tuple with the EncryptedCardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedCardNumber

`func (o *CostEstimateRequest) SetEncryptedCardNumber(v string)`

SetEncryptedCardNumber sets EncryptedCardNumber field to given value.

### HasEncryptedCardNumber

`func (o *CostEstimateRequest) HasEncryptedCardNumber() bool`

HasEncryptedCardNumber returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *CostEstimateRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *CostEstimateRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *CostEstimateRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetMerchantDetails

`func (o *CostEstimateRequest) GetMerchantDetails() MerchantDetails`

GetMerchantDetails returns the MerchantDetails field if non-nil, zero value otherwise.

### GetMerchantDetailsOk

`func (o *CostEstimateRequest) GetMerchantDetailsOk() (*MerchantDetails, bool)`

GetMerchantDetailsOk returns a tuple with the MerchantDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantDetails

`func (o *CostEstimateRequest) SetMerchantDetails(v MerchantDetails)`

SetMerchantDetails sets MerchantDetails field to given value.

### HasMerchantDetails

`func (o *CostEstimateRequest) HasMerchantDetails() bool`

HasMerchantDetails returns a boolean if a field has been set.

### GetRecurring

`func (o *CostEstimateRequest) GetRecurring() Recurring`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *CostEstimateRequest) GetRecurringOk() (*Recurring, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *CostEstimateRequest) SetRecurring(v Recurring)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *CostEstimateRequest) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.

### GetSelectedRecurringDetailReference

`func (o *CostEstimateRequest) GetSelectedRecurringDetailReference() string`

GetSelectedRecurringDetailReference returns the SelectedRecurringDetailReference field if non-nil, zero value otherwise.

### GetSelectedRecurringDetailReferenceOk

`func (o *CostEstimateRequest) GetSelectedRecurringDetailReferenceOk() (*string, bool)`

GetSelectedRecurringDetailReferenceOk returns a tuple with the SelectedRecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRecurringDetailReference

`func (o *CostEstimateRequest) SetSelectedRecurringDetailReference(v string)`

SetSelectedRecurringDetailReference sets SelectedRecurringDetailReference field to given value.

### HasSelectedRecurringDetailReference

`func (o *CostEstimateRequest) HasSelectedRecurringDetailReference() bool`

HasSelectedRecurringDetailReference returns a boolean if a field has been set.

### GetShopperInteraction

`func (o *CostEstimateRequest) GetShopperInteraction() string`

GetShopperInteraction returns the ShopperInteraction field if non-nil, zero value otherwise.

### GetShopperInteractionOk

`func (o *CostEstimateRequest) GetShopperInteractionOk() (*string, bool)`

GetShopperInteractionOk returns a tuple with the ShopperInteraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperInteraction

`func (o *CostEstimateRequest) SetShopperInteraction(v string)`

SetShopperInteraction sets ShopperInteraction field to given value.

### HasShopperInteraction

`func (o *CostEstimateRequest) HasShopperInteraction() bool`

HasShopperInteraction returns a boolean if a field has been set.

### GetShopperReference

`func (o *CostEstimateRequest) GetShopperReference() string`

GetShopperReference returns the ShopperReference field if non-nil, zero value otherwise.

### GetShopperReferenceOk

`func (o *CostEstimateRequest) GetShopperReferenceOk() (*string, bool)`

GetShopperReferenceOk returns a tuple with the ShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperReference

`func (o *CostEstimateRequest) SetShopperReference(v string)`

SetShopperReference sets ShopperReference field to given value.

### HasShopperReference

`func (o *CostEstimateRequest) HasShopperReference() bool`

HasShopperReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


