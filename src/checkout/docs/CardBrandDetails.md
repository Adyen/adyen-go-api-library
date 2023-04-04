# CardBrandDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supported** | Pointer to **bool** | Indicates if you support the card brand. | [optional] 
**Type** | Pointer to **string** | The name of the card brand. | [optional] 

## Methods

### NewCardBrandDetails

`func NewCardBrandDetails() *CardBrandDetails`

NewCardBrandDetails instantiates a new CardBrandDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardBrandDetailsWithDefaults

`func NewCardBrandDetailsWithDefaults() *CardBrandDetails`

NewCardBrandDetailsWithDefaults instantiates a new CardBrandDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupported

`func (o *CardBrandDetails) GetSupported() bool`

GetSupported returns the Supported field if non-nil, zero value otherwise.

### GetSupportedOk

`func (o *CardBrandDetails) GetSupportedOk() (*bool, bool)`

GetSupportedOk returns a tuple with the Supported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupported

`func (o *CardBrandDetails) SetSupported(v bool)`

SetSupported sets Supported field to given value.

### HasSupported

`func (o *CardBrandDetails) HasSupported() bool`

HasSupported returns a boolean if a field has been set.

### GetType

`func (o *CardBrandDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CardBrandDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CardBrandDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CardBrandDetails) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


