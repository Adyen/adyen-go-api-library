# ResponseAdditionalDataCommon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcquirerAccountCode** | Pointer to **string** | The name of the Adyen acquirer account.  Example: PayPalSandbox_TestAcquirer  &gt; Only relevant for PayPal transactions. | [optional] 
**AcquirerCode** | Pointer to **string** | The name of the acquirer processing the payment request.  Example: TestPmmAcquirer | [optional] 
**AcquirerReference** | Pointer to **string** | The reference number that can be used for reconciliation in case a non-Adyen acquirer is used for settlement.  Example: 7C9N3FNBKT9 | [optional] 
**Alias** | Pointer to **string** | The Adyen alias of the card.  Example: H167852639363479 | [optional] 
**AliasType** | Pointer to **string** | The type of the card alias.  Example: Default | [optional] 
**AuthCode** | Pointer to **string** | Authorisation code: * When the payment is authorised successfully, this field holds the authorisation code for the payment. * When the payment is not authorised, this field is empty.  Example: 58747 | [optional] 
**AuthorisationMid** | Pointer to **string** | Merchant ID known by the acquirer. | [optional] 
**AuthorisedAmountCurrency** | Pointer to **string** | The currency of the authorised amount, as a three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes). | [optional] 
**AuthorisedAmountValue** | Pointer to **string** | Value of the amount authorised.  This amount is represented in minor units according to the [following table](https://docs.adyen.com/development-resources/currency-codes). | [optional] 
**AvsResult** | Pointer to **string** | The AVS result code of the payment, which provides information about the outcome of the AVS check.  For possible values, see [AVS](https://docs.adyen.com/risk-management/configure-standard-risk-rules/consistency-rules#billing-address-does-not-match-cardholder-address-avs). | [optional] 
**AvsResultRaw** | Pointer to **string** | Raw AVS result received from the acquirer, where available.  Example: D | [optional] 
**Bic** | Pointer to **string** | BIC of a bank account.  Example: TESTNL01  &gt; Only relevant for SEPA Direct Debit transactions. | [optional] 
**CoBrandedWith** | Pointer to **string** | Includes the co-branded card information. | [optional] 
**CvcResult** | Pointer to **string** | The result of CVC verification. | [optional] 
**CvcResultRaw** | Pointer to **string** | The raw result of CVC verification. | [optional] 
**DsTransID** | Pointer to **string** | Supported for 3D Secure 2. The unique transaction identifier assigned by the DS to identify a single transaction. | [optional] 
**Eci** | Pointer to **string** | The Electronic Commerce Indicator returned from the schemes for the 3DS payment session.  Example: 02 | [optional] 
**ExpiryDate** | Pointer to **string** | The expiry date on the card.  Example: 6/2016  &gt; Returned only in case of a card payment. | [optional] 
**ExtraCostsCurrency** | Pointer to **string** | The currency of the extra amount charged due to additional amounts set in the skin used in the HPP payment request.  Example: EUR | [optional] 
**ExtraCostsValue** | Pointer to **string** | The value of the extra amount charged due to additional amounts set in the skin used in the HPP payment request. The amount is in minor units. | [optional] 
**FraudCheckItemNrFraudCheckname** | Pointer to **string** | The fraud score due to a particular fraud check. The fraud check name is found in the key of the key-value pair. | [optional] 
**FraudManualReview** | Pointer to **string** | Indicates if the payment is sent to manual review. | [optional] 
**FraudResultType** | Pointer to **string** | The fraud result properties of the payment. | [optional] 
**FundingSource** | Pointer to **string** | Information regarding the funding type of the card. The possible return values are: * CHARGE * CREDIT * DEBIT * PREPAID * PREPAID_RELOADABLE  * PREPAID_NONRELOADABLE * DEFFERED_DEBIT  &gt; This functionality requires additional configuration on Adyen&#39;s end. To enable it, contact the Support Team.  For receiving this field in the notification, enable **Include Funding Source** in **Notifications** &gt; **Additional settings**. | [optional] 
**FundsAvailability** | Pointer to **string** | Indicates availability of funds.  Visa: * \&quot;I\&quot; (fast funds are supported) * \&quot;N\&quot; (otherwise)  Mastercard: * \&quot;I\&quot; (product type is Prepaid or Debit, or issuing country is in CEE/HGEM list) * \&quot;N\&quot; (otherwise)  &gt; Returned when you verify a card BIN or estimate costs, and only if payoutEligible is \&quot;Y\&quot; or \&quot;D\&quot;. | [optional] 
**InferredRefusalReason** | Pointer to **string** | Provides the more granular indication of why a transaction was refused. When a transaction fails with either \&quot;Refused\&quot;, \&quot;Restricted Card\&quot;, \&quot;Transaction Not Permitted\&quot;, \&quot;Not supported\&quot; or \&quot;DeclinedNon Generic\&quot; refusalReason from the issuer, Adyen cross references its PSP-wide data for extra insight into the refusal reason. If an inferred refusal reason is available, the &#x60;inferredRefusalReason&#x60;, field is populated and the &#x60;refusalReason&#x60;, is set to \&quot;Not Supported\&quot;.  Possible values:  * 3D Secure Mandated * Closed Account * ContAuth Not Supported * CVC Mandated * Ecommerce Not Allowed * Crossborder Not Supported * Card Updated  * Low Authrate Bin * Non-reloadable prepaid card | [optional] 
**IsCardCommercial** | Pointer to **string** | Indicates if the card is used for business purposes only. | [optional] 
**IssuerCountry** | Pointer to **string** | The issuing country of the card based on the BIN list that Adyen maintains.  Example: JP | [optional] 
**LiabilityShift** | Pointer to **string** | A Boolean value indicating whether a liability shift was offered for this payment. | [optional] 
**McBankNetReferenceNumber** | Pointer to **string** | The &#x60;mcBankNetReferenceNumber&#x60;, is a minimum of six characters and a maximum of nine characters long.  &gt; Contact Support Team to enable this field. | [optional] 
**MerchantAdviceCode** | Pointer to **string** | The Merchant Advice Code (MAC) can be returned by Mastercard issuers for refused payments. If present, the MAC contains information about why the payment failed, and whether it can be retried.  For more information see [Mastercard Merchant Advice Codes](https://docs.adyen.com/development-resources/raw-acquirer-responses#mastercard-merchant-advice-codes). | [optional] 
**MerchantReference** | Pointer to **string** | The reference provided for the transaction. | [optional] 
**NetworkTxReference** | Pointer to **string** | Returned in the response if you are not tokenizing with Adyen and are using the Merchant-initiated transactions (MIT) framework from Mastercard or Visa.  This contains either the Mastercard Trace ID or the Visa Transaction ID. | [optional] 
**OwnerName** | Pointer to **string** | The owner name of a bank account.  Only relevant for SEPA Direct Debit transactions. | [optional] 
**PaymentAccountReference** | Pointer to **string** | The Payment Account Reference (PAR) value links a network token with the underlying primary account number (PAN). The PAR value consists of 29 uppercase alphanumeric characters. | [optional] 
**PaymentMethod** | Pointer to **string** | The payment method used in the transaction. | [optional] 
**PaymentMethodVariant** | Pointer to **string** | The Adyen sub-variant of the payment method used for the payment request.  For more information, refer to [PaymentMethodVariant](https://docs.adyen.com/development-resources/paymentmethodvariant).  Example: mcpro | [optional] 
**PayoutEligible** | Pointer to **string** | Indicates whether a payout is eligible or not for this card.  Visa: * \&quot;Y\&quot; * \&quot;N\&quot;  Mastercard: * \&quot;Y\&quot; (domestic and cross-border)  * \&quot;D\&quot; (only domestic) * \&quot;N\&quot; (no MoneySend) * \&quot;U\&quot; (unknown) | [optional] 
**RealtimeAccountUpdaterStatus** | Pointer to **string** | The response code from the Real Time Account Updater service.  Possible return values are: * CardChanged * CardExpiryChanged * CloseAccount  * ContactCardAccountHolder | [optional] 
**ReceiptFreeText** | Pointer to **string** | Message to be displayed on the terminal. | [optional] 
**RecurringContractTypes** | Pointer to **string** | The recurring contract types applicable to the transaction. | [optional] 
**RecurringFirstPspReference** | Pointer to **string** | The &#x60;pspReference&#x60;, of the first recurring payment that created the recurring detail.  This functionality requires additional configuration on Adyen&#39;s end. To enable it, contact the Support Team. | [optional] 
**RecurringRecurringDetailReference** | Pointer to **string** | The reference that uniquely identifies the recurring transaction. | [optional] 
**RecurringShopperReference** | Pointer to **string** | The provided reference of the shopper for a recurring transaction. | [optional] 
**RecurringProcessingModel** | Pointer to **string** | The processing model used for the recurring transaction. | [optional] 
**Referred** | Pointer to **string** | If the payment is referred, this field is set to true.  This field is unavailable if the payment is referred and is usually not returned with ecommerce transactions.  Example: true | [optional] 
**RefusalReasonRaw** | Pointer to **string** | Raw refusal reason received from the acquirer, where available.  Example: AUTHORISED | [optional] 
**RequestAmount** | Pointer to **string** | The amount of the payment request. | [optional] 
**RequestCurrencyCode** | Pointer to **string** | The currency of the payment request. | [optional] 
**ShopperInteraction** | Pointer to **string** | The shopper interaction type of the payment request.  Example: Ecommerce | [optional] 
**ShopperReference** | Pointer to **string** | The shopperReference passed in the payment request.  Example: AdyenTestShopperXX | [optional] 
**TerminalId** | Pointer to **string** | The terminal ID used in a point-of-sale payment.  Example: 06022622 | [optional] 
**ThreeDAuthenticated** | Pointer to **string** | A Boolean value indicating whether 3DS authentication was completed on this payment.  Example: true | [optional] 
**ThreeDAuthenticatedResponse** | Pointer to **string** | The raw 3DS authentication result from the card issuer.  Example: N | [optional] 
**ThreeDOffered** | Pointer to **string** | A Boolean value indicating whether 3DS was offered for this payment.  Example: true | [optional] 
**ThreeDOfferedResponse** | Pointer to **string** | The raw enrollment result from the 3DS directory services of the card schemes.  Example: Y | [optional] 
**ThreeDSVersion** | Pointer to **string** | The 3D Secure 2 version. | [optional] 
**VisaTransactionId** | Pointer to **string** | The &#x60;visaTransactionId&#x60;, has a fixed length of 15 numeric characters.  &gt; Contact Support Team to enable this field. | [optional] 
**Xid** | Pointer to **string** | The 3DS transaction ID of the 3DS session sent in notifications. The value is Base64-encoded and is returned for transactions with directoryResponse &#39;N&#39; or &#39;Y&#39;. If you want to submit the xid in your 3D Secure 1 request, use the &#x60;mpiData.xid&#x60;, field.  Example: ODgxNDc2MDg2MDExODk5MAAAAAA&#x3D; | [optional] 

## Methods

### NewResponseAdditionalDataCommon

`func NewResponseAdditionalDataCommon() *ResponseAdditionalDataCommon`

NewResponseAdditionalDataCommon instantiates a new ResponseAdditionalDataCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAdditionalDataCommonWithDefaults

`func NewResponseAdditionalDataCommonWithDefaults() *ResponseAdditionalDataCommon`

NewResponseAdditionalDataCommonWithDefaults instantiates a new ResponseAdditionalDataCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcquirerAccountCode

`func (o *ResponseAdditionalDataCommon) GetAcquirerAccountCode() string`

GetAcquirerAccountCode returns the AcquirerAccountCode field if non-nil, zero value otherwise.

### GetAcquirerAccountCodeOk

`func (o *ResponseAdditionalDataCommon) GetAcquirerAccountCodeOk() (*string, bool)`

GetAcquirerAccountCodeOk returns a tuple with the AcquirerAccountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerAccountCode

`func (o *ResponseAdditionalDataCommon) SetAcquirerAccountCode(v string)`

SetAcquirerAccountCode sets AcquirerAccountCode field to given value.

### HasAcquirerAccountCode

`func (o *ResponseAdditionalDataCommon) HasAcquirerAccountCode() bool`

HasAcquirerAccountCode returns a boolean if a field has been set.

### GetAcquirerCode

`func (o *ResponseAdditionalDataCommon) GetAcquirerCode() string`

GetAcquirerCode returns the AcquirerCode field if non-nil, zero value otherwise.

### GetAcquirerCodeOk

`func (o *ResponseAdditionalDataCommon) GetAcquirerCodeOk() (*string, bool)`

GetAcquirerCodeOk returns a tuple with the AcquirerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerCode

`func (o *ResponseAdditionalDataCommon) SetAcquirerCode(v string)`

SetAcquirerCode sets AcquirerCode field to given value.

### HasAcquirerCode

`func (o *ResponseAdditionalDataCommon) HasAcquirerCode() bool`

HasAcquirerCode returns a boolean if a field has been set.

### GetAcquirerReference

`func (o *ResponseAdditionalDataCommon) GetAcquirerReference() string`

GetAcquirerReference returns the AcquirerReference field if non-nil, zero value otherwise.

### GetAcquirerReferenceOk

`func (o *ResponseAdditionalDataCommon) GetAcquirerReferenceOk() (*string, bool)`

GetAcquirerReferenceOk returns a tuple with the AcquirerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerReference

`func (o *ResponseAdditionalDataCommon) SetAcquirerReference(v string)`

SetAcquirerReference sets AcquirerReference field to given value.

### HasAcquirerReference

`func (o *ResponseAdditionalDataCommon) HasAcquirerReference() bool`

HasAcquirerReference returns a boolean if a field has been set.

### GetAlias

`func (o *ResponseAdditionalDataCommon) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ResponseAdditionalDataCommon) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ResponseAdditionalDataCommon) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ResponseAdditionalDataCommon) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetAliasType

`func (o *ResponseAdditionalDataCommon) GetAliasType() string`

GetAliasType returns the AliasType field if non-nil, zero value otherwise.

### GetAliasTypeOk

`func (o *ResponseAdditionalDataCommon) GetAliasTypeOk() (*string, bool)`

GetAliasTypeOk returns a tuple with the AliasType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasType

`func (o *ResponseAdditionalDataCommon) SetAliasType(v string)`

SetAliasType sets AliasType field to given value.

### HasAliasType

`func (o *ResponseAdditionalDataCommon) HasAliasType() bool`

HasAliasType returns a boolean if a field has been set.

### GetAuthCode

`func (o *ResponseAdditionalDataCommon) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *ResponseAdditionalDataCommon) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *ResponseAdditionalDataCommon) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *ResponseAdditionalDataCommon) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetAuthorisationMid

`func (o *ResponseAdditionalDataCommon) GetAuthorisationMid() string`

GetAuthorisationMid returns the AuthorisationMid field if non-nil, zero value otherwise.

### GetAuthorisationMidOk

`func (o *ResponseAdditionalDataCommon) GetAuthorisationMidOk() (*string, bool)`

GetAuthorisationMidOk returns a tuple with the AuthorisationMid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorisationMid

`func (o *ResponseAdditionalDataCommon) SetAuthorisationMid(v string)`

SetAuthorisationMid sets AuthorisationMid field to given value.

### HasAuthorisationMid

`func (o *ResponseAdditionalDataCommon) HasAuthorisationMid() bool`

HasAuthorisationMid returns a boolean if a field has been set.

### GetAuthorisedAmountCurrency

`func (o *ResponseAdditionalDataCommon) GetAuthorisedAmountCurrency() string`

GetAuthorisedAmountCurrency returns the AuthorisedAmountCurrency field if non-nil, zero value otherwise.

### GetAuthorisedAmountCurrencyOk

`func (o *ResponseAdditionalDataCommon) GetAuthorisedAmountCurrencyOk() (*string, bool)`

GetAuthorisedAmountCurrencyOk returns a tuple with the AuthorisedAmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorisedAmountCurrency

`func (o *ResponseAdditionalDataCommon) SetAuthorisedAmountCurrency(v string)`

SetAuthorisedAmountCurrency sets AuthorisedAmountCurrency field to given value.

### HasAuthorisedAmountCurrency

`func (o *ResponseAdditionalDataCommon) HasAuthorisedAmountCurrency() bool`

HasAuthorisedAmountCurrency returns a boolean if a field has been set.

### GetAuthorisedAmountValue

`func (o *ResponseAdditionalDataCommon) GetAuthorisedAmountValue() string`

GetAuthorisedAmountValue returns the AuthorisedAmountValue field if non-nil, zero value otherwise.

### GetAuthorisedAmountValueOk

`func (o *ResponseAdditionalDataCommon) GetAuthorisedAmountValueOk() (*string, bool)`

GetAuthorisedAmountValueOk returns a tuple with the AuthorisedAmountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorisedAmountValue

`func (o *ResponseAdditionalDataCommon) SetAuthorisedAmountValue(v string)`

SetAuthorisedAmountValue sets AuthorisedAmountValue field to given value.

### HasAuthorisedAmountValue

`func (o *ResponseAdditionalDataCommon) HasAuthorisedAmountValue() bool`

HasAuthorisedAmountValue returns a boolean if a field has been set.

### GetAvsResult

`func (o *ResponseAdditionalDataCommon) GetAvsResult() string`

GetAvsResult returns the AvsResult field if non-nil, zero value otherwise.

### GetAvsResultOk

`func (o *ResponseAdditionalDataCommon) GetAvsResultOk() (*string, bool)`

GetAvsResultOk returns a tuple with the AvsResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsResult

`func (o *ResponseAdditionalDataCommon) SetAvsResult(v string)`

SetAvsResult sets AvsResult field to given value.

### HasAvsResult

`func (o *ResponseAdditionalDataCommon) HasAvsResult() bool`

HasAvsResult returns a boolean if a field has been set.

### GetAvsResultRaw

`func (o *ResponseAdditionalDataCommon) GetAvsResultRaw() string`

GetAvsResultRaw returns the AvsResultRaw field if non-nil, zero value otherwise.

### GetAvsResultRawOk

`func (o *ResponseAdditionalDataCommon) GetAvsResultRawOk() (*string, bool)`

GetAvsResultRawOk returns a tuple with the AvsResultRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsResultRaw

`func (o *ResponseAdditionalDataCommon) SetAvsResultRaw(v string)`

SetAvsResultRaw sets AvsResultRaw field to given value.

### HasAvsResultRaw

`func (o *ResponseAdditionalDataCommon) HasAvsResultRaw() bool`

HasAvsResultRaw returns a boolean if a field has been set.

### GetBic

`func (o *ResponseAdditionalDataCommon) GetBic() string`

GetBic returns the Bic field if non-nil, zero value otherwise.

### GetBicOk

`func (o *ResponseAdditionalDataCommon) GetBicOk() (*string, bool)`

GetBicOk returns a tuple with the Bic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBic

`func (o *ResponseAdditionalDataCommon) SetBic(v string)`

SetBic sets Bic field to given value.

### HasBic

`func (o *ResponseAdditionalDataCommon) HasBic() bool`

HasBic returns a boolean if a field has been set.

### GetCoBrandedWith

`func (o *ResponseAdditionalDataCommon) GetCoBrandedWith() string`

GetCoBrandedWith returns the CoBrandedWith field if non-nil, zero value otherwise.

### GetCoBrandedWithOk

`func (o *ResponseAdditionalDataCommon) GetCoBrandedWithOk() (*string, bool)`

GetCoBrandedWithOk returns a tuple with the CoBrandedWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoBrandedWith

`func (o *ResponseAdditionalDataCommon) SetCoBrandedWith(v string)`

SetCoBrandedWith sets CoBrandedWith field to given value.

### HasCoBrandedWith

`func (o *ResponseAdditionalDataCommon) HasCoBrandedWith() bool`

HasCoBrandedWith returns a boolean if a field has been set.

### GetCvcResult

`func (o *ResponseAdditionalDataCommon) GetCvcResult() string`

GetCvcResult returns the CvcResult field if non-nil, zero value otherwise.

### GetCvcResultOk

`func (o *ResponseAdditionalDataCommon) GetCvcResultOk() (*string, bool)`

GetCvcResultOk returns a tuple with the CvcResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvcResult

`func (o *ResponseAdditionalDataCommon) SetCvcResult(v string)`

SetCvcResult sets CvcResult field to given value.

### HasCvcResult

`func (o *ResponseAdditionalDataCommon) HasCvcResult() bool`

HasCvcResult returns a boolean if a field has been set.

### GetCvcResultRaw

`func (o *ResponseAdditionalDataCommon) GetCvcResultRaw() string`

GetCvcResultRaw returns the CvcResultRaw field if non-nil, zero value otherwise.

### GetCvcResultRawOk

`func (o *ResponseAdditionalDataCommon) GetCvcResultRawOk() (*string, bool)`

GetCvcResultRawOk returns a tuple with the CvcResultRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvcResultRaw

`func (o *ResponseAdditionalDataCommon) SetCvcResultRaw(v string)`

SetCvcResultRaw sets CvcResultRaw field to given value.

### HasCvcResultRaw

`func (o *ResponseAdditionalDataCommon) HasCvcResultRaw() bool`

HasCvcResultRaw returns a boolean if a field has been set.

### GetDsTransID

`func (o *ResponseAdditionalDataCommon) GetDsTransID() string`

GetDsTransID returns the DsTransID field if non-nil, zero value otherwise.

### GetDsTransIDOk

`func (o *ResponseAdditionalDataCommon) GetDsTransIDOk() (*string, bool)`

GetDsTransIDOk returns a tuple with the DsTransID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsTransID

`func (o *ResponseAdditionalDataCommon) SetDsTransID(v string)`

SetDsTransID sets DsTransID field to given value.

### HasDsTransID

`func (o *ResponseAdditionalDataCommon) HasDsTransID() bool`

HasDsTransID returns a boolean if a field has been set.

### GetEci

`func (o *ResponseAdditionalDataCommon) GetEci() string`

GetEci returns the Eci field if non-nil, zero value otherwise.

### GetEciOk

`func (o *ResponseAdditionalDataCommon) GetEciOk() (*string, bool)`

GetEciOk returns a tuple with the Eci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEci

`func (o *ResponseAdditionalDataCommon) SetEci(v string)`

SetEci sets Eci field to given value.

### HasEci

`func (o *ResponseAdditionalDataCommon) HasEci() bool`

HasEci returns a boolean if a field has been set.

### GetExpiryDate

`func (o *ResponseAdditionalDataCommon) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *ResponseAdditionalDataCommon) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *ResponseAdditionalDataCommon) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *ResponseAdditionalDataCommon) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetExtraCostsCurrency

`func (o *ResponseAdditionalDataCommon) GetExtraCostsCurrency() string`

GetExtraCostsCurrency returns the ExtraCostsCurrency field if non-nil, zero value otherwise.

### GetExtraCostsCurrencyOk

`func (o *ResponseAdditionalDataCommon) GetExtraCostsCurrencyOk() (*string, bool)`

GetExtraCostsCurrencyOk returns a tuple with the ExtraCostsCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraCostsCurrency

`func (o *ResponseAdditionalDataCommon) SetExtraCostsCurrency(v string)`

SetExtraCostsCurrency sets ExtraCostsCurrency field to given value.

### HasExtraCostsCurrency

`func (o *ResponseAdditionalDataCommon) HasExtraCostsCurrency() bool`

HasExtraCostsCurrency returns a boolean if a field has been set.

### GetExtraCostsValue

`func (o *ResponseAdditionalDataCommon) GetExtraCostsValue() string`

GetExtraCostsValue returns the ExtraCostsValue field if non-nil, zero value otherwise.

### GetExtraCostsValueOk

`func (o *ResponseAdditionalDataCommon) GetExtraCostsValueOk() (*string, bool)`

GetExtraCostsValueOk returns a tuple with the ExtraCostsValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraCostsValue

`func (o *ResponseAdditionalDataCommon) SetExtraCostsValue(v string)`

SetExtraCostsValue sets ExtraCostsValue field to given value.

### HasExtraCostsValue

`func (o *ResponseAdditionalDataCommon) HasExtraCostsValue() bool`

HasExtraCostsValue returns a boolean if a field has been set.

### GetFraudCheckItemNrFraudCheckname

`func (o *ResponseAdditionalDataCommon) GetFraudCheckItemNrFraudCheckname() string`

GetFraudCheckItemNrFraudCheckname returns the FraudCheckItemNrFraudCheckname field if non-nil, zero value otherwise.

### GetFraudCheckItemNrFraudChecknameOk

`func (o *ResponseAdditionalDataCommon) GetFraudCheckItemNrFraudChecknameOk() (*string, bool)`

GetFraudCheckItemNrFraudChecknameOk returns a tuple with the FraudCheckItemNrFraudCheckname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudCheckItemNrFraudCheckname

`func (o *ResponseAdditionalDataCommon) SetFraudCheckItemNrFraudCheckname(v string)`

SetFraudCheckItemNrFraudCheckname sets FraudCheckItemNrFraudCheckname field to given value.

### HasFraudCheckItemNrFraudCheckname

`func (o *ResponseAdditionalDataCommon) HasFraudCheckItemNrFraudCheckname() bool`

HasFraudCheckItemNrFraudCheckname returns a boolean if a field has been set.

### GetFraudManualReview

`func (o *ResponseAdditionalDataCommon) GetFraudManualReview() string`

GetFraudManualReview returns the FraudManualReview field if non-nil, zero value otherwise.

### GetFraudManualReviewOk

`func (o *ResponseAdditionalDataCommon) GetFraudManualReviewOk() (*string, bool)`

GetFraudManualReviewOk returns a tuple with the FraudManualReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudManualReview

`func (o *ResponseAdditionalDataCommon) SetFraudManualReview(v string)`

SetFraudManualReview sets FraudManualReview field to given value.

### HasFraudManualReview

`func (o *ResponseAdditionalDataCommon) HasFraudManualReview() bool`

HasFraudManualReview returns a boolean if a field has been set.

### GetFraudResultType

`func (o *ResponseAdditionalDataCommon) GetFraudResultType() string`

GetFraudResultType returns the FraudResultType field if non-nil, zero value otherwise.

### GetFraudResultTypeOk

`func (o *ResponseAdditionalDataCommon) GetFraudResultTypeOk() (*string, bool)`

GetFraudResultTypeOk returns a tuple with the FraudResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudResultType

`func (o *ResponseAdditionalDataCommon) SetFraudResultType(v string)`

SetFraudResultType sets FraudResultType field to given value.

### HasFraudResultType

`func (o *ResponseAdditionalDataCommon) HasFraudResultType() bool`

HasFraudResultType returns a boolean if a field has been set.

### GetFundingSource

`func (o *ResponseAdditionalDataCommon) GetFundingSource() string`

GetFundingSource returns the FundingSource field if non-nil, zero value otherwise.

### GetFundingSourceOk

`func (o *ResponseAdditionalDataCommon) GetFundingSourceOk() (*string, bool)`

GetFundingSourceOk returns a tuple with the FundingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSource

`func (o *ResponseAdditionalDataCommon) SetFundingSource(v string)`

SetFundingSource sets FundingSource field to given value.

### HasFundingSource

`func (o *ResponseAdditionalDataCommon) HasFundingSource() bool`

HasFundingSource returns a boolean if a field has been set.

### GetFundsAvailability

`func (o *ResponseAdditionalDataCommon) GetFundsAvailability() string`

GetFundsAvailability returns the FundsAvailability field if non-nil, zero value otherwise.

### GetFundsAvailabilityOk

`func (o *ResponseAdditionalDataCommon) GetFundsAvailabilityOk() (*string, bool)`

GetFundsAvailabilityOk returns a tuple with the FundsAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundsAvailability

`func (o *ResponseAdditionalDataCommon) SetFundsAvailability(v string)`

SetFundsAvailability sets FundsAvailability field to given value.

### HasFundsAvailability

`func (o *ResponseAdditionalDataCommon) HasFundsAvailability() bool`

HasFundsAvailability returns a boolean if a field has been set.

### GetInferredRefusalReason

`func (o *ResponseAdditionalDataCommon) GetInferredRefusalReason() string`

GetInferredRefusalReason returns the InferredRefusalReason field if non-nil, zero value otherwise.

### GetInferredRefusalReasonOk

`func (o *ResponseAdditionalDataCommon) GetInferredRefusalReasonOk() (*string, bool)`

GetInferredRefusalReasonOk returns a tuple with the InferredRefusalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferredRefusalReason

`func (o *ResponseAdditionalDataCommon) SetInferredRefusalReason(v string)`

SetInferredRefusalReason sets InferredRefusalReason field to given value.

### HasInferredRefusalReason

`func (o *ResponseAdditionalDataCommon) HasInferredRefusalReason() bool`

HasInferredRefusalReason returns a boolean if a field has been set.

### GetIsCardCommercial

`func (o *ResponseAdditionalDataCommon) GetIsCardCommercial() string`

GetIsCardCommercial returns the IsCardCommercial field if non-nil, zero value otherwise.

### GetIsCardCommercialOk

`func (o *ResponseAdditionalDataCommon) GetIsCardCommercialOk() (*string, bool)`

GetIsCardCommercialOk returns a tuple with the IsCardCommercial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCardCommercial

`func (o *ResponseAdditionalDataCommon) SetIsCardCommercial(v string)`

SetIsCardCommercial sets IsCardCommercial field to given value.

### HasIsCardCommercial

`func (o *ResponseAdditionalDataCommon) HasIsCardCommercial() bool`

HasIsCardCommercial returns a boolean if a field has been set.

### GetIssuerCountry

`func (o *ResponseAdditionalDataCommon) GetIssuerCountry() string`

GetIssuerCountry returns the IssuerCountry field if non-nil, zero value otherwise.

### GetIssuerCountryOk

`func (o *ResponseAdditionalDataCommon) GetIssuerCountryOk() (*string, bool)`

GetIssuerCountryOk returns a tuple with the IssuerCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCountry

`func (o *ResponseAdditionalDataCommon) SetIssuerCountry(v string)`

SetIssuerCountry sets IssuerCountry field to given value.

### HasIssuerCountry

`func (o *ResponseAdditionalDataCommon) HasIssuerCountry() bool`

HasIssuerCountry returns a boolean if a field has been set.

### GetLiabilityShift

`func (o *ResponseAdditionalDataCommon) GetLiabilityShift() string`

GetLiabilityShift returns the LiabilityShift field if non-nil, zero value otherwise.

### GetLiabilityShiftOk

`func (o *ResponseAdditionalDataCommon) GetLiabilityShiftOk() (*string, bool)`

GetLiabilityShiftOk returns a tuple with the LiabilityShift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiabilityShift

`func (o *ResponseAdditionalDataCommon) SetLiabilityShift(v string)`

SetLiabilityShift sets LiabilityShift field to given value.

### HasLiabilityShift

`func (o *ResponseAdditionalDataCommon) HasLiabilityShift() bool`

HasLiabilityShift returns a boolean if a field has been set.

### GetMcBankNetReferenceNumber

`func (o *ResponseAdditionalDataCommon) GetMcBankNetReferenceNumber() string`

GetMcBankNetReferenceNumber returns the McBankNetReferenceNumber field if non-nil, zero value otherwise.

### GetMcBankNetReferenceNumberOk

`func (o *ResponseAdditionalDataCommon) GetMcBankNetReferenceNumberOk() (*string, bool)`

GetMcBankNetReferenceNumberOk returns a tuple with the McBankNetReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcBankNetReferenceNumber

`func (o *ResponseAdditionalDataCommon) SetMcBankNetReferenceNumber(v string)`

SetMcBankNetReferenceNumber sets McBankNetReferenceNumber field to given value.

### HasMcBankNetReferenceNumber

`func (o *ResponseAdditionalDataCommon) HasMcBankNetReferenceNumber() bool`

HasMcBankNetReferenceNumber returns a boolean if a field has been set.

### GetMerchantAdviceCode

`func (o *ResponseAdditionalDataCommon) GetMerchantAdviceCode() string`

GetMerchantAdviceCode returns the MerchantAdviceCode field if non-nil, zero value otherwise.

### GetMerchantAdviceCodeOk

`func (o *ResponseAdditionalDataCommon) GetMerchantAdviceCodeOk() (*string, bool)`

GetMerchantAdviceCodeOk returns a tuple with the MerchantAdviceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAdviceCode

`func (o *ResponseAdditionalDataCommon) SetMerchantAdviceCode(v string)`

SetMerchantAdviceCode sets MerchantAdviceCode field to given value.

### HasMerchantAdviceCode

`func (o *ResponseAdditionalDataCommon) HasMerchantAdviceCode() bool`

HasMerchantAdviceCode returns a boolean if a field has been set.

### GetMerchantReference

`func (o *ResponseAdditionalDataCommon) GetMerchantReference() string`

GetMerchantReference returns the MerchantReference field if non-nil, zero value otherwise.

### GetMerchantReferenceOk

`func (o *ResponseAdditionalDataCommon) GetMerchantReferenceOk() (*string, bool)`

GetMerchantReferenceOk returns a tuple with the MerchantReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantReference

`func (o *ResponseAdditionalDataCommon) SetMerchantReference(v string)`

SetMerchantReference sets MerchantReference field to given value.

### HasMerchantReference

`func (o *ResponseAdditionalDataCommon) HasMerchantReference() bool`

HasMerchantReference returns a boolean if a field has been set.

### GetNetworkTxReference

`func (o *ResponseAdditionalDataCommon) GetNetworkTxReference() string`

GetNetworkTxReference returns the NetworkTxReference field if non-nil, zero value otherwise.

### GetNetworkTxReferenceOk

`func (o *ResponseAdditionalDataCommon) GetNetworkTxReferenceOk() (*string, bool)`

GetNetworkTxReferenceOk returns a tuple with the NetworkTxReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTxReference

`func (o *ResponseAdditionalDataCommon) SetNetworkTxReference(v string)`

SetNetworkTxReference sets NetworkTxReference field to given value.

### HasNetworkTxReference

`func (o *ResponseAdditionalDataCommon) HasNetworkTxReference() bool`

HasNetworkTxReference returns a boolean if a field has been set.

### GetOwnerName

`func (o *ResponseAdditionalDataCommon) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *ResponseAdditionalDataCommon) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *ResponseAdditionalDataCommon) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *ResponseAdditionalDataCommon) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### GetPaymentAccountReference

