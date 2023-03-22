# StoredPaymentMethodResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to **string** | The brand of the card. | [optional] 
**ExpiryMonth** | Pointer to **string** | The month the card expires. | [optional] 
**ExpiryYear** | Pointer to **string** | The last two digits of the year the card expires. For example, **22** for the year 2022. | [optional] 
**ExternalResponseCode** | Pointer to **string** | The response code returned by an external system (for example after a provisioning operation). | [optional] 
**ExternalTokenReference** | Pointer to **string** | The token reference of a linked token in an external system (for example a network token reference). | [optional] 
**HolderName** | Pointer to **string** | The unique payment method code. | [optional] 
**Iban** | Pointer to **string** | The IBAN of the bank account. | [optional] 
**Id** | Pointer to **string** | A unique identifier of this stored payment method. | [optional] 
**IssuerName** | Pointer to **string** | The name of the issuer of token or card. | [optional] 
**LastFour** | Pointer to **string** | The last four digits of the PAN. | [optional] 
**Name** | Pointer to **string** | The display name of the stored payment method. | [optional] 
**NetworkTxReference** | Pointer to **string** | Returned in the response if you are not tokenizing with Adyen and are using the Merchant-initiated transactions (MIT) framework from Mastercard or Visa.  This contains either the Mastercard Trace ID or the Visa Transaction ID. | [optional] 
**OwnerName** | Pointer to **string** | The name of the bank account holder. | [optional] 
**ShopperEmail** | Pointer to **string** | The shopper’s email address. | [optional] 
**ShopperReference** | Pointer to **string** | Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. &gt; Your reference must not include personally identifiable information (PII), for example name or email address. | [optional] 
**SupportedRecurringProcessingModels** | Pointer to **[]string** | Defines a recurring payment type. Allowed values: * &#x60;Subscription&#x60; – A transaction for a fixed or variable amount, which follows a fixed schedule. * &#x60;CardOnFile&#x60; – With a card-on-file (CoF) transaction, card details are stored to enable one-click or omnichannel journeys, or simply to streamline the checkout process. Any subscription not following a fixed schedule is also considered a card-on-file transaction. * &#x60;UnscheduledCardOnFile&#x60; – An unscheduled card-on-file (UCoF) transaction is a transaction that occurs on a non-fixed schedule and/or have variable amounts. For example, automatic top-ups when a cardholder&#39;s balance drops below a certain amount. | [optional] 
**Type** | Pointer to **string** | The type of payment method. | [optional] 

## Methods

### NewStoredPaymentMethodResource

`func NewStoredPaymentMethodResource() *StoredPaymentMethodResource`

NewStoredPaymentMethodResource instantiates a new StoredPaymentMethodResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoredPaymentMethodResourceWithDefaults

`func NewStoredPaymentMethodResourceWithDefaults() *StoredPaymentMethodResource`

NewStoredPaymentMethodResourceWithDefaults instantiates a new StoredPaymentMethodResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *StoredPaymentMethodResource) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *StoredPaymentMethodResource) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *StoredPaymentMethodResource) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *StoredPaymentMethodResource) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetExpiryMonth

`func (o *StoredPaymentMethodResource) GetExpiryMonth() string`

GetExpiryMonth returns the ExpiryMonth field if non-nil, zero value otherwise.

### GetExpiryMonthOk

`func (o *StoredPaymentMethodResource) GetExpiryMonthOk() (*string, bool)`

GetExpiryMonthOk returns a tuple with the ExpiryMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryMonth

`func (o *StoredPaymentMethodResource) SetExpiryMonth(v string)`

SetExpiryMonth sets ExpiryMonth field to given value.

### HasExpiryMonth

`func (o *StoredPaymentMethodResource) HasExpiryMonth() bool`

HasExpiryMonth returns a boolean if a field has been set.

### GetExpiryYear

`func (o *StoredPaymentMethodResource) GetExpiryYear() string`

GetExpiryYear returns the ExpiryYear field if non-nil, zero value otherwise.

### GetExpiryYearOk

`func (o *StoredPaymentMethodResource) GetExpiryYearOk() (*string, bool)`

GetExpiryYearOk returns a tuple with the ExpiryYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryYear

`func (o *StoredPaymentMethodResource) SetExpiryYear(v string)`

SetExpiryYear sets ExpiryYear field to given value.

### HasExpiryYear

`func (o *StoredPaymentMethodResource) HasExpiryYear() bool`

HasExpiryYear returns a boolean if a field has been set.

### GetExternalResponseCode

`func (o *StoredPaymentMethodResource) GetExternalResponseCode() string`

GetExternalResponseCode returns the ExternalResponseCode field if non-nil, zero value otherwise.

### GetExternalResponseCodeOk

`func (o *StoredPaymentMethodResource) GetExternalResponseCodeOk() (*string, bool)`

GetExternalResponseCodeOk returns a tuple with the ExternalResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalResponseCode

`func (o *StoredPaymentMethodResource) SetExternalResponseCode(v string)`

SetExternalResponseCode sets ExternalResponseCode field to given value.

### HasExternalResponseCode

`func (o *StoredPaymentMethodResource) HasExternalResponseCode() bool`

HasExternalResponseCode returns a boolean if a field has been set.

### GetExternalTokenReference

`func (o *StoredPaymentMethodResource) GetExternalTokenReference() string`

GetExternalTokenReference returns the ExternalTokenReference field if non-nil, zero value otherwise.

### GetExternalTokenReferenceOk

`func (o *StoredPaymentMethodResource) GetExternalTokenReferenceOk() (*string, bool)`

