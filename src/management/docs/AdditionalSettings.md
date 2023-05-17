# AdditionalSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeEventCodes** | Pointer to **[]string** | Object containing list of event codes for which the notifcation will be sent.  | [optional] 
**Properties** | Pointer to **map[string]bool** | Object containing boolean key-value pairs. The key can be any [standard webhook additional setting](https://docs.adyen.com/development-resources/webhooks/additional-settings), and the value indicates if the setting is enabled. For example, &#x60;captureDelayHours&#x60;: **true** means the standard notifications you get will contain the number of hours remaining until the payment will be captured. | [optional] 

## Methods

### NewAdditionalSettings

`func NewAdditionalSettings() *AdditionalSettings`

NewAdditionalSettings instantiates a new AdditionalSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalSettingsWithDefaults

`func NewAdditionalSettingsWithDefaults() *AdditionalSettings`

NewAdditionalSettingsWithDefaults instantiates a new AdditionalSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeEventCodes

`func (o *AdditionalSettings) GetIncludeEventCodes() []string`

GetIncludeEventCodes returns the IncludeEventCodes field if non-nil, zero value otherwise.

### GetIncludeEventCodesOk

`func (o *AdditionalSettings) GetIncludeEventCodesOk() (*[]string, bool)`

GetIncludeEventCodesOk returns a tuple with the IncludeEventCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeEventCodes

`func (o *AdditionalSettings) SetIncludeEventCodes(v []string)`

SetIncludeEventCodes sets IncludeEventCodes field to given value.

### HasIncludeEventCodes

`func (o *AdditionalSettings) HasIncludeEventCodes() bool`

HasIncludeEventCodes returns a boolean if a field has been set.

### GetProperties

`func (o *AdditionalSettings) GetProperties() map[string]bool`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *AdditionalSettings) GetPropertiesOk() (*map[string]bool, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *AdditionalSettings) SetProperties(v map[string]bool)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *AdditionalSettings) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


