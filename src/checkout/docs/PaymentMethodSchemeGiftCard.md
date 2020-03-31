# PaymentMethodSchemeGiftCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Must be set to **scheme**. | 
**Cvc** | **string** | The [card verification code](https://docs.adyen.com/payments-fundamentals/payment-glossary#card-security-code-cvc-cvv-cid) (1-20 characters). Depending on the card brand, it is known also as: * CVV2/CVC2 – length: 3 digits * CID – length: 4 digits | [optional] 
**ExpiryMonth** | **string** | The card expiry month. Format: 2 digits, zero-padded for single digits. For example: * 03 &#x3D; March * 11 &#x3D; November | 
**ExpiryYear** | **string** | The card expiry year. Format: 4 digits. For example: 2020 | 
**HolderName** | **string** | The name of the cardholder, as printed on the card. | 
**Number** | **string** | The card number (4-19 characters). Do not use any separators. | 
**EncryptedSecurityCode** | **string** | The encrypted [card verification code](https://docs.adyen.com/payments-fundamentals/payment-glossary#card-security-code-cvc-cvv-cid) (1-20 characters). Depending on the card brand, it is known also as: * CVV2/CVC2 – length: 3 digits * CID – length: 4 digits | 
**EncryptedExpiryMonth** | **string** | The encrypted card expiry month. Format: 2 digits, zero-padded for single digits. For example: * 03 &#x3D; March * 11 &#x3D; November | 
**EncryptedExpiryYear** | **string** | The encrypted card expiry year. Format: 4 digits. For example: 2020 | 
**EncryptedCardNumber** | **string** | The encrypted card number (4-19 characters). Do not use any separators. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


