# PaymentDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **map[string]string** | Contains additional information about the payment. Some data fields are included only if you select them first: Go to **Customer Area** &gt; **Developers** &gt; **Additional data**. | [optional] 
**Amount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**DonationToken** | Pointer to **string** | Donation Token containing payment details for Adyen Giving. | [optional] 
**FraudResult** | Pointer to [**FraudResult**](FraudResult.md) |  | [optional] 
**MerchantReference** | Pointer to **string** | The reference used during the /payments request. | [optional] 
**Order** | Pointer to [**CheckoutOrderResponse**](CheckoutOrderResponse.md) |  | [optional] 
**PaymentMethod** | Pointer to [**ResponsePaymentMethod**](ResponsePaymentMethod.md) |  | [optional] 
**PspReference** | Pointer to **string** | Adyen&#39;s 16-character string reference associated with the transaction/request. This value is globally unique; quote it when communicating with us about this request. | [optional] 
**RefusalReason** | Pointer to **string** | If the payment&#39;s authorisation is refused or an error occurs during authorisation, this field holds Adyen&#39;s mapped reason for the refusal or a description of the error. When a transaction fails, the authorisation response includes &#x60;resultCode&#x60; and &#x60;refusalReason&#x60; values.  For more information, see [Refusal reasons](https://docs.adyen.com/development-resources/refusal-reasons). | [optional] 
**RefusalReasonCode** | Pointer to **string** | Code that specifies the refusal reason. For more information, see [Authorisation refusal reasons](https://docs.adyen.com/development-resources/refusal-reasons). | [optional] 
**ResultCode** | Pointer to **string** | The result of the payment. For more information, see [Result codes](https://docs.adyen.com/online-payments/payment-result-codes).  Possible values:  * **AuthenticationFinished** – The payment has been successfully authenticated with 3D Secure 2. Returned for 3D Secure 2 authentication-only transactions. * **AuthenticationNotRequired** – The transaction does not require 3D Secure authentication. Returned for [standalone authentication-only integrations](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only). * **Authorised** – The payment was successfully authorised. This state serves as an indicator to proceed with the delivery of goods and services. This is a final state. * **Cancelled** – Indicates the payment has been cancelled (either by the shopper or the merchant) before processing was completed. This is a final state. * **ChallengeShopper** – The issuer requires further shopper interaction before the payment can be authenticated. Returned for 3D Secure 2 transactions. * **Error** – There was an error when the payment was being processed. The reason is given in the &#x60;refusalReason&#x60; field. This is a final state. * **IdentifyShopper** – The issuer requires the shopper&#39;s device fingerprint before the payment can be authenticated. Returned for 3D Secure 2 transactions. * **Pending** – Indicates that it is not possible to obtain the final status of the payment. This can happen if the systems providing final status information for the payment are unavailable, or if the shopper needs to take further action to complete the payment. * **PresentToShopper** – Indicates that the response contains additional information that you need to present to a shopper, so that they can use it to complete a payment. * **Received** – Indicates the payment has successfully been received by Adyen, and will be processed. This is the initial state for all payments. * **RedirectShopper** – Indicates the shopper should be redirected to an external web page or app to complete the authorisation. * **Refused** – Indicates the payment was refused. The reason is given in the &#x60;refusalReason&#x60; field. This is a final state. | [optional] 
**ShopperLocale** | Pointer to **string** | The shopperLocale. | [optional] 
**ThreeDS2ResponseData** | Pointer to [**ThreeDS2ResponseData**](ThreeDS2ResponseData.md) |  | [optional] 
**ThreeDS2Result** | Pointer to [**ThreeDS2Result**](ThreeDS2Result.md) |  | [optional] 
**ThreeDSPaymentData** | Pointer to **string** | When non-empty, contains a value that you must submit to the &#x60;/payments/details&#x60; endpoint as &#x60;paymentData&#x60;. | [optional] 

## Methods

### NewPaymentDetailsResponse

`func NewPaymentDetailsResponse() *PaymentDetailsResponse`

NewPaymentDetailsResponse instantiates a new PaymentDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentDetailsResponseWithDefaults

`func NewPaymentDetailsResponseWithDefaults() *PaymentDetailsResponse`

NewPaymentDetailsResponseWithDefaults instantiates a new PaymentDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *PaymentDetailsResponse) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *PaymentDetailsResponse) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *PaymentDetailsResponse) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *PaymentDetailsResponse) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetAmount

`func (o *PaymentDetailsResponse) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentDetailsResponse) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentDetailsResponse) SetAmount(v Amount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentDetailsResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDonationToken

`func (o *PaymentDetailsResponse) GetDonationToken() string`

GetDonationToken returns the DonationToken field if non-nil, zero value otherwise.

### GetDonationTokenOk

`func (o *PaymentDetailsResponse) GetDonationTokenOk() (*string, bool)`

GetDonationTokenOk returns a tuple with the DonationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDonationToken

`func (o *PaymentDetailsResponse) SetDonationToken(v string)`

SetDonationToken sets DonationToken field to given value.

### HasDonationToken

`func (o *PaymentDetailsResponse) HasDonationToken() bool`

HasDonationToken returns a boolean if a field has been set.

### GetFraudResult

`func (o *PaymentDetailsResponse) GetFraudResult() FraudResult`

GetFraudResult returns the FraudResult field if non-nil, zero value otherwise.

### GetFraudResultOk

`func (o *PaymentDetailsResponse) GetFraudResultOk() (*FraudResult, bool)`

GetFraudResultOk returns a tuple with the FraudResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudResult

`func (o *PaymentDetailsResponse) SetFraudResult(v FraudResult)`

SetFraudResult sets FraudResult field to given value.

### HasFraudResult

`func (o *PaymentDetailsResponse) HasFraudResult() bool`

HasFraudResult returns a boolean if a field has been set.

### GetMerchantReference

`func (o *PaymentDetailsResponse) GetMerchantReference() string`

GetMerchantReference returns the MerchantReference field if non-nil, zero value otherwise.

### GetMerchantReferenceOk

`func (o *PaymentDetailsResponse) GetMerchantReferenceOk() (*string, bool)`

GetMerchantReferenceOk returns a tuple with the MerchantReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantReference

`func (o *PaymentDetailsResponse) SetMerchantReference(v string)`

SetMerchantReference sets MerchantReference field to given value.

### HasMerchantReference

`func (o *PaymentDetailsResponse) HasMerchantReference() bool`

HasMerchantReference returns a boolean if a field has been set.

### GetOrder

`func (o *PaymentDetailsResponse) GetOrder() CheckoutOrderResponse`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PaymentDetailsResponse) GetOrderOk() (*CheckoutOrderResponse, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PaymentDetailsResponse) SetOrder(v CheckoutOrderResponse)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PaymentDetailsResponse) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *PaymentDetailsResponse) GetPaymentMethod() ResponsePaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PaymentDetailsResponse) GetPaymentMethodOk() (*ResponsePaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PaymentDetailsResponse) SetPaymentMethod(v ResponsePaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *PaymentDetailsResponse) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPspReference

`func (o *PaymentDetailsResponse) GetPspReference() string`

GetPspReference returns the PspReference field if non-nil, zero value otherwise.

### GetPspReferenceOk

`func (o *PaymentDetailsResponse) GetPspReferenceOk() (*string, bool)`

GetPspReferenceOk returns a tuple with the PspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspReference

`func (o *PaymentDetailsResponse) SetPspReference(v string)`

SetPspReference sets PspReference field to given value.

### HasPspReference

`func (o *PaymentDetailsResponse) HasPspReference() bool`

HasPspReference returns a boolean if a field has been set.

### GetRefusalReason

`func (o *PaymentDetailsResponse) GetRefusalReason() string`

GetRefusalReason returns the RefusalReason field if non-nil, zero value otherwise.

### GetRefusalReasonOk

`func (o *PaymentDetailsResponse) GetRefusalReasonOk() (*string, bool)`

GetRefusalReasonOk returns a tuple with the RefusalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusalReason

`func (o *PaymentDetailsResponse) SetRefusalReason(v string)`

SetRefusalReason sets RefusalReason field to given value.

### HasRefusalReason

`func (o *PaymentDetailsResponse) HasRefusalReason() bool`

HasRefusalReason returns a boolean if a field has been set.

### GetRefusalReasonCode

`func (o *PaymentDetailsResponse) GetRefusalReasonCode() string`

GetRefusalReasonCode returns the RefusalReasonCode field if non-nil, zero value otherwise.

### GetRefusalReasonCodeOk

`func (o *PaymentDetailsResponse) GetRefusalReasonCodeOk() (*string, bool)`

GetRefusalReasonCodeOk returns a tuple with the RefusalReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusalReasonCode

`func (o *PaymentDetailsResponse) SetRefusalReasonCode(v string)`

SetRefusalReasonCode sets RefusalReasonCode field to given value.

### HasRefusalReasonCode

`func (o *PaymentDetailsResponse) HasRefusalReasonCode() bool`

HasRefusalReasonCode returns a boolean if a field has been set.

### GetResultCode

`func (o *PaymentDetailsResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *PaymentDetailsResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *PaymentDetailsResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *PaymentDetailsResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetShopperLocale

`func (o *PaymentDetailsResponse) GetShopperLocale() string`

GetShopperLocale returns the ShopperLocale field if non-nil, zero value otherwise.

### GetShopperLocaleOk

`func (o *PaymentDetailsResponse) GetShopperLocaleOk() (*string, bool)`

GetShopperLocaleOk returns a tuple with the ShopperLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperLocale

`func (o *PaymentDetailsResponse) SetShopperLocale(v string)`

SetShopperLocale sets ShopperLocale field to given value.

### HasShopperLocale

`func (o *PaymentDetailsResponse) HasShopperLocale() bool`

HasShopperLocale returns a boolean if a field has been set.

### GetThreeDS2ResponseData

`func (o *PaymentDetailsResponse) GetThreeDS2ResponseData() ThreeDS2ResponseData`

GetThreeDS2ResponseData returns the ThreeDS2ResponseData field if non-nil, zero value otherwise.

### GetThreeDS2ResponseDataOk

`func (o *PaymentDetailsResponse) GetThreeDS2ResponseDataOk() (*ThreeDS2ResponseData, bool)`

GetThreeDS2ResponseDataOk returns a tuple with the ThreeDS2ResponseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDS2ResponseData

`func (o *PaymentDetailsResponse) SetThreeDS2ResponseData(v ThreeDS2ResponseData)`

SetThreeDS2ResponseData sets ThreeDS2ResponseData field to given value.

### HasThreeDS2ResponseData

`func (o *PaymentDetailsResponse) HasThreeDS2ResponseData() bool`

HasThreeDS2ResponseData returns a boolean if a field has been set.

### GetThreeDS2Result

`func (o *PaymentDetailsResponse) GetThreeDS2Result() ThreeDS2Result`

GetThreeDS2Result returns the ThreeDS2Result field if non-nil, zero value otherwise.

### GetThreeDS2ResultOk

`func (o *PaymentDetailsResponse) GetThreeDS2ResultOk() (*ThreeDS2Result, bool)`

GetThreeDS2ResultOk returns a tuple with the ThreeDS2Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDS2Result

`func (o *PaymentDetailsResponse) SetThreeDS2Result(v ThreeDS2Result)`

SetThreeDS2Result sets ThreeDS2Result field to given value.

### HasThreeDS2Result

`func (o *PaymentDetailsResponse) HasThreeDS2Result() bool`

HasThreeDS2Result returns a boolean if a field has been set.

### GetThreeDSPaymentData

`func (o *PaymentDetailsResponse) GetThreeDSPaymentData() string`

GetThreeDSPaymentData returns the ThreeDSPaymentData field if non-nil, zero value otherwise.

### GetThreeDSPaymentDataOk

`func (o *PaymentDetailsResponse) GetThreeDSPaymentDataOk() (*string, bool)`

GetThreeDSPaymentDataOk returns a tuple with the ThreeDSPaymentData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSPaymentData

`func (o *PaymentDetailsResponse) SetThreeDSPaymentData(v string)`

SetThreeDSPaymentData sets ThreeDSPaymentData field to given value.

### HasThreeDSPaymentData

`func (o *PaymentDetailsResponse) HasThreeDSPaymentData() bool`

HasThreeDSPaymentData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


