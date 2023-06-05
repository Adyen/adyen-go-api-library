# RelayedAuthorisationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]string** | The &#x60;metadata&#x60; object from the relayed authorisation response from your server. | [optional] 
**Reference** | Pointer to **string** | The &#x60;reference&#x60; from the relayed authorisation response from your server. | [optional] 
**Status** | Pointer to **string** | The value can be **Authorised** or **Refused**, based on the &#x60;authorisationDecision.status&#x60; in the relayed authorisation response from your server. | [optional] 

## Methods

### NewRelayedAuthorisationData

`func NewRelayedAuthorisationData() *RelayedAuthorisationData`

NewRelayedAuthorisationData instantiates a new RelayedAuthorisationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelayedAuthorisationDataWithDefaults

`func NewRelayedAuthorisationDataWithDefaults() *RelayedAuthorisationData`

NewRelayedAuthorisationDataWithDefaults instantiates a new RelayedAuthorisationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *RelayedAuthorisationData) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RelayedAuthorisationData) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RelayedAuthorisationData) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RelayedAuthorisationData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReference

`func (o *RelayedAuthorisationData) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *RelayedAuthorisationData) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *RelayedAuthorisationData) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *RelayedAuthorisationData) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetStatus

`func (o *RelayedAuthorisationData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RelayedAuthorisationData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RelayedAuthorisationData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RelayedAuthorisationData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