`func (o *ResponseAdditionalDataCommon) GetPaymentAccountReference() string`

GetPaymentAccountReference returns the PaymentAccountReference field if non-nil, zero value otherwise.

### GetPaymentAccountReferenceOk

`func (o *ResponseAdditionalDataCommon) GetPaymentAccountReferenceOk() (*string, bool)`

GetPaymentAccountReferenceOk returns a tuple with the PaymentAccountReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccountReference

`func (o *ResponseAdditionalDataCommon) SetPaymentAccountReference(v string)`

SetPaymentAccountReference sets PaymentAccountReference field to given value.

### HasPaymentAccountReference

`func (o *ResponseAdditionalDataCommon) HasPaymentAccountReference() bool`

HasPaymentAccountReference returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *ResponseAdditionalDataCommon) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *ResponseAdditionalDataCommon) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *ResponseAdditionalDataCommon) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *ResponseAdditionalDataCommon) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPaymentMethodVariant

`func (o *ResponseAdditionalDataCommon) GetPaymentMethodVariant() string`

GetPaymentMethodVariant returns the PaymentMethodVariant field if non-nil, zero value otherwise.

### GetPaymentMethodVariantOk

`func (o *ResponseAdditionalDataCommon) GetPaymentMethodVariantOk() (*string, bool)`

