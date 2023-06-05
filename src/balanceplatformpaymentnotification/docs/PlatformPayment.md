# PlatformPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | The account given in the related split. | [optional] 
**Description** | Pointer to **string** | The description of the related split. | [optional] 
**ModificationMerchantReference** | Pointer to **string** | The merchant reference of the modification. | [optional] 
**ModificationPspReference** | Pointer to **string** | The pspReference of the modification. | [optional] 
**PaymentMerchantReference** | Pointer to **string** | The merchant reference of the payment. | [optional] 
**PaymentPspReference** | Pointer to **string** | The pspReference of the payment. | [optional] 
**Reference** | Pointer to **string** | The reference of the related split. | [optional] 
**Type** | Pointer to **string** | The type of the related split. | [optional] 

## Methods

### NewPlatformPayment

`func NewPlatformPayment() *PlatformPayment`

NewPlatformPayment instantiates a new PlatformPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformPaymentWithDefaults

`func NewPlatformPaymentWithDefaults() *PlatformPayment`

NewPlatformPaymentWithDefaults instantiates a new PlatformPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *PlatformPayment) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PlatformPayment) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PlatformPayment) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PlatformPayment) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDescription

`func (o *PlatformPayment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlatformPayment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlatformPayment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlatformPayment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetModificationMerchantReference

`func (o *PlatformPayment) GetModificationMerchantReference() string`

GetModificationMerchantReference returns the ModificationMerchantReference field if non-nil, zero value otherwise.

### GetModificationMerchantReferenceOk

`func (o *PlatformPayment) GetModificationMerchantReferenceOk() (*string, bool)`

GetModificationMerchantReferenceOk returns a tuple with the ModificationMerchantReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationMerchantReference

`func (o *PlatformPayment) SetModificationMerchantReference(v string)`

SetModificationMerchantReference sets ModificationMerchantReference field to given value.

### HasModificationMerchantReference

`func (o *PlatformPayment) HasModificationMerchantReference() bool`

HasModificationMerchantReference returns a boolean if a field has been set.

### GetModificationPspReference

`func (o *PlatformPayment) GetModificationPspReference() string`

GetModificationPspReference returns the ModificationPspReference field if non-nil, zero value otherwise.

### GetModificationPspReferenceOk

`func (o *PlatformPayment) GetModificationPspReferenceOk() (*string, bool)`

GetModificationPspReferenceOk returns a tuple with the ModificationPspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationPspReference

`func (o *PlatformPayment) SetModificationPspReference(v string)`

SetModificationPspReference sets ModificationPspReference field to given value.

### HasModificationPspReference

`func (o *PlatformPayment) HasModificationPspReference() bool`

HasModificationPspReference returns a boolean if a field has been set.

### GetPaymentMerchantReference

`func (o *PlatformPayment) GetPaymentMerchantReference() string`

GetPaymentMerchantReference returns the PaymentMerchantReference field if non-nil, zero value otherwise.

### GetPaymentMerchantReferenceOk

`func (o *PlatformPayment) GetPaymentMerchantReferenceOk() (*string, bool)`

GetPaymentMerchantReferenceOk returns a tuple with the PaymentMerchantReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMerchantReference

`func (o *PlatformPayment) SetPaymentMerchantReference(v string)`

SetPaymentMerchantReference sets PaymentMerchantReference field to given value.

### HasPaymentMerchantReference

`func (o *PlatformPayment) HasPaymentMerchantReference() bool`

HasPaymentMerchantReference returns a boolean if a field has been set.

### GetPaymentPspReference

`func (o *PlatformPayment) GetPaymentPspReference() string`

GetPaymentPspReference returns the PaymentPspReference field if non-nil, zero value otherwise.

### GetPaymentPspReferenceOk

`func (o *PlatformPayment) GetPaymentPspReferenceOk() (*string, bool)`

GetPaymentPspReferenceOk returns a tuple with the PaymentPspReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentPspReference

`func (o *PlatformPayment) SetPaymentPspReference(v string)`

SetPaymentPspReference sets PaymentPspReference field to given value.

### HasPaymentPspReference

`func (o *PlatformPayment) HasPaymentPspReference() bool`

HasPaymentPspReference returns a boolean if a field has been set.

### GetReference

`func (o *PlatformPayment) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PlatformPayment) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PlatformPayment) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PlatformPayment) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetType

`func (o *PlatformPayment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlatformPayment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlatformPayment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PlatformPayment) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


