# AdditionalDataOpenInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpeninvoicedataMerchantData** | Pointer to **string** | Holds different merchant data points like product, purchase, customer, and so on. It takes data in a Base64 encoded string.  The &#x60;merchantData&#x60; parameter needs to be added to the &#x60;openinvoicedata&#x60; signature at the end.  Since the field is optional, if it&#39;s not included it does not impact computing the merchant signature.  Applies only to Klarna.  You can contact Klarna for the format and structure of the string. | [optional] 
**OpeninvoicedataNumberOfLines** | Pointer to **string** | The number of invoice lines included in &#x60;openinvoicedata&#x60;.  There needs to be at least one line, so &#x60;numberOfLines&#x60; needs to be at least 1. | [optional] 
**OpeninvoicedataRecipientFirstName** | Pointer to **string** | First name of the recipient. If the delivery address and the billing address are different, specify the &#x60;recipientFirstName&#x60; and &#x60;recipientLastName&#x60; to share the delivery address with Klarna. Otherwise, only the billing address is shared with Klarna. | [optional] 
**OpeninvoicedataRecipientLastName** | Pointer to **string** | Last name of the recipient. If the delivery address and the billing address are different, specify the &#x60;recipientFirstName&#x60; and &#x60;recipientLastName&#x60; to share the delivery address with Klarna. Otherwise, only the billing address is shared with Klarna. | [optional] 
**OpeninvoicedataLineItemNrCurrencyCode** | Pointer to **string** | The three-character ISO currency code. | [optional] 
**OpeninvoicedataLineItemNrDescription** | Pointer to **string** | A text description of the product the invoice line refers to. | [optional] 
**OpeninvoicedataLineItemNrItemAmount** | Pointer to **string** | The price for one item in the invoice line, represented in minor units.  The due amount for the item, VAT excluded. | [optional] 
**OpeninvoicedataLineItemNrItemId** | Pointer to **string** | A unique id for this item. Required for RatePay if the description of each item is not unique. | [optional] 
**OpeninvoicedataLineItemNrItemVatAmount** | Pointer to **string** | The VAT due for one item in the invoice line, represented in minor units. | [optional] 
**OpeninvoicedataLineItemNrItemVatPercentage** | Pointer to **string** | The VAT percentage for one item in the invoice line, represented in minor units.  For example, 19% VAT is specified as 1900. | [optional] 
**OpeninvoicedataLineItemNrNumberOfItems** | Pointer to **string** | The number of units purchased of a specific product. | [optional] 
**OpeninvoicedataLineItemNrReturnShippingCompany** | Pointer to **string** | Name of the shipping company handling the the return shipment. | [optional] 
**OpeninvoicedataLineItemNrReturnTrackingNumber** | Pointer to **string** | The tracking number for the return of the shipment. | [optional] 
**OpeninvoicedataLineItemNrReturnTrackingUri** | Pointer to **string** | URI where the customer can track the return of their shipment. | [optional] 
**OpeninvoicedataLineItemNrShippingCompany** | Pointer to **string** | Name of the shipping company handling the delivery. | [optional] 
**OpeninvoicedataLineItemNrShippingMethod** | Pointer to **string** | Shipping method. | [optional] 
**OpeninvoicedataLineItemNrTrackingNumber** | Pointer to **string** | The tracking number for the shipment. | [optional] 
**OpeninvoicedataLineItemNrTrackingUri** | Pointer to **string** | URI where the customer can track their shipment. | [optional] 

## Methods

### NewAdditionalDataOpenInvoice

`func NewAdditionalDataOpenInvoice() *AdditionalDataOpenInvoice`

NewAdditionalDataOpenInvoice instantiates a new AdditionalDataOpenInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalDataOpenInvoiceWithDefaults

`func NewAdditionalDataOpenInvoiceWithDefaults() *AdditionalDataOpenInvoice`

NewAdditionalDataOpenInvoiceWithDefaults instantiates a new AdditionalDataOpenInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpeninvoicedataMerchantData

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataMerchantData() string`

GetOpeninvoicedataMerchantData returns the OpeninvoicedataMerchantData field if non-nil, zero value otherwise.

### GetOpeninvoicedataMerchantDataOk

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataMerchantDataOk() (*string, bool)`

