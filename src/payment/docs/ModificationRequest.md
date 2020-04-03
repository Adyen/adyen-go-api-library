# ModificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | **map[string]string** | This field contains additional data, which may be required for a particular modification request.  The additionalData object consists of entries, each of which includes the key and value. | [optional] 
**MerchantAccount** | **string** | The merchant account that is used to process the payment. | 
**ModificationAmount** |  Pointer to [**Amount**](Amount.md) |  | [optional] 
**MpiData** |  Pointer to [**ThreeDSecureData**](ThreeDSecureData.md) |  | [optional] 
**OriginalMerchantReference** | **string** | The original merchant reference to cancel. | [optional] 
**OriginalReference** | **string** | The original pspReference of the payment to modify. This reference is returned in: * authorisation response * authorisation notification | 
**Reference** | **string** | Optionally, you can specify your reference for the payment modification. This reference is visible in Customer Area and in reports. Maximum length: 80 characters. | [optional] 
**Splits** |  Pointer to [**[]Split**](Split.md) | Information on how the payment should be split between accounts when using [Adyen for Platforms](https://docs.adyen.com/marketpay/processing-payments#providing-split-information). | [optional] 
**TenderReference** | **string** | The transaction reference provided by the PED. For point-of-sale integrations only. | [optional] 
**UniqueTerminalId** | **string** | Unique terminal ID for the PED that originally processed the request. For point-of-sale integrations only. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


