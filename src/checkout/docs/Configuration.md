# Configuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avs** | Pointer to [**Avs**](Avs.md) |  | [optional] 
**CardHolderName** | Pointer to **string** | Determines whether the cardholder name should be provided or not.  Permitted values: * NONE * OPTIONAL * REQUIRED | [optional] 
**Installments** | Pointer to [**InstallmentsNumber**](InstallmentsNumber.md) |  | [optional] 
**ShopperInput** | Pointer to [**ShopperInput**](ShopperInput.md) |  | [optional] 

## Methods

### NewConfiguration

`func NewConfiguration() *Configuration`

NewConfiguration instantiates a new Configuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationWithDefaults

`func NewConfigurationWithDefaults() *Configuration`

NewConfigurationWithDefaults instantiates a new Configuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvs

`func (o *Configuration) GetAvs() Avs`

GetAvs returns the Avs field if non-nil, zero value otherwise.

### GetAvsOk

`func (o *Configuration) GetAvsOk() (*Avs, bool)`

GetAvsOk returns a tuple with the Avs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvs

`func (o *Configuration) SetAvs(v Avs)`

SetAvs sets Avs field to given value.

### HasAvs

`func (o *Configuration) HasAvs() bool`

HasAvs returns a boolean if a field has been set.

### GetCardHolderName

`func (o *Configuration) GetCardHolderName() string`

GetCardHolderName returns the CardHolderName field if non-nil, zero value otherwise.

### GetCardHolderNameOk

`func (o *Configuration) GetCardHolderNameOk() (*string, bool)`

GetCardHolderNameOk returns a tuple with the CardHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardHolderName

`func (o *Configuration) SetCardHolderName(v string)`

SetCardHolderName sets CardHolderName field to given value.

### HasCardHolderName

`func (o *Configuration) HasCardHolderName() bool`

HasCardHolderName returns a boolean if a field has been set.

### GetInstallments

`func (o *Configuration) GetInstallments() InstallmentsNumber`

GetInstallments returns the Installments field if non-nil, zero value otherwise.

### GetInstallmentsOk

`func (o *Configuration) GetInstallmentsOk() (*InstallmentsNumber, bool)`

GetInstallmentsOk returns a tuple with the Installments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallments

`func (o *Configuration) SetInstallments(v InstallmentsNumber)`

SetInstallments sets Installments field to given value.

### HasInstallments

`func (o *Configuration) HasInstallments() bool`

HasInstallments returns a boolean if a field has been set.

### GetShopperInput

`func (o *Configuration) GetShopperInput() ShopperInput`

GetShopperInput returns the ShopperInput field if non-nil, zero value otherwise.

### GetShopperInputOk

`func (o *Configuration) GetShopperInputOk() (*ShopperInput, bool)`

GetShopperInputOk returns a tuple with the ShopperInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperInput

`func (o *Configuration) SetShopperInput(v ShopperInput)`

SetShopperInput sets ShopperInput field to given value.

### HasShopperInput

`func (o *Configuration) HasShopperInput() bool`

HasShopperInput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


