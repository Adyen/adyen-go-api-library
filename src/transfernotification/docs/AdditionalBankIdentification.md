# AdditionalBankIdentification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The value of the additional bank identification. | [optional] 
**Type** | Pointer to **string** | The type of additional bank identification, depending on the country.  Possible values:   * **gbSortCode**: The 6-digit [UK sort code](https://en.wikipedia.org/wiki/Sort_code), without separators or spaces  * **usRoutingNumber**: The 9-digit [routing number](https://en.wikipedia.org/wiki/ABA_routing_transit_number), without separators or spaces. | [optional] 

## Methods

### NewAdditionalBankIdentification

`func NewAdditionalBankIdentification() *AdditionalBankIdentification`

NewAdditionalBankIdentification instantiates a new AdditionalBankIdentification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalBankIdentificationWithDefaults

`func NewAdditionalBankIdentificationWithDefaults() *AdditionalBankIdentification`

NewAdditionalBankIdentificationWithDefaults instantiates a new AdditionalBankIdentification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AdditionalBankIdentification) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AdditionalBankIdentification) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AdditionalBankIdentification) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AdditionalBankIdentification) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *AdditionalBankIdentification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdditionalBankIdentification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdditionalBankIdentification) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AdditionalBankIdentification) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


