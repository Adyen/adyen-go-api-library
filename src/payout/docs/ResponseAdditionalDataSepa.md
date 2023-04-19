# ResponseAdditionalDataSepa

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SepadirectdebitDateOfSignature** | Pointer to **string** | The transaction signature date.  Format: yyyy-MM-dd | [optional] 
**SepadirectdebitMandateId** | Pointer to **string** | Its value corresponds to the pspReference value of the transaction. | [optional] 
**SepadirectdebitSequenceType** | Pointer to **string** | This field can take one of the following values: * OneOff: (OOFF) Direct debit instruction to initiate exactly one direct debit transaction.  * First: (FRST) Initial/first collection in a series of direct debit instructions. * Recurring: (RCUR) Direct debit instruction to carry out regular direct debit transactions initiated by the creditor. * Final: (FNAL) Last/final collection in a series of direct debit instructions.  Example: OOFF | [optional] 

## Methods

### NewResponseAdditionalDataSepa

`func NewResponseAdditionalDataSepa() *ResponseAdditionalDataSepa`

NewResponseAdditionalDataSepa instantiates a new ResponseAdditionalDataSepa object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAdditionalDataSepaWithDefaults

`func NewResponseAdditionalDataSepaWithDefaults() *ResponseAdditionalDataSepa`

NewResponseAdditionalDataSepaWithDefaults instantiates a new ResponseAdditionalDataSepa object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSepadirectdebitDateOfSignature

`func (o *ResponseAdditionalDataSepa) GetSepadirectdebitDateOfSignature() string`

GetSepadirectdebitDateOfSignature returns the SepadirectdebitDateOfSignature field if non-nil, zero value otherwise.

### GetSepadirectdebitDateOfSignatureOk

`func (o *ResponseAdditionalDataSepa) GetSepadirectdebitDateOfSignatureOk() (*string, bool)`

GetSepadirectdebitDateOfSignatureOk returns a tuple with the SepadirectdebitDateOfSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepadirectdebitDateOfSignature

`func (o *ResponseAdditionalDataSepa) SetSepadirectdebitDateOfSignature(v string)`

SetSepadirectdebitDateOfSignature sets SepadirectdebitDateOfSignature field to given value.

### HasSepadirectdebitDateOfSignature

`func (o *ResponseAdditionalDataSepa) HasSepadirectdebitDateOfSignature() bool`

HasSepadirectdebitDateOfSignature returns a boolean if a field has been set.

### GetSepadirectdebitMandateId

`func (o *ResponseAdditionalDataSepa) GetSepadirectdebitMandateId() string`

GetSepadirectdebitMandateId returns the SepadirectdebitMandateId field if non-nil, zero value otherwise.

### GetSepadirectdebitMandateIdOk

`func (o *ResponseAdditionalDataSepa) GetSepadirectdebitMandateIdOk() (*string, bool)`

GetSepadirectdebitMandateIdOk returns a tuple with the SepadirectdebitMandateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepadirectdebitMandateId

`func (o *ResponseAdditionalDataSepa) SetSepadirectdebitMandateId(v string)`

SetSepadirectdebitMandateId sets SepadirectdebitMandateId field to given value.

### HasSepadirectdebitMandateId

`func (o *ResponseAdditionalDataSepa) HasSepadirectdebitMandateId() bool`

HasSepadirectdebitMandateId returns a boolean if a field has been set.

### GetSepadirectdebitSequenceType

`func (o *ResponseAdditionalDataSepa) GetSepadirectdebitSequenceType() string`

GetSepadirectdebitSequenceType returns the SepadirectdebitSequenceType field if non-nil, zero value otherwise.

### GetSepadirectdebitSequenceTypeOk

`func (o *ResponseAdditionalDataSepa) GetSepadirectdebitSequenceTypeOk() (*string, bool)`

GetSepadirectdebitSequenceTypeOk returns a tuple with the SepadirectdebitSequenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepadirectdebitSequenceType

`func (o *ResponseAdditionalDataSepa) SetSepadirectdebitSequenceType(v string)`

SetSepadirectdebitSequenceType sets SepadirectdebitSequenceType field to given value.

### HasSepadirectdebitSequenceType

`func (o *ResponseAdditionalDataSepa) HasSepadirectdebitSequenceType() bool`

HasSepadirectdebitSequenceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


