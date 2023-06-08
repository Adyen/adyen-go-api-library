# SplitConfigurationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SplitConfiguration**](SplitConfiguration.md) | List of split configurations applied to the stores under the merchant account. | [optional] 

## Methods

### NewSplitConfigurationList

`func NewSplitConfigurationList() *SplitConfigurationList`

NewSplitConfigurationList instantiates a new SplitConfigurationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSplitConfigurationListWithDefaults

`func NewSplitConfigurationListWithDefaults() *SplitConfigurationList`

NewSplitConfigurationListWithDefaults instantiates a new SplitConfigurationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SplitConfigurationList) GetData() []SplitConfiguration`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SplitConfigurationList) GetDataOk() (*[]SplitConfiguration, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SplitConfigurationList) SetData(v []SplitConfiguration)`

SetData sets Data field to given value.

### HasData

`func (o *SplitConfigurationList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


