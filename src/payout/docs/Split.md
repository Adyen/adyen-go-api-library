# Split

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | The account to which this split applies.  &gt;Required if the type is &#x60;MarketPlace&#x60;. | [optional] 
**Amount** |  Pointer to [**SplitAmount**](SplitAmount.md) |  | 
**Description** | **string** | A description of this split. | [optional] 
**Reference** | **string** | The reference of this split. Used to link other operations (e.g. captures and refunds) to this split.  &gt;Required if the type is &#x60;MarketPlace&#x60;. | [optional] 
**Type** | **string** | The type of this split.  &gt;Permitted values: &#x60;Default&#x60;, &#x60;PaymentFee&#x60;, &#x60;VAT&#x60;, &#x60;Commission&#x60;, &#x60;MarketPlace&#x60;, &#x60;BalanceAccount&#x60;. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


