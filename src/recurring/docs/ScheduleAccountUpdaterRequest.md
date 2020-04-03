# ScheduleAccountUpdaterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** |  Pointer to [**map[string]interface{}**](.md) | This field contains additional data, which may be required for a particular request. | [optional] 
**Card** |  Pointer to [**Card**](Card.md) |  | [optional] 
**MerchantAccount** | **string** | Account of the merchant. | 
**Reference** | **string** | A reference that merchants can apply for the call. | 
**SelectedRecurringDetailReference** | **string** | The selected detail recurring reference.  Optional if &#x60;card&#x60; is provided. | [optional] 
**ShopperReference** | **string** | The reference of the shopper that owns the recurring contract.  Optional if &#x60;card&#x60; is provided. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


