# StoredValueBalanceMergeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthCode** | Pointer to **string** | Authorisation code: * When the payment is authorised, this field holds the authorisation code for the payment. * When the payment is not authorised, this field is empty. | [optional] 
**CurrentBalance** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**PspReference** | Pointer to **string** | Adyen&#39;s 16-character string reference associated with the transaction/request. This value is globally unique; quote it when communicating with us about this request. | [optional] 
**RefusalReason** | Pointer to **string** | If the transaction is refused or an error occurs, this field holds Adyen&#39;s mapped reason for the refusal or a description of the error.  When a transaction fails, the authorisation response includes &#x60;resultCode&#x60; and &#x60;refusalReason&#x60; values. | [optional] 
**ResultCode** | Pointer to **string** | The result of the payment. Possible values:  * **Success** – The operation has been completed successfully.  * **Refused** – The operation was refused. The reason is given in the &#x60;refusalReason&#x60; field.  * **Error** – There was an error when the operation was processed. The reason is given in the &#x60;refusalReason&#x60; field.  * **NotEnoughBalance** – The amount on the payment method is lower than the amount given in the request. Only applicable to balance checks.   | [optional] 
**ThirdPartyRefusalReason** | Pointer to **string** | Raw refusal reason received from the third party, where available | [optional] 

## Methods

### NewStoredValueBalanceMergeResponse

`func NewStoredValueBalanceMergeResponse() *StoredValueBalanceMergeResponse`

NewStoredValueBalanceMergeResponse instantiates a new StoredValueBalanceMergeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoredValueBalanceMergeResponseWithDefaults

`func NewStoredValueBalanceMergeResponseWithDefaults() *StoredValueBalanceMergeResponse`

NewStoredValueBalanceMergeResponseWithDefaults instantiates a new StoredValueBalanceMergeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthCode

`func (o *StoredValueBalanceMergeResponse) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *StoredValueBalanceMergeResponse) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *StoredValueBalanceMergeResponse) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *StoredValueBalanceMergeResponse) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetCurrentBalance

`func (o *StoredValueBalanceMergeResponse) GetCurrentBalance() Amount`

GetCurrentBalance returns the CurrentBalance field if non-nil, zero value otherwise.

### GetCurrentBalanceOk

`func (o *StoredValueBalanceMergeResponse) GetCurrentBalanceOk() (*Amount, bool)`

GetCurrentBalanceOk returns a tuple with the CurrentBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBalance

`func (o *StoredValueBalanceMergeResponse) SetCurrentBalance(v Amount)`

SetCurrentBalance sets CurrentBalance field to given value.

### HasCurrentBalance

`func (o *StoredValueBalanceMergeResponse) HasCurrentBalance() bool`

HasCurrentBalance returns a boolean if a field has been set.

### GetPspReference

`func (o *StoredValueBalanceMergeResponse) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *StoredValueBalanceMergeResponse) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *StoredValueBalanceMergeResponse) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.

### HasPspReference

`func (o *StoredValueBalanceMergeResponse) HasPspReference() bool`

HasPspReference returns a boolean if a field has been set.

### GetRefusalReason

`func (o *StoredValueBalanceMergeResponse) GetRefusalReason() string`

GetRefusalReason returns the RefusalReason field if non-nil, zero value otherwise.

### GetRefusalReasonOk

`func (o *StoredValueBalanceMergeResponse) GetRefusalReasonOk() (*string, bool)`

GetRefusalReasonOk returns a tuple with the RefusalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusalReason

`func (o *StoredValueBalanceMergeResponse) SetRefusalReason(v string)`

SetRefusalReason sets RefusalReason field to given value.

### HasRefusalReason

`func (o *StoredValueBalanceMergeResponse) HasRefusalReason() bool`

HasRefusalReason returns a boolean if a field has been set.

### GetResultCode

`func (o *StoredValueBalanceMergeResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *StoredValueBalanceMergeResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *StoredValueBalanceMergeResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *StoredValueBalanceMergeResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetThirdPartyRefusalReason

`func (o *StoredValueBalanceMergeResponse) GetThirdPartyRefusalReason() string`

GetThirdPartyRefusalReason returns the ThirdPartyRefusalReason field if non-nil, zero value otherwise.

### GetThirdPartyRefusalReasonOk

`func (o *StoredValueBalanceMergeResponse) GetThirdPartyRefusalReasonOk() (*string, bool)`

GetThirdPartyRefusalReasonOk returns a tuple with the ThirdPartyRefusalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyRefusalReason

`func (o *StoredValueBalanceMergeResponse) SetThirdPartyRefusalReason(v string)`

SetThirdPartyRefusalReason sets ThirdPartyRefusalReason field to given value.

### HasThirdPartyRefusalReason

`func (o *StoredValueBalanceMergeResponse) HasThirdPartyRefusalReason() bool`

HasThirdPartyRefusalReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


