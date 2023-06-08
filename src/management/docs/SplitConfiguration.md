# SplitConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Your description for the split configuration. | 
**Rules** | [**[]SplitConfigurationRule**](SplitConfigurationRule.md) | Array of rules that define the split configuration behavior. | 
**SplitConfigurationId** | Pointer to **string** | Unique identifier of the split configuration. | [optional] [readonly] 
**Stores** | Pointer to **[]string** | List of stores to which the split configuration applies. | [optional] [readonly] 

## Methods

### NewSplitConfiguration

`func NewSplitConfiguration(description string, rules []SplitConfigurationRule, ) *SplitConfiguration`

NewSplitConfiguration instantiates a new SplitConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSplitConfigurationWithDefaults

`func NewSplitConfigurationWithDefaults() *SplitConfiguration`

NewSplitConfigurationWithDefaults instantiates a new SplitConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SplitConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SplitConfiguration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SplitConfiguration) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRules

`func (o *SplitConfiguration) GetRules() []SplitConfigurationRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *SplitConfiguration) GetRulesOk() (*[]SplitConfigurationRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *SplitConfiguration) SetRules(v []SplitConfigurationRule)`

SetRules sets Rules field to given value.


### GetSplitConfigurationId

`func (o *SplitConfiguration) GetSplitConfigurationId() string`

GetSplitConfigurationId returns the SplitConfigurationId field if non-nil, zero value otherwise.

### GetSplitConfigurationIdOk

`func (o *SplitConfiguration) GetSplitConfigurationIdOk() (*string, bool)`

GetSplitConfigurationIdOk returns a tuple with the SplitConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitConfigurationId

`func (o *SplitConfiguration) SetSplitConfigurationId(v string)`

SetSplitConfigurationId sets SplitConfigurationId field to given value.

### HasSplitConfigurationId

`func (o *SplitConfiguration) HasSplitConfigurationId() bool`

HasSplitConfigurationId returns a boolean if a field has been set.

### GetStores

`func (o *SplitConfiguration) GetStores() []string`

GetStores returns the Stores field if non-nil, zero value otherwise.

### GetStoresOk

`func (o *SplitConfiguration) GetStoresOk() (*[]string, bool)`

GetStoresOk returns a tuple with the Stores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStores

`func (o *SplitConfiguration) SetStores(v []string)`

SetStores sets Stores field to given value.

### HasStores

`func (o *SplitConfiguration) HasStores() bool`

HasStores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


