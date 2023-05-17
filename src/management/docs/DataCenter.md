# DataCenter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LivePrefix** | Pointer to **string** | The unique [live URL prefix](https://docs.adyen.com/development-resources/live-endpoints#live-url-prefix) for your live endpoint. Each data center has its own live URL prefix.  This field is empty for requests made in the test environment. | [optional] 
**Name** | Pointer to **string** | The name assigned to a data center, for example **EU** for the European data center. Possible values are:  * **default**: the European data center. This value is always returned in the test environment.  * **AU** * **EU** * **US** | [optional] 

## Methods

### NewDataCenter

`func NewDataCenter() *DataCenter`

NewDataCenter instantiates a new DataCenter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataCenterWithDefaults

`func NewDataCenterWithDefaults() *DataCenter`

NewDataCenterWithDefaults instantiates a new DataCenter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLivePrefix

`func (o *DataCenter) GetLivePrefix() string`

GetLivePrefix returns the LivePrefix field if non-nil, zero value otherwise.

### GetLivePrefixOk

`func (o *DataCenter) GetLivePrefixOk() (*string, bool)`

GetLivePrefixOk returns a tuple with the LivePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivePrefix

`func (o *DataCenter) SetLivePrefix(v string)`

SetLivePrefix sets LivePrefix field to given value.

### HasLivePrefix

`func (o *DataCenter) HasLivePrefix() bool`

HasLivePrefix returns a boolean if a field has been set.

### GetName

`func (o *DataCenter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataCenter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataCenter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataCenter) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


