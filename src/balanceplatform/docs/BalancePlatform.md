# BalancePlatform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Your description of the balance platform, maximum 300 characters. | [optional] 
**Id** | **string** | The unique identifier of the balance platform. | 
**Status** | Pointer to **string** | The status of the balance platform.  Possible values: **Active**, **Inactive**, **Closed**, **Suspended**. | [optional] 

## Methods

### NewBalancePlatform

`func NewBalancePlatform(id string, ) *BalancePlatform`

NewBalancePlatform instantiates a new BalancePlatform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalancePlatformWithDefaults

`func NewBalancePlatformWithDefaults() *BalancePlatform`

NewBalancePlatformWithDefaults instantiates a new BalancePlatform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BalancePlatform) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BalancePlatform) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BalancePlatform) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BalancePlatform) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *BalancePlatform) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BalancePlatform) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BalancePlatform) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *BalancePlatform) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BalancePlatform) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BalancePlatform) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BalancePlatform) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


