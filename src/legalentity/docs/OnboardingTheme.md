# OnboardingTheme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | The creation date of the theme. | 
**Description** | Pointer to **string** | The description of the theme. | [optional] 
**Id** | **string** | The unique identifier of the theme. | 
**Properties** | **map[string]string** | The properties of the theme. | 
**UpdatedAt** | Pointer to **time.Time** | The date when the theme was last updated. | [optional] 

## Methods

### NewOnboardingTheme

`func NewOnboardingTheme(createdAt time.Time, id string, properties map[string]string, ) *OnboardingTheme`

NewOnboardingTheme instantiates a new OnboardingTheme object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingThemeWithDefaults

`func NewOnboardingThemeWithDefaults() *OnboardingTheme`

NewOnboardingThemeWithDefaults instantiates a new OnboardingTheme object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *OnboardingTheme) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OnboardingTheme) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OnboardingTheme) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *OnboardingTheme) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OnboardingTheme) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OnboardingTheme) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OnboardingTheme) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *OnboardingTheme) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OnboardingTheme) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OnboardingTheme) SetId(v string)`

SetId sets Id field to given value.


### GetProperties

`func (o *OnboardingTheme) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *OnboardingTheme) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *OnboardingTheme) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.


### GetUpdatedAt

`func (o *OnboardingTheme) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OnboardingTheme) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OnboardingTheme) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OnboardingTheme) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


