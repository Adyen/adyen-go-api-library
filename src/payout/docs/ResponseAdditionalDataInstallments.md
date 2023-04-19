# ResponseAdditionalDataInstallments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallmentPaymentDataInstallmentType** | Pointer to **string** | Type of installment. The value of &#x60;installmentType&#x60; should be **IssuerFinanced**. | [optional] 
**InstallmentPaymentDataOptionItemNrAnnualPercentageRate** | Pointer to **string** | Annual interest rate. | [optional] 
**InstallmentPaymentDataOptionItemNrFirstInstallmentAmount** | Pointer to **string** | First Installment Amount in minor units. | [optional] 
**InstallmentPaymentDataOptionItemNrInstallmentFee** | Pointer to **string** | Installment fee amount in minor units. | [optional] 
**InstallmentPaymentDataOptionItemNrInterestRate** | Pointer to **string** | Interest rate for the installment period. | [optional] 
**InstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments** | Pointer to **string** | Maximum number of installments possible for this payment. | [optional] 
**InstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments** | Pointer to **string** | Minimum number of installments possible for this payment. | [optional] 
**InstallmentPaymentDataOptionItemNrNumberOfInstallments** | Pointer to **string** | Total number of installments possible for this payment. | [optional] 
**InstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount** | Pointer to **string** | Subsequent Installment Amount in minor units. | [optional] 
**InstallmentPaymentDataOptionItemNrTotalAmountDue** | Pointer to **string** | Total amount in minor units. | [optional] 
**InstallmentPaymentDataPaymentOptions** | Pointer to **string** | Possible values: * PayInInstallmentsOnly * PayInFullOnly * PayInFullOrInstallments | [optional] 
**InstallmentsValue** | Pointer to **string** | The number of installments that the payment amount should be charged with.  Example: 5 &gt; Only relevant for card payments in countries that support installments. | [optional] 

## Methods

### NewResponseAdditionalDataInstallments

`func NewResponseAdditionalDataInstallments() *ResponseAdditionalDataInstallments`

NewResponseAdditionalDataInstallments instantiates a new ResponseAdditionalDataInstallments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAdditionalDataInstallmentsWithDefaults

`func NewResponseAdditionalDataInstallmentsWithDefaults() *ResponseAdditionalDataInstallments`

NewResponseAdditionalDataInstallmentsWithDefaults instantiates a new ResponseAdditionalDataInstallments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallmentPaymentDataInstallmentType

`func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataInstallmentType() string`

GetInstallmentPaymentDataInstallmentType returns the InstallmentPaymentDataInstallmentType field if non-nil, zero value otherwise.

### GetInstallmentPaymentDataInstallmentTypeOk

`func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataInstallmentTypeOk() (*string, bool)`

GetInstallmentPaymentDataInstallmentTypeOk returns a tuple with the InstallmentPaymentDataInstallmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallmentPaymentDataInstallmentType

`func (o *ResponseAdditionalDataInstallments) SetInstallmentPaymentDataInstallmentType(v string)`

SetInstallmentPaymentDataInstallmentType sets InstallmentPaymentDataInstallmentType field to given value.

### HasInstallmentPaymentDataInstallmentType

`func (o *ResponseAdditionalDataInstallments) HasInstallmentPaymentDataInstallmentType() bool`

HasInstallmentPaymentDataInstallmentType returns a boolean if a field has been set.

### GetInstallmentPaymentDataOptionItemNrAnnualPercentageRate

`func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrAnnualPercentageRate() string`

GetInstallmentPaymentDataOptionItemNrAnnualPercentageRate returns the InstallmentPaymentDataOptionItemNrAnnualPercentageRate field if non-nil, zero value otherwise.

### GetInstallmentPaymentDataOptionItemNrAnnualPercentageRateOk

`func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrAnnualPercentageRateOk() (*string, bool)`

GetInstallmentPaymentDataOptionItemNrAnnualPercentageRateOk returns a tuple with the InstallmentPaymentDataOptionItemNrAnnualPercentageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallmentPaymentDataOptionItemNrAnnualPercentageRate

`func (o *ResponseAdditionalDataInstallments) SetInstallmentPaymentDataOptionItemNrAnnualPercentageRate(v string)`

SetInstallmentPaymentDataOptionItemNrAnnualPercentageRate sets InstallmentPaymentDataOptionItemNrAnnualPercentageRate field to given value.

### HasInstallmentPaymentDataOptionItemNrAnnualPercentageRate

`func (o *ResponseAdditionalDataInstallments) HasInstallmentPaymentDataOptionItemNrAnnualPercentageRate() bool`

HasInstallmentPaymentDataOptionItemNrAnnualPercentageRate returns a boolean if a field has been set.

### GetInstallmentPaymentDataOptionItemNrFirstInstallmentAmount

`func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrFirstInstallmentAmount() string`

GetInstallmentPaymentDataOptionItemNrFirstInstallmentAmount returns the InstallmentPaymentDataOptionItemNrFirstInstallmentAmount field if non-nil, zero value otherwise.

