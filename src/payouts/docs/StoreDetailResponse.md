# StoreDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **map[string]string** | This field contains additional data, which may be returned in a particular response. | [optional] 
**PspReference** | **string** | A new reference to uniquely identify this request. | 
**RecurringDetailReference** | **string** | The token which you can use later on for submitting the payout. | 
**ResultCode** | **string** | The result code of the transaction. &#x60;Success&#x60; indicates that the details were stored successfully. | 

## Methods

### NewStoreDetailResponse

`func NewStoreDetailResponse(pspReference string, recurringDetailReference string, resultCode string, ) *StoreDetailResponse`

NewStoreDetailResponse instantiates a new StoreDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreDetailResponseWithDefaults

`func NewStoreDetailResponseWithDefaults() *StoreDetailResponse`

NewStoreDetailResponseWithDefaults instantiates a new StoreDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *StoreDetailResponse) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *StoreDetailResponse) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *StoreDetailResponse) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *StoreDetailResponse) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetPspReference

`func (o *StoreDetailResponse) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *StoreDetailResponse) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *StoreDetailResponse) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.


### GetRecurringDetailReference

`func (o *StoreDetailResponse) GetRecurringDetailReference() string`

GetRecurringDetailReference returns the RecurringDetailReference field if non-nil, zero value otherwise.

### GetRecurringDetailReferenceOk

`func (o *StoreDetailResponse) GetRecurringDetailReferenceOk() (*string, bool)`

GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDetailReference

`func (o *StoreDetailResponse) SetRecurringDetailReference(v string)`

SetRecurringDetailReference sets RecurringDetailReference field to given value.


### GetResultCode

`func (o *StoreDetailResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *StoreDetailResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *StoreDetailResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


