# AdditionalDataLodging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LodgingCheckInDate** | Pointer to **string** | The arrival date. * Date format: **yyyyMmDd**. For example, for 2023 April 22, **20230422**. | [optional] 
**LodgingCheckOutDate** | Pointer to **string** | The departure date. * Date format: **yyyyMmDd**. For example, for 2023 April 22, **20230422**. | [optional] 
**LodgingCustomerServiceTollFreeNumber** | Pointer to **string** | The toll-free phone number for the lodging. * Format: alphanumeric. * Max length: 17 characters. * For US numbers: must start with 3 digits and be at least 10 characters in length. Otherwise, the capture can fail. | [optional] 
**LodgingFireSafetyActIndicator** | Pointer to **string** | Identifies that the facility complies with the Hotel and Motel Fire Safety Act of 1990. Values can be: &#39;Y&#39; or &#39;N&#39;. * Format: alphabetic. * Max length: 1 character. | [optional] 
**LodgingFolioCashAdvances** | Pointer to **string** | The folio cash advances. * Format: numeric. * Max length: 12 characters. | [optional] 
**LodgingFolioNumber** | Pointer to **string** | The card acceptor’s internal invoice or billing ID reference number. * Format: alphanumeric. * Max length: 25 characters. | [optional] 
**LodgingFoodBeverageCharges** | Pointer to **string** | The additional charges for food and beverages associated with the booking. * Format: numeric. * Max length: 12 characters. | [optional] 
**LodgingNoShowIndicator** | Pointer to **string** | Indicates if the customer didn&#39;t check in for their booking.  Possible values:   * **Y**: the customer didn&#39;t check in.   **N**: the customer checked in. | [optional] 
**LodgingPrepaidExpenses** | Pointer to **string** | The prepaid expenses for the booking. * Format: numeric. * Max length: 12 characters. | [optional] 
**LodgingPropertyPhoneNumber** | Pointer to **string** | Identifies the location of the lodging by its local phone number. * Format: alphanumeric. * Max length: 17 characters. * For US numbers: must start with 3 digits and be at least 10 characters in length. Otherwise, the capture can fail. | [optional] 
**LodgingRoom1NumberOfNights** | Pointer to **string** | The total number of nights the room is booked for. * Format: numeric. * Max length: 4 characters. | [optional] 
**LodgingRoom1Rate** | Pointer to **string** | The rate of the room. * Format: numeric. * Max length: 12 characters. * Must be in [minor units](https://docs.adyen.com/development-resources/currency-codes). | [optional] 
**LodgingRoom1Tax** | Pointer to **string** | The total amount of tax to be paid. * Format: numeric. * Max length: 12 chracters. * Must be in [minor units](https://docs.adyen.com/development-resources/currency-codes). | [optional] 
**LodgingTotalRoomTax** | Pointer to **string** | The total room tax amount. * Format: numeric. * Max length: 12 characters. * Must be in [minor units](https://docs.adyen.com/development-resources/currency-codes). | [optional] 
**LodgingTotalTax** | Pointer to **string** | The total tax amount. * Format: numeric. * Max length: 12 characters. * Must be in [minor units](https://docs.adyen.com/development-resources/currency-codes). | [optional] 
**TravelEntertainmentAuthDataDuration** | Pointer to **string** | The number of nights. This should be included in the auth message. * Format: numeric. * Max length: 2 characters. | [optional] 
**TravelEntertainmentAuthDataMarket** | Pointer to **string** | Indicates what market-specific dataset will be submitted or is being submitted. Value should be \&quot;H\&quot; for Hotel. This should be included in the auth message.  * Format: alphanumeric. * Max length: 1 character. | [optional] 

## Methods

### NewAdditionalDataLodging

`func NewAdditionalDataLodging() *AdditionalDataLodging`

NewAdditionalDataLodging instantiates a new AdditionalDataLodging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalDataLodgingWithDefaults

`func NewAdditionalDataLodgingWithDefaults() *AdditionalDataLodging`

NewAdditionalDataLodgingWithDefaults instantiates a new AdditionalDataLodging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLodgingCheckInDate

`func (o *AdditionalDataLodging) GetLodgingCheckInDate() string`

GetLodgingCheckInDate returns the LodgingCheckInDate field if non-nil, zero value otherwise.

### GetLodgingCheckInDateOk

`func (o *AdditionalDataLodging) GetLodgingCheckInDateOk() (*string, bool)`

GetLodgingCheckInDateOk returns a tuple with the LodgingCheckInDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLodgingCheckInDate

`func (o *AdditionalDataLodging) SetLodgingCheckInDate(v string)`

SetLodgingCheckInDate sets LodgingCheckInDate field to given value.

### HasLodgingCheckInDate

`func (o *AdditionalDataLodging) HasLodgingCheckInDate() bool`

HasLodgingCheckInDate returns a boolean if a field has been set.

### GetLodgingCheckOutDate

`func (o *AdditionalDataLodging) GetLodgingCheckOutDate() string`

GetLodgingCheckOutDate returns the LodgingCheckOutDate field if non-nil, zero value otherwise.

### GetLodgingCheckOutDateOk

`func (o *AdditionalDataLodging) GetLodgingCheckOutDateOk() (*string, bool)`

GetLodgingCheckOutDateOk returns a tuple with the LodgingCheckOutDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLodgingCheckOutDate

`func (o *AdditionalDataLodging) SetLodgingCheckOutDate(v string)`

SetLodgingCheckOutDate sets LodgingCheckOutDate field to given value.

### HasLodgingCheckOutDate

`func (o *AdditionalDataLodging) HasLodgingCheckOutDate() bool`

HasLodgingCheckOutDate returns a boolean if a field has been set.

### GetLodgingCustomerServiceTollFreeNumber

`func (o *AdditionalDataLodging) GetLodgingCustomerServiceTollFreeNumber() string`

GetLodgingCustomerServiceTollFreeNumber returns the LodgingCustomerServiceTollFreeNumber field if non-nil, zero value otherwise.

### GetLodgingCustomerServiceTollFreeNumberOk

`func (o *AdditionalDataLodging) GetLodgingCustomerServiceTollFreeNumberOk() (*string, bool)`

GetLodgingCustomerServiceTollFreeNumberOk returns a tuple with the LodgingCustomerServiceTollFreeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLodgingCustomerServiceTollFreeNumber

`func (o *AdditionalDataLodging) SetLodgingCustomerServiceTollFreeNumber(v string)`

SetLodgingCustomerServiceTollFreeNumber sets LodgingCustomerServiceTollFreeNumber field to given value.

### HasLodgingCustomerServiceTollFreeNumber

`func (o *AdditionalDataLodging) HasLodgingCustomerServiceTollFreeNumber() bool`

HasLodgingCustomerServiceTollFreeNumber returns a boolean if a field has been set.

### GetLodgingFireSafetyActIndicator

`func (o *AdditionalDataLodging) GetLodgingFireSafetyActIndicator() string`

GetLodgingFireSafetyActIndicator returns the LodgingFireSafetyActIndicator field if non-nil, zero value otherwise.

### GetLodgingFireSafetyActIndicatorOk

`func (o *AdditionalDataLodging) GetLodgingFireSafetyActIndicatorOk() (*string, bool)`

GetLodgingFireSafetyActIndicatorOk returns a tuple with the LodgingFireSafetyActIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLodgingFireSafetyActIndicator

`func (o *AdditionalDataLodging) SetLodgingFireSafetyActIndicator(v string)`

SetLodgingFireSafetyActIndicator sets LodgingFireSafetyActIndicator field to given value.

### HasLodgingFireSafetyActIndicator

`func (o *AdditionalDataLodging) HasLodgingFireSafetyActIndicator() bool`

HasLodgingFireSafetyActIndicator returns a boolean if a field has been set.

### GetLodgingFolioCashAdvances

`func (o *AdditionalDataLodging) GetLodgingFolioCashAdvances() string`

GetLodgingFolioCashAdvances returns the LodgingFolioCashAdvances field if non-nil, zero value otherwise.

### GetLodgingFolioCashAdvancesOk

`func (o *AdditionalDataLodging) GetLodgingFolioCashAdvancesOk() (*string, bool)`

GetLodgingFolioCashAdvancesOk returns a tuple with the LodgingFolioCashAdvances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLodgingFolioCashAdvances

`func (o *AdditionalDataLodging) SetLodgingFolioCashAdvances(v string)`

SetLodgingFolioCashAdvances sets LodgingFolioCashAdvances field to given value.

### HasLodgingFolioCashAdvances

`func (o *AdditionalDataLodging) HasLodgingFolioCashAdvances() bool`

HasLodgingFolioCashAdvances returns a boolean if a field has been set.

### GetLodgingFolioNumber

`func (o *AdditionalDataLodging) GetLodgingFolioNumber() string`

GetLodgingFolioNumber returns the LodgingFolioNumber field if non-nil, zero value otherwise.

### GetLodgingFolioNumberOk

`func (o *AdditionalDataLodging) GetLodgingFolioNumberOk() (*string, bool)`

GetLodgingFolioNumberOk returns a tuple with the LodgingFolioNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLodgingFolioNumber

`func (o *AdditionalDataLodging) SetLodgingFolioNumber(v string)`

SetLodgingFolioNumber sets LodgingFolioNumber field to given value.

### HasLodgingFolioNumber

`func (o *AdditionalDataLodging) HasLodgingFolioNumber() bool`

HasLodgingFolioNumber returns a boolean if a field has been set.

### GetLodgingFoodBeverageCharges

`func (o *AdditionalDataLodging) GetLodgingFoodBeverageCharges() string`

GetLodgingFoodBeverageCharges returns the LodgingFoodBeverageCharges field if non-nil, zero value otherwise.

### GetLodgingFoodBeverageChargesOk

`func (o *AdditionalDataLodging) GetLodgingFoodBeverageChargesOk() (*string, bool)`

GetLodgingFoodBeverageChargesOk returns a tuple with the LodgingFoodBeverageCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLodgingFoodBeverageCharges

`func (o *AdditionalDataLodging) SetLodgingFoodBeverageCharges(v string)`

SetLodgingFoodBeverageCharges sets LodgingFoodBeverageCharges field to given value.

### HasLodgingFoodBeverageCharges

`func (o *AdditionalDataLodging) HasLodgingFoodBeverageCharges() bool`

HasLodgingFoodBeverageCharges returns a boolean if a field has been set.

### GetLodgingNoShowIndicator

`func (o *AdditionalDataLodging) GetLodgingNoShowIndicator() string`

GetLodgingNoShowIndicator returns the LodgingNoShowIndicator field if non-nil, zero value otherwise.

### GetLodgingNoShowIndicatorOk

`func (o *AdditionalDataLodging) GetLodgingNoShowIndicatorOk() (*string, bool)`

GetLodgingNoShowIndicatorOk returns a tuple with the LodgingNoShowIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLodgingNoShowIndicator

`func (o *AdditionalDataLodging) SetLodgingNoShowIndicator(v string)`

SetLodgingNoShowIndicator sets LodgingNoShowIndicator field to given value.

### HasLodgingNoShowIndicator

`func (o *AdditionalDataLodging) HasLodgingNoShowIndicator() bool`

HasLodgingNoShowIndicator returns a boolean if a field has been set.

### GetLodgingPrepaidExpenses

`func (o *AdditionalDataLodging) GetLodgingPrepaidExpenses() string`

GetLodgingPrepaidExpenses returns the LodgingPrepaidExpenses field if non-nil, zero value otherwise.

### GetLodgingPrepaidExpensesOk

`func (o *AdditionalDataLodging) GetLodgingPrepaidExpensesOk() (*string, bool)`

GetLodgingPrepaidExpensesOk returns a tuple with the LodgingPrepaidExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLodgingPrepaidExpenses

`func (o *AdditionalDataLodging) SetLodgingPrepaidExpenses(v string)`

SetLodgingPrepaidExpenses sets LodgingPrepaidExpenses field to given value.

### HasLodgingPrepaidExpenses

`func (o *AdditionalDataLodging) HasLodgingPrepaidExpenses() bool`

HasLodgingPrepaidExpenses returns a boolean if a field has been set.

### GetLodgingPropertyPhoneNumber

`func (o *AdditionalDataLodging) GetLodgingPropertyPhoneNumber() string`

GetLodgingPropertyPhoneNumber returns the LodgingPropertyPhoneNumber field if non-nil, zero value otherwise.

### GetLodgingPropertyPhoneNumberOk

`func (o *AdditionalDataLodging) GetLodgingPropertyPhoneNumberOk() (*string, bool)`

GetLodgingPropertyPhoneNumberOk returns a tuple with the LodgingPropertyPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLodgingPropertyPhoneNumber

`func (o *AdditionalDataLodging) SetLodgingPropertyPhoneNumber(v string)`

SetLodgingPropertyPhoneNumber sets LodgingPropertyPhoneNumber field to given value.

### HasLodgingPropertyPhoneNumber

`func (o *AdditionalDataLodging) HasLodgingPropertyPhoneNumber() bool`

HasLodgingPropertyPhoneNumber returns a boolean if a field has been set.

### GetLodgingRoom1NumberOfNights

`func (o *AdditionalDataLodging) GetLodgingRoom1NumberOfNights() string`

GetLodgingRoom1NumberOfNights returns the LodgingRoom1NumberOfNights field if non-nil, zero value otherwise.

### GetLodgingRoom1NumberOfNightsOk

`func (o *AdditionalDataLodging) GetLodgingRoom1NumberOfNightsOk() (*string, bool)`

GetLodgingRoom1NumberOfNightsOk returns a tuple with the LodgingRoom1NumberOfNights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLodgingRoom1NumberOfNights

`func (o *AdditionalDataLodging) SetLodgingRoom1NumberOfNights(v string)`

SetLodgingRoom1NumberOfNights sets LodgingRoom1NumberOfNights field to given value.

### HasLodgingRoom1NumberOfNights

`func (o *AdditionalDataLodging) HasLodgingRoom1NumberOfNights() bool`

HasLodgingRoom1NumberOfNights returns a boolean if a field has been set.

### GetLodgingRoom1Rate

`func (o *AdditionalDataLodging) GetLodgingRoom1Rate() string`

GetLodgingRoom1Rate returns the LodgingRoom1Rate field if non-nil, zero value otherwise.

### GetLodgingRoom1RateOk

`func (o *AdditionalDataLodging) GetLodgingRoom1RateOk() (*string, bool)`

GetLodgingRoom1RateOk returns a tuple with the LodgingRoom1Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLodgingRoom1Rate

`func (o *AdditionalDataLodging) SetLodgingRoom1Rate(v string)`

SetLodgingRoom1Rate sets LodgingRoom1Rate field to given value.

### HasLodgingRoom1Rate

`func (o *AdditionalDataLodging) HasLodgingRoom1Rate() bool`

HasLodgingRoom1Rate returns a boolean if a field has been set.

### GetLodgingRoom1Tax

`func (o *AdditionalDataLodging) GetLodgingRoom1Tax() string`

GetLodgingRoom1Tax returns the LodgingRoom1Tax field if non-nil, zero value otherwise.

### GetLodgingRoom1TaxOk

`func (o *AdditionalDataLodging) GetLodgingRoom1TaxOk() (*string, bool)`

GetLodgingRoom1TaxOk returns a tuple with the LodgingRoom1Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLodgingRoom1Tax

`func (o *AdditionalDataLodging) SetLodgingRoom1Tax(v string)`

SetLodgingRoom1Tax sets LodgingRoom1Tax field to given value.

### HasLodgingRoom1Tax

`func (o *AdditionalDataLodging) HasLodgingRoom1Tax() bool`

HasLodgingRoom1Tax returns a boolean if a field has been set.

### GetLodgingTotalRoomTax

`func (o *AdditionalDataLodging) GetLodgingTotalRoomTax() string`

GetLodgingTotalRoomTax returns the LodgingTotalRoomTax field if non-nil, zero value otherwise.

### GetLodgingTotalRoomTaxOk

`func (o *AdditionalDataLodging) GetLodgingTotalRoomTaxOk() (*string, bool)`

GetLodgingTotalRoomTaxOk returns a tuple with the LodgingTotalRoomTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLodgingTotalRoomTax

`func (o *AdditionalDataLodging) SetLodgingTotalRoomTax(v string)`

SetLodgingTotalRoomTax sets LodgingTotalRoomTax field to given value.

### HasLodgingTotalRoomTax

`func (o *AdditionalDataLodging) HasLodgingTotalRoomTax() bool`

HasLodgingTotalRoomTax returns a boolean if a field has been set.

### GetLodgingTotalTax

`func (o *AdditionalDataLodging) GetLodgingTotalTax() string`

GetLodgingTotalTax returns the LodgingTotalTax field if non-nil, zero value otherwise.

### GetLodgingTotalTaxOk

`func (o *AdditionalDataLodging) GetLodgingTotalTaxOk() (*string, bool)`

GetLodgingTotalTaxOk returns a tuple with the LodgingTotalTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLodgingTotalTax

`func (o *AdditionalDataLodging) SetLodgingTotalTax(v string)`

SetLodgingTotalTax sets LodgingTotalTax field to given value.

### HasLodgingTotalTax

`func (o *AdditionalDataLodging) HasLodgingTotalTax() bool`

HasLodgingTotalTax returns a boolean if a field has been set.

### GetTravelEntertainmentAuthDataDuration

`func (o *AdditionalDataLodging) GetTravelEntertainmentAuthDataDuration() string`

GetTravelEntertainmentAuthDataDuration returns the TravelEntertainmentAuthDataDuration field if non-nil, zero value otherwise.

### GetTravelEntertainmentAuthDataDurationOk

`func (o *AdditionalDataLodging) GetTravelEntertainmentAuthDataDurationOk() (*string, bool)`

GetTravelEntertainmentAuthDataDurationOk returns a tuple with the TravelEntertainmentAuthDataDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelEntertainmentAuthDataDuration

`func (o *AdditionalDataLodging) SetTravelEntertainmentAuthDataDuration(v string)`

SetTravelEntertainmentAuthDataDuration sets TravelEntertainmentAuthDataDuration field to given value.

### HasTravelEntertainmentAuthDataDuration

`func (o *AdditionalDataLodging) HasTravelEntertainmentAuthDataDuration() bool`

HasTravelEntertainmentAuthDataDuration returns a boolean if a field has been set.

### GetTravelEntertainmentAuthDataMarket

`func (o *AdditionalDataLodging) GetTravelEntertainmentAuthDataMarket() string`

GetTravelEntertainmentAuthDataMarket returns the TravelEntertainmentAuthDataMarket field if non-nil, zero value otherwise.

### GetTravelEntertainmentAuthDataMarketOk

`func (o *AdditionalDataLodging) GetTravelEntertainmentAuthDataMarketOk() (*string, bool)`

GetTravelEntertainmentAuthDataMarketOk returns a tuple with the TravelEntertainmentAuthDataMarket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelEntertainmentAuthDataMarket

`func (o *AdditionalDataLodging) SetTravelEntertainmentAuthDataMarket(v string)`

SetTravelEntertainmentAuthDataMarket sets TravelEntertainmentAuthDataMarket field to given value.

### HasTravelEntertainmentAuthDataMarket

`func (o *AdditionalDataLodging) HasTravelEntertainmentAuthDataMarket() bool`

HasTravelEntertainmentAuthDataMarket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


