# PayoutResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **map[string]string** | Contains additional information about the payment. Some data fields are included only if you select them first: Go to **Customer Area** &gt; **Developers** &gt; **Additional data**. | [optional] 
**AuthCode** | Pointer to **string** | Authorisation code: * When the payment is authorised successfully, this field holds the authorisation code for the payment. * When the payment is not authorised, this field is empty. | [optional] 
**DccAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**DccSignature** | Pointer to **string** | Cryptographic signature used to verify &#x60;dccQuote&#x60;. &gt; This value only applies if you have implemented Dynamic Currency Conversion. For more information, [contact Support](https://www.adyen.help/hc/en-us/requests/new). | [optional] 
**FraudResult** | Pointer to [**FraudResult**](FraudResult.md) |  | [optional] 
**IssuerUrl** | Pointer to **string** | The URL to direct the shopper to. &gt; In case of SecurePlus, do not redirect a shopper to this URL. | [optional] 
**Md** | Pointer to **string** | The payment session. | [optional] 
**PaRequest** | Pointer to **string** | The 3D request data for the issuer.  If the value is **CUPSecurePlus-CollectSMSVerificationCode**, collect an SMS code from the shopper and pass it in the &#x60;/authorise3D&#x60; request. For more information, see [3D Secure](https://docs.adyen.com/classic-integration/3d-secure). | [optional] 
**PspReference** | Pointer to **string** | Adyen&#39;s 16-character reference associated with the transaction/request. This value is globally unique; quote it when communicating with us about this request. | [optional] 
**RefusalReason** | Pointer to **string** | If the payment&#39;s authorisation is refused or an error occurs during authorisation, this field holds Adyen&#39;s mapped reason for the refusal or a description of the error. When a transaction fails, the authorisation response includes &#x60;resultCode&#x60; and &#x60;refusalReason&#x60; values.  For more information, see [Refusal reasons](https://docs.adyen.com/development-resources/refusal-reasons). | [optional] 
**ResultCode** | Pointer to **string** | The result of the payment. For more information, see [Result codes](https://docs.adyen.com/online-payments/payment-result-codes).  Possible values:  * **AuthenticationFinished** – The payment has been successfully authenticated with 3D Secure 2. Returned for 3D Secure 2 authentication-only transactions. * **AuthenticationNotRequired** – The transaction does not require 3D Secure authentication. Returned for [standalone authentication-only integrations](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only). * **Authorised** – The payment was successfully authorised. This state serves as an indicator to proceed with the delivery of goods and services. This is a final state. * **Cancelled** – Indicates the payment has been cancelled (either by the shopper or the merchant) before processing was completed. This is a final state. * **ChallengeShopper** – The issuer requires further shopper interaction before the payment can be authenticated. Returned for 3D Secure 2 transactions. * **Error** – There was an error when the payment was being processed. The reason is given in the &#x60;refusalReason&#x60; field. This is a final state. * **IdentifyShopper** – The issuer requires the shopper&#39;s device fingerprint before the payment can be authenticated. Returned for 3D Secure 2 transactions. * **Pending** – Indicates that it is not possible to obtain the final status of the payment. This can happen if the systems providing final status information for the payment are unavailable, or if the shopper needs to take further action to complete the payment. * **PresentToShopper** – Indicates that the response contains additional information that you need to present to a shopper, so that they can use it to complete a payment. * **Received** – Indicates the payment has successfully been received by Adyen, and will be processed. This is the initial state for all payments. * **RedirectShopper** – Indicates the shopper should be redirected to an external web page or app to complete the authorisation. * **Refused** – Indicates the payment was refused. The reason is given in the &#x60;refusalReason&#x60; field. This is a final state. | [optional] 

## Methods

### NewPayoutResponse

`func NewPayoutResponse() *PayoutResponse`

NewPayoutResponse instantiates a new PayoutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutResponseWithDefaults

`func NewPayoutResponseWithDefaults() *PayoutResponse`

NewPayoutResponseWithDefaults instantiates a new PayoutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *PayoutResponse) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *PayoutResponse) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *PayoutResponse) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *PayoutResponse) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetAuthCode

