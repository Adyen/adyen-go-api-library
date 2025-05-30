/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the Profile type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Profile{}

// Profile struct for Profile
type Profile struct {
	// The type of Wi-Fi network. Possible values: **wpa-psk**, **wpa2-psk**, **wpa-eap**, **wpa2-eap**.
	AuthType string `json:"authType"`
	// Indicates whether to automatically select the best authentication method available. Does not work on older terminal models.
	AutoWifi *bool `json:"autoWifi,omitempty"`
	// Use **infra** for infrastructure-based networks. This applies to most networks. Use **adhoc** only if the communication is p2p-based between base stations.
	BssType string `json:"bssType"`
	// The channel number of the Wi-Fi network. The recommended setting is **0** for automatic channel selection.
	Channel *int32 `json:"channel,omitempty"`
	// Indicates whether this is your preferred wireless network. If **true**, the terminal will try connecting to this network first.
	DefaultProfile *bool `json:"defaultProfile,omitempty"`
	// Specifies the server domain name for EAP-TLS and EAP-PEAP WiFi profiles on Android 11 and above.
	DomainSuffix *string `json:"domainSuffix,omitempty"`
	// For `authType` **wpa-eap** or **wpa2-eap**. Possible values: **tls**, **peap**, **leap**, **fast**
	Eap           *string `json:"eap,omitempty"`
	EapCaCert     *File   `json:"eapCaCert,omitempty"`
	EapClientCert *File   `json:"eapClientCert,omitempty"`
	EapClientKey  *File   `json:"eapClientKey,omitempty"`
	// For `eap` **tls**. The password of the RSA key file, if that file is password-protected.
	EapClientPwd *string `json:"eapClientPwd,omitempty"`
	// For `authType` **wpa-eap** or **wpa2-eap**. The EAP-PEAP username from your MS-CHAP account. Must match the configuration of your RADIUS server.
	EapIdentity         *string `json:"eapIdentity,omitempty"`
	EapIntermediateCert *File   `json:"eapIntermediateCert,omitempty"`
	// For `eap` **peap**. The EAP-PEAP password from your MS-CHAP account. Must match the configuration of your RADIUS server.
	EapPwd *string `json:"eapPwd,omitempty"`
	// Indicates if the network doesn't broadcast its SSID. Mandatory for Android terminals, because these terminals rely on this setting to be able to connect to any network.
	HiddenSsid *bool `json:"hiddenSsid,omitempty"`
	// Your name for the Wi-Fi profile.
	Name *string `json:"name,omitempty"`
	// For `authType` **wpa-psk or **wpa2-psk**. The password to the wireless network.
	Psk *string `json:"psk,omitempty"`
	// The name of the wireless network.
	Ssid string `json:"ssid"`
	// The type of encryption. Possible values: **auto**, **ccmp** (recommended), **tkip**
	Wsec string `json:"wsec"`
}

// NewProfile instantiates a new Profile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfile(authType string, bssType string, ssid string, wsec string) *Profile {
	this := Profile{}
	this.AuthType = authType
	this.BssType = bssType
	this.Ssid = ssid
	this.Wsec = wsec
	return &this
}

// NewProfileWithDefaults instantiates a new Profile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileWithDefaults() *Profile {
	this := Profile{}
	return &this
}

// GetAuthType returns the AuthType field value
func (o *Profile) GetAuthType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *Profile) GetAuthTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value
func (o *Profile) SetAuthType(v string) {
	o.AuthType = v
}

// GetAutoWifi returns the AutoWifi field value if set, zero value otherwise.
func (o *Profile) GetAutoWifi() bool {
	if o == nil || common.IsNil(o.AutoWifi) {
		var ret bool
		return ret
	}
	return *o.AutoWifi
}

// GetAutoWifiOk returns a tuple with the AutoWifi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetAutoWifiOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AutoWifi) {
		return nil, false
	}
	return o.AutoWifi, true
}

// HasAutoWifi returns a boolean if a field has been set.
func (o *Profile) HasAutoWifi() bool {
	if o != nil && !common.IsNil(o.AutoWifi) {
		return true
	}

	return false
}

// SetAutoWifi gets a reference to the given bool and assigns it to the AutoWifi field.
func (o *Profile) SetAutoWifi(v bool) {
	o.AutoWifi = &v
}

// GetBssType returns the BssType field value
func (o *Profile) GetBssType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BssType
}

// GetBssTypeOk returns a tuple with the BssType field value
// and a boolean to check if the value has been set.
func (o *Profile) GetBssTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BssType, true
}

