# TokenDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenData** | Pointer to **map[string]string** |  | [optional] 
**TokenDataType** | Pointer to **string** |  | [optional] 

## Methods

### NewTokenDetails

`func NewTokenDetails() *TokenDetails`

NewTokenDetails instantiates a new TokenDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenDetailsWithDefaults

`func NewTokenDetailsWithDefaults() *TokenDetails`

NewTokenDetailsWithDefaults instantiates a new TokenDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenData

`func (o *TokenDetails) GetTokenData() map[string]string`

GetTokenData returns the TokenData field if non-nil, zero value otherwise.

### GetTokenDataOk

`func (o *TokenDetails) GetTokenDataOk() (*map[string]string, bool)`

GetTokenDataOk returns a tuple with the TokenData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenData

`func (o *TokenDetails) SetTokenData(v map[string]string)`

SetTokenData sets TokenData field to given value.

### HasTokenData

`func (o *TokenDetails) HasTokenData() bool`

HasTokenData returns a boolean if a field has been set.

### GetTokenDataType

`func (o *TokenDetails) GetTokenDataType() string`

GetTokenDataType returns the TokenDataType field if non-nil, zero value otherwise.

### GetTokenDataTypeOk

`func (o *TokenDetails) GetTokenDataTypeOk() (*string, bool)`

GetTokenDataTypeOk returns a tuple with the TokenDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenDataType

`func (o *TokenDetails) SetTokenDataType(v string)`

SetTokenDataType sets TokenDataType field to given value.

### HasTokenDataType

`func (o *TokenDetails) HasTokenDataType() bool`

HasTokenDataType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