GetOpeninvoicedataMerchantDataOk returns a tuple with the OpeninvoicedataMerchantData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeninvoicedataMerchantData

`func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataMerchantData(v string)`

SetOpeninvoicedataMerchantData sets OpeninvoicedataMerchantData field to given value.

### HasOpeninvoicedataMerchantData

`func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataMerchantData() bool`

HasOpeninvoicedataMerchantData returns a boolean if a field has been set.

### GetOpeninvoicedataNumberOfLines

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataNumberOfLines() string`

GetOpeninvoicedataNumberOfLines returns the OpeninvoicedataNumberOfLines field if non-nil, zero value otherwise.

### GetOpeninvoicedataNumberOfLinesOk

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataNumberOfLinesOk() (*string, bool)`

GetOpeninvoicedataNumberOfLinesOk returns a tuple with the OpeninvoicedataNumberOfLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeninvoicedataNumberOfLines

`func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataNumberOfLines(v string)`

SetOpeninvoicedataNumberOfLines sets OpeninvoicedataNumberOfLines field to given value.

### HasOpeninvoicedataNumberOfLines

`func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataNumberOfLines() bool`

HasOpeninvoicedataNumberOfLines returns a boolean if a field has been set.

### GetOpeninvoicedataRecipientFirstName

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataRecipientFirstName() string`

GetOpeninvoicedataRecipientFirstName returns the OpeninvoicedataRecipientFirstName field if non-nil, zero value otherwise.

### GetOpeninvoicedataRecipientFirstNameOk

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataRecipientFirstNameOk() (*string, bool)`

GetOpeninvoicedataRecipientFirstNameOk returns a tuple with the OpeninvoicedataRecipientFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeninvoicedataRecipientFirstName

`func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataRecipientFirstName(v string)`

SetOpeninvoicedataRecipientFirstName sets OpeninvoicedataRecipientFirstName field to given value.

### HasOpeninvoicedataRecipientFirstName

`func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataRecipientFirstName() bool`

HasOpeninvoicedataRecipientFirstName returns a boolean if a field has been set.

### GetOpeninvoicedataRecipientLastName

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataRecipientLastName() string`

GetOpeninvoicedataRecipientLastName returns the OpeninvoicedataRecipientLastName field if non-nil, zero value otherwise.

### GetOpeninvoicedataRecipientLastNameOk

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataRecipientLastNameOk() (*string, bool)`

GetOpeninvoicedataRecipientLastNameOk returns a tuple with the OpeninvoicedataRecipientLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeninvoicedataRecipientLastName

`func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataRecipientLastName(v string)`

SetOpeninvoicedataRecipientLastName sets OpeninvoicedataRecipientLastName field to given value.

### HasOpeninvoicedataRecipientLastName

`func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataRecipientLastName() bool`

HasOpeninvoicedataRecipientLastName returns a boolean if a field has been set.

### GetOpeninvoicedataLineItemNrCurrencyCode

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrCurrencyCode() string`

GetOpeninvoicedataLineItemNrCurrencyCode returns the OpeninvoicedataLineItemNrCurrencyCode field if non-nil, zero value otherwise.

### GetOpeninvoicedataLineItemNrCurrencyCodeOk

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrCurrencyCodeOk() (*string, bool)`

GetOpeninvoicedataLineItemNrCurrencyCodeOk returns a tuple with the OpeninvoicedataLineItemNrCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeninvoicedataLineItemNrCurrencyCode

`func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrCurrencyCode(v string)`

SetOpeninvoicedataLineItemNrCurrencyCode sets OpeninvoicedataLineItemNrCurrencyCode field to given value.

### HasOpeninvoicedataLineItemNrCurrencyCode

`func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrCurrencyCode() bool`

HasOpeninvoicedataLineItemNrCurrencyCode returns a boolean if a field has been set.

