# AdditionalDataRatepay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RatepayInstallmentAmount** | Pointer to **string** | Amount the customer has to pay each month. | [optional] 
**RatepayInterestRate** | Pointer to **string** | Interest rate of this installment. | [optional] 
**RatepayLastInstallmentAmount** | Pointer to **string** | Amount of the last installment. | [optional] 
**RatepayPaymentFirstday** | Pointer to **string** | Calendar day of the first payment. | [optional] 
**RatepaydataDeliveryDate** | Pointer to **string** | Date the merchant delivered the goods to the customer. | [optional] 
**RatepaydataDueDate** | Pointer to **string** | Date by which the customer must settle the payment. | [optional] 
**RatepaydataInvoiceDate** | Pointer to **string** | Invoice date, defined by the merchant. If not included, the invoice date is set to the delivery date. | [optional] 
**RatepaydataInvoiceId** | Pointer to **string** | Identification name or number for the invoice, defined by the merchant. | [optional] 

## Methods

### NewAdditionalDataRatepay

`func NewAdditionalDataRatepay() *AdditionalDataRatepay`

NewAdditionalDataRatepay instantiates a new AdditionalDataRatepay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalDataRatepayWithDefaults

`func NewAdditionalDataRatepayWithDefaults() *AdditionalDataRatepay`

NewAdditionalDataRatepayWithDefaults instantiates a new AdditionalDataRatepay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRatepayInstallmentAmount

`func (o *AdditionalDataRatepay) GetRatepayInstallmentAmount() string`

GetRatepayInstallmentAmount returns the RatepayInstallmentAmount field if non-nil, zero value otherwise.

### GetRatepayInstallmentAmountOk

`func (o *AdditionalDataRatepay) GetRatepayInstallmentAmountOk() (*string, bool)`

GetRatepayInstallmentAmountOk returns a tuple with the RatepayInstallmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatepayInstallmentAmount

`func (o *AdditionalDataRatepay) SetRatepayInstallmentAmount(v string)`

SetRatepayInstallmentAmount sets RatepayInstallmentAmount field to given value.

### HasRatepayInstallmentAmount

`func (o *AdditionalDataRatepay) HasRatepayInstallmentAmount() bool`

HasRatepayInstallmentAmount returns a boolean if a field has been set.

### GetRatepayInterestRate

`func (o *AdditionalDataRatepay) GetRatepayInterestRate() string`

GetRatepayInterestRate returns the RatepayInterestRate field if non-nil, zero value otherwise.

### GetRatepayInterestRateOk

`func (o *AdditionalDataRatepay) GetRatepayInterestRateOk() (*string, bool)`

GetRatepayInterestRateOk returns a tuple with the RatepayInterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatepayInterestRate

`func (o *AdditionalDataRatepay) SetRatepayInterestRate(v string)`

SetRatepayInterestRate sets RatepayInterestRate field to given value.

### HasRatepayInterestRate

`func (o *AdditionalDataRatepay) HasRatepayInterestRate() bool`

HasRatepayInterestRate returns a boolean if a field has been set.

### GetRatepayLastInstallmentAmount

`func (o *AdditionalDataRatepay) GetRatepayLastInstallmentAmount() string`

GetRatepayLastInstallmentAmount returns the RatepayLastInstallmentAmount field if non-nil, zero value otherwise.

### GetRatepayLastInstallmentAmountOk

`func (o *AdditionalDataRatepay) GetRatepayLastInstallmentAmountOk() (*string, bool)`

GetRatepayLastInstallmentAmountOk returns a tuple with the RatepayLastInstallmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatepayLastInstallmentAmount

`func (o *AdditionalDataRatepay) SetRatepayLastInstallmentAmount(v string)`

SetRatepayLastInstallmentAmount sets RatepayLastInstallmentAmount field to given value.

### HasRatepayLastInstallmentAmount

`func (o *AdditionalDataRatepay) HasRatepayLastInstallmentAmount() bool`

HasRatepayLastInstallmentAmount returns a boolean if a field has been set.

### GetRatepayPaymentFirstday

`func (o *AdditionalDataRatepay) GetRatepayPaymentFirstday() string`

GetRatepayPaymentFirstday returns the RatepayPaymentFirstday field if non-nil, zero value otherwise.

