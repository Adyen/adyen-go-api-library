# Profile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | **string** | The type of Wi-Fi network. Possible values: **wpa-psk**, **wpa2-psk**, **wpa-eap**, **wpa2-eap**. | 
**AutoWifi** | Pointer to **bool** | Indicates whether to automatically select the best authentication method available. Does not work on older terminal models. | [optional] 
**BssType** | **string** | Use **infra** for infrastructure-based networks. This applies to most networks. Use **adhoc** only if the communication is p2p-based between base stations. | 
**Channel** | Pointer to **int32** | The channel number of the Wi-Fi network. The recommended setting is **0** for automatic channel selection. | [optional] 
**DefaultProfile** | Pointer to **bool** | Indicates whether this is your preferred wireless network. If **true**, the terminal will try connecting to this network first. | [optional] 
**Eap** | Pointer to **string** | For &#x60;authType&#x60; **wpa-eap** or **wpa2-eap**. Possible values: **tls**, **peap**, **leap**, **fast** | [optional] 
**EapCaCert** | Pointer to [**File**](File.md) |  | [optional] 
**EapClientCert** | Pointer to [**File**](File.md) |  | [optional] 
**EapClientKey** | Pointer to [**File**](File.md) |  | [optional] 
**EapClientPwd** | Pointer to **string** | For &#x60;eap&#x60; **tls**. The password of the RSA key file, if that file is password-protected. | [optional] 
**EapIdentity** | Pointer to **string** | For &#x60;authType&#x60; **wpa-eap** or **wpa2-eap**. The EAP-PEAP username from your MS-CHAP account. Must match the configuration of your RADIUS server. | [optional] 
**EapIntermediateCert** | Pointer to [**File**](File.md) |  | [optional] 
**EapPwd** | Pointer to **string** | For &#x60;eap&#x60; **peap**. The EAP-PEAP password from your MS-CHAP account. Must match the configuration of your RADIUS server. | [optional] 
**HiddenSsid** | Pointer to **bool** | Indicates if the network doesn&#39;t broadcast its SSID. Mandatory for Android terminals, because these terminals rely on this setting to be able to connect to any network. | [optional] 
**Name** | Pointer to **string** | Your name for the Wi-Fi profile. | [optional] 
**Psk** | Pointer to **string** | For &#x60;authType&#x60; **wpa-psk or **wpa2-psk**. The password to the wireless network. | [optional] 
**Ssid** | **string** | The name of the wireless network. | 
**Wsec** | **string** | The type of encryption. Possible values: **auto**, **ccmp** (recommended), **tkip** | 

## Methods

### NewProfile

`func NewProfile(authType string, bssType string, ssid string, wsec string, ) *Profile`

NewProfile instantiates a new Profile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileWithDefaults

`func NewProfileWithDefaults() *Profile`

NewProfileWithDefaults instantiates a new Profile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *Profile) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *Profile) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *Profile) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### GetAutoWifi

`func (o *Profile) GetAutoWifi() bool`

GetAutoWifi returns the AutoWifi field if non-nil, zero value otherwise.

### GetAutoWifiOk

`func (o *Profile) GetAutoWifiOk() (*bool, bool)`

GetAutoWifiOk returns a tuple with the AutoWifi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoWifi

`func (o *Profile) SetAutoWifi(v bool)`

SetAutoWifi sets AutoWifi field to given value.

### HasAutoWifi

`func (o *Profile) HasAutoWifi() bool`

HasAutoWifi returns a boolean if a field has been set.

### GetBssType

`func (o *Profile) GetBssType() string`

GetBssType returns the BssType field if non-nil, zero value otherwise.

### GetBssTypeOk

`func (o *Profile) GetBssTypeOk() (*string, bool)`

GetBssTypeOk returns a tuple with the BssType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBssType

`func (o *Profile) SetBssType(v string)`

SetBssType sets BssType field to given value.


### GetChannel

