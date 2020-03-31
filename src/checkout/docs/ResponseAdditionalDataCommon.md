# ResponseAdditionalDataCommon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcquirerAccountCode** | **string** | The name of the Adyen acquirer account.  Example: PayPalSandbox_TestAcquirer &gt; Only relevant for PayPal transactions. | [optional] 
**AcquirerCode** | **string** | The name of the acquirer processing the payment request.  Example: TestPmmAcquirer | [optional] 
**AcquirerReference** | **string** | The reference number that can be used for reconciliation in case a non-Adyen acquirer is used for settlement.  Example: 7C9N3FNBKT9 | [optional] 
**Alias** | **string** | The Adyen alias of the card.  Example: H167852639363479 | [optional] 
**AliasType** | **string** | The type of the card alias.  Example: Default | [optional] 
**AuthCode** | **string** | Authorisation code: * When the payment is authorised successfully, this field holds the authorisation code for the payment. * When the payment is not authorised, this field is empty.  Example: 58747 | [optional] 
**AuthorisedAmountCurrency** | **string** | The currency of the authorised amount, as a three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes). | [optional] 
**AuthorisedAmountValue** | **string** | Value of the amount authorised.  This amount is represented in minor units according to the [following table](https://docs.adyen.com/development-resources/currency-codes). | [optional] 
**AvsResult** | **string** | The AVS result code of the payment, which provides information about the outcome of the AVS check.  For possible values, see [AVS](https://docs.adyen.com/risk-management/configure-standard-risk-rules/consistency-rules#billing-address-does-not-match-cardholder-address-avs). | [optional] 
**AvsResultRaw** | **string** | Raw AVS result received from the acquirer, where available.  Example: D | [optional] 
**Bic** | **string** | BIC of a bank account.  Example: TESTNL01 &gt; Only relevant for SEPA Direct Debit transactions. | [optional] 
**DsTransID** | **string** | Supported for 3D Secure 2. The unique transaction identifier assigned by the DS to identify a single transaction. | [optional] 
**Eci** | **string** | The Electronic Commerce Indicator returned from the schemes for the 3DS payment session.  Example: 02 | [optional] 
**ExpiryDate** | **string** | The expiry date on the card.  Example: 6/2016 &gt; Returned only in case of a card payment. | [optional] 
**ExtraCostsCurrency** | **string** | The currency of the extra amount charged due to additional amounts set in the skin used in the HPP payment request.  Example: EUR | [optional] 
**ExtraCostsValue** | **string** | The value of the extra amount charged due to additional amounts set in the skin used in the HPP payment request. The amount is in minor units. | [optional] 
**FraudCheckItemNrFraudCheckname** | **string** | The fraud score due to a particular fraud check. The fraud check name is found in the key of the key-value pair. | [optional] 
**FundingSource** | **string** | Information regarding the funding type of the card. The possible return values are: * CHARGE * CREDIT * DEBIT * PREPAID * PREPAID_RELOADABLE * PREPAID_NONRELOADABLE * DEFFERED_DEBIT &gt; This functionality requires additional configuration on Adyen&#39;s end. To enable it, contact the Support Team.  For receiving this field in the notification, enable **Include Funding Source** in **Notifications** &gt; **Additional settings**. | [optional] 
**FundsAvailability** | **string** | Indicates availability of funds.  Visa: * \&quot;I\&quot; (fast funds are supported) * \&quot;N\&quot; (otherwise)  Mastercard: * \&quot;I\&quot; (product type is Prepaid or Debit, or issuing country is in CEE/HGEM list) * \&quot;N\&quot; (otherwise)  &gt; Returned when you verify a card BIN or estimate costs, and only if payoutEligible is \&quot;Y\&quot; or \&quot;D\&quot;. | [optional] 
**InferredRefusalReason** | **string** | Provides the more granular indication of why a transaction was refused. When a transaction fails with either \&quot;Refused\&quot;, \&quot;Restricted Card\&quot;, \&quot;Transaction Not Permitted\&quot;, \&quot;Not supported\&quot; or \&quot;DeclinedNon Generic\&quot; refusalReason from the issuer, Adyen cross references its PSP-wide data for extra insight into the refusal reason. If an inferred refusal reason is available, the &#x60;inferredRefusalReason&#x60;, field is populated and the &#x60;refusalReason&#x60;, is set to \&quot;Not Supported\&quot;.  Possible values: * 3D Secure Mandated * Closed Account * ContAuth Not Supported * CVC Mandated * Ecommerce Not Allowed * Crossborder Not Supported * Card Updated * Low Authrate Bin * Non-reloadable prepaid card | [optional] 
**IssuerCountry** | **string** | The issuing country of the card based on the BIN list that Adyen maintains.  Example: JP | [optional] 
**McBankNetReferenceNumber** | **string** | The &#x60;mcBankNetReferenceNumber&#x60;, is a minimum of six characters and a maximum of nine characters long. &gt; Contact Support Team to enable this field. | [optional] 
**NetworkTxReference** | **string** | Returned in the response if you are not tokenizing with Adyen and are using the Merchant-initiated transactions (MIT) framework from Mastercard or Visa.  This contains either the Mastercard Trace ID or the Visa Transaction ID. | [optional] 
**OwnerName** | **string** | The owner name of a bank account.  Only relevant for SEPA Direct Debit transactions. | [optional] 
**PaymentAccountReference** | **string** | The Payment Account Reference (PAR) value links a network token with the underlying primary account number (PAN). The PAR value consists of 29 uppercase alphanumeric characters. | [optional] 
**PaymentMethodVariant** | **string** | The Adyen sub-variant of the payment method used for the payment request.  For more information, refer to [PaymentMethodVariant](https://docs.adyen.com/development-resources/paymentmethodvariant).  Example: mcpro | [optional] 
**PayoutEligible** | **string** | Indicates whether a payout is eligible or not for this card.  Visa: * \&quot;Y\&quot; * \&quot;N\&quot;  Mastercard: * \&quot;Y\&quot; (domestic and cross-border) * \&quot;D\&quot; (only domestic) * \&quot;N\&quot; (no MoneySend) * \&quot;U\&quot; (unknown) | [optional] 
**RealtimeAccountUpdaterStatus** | **string** | The response code from the Real Time Account Updater service.  Possible return values are: * CardChanged * CardExpiryChanged * CloseAccount * ContactCardAccountHolder | [optional] 
**ReceiptFreeText** | **string** | Message to be displayed on the terminal. | [optional] 
**RecurringFirstPspReference** | **string** | The &#x60;pspReference&#x60;, of the first recurring payment that created the recurring detail.  This functionality requires additional configuration on Adyen&#39;s end. To enable it, contact the Support Team. | [optional] 
**RecurringRecurringDetailReference** | **string** | The reference that uniquely identifies the recurring transaction. | [optional] 
**Referred** | **string** | If the payment is referred, this field is set to true.  This field is unavailable if the payment is referred and is usually not returned with ecommerce transactions.  Example: true | [optional] 
**RefusalReasonRaw** | **string** | Raw refusal reason received from the acquirer, where available.  Example: AUTHORISED | [optional] 
**ShopperInteraction** | **string** | The shopper interaction type of the payment request.  Example: Ecommerce | [optional] 
**ShopperReference** | **string** | The shopperReference passed in the payment request.  Example: AdyenTestShopperXX | [optional] 
**TerminalId** | **string** | The terminal ID used in a point-of-sale payment.  Example: 06022622 | [optional] 
**ThreeDAuthenticated** | **string** | A Boolean value indicating whether 3DS authentication was completed on this payment.  Example: true | [optional] 
**ThreeDAuthenticatedResponse** | **string** | The raw 3DS authentication result from the card issuer.  Example: N | [optional] 
**ThreeDOffered** | **string** | A Boolean value indicating whether 3DS was offered for this payment.  Example: true | [optional] 
**ThreeDOfferedResponse** | **string** | The raw enrollment result from the 3DS directory services of the card schemes.  Example: Y | [optional] 
**ThreeDSVersion** | **string** | The 3D Secure 2 version. | [optional] 
**VisaTransactionId** | **string** | The &#x60;visaTransactionId&#x60;, has a fixed length of 15 numeric characters. &gt; Contact Support Team to enable this field. | [optional] 
**Xid** | **string** | The 3DS transaction ID of the 3DS session sent in notifications. The value is Base64-encoded and is returned for transactions with directoryResponse &#39;N&#39; or &#39;Y&#39;. If you want to submit the xid in your 3D Secure 1 request, use the &#x60;mpiData.xid&#x60;, field.  Example: ODgxNDc2MDg2MDExODk5MAAAAAA&#x3D; | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