### GetOpeninvoicedataLineItemNrDescription

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrDescription() string`

GetOpeninvoicedataLineItemNrDescription returns the OpeninvoicedataLineItemNrDescription field if non-nil, zero value otherwise.

### GetOpeninvoicedataLineItemNrDescriptionOk

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrDescriptionOk() (*string, bool)`

GetOpeninvoicedataLineItemNrDescriptionOk returns a tuple with the OpeninvoicedataLineItemNrDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeninvoicedataLineItemNrDescription

`func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrDescription(v string)`

SetOpeninvoicedataLineItemNrDescription sets OpeninvoicedataLineItemNrDescription field to given value.

### HasOpeninvoicedataLineItemNrDescription

`func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrDescription() bool`

HasOpeninvoicedataLineItemNrDescription returns a boolean if a field has been set.

### GetOpeninvoicedataLineItemNrItemAmount

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrItemAmount() string`

GetOpeninvoicedataLineItemNrItemAmount returns the OpeninvoicedataLineItemNrItemAmount field if non-nil, zero value otherwise.

### GetOpeninvoicedataLineItemNrItemAmountOk

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrItemAmountOk() (*string, bool)`

GetOpeninvoicedataLineItemNrItemAmountOk returns a tuple with the OpeninvoicedataLineItemNrItemAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeninvoicedataLineItemNrItemAmount

`func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrItemAmount(v string)`

SetOpeninvoicedataLineItemNrItemAmount sets OpeninvoicedataLineItemNrItemAmount field to given value.

### HasOpeninvoicedataLineItemNrItemAmount

`func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrItemAmount() bool`

HasOpeninvoicedataLineItemNrItemAmount returns a boolean if a field has been set.

### GetOpeninvoicedataLineItemNrItemId

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrItemId() string`

GetOpeninvoicedataLineItemNrItemId returns the OpeninvoicedataLineItemNrItemId field if non-nil, zero value otherwise.

### GetOpeninvoicedataLineItemNrItemIdOk

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrItemIdOk() (*string, bool)`

GetOpeninvoicedataLineItemNrItemIdOk returns a tuple with the OpeninvoicedataLineItemNrItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeninvoicedataLineItemNrItemId

`func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrItemId(v string)`

SetOpeninvoicedataLineItemNrItemId sets OpeninvoicedataLineItemNrItemId field to given value.

### HasOpeninvoicedataLineItemNrItemId

`func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrItemId() bool`

HasOpeninvoicedataLineItemNrItemId returns a boolean if a field has been set.

### GetOpeninvoicedataLineItemNrItemVatAmount

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrItemVatAmount() string`

GetOpeninvoicedataLineItemNrItemVatAmount returns the OpeninvoicedataLineItemNrItemVatAmount field if non-nil, zero value otherwise.

### GetOpeninvoicedataLineItemNrItemVatAmountOk

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrItemVatAmountOk() (*string, bool)`

GetOpeninvoicedataLineItemNrItemVatAmountOk returns a tuple with the OpeninvoicedataLineItemNrItemVatAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeninvoicedataLineItemNrItemVatAmount

`func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrItemVatAmount(v string)`

SetOpeninvoicedataLineItemNrItemVatAmount sets OpeninvoicedataLineItemNrItemVatAmount field to given value.

### HasOpeninvoicedataLineItemNrItemVatAmount

`func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrItemVatAmount() bool`

HasOpeninvoicedataLineItemNrItemVatAmount returns a boolean if a field has been set.

### GetOpeninvoicedataLineItemNrItemVatPercentage

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrItemVatPercentage() string`

GetOpeninvoicedataLineItemNrItemVatPercentage returns the OpeninvoicedataLineItemNrItemVatPercentage field if non-nil, zero value otherwise.

### GetOpeninvoicedataLineItemNrItemVatPercentageOk

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrItemVatPercentageOk() (*string, bool)`

GetOpeninvoicedataLineItemNrItemVatPercentageOk returns a tuple with the OpeninvoicedataLineItemNrItemVatPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeninvoicedataLineItemNrItemVatPercentage

`func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrItemVatPercentage(v string)`

