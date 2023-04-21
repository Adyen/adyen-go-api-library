# SubmitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **map[string]string** | This field contains additional data, which may be returned in a particular response. | [optional] 
**PspReference** | **string** | A new reference to uniquely identify this request. | 
**RefusalReason** | Pointer to **string** | In case of refusal, an informational message for the reason. | [optional] 
**ResultCode** | **string** | The response: * In case of success, it is &#x60;payout-submit-received&#x60;. * In case of an error, an informational message is returned. | 

## Methods

### NewSubmitResponse

`func NewSubmitResponse(pspReference string, resultCode string, ) *SubmitResponse`

NewSubmitResponse instantiates a new SubmitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitResponseWithDefaults

`func NewSubmitResponseWithDefaults() *SubmitResponse`

NewSubmitResponseWithDefaults instantiates a new SubmitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *SubmitResponse) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *SubmitResponse) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *SubmitResponse) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *SubmitResponse) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetPspReference

`func (o *SubmitResponse) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *SubmitResponse) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *SubmitResponse) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.


### GetRefusalReason

`func (o *SubmitResponse) GetRefusalReason() string`

GetRefusalReason returns the RefusalReason field if non-nil, zero value otherwise.

### GetRefusalReasonOk

`func (o *SubmitResponse) GetRefusalReasonOk() (*string, bool)`

GetRefusalReasonOk returns a tuple with the RefusalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusalReason

`func (o *SubmitResponse) SetRefusalReason(v string)`

SetRefusalReason sets RefusalReason field to given value.

### HasRefusalReason

`func (o *SubmitResponse) HasRefusalReason() bool`

HasRefusalReason returns a boolean if a field has been set.

### GetResultCode

`func (o *SubmitResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *SubmitResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *SubmitResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


