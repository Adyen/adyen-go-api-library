# ApplePaySessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **string** | Base64 encoded data you need to [complete the Apple Pay merchant validation](https://docs.adyen.com/payment-methods/apple-pay/api-only?tab&#x3D;adyen-certificate-validation_1#complete-apple-pay-session-validation). | 

## Methods

### NewApplePaySessionResponse

`func NewApplePaySessionResponse(data string, ) *ApplePaySessionResponse`

NewApplePaySessionResponse instantiates a new ApplePaySessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplePaySessionResponseWithDefaults

`func NewApplePaySessionResponseWithDefaults() *ApplePaySessionResponse`

NewApplePaySessionResponseWithDefaults instantiates a new ApplePaySessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ApplePaySessionResponse) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApplePaySessionResponse) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApplePaySessionResponse) SetData(v string)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


