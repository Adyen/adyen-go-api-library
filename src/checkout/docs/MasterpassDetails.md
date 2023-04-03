# MasterpassDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckoutAttemptId** | Pointer to **string** | The checkout attempt identifier. | [optional] 
**FundingSource** | Pointer to **string** | The funding source that should be used when multiple sources are available. For Brazilian combo cards, by default the funding source is credit. To use debit, set this value to **debit**. | [optional] 
**MasterpassTransactionId** | **string** | The Masterpass transaction ID. | 
**Type** | Pointer to **string** | **masterpass** | [optional] [default to "masterpass"]

## Methods

### NewMasterpassDetails

`func NewMasterpassDetails(masterpassTransactionId string, ) *MasterpassDetails`

NewMasterpassDetails instantiates a new MasterpassDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMasterpassDetailsWithDefaults

`func NewMasterpassDetailsWithDefaults() *MasterpassDetails`

NewMasterpassDetailsWithDefaults instantiates a new MasterpassDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckoutAttemptId

`func (o *MasterpassDetails) GetCheckoutAttemptId() string`

GetCheckoutAttemptId returns the CheckoutAttemptId field if non-nil, zero value otherwise.

### GetCheckoutAttemptIdOk

`func (o *MasterpassDetails) GetCheckoutAttemptIdOk() (*string, bool)`

GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutAttemptId

`func (o *MasterpassDetails) SetCheckoutAttemptId(v string)`

SetCheckoutAttemptId sets CheckoutAttemptId field to given value.

### HasCheckoutAttemptId

`func (o *MasterpassDetails) HasCheckoutAttemptId() bool`

HasCheckoutAttemptId returns a boolean if a field has been set.

### GetFundingSource

`func (o *MasterpassDetails) GetFundingSource() string`

GetFundingSource returns the FundingSource field if non-nil, zero value otherwise.

### GetFundingSourceOk

`func (o *MasterpassDetails) GetFundingSourceOk() (*string, bool)`

GetFundingSourceOk returns a tuple with the FundingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSource

`func (o *MasterpassDetails) SetFundingSource(v string)`

SetFundingSource sets FundingSource field to given value.

### HasFundingSource

`func (o *MasterpassDetails) HasFundingSource() bool`

HasFundingSource returns a boolean if a field has been set.

### GetMasterpassTransactionId

`func (o *MasterpassDetails) GetMasterpassTransactionId() string`

GetMasterpassTransactionId returns the MasterpassTransactionId field if non-nil, zero value otherwise.

### GetMasterpassTransactionIdOk

`func (o *MasterpassDetails) GetMasterpassTransactionIdOk() (*string, bool)`

GetMasterpassTransactionIdOk returns a tuple with the MasterpassTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterpassTransactionId

`func (o *MasterpassDetails) SetMasterpassTransactionId(v string)`

SetMasterpassTransactionId sets MasterpassTransactionId field to given value.


### GetType

`func (o *MasterpassDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MasterpassDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MasterpassDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MasterpassDetails) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


