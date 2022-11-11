# InstallAndroidAppDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | The unique identifier of the app to be installed. | [optional] 
**Type** | Pointer to **string** | Type of terminal action: Install an Android app. | [optional] [default to "InstallAndroidApp"]

## Methods

### NewInstallAndroidAppDetails

`func NewInstallAndroidAppDetails() *InstallAndroidAppDetails`

NewInstallAndroidAppDetails instantiates a new InstallAndroidAppDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallAndroidAppDetailsWithDefaults

`func NewInstallAndroidAppDetailsWithDefaults() *InstallAndroidAppDetails`

NewInstallAndroidAppDetailsWithDefaults instantiates a new InstallAndroidAppDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *InstallAndroidAppDetails) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *InstallAndroidAppDetails) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *InstallAndroidAppDetails) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *InstallAndroidAppDetails) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetType

`func (o *InstallAndroidAppDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstallAndroidAppDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstallAndroidAppDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InstallAndroidAppDetails) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


