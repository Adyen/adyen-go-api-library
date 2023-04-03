# AdditionalDataCommon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestedTestErrorResponseCode** | Pointer to **string** | Triggers test scenarios that allow to replicate certain communication errors.  Allowed values: * **NO_CONNECTION_AVAILABLE** – There wasn&#39;t a connection available to service the outgoing communication. This is a transient, retriable error since no messaging could be initiated to an issuing system (or third-party acquiring system). Therefore, the header Transient-Error: true is returned in the response. A subsequent request using the same idempotency key will be processed as if it was the first request. * **IOEXCEPTION_RECEIVED** – Something went wrong during transmission of the message or receiving the response. This is a classified as non-transient because the message could have been received by the issuing party and been acted upon. No transient error header is returned. If using idempotency, the (error) response is stored as the final result for the idempotency key. Subsequent messages with the same idempotency key not be processed beyond returning the stored response. | [optional] 
**AllowPartialAuth** | Pointer to **string** | Set to true to authorise a part of the requested amount in case the cardholder does not have enough funds on their account.  If a payment was partially authorised, the response includes resultCode: PartiallyAuthorised and the authorised amount in additionalData.authorisedAmountValue. To enable this functionality, contact our Support Team. | [optional] 
**AuthorisationType** | Pointer to **string** | Flags a card payment request for either pre-authorisation or final authorisation. For more information, refer to [Authorisation types](https://docs.adyen.com/online-payments/adjust-authorisation#authorisation-types).  Allowed values: * **PreAuth** – flags the payment request to be handled as a pre-authorisation. * **FinalAuth** – flags the payment request to be handled as a final authorisation. | [optional] 
**CustomRoutingFlag** | Pointer to **string** | Allows you to determine or override the acquirer account that should be used for the transaction.  If you need to process a payment with an acquirer different from a default one, you can set up a corresponding configuration on the Adyen payments platform. Then you can pass a custom routing flag in a payment request&#39;s additional data to target a specific acquirer.  To enable this functionality, contact [Support](https://www.adyen.help/hc/en-us/requests/new). | [optional] 
**IndustryUsage** | Pointer to **string** | In case of [asynchronous authorisation adjustment](https://docs.adyen.com/online-payments/adjust-authorisation#adjust-authorisation), this field denotes why the additional payment is made.  Possible values:   * **NoShow**: An incremental charge is carried out because of a no-show for a guaranteed reservation.   * **DelayedCharge**: An incremental charge is carried out to process an additional payment after the original services have been rendered and the respective payment has been processed. | [optional] 
**ManualCapture** | Pointer to **string** | Set to **true** to require [manual capture](https://docs.adyen.com/online-payments/capture) for the transaction. | [optional] 
**NetworkTxReference** | Pointer to **string** | Allows you to link the transaction to the original or previous one in a subscription/card-on-file chain. This field is required for token-based transactions where Adyen does not tokenize the card.  Transaction identifier from card schemes, for example, Mastercard Trace ID or the Visa Transaction ID.  Submit the original transaction ID of the contract in your payment request if you are not tokenizing card details with Adyen and are making a merchant-initiated transaction (MIT) for subsequent charges.  Make sure you are sending &#x60;shopperInteraction&#x60; **ContAuth** and &#x60;recurringProcessingModel&#x60; **Subscription** or **UnscheduledCardOnFile** to ensure that the transaction is classified as MIT. | [optional] 
**OverwriteBrand** | Pointer to **string** | Boolean indicator that can be optionally used for performing debit transactions on combo cards (for example, combo cards in Brazil). This is not mandatory but we recommend that you set this to true if you want to use the &#x60;selectedBrand&#x60; value to specify how to process the transaction. | [optional] 
**SubMerchantCity** | Pointer to **string** | This field is required if the transaction is performed by a registered payment facilitator. This field must contain the city of the actual merchant&#39;s address. * Format: alpha-numeric. * Maximum length: 13 characters. | [optional] 
**SubMerchantCountry** | Pointer to **string** | This field is required if the transaction is performed by a registered payment facilitator. This field must contain the three-letter country code of the actual merchant&#39;s address. * Format: alpha-numeric. * Fixed length: 3 characters. | [optional] 
**SubMerchantID** | Pointer to **string** | This field contains an identifier of the actual merchant when a transaction is submitted via a payment facilitator. The payment facilitator must send in this unique ID.  A unique identifier per submerchant that is required if the transaction is performed by a registered payment facilitator. * Format: alpha-numeric. * Fixed length: 15 characters. | [optional] 
**SubMerchantName** | Pointer to **string** | This field is required if the transaction is performed by a registered payment facilitator. This field must contain the name of the actual merchant. * Format: alpha-numeric. * Maximum length: 22 characters. | [optional] 
**SubMerchantPostalCode** | Pointer to **string** | This field is required if the transaction is performed by a registered payment facilitator. This field must contain the postal code of the actual merchant&#39;s address. * Format: alpha-numeric. * Maximum length: 10 characters. | [optional] 
**SubMerchantState** | Pointer to **string** | This field is required if the transaction is performed by a registered payment facilitator, and if applicable to the country. This field must contain the state code of the actual merchant&#39;s address. * Format: alpha-numeric. * Maximum length: 3 characters. | [optional] 
**SubMerchantStreet** | Pointer to **string** | This field is required if the transaction is performed by a registered payment facilitator. This field must contain the street of the actual merchant&#39;s address. * Format: alpha-numeric. * Maximum length: 60 characters. | [optional] 
**SubMerchantTaxId** | Pointer to **string** | This field is required if the transaction is performed by a registered payment facilitator. This field must contain the tax ID of the actual merchant. * Format: alpha-numeric. * Fixed length: 11 or 14 characters. | [optional] 

## Methods

### NewAdditionalDataCommon

`func NewAdditionalDataCommon() *AdditionalDataCommon`

NewAdditionalDataCommon instantiates a new AdditionalDataCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalDataCommonWithDefaults

`func NewAdditionalDataCommonWithDefaults() *AdditionalDataCommon`

NewAdditionalDataCommonWithDefaults instantiates a new AdditionalDataCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestedTestErrorResponseCode

`func (o *AdditionalDataCommon) GetRequestedTestErrorResponseCode() string`

GetRequestedTestErrorResponseCode returns the RequestedTestErrorResponseCode field if non-nil, zero value otherwise.

### GetRequestedTestErrorResponseCodeOk

`func (o *AdditionalDataCommon) GetRequestedTestErrorResponseCodeOk() (*string, bool)`

GetRequestedTestErrorResponseCodeOk returns a tuple with the RequestedTestErrorResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTestErrorResponseCode

`func (o *AdditionalDataCommon) SetRequestedTestErrorResponseCode(v string)`

SetRequestedTestErrorResponseCode sets RequestedTestErrorResponseCode field to given value.

### HasRequestedTestErrorResponseCode

`func (o *AdditionalDataCommon) HasRequestedTestErrorResponseCode() bool`

HasRequestedTestErrorResponseCode returns a boolean if a field has been set.

### GetAllowPartialAuth

`func (o *AdditionalDataCommon) GetAllowPartialAuth() string`

GetAllowPartialAuth returns the AllowPartialAuth field if non-nil, zero value otherwise.

### GetAllowPartialAuthOk

`func (o *AdditionalDataCommon) GetAllowPartialAuthOk() (*string, bool)`

GetAllowPartialAuthOk returns a tuple with the AllowPartialAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPartialAuth

`func (o *AdditionalDataCommon) SetAllowPartialAuth(v string)`

SetAllowPartialAuth sets AllowPartialAuth field to given value.

### HasAllowPartialAuth

`func (o *AdditionalDataCommon) HasAllowPartialAuth() bool`

HasAllowPartialAuth returns a boolean if a field has been set.

### GetAuthorisationType

`func (o *AdditionalDataCommon) GetAuthorisationType() string`

GetAuthorisationType returns the AuthorisationType field if non-nil, zero value otherwise.

### GetAuthorisationTypeOk

`func (o *AdditionalDataCommon) GetAuthorisationTypeOk() (*string, bool)`

GetAuthorisationTypeOk returns a tuple with the AuthorisationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorisationType

`func (o *AdditionalDataCommon) SetAuthorisationType(v string)`

SetAuthorisationType sets AuthorisationType field to given value.

### HasAuthorisationType

`func (o *AdditionalDataCommon) HasAuthorisationType() bool`

HasAuthorisationType returns a boolean if a field has been set.

### GetCustomRoutingFlag

`func (o *AdditionalDataCommon) GetCustomRoutingFlag() string`

GetCustomRoutingFlag returns the CustomRoutingFlag field if non-nil, zero value otherwise.

### GetCustomRoutingFlagOk

`func (o *AdditionalDataCommon) GetCustomRoutingFlagOk() (*string, bool)`

GetCustomRoutingFlagOk returns a tuple with the CustomRoutingFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRoutingFlag

`func (o *AdditionalDataCommon) SetCustomRoutingFlag(v string)`

SetCustomRoutingFlag sets CustomRoutingFlag field to given value.

### HasCustomRoutingFlag

`func (o *AdditionalDataCommon) HasCustomRoutingFlag() bool`

HasCustomRoutingFlag returns a boolean if a field has been set.

### GetIndustryUsage

`func (o *AdditionalDataCommon) GetIndustryUsage() string`

GetIndustryUsage returns the IndustryUsage field if non-nil, zero value otherwise.

### GetIndustryUsageOk

`func (o *AdditionalDataCommon) GetIndustryUsageOk() (*string, bool)`

GetIndustryUsageOk returns a tuple with the IndustryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryUsage

`func (o *AdditionalDataCommon) SetIndustryUsage(v string)`

SetIndustryUsage sets IndustryUsage field to given value.

### HasIndustryUsage

`func (o *AdditionalDataCommon) HasIndustryUsage() bool`

HasIndustryUsage returns a boolean if a field has been set.

### GetManualCapture

`func (o *AdditionalDataCommon) GetManualCapture() string`

GetManualCapture returns the ManualCapture field if non-nil, zero value otherwise.

### GetManualCaptureOk

`func (o *AdditionalDataCommon) GetManualCaptureOk() (*string, bool)`

GetManualCaptureOk returns a tuple with the ManualCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualCapture

`func (o *AdditionalDataCommon) SetManualCapture(v string)`

SetManualCapture sets ManualCapture field to given value.

### HasManualCapture

`func (o *AdditionalDataCommon) HasManualCapture() bool`

HasManualCapture returns a boolean if a field has been set.

### GetNetworkTxReference

`func (o *AdditionalDataCommon) GetNetworkTxReference() string`

GetNetworkTxReference returns the NetworkTxReference field if non-nil, zero value otherwise.

### GetNetworkTxReferenceOk

`func (o *AdditionalDataCommon) GetNetworkTxReferenceOk() (*string, bool)`

GetNetworkTxReferenceOk returns a tuple with the NetworkTxReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTxReference

`func (o *AdditionalDataCommon) SetNetworkTxReference(v string)`

SetNetworkTxReference sets NetworkTxReference field to given value.

### HasNetworkTxReference

`func (o *AdditionalDataCommon) HasNetworkTxReference() bool`

HasNetworkTxReference returns a boolean if a field has been set.

### GetOverwriteBrand

`func (o *AdditionalDataCommon) GetOverwriteBrand() string`

GetOverwriteBrand returns the OverwriteBrand field if non-nil, zero value otherwise.

### GetOverwriteBrandOk

`func (o *AdditionalDataCommon) GetOverwriteBrandOk() (*string, bool)`

GetOverwriteBrandOk returns a tuple with the OverwriteBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteBrand

`func (o *AdditionalDataCommon) SetOverwriteBrand(v string)`

SetOverwriteBrand sets OverwriteBrand field to given value.

### HasOverwriteBrand

`func (o *AdditionalDataCommon) HasOverwriteBrand() bool`

HasOverwriteBrand returns a boolean if a field has been set.

### GetSubMerchantCity

`func (o *AdditionalDataCommon) GetSubMerchantCity() string`

GetSubMerchantCity returns the SubMerchantCity field if non-nil, zero value otherwise.

### GetSubMerchantCityOk

`func (o *AdditionalDataCommon) GetSubMerchantCityOk() (*string, bool)`

GetSubMerchantCityOk returns a tuple with the SubMerchantCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantCity

`func (o *AdditionalDataCommon) SetSubMerchantCity(v string)`

SetSubMerchantCity sets SubMerchantCity field to given value.

### HasSubMerchantCity

`func (o *AdditionalDataCommon) HasSubMerchantCity() bool`

HasSubMerchantCity returns a boolean if a field has been set.

### GetSubMerchantCountry

`func (o *AdditionalDataCommon) GetSubMerchantCountry() string`

GetSubMerchantCountry returns the SubMerchantCountry field if non-nil, zero value otherwise.

### GetSubMerchantCountryOk

`func (o *AdditionalDataCommon) GetSubMerchantCountryOk() (*string, bool)`

GetSubMerchantCountryOk returns a tuple with the SubMerchantCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantCountry

`func (o *AdditionalDataCommon) SetSubMerchantCountry(v string)`

SetSubMerchantCountry sets SubMerchantCountry field to given value.

### HasSubMerchantCountry

`func (o *AdditionalDataCommon) HasSubMerchantCountry() bool`

HasSubMerchantCountry returns a boolean if a field has been set.

### GetSubMerchantID

`func (o *AdditionalDataCommon) GetSubMerchantID() string`

GetSubMerchantID returns the SubMerchantID field if non-nil, zero value otherwise.

### GetSubMerchantIDOk

`func (o *AdditionalDataCommon) GetSubMerchantIDOk() (*string, bool)`

GetSubMerchantIDOk returns a tuple with the SubMerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantID

`func (o *AdditionalDataCommon) SetSubMerchantID(v string)`

SetSubMerchantID sets SubMerchantID field to given value.

### HasSubMerchantID

`func (o *AdditionalDataCommon) HasSubMerchantID() bool`

HasSubMerchantID returns a boolean if a field has been set.

### GetSubMerchantName

`func (o *AdditionalDataCommon) GetSubMerchantName() string`

GetSubMerchantName returns the SubMerchantName field if non-nil, zero value otherwise.

### GetSubMerchantNameOk

`func (o *AdditionalDataCommon) GetSubMerchantNameOk() (*string, bool)`

GetSubMerchantNameOk returns a tuple with the SubMerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantName

`func (o *AdditionalDataCommon) SetSubMerchantName(v string)`

SetSubMerchantName sets SubMerchantName field to given value.

### HasSubMerchantName

`func (o *AdditionalDataCommon) HasSubMerchantName() bool`

HasSubMerchantName returns a boolean if a field has been set.

### GetSubMerchantPostalCode

`func (o *AdditionalDataCommon) GetSubMerchantPostalCode() string`

GetSubMerchantPostalCode returns the SubMerchantPostalCode field if non-nil, zero value otherwise.

### GetSubMerchantPostalCodeOk

`func (o *AdditionalDataCommon) GetSubMerchantPostalCodeOk() (*string, bool)`

GetSubMerchantPostalCodeOk returns a tuple with the SubMerchantPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantPostalCode

`func (o *AdditionalDataCommon) SetSubMerchantPostalCode(v string)`

SetSubMerchantPostalCode sets SubMerchantPostalCode field to given value.

### HasSubMerchantPostalCode

`func (o *AdditionalDataCommon) HasSubMerchantPostalCode() bool`

HasSubMerchantPostalCode returns a boolean if a field has been set.

### GetSubMerchantState

`func (o *AdditionalDataCommon) GetSubMerchantState() string`

GetSubMerchantState returns the SubMerchantState field if non-nil, zero value otherwise.

### GetSubMerchantStateOk

`func (o *AdditionalDataCommon) GetSubMerchantStateOk() (*string, bool)`

GetSubMerchantStateOk returns a tuple with the SubMerchantState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantState

`func (o *AdditionalDataCommon) SetSubMerchantState(v string)`

SetSubMerchantState sets SubMerchantState field to given value.

### HasSubMerchantState

`func (o *AdditionalDataCommon) HasSubMerchantState() bool`

HasSubMerchantState returns a boolean if a field has been set.

### GetSubMerchantStreet

`func (o *AdditionalDataCommon) GetSubMerchantStreet() string`

GetSubMerchantStreet returns the SubMerchantStreet field if non-nil, zero value otherwise.

### GetSubMerchantStreetOk

`func (o *AdditionalDataCommon) GetSubMerchantStreetOk() (*string, bool)`

GetSubMerchantStreetOk returns a tuple with the SubMerchantStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantStreet

`func (o *AdditionalDataCommon) SetSubMerchantStreet(v string)`

SetSubMerchantStreet sets SubMerchantStreet field to given value.

### HasSubMerchantStreet

`func (o *AdditionalDataCommon) HasSubMerchantStreet() bool`

HasSubMerchantStreet returns a boolean if a field has been set.

### GetSubMerchantTaxId

`func (o *AdditionalDataCommon) GetSubMerchantTaxId() string`

GetSubMerchantTaxId returns the SubMerchantTaxId field if non-nil, zero value otherwise.

### GetSubMerchantTaxIdOk

`func (o *AdditionalDataCommon) GetSubMerchantTaxIdOk() (*string, bool)`

GetSubMerchantTaxIdOk returns a tuple with the SubMerchantTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantTaxId

`func (o *AdditionalDataCommon) SetSubMerchantTaxId(v string)`

SetSubMerchantTaxId sets SubMerchantTaxId field to given value.

### HasSubMerchantTaxId

`func (o *AdditionalDataCommon) HasSubMerchantTaxId() bool`

HasSubMerchantTaxId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


