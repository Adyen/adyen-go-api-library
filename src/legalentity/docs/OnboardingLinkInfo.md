# OnboardingLinkInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locale** | Pointer to **string** | The language that will be used for the page, specified by a combination of two letter [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language and [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country codes. See [possible values](https://docs.adyen.com/marketplaces-and-platforms/onboarding/hosted#supported-languages).   If not specified in the request or if the language is not supported, the page uses the browser language. If the browser language is not supported, the page uses **en-US** by default. | [optional] 
**RedirectUrl** | Pointer to **string** | The URL where the user is redirected after they complete hosted onboarding. | [optional] 
**Settings** | Pointer to **map[string]bool** | The enabled/disabled settings in the hosted onboarding webpage. | [optional] 
**ThemeId** | Pointer to **string** | The unique identifier of the hosted onboarding theme. | [optional] 

## Methods

### NewOnboardingLinkInfo

`func NewOnboardingLinkInfo() *OnboardingLinkInfo`

NewOnboardingLinkInfo instantiates a new OnboardingLinkInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingLinkInfoWithDefaults

`func NewOnboardingLinkInfoWithDefaults() *OnboardingLinkInfo`

NewOnboardingLinkInfoWithDefaults instantiates a new OnboardingLinkInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocale

`func (o *OnboardingLinkInfo) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *OnboardingLinkInfo) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *OnboardingLinkInfo) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *OnboardingLinkInfo) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *OnboardingLinkInfo) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *OnboardingLinkInfo) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *OnboardingLinkInfo) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *OnboardingLinkInfo) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetSettings

`func (o *OnboardingLinkInfo) GetSettings() map[string]bool`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *OnboardingLinkInfo) GetSettingsOk() (*map[string]bool, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *OnboardingLinkInfo) SetSettings(v map[string]bool)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *OnboardingLinkInfo) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetThemeId

`func (o *OnboardingLinkInfo) GetThemeId() string`

GetThemeId returns the ThemeId field if non-nil, zero value otherwise.

### GetThemeIdOk

`func (o *OnboardingLinkInfo) GetThemeIdOk() (*string, bool)`

GetThemeIdOk returns a tuple with the ThemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeId

`func (o *OnboardingLinkInfo) SetThemeId(v string)`

SetThemeId sets ThemeId field to given value.

### HasThemeId

`func (o *OnboardingLinkInfo) HasThemeId() bool`

HasThemeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


