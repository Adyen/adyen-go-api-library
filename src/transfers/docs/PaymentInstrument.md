# PaymentInstrument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the resource. | [optional] 
**Id** | Pointer to **string** | The unique identifier of the resource. | [optional] 
**Reference** | Pointer to **string** | The reference for the resource. | [optional] 
**TokenType** | Pointer to **string** | The type of wallet the network token is associated with. | [optional] 

## Methods

### NewPaymentInstrument

`func NewPaymentInstrument() *PaymentInstrument`

NewPaymentInstrument instantiates a new PaymentInstrument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInstrumentWithDefaults

`func NewPaymentInstrumentWithDefaults() *PaymentInstrument`

NewPaymentInstrumentWithDefaults instantiates a new PaymentInstrument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PaymentInstrument) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentInstrument) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentInstrument) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentInstrument) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *PaymentInstrument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentInstrument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentInstrument) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentInstrument) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReference

`func (o *PaymentInstrument) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaymentInstrument) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaymentInstrument) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PaymentInstrument) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetTokenType

`func (o *PaymentInstrument) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *PaymentInstrument) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *PaymentInstrument) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *PaymentInstrument) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


