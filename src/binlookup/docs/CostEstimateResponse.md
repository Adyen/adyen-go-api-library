# CostEstimateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardBin** | Pointer to [**CardBin**](CardBin.md) |  | [optional] 
**CostEstimateAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**CostEstimateReference** | Pointer to **string** | Adyen&#39;s 16-character reference associated with the request. | [optional] 
**ResultCode** | Pointer to **string** | The result of the cost estimation. | [optional] 
**SurchargeType** | Pointer to **string** | Indicates the way the charges can be passed on to the cardholder. The following values are possible: * &#x60;ZERO&#x60; - the charges are not allowed to pass on * &#x60;PASSTHROUGH&#x60; - the charges can be passed on * &#x60;UNLIMITED&#x60; - there is no limit on how much surcharge is passed on | [optional] 

## Methods

### NewCostEstimateResponse

`func NewCostEstimateResponse() *CostEstimateResponse`

NewCostEstimateResponse instantiates a new CostEstimateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostEstimateResponseWithDefaults

`func NewCostEstimateResponseWithDefaults() *CostEstimateResponse`

NewCostEstimateResponseWithDefaults instantiates a new CostEstimateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardBin

`func (o *CostEstimateResponse) GetCardBin() CardBin`

GetCardBin returns the CardBin field if non-nil, zero value otherwise.

### GetCardBinOk

`func (o *CostEstimateResponse) GetCardBinOk() (*CardBin, bool)`

GetCardBinOk returns a tuple with the CardBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBin

`func (o *CostEstimateResponse) SetCardBin(v CardBin)`

SetCardBin sets CardBin field to given value.

### HasCardBin

`func (o *CostEstimateResponse) HasCardBin() bool`

HasCardBin returns a boolean if a field has been set.

### GetCostEstimateAmount

`func (o *CostEstimateResponse) GetCostEstimateAmount() Amount`

GetCostEstimateAmount returns the CostEstimateAmount field if non-nil, zero value otherwise.

### GetCostEstimateAmountOk

`func (o *CostEstimateResponse) GetCostEstimateAmountOk() (*Amount, bool)`

GetCostEstimateAmountOk returns a tuple with the CostEstimateAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostEstimateAmount

`func (o *CostEstimateResponse) SetCostEstimateAmount(v Amount)`

SetCostEstimateAmount sets CostEstimateAmount field to given value.

### HasCostEstimateAmount

`func (o *CostEstimateResponse) HasCostEstimateAmount() bool`

HasCostEstimateAmount returns a boolean if a field has been set.

### GetCostEstimateReference

`func (o *CostEstimateResponse) GetCostEstimateReference() string`

GetCostEstimateReference returns the CostEstimateReference field if non-nil, zero value otherwise.

### GetCostEstimateReferenceOk

`func (o *CostEstimateResponse) GetCostEstimateReferenceOk() (*string, bool)`

GetCostEstimateReferenceOk returns a tuple with the CostEstimateReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostEstimateReference

`func (o *CostEstimateResponse) SetCostEstimateReference(v string)`

SetCostEstimateReference sets CostEstimateReference field to given value.

### HasCostEstimateReference

`func (o *CostEstimateResponse) HasCostEstimateReference() bool`

HasCostEstimateReference returns a boolean if a field has been set.

### GetResultCode

`func (o *CostEstimateResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *CostEstimateResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *CostEstimateResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *CostEstimateResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetSurchargeType

`func (o *CostEstimateResponse) GetSurchargeType() string`

GetSurchargeType returns the SurchargeType field if non-nil, zero value otherwise.

### GetSurchargeTypeOk

`func (o *CostEstimateResponse) GetSurchargeTypeOk() (*string, bool)`

GetSurchargeTypeOk returns a tuple with the SurchargeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurchargeType

`func (o *CostEstimateResponse) SetSurchargeType(v string)`

SetSurchargeType sets SurchargeType field to given value.

### HasSurchargeType

`func (o *CostEstimateResponse) HasSurchargeType() bool`

HasSurchargeType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