### GetInstallmentPaymentDataOptionItemNrFirstInstallmentAmountOk

`func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrFirstInstallmentAmountOk() (*string, bool)`

GetInstallmentPaymentDataOptionItemNrFirstInstallmentAmountOk returns a tuple with the InstallmentPaymentDataOptionItemNrFirstInstallmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallmentPaymentDataOptionItemNrFirstInstallmentAmount

`func (o *ResponseAdditionalDataInstallments) SetInstallmentPaymentDataOptionItemNrFirstInstallmentAmount(v string)`

SetInstallmentPaymentDataOptionItemNrFirstInstallmentAmount sets InstallmentPaymentDataOptionItemNrFirstInstallmentAmount field to given value.

### HasInstallmentPaymentDataOptionItemNrFirstInstallmentAmount

`func (o *ResponseAdditionalDataInstallments) HasInstallmentPaymentDataOptionItemNrFirstInstallmentAmount() bool`

HasInstallmentPaymentDataOptionItemNrFirstInstallmentAmount returns a boolean if a field has been set.

### GetInstallmentPaymentDataOptionItemNrInstallmentFee

`func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrInstallmentFee() string`

GetInstallmentPaymentDataOptionItemNrInstallmentFee returns the InstallmentPaymentDataOptionItemNrInstallmentFee field if non-nil, zero value otherwise.

### GetInstallmentPaymentDataOptionItemNrInstallmentFeeOk

`func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrInstallmentFeeOk() (*string, bool)`

GetInstallmentPaymentDataOptionItemNrInstallmentFeeOk returns a tuple with the InstallmentPaymentDataOptionItemNrInstallmentFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallmentPaymentDataOptionItemNrInstallmentFee

`func (o *ResponseAdditionalDataInstallments) SetInstallmentPaymentDataOptionItemNrInstallmentFee(v string)`

SetInstallmentPaymentDataOptionItemNrInstallmentFee sets InstallmentPaymentDataOptionItemNrInstallmentFee field to given value.

### HasInstallmentPaymentDataOptionItemNrInstallmentFee

`func (o *ResponseAdditionalDataInstallments) HasInstallmentPaymentDataOptionItemNrInstallmentFee() bool`

HasInstallmentPaymentDataOptionItemNrInstallmentFee returns a boolean if a field has been set.

### GetInstallmentPaymentDataOptionItemNrInterestRate

`func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrInterestRate() string`

GetInstallmentPaymentDataOptionItemNrInterestRate returns the InstallmentPaymentDataOptionItemNrInterestRate field if non-nil, zero value otherwise.

### GetInstallmentPaymentDataOptionItemNrInterestRateOk

`func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrInterestRateOk() (*string, bool)`

GetInstallmentPaymentDataOptionItemNrInterestRateOk returns a tuple with the InstallmentPaymentDataOptionItemNrInterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallmentPaymentDataOptionItemNrInterestRate

`func (o *ResponseAdditionalDataInstallments) SetInstallmentPaymentDataOptionItemNrInterestRate(v string)`

SetInstallmentPaymentDataOptionItemNrInterestRate sets InstallmentPaymentDataOptionItemNrInterestRate field to given value.

### HasInstallmentPaymentDataOptionItemNrInterestRate

`func (o *ResponseAdditionalDataInstallments) HasInstallmentPaymentDataOptionItemNrInterestRate() bool`

HasInstallmentPaymentDataOptionItemNrInterestRate returns a boolean if a field has been set.

### GetInstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments

`func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments() string`

GetInstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments returns the InstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments field if non-nil, zero value otherwise.

### GetInstallmentPaymentDataOptionItemNrMaximumNumberOfInstallmentsOk

`func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrMaximumNumberOfInstallmentsOk() (*string, bool)`

GetInstallmentPaymentDataOptionItemNrMaximumNumberOfInstallmentsOk returns a tuple with the InstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments

`func (o *ResponseAdditionalDataInstallments) SetInstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments(v string)`

SetInstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments sets InstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments field to given value.

### HasInstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments

`func (o *ResponseAdditionalDataInstallments) HasInstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments() bool`

HasInstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments returns a boolean if a field has been set.

### GetInstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments

`func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments() string`

GetInstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments returns the InstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments field if non-nil, zero value otherwise.

### GetInstallmentPaymentDataOptionItemNrMinimumNumberOfInstallmentsOk

`func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrMinimumNumberOfInstallmentsOk() (*string, bool)`

GetInstallmentPaymentDataOptionItemNrMinimumNumberOfInstallmentsOk returns a tuple with the InstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments

`func (o *ResponseAdditionalDataInstallments) SetInstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments(v string)`

SetInstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments sets InstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments field to given value.

### HasInstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments

`func (o *ResponseAdditionalDataInstallments) HasInstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments() bool`

HasInstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments returns a boolean if a field has been set.

### GetInstallmentPaymentDataOptionItemNrNumberOfInstallments

`func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrNumberOfInstallments() string`

