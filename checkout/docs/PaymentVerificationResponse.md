# PaymentVerificationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **map[string]string** | Contains additional information about the payment. Some data fields are included only if you select them first: Go to **Customer Area** &gt; **Developers** &gt; **Additional data**. | [optional] 
**FraudResult** | Pointer to [**FraudResult**](FraudResult.md) |  | [optional] 
**MerchantReference** | **string** | A unique value that you provided in the initial &#x60;/paymentSession&#x60; request as a &#x60;reference&#x60; field. | 
**Order** | Pointer to [**CheckoutOrderResponse**](CheckoutOrderResponse.md) |  | [optional] 
**PspReference** | Pointer to **string** | Adyen&#39;s 16-character reference associated with the transaction/request. This value is globally unique; quote it when communicating with us about this request. | [optional] 
**RefusalReason** | Pointer to **string** | If the payment&#39;s authorisation is refused or an error occurs during authorisation, this field holds Adyen&#39;s mapped reason for the refusal or a description of the error. When a transaction fails, the authorisation response includes &#x60;resultCode&#x60; and &#x60;refusalReason&#x60; values.  For more information, see [Refusal reasons](https://docs.adyen.com/development-resources/refusal-reasons). | [optional] 
**RefusalReasonCode** | Pointer to **string** | Code that specifies the refusal reason. For more information, see [Authorisation refusal reasons](https://docs.adyen.com/development-resources/refusal-reasons). | [optional] 
**ResultCode** | Pointer to **string** | The result of the payment. For more information, see [Result codes](https://docs.adyen.com/online-payments/payment-result-codes).  Possible values:  * **AuthenticationFinished** – The payment has been successfully authenticated with 3D Secure 2. Returned for 3D Secure 2 authentication-only transactions. * **AuthenticationNotRequired** – The transaction does not require 3D Secure authentication. Returned for [standalone authentication-only integrations](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only). * **Authorised** – The payment was successfully authorised. This state serves as an indicator to proceed with the delivery of goods and services. This is a final state. * **Cancelled** – Indicates the payment has been cancelled (either by the shopper or the merchant) before processing was completed. This is a final state. * **ChallengeShopper** – The issuer requires further shopper interaction before the payment can be authenticated. Returned for 3D Secure 2 transactions. * **Error** – There was an error when the payment was being processed. The reason is given in the &#x60;refusalReason&#x60; field. This is a final state. * **IdentifyShopper** – The issuer requires the shopper&#39;s device fingerprint before the payment can be authenticated. Returned for 3D Secure 2 transactions. * **Pending** – Indicates that it is not possible to obtain the final status of the payment. This can happen if the systems providing final status information for the payment are unavailable, or if the shopper needs to take further action to complete the payment. * **PresentToShopper** – Indicates that the response contains additional information that you need to present to a shopper, so that they can use it to complete a payment. * **Received** – Indicates the payment has successfully been received by Adyen, and will be processed. This is the initial state for all payments. * **RedirectShopper** – Indicates the shopper should be redirected to an external web page or app to complete the authorisation. * **Refused** – Indicates the payment was refused. The reason is given in the &#x60;refusalReason&#x60; field. This is a final state. | [optional] 
**ServiceError** | Pointer to [**ServiceError2**](ServiceError2.md) |  | [optional] 
**ShopperLocale** | **string** | The shopperLocale value provided in the payment request. | 

## Methods

### NewPaymentVerificationResponse

`func NewPaymentVerificationResponse(merchantReference string, shopperLocale string, ) *PaymentVerificationResponse`

NewPaymentVerificationResponse instantiates a new PaymentVerificationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentVerificationResponseWithDefaults

`func NewPaymentVerificationResponseWithDefaults() *PaymentVerificationResponse`

NewPaymentVerificationResponseWithDefaults instantiates a new PaymentVerificationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *PaymentVerificationResponse) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *PaymentVerificationResponse) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *PaymentVerificationResponse) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *PaymentVerificationResponse) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetFraudResult

`func (o *PaymentVerificationResponse) GetFraudResult() FraudResult`

GetFraudResult returns the FraudResult field if non-nil, zero value otherwise.

### GetFraudResultOk

`func (o *PaymentVerificationResponse) GetFraudResultOk() (*FraudResult, bool)`