GetPaymentMethodVariantOk returns a tuple with the PaymentMethodVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodVariant

`func (o *ResponseAdditionalDataCommon) SetPaymentMethodVariant(v string)`

SetPaymentMethodVariant sets PaymentMethodVariant field to given value.

### HasPaymentMethodVariant

`func (o *ResponseAdditionalDataCommon) HasPaymentMethodVariant() bool`

HasPaymentMethodVariant returns a boolean if a field has been set.

### GetPayoutEligible

`func (o *ResponseAdditionalDataCommon) GetPayoutEligible() string`

GetPayoutEligible returns the PayoutEligible field if non-nil, zero value otherwise.

### GetPayoutEligibleOk

`func (o *ResponseAdditionalDataCommon) GetPayoutEligibleOk() (*string, bool)`

GetPayoutEligibleOk returns a tuple with the PayoutEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutEligible

`func (o *ResponseAdditionalDataCommon) SetPayoutEligible(v string)`

SetPayoutEligible sets PayoutEligible field to given value.

### HasPayoutEligible

`func (o *ResponseAdditionalDataCommon) HasPayoutEligible() bool`

HasPayoutEligible returns a boolean if a field has been set.

### GetRealtimeAccountUpdaterStatus

`func (o *ResponseAdditionalDataCommon) GetRealtimeAccountUpdaterStatus() string`

