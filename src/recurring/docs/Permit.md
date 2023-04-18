# Permit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerId** | Pointer to **string** | Partner ID (when using the permit-per-partner token sharing model). | [optional] 
**ProfileReference** | Pointer to **string** | The profile to apply to this permit (when using the shared permits model). | [optional] 
**Restriction** | Pointer to [**PermitRestriction**](PermitRestriction.md) |  | [optional] 
**ResultKey** | Pointer to **string** | The key to link permit requests to permit results. | [optional] 
**ValidTillDate** | Pointer to **time.Time** | The expiry date for this permit. | [optional] 

## Methods

### NewPermit

`func NewPermit() *Permit`

NewPermit instantiates a new Permit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermitWithDefaults

`func NewPermitWithDefaults() *Permit`

NewPermitWithDefaults instantiates a new Permit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartnerId

`func (o *Permit) GetPartnerId() string`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *Permit) GetPartnerIdOk() (*string, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *Permit) SetPartnerId(v string)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *Permit) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetProfileReference

`func (o *Permit) GetProfileReference() string`

GetProfileReference returns the ProfileReference field if non-nil, zero value otherwise.

### GetProfileReferenceOk

`func (o *Permit) GetProfileReferenceOk() (*string, bool)`

GetProfileReferenceOk returns a tuple with the ProfileReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileReference

`func (o *Permit) SetProfileReference(v string)`

SetProfileReference sets ProfileReference field to given value.

### HasProfileReference

`func (o *Permit) HasProfileReference() bool`

HasProfileReference returns a boolean if a field has been set.

### GetRestriction

`func (o *Permit) GetRestriction() PermitRestriction`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *Permit) GetRestrictionOk() (*PermitRestriction, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *Permit) SetRestriction(v PermitRestriction)`

SetRestriction sets Restriction field to given value.

### HasRestriction

`func (o *Permit) HasRestriction() bool`

HasRestriction returns a boolean if a field has been set.

### GetResultKey

`func (o *Permit) GetResultKey() string`

GetResultKey returns the ResultKey field if non-nil, zero value otherwise.

### GetResultKeyOk

`func (o *Permit) GetResultKeyOk() (*string, bool)`

GetResultKeyOk returns a tuple with the ResultKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultKey

`func (o *Permit) SetResultKey(v string)`

SetResultKey sets ResultKey field to given value.

### HasResultKey

`func (o *Permit) HasResultKey() bool`

HasResultKey returns a boolean if a field has been set.

### GetValidTillDate

`func (o *Permit) GetValidTillDate() time.Time`

GetValidTillDate returns the ValidTillDate field if non-nil, zero value otherwise.

### GetValidTillDateOk

`func (o *Permit) GetValidTillDateOk() (*time.Time, bool)`

GetValidTillDateOk returns a tuple with the ValidTillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTillDate

`func (o *Permit) SetValidTillDate(v time.Time)`

SetValidTillDate sets ValidTillDate field to given value.

### HasValidTillDate

`func (o *Permit) HasValidTillDate() bool`

HasValidTillDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


