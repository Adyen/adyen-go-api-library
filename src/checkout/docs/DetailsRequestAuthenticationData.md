# DetailsRequestAuthenticationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationOnly** | Pointer to **bool** | If set to true, you will only perform the [3D Secure 2 authentication](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only), and not the payment authorisation. Default: *false**. | [optional] [default to false]

## Methods

### NewDetailsRequestAuthenticationData

`func NewDetailsRequestAuthenticationData() *DetailsRequestAuthenticationData`

NewDetailsRequestAuthenticationData instantiates a new DetailsRequestAuthenticationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailsRequestAuthenticationDataWithDefaults

`func NewDetailsRequestAuthenticationDataWithDefaults() *DetailsRequestAuthenticationData`

NewDetailsRequestAuthenticationDataWithDefaults instantiates a new DetailsRequestAuthenticationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationOnly

`func (o *DetailsRequestAuthenticationData) GetAuthenticationOnly() bool`

GetAuthenticationOnly returns the AuthenticationOnly field if non-nil, zero value otherwise.

### GetAuthenticationOnlyOk

`func (o *DetailsRequestAuthenticationData) GetAuthenticationOnlyOk() (*bool, bool)`

GetAuthenticationOnlyOk returns a tuple with the AuthenticationOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOnly

`func (o *DetailsRequestAuthenticationData) SetAuthenticationOnly(v bool)`

SetAuthenticationOnly sets AuthenticationOnly field to given value.

### HasAuthenticationOnly

`func (o *DetailsRequestAuthenticationData) HasAuthenticationOnly() bool`

HasAuthenticationOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