GetRealtimeAccountUpdaterStatus returns the RealtimeAccountUpdaterStatus field if non-nil, zero value otherwise.

### GetRealtimeAccountUpdaterStatusOk

`func (o *ResponseAdditionalDataCommon) GetRealtimeAccountUpdaterStatusOk() (*string, bool)`

GetRealtimeAccountUpdaterStatusOk returns a tuple with the RealtimeAccountUpdaterStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealtimeAccountUpdaterStatus

`func (o *ResponseAdditionalDataCommon) SetRealtimeAccountUpdaterStatus(v string)`

SetRealtimeAccountUpdaterStatus sets RealtimeAccountUpdaterStatus field to given value.

### HasRealtimeAccountUpdaterStatus

`func (o *ResponseAdditionalDataCommon) HasRealtimeAccountUpdaterStatus() bool`

HasRealtimeAccountUpdaterStatus returns a boolean if a field has been set.

### GetReceiptFreeText

`func (o *ResponseAdditionalDataCommon) GetReceiptFreeText() string`

GetReceiptFreeText returns the ReceiptFreeText field if non-nil, zero value otherwise.

### GetReceiptFreeTextOk

`func (o *ResponseAdditionalDataCommon) GetReceiptFreeTextOk() (*string, bool)`

GetReceiptFreeTextOk returns a tuple with the ReceiptFreeText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptFreeText

