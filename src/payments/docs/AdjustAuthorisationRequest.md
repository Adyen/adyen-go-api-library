# AdjustAuthorisationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **map[string]string** | This field contains additional data, which may be required for a particular modification request.  The additionalData object consists of entries, each of which includes the key and value. | [optional] 
**MerchantAccount** | **string** | The merchant account that is used to process the payment. | 
**ModificationAmount** | [**Amount**](Amount.md) |  | 
**MpiData** | Pointer to [**ThreeDSecureData**](ThreeDSecureData.md) |  | [optional] 
**OriginalMerchantReference** | Pointer to **string** | The original merchant reference to cancel. | [optional] 
**OriginalReference** | **string** | The original pspReference of the payment to modify. This reference is returned in: * authorisation response * authorisation notification   | 
**PlatformChargebackLogic** | Pointer to [**PlatformChargebackLogic**](PlatformChargebackLogic.md) |  | [optional] 
**Reference** | Pointer to **string** | Your reference for the payment modification. This reference is visible in Customer Area and in reports. Maximum length: 80 characters. | [optional] 
**Splits** | Pointer to [**[]Split**](Split.md) | An array of objects specifying how the amount should be split between accounts when using Adyen for Platforms. For details, refer to [Providing split information](https://docs.adyen.com/marketplaces-and-platforms/processing-payments#providing-split-information). | [optional] 
**TenderReference** | Pointer to **string** | The transaction reference provided by the PED. For point-of-sale integrations only. | [optional] 
**UniqueTerminalId** | Pointer to **string** | Unique terminal ID for the PED that originally processed the request. For point-of-sale integrations only. | [optional] 

## Methods

### NewAdjustAuthorisationRequest

`func NewAdjustAuthorisationRequest(merchantAccount string, modificationAmount Amount, originalReference string, ) *AdjustAuthorisationRequest`

NewAdjustAuthorisationRequest instantiates a new AdjustAuthorisationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdjustAuthorisationRequestWithDefaults

`func NewAdjustAuthorisationRequestWithDefaults() *AdjustAuthorisationRequest`

NewAdjustAuthorisationRequestWithDefaults instantiates a new AdjustAuthorisationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *AdjustAuthorisationRequest) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *AdjustAuthorisationRequest) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *AdjustAuthorisationRequest) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *AdjustAuthorisationRequest) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *AdjustAuthorisationRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *AdjustAuthorisationRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *AdjustAuthorisationRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetModificationAmount

`func (o *AdjustAuthorisationRequest) GetModificationAmount() Amount`

GetModificationAmount returns the ModificationAmount field if non-nil, zero value otherwise.

### GetModificationAmountOk

`func (o *AdjustAuthorisationRequest) GetModificationAmountOk() (*Amount, bool)`

GetModificationAmountOk returns a tuple with the ModificationAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationAmount

`func (o *AdjustAuthorisationRequest) SetModificationAmount(v Amount)`

SetModificationAmount sets ModificationAmount field to given value.


### GetMpiData

`func (o *AdjustAuthorisationRequest) GetMpiData() ThreeDSecureData`

GetMpiData returns the MpiData field if non-nil, zero value otherwise.

### GetMpiDataOk

`func (o *AdjustAuthorisationRequest) GetMpiDataOk() (*ThreeDSecureData, bool)`

GetMpiDataOk returns a tuple with the MpiData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpiData

`func (o *AdjustAuthorisationRequest) SetMpiData(v ThreeDSecureData)`

SetMpiData sets MpiData field to given value.

### HasMpiData

`func (o *AdjustAuthorisationRequest) HasMpiData() bool`

HasMpiData returns a boolean if a field has been set.

### GetOriginalMerchantReference

`func (o *AdjustAuthorisationRequest) GetOriginalMerchantReference() string`

GetOriginalMerchantReference returns the OriginalMerchantReference field if non-nil, zero value otherwise.

### GetOriginalMerchantReferenceOk

`func (o *AdjustAuthorisationRequest) GetOriginalMerchantReferenceOk() (*string, bool)`

