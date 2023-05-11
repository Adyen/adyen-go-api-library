# AndroidApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description that was provided when uploading the app. The description is not shown on the terminal. | [optional] 
**Id** | **string** | The unique identifier of the app. | 
**Label** | Pointer to **string** | The app name that is shown on the terminal. | [optional] 
**PackageName** | Pointer to **string** | The package name of the app. | [optional] 
**Status** | **string** | The status of the app. Possible values:  * &#x60;processing&#x60;: The app is being signed and converted to a format that the terminal can handle. * &#x60;error&#x60;: Something went wrong. Check that the app matches the [requirements](https://docs.adyen.com/point-of-sale/android-terminals/app-requirements). * &#x60;invalid&#x60;: There is something wrong with the APK file of the app. * &#x60;ready&#x60;: The app has been signed and converted. * &#x60;archived&#x60;: The app is no longer available. | 
**VersionCode** | Pointer to **int32** | The internal version number of the app. | [optional] 
**VersionName** | Pointer to **string** | The app version number that is shown on the terminal. | [optional] 

## Methods

### NewAndroidApp

`func NewAndroidApp(id string, status string, ) *AndroidApp`

NewAndroidApp instantiates a new AndroidApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAndroidAppWithDefaults

`func NewAndroidAppWithDefaults() *AndroidApp`

NewAndroidAppWithDefaults instantiates a new AndroidApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AndroidApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AndroidApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AndroidApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AndroidApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *AndroidApp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AndroidApp) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AndroidApp) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *AndroidApp) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AndroidApp) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AndroidApp) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AndroidApp) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPackageName

`func (o *AndroidApp) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *AndroidApp) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *AndroidApp) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *AndroidApp) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetStatus

`func (o *AndroidApp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AndroidApp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AndroidApp) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetVersionCode

`func (o *AndroidApp) GetVersionCode() int32`

GetVersionCode returns the VersionCode field if non-nil, zero value otherwise.

### GetVersionCodeOk

`func (o *AndroidApp) GetVersionCodeOk() (*int32, bool)`

GetVersionCodeOk returns a tuple with the VersionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionCode

`func (o *AndroidApp) SetVersionCode(v int32)`

SetVersionCode sets VersionCode field to given value.

### HasVersionCode

`func (o *AndroidApp) HasVersionCode() bool`

HasVersionCode returns a boolean if a field has been set.

### GetVersionName

`func (o *AndroidApp) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *AndroidApp) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *AndroidApp) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.

### HasVersionName

`func (o *AndroidApp) HasVersionName() bool`

HasVersionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