`func (o *ResponseAdditionalDataCommon) SetReceiptFreeText(v string)`

SetReceiptFreeText sets ReceiptFreeText field to given value.

### HasReceiptFreeText

`func (o *ResponseAdditionalDataCommon) HasReceiptFreeText() bool`

HasReceiptFreeText returns a boolean if a field has been set.

### GetRecurringContractTypes

`func (o *ResponseAdditionalDataCommon) GetRecurringContractTypes() string`

GetRecurringContractTypes returns the RecurringContractTypes field if non-nil, zero value otherwise.

### GetRecurringContractTypesOk

`func (o *ResponseAdditionalDataCommon) GetRecurringContractTypesOk() (*string, bool)`

GetRecurringContractTypesOk returns a tuple with the RecurringContractTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringContractTypes

`func (o *ResponseAdditionalDataCommon) SetRecurringContractTypes(v string)`

SetRecurringContractTypes sets RecurringContractTypes field to given value.

### HasRecurringContractTypes

`func (o *ResponseAdditionalDataCommon) HasRecurringContractTypes() bool`

HasRecurringContractTypes returns a boolean if a field has been set.

### GetRecurringFirstPspReference

`func (o *ResponseAdditionalDataCommon) GetRecurringFirstPspReference() string`

GetRecurringFirstPspReference returns the RecurringFirstPspReference field if non-nil, zero value otherwise.

