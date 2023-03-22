# VisaCheckoutDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckoutAttemptId** | Pointer to **string** | The checkout attempt identifier. | [optional] 
**FundingSource** | Pointer to **string** | The funding source that should be used when multiple sources are available. For Brazilian combo cards, by default the funding source is credit. To use debit, set this value to **debit**. | [optional] 
**Type** | Pointer to **string** | **visacheckout** | [optional] [default to "visacheckout"]
**VisaCheckoutCallId** | **string** | The Visa Click to Pay Call ID value. When your shopper selects a payment and/or a shipping address from Visa Click to Pay, you will receive a Visa Click to Pay Call ID. | 

## Methods

### NewVisaCheckoutDetails

`func NewVisaCheckoutDetails(visaCheckoutCallId string, ) *VisaCheckoutDetails`

NewVisaCheckoutDetails instantiates a new VisaCheckoutDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisaCheckoutDetailsWithDefaults

`func NewVisaCheckoutDetailsWithDefaults() *VisaCheckoutDetails`

NewVisaCheckoutDetailsWithDefaults instantiates a new VisaCheckoutDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckoutAttemptId

`func (o *VisaCheckoutDetails) GetCheckoutAttemptId() string`

GetCheckoutAttemptId returns the CheckoutAttemptId field if non-nil, zero value otherwise.

### GetCheckoutAttemptIdOk

`func (o *VisaCheckoutDetails) GetCheckoutAttemptIdOk() (*string, bool)`

GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutAttemptId

`func (o *VisaCheckoutDetails) SetCheckoutAttemptId(v string)`

SetCheckoutAttemptId sets CheckoutAttemptId field to given value.

### HasCheckoutAttemptId

`func (o *VisaCheckoutDetails) HasCheckoutAttemptId() bool`

HasCheckoutAttemptId returns a boolean if a field has been set.

### GetFundingSource

`func (o *VisaCheckoutDetails) GetFundingSource() string`

GetFundingSource returns the FundingSource field if non-nil, zero value otherwise.

### GetFundingSourceOk

`func (o *VisaCheckoutDetails) GetFundingSourceOk() (*string, bool)`

GetFundingSourceOk returns a tuple with the FundingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSource

`func (o *VisaCheckoutDetails) SetFundingSource(v string)`

SetFundingSource sets FundingSource field to given value.

### HasFundingSource

`func (o *VisaCheckoutDetails) HasFundingSource() bool`

HasFundingSource returns a boolean if a field has been set.

### GetType

`func (o *VisaCheckoutDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VisaCheckoutDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VisaCheckoutDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VisaCheckoutDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVisaCheckoutCallId

`func (o *VisaCheckoutDetails) GetVisaCheckoutCallId() string`

GetVisaCheckoutCallId returns the VisaCheckoutCallId field if non-nil, zero value otherwise.

### GetVisaCheckoutCallIdOk

`func (o *VisaCheckoutDetails) GetVisaCheckoutCallIdOk() (*string, bool)`

GetVisaCheckoutCallIdOk returns a tuple with the VisaCheckoutCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisaCheckoutCallId

`func (o *VisaCheckoutDetails) SetVisaCheckoutCallId(v string)`

SetVisaCheckoutCallId sets VisaCheckoutCallId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