// SetBssType sets field value
func (o *Profile) SetBssType(v string) {
	o.BssType = v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *Profile) GetChannel() int32 {
	if o == nil || common.IsNil(o.Channel) {
		var ret int32
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetChannelOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *Profile) HasChannel() bool {
	if o != nil && !common.IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given int32 and assigns it to the Channel field.
func (o *Profile) SetChannel(v int32) {
	o.Channel = &v
}

// GetDefaultProfile returns the DefaultProfile field value if set, zero value otherwise.
func (o *Profile) GetDefaultProfile() bool {
	if o == nil || common.IsNil(o.DefaultProfile) {
		var ret bool
		return ret
	}
	return *o.DefaultProfile
}

// GetDefaultProfileOk returns a tuple with the DefaultProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetDefaultProfileOk() (*bool, bool) {
	if o == nil || common.IsNil(o.DefaultProfile) {
		return nil, false
	}
	return o.DefaultProfile, true
}

// HasDefaultProfile returns a boolean if a field has been set.
func (o *Profile) HasDefaultProfile() bool {
	if o != nil && !common.IsNil(o.DefaultProfile) {
		return true
	}

	return false
}

// SetDefaultProfile gets a reference to the given bool and assigns it to the DefaultProfile field.
func (o *Profile) SetDefaultProfile(v bool) {
	o.DefaultProfile = &v
}

// GetDomainSuffix returns the DomainSuffix field value if set, zero value otherwise.
func (o *Profile) GetDomainSuffix() string {
	if o == nil || common.IsNil(o.DomainSuffix) {
		var ret string
		return ret
	}
	return *o.DomainSuffix
}

// GetDomainSuffixOk returns a tuple with the DomainSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetDomainSuffixOk() (*string, bool) {
	if o == nil || common.IsNil(o.DomainSuffix) {
		return nil, false
	}
	return o.DomainSuffix, true
}

// HasDomainSuffix returns a boolean if a field has been set.
func (o *Profile) HasDomainSuffix() bool {
	if o != nil && !common.IsNil(o.DomainSuffix) {
		return true
	}

	return false
}

// SetDomainSuffix gets a reference to the given string and assigns it to the DomainSuffix field.
func (o *Profile) SetDomainSuffix(v string) {
	o.DomainSuffix = &v
}

// GetEap returns the Eap field value if set, zero value otherwise.
func (o *Profile) GetEap() string {
	if o == nil || common.IsNil(o.Eap) {
		var ret string
		return ret
	}
	return *o.Eap
}

// GetEapOk returns a tuple with the Eap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetEapOk() (*string, bool) {
	if o == nil || common.IsNil(o.Eap) {
		return nil, false
	}
	return o.Eap, true
}

// HasEap returns a boolean if a field has been set.
func (o *Profile) HasEap() bool {
	if o != nil && !common.IsNil(o.Eap) {
		return true
	}

	return false
}

// SetEap gets a reference to the given string and assigns it to the Eap field.
func (o *Profile) SetEap(v string) {
	o.Eap = &v
}

// GetEapCaCert returns the EapCaCert field value if set, zero value otherwise.
func (o *Profile) GetEapCaCert() File {
	if o == nil || common.IsNil(o.EapCaCert) {
		var ret File
		return ret
	}
	return *o.EapCaCert
}

// GetEapCaCertOk returns a tuple with the EapCaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetEapCaCertOk() (*File, bool) {
	if o == nil || common.IsNil(o.EapCaCert) {
		return nil, false
	}
	return o.EapCaCert, true
}

// HasEapCaCert returns a boolean if a field has been set.
func (o *Profile) HasEapCaCert() bool {
	if o != nil && !common.IsNil(o.EapCaCert) {
		return true
	}

	return false
}

// SetEapCaCert gets a reference to the given File and assigns it to the EapCaCert field.
func (o *Profile) SetEapCaCert(v File) {
	o.EapCaCert = &v
}

// GetEapClientCert returns the EapClientCert field value if set, zero value otherwise.
func (o *Profile) GetEapClientCert() File {
	if o == nil || common.IsNil(o.EapClientCert) {
		var ret File
		return ret
	}
	return *o.EapClientCert
}

// GetEapClientCertOk returns a tuple with the EapClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetEapClientCertOk() (*File, bool) {
	if o == nil || common.IsNil(o.EapClientCert) {
		return nil, false
	}
	return o.EapClientCert, true
}

// HasEapClientCert returns a boolean if a field has been set.
func (o *Profile) HasEapClientCert() bool {
	if o != nil && !common.IsNil(o.EapClientCert) {
		return true
	}

	return false
}

// SetEapClientCert gets a reference to the given File and assigns it to the EapClientCert field.
func (o *Profile) SetEapClientCert(v File) {
	o.EapClientCert = &v
}