### GetRecurringFirstPspReferenceOk

`func (o *ResponseAdditionalDataCommon) GetRecurringFirstPspReferenceOk() (*string, bool)`

GetRecurringFirstPspReferenceOk returns a tuple with the RecurringFirstPspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringFirstPspReference

`func (o *ResponseAdditionalDataCommon) SetRecurringFirstPspReference(v string)`

SetRecurringFirstPspReference sets RecurringFirstPspReference field to given value.

### HasRecurringFirstPspReference

`func (o *ResponseAdditionalDataCommon) HasRecurringFirstPspReference() bool`

HasRecurringFirstPspReference returns a boolean if a field has been set.

### GetRecurringRecurringDetailReference

`func (o *ResponseAdditionalDataCommon) GetRecurringRecurringDetailReference() string`

GetRecurringRecurringDetailReference returns the RecurringRecurringDetailReference field if non-nil, zero value otherwise.

### GetRecurringRecurringDetailReferenceOk

`func (o *ResponseAdditionalDataCommon) GetRecurringRecurringDetailReferenceOk() (*string, bool)`

GetRecurringRecurringDetailReferenceOk returns a tuple with the RecurringRecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringRecurringDetailReference

`func (o *ResponseAdditionalDataCommon) SetRecurringRecurringDetailReference(v string)`

