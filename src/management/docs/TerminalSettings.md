# TerminalSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardholderReceipt** | Pointer to [**CardholderReceipt**](CardholderReceipt.md) |  | [optional] 
**Connectivity** | Pointer to [**Connectivity**](Connectivity.md) |  | [optional] 
**Gratuities** | Pointer to [**[]Gratuity**](Gratuity.md) | Settings for tipping with or without predefined options to choose from. The maximum number of predefined options is four, or three plus the option to enter a custom tip. | [optional] 
**Hardware** | Pointer to [**Hardware**](Hardware.md) |  | [optional] 
**Nexo** | Pointer to [**Nexo**](Nexo.md) |  | [optional] 
**OfflineProcessing** | Pointer to [**OfflineProcessing**](OfflineProcessing.md) |  | [optional] 
**Opi** | Pointer to [**Opi**](Opi.md) |  | [optional] 
**ReceiptOptions** | Pointer to [**ReceiptOptions**](ReceiptOptions.md) |  | [optional] 
**ReceiptPrinting** | Pointer to [**ReceiptPrinting**](ReceiptPrinting.md) |  | [optional] 
**Signature** | Pointer to [**Signature**](Signature.md) |  | [optional] 
**Surcharge** | Pointer to [**Surcharge**](Surcharge.md) |  | [optional] 
**Timeouts** | Pointer to [**Timeouts**](Timeouts.md) |  | [optional] 
**WifiProfiles** | Pointer to [**WifiProfiles**](WifiProfiles.md) |  | [optional] 

## Methods

### NewTerminalSettings

`func NewTerminalSettings() *TerminalSettings`

NewTerminalSettings instantiates a new TerminalSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminalSettingsWithDefaults

`func NewTerminalSettingsWithDefaults() *TerminalSettings`

NewTerminalSettingsWithDefaults instantiates a new TerminalSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardholderReceipt

`func (o *TerminalSettings) GetCardholderReceipt() CardholderReceipt`

GetCardholderReceipt returns the CardholderReceipt field if non-nil, zero value otherwise.

### GetCardholderReceiptOk

`func (o *TerminalSettings) GetCardholderReceiptOk() (*CardholderReceipt, bool)`

GetCardholderReceiptOk returns a tuple with the CardholderReceipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderReceipt

`func (o *TerminalSettings) SetCardholderReceipt(v CardholderReceipt)`

SetCardholderReceipt sets CardholderReceipt field to given value.

### HasCardholderReceipt

`func (o *TerminalSettings) HasCardholderReceipt() bool`

HasCardholderReceipt returns a boolean if a field has been set.

### GetConnectivity

`func (o *TerminalSettings) GetConnectivity() Connectivity`

GetConnectivity returns the Connectivity field if non-nil, zero value otherwise.

### GetConnectivityOk

`func (o *TerminalSettings) GetConnectivityOk() (*Connectivity, bool)`

GetConnectivityOk returns a tuple with the Connectivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivity

`func (o *TerminalSettings) SetConnectivity(v Connectivity)`

SetConnectivity sets Connectivity field to given value.

### HasConnectivity

`func (o *TerminalSettings) HasConnectivity() bool`

HasConnectivity returns a boolean if a field has been set.

### GetGratuities

`func (o *TerminalSettings) GetGratuities() []Gratuity`

GetGratuities returns the Gratuities field if non-nil, zero value otherwise.

### GetGratuitiesOk

`func (o *TerminalSettings) GetGratuitiesOk() (*[]Gratuity, bool)`

GetGratuitiesOk returns a tuple with the Gratuities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGratuities

`func (o *TerminalSettings) SetGratuities(v []Gratuity)`

SetGratuities sets Gratuities field to given value.

### HasGratuities

`func (o *TerminalSettings) HasGratuities() bool`

HasGratuities returns a boolean if a field has been set.

### GetHardware