SetOpeninvoicedataLineItemNrItemVatPercentage sets OpeninvoicedataLineItemNrItemVatPercentage field to given value.

### HasOpeninvoicedataLineItemNrItemVatPercentage

`func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrItemVatPercentage() bool`

HasOpeninvoicedataLineItemNrItemVatPercentage returns a boolean if a field has been set.

### GetOpeninvoicedataLineItemNrNumberOfItems

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrNumberOfItems() string`

GetOpeninvoicedataLineItemNrNumberOfItems returns the OpeninvoicedataLineItemNrNumberOfItems field if non-nil, zero value otherwise.

### GetOpeninvoicedataLineItemNrNumberOfItemsOk

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrNumberOfItemsOk() (*string, bool)`

GetOpeninvoicedataLineItemNrNumberOfItemsOk returns a tuple with the OpeninvoicedataLineItemNrNumberOfItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeninvoicedataLineItemNrNumberOfItems

`func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrNumberOfItems(v string)`

SetOpeninvoicedataLineItemNrNumberOfItems sets OpeninvoicedataLineItemNrNumberOfItems field to given value.

### HasOpeninvoicedataLineItemNrNumberOfItems

`func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrNumberOfItems() bool`

HasOpeninvoicedataLineItemNrNumberOfItems returns a boolean if a field has been set.

### GetOpeninvoicedataLineItemNrReturnShippingCompany

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrReturnShippingCompany() string`

GetOpeninvoicedataLineItemNrReturnShippingCompany returns the OpeninvoicedataLineItemNrReturnShippingCompany field if non-nil, zero value otherwise.

### GetOpeninvoicedataLineItemNrReturnShippingCompanyOk

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrReturnShippingCompanyOk() (*string, bool)`

GetOpeninvoicedataLineItemNrReturnShippingCompanyOk returns a tuple with the OpeninvoicedataLineItemNrReturnShippingCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeninvoicedataLineItemNrReturnShippingCompany

`func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrReturnShippingCompany(v string)`

SetOpeninvoicedataLineItemNrReturnShippingCompany sets OpeninvoicedataLineItemNrReturnShippingCompany field to given value.

### HasOpeninvoicedataLineItemNrReturnShippingCompany

`func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrReturnShippingCompany() bool`

HasOpeninvoicedataLineItemNrReturnShippingCompany returns a boolean if a field has been set.

### GetOpeninvoicedataLineItemNrReturnTrackingNumber

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrReturnTrackingNumber() string`

GetOpeninvoicedataLineItemNrReturnTrackingNumber returns the OpeninvoicedataLineItemNrReturnTrackingNumber field if non-nil, zero value otherwise.

### GetOpeninvoicedataLineItemNrReturnTrackingNumberOk

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrReturnTrackingNumberOk() (*string, bool)`

GetOpeninvoicedataLineItemNrReturnTrackingNumberOk returns a tuple with the OpeninvoicedataLineItemNrReturnTrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeninvoicedataLineItemNrReturnTrackingNumber

`func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrReturnTrackingNumber(v string)`

SetOpeninvoicedataLineItemNrReturnTrackingNumber sets OpeninvoicedataLineItemNrReturnTrackingNumber field to given value.

### HasOpeninvoicedataLineItemNrReturnTrackingNumber

`func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrReturnTrackingNumber() bool`

HasOpeninvoicedataLineItemNrReturnTrackingNumber returns a boolean if a field has been set.

### GetOpeninvoicedataLineItemNrReturnTrackingUri

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrReturnTrackingUri() string`

GetOpeninvoicedataLineItemNrReturnTrackingUri returns the OpeninvoicedataLineItemNrReturnTrackingUri field if non-nil, zero value otherwise.

### GetOpeninvoicedataLineItemNrReturnTrackingUriOk

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrReturnTrackingUriOk() (*string, bool)`

GetOpeninvoicedataLineItemNrReturnTrackingUriOk returns a tuple with the OpeninvoicedataLineItemNrReturnTrackingUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeninvoicedataLineItemNrReturnTrackingUri

