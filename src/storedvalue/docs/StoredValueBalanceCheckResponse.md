# StoredValueBalanceCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentBalance** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**PspReference** | Pointer to **string** | Adyen&#39;s 16-character string reference associated with the transaction/request. This value is globally unique; quote it when communicating with us about this request. | [optional] 
**RefusalReason** | Pointer to **string** | If the transaction is refused or an error occurs, this field holds Adyen&#39;s mapped reason for the refusal or a description of the error.  When a transaction fails, the authorisation response includes &#x60;resultCode&#x60; and &#x60;refusalReason&#x60; values. | [optional] 
**ResultCode** | Pointer to **string** | The result of the payment. Possible values:  * **Success** – The operation has been completed successfully.  * **Refused** – The operation was refused. The reason is given in the &#x60;refusalReason&#x60; field.  * **Error** – There was an error when the operation was processed. The reason is given in the &#x60;refusalReason&#x60; field.  * **NotEnoughBalance** – The amount on the payment method is lower than the amount given in the request. Only applicable to balance checks.   | [optional] 
**ThirdPartyRefusalReason** | Pointer to **string** | Raw refusal reason received from the third party, where available | [optional] 

## Methods

### NewStoredValueBalanceCheckResponse

`func NewStoredValueBalanceCheckResponse() *StoredValueBalanceCheckResponse`

NewStoredValueBalanceCheckResponse instantiates a new StoredValueBalanceCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoredValueBalanceCheckResponseWithDefaults

`func NewStoredValueBalanceCheckResponseWithDefaults() *StoredValueBalanceCheckResponse`

NewStoredValueBalanceCheckResponseWithDefaults instantiates a new StoredValueBalanceCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentBalance

`func (o *StoredValueBalanceCheckResponse) GetCurrentBalance() Amount`

GetCurrentBalance returns the CurrentBalance field if non-nil, zero value otherwise.

### GetCurrentBalanceOk

`func (o *StoredValueBalanceCheckResponse) GetCurrentBalanceOk() (*Amount, bool)`

GetCurrentBalanceOk returns a tuple with the CurrentBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBalance

`func (o *StoredValueBalanceCheckResponse) SetCurrentBalance(v Amount)`

SetCurrentBalance sets CurrentBalance field to given value.

### HasCurrentBalance

`func (o *StoredValueBalanceCheckResponse) HasCurrentBalance() bool`

HasCurrentBalance returns a boolean if a field has been set.

### GetPspReference

`func (o *StoredValueBalanceCheckResponse) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *StoredValueBalanceCheckResponse) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *StoredValueBalanceCheckResponse) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.

### HasPspReference

`func (o *StoredValueBalanceCheckResponse) HasPspReference() bool`

HasPspReference returns a boolean if a field has been set.

### GetRefusalReason

`func (o *StoredValueBalanceCheckResponse) GetRefusalReason() string`

GetRefusalReason returns the RefusalReason field if non-nil, zero value otherwise.

### GetRefusalReasonOk

`func (o *StoredValueBalanceCheckResponse) GetRefusalReasonOk() (*string, bool)`

GetRefusalReasonOk returns a tuple with the RefusalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusalReason

`func (o *StoredValueBalanceCheckResponse) SetRefusalReason(v string)`

SetRefusalReason sets RefusalReason field to given value.

### HasRefusalReason

`func (o *StoredValueBalanceCheckResponse) HasRefusalReason() bool`

HasRefusalReason returns a boolean if a field has been set.

### GetResultCode

`func (o *StoredValueBalanceCheckResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *StoredValueBalanceCheckResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *StoredValueBalanceCheckResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *StoredValueBalanceCheckResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetThirdPartyRefusalReason

`func (o *StoredValueBalanceCheckResponse) GetThirdPartyRefusalReason() string`

GetThirdPartyRefusalReason returns the ThirdPartyRefusalReason field if non-nil, zero value otherwise.

### GetThirdPartyRefusalReasonOk

`func (o *StoredValueBalanceCheckResponse) GetThirdPartyRefusalReasonOk() (*string, bool)`

GetThirdPartyRefusalReasonOk returns a tuple with the ThirdPartyRefusalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyRefusalReason

`func (o *StoredValueBalanceCheckResponse) SetThirdPartyRefusalReason(v string)`

SetThirdPartyRefusalReason sets ThirdPartyRefusalReason field to given value.

### HasThirdPartyRefusalReason

`func (o *StoredValueBalanceCheckResponse) HasThirdPartyRefusalReason() bool`

HasThirdPartyRefusalReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


