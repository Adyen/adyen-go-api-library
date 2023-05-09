# VerificationDeadline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | **[]string** | The names of the capabilities to be disallowed. | [readonly] 
**ExpiresAt** | **time.Time** | The date that verification is due by before capabilities are disallowed. | [readonly] 

## Methods

### NewVerificationDeadline

`func NewVerificationDeadline(capabilities []string, expiresAt time.Time, ) *VerificationDeadline`

NewVerificationDeadline instantiates a new VerificationDeadline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationDeadlineWithDefaults

`func NewVerificationDeadlineWithDefaults() *VerificationDeadline`

NewVerificationDeadlineWithDefaults instantiates a new VerificationDeadline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *VerificationDeadline) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *VerificationDeadline) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *VerificationDeadline) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.


### GetExpiresAt

`func (o *VerificationDeadline) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *VerificationDeadline) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *VerificationDeadline) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