`func (o *TerminalSettings) GetHardware() Hardware`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *TerminalSettings) GetHardwareOk() (*Hardware, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *TerminalSettings) SetHardware(v Hardware)`

SetHardware sets Hardware field to given value.

### HasHardware

`func (o *TerminalSettings) HasHardware() bool`

HasHardware returns a boolean if a field has been set.

### GetNexo

`func (o *TerminalSettings) GetNexo() Nexo`

GetNexo returns the Nexo field if non-nil, zero value otherwise.

### GetNexoOk

`func (o *TerminalSettings) GetNexoOk() (*Nexo, bool)`

GetNexoOk returns a tuple with the Nexo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexo

`func (o *TerminalSettings) SetNexo(v Nexo)`

SetNexo sets Nexo field to given value.

### HasNexo

`func (o *TerminalSettings) HasNexo() bool`

HasNexo returns a boolean if a field has been set.

### GetOfflineProcessing

`func (o *TerminalSettings) GetOfflineProcessing() OfflineProcessing`

GetOfflineProcessing returns the OfflineProcessing field if non-nil, zero value otherwise.

### GetOfflineProcessingOk

`func (o *TerminalSettings) GetOfflineProcessingOk() (*OfflineProcessing, bool)`

GetOfflineProcessingOk returns a tuple with the OfflineProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineProcessing

`func (o *TerminalSettings) SetOfflineProcessing(v OfflineProcessing)`

SetOfflineProcessing sets OfflineProcessing field to given value.

### HasOfflineProcessing

`func (o *TerminalSettings) HasOfflineProcessing() bool`

HasOfflineProcessing returns a boolean if a field has been set.

### GetOpi

`func (o *TerminalSettings) GetOpi() Opi`

GetOpi returns the Opi field if non-nil, zero value otherwise.

### GetOpiOk

`func (o *TerminalSettings) GetOpiOk() (*Opi, bool)`

GetOpiOk returns a tuple with the Opi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpi

`func (o *TerminalSettings) SetOpi(v Opi)`

SetOpi sets Opi field to given value.

### HasOpi

`func (o *TerminalSettings) HasOpi() bool`

HasOpi returns a boolean if a field has been set.

### GetReceiptOptions

`func (o *TerminalSettings) GetReceiptOptions() ReceiptOptions`

GetReceiptOptions returns the ReceiptOptions field if non-nil, zero value otherwise.

### GetReceiptOptionsOk

`func (o *TerminalSettings) GetReceiptOptionsOk() (*ReceiptOptions, bool)`

GetReceiptOptionsOk returns a tuple with the ReceiptOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptOptions

`func (o *TerminalSettings) SetReceiptOptions(v ReceiptOptions)`

SetReceiptOptions sets ReceiptOptions field to given value.

### HasReceiptOptions

`func (o *TerminalSettings) HasReceiptOptions() bool`

HasReceiptOptions returns a boolean if a field has been set.

### GetReceiptPrinting

`func (o *TerminalSettings) GetReceiptPrinting() ReceiptPrinting`

GetReceiptPrinting returns the ReceiptPrinting field if non-nil, zero value otherwise.

### GetReceiptPrintingOk

`func (o *TerminalSettings) GetReceiptPrintingOk() (*ReceiptPrinting, bool)`

GetReceiptPrintingOk returns a tuple with the ReceiptPrinting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptPrinting

`func (o *TerminalSettings) SetReceiptPrinting(v ReceiptPrinting)`

SetReceiptPrinting sets ReceiptPrinting field to given value.

### HasReceiptPrinting

`func (o *TerminalSettings) HasReceiptPrinting() bool`

HasReceiptPrinting returns a boolean if a field has been set.

### GetSignature

`func (o *TerminalSettings) GetSignature() Signature`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *TerminalSettings) GetSignatureOk() (*Signature, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *TerminalSettings) SetSignature(v Signature)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *TerminalSettings) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetSurcharge

`func (o *TerminalSettings) GetSurcharge() Surcharge`

GetSurcharge returns the Surcharge field if non-nil, zero value otherwise.

### GetSurchargeOk

`func (o *TerminalSettings) GetSurchargeOk() (*Surcharge, bool)`

GetSurchargeOk returns a tuple with the Surcharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurcharge

`func (o *TerminalSettings) SetSurcharge(v Surcharge)`

SetSurcharge sets Surcharge field to given value.

### HasSurcharge

`func (o *TerminalSettings) HasSurcharge() bool`

HasSurcharge returns a boolean if a field has been set.

### GetTimeouts

`func (o *TerminalSettings) GetTimeouts() Timeouts`

GetTimeouts returns the Timeouts field if non-nil, zero value otherwise.

### GetTimeoutsOk

`func (o *TerminalSettings) GetTimeoutsOk() (*Timeouts, bool)`

GetTimeoutsOk returns a tuple with the Timeouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeouts

`func (o *TerminalSettings) SetTimeouts(v Timeouts)`

SetTimeouts sets Timeouts field to given value.

### HasTimeouts

`func (o *TerminalSettings) HasTimeouts() bool`

HasTimeouts returns a boolean if a field has been set.

### GetWifiProfiles

`func (o *TerminalSettings) GetWifiProfiles() WifiProfiles`

GetWifiProfiles returns the WifiProfiles field if non-nil, zero value otherwise.

### GetWifiProfilesOk

`func (o *TerminalSettings) GetWifiProfilesOk() (*WifiProfiles, bool)`

GetWifiProfilesOk returns a tuple with the WifiProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiProfiles

`func (o *TerminalSettings) SetWifiProfiles(v WifiProfiles)`

SetWifiProfiles sets WifiProfiles field to given value.

### HasWifiProfiles

`func (o *TerminalSettings) HasWifiProfiles() bool`

HasWifiProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


