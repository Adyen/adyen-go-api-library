# Recurring

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contract** | Pointer to **string** | The type of recurring contract to be used. Possible values: * &#x60;ONECLICK&#x60; – Payment details can be used to initiate a one-click payment, where the shopper enters the [card security code (CVC/CVV)](https://docs.adyen.com/payments-fundamentals/payment-glossary#card-security-code-cvc-cvv-cid). * &#x60;RECURRING&#x60; – Payment details can be used without the card security code to initiate [card-not-present transactions](https://docs.adyen.com/payments-fundamentals/payment-glossary#card-not-present-cnp). * &#x60;ONECLICK,RECURRING&#x60; – Payment details can be used regardless of whether the shopper is on your site or not. * &#x60;PAYOUT&#x60; – Payment details can be used to [make a payout](https://docs.adyen.com/online-payments/online-payouts). | [optional] 
**RecurringDetailName** | Pointer to **string** | A descriptive name for this detail. | [optional] 
**RecurringExpiry** | Pointer to **time.Time** | Date after which no further authorisations shall be performed. Only for 3D Secure 2. | [optional] 
**RecurringFrequency** | Pointer to **string** | Minimum number of days between authorisations. Only for 3D Secure 2. | [optional] 
**TokenService** | Pointer to **string** | The name of the token service. | [optional] 

## Methods

### NewRecurring

`func NewRecurring() *Recurring`

NewRecurring instantiates a new Recurring object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurringWithDefaults

`func NewRecurringWithDefaults() *Recurring`

NewRecurringWithDefaults instantiates a new Recurring object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContract

`func (o *Recurring) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *Recurring) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *Recurring) SetContract(v string)`

SetContract sets Contract field to given value.

### HasContract

`func (o *Recurring) HasContract() bool`

HasContract returns a boolean if a field has been set.

### GetRecurringDetailName

`func (o *Recurring) GetRecurringDetailName() string`

GetRecurringDetailName returns the RecurringDetailName field if non-nil, zero value otherwise.

### GetRecurringDetailNameOk

`func (o *Recurring) GetRecurringDetailNameOk() (*string, bool)`

GetRecurringDetailNameOk returns a tuple with the RecurringDetailName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDetailName

`func (o *Recurring) SetRecurringDetailName(v string)`

SetRecurringDetailName sets RecurringDetailName field to given value.

### HasRecurringDetailName

`func (o *Recurring) HasRecurringDetailName() bool`

HasRecurringDetailName returns a boolean if a field has been set.

### GetRecurringExpiry

`func (o *Recurring) GetRecurringExpiry() time.Time`

GetRecurringExpiry returns the RecurringExpiry field if non-nil, zero value otherwise.

### GetRecurringExpiryOk

`func (o *Recurring) GetRecurringExpiryOk() (*time.Time, bool)`

GetRecurringExpiryOk returns a tuple with the RecurringExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringExpiry

`func (o *Recurring) SetRecurringExpiry(v time.Time)`

SetRecurringExpiry sets RecurringExpiry field to given value.

### HasRecurringExpiry

`func (o *Recurring) HasRecurringExpiry() bool`

HasRecurringExpiry returns a boolean if a field has been set.

### GetRecurringFrequency

`func (o *Recurring) GetRecurringFrequency() string`

GetRecurringFrequency returns the RecurringFrequency field if non-nil, zero value otherwise.

### GetRecurringFrequencyOk

`func (o *Recurring) GetRecurringFrequencyOk() (*string, bool)`

GetRecurringFrequencyOk returns a tuple with the RecurringFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringFrequency

`func (o *Recurring) SetRecurringFrequency(v string)`

SetRecurringFrequency sets RecurringFrequency field to given value.

### HasRecurringFrequency

`func (o *Recurring) HasRecurringFrequency() bool`

HasRecurringFrequency returns a boolean if a field has been set.

### GetTokenService

`func (o *Recurring) GetTokenService() string`

GetTokenService returns the TokenService field if non-nil, zero value otherwise.

### GetTokenServiceOk

`func (o *Recurring) GetTokenServiceOk() (*string, bool)`

GetTokenServiceOk returns a tuple with the TokenService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenService

`func (o *Recurring) SetTokenService(v string)`

SetTokenService sets TokenService field to given value.

### HasTokenService

`func (o *Recurring) HasTokenService() bool`

HasTokenService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


