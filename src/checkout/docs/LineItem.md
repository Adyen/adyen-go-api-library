# LineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountExcludingTax** | Pointer to **int64** | Item amount excluding the tax, in minor units. | [optional] 
**AmountIncludingTax** | Pointer to **int64** | Item amount including the tax, in minor units. | [optional] 
**Brand** | Pointer to **string** | Brand of the item. | [optional] 
**Color** | Pointer to **string** | Color of the item. | [optional] 
**Description** | Pointer to **string** | Description of the line item. | [optional] 
**Id** | Pointer to **string** | ID of the line item. | [optional] 
**ImageUrl** | Pointer to **string** | Link to the picture of the purchased item. | [optional] 
**ItemCategory** | Pointer to **string** | Item category, used by the RatePay payment method. | [optional] 
**Manufacturer** | Pointer to **string** | Manufacturer of the item. | [optional] 
**ProductUrl** | Pointer to **string** | Link to the purchased item. | [optional] 
**Quantity** | Pointer to **int64** | Number of items. | [optional] 
**ReceiverEmail** | Pointer to **string** | Email associated with the given product in the basket (usually in electronic gift cards). | [optional] 
**Size** | Pointer to **string** | Size of the item. | [optional] 
**Sku** | Pointer to **string** | Stock keeping unit. | [optional] 
**TaxAmount** | Pointer to **int64** | Tax amount, in minor units. | [optional] 
**TaxPercentage** | Pointer to **int64** | Tax percentage, in minor units. | [optional] 
**Upc** | Pointer to **string** | Universal Product Code. | [optional] 

## Methods

### NewLineItem

`func NewLineItem() *LineItem`

NewLineItem instantiates a new LineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemWithDefaults

`func NewLineItemWithDefaults() *LineItem`

NewLineItemWithDefaults instantiates a new LineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountExcludingTax

`func (o *LineItem) GetAmountExcludingTax() int64`

GetAmountExcludingTax returns the AmountExcludingTax field if non-nil, zero value otherwise.

### GetAmountExcludingTaxOk

`func (o *LineItem) GetAmountExcludingTaxOk() (*int64, bool)`

GetAmountExcludingTaxOk returns a tuple with the AmountExcludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountExcludingTax

`func (o *LineItem) SetAmountExcludingTax(v int64)`

SetAmountExcludingTax sets AmountExcludingTax field to given value.

### HasAmountExcludingTax

`func (o *LineItem) HasAmountExcludingTax() bool`

HasAmountExcludingTax returns a boolean if a field has been set.

### GetAmountIncludingTax

`func (o *LineItem) GetAmountIncludingTax() int64`

GetAmountIncludingTax returns the AmountIncludingTax field if non-nil, zero value otherwise.

### GetAmountIncludingTaxOk

`func (o *LineItem) GetAmountIncludingTaxOk() (*int64, bool)`

GetAmountIncludingTaxOk returns a tuple with the AmountIncludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountIncludingTax

`func (o *LineItem) SetAmountIncludingTax(v int64)`

SetAmountIncludingTax sets AmountIncludingTax field to given value.

### HasAmountIncludingTax

`func (o *LineItem) HasAmountIncludingTax() bool`

HasAmountIncludingTax returns a boolean if a field has been set.

### GetBrand

`func (o *LineItem) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *LineItem) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *LineItem) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *LineItem) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetColor

`func (o *LineItem) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *LineItem) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *LineItem) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *LineItem) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDescription

`func (o *LineItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LineItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LineItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LineItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *LineItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LineItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LineItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LineItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageUrl

`func (o *LineItem) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *LineItem) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *LineItem) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *LineItem) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetItemCategory

`func (o *LineItem) GetItemCategory() string`

GetItemCategory returns the ItemCategory field if non-nil, zero value otherwise.

### GetItemCategoryOk

`func (o *LineItem) GetItemCategoryOk() (*string, bool)`

GetItemCategoryOk returns a tuple with the ItemCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCategory

`func (o *LineItem) SetItemCategory(v string)`

SetItemCategory sets ItemCategory field to given value.

### HasItemCategory

`func (o *LineItem) HasItemCategory() bool`

HasItemCategory returns a boolean if a field has been set.

### GetManufacturer

`func (o *LineItem) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *LineItem) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *LineItem) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *LineItem) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetProductUrl

`func (o *LineItem) GetProductUrl() string`

GetProductUrl returns the ProductUrl field if non-nil, zero value otherwise.

### GetProductUrlOk

`func (o *LineItem) GetProductUrlOk() (*string, bool)`

GetProductUrlOk returns a tuple with the ProductUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductUrl

`func (o *LineItem) SetProductUrl(v string)`

SetProductUrl sets ProductUrl field to given value.

### HasProductUrl

`func (o *LineItem) HasProductUrl() bool`

HasProductUrl returns a boolean if a field has been set.

### GetQuantity

`func (o *LineItem) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *LineItem) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *LineItem) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *LineItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetReceiverEmail

`func (o *LineItem) GetReceiverEmail() string`

GetReceiverEmail returns the ReceiverEmail field if non-nil, zero value otherwise.

### GetReceiverEmailOk

`func (o *LineItem) GetReceiverEmailOk() (*string, bool)`

GetReceiverEmailOk returns a tuple with the ReceiverEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverEmail

`func (o *LineItem) SetReceiverEmail(v string)`

SetReceiverEmail sets ReceiverEmail field to given value.

### HasReceiverEmail

`func (o *LineItem) HasReceiverEmail() bool`

HasReceiverEmail returns a boolean if a field has been set.

### GetSize

`func (o *LineItem) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *LineItem) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *LineItem) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *LineItem) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSku

`func (o *LineItem) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *LineItem) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *LineItem) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *LineItem) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetTaxAmount

`func (o *LineItem) GetTaxAmount() int64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *LineItem) GetTaxAmountOk() (*int64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *LineItem) SetTaxAmount(v int64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *LineItem) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *LineItem) GetTaxPercentage() int64`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *LineItem) GetTaxPercentageOk() (*int64, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *LineItem) SetTaxPercentage(v int64)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *LineItem) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetUpc

`func (o *LineItem) GetUpc() string`

GetUpc returns the Upc field if non-nil, zero value otherwise.

### GetUpcOk

`func (o *LineItem) GetUpcOk() (*string, bool)`

GetUpcOk returns a tuple with the Upc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpc

`func (o *LineItem) SetUpc(v string)`

SetUpc sets Upc field to given value.

### HasUpc

`func (o *LineItem) HasUpc() bool`

HasUpc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