// GetEapClientKey returns the EapClientKey field value if set, zero value otherwise.
func (o *Profile) GetEapClientKey() File {
	if o == nil || common.IsNil(o.EapClientKey) {
		var ret File
		return ret
	}
	return *o.EapClientKey
}

// GetEapClientKeyOk returns a tuple with the EapClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetEapClientKeyOk() (*File, bool) {
	if o == nil || common.IsNil(o.EapClientKey) {
		return nil, false
	}
	return o.EapClientKey, true
}

// HasEapClientKey returns a boolean if a field has been set.
func (o *Profile) HasEapClientKey() bool {
	if o != nil && !common.IsNil(o.EapClientKey) {
		return true
	}

	return false
}

// SetEapClientKey gets a reference to the given File and assigns it to the EapClientKey field.
func (o *Profile) SetEapClientKey(v File) {
	o.EapClientKey = &v
}

// GetEapClientPwd returns the EapClientPwd field value if set, zero value otherwise.
func (o *Profile) GetEapClientPwd() string {
	if o == nil || common.IsNil(o.EapClientPwd) {
		var ret string
		return ret
	}
	return *o.EapClientPwd
}

// GetEapClientPwdOk returns a tuple with the EapClientPwd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetEapClientPwdOk() (*string, bool) {
	if o == nil || common.IsNil(o.EapClientPwd) {
		return nil, false
	}
	return o.EapClientPwd, true
}

// HasEapClientPwd returns a boolean if a field has been set.
func (o *Profile) HasEapClientPwd() bool {
	if o != nil && !common.IsNil(o.EapClientPwd) {
		return true
	}

	return false
}

// SetEapClientPwd gets a reference to the given string and assigns it to the EapClientPwd field.
func (o *Profile) SetEapClientPwd(v string) {
	o.EapClientPwd = &v
}

// GetEapIdentity returns the EapIdentity field value if set, zero value otherwise.
func (o *Profile) GetEapIdentity() string {
	if o == nil || common.IsNil(o.EapIdentity) {
		var ret string
		return ret
	}
	return *o.EapIdentity
}

// GetEapIdentityOk returns a tuple with the EapIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetEapIdentityOk() (*string, bool) {
	if o == nil || common.IsNil(o.EapIdentity) {
		return nil, false
	}
	return o.EapIdentity, true
}

// HasEapIdentity returns a boolean if a field has been set.
func (o *Profile) HasEapIdentity() bool {
	if o != nil && !common.IsNil(o.EapIdentity) {
		return true
	}

	return false
}

// SetEapIdentity gets a reference to the given string and assigns it to the EapIdentity field.
func (o *Profile) SetEapIdentity(v string) {
	o.EapIdentity = &v
}

// GetEapIntermediateCert returns the EapIntermediateCert field value if set, zero value otherwise.
func (o *Profile) GetEapIntermediateCert() File {
	if o == nil || common.IsNil(o.EapIntermediateCert) {
		var ret File
		return ret
	}
	return *o.EapIntermediateCert
}

// GetEapIntermediateCertOk returns a tuple with the EapIntermediateCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetEapIntermediateCertOk() (*File, bool) {
	if o == nil || common.IsNil(o.EapIntermediateCert) {
		return nil, false
	}
	return o.EapIntermediateCert, true
}

// HasEapIntermediateCert returns a boolean if a field has been set.
func (o *Profile) HasEapIntermediateCert() bool {
	if o != nil && !common.IsNil(o.EapIntermediateCert) {
		return true
	}

	return false
}

// SetEapIntermediateCert gets a reference to the given File and assigns it to the EapIntermediateCert field.
func (o *Profile) SetEapIntermediateCert(v File) {
	o.EapIntermediateCert = &v
}

// GetEapPwd returns the EapPwd field value if set, zero value otherwise.
func (o *Profile) GetEapPwd() string {
	if o == nil || common.IsNil(o.EapPwd) {
		var ret string
		return ret
	}
	return *o.EapPwd
}

// GetEapPwdOk returns a tuple with the EapPwd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetEapPwdOk() (*string, bool) {
	if o == nil || common.IsNil(o.EapPwd) {
		return nil, false
	}
	return o.EapPwd, true
}

// HasEapPwd returns a boolean if a field has been set.
func (o *Profile) HasEapPwd() bool {
	if o != nil && !common.IsNil(o.EapPwd) {
		return true
	}

	return false
}

// SetEapPwd gets a reference to the given string and assigns it to the EapPwd field.
func (o *Profile) SetEapPwd(v string) {
	o.EapPwd = &v
}

