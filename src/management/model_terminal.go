/*
Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// checks if the Terminal type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Terminal{}

// Terminal struct for Terminal
type Terminal struct {
	// The [assignment status](https://docs.adyen.com/point-of-sale/automating-terminal-management/assign-terminals-api) of the terminal. If true, the terminal is assigned. If false, the terminal is in inventory and can't be boarded.
	Assigned *bool `json:"assigned,omitempty"`
	// The Bluetooth IP address of the terminal.
	BluetoothIp *string `json:"bluetoothIp,omitempty"`
	// The Bluetooth MAC address of the terminal.
	BluetoothMac *string `json:"bluetoothMac,omitempty"`
	// The city where the terminal is located.
	City *string `json:"city,omitempty"`
	// The company account that the terminal is associated with. If this is the only account level shown in the response, the terminal is assigned to the inventory of the company account.
	CompanyAccount *string `json:"companyAccount,omitempty"`
	// The country code of the country where the terminal is located.
	CountryCode *string `json:"countryCode,omitempty"`
	// The model name of the terminal.
	DeviceModel *string `json:"deviceModel,omitempty"`
	// The ethernet IP address of the terminal.
	EthernetIp *string `json:"ethernetIp,omitempty"`
	// The ethernet MAC address of the terminal.
	EthernetMac *string `json:"ethernetMac,omitempty"`
	// The software release currently in use on the terminal.
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`
	// The integrated circuit card identifier (ICCID) of the SIM card in the terminal.
	Iccid *string `json:"iccid,omitempty"`
	// The unique identifier of the terminal.
	Id *string `json:"id,omitempty"`
	// Date and time of the last activity on the terminal. Not included when the last activity was more than 14 days ago.
	LastActivityDateTime *time.Time `json:"lastActivityDateTime,omitempty"`
	// Date and time of the last transaction on the terminal. Not included when the last transaction was more than 14 days ago.
	LastTransactionDateTime *time.Time `json:"lastTransactionDateTime,omitempty"`
	// The Ethernet link negotiation that the terminal uses:  - `auto`: Auto-negotiation  - `100full`: 100 Mbps full duplex
	LinkNegotiation *string `json:"linkNegotiation,omitempty"`
	// The serial number of the terminal.
	SerialNumber *string `json:"serialNumber,omitempty"`
	// On a terminal that supports 3G or 4G connectivity, indicates the status of the SIM card in the terminal: ACTIVE or INVENTORY.
	SimStatus *string `json:"simStatus,omitempty"`
	// Indicates when the terminal was last online, whether the terminal is being reassigned, or whether the terminal is turned off. If the terminal was last online more that a week ago, it is also shown as turned off.
	Status *string `json:"status,omitempty"`
	// The status of the store that the terminal is assigned to.
	StoreStatus *string `json:"storeStatus,omitempty"`
	// The terminal's IP address in your Wi-Fi network.
	WifiIp *string `json:"wifiIp,omitempty"`
	// The terminal's MAC address in your Wi-Fi network.
	WifiMac *string `json:"wifiMac,omitempty"`
	// The SSID of the Wi-Fi network that your terminal is connected to.
	WifiSsid *string `json:"wifiSsid,omitempty"`
}

// NewTerminal instantiates a new Terminal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminal() *Terminal {
	this := Terminal{}
	return &this
}

// NewTerminalWithDefaults instantiates a new Terminal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalWithDefaults() *Terminal {
	this := Terminal{}
	return &this
}

// GetAssigned returns the Assigned field value if set, zero value otherwise.
func (o *Terminal) GetAssigned() bool {
	if o == nil || common.IsNil(o.Assigned) {
		var ret bool
		return ret
	}
	return *o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetAssignedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Assigned) {
		return nil, false
	}
	return o.Assigned, true
}

// HasAssigned returns a boolean if a field has been set.
func (o *Terminal) HasAssigned() bool {
	if o != nil && !common.IsNil(o.Assigned) {
		return true
	}

	return false
}

// SetAssigned gets a reference to the given bool and assigns it to the Assigned field.
func (o *Terminal) SetAssigned(v bool) {
	o.Assigned = &v
}

// GetBluetoothIp returns the BluetoothIp field value if set, zero value otherwise.
func (o *Terminal) GetBluetoothIp() string {
	if o == nil || common.IsNil(o.BluetoothIp) {
		var ret string
		return ret
	}
	return *o.BluetoothIp
}

// GetBluetoothIpOk returns a tuple with the BluetoothIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetBluetoothIpOk() (*string, bool) {
	if o == nil || common.IsNil(o.BluetoothIp) {
		return nil, false
	}
	return o.BluetoothIp, true
}

// HasBluetoothIp returns a boolean if a field has been set.
func (o *Terminal) HasBluetoothIp() bool {
	if o != nil && !common.IsNil(o.BluetoothIp) {
		return true
	}

	return false
}

// SetBluetoothIp gets a reference to the given string and assigns it to the BluetoothIp field.
func (o *Terminal) SetBluetoothIp(v string) {
	o.BluetoothIp = &v
}

// GetBluetoothMac returns the BluetoothMac field value if set, zero value otherwise.
func (o *Terminal) GetBluetoothMac() string {
	if o == nil || common.IsNil(o.BluetoothMac) {
		var ret string
		return ret
	}
	return *o.BluetoothMac
}

// GetBluetoothMacOk returns a tuple with the BluetoothMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetBluetoothMacOk() (*string, bool) {
	if o == nil || common.IsNil(o.BluetoothMac) {
		return nil, false
	}
	return o.BluetoothMac, true
}

// HasBluetoothMac returns a boolean if a field has been set.
func (o *Terminal) HasBluetoothMac() bool {
	if o != nil && !common.IsNil(o.BluetoothMac) {
		return true
	}

	return false
}

// SetBluetoothMac gets a reference to the given string and assigns it to the BluetoothMac field.
func (o *Terminal) SetBluetoothMac(v string) {
	o.BluetoothMac = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *Terminal) GetCity() string {
	if o == nil || common.IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetCityOk() (*string, bool) {
	if o == nil || common.IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *Terminal) HasCity() bool {
	if o != nil && !common.IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *Terminal) SetCity(v string) {
	o.City = &v
}

// GetCompanyAccount returns the CompanyAccount field value if set, zero value otherwise.
func (o *Terminal) GetCompanyAccount() string {
	if o == nil || common.IsNil(o.CompanyAccount) {
		var ret string
		return ret
	}
	return *o.CompanyAccount
}

// GetCompanyAccountOk returns a tuple with the CompanyAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetCompanyAccountOk() (*string, bool) {
	if o == nil || common.IsNil(o.CompanyAccount) {
		return nil, false
	}
	return o.CompanyAccount, true
}

// HasCompanyAccount returns a boolean if a field has been set.
func (o *Terminal) HasCompanyAccount() bool {
	if o != nil && !common.IsNil(o.CompanyAccount) {
		return true
	}

	return false
}

// SetCompanyAccount gets a reference to the given string and assigns it to the CompanyAccount field.
func (o *Terminal) SetCompanyAccount(v string) {
	o.CompanyAccount = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *Terminal) GetCountryCode() string {
	if o == nil || common.IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetCountryCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *Terminal) HasCountryCode() bool {
	if o != nil && !common.IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *Terminal) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetDeviceModel returns the DeviceModel field value if set, zero value otherwise.
func (o *Terminal) GetDeviceModel() string {
	if o == nil || common.IsNil(o.DeviceModel) {
		var ret string
		return ret
	}
	return *o.DeviceModel
}

// GetDeviceModelOk returns a tuple with the DeviceModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetDeviceModelOk() (*string, bool) {
	if o == nil || common.IsNil(o.DeviceModel) {
		return nil, false
	}
	return o.DeviceModel, true
}

// HasDeviceModel returns a boolean if a field has been set.
func (o *Terminal) HasDeviceModel() bool {
	if o != nil && !common.IsNil(o.DeviceModel) {
		return true
	}

	return false
}

// SetDeviceModel gets a reference to the given string and assigns it to the DeviceModel field.
func (o *Terminal) SetDeviceModel(v string) {
	o.DeviceModel = &v
}

// GetEthernetIp returns the EthernetIp field value if set, zero value otherwise.
func (o *Terminal) GetEthernetIp() string {
	if o == nil || common.IsNil(o.EthernetIp) {
		var ret string
		return ret
	}
	return *o.EthernetIp
}

// GetEthernetIpOk returns a tuple with the EthernetIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetEthernetIpOk() (*string, bool) {
	if o == nil || common.IsNil(o.EthernetIp) {
		return nil, false
	}
	return o.EthernetIp, true
}

// HasEthernetIp returns a boolean if a field has been set.
func (o *Terminal) HasEthernetIp() bool {
	if o != nil && !common.IsNil(o.EthernetIp) {
		return true
	}

	return false
}

// SetEthernetIp gets a reference to the given string and assigns it to the EthernetIp field.
func (o *Terminal) SetEthernetIp(v string) {
	o.EthernetIp = &v
}

// GetEthernetMac returns the EthernetMac field value if set, zero value otherwise.
func (o *Terminal) GetEthernetMac() string {
	if o == nil || common.IsNil(o.EthernetMac) {
		var ret string
		return ret
	}
	return *o.EthernetMac
}

// GetEthernetMacOk returns a tuple with the EthernetMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetEthernetMacOk() (*string, bool) {
	if o == nil || common.IsNil(o.EthernetMac) {
		return nil, false
	}
	return o.EthernetMac, true
}

// HasEthernetMac returns a boolean if a field has been set.
func (o *Terminal) HasEthernetMac() bool {
	if o != nil && !common.IsNil(o.EthernetMac) {
		return true
	}

	return false
}

// SetEthernetMac gets a reference to the given string and assigns it to the EthernetMac field.
func (o *Terminal) SetEthernetMac(v string) {
	o.EthernetMac = &v
}

// GetFirmwareVersion returns the FirmwareVersion field value if set, zero value otherwise.
func (o *Terminal) GetFirmwareVersion() string {
	if o == nil || common.IsNil(o.FirmwareVersion) {
		var ret string
		return ret
	}
	return *o.FirmwareVersion
}

// GetFirmwareVersionOk returns a tuple with the FirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetFirmwareVersionOk() (*string, bool) {
	if o == nil || common.IsNil(o.FirmwareVersion) {
		return nil, false
	}
	return o.FirmwareVersion, true
}

// HasFirmwareVersion returns a boolean if a field has been set.
func (o *Terminal) HasFirmwareVersion() bool {
	if o != nil && !common.IsNil(o.FirmwareVersion) {
		return true
	}

	return false
}

// SetFirmwareVersion gets a reference to the given string and assigns it to the FirmwareVersion field.
func (o *Terminal) SetFirmwareVersion(v string) {
	o.FirmwareVersion = &v
}

// GetIccid returns the Iccid field value if set, zero value otherwise.
func (o *Terminal) GetIccid() string {
	if o == nil || common.IsNil(o.Iccid) {
		var ret string
		return ret
	}
	return *o.Iccid
}

// GetIccidOk returns a tuple with the Iccid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetIccidOk() (*string, bool) {
	if o == nil || common.IsNil(o.Iccid) {
		return nil, false
	}
	return o.Iccid, true
}

// HasIccid returns a boolean if a field has been set.
func (o *Terminal) HasIccid() bool {
	if o != nil && !common.IsNil(o.Iccid) {
		return true
	}

	return false
}

// SetIccid gets a reference to the given string and assigns it to the Iccid field.
func (o *Terminal) SetIccid(v string) {
	o.Iccid = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Terminal) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Terminal) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Terminal) SetId(v string) {
	o.Id = &v
}

// GetLastActivityDateTime returns the LastActivityDateTime field value if set, zero value otherwise.
func (o *Terminal) GetLastActivityDateTime() time.Time {
	if o == nil || common.IsNil(o.LastActivityDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastActivityDateTime
}

// GetLastActivityDateTimeOk returns a tuple with the LastActivityDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetLastActivityDateTimeOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.LastActivityDateTime) {
		return nil, false
	}
	return o.LastActivityDateTime, true
}

// HasLastActivityDateTime returns a boolean if a field has been set.
func (o *Terminal) HasLastActivityDateTime() bool {
	if o != nil && !common.IsNil(o.LastActivityDateTime) {
		return true
	}

	return false
}

// SetLastActivityDateTime gets a reference to the given time.Time and assigns it to the LastActivityDateTime field.
func (o *Terminal) SetLastActivityDateTime(v time.Time) {
	o.LastActivityDateTime = &v
}

// GetLastTransactionDateTime returns the LastTransactionDateTime field value if set, zero value otherwise.
func (o *Terminal) GetLastTransactionDateTime() time.Time {
	if o == nil || common.IsNil(o.LastTransactionDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastTransactionDateTime
}

// GetLastTransactionDateTimeOk returns a tuple with the LastTransactionDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetLastTransactionDateTimeOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.LastTransactionDateTime) {
		return nil, false
	}
	return o.LastTransactionDateTime, true
}

// HasLastTransactionDateTime returns a boolean if a field has been set.
func (o *Terminal) HasLastTransactionDateTime() bool {
	if o != nil && !common.IsNil(o.LastTransactionDateTime) {
		return true
	}

	return false
}

// SetLastTransactionDateTime gets a reference to the given time.Time and assigns it to the LastTransactionDateTime field.
func (o *Terminal) SetLastTransactionDateTime(v time.Time) {
	o.LastTransactionDateTime = &v
}

// GetLinkNegotiation returns the LinkNegotiation field value if set, zero value otherwise.
func (o *Terminal) GetLinkNegotiation() string {
	if o == nil || common.IsNil(o.LinkNegotiation) {
		var ret string
		return ret
	}
	return *o.LinkNegotiation
}

// GetLinkNegotiationOk returns a tuple with the LinkNegotiation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetLinkNegotiationOk() (*string, bool) {
	if o == nil || common.IsNil(o.LinkNegotiation) {
		return nil, false
	}
	return o.LinkNegotiation, true
}

// HasLinkNegotiation returns a boolean if a field has been set.
func (o *Terminal) HasLinkNegotiation() bool {
	if o != nil && !common.IsNil(o.LinkNegotiation) {
		return true
	}

	return false
}

// SetLinkNegotiation gets a reference to the given string and assigns it to the LinkNegotiation field.
func (o *Terminal) SetLinkNegotiation(v string) {
	o.LinkNegotiation = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *Terminal) GetSerialNumber() string {
	if o == nil || common.IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetSerialNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *Terminal) HasSerialNumber() bool {
	if o != nil && !common.IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *Terminal) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSimStatus returns the SimStatus field value if set, zero value otherwise.
func (o *Terminal) GetSimStatus() string {
	if o == nil || common.IsNil(o.SimStatus) {
		var ret string
		return ret
	}
	return *o.SimStatus
}

// GetSimStatusOk returns a tuple with the SimStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetSimStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.SimStatus) {
		return nil, false
	}
	return o.SimStatus, true
}

// HasSimStatus returns a boolean if a field has been set.
func (o *Terminal) HasSimStatus() bool {
	if o != nil && !common.IsNil(o.SimStatus) {
		return true
	}

	return false
}

// SetSimStatus gets a reference to the given string and assigns it to the SimStatus field.
func (o *Terminal) SetSimStatus(v string) {
	o.SimStatus = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Terminal) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Terminal) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Terminal) SetStatus(v string) {
	o.Status = &v
}

// GetStoreStatus returns the StoreStatus field value if set, zero value otherwise.
func (o *Terminal) GetStoreStatus() string {
	if o == nil || common.IsNil(o.StoreStatus) {
		var ret string
		return ret
	}
	return *o.StoreStatus
}

// GetStoreStatusOk returns a tuple with the StoreStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetStoreStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.StoreStatus) {
		return nil, false
	}
	return o.StoreStatus, true
}

// HasStoreStatus returns a boolean if a field has been set.
func (o *Terminal) HasStoreStatus() bool {
	if o != nil && !common.IsNil(o.StoreStatus) {
		return true
	}

	return false
}

// SetStoreStatus gets a reference to the given string and assigns it to the StoreStatus field.
func (o *Terminal) SetStoreStatus(v string) {
	o.StoreStatus = &v
}

// GetWifiIp returns the WifiIp field value if set, zero value otherwise.
func (o *Terminal) GetWifiIp() string {
	if o == nil || common.IsNil(o.WifiIp) {
		var ret string
		return ret
	}
	return *o.WifiIp
}

// GetWifiIpOk returns a tuple with the WifiIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetWifiIpOk() (*string, bool) {
	if o == nil || common.IsNil(o.WifiIp) {
		return nil, false
	}
	return o.WifiIp, true
}

// HasWifiIp returns a boolean if a field has been set.
func (o *Terminal) HasWifiIp() bool {
	if o != nil && !common.IsNil(o.WifiIp) {
		return true
	}

	return false
}

// SetWifiIp gets a reference to the given string and assigns it to the WifiIp field.
func (o *Terminal) SetWifiIp(v string) {
	o.WifiIp = &v
}

// GetWifiMac returns the WifiMac field value if set, zero value otherwise.
func (o *Terminal) GetWifiMac() string {
	if o == nil || common.IsNil(o.WifiMac) {
		var ret string
		return ret
	}
	return *o.WifiMac
}

// GetWifiMacOk returns a tuple with the WifiMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetWifiMacOk() (*string, bool) {
	if o == nil || common.IsNil(o.WifiMac) {
		return nil, false
	}
	return o.WifiMac, true
}

// HasWifiMac returns a boolean if a field has been set.
func (o *Terminal) HasWifiMac() bool {
	if o != nil && !common.IsNil(o.WifiMac) {
		return true
	}

	return false
}

// SetWifiMac gets a reference to the given string and assigns it to the WifiMac field.
func (o *Terminal) SetWifiMac(v string) {
	o.WifiMac = &v
}

// GetWifiSsid returns the WifiSsid field value if set, zero value otherwise.
func (o *Terminal) GetWifiSsid() string {
	if o == nil || common.IsNil(o.WifiSsid) {
		var ret string
		return ret
	}
	return *o.WifiSsid
}

// GetWifiSsidOk returns a tuple with the WifiSsid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetWifiSsidOk() (*string, bool) {
	if o == nil || common.IsNil(o.WifiSsid) {
		return nil, false
	}
	return o.WifiSsid, true
}

// HasWifiSsid returns a boolean if a field has been set.
func (o *Terminal) HasWifiSsid() bool {
	if o != nil && !common.IsNil(o.WifiSsid) {
		return true
	}

	return false
}

// SetWifiSsid gets a reference to the given string and assigns it to the WifiSsid field.
func (o *Terminal) SetWifiSsid(v string) {
	o.WifiSsid = &v
}

func (o Terminal) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Terminal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Assigned) {
		toSerialize["assigned"] = o.Assigned
	}
	if !common.IsNil(o.BluetoothIp) {
		toSerialize["bluetoothIp"] = o.BluetoothIp
	}
	if !common.IsNil(o.BluetoothMac) {
		toSerialize["bluetoothMac"] = o.BluetoothMac
	}
	if !common.IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !common.IsNil(o.CompanyAccount) {
		toSerialize["companyAccount"] = o.CompanyAccount
	}
	if !common.IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !common.IsNil(o.DeviceModel) {
		toSerialize["deviceModel"] = o.DeviceModel
	}
	if !common.IsNil(o.EthernetIp) {
		toSerialize["ethernetIp"] = o.EthernetIp
	}
	if !common.IsNil(o.EthernetMac) {
		toSerialize["ethernetMac"] = o.EthernetMac
	}
	if !common.IsNil(o.FirmwareVersion) {
		toSerialize["firmwareVersion"] = o.FirmwareVersion
	}
	if !common.IsNil(o.Iccid) {
		toSerialize["iccid"] = o.Iccid
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.LastActivityDateTime) {
		toSerialize["lastActivityDateTime"] = o.LastActivityDateTime
	}
	if !common.IsNil(o.LastTransactionDateTime) {
		toSerialize["lastTransactionDateTime"] = o.LastTransactionDateTime
	}
	if !common.IsNil(o.LinkNegotiation) {
		toSerialize["linkNegotiation"] = o.LinkNegotiation
	}
	if !common.IsNil(o.SerialNumber) {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if !common.IsNil(o.SimStatus) {
		toSerialize["simStatus"] = o.SimStatus
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.StoreStatus) {
		toSerialize["storeStatus"] = o.StoreStatus
	}
	if !common.IsNil(o.WifiIp) {
		toSerialize["wifiIp"] = o.WifiIp
	}
	if !common.IsNil(o.WifiMac) {
		toSerialize["wifiMac"] = o.WifiMac
	}
	if !common.IsNil(o.WifiSsid) {
		toSerialize["wifiSsid"] = o.WifiSsid
	}
	return toSerialize, nil
}

type NullableTerminal struct {
	value *Terminal
	isSet bool
}

func (v NullableTerminal) Get() *Terminal {
	return v.value
}

func (v *NullableTerminal) Set(val *Terminal) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminal) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminal(val *Terminal) *NullableTerminal {
	return &NullableTerminal{value: val, isSet: true}
}

func (v NullableTerminal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}