# StoredPaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to **string** | The brand of the card. | [optional] 
**ExpiryMonth** | Pointer to **string** | The month the card expires. | [optional] 
**ExpiryYear** | Pointer to **string** | The last two digits of the year the card expires. For example, **22** for the year 2022. | [optional] 
**HolderName** | Pointer to **string** | The unique payment method code. | [optional] 
**Iban** | Pointer to **string** | The IBAN of the bank account. | [optional] 
**Id** | Pointer to **string** | A unique identifier of this stored payment method. | [optional] 
**LastFour** | Pointer to **string** | The last four digits of the PAN. | [optional] 
**Name** | Pointer to **string** | The display name of the stored payment method. | [optional] 
**NetworkTxReference** | Pointer to **string** | Returned in the response if you are not tokenizing with Adyen and are using the Merchant-initiated transactions (MIT) framework from Mastercard or Visa.  This contains either the Mastercard Trace ID or the Visa Transaction ID. | [optional] 
**OwnerName** | Pointer to **string** | The name of the bank account holder. | [optional] 
**ShopperEmail** | Pointer to **string** | The shopper’s email address. | [optional] 
**SupportedRecurringProcessingModels** | Pointer to **[]string** | The supported recurring processing models for this stored payment method. | [optional] 
**SupportedShopperInteractions** | Pointer to **[]string** | The supported shopper interactions for this stored payment method. | [optional] 
**Type** | Pointer to **string** | The type of payment method. | [optional] 

## Methods

### NewStoredPaymentMethod

`func NewStoredPaymentMethod() *StoredPaymentMethod`

NewStoredPaymentMethod instantiates a new StoredPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoredPaymentMethodWithDefaults

`func NewStoredPaymentMethodWithDefaults() *StoredPaymentMethod`

NewStoredPaymentMethodWithDefaults instantiates a new StoredPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *StoredPaymentMethod) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *StoredPaymentMethod) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *StoredPaymentMethod) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *StoredPaymentMethod) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetExpiryMonth

`func (o *StoredPaymentMethod) GetExpiryMonth() string`

GetExpiryMonth returns the ExpiryMonth field if non-nil, zero value otherwise.

### GetExpiryMonthOk

`func (o *StoredPaymentMethod) GetExpiryMonthOk() (*string, bool)`

GetExpiryMonthOk returns a tuple with the ExpiryMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryMonth

`func (o *StoredPaymentMethod) SetExpiryMonth(v string)`

SetExpiryMonth sets ExpiryMonth field to given value.

### HasExpiryMonth

`func (o *StoredPaymentMethod) HasExpiryMonth() bool`

HasExpiryMonth returns a boolean if a field has been set.

### GetExpiryYear

`func (o *StoredPaymentMethod) GetExpiryYear() string`

GetExpiryYear returns the ExpiryYear field if non-nil, zero value otherwise.

### GetExpiryYearOk

`func (o *StoredPaymentMethod) GetExpiryYearOk() (*string, bool)`

GetExpiryYearOk returns a tuple with the ExpiryYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryYear

`func (o *StoredPaymentMethod) SetExpiryYear(v string)`

SetExpiryYear sets ExpiryYear field to given value.

### HasExpiryYear

`func (o *StoredPaymentMethod) HasExpiryYear() bool`

HasExpiryYear returns a boolean if a field has been set.

### GetHolderName

`func (o *StoredPaymentMethod) GetHolderName() string`

GetHolderName returns the HolderName field if non-nil, zero value otherwise.

### GetHolderNameOk

`func (o *StoredPaymentMethod) GetHolderNameOk() (*string, bool)`

GetHolderNameOk returns a tuple with the HolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderName

`func (o *StoredPaymentMethod) SetHolderName(v string)`

SetHolderName sets HolderName field to given value.

### HasHolderName

`func (o *StoredPaymentMethod) HasHolderName() bool`

HasHolderName returns a boolean if a field has been set.

### GetIban

`func (o *StoredPaymentMethod) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *StoredPaymentMethod) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *StoredPaymentMethod) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *StoredPaymentMethod) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetId

`func (o *StoredPaymentMethod) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StoredPaymentMethod) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StoredPaymentMethod) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StoredPaymentMethod) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastFour

