# TechnicalCancelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **map[string]string** | This field contains additional data, which may be required for a particular modification request.  The additionalData object consists of entries, each of which includes the key and value. | [optional] 
**MerchantAccount** | **string** | The merchant account that is used to process the payment. | 
**ModificationAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**MpiData** | Pointer to [**ThreeDSecureData**](ThreeDSecureData.md) |  | [optional] 
**OriginalMerchantReference** | **string** | The original merchant reference to cancel. | 
**PlatformChargebackLogic** | Pointer to [**PlatformChargebackLogic**](PlatformChargebackLogic.md) |  | [optional] 
**Reference** | Pointer to **string** | Your reference for the payment modification. This reference is visible in Customer Area and in reports. Maximum length: 80 characters. | [optional] 
**Splits** | Pointer to [**[]Split**](Split.md) | An array of objects specifying how the amount should be split between accounts when using Adyen for Platforms. For details, refer to [Providing split information](https://docs.adyen.com/marketplaces-and-platforms/processing-payments#providing-split-information). | [optional] 
**TenderReference** | Pointer to **string** | The transaction reference provided by the PED. For point-of-sale integrations only. | [optional] 
**UniqueTerminalId** | Pointer to **string** | Unique terminal ID for the PED that originally processed the request. For point-of-sale integrations only. | [optional] 

## Methods

### NewTechnicalCancelRequest

`func NewTechnicalCancelRequest(merchantAccount string, originalMerchantReference string, ) *TechnicalCancelRequest`

NewTechnicalCancelRequest instantiates a new TechnicalCancelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechnicalCancelRequestWithDefaults

`func NewTechnicalCancelRequestWithDefaults() *TechnicalCancelRequest`

NewTechnicalCancelRequestWithDefaults instantiates a new TechnicalCancelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *TechnicalCancelRequest) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *TechnicalCancelRequest) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *TechnicalCancelRequest) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *TechnicalCancelRequest) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *TechnicalCancelRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *TechnicalCancelRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *TechnicalCancelRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetModificationAmount

`func (o *TechnicalCancelRequest) GetModificationAmount() Amount`

GetModificationAmount returns the ModificationAmount field if non-nil, zero value otherwise.

### GetModificationAmountOk

`func (o *TechnicalCancelRequest) GetModificationAmountOk() (*Amount, bool)`

GetModificationAmountOk returns a tuple with the ModificationAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationAmount

`func (o *TechnicalCancelRequest) SetModificationAmount(v Amount)`

SetModificationAmount sets ModificationAmount field to given value.

### HasModificationAmount

`func (o *TechnicalCancelRequest) HasModificationAmount() bool`

HasModificationAmount returns a boolean if a field has been set.

### GetMpiData

`func (o *TechnicalCancelRequest) GetMpiData() ThreeDSecureData`

GetMpiData returns the MpiData field if non-nil, zero value otherwise.

### GetMpiDataOk

`func (o *TechnicalCancelRequest) GetMpiDataOk() (*ThreeDSecureData, bool)`

GetMpiDataOk returns a tuple with the MpiData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpiData

`func (o *TechnicalCancelRequest) SetMpiData(v ThreeDSecureData)`

SetMpiData sets MpiData field to given value.

### HasMpiData

`func (o *TechnicalCancelRequest) HasMpiData() bool`

HasMpiData returns a boolean if a field has been set.

### GetOriginalMerchantReference

`func (o *TechnicalCancelRequest) GetOriginalMerchantReference() string`

GetOriginalMerchantReference returns the OriginalMerchantReference field if non-nil, zero value otherwise.

### GetOriginalMerchantReferenceOk

`func (o *TechnicalCancelRequest) GetOriginalMerchantReferenceOk() (*string, bool)`

GetOriginalMerchantReferenceOk returns a tuple with the OriginalMerchantReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalMerchantReference

`func (o *TechnicalCancelRequest) SetOriginalMerchantReference(v string)`

SetOriginalMerchantReference sets OriginalMerchantReference field to given value.


### GetPlatformChargebackLogic

`func (o *TechnicalCancelRequest) GetPlatformChargebackLogic() PlatformChargebackLogic`

GetPlatformChargebackLogic returns the PlatformChargebackLogic field if non-nil, zero value otherwise.

### GetPlatformChargebackLogicOk

`func (o *TechnicalCancelRequest) GetPlatformChargebackLogicOk() (*PlatformChargebackLogic, bool)`

GetPlatformChargebackLogicOk returns a tuple with the PlatformChargebackLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformChargebackLogic

`func (o *TechnicalCancelRequest) SetPlatformChargebackLogic(v PlatformChargebackLogic)`

SetPlatformChargebackLogic sets PlatformChargebackLogic field to given value.

### HasPlatformChargebackLogic

`func (o *TechnicalCancelRequest) HasPlatformChargebackLogic() bool`

HasPlatformChargebackLogic returns a boolean if a field has been set.

### GetReference

`func (o *TechnicalCancelRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *TechnicalCancelRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *TechnicalCancelRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *TechnicalCancelRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSplits

`func (o *TechnicalCancelRequest) GetSplits() []Split`

GetSplits returns the Splits field if non-nil, zero value otherwise.

### GetSplitsOk

`func (o *TechnicalCancelRequest) GetSplitsOk() (*[]Split, bool)`

GetSplitsOk returns a tuple with the Splits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplits

`func (o *TechnicalCancelRequest) SetSplits(v []Split)`

SetSplits sets Splits field to given value.

### HasSplits

`func (o *TechnicalCancelRequest) HasSplits() bool`

HasSplits returns a boolean if a field has been set.

### GetTenderReference

`func (o *TechnicalCancelRequest) GetTenderReference() string`

GetTenderReference returns the TenderReference field if non-nil, zero value otherwise.

### GetTenderReferenceOk

`func (o *TechnicalCancelRequest) GetTenderReferenceOk() (*string, bool)`

GetTenderReferenceOk returns a tuple with the TenderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenderReference

`func (o *TechnicalCancelRequest) SetTenderReference(v string)`

SetTenderReference sets TenderReference field to given value.

### HasTenderReference

`func (o *TechnicalCancelRequest) HasTenderReference() bool`

HasTenderReference returns a boolean if a field has been set.

### GetUniqueTerminalId

`func (o *TechnicalCancelRequest) GetUniqueTerminalId() string`

GetUniqueTerminalId returns the UniqueTerminalId field if non-nil, zero value otherwise.

### GetUniqueTerminalIdOk

`func (o *TechnicalCancelRequest) GetUniqueTerminalIdOk() (*string, bool)`

GetUniqueTerminalIdOk returns a tuple with the UniqueTerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueTerminalId

`func (o *TechnicalCancelRequest) SetUniqueTerminalId(v string)`

SetUniqueTerminalId sets UniqueTerminalId field to given value.

### HasUniqueTerminalId

`func (o *TechnicalCancelRequest) HasUniqueTerminalId() bool`

HasUniqueTerminalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


