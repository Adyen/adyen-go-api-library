# USLocalAccountIdentification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | The bank account number, without separators or whitespace. | 
**AccountType** | Pointer to **string** | The bank account type.  Possible values: **checking** or **savings**. Defaults to **checking**. | [optional] [default to "checking"]
**RoutingNumber** | **string** | The 9-digit [routing number](https://en.wikipedia.org/wiki/ABA_routing_transit_number), without separators or whitespace. | 
**Type** | **string** | **usLocal** | [default to "usLocal"]

## Methods

### NewUSLocalAccountIdentification

`func NewUSLocalAccountIdentification(accountNumber string, routingNumber string, type_ string, ) *USLocalAccountIdentification`

NewUSLocalAccountIdentification instantiates a new USLocalAccountIdentification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUSLocalAccountIdentificationWithDefaults

`func NewUSLocalAccountIdentificationWithDefaults() *USLocalAccountIdentification`

NewUSLocalAccountIdentificationWithDefaults instantiates a new USLocalAccountIdentification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *USLocalAccountIdentification) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *USLocalAccountIdentification) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *USLocalAccountIdentification) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetAccountType

`func (o *USLocalAccountIdentification) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *USLocalAccountIdentification) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *USLocalAccountIdentification) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *USLocalAccountIdentification) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetRoutingNumber

`func (o *USLocalAccountIdentification) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *USLocalAccountIdentification) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *USLocalAccountIdentification) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.


### GetType

`func (o *USLocalAccountIdentification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *USLocalAccountIdentification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *USLocalAccountIdentification) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