`func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrReturnTrackingUri(v string)`

SetOpeninvoicedataLineItemNrReturnTrackingUri sets OpeninvoicedataLineItemNrReturnTrackingUri field to given value.

### HasOpeninvoicedataLineItemNrReturnTrackingUri

`func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrReturnTrackingUri() bool`

HasOpeninvoicedataLineItemNrReturnTrackingUri returns a boolean if a field has been set.

### GetOpeninvoicedataLineItemNrShippingCompany

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrShippingCompany() string`

GetOpeninvoicedataLineItemNrShippingCompany returns the OpeninvoicedataLineItemNrShippingCompany field if non-nil, zero value otherwise.

### GetOpeninvoicedataLineItemNrShippingCompanyOk

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrShippingCompanyOk() (*string, bool)`

GetOpeninvoicedataLineItemNrShippingCompanyOk returns a tuple with the OpeninvoicedataLineItemNrShippingCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeninvoicedataLineItemNrShippingCompany

`func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrShippingCompany(v string)`

SetOpeninvoicedataLineItemNrShippingCompany sets OpeninvoicedataLineItemNrShippingCompany field to given value.

### HasOpeninvoicedataLineItemNrShippingCompany

`func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrShippingCompany() bool`

HasOpeninvoicedataLineItemNrShippingCompany returns a boolean if a field has been set.

### GetOpeninvoicedataLineItemNrShippingMethod

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrShippingMethod() string`

GetOpeninvoicedataLineItemNrShippingMethod returns the OpeninvoicedataLineItemNrShippingMethod field if non-nil, zero value otherwise.

### GetOpeninvoicedataLineItemNrShippingMethodOk

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrShippingMethodOk() (*string, bool)`

GetOpeninvoicedataLineItemNrShippingMethodOk returns a tuple with the OpeninvoicedataLineItemNrShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeninvoicedataLineItemNrShippingMethod

`func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrShippingMethod(v string)`

SetOpeninvoicedataLineItemNrShippingMethod sets OpeninvoicedataLineItemNrShippingMethod field to given value.

### HasOpeninvoicedataLineItemNrShippingMethod

`func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrShippingMethod() bool`

HasOpeninvoicedataLineItemNrShippingMethod returns a boolean if a field has been set.

### GetOpeninvoicedataLineItemNrTrackingNumber

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrTrackingNumber() string`

GetOpeninvoicedataLineItemNrTrackingNumber returns the OpeninvoicedataLineItemNrTrackingNumber field if non-nil, zero value otherwise.

### GetOpeninvoicedataLineItemNrTrackingNumberOk

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrTrackingNumberOk() (*string, bool)`

GetOpeninvoicedataLineItemNrTrackingNumberOk returns a tuple with the OpeninvoicedataLineItemNrTrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeninvoicedataLineItemNrTrackingNumber

`func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrTrackingNumber(v string)`

SetOpeninvoicedataLineItemNrTrackingNumber sets OpeninvoicedataLineItemNrTrackingNumber field to given value.

### HasOpeninvoicedataLineItemNrTrackingNumber

`func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrTrackingNumber() bool`

HasOpeninvoicedataLineItemNrTrackingNumber returns a boolean if a field has been set.

### GetOpeninvoicedataLineItemNrTrackingUri

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrTrackingUri() string`

GetOpeninvoicedataLineItemNrTrackingUri returns the OpeninvoicedataLineItemNrTrackingUri field if non-nil, zero value otherwise.

### GetOpeninvoicedataLineItemNrTrackingUriOk

`func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrTrackingUriOk() (*string, bool)`

GetOpeninvoicedataLineItemNrTrackingUriOk returns a tuple with the OpeninvoicedataLineItemNrTrackingUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeninvoicedataLineItemNrTrackingUri

`func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrTrackingUri(v string)`

SetOpeninvoicedataLineItemNrTrackingUri sets OpeninvoicedataLineItemNrTrackingUri field to given value.

### HasOpeninvoicedataLineItemNrTrackingUri

`func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrTrackingUri() bool`

HasOpeninvoicedataLineItemNrTrackingUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


