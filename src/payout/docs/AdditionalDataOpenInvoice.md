# AdditionalDataOpenInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpeninvoicedataNumberOfLines** | **int32** | The number of invoice lines included in &#x60;openinvoicedata&#x60;.  There needs to be at least one line, so &#x60;numberOfLines&#x60; needs to be at least 1. | [optional] 
**OpeninvoicedataMerchantData** | **string** | Holds different merchant data points like product, purchase, customer, and so on. It takes data in a Base64 encoded string.  The &#x60;merchantData&#x60; parameter needs to be added to the &#x60;openinvoicedata&#x60; signature at the end.  Since the field is optional, if it&#39;s not included it does not impact computing the merchant signature.  Applies only to Klarna.  You can contact Klarna for the format and structure of the string. | [optional] 
**OpeninvoicedataLineItemNrCurrencyCode** | **string** | The three-character ISO currency code. | [optional] 
**OpeninvoicedataLineItemNrDescription** | **string** | A text description of the product the invoice line refers to. | [optional] 
**OpeninvoicedataLineItemNrItemAmount** | **int32** | The price for one item in the invoice line, represented in minor units.  The due amount for the item, VAT excluded. | [optional] 
**OpeninvoicedataLineItemNrItemVatAmount** | **int32** | The VAT due for one item in the invoice line, represented in minor units. | [optional] 
**OpeninvoicedataLineItemNrItemVatPercentage** | **int32** | The VAT percentage for one item in the invoice line, represented in minor units.  For example, 19% VAT is specified as 1900. | [optional] 
**OpeninvoicedataLineItemNrItemId** | **string** | A unique id for this item. Required for RatePay if the description of each item is not unique. | [optional] 
**OpeninvoicedataLineItemNrNumberOfItems** | **int32** | The number of units purchased of a specific product. | [optional] 
**OpeninvoicedataLineItemNrVatCategory** | **string** | Required for AfterPay. The country-specific VAT category a product falls under.  Allowed values: * High * Low * None. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


