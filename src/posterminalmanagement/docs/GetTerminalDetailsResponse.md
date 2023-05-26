# GetTerminalDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BluetoothIp** | Pointer to **string** | The Bluetooth IP address of the terminal. | [optional] 
**BluetoothMac** | Pointer to **string** | The Bluetooth MAC address of the terminal. | [optional] 
**CompanyAccount** | **string** | The company account that the terminal is associated with. If this is the only account level shown in the response, the terminal is assigned to the inventory of the company account. | 
**Country** | Pointer to **string** | The country where the terminal is used. | [optional] 
**DeviceModel** | Pointer to **string** | The model name of the terminal. | [optional] 
**DhcpEnabled** | Pointer to **bool** | Indicates whether assigning IP addresses through a DHCP server is enabled on the terminal. | [optional] 
**DisplayLabel** | Pointer to **string** | The label shown on the status bar of the display. This label (if any) is specified in your Customer Area. | [optional] 
**EthernetIp** | Pointer to **string** | The terminal&#39;s IP address in your Ethernet network. | [optional] 
**EthernetMac** | Pointer to **string** | The terminal&#39;s MAC address in your Ethernet network. | [optional] 
**FirmwareVersion** | Pointer to **string** | The software release currently in use on the terminal. | [optional] 
**Iccid** | Pointer to **string** | The integrated circuit card identifier (ICCID) of the SIM card in the terminal. | [optional] 
**LastActivityDateTime** | Pointer to **time.Time** | Date and time of the last activity on the terminal. Not included when the last activity was more than 14 days ago. | [optional] 
**LastTransactionDateTime** | Pointer to **time.Time** | Date and time of the last transaction on the terminal. Not included when the last transaction was more than 14 days ago. | [optional] 
**LinkNegotiation** | Pointer to **string** | The Ethernet link negotiation that the terminal uses:   - &#x60;auto&#x60;: Auto-negotiation  - &#x60;100full&#x60;: 100 Mbps full duplex | [optional] 
**MerchantAccount** | Pointer to **string** | The merchant account that the terminal is associated with. If the response doesn&#39;t contain a &#x60;store&#x60; the terminal is assigned to this merchant account. | [optional] 
**MerchantInventory** | Pointer to **bool** | Boolean that indicates if the terminal is assigned to the merchant inventory. This is returned when the terminal is assigned to a merchant account.  - If **true**, this indicates that the terminal is in the merchant inventory. This also means that the terminal cannot be boarded.  - If **false**, this indicates that the terminal is assigned to the merchant account as an in-store terminal. This means that the terminal is ready to be boarded, or is already boarded. | [optional] 
**PermanentTerminalId** | Pointer to **string** | The permanent terminal ID. | [optional] 
**SerialNumber** | Pointer to **string** | The serial number of the terminal. | [optional] 
**SimStatus** | Pointer to **string** | On a terminal that supports 3G or 4G connectivity, indicates the status of the SIM card in the terminal: ACTIVE or INVENTORY. | [optional] 
**Store** | Pointer to **string** | The store code of the store that the terminal is assigned to. | [optional] 
**StoreDetails** | Pointer to [**Store**](Store.md) |  | [optional] 
**Terminal** | **string** | The unique terminal ID. | 
**TerminalStatus** | Pointer to **string** | The status of the terminal:   - &#x60;OnlineToday&#x60;, &#x60;OnlineLast1Day&#x60;, &#x60;OnlineLast2Days&#x60; etcetera to &#x60;OnlineLast7Days&#x60;: Indicates when in the past week the terminal was last online.   - &#x60;SwitchedOff&#x60;: Indicates it was more than a week ago that the terminal was last online.   - &#x60;ReAssignToInventoryPending&#x60;, &#x60;ReAssignToStorePending&#x60;, &#x60;ReAssignToMerchantInventoryPending&#x60;: Indicates the terminal is scheduled to be reassigned. | [optional] 
**WifiIp** | Pointer to **string** | The terminal&#39;s IP address in your Wi-Fi network. | [optional] 
**WifiMac** | Pointer to **string** | The terminal&#39;s MAC address in your Wi-Fi network. | [optional] 

## Methods

### NewGetTerminalDetailsResponse

`func NewGetTerminalDetailsResponse(companyAccount string, terminal string, ) *GetTerminalDetailsResponse`

NewGetTerminalDetailsResponse instantiates a new GetTerminalDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTerminalDetailsResponseWithDefaults

`func NewGetTerminalDetailsResponseWithDefaults() *GetTerminalDetailsResponse`

NewGetTerminalDetailsResponseWithDefaults instantiates a new GetTerminalDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBluetoothIp

`func (o *GetTerminalDetailsResponse) GetBluetoothIp() string`

GetBluetoothIp returns the BluetoothIp field if non-nil, zero value otherwise.

### GetBluetoothIpOk

`func (o *GetTerminalDetailsResponse) GetBluetoothIpOk() (*string, bool)`

GetBluetoothIpOk returns a tuple with the BluetoothIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBluetoothIp

`func (o *GetTerminalDetailsResponse) SetBluetoothIp(v string)`

SetBluetoothIp sets BluetoothIp field to given value.

### HasBluetoothIp

`func (o *GetTerminalDetailsResponse) HasBluetoothIp() bool`

HasBluetoothIp returns a boolean if a field has been set.

### GetBluetoothMac

`func (o *GetTerminalDetailsResponse) GetBluetoothMac() string`

GetBluetoothMac returns the BluetoothMac field if non-nil, zero value otherwise.

### GetBluetoothMacOk

`func (o *GetTerminalDetailsResponse) GetBluetoothMacOk() (*string, bool)`

GetBluetoothMacOk returns a tuple with the BluetoothMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBluetoothMac

`func (o *GetTerminalDetailsResponse) SetBluetoothMac(v string)`

SetBluetoothMac sets BluetoothMac field to given value.

### HasBluetoothMac

`func (o *GetTerminalDetailsResponse) HasBluetoothMac() bool`

HasBluetoothMac returns a boolean if a field has been set.

### GetCompanyAccount

`func (o *GetTerminalDetailsResponse) GetCompanyAccount() string`

GetCompanyAccount returns the CompanyAccount field if non-nil, zero value otherwise.

### GetCompanyAccountOk

`func (o *GetTerminalDetailsResponse) GetCompanyAccountOk() (*string, bool)`

GetCompanyAccountOk returns a tuple with the CompanyAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAccount

`func (o *GetTerminalDetailsResponse) SetCompanyAccount(v string)`

SetCompanyAccount sets CompanyAccount field to given value.


### GetCountry

`func (o *GetTerminalDetailsResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GetTerminalDetailsResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GetTerminalDetailsResponse) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GetTerminalDetailsResponse) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDeviceModel

`func (o *GetTerminalDetailsResponse) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *GetTerminalDetailsResponse) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *GetTerminalDetailsResponse) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *GetTerminalDetailsResponse) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetDhcpEnabled

`func (o *GetTerminalDetailsResponse) GetDhcpEnabled() bool`

GetDhcpEnabled returns the DhcpEnabled field if non-nil, zero value otherwise.

### GetDhcpEnabledOk

`func (o *GetTerminalDetailsResponse) GetDhcpEnabledOk() (*bool, bool)`

GetDhcpEnabledOk returns a tuple with the DhcpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpEnabled

`func (o *GetTerminalDetailsResponse) SetDhcpEnabled(v bool)`

SetDhcpEnabled sets DhcpEnabled field to given value.

### HasDhcpEnabled

`func (o *GetTerminalDetailsResponse) HasDhcpEnabled() bool`

HasDhcpEnabled returns a boolean if a field has been set.

### GetDisplayLabel

