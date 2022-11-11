# WifiProfiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profiles** | Pointer to [**[]Profile**](Profile.md) | List of remote Wi-Fi profiles | [optional] 
**Settings** | Pointer to [**Settings**](Settings.md) |  | [optional] 

## Methods

### NewWifiProfiles

`func NewWifiProfiles() *WifiProfiles`

NewWifiProfiles instantiates a new WifiProfiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWifiProfilesWithDefaults

`func NewWifiProfilesWithDefaults() *WifiProfiles`

NewWifiProfilesWithDefaults instantiates a new WifiProfiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfiles

`func (o *WifiProfiles) GetProfiles() []Profile`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *WifiProfiles) GetProfilesOk() (*[]Profile, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *WifiProfiles) SetProfiles(v []Profile)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *WifiProfiles) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetSettings

`func (o *WifiProfiles) GetSettings() Settings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *WifiProfiles) GetSettingsOk() (*Settings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *WifiProfiles) SetSettings(v Settings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *WifiProfiles) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


