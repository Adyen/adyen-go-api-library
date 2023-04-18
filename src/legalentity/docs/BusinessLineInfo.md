# BusinessLineInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capability** | Pointer to **string** | The capability for which you are creating the business line. For example, **receivePayments**. | [optional] 
**IndustryCode** | **string** | A code that represents the industry of the legal entity. For example, **4431A** for computer software stores. | 
**LegalEntityId** | **string** | Unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/legalEntities__resParam_id) that owns the business line. | 
**SalesChannels** | Pointer to **[]string** | A list of channels where goods or services are sold.  Possible values: **pos**, **posMoto**, **eCommerce**, **ecomMoto**, **payByLink**.  Required only in combination with the &#x60;service&#x60; **paymentProcessing**. | [optional] 
**Service** | **string** | The service for which you are creating the business line.  Possible values:**paymentProcessing**, **issuing**, **banking** | 
**SourceOfFunds** | Pointer to [**SourceOfFunds**](SourceOfFunds.md) |  | [optional] 
**WebData** | Pointer to [**[]WebData**](WebData.md) | List of website URLs where your user&#39;s goods or services are sold. When this is required for a service but your user does not have an online presence, provide the reason in the &#x60;webDataExemption&#x60; object. | [optional] 
**WebDataExemption** | Pointer to [**WebDataExemption**](WebDataExemption.md) |  | [optional] 

## Methods

### NewBusinessLineInfo

`func NewBusinessLineInfo(industryCode string, legalEntityId string, service string, ) *BusinessLineInfo`

NewBusinessLineInfo instantiates a new BusinessLineInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessLineInfoWithDefaults

`func NewBusinessLineInfoWithDefaults() *BusinessLineInfo`

NewBusinessLineInfoWithDefaults instantiates a new BusinessLineInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapability

`func (o *BusinessLineInfo) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *BusinessLineInfo) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *BusinessLineInfo) SetCapability(v string)`

SetCapability sets Capability field to given value.

### HasCapability

`func (o *BusinessLineInfo) HasCapability() bool`

HasCapability returns a boolean if a field has been set.

### GetIndustryCode

`func (o *BusinessLineInfo) GetIndustryCode() string`

GetIndustryCode returns the IndustryCode field if non-nil, zero value otherwise.

### GetIndustryCodeOk

`func (o *BusinessLineInfo) GetIndustryCodeOk() (*string, bool)`

GetIndustryCodeOk returns a tuple with the IndustryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryCode

`func (o *BusinessLineInfo) SetIndustryCode(v string)`

SetIndustryCode sets IndustryCode field to given value.


### GetLegalEntityId

`func (o *BusinessLineInfo) GetLegalEntityId() string`

GetLegalEntityId returns the LegalEntityId field if non-nil, zero value otherwise.

### GetLegalEntityIdOk

`func (o *BusinessLineInfo) GetLegalEntityIdOk() (*string, bool)`

GetLegalEntityIdOk returns a tuple with the LegalEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalEntityId

`func (o *BusinessLineInfo) SetLegalEntityId(v string)`

SetLegalEntityId sets LegalEntityId field to given value.


### GetSalesChannels

`func (o *BusinessLineInfo) GetSalesChannels() []string`

GetSalesChannels returns the SalesChannels field if non-nil, zero value otherwise.

### GetSalesChannelsOk

`func (o *BusinessLineInfo) GetSalesChannelsOk() (*[]string, bool)`

GetSalesChannelsOk returns a tuple with the SalesChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesChannels

`func (o *BusinessLineInfo) SetSalesChannels(v []string)`

SetSalesChannels sets SalesChannels field to given value.

### HasSalesChannels

`func (o *BusinessLineInfo) HasSalesChannels() bool`

HasSalesChannels returns a boolean if a field has been set.

### GetService

`func (o *BusinessLineInfo) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *BusinessLineInfo) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *BusinessLineInfo) SetService(v string)`

SetService sets Service field to given value.


### GetSourceOfFunds

`func (o *BusinessLineInfo) GetSourceOfFunds() SourceOfFunds`

GetSourceOfFunds returns the SourceOfFunds field if non-nil, zero value otherwise.

### GetSourceOfFundsOk

`func (o *BusinessLineInfo) GetSourceOfFundsOk() (*SourceOfFunds, bool)`

GetSourceOfFundsOk returns a tuple with the SourceOfFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfFunds

`func (o *BusinessLineInfo) SetSourceOfFunds(v SourceOfFunds)`

SetSourceOfFunds sets SourceOfFunds field to given value.

### HasSourceOfFunds

`func (o *BusinessLineInfo) HasSourceOfFunds() bool`

HasSourceOfFunds returns a boolean if a field has been set.

### GetWebData

`func (o *BusinessLineInfo) GetWebData() []WebData`

GetWebData returns the WebData field if non-nil, zero value otherwise.

### GetWebDataOk

`func (o *BusinessLineInfo) GetWebDataOk() (*[]WebData, bool)`

GetWebDataOk returns a tuple with the WebData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebData

`func (o *BusinessLineInfo) SetWebData(v []WebData)`

SetWebData sets WebData field to given value.

### HasWebData

`func (o *BusinessLineInfo) HasWebData() bool`

HasWebData returns a boolean if a field has been set.

### GetWebDataExemption

`func (o *BusinessLineInfo) GetWebDataExemption() WebDataExemption`

GetWebDataExemption returns the WebDataExemption field if non-nil, zero value otherwise.

### GetWebDataExemptionOk

`func (o *BusinessLineInfo) GetWebDataExemptionOk() (*WebDataExemption, bool)`

GetWebDataExemptionOk returns a tuple with the WebDataExemption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebDataExemption

`func (o *BusinessLineInfo) SetWebDataExemption(v WebDataExemption)`

SetWebDataExemption sets WebDataExemption field to given value.

### HasWebDataExemption

`func (o *BusinessLineInfo) HasWebDataExemption() bool`

HasWebDataExemption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


