# PaymentCompletionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MD** | Pointer to **string** | A payment session identifier returned by the card issuer. | [optional] 
**PaReq** | Pointer to **string** | (3D) Payment Authentication Request data for the card issuer. | [optional] 
**PaRes** | Pointer to **string** | (3D) Payment Authentication Response data by the card issuer. | [optional] 
**BillingToken** | Pointer to **string** | PayPal-generated token for recurring payments. | [optional] 
**CupsecureplusSmscode** | Pointer to **string** | The SMS verification code collected from the shopper. | [optional] 
**FacilitatorAccessToken** | Pointer to **string** | PayPal-generated third party access token. | [optional] 
**OneTimePasscode** | Pointer to **string** | A random number sent to the mobile phone number of the shopper to verify the payment. | [optional] 
**OrderID** | Pointer to **string** | PayPal-assigned ID for the order. | [optional] 
**PayerID** | Pointer to **string** | PayPal-assigned ID for the payer (shopper). | [optional] 
**Payload** | Pointer to **string** | Payload appended to the &#x60;returnURL&#x60; as a result of the redirect. | [optional] 
**PaymentID** | Pointer to **string** | PayPal-generated ID for the payment. | [optional] 
**PaymentStatus** | Pointer to **string** | Value passed from the WeChat MiniProgram &#x60;wx.requestPayment&#x60; **complete** callback. Possible values: any value starting with &#x60;requestPayment:&#x60;. | [optional] 
**RedirectResult** | Pointer to **string** | The result of the redirect as appended to the &#x60;returnURL&#x60;. | [optional] 
**ResultCode** | Pointer to **string** | Value you received from the WeChat Pay SDK. | [optional] 
**ThreeDSResult** | Pointer to **string** | Base64-encoded string returned by the Component after the challenge flow. It contains the following parameters: &#x60;transStatus&#x60;, &#x60;authorisationToken&#x60;. | [optional] 
**Threeds2ChallengeResult** | Pointer to **string** | Base64-encoded string returned by the Component after the challenge flow. It contains the following parameter: &#x60;transStatus&#x60;. | [optional] 
**Threeds2Fingerprint** | Pointer to **string** | Base64-encoded string returned by the Component after the challenge flow. It contains the following parameter: &#x60;threeDSCompInd&#x60;. | [optional] 

## Methods

### NewPaymentCompletionDetails

`func NewPaymentCompletionDetails() *PaymentCompletionDetails`

NewPaymentCompletionDetails instantiates a new PaymentCompletionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentCompletionDetailsWithDefaults

`func NewPaymentCompletionDetailsWithDefaults() *PaymentCompletionDetails`

NewPaymentCompletionDetailsWithDefaults instantiates a new PaymentCompletionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMD

`func (o *PaymentCompletionDetails) GetMD() string`

GetMD returns the MD field if non-nil, zero value otherwise.

### GetMDOk

`func (o *PaymentCompletionDetails) GetMDOk() (*string, bool)`

GetMDOk returns a tuple with the MD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMD

`func (o *PaymentCompletionDetails) SetMD(v string)`

SetMD sets MD field to given value.

### HasMD

`func (o *PaymentCompletionDetails) HasMD() bool`

HasMD returns a boolean if a field has been set.

### GetPaReq

`func (o *PaymentCompletionDetails) GetPaReq() string`

GetPaReq returns the PaReq field if non-nil, zero value otherwise.

### GetPaReqOk

`func (o *PaymentCompletionDetails) GetPaReqOk() (*string, bool)`

GetPaReqOk returns a tuple with the PaReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaReq

`func (o *PaymentCompletionDetails) SetPaReq(v string)`

SetPaReq sets PaReq field to given value.

### HasPaReq

`func (o *PaymentCompletionDetails) HasPaReq() bool`

HasPaReq returns a boolean if a field has been set.

### GetPaRes

`func (o *PaymentCompletionDetails) GetPaRes() string`

GetPaRes returns the PaRes field if non-nil, zero value otherwise.

### GetPaResOk

`func (o *PaymentCompletionDetails) GetPaResOk() (*string, bool)`

GetPaResOk returns a tuple with the PaRes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaRes

`func (o *PaymentCompletionDetails) SetPaRes(v string)`

SetPaRes sets PaRes field to given value.

### HasPaRes

