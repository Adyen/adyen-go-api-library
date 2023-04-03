# RecurringDetail

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
**RecurringDetailReference** | Pointer to **string** | The reference that uniquely identifies the recurring detail. | [optional] 
**StoredDetails** | Pointer to [**StoredDetails**](StoredDetails.md) |  | [optional] 
**Type** | Pointer to **string** | The unique payment method code. | [optional] 

## Methods

### NewRecurringDetail

`func NewRecurringDetail() *RecurringDetail`

NewRecurringDetail instantiates a new RecurringDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurringDetailWithDefaults

`func NewRecurringDetailWithDefaults() *RecurringDetail`

NewRecurringDetailWithDefaults instantiates a new RecurringDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *RecurringDetail) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *RecurringDetail) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *RecurringDetail) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *RecurringDetail) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetBrands

`func (o *RecurringDetail) GetBrands() []string`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *RecurringDetail) GetBrandsOk() (*[]string, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *RecurringDetail) SetBrands(v []string)`

SetBrands sets Brands field to given value.

### HasBrands

`func (o *RecurringDetail) HasBrands() bool`

HasBrands returns a boolean if a field has been set.

### GetConfiguration

`func (o *RecurringDetail) GetConfiguration() map[string]string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *RecurringDetail) GetConfigurationOk() (*map[string]string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *RecurringDetail) SetConfiguration(v map[string]string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *RecurringDetail) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetFundingSource

`func (o *RecurringDetail) GetFundingSource() string`

GetFundingSource returns the FundingSource field if non-nil, zero value otherwise.

### GetFundingSourceOk

`func (o *RecurringDetail) GetFundingSourceOk() (*string, bool)`

GetFundingSourceOk returns a tuple with the FundingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSource

`func (o *RecurringDetail) SetFundingSource(v string)`

SetFundingSource sets FundingSource field to given value.

### HasFundingSource

`func (o *RecurringDetail) HasFundingSource() bool`

HasFundingSource returns a boolean if a field has been set.

### GetGroup

`func (o *RecurringDetail) GetGroup() PaymentMethodGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *RecurringDetail) GetGroupOk() (*PaymentMethodGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *RecurringDetail) SetGroup(v PaymentMethodGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *RecurringDetail) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetInputDetails

`func (o *RecurringDetail) GetInputDetails() []InputDetail`

GetInputDetails returns the InputDetails field if non-nil, zero value otherwise.

### GetInputDetailsOk

`func (o *RecurringDetail) GetInputDetailsOk() (*[]InputDetail, bool)`

GetInputDetailsOk returns a tuple with the InputDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputDetails

`func (o *RecurringDetail) SetInputDetails(v []InputDetail)`

SetInputDetails sets InputDetails field to given value.

### HasInputDetails

`func (o *RecurringDetail) HasInputDetails() bool`

HasInputDetails returns a boolean if a field has been set.

### GetIssuers

`func (o *RecurringDetail) GetIssuers() []PaymentMethodIssuer`

GetIssuers returns the Issuers field if non-nil, zero value otherwise.

### GetIssuersOk

`func (o *RecurringDetail) GetIssuersOk() (*[]PaymentMethodIssuer, bool)`

GetIssuersOk returns a tuple with the Issuers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuers

`func (o *RecurringDetail) SetIssuers(v []PaymentMethodIssuer)`

SetIssuers sets Issuers field to given value.

### HasIssuers

`func (o *RecurringDetail) HasIssuers() bool`

HasIssuers returns a boolean if a field has been set.

### GetName

`func (o *RecurringDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecurringDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecurringDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecurringDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecurringDetailReference

`func (o *RecurringDetail) GetRecurringDetailReference() string`

GetRecurringDetailReference returns the RecurringDetailReference field if non-nil, zero value otherwise.

### GetRecurringDetailReferenceOk

`func (o *RecurringDetail) GetRecurringDetailReferenceOk() (*string, bool)`

GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDetailReference

`func (o *RecurringDetail) SetRecurringDetailReference(v string)`

SetRecurringDetailReference sets RecurringDetailReference field to given value.

### HasRecurringDetailReference

`func (o *RecurringDetail) HasRecurringDetailReference() bool`

HasRecurringDetailReference returns a boolean if a field has been set.

### GetStoredDetails

`func (o *RecurringDetail) GetStoredDetails() StoredDetails`

GetStoredDetails returns the StoredDetails field if non-nil, zero value otherwise.

### GetStoredDetailsOk

`func (o *RecurringDetail) GetStoredDetailsOk() (*StoredDetails, bool)`

GetStoredDetailsOk returns a tuple with the StoredDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredDetails

`func (o *RecurringDetail) SetStoredDetails(v StoredDetails)`

SetStoredDetails sets StoredDetails field to given value.

### HasStoredDetails

`func (o *RecurringDetail) HasStoredDetails() bool`

HasStoredDetails returns a boolean if a field has been set.

### GetType

`func (o *RecurringDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecurringDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecurringDetail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RecurringDetail) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


