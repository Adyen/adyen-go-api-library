# ThreeDS2ResultRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantAccount** | **string** | The merchant account identifier, with which you want to process the transaction. | 
**PspReference** | **string** | The pspReference returned in the /authorise call. | 

## Methods

### NewThreeDS2ResultRequest

`func NewThreeDS2ResultRequest(merchantAccount string, pspReference string, ) *ThreeDS2ResultRequest`

NewThreeDS2ResultRequest instantiates a new ThreeDS2ResultRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDS2ResultRequestWithDefaults

`func NewThreeDS2ResultRequestWithDefaults() *ThreeDS2ResultRequest`

NewThreeDS2ResultRequestWithDefaults instantiates a new ThreeDS2ResultRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantAccount

`func (o *ThreeDS2ResultRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *ThreeDS2ResultRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *ThreeDS2ResultRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetPspReference

`func (o *ThreeDS2ResultRequest) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *ThreeDS2ResultRequest) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *ThreeDS2ResultRequest) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