`func (o *GetTerminalDetailsResponse) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *GetTerminalDetailsResponse) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *GetTerminalDetailsResponse) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *GetTerminalDetailsResponse) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### GetEthernetIp

`func (o *GetTerminalDetailsResponse) GetEthernetIp() string`

GetEthernetIp returns the EthernetIp field if non-nil, zero value otherwise.

### GetEthernetIpOk

`func (o *GetTerminalDetailsResponse) GetEthernetIpOk() (*string, bool)`

GetEthernetIpOk returns a tuple with the EthernetIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetIp

`func (o *GetTerminalDetailsResponse) SetEthernetIp(v string)`

SetEthernetIp sets EthernetIp field to given value.

### HasEthernetIp

`func (o *GetTerminalDetailsResponse) HasEthernetIp() bool`

HasEthernetIp returns a boolean if a field has been set.

### GetEthernetMac

`func (o *GetTerminalDetailsResponse) GetEthernetMac() string`

GetEthernetMac returns the EthernetMac field if non-nil, zero value otherwise.

### GetEthernetMacOk

`func (o *GetTerminalDetailsResponse) GetEthernetMacOk() (*string, bool)`

GetEthernetMacOk returns a tuple with the EthernetMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetMac

`func (o *GetTerminalDetailsResponse) SetEthernetMac(v string)`

SetEthernetMac sets EthernetMac field to given value.

### HasEthernetMac

`func (o *GetTerminalDetailsResponse) HasEthernetMac() bool`

HasEthernetMac returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *GetTerminalDetailsResponse) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *GetTerminalDetailsResponse) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *GetTerminalDetailsResponse) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *GetTerminalDetailsResponse) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetIccid

`func (o *GetTerminalDetailsResponse) GetIccid() string`

GetIccid returns the Iccid field if non-nil, zero value otherwise.

### GetIccidOk

`func (o *GetTerminalDetailsResponse) GetIccidOk() (*string, bool)`

GetIccidOk returns a tuple with the Iccid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIccid

`func (o *GetTerminalDetailsResponse) SetIccid(v string)`

SetIccid sets Iccid field to given value.

### HasIccid

`func (o *GetTerminalDetailsResponse) HasIccid() bool`

HasIccid returns a boolean if a field has been set.

### GetLastActivityDateTime

`func (o *GetTerminalDetailsResponse) GetLastActivityDateTime() time.Time`

GetLastActivityDateTime returns the LastActivityDateTime field if non-nil, zero value otherwise.

### GetLastActivityDateTimeOk

`func (o *GetTerminalDetailsResponse) GetLastActivityDateTimeOk() (*time.Time, bool)`

GetLastActivityDateTimeOk returns a tuple with the LastActivityDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityDateTime

`func (o *GetTerminalDetailsResponse) SetLastActivityDateTime(v time.Time)`

SetLastActivityDateTime sets LastActivityDateTime field to given value.

### HasLastActivityDateTime

`func (o *GetTerminalDetailsResponse) HasLastActivityDateTime() bool`

HasLastActivityDateTime returns a boolean if a field has been set.

### GetLastTransactionDateTime

`func (o *GetTerminalDetailsResponse) GetLastTransactionDateTime() time.Time`

GetLastTransactionDateTime returns the LastTransactionDateTime field if non-nil, zero value otherwise.

### GetLastTransactionDateTimeOk

`func (o *GetTerminalDetailsResponse) GetLastTransactionDateTimeOk() (*time.Time, bool)`

GetLastTransactionDateTimeOk returns a tuple with the LastTransactionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransactionDateTime

`func (o *GetTerminalDetailsResponse) SetLastTransactionDateTime(v time.Time)`

SetLastTransactionDateTime sets LastTransactionDateTime field to given value.

### HasLastTransactionDateTime

`func (o *GetTerminalDetailsResponse) HasLastTransactionDateTime() bool`

HasLastTransactionDateTime returns a boolean if a field has been set.

### GetLinkNegotiation

`func (o *GetTerminalDetailsResponse) GetLinkNegotiation() string`

GetLinkNegotiation returns the LinkNegotiation field if non-nil, zero value otherwise.

### GetLinkNegotiationOk

`func (o *GetTerminalDetailsResponse) GetLinkNegotiationOk() (*string, bool)`

GetLinkNegotiationOk returns a tuple with the LinkNegotiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNegotiation

`func (o *GetTerminalDetailsResponse) SetLinkNegotiation(v string)`

SetLinkNegotiation sets LinkNegotiation field to given value.

### HasLinkNegotiation

`func (o *GetTerminalDetailsResponse) HasLinkNegotiation() bool`

HasLinkNegotiation returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *GetTerminalDetailsResponse) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *GetTerminalDetailsResponse) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *GetTerminalDetailsResponse) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.

### HasMerchantAccount

`func (o *GetTerminalDetailsResponse) HasMerchantAccount() bool`

HasMerchantAccount returns a boolean if a field has been set.

### GetMerchantInventory

`func (o *GetTerminalDetailsResponse) GetMerchantInventory() bool`

GetMerchantInventory returns the MerchantInventory field if non-nil, zero value otherwise.

### GetMerchantInventoryOk

`func (o *GetTerminalDetailsResponse) GetMerchantInventoryOk() (*bool, bool)`

GetMerchantInventoryOk returns a tuple with the MerchantInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantInventory

`func (o *GetTerminalDetailsResponse) SetMerchantInventory(v bool)`

SetMerchantInventory sets MerchantInventory field to given value.

### HasMerchantInventory

`func (o *GetTerminalDetailsResponse) HasMerchantInventory() bool`

HasMerchantInventory returns a boolean if a field has been set.

### GetPermanentTerminalId

`func (o *GetTerminalDetailsResponse) GetPermanentTerminalId() string`

GetPermanentTerminalId returns the PermanentTerminalId field if non-nil, zero value otherwise.

### GetPermanentTerminalIdOk

`func (o *GetTerminalDetailsResponse) GetPermanentTerminalIdOk() (*string, bool)`

GetPermanentTerminalIdOk returns a tuple with the PermanentTerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanentTerminalId

`func (o *GetTerminalDetailsResponse) SetPermanentTerminalId(v string)`

SetPermanentTerminalId sets PermanentTerminalId field to given value.

### HasPermanentTerminalId

`func (o *GetTerminalDetailsResponse) HasPermanentTerminalId() bool`

HasPermanentTerminalId returns a boolean if a field has been set.

### GetSerialNumber

`func (o *GetTerminalDetailsResponse) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *GetTerminalDetailsResponse) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *GetTerminalDetailsResponse) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *GetTerminalDetailsResponse) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSimStatus

