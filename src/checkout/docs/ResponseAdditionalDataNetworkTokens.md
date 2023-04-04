# ResponseAdditionalDataNetworkTokens

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkTokenAvailable** | Pointer to **string** | Indicates whether a network token is available for the specified card. | [optional] 
**NetworkTokenBin** | Pointer to **string** | The Bank Identification Number of a tokenized card, which is the first six digits of a card number. | [optional] 
**NetworkTokenTokenSummary** | Pointer to **string** | The last four digits of a network token. | [optional] 

## Methods

### NewResponseAdditionalDataNetworkTokens

`func NewResponseAdditionalDataNetworkTokens() *ResponseAdditionalDataNetworkTokens`

NewResponseAdditionalDataNetworkTokens instantiates a new ResponseAdditionalDataNetworkTokens object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAdditionalDataNetworkTokensWithDefaults

`func NewResponseAdditionalDataNetworkTokensWithDefaults() *ResponseAdditionalDataNetworkTokens`

NewResponseAdditionalDataNetworkTokensWithDefaults instantiates a new ResponseAdditionalDataNetworkTokens object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkTokenAvailable

`func (o *ResponseAdditionalDataNetworkTokens) GetNetworkTokenAvailable() string`

GetNetworkTokenAvailable returns the NetworkTokenAvailable field if non-nil, zero value otherwise.

### GetNetworkTokenAvailableOk

`func (o *ResponseAdditionalDataNetworkTokens) GetNetworkTokenAvailableOk() (*string, bool)`

GetNetworkTokenAvailableOk returns a tuple with the NetworkTokenAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTokenAvailable

`func (o *ResponseAdditionalDataNetworkTokens) SetNetworkTokenAvailable(v string)`

SetNetworkTokenAvailable sets NetworkTokenAvailable field to given value.

### HasNetworkTokenAvailable

`func (o *ResponseAdditionalDataNetworkTokens) HasNetworkTokenAvailable() bool`

HasNetworkTokenAvailable returns a boolean if a field has been set.

### GetNetworkTokenBin

`func (o *ResponseAdditionalDataNetworkTokens) GetNetworkTokenBin() string`

GetNetworkTokenBin returns the NetworkTokenBin field if non-nil, zero value otherwise.

### GetNetworkTokenBinOk

`func (o *ResponseAdditionalDataNetworkTokens) GetNetworkTokenBinOk() (*string, bool)`

GetNetworkTokenBinOk returns a tuple with the NetworkTokenBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTokenBin

`func (o *ResponseAdditionalDataNetworkTokens) SetNetworkTokenBin(v string)`

SetNetworkTokenBin sets NetworkTokenBin field to given value.

### HasNetworkTokenBin

`func (o *ResponseAdditionalDataNetworkTokens) HasNetworkTokenBin() bool`

HasNetworkTokenBin returns a boolean if a field has been set.

### GetNetworkTokenTokenSummary

`func (o *ResponseAdditionalDataNetworkTokens) GetNetworkTokenTokenSummary() string`

GetNetworkTokenTokenSummary returns the NetworkTokenTokenSummary field if non-nil, zero value otherwise.

### GetNetworkTokenTokenSummaryOk

`func (o *ResponseAdditionalDataNetworkTokens) GetNetworkTokenTokenSummaryOk() (*string, bool)`

GetNetworkTokenTokenSummaryOk returns a tuple with the NetworkTokenTokenSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTokenTokenSummary

`func (o *ResponseAdditionalDataNetworkTokens) SetNetworkTokenTokenSummary(v string)`

SetNetworkTokenTokenSummary sets NetworkTokenTokenSummary field to given value.

### HasNetworkTokenTokenSummary

`func (o *ResponseAdditionalDataNetworkTokens) HasNetworkTokenTokenSummary() bool`

HasNetworkTokenTokenSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


