# CapabilitySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountPerIndustry** | Pointer to [**map[string]Amount**](Amount.md) |  | [optional] 
**AuthorizedCardUsers** | Pointer to **bool** |  | [optional] 
**FundingSource** | Pointer to **[]string** |  | [optional] 
**Interval** | Pointer to **string** |  | [optional] 
**MaxAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 

## Methods

### NewCapabilitySettings

`func NewCapabilitySettings() *CapabilitySettings`

NewCapabilitySettings instantiates a new CapabilitySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitySettingsWithDefaults

`func NewCapabilitySettingsWithDefaults() *CapabilitySettings`

NewCapabilitySettingsWithDefaults instantiates a new CapabilitySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountPerIndustry

`func (o *CapabilitySettings) GetAmountPerIndustry() map[string]Amount`

GetAmountPerIndustry returns the AmountPerIndustry field if non-nil, zero value otherwise.

### GetAmountPerIndustryOk

`func (o *CapabilitySettings) GetAmountPerIndustryOk() (*map[string]Amount, bool)`

GetAmountPerIndustryOk returns a tuple with the AmountPerIndustry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPerIndustry

`func (o *CapabilitySettings) SetAmountPerIndustry(v map[string]Amount)`

SetAmountPerIndustry sets AmountPerIndustry field to given value.

### HasAmountPerIndustry

`func (o *CapabilitySettings) HasAmountPerIndustry() bool`

HasAmountPerIndustry returns a boolean if a field has been set.

### GetAuthorizedCardUsers

`func (o *CapabilitySettings) GetAuthorizedCardUsers() bool`

GetAuthorizedCardUsers returns the AuthorizedCardUsers field if non-nil, zero value otherwise.

### GetAuthorizedCardUsersOk

`func (o *CapabilitySettings) GetAuthorizedCardUsersOk() (*bool, bool)`

GetAuthorizedCardUsersOk returns a tuple with the AuthorizedCardUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedCardUsers

`func (o *CapabilitySettings) SetAuthorizedCardUsers(v bool)`

SetAuthorizedCardUsers sets AuthorizedCardUsers field to given value.

### HasAuthorizedCardUsers

`func (o *CapabilitySettings) HasAuthorizedCardUsers() bool`

HasAuthorizedCardUsers returns a boolean if a field has been set.

### GetFundingSource

`func (o *CapabilitySettings) GetFundingSource() []string`

GetFundingSource returns the FundingSource field if non-nil, zero value otherwise.

### GetFundingSourceOk

`func (o *CapabilitySettings) GetFundingSourceOk() (*[]string, bool)`

GetFundingSourceOk returns a tuple with the FundingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSource

`func (o *CapabilitySettings) SetFundingSource(v []string)`

SetFundingSource sets FundingSource field to given value.

### HasFundingSource

`func (o *CapabilitySettings) HasFundingSource() bool`

HasFundingSource returns a boolean if a field has been set.

### GetInterval

`func (o *CapabilitySettings) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CapabilitySettings) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CapabilitySettings) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *CapabilitySettings) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetMaxAmount

`func (o *CapabilitySettings) GetMaxAmount() Amount`

GetMaxAmount returns the MaxAmount field if non-nil, zero value otherwise.

### GetMaxAmountOk

`func (o *CapabilitySettings) GetMaxAmountOk() (*Amount, bool)`

GetMaxAmountOk returns a tuple with the MaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmount

`func (o *CapabilitySettings) SetMaxAmount(v Amount)`

SetMaxAmount sets MaxAmount field to given value.

### HasMaxAmount

`func (o *CapabilitySettings) HasMaxAmount() bool`

HasMaxAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