SetRecurringRecurringDetailReference sets RecurringRecurringDetailReference field to given value.

### HasRecurringRecurringDetailReference

`func (o *ResponseAdditionalDataCommon) HasRecurringRecurringDetailReference() bool`

HasRecurringRecurringDetailReference returns a boolean if a field has been set.

### GetRecurringShopperReference

`func (o *ResponseAdditionalDataCommon) GetRecurringShopperReference() string`

GetRecurringShopperReference returns the RecurringShopperReference field if non-nil, zero value otherwise.

### GetRecurringShopperReferenceOk

`func (o *ResponseAdditionalDataCommon) GetRecurringShopperReferenceOk() (*string, bool)`

GetRecurringShopperReferenceOk returns a tuple with the RecurringShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringShopperReference

`func (o *ResponseAdditionalDataCommon) SetRecurringShopperReference(v string)`

SetRecurringShopperReference sets RecurringShopperReference field to given value.

### HasRecurringShopperReference

`func (o *ResponseAdditionalDataCommon) HasRecurringShopperReference() bool`

HasRecurringShopperReference returns a boolean if a field has been set.

### GetRecurringProcessingModel

`func (o *ResponseAdditionalDataCommon) GetRecurringProcessingModel() string`

GetRecurringProcessingModel returns the RecurringProcessingModel field if non-nil, zero value otherwise.

### GetRecurringProcessingModelOk

`func (o *ResponseAdditionalDataCommon) GetRecurringProcessingModelOk() (*string, bool)`

GetRecurringProcessingModelOk returns a tuple with the RecurringProcessingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringProcessingModel

`func (o *ResponseAdditionalDataCommon) SetRecurringProcessingModel(v string)`

SetRecurringProcessingModel sets RecurringProcessingModel field to given value.

### HasRecurringProcessingModel

`func (o *ResponseAdditionalDataCommon) HasRecurringProcessingModel() bool`

HasRecurringProcessingModel returns a boolean if a field has been set.

### GetReferred

`func (o *ResponseAdditionalDataCommon) GetReferred() string`

GetReferred returns the Referred field if non-nil, zero value otherwise.

### GetReferredOk

`func (o *ResponseAdditionalDataCommon) GetReferredOk() (*string, bool)`

GetReferredOk returns a tuple with the Referred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferred

`func (o *ResponseAdditionalDataCommon) SetReferred(v string)`

SetReferred sets Referred field to given value.

### HasReferred

`func (o *ResponseAdditionalDataCommon) HasReferred() bool`

HasReferred returns a boolean if a field has been set.

### GetRefusalReasonRaw

`func (o *ResponseAdditionalDataCommon) GetRefusalReasonRaw() string`

GetRefusalReasonRaw returns the RefusalReasonRaw field if non-nil, zero value otherwise.

### GetRefusalReasonRawOk

`func (o *ResponseAdditionalDataCommon) GetRefusalReasonRawOk() (*string, bool)`

GetRefusalReasonRawOk returns a tuple with the RefusalReasonRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusalReasonRaw

`func (o *ResponseAdditionalDataCommon) SetRefusalReasonRaw(v string)`

SetRefusalReasonRaw sets RefusalReasonRaw field to given value.

### HasRefusalReasonRaw

`func (o *ResponseAdditionalDataCommon) HasRefusalReasonRaw() bool`

HasRefusalReasonRaw returns a boolean if a field has been set.

### GetRequestAmount

`func (o *ResponseAdditionalDataCommon) GetRequestAmount() string`

GetRequestAmount returns the RequestAmount field if non-nil, zero value otherwise.

### GetRequestAmountOk

`func (o *ResponseAdditionalDataCommon) GetRequestAmountOk() (*string, bool)`

GetRequestAmountOk returns a tuple with the RequestAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestAmount

`func (o *ResponseAdditionalDataCommon) SetRequestAmount(v string)`

SetRequestAmount sets RequestAmount field to given value.

### HasRequestAmount

`func (o *ResponseAdditionalDataCommon) HasRequestAmount() bool`

HasRequestAmount returns a boolean if a field has been set.

### GetRequestCurrencyCode

`func (o *ResponseAdditionalDataCommon) GetRequestCurrencyCode() string`

GetRequestCurrencyCode returns the RequestCurrencyCode field if non-nil, zero value otherwise.

### GetRequestCurrencyCodeOk

`func (o *ResponseAdditionalDataCommon) GetRequestCurrencyCodeOk() (*string, bool)`

GetRequestCurrencyCodeOk returns a tuple with the RequestCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCurrencyCode

`func (o *ResponseAdditionalDataCommon) SetRequestCurrencyCode(v string)`

SetRequestCurrencyCode sets RequestCurrencyCode field to given value.

### HasRequestCurrencyCode

`func (o *ResponseAdditionalDataCommon) HasRequestCurrencyCode() bool`

HasRequestCurrencyCode returns a boolean if a field has been set.

### GetShopperInteraction

`func (o *ResponseAdditionalDataCommon) GetShopperInteraction() string`

GetShopperInteraction returns the ShopperInteraction field if non-nil, zero value otherwise.

### GetShopperInteractionOk

`func (o *ResponseAdditionalDataCommon) GetShopperInteractionOk() (*string, bool)`

GetShopperInteractionOk returns a tuple with the ShopperInteraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperInteraction

`func (o *ResponseAdditionalDataCommon) SetShopperInteraction(v string)`

SetShopperInteraction sets ShopperInteraction field to given value.

### HasShopperInteraction

`func (o *ResponseAdditionalDataCommon) HasShopperInteraction() bool`

HasShopperInteraction returns a boolean if a field has been set.

### GetShopperReference

`func (o *ResponseAdditionalDataCommon) GetShopperReference() string`

