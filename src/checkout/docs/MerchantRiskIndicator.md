# MerchantRiskIndicator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressMatch** | **bool** | Whether the chosen delivery address is identical to the billing address. | [optional] 
**DeliveryAddressIndicator** | **string** | Indicator regarding the delivery address. Allowed values: * &#x60;shipToBillingAddress&#x60; * &#x60;shipToVerifiedAddress&#x60; * &#x60;shipToNewAddress&#x60; * &#x60;shipToStore&#x60; * &#x60;digitalGoods&#x60; * &#x60;goodsNotShipped&#x60; * &#x60;other&#x60; | [optional] 
**DeliveryEmail** | **string** | The delivery email address (for digital goods). | [optional] 
**DeliveryTimeframe** | **string** | The estimated delivery time for the shopper to receive the goods. Allowed values: * &#x60;electronicDelivery&#x60; * &#x60;sameDayShipping&#x60; * &#x60;overnightShipping&#x60; * &#x60;twoOrMoreDaysShipping&#x60; | [optional] 
**GiftCardAmount** |  Pointer to [**Amount**](Amount.md) |  | [optional] 
**GiftCardCount** | **int32** | Number of individual prepaid or gift cards used for this purchase. | [optional] 
**PreOrderDate** |  Pointer to [**time.Time**](time.Time.md) | For pre-order purchases, the expected date this product will be available to the shopper. | [optional] 
**PreOrderPurchase** | **bool** | Indicator for whether this transaction is for pre-ordering a product. | [optional] 
**ReorderItems** | **bool** | Indicator for whether the shopper has already purchased the same items in the past. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