### GetRatepayPaymentFirstdayOk

`func (o *AdditionalDataRatepay) GetRatepayPaymentFirstdayOk() (*string, bool)`

GetRatepayPaymentFirstdayOk returns a tuple with the RatepayPaymentFirstday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatepayPaymentFirstday

`func (o *AdditionalDataRatepay) SetRatepayPaymentFirstday(v string)`

SetRatepayPaymentFirstday sets RatepayPaymentFirstday field to given value.

### HasRatepayPaymentFirstday

`func (o *AdditionalDataRatepay) HasRatepayPaymentFirstday() bool`

HasRatepayPaymentFirstday returns a boolean if a field has been set.

### GetRatepaydataDeliveryDate

`func (o *AdditionalDataRatepay) GetRatepaydataDeliveryDate() string`

GetRatepaydataDeliveryDate returns the RatepaydataDeliveryDate field if non-nil, zero value otherwise.

### GetRatepaydataDeliveryDateOk

`func (o *AdditionalDataRatepay) GetRatepaydataDeliveryDateOk() (*string, bool)`

GetRatepaydataDeliveryDateOk returns a tuple with the RatepaydataDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatepaydataDeliveryDate

`func (o *AdditionalDataRatepay) SetRatepaydataDeliveryDate(v string)`

SetRatepaydataDeliveryDate sets RatepaydataDeliveryDate field to given value.

### HasRatepaydataDeliveryDate

`func (o *AdditionalDataRatepay) HasRatepaydataDeliveryDate() bool`

HasRatepaydataDeliveryDate returns a boolean if a field has been set.

### GetRatepaydataDueDate

`func (o *AdditionalDataRatepay) GetRatepaydataDueDate() string`

GetRatepaydataDueDate returns the RatepaydataDueDate field if non-nil, zero value otherwise.

### GetRatepaydataDueDateOk

`func (o *AdditionalDataRatepay) GetRatepaydataDueDateOk() (*string, bool)`

GetRatepaydataDueDateOk returns a tuple with the RatepaydataDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatepaydataDueDate

`func (o *AdditionalDataRatepay) SetRatepaydataDueDate(v string)`

SetRatepaydataDueDate sets RatepaydataDueDate field to given value.

### HasRatepaydataDueDate

`func (o *AdditionalDataRatepay) HasRatepaydataDueDate() bool`

HasRatepaydataDueDate returns a boolean if a field has been set.

### GetRatepaydataInvoiceDate

`func (o *AdditionalDataRatepay) GetRatepaydataInvoiceDate() string`

GetRatepaydataInvoiceDate returns the RatepaydataInvoiceDate field if non-nil, zero value otherwise.

### GetRatepaydataInvoiceDateOk

`func (o *AdditionalDataRatepay) GetRatepaydataInvoiceDateOk() (*string, bool)`

GetRatepaydataInvoiceDateOk returns a tuple with the RatepaydataInvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatepaydataInvoiceDate

`func (o *AdditionalDataRatepay) SetRatepaydataInvoiceDate(v string)`

SetRatepaydataInvoiceDate sets RatepaydataInvoiceDate field to given value.

### HasRatepaydataInvoiceDate

`func (o *AdditionalDataRatepay) HasRatepaydataInvoiceDate() bool`

HasRatepaydataInvoiceDate returns a boolean if a field has been set.

### GetRatepaydataInvoiceId

`func (o *AdditionalDataRatepay) GetRatepaydataInvoiceId() string`

GetRatepaydataInvoiceId returns the RatepaydataInvoiceId field if non-nil, zero value otherwise.

### GetRatepaydataInvoiceIdOk

`func (o *AdditionalDataRatepay) GetRatepaydataInvoiceIdOk() (*string, bool)`

GetRatepaydataInvoiceIdOk returns a tuple with the RatepaydataInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatepaydataInvoiceId

`func (o *AdditionalDataRatepay) SetRatepaydataInvoiceId(v string)`

SetRatepaydataInvoiceId sets RatepaydataInvoiceId field to given value.

### HasRatepaydataInvoiceId

`func (o *AdditionalDataRatepay) HasRatepaydataInvoiceId() bool`

HasRatepaydataInvoiceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


