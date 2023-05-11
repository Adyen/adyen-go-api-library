# ReleaseUpdateDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type of terminal action: Update Release. | [optional] [default to "ReleaseUpdate"]
**UpdateAtFirstMaintenanceCall** | Pointer to **bool** | Boolean flag that tells if the terminal should update at the first next maintenance call. If false, terminal will update on its configured reboot time. | [optional] 

## Methods

### NewReleaseUpdateDetails

`func NewReleaseUpdateDetails() *ReleaseUpdateDetails`

NewReleaseUpdateDetails instantiates a new ReleaseUpdateDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseUpdateDetailsWithDefaults

`func NewReleaseUpdateDetailsWithDefaults() *ReleaseUpdateDetails`

NewReleaseUpdateDetailsWithDefaults instantiates a new ReleaseUpdateDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ReleaseUpdateDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReleaseUpdateDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReleaseUpdateDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ReleaseUpdateDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateAtFirstMaintenanceCall

`func (o *ReleaseUpdateDetails) GetUpdateAtFirstMaintenanceCall() bool`

GetUpdateAtFirstMaintenanceCall returns the UpdateAtFirstMaintenanceCall field if non-nil, zero value otherwise.

### GetUpdateAtFirstMaintenanceCallOk

`func (o *ReleaseUpdateDetails) GetUpdateAtFirstMaintenanceCallOk() (*bool, bool)`

GetUpdateAtFirstMaintenanceCallOk returns a tuple with the UpdateAtFirstMaintenanceCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAtFirstMaintenanceCall

`func (o *ReleaseUpdateDetails) SetUpdateAtFirstMaintenanceCall(v bool)`

SetUpdateAtFirstMaintenanceCall sets UpdateAtFirstMaintenanceCall field to given value.

### HasUpdateAtFirstMaintenanceCall

`func (o *ReleaseUpdateDetails) HasUpdateAtFirstMaintenanceCall() bool`

HasUpdateAtFirstMaintenanceCall returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


