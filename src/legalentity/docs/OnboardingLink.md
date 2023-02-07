# OnboardingLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The URL of the hosted onboarding page where you need to redirect your user. This URL expires after 4 minutes and can only be used once.  If the link expires, you need to create a new link. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


