# CreateOnboardingLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminUser** | [**User2**](User2.md) |  | 
**CompanyId** | Pointer to **string** | Id of the company to be onboarded | [optional] 
**Contacts** | Pointer to [**[]Contact2**](Contact2.md) | Contacts | [optional] 
**Country** | **string** | Country | 
**MerchantId** | Pointer to **string** | Id of the merchant to be onboarded | [optional] 
**PricingPlan** | Pointer to **string** | Pricing Plan for the merchant | [optional] 
**ReturnUrl** | Pointer to **string** | If supplied, merchant will be redirected to this URL after onboarding form submission | [optional] 
**TermsAndConditionsUrl** | Pointer to **string** | Terms and conditions URL | [optional] 
**WebAddress** | Pointer to **string** | Web address | [optional] 
**Webhook** | Pointer to [**Webhook2**](Webhook2.md) |  | [optional] 

## Methods

### NewCreateOnboardingLinkRequest

`func NewCreateOnboardingLinkRequest(adminUser User2, country string, ) *CreateOnboardingLinkRequest`

NewCreateOnboardingLinkRequest instantiates a new CreateOnboardingLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOnboardingLinkRequestWithDefaults

`func NewCreateOnboardingLinkRequestWithDefaults() *CreateOnboardingLinkRequest`

NewCreateOnboardingLinkRequestWithDefaults instantiates a new CreateOnboardingLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminUser

`func (o *CreateOnboardingLinkRequest) GetAdminUser() User2`

GetAdminUser returns the AdminUser field if non-nil, zero value otherwise.

### GetAdminUserOk

`func (o *CreateOnboardingLinkRequest) GetAdminUserOk() (*User2, bool)`

GetAdminUserOk returns a tuple with the AdminUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUser

`func (o *CreateOnboardingLinkRequest) SetAdminUser(v User2)`

SetAdminUser sets AdminUser field to given value.


### GetCompanyId

`func (o *CreateOnboardingLinkRequest) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CreateOnboardingLinkRequest) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CreateOnboardingLinkRequest) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CreateOnboardingLinkRequest) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetContacts

`func (o *CreateOnboardingLinkRequest) GetContacts() []Contact2`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *CreateOnboardingLinkRequest) GetContactsOk() (*[]Contact2, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *CreateOnboardingLinkRequest) SetContacts(v []Contact2)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *CreateOnboardingLinkRequest) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetCountry

`func (o *CreateOnboardingLinkRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CreateOnboardingLinkRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CreateOnboardingLinkRequest) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetMerchantId

`func (o *CreateOnboardingLinkRequest) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *CreateOnboardingLinkRequest) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *CreateOnboardingLinkRequest) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *CreateOnboardingLinkRequest) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetPricingPlan

`func (o *CreateOnboardingLinkRequest) GetPricingPlan() string`

GetPricingPlan returns the PricingPlan field if non-nil, zero value otherwise.

### GetPricingPlanOk

`func (o *CreateOnboardingLinkRequest) GetPricingPlanOk() (*string, bool)`

GetPricingPlanOk returns a tuple with the PricingPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingPlan

`func (o *CreateOnboardingLinkRequest) SetPricingPlan(v string)`

SetPricingPlan sets PricingPlan field to given value.

### HasPricingPlan

`func (o *CreateOnboardingLinkRequest) HasPricingPlan() bool`

HasPricingPlan returns a boolean if a field has been set.

### GetReturnUrl

`func (o *CreateOnboardingLinkRequest) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *CreateOnboardingLinkRequest) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *CreateOnboardingLinkRequest) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *CreateOnboardingLinkRequest) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetTermsAndConditionsUrl

`func (o *CreateOnboardingLinkRequest) GetTermsAndConditionsUrl() string`

GetTermsAndConditionsUrl returns the TermsAndConditionsUrl field if non-nil, zero value otherwise.

### GetTermsAndConditionsUrlOk

`func (o *CreateOnboardingLinkRequest) GetTermsAndConditionsUrlOk() (*string, bool)`

GetTermsAndConditionsUrlOk returns a tuple with the TermsAndConditionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsAndConditionsUrl

`func (o *CreateOnboardingLinkRequest) SetTermsAndConditionsUrl(v string)`

SetTermsAndConditionsUrl sets TermsAndConditionsUrl field to given value.

### HasTermsAndConditionsUrl

`func (o *CreateOnboardingLinkRequest) HasTermsAndConditionsUrl() bool`

HasTermsAndConditionsUrl returns a boolean if a field has been set.

### GetWebAddress

`func (o *CreateOnboardingLinkRequest) GetWebAddress() string`

GetWebAddress returns the WebAddress field if non-nil, zero value otherwise.

### GetWebAddressOk

`func (o *CreateOnboardingLinkRequest) GetWebAddressOk() (*string, bool)`

GetWebAddressOk returns a tuple with the WebAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAddress

`func (o *CreateOnboardingLinkRequest) SetWebAddress(v string)`

SetWebAddress sets WebAddress field to given value.

### HasWebAddress

`func (o *CreateOnboardingLinkRequest) HasWebAddress() bool`

HasWebAddress returns a boolean if a field has been set.

### GetWebhook

`func (o *CreateOnboardingLinkRequest) GetWebhook() Webhook2`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *CreateOnboardingLinkRequest) GetWebhookOk() (*Webhook2, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *CreateOnboardingLinkRequest) SetWebhook(v Webhook2)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *CreateOnboardingLinkRequest) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