// GetHiddenSsid returns the HiddenSsid field value if set, zero value otherwise.
func (o *Profile) GetHiddenSsid() bool {
	if o == nil || common.IsNil(o.HiddenSsid) {
		var ret bool
		return ret
	}
	return *o.HiddenSsid
}

// GetHiddenSsidOk returns a tuple with the HiddenSsid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetHiddenSsidOk() (*bool, bool) {
	if o == nil || common.IsNil(o.HiddenSsid) {
		return nil, false
	}
	return o.HiddenSsid, true
}

// HasHiddenSsid returns a boolean if a field has been set.
func (o *Profile) HasHiddenSsid() bool {
	if o != nil && !common.IsNil(o.HiddenSsid) {
		return true
	}

	return false
}

// SetHiddenSsid gets a reference to the given bool and assigns it to the HiddenSsid field.
func (o *Profile) SetHiddenSsid(v bool) {
	o.HiddenSsid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Profile) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Profile) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Profile) SetName(v string) {
	o.Name = &v
}

// GetPsk returns the Psk field value if set, zero value otherwise.
func (o *Profile) GetPsk() string {
	if o == nil || common.IsNil(o.Psk) {
		var ret string
		return ret
	}
	return *o.Psk
}

// GetPskOk returns a tuple with the Psk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetPskOk() (*string, bool) {
	if o == nil || common.IsNil(o.Psk) {
		return nil, false
	}
	return o.Psk, true
}

// HasPsk returns a boolean if a field has been set.
func (o *Profile) HasPsk() bool {
	if o != nil && !common.IsNil(o.Psk) {
		return true
	}

	return false
}

// SetPsk gets a reference to the given string and assigns it to the Psk field.
func (o *Profile) SetPsk(v string) {
	o.Psk = &v
}

// GetSsid returns the Ssid field value
func (o *Profile) GetSsid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value
// and a boolean to check if the value has been set.
func (o *Profile) GetSsidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ssid, true
}

// SetSsid sets field value
func (o *Profile) SetSsid(v string) {
	o.Ssid = v
}

// GetWsec returns the Wsec field value
func (o *Profile) GetWsec() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Wsec
}

// GetWsecOk returns a tuple with the Wsec field value
// and a boolean to check if the value has been set.
func (o *Profile) GetWsecOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Wsec, true
}

// SetWsec sets field value
func (o *Profile) SetWsec(v string) {
	o.Wsec = v
}

func (o Profile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Profile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authType"] = o.AuthType
	if !common.IsNil(o.AutoWifi) {
		toSerialize["autoWifi"] = o.AutoWifi
	}
	toSerialize["bssType"] = o.BssType
	if !common.IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !common.IsNil(o.DefaultProfile) {
		toSerialize["defaultProfile"] = o.DefaultProfile
	}
	if !common.IsNil(o.DomainSuffix) {
		toSerialize["domainSuffix"] = o.DomainSuffix
	}
	if !common.IsNil(o.Eap) {
		toSerialize["eap"] = o.Eap
	}
	if !common.IsNil(o.EapCaCert) {
		toSerialize["eapCaCert"] = o.EapCaCert
	}
	if !common.IsNil(o.EapClientCert) {
		toSerialize["eapClientCert"] = o.EapClientCert
	}
	if !common.IsNil(o.EapClientKey) {
		toSerialize["eapClientKey"] = o.EapClientKey
	}
	if !common.IsNil(o.EapClientPwd) {
		toSerialize["eapClientPwd"] = o.EapClientPwd
	}
	if !common.IsNil(o.EapIdentity) {
		toSerialize["eapIdentity"] = o.EapIdentity
	}
	if !common.IsNil(o.EapIntermediateCert) {
		toSerialize["eapIntermediateCert"] = o.EapIntermediateCert
	}
	if !common.IsNil(o.EapPwd) {
		toSerialize["eapPwd"] = o.EapPwd
	}
	if !common.IsNil(o.HiddenSsid) {
		toSerialize["hiddenSsid"] = o.HiddenSsid
	}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.Psk) {
		toSerialize["psk"] = o.Psk
	}
	toSerialize["ssid"] = o.Ssid
	toSerialize["wsec"] = o.Wsec
	return toSerialize, nil
}

type NullableProfile struct {
	value *Profile
	isSet bool
}

func (v NullableProfile) Get() *Profile {
	return v.value
}

func (v *NullableProfile) Set(val *Profile) {
	v.value = val
	v.isSet = true
}

func (v NullableProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfile(val *Profile) *NullableProfile {
	return &NullableProfile{value: val, isSet: true}
}

func (v NullableProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
