# DeviceRenderOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SdkInterface** | Pointer to **string** | Supported SDK interface types. Allowed values: * native * html * both | [optional] [default to "both"]
**SdkUiType** | Pointer to **[]string** | UI types supported for displaying specific challenges. Allowed values: * text * singleSelect * outOfBand * otherHtml * multiSelect | [optional] 

## Methods

### NewDeviceRenderOptions

`func NewDeviceRenderOptions() *DeviceRenderOptions`

NewDeviceRenderOptions instantiates a new DeviceRenderOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceRenderOptionsWithDefaults

`func NewDeviceRenderOptionsWithDefaults() *DeviceRenderOptions`

NewDeviceRenderOptionsWithDefaults instantiates a new DeviceRenderOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSdkInterface

`func (o *DeviceRenderOptions) GetSdkInterface() string`

GetSdkInterface returns the SdkInterface field if non-nil, zero value otherwise.

### GetSdkInterfaceOk

`func (o *DeviceRenderOptions) GetSdkInterfaceOk() (*string, bool)`

GetSdkInterfaceOk returns a tuple with the SdkInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkInterface

`func (o *DeviceRenderOptions) SetSdkInterface(v string)`

SetSdkInterface sets SdkInterface field to given value.

### HasSdkInterface

`func (o *DeviceRenderOptions) HasSdkInterface() bool`

HasSdkInterface returns a boolean if a field has been set.

### GetSdkUiType

`func (o *DeviceRenderOptions) GetSdkUiType() []string`

GetSdkUiType returns the SdkUiType field if non-nil, zero value otherwise.

### GetSdkUiTypeOk

`func (o *DeviceRenderOptions) GetSdkUiTypeOk() (*[]string, bool)`

GetSdkUiTypeOk returns a tuple with the SdkUiType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkUiType

`func (o *DeviceRenderOptions) SetSdkUiType(v []string)`

SetSdkUiType sets SdkUiType field to given value.

### HasSdkUiType

`func (o *DeviceRenderOptions) HasSdkUiType() bool`

HasSdkUiType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


