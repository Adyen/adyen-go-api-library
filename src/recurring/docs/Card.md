# Card

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cvc** | **string** | The [card verification code](https://docs.adyen.com/payments-fundamentals/payment-glossary#card-security-code-cvc-cvv-cid) (1-20 characters). Depending on the card brand, it is known also as: * CVV2/CVC2 – length: 3 digits * CID – length: 4 digits &gt; If you are using [Client-Side Encryption](https://docs.adyen.com/classic-integration/cse-integration-ecommerce), the CVC code is present in the encrypted data. You must never post the card details to the server. &gt; This field must be always present in a [one-click payment request](https://docs.adyen.com/classic-integration/recurring-payments). &gt; When this value is returned in a response, it is always empty because it is not stored. | [optional] 
**ExpiryMonth** | **string** | The card expiry month. Format: 2 digits, zero-padded for single digits. For example: * 03 &#x3D; March * 11 &#x3D; November | 
**ExpiryYear** | **string** | The card expiry year. Format: 4 digits. For example: 2020 | 
**HolderName** | **string** | The name of the cardholder, as printed on the card. | 
**IssueNumber** | **string** | The issue number of the card (for some UK debit cards only). | [optional] 
**Number** | **string** | The card number (4-19 characters). Do not use any separators. When this value is returned in a response, only the last 4 digits of the card number are returned. | 
**StartMonth** | **string** | The month component of the start date (for some UK debit cards only). | [optional] 
**StartYear** | **string** | The year component of the start date (for some UK debit cards only). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