`func (o *StoredPaymentMethod) GetLastFour() string`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *StoredPaymentMethod) GetLastFourOk() (*string, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *StoredPaymentMethod) SetLastFour(v string)`

SetLastFour sets LastFour field to given value.

### HasLastFour

`func (o *StoredPaymentMethod) HasLastFour() bool`

HasLastFour returns a boolean if a field has been set.

### GetName

`func (o *StoredPaymentMethod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoredPaymentMethod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoredPaymentMethod) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoredPaymentMethod) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkTxReference

`func (o *StoredPaymentMethod) GetNetworkTxReference() string`

GetNetworkTxReference returns the NetworkTxReference field if non-nil, zero value otherwise.

### GetNetworkTxReferenceOk

`func (o *StoredPaymentMethod) GetNetworkTxReferenceOk() (*string, bool)`

GetNetworkTxReferenceOk returns a tuple with the NetworkTxReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTxReference

`func (o *StoredPaymentMethod) SetNetworkTxReference(v string)`

SetNetworkTxReference sets NetworkTxReference field to given value.

### HasNetworkTxReference

`func (o *StoredPaymentMethod) HasNetworkTxReference() bool`

HasNetworkTxReference returns a boolean if a field has been set.

### GetOwnerName

`func (o *StoredPaymentMethod) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *StoredPaymentMethod) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *StoredPaymentMethod) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *StoredPaymentMethod) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### GetShopperEmail

`func (o *StoredPaymentMethod) GetShopperEmail() string`

GetShopperEmail returns the ShopperEmail field if non-nil, zero value otherwise.

### GetShopperEmailOk

`func (o *StoredPaymentMethod) GetShopperEmailOk() (*string, bool)`

GetShopperEmailOk returns a tuple with the ShopperEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperEmail

`func (o *StoredPaymentMethod) SetShopperEmail(v string)`

SetShopperEmail sets ShopperEmail field to given value.

### HasShopperEmail

`func (o *StoredPaymentMethod) HasShopperEmail() bool`

HasShopperEmail returns a boolean if a field has been set.

### GetSupportedRecurringProcessingModels

`func (o *StoredPaymentMethod) GetSupportedRecurringProcessingModels() []string`

GetSupportedRecurringProcessingModels returns the SupportedRecurringProcessingModels field if non-nil, zero value otherwise.

### GetSupportedRecurringProcessingModelsOk

`func (o *StoredPaymentMethod) GetSupportedRecurringProcessingModelsOk() (*[]string, bool)`

GetSupportedRecurringProcessingModelsOk returns a tuple with the SupportedRecurringProcessingModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedRecurringProcessingModels

`func (o *StoredPaymentMethod) SetSupportedRecurringProcessingModels(v []string)`

SetSupportedRecurringProcessingModels sets SupportedRecurringProcessingModels field to given value.

### HasSupportedRecurringProcessingModels

`func (o *StoredPaymentMethod) HasSupportedRecurringProcessingModels() bool`

HasSupportedRecurringProcessingModels returns a boolean if a field has been set.

### GetSupportedShopperInteractions

`func (o *StoredPaymentMethod) GetSupportedShopperInteractions() []string`

GetSupportedShopperInteractions returns the SupportedShopperInteractions field if non-nil, zero value otherwise.

### GetSupportedShopperInteractionsOk

`func (o *StoredPaymentMethod) GetSupportedShopperInteractionsOk() (*[]string, bool)`

GetSupportedShopperInteractionsOk returns a tuple with the SupportedShopperInteractions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedShopperInteractions

`func (o *StoredPaymentMethod) SetSupportedShopperInteractions(v []string)`

SetSupportedShopperInteractions sets SupportedShopperInteractions field to given value.

### HasSupportedShopperInteractions

`func (o *StoredPaymentMethod) HasSupportedShopperInteractions() bool`

HasSupportedShopperInteractions returns a boolean if a field has been set.

### GetType

`func (o *StoredPaymentMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StoredPaymentMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StoredPaymentMethod) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StoredPaymentMethod) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


