# AdditionalDataRisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RiskdataCustomFieldName** | Pointer to **string** | The data for your custom risk field. For more information, refer to [Create custom risk fields](https://docs.adyen.com/risk-management/configure-custom-risk-rules#step-1-create-custom-risk-fields). | [optional] 
**RiskdataBasketItemItemNrAmountPerItem** | Pointer to **string** | The price of item in the basket, represented in [minor units](https://docs.adyen.com/development-resources/currency-codes). | [optional] 
**RiskdataBasketItemItemNrBrand** | Pointer to **string** | Brand of the item. | [optional] 
**RiskdataBasketItemItemNrCategory** | Pointer to **string** | Category of the item. | [optional] 
**RiskdataBasketItemItemNrColor** | Pointer to **string** | Color of the item. | [optional] 
**RiskdataBasketItemItemNrCurrency** | Pointer to **string** | The three-character [ISO currency code](https://en.wikipedia.org/wiki/ISO_4217). | [optional] 
**RiskdataBasketItemItemNrItemID** | Pointer to **string** | ID of the item. | [optional] 
**RiskdataBasketItemItemNrManufacturer** | Pointer to **string** | Manufacturer of the item. | [optional] 
**RiskdataBasketItemItemNrProductTitle** | Pointer to **string** | A text description of the product the invoice line refers to. | [optional] 
**RiskdataBasketItemItemNrQuantity** | Pointer to **string** | Quantity of the item purchased. | [optional] 
**RiskdataBasketItemItemNrReceiverEmail** | Pointer to **string** | Email associated with the given product in the basket (usually in electronic gift cards). | [optional] 
**RiskdataBasketItemItemNrSize** | Pointer to **string** | Size of the item. | [optional] 
**RiskdataBasketItemItemNrSku** | Pointer to **string** | [Stock keeping unit](https://en.wikipedia.org/wiki/Stock_keeping_unit). | [optional] 
**RiskdataBasketItemItemNrUpc** | Pointer to **string** | [Universal Product Code](https://en.wikipedia.org/wiki/Universal_Product_Code). | [optional] 
**RiskdataPromotionsPromotionItemNrPromotionCode** | Pointer to **string** | Code of the promotion. | [optional] 
**RiskdataPromotionsPromotionItemNrPromotionDiscountAmount** | Pointer to **string** | The discount amount of the promotion, represented in [minor units](https://docs.adyen.com/development-resources/currency-codes). | [optional] 
**RiskdataPromotionsPromotionItemNrPromotionDiscountCurrency** | Pointer to **string** | The three-character [ISO currency code](https://en.wikipedia.org/wiki/ISO_4217). | [optional] 
**RiskdataPromotionsPromotionItemNrPromotionDiscountPercentage** | Pointer to **string** | Promotion&#39;s percentage discount. It is represented in percentage value and there is no need to include the &#39;%&#39; sign.  e.g. for a promotion discount of 30%, the value of the field should be 30. | [optional] 
**RiskdataPromotionsPromotionItemNrPromotionName** | Pointer to **string** | Name of the promotion. | [optional] 
**RiskdataRiskProfileReference** | Pointer to **string** | Reference number of the risk profile that you want to apply to the payment. If not provided or left blank, the merchant-level account&#39;s default risk profile will be applied to the payment. For more information, see [dynamically assign a risk profile to a payment](https://docs.adyen.com/risk-management/create-and-use-risk-profiles#dynamically-assign-a-risk-profile-to-a-payment). | [optional] 
**RiskdataSkipRisk** | Pointer to **string** | If this parameter is provided with the value **true**, risk checks for the payment request are skipped and the transaction will not get a risk score. | [optional] 

## Methods

### NewAdditionalDataRisk

`func NewAdditionalDataRisk() *AdditionalDataRisk`

NewAdditionalDataRisk instantiates a new AdditionalDataRisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalDataRiskWithDefaults

`func NewAdditionalDataRiskWithDefaults() *AdditionalDataRisk`

NewAdditionalDataRiskWithDefaults instantiates a new AdditionalDataRisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRiskdataCustomFieldName

`func (o *AdditionalDataRisk) GetRiskdataCustomFieldName() string`

GetRiskdataCustomFieldName returns the RiskdataCustomFieldName field if non-nil, zero value otherwise.

### GetRiskdataCustomFieldNameOk

`func (o *AdditionalDataRisk) GetRiskdataCustomFieldNameOk() (*string, bool)`

GetRiskdataCustomFieldNameOk returns a tuple with the RiskdataCustomFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskdataCustomFieldName

`func (o *AdditionalDataRisk) SetRiskdataCustomFieldName(v string)`

SetRiskdataCustomFieldName sets RiskdataCustomFieldName field to given value.

### HasRiskdataCustomFieldName

`func (o *AdditionalDataRisk) HasRiskdataCustomFieldName() bool`

HasRiskdataCustomFieldName returns a boolean if a field has been set.

### GetRiskdataBasketItemItemNrAmountPerItem

`func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrAmountPerItem() string`

GetRiskdataBasketItemItemNrAmountPerItem returns the RiskdataBasketItemItemNrAmountPerItem field if non-nil, zero value otherwise.

### GetRiskdataBasketItemItemNrAmountPerItemOk

`func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrAmountPerItemOk() (*string, bool)`

GetRiskdataBasketItemItemNrAmountPerItemOk returns a tuple with the RiskdataBasketItemItemNrAmountPerItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskdataBasketItemItemNrAmountPerItem

`func (o *AdditionalDataRisk) SetRiskdataBasketItemItemNrAmountPerItem(v string)`

SetRiskdataBasketItemItemNrAmountPerItem sets RiskdataBasketItemItemNrAmountPerItem field to given value.

### HasRiskdataBasketItemItemNrAmountPerItem

`func (o *AdditionalDataRisk) HasRiskdataBasketItemItemNrAmountPerItem() bool`

HasRiskdataBasketItemItemNrAmountPerItem returns a boolean if a field has been set.

### GetRiskdataBasketItemItemNrBrand

`func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrBrand() string`

GetRiskdataBasketItemItemNrBrand returns the RiskdataBasketItemItemNrBrand field if non-nil, zero value otherwise.

### GetRiskdataBasketItemItemNrBrandOk

`func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrBrandOk() (*string, bool)`

GetRiskdataBasketItemItemNrBrandOk returns a tuple with the RiskdataBasketItemItemNrBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskdataBasketItemItemNrBrand

`func (o *AdditionalDataRisk) SetRiskdataBasketItemItemNrBrand(v string)`

SetRiskdataBasketItemItemNrBrand sets RiskdataBasketItemItemNrBrand field to given value.

### HasRiskdataBasketItemItemNrBrand

`func (o *AdditionalDataRisk) HasRiskdataBasketItemItemNrBrand() bool`

HasRiskdataBasketItemItemNrBrand returns a boolean if a field has been set.

### GetRiskdataBasketItemItemNrCategory

`func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrCategory() string`

GetRiskdataBasketItemItemNrCategory returns the RiskdataBasketItemItemNrCategory field if non-nil, zero value otherwise.

### GetRiskdataBasketItemItemNrCategoryOk

`func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrCategoryOk() (*string, bool)`

GetRiskdataBasketItemItemNrCategoryOk returns a tuple with the RiskdataBasketItemItemNrCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskdataBasketItemItemNrCategory

`func (o *AdditionalDataRisk) SetRiskdataBasketItemItemNrCategory(v string)`

SetRiskdataBasketItemItemNrCategory sets RiskdataBasketItemItemNrCategory field to given value.

### HasRiskdataBasketItemItemNrCategory

`func (o *AdditionalDataRisk) HasRiskdataBasketItemItemNrCategory() bool`

HasRiskdataBasketItemItemNrCategory returns a boolean if a field has been set.

### GetRiskdataBasketItemItemNrColor

`func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrColor() string`

GetRiskdataBasketItemItemNrColor returns the RiskdataBasketItemItemNrColor field if non-nil, zero value otherwise.

### GetRiskdataBasketItemItemNrColorOk

`func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrColorOk() (*string, bool)`

GetRiskdataBasketItemItemNrColorOk returns a tuple with the RiskdataBasketItemItemNrColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskdataBasketItemItemNrColor

`func (o *AdditionalDataRisk) SetRiskdataBasketItemItemNrColor(v string)`

SetRiskdataBasketItemItemNrColor sets RiskdataBasketItemItemNrColor field to given value.

### HasRiskdataBasketItemItemNrColor

`func (o *AdditionalDataRisk) HasRiskdataBasketItemItemNrColor() bool`

HasRiskdataBasketItemItemNrColor returns a boolean if a field has been set.

### GetRiskdataBasketItemItemNrCurrency

`func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrCurrency() string`

GetRiskdataBasketItemItemNrCurrency returns the RiskdataBasketItemItemNrCurrency field if non-nil, zero value otherwise.

### GetRiskdataBasketItemItemNrCurrencyOk

`func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrCurrencyOk() (*string, bool)`

GetRiskdataBasketItemItemNrCurrencyOk returns a tuple with the RiskdataBasketItemItemNrCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskdataBasketItemItemNrCurrency

`func (o *AdditionalDataRisk) SetRiskdataBasketItemItemNrCurrency(v string)`

SetRiskdataBasketItemItemNrCurrency sets RiskdataBasketItemItemNrCurrency field to given value.

### HasRiskdataBasketItemItemNrCurrency

`func (o *AdditionalDataRisk) HasRiskdataBasketItemItemNrCurrency() bool`

HasRiskdataBasketItemItemNrCurrency returns a boolean if a field has been set.

### GetRiskdataBasketItemItemNrItemID

`func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrItemID() string`

GetRiskdataBasketItemItemNrItemID returns the RiskdataBasketItemItemNrItemID field if non-nil, zero value otherwise.

### GetRiskdataBasketItemItemNrItemIDOk

`func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrItemIDOk() (*string, bool)`

GetRiskdataBasketItemItemNrItemIDOk returns a tuple with the RiskdataBasketItemItemNrItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskdataBasketItemItemNrItemID

`func (o *AdditionalDataRisk) SetRiskdataBasketItemItemNrItemID(v string)`

SetRiskdataBasketItemItemNrItemID sets RiskdataBasketItemItemNrItemID field to given value.

### HasRiskdataBasketItemItemNrItemID

`func (o *AdditionalDataRisk) HasRiskdataBasketItemItemNrItemID() bool`

HasRiskdataBasketItemItemNrItemID returns a boolean if a field has been set.

### GetRiskdataBasketItemItemNrManufacturer

`func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrManufacturer() string`

GetRiskdataBasketItemItemNrManufacturer returns the RiskdataBasketItemItemNrManufacturer field if non-nil, zero value otherwise.

### GetRiskdataBasketItemItemNrManufacturerOk

`func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrManufacturerOk() (*string, bool)`

GetRiskdataBasketItemItemNrManufacturerOk returns a tuple with the RiskdataBasketItemItemNrManufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskdataBasketItemItemNrManufacturer

`func (o *AdditionalDataRisk) SetRiskdataBasketItemItemNrManufacturer(v string)`

SetRiskdataBasketItemItemNrManufacturer sets RiskdataBasketItemItemNrManufacturer field to given value.

### HasRiskdataBasketItemItemNrManufacturer

`func (o *AdditionalDataRisk) HasRiskdataBasketItemItemNrManufacturer() bool`

HasRiskdataBasketItemItemNrManufacturer returns a boolean if a field has been set.

### GetRiskdataBasketItemItemNrProductTitle

`func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrProductTitle() string`

GetRiskdataBasketItemItemNrProductTitle returns the RiskdataBasketItemItemNrProductTitle field if non-nil, zero value otherwise.

### GetRiskdataBasketItemItemNrProductTitleOk

`func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrProductTitleOk() (*string, bool)`

GetRiskdataBasketItemItemNrProductTitleOk returns a tuple with the RiskdataBasketItemItemNrProductTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskdataBasketItemItemNrProductTitle

`func (o *AdditionalDataRisk) SetRiskdataBasketItemItemNrProductTitle(v string)`

SetRiskdataBasketItemItemNrProductTitle sets RiskdataBasketItemItemNrProductTitle field to given value.

### HasRiskdataBasketItemItemNrProductTitle

`func (o *AdditionalDataRisk) HasRiskdataBasketItemItemNrProductTitle() bool`

HasRiskdataBasketItemItemNrProductTitle returns a boolean if a field has been set.

### GetRiskdataBasketItemItemNrQuantity

`func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrQuantity() string`

GetRiskdataBasketItemItemNrQuantity returns the RiskdataBasketItemItemNrQuantity field if non-nil, zero value otherwise.

### GetRiskdataBasketItemItemNrQuantityOk

`func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrQuantityOk() (*string, bool)`

GetRiskdataBasketItemItemNrQuantityOk returns a tuple with the RiskdataBasketItemItemNrQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskdataBasketItemItemNrQuantity

`func (o *AdditionalDataRisk) SetRiskdataBasketItemItemNrQuantity(v string)`

SetRiskdataBasketItemItemNrQuantity sets RiskdataBasketItemItemNrQuantity field to given value.

### HasRiskdataBasketItemItemNrQuantity

`func (o *AdditionalDataRisk) HasRiskdataBasketItemItemNrQuantity() bool`

HasRiskdataBasketItemItemNrQuantity returns a boolean if a field has been set.

### GetRiskdataBasketItemItemNrReceiverEmail

`func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrReceiverEmail() string`

GetRiskdataBasketItemItemNrReceiverEmail returns the RiskdataBasketItemItemNrReceiverEmail field if non-nil, zero value otherwise.

### GetRiskdataBasketItemItemNrReceiverEmailOk

`func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrReceiverEmailOk() (*string, bool)`

GetRiskdataBasketItemItemNrReceiverEmailOk returns a tuple with the RiskdataBasketItemItemNrReceiverEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskdataBasketItemItemNrReceiverEmail

`func (o *AdditionalDataRisk) SetRiskdataBasketItemItemNrReceiverEmail(v string)`

SetRiskdataBasketItemItemNrReceiverEmail sets RiskdataBasketItemItemNrReceiverEmail field to given value.

### HasRiskdataBasketItemItemNrReceiverEmail

`func (o *AdditionalDataRisk) HasRiskdataBasketItemItemNrReceiverEmail() bool`

HasRiskdataBasketItemItemNrReceiverEmail returns a boolean if a field has been set.

### GetRiskdataBasketItemItemNrSize

`func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrSize() string`

GetRiskdataBasketItemItemNrSize returns the RiskdataBasketItemItemNrSize field if non-nil, zero value otherwise.

### GetRiskdataBasketItemItemNrSizeOk

`func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrSizeOk() (*string, bool)`

GetRiskdataBasketItemItemNrSizeOk returns a tuple with the RiskdataBasketItemItemNrSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskdataBasketItemItemNrSize

`func (o *AdditionalDataRisk) SetRiskdataBasketItemItemNrSize(v string)`

SetRiskdataBasketItemItemNrSize sets RiskdataBasketItemItemNrSize field to given value.

### HasRiskdataBasketItemItemNrSize

`func (o *AdditionalDataRisk) HasRiskdataBasketItemItemNrSize() bool`

HasRiskdataBasketItemItemNrSize returns a boolean if a field has been set.

### GetRiskdataBasketItemItemNrSku

`func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrSku() string`

GetRiskdataBasketItemItemNrSku returns the RiskdataBasketItemItemNrSku field if non-nil, zero value otherwise.

### GetRiskdataBasketItemItemNrSkuOk

`func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrSkuOk() (*string, bool)`

GetRiskdataBasketItemItemNrSkuOk returns a tuple with the RiskdataBasketItemItemNrSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskdataBasketItemItemNrSku

`func (o *AdditionalDataRisk) SetRiskdataBasketItemItemNrSku(v string)`

SetRiskdataBasketItemItemNrSku sets RiskdataBasketItemItemNrSku field to given value.

### HasRiskdataBasketItemItemNrSku

`func (o *AdditionalDataRisk) HasRiskdataBasketItemItemNrSku() bool`

HasRiskdataBasketItemItemNrSku returns a boolean if a field has been set.

### GetRiskdataBasketItemItemNrUpc

`func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrUpc() string`

GetRiskdataBasketItemItemNrUpc returns the RiskdataBasketItemItemNrUpc field if non-nil, zero value otherwise.

### GetRiskdataBasketItemItemNrUpcOk

`func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrUpcOk() (*string, bool)`

GetRiskdataBasketItemItemNrUpcOk returns a tuple with the RiskdataBasketItemItemNrUpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskdataBasketItemItemNrUpc

`func (o *AdditionalDataRisk) SetRiskdataBasketItemItemNrUpc(v string)`

SetRiskdataBasketItemItemNrUpc sets RiskdataBasketItemItemNrUpc field to given value.

### HasRiskdataBasketItemItemNrUpc

`func (o *AdditionalDataRisk) HasRiskdataBasketItemItemNrUpc() bool`

HasRiskdataBasketItemItemNrUpc returns a boolean if a field has been set.

### GetRiskdataPromotionsPromotionItemNrPromotionCode

`func (o *AdditionalDataRisk) GetRiskdataPromotionsPromotionItemNrPromotionCode() string`

GetRiskdataPromotionsPromotionItemNrPromotionCode returns the RiskdataPromotionsPromotionItemNrPromotionCode field if non-nil, zero value otherwise.

### GetRiskdataPromotionsPromotionItemNrPromotionCodeOk

`func (o *AdditionalDataRisk) GetRiskdataPromotionsPromotionItemNrPromotionCodeOk() (*string, bool)`

GetRiskdataPromotionsPromotionItemNrPromotionCodeOk returns a tuple with the RiskdataPromotionsPromotionItemNrPromotionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskdataPromotionsPromotionItemNrPromotionCode

`func (o *AdditionalDataRisk) SetRiskdataPromotionsPromotionItemNrPromotionCode(v string)`

SetRiskdataPromotionsPromotionItemNrPromotionCode sets RiskdataPromotionsPromotionItemNrPromotionCode field to given value.

### HasRiskdataPromotionsPromotionItemNrPromotionCode

`func (o *AdditionalDataRisk) HasRiskdataPromotionsPromotionItemNrPromotionCode() bool`

HasRiskdataPromotionsPromotionItemNrPromotionCode returns a boolean if a field has been set.

### GetRiskdataPromotionsPromotionItemNrPromotionDiscountAmount

`func (o *AdditionalDataRisk) GetRiskdataPromotionsPromotionItemNrPromotionDiscountAmount() string`

GetRiskdataPromotionsPromotionItemNrPromotionDiscountAmount returns the RiskdataPromotionsPromotionItemNrPromotionDiscountAmount field if non-nil, zero value otherwise.

### GetRiskdataPromotionsPromotionItemNrPromotionDiscountAmountOk

`func (o *AdditionalDataRisk) GetRiskdataPromotionsPromotionItemNrPromotionDiscountAmountOk() (*string, bool)`

GetRiskdataPromotionsPromotionItemNrPromotionDiscountAmountOk returns a tuple with the RiskdataPromotionsPromotionItemNrPromotionDiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskdataPromotionsPromotionItemNrPromotionDiscountAmount

`func (o *AdditionalDataRisk) SetRiskdataPromotionsPromotionItemNrPromotionDiscountAmount(v string)`

SetRiskdataPromotionsPromotionItemNrPromotionDiscountAmount sets RiskdataPromotionsPromotionItemNrPromotionDiscountAmount field to given value.

### HasRiskdataPromotionsPromotionItemNrPromotionDiscountAmount

`func (o *AdditionalDataRisk) HasRiskdataPromotionsPromotionItemNrPromotionDiscountAmount() bool`

HasRiskdataPromotionsPromotionItemNrPromotionDiscountAmount returns a boolean if a field has been set.

### GetRiskdataPromotionsPromotionItemNrPromotionDiscountCurrency

`func (o *AdditionalDataRisk) GetRiskdataPromotionsPromotionItemNrPromotionDiscountCurrency() string`

GetRiskdataPromotionsPromotionItemNrPromotionDiscountCurrency returns the RiskdataPromotionsPromotionItemNrPromotionDiscountCurrency field if non-nil, zero value otherwise.

### GetRiskdataPromotionsPromotionItemNrPromotionDiscountCurrencyOk

`func (o *AdditionalDataRisk) GetRiskdataPromotionsPromotionItemNrPromotionDiscountCurrencyOk() (*string, bool)`

GetRiskdataPromotionsPromotionItemNrPromotionDiscountCurrencyOk returns a tuple with the RiskdataPromotionsPromotionItemNrPromotionDiscountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskdataPromotionsPromotionItemNrPromotionDiscountCurrency

`func (o *AdditionalDataRisk) SetRiskdataPromotionsPromotionItemNrPromotionDiscountCurrency(v string)`

SetRiskdataPromotionsPromotionItemNrPromotionDiscountCurrency sets RiskdataPromotionsPromotionItemNrPromotionDiscountCurrency field to given value.

### HasRiskdataPromotionsPromotionItemNrPromotionDiscountCurrency

`func (o *AdditionalDataRisk) HasRiskdataPromotionsPromotionItemNrPromotionDiscountCurrency() bool`

HasRiskdataPromotionsPromotionItemNrPromotionDiscountCurrency returns a boolean if a field has been set.

### GetRiskdataPromotionsPromotionItemNrPromotionDiscountPercentage

`func (o *AdditionalDataRisk) GetRiskdataPromotionsPromotionItemNrPromotionDiscountPercentage() string`

GetRiskdataPromotionsPromotionItemNrPromotionDiscountPercentage returns the RiskdataPromotionsPromotionItemNrPromotionDiscountPercentage field if non-nil, zero value otherwise.

### GetRiskdataPromotionsPromotionItemNrPromotionDiscountPercentageOk

`func (o *AdditionalDataRisk) GetRiskdataPromotionsPromotionItemNrPromotionDiscountPercentageOk() (*string, bool)`

GetRiskdataPromotionsPromotionItemNrPromotionDiscountPercentageOk returns a tuple with the RiskdataPromotionsPromotionItemNrPromotionDiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskdataPromotionsPromotionItemNrPromotionDiscountPercentage

`func (o *AdditionalDataRisk) SetRiskdataPromotionsPromotionItemNrPromotionDiscountPercentage(v string)`

SetRiskdataPromotionsPromotionItemNrPromotionDiscountPercentage sets RiskdataPromotionsPromotionItemNrPromotionDiscountPercentage field to given value.

### HasRiskdataPromotionsPromotionItemNrPromotionDiscountPercentage

`func (o *AdditionalDataRisk) HasRiskdataPromotionsPromotionItemNrPromotionDiscountPercentage() bool`

HasRiskdataPromotionsPromotionItemNrPromotionDiscountPercentage returns a boolean if a field has been set.

### GetRiskdataPromotionsPromotionItemNrPromotionName

`func (o *AdditionalDataRisk) GetRiskdataPromotionsPromotionItemNrPromotionName() string`

GetRiskdataPromotionsPromotionItemNrPromotionName returns the RiskdataPromotionsPromotionItemNrPromotionName field if non-nil, zero value otherwise.

### GetRiskdataPromotionsPromotionItemNrPromotionNameOk

`func (o *AdditionalDataRisk) GetRiskdataPromotionsPromotionItemNrPromotionNameOk() (*string, bool)`

GetRiskdataPromotionsPromotionItemNrPromotionNameOk returns a tuple with the RiskdataPromotionsPromotionItemNrPromotionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskdataPromotionsPromotionItemNrPromotionName

`func (o *AdditionalDataRisk) SetRiskdataPromotionsPromotionItemNrPromotionName(v string)`

SetRiskdataPromotionsPromotionItemNrPromotionName sets RiskdataPromotionsPromotionItemNrPromotionName field to given value.

### HasRiskdataPromotionsPromotionItemNrPromotionName

`func (o *AdditionalDataRisk) HasRiskdataPromotionsPromotionItemNrPromotionName() bool`

HasRiskdataPromotionsPromotionItemNrPromotionName returns a boolean if a field has been set.

### GetRiskdataRiskProfileReference

`func (o *AdditionalDataRisk) GetRiskdataRiskProfileReference() string`

GetRiskdataRiskProfileReference returns the RiskdataRiskProfileReference field if non-nil, zero value otherwise.

### GetRiskdataRiskProfileReferenceOk

`func (o *AdditionalDataRisk) GetRiskdataRiskProfileReferenceOk() (*string, bool)`

GetRiskdataRiskProfileReferenceOk returns a tuple with the RiskdataRiskProfileReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskdataRiskProfileReference

`func (o *AdditionalDataRisk) SetRiskdataRiskProfileReference(v string)`

SetRiskdataRiskProfileReference sets RiskdataRiskProfileReference field to given value.

### HasRiskdataRiskProfileReference

`func (o *AdditionalDataRisk) HasRiskdataRiskProfileReference() bool`

HasRiskdataRiskProfileReference returns a boolean if a field has been set.

### GetRiskdataSkipRisk

`func (o *AdditionalDataRisk) GetRiskdataSkipRisk() string`

GetRiskdataSkipRisk returns the RiskdataSkipRisk field if non-nil, zero value otherwise.

### GetRiskdataSkipRiskOk

`func (o *AdditionalDataRisk) GetRiskdataSkipRiskOk() (*string, bool)`

GetRiskdataSkipRiskOk returns a tuple with the RiskdataSkipRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskdataSkipRisk

`func (o *AdditionalDataRisk) SetRiskdataSkipRisk(v string)`

SetRiskdataSkipRisk sets RiskdataSkipRisk field to given value.

### HasRiskdataSkipRisk

`func (o *AdditionalDataRisk) HasRiskdataSkipRisk() bool`

HasRiskdataSkipRisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


