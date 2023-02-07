# TermsOfServiceAcceptanceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedBy** | Pointer to **string** | The unique identifier of the user that accepted the Terms of Service. | [optional] 
**AcceptedFor** | Pointer to **string** | The unique identifier of the legal entity for which the Terms of Service are accepted. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date when the Terms of Service were accepted. | [optional] 
**Id** | Pointer to **string** | An Adyen-generated reference for the accepted Terms of Service. | [optional] 
**Type** | Pointer to **string** | The type of Terms of Service. | [optional] 

## Methods

### NewTermsOfServiceAcceptanceInfo

`func NewTermsOfServiceAcceptanceInfo() *TermsOfServiceAcceptanceInfo`

NewTermsOfServiceAcceptanceInfo instantiates a new TermsOfServiceAcceptanceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTermsOfServiceAcceptanceInfoWithDefaults

`func NewTermsOfServiceAcceptanceInfoWithDefaults() *TermsOfServiceAcceptanceInfo`

NewTermsOfServiceAcceptanceInfoWithDefaults instantiates a new TermsOfServiceAcceptanceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedBy

`func (o *TermsOfServiceAcceptanceInfo) GetAcceptedBy() string`

GetAcceptedBy returns the AcceptedBy field if non-nil, zero value otherwise.

### GetAcceptedByOk

`func (o *TermsOfServiceAcceptanceInfo) GetAcceptedByOk() (*string, bool)`

GetAcceptedByOk returns a tuple with the AcceptedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedBy

`func (o *TermsOfServiceAcceptanceInfo) SetAcceptedBy(v string)`

SetAcceptedBy sets AcceptedBy field to given value.

### HasAcceptedBy

`func (o *TermsOfServiceAcceptanceInfo) HasAcceptedBy() bool`

HasAcceptedBy returns a boolean if a field has been set.

### GetAcceptedFor

`func (o *TermsOfServiceAcceptanceInfo) GetAcceptedFor() string`

GetAcceptedFor returns the AcceptedFor field if non-nil, zero value otherwise.

### GetAcceptedForOk

`func (o *TermsOfServiceAcceptanceInfo) GetAcceptedForOk() (*string, bool)`

GetAcceptedForOk returns a tuple with the AcceptedFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedFor

`func (o *TermsOfServiceAcceptanceInfo) SetAcceptedFor(v string)`

SetAcceptedFor sets AcceptedFor field to given value.

### HasAcceptedFor

`func (o *TermsOfServiceAcceptanceInfo) HasAcceptedFor() bool`

HasAcceptedFor returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TermsOfServiceAcceptanceInfo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TermsOfServiceAcceptanceInfo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TermsOfServiceAcceptanceInfo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TermsOfServiceAcceptanceInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *TermsOfServiceAcceptanceInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TermsOfServiceAcceptanceInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TermsOfServiceAcceptanceInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TermsOfServiceAcceptanceInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TermsOfServiceAcceptanceInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TermsOfServiceAcceptanceInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TermsOfServiceAcceptanceInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TermsOfServiceAcceptanceInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


