# AdditionalDataOpi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpiIncludeTransToken** | Pointer to **string** | Optional boolean indicator. Set to **true** if you want an ecommerce transaction to return an &#x60;opi.transToken&#x60; as additional data in the response.  You can store this Oracle Payment Interface token in your Oracle Opera database. For more information and required settings, see [Oracle Opera](https://docs.adyen.com/plugins/oracle-opera#opi-token-ecommerce). | [optional] 

## Methods

### NewAdditionalDataOpi

`func NewAdditionalDataOpi() *AdditionalDataOpi`

NewAdditionalDataOpi instantiates a new AdditionalDataOpi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalDataOpiWithDefaults

`func NewAdditionalDataOpiWithDefaults() *AdditionalDataOpi`

NewAdditionalDataOpiWithDefaults instantiates a new AdditionalDataOpi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpiIncludeTransToken

`func (o *AdditionalDataOpi) GetOpiIncludeTransToken() string`

GetOpiIncludeTransToken returns the OpiIncludeTransToken field if non-nil, zero value otherwise.

### GetOpiIncludeTransTokenOk

`func (o *AdditionalDataOpi) GetOpiIncludeTransTokenOk() (*string, bool)`

GetOpiIncludeTransTokenOk returns a tuple with the OpiIncludeTransToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpiIncludeTransToken

`func (o *AdditionalDataOpi) SetOpiIncludeTransToken(v string)`

SetOpiIncludeTransToken sets OpiIncludeTransToken field to given value.

### HasOpiIncludeTransToken

`func (o *AdditionalDataOpi) HasOpiIncludeTransToken() bool`

HasOpiIncludeTransToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


