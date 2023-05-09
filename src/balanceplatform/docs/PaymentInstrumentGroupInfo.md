# PaymentInstrumentGroupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalancePlatform** | **string** | The unique identifier of the [balance platform](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/get/balancePlatforms/{id}__queryParam_id) to which the payment instrument group belongs. | 
**Description** | Pointer to **string** | Your description for the payment instrument group, maximum 300 characters. | [optional] 
**Properties** | Pointer to **map[string]string** | Properties of the payment instrument group. | [optional] 
**Reference** | Pointer to **string** | Your reference for the payment instrument group, maximum 150 characters. | [optional] 
**TxVariant** | **string** | The tx variant of the payment instrument group. | 

## Methods

### NewPaymentInstrumentGroupInfo

`func NewPaymentInstrumentGroupInfo(balancePlatform string, txVariant string, ) *PaymentInstrumentGroupInfo`

NewPaymentInstrumentGroupInfo instantiates a new PaymentInstrumentGroupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInstrumentGroupInfoWithDefaults

`func NewPaymentInstrumentGroupInfoWithDefaults() *PaymentInstrumentGroupInfo`

NewPaymentInstrumentGroupInfoWithDefaults instantiates a new PaymentInstrumentGroupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalancePlatform

`func (o *PaymentInstrumentGroupInfo) GetBalancePlatform() string`

GetBalancePlatform returns the BalancePlatform field if non-nil, zero value otherwise.

### GetBalancePlatformOk

`func (o *PaymentInstrumentGroupInfo) GetBalancePlatformOk() (*string, bool)`

GetBalancePlatformOk returns a tuple with the BalancePlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancePlatform

`func (o *PaymentInstrumentGroupInfo) SetBalancePlatform(v string)`

SetBalancePlatform sets BalancePlatform field to given value.


### GetDescription

`func (o *PaymentInstrumentGroupInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentInstrumentGroupInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentInstrumentGroupInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentInstrumentGroupInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProperties

`func (o *PaymentInstrumentGroupInfo) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PaymentInstrumentGroupInfo) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PaymentInstrumentGroupInfo) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PaymentInstrumentGroupInfo) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetReference

`func (o *PaymentInstrumentGroupInfo) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaymentInstrumentGroupInfo) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaymentInstrumentGroupInfo) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PaymentInstrumentGroupInfo) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetTxVariant

`func (o *PaymentInstrumentGroupInfo) GetTxVariant() string`

GetTxVariant returns the TxVariant field if non-nil, zero value otherwise.

### GetTxVariantOk

`func (o *PaymentInstrumentGroupInfo) GetTxVariantOk() (*string, bool)`

GetTxVariantOk returns a tuple with the TxVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxVariant

`func (o *PaymentInstrumentGroupInfo) SetTxVariant(v string)`

SetTxVariant sets TxVariant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


