# PaymentVerificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | **string** | Encrypted and signed payment result data. You should receive this value from the Checkout SDK after the shopper completes the payment. | 

## Methods

### NewPaymentVerificationRequest

`func NewPaymentVerificationRequest(payload string, ) *PaymentVerificationRequest`

NewPaymentVerificationRequest instantiates a new PaymentVerificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentVerificationRequestWithDefaults

`func NewPaymentVerificationRequestWithDefaults() *PaymentVerificationRequest`

NewPaymentVerificationRequestWithDefaults instantiates a new PaymentVerificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayload

`func (o *PaymentVerificationRequest) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *PaymentVerificationRequest) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *PaymentVerificationRequest) SetPayload(v string)`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


