# OnboardingThemes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **string** | The next page. Only present if there is a next page. | [optional] 
**Previous** | Pointer to **string** | The previous page. Only present if there is a previous page. | [optional] 
**Themes** | [**[]OnboardingTheme**](OnboardingTheme.md) | List of onboarding themes. | 

## Methods

### NewOnboardingThemes

`func NewOnboardingThemes(themes []OnboardingTheme, ) *OnboardingThemes`

NewOnboardingThemes instantiates a new OnboardingThemes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingThemesWithDefaults

`func NewOnboardingThemesWithDefaults() *OnboardingThemes`

NewOnboardingThemesWithDefaults instantiates a new OnboardingThemes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *OnboardingThemes) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *OnboardingThemes) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *OnboardingThemes) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *OnboardingThemes) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *OnboardingThemes) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *OnboardingThemes) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *OnboardingThemes) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *OnboardingThemes) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetThemes

`func (o *OnboardingThemes) GetThemes() []OnboardingTheme`

GetThemes returns the Themes field if non-nil, zero value otherwise.

### GetThemesOk

`func (o *OnboardingThemes) GetThemesOk() (*[]OnboardingTheme, bool)`

GetThemesOk returns a tuple with the Themes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemes

`func (o *OnboardingThemes) SetThemes(v []OnboardingTheme)`

SetThemes sets Themes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


