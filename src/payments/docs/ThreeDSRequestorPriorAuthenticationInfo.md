# ThreeDSRequestorPriorAuthenticationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreeDSReqPriorAuthData** | Pointer to **string** | Data that documents and supports a specific authentication process. Maximum length: 2048 bytes. | [optional] 
**ThreeDSReqPriorAuthMethod** | Pointer to **string** | Mechanism used by the Cardholder to previously authenticate to the 3DS Requestor. Allowed values: * **01** — Frictionless authentication occurred by ACS. * **02** — Cardholder challenge occurred by ACS. * **03** — AVS verified. * **04** — Other issuer methods. | [optional] 
**ThreeDSReqPriorAuthTimestamp** | Pointer to **string** | Date and time in UTC of the prior cardholder authentication. Format: YYYYMMDDHHMM | [optional] 
**ThreeDSReqPriorRef** | Pointer to **string** | This data element provides additional information to the ACS to determine the best approach for handing a request. This data element contains an ACS Transaction ID for a prior authenticated transaction. For example, the first recurring transaction that was authenticated with the cardholder. Length: 30 characters. | [optional] 

## Methods

### NewThreeDSRequestorPriorAuthenticationInfo

`func NewThreeDSRequestorPriorAuthenticationInfo() *ThreeDSRequestorPriorAuthenticationInfo`

NewThreeDSRequestorPriorAuthenticationInfo instantiates a new ThreeDSRequestorPriorAuthenticationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSRequestorPriorAuthenticationInfoWithDefaults

`func NewThreeDSRequestorPriorAuthenticationInfoWithDefaults() *ThreeDSRequestorPriorAuthenticationInfo`

NewThreeDSRequestorPriorAuthenticationInfoWithDefaults instantiates a new ThreeDSRequestorPriorAuthenticationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreeDSReqPriorAuthData

`func (o *ThreeDSRequestorPriorAuthenticationInfo) GetThreeDSReqPriorAuthData() string`

GetThreeDSReqPriorAuthData returns the ThreeDSReqPriorAuthData field if non-nil, zero value otherwise.

### GetThreeDSReqPriorAuthDataOk

`func (o *ThreeDSRequestorPriorAuthenticationInfo) GetThreeDSReqPriorAuthDataOk() (*string, bool)`

GetThreeDSReqPriorAuthDataOk returns a tuple with the ThreeDSReqPriorAuthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSReqPriorAuthData

`func (o *ThreeDSRequestorPriorAuthenticationInfo) SetThreeDSReqPriorAuthData(v string)`

SetThreeDSReqPriorAuthData sets ThreeDSReqPriorAuthData field to given value.

### HasThreeDSReqPriorAuthData

`func (o *ThreeDSRequestorPriorAuthenticationInfo) HasThreeDSReqPriorAuthData() bool`

HasThreeDSReqPriorAuthData returns a boolean if a field has been set.

### GetThreeDSReqPriorAuthMethod

`func (o *ThreeDSRequestorPriorAuthenticationInfo) GetThreeDSReqPriorAuthMethod() string`

GetThreeDSReqPriorAuthMethod returns the ThreeDSReqPriorAuthMethod field if non-nil, zero value otherwise.

### GetThreeDSReqPriorAuthMethodOk

`func (o *ThreeDSRequestorPriorAuthenticationInfo) GetThreeDSReqPriorAuthMethodOk() (*string, bool)`

GetThreeDSReqPriorAuthMethodOk returns a tuple with the ThreeDSReqPriorAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSReqPriorAuthMethod

`func (o *ThreeDSRequestorPriorAuthenticationInfo) SetThreeDSReqPriorAuthMethod(v string)`

SetThreeDSReqPriorAuthMethod sets ThreeDSReqPriorAuthMethod field to given value.

### HasThreeDSReqPriorAuthMethod

`func (o *ThreeDSRequestorPriorAuthenticationInfo) HasThreeDSReqPriorAuthMethod() bool`

HasThreeDSReqPriorAuthMethod returns a boolean if a field has been set.

### GetThreeDSReqPriorAuthTimestamp

`func (o *ThreeDSRequestorPriorAuthenticationInfo) GetThreeDSReqPriorAuthTimestamp() string`

GetThreeDSReqPriorAuthTimestamp returns the ThreeDSReqPriorAuthTimestamp field if non-nil, zero value otherwise.

### GetThreeDSReqPriorAuthTimestampOk

`func (o *ThreeDSRequestorPriorAuthenticationInfo) GetThreeDSReqPriorAuthTimestampOk() (*string, bool)`

GetThreeDSReqPriorAuthTimestampOk returns a tuple with the ThreeDSReqPriorAuthTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSReqPriorAuthTimestamp

`func (o *ThreeDSRequestorPriorAuthenticationInfo) SetThreeDSReqPriorAuthTimestamp(v string)`

SetThreeDSReqPriorAuthTimestamp sets ThreeDSReqPriorAuthTimestamp field to given value.

### HasThreeDSReqPriorAuthTimestamp

`func (o *ThreeDSRequestorPriorAuthenticationInfo) HasThreeDSReqPriorAuthTimestamp() bool`

HasThreeDSReqPriorAuthTimestamp returns a boolean if a field has been set.

### GetThreeDSReqPriorRef

`func (o *ThreeDSRequestorPriorAuthenticationInfo) GetThreeDSReqPriorRef() string`

GetThreeDSReqPriorRef returns the ThreeDSReqPriorRef field if non-nil, zero value otherwise.

### GetThreeDSReqPriorRefOk

`func (o *ThreeDSRequestorPriorAuthenticationInfo) GetThreeDSReqPriorRefOk() (*string, bool)`

GetThreeDSReqPriorRefOk returns a tuple with the ThreeDSReqPriorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSReqPriorRef

`func (o *ThreeDSRequestorPriorAuthenticationInfo) SetThreeDSReqPriorRef(v string)`

SetThreeDSReqPriorRef sets ThreeDSReqPriorRef field to given value.

### HasThreeDSReqPriorRef

`func (o *ThreeDSRequestorPriorAuthenticationInfo) HasThreeDSReqPriorRef() bool`

HasThreeDSReqPriorRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


