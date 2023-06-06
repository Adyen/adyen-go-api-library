# SweepconfigurationwebhookData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The unique identifier of the balance account for which the sweep was configured. | [optional] 
**BalancePlatform** | Pointer to **string** | The unique identifier of the balance platform. | [optional] 
**Sweep** | Pointer to [**SweepConfigurationV2**](SweepConfigurationV2.md) |  | [optional] 

## Methods

### NewSweepconfigurationwebhookData

`func NewSweepconfigurationwebhookData() *SweepconfigurationwebhookData`

NewSweepconfigurationwebhookData instantiates a new SweepconfigurationwebhookData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSweepconfigurationwebhookDataWithDefaults

`func NewSweepconfigurationwebhookDataWithDefaults() *SweepconfigurationwebhookData`

NewSweepconfigurationwebhookDataWithDefaults instantiates a new SweepconfigurationwebhookData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SweepconfigurationwebhookData) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SweepconfigurationwebhookData) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SweepconfigurationwebhookData) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *SweepconfigurationwebhookData) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetBalancePlatform

`func (o *SweepconfigurationwebhookData) GetBalancePlatform() string`

GetBalancePlatform returns the BalancePlatform field if non-nil, zero value otherwise.

### GetBalancePlatformOk

`func (o *SweepconfigurationwebhookData) GetBalancePlatformOk() (*string, bool)`

GetBalancePlatformOk returns a tuple with the BalancePlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancePlatform

`func (o *SweepconfigurationwebhookData) SetBalancePlatform(v string)`

SetBalancePlatform sets BalancePlatform field to given value.

### HasBalancePlatform

`func (o *SweepconfigurationwebhookData) HasBalancePlatform() bool`

HasBalancePlatform returns a boolean if a field has been set.

### GetSweep

`func (o *SweepconfigurationwebhookData) GetSweep() SweepConfigurationV2`

GetSweep returns the Sweep field if non-nil, zero value otherwise.

### GetSweepOk

`func (o *SweepconfigurationwebhookData) GetSweepOk() (*SweepConfigurationV2, bool)`

GetSweepOk returns a tuple with the Sweep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSweep

`func (o *SweepconfigurationwebhookData) SetSweep(v SweepConfigurationV2)`

SetSweep sets Sweep field to given value.

### HasSweep

`func (o *SweepconfigurationwebhookData) HasSweep() bool`

HasSweep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


