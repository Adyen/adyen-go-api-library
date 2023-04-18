# SourceOfFunds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcquiringBusinessLineId** | Pointer to **string** | The unique identifier of the business line that will be the source of funds.This must be a business line for a **receivePayments** or **receiveFromPlatformPayments** capability. | [optional] 
**AdyenProcessedFunds** | Pointer to **bool** | Indicates whether the funds are coming from transactions processed by Adyen. If **false**, a &#x60;description&#x60; is required. | [optional] 
**Description** | Pointer to **string** | Text describing the source of funds. For example, for &#x60;type&#x60; **business**, provide a description of where the business transactions come from, such as payments through bank transfer. Required when &#x60;adyenProcessedFunds&#x60; is **false**. | [optional] 
**Type** | Pointer to **string** | The type of the source of funds. Possible value: **business**. | [optional] 

## Methods

### NewSourceOfFunds

`func NewSourceOfFunds() *SourceOfFunds`

NewSourceOfFunds instantiates a new SourceOfFunds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceOfFundsWithDefaults

`func NewSourceOfFundsWithDefaults() *SourceOfFunds`

NewSourceOfFundsWithDefaults instantiates a new SourceOfFunds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcquiringBusinessLineId

`func (o *SourceOfFunds) GetAcquiringBusinessLineId() string`

GetAcquiringBusinessLineId returns the AcquiringBusinessLineId field if non-nil, zero value otherwise.

### GetAcquiringBusinessLineIdOk

`func (o *SourceOfFunds) GetAcquiringBusinessLineIdOk() (*string, bool)`

GetAcquiringBusinessLineIdOk returns a tuple with the AcquiringBusinessLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquiringBusinessLineId

`func (o *SourceOfFunds) SetAcquiringBusinessLineId(v string)`

SetAcquiringBusinessLineId sets AcquiringBusinessLineId field to given value.

### HasAcquiringBusinessLineId

`func (o *SourceOfFunds) HasAcquiringBusinessLineId() bool`

HasAcquiringBusinessLineId returns a boolean if a field has been set.

### GetAdyenProcessedFunds

`func (o *SourceOfFunds) GetAdyenProcessedFunds() bool`

GetAdyenProcessedFunds returns the AdyenProcessedFunds field if non-nil, zero value otherwise.

### GetAdyenProcessedFundsOk

`func (o *SourceOfFunds) GetAdyenProcessedFundsOk() (*bool, bool)`

GetAdyenProcessedFundsOk returns a tuple with the AdyenProcessedFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdyenProcessedFunds

`func (o *SourceOfFunds) SetAdyenProcessedFunds(v bool)`

SetAdyenProcessedFunds sets AdyenProcessedFunds field to given value.

### HasAdyenProcessedFunds

`func (o *SourceOfFunds) HasAdyenProcessedFunds() bool`

HasAdyenProcessedFunds returns a boolean if a field has been set.

### GetDescription

`func (o *SourceOfFunds) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SourceOfFunds) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SourceOfFunds) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SourceOfFunds) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *SourceOfFunds) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SourceOfFunds) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SourceOfFunds) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SourceOfFunds) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


