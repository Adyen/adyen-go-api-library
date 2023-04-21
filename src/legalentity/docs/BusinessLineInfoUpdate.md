# BusinessLineInfoUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capability** | Pointer to **string** | The capability for which you are creating the business line. For example, **receivePayments**. | [optional] 
**IndustryCode** | Pointer to **string** | A code that represents the industry of your legal entity. For example, **4431A** for computer software stores. | [optional] 
**LegalEntityId** | Pointer to **string** | Unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/legalEntities__resParam_id) that owns the business line. | [optional] 
**SalesChannels** | Pointer to **[]string** | A list of channels where goods or services are sold.  Possible values: **pos**, **posMoto**, **eCommerce**, **ecomMoto**, **payByLink**.  Required only in combination with the &#x60;service&#x60; **paymentProcessing**. | [optional] 
**Service** | **string** | The service for which you are creating the business line.  Possible values:**paymentProcessing**, **issuing**, **banking** | 
**SourceOfFunds** | Pointer to [**SourceOfFunds**](SourceOfFunds.md) |  | [optional] 
**WebData** | Pointer to [**[]WebData**](WebData.md) | List of website URLs where your user&#39;s goods or services are sold. When this is required for a service but your user does not have an online presence, provide the reason in the &#x60;webDataExemption&#x60; object. | [optional] 
**WebDataExemption** | Pointer to [**WebDataExemption**](WebDataExemption.md) |  | [optional] 

## Methods

### NewBusinessLineInfoUpdate

`func NewBusinessLineInfoUpdate(service string, ) *BusinessLineInfoUpdate`

NewBusinessLineInfoUpdate instantiates a new BusinessLineInfoUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessLineInfoUpdateWithDefaults

`func NewBusinessLineInfoUpdateWithDefaults() *BusinessLineInfoUpdate`

NewBusinessLineInfoUpdateWithDefaults instantiates a new BusinessLineInfoUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapability

`func (o *BusinessLineInfoUpdate) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *BusinessLineInfoUpdate) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *BusinessLineInfoUpdate) SetCapability(v string)`

SetCapability sets Capability field to given value.

### HasCapability

`func (o *BusinessLineInfoUpdate) HasCapability() bool`

HasCapability returns a boolean if a field has been set.

### GetIndustryCode

`func (o *BusinessLineInfoUpdate) GetIndustryCode() string`

GetIndustryCode returns the IndustryCode field if non-nil, zero value otherwise.

### GetIndustryCodeOk

`func (o *BusinessLineInfoUpdate) GetIndustryCodeOk() (*string, bool)`

GetIndustryCodeOk returns a tuple with the IndustryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryCode

`func (o *BusinessLineInfoUpdate) SetIndustryCode(v string)`

SetIndustryCode sets IndustryCode field to given value.

### HasIndustryCode

`func (o *BusinessLineInfoUpdate) HasIndustryCode() bool`

HasIndustryCode returns a boolean if a field has been set.

### GetLegalEntityId

`func (o *BusinessLineInfoUpdate) GetLegalEntityId() string`

GetLegalEntityId returns the LegalEntityId field if non-nil, zero value otherwise.

### GetLegalEntityIdOk

`func (o *BusinessLineInfoUpdate) GetLegalEntityIdOk() (*string, bool)`

GetLegalEntityIdOk returns a tuple with the LegalEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalEntityId

`func (o *BusinessLineInfoUpdate) SetLegalEntityId(v string)`

SetLegalEntityId sets LegalEntityId field to given value.

### HasLegalEntityId

`func (o *BusinessLineInfoUpdate) HasLegalEntityId() bool`

HasLegalEntityId returns a boolean if a field has been set.

### GetSalesChannels

`func (o *BusinessLineInfoUpdate) GetSalesChannels() []string`

GetSalesChannels returns the SalesChannels field if non-nil, zero value otherwise.

### GetSalesChannelsOk

`func (o *BusinessLineInfoUpdate) GetSalesChannelsOk() (*[]string, bool)`

GetSalesChannelsOk returns a tuple with the SalesChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesChannels

`func (o *BusinessLineInfoUpdate) SetSalesChannels(v []string)`

SetSalesChannels sets SalesChannels field to given value.

### HasSalesChannels

`func (o *BusinessLineInfoUpdate) HasSalesChannels() bool`

HasSalesChannels returns a boolean if a field has been set.

### GetService

`func (o *BusinessLineInfoUpdate) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *BusinessLineInfoUpdate) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *BusinessLineInfoUpdate) SetService(v string)`

SetService sets Service field to given value.


### GetSourceOfFunds

`func (o *BusinessLineInfoUpdate) GetSourceOfFunds() SourceOfFunds`

GetSourceOfFunds returns the SourceOfFunds field if non-nil, zero value otherwise.

### GetSourceOfFundsOk

`func (o *BusinessLineInfoUpdate) GetSourceOfFundsOk() (*SourceOfFunds, bool)`

GetSourceOfFundsOk returns a tuple with the SourceOfFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfFunds

`func (o *BusinessLineInfoUpdate) SetSourceOfFunds(v SourceOfFunds)`

SetSourceOfFunds sets SourceOfFunds field to given value.

### HasSourceOfFunds

`func (o *BusinessLineInfoUpdate) HasSourceOfFunds() bool`

HasSourceOfFunds returns a boolean if a field has been set.

### GetWebData

`func (o *BusinessLineInfoUpdate) GetWebData() []WebData`

GetWebData returns the WebData field if non-nil, zero value otherwise.

### GetWebDataOk

`func (o *BusinessLineInfoUpdate) GetWebDataOk() (*[]WebData, bool)`

GetWebDataOk returns a tuple with the WebData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebData

`func (o *BusinessLineInfoUpdate) SetWebData(v []WebData)`

SetWebData sets WebData field to given value.

### HasWebData

`func (o *BusinessLineInfoUpdate) HasWebData() bool`

HasWebData returns a boolean if a field has been set.

### GetWebDataExemption

`func (o *BusinessLineInfoUpdate) GetWebDataExemption() WebDataExemption`

GetWebDataExemption returns the WebDataExemption field if non-nil, zero value otherwise.

### GetWebDataExemptionOk

`func (o *BusinessLineInfoUpdate) GetWebDataExemptionOk() (*WebDataExemption, bool)`

GetWebDataExemptionOk returns a tuple with the WebDataExemption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebDataExemption

`func (o *BusinessLineInfoUpdate) SetWebDataExemption(v WebDataExemption)`

SetWebDataExemption sets WebDataExemption field to given value.

### HasWebDataExemption

`func (o *BusinessLineInfoUpdate) HasWebDataExemption() bool`

HasWebDataExemption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


