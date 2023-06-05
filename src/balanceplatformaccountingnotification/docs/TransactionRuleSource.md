# TransactionRuleSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the resource, when applicable. | [optional] 
**Type** | Pointer to **string** | Indicates the type of resource for which the transaction rule is defined.  Possible values:   * **PaymentInstrumentGroup**  * **PaymentInstrument**  * **BalancePlatform**  * **EntityUsageConfiguration**  * **PlatformRule**: The transaction rule is a platform-wide rule imposed by Adyen. | [optional] 

## Methods

### NewTransactionRuleSource

`func NewTransactionRuleSource() *TransactionRuleSource`

NewTransactionRuleSource instantiates a new TransactionRuleSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRuleSourceWithDefaults

`func NewTransactionRuleSourceWithDefaults() *TransactionRuleSource`

NewTransactionRuleSourceWithDefaults instantiates a new TransactionRuleSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransactionRuleSource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionRuleSource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionRuleSource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransactionRuleSource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TransactionRuleSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionRuleSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionRuleSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransactionRuleSource) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