GetInstallmentPaymentDataOptionItemNrNumberOfInstallments returns the InstallmentPaymentDataOptionItemNrNumberOfInstallments field if non-nil, zero value otherwise.

### GetInstallmentPaymentDataOptionItemNrNumberOfInstallmentsOk

`func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrNumberOfInstallmentsOk() (*string, bool)`

GetInstallmentPaymentDataOptionItemNrNumberOfInstallmentsOk returns a tuple with the InstallmentPaymentDataOptionItemNrNumberOfInstallments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallmentPaymentDataOptionItemNrNumberOfInstallments

`func (o *ResponseAdditionalDataInstallments) SetInstallmentPaymentDataOptionItemNrNumberOfInstallments(v string)`

SetInstallmentPaymentDataOptionItemNrNumberOfInstallments sets InstallmentPaymentDataOptionItemNrNumberOfInstallments field to given value.

### HasInstallmentPaymentDataOptionItemNrNumberOfInstallments

`func (o *ResponseAdditionalDataInstallments) HasInstallmentPaymentDataOptionItemNrNumberOfInstallments() bool`

HasInstallmentPaymentDataOptionItemNrNumberOfInstallments returns a boolean if a field has been set.

### GetInstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount

`func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount() string`

GetInstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount returns the InstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount field if non-nil, zero value otherwise.

### GetInstallmentPaymentDataOptionItemNrSubsequentInstallmentAmountOk

`func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrSubsequentInstallmentAmountOk() (*string, bool)`

GetInstallmentPaymentDataOptionItemNrSubsequentInstallmentAmountOk returns a tuple with the InstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount

`func (o *ResponseAdditionalDataInstallments) SetInstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount(v string)`

SetInstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount sets InstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount field to given value.

### HasInstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount

`func (o *ResponseAdditionalDataInstallments) HasInstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount() bool`

HasInstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount returns a boolean if a field has been set.

### GetInstallmentPaymentDataOptionItemNrTotalAmountDue

`func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrTotalAmountDue() string`

GetInstallmentPaymentDataOptionItemNrTotalAmountDue returns the InstallmentPaymentDataOptionItemNrTotalAmountDue field if non-nil, zero value otherwise.

### GetInstallmentPaymentDataOptionItemNrTotalAmountDueOk

`func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataOptionItemNrTotalAmountDueOk() (*string, bool)`

GetInstallmentPaymentDataOptionItemNrTotalAmountDueOk returns a tuple with the InstallmentPaymentDataOptionItemNrTotalAmountDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallmentPaymentDataOptionItemNrTotalAmountDue

`func (o *ResponseAdditionalDataInstallments) SetInstallmentPaymentDataOptionItemNrTotalAmountDue(v string)`

SetInstallmentPaymentDataOptionItemNrTotalAmountDue sets InstallmentPaymentDataOptionItemNrTotalAmountDue field to given value.

### HasInstallmentPaymentDataOptionItemNrTotalAmountDue

`func (o *ResponseAdditionalDataInstallments) HasInstallmentPaymentDataOptionItemNrTotalAmountDue() bool`

HasInstallmentPaymentDataOptionItemNrTotalAmountDue returns a boolean if a field has been set.

### GetInstallmentPaymentDataPaymentOptions

`func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataPaymentOptions() string`

GetInstallmentPaymentDataPaymentOptions returns the InstallmentPaymentDataPaymentOptions field if non-nil, zero value otherwise.

### GetInstallmentPaymentDataPaymentOptionsOk

`func (o *ResponseAdditionalDataInstallments) GetInstallmentPaymentDataPaymentOptionsOk() (*string, bool)`

GetInstallmentPaymentDataPaymentOptionsOk returns a tuple with the InstallmentPaymentDataPaymentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallmentPaymentDataPaymentOptions

`func (o *ResponseAdditionalDataInstallments) SetInstallmentPaymentDataPaymentOptions(v string)`

SetInstallmentPaymentDataPaymentOptions sets InstallmentPaymentDataPaymentOptions field to given value.

### HasInstallmentPaymentDataPaymentOptions

`func (o *ResponseAdditionalDataInstallments) HasInstallmentPaymentDataPaymentOptions() bool`

HasInstallmentPaymentDataPaymentOptions returns a boolean if a field has been set.

### GetInstallmentsValue

`func (o *ResponseAdditionalDataInstallments) GetInstallmentsValue() string`

GetInstallmentsValue returns the InstallmentsValue field if non-nil, zero value otherwise.

### GetInstallmentsValueOk

`func (o *ResponseAdditionalDataInstallments) GetInstallmentsValueOk() (*string, bool)`

GetInstallmentsValueOk returns a tuple with the InstallmentsValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallmentsValue

`func (o *ResponseAdditionalDataInstallments) SetInstallmentsValue(v string)`

SetInstallmentsValue sets InstallmentsValue field to given value.

### HasInstallmentsValue

`func (o *ResponseAdditionalDataInstallments) HasInstallmentsValue() bool`

HasInstallmentsValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


