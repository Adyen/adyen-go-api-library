# GooglePayInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | **string** | Google Pay [Merchant ID](https://support.google.com/paymentscenter/answer/7163092?hl&#x3D;en). Character length and limitations: 16 alphanumeric characters or 20 numeric characters. | 
**ReuseMerchantId** | Pointer to **bool** | Indicates whether the Google Pay Merchant ID is used for several merchant accounts. Default value: **false**. | [optional] 

## Methods

### NewGooglePayInfo

`func NewGooglePayInfo(merchantId string, ) *GooglePayInfo`

NewGooglePayInfo instantiates a new GooglePayInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGooglePayInfoWithDefaults

`func NewGooglePayInfoWithDefaults() *GooglePayInfo`

NewGooglePayInfoWithDefaults instantiates a new GooglePayInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *GooglePayInfo) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *GooglePayInfo) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *GooglePayInfo) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetReuseMerchantId

`func (o *GooglePayInfo) GetReuseMerchantId() bool`

GetReuseMerchantId returns the ReuseMerchantId field if non-nil, zero value otherwise.

### GetReuseMerchantIdOk

`func (o *GooglePayInfo) GetReuseMerchantIdOk() (*bool, bool)`

GetReuseMerchantIdOk returns a tuple with the ReuseMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReuseMerchantId

`func (o *GooglePayInfo) SetReuseMerchantId(v bool)`

SetReuseMerchantId sets ReuseMerchantId field to given value.

### HasReuseMerchantId

`func (o *GooglePayInfo) HasReuseMerchantId() bool`

HasReuseMerchantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


