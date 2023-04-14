# DonationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DonationAccount** | **string** | The Adyen account name of the charity. | 
**MerchantAccount** | **string** | The merchant account that is used to process the payment. | 
**ModificationAmount** | [**Amount**](Amount.md) |  | 
**OriginalReference** | Pointer to **string** | The original pspReference of the payment to modify. This reference is returned in: * authorisation response * authorisation notification   | [optional] 
**PlatformChargebackLogic** | Pointer to [**PlatformChargebackLogic**](PlatformChargebackLogic.md) |  | [optional] 
**Reference** | Pointer to **string** | Your reference for the payment modification. This reference is visible in Customer Area and in reports. Maximum length: 80 characters. | [optional] 

## Methods

### NewDonationRequest

`func NewDonationRequest(donationAccount string, merchantAccount string, modificationAmount Amount, ) *DonationRequest`

NewDonationRequest instantiates a new DonationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDonationRequestWithDefaults

`func NewDonationRequestWithDefaults() *DonationRequest`

NewDonationRequestWithDefaults instantiates a new DonationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDonationAccount

`func (o *DonationRequest) GetDonationAccount() string`

GetDonationAccount returns the DonationAccount field if non-nil, zero value otherwise.

### GetDonationAccountOk

`func (o *DonationRequest) GetDonationAccountOk() (*string, bool)`

GetDonationAccountOk returns a tuple with the DonationAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDonationAccount

`func (o *DonationRequest) SetDonationAccount(v string)`

SetDonationAccount sets DonationAccount field to given value.


### GetMerchantAccount

`func (o *DonationRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *DonationRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *DonationRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetModificationAmount

`func (o *DonationRequest) GetModificationAmount() Amount`

GetModificationAmount returns the ModificationAmount field if non-nil, zero value otherwise.

### GetModificationAmountOk

`func (o *DonationRequest) GetModificationAmountOk() (*Amount, bool)`

GetModificationAmountOk returns a tuple with the ModificationAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationAmount

`func (o *DonationRequest) SetModificationAmount(v Amount)`

SetModificationAmount sets ModificationAmount field to given value.


### GetOriginalReference

`func (o *DonationRequest) GetOriginalReference() string`

GetOriginalReference returns the OriginalReference field if non-nil, zero value otherwise.

### GetOriginalReferenceOk

`func (o *DonationRequest) GetOriginalReferenceOk() (*string, bool)`

GetOriginalReferenceOk returns a tuple with the OriginalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalReference

`func (o *DonationRequest) SetOriginalReference(v string)`

SetOriginalReference sets OriginalReference field to given value.

### HasOriginalReference

`func (o *DonationRequest) HasOriginalReference() bool`

HasOriginalReference returns a boolean if a field has been set.

### GetPlatformChargebackLogic

`func (o *DonationRequest) GetPlatformChargebackLogic() PlatformChargebackLogic`

GetPlatformChargebackLogic returns the PlatformChargebackLogic field if non-nil, zero value otherwise.

### GetPlatformChargebackLogicOk

`func (o *DonationRequest) GetPlatformChargebackLogicOk() (*PlatformChargebackLogic, bool)`

GetPlatformChargebackLogicOk returns a tuple with the PlatformChargebackLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformChargebackLogic

`func (o *DonationRequest) SetPlatformChargebackLogic(v PlatformChargebackLogic)`

SetPlatformChargebackLogic sets PlatformChargebackLogic field to given value.

### HasPlatformChargebackLogic

`func (o *DonationRequest) HasPlatformChargebackLogic() bool`

HasPlatformChargebackLogic returns a boolean if a field has been set.

### GetReference

`func (o *DonationRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *DonationRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *DonationRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *DonationRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


