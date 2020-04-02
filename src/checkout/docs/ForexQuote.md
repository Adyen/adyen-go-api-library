# ForexQuote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | The account name. | [optional] 
**AccountType** | **string** | The account type. | [optional] 
**BaseAmount** |  Pointer to [**Amount**](Amount.md) |  | [optional] 
**BasePoints** | **int32** | The base points. | 
**Buy** |  Pointer to [**Amount**](Amount.md) |  | [optional] 
**Interbank** |  Pointer to [**Amount**](Amount.md) |  | [optional] 
**Reference** | **string** | The reference assigned to the forex quote request. | [optional] 
**Sell** |  Pointer to [**Amount**](Amount.md) |  | [optional] 
**Signature** | **string** | The signature to validate the integrity. | [optional] 
**Source** | **string** | The source of the forex quote. | [optional] 
**Type** | **string** | The type of forex. | [optional] 
**ValidTill** |  Pointer to [**time.Time**](time.Time.md) | The date until which the forex quote is valid. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


