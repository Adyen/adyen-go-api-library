# PaymentInstrumentGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalancePlatform** | **string** | The unique identifier of the [balance platform](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/get/balancePlatforms/{id}__queryParam_id) to which the payment instrument group belongs. | 
**Description** | Pointer to **string** | Your description for the payment instrument group, maximum 300 characters. | [optional] 
**Id** | Pointer to **string** | The unique identifier of the payment instrument group. | [optional] 
**Properties** | Pointer to **map[string]string** | Properties of the payment instrument group. | [optional] 
**Reference** | Pointer to **string** | Your reference for the payment instrument group, maximum 150 characters. | [optional] 
**TxVariant** | **string** | The tx variant of the payment instrument group. | 

## Methods

### NewPaymentInstrumentGroup

`func NewPaymentInstrumentGroup(balancePlatform string, txVariant string, ) *PaymentInstrumentGroup`

NewPaymentInstrumentGroup instantiates a new PaymentInstrumentGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInstrumentGroupWithDefaults

`func NewPaymentInstrumentGroupWithDefaults() *PaymentInstrumentGroup`

NewPaymentInstrumentGroupWithDefaults instantiates a new PaymentInstrumentGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalancePlatform

`func (o *PaymentInstrumentGroup) GetBalancePlatform() string`

GetBalancePlatform returns the BalancePlatform field if non-nil, zero value otherwise.

### GetBalancePlatformOk

`func (o *PaymentInstrumentGroup) GetBalancePlatformOk() (*string, bool)`

GetBalancePlatformOk returns a tuple with the BalancePlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancePlatform

`func (o *PaymentInstrumentGroup) SetBalancePlatform(v string)`

SetBalancePlatform sets BalancePlatform field to given value.


### GetDescription

`func (o *PaymentInstrumentGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentInstrumentGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentInstrumentGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentInstrumentGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *PaymentInstrumentGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentInstrumentGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentInstrumentGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentInstrumentGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProperties

`func (o *PaymentInstrumentGroup) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PaymentInstrumentGroup) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PaymentInstrumentGroup) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PaymentInstrumentGroup) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetReference

`func (o *PaymentInstrumentGroup) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaymentInstrumentGroup) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaymentInstrumentGroup) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PaymentInstrumentGroup) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetTxVariant

`func (o *PaymentInstrumentGroup) GetTxVariant() string`

GetTxVariant returns the TxVariant field if non-nil, zero value otherwise.

### GetTxVariantOk

`func (o *PaymentInstrumentGroup) GetTxVariantOk() (*string, bool)`

GetTxVariantOk returns a tuple with the TxVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxVariant

`func (o *PaymentInstrumentGroup) SetTxVariant(v string)`

SetTxVariant sets TxVariant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