`func (o *PaymentCompletionDetails) HasPaRes() bool`

HasPaRes returns a boolean if a field has been set.

### GetBillingToken

`func (o *PaymentCompletionDetails) GetBillingToken() string`

GetBillingToken returns the BillingToken field if non-nil, zero value otherwise.

### GetBillingTokenOk

`func (o *PaymentCompletionDetails) GetBillingTokenOk() (*string, bool)`

GetBillingTokenOk returns a tuple with the BillingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingToken

`func (o *PaymentCompletionDetails) SetBillingToken(v string)`

SetBillingToken sets BillingToken field to given value.

### HasBillingToken

`func (o *PaymentCompletionDetails) HasBillingToken() bool`

HasBillingToken returns a boolean if a field has been set.

### GetCupsecureplusSmscode

`func (o *PaymentCompletionDetails) GetCupsecureplusSmscode() string`

GetCupsecureplusSmscode returns the CupsecureplusSmscode field if non-nil, zero value otherwise.

### GetCupsecureplusSmscodeOk

`func (o *PaymentCompletionDetails) GetCupsecureplusSmscodeOk() (*string, bool)`

GetCupsecureplusSmscodeOk returns a tuple with the CupsecureplusSmscode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCupsecureplusSmscode

`func (o *PaymentCompletionDetails) SetCupsecureplusSmscode(v string)`

SetCupsecureplusSmscode sets CupsecureplusSmscode field to given value.

### HasCupsecureplusSmscode

`func (o *PaymentCompletionDetails) HasCupsecureplusSmscode() bool`

HasCupsecureplusSmscode returns a boolean if a field has been set.

### GetFacilitatorAccessToken

`func (o *PaymentCompletionDetails) GetFacilitatorAccessToken() string`

GetFacilitatorAccessToken returns the FacilitatorAccessToken field if non-nil, zero value otherwise.

### GetFacilitatorAccessTokenOk

`func (o *PaymentCompletionDetails) GetFacilitatorAccessTokenOk() (*string, bool)`

GetFacilitatorAccessTokenOk returns a tuple with the FacilitatorAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilitatorAccessToken

`func (o *PaymentCompletionDetails) SetFacilitatorAccessToken(v string)`

SetFacilitatorAccessToken sets FacilitatorAccessToken field to given value.

### HasFacilitatorAccessToken

`func (o *PaymentCompletionDetails) HasFacilitatorAccessToken() bool`

HasFacilitatorAccessToken returns a boolean if a field has been set.

### GetOneTimePasscode

`func (o *PaymentCompletionDetails) GetOneTimePasscode() string`

GetOneTimePasscode returns the OneTimePasscode field if non-nil, zero value otherwise.

### GetOneTimePasscodeOk

`func (o *PaymentCompletionDetails) GetOneTimePasscodeOk() (*string, bool)`

GetOneTimePasscodeOk returns a tuple with the OneTimePasscode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneTimePasscode

`func (o *PaymentCompletionDetails) SetOneTimePasscode(v string)`

SetOneTimePasscode sets OneTimePasscode field to given value.

### HasOneTimePasscode

`func (o *PaymentCompletionDetails) HasOneTimePasscode() bool`

HasOneTimePasscode returns a boolean if a field has been set.

### GetOrderID

`func (o *PaymentCompletionDetails) GetOrderID() string`

GetOrderID returns the OrderID field if non-nil, zero value otherwise.

### GetOrderIDOk

`func (o *PaymentCompletionDetails) GetOrderIDOk() (*string, bool)`

GetOrderIDOk returns a tuple with the OrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderID

`func (o *PaymentCompletionDetails) SetOrderID(v string)`

SetOrderID sets OrderID field to given value.

### HasOrderID

`func (o *PaymentCompletionDetails) HasOrderID() bool`

HasOrderID returns a boolean if a field has been set.

### GetPayerID

`func (o *PaymentCompletionDetails) GetPayerID() string`

GetPayerID returns the PayerID field if non-nil, zero value otherwise.

### GetPayerIDOk

`func (o *PaymentCompletionDetails) GetPayerIDOk() (*string, bool)`

GetPayerIDOk returns a tuple with the PayerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerID

`func (o *PaymentCompletionDetails) SetPayerID(v string)`

SetPayerID sets PayerID field to given value.

