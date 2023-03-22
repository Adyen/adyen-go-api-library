# PaymentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**PaymentResponseAction**](PaymentResponseAction.md) |  | [optional] 
**AdditionalData** | Pointer to **map[string]string** | Contains additional information about the payment. Some data fields are included only if you select them first: Go to **Customer Area** &gt; **Developers** &gt; **Additional data**. | [optional] 
**Amount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**DonationToken** | Pointer to **string** | Donation Token containing payment details for Adyen Giving. | [optional] 
**FraudResult** | Pointer to [**FraudResult**](FraudResult.md) |  | [optional] 
**MerchantReference** | Pointer to **string** | The reference to uniquely identify a payment. This reference is used in all communication with you about the payment status. We recommend using a unique value per payment; however, it is not a requirement. If you need to provide multiple references for a transaction, separate them with hyphens (\&quot;-\&quot;). Maximum length: 80 characters. | [optional] 
**Order** | Pointer to [**CheckoutOrderResponse**](CheckoutOrderResponse.md) |  | [optional] 
**PaymentMethod** | Pointer to [**ResponsePaymentMethod**](ResponsePaymentMethod.md) |  | [optional] 
**PspReference** | Pointer to **string** | Adyen&#39;s 16-character string reference associated with the transaction/request. This value is globally unique; quote it when communicating with us about this request.  &gt; For payment methods that require a redirect or additional action, you will get this value in the &#x60;/payments/details&#x60; response. | [optional] 
**RefusalReason** | Pointer to **string** | If the payment&#39;s authorisation is refused or an error occurs during authorisation, this field holds Adyen&#39;s mapped reason for the refusal or a description of the error. When a transaction fails, the authorisation response includes &#x60;resultCode&#x60; and &#x60;refusalReason&#x60; values.  For more information, see [Refusal reasons](https://docs.adyen.com/development-resources/refusal-reasons). | [optional] 
**RefusalReasonCode** | Pointer to **string** | Code that specifies the refusal reason. For more information, see [Authorisation refusal reasons](https://docs.adyen.com/development-resources/refusal-reasons). | [optional] 
**ResultCode** | Pointer to **string** | The result of the payment. For more information, see [Result codes](https://docs.adyen.com/online-payments/payment-result-codes).  Possible values:  * **AuthenticationFinished** – The payment has been successfully authenticated with 3D Secure 2. Returned for 3D Secure 2 authentication-only transactions. * **AuthenticationNotRequired** – The transaction does not require 3D Secure authentication. Returned for [standalone authentication-only integrations](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only). * **Authorised** – The payment was successfully authorised. This state serves as an indicator to proceed with the delivery of goods and services. This is a final state. * **Cancelled** – Indicates the payment has been cancelled (either by the shopper or the merchant) before processing was completed. This is a final state. * **ChallengeShopper** – The issuer requires further shopper interaction before the payment can be authenticated. Returned for 3D Secure 2 transactions. * **Error** – There was an error when the payment was being processed. The reason is given in the &#x60;refusalReason&#x60; field. This is a final state. * **IdentifyShopper** – The issuer requires the shopper&#39;s device fingerprint before the payment can be authenticated. Returned for 3D Secure 2 transactions. * **Pending** – Indicates that it is not possible to obtain the final status of the payment. This can happen if the systems providing final status information for the payment are unavailable, or if the shopper needs to take further action to complete the payment. * **PresentToShopper** – Indicates that the response contains additional information that you need to present to a shopper, so that they can use it to complete a payment. * **Received** – Indicates the payment has successfully been received by Adyen, and will be processed. This is the initial state for all payments. * **RedirectShopper** – Indicates the shopper should be redirected to an external web page or app to complete the authorisation. * **Refused** – Indicates the payment was refused. The reason is given in the &#x60;refusalReason&#x60; field. This is a final state. | [optional] 
**ThreeDS2ResponseData** | Pointer to [**ThreeDS2ResponseData**](ThreeDS2ResponseData.md) |  | [optional] 
**ThreeDS2Result** | Pointer to [**ThreeDS2Result**](ThreeDS2Result.md) |  | [optional] 
**ThreeDSPaymentData** | Pointer to **string** | When non-empty, contains a value that you must submit to the &#x60;/payments/details&#x60; endpoint as &#x60;paymentData&#x60;. | [optional] 

## Methods

### NewPaymentResponse

`func NewPaymentResponse() *PaymentResponse`

NewPaymentResponse instantiates a new PaymentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentResponseWithDefaults

`func NewPaymentResponseWithDefaults() *PaymentResponse`

NewPaymentResponseWithDefaults instantiates a new PaymentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PaymentResponse) GetAction() PaymentResponseAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PaymentResponse) GetActionOk() (*PaymentResponseAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PaymentResponse) SetAction(v PaymentResponseAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *PaymentResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAdditionalData

`func (o *PaymentResponse) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *PaymentResponse) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *PaymentResponse) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *PaymentResponse) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetAmount

`func (o *PaymentResponse) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentResponse) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentResponse) SetAmount(v Amount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDonationToken

`func (o *PaymentResponse) GetDonationToken() string`

GetDonationToken returns the DonationToken field if non-nil, zero value otherwise.

### GetDonationTokenOk

`func (o *PaymentResponse) GetDonationTokenOk() (*string, bool)`

GetDonationTokenOk returns a tuple with the DonationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDonationToken

`func (o *PaymentResponse) SetDonationToken(v string)`

SetDonationToken sets DonationToken field to given value.

### HasDonationToken

`func (o *PaymentResponse) HasDonationToken() bool`

HasDonationToken returns a boolean if a field has been set.

### GetFraudResult

`func (o *PaymentResponse) GetFraudResult() FraudResult`

GetFraudResult returns the FraudResult field if non-nil, zero value otherwise.

### GetFraudResultOk

`func (o *PaymentResponse) GetFraudResultOk() (*FraudResult, bool)`

GetFraudResultOk returns a tuple with the FraudResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudResult

`func (o *PaymentResponse) SetFraudResult(v FraudResult)`

SetFraudResult sets FraudResult field to given value.

### HasFraudResult

`func (o *PaymentResponse) HasFraudResult() bool`

HasFraudResult returns a boolean if a field has been set.

### GetMerchantReference

`func (o *PaymentResponse) GetMerchantReference() string`

GetMerchantReference returns the MerchantReference field if non-nil, zero value otherwise.

### GetMerchantReferenceOk

`func (o *PaymentResponse) GetMerchantReferenceOk() (*string, bool)`

GetMerchantReferenceOk returns a tuple with the MerchantReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantReference

`func (o *PaymentResponse) SetMerchantReference(v string)`

SetMerchantReference sets MerchantReference field to given value.

### HasMerchantReference

`func (o *PaymentResponse) HasMerchantReference() bool`

HasMerchantReference returns a boolean if a field has been set.

### GetOrder

`func (o *PaymentResponse) GetOrder() CheckoutOrderResponse`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PaymentResponse) GetOrderOk() (*CheckoutOrderResponse, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PaymentResponse) SetOrder(v CheckoutOrderResponse)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PaymentResponse) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *PaymentResponse) GetPaymentMethod() ResponsePaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PaymentResponse) GetPaymentMethodOk() (*ResponsePaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PaymentResponse) SetPaymentMethod(v ResponsePaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *PaymentResponse) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPspReference

`func (o *PaymentResponse) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *PaymentResponse) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *PaymentResponse) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.

### HasPspReference

`func (o *PaymentResponse) HasPspReference() bool`

HasPspReference returns a boolean if a field has been set.

### GetRefusalReason

`func (o *PaymentResponse) GetRefusalReason() string`

GetRefusalReason returns the RefusalReason field if non-nil, zero value otherwise.

### GetRefusalReasonOk

`func (o *PaymentResponse) GetRefusalReasonOk() (*string, bool)`

GetRefusalReasonOk returns a tuple with the RefusalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusalReason

`func (o *PaymentResponse) SetRefusalReason(v string)`

SetRefusalReason sets RefusalReason field to given value.

### HasRefusalReason

`func (o *PaymentResponse) HasRefusalReason() bool`

HasRefusalReason returns a boolean if a field has been set.

### GetRefusalReasonCode

`func (o *PaymentResponse) GetRefusalReasonCode() string`

GetRefusalReasonCode returns the RefusalReasonCode field if non-nil, zero value otherwise.

### GetRefusalReasonCodeOk

`func (o *PaymentResponse) GetRefusalReasonCodeOk() (*string, bool)`

GetRefusalReasonCodeOk returns a tuple with the RefusalReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusalReasonCode

`func (o *PaymentResponse) SetRefusalReasonCode(v string)`

SetRefusalReasonCode sets RefusalReasonCode field to given value.

### HasRefusalReasonCode

`func (o *PaymentResponse) HasRefusalReasonCode() bool`

HasRefusalReasonCode returns a boolean if a field has been set.

### GetResultCode

`func (o *PaymentResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *PaymentResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *PaymentResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *PaymentResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetThreeDS2ResponseData

`func (o *PaymentResponse) GetThreeDS2ResponseData() ThreeDS2ResponseData`

GetThreeDS2ResponseData returns the ThreeDS2ResponseData field if non-nil, zero value otherwise.

### GetThreeDS2ResponseDataOk

`func (o *PaymentResponse) GetThreeDS2ResponseDataOk() (*ThreeDS2ResponseData, bool)`

GetThreeDS2ResponseDataOk returns a tuple with the ThreeDS2ResponseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDS2ResponseData

`func (o *PaymentResponse) SetThreeDS2ResponseData(v ThreeDS2ResponseData)`

SetThreeDS2ResponseData sets ThreeDS2ResponseData field to given value.

### HasThreeDS2ResponseData

`func (o *PaymentResponse) HasThreeDS2ResponseData() bool`

HasThreeDS2ResponseData returns a boolean if a field has been set.

### GetThreeDS2Result

`func (o *PaymentResponse) GetThreeDS2Result() ThreeDS2Result`

GetThreeDS2Result returns the ThreeDS2Result field if non-nil, zero value otherwise.

### GetThreeDS2ResultOk

`func (o *PaymentResponse) GetThreeDS2ResultOk() (*ThreeDS2Result, bool)`

GetThreeDS2ResultOk returns a tuple with the ThreeDS2Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDS2Result

`func (o *PaymentResponse) SetThreeDS2Result(v ThreeDS2Result)`

SetThreeDS2Result sets ThreeDS2Result field to given value.

### HasThreeDS2Result

`func (o *PaymentResponse) HasThreeDS2Result() bool`

HasThreeDS2Result returns a boolean if a field has been set.

### GetThreeDSPaymentData

`func (o *PaymentResponse) GetThreeDSPaymentData() string`

GetThreeDSPaymentData returns the ThreeDSPaymentData field if non-nil, zero value otherwise.

### GetThreeDSPaymentDataOk

`func (o *PaymentResponse) GetThreeDSPaymentDataOk() (*string, bool)`

GetThreeDSPaymentDataOk returns a tuple with the ThreeDSPaymentData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSPaymentData

`func (o *PaymentResponse) SetThreeDSPaymentData(v string)`

SetThreeDSPaymentData sets ThreeDSPaymentData field to given value.

### HasThreeDSPaymentData

`func (o *PaymentResponse) HasThreeDSPaymentData() bool`

HasThreeDSPaymentData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


