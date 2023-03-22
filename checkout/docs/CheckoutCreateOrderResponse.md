# CheckoutCreateOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **map[string]string** | Contains additional information about the payment. Some data fields are included only if you select them first: Go to **Customer Area** &gt; **Developers** &gt; **Additional data**. | [optional] 
**Amount** | [**Amount**](Amount.md) |  | 
**ExpiresAt** | **string** | The date that the order will expire. | 
**FraudResult** | Pointer to [**FraudResult**](FraudResult.md) |  | [optional] 
**OrderData** | **string** | The encrypted data that will be used by merchant for adding payments to the order. | 
**PspReference** | Pointer to **string** | Adyen&#39;s 16-character reference associated with the transaction/request. This value is globally unique; quote it when communicating with us about this request. | [optional] 
**Reference** | Pointer to **string** | The reference provided by merchant for creating the order. | [optional] 
**RefusalReason** | Pointer to **string** | If the payment&#39;s authorisation is refused or an error occurs during authorisation, this field holds Adyen&#39;s mapped reason for the refusal or a description of the error. When a transaction fails, the authorisation response includes &#x60;resultCode&#x60; and &#x60;refusalReason&#x60; values.  For more information, see [Refusal reasons](https://docs.adyen.com/development-resources/refusal-reasons). | [optional] 
**RemainingAmount** | [**Amount**](Amount.md) |  | 
**ResultCode** | **string** | The result of the order creation request.  The value is always **Success**. | 

## Methods

### NewCheckoutCreateOrderResponse

`func NewCheckoutCreateOrderResponse(amount Amount, expiresAt string, orderData string, remainingAmount Amount, resultCode string, ) *CheckoutCreateOrderResponse`

NewCheckoutCreateOrderResponse instantiates a new CheckoutCreateOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutCreateOrderResponseWithDefaults

`func NewCheckoutCreateOrderResponseWithDefaults() *CheckoutCreateOrderResponse`

NewCheckoutCreateOrderResponseWithDefaults instantiates a new CheckoutCreateOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *CheckoutCreateOrderResponse) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *CheckoutCreateOrderResponse) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *CheckoutCreateOrderResponse) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *CheckoutCreateOrderResponse) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetAmount

`func (o *CheckoutCreateOrderResponse) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CheckoutCreateOrderResponse) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CheckoutCreateOrderResponse) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetExpiresAt

`func (o *CheckoutCreateOrderResponse) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CheckoutCreateOrderResponse) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CheckoutCreateOrderResponse) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### GetFraudResult

`func (o *CheckoutCreateOrderResponse) GetFraudResult() FraudResult`

GetFraudResult returns the FraudResult field if non-nil, zero value otherwise.

### GetFraudResultOk

`func (o *CheckoutCreateOrderResponse) GetFraudResultOk() (*FraudResult, bool)`

GetFraudResultOk returns a tuple with the FraudResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudResult

`func (o *CheckoutCreateOrderResponse) SetFraudResult(v FraudResult)`

SetFraudResult sets FraudResult field to given value.

### HasFraudResult

`func (o *CheckoutCreateOrderResponse) HasFraudResult() bool`

HasFraudResult returns a boolean if a field has been set.

### GetOrderData

`func (o *CheckoutCreateOrderResponse) GetOrderData() string`

GetOrderData returns the OrderData field if non-nil, zero value otherwise.

### GetOrderDataOk

`func (o *CheckoutCreateOrderResponse) GetOrderDataOk() (*string, bool)`

GetOrderDataOk returns a tuple with the OrderData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderData

`func (o *CheckoutCreateOrderResponse) SetOrderData(v string)`

SetOrderData sets OrderData field to given value.


### GetPspReference

`func (o *CheckoutCreateOrderResponse) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *CheckoutCreateOrderResponse) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *CheckoutCreateOrderResponse) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.

### HasPspReference

`func (o *CheckoutCreateOrderResponse) HasPspReference() bool`

HasPspReference returns a boolean if a field has been set.

### GetReference

`func (o *CheckoutCreateOrderResponse) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CheckoutCreateOrderResponse) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CheckoutCreateOrderResponse) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *CheckoutCreateOrderResponse) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetRefusalReason

`func (o *CheckoutCreateOrderResponse) GetRefusalReason() string`

GetRefusalReason returns the RefusalReason field if non-nil, zero value otherwise.

### GetRefusalReasonOk

`func (o *CheckoutCreateOrderResponse) GetRefusalReasonOk() (*string, bool)`

GetRefusalReasonOk returns a tuple with the RefusalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusalReason

`func (o *CheckoutCreateOrderResponse) SetRefusalReason(v string)`

SetRefusalReason sets RefusalReason field to given value.

### HasRefusalReason

`func (o *CheckoutCreateOrderResponse) HasRefusalReason() bool`

HasRefusalReason returns a boolean if a field has been set.

### GetRemainingAmount

`func (o *CheckoutCreateOrderResponse) GetRemainingAmount() Amount`

GetRemainingAmount returns the RemainingAmount field if non-nil, zero value otherwise.

### GetRemainingAmountOk

`func (o *CheckoutCreateOrderResponse) GetRemainingAmountOk() (*Amount, bool)`

GetRemainingAmountOk returns a tuple with the RemainingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingAmount

`func (o *CheckoutCreateOrderResponse) SetRemainingAmount(v Amount)`

SetRemainingAmount sets RemainingAmount field to given value.


### GetResultCode

`func (o *CheckoutCreateOrderResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *CheckoutCreateOrderResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *CheckoutCreateOrderResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


