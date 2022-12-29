# OnboardingLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminUser** | Pointer to [**User2**](User2.md) |  | [optional] 
**CompanyId** | Pointer to **string** | Company id | [optional] 
**Contacts** | Pointer to [**[]Contact2**](Contact2.md) | Contacts | [optional] 
**Country** | Pointer to **string** | Country | [optional] 
**ExpiryDate** | Pointer to **string** | Onboarding URL expiry date | [optional] 
**Id** | Pointer to **string** | Onboarding link id | [optional] 
**MerchantId** | Pointer to **string** | Merchant id | [optional] 
**ReturnUrl** | Pointer to **string** | If supplied, merchant will be redirected to this URL after onboarding form submission | [optional] 
**TermsAndConditionsUrl** | Pointer to **string** | Terms and conditions URL | [optional] 
**Url** | Pointer to **string** | Onboarding URL | [optional] 
**WebAddress** | Pointer to **string** | Web address | [optional] 
**Webhook** | Pointer to [**Webhook2**](Webhook2.md) |  | [optional] 

## Methods

### NewOnboardingLink

`func NewOnboardingLink() *OnboardingLink`

NewOnboardingLink instantiates a new OnboardingLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingLinkWithDefaults

`func NewOnboardingLinkWithDefaults() *OnboardingLink`

NewOnboardingLinkWithDefaults instantiates a new OnboardingLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminUser

`func (o *OnboardingLink) GetAdminUser() User2`

GetAdminUser returns the AdminUser field if non-nil, zero value otherwise.

### GetAdminUserOk

`func (o *OnboardingLink) GetAdminUserOk() (*User2, bool)`

GetAdminUserOk returns a tuple with the AdminUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUser

`func (o *OnboardingLink) SetAdminUser(v User2)`

SetAdminUser sets AdminUser field to given value.

### HasAdminUser

`func (o *OnboardingLink) HasAdminUser() bool`

HasAdminUser returns a boolean if a field has been set.

### GetCompanyId

`func (o *OnboardingLink) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *OnboardingLink) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *OnboardingLink) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *OnboardingLink) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetContacts

`func (o *OnboardingLink) GetContacts() []Contact2`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *OnboardingLink) GetContactsOk() (*[]Contact2, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *OnboardingLink) SetContacts(v []Contact2)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *OnboardingLink) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetCountry

`func (o *OnboardingLink) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *OnboardingLink) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *OnboardingLink) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *OnboardingLink) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetExpiryDate

`func (o *OnboardingLink) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *OnboardingLink) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *OnboardingLink) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *OnboardingLink) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetId

`func (o *OnboardingLink) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OnboardingLink) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OnboardingLink) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OnboardingLink) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantId

`func (o *OnboardingLink) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *OnboardingLink) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *OnboardingLink) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *OnboardingLink) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetReturnUrl

`func (o *OnboardingLink) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *OnboardingLink) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *OnboardingLink) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *OnboardingLink) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetTermsAndConditionsUrl

`func (o *OnboardingLink) GetTermsAndConditionsUrl() string`

GetTermsAndConditionsUrl returns the TermsAndConditionsUrl field if non-nil, zero value otherwise.

### GetTermsAndConditionsUrlOk

`func (o *OnboardingLink) GetTermsAndConditionsUrlOk() (*string, bool)`

GetTermsAndConditionsUrlOk returns a tuple with the TermsAndConditionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsAndConditionsUrl

`func (o *OnboardingLink) SetTermsAndConditionsUrl(v string)`

SetTermsAndConditionsUrl sets TermsAndConditionsUrl field to given value.

### HasTermsAndConditionsUrl

`func (o *OnboardingLink) HasTermsAndConditionsUrl() bool`

HasTermsAndConditionsUrl returns a boolean if a field has been set.

### GetUrl

`func (o *OnboardingLink) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OnboardingLink) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OnboardingLink) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OnboardingLink) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWebAddress

`func (o *OnboardingLink) GetWebAddress() string`

GetWebAddress returns the WebAddress field if non-nil, zero value otherwise.

### GetWebAddressOk

`func (o *OnboardingLink) GetWebAddressOk() (*string, bool)`

GetWebAddressOk returns a tuple with the WebAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAddress

`func (o *OnboardingLink) SetWebAddress(v string)`

SetWebAddress sets WebAddress field to given value.

### HasWebAddress

`func (o *OnboardingLink) HasWebAddress() bool`

HasWebAddress returns a boolean if a field has been set.

### GetWebhook

`func (o *OnboardingLink) GetWebhook() Webhook2`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *OnboardingLink) GetWebhookOk() (*Webhook2, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *OnboardingLink) SetWebhook(v Webhook2)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *OnboardingLink) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


