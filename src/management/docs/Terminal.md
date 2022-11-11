# Terminal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assigned** | Pointer to **bool** | The [assignment status](https://docs.adyen.com/point-of-sale/automating-terminal-management/assign-terminals-api) of the terminal. If true, the terminal is assigned. If false, the terminal is in inventory and can&#39;t be boarded. | [optional] 
**BluetoothIp** | Pointer to **string** | The Bluetooth IP address of the terminal. | [optional] 
**BluetoothMac** | Pointer to **string** | The Bluetooth MAC address of the terminal. | [optional] 
**City** | Pointer to **string** | The city where the terminal is located. | [optional] 
**CompanyAccount** | Pointer to **string** | The company account of the terminal. | [optional] 
**CountryCode** | Pointer to **string** | The country code where the terminal is located. | [optional] 
**DeviceModel** | Pointer to **string** | The terminal model of the device. | [optional] 
**EthernetIp** | Pointer to **string** | The ethernet IP address of the terminal. | [optional] 
**EthernetMac** | Pointer to **string** | The ethernet MAC address of the terminal. | [optional] 
**FirmwareVersion** | Pointer to **string** | The firmware Version of the terminal. | [optional] 
**Iccid** | Pointer to **string** | The ICCID number of the cellular communications card. | [optional] 
**Id** | Pointer to **string** | The unique identifier of the terminal. | [optional] 
**LastActivityDateTime** | Pointer to **time.Time** | The last Activity Date and Time of the terminal. | [optional] 
**LastTransactionDateTime** | Pointer to **time.Time** | The last Transaction Date and Time of the terminal. | [optional] 
**LinkNegotiation** | Pointer to **string** | The ethernet link speed of the terminal that was negotiated. | [optional] 
**SerialNumber** | Pointer to **string** | The serial number of the terminal. | [optional] 
**SimStatus** | Pointer to **string** | On a terminal that supports 3G or 4G connectivity, indicates the status of the SIM card in the terminal: ACTIVE or INVENTORY. | [optional] 
**Status** | Pointer to **string** | Indicates when the terminal was last online, whether the terminal is being reassigned, or whether the terminal is turned off. If the terminal was last online more that a week ago, it is also shown as turned off. | [optional] 
**StoreStatus** | Pointer to **string** | The Status of store where the terminal is located. | [optional] 
**WifiIp** | Pointer to **string** | The WiFi IP address of the terminal. | [optional] 
**WifiMac** | Pointer to **string** | The WiFi MAC address of the terminal. | [optional] 
**WifiSsid** | Pointer to **string** | The WIFI SSID of the terminal. | [optional] 

## Methods

### NewTerminal

`func NewTerminal() *Terminal`

NewTerminal instantiates a new Terminal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminalWithDefaults

`func NewTerminalWithDefaults() *Terminal`

NewTerminalWithDefaults instantiates a new Terminal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssigned

`func (o *Terminal) GetAssigned() bool`

GetAssigned returns the Assigned field if non-nil, zero value otherwise.

### GetAssignedOk

`func (o *Terminal) GetAssignedOk() (*bool, bool)`

GetAssignedOk returns a tuple with the Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigned

`func (o *Terminal) SetAssigned(v bool)`

SetAssigned sets Assigned field to given value.

### HasAssigned

`func (o *Terminal) HasAssigned() bool`

HasAssigned returns a boolean if a field has been set.

### GetBluetoothIp

`func (o *Terminal) GetBluetoothIp() string`

GetBluetoothIp returns the BluetoothIp field if non-nil, zero value otherwise.

### GetBluetoothIpOk

`func (o *Terminal) GetBluetoothIpOk() (*string, bool)`

GetBluetoothIpOk returns a tuple with the BluetoothIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBluetoothIp

`func (o *Terminal) SetBluetoothIp(v string)`

SetBluetoothIp sets BluetoothIp field to given value.

### HasBluetoothIp

`func (o *Terminal) HasBluetoothIp() bool`

HasBluetoothIp returns a boolean if a field has been set.

### GetBluetoothMac

`func (o *Terminal) GetBluetoothMac() string`

GetBluetoothMac returns the BluetoothMac field if non-nil, zero value otherwise.

