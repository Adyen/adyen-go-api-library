# Connectivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SimcardStatus** | Pointer to **string** | Indicates the status of the SIM card in the payment terminal. Can be updated and received only at terminal level, and only for models that support cellular connectivity.  Possible values: * **ACTIVATED**: the SIM card is activated. Cellular connectivity may still need to be enabled on the terminal itself, in the **Network** settings. * **INVENTORY**: the SIM card is not activated. The terminal can&#39;t use cellular connectivity. | [optional] 

## Methods

### NewConnectivity

`func NewConnectivity() *Connectivity`

NewConnectivity instantiates a new Connectivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectivityWithDefaults

`func NewConnectivityWithDefaults() *Connectivity`

NewConnectivityWithDefaults instantiates a new Connectivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSimcardStatus

`func (o *Connectivity) GetSimcardStatus() string`

GetSimcardStatus returns the SimcardStatus field if non-nil, zero value otherwise.

### GetSimcardStatusOk

`func (o *Connectivity) GetSimcardStatusOk() (*string, bool)`

GetSimcardStatusOk returns a tuple with the SimcardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimcardStatus

`func (o *Connectivity) SetSimcardStatus(v string)`

SetSimcardStatus sets SimcardStatus field to given value.

### HasSimcardStatus

`func (o *Connectivity) HasSimcardStatus() bool`

HasSimcardStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


