# TransactionRuleEntityKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityReference** | Pointer to **string** | The unique identifier of the resource. | [optional] 
**EntityType** | Pointer to **string** | The type of resource.  Possible values: **balancePlatform**, **paymentInstrumentGroup**, **accountHolder**, **balanceAccount**, or **paymentInstrument**. | [optional] 

## Methods

### NewTransactionRuleEntityKey

`func NewTransactionRuleEntityKey() *TransactionRuleEntityKey`

NewTransactionRuleEntityKey instantiates a new TransactionRuleEntityKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRuleEntityKeyWithDefaults

`func NewTransactionRuleEntityKeyWithDefaults() *TransactionRuleEntityKey`

NewTransactionRuleEntityKeyWithDefaults instantiates a new TransactionRuleEntityKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityReference

`func (o *TransactionRuleEntityKey) GetEntityReference() string`

GetEntityReference returns the EntityReference field if non-nil, zero value otherwise.

### GetEntityReferenceOk

`func (o *TransactionRuleEntityKey) GetEntityReferenceOk() (*string, bool)`

GetEntityReferenceOk returns a tuple with the EntityReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityReference

`func (o *TransactionRuleEntityKey) SetEntityReference(v string)`

SetEntityReference sets EntityReference field to given value.

### HasEntityReference

`func (o *TransactionRuleEntityKey) HasEntityReference() bool`

HasEntityReference returns a boolean if a field has been set.

### GetEntityType

`func (o *TransactionRuleEntityKey) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *TransactionRuleEntityKey) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *TransactionRuleEntityKey) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *TransactionRuleEntityKey) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