### GetBluetoothMacOk

`func (o *Terminal) GetBluetoothMacOk() (*string, bool)`

GetBluetoothMacOk returns a tuple with the BluetoothMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBluetoothMac

`func (o *Terminal) SetBluetoothMac(v string)`

SetBluetoothMac sets BluetoothMac field to given value.

### HasBluetoothMac

`func (o *Terminal) HasBluetoothMac() bool`

HasBluetoothMac returns a boolean if a field has been set.

### GetCity

`func (o *Terminal) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Terminal) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Terminal) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Terminal) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanyAccount

`func (o *Terminal) GetCompanyAccount() string`

GetCompanyAccount returns the CompanyAccount field if non-nil, zero value otherwise.

### GetCompanyAccountOk

`func (o *Terminal) GetCompanyAccountOk() (*string, bool)`

GetCompanyAccountOk returns a tuple with the CompanyAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAccount

`func (o *Terminal) SetCompanyAccount(v string)`

SetCompanyAccount sets CompanyAccount field to given value.

### HasCompanyAccount

`func (o *Terminal) HasCompanyAccount() bool`

HasCompanyAccount returns a boolean if a field has been set.

### GetCountryCode

`func (o *Terminal) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Terminal) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Terminal) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *Terminal) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetDeviceModel

`func (o *Terminal) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *Terminal) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *Terminal) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *Terminal) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetEthernetIp

`func (o *Terminal) GetEthernetIp() string`

GetEthernetIp returns the EthernetIp field if non-nil, zero value otherwise.

### GetEthernetIpOk

`func (o *Terminal) GetEthernetIpOk() (*string, bool)`

GetEthernetIpOk returns a tuple with the EthernetIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetIp

`func (o *Terminal) SetEthernetIp(v string)`

SetEthernetIp sets EthernetIp field to given value.

### HasEthernetIp

`func (o *Terminal) HasEthernetIp() bool`

HasEthernetIp returns a boolean if a field has been set.

### GetEthernetMac

`func (o *Terminal) GetEthernetMac() string`

GetEthernetMac returns the EthernetMac field if non-nil, zero value otherwise.

### GetEthernetMacOk

`func (o *Terminal) GetEthernetMacOk() (*string, bool)`

GetEthernetMacOk returns a tuple with the EthernetMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetMac

`func (o *Terminal) SetEthernetMac(v string)`

SetEthernetMac sets EthernetMac field to given value.

### HasEthernetMac

`func (o *Terminal) HasEthernetMac() bool`

HasEthernetMac returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *Terminal) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *Terminal) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *Terminal) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *Terminal) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetIccid

`func (o *Terminal) GetIccid() string`

GetIccid returns the Iccid field if non-nil, zero value otherwise.

### GetIccidOk

`func (o *Terminal) GetIccidOk() (*string, bool)`

GetIccidOk returns a tuple with the Iccid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIccid

`func (o *Terminal) SetIccid(v string)`

SetIccid sets Iccid field to given value.

### HasIccid

`func (o *Terminal) HasIccid() bool`

HasIccid returns a boolean if a field has been set.

### GetId

`func (o *Terminal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Terminal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Terminal) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Terminal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastActivityDateTime

`func (o *Terminal) GetLastActivityDateTime() time.Time`

GetLastActivityDateTime returns the LastActivityDateTime field if non-nil, zero value otherwise.

### GetLastActivityDateTimeOk

`func (o *Terminal) GetLastActivityDateTimeOk() (*time.Time, bool)`

GetLastActivityDateTimeOk returns a tuple with the LastActivityDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityDateTime

`func (o *Terminal) SetLastActivityDateTime(v time.Time)`

SetLastActivityDateTime sets LastActivityDateTime field to given value.

### HasLastActivityDateTime

`func (o *Terminal) HasLastActivityDateTime() bool`

HasLastActivityDateTime returns a boolean if a field has been set.

### GetLastTransactionDateTime

`func (o *Terminal) GetLastTransactionDateTime() time.Time`

GetLastTransactionDateTime returns the LastTransactionDateTime field if non-nil, zero value otherwise.

### GetLastTransactionDateTimeOk

