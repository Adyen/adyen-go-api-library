# DSPublicKeyDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to **string** | Card brand. | [optional] 
**DirectoryServerId** | Pointer to **string** | Directory Server (DS) identifier. | [optional] 
**FromSDKVersion** | Pointer to **string** | The version of the mobile 3D Secure 2 SDK. For the possible values, refer to the versions in [Adyen 3DS2 Android](https://github.com/Adyen/adyen-3ds2-android/releases) and [Adyen 3DS2 iOS](https://github.com/Adyen/adyen-3ds2-ios/releases). | [optional] 
**PublicKey** | Pointer to **string** | Public key. The 3D Secure 2 SDK encrypts the device information by using the DS public key. | [optional] 

## Methods

### NewDSPublicKeyDetail

`func NewDSPublicKeyDetail() *DSPublicKeyDetail`

NewDSPublicKeyDetail instantiates a new DSPublicKeyDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDSPublicKeyDetailWithDefaults

`func NewDSPublicKeyDetailWithDefaults() *DSPublicKeyDetail`

NewDSPublicKeyDetailWithDefaults instantiates a new DSPublicKeyDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *DSPublicKeyDetail) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *DSPublicKeyDetail) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *DSPublicKeyDetail) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *DSPublicKeyDetail) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetDirectoryServerId

`func (o *DSPublicKeyDetail) GetDirectoryServerId() string`

GetDirectoryServerId returns the DirectoryServerId field if non-nil, zero value otherwise.

### GetDirectoryServerIdOk

`func (o *DSPublicKeyDetail) GetDirectoryServerIdOk() (*string, bool)`

GetDirectoryServerIdOk returns a tuple with the DirectoryServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryServerId

`func (o *DSPublicKeyDetail) SetDirectoryServerId(v string)`

SetDirectoryServerId sets DirectoryServerId field to given value.

### HasDirectoryServerId

`func (o *DSPublicKeyDetail) HasDirectoryServerId() bool`

HasDirectoryServerId returns a boolean if a field has been set.

### GetFromSDKVersion

`func (o *DSPublicKeyDetail) GetFromSDKVersion() string`

GetFromSDKVersion returns the FromSDKVersion field if non-nil, zero value otherwise.

### GetFromSDKVersionOk

`func (o *DSPublicKeyDetail) GetFromSDKVersionOk() (*string, bool)`

GetFromSDKVersionOk returns a tuple with the FromSDKVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromSDKVersion

`func (o *DSPublicKeyDetail) SetFromSDKVersion(v string)`

SetFromSDKVersion sets FromSDKVersion field to given value.

### HasFromSDKVersion

`func (o *DSPublicKeyDetail) HasFromSDKVersion() bool`

HasFromSDKVersion returns a boolean if a field has been set.

### GetPublicKey

`func (o *DSPublicKeyDetail) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *DSPublicKeyDetail) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *DSPublicKeyDetail) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *DSPublicKeyDetail) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


