# RelayedAuthorisationData2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]string** | Contains key-value pairs of your references and descriptions, for example, &#x60;customId&#x60;:&#x60;your-own-custom-field-12345&#x60;. | [optional] 
**Reference** | Pointer to **string** | Your reference for the relayed authorisation data. | [optional] 

## Methods

### NewRelayedAuthorisationData2

`func NewRelayedAuthorisationData2() *RelayedAuthorisationData2`

NewRelayedAuthorisationData2 instantiates a new RelayedAuthorisationData2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelayedAuthorisationData2WithDefaults

`func NewRelayedAuthorisationData2WithDefaults() *RelayedAuthorisationData2`

NewRelayedAuthorisationData2WithDefaults instantiates a new RelayedAuthorisationData2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *RelayedAuthorisationData2) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RelayedAuthorisationData2) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RelayedAuthorisationData2) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RelayedAuthorisationData2) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReference

`func (o *RelayedAuthorisationData2) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *RelayedAuthorisationData2) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *RelayedAuthorisationData2) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *RelayedAuthorisationData2) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


