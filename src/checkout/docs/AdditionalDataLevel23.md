# AdditionalDataLevel23

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnhancedSchemeDataCustomerReference** | Pointer to **string** | The customer code, if supplied by a customer.  Encoding: ASCII  Max length: 25 characters  Must not start with a space or be all spaces  Must not be all zeros | [optional] 
**EnhancedSchemeDataDestinationCountryCode** | Pointer to **string** | The three-letter [ISO 3166-1 alpha-3 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3) for the destination address.  Encoding: ASCII  Fixed length: 3 characters | [optional] 
**EnhancedSchemeDataDestinationPostalCode** | Pointer to **string** | The postal code of the destination address.  Encoding: ASCII  Max length: 10 characters  Must not start with a space | [optional] 
**EnhancedSchemeDataDestinationStateProvinceCode** | Pointer to **string** | Destination state or province code.  Encoding: ASCII  Max length: 3 characters  Must not start with a space | [optional] 
**EnhancedSchemeDataDutyAmount** | Pointer to **string** | The duty amount, in [minor units](https://docs.adyen.com/development-resources/currency-codes).  For example, 2000 means USD 20.00.  Encoding: Numeric  Max length: 12 characters | [optional] 
**EnhancedSchemeDataFreightAmount** | Pointer to **string** | The shipping amount, in [minor units](https://docs.adyen.com/development-resources/currency-codes).  For example, 2000 means USD 20.00.  Encoding: Numeric  Max length: 12 characters | [optional] 
**EnhancedSchemeDataItemDetailLineItemNrCommodityCode** | Pointer to **string** | The [UNSPC commodity code](https://www.unspsc.org/) of the item.  Encoding: ASCII  Max length: 12 characters  Must not start with a space or be all spaces  Must not be all zeros | [optional] 
**EnhancedSchemeDataItemDetailLineItemNrDescription** | Pointer to **string** | A description of the item.  Encoding: ASCII  Max length: 26 characters  Must not start with a space or be all spaces  Must not be all zeros | [optional] 
**EnhancedSchemeDataItemDetailLineItemNrDiscountAmount** | Pointer to **string** | The discount amount, in [minor units](https://docs.adyen.com/development-resources/currency-codes).  For example, 2000 means USD 20.00.  Encoding: Numeric  Max length: 12 characters | [optional] 
**EnhancedSchemeDataItemDetailLineItemNrProductCode** | Pointer to **string** | The product code.  Encoding: ASCII.  Max length: 12 characters  Must not start with a space or be all spaces  Must not be all zeros | [optional] 
**EnhancedSchemeDataItemDetailLineItemNrQuantity** | Pointer to **string** | The number of items. Must be an integer greater than zero.  Encoding: Numeric  Max length: 12 characters  Must not start with a space or be all spaces   | [optional] 
**EnhancedSchemeDataItemDetailLineItemNrTotalAmount** | Pointer to **string** | The total amount, in [minor units](https://docs.adyen.com/development-resources/currency-codes).  For example, 2000 means USD 20.00.  Max length: 12 characters  Must not start with a space or be all spaces  Must not be all zeros | [optional] 
**EnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure** | Pointer to **string** | The unit of measurement for an item.  Encoding: ASCII  Max length: 3 characters  Must not start with a space or be all spaces  Must not be all zeros | [optional] 
**EnhancedSchemeDataItemDetailLineItemNrUnitPrice** | Pointer to **string** | The unit price in [minor units](https://docs.adyen.com/development-resources/currency-codes).  For example, 2000 means USD 20.00.  Encoding: Numeric  Max length: 12 characters | [optional] 
**EnhancedSchemeDataOrderDate** | Pointer to **string** | The order date. * Format: &#x60;ddMMyy&#x60;  Encoding: ASCII  Max length: 6 characters | [optional] 
**EnhancedSchemeDataShipFromPostalCode** | Pointer to **string** | The postal code of the address the item is shipped from.  Encoding: ASCII  Max length: 10 characters  Must not start with a space or be all spaces  Must not be all zeros | [optional] 
**EnhancedSchemeDataTotalTaxAmount** | Pointer to **string** | The total tax amount, in [minor units](https://docs.adyen.com/development-resources/currency-codes).  For example, 2000 means USD 20.00.  Encoding: Numeric  Max length: 12 characters   | [optional] 

## Methods

### NewAdditionalDataLevel23

`func NewAdditionalDataLevel23() *AdditionalDataLevel23`

NewAdditionalDataLevel23 instantiates a new AdditionalDataLevel23 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalDataLevel23WithDefaults

`func NewAdditionalDataLevel23WithDefaults() *AdditionalDataLevel23`

NewAdditionalDataLevel23WithDefaults instantiates a new AdditionalDataLevel23 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnhancedSchemeDataCustomerReference

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataCustomerReference() string`

GetEnhancedSchemeDataCustomerReference returns the EnhancedSchemeDataCustomerReference field if non-nil, zero value otherwise.

### GetEnhancedSchemeDataCustomerReferenceOk

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataCustomerReferenceOk() (*string, bool)`

GetEnhancedSchemeDataCustomerReferenceOk returns a tuple with the EnhancedSchemeDataCustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSchemeDataCustomerReference

`func (o *AdditionalDataLevel23) SetEnhancedSchemeDataCustomerReference(v string)`

SetEnhancedSchemeDataCustomerReference sets EnhancedSchemeDataCustomerReference field to given value.

### HasEnhancedSchemeDataCustomerReference

`func (o *AdditionalDataLevel23) HasEnhancedSchemeDataCustomerReference() bool`

HasEnhancedSchemeDataCustomerReference returns a boolean if a field has been set.

### GetEnhancedSchemeDataDestinationCountryCode

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataDestinationCountryCode() string`

GetEnhancedSchemeDataDestinationCountryCode returns the EnhancedSchemeDataDestinationCountryCode field if non-nil, zero value otherwise.

### GetEnhancedSchemeDataDestinationCountryCodeOk

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataDestinationCountryCodeOk() (*string, bool)`

GetEnhancedSchemeDataDestinationCountryCodeOk returns a tuple with the EnhancedSchemeDataDestinationCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSchemeDataDestinationCountryCode

`func (o *AdditionalDataLevel23) SetEnhancedSchemeDataDestinationCountryCode(v string)`

SetEnhancedSchemeDataDestinationCountryCode sets EnhancedSchemeDataDestinationCountryCode field to given value.

### HasEnhancedSchemeDataDestinationCountryCode

`func (o *AdditionalDataLevel23) HasEnhancedSchemeDataDestinationCountryCode() bool`

HasEnhancedSchemeDataDestinationCountryCode returns a boolean if a field has been set.

### GetEnhancedSchemeDataDestinationPostalCode

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataDestinationPostalCode() string`

GetEnhancedSchemeDataDestinationPostalCode returns the EnhancedSchemeDataDestinationPostalCode field if non-nil, zero value otherwise.

### GetEnhancedSchemeDataDestinationPostalCodeOk

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataDestinationPostalCodeOk() (*string, bool)`

GetEnhancedSchemeDataDestinationPostalCodeOk returns a tuple with the EnhancedSchemeDataDestinationPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSchemeDataDestinationPostalCode

`func (o *AdditionalDataLevel23) SetEnhancedSchemeDataDestinationPostalCode(v string)`

SetEnhancedSchemeDataDestinationPostalCode sets EnhancedSchemeDataDestinationPostalCode field to given value.

### HasEnhancedSchemeDataDestinationPostalCode

`func (o *AdditionalDataLevel23) HasEnhancedSchemeDataDestinationPostalCode() bool`

HasEnhancedSchemeDataDestinationPostalCode returns a boolean if a field has been set.

### GetEnhancedSchemeDataDestinationStateProvinceCode

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataDestinationStateProvinceCode() string`

GetEnhancedSchemeDataDestinationStateProvinceCode returns the EnhancedSchemeDataDestinationStateProvinceCode field if non-nil, zero value otherwise.

### GetEnhancedSchemeDataDestinationStateProvinceCodeOk

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataDestinationStateProvinceCodeOk() (*string, bool)`

GetEnhancedSchemeDataDestinationStateProvinceCodeOk returns a tuple with the EnhancedSchemeDataDestinationStateProvinceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSchemeDataDestinationStateProvinceCode

`func (o *AdditionalDataLevel23) SetEnhancedSchemeDataDestinationStateProvinceCode(v string)`

SetEnhancedSchemeDataDestinationStateProvinceCode sets EnhancedSchemeDataDestinationStateProvinceCode field to given value.

### HasEnhancedSchemeDataDestinationStateProvinceCode

`func (o *AdditionalDataLevel23) HasEnhancedSchemeDataDestinationStateProvinceCode() bool`

HasEnhancedSchemeDataDestinationStateProvinceCode returns a boolean if a field has been set.

### GetEnhancedSchemeDataDutyAmount

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataDutyAmount() string`

GetEnhancedSchemeDataDutyAmount returns the EnhancedSchemeDataDutyAmount field if non-nil, zero value otherwise.

### GetEnhancedSchemeDataDutyAmountOk

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataDutyAmountOk() (*string, bool)`

GetEnhancedSchemeDataDutyAmountOk returns a tuple with the EnhancedSchemeDataDutyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSchemeDataDutyAmount

`func (o *AdditionalDataLevel23) SetEnhancedSchemeDataDutyAmount(v string)`

SetEnhancedSchemeDataDutyAmount sets EnhancedSchemeDataDutyAmount field to given value.

### HasEnhancedSchemeDataDutyAmount

`func (o *AdditionalDataLevel23) HasEnhancedSchemeDataDutyAmount() bool`

HasEnhancedSchemeDataDutyAmount returns a boolean if a field has been set.

### GetEnhancedSchemeDataFreightAmount

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataFreightAmount() string`

GetEnhancedSchemeDataFreightAmount returns the EnhancedSchemeDataFreightAmount field if non-nil, zero value otherwise.

### GetEnhancedSchemeDataFreightAmountOk

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataFreightAmountOk() (*string, bool)`

GetEnhancedSchemeDataFreightAmountOk returns a tuple with the EnhancedSchemeDataFreightAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSchemeDataFreightAmount

`func (o *AdditionalDataLevel23) SetEnhancedSchemeDataFreightAmount(v string)`

SetEnhancedSchemeDataFreightAmount sets EnhancedSchemeDataFreightAmount field to given value.

### HasEnhancedSchemeDataFreightAmount

`func (o *AdditionalDataLevel23) HasEnhancedSchemeDataFreightAmount() bool`

HasEnhancedSchemeDataFreightAmount returns a boolean if a field has been set.

### GetEnhancedSchemeDataItemDetailLineItemNrCommodityCode

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrCommodityCode() string`

GetEnhancedSchemeDataItemDetailLineItemNrCommodityCode returns the EnhancedSchemeDataItemDetailLineItemNrCommodityCode field if non-nil, zero value otherwise.

### GetEnhancedSchemeDataItemDetailLineItemNrCommodityCodeOk

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrCommodityCodeOk() (*string, bool)`

GetEnhancedSchemeDataItemDetailLineItemNrCommodityCodeOk returns a tuple with the EnhancedSchemeDataItemDetailLineItemNrCommodityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSchemeDataItemDetailLineItemNrCommodityCode

`func (o *AdditionalDataLevel23) SetEnhancedSchemeDataItemDetailLineItemNrCommodityCode(v string)`

SetEnhancedSchemeDataItemDetailLineItemNrCommodityCode sets EnhancedSchemeDataItemDetailLineItemNrCommodityCode field to given value.

### HasEnhancedSchemeDataItemDetailLineItemNrCommodityCode

`func (o *AdditionalDataLevel23) HasEnhancedSchemeDataItemDetailLineItemNrCommodityCode() bool`

HasEnhancedSchemeDataItemDetailLineItemNrCommodityCode returns a boolean if a field has been set.

### GetEnhancedSchemeDataItemDetailLineItemNrDescription

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrDescription() string`

GetEnhancedSchemeDataItemDetailLineItemNrDescription returns the EnhancedSchemeDataItemDetailLineItemNrDescription field if non-nil, zero value otherwise.

### GetEnhancedSchemeDataItemDetailLineItemNrDescriptionOk

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrDescriptionOk() (*string, bool)`

GetEnhancedSchemeDataItemDetailLineItemNrDescriptionOk returns a tuple with the EnhancedSchemeDataItemDetailLineItemNrDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSchemeDataItemDetailLineItemNrDescription

`func (o *AdditionalDataLevel23) SetEnhancedSchemeDataItemDetailLineItemNrDescription(v string)`

SetEnhancedSchemeDataItemDetailLineItemNrDescription sets EnhancedSchemeDataItemDetailLineItemNrDescription field to given value.

### HasEnhancedSchemeDataItemDetailLineItemNrDescription

`func (o *AdditionalDataLevel23) HasEnhancedSchemeDataItemDetailLineItemNrDescription() bool`

HasEnhancedSchemeDataItemDetailLineItemNrDescription returns a boolean if a field has been set.

### GetEnhancedSchemeDataItemDetailLineItemNrDiscountAmount

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrDiscountAmount() string`

GetEnhancedSchemeDataItemDetailLineItemNrDiscountAmount returns the EnhancedSchemeDataItemDetailLineItemNrDiscountAmount field if non-nil, zero value otherwise.

### GetEnhancedSchemeDataItemDetailLineItemNrDiscountAmountOk

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrDiscountAmountOk() (*string, bool)`

GetEnhancedSchemeDataItemDetailLineItemNrDiscountAmountOk returns a tuple with the EnhancedSchemeDataItemDetailLineItemNrDiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSchemeDataItemDetailLineItemNrDiscountAmount

`func (o *AdditionalDataLevel23) SetEnhancedSchemeDataItemDetailLineItemNrDiscountAmount(v string)`

SetEnhancedSchemeDataItemDetailLineItemNrDiscountAmount sets EnhancedSchemeDataItemDetailLineItemNrDiscountAmount field to given value.

### HasEnhancedSchemeDataItemDetailLineItemNrDiscountAmount

`func (o *AdditionalDataLevel23) HasEnhancedSchemeDataItemDetailLineItemNrDiscountAmount() bool`

HasEnhancedSchemeDataItemDetailLineItemNrDiscountAmount returns a boolean if a field has been set.

### GetEnhancedSchemeDataItemDetailLineItemNrProductCode

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrProductCode() string`

GetEnhancedSchemeDataItemDetailLineItemNrProductCode returns the EnhancedSchemeDataItemDetailLineItemNrProductCode field if non-nil, zero value otherwise.

### GetEnhancedSchemeDataItemDetailLineItemNrProductCodeOk

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrProductCodeOk() (*string, bool)`

GetEnhancedSchemeDataItemDetailLineItemNrProductCodeOk returns a tuple with the EnhancedSchemeDataItemDetailLineItemNrProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSchemeDataItemDetailLineItemNrProductCode

`func (o *AdditionalDataLevel23) SetEnhancedSchemeDataItemDetailLineItemNrProductCode(v string)`

SetEnhancedSchemeDataItemDetailLineItemNrProductCode sets EnhancedSchemeDataItemDetailLineItemNrProductCode field to given value.

### HasEnhancedSchemeDataItemDetailLineItemNrProductCode

`func (o *AdditionalDataLevel23) HasEnhancedSchemeDataItemDetailLineItemNrProductCode() bool`

HasEnhancedSchemeDataItemDetailLineItemNrProductCode returns a boolean if a field has been set.

### GetEnhancedSchemeDataItemDetailLineItemNrQuantity

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrQuantity() string`

GetEnhancedSchemeDataItemDetailLineItemNrQuantity returns the EnhancedSchemeDataItemDetailLineItemNrQuantity field if non-nil, zero value otherwise.

### GetEnhancedSchemeDataItemDetailLineItemNrQuantityOk

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrQuantityOk() (*string, bool)`

GetEnhancedSchemeDataItemDetailLineItemNrQuantityOk returns a tuple with the EnhancedSchemeDataItemDetailLineItemNrQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSchemeDataItemDetailLineItemNrQuantity

`func (o *AdditionalDataLevel23) SetEnhancedSchemeDataItemDetailLineItemNrQuantity(v string)`

SetEnhancedSchemeDataItemDetailLineItemNrQuantity sets EnhancedSchemeDataItemDetailLineItemNrQuantity field to given value.

### HasEnhancedSchemeDataItemDetailLineItemNrQuantity

`func (o *AdditionalDataLevel23) HasEnhancedSchemeDataItemDetailLineItemNrQuantity() bool`

HasEnhancedSchemeDataItemDetailLineItemNrQuantity returns a boolean if a field has been set.

### GetEnhancedSchemeDataItemDetailLineItemNrTotalAmount

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrTotalAmount() string`

GetEnhancedSchemeDataItemDetailLineItemNrTotalAmount returns the EnhancedSchemeDataItemDetailLineItemNrTotalAmount field if non-nil, zero value otherwise.

### GetEnhancedSchemeDataItemDetailLineItemNrTotalAmountOk

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrTotalAmountOk() (*string, bool)`

GetEnhancedSchemeDataItemDetailLineItemNrTotalAmountOk returns a tuple with the EnhancedSchemeDataItemDetailLineItemNrTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSchemeDataItemDetailLineItemNrTotalAmount

`func (o *AdditionalDataLevel23) SetEnhancedSchemeDataItemDetailLineItemNrTotalAmount(v string)`

SetEnhancedSchemeDataItemDetailLineItemNrTotalAmount sets EnhancedSchemeDataItemDetailLineItemNrTotalAmount field to given value.

### HasEnhancedSchemeDataItemDetailLineItemNrTotalAmount

`func (o *AdditionalDataLevel23) HasEnhancedSchemeDataItemDetailLineItemNrTotalAmount() bool`

HasEnhancedSchemeDataItemDetailLineItemNrTotalAmount returns a boolean if a field has been set.

### GetEnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure() string`

GetEnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure returns the EnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure field if non-nil, zero value otherwise.

### GetEnhancedSchemeDataItemDetailLineItemNrUnitOfMeasureOk

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrUnitOfMeasureOk() (*string, bool)`

GetEnhancedSchemeDataItemDetailLineItemNrUnitOfMeasureOk returns a tuple with the EnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure

`func (o *AdditionalDataLevel23) SetEnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure(v string)`

SetEnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure sets EnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure field to given value.

### HasEnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure

`func (o *AdditionalDataLevel23) HasEnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure() bool`

HasEnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure returns a boolean if a field has been set.

### GetEnhancedSchemeDataItemDetailLineItemNrUnitPrice

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrUnitPrice() string`

GetEnhancedSchemeDataItemDetailLineItemNrUnitPrice returns the EnhancedSchemeDataItemDetailLineItemNrUnitPrice field if non-nil, zero value otherwise.

### GetEnhancedSchemeDataItemDetailLineItemNrUnitPriceOk

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataItemDetailLineItemNrUnitPriceOk() (*string, bool)`

GetEnhancedSchemeDataItemDetailLineItemNrUnitPriceOk returns a tuple with the EnhancedSchemeDataItemDetailLineItemNrUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSchemeDataItemDetailLineItemNrUnitPrice

`func (o *AdditionalDataLevel23) SetEnhancedSchemeDataItemDetailLineItemNrUnitPrice(v string)`

SetEnhancedSchemeDataItemDetailLineItemNrUnitPrice sets EnhancedSchemeDataItemDetailLineItemNrUnitPrice field to given value.

### HasEnhancedSchemeDataItemDetailLineItemNrUnitPrice

`func (o *AdditionalDataLevel23) HasEnhancedSchemeDataItemDetailLineItemNrUnitPrice() bool`

HasEnhancedSchemeDataItemDetailLineItemNrUnitPrice returns a boolean if a field has been set.

### GetEnhancedSchemeDataOrderDate

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataOrderDate() string`

GetEnhancedSchemeDataOrderDate returns the EnhancedSchemeDataOrderDate field if non-nil, zero value otherwise.

### GetEnhancedSchemeDataOrderDateOk

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataOrderDateOk() (*string, bool)`

GetEnhancedSchemeDataOrderDateOk returns a tuple with the EnhancedSchemeDataOrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSchemeDataOrderDate

`func (o *AdditionalDataLevel23) SetEnhancedSchemeDataOrderDate(v string)`

SetEnhancedSchemeDataOrderDate sets EnhancedSchemeDataOrderDate field to given value.

### HasEnhancedSchemeDataOrderDate

`func (o *AdditionalDataLevel23) HasEnhancedSchemeDataOrderDate() bool`

HasEnhancedSchemeDataOrderDate returns a boolean if a field has been set.

### GetEnhancedSchemeDataShipFromPostalCode

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataShipFromPostalCode() string`

GetEnhancedSchemeDataShipFromPostalCode returns the EnhancedSchemeDataShipFromPostalCode field if non-nil, zero value otherwise.

### GetEnhancedSchemeDataShipFromPostalCodeOk

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataShipFromPostalCodeOk() (*string, bool)`

GetEnhancedSchemeDataShipFromPostalCodeOk returns a tuple with the EnhancedSchemeDataShipFromPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSchemeDataShipFromPostalCode

`func (o *AdditionalDataLevel23) SetEnhancedSchemeDataShipFromPostalCode(v string)`

SetEnhancedSchemeDataShipFromPostalCode sets EnhancedSchemeDataShipFromPostalCode field to given value.

### HasEnhancedSchemeDataShipFromPostalCode

`func (o *AdditionalDataLevel23) HasEnhancedSchemeDataShipFromPostalCode() bool`

HasEnhancedSchemeDataShipFromPostalCode returns a boolean if a field has been set.

### GetEnhancedSchemeDataTotalTaxAmount

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataTotalTaxAmount() string`

GetEnhancedSchemeDataTotalTaxAmount returns the EnhancedSchemeDataTotalTaxAmount field if non-nil, zero value otherwise.

### GetEnhancedSchemeDataTotalTaxAmountOk

`func (o *AdditionalDataLevel23) GetEnhancedSchemeDataTotalTaxAmountOk() (*string, bool)`

GetEnhancedSchemeDataTotalTaxAmountOk returns a tuple with the EnhancedSchemeDataTotalTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSchemeDataTotalTaxAmount

`func (o *AdditionalDataLevel23) SetEnhancedSchemeDataTotalTaxAmount(v string)`

SetEnhancedSchemeDataTotalTaxAmount sets EnhancedSchemeDataTotalTaxAmount field to given value.

### HasEnhancedSchemeDataTotalTaxAmount

`func (o *AdditionalDataLevel23) HasEnhancedSchemeDataTotalTaxAmount() bool`

HasEnhancedSchemeDataTotalTaxAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