GetFraudResultOk returns a tuple with the FraudResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudResult

`func (o *PaymentVerificationResponse) SetFraudResult(v FraudResult)`

SetFraudResult sets FraudResult field to given value.

### HasFraudResult

`func (o *PaymentVerificationResponse) HasFraudResult() bool`

HasFraudResult returns a boolean if a field has been set.

### GetMerchantReference

`func (o *PaymentVerificationResponse) GetMerchantReference() string`

GetMerchantReference returns the MerchantReference field if non-nil, zero value otherwise.

### GetMerchantReferenceOk

`func (o *PaymentVerificationResponse) GetMerchantReferenceOk() (*string, bool)`

GetMerchantReferenceOk returns a tuple with the MerchantReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantReference

`func (o *PaymentVerificationResponse) SetMerchantReference(v string)`

SetMerchantReference sets MerchantReference field to given value.


### GetOrder

`func (o *PaymentVerificationResponse) GetOrder() CheckoutOrderResponse`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PaymentVerificationResponse) GetOrderOk() (*CheckoutOrderResponse, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PaymentVerificationResponse) SetOrder(v CheckoutOrderResponse)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PaymentVerificationResponse) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPspReference

`func (o *PaymentVerificationResponse) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *PaymentVerificationResponse) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *PaymentVerificationResponse) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.

### HasPspReference

`func (o *PaymentVerificationResponse) HasPspReference() bool`

HasPspReference returns a boolean if a field has been set.

### GetRefusalReason

`func (o *PaymentVerificationResponse) GetRefusalReason() string`

GetRefusalReason returns the RefusalReason field if non-nil, zero value otherwise.

### GetRefusalReasonOk

`func (o *PaymentVerificationResponse) GetRefusalReasonOk() (*string, bool)`

GetRefusalReasonOk returns a tuple with the RefusalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusalReason

`func (o *PaymentVerificationResponse) SetRefusalReason(v string)`

SetRefusalReason sets RefusalReason field to given value.

### HasRefusalReason

`func (o *PaymentVerificationResponse) HasRefusalReason() bool`

HasRefusalReason returns a boolean if a field has been set.

### GetRefusalReasonCode

`func (o *PaymentVerificationResponse) GetRefusalReasonCode() string`

GetRefusalReasonCode returns the RefusalReasonCode field if non-nil, zero value otherwise.

### GetRefusalReasonCodeOk

`func (o *PaymentVerificationResponse) GetRefusalReasonCodeOk() (*string, bool)`

GetRefusalReasonCodeOk returns a tuple with the RefusalReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusalReasonCode

`func (o *PaymentVerificationResponse) SetRefusalReasonCode(v string)`

SetRefusalReasonCode sets RefusalReasonCode field to given value.

### HasRefusalReasonCode

`func (o *PaymentVerificationResponse) HasRefusalReasonCode() bool`

HasRefusalReasonCode returns a boolean if a field has been set.

### GetResultCode

`func (o *PaymentVerificationResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *PaymentVerificationResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *PaymentVerificationResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *PaymentVerificationResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetServiceError

`func (o *PaymentVerificationResponse) GetServiceError() ServiceError2`

GetServiceError returns the ServiceError field if non-nil, zero value otherwise.

### GetServiceErrorOk

`func (o *PaymentVerificationResponse) GetServiceErrorOk() (*ServiceError2, bool)`

GetServiceErrorOk returns a tuple with the ServiceError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceError

`func (o *PaymentVerificationResponse) SetServiceError(v ServiceError2)`

SetServiceError sets ServiceError field to given value.

### HasServiceError

`func (o *PaymentVerificationResponse) HasServiceError() bool`

HasServiceError returns a boolean if a field has been set.

### GetShopperLocale

`func (o *PaymentVerificationResponse) GetShopperLocale() string`

GetShopperLocale returns the ShopperLocale field if non-nil, zero value otherwise.

### GetShopperLocaleOk

`func (o *PaymentVerificationResponse) GetShopperLocaleOk() (*string, bool)`

GetShopperLocaleOk returns a tuple with the ShopperLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperLocale

`func (o *PaymentVerificationResponse) SetShopperLocale(v string)`

SetShopperLocale sets ShopperLocale field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


