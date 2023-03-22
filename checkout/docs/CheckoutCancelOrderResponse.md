# CheckoutCancelOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PspReference** | **string** | A unique reference of the cancellation request. | 
**ResultCode** | **string** | The result of the cancellation request.  Possible values:  * **Received** – Indicates the cancellation has successfully been received by Adyen, and will be processed. | 

## Methods

### NewCheckoutCancelOrderResponse

`func NewCheckoutCancelOrderResponse(pspReference string, resultCode string, ) *CheckoutCancelOrderResponse`

NewCheckoutCancelOrderResponse instantiates a new CheckoutCancelOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutCancelOrderResponseWithDefaults

`func NewCheckoutCancelOrderResponseWithDefaults() *CheckoutCancelOrderResponse`

NewCheckoutCancelOrderResponseWithDefaults instantiates a new CheckoutCancelOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPspReference

`func (o *CheckoutCancelOrderResponse) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *CheckoutCancelOrderResponse) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *CheckoutCancelOrderResponse) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.


### GetResultCode

`func (o *CheckoutCancelOrderResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *CheckoutCancelOrderResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *CheckoutCancelOrderResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


