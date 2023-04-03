# MerchantRiskIndicator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressMatch** | Pointer to **bool** | Whether the chosen delivery address is identical to the billing address. | [optional] 
**DeliveryAddressIndicator** | Pointer to **string** | Indicator regarding the delivery address. Allowed values: * &#x60;shipToBillingAddress&#x60; * &#x60;shipToVerifiedAddress&#x60; * &#x60;shipToNewAddress&#x60; * &#x60;shipToStore&#x60; * &#x60;digitalGoods&#x60; * &#x60;goodsNotShipped&#x60; * &#x60;other&#x60; | [optional] 
**DeliveryEmail** | Pointer to **string** | The delivery email address (for digital goods). | [optional] 
**DeliveryEmailAddress** | Pointer to **string** | For Electronic delivery, the email address to which the merchandise was delivered. Maximum length: 254 characters. | [optional] 
**DeliveryTimeframe** | Pointer to **string** | The estimated delivery time for the shopper to receive the goods. Allowed values: * &#x60;electronicDelivery&#x60; * &#x60;sameDayShipping&#x60; * &#x60;overnightShipping&#x60; * &#x60;twoOrMoreDaysShipping&#x60; | [optional] 
**GiftCardAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**GiftCardCount** | Pointer to **int32** | For prepaid or gift card purchase, total count of individual prepaid or gift cards/codes purchased. | [optional] 
**GiftCardCurr** | Pointer to **string** | For prepaid or gift card purchase, [ISO 4217](https://www.iso.org/iso-4217-currency-codes.html) three-digit currency code of the gift card, other than those listed in Table A.5 of the EMVCo 3D Secure Protocol and Core Functions Specification. | [optional] 
**PreOrderDate** | Pointer to **time.Time** | For pre-order purchases, the expected date this product will be available to the shopper. | [optional] 
**PreOrderPurchase** | Pointer to **bool** | Indicator for whether this transaction is for pre-ordering a product. | [optional] 
**PreOrderPurchaseInd** | Pointer to **string** | Indicates whether Cardholder is placing an order for merchandise with a future availability or release date. | [optional] 
**ReorderItems** | Pointer to **bool** | Indicator for whether the shopper has already purchased the same items in the past. | [optional] 
**ReorderItemsInd** | Pointer to **string** | Indicates whether the cardholder is reordering previously purchased merchandise. | [optional] 
**ShipIndicator** | Pointer to **string** | Indicates shipping method chosen for the transaction. | [optional] 

## Methods

### NewMerchantRiskIndicator

`func NewMerchantRiskIndicator() *MerchantRiskIndicator`

NewMerchantRiskIndicator instantiates a new MerchantRiskIndicator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantRiskIndicatorWithDefaults

`func NewMerchantRiskIndicatorWithDefaults() *MerchantRiskIndicator`

NewMerchantRiskIndicatorWithDefaults instantiates a new MerchantRiskIndicator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressMatch

`func (o *MerchantRiskIndicator) GetAddressMatch() bool`

GetAddressMatch returns the AddressMatch field if non-nil, zero value otherwise.

### GetAddressMatchOk

`func (o *MerchantRiskIndicator) GetAddressMatchOk() (*bool, bool)`

GetAddressMatchOk returns a tuple with the AddressMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressMatch

`func (o *MerchantRiskIndicator) SetAddressMatch(v bool)`

SetAddressMatch sets AddressMatch field to given value.

### HasAddressMatch

`func (o *MerchantRiskIndicator) HasAddressMatch() bool`

HasAddressMatch returns a boolean if a field has been set.

### GetDeliveryAddressIndicator

`func (o *MerchantRiskIndicator) GetDeliveryAddressIndicator() string`

GetDeliveryAddressIndicator returns the DeliveryAddressIndicator field if non-nil, zero value otherwise.

### GetDeliveryAddressIndicatorOk

`func (o *MerchantRiskIndicator) GetDeliveryAddressIndicatorOk() (*string, bool)`

GetDeliveryAddressIndicatorOk returns a tuple with the DeliveryAddressIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddressIndicator

`func (o *MerchantRiskIndicator) SetDeliveryAddressIndicator(v string)`

SetDeliveryAddressIndicator sets DeliveryAddressIndicator field to given value.

### HasDeliveryAddressIndicator

`func (o *MerchantRiskIndicator) HasDeliveryAddressIndicator() bool`

HasDeliveryAddressIndicator returns a boolean if a field has been set.

### GetDeliveryEmail

`func (o *MerchantRiskIndicator) GetDeliveryEmail() string`

GetDeliveryEmail returns the DeliveryEmail field if non-nil, zero value otherwise.

### GetDeliveryEmailOk

`func (o *MerchantRiskIndicator) GetDeliveryEmailOk() (*string, bool)`

GetDeliveryEmailOk returns a tuple with the DeliveryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryEmail

`func (o *MerchantRiskIndicator) SetDeliveryEmail(v string)`

SetDeliveryEmail sets DeliveryEmail field to given value.

### HasDeliveryEmail

`func (o *MerchantRiskIndicator) HasDeliveryEmail() bool`

HasDeliveryEmail returns a boolean if a field has been set.

### GetDeliveryEmailAddress

`func (o *MerchantRiskIndicator) GetDeliveryEmailAddress() string`

GetDeliveryEmailAddress returns the DeliveryEmailAddress field if non-nil, zero value otherwise.

### GetDeliveryEmailAddressOk

`func (o *MerchantRiskIndicator) GetDeliveryEmailAddressOk() (*string, bool)`

GetDeliveryEmailAddressOk returns a tuple with the DeliveryEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryEmailAddress

`func (o *MerchantRiskIndicator) SetDeliveryEmailAddress(v string)`

SetDeliveryEmailAddress sets DeliveryEmailAddress field to given value.

### HasDeliveryEmailAddress

`func (o *MerchantRiskIndicator) HasDeliveryEmailAddress() bool`

HasDeliveryEmailAddress returns a boolean if a field has been set.

### GetDeliveryTimeframe

`func (o *MerchantRiskIndicator) GetDeliveryTimeframe() string`

GetDeliveryTimeframe returns the DeliveryTimeframe field if non-nil, zero value otherwise.

### GetDeliveryTimeframeOk

`func (o *MerchantRiskIndicator) GetDeliveryTimeframeOk() (*string, bool)`

GetDeliveryTimeframeOk returns a tuple with the DeliveryTimeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTimeframe

`func (o *MerchantRiskIndicator) SetDeliveryTimeframe(v string)`

SetDeliveryTimeframe sets DeliveryTimeframe field to given value.

### HasDeliveryTimeframe

`func (o *MerchantRiskIndicator) HasDeliveryTimeframe() bool`

HasDeliveryTimeframe returns a boolean if a field has been set.

### GetGiftCardAmount

`func (o *MerchantRiskIndicator) GetGiftCardAmount() Amount`

GetGiftCardAmount returns the GiftCardAmount field if non-nil, zero value otherwise.

### GetGiftCardAmountOk

`func (o *MerchantRiskIndicator) GetGiftCardAmountOk() (*Amount, bool)`

GetGiftCardAmountOk returns a tuple with the GiftCardAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardAmount

`func (o *MerchantRiskIndicator) SetGiftCardAmount(v Amount)`

SetGiftCardAmount sets GiftCardAmount field to given value.

### HasGiftCardAmount

`func (o *MerchantRiskIndicator) HasGiftCardAmount() bool`

HasGiftCardAmount returns a boolean if a field has been set.

### GetGiftCardCount

`func (o *MerchantRiskIndicator) GetGiftCardCount() int32`

GetGiftCardCount returns the GiftCardCount field if non-nil, zero value otherwise.

### GetGiftCardCountOk

`func (o *MerchantRiskIndicator) GetGiftCardCountOk() (*int32, bool)`

GetGiftCardCountOk returns a tuple with the GiftCardCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardCount

`func (o *MerchantRiskIndicator) SetGiftCardCount(v int32)`

SetGiftCardCount sets GiftCardCount field to given value.

### HasGiftCardCount

`func (o *MerchantRiskIndicator) HasGiftCardCount() bool`

HasGiftCardCount returns a boolean if a field has been set.

### GetGiftCardCurr

`func (o *MerchantRiskIndicator) GetGiftCardCurr() string`

GetGiftCardCurr returns the GiftCardCurr field if non-nil, zero value otherwise.

### GetGiftCardCurrOk

`func (o *MerchantRiskIndicator) GetGiftCardCurrOk() (*string, bool)`

GetGiftCardCurrOk returns a tuple with the GiftCardCurr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardCurr

`func (o *MerchantRiskIndicator) SetGiftCardCurr(v string)`

SetGiftCardCurr sets GiftCardCurr field to given value.

### HasGiftCardCurr

`func (o *MerchantRiskIndicator) HasGiftCardCurr() bool`

HasGiftCardCurr returns a boolean if a field has been set.

### GetPreOrderDate

`func (o *MerchantRiskIndicator) GetPreOrderDate() time.Time`

GetPreOrderDate returns the PreOrderDate field if non-nil, zero value otherwise.

### GetPreOrderDateOk

`func (o *MerchantRiskIndicator) GetPreOrderDateOk() (*time.Time, bool)`

GetPreOrderDateOk returns a tuple with the PreOrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreOrderDate

`func (o *MerchantRiskIndicator) SetPreOrderDate(v time.Time)`

SetPreOrderDate sets PreOrderDate field to given value.

### HasPreOrderDate

`func (o *MerchantRiskIndicator) HasPreOrderDate() bool`

HasPreOrderDate returns a boolean if a field has been set.

### GetPreOrderPurchase

`func (o *MerchantRiskIndicator) GetPreOrderPurchase() bool`

GetPreOrderPurchase returns the PreOrderPurchase field if non-nil, zero value otherwise.

### GetPreOrderPurchaseOk

`func (o *MerchantRiskIndicator) GetPreOrderPurchaseOk() (*bool, bool)`

GetPreOrderPurchaseOk returns a tuple with the PreOrderPurchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreOrderPurchase

`func (o *MerchantRiskIndicator) SetPreOrderPurchase(v bool)`

SetPreOrderPurchase sets PreOrderPurchase field to given value.

### HasPreOrderPurchase

`func (o *MerchantRiskIndicator) HasPreOrderPurchase() bool`

HasPreOrderPurchase returns a boolean if a field has been set.

### GetPreOrderPurchaseInd

`func (o *MerchantRiskIndicator) GetPreOrderPurchaseInd() string`

GetPreOrderPurchaseInd returns the PreOrderPurchaseInd field if non-nil, zero value otherwise.

### GetPreOrderPurchaseIndOk

`func (o *MerchantRiskIndicator) GetPreOrderPurchaseIndOk() (*string, bool)`

GetPreOrderPurchaseIndOk returns a tuple with the PreOrderPurchaseInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreOrderPurchaseInd

`func (o *MerchantRiskIndicator) SetPreOrderPurchaseInd(v string)`

SetPreOrderPurchaseInd sets PreOrderPurchaseInd field to given value.

### HasPreOrderPurchaseInd

`func (o *MerchantRiskIndicator) HasPreOrderPurchaseInd() bool`

HasPreOrderPurchaseInd returns a boolean if a field has been set.

### GetReorderItems

`func (o *MerchantRiskIndicator) GetReorderItems() bool`

GetReorderItems returns the ReorderItems field if non-nil, zero value otherwise.

### GetReorderItemsOk

`func (o *MerchantRiskIndicator) GetReorderItemsOk() (*bool, bool)`

GetReorderItemsOk returns a tuple with the ReorderItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReorderItems

`func (o *MerchantRiskIndicator) SetReorderItems(v bool)`

SetReorderItems sets ReorderItems field to given value.

### HasReorderItems

`func (o *MerchantRiskIndicator) HasReorderItems() bool`

HasReorderItems returns a boolean if a field has been set.

### GetReorderItemsInd

`func (o *MerchantRiskIndicator) GetReorderItemsInd() string`

GetReorderItemsInd returns the ReorderItemsInd field if non-nil, zero value otherwise.

### GetReorderItemsIndOk

`func (o *MerchantRiskIndicator) GetReorderItemsIndOk() (*string, bool)`

GetReorderItemsIndOk returns a tuple with the ReorderItemsInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReorderItemsInd

`func (o *MerchantRiskIndicator) SetReorderItemsInd(v string)`

SetReorderItemsInd sets ReorderItemsInd field to given value.

### HasReorderItemsInd

`func (o *MerchantRiskIndicator) HasReorderItemsInd() bool`

HasReorderItemsInd returns a boolean if a field has been set.

### GetShipIndicator

`func (o *MerchantRiskIndicator) GetShipIndicator() string`

GetShipIndicator returns the ShipIndicator field if non-nil, zero value otherwise.

### GetShipIndicatorOk

`func (o *MerchantRiskIndicator) GetShipIndicatorOk() (*string, bool)`

GetShipIndicatorOk returns a tuple with the ShipIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipIndicator

`func (o *MerchantRiskIndicator) SetShipIndicator(v string)`

SetShipIndicator sets ShipIndicator field to given value.

### HasShipIndicator

`func (o *MerchantRiskIndicator) HasShipIndicator() bool`

HasShipIndicator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