GetExternalTokenReferenceOk returns a tuple with the ExternalTokenReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTokenReference

`func (o *StoredPaymentMethodResource) SetExternalTokenReference(v string)`

SetExternalTokenReference sets ExternalTokenReference field to given value.

### HasExternalTokenReference

`func (o *StoredPaymentMethodResource) HasExternalTokenReference() bool`

HasExternalTokenReference returns a boolean if a field has been set.

### GetHolderName

`func (o *StoredPaymentMethodResource) GetHolderName() string`

GetHolderName returns the HolderName field if non-nil, zero value otherwise.

### GetHolderNameOk

`func (o *StoredPaymentMethodResource) GetHolderNameOk() (*string, bool)`

GetHolderNameOk returns a tuple with the HolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderName

`func (o *StoredPaymentMethodResource) SetHolderName(v string)`

SetHolderName sets HolderName field to given value.

### HasHolderName

`func (o *StoredPaymentMethodResource) HasHolderName() bool`

HasHolderName returns a boolean if a field has been set.

### GetIban

`func (o *StoredPaymentMethodResource) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *StoredPaymentMethodResource) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *StoredPaymentMethodResource) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *StoredPaymentMethodResource) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetId

`func (o *StoredPaymentMethodResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StoredPaymentMethodResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StoredPaymentMethodResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StoredPaymentMethodResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssuerName

`func (o *StoredPaymentMethodResource) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *StoredPaymentMethodResource) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *StoredPaymentMethodResource) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.

### HasIssuerName

`func (o *StoredPaymentMethodResource) HasIssuerName() bool`

HasIssuerName returns a boolean if a field has been set.

### GetLastFour

`func (o *StoredPaymentMethodResource) GetLastFour() string`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *StoredPaymentMethodResource) GetLastFourOk() (*string, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *StoredPaymentMethodResource) SetLastFour(v string)`

SetLastFour sets LastFour field to given value.

### HasLastFour

`func (o *StoredPaymentMethodResource) HasLastFour() bool`

HasLastFour returns a boolean if a field has been set.

### GetName

`func (o *StoredPaymentMethodResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoredPaymentMethodResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoredPaymentMethodResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoredPaymentMethodResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkTxReference

`func (o *StoredPaymentMethodResource) GetNetworkTxReference() string`

GetNetworkTxReference returns the NetworkTxReference field if non-nil, zero value otherwise.

### GetNetworkTxReferenceOk

`func (o *StoredPaymentMethodResource) GetNetworkTxReferenceOk() (*string, bool)`

GetNetworkTxReferenceOk returns a tuple with the NetworkTxReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTxReference

`func (o *StoredPaymentMethodResource) SetNetworkTxReference(v string)`

SetNetworkTxReference sets NetworkTxReference field to given value.

### HasNetworkTxReference

`func (o *StoredPaymentMethodResource) HasNetworkTxReference() bool`

HasNetworkTxReference returns a boolean if a field has been set.

### GetOwnerName

`func (o *StoredPaymentMethodResource) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *StoredPaymentMethodResource) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *StoredPaymentMethodResource) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *StoredPaymentMethodResource) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### GetShopperEmail

`func (o *StoredPaymentMethodResource) GetShopperEmail() string`

GetShopperEmail returns the ShopperEmail field if non-nil, zero value otherwise.

### GetShopperEmailOk

`func (o *StoredPaymentMethodResource) GetShopperEmailOk() (*string, bool)`

GetShopperEmailOk returns a tuple with the ShopperEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperEmail

`func (o *StoredPaymentMethodResource) SetShopperEmail(v string)`

SetShopperEmail sets ShopperEmail field to given value.

### HasShopperEmail

`func (o *StoredPaymentMethodResource) HasShopperEmail() bool`

HasShopperEmail returns a boolean if a field has been set.

### GetShopperReference

`func (o *StoredPaymentMethodResource) GetShopperReference() string`

GetShopperReference returns the ShopperReference field if non-nil, zero value otherwise.

### GetShopperReferenceOk

`func (o *StoredPaymentMethodResource) GetShopperReferenceOk() (*string, bool)`

GetShopperReferenceOk returns a tuple with the ShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperReference

`func (o *StoredPaymentMethodResource) SetShopperReference(v string)`

SetShopperReference sets ShopperReference field to given value.

### HasShopperReference

`func (o *StoredPaymentMethodResource) HasShopperReference() bool`

HasShopperReference returns a boolean if a field has been set.

### GetSupportedRecurringProcessingModels

`func (o *StoredPaymentMethodResource) GetSupportedRecurringProcessingModels() []string`

GetSupportedRecurringProcessingModels returns the SupportedRecurringProcessingModels field if non-nil, zero value otherwise.

### GetSupportedRecurringProcessingModelsOk

`func (o *StoredPaymentMethodResource) GetSupportedRecurringProcessingModelsOk() (*[]string, bool)`

GetSupportedRecurringProcessingModelsOk returns a tuple with the SupportedRecurringProcessingModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedRecurringProcessingModels

`func (o *StoredPaymentMethodResource) SetSupportedRecurringProcessingModels(v []string)`

SetSupportedRecurringProcessingModels sets SupportedRecurringProcessingModels field to given value.

### HasSupportedRecurringProcessingModels

`func (o *StoredPaymentMethodResource) HasSupportedRecurringProcessingModels() bool`

HasSupportedRecurringProcessingModels returns a boolean if a field has been set.

### GetType

`func (o *StoredPaymentMethodResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StoredPaymentMethodResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StoredPaymentMethodResource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StoredPaymentMethodResource) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


