# UpdateSplitConfigurationLogicRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalCommission** | Pointer to [**AdditionalCommission**](AdditionalCommission.md) |  | [optional] 
**Chargeback** | Pointer to **string** | Specifies the logic to apply when booking the chargeback amount.  Possible values: **deductFromLiableAccount**, **deductFromOneBalanceAccount**, **deductAccordingToSplitRatio**. | [optional] 
**Commission** | [**Commission**](Commission.md) |  | 
**PaymentFee** | **string** | Specifies the logic to apply when booking the transaction fees.  Possible values: **deductFromLiableAccount**, **deductFromOneBalanceAccount**. | 
**Remainder** | Pointer to **string** | Specifies the logic to apply when booking the amount left over after currency conversion.  Possible values: **addToLiableAccount**, **addToOneBalanceAccount**. | [optional] 
**SplitLogicId** | Pointer to **string** | Unique identifier of the split logic that is applied when the split configuration conditions are met. | [optional] [readonly] 
**Surcharge** | Pointer to **string** | Specifies the logic to apply when booking the surcharge amount.  Possible values: **addToLiableAccount**, **addToOneBalanceAccount** | [optional] 
**Tip** | Pointer to **string** | Specifies the logic to apply when booking tips (gratuity).  Possible values: **addToLiableAccount**, **addToOneBalanceAccount**. | [optional] 

## Methods

### NewUpdateSplitConfigurationLogicRequest

`func NewUpdateSplitConfigurationLogicRequest(commission Commission, paymentFee string, ) *UpdateSplitConfigurationLogicRequest`

NewUpdateSplitConfigurationLogicRequest instantiates a new UpdateSplitConfigurationLogicRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSplitConfigurationLogicRequestWithDefaults

`func NewUpdateSplitConfigurationLogicRequestWithDefaults() *UpdateSplitConfigurationLogicRequest`

NewUpdateSplitConfigurationLogicRequestWithDefaults instantiates a new UpdateSplitConfigurationLogicRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalCommission

`func (o *UpdateSplitConfigurationLogicRequest) GetAdditionalCommission() AdditionalCommission`

GetAdditionalCommission returns the AdditionalCommission field if non-nil, zero value otherwise.

### GetAdditionalCommissionOk

`func (o *UpdateSplitConfigurationLogicRequest) GetAdditionalCommissionOk() (*AdditionalCommission, bool)`

GetAdditionalCommissionOk returns a tuple with the AdditionalCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalCommission

`func (o *UpdateSplitConfigurationLogicRequest) SetAdditionalCommission(v AdditionalCommission)`

SetAdditionalCommission sets AdditionalCommission field to given value.

### HasAdditionalCommission

`func (o *UpdateSplitConfigurationLogicRequest) HasAdditionalCommission() bool`

HasAdditionalCommission returns a boolean if a field has been set.

### GetChargeback

`func (o *UpdateSplitConfigurationLogicRequest) GetChargeback() string`

GetChargeback returns the Chargeback field if non-nil, zero value otherwise.

### GetChargebackOk

`func (o *UpdateSplitConfigurationLogicRequest) GetChargebackOk() (*string, bool)`

GetChargebackOk returns a tuple with the Chargeback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeback

`func (o *UpdateSplitConfigurationLogicRequest) SetChargeback(v string)`

SetChargeback sets Chargeback field to given value.

### HasChargeback

`func (o *UpdateSplitConfigurationLogicRequest) HasChargeback() bool`

HasChargeback returns a boolean if a field has been set.

### GetCommission

`func (o *UpdateSplitConfigurationLogicRequest) GetCommission() Commission`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *UpdateSplitConfigurationLogicRequest) GetCommissionOk() (*Commission, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *UpdateSplitConfigurationLogicRequest) SetCommission(v Commission)`

SetCommission sets Commission field to given value.


### GetPaymentFee

`func (o *UpdateSplitConfigurationLogicRequest) GetPaymentFee() string`

GetPaymentFee returns the PaymentFee field if non-nil, zero value otherwise.

### GetPaymentFeeOk

`func (o *UpdateSplitConfigurationLogicRequest) GetPaymentFeeOk() (*string, bool)`

GetPaymentFeeOk returns a tuple with the PaymentFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFee

`func (o *UpdateSplitConfigurationLogicRequest) SetPaymentFee(v string)`

SetPaymentFee sets PaymentFee field to given value.


### GetRemainder

`func (o *UpdateSplitConfigurationLogicRequest) GetRemainder() string`

GetRemainder returns the Remainder field if non-nil, zero value otherwise.

### GetRemainderOk

`func (o *UpdateSplitConfigurationLogicRequest) GetRemainderOk() (*string, bool)`

GetRemainderOk returns a tuple with the Remainder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainder

`func (o *UpdateSplitConfigurationLogicRequest) SetRemainder(v string)`

SetRemainder sets Remainder field to given value.

### HasRemainder

`func (o *UpdateSplitConfigurationLogicRequest) HasRemainder() bool`

HasRemainder returns a boolean if a field has been set.

### GetSplitLogicId

`func (o *UpdateSplitConfigurationLogicRequest) GetSplitLogicId() string`

GetSplitLogicId returns the SplitLogicId field if non-nil, zero value otherwise.

### GetSplitLogicIdOk

`func (o *UpdateSplitConfigurationLogicRequest) GetSplitLogicIdOk() (*string, bool)`

GetSplitLogicIdOk returns a tuple with the SplitLogicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitLogicId

`func (o *UpdateSplitConfigurationLogicRequest) SetSplitLogicId(v string)`

SetSplitLogicId sets SplitLogicId field to given value.

### HasSplitLogicId

`func (o *UpdateSplitConfigurationLogicRequest) HasSplitLogicId() bool`

HasSplitLogicId returns a boolean if a field has been set.

### GetSurcharge

`func (o *UpdateSplitConfigurationLogicRequest) GetSurcharge() string`

GetSurcharge returns the Surcharge field if non-nil, zero value otherwise.

### GetSurchargeOk

`func (o *UpdateSplitConfigurationLogicRequest) GetSurchargeOk() (*string, bool)`

GetSurchargeOk returns a tuple with the Surcharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurcharge

`func (o *UpdateSplitConfigurationLogicRequest) SetSurcharge(v string)`

SetSurcharge sets Surcharge field to given value.

### HasSurcharge

`func (o *UpdateSplitConfigurationLogicRequest) HasSurcharge() bool`

HasSurcharge returns a boolean if a field has been set.

### GetTip

`func (o *UpdateSplitConfigurationLogicRequest) GetTip() string`

GetTip returns the Tip field if non-nil, zero value otherwise.

### GetTipOk

`func (o *UpdateSplitConfigurationLogicRequest) GetTipOk() (*string, bool)`

GetTipOk returns a tuple with the Tip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTip

`func (o *UpdateSplitConfigurationLogicRequest) SetTip(v string)`

SetTip sets Tip field to given value.

### HasTip

`func (o *UpdateSplitConfigurationLogicRequest) HasTip() bool`

HasTip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


