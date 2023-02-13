# BusinessLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capability** | **string** | The capability for which you are creating the business line. For example, **receivePayments**. | 
**Id** | **string** | The unique identifier of the business line. | [readonly] 
**IndustryCode** | **string** | A code that represents the industry of the legal entity. For example, **4431A** for computer software stores. | 
**LegalEntityId** | **string** | Unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/legalEntities__resParam_id) that owns the business line. | 
**SalesChannels** | Pointer to **[]string** | A list of channels where goods or services are sold. You cannot combine point of sale and ecommerce in one business line.  Possible values and combinations:  - For point of sale: **pos** and **posMoto**  - For ecommerce: **eCommerce** and **ecomMoto**  - For Pay by Link: **payByLink**  Required only in combination with the &#x60;capability&#x60; to **receivePayments** or **receiveFromPlatformPayments**. | [optional] 
**SourceOfFunds** | Pointer to [**SourceOfFunds**](SourceOfFunds.md) |  | [optional] 
**WebData** | Pointer to [**[]WebData**](WebData.md) | List of website URLs where your user&#39;s goods or services are sold. When this is required for a capability but your user does not have an online presence, provide the reason in the &#x60;webDataExemption&#x60; object. | [optional] 
**WebDataExemption** | Pointer to [**WebDataExemption**](WebDataExemption.md) |  | [optional] 

## Methods

### NewBusinessLine

`func NewBusinessLine(capability string, id string, industryCode string, legalEntityId string, ) *BusinessLine`

NewBusinessLine instantiates a new BusinessLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessLineWithDefaults

`func NewBusinessLineWithDefaults() *BusinessLine`

NewBusinessLineWithDefaults instantiates a new BusinessLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapability

`func (o *BusinessLine) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *BusinessLine) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *BusinessLine) SetCapability(v string)`

SetCapability sets Capability field to given value.


### GetId

`func (o *BusinessLine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BusinessLine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BusinessLine) SetId(v string)`

SetId sets Id field to given value.


### GetIndustryCode

`func (o *BusinessLine) GetIndustryCode() string`

GetIndustryCode returns the IndustryCode field if non-nil, zero value otherwise.

### GetIndustryCodeOk

`func (o *BusinessLine) GetIndustryCodeOk() (*string, bool)`

GetIndustryCodeOk returns a tuple with the IndustryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryCode

`func (o *BusinessLine) SetIndustryCode(v string)`

SetIndustryCode sets IndustryCode field to given value.


### GetLegalEntityId

`func (o *BusinessLine) GetLegalEntityId() string`

GetLegalEntityId returns the LegalEntityId field if non-nil, zero value otherwise.

### GetLegalEntityIdOk

`func (o *BusinessLine) GetLegalEntityIdOk() (*string, bool)`

GetLegalEntityIdOk returns a tuple with the LegalEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalEntityId

`func (o *BusinessLine) SetLegalEntityId(v string)`

SetLegalEntityId sets LegalEntityId field to given value.


### GetSalesChannels

`func (o *BusinessLine) GetSalesChannels() []string`

GetSalesChannels returns the SalesChannels field if non-nil, zero value otherwise.

### GetSalesChannelsOk

`func (o *BusinessLine) GetSalesChannelsOk() (*[]string, bool)`

GetSalesChannelsOk returns a tuple with the SalesChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesChannels

`func (o *BusinessLine) SetSalesChannels(v []string)`

SetSalesChannels sets SalesChannels field to given value.

### HasSalesChannels

`func (o *BusinessLine) HasSalesChannels() bool`

HasSalesChannels returns a boolean if a field has been set.

### GetSourceOfFunds

`func (o *BusinessLine) GetSourceOfFunds() SourceOfFunds`

GetSourceOfFunds returns the SourceOfFunds field if non-nil, zero value otherwise.

### GetSourceOfFundsOk

`func (o *BusinessLine) GetSourceOfFundsOk() (*SourceOfFunds, bool)`

GetSourceOfFundsOk returns a tuple with the SourceOfFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfFunds

`func (o *BusinessLine) SetSourceOfFunds(v SourceOfFunds)`

SetSourceOfFunds sets SourceOfFunds field to given value.

### HasSourceOfFunds

`func (o *BusinessLine) HasSourceOfFunds() bool`

HasSourceOfFunds returns a boolean if a field has been set.

### GetWebData

`func (o *BusinessLine) GetWebData() []WebData`

GetWebData returns the WebData field if non-nil, zero value otherwise.

### GetWebDataOk

`func (o *BusinessLine) GetWebDataOk() (*[]WebData, bool)`

GetWebDataOk returns a tuple with the WebData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebData

`func (o *BusinessLine) SetWebData(v []WebData)`

SetWebData sets WebData field to given value.

### HasWebData

`func (o *BusinessLine) HasWebData() bool`

HasWebData returns a boolean if a field has been set.

### GetWebDataExemption

`func (o *BusinessLine) GetWebDataExemption() WebDataExemption`

GetWebDataExemption returns the WebDataExemption field if non-nil, zero value otherwise.

### GetWebDataExemptionOk

`func (o *BusinessLine) GetWebDataExemptionOk() (*WebDataExemption, bool)`

GetWebDataExemptionOk returns a tuple with the WebDataExemption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebDataExemption

`func (o *BusinessLine) SetWebDataExemption(v WebDataExemption)`

SetWebDataExemption sets WebDataExemption field to given value.

### HasWebDataExemption

`func (o *BusinessLine) HasWebDataExemption() bool`

HasWebDataExemption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


