# ThreeDSRequestorAuthenticationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreeDSReqAuthData** | Pointer to **string** | Data that documents and supports a specific authentication process. Maximum length: 2048 bytes. | [optional] 
**ThreeDSReqAuthMethod** | Pointer to **string** | Mechanism used by the Cardholder to authenticate to the 3DS Requestor. Allowed values: * **01** — No 3DS Requestor authentication occurred (for example, cardholder “logged in” as guest). * **02** — Login to the cardholder account at the 3DS Requestor system using 3DS Requestor’s own credentials. * **03** — Login to the cardholder account at the 3DS Requestor system using federated ID. * **04** — Login to the cardholder account at the 3DS Requestor system using issuer credentials. * **05** — Login to the cardholder account at the 3DS Requestor system using third-party authentication. * **06** — Login to the cardholder account at the 3DS Requestor system using FIDO Authenticator. | [optional] 
**ThreeDSReqAuthTimestamp** | Pointer to **string** | Date and time in UTC of the cardholder authentication. Format: YYYYMMDDHHMM | [optional] 

## Methods

### NewThreeDSRequestorAuthenticationInfo

`func NewThreeDSRequestorAuthenticationInfo() *ThreeDSRequestorAuthenticationInfo`

NewThreeDSRequestorAuthenticationInfo instantiates a new ThreeDSRequestorAuthenticationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSRequestorAuthenticationInfoWithDefaults

`func NewThreeDSRequestorAuthenticationInfoWithDefaults() *ThreeDSRequestorAuthenticationInfo`

NewThreeDSRequestorAuthenticationInfoWithDefaults instantiates a new ThreeDSRequestorAuthenticationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreeDSReqAuthData

`func (o *ThreeDSRequestorAuthenticationInfo) GetThreeDSReqAuthData() string`

GetThreeDSReqAuthData returns the ThreeDSReqAuthData field if non-nil, zero value otherwise.

### GetThreeDSReqAuthDataOk

`func (o *ThreeDSRequestorAuthenticationInfo) GetThreeDSReqAuthDataOk() (*string, bool)`

GetThreeDSReqAuthDataOk returns a tuple with the ThreeDSReqAuthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSReqAuthData

`func (o *ThreeDSRequestorAuthenticationInfo) SetThreeDSReqAuthData(v string)`

SetThreeDSReqAuthData sets ThreeDSReqAuthData field to given value.

### HasThreeDSReqAuthData

`func (o *ThreeDSRequestorAuthenticationInfo) HasThreeDSReqAuthData() bool`

HasThreeDSReqAuthData returns a boolean if a field has been set.

### GetThreeDSReqAuthMethod

`func (o *ThreeDSRequestorAuthenticationInfo) GetThreeDSReqAuthMethod() string`

GetThreeDSReqAuthMethod returns the ThreeDSReqAuthMethod field if non-nil, zero value otherwise.

### GetThreeDSReqAuthMethodOk

`func (o *ThreeDSRequestorAuthenticationInfo) GetThreeDSReqAuthMethodOk() (*string, bool)`

GetThreeDSReqAuthMethodOk returns a tuple with the ThreeDSReqAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSReqAuthMethod

`func (o *ThreeDSRequestorAuthenticationInfo) SetThreeDSReqAuthMethod(v string)`

SetThreeDSReqAuthMethod sets ThreeDSReqAuthMethod field to given value.

### HasThreeDSReqAuthMethod

`func (o *ThreeDSRequestorAuthenticationInfo) HasThreeDSReqAuthMethod() bool`

HasThreeDSReqAuthMethod returns a boolean if a field has been set.

### GetThreeDSReqAuthTimestamp

`func (o *ThreeDSRequestorAuthenticationInfo) GetThreeDSReqAuthTimestamp() string`

GetThreeDSReqAuthTimestamp returns the ThreeDSReqAuthTimestamp field if non-nil, zero value otherwise.

### GetThreeDSReqAuthTimestampOk

`func (o *ThreeDSRequestorAuthenticationInfo) GetThreeDSReqAuthTimestampOk() (*string, bool)`

GetThreeDSReqAuthTimestampOk returns a tuple with the ThreeDSReqAuthTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSReqAuthTimestamp

`func (o *ThreeDSRequestorAuthenticationInfo) SetThreeDSReqAuthTimestamp(v string)`

SetThreeDSReqAuthTimestamp sets ThreeDSReqAuthTimestamp field to given value.

### HasThreeDSReqAuthTimestamp

`func (o *ThreeDSRequestorAuthenticationInfo) HasThreeDSReqAuthTimestamp() bool`

HasThreeDSReqAuthTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


