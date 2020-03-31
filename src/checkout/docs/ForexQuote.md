# ForexQuote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | The account name. | [optional] 
**AccountType** | **string** | The account type. | [optional] 
**BaseAmount** | [**Amount**](Amount.md) |  | [optional] 
**BasePoints** | **int32** | The base points. | 
**Buy** | [**Amount**](Amount.md) |  | [optional] 
**Interbank** | [**Amount**](Amount.md) |  | [optional] 
**Reference** | **string** | The reference assigned to the forex quote request. | [optional] 
**Sell** | [**Amount**](Amount.md) |  | [optional] 
**Signature** | **string** | The signature to validate the integrity. | [optional] 
**Source** | **string** | The source of the forex quote. | [optional] 
**Type** | **string** | The type of forex. | [optional] 
**ValidTill** | [**time.Time**](time.Time.md) | The date until which the forex quote is valid. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


