# WebhookLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to [**LinksElement**](LinksElement.md) |  | [optional] 
**GenerateHmac** | [**LinksElement**](LinksElement.md) |  | 
**Merchant** | Pointer to [**LinksElement**](LinksElement.md) |  | [optional] 
**Self** | [**LinksElement**](LinksElement.md) |  | 
**TestWebhook** | [**LinksElement**](LinksElement.md) |  | 

## Methods

### NewWebhookLinks

`func NewWebhookLinks(generateHmac LinksElement, self LinksElement, testWebhook LinksElement, ) *WebhookLinks`

NewWebhookLinks instantiates a new WebhookLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookLinksWithDefaults

`func NewWebhookLinksWithDefaults() *WebhookLinks`

NewWebhookLinksWithDefaults instantiates a new WebhookLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *WebhookLinks) GetCompany() LinksElement`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *WebhookLinks) GetCompanyOk() (*LinksElement, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *WebhookLinks) SetCompany(v LinksElement)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *WebhookLinks) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetGenerateHmac

`func (o *WebhookLinks) GetGenerateHmac() LinksElement`

GetGenerateHmac returns the GenerateHmac field if non-nil, zero value otherwise.

### GetGenerateHmacOk

`func (o *WebhookLinks) GetGenerateHmacOk() (*LinksElement, bool)`

GetGenerateHmacOk returns a tuple with the GenerateHmac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateHmac

`func (o *WebhookLinks) SetGenerateHmac(v LinksElement)`

SetGenerateHmac sets GenerateHmac field to given value.


### GetMerchant

`func (o *WebhookLinks) GetMerchant() LinksElement`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *WebhookLinks) GetMerchantOk() (*LinksElement, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *WebhookLinks) SetMerchant(v LinksElement)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *WebhookLinks) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetSelf

`func (o *WebhookLinks) GetSelf() LinksElement`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *WebhookLinks) GetSelfOk() (*LinksElement, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *WebhookLinks) SetSelf(v LinksElement)`

SetSelf sets Self field to given value.


### GetTestWebhook

`func (o *WebhookLinks) GetTestWebhook() LinksElement`

GetTestWebhook returns the TestWebhook field if non-nil, zero value otherwise.

### GetTestWebhookOk

`func (o *WebhookLinks) GetTestWebhookOk() (*LinksElement, bool)`

GetTestWebhookOk returns a tuple with the TestWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestWebhook

`func (o *WebhookLinks) SetTestWebhook(v LinksElement)`

SetTestWebhook sets TestWebhook field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


