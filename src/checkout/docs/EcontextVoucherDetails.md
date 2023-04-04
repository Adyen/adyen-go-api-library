# EcontextVoucherDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckoutAttemptId** | Pointer to **string** | The checkout attempt identifier. | [optional] 
**FirstName** | **string** | The shopper&#39;s first name. | 
**LastName** | **string** | The shopper&#39;s last name. | 
**ShopperEmail** | **string** | The shopper&#39;s email. | 
**TelephoneNumber** | **string** | The shopper&#39;s contact number. It must have an international number format, for example **+31 20 779 1846**. Formats like **+31 (0)20 779 1846** or **0031 20 779 1846** are not accepted. | 
**Type** | **string** | **econtextvoucher** | 

## Methods

### NewEcontextVoucherDetails

`func NewEcontextVoucherDetails(firstName string, lastName string, shopperEmail string, telephoneNumber string, type_ string, ) *EcontextVoucherDetails`

NewEcontextVoucherDetails instantiates a new EcontextVoucherDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEcontextVoucherDetailsWithDefaults

`func NewEcontextVoucherDetailsWithDefaults() *EcontextVoucherDetails`

NewEcontextVoucherDetailsWithDefaults instantiates a new EcontextVoucherDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckoutAttemptId

`func (o *EcontextVoucherDetails) GetCheckoutAttemptId() string`

GetCheckoutAttemptId returns the CheckoutAttemptId field if non-nil, zero value otherwise.

### GetCheckoutAttemptIdOk

`func (o *EcontextVoucherDetails) GetCheckoutAttemptIdOk() (*string, bool)`

GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutAttemptId

`func (o *EcontextVoucherDetails) SetCheckoutAttemptId(v string)`

SetCheckoutAttemptId sets CheckoutAttemptId field to given value.

### HasCheckoutAttemptId

`func (o *EcontextVoucherDetails) HasCheckoutAttemptId() bool`

HasCheckoutAttemptId returns a boolean if a field has been set.

### GetFirstName

`func (o *EcontextVoucherDetails) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *EcontextVoucherDetails) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *EcontextVoucherDetails) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *EcontextVoucherDetails) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *EcontextVoucherDetails) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *EcontextVoucherDetails) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetShopperEmail

`func (o *EcontextVoucherDetails) GetShopperEmail() string`

GetShopperEmail returns the ShopperEmail field if non-nil, zero value otherwise.

### GetShopperEmailOk

`func (o *EcontextVoucherDetails) GetShopperEmailOk() (*string, bool)`

GetShopperEmailOk returns a tuple with the ShopperEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperEmail

`func (o *EcontextVoucherDetails) SetShopperEmail(v string)`

SetShopperEmail sets ShopperEmail field to given value.


### GetTelephoneNumber

`func (o *EcontextVoucherDetails) GetTelephoneNumber() string`

GetTelephoneNumber returns the TelephoneNumber field if non-nil, zero value otherwise.

### GetTelephoneNumberOk

`func (o *EcontextVoucherDetails) GetTelephoneNumberOk() (*string, bool)`

GetTelephoneNumberOk returns a tuple with the TelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumber

`func (o *EcontextVoucherDetails) SetTelephoneNumber(v string)`

SetTelephoneNumber sets TelephoneNumber field to given value.


### GetType

`func (o *EcontextVoucherDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EcontextVoucherDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EcontextVoucherDetails) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