`func (o *PayoutResponse) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *PayoutResponse) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *PayoutResponse) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *PayoutResponse) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetDccAmount

`func (o *PayoutResponse) GetDccAmount() Amount`

GetDccAmount returns the DccAmount field if non-nil, zero value otherwise.

### GetDccAmountOk

`func (o *PayoutResponse) GetDccAmountOk() (*Amount, bool)`

GetDccAmountOk returns a tuple with the DccAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDccAmount

`func (o *PayoutResponse) SetDccAmount(v Amount)`

SetDccAmount sets DccAmount field to given value.

### HasDccAmount

`func (o *PayoutResponse) HasDccAmount() bool`

HasDccAmount returns a boolean if a field has been set.

### GetDccSignature

`func (o *PayoutResponse) GetDccSignature() string`

GetDccSignature returns the DccSignature field if non-nil, zero value otherwise.

### GetDccSignatureOk

`func (o *PayoutResponse) GetDccSignatureOk() (*string, bool)`

GetDccSignatureOk returns a tuple with the DccSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDccSignature

`func (o *PayoutResponse) SetDccSignature(v string)`

SetDccSignature sets DccSignature field to given value.

### HasDccSignature

`func (o *PayoutResponse) HasDccSignature() bool`

HasDccSignature returns a boolean if a field has been set.

### GetFraudResult

`func (o *PayoutResponse) GetFraudResult() FraudResult`

GetFraudResult returns the FraudResult field if non-nil, zero value otherwise.

### GetFraudResultOk

`func (o *PayoutResponse) GetFraudResultOk() (*FraudResult, bool)`

GetFraudResultOk returns a tuple with the FraudResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudResult

`func (o *PayoutResponse) SetFraudResult(v FraudResult)`

SetFraudResult sets FraudResult field to given value.

### HasFraudResult

`func (o *PayoutResponse) HasFraudResult() bool`

HasFraudResult returns a boolean if a field has been set.

### GetIssuerUrl

`func (o *PayoutResponse) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *PayoutResponse) GetIssuerUrlOk() (*string, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *PayoutResponse) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.

### HasIssuerUrl

`func (o *PayoutResponse) HasIssuerUrl() bool`

HasIssuerUrl returns a boolean if a field has been set.

### GetMd

`func (o *PayoutResponse) GetMd() string`

GetMd returns the Md field if non-nil, zero value otherwise.

### GetMdOk

`func (o *PayoutResponse) GetMdOk() (*string, bool)`

GetMdOk returns a tuple with the Md field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd

`func (o *PayoutResponse) SetMd(v string)`

SetMd sets Md field to given value.

### HasMd

`func (o *PayoutResponse) HasMd() bool`

HasMd returns a boolean if a field has been set.

### GetPaRequest

`func (o *PayoutResponse) GetPaRequest() string`

GetPaRequest returns the PaRequest field if non-nil, zero value otherwise.

### GetPaRequestOk

`func (o *PayoutResponse) GetPaRequestOk() (*string, bool)`

GetPaRequestOk returns a tuple with the PaRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaRequest

`func (o *PayoutResponse) SetPaRequest(v string)`

SetPaRequest sets PaRequest field to given value.

### HasPaRequest

`func (o *PayoutResponse) HasPaRequest() bool`

HasPaRequest returns a boolean if a field has been set.

### GetPspReference

`func (o *PayoutResponse) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *PayoutResponse) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *PayoutResponse) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.

### HasPspReference

`func (o *PayoutResponse) HasPspReference() bool`

HasPspReference returns a boolean if a field has been set.

### GetRefusalReason

`func (o *PayoutResponse) GetRefusalReason() string`

GetRefusalReason returns the RefusalReason field if non-nil, zero value otherwise.

### GetRefusalReasonOk

`func (o *PayoutResponse) GetRefusalReasonOk() (*string, bool)`

GetRefusalReasonOk returns a tuple with the RefusalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusalReason

`func (o *PayoutResponse) SetRefusalReason(v string)`

SetRefusalReason sets RefusalReason field to given value.

### HasRefusalReason

`func (o *PayoutResponse) HasRefusalReason() bool`

HasRefusalReason returns a boolean if a field has been set.

### GetResultCode

`func (o *PayoutResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *PayoutResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *PayoutResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *PayoutResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


