# DonationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**DonationAccount** | Pointer to **string** | The Adyen account name of your charity. We will provide you with this account name once your chosen charity has been [onboarded](https://docs.adyen.com/online-payments/donations#onboarding). | [optional] 
**Id** | Pointer to **string** | Your unique resource identifier. | [optional] 
**MerchantAccount** | Pointer to **string** | The merchant account identifier, with which you want to process the transaction. | [optional] 
**Payment** | Pointer to [**PaymentResponse**](PaymentResponse.md) |  | [optional] 
**Reference** | Pointer to **string** | The reference to uniquely identify a payment. This reference is used in all communication with you about the payment status. We recommend using a unique value per payment; however, it is not a requirement. If you need to provide multiple references for a transaction, separate them with hyphens (\&quot;-\&quot;). Maximum length: 80 characters. | [optional] 
**Status** | Pointer to **string** | The status of the donation transaction.  Possible values: * **completed** * **pending** * **refused** | [optional] 

## Methods

### NewDonationResponse

`func NewDonationResponse() *DonationResponse`

NewDonationResponse instantiates a new DonationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDonationResponseWithDefaults

`func NewDonationResponseWithDefaults() *DonationResponse`

NewDonationResponseWithDefaults instantiates a new DonationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DonationResponse) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DonationResponse) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DonationResponse) SetAmount(v Amount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DonationResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDonationAccount

`func (o *DonationResponse) GetDonationAccount() string`

GetDonationAccount returns the DonationAccount field if non-nil, zero value otherwise.

### GetDonationAccountOk

`func (o *DonationResponse) GetDonationAccountOk() (*string, bool)`

GetDonationAccountOk returns a tuple with the DonationAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDonationAccount

`func (o *DonationResponse) SetDonationAccount(v string)`

SetDonationAccount sets DonationAccount field to given value.

### HasDonationAccount

`func (o *DonationResponse) HasDonationAccount() bool`

HasDonationAccount returns a boolean if a field has been set.

### GetId

`func (o *DonationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DonationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DonationResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DonationResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *DonationResponse) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *DonationResponse) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *DonationResponse) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.

### HasMerchantAccount

`func (o *DonationResponse) HasMerchantAccount() bool`

HasMerchantAccount returns a boolean if a field has been set.

### GetPayment

`func (o *DonationResponse) GetPayment() PaymentResponse`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *DonationResponse) GetPaymentOk() (*PaymentResponse, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *DonationResponse) SetPayment(v PaymentResponse)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *DonationResponse) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetReference

`func (o *DonationResponse) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *DonationResponse) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *DonationResponse) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *DonationResponse) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetStatus

`func (o *DonationResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DonationResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DonationResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DonationResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


