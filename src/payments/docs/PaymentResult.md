# PaymentResult

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

### NewPaymentResult

`func NewPaymentResult() *PaymentResult`

NewPaymentResult instantiates a new PaymentResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentResultWithDefaults

`func NewPaymentResultWithDefaults() *PaymentResult`

NewPaymentResultWithDefaults instantiates a new PaymentResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *PaymentResult) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *PaymentResult) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *PaymentResult) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *PaymentResult) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetAuthCode

`func (o *PaymentResult) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *PaymentResult) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *PaymentResult) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *PaymentResult) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetDccAmount

`func (o *PaymentResult) GetDccAmount() Amount`

GetDccAmount returns the DccAmount field if non-nil, zero value otherwise.

### GetDccAmountOk

`func (o *PaymentResult) GetDccAmountOk() (*Amount, bool)`

GetDccAmountOk returns a tuple with the DccAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDccAmount

`func (o *PaymentResult) SetDccAmount(v Amount)`

SetDccAmount sets DccAmount field to given value.

### HasDccAmount

`func (o *PaymentResult) HasDccAmount() bool`

HasDccAmount returns a boolean if a field has been set.

### GetDccSignature

`func (o *PaymentResult) GetDccSignature() string`

GetDccSignature returns the DccSignature field if non-nil, zero value otherwise.

### GetDccSignatureOk

`func (o *PaymentResult) GetDccSignatureOk() (*string, bool)`

GetDccSignatureOk returns a tuple with the DccSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDccSignature

`func (o *PaymentResult) SetDccSignature(v string)`

SetDccSignature sets DccSignature field to given value.

### HasDccSignature

`func (o *PaymentResult) HasDccSignature() bool`

HasDccSignature returns a boolean if a field has been set.

### GetFraudResult

`func (o *PaymentResult) GetFraudResult() FraudResult`

GetFraudResult returns the FraudResult field if non-nil, zero value otherwise.

### GetFraudResultOk

`func (o *PaymentResult) GetFraudResultOk() (*FraudResult, bool)`

GetFraudResultOk returns a tuple with the FraudResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudResult

`func (o *PaymentResult) SetFraudResult(v FraudResult)`

SetFraudResult sets FraudResult field to given value.

### HasFraudResult

`func (o *PaymentResult) HasFraudResult() bool`

HasFraudResult returns a boolean if a field has been set.

### GetIssuerUrl

`func (o *PaymentResult) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *PaymentResult) GetIssuerUrlOk() (*string, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *PaymentResult) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.

### HasIssuerUrl

`func (o *PaymentResult) HasIssuerUrl() bool`

HasIssuerUrl returns a boolean if a field has been set.

### GetMd

`func (o *PaymentResult) GetMd() string`

GetMd returns the Md field if non-nil, zero value otherwise.

### GetMdOk

`func (o *PaymentResult) GetMdOk() (*string, bool)`

GetMdOk returns a tuple with the Md field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd

`func (o *PaymentResult) SetMd(v string)`

SetMd sets Md field to given value.

### HasMd

`func (o *PaymentResult) HasMd() bool`

HasMd returns a boolean if a field has been set.

### GetPaRequest

`func (o *PaymentResult) GetPaRequest() string`

GetPaRequest returns the PaRequest field if non-nil, zero value otherwise.

### GetPaRequestOk

`func (o *PaymentResult) GetPaRequestOk() (*string, bool)`

GetPaRequestOk returns a tuple with the PaRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaRequest

`func (o *PaymentResult) SetPaRequest(v string)`

SetPaRequest sets PaRequest field to given value.

### HasPaRequest

`func (o *PaymentResult) HasPaRequest() bool`

HasPaRequest returns a boolean if a field has been set.

### GetPspReference

`func (o *PaymentResult) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *PaymentResult) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *PaymentResult) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.

### HasPspReference

`func (o *PaymentResult) HasPspReference() bool`

HasPspReference returns a boolean if a field has been set.

### GetRefusalReason

`func (o *PaymentResult) GetRefusalReason() string`

GetRefusalReason returns the RefusalReason field if non-nil, zero value otherwise.

### GetRefusalReasonOk

`func (o *PaymentResult) GetRefusalReasonOk() (*string, bool)`

GetRefusalReasonOk returns a tuple with the RefusalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusalReason

`func (o *PaymentResult) SetRefusalReason(v string)`

SetRefusalReason sets RefusalReason field to given value.

### HasRefusalReason

`func (o *PaymentResult) HasRefusalReason() bool`

HasRefusalReason returns a boolean if a field has been set.

### GetResultCode

`func (o *PaymentResult) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *PaymentResult) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *PaymentResult) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *PaymentResult) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


