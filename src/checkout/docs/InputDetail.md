# InputDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to **map[string]string** | Configuration parameters for the required input. | [optional] 
**Details** | Pointer to [**[]SubInputDetail**](SubInputDetail.md) | Input details can also be provided recursively. | [optional] 
**InputDetails** | Pointer to [**[]SubInputDetail**](SubInputDetail.md) | Input details can also be provided recursively (deprecated). | [optional] 
**ItemSearchUrl** | Pointer to **string** | In case of a select, the URL from which to query the items. | [optional] 
**Items** | Pointer to [**[]Item**](Item.md) | In case of a select, the items to choose from. | [optional] 
**Key** | Pointer to **string** | The value to provide in the result. | [optional] 
**Optional** | Pointer to **bool** | True if this input value is optional. | [optional] 
**Type** | Pointer to **string** | The type of the required input. | [optional] 
**Value** | Pointer to **string** | The value can be pre-filled, if available. | [optional] 

## Methods

### NewInputDetail

`func NewInputDetail() *InputDetail`

NewInputDetail instantiates a new InputDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputDetailWithDefaults

`func NewInputDetailWithDefaults() *InputDetail`

NewInputDetailWithDefaults instantiates a new InputDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *InputDetail) GetConfiguration() map[string]string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *InputDetail) GetConfigurationOk() (*map[string]string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *InputDetail) SetConfiguration(v map[string]string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *InputDetail) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDetails

`func (o *InputDetail) GetDetails() []SubInputDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *InputDetail) GetDetailsOk() (*[]SubInputDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *InputDetail) SetDetails(v []SubInputDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *InputDetail) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetInputDetails

`func (o *InputDetail) GetInputDetails() []SubInputDetail`

GetInputDetails returns the InputDetails field if non-nil, zero value otherwise.

### GetInputDetailsOk

`func (o *InputDetail) GetInputDetailsOk() (*[]SubInputDetail, bool)`

GetInputDetailsOk returns a tuple with the InputDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputDetails

`func (o *InputDetail) SetInputDetails(v []SubInputDetail)`

SetInputDetails sets InputDetails field to given value.

### HasInputDetails

`func (o *InputDetail) HasInputDetails() bool`

HasInputDetails returns a boolean if a field has been set.

### GetItemSearchUrl

`func (o *InputDetail) GetItemSearchUrl() string`

GetItemSearchUrl returns the ItemSearchUrl field if non-nil, zero value otherwise.

### GetItemSearchUrlOk

`func (o *InputDetail) GetItemSearchUrlOk() (*string, bool)`

GetItemSearchUrlOk returns a tuple with the ItemSearchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSearchUrl

`func (o *InputDetail) SetItemSearchUrl(v string)`

SetItemSearchUrl sets ItemSearchUrl field to given value.

### HasItemSearchUrl

`func (o *InputDetail) HasItemSearchUrl() bool`

HasItemSearchUrl returns a boolean if a field has been set.

### GetItems

`func (o *InputDetail) GetItems() []Item`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InputDetail) GetItemsOk() (*[]Item, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InputDetail) SetItems(v []Item)`

SetItems sets Items field to given value.

### HasItems

`func (o *InputDetail) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKey

`func (o *InputDetail) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *InputDetail) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *InputDetail) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *InputDetail) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetOptional

`func (o *InputDetail) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *InputDetail) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *InputDetail) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *InputDetail) HasOptional() bool`

HasOptional returns a boolean if a field has been set.

### GetType

`func (o *InputDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputDetail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InputDetail) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *InputDetail) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InputDetail) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InputDetail) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *InputDetail) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


