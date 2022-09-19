# CardholderReceipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeaderForAuthorizedReceipt** | Pointer to **string** | A custom header to show on the shopper receipt for an authorised transaction. Allows one or two comma-separated header lines, and blank lines. For example, &#x60;header,header,filler&#x60; | [optional] 

## Methods

### NewCardholderReceipt

`func NewCardholderReceipt() *CardholderReceipt`

NewCardholderReceipt instantiates a new CardholderReceipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardholderReceiptWithDefaults

`func NewCardholderReceiptWithDefaults() *CardholderReceipt`

NewCardholderReceiptWithDefaults instantiates a new CardholderReceipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaderForAuthorizedReceipt

`func (o *CardholderReceipt) GetHeaderForAuthorizedReceipt() string`

GetHeaderForAuthorizedReceipt returns the HeaderForAuthorizedReceipt field if non-nil, zero value otherwise.

### GetHeaderForAuthorizedReceiptOk

`func (o *CardholderReceipt) GetHeaderForAuthorizedReceiptOk() (*string, bool)`

GetHeaderForAuthorizedReceiptOk returns a tuple with the HeaderForAuthorizedReceipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderForAuthorizedReceipt

`func (o *CardholderReceipt) SetHeaderForAuthorizedReceipt(v string)`

SetHeaderForAuthorizedReceipt sets HeaderForAuthorizedReceipt field to given value.

### HasHeaderForAuthorizedReceipt

`func (o *CardholderReceipt) HasHeaderForAuthorizedReceipt() bool`

HasHeaderForAuthorizedReceipt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


