# ModifyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **map[string]string** | This field contains additional data, which may be returned in a particular response. | [optional] 
**PspReference** | **string** | Adyen&#39;s 16-character string reference associated with the transaction. This value is globally unique; quote it when communicating with us about this response. | 
**Response** | **string** | The response: * In case of success, it is either &#x60;payout-confirm-received&#x60; or &#x60;payout-decline-received&#x60;. * In case of an error, an informational message is returned. | 

## Methods

### NewModifyResponse

`func NewModifyResponse(pspReference string, response string, ) *ModifyResponse`

NewModifyResponse instantiates a new ModifyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyResponseWithDefaults

`func NewModifyResponseWithDefaults() *ModifyResponse`

NewModifyResponseWithDefaults instantiates a new ModifyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *ModifyResponse) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *ModifyResponse) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *ModifyResponse) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *ModifyResponse) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetPspReference

`func (o *ModifyResponse) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *ModifyResponse) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *ModifyResponse) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.


### GetResponse

`func (o *ModifyResponse) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ModifyResponse) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ModifyResponse) SetResponse(v string)`

SetResponse sets Response field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