GetOriginalMerchantReferenceOk returns a tuple with the OriginalMerchantReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalMerchantReference

`func (o *AdjustAuthorisationRequest) SetOriginalMerchantReference(v string)`

SetOriginalMerchantReference sets OriginalMerchantReference field to given value.

### HasOriginalMerchantReference

`func (o *AdjustAuthorisationRequest) HasOriginalMerchantReference() bool`

HasOriginalMerchantReference returns a boolean if a field has been set.

### GetOriginalReference

`func (o *AdjustAuthorisationRequest) GetOriginalReference() string`

GetOriginalReference returns the OriginalReference field if non-nil, zero value otherwise.

### GetOriginalReferenceOk

`func (o *AdjustAuthorisationRequest) GetOriginalReferenceOk() (*string, bool)`

GetOriginalReferenceOk returns a tuple with the OriginalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalReference

`func (o *AdjustAuthorisationRequest) SetOriginalReference(v string)`

SetOriginalReference sets OriginalReference field to given value.


### GetPlatformChargebackLogic

`func (o *AdjustAuthorisationRequest) GetPlatformChargebackLogic() PlatformChargebackLogic`

GetPlatformChargebackLogic returns the PlatformChargebackLogic field if non-nil, zero value otherwise.

### GetPlatformChargebackLogicOk

`func (o *AdjustAuthorisationRequest) GetPlatformChargebackLogicOk() (*PlatformChargebackLogic, bool)`

GetPlatformChargebackLogicOk returns a tuple with the PlatformChargebackLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformChargebackLogic

`func (o *AdjustAuthorisationRequest) SetPlatformChargebackLogic(v PlatformChargebackLogic)`

SetPlatformChargebackLogic sets PlatformChargebackLogic field to given value.

### HasPlatformChargebackLogic

`func (o *AdjustAuthorisationRequest) HasPlatformChargebackLogic() bool`

HasPlatformChargebackLogic returns a boolean if a field has been set.

### GetReference

`func (o *AdjustAuthorisationRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *AdjustAuthorisationRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *AdjustAuthorisationRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *AdjustAuthorisationRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSplits

`func (o *AdjustAuthorisationRequest) GetSplits() []Split`

GetSplits returns the Splits field if non-nil, zero value otherwise.

### GetSplitsOk

`func (o *AdjustAuthorisationRequest) GetSplitsOk() (*[]Split, bool)`

GetSplitsOk returns a tuple with the Splits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplits

`func (o *AdjustAuthorisationRequest) SetSplits(v []Split)`

SetSplits sets Splits field to given value.

### HasSplits

`func (o *AdjustAuthorisationRequest) HasSplits() bool`

HasSplits returns a boolean if a field has been set.

### GetTenderReference

`func (o *AdjustAuthorisationRequest) GetTenderReference() string`

GetTenderReference returns the TenderReference field if non-nil, zero value otherwise.

### GetTenderReferenceOk

`func (o *AdjustAuthorisationRequest) GetTenderReferenceOk() (*string, bool)`

GetTenderReferenceOk returns a tuple with the TenderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenderReference

`func (o *AdjustAuthorisationRequest) SetTenderReference(v string)`

SetTenderReference sets TenderReference field to given value.

### HasTenderReference

`func (o *AdjustAuthorisationRequest) HasTenderReference() bool`

HasTenderReference returns a boolean if a field has been set.

### GetUniqueTerminalId

`func (o *AdjustAuthorisationRequest) GetUniqueTerminalId() string`

GetUniqueTerminalId returns the UniqueTerminalId field if non-nil, zero value otherwise.

### GetUniqueTerminalIdOk

`func (o *AdjustAuthorisationRequest) GetUniqueTerminalIdOk() (*string, bool)`

GetUniqueTerminalIdOk returns a tuple with the UniqueTerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueTerminalId

`func (o *AdjustAuthorisationRequest) SetUniqueTerminalId(v string)`

SetUniqueTerminalId sets UniqueTerminalId field to given value.

### HasUniqueTerminalId

`func (o *AdjustAuthorisationRequest) HasUniqueTerminalId() bool`

HasUniqueTerminalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


