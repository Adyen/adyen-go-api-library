# ScheduleTerminalActionsRequestActionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | The unique identifier of the app to be uninstalled. | [optional] 
**Type** | Pointer to **string** | Type of terminal action: Uninstall an Android certificate. | [optional] [default to "UninstallAndroidCertificate"]
**CertificateId** | Pointer to **string** | The unique identifier of the certificate to be uninstalled. | [optional] 
**UpdateAtFirstMaintenanceCall** | Pointer to **bool** | Boolean flag that tells if the terminal should update at the first next maintenance call. If false, terminal will update on its configured reboot time. | [optional] 

## Methods

### NewScheduleTerminalActionsRequestActionDetails

`func NewScheduleTerminalActionsRequestActionDetails() *ScheduleTerminalActionsRequestActionDetails`

NewScheduleTerminalActionsRequestActionDetails instantiates a new ScheduleTerminalActionsRequestActionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleTerminalActionsRequestActionDetailsWithDefaults

`func NewScheduleTerminalActionsRequestActionDetailsWithDefaults() *ScheduleTerminalActionsRequestActionDetails`

NewScheduleTerminalActionsRequestActionDetailsWithDefaults instantiates a new ScheduleTerminalActionsRequestActionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *ScheduleTerminalActionsRequestActionDetails) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ScheduleTerminalActionsRequestActionDetails) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ScheduleTerminalActionsRequestActionDetails) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ScheduleTerminalActionsRequestActionDetails) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetType

`func (o *ScheduleTerminalActionsRequestActionDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScheduleTerminalActionsRequestActionDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScheduleTerminalActionsRequestActionDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ScheduleTerminalActionsRequestActionDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCertificateId

`func (o *ScheduleTerminalActionsRequestActionDetails) GetCertificateId() string`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *ScheduleTerminalActionsRequestActionDetails) GetCertificateIdOk() (*string, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *ScheduleTerminalActionsRequestActionDetails) SetCertificateId(v string)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *ScheduleTerminalActionsRequestActionDetails) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### GetUpdateAtFirstMaintenanceCall

`func (o *ScheduleTerminalActionsRequestActionDetails) GetUpdateAtFirstMaintenanceCall() bool`

GetUpdateAtFirstMaintenanceCall returns the UpdateAtFirstMaintenanceCall field if non-nil, zero value otherwise.

### GetUpdateAtFirstMaintenanceCallOk

`func (o *ScheduleTerminalActionsRequestActionDetails) GetUpdateAtFirstMaintenanceCallOk() (*bool, bool)`

GetUpdateAtFirstMaintenanceCallOk returns a tuple with the UpdateAtFirstMaintenanceCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAtFirstMaintenanceCall

`func (o *ScheduleTerminalActionsRequestActionDetails) SetUpdateAtFirstMaintenanceCall(v bool)`

SetUpdateAtFirstMaintenanceCall sets UpdateAtFirstMaintenanceCall field to given value.

### HasUpdateAtFirstMaintenanceCall

`func (o *ScheduleTerminalActionsRequestActionDetails) HasUpdateAtFirstMaintenanceCall() bool`

HasUpdateAtFirstMaintenanceCall returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