`func (o *GetTerminalDetailsResponse) GetSimStatus() string`

GetSimStatus returns the SimStatus field if non-nil, zero value otherwise.

### GetSimStatusOk

`func (o *GetTerminalDetailsResponse) GetSimStatusOk() (*string, bool)`

GetSimStatusOk returns a tuple with the SimStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimStatus

`func (o *GetTerminalDetailsResponse) SetSimStatus(v string)`

SetSimStatus sets SimStatus field to given value.

### HasSimStatus

`func (o *GetTerminalDetailsResponse) HasSimStatus() bool`

HasSimStatus returns a boolean if a field has been set.

### GetStore

`func (o *GetTerminalDetailsResponse) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *GetTerminalDetailsResponse) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *GetTerminalDetailsResponse) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *GetTerminalDetailsResponse) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetStoreDetails

`func (o *GetTerminalDetailsResponse) GetStoreDetails() Store`

GetStoreDetails returns the StoreDetails field if non-nil, zero value otherwise.

### GetStoreDetailsOk

`func (o *GetTerminalDetailsResponse) GetStoreDetailsOk() (*Store, bool)`

GetStoreDetailsOk returns a tuple with the StoreDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreDetails

`func (o *GetTerminalDetailsResponse) SetStoreDetails(v Store)`

SetStoreDetails sets StoreDetails field to given value.

### HasStoreDetails

`func (o *GetTerminalDetailsResponse) HasStoreDetails() bool`

HasStoreDetails returns a boolean if a field has been set.

### GetTerminal

`func (o *GetTerminalDetailsResponse) GetTerminal() string`

GetTerminal returns the Terminal field if non-nil, zero value otherwise.

### GetTerminalOk

`func (o *GetTerminalDetailsResponse) GetTerminalOk() (*string, bool)`

GetTerminalOk returns a tuple with the Terminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminal

`func (o *GetTerminalDetailsResponse) SetTerminal(v string)`

SetTerminal sets Terminal field to given value.


### GetTerminalStatus

`func (o *GetTerminalDetailsResponse) GetTerminalStatus() string`

GetTerminalStatus returns the TerminalStatus field if non-nil, zero value otherwise.

### GetTerminalStatusOk

`func (o *GetTerminalDetailsResponse) GetTerminalStatusOk() (*string, bool)`

GetTerminalStatusOk returns a tuple with the TerminalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalStatus

`func (o *GetTerminalDetailsResponse) SetTerminalStatus(v string)`

SetTerminalStatus sets TerminalStatus field to given value.

### HasTerminalStatus

`func (o *GetTerminalDetailsResponse) HasTerminalStatus() bool`

HasTerminalStatus returns a boolean if a field has been set.

### GetWifiIp

`func (o *GetTerminalDetailsResponse) GetWifiIp() string`

GetWifiIp returns the WifiIp field if non-nil, zero value otherwise.

### GetWifiIpOk

`func (o *GetTerminalDetailsResponse) GetWifiIpOk() (*string, bool)`

GetWifiIpOk returns a tuple with the WifiIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiIp

`func (o *GetTerminalDetailsResponse) SetWifiIp(v string)`

SetWifiIp sets WifiIp field to given value.

### HasWifiIp

`func (o *GetTerminalDetailsResponse) HasWifiIp() bool`

HasWifiIp returns a boolean if a field has been set.

### GetWifiMac

`func (o *GetTerminalDetailsResponse) GetWifiMac() string`

GetWifiMac returns the WifiMac field if non-nil, zero value otherwise.

### GetWifiMacOk

`func (o *GetTerminalDetailsResponse) GetWifiMacOk() (*string, bool)`

GetWifiMacOk returns a tuple with the WifiMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMac

`func (o *GetTerminalDetailsResponse) SetWifiMac(v string)`

SetWifiMac sets WifiMac field to given value.

### HasWifiMac

`func (o *GetTerminalDetailsResponse) HasWifiMac() bool`

HasWifiMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


