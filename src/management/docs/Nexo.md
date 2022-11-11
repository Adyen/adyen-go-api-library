# Nexo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventUrls** | Pointer to [**EventUrl**](EventUrl.md) |  | [optional] 
**NexoEventUrls** | Pointer to **[]string** | @deprecated One or more URLs to send event messages to when using Terminal API. | [optional] 

## Methods

### NewNexo

`func NewNexo() *Nexo`

NewNexo instantiates a new Nexo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNexoWithDefaults

`func NewNexoWithDefaults() *Nexo`

NewNexoWithDefaults instantiates a new Nexo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventUrls

`func (o *Nexo) GetEventUrls() EventUrl`

GetEventUrls returns the EventUrls field if non-nil, zero value otherwise.

### GetEventUrlsOk

`func (o *Nexo) GetEventUrlsOk() (*EventUrl, bool)`

GetEventUrlsOk returns a tuple with the EventUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventUrls

`func (o *Nexo) SetEventUrls(v EventUrl)`

SetEventUrls sets EventUrls field to given value.

### HasEventUrls

`func (o *Nexo) HasEventUrls() bool`

HasEventUrls returns a boolean if a field has been set.

### GetNexoEventUrls

`func (o *Nexo) GetNexoEventUrls() []string`

GetNexoEventUrls returns the NexoEventUrls field if non-nil, zero value otherwise.

### GetNexoEventUrlsOk

`func (o *Nexo) GetNexoEventUrlsOk() (*[]string, bool)`

GetNexoEventUrlsOk returns a tuple with the NexoEventUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexoEventUrls

`func (o *Nexo) SetNexoEventUrls(v []string)`

SetNexoEventUrls sets NexoEventUrls field to given value.

### HasNexoEventUrls

`func (o *Nexo) HasNexoEventUrls() bool`

HasNexoEventUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


