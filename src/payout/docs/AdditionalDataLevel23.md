# AdditionalDataLevel23

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnhancedSchemeDataCustomerReference** | **string** | Customer code, if supplied by a customer. Encoding: ASCII. Max length: 25 characters. &gt; Required for Level 2 and Level 3 data. | [optional] 
**EnhancedSchemeDataTotalTaxAmount** | **float32** | Total tax amount, in minor units.  For example, 2000 means USD 20.00. Max length: 12 characters. &gt; Required for Level 2 and Level 3 data. | [optional] 
**EnhancedSchemeDataFreightAmount** | **float32** | Shipping amount, in minor units.  For example, 2000 means USD 20.00. Max length: 12 characters. | [optional] 
**EnhancedSchemeDataDutyAmount** | **float32** | Duty amount, in minor units.  For example, 2000 means USD 20.00. Max length: 12 characters. | [optional] 
**EnhancedSchemeDataDestinationPostalCode** | **string** | The postal code of a destination address.  Encoding: ASCII. Max length: 10 characters. &gt; Required for American Express. | [optional] 
**EnhancedSchemeDataDestinationStateProvinceCode** | **string** | Destination state or province code.  Encoding: ASCII.Max length: 3 characters. | [optional] 
**EnhancedSchemeDataShipFromPostalCode** | **string** | The postal code of a \&quot;ship-from\&quot; address.  Encoding: ASCII. Max length: 10 characters. | [optional] 
**EnhancedSchemeDataDestinationCountryCode** | **string** | Destination country code.  Encoding: ASCII. Max length: 3 characters. | [optional] 
**EnhancedSchemeDataOrderDate** | **string** | Order date. * Format: &#x60;ddMMyy&#x60;  Encoding: ASCII. Max length: 6 characters. | [optional] 
**EnhancedSchemeDataItemDetailLineItemNrCommodityCode** | **string** | Item commodity code. Encoding: ASCII. Max length: 12 characters. | [optional] 
**EnhancedSchemeDataItemDetailLineItemNrDescription** | **string** | Item description. Encoding: ASCII. Max length: 26 characters. | [optional] 
**EnhancedSchemeDataItemDetailLineItemNrProductCode** | **string** | Product code. Encoding: ASCII. Max length: 12 characters. | [optional] 
**EnhancedSchemeDataItemDetailLineItemNrQuantity** | **float32** | Quantity, specified as an integer value. Value must be greater than 0. Max length: 12 characters. | [optional] 
**EnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure** | **string** | Item unit of measurement. Encoding: ASCII. Max length: 3 characters. | [optional] 
**EnhancedSchemeDataItemDetailLineItemNrUnitPrice** | **float32** | Unit price, specified in [minor units](https://docs.adyen.com/development-resources/currency-codes). Max length: 12 characters. | [optional] 
**EnhancedSchemeDataItemDetailLineItemNrDiscountAmount** | **float32** | Discount amount, in minor units.  For example, 2000 means USD 20.00. Max length: 12 characters. | [optional] 
**EnhancedSchemeDataItemDetailLineItemNrTotalAmount** | **float32** | Total amount, in minor units. For example, 2000 means USD 20.00. Max length: 12 characters. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


