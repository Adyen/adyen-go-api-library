# InputDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** |  Pointer to [**map[string]interface{}**](.md) | Configuration parameters for the required input. | [optional] 
**Details** |  Pointer to [**[]SubInputDetail**](SubInputDetail.md) | Input details can also be provided recursively. | [optional] 
**InputDetails** |  Pointer to [**[]SubInputDetail**](SubInputDetail.md) | Input details can also be provided recursively (deprecated). | [optional] 
**ItemSearchUrl** | **string** | In case of a select, the URL from which to query the items. | [optional] 
**Items** |  Pointer to [**[]Item**](Item.md) | In case of a select, the items to choose from. | [optional] 
**Key** | **string** | The value to provide in the result. | [optional] 
**Optional** | **bool** | True if this input value is optional. | [optional] 
**Type** | **string** | The type of the required input. | [optional] 
**Value** | **string** | The value can be pre-filled, if available. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


