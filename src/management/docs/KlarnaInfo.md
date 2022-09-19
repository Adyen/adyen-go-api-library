# KlarnaInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoCapture** | Pointer to **bool** | Indicates the status of [Automatic capture](https://docs.adyen.com/online-payments/capture#automatic-capture). Default value: **false**. | [optional] 
**DisputeEmail** | **string** | The email address for disputes. | 
**Region** | **string** | The region of operation. For example, **NA**, **EU**, **CH**, **AU**. | 
**SupportEmail** | **string** | The email address of merchant support. | 

## Methods

### NewKlarnaInfo

`func NewKlarnaInfo(disputeEmail string, region string, supportEmail string, ) *KlarnaInfo`

NewKlarnaInfo instantiates a new KlarnaInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKlarnaInfoWithDefaults

`func NewKlarnaInfoWithDefaults() *KlarnaInfo`

NewKlarnaInfoWithDefaults instantiates a new KlarnaInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoCapture

`func (o *KlarnaInfo) GetAutoCapture() bool`

GetAutoCapture returns the AutoCapture field if non-nil, zero value otherwise.

### GetAutoCaptureOk

`func (o *KlarnaInfo) GetAutoCaptureOk() (*bool, bool)`

GetAutoCaptureOk returns a tuple with the AutoCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCapture

`func (o *KlarnaInfo) SetAutoCapture(v bool)`

SetAutoCapture sets AutoCapture field to given value.

### HasAutoCapture

`func (o *KlarnaInfo) HasAutoCapture() bool`

HasAutoCapture returns a boolean if a field has been set.

### GetDisputeEmail

`func (o *KlarnaInfo) GetDisputeEmail() string`

GetDisputeEmail returns the DisputeEmail field if non-nil, zero value otherwise.

### GetDisputeEmailOk

`func (o *KlarnaInfo) GetDisputeEmailOk() (*string, bool)`

GetDisputeEmailOk returns a tuple with the DisputeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeEmail

`func (o *KlarnaInfo) SetDisputeEmail(v string)`

SetDisputeEmail sets DisputeEmail field to given value.


### GetRegion

`func (o *KlarnaInfo) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *KlarnaInfo) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *KlarnaInfo) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetSupportEmail

`func (o *KlarnaInfo) GetSupportEmail() string`

GetSupportEmail returns the SupportEmail field if non-nil, zero value otherwise.

### GetSupportEmailOk

`func (o *KlarnaInfo) GetSupportEmailOk() (*string, bool)`

GetSupportEmailOk returns a tuple with the SupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmail

`func (o *KlarnaInfo) SetSupportEmail(v string)`

SetSupportEmail sets SupportEmail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