GetShopperReference returns the ShopperReference field if non-nil, zero value otherwise.

### GetShopperReferenceOk

`func (o *ResponseAdditionalDataCommon) GetShopperReferenceOk() (*string, bool)`

GetShopperReferenceOk returns a tuple with the ShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperReference

`func (o *ResponseAdditionalDataCommon) SetShopperReference(v string)`

SetShopperReference sets ShopperReference field to given value.

### HasShopperReference

`func (o *ResponseAdditionalDataCommon) HasShopperReference() bool`

HasShopperReference returns a boolean if a field has been set.

### GetTerminalId

`func (o *ResponseAdditionalDataCommon) GetTerminalId() string`

GetTerminalId returns the TerminalId field if non-nil, zero value otherwise.

### GetTerminalIdOk

`func (o *ResponseAdditionalDataCommon) GetTerminalIdOk() (*string, bool)`

GetTerminalIdOk returns a tuple with the TerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalId

`func (o *ResponseAdditionalDataCommon) SetTerminalId(v string)`

SetTerminalId sets TerminalId field to given value.

### HasTerminalId

`func (o *ResponseAdditionalDataCommon) HasTerminalId() bool`

HasTerminalId returns a boolean if a field has been set.

### GetThreeDAuthenticated

`func (o *ResponseAdditionalDataCommon) GetThreeDAuthenticated() string`

GetThreeDAuthenticated returns the ThreeDAuthenticated field if non-nil, zero value otherwise.

### GetThreeDAuthenticatedOk

`func (o *ResponseAdditionalDataCommon) GetThreeDAuthenticatedOk() (*string, bool)`

GetThreeDAuthenticatedOk returns a tuple with the ThreeDAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDAuthenticated

`func (o *ResponseAdditionalDataCommon) SetThreeDAuthenticated(v string)`

SetThreeDAuthenticated sets ThreeDAuthenticated field to given value.

### HasThreeDAuthenticated

`func (o *ResponseAdditionalDataCommon) HasThreeDAuthenticated() bool`

HasThreeDAuthenticated returns a boolean if a field has been set.

### GetThreeDAuthenticatedResponse

`func (o *ResponseAdditionalDataCommon) GetThreeDAuthenticatedResponse() string`

GetThreeDAuthenticatedResponse returns the ThreeDAuthenticatedResponse field if non-nil, zero value otherwise.

### GetThreeDAuthenticatedResponseOk

`func (o *ResponseAdditionalDataCommon) GetThreeDAuthenticatedResponseOk() (*string, bool)`

GetThreeDAuthenticatedResponseOk returns a tuple with the ThreeDAuthenticatedResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDAuthenticatedResponse

`func (o *ResponseAdditionalDataCommon) SetThreeDAuthenticatedResponse(v string)`

SetThreeDAuthenticatedResponse sets ThreeDAuthenticatedResponse field to given value.

### HasThreeDAuthenticatedResponse

`func (o *ResponseAdditionalDataCommon) HasThreeDAuthenticatedResponse() bool`

HasThreeDAuthenticatedResponse returns a boolean if a field has been set.

### GetThreeDOffered

`func (o *ResponseAdditionalDataCommon) GetThreeDOffered() string`

GetThreeDOffered returns the ThreeDOffered field if non-nil, zero value otherwise.

### GetThreeDOfferedOk

`func (o *ResponseAdditionalDataCommon) GetThreeDOfferedOk() (*string, bool)`

GetThreeDOfferedOk returns a tuple with the ThreeDOffered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDOffered

`func (o *ResponseAdditionalDataCommon) SetThreeDOffered(v string)`

SetThreeDOffered sets ThreeDOffered field to given value.

### HasThreeDOffered

`func (o *ResponseAdditionalDataCommon) HasThreeDOffered() bool`

HasThreeDOffered returns a boolean if a field has been set.

### GetThreeDOfferedResponse

`func (o *ResponseAdditionalDataCommon) GetThreeDOfferedResponse() string`

GetThreeDOfferedResponse returns the ThreeDOfferedResponse field if non-nil, zero value otherwise.

### GetThreeDOfferedResponseOk

`func (o *ResponseAdditionalDataCommon) GetThreeDOfferedResponseOk() (*string, bool)`

GetThreeDOfferedResponseOk returns a tuple with the ThreeDOfferedResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDOfferedResponse

`func (o *ResponseAdditionalDataCommon) SetThreeDOfferedResponse(v string)`

SetThreeDOfferedResponse sets ThreeDOfferedResponse field to given value.

### HasThreeDOfferedResponse

`func (o *ResponseAdditionalDataCommon) HasThreeDOfferedResponse() bool`

HasThreeDOfferedResponse returns a boolean if a field has been set.

### GetThreeDSVersion

`func (o *ResponseAdditionalDataCommon) GetThreeDSVersion() string`

GetThreeDSVersion returns the ThreeDSVersion field if non-nil, zero value otherwise.

### GetThreeDSVersionOk

`func (o *ResponseAdditionalDataCommon) GetThreeDSVersionOk() (*string, bool)`

GetThreeDSVersionOk returns a tuple with the ThreeDSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSVersion

`func (o *ResponseAdditionalDataCommon) SetThreeDSVersion(v string)`

SetThreeDSVersion sets ThreeDSVersion field to given value.

### HasThreeDSVersion

`func (o *ResponseAdditionalDataCommon) HasThreeDSVersion() bool`

HasThreeDSVersion returns a boolean if a field has been set.

### GetVisaTransactionId

`func (o *ResponseAdditionalDataCommon) GetVisaTransactionId() string`

GetVisaTransactionId returns the VisaTransactionId field if non-nil, zero value otherwise.

### GetVisaTransactionIdOk

`func (o *ResponseAdditionalDataCommon) GetVisaTransactionIdOk() (*string, bool)`

GetVisaTransactionIdOk returns a tuple with the VisaTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisaTransactionId

`func (o *ResponseAdditionalDataCommon) SetVisaTransactionId(v string)`

SetVisaTransactionId sets VisaTransactionId field to given value.

### HasVisaTransactionId

`func (o *ResponseAdditionalDataCommon) HasVisaTransactionId() bool`

HasVisaTransactionId returns a boolean if a field has been set.

### GetXid

`func (o *ResponseAdditionalDataCommon) GetXid() string`

GetXid returns the Xid field if non-nil, zero value otherwise.

### GetXidOk

`func (o *ResponseAdditionalDataCommon) GetXidOk() (*string, bool)`

GetXidOk returns a tuple with the Xid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXid

`func (o *ResponseAdditionalDataCommon) SetXid(v string)`

SetXid sets Xid field to given value.

### HasXid

`func (o *ResponseAdditionalDataCommon) HasXid() bool`

HasXid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


