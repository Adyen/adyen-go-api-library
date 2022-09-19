# Surcharge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AskConfirmation** | Pointer to **bool** | Show the surcharge details on the terminal, so the shopper can confirm. | [optional] 
**Configurations** | Pointer to [**[]Configuration**](Configuration.md) | Surcharge fees or percentages for specific payment methods, funding sources (credit or debit), and currencies. | [optional] 

## Methods

### NewSurcharge

`func NewSurcharge() *Surcharge`

NewSurcharge instantiates a new Surcharge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSurchargeWithDefaults

`func NewSurchargeWithDefaults() *Surcharge`

NewSurchargeWithDefaults instantiates a new Surcharge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAskConfirmation

`func (o *Surcharge) GetAskConfirmation() bool`

GetAskConfirmation returns the AskConfirmation field if non-nil, zero value otherwise.

### GetAskConfirmationOk

`func (o *Surcharge) GetAskConfirmationOk() (*bool, bool)`

GetAskConfirmationOk returns a tuple with the AskConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskConfirmation

`func (o *Surcharge) SetAskConfirmation(v bool)`

SetAskConfirmation sets AskConfirmation field to given value.

### HasAskConfirmation

`func (o *Surcharge) HasAskConfirmation() bool`

HasAskConfirmation returns a boolean if a field has been set.

### GetConfigurations

`func (o *Surcharge) GetConfigurations() []Configuration`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *Surcharge) GetConfigurationsOk() (*[]Configuration, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *Surcharge) SetConfigurations(v []Configuration)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *Surcharge) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


