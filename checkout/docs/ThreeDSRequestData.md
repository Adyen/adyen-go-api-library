# ThreeDSRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChallengeWindowSize** | Pointer to **string** | Dimensions of the 3DS2 challenge window to be displayed to the cardholder.  Possible values:  * **01** - size of 250x400  * **02** - size of 390x400 * **03** - size of 500x600 * **04** - size of 600x400 * **05** - Fullscreen | [optional] 
**DataOnly** | Pointer to **string** | Flag for data only flow. | [optional] 
**NativeThreeDS** | Pointer to **string** | Indicates if [native 3D Secure authentication](https://docs.adyen.com/online-payments/3d-secure/native-3ds2) should be used when available.  Possible values: * **preferred**: Use native 3D Secure authentication when available. | [optional] 
**ThreeDSVersion** | Pointer to **string** | The version of 3D Secure to use.  Possible values:  * **2.1.0** * **2.2.0** | [optional] 

## Methods

### NewThreeDSRequestData

`func NewThreeDSRequestData() *ThreeDSRequestData`

NewThreeDSRequestData instantiates a new ThreeDSRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSRequestDataWithDefaults

`func NewThreeDSRequestDataWithDefaults() *ThreeDSRequestData`

NewThreeDSRequestDataWithDefaults instantiates a new ThreeDSRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChallengeWindowSize

`func (o *ThreeDSRequestData) GetChallengeWindowSize() string`

GetChallengeWindowSize returns the ChallengeWindowSize field if non-nil, zero value otherwise.

### GetChallengeWindowSizeOk

`func (o *ThreeDSRequestData) GetChallengeWindowSizeOk() (*string, bool)`

GetChallengeWindowSizeOk returns a tuple with the ChallengeWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeWindowSize

`func (o *ThreeDSRequestData) SetChallengeWindowSize(v string)`

SetChallengeWindowSize sets ChallengeWindowSize field to given value.

### HasChallengeWindowSize

`func (o *ThreeDSRequestData) HasChallengeWindowSize() bool`

HasChallengeWindowSize returns a boolean if a field has been set.

### GetDataOnly

`func (o *ThreeDSRequestData) GetDataOnly() string`

GetDataOnly returns the DataOnly field if non-nil, zero value otherwise.

### GetDataOnlyOk

`func (o *ThreeDSRequestData) GetDataOnlyOk() (*string, bool)`

GetDataOnlyOk returns a tuple with the DataOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataOnly

`func (o *ThreeDSRequestData) SetDataOnly(v string)`

SetDataOnly sets DataOnly field to given value.

### HasDataOnly

`func (o *ThreeDSRequestData) HasDataOnly() bool`

HasDataOnly returns a boolean if a field has been set.

### GetNativeThreeDS

`func (o *ThreeDSRequestData) GetNativeThreeDS() string`

GetNativeThreeDS returns the NativeThreeDS field if non-nil, zero value otherwise.

### GetNativeThreeDSOk

`func (o *ThreeDSRequestData) GetNativeThreeDSOk() (*string, bool)`

GetNativeThreeDSOk returns a tuple with the NativeThreeDS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeThreeDS

`func (o *ThreeDSRequestData) SetNativeThreeDS(v string)`

SetNativeThreeDS sets NativeThreeDS field to given value.

### HasNativeThreeDS

`func (o *ThreeDSRequestData) HasNativeThreeDS() bool`

HasNativeThreeDS returns a boolean if a field has been set.

### GetThreeDSVersion

`func (o *ThreeDSRequestData) GetThreeDSVersion() string`

GetThreeDSVersion returns the ThreeDSVersion field if non-nil, zero value otherwise.

### GetThreeDSVersionOk

`func (o *ThreeDSRequestData) GetThreeDSVersionOk() (*string, bool)`

GetThreeDSVersionOk returns a tuple with the ThreeDSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSVersion

`func (o *ThreeDSRequestData) SetThreeDSVersion(v string)`

SetThreeDSVersion sets ThreeDSVersion field to given value.

### HasThreeDSVersion

`func (o *ThreeDSRequestData) HasThreeDSVersion() bool`

HasThreeDSVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