### HasPayerID

`func (o *PaymentCompletionDetails) HasPayerID() bool`

HasPayerID returns a boolean if a field has been set.

### GetPayload

`func (o *PaymentCompletionDetails) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *PaymentCompletionDetails) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *PaymentCompletionDetails) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *PaymentCompletionDetails) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetPaymentID

`func (o *PaymentCompletionDetails) GetPaymentID() string`

GetPaymentID returns the PaymentID field if non-nil, zero value otherwise.

### GetPaymentIDOk

`func (o *PaymentCompletionDetails) GetPaymentIDOk() (*string, bool)`

GetPaymentIDOk returns a tuple with the PaymentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentID

`func (o *PaymentCompletionDetails) SetPaymentID(v string)`

SetPaymentID sets PaymentID field to given value.

### HasPaymentID

`func (o *PaymentCompletionDetails) HasPaymentID() bool`

HasPaymentID returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *PaymentCompletionDetails) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *PaymentCompletionDetails) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *PaymentCompletionDetails) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *PaymentCompletionDetails) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetRedirectResult

`func (o *PaymentCompletionDetails) GetRedirectResult() string`

GetRedirectResult returns the RedirectResult field if non-nil, zero value otherwise.

### GetRedirectResultOk

`func (o *PaymentCompletionDetails) GetRedirectResultOk() (*string, bool)`

GetRedirectResultOk returns a tuple with the RedirectResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectResult

`func (o *PaymentCompletionDetails) SetRedirectResult(v string)`

SetRedirectResult sets RedirectResult field to given value.

### HasRedirectResult

`func (o *PaymentCompletionDetails) HasRedirectResult() bool`

HasRedirectResult returns a boolean if a field has been set.

### GetResultCode

`func (o *PaymentCompletionDetails) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *PaymentCompletionDetails) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *PaymentCompletionDetails) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *PaymentCompletionDetails) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetThreeDSResult

`func (o *PaymentCompletionDetails) GetThreeDSResult() string`

GetThreeDSResult returns the ThreeDSResult field if non-nil, zero value otherwise.

### GetThreeDSResultOk

`func (o *PaymentCompletionDetails) GetThreeDSResultOk() (*string, bool)`

GetThreeDSResultOk returns a tuple with the ThreeDSResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSResult

`func (o *PaymentCompletionDetails) SetThreeDSResult(v string)`

SetThreeDSResult sets ThreeDSResult field to given value.

### HasThreeDSResult

`func (o *PaymentCompletionDetails) HasThreeDSResult() bool`

HasThreeDSResult returns a boolean if a field has been set.

### GetThreeds2ChallengeResult

`func (o *PaymentCompletionDetails) GetThreeds2ChallengeResult() string`

GetThreeds2ChallengeResult returns the Threeds2ChallengeResult field if non-nil, zero value otherwise.

### GetThreeds2ChallengeResultOk

`func (o *PaymentCompletionDetails) GetThreeds2ChallengeResultOk() (*string, bool)`

GetThreeds2ChallengeResultOk returns a tuple with the Threeds2ChallengeResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeds2ChallengeResult

`func (o *PaymentCompletionDetails) SetThreeds2ChallengeResult(v string)`

SetThreeds2ChallengeResult sets Threeds2ChallengeResult field to given value.

### HasThreeds2ChallengeResult

`func (o *PaymentCompletionDetails) HasThreeds2ChallengeResult() bool`

HasThreeds2ChallengeResult returns a boolean if a field has been set.

### GetThreeds2Fingerprint

`func (o *PaymentCompletionDetails) GetThreeds2Fingerprint() string`

GetThreeds2Fingerprint returns the Threeds2Fingerprint field if non-nil, zero value otherwise.

### GetThreeds2FingerprintOk

`func (o *PaymentCompletionDetails) GetThreeds2FingerprintOk() (*string, bool)`

GetThreeds2FingerprintOk returns a tuple with the Threeds2Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeds2Fingerprint

`func (o *PaymentCompletionDetails) SetThreeds2Fingerprint(v string)`

SetThreeds2Fingerprint sets Threeds2Fingerprint field to given value.

### HasThreeds2Fingerprint

`func (o *PaymentCompletionDetails) HasThreeds2Fingerprint() bool`

HasThreeds2Fingerprint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


