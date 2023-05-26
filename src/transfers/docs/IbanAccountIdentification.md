# IbanAccountIdentification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iban** | **string** | The international bank account number as defined in the [ISO-13616](https://www.iso.org/standard/81090.html) standard. | 
**Type** | **string** | **iban** | [default to "iban"]

## Methods

### NewIbanAccountIdentification

`func NewIbanAccountIdentification(iban string, type_ string, ) *IbanAccountIdentification`

NewIbanAccountIdentification instantiates a new IbanAccountIdentification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIbanAccountIdentificationWithDefaults

`func NewIbanAccountIdentificationWithDefaults() *IbanAccountIdentification`

NewIbanAccountIdentificationWithDefaults instantiates a new IbanAccountIdentification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIban

`func (o *IbanAccountIdentification) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *IbanAccountIdentification) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *IbanAccountIdentification) SetIban(v string)`

SetIban sets Iban field to given value.


### GetType

`func (o *IbanAccountIdentification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IbanAccountIdentification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IbanAccountIdentification) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


