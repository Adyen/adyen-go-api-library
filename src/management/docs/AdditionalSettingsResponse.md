# AdditionalSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeEventCodes** | Pointer to **[]string** | Object containing list of event codes for which the notifcation will not be sent.  | [optional] 
**IncludeEventCodes** | Pointer to **[]string** | Object containing list of event codes for which the notifcation will be sent.  | [optional] 
**Properties** | Pointer to **map[string]bool** | Object containing boolean key-value pairs. The key can be any [standard webhook additional setting](https://docs.adyen.com/development-resources/webhooks/additional-settings), and the value indicates if the setting is enabled. For example, &#x60;captureDelayHours&#x60;: **true** means the standard notifications you get will contain the number of hours remaining until the payment will be captured. | [optional] 

## Methods

### NewAdditionalSettingsResponse

`func NewAdditionalSettingsResponse() *AdditionalSettingsResponse`

NewAdditionalSettingsResponse instantiates a new AdditionalSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalSettingsResponseWithDefaults

`func NewAdditionalSettingsResponseWithDefaults() *AdditionalSettingsResponse`

NewAdditionalSettingsResponseWithDefaults instantiates a new AdditionalSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeEventCodes

`func (o *AdditionalSettingsResponse) GetExcludeEventCodes() []string`

GetExcludeEventCodes returns the ExcludeEventCodes field if non-nil, zero value otherwise.

### GetExcludeEventCodesOk

`func (o *AdditionalSettingsResponse) GetExcludeEventCodesOk() (*[]string, bool)`

GetExcludeEventCodesOk returns a tuple with the ExcludeEventCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeEventCodes

`func (o *AdditionalSettingsResponse) SetExcludeEventCodes(v []string)`

SetExcludeEventCodes sets ExcludeEventCodes field to given value.

### HasExcludeEventCodes

`func (o *AdditionalSettingsResponse) HasExcludeEventCodes() bool`

HasExcludeEventCodes returns a boolean if a field has been set.

### GetIncludeEventCodes

`func (o *AdditionalSettingsResponse) GetIncludeEventCodes() []string`

GetIncludeEventCodes returns the IncludeEventCodes field if non-nil, zero value otherwise.

### GetIncludeEventCodesOk

`func (o *AdditionalSettingsResponse) GetIncludeEventCodesOk() (*[]string, bool)`

GetIncludeEventCodesOk returns a tuple with the IncludeEventCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeEventCodes

`func (o *AdditionalSettingsResponse) SetIncludeEventCodes(v []string)`

SetIncludeEventCodes sets IncludeEventCodes field to given value.

### HasIncludeEventCodes

`func (o *AdditionalSettingsResponse) HasIncludeEventCodes() bool`

HasIncludeEventCodes returns a boolean if a field has been set.

### GetProperties

`func (o *AdditionalSettingsResponse) GetProperties() map[string]bool`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *AdditionalSettingsResponse) GetPropertiesOk() (*map[string]bool, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *AdditionalSettingsResponse) SetProperties(v map[string]bool)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *AdditionalSettingsResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


