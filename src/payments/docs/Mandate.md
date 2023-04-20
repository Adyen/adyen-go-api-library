# Mandate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | The billing amount (in minor units) of the recurring transactions. | 
**AmountRule** | Pointer to **string** | The limitation rule of the billing amount.  Possible values:  * **max**: The transaction amount can not exceed the &#x60;amount&#x60;.   * **exact**: The transaction amount should be the same as the &#x60;amount&#x60;.   | [optional] 
**BillingAttemptsRule** | Pointer to **string** | The rule to specify the period, within which the recurring debit can happen, relative to the mandate recurring date.  Possible values:   * **on**: On a specific date.   * **before**:  Before and on a specific date.   * **after**: On and after a specific date.   | [optional] 
**BillingDay** | Pointer to **string** | The number of the day, on which the recurring debit can happen. Should be within the same calendar month as the mandate recurring date.  Possible values: 1-31 based on the &#x60;frequency&#x60;. | [optional] 
**EndsAt** | **string** | End date of the billing plan, in YYYY-MM-DD format. | 
**Frequency** | **string** | The frequency with which a shopper should be charged.  Possible values: **daily**, **weekly**, **biWeekly**, **monthly**, **quarterly**, **halfYearly**, **yearly**. | 
**Remarks** | Pointer to **string** | The message shown by UPI to the shopper on the approval screen. | [optional] 
**StartsAt** | Pointer to **string** | Start date of the billing plan, in YYYY-MM-DD format. By default, the transaction date. | [optional] 

## Methods

### NewMandate

`func NewMandate(amount string, endsAt string, frequency string, ) *Mandate`

NewMandate instantiates a new Mandate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMandateWithDefaults

`func NewMandateWithDefaults() *Mandate`

NewMandateWithDefaults instantiates a new Mandate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Mandate) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Mandate) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Mandate) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetAmountRule

`func (o *Mandate) GetAmountRule() string`

GetAmountRule returns the AmountRule field if non-nil, zero value otherwise.

### GetAmountRuleOk

`func (o *Mandate) GetAmountRuleOk() (*string, bool)`

GetAmountRuleOk returns a tuple with the AmountRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRule

`func (o *Mandate) SetAmountRule(v string)`

SetAmountRule sets AmountRule field to given value.

### HasAmountRule

`func (o *Mandate) HasAmountRule() bool`

HasAmountRule returns a boolean if a field has been set.

### GetBillingAttemptsRule

`func (o *Mandate) GetBillingAttemptsRule() string`

GetBillingAttemptsRule returns the BillingAttemptsRule field if non-nil, zero value otherwise.

### GetBillingAttemptsRuleOk

`func (o *Mandate) GetBillingAttemptsRuleOk() (*string, bool)`

GetBillingAttemptsRuleOk returns a tuple with the BillingAttemptsRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAttemptsRule

`func (o *Mandate) SetBillingAttemptsRule(v string)`

SetBillingAttemptsRule sets BillingAttemptsRule field to given value.

### HasBillingAttemptsRule

`func (o *Mandate) HasBillingAttemptsRule() bool`

HasBillingAttemptsRule returns a boolean if a field has been set.

### GetBillingDay

`func (o *Mandate) GetBillingDay() string`

GetBillingDay returns the BillingDay field if non-nil, zero value otherwise.

### GetBillingDayOk

`func (o *Mandate) GetBillingDayOk() (*string, bool)`

GetBillingDayOk returns a tuple with the BillingDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingDay

`func (o *Mandate) SetBillingDay(v string)`

SetBillingDay sets BillingDay field to given value.

### HasBillingDay

`func (o *Mandate) HasBillingDay() bool`

HasBillingDay returns a boolean if a field has been set.

### GetEndsAt

`func (o *Mandate) GetEndsAt() string`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *Mandate) GetEndsAtOk() (*string, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *Mandate) SetEndsAt(v string)`

SetEndsAt sets EndsAt field to given value.


### GetFrequency

`func (o *Mandate) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *Mandate) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *Mandate) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.


### GetRemarks

`func (o *Mandate) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *Mandate) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *Mandate) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *Mandate) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetStartsAt

`func (o *Mandate) GetStartsAt() string`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *Mandate) GetStartsAtOk() (*string, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *Mandate) SetStartsAt(v string)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *Mandate) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


