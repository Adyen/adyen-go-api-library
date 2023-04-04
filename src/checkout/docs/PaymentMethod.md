# PaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to **string** | Brand for the selected gift card. For example: plastix, hmclub. | [optional] 
**Brands** | Pointer to **[]string** | List of possible brands. For example: visa, mc. | [optional] 
**Configuration** | Pointer to **map[string]string** | The configuration of the payment method. | [optional] 
**FundingSource** | Pointer to **string** | The funding source of the payment method. | [optional] 
**Group** | Pointer to [**PaymentMethodGroup**](PaymentMethodGroup.md) |  | [optional] 
**InputDetails** | Pointer to [**[]InputDetail**](InputDetail.md) | All input details to be provided to complete the payment with this payment method. | [optional] 
**Issuers** | Pointer to [**[]PaymentMethodIssuer**](PaymentMethodIssuer.md) | A list of issuers for this payment method. | [optional] 
**Name** | Pointer to **string** | The displayable name of this payment method. | [optional] 
**Type** | Pointer to **string** | The unique payment method code. | [optional] 

## Methods

### NewPaymentMethod

`func NewPaymentMethod() *PaymentMethod`

NewPaymentMethod instantiates a new PaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodWithDefaults

`func NewPaymentMethodWithDefaults() *PaymentMethod`

NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *PaymentMethod) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *PaymentMethod) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *PaymentMethod) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *PaymentMethod) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetBrands

`func (o *PaymentMethod) GetBrands() []string`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *PaymentMethod) GetBrandsOk() (*[]string, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *PaymentMethod) SetBrands(v []string)`

SetBrands sets Brands field to given value.

### HasBrands

`func (o *PaymentMethod) HasBrands() bool`

HasBrands returns a boolean if a field has been set.

### GetConfiguration

`func (o *PaymentMethod) GetConfiguration() map[string]string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PaymentMethod) GetConfigurationOk() (*map[string]string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PaymentMethod) SetConfiguration(v map[string]string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *PaymentMethod) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetFundingSource

`func (o *PaymentMethod) GetFundingSource() string`

GetFundingSource returns the FundingSource field if non-nil, zero value otherwise.

### GetFundingSourceOk

`func (o *PaymentMethod) GetFundingSourceOk() (*string, bool)`

GetFundingSourceOk returns a tuple with the FundingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSource

`func (o *PaymentMethod) SetFundingSource(v string)`

SetFundingSource sets FundingSource field to given value.

### HasFundingSource

`func (o *PaymentMethod) HasFundingSource() bool`

HasFundingSource returns a boolean if a field has been set.

### GetGroup

`func (o *PaymentMethod) GetGroup() PaymentMethodGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PaymentMethod) GetGroupOk() (*PaymentMethodGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PaymentMethod) SetGroup(v PaymentMethodGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PaymentMethod) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetInputDetails

`func (o *PaymentMethod) GetInputDetails() []InputDetail`

GetInputDetails returns the InputDetails field if non-nil, zero value otherwise.

### GetInputDetailsOk

`func (o *PaymentMethod) GetInputDetailsOk() (*[]InputDetail, bool)`

GetInputDetailsOk returns a tuple with the InputDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputDetails

`func (o *PaymentMethod) SetInputDetails(v []InputDetail)`

SetInputDetails sets InputDetails field to given value.

### HasInputDetails

`func (o *PaymentMethod) HasInputDetails() bool`

HasInputDetails returns a boolean if a field has been set.

### GetIssuers

`func (o *PaymentMethod) GetIssuers() []PaymentMethodIssuer`

GetIssuers returns the Issuers field if non-nil, zero value otherwise.

### GetIssuersOk

`func (o *PaymentMethod) GetIssuersOk() (*[]PaymentMethodIssuer, bool)`

GetIssuersOk returns a tuple with the Issuers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuers

`func (o *PaymentMethod) SetIssuers(v []PaymentMethodIssuer)`

SetIssuers sets Issuers field to given value.

### HasIssuers

`func (o *PaymentMethod) HasIssuers() bool`

HasIssuers returns a boolean if a field has been set.

### GetName

`func (o *PaymentMethod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentMethod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentMethod) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaymentMethod) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *PaymentMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethod) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentMethod) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


