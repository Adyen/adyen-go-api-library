# AdditionalDataRiskStandalone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvsResultRaw** | **string** | Raw AVS result received from the acquirer, where available. Example: D | [optional] 
**Bin** | **string** | The Bank Identification Number of a credit card, which is the first six digits of a card number. Required for [tokenized card request](https://docs.adyen.com/risk-management/standalone-risk#tokenised-pan-request). | [optional] 
**CvcResultRaw** | **string** | Raw CVC result received from the acquirer, where available. Example: 1 | [optional] 
**RiskToken** | **string** | Unique identifier or token for the shopper&#39;s card details. | [optional] 
**ThreeDAuthenticated** | **string** | A Boolean value indicating whether 3DS authentication was completed on this payment. Example: true | [optional] 
**ThreeDOffered** | **string** | A Boolean value indicating whether 3DS was offered for this payment. Example: true | [optional] 
**TokenDataType** | **string** | Required for PayPal payments only. The only supported value is: **paypal**. | [optional] 
**PayPalProtectionEligibility** | **string** | Allowed values: * **Eligible** — Merchant is protected by PayPal&#39;s Seller Protection Policy for Unauthorized Payments and Item Not Received. * **PartiallyEligible** — Merchant is protected by PayPal&#39;s Seller Protection Policy for Item Not Received. * **Ineligible** — Merchant is not protected under the Seller Protection Policy. | [optional] 
**PayPalPayerId** | **string** | Unique PayPal Customer Account identification number. Character length and limitations: 13 single-byte alphanumeric characters. | [optional] 
**PayPalTransactionId** | **string** | Unique transaction ID of the payment. | [optional] 
**PayPalCountryCode** | **string** | Shopper&#39;s country of residence in the form of ISO standard 3166 2-character country codes. | [optional] 
**PayPalFirstName** | **string** | Shopper&#39;s first name. | [optional] 
**PayPalLastName** | **string** | Shopper&#39;s last name. | [optional] 
**PayPalPhone** | **string** | Shopper&#39;s phone number. | [optional] 
**PayPalEmailId** | **string** | Shopper&#39;s email. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


