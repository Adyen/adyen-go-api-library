# StoreDetailAndSubmitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **map[string]string** | This field contains additional data, which may be returned in a particular response. | [optional] 
**PspReference** | **string** | A new reference to uniquely identify this request. | 
**RefusalReason** | Pointer to **string** | In case of refusal, an informational message for the reason. | [optional] 
**ResultCode** | **string** | The response:  * In case of success is payout-submit-received. * In case of an error, an informational message is returned. | 

## Methods

### NewStoreDetailAndSubmitResponse

`func NewStoreDetailAndSubmitResponse(pspReference string, resultCode string, ) *StoreDetailAndSubmitResponse`

NewStoreDetailAndSubmitResponse instantiates a new StoreDetailAndSubmitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreDetailAndSubmitResponseWithDefaults

`func NewStoreDetailAndSubmitResponseWithDefaults() *StoreDetailAndSubmitResponse`

NewStoreDetailAndSubmitResponseWithDefaults instantiates a new StoreDetailAndSubmitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *StoreDetailAndSubmitResponse) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *StoreDetailAndSubmitResponse) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *StoreDetailAndSubmitResponse) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *StoreDetailAndSubmitResponse) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetPspReference

`func (o *StoreDetailAndSubmitResponse) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *StoreDetailAndSubmitResponse) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *StoreDetailAndSubmitResponse) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.


### GetRefusalReason

`func (o *StoreDetailAndSubmitResponse) GetRefusalReason() string`

GetRefusalReason returns the RefusalReason field if non-nil, zero value otherwise.

### GetRefusalReasonOk

`func (o *StoreDetailAndSubmitResponse) GetRefusalReasonOk() (*string, bool)`

GetRefusalReasonOk returns a tuple with the RefusalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusalReason

`func (o *StoreDetailAndSubmitResponse) SetRefusalReason(v string)`

SetRefusalReason sets RefusalReason field to given value.

### HasRefusalReason

`func (o *StoreDetailAndSubmitResponse) HasRefusalReason() bool`

HasRefusalReason returns a boolean if a field has been set.

### GetResultCode

`func (o *StoreDetailAndSubmitResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *StoreDetailAndSubmitResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *StoreDetailAndSubmitResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


