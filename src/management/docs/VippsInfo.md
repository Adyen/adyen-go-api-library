# VippsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logo** | **string** | Vipps logo. Format: Base64-encoded string. | 
**SubscriptionCancelUrl** | Pointer to **string** | Vipps subscription cancel url (required in case of [recurring payments](https://docs.adyen.com/online-payments/tokenization)) | [optional] 

## Methods

### NewVippsInfo

`func NewVippsInfo(logo string, ) *VippsInfo`

NewVippsInfo instantiates a new VippsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVippsInfoWithDefaults

`func NewVippsInfoWithDefaults() *VippsInfo`

NewVippsInfoWithDefaults instantiates a new VippsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogo

`func (o *VippsInfo) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *VippsInfo) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *VippsInfo) SetLogo(v string)`

SetLogo sets Logo field to given value.


### GetSubscriptionCancelUrl

`func (o *VippsInfo) GetSubscriptionCancelUrl() string`

GetSubscriptionCancelUrl returns the SubscriptionCancelUrl field if non-nil, zero value otherwise.

### GetSubscriptionCancelUrlOk

`func (o *VippsInfo) GetSubscriptionCancelUrlOk() (*string, bool)`

GetSubscriptionCancelUrlOk returns a tuple with the SubscriptionCancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionCancelUrl

`func (o *VippsInfo) SetSubscriptionCancelUrl(v string)`

SetSubscriptionCancelUrl sets SubscriptionCancelUrl field to given value.

### HasSubscriptionCancelUrl

`func (o *VippsInfo) HasSubscriptionCancelUrl() bool`

HasSubscriptionCancelUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


