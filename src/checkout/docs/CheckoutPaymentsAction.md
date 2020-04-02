# CheckoutPaymentsAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternativeReference** | **string** | The voucher alternative reference code. | [optional] 
**Data** |  Pointer to [**map[string]interface{}**](.md) | When the redirect URL must be accessed via POST, use this data to post to the redirect URL. | [optional] 
**DownloadUrl** | **string** | The URL to download the voucher. | [optional] 
**ExpiresAt** | **string** | The date time of the voucher expiry. | [optional] 
**InitialAmount** |  Pointer to [**Amount**](Amount.md) |  | [optional] 
**InstructionsUrl** | **string** | The URL to the detailed instructions to make payment using the voucher. | [optional] 
**Issuer** | **string** | The issuer of the voucher. | [optional] 
**MaskedTelephoneNumber** | **string** | The shopper telephone number (partially masked). | [optional] 
**MerchantName** | **string** | The merchant name. | [optional] 
**MerchantReference** | **string** | The merchant reference. | [optional] 
**Method** | **string** | Specifies the HTTP method, for example GET or POST. | [optional] 
**PaymentData** | **string** | When non-empty, contains a value that you must submit to the &#x60;/payments/details&#x60; endpoint. In some cases, required for polling. | [optional] 
**PaymentMethodType** | **string** | Specifies the payment method. | [optional] 
**QrCodeData** | **string** | The contents of the QR code as a UTF8 string. | [optional] 
**Reference** | **string** | The voucher reference code. | [optional] 
**ShopperEmail** | **string** | The shopper email. | [optional] 
**ShopperName** | **string** | The shopper name. | [optional] 
**Surcharge** |  Pointer to [**Amount**](Amount.md) |  | [optional] 
**Token** | **string** | A token to pass to the 3DS2 Component to get the fingerprint/challenge. | [optional] 
**TotalAmount** |  Pointer to [**Amount**](Amount.md) |  | [optional] 
**Type** | **string** | Enum that specifies the action that needs to be taken by the client. | [optional] 
**Url** | **string** | Specifies the URL to redirect to. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