`func (o *Terminal) GetLastTransactionDateTimeOk() (*time.Time, bool)`

GetLastTransactionDateTimeOk returns a tuple with the LastTransactionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransactionDateTime

`func (o *Terminal) SetLastTransactionDateTime(v time.Time)`

SetLastTransactionDateTime sets LastTransactionDateTime field to given value.

### HasLastTransactionDateTime

`func (o *Terminal) HasLastTransactionDateTime() bool`

HasLastTransactionDateTime returns a boolean if a field has been set.

### GetLinkNegotiation

`func (o *Terminal) GetLinkNegotiation() string`

GetLinkNegotiation returns the LinkNegotiation field if non-nil, zero value otherwise.

### GetLinkNegotiationOk

`func (o *Terminal) GetLinkNegotiationOk() (*string, bool)`

GetLinkNegotiationOk returns a tuple with the LinkNegotiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNegotiation

`func (o *Terminal) SetLinkNegotiation(v string)`

SetLinkNegotiation sets LinkNegotiation field to given value.

### HasLinkNegotiation

`func (o *Terminal) HasLinkNegotiation() bool`

HasLinkNegotiation returns a boolean if a field has been set.

### GetSerialNumber

`func (o *Terminal) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Terminal) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Terminal) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *Terminal) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSimStatus

`func (o *Terminal) GetSimStatus() string`

GetSimStatus returns the SimStatus field if non-nil, zero value otherwise.

### GetSimStatusOk

`func (o *Terminal) GetSimStatusOk() (*string, bool)`

GetSimStatusOk returns a tuple with the SimStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimStatus

`func (o *Terminal) SetSimStatus(v string)`

SetSimStatus sets SimStatus field to given value.

### HasSimStatus

`func (o *Terminal) HasSimStatus() bool`

HasSimStatus returns a boolean if a field has been set.

### GetStatus

`func (o *Terminal) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Terminal) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Terminal) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Terminal) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStoreStatus

`func (o *Terminal) GetStoreStatus() string`

GetStoreStatus returns the StoreStatus field if non-nil, zero value otherwise.

### GetStoreStatusOk

`func (o *Terminal) GetStoreStatusOk() (*string, bool)`

GetStoreStatusOk returns a tuple with the StoreStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreStatus

`func (o *Terminal) SetStoreStatus(v string)`

SetStoreStatus sets StoreStatus field to given value.

### HasStoreStatus

`func (o *Terminal) HasStoreStatus() bool`

HasStoreStatus returns a boolean if a field has been set.

### GetWifiIp

`func (o *Terminal) GetWifiIp() string`

GetWifiIp returns the WifiIp field if non-nil, zero value otherwise.

### GetWifiIpOk

`func (o *Terminal) GetWifiIpOk() (*string, bool)`

GetWifiIpOk returns a tuple with the WifiIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiIp

`func (o *Terminal) SetWifiIp(v string)`

SetWifiIp sets WifiIp field to given value.

### HasWifiIp

`func (o *Terminal) HasWifiIp() bool`

HasWifiIp returns a boolean if a field has been set.

### GetWifiMac

`func (o *Terminal) GetWifiMac() string`

GetWifiMac returns the WifiMac field if non-nil, zero value otherwise.

### GetWifiMacOk

`func (o *Terminal) GetWifiMacOk() (*string, bool)`

GetWifiMacOk returns a tuple with the WifiMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMac

`func (o *Terminal) SetWifiMac(v string)`

SetWifiMac sets WifiMac field to given value.

### HasWifiMac

`func (o *Terminal) HasWifiMac() bool`

HasWifiMac returns a boolean if a field has been set.

### GetWifiSsid

`func (o *Terminal) GetWifiSsid() string`

GetWifiSsid returns the WifiSsid field if non-nil, zero value otherwise.

### GetWifiSsidOk

`func (o *Terminal) GetWifiSsidOk() (*string, bool)`

GetWifiSsidOk returns a tuple with the WifiSsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiSsid

`func (o *Terminal) SetWifiSsid(v string)`

SetWifiSsid sets WifiSsid field to given value.

### HasWifiSsid

`func (o *Terminal) HasWifiSsid() bool`

HasWifiSsid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


