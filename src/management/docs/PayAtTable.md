# PayAtTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationMethod** | Pointer to **string** | Allowed authentication methods: Magswipe, Manual Entry. | [optional] 
**EnablePayAtTable** | Pointer to **bool** | Enable Pay at table. | [optional] 

## Methods

### NewPayAtTable

`func NewPayAtTable() *PayAtTable`

NewPayAtTable instantiates a new PayAtTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayAtTableWithDefaults

`func NewPayAtTableWithDefaults() *PayAtTable`

NewPayAtTableWithDefaults instantiates a new PayAtTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationMethod

`func (o *PayAtTable) GetAuthenticationMethod() string`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *PayAtTable) GetAuthenticationMethodOk() (*string, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *PayAtTable) SetAuthenticationMethod(v string)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.

### HasAuthenticationMethod

`func (o *PayAtTable) HasAuthenticationMethod() bool`

HasAuthenticationMethod returns a boolean if a field has been set.

### GetEnablePayAtTable

`func (o *PayAtTable) GetEnablePayAtTable() bool`

GetEnablePayAtTable returns the EnablePayAtTable field if non-nil, zero value otherwise.

### GetEnablePayAtTableOk

`func (o *PayAtTable) GetEnablePayAtTableOk() (*bool, bool)`

GetEnablePayAtTableOk returns a tuple with the EnablePayAtTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePayAtTable

`func (o *PayAtTable) SetEnablePayAtTable(v bool)`

SetEnablePayAtTable sets EnablePayAtTable field to given value.

### HasEnablePayAtTable

`func (o *PayAtTable) HasEnablePayAtTable() bool`

HasEnablePayAtTable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


