# CheckoutBalanceCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **map[string]string** | Contains additional information about the payment. Some data fields are included only if you select them first: Go to **Customer Area** &gt; **Developers** &gt; **Additional data**. | [optional] 
**Balance** | [**Amount**](Amount.md) |  | 
**FraudResult** | Pointer to [**FraudResult**](FraudResult.md) |  | [optional] 
**PspReference** | Pointer to **string** | Adyen&#39;s 16-character reference associated with the transaction/request. This value is globally unique; quote it when communicating with us about this request. | [optional] 
**RefusalReason** | Pointer to **string** | If the payment&#39;s authorisation is refused or an error occurs during authorisation, this field holds Adyen&#39;s mapped reason for the refusal or a description of the error. When a transaction fails, the authorisation response includes &#x60;resultCode&#x60; and &#x60;refusalReason&#x60; values.  For more information, see [Refusal reasons](https://docs.adyen.com/development-resources/refusal-reasons). | [optional] 
**ResultCode** | **string** | The result of the cancellation request.  Possible values:  * **Success** – Indicates that the balance check was successful. * **NotEnoughBalance** – Commonly indicates that the card did not have enough balance to pay the amount in the request, or that the currency of the balance on the card did not match the currency of the requested amount. * **Failed** – Indicates that the balance check failed. | 
**TransactionLimit** | Pointer to [**Amount**](Amount.md) |  | [optional] 

## Methods

### NewCheckoutBalanceCheckResponse

`func NewCheckoutBalanceCheckResponse(balance Amount, resultCode string, ) *CheckoutBalanceCheckResponse`

NewCheckoutBalanceCheckResponse instantiates a new CheckoutBalanceCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutBalanceCheckResponseWithDefaults

`func NewCheckoutBalanceCheckResponseWithDefaults() *CheckoutBalanceCheckResponse`

NewCheckoutBalanceCheckResponseWithDefaults instantiates a new CheckoutBalanceCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *CheckoutBalanceCheckResponse) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *CheckoutBalanceCheckResponse) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *CheckoutBalanceCheckResponse) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *CheckoutBalanceCheckResponse) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetBalance

`func (o *CheckoutBalanceCheckResponse) GetBalance() Amount`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *CheckoutBalanceCheckResponse) GetBalanceOk() (*Amount, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *CheckoutBalanceCheckResponse) SetBalance(v Amount)`

SetBalance sets Balance field to given value.


### GetFraudResult

`func (o *CheckoutBalanceCheckResponse) GetFraudResult() FraudResult`

GetFraudResult returns the FraudResult field if non-nil, zero value otherwise.

### GetFraudResultOk

`func (o *CheckoutBalanceCheckResponse) GetFraudResultOk() (*FraudResult, bool)`

GetFraudResultOk returns a tuple with the FraudResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudResult

`func (o *CheckoutBalanceCheckResponse) SetFraudResult(v FraudResult)`

SetFraudResult sets FraudResult field to given value.

### HasFraudResult

`func (o *CheckoutBalanceCheckResponse) HasFraudResult() bool`

HasFraudResult returns a boolean if a field has been set.

### GetPspReference

`func (o *CheckoutBalanceCheckResponse) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *CheckoutBalanceCheckResponse) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *CheckoutBalanceCheckResponse) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.

### HasPspReference

`func (o *CheckoutBalanceCheckResponse) HasPspReference() bool`

HasPspReference returns a boolean if a field has been set.

### GetRefusalReason

`func (o *CheckoutBalanceCheckResponse) GetRefusalReason() string`

GetRefusalReason returns the RefusalReason field if non-nil, zero value otherwise.

### GetRefusalReasonOk

`func (o *CheckoutBalanceCheckResponse) GetRefusalReasonOk() (*string, bool)`

GetRefusalReasonOk returns a tuple with the RefusalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusalReason

`func (o *CheckoutBalanceCheckResponse) SetRefusalReason(v string)`

SetRefusalReason sets RefusalReason field to given value.

### HasRefusalReason

`func (o *CheckoutBalanceCheckResponse) HasRefusalReason() bool`

HasRefusalReason returns a boolean if a field has been set.

### GetResultCode

`func (o *CheckoutBalanceCheckResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *CheckoutBalanceCheckResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *CheckoutBalanceCheckResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.


### GetTransactionLimit

`func (o *CheckoutBalanceCheckResponse) GetTransactionLimit() Amount`

GetTransactionLimit returns the TransactionLimit field if non-nil, zero value otherwise.

### GetTransactionLimitOk

`func (o *CheckoutBalanceCheckResponse) GetTransactionLimitOk() (*Amount, bool)`

GetTransactionLimitOk returns a tuple with the TransactionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionLimit

`func (o *CheckoutBalanceCheckResponse) SetTransactionLimit(v Amount)`

SetTransactionLimit sets TransactionLimit field to given value.

### HasTransactionLimit

`func (o *CheckoutBalanceCheckResponse) HasTransactionLimit() bool`

HasTransactionLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


