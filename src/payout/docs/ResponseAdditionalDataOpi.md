# ResponseAdditionalDataOpi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpiTransToken** | Pointer to **string** | Returned in the response if you included &#x60;opi.includeTransToken: true&#x60; in an ecommerce payment request. This contains an Oracle Payment Interface token that you can store in your Oracle Opera database to identify tokenized ecommerce transactions. For more information and required settings, see [Oracle Opera](https://docs.adyen.com/plugins/oracle-opera#opi-token-ecommerce). | [optional] 

## Methods

### NewResponseAdditionalDataOpi

`func NewResponseAdditionalDataOpi() *ResponseAdditionalDataOpi`

NewResponseAdditionalDataOpi instantiates a new ResponseAdditionalDataOpi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAdditionalDataOpiWithDefaults

`func NewResponseAdditionalDataOpiWithDefaults() *ResponseAdditionalDataOpi`

NewResponseAdditionalDataOpiWithDefaults instantiates a new ResponseAdditionalDataOpi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpiTransToken

`func (o *ResponseAdditionalDataOpi) GetOpiTransToken() string`

GetOpiTransToken returns the OpiTransToken field if non-nil, zero value otherwise.

### GetOpiTransTokenOk

`func (o *ResponseAdditionalDataOpi) GetOpiTransTokenOk() (*string, bool)`

GetOpiTransTokenOk returns a tuple with the OpiTransToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpiTransToken

`func (o *ResponseAdditionalDataOpi) SetOpiTransToken(v string)`

SetOpiTransToken sets OpiTransToken field to given value.

### HasOpiTransToken

`func (o *ResponseAdditionalDataOpi) HasOpiTransToken() bool`

HasOpiTransToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