`func (o *Profile) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *Profile) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *Profile) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *Profile) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetDefaultProfile

`func (o *Profile) GetDefaultProfile() bool`

GetDefaultProfile returns the DefaultProfile field if non-nil, zero value otherwise.

### GetDefaultProfileOk

`func (o *Profile) GetDefaultProfileOk() (*bool, bool)`

GetDefaultProfileOk returns a tuple with the DefaultProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultProfile

`func (o *Profile) SetDefaultProfile(v bool)`

SetDefaultProfile sets DefaultProfile field to given value.

### HasDefaultProfile

`func (o *Profile) HasDefaultProfile() bool`

HasDefaultProfile returns a boolean if a field has been set.

### GetEap

`func (o *Profile) GetEap() string`

GetEap returns the Eap field if non-nil, zero value otherwise.

### GetEapOk

`func (o *Profile) GetEapOk() (*string, bool)`

GetEapOk returns a tuple with the Eap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEap

`func (o *Profile) SetEap(v string)`

SetEap sets Eap field to given value.

### HasEap

`func (o *Profile) HasEap() bool`

HasEap returns a boolean if a field has been set.

### GetEapCaCert

`func (o *Profile) GetEapCaCert() File`

GetEapCaCert returns the EapCaCert field if non-nil, zero value otherwise.

### GetEapCaCertOk

`func (o *Profile) GetEapCaCertOk() (*File, bool)`

GetEapCaCertOk returns a tuple with the EapCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapCaCert

`func (o *Profile) SetEapCaCert(v File)`

SetEapCaCert sets EapCaCert field to given value.

### HasEapCaCert

`func (o *Profile) HasEapCaCert() bool`

HasEapCaCert returns a boolean if a field has been set.

### GetEapClientCert

`func (o *Profile) GetEapClientCert() File`

GetEapClientCert returns the EapClientCert field if non-nil, zero value otherwise.

### GetEapClientCertOk

`func (o *Profile) GetEapClientCertOk() (*File, bool)`

GetEapClientCertOk returns a tuple with the EapClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapClientCert

`func (o *Profile) SetEapClientCert(v File)`

SetEapClientCert sets EapClientCert field to given value.

### HasEapClientCert

`func (o *Profile) HasEapClientCert() bool`

HasEapClientCert returns a boolean if a field has been set.

### GetEapClientKey

`func (o *Profile) GetEapClientKey() File`

GetEapClientKey returns the EapClientKey field if non-nil, zero value otherwise.

### GetEapClientKeyOk

`func (o *Profile) GetEapClientKeyOk() (*File, bool)`

GetEapClientKeyOk returns a tuple with the EapClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapClientKey

`func (o *Profile) SetEapClientKey(v File)`

SetEapClientKey sets EapClientKey field to given value.

### HasEapClientKey

`func (o *Profile) HasEapClientKey() bool`

HasEapClientKey returns a boolean if a field has been set.

### GetEapClientPwd

`func (o *Profile) GetEapClientPwd() string`

GetEapClientPwd returns the EapClientPwd field if non-nil, zero value otherwise.

### GetEapClientPwdOk

`func (o *Profile) GetEapClientPwdOk() (*string, bool)`

GetEapClientPwdOk returns a tuple with the EapClientPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapClientPwd

`func (o *Profile) SetEapClientPwd(v string)`

SetEapClientPwd sets EapClientPwd field to given value.

### HasEapClientPwd

`func (o *Profile) HasEapClientPwd() bool`

HasEapClientPwd returns a boolean if a field has been set.

### GetEapIdentity

`func (o *Profile) GetEapIdentity() string`

GetEapIdentity returns the EapIdentity field if non-nil, zero value otherwise.

### GetEapIdentityOk

`func (o *Profile) GetEapIdentityOk() (*string, bool)`

GetEapIdentityOk returns a tuple with the EapIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapIdentity

`func (o *Profile) SetEapIdentity(v string)`

SetEapIdentity sets EapIdentity field to given value.

### HasEapIdentity

`func (o *Profile) HasEapIdentity() bool`

HasEapIdentity returns a boolean if a field has been set.

### GetEapIntermediateCert

`func (o *Profile) GetEapIntermediateCert() File`

GetEapIntermediateCert returns the EapIntermediateCert field if non-nil, zero value otherwise.

### GetEapIntermediateCertOk

`func (o *Profile) GetEapIntermediateCertOk() (*File, bool)`

GetEapIntermediateCertOk returns a tuple with the EapIntermediateCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapIntermediateCert

`func (o *Profile) SetEapIntermediateCert(v File)`

SetEapIntermediateCert sets EapIntermediateCert field to given value.

### HasEapIntermediateCert

`func (o *Profile) HasEapIntermediateCert() bool`

HasEapIntermediateCert returns a boolean if a field has been set.

### GetEapPwd

`func (o *Profile) GetEapPwd() string`

GetEapPwd returns the EapPwd field if non-nil, zero value otherwise.

### GetEapPwdOk

`func (o *Profile) GetEapPwdOk() (*string, bool)`

GetEapPwdOk returns a tuple with the EapPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapPwd

`func (o *Profile) SetEapPwd(v string)`

SetEapPwd sets EapPwd field to given value.

### HasEapPwd

`func (o *Profile) HasEapPwd() bool`

HasEapPwd returns a boolean if a field has been set.

### GetHiddenSsid

`func (o *Profile) GetHiddenSsid() bool`

GetHiddenSsid returns the HiddenSsid field if non-nil, zero value otherwise.

### GetHiddenSsidOk

`func (o *Profile) GetHiddenSsidOk() (*bool, bool)`

GetHiddenSsidOk returns a tuple with the HiddenSsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenSsid

`func (o *Profile) SetHiddenSsid(v bool)`

SetHiddenSsid sets HiddenSsid field to given value.

### HasHiddenSsid

`func (o *Profile) HasHiddenSsid() bool`

HasHiddenSsid returns a boolean if a field has been set.

### GetName

`func (o *Profile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Profile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Profile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Profile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPsk

`func (o *Profile) GetPsk() string`

GetPsk returns the Psk field if non-nil, zero value otherwise.

### GetPskOk

`func (o *Profile) GetPskOk() (*string, bool)`

GetPskOk returns a tuple with the Psk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsk

`func (o *Profile) SetPsk(v string)`

SetPsk sets Psk field to given value.

### HasPsk

`func (o *Profile) HasPsk() bool`

HasPsk returns a boolean if a field has been set.

### GetSsid

`func (o *Profile) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *Profile) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *Profile) SetSsid(v string)`

SetSsid sets Ssid field to given value.


### GetWsec

`func (o *Profile) GetWsec() string`

GetWsec returns the Wsec field if non-nil, zero value otherwise.

### GetWsecOk

`func (o *Profile) GetWsecOk() (*string, bool)`

GetWsecOk returns a tuple with the Wsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWsec

`func (o *Profile) SetWsec(v string)`

SetWsec sets Wsec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


