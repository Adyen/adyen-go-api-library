# AdditionalDataCarRental

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarRentalCheckOutDate** | Pointer to **string** | The pick-up date. * Date format: &#x60;yyyyMMdd&#x60; | [optional] 
**CarRentalCustomerServiceTollFreeNumber** | Pointer to **string** | The customer service phone number of the car rental company. * Format: Alphanumeric * maxLength: 17 * For US and CA numbers must be 10 characters in length * Must not start with a space * Must not be all zeros * Must not contain any special characters such as + or - | [optional] 
**CarRentalDaysRented** | Pointer to **string** | Number of days for which the car is being rented. * Format: Numeric * maxLength: 2 * Must not be all spaces | [optional] 
**CarRentalFuelCharges** | Pointer to **string** | Any fuel charges associated with the rental, in [minor units](https://docs.adyen.com/development-resources/currency-codes). * Format: Numeric * maxLength: 12 | [optional] 
**CarRentalInsuranceCharges** | Pointer to **string** | Any insurance charges associated with the rental, in [minor units](https://docs.adyen.com/development-resources/currency-codes). * Format: Numeric * maxLength: 12 * Must not be all spaces * Must not be all zeros | [optional] 
**CarRentalLocationCity** | Pointer to **string** | The city where the car is rented. * Format: Alphanumeric * maxLength: 18 * Must not start with a space or be all spaces * Must not be all zeros | [optional] 
**CarRentalLocationCountry** | Pointer to **string** | The country where the car is rented, in [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) format. * Format: Alphanumeric * maxLength: 2 | [optional] 
**CarRentalLocationStateProvince** | Pointer to **string** | The state or province where the car is rented. * Format: Alphanumeric * maxLength: 2 * Must not start with a space or be all spaces * Must not be all zeros | [optional] 
**CarRentalNoShowIndicator** | Pointer to **string** | Indicates if the customer didn&#39;t pick up their rental car. * Y - Customer did not pick up their car * N - Not applicable | [optional] 
**CarRentalOneWayDropOffCharges** | Pointer to **string** | The charge for not returning a car to the original rental location, in [minor units](https://docs.adyen.com/development-resources/currency-codes). * maxLength: 12 | [optional] 
**CarRentalRate** | Pointer to **string** | The daily rental rate, in [minor units](https://docs.adyen.com/development-resources/currency-codes). * Format: Alphanumeric * maxLength: 12 | [optional] 
**CarRentalRateIndicator** | Pointer to **string** | Specifies whether the given rate is applied daily or weekly. * D - Daily rate * W - Weekly rate | [optional] 
**CarRentalRentalAgreementNumber** | Pointer to **string** | The rental agreement number for the car rental. * Format: Alphanumeric * maxLength: 9 * Must not start with a space or be all spaces * Must not be all zeros | [optional] 
**CarRentalRentalClassId** | Pointer to **string** | The classification of the rental car. * Format: Alphanumeric * maxLength: 4 * Must not start with a space or be all spaces * Must not be all zeros | [optional] 
**CarRentalRenterName** | Pointer to **string** | The name of the person renting the car. * Format: Alphanumeric * maxLength: 26 * If you send more than 26 characters, the name is truncated * Must not start with a space or be all spaces * Must not be all zeros | [optional] 
**CarRentalReturnCity** | Pointer to **string** | The city where the car must be returned. * Format: Alphanumeric * maxLength: 18 * Must not start with a space or be all spaces * Must not be all zeros | [optional] 
**CarRentalReturnCountry** | Pointer to **string** | The country where the car must be returned, in [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) format. * Format: Alphanumeric * maxLength: 2 | [optional] 
**CarRentalReturnDate** | Pointer to **string** | The last date to return the car by. * Date format: &#x60;yyyyMMdd&#x60; * maxLength: 8 | [optional] 
**CarRentalReturnLocationId** | Pointer to **string** | The agency code, phone number, or address abbreviation * Format: Alphanumeric * maxLength: 10 * Must not start with a space or be all spaces * Must not be all zeros | [optional] 
**CarRentalReturnStateProvince** | Pointer to **string** | The state or province where the car must be returned. * Format: Alphanumeric * maxLength: 3 * Must not start with a space or be all spaces * Must not be all zeros | [optional] 
**CarRentalTaxExemptIndicator** | Pointer to **string** | Indicates if the goods or services were tax-exempt, or if tax was not paid on them.  Values: * Y - Goods or services were tax exempt * N - Tax was not collected | [optional] 
**TravelEntertainmentAuthDataDuration** | Pointer to **string** | Number of days the car is rented for. This should be included in the auth message. * Format: Numeric * maxLength: 2 | [optional] 
**TravelEntertainmentAuthDataMarket** | Pointer to **string** | Indicates what market-specific dataset will be submitted or is being submitted. Value should be &#39;A&#39; for car rental. This should be included in the auth message. * Format: Alphanumeric * maxLength: 1 | [optional] 

## Methods

### NewAdditionalDataCarRental

`func NewAdditionalDataCarRental() *AdditionalDataCarRental`

NewAdditionalDataCarRental instantiates a new AdditionalDataCarRental object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalDataCarRentalWithDefaults

`func NewAdditionalDataCarRentalWithDefaults() *AdditionalDataCarRental`

NewAdditionalDataCarRentalWithDefaults instantiates a new AdditionalDataCarRental object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarRentalCheckOutDate

`func (o *AdditionalDataCarRental) GetCarRentalCheckOutDate() string`

GetCarRentalCheckOutDate returns the CarRentalCheckOutDate field if non-nil, zero value otherwise.

### GetCarRentalCheckOutDateOk

`func (o *AdditionalDataCarRental) GetCarRentalCheckOutDateOk() (*string, bool)`

GetCarRentalCheckOutDateOk returns a tuple with the CarRentalCheckOutDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarRentalCheckOutDate

`func (o *AdditionalDataCarRental) SetCarRentalCheckOutDate(v string)`

SetCarRentalCheckOutDate sets CarRentalCheckOutDate field to given value.

### HasCarRentalCheckOutDate

`func (o *AdditionalDataCarRental) HasCarRentalCheckOutDate() bool`

HasCarRentalCheckOutDate returns a boolean if a field has been set.

### GetCarRentalCustomerServiceTollFreeNumber

`func (o *AdditionalDataCarRental) GetCarRentalCustomerServiceTollFreeNumber() string`

GetCarRentalCustomerServiceTollFreeNumber returns the CarRentalCustomerServiceTollFreeNumber field if non-nil, zero value otherwise.

### GetCarRentalCustomerServiceTollFreeNumberOk

`func (o *AdditionalDataCarRental) GetCarRentalCustomerServiceTollFreeNumberOk() (*string, bool)`

GetCarRentalCustomerServiceTollFreeNumberOk returns a tuple with the CarRentalCustomerServiceTollFreeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarRentalCustomerServiceTollFreeNumber

`func (o *AdditionalDataCarRental) SetCarRentalCustomerServiceTollFreeNumber(v string)`

SetCarRentalCustomerServiceTollFreeNumber sets CarRentalCustomerServiceTollFreeNumber field to given value.

### HasCarRentalCustomerServiceTollFreeNumber

`func (o *AdditionalDataCarRental) HasCarRentalCustomerServiceTollFreeNumber() bool`

HasCarRentalCustomerServiceTollFreeNumber returns a boolean if a field has been set.

### GetCarRentalDaysRented

`func (o *AdditionalDataCarRental) GetCarRentalDaysRented() string`

GetCarRentalDaysRented returns the CarRentalDaysRented field if non-nil, zero value otherwise.

### GetCarRentalDaysRentedOk

`func (o *AdditionalDataCarRental) GetCarRentalDaysRentedOk() (*string, bool)`

GetCarRentalDaysRentedOk returns a tuple with the CarRentalDaysRented field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarRentalDaysRented

`func (o *AdditionalDataCarRental) SetCarRentalDaysRented(v string)`

SetCarRentalDaysRented sets CarRentalDaysRented field to given value.

### HasCarRentalDaysRented

`func (o *AdditionalDataCarRental) HasCarRentalDaysRented() bool`

HasCarRentalDaysRented returns a boolean if a field has been set.

### GetCarRentalFuelCharges

`func (o *AdditionalDataCarRental) GetCarRentalFuelCharges() string`

GetCarRentalFuelCharges returns the CarRentalFuelCharges field if non-nil, zero value otherwise.

### GetCarRentalFuelChargesOk

`func (o *AdditionalDataCarRental) GetCarRentalFuelChargesOk() (*string, bool)`

GetCarRentalFuelChargesOk returns a tuple with the CarRentalFuelCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarRentalFuelCharges

`func (o *AdditionalDataCarRental) SetCarRentalFuelCharges(v string)`

SetCarRentalFuelCharges sets CarRentalFuelCharges field to given value.

### HasCarRentalFuelCharges

`func (o *AdditionalDataCarRental) HasCarRentalFuelCharges() bool`

HasCarRentalFuelCharges returns a boolean if a field has been set.

### GetCarRentalInsuranceCharges

`func (o *AdditionalDataCarRental) GetCarRentalInsuranceCharges() string`

GetCarRentalInsuranceCharges returns the CarRentalInsuranceCharges field if non-nil, zero value otherwise.

### GetCarRentalInsuranceChargesOk

`func (o *AdditionalDataCarRental) GetCarRentalInsuranceChargesOk() (*string, bool)`

GetCarRentalInsuranceChargesOk returns a tuple with the CarRentalInsuranceCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarRentalInsuranceCharges

`func (o *AdditionalDataCarRental) SetCarRentalInsuranceCharges(v string)`

SetCarRentalInsuranceCharges sets CarRentalInsuranceCharges field to given value.

### HasCarRentalInsuranceCharges

`func (o *AdditionalDataCarRental) HasCarRentalInsuranceCharges() bool`

HasCarRentalInsuranceCharges returns a boolean if a field has been set.

### GetCarRentalLocationCity

`func (o *AdditionalDataCarRental) GetCarRentalLocationCity() string`

GetCarRentalLocationCity returns the CarRentalLocationCity field if non-nil, zero value otherwise.

### GetCarRentalLocationCityOk

`func (o *AdditionalDataCarRental) GetCarRentalLocationCityOk() (*string, bool)`

GetCarRentalLocationCityOk returns a tuple with the CarRentalLocationCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarRentalLocationCity

`func (o *AdditionalDataCarRental) SetCarRentalLocationCity(v string)`

SetCarRentalLocationCity sets CarRentalLocationCity field to given value.

### HasCarRentalLocationCity

`func (o *AdditionalDataCarRental) HasCarRentalLocationCity() bool`

HasCarRentalLocationCity returns a boolean if a field has been set.

### GetCarRentalLocationCountry

`func (o *AdditionalDataCarRental) GetCarRentalLocationCountry() string`

GetCarRentalLocationCountry returns the CarRentalLocationCountry field if non-nil, zero value otherwise.

### GetCarRentalLocationCountryOk

`func (o *AdditionalDataCarRental) GetCarRentalLocationCountryOk() (*string, bool)`

GetCarRentalLocationCountryOk returns a tuple with the CarRentalLocationCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarRentalLocationCountry

`func (o *AdditionalDataCarRental) SetCarRentalLocationCountry(v string)`

SetCarRentalLocationCountry sets CarRentalLocationCountry field to given value.

### HasCarRentalLocationCountry

`func (o *AdditionalDataCarRental) HasCarRentalLocationCountry() bool`

HasCarRentalLocationCountry returns a boolean if a field has been set.

### GetCarRentalLocationStateProvince

`func (o *AdditionalDataCarRental) GetCarRentalLocationStateProvince() string`

GetCarRentalLocationStateProvince returns the CarRentalLocationStateProvince field if non-nil, zero value otherwise.

### GetCarRentalLocationStateProvinceOk

`func (o *AdditionalDataCarRental) GetCarRentalLocationStateProvinceOk() (*string, bool)`

GetCarRentalLocationStateProvinceOk returns a tuple with the CarRentalLocationStateProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarRentalLocationStateProvince

`func (o *AdditionalDataCarRental) SetCarRentalLocationStateProvince(v string)`

SetCarRentalLocationStateProvince sets CarRentalLocationStateProvince field to given value.

### HasCarRentalLocationStateProvince

`func (o *AdditionalDataCarRental) HasCarRentalLocationStateProvince() bool`

HasCarRentalLocationStateProvince returns a boolean if a field has been set.

### GetCarRentalNoShowIndicator

`func (o *AdditionalDataCarRental) GetCarRentalNoShowIndicator() string`

GetCarRentalNoShowIndicator returns the CarRentalNoShowIndicator field if non-nil, zero value otherwise.

### GetCarRentalNoShowIndicatorOk

`func (o *AdditionalDataCarRental) GetCarRentalNoShowIndicatorOk() (*string, bool)`

GetCarRentalNoShowIndicatorOk returns a tuple with the CarRentalNoShowIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarRentalNoShowIndicator

`func (o *AdditionalDataCarRental) SetCarRentalNoShowIndicator(v string)`

SetCarRentalNoShowIndicator sets CarRentalNoShowIndicator field to given value.

### HasCarRentalNoShowIndicator

`func (o *AdditionalDataCarRental) HasCarRentalNoShowIndicator() bool`

HasCarRentalNoShowIndicator returns a boolean if a field has been set.

### GetCarRentalOneWayDropOffCharges

`func (o *AdditionalDataCarRental) GetCarRentalOneWayDropOffCharges() string`

GetCarRentalOneWayDropOffCharges returns the CarRentalOneWayDropOffCharges field if non-nil, zero value otherwise.

### GetCarRentalOneWayDropOffChargesOk

`func (o *AdditionalDataCarRental) GetCarRentalOneWayDropOffChargesOk() (*string, bool)`

GetCarRentalOneWayDropOffChargesOk returns a tuple with the CarRentalOneWayDropOffCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarRentalOneWayDropOffCharges

`func (o *AdditionalDataCarRental) SetCarRentalOneWayDropOffCharges(v string)`

SetCarRentalOneWayDropOffCharges sets CarRentalOneWayDropOffCharges field to given value.

### HasCarRentalOneWayDropOffCharges

`func (o *AdditionalDataCarRental) HasCarRentalOneWayDropOffCharges() bool`

HasCarRentalOneWayDropOffCharges returns a boolean if a field has been set.

### GetCarRentalRate

`func (o *AdditionalDataCarRental) GetCarRentalRate() string`

GetCarRentalRate returns the CarRentalRate field if non-nil, zero value otherwise.

### GetCarRentalRateOk

`func (o *AdditionalDataCarRental) GetCarRentalRateOk() (*string, bool)`

GetCarRentalRateOk returns a tuple with the CarRentalRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarRentalRate

`func (o *AdditionalDataCarRental) SetCarRentalRate(v string)`

SetCarRentalRate sets CarRentalRate field to given value.

### HasCarRentalRate

`func (o *AdditionalDataCarRental) HasCarRentalRate() bool`

HasCarRentalRate returns a boolean if a field has been set.

### GetCarRentalRateIndicator

`func (o *AdditionalDataCarRental) GetCarRentalRateIndicator() string`

GetCarRentalRateIndicator returns the CarRentalRateIndicator field if non-nil, zero value otherwise.

### GetCarRentalRateIndicatorOk

`func (o *AdditionalDataCarRental) GetCarRentalRateIndicatorOk() (*string, bool)`

GetCarRentalRateIndicatorOk returns a tuple with the CarRentalRateIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarRentalRateIndicator

`func (o *AdditionalDataCarRental) SetCarRentalRateIndicator(v string)`

SetCarRentalRateIndicator sets CarRentalRateIndicator field to given value.

### HasCarRentalRateIndicator

`func (o *AdditionalDataCarRental) HasCarRentalRateIndicator() bool`

HasCarRentalRateIndicator returns a boolean if a field has been set.

### GetCarRentalRentalAgreementNumber

`func (o *AdditionalDataCarRental) GetCarRentalRentalAgreementNumber() string`

GetCarRentalRentalAgreementNumber returns the CarRentalRentalAgreementNumber field if non-nil, zero value otherwise.

### GetCarRentalRentalAgreementNumberOk

`func (o *AdditionalDataCarRental) GetCarRentalRentalAgreementNumberOk() (*string, bool)`

GetCarRentalRentalAgreementNumberOk returns a tuple with the CarRentalRentalAgreementNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarRentalRentalAgreementNumber

`func (o *AdditionalDataCarRental) SetCarRentalRentalAgreementNumber(v string)`

SetCarRentalRentalAgreementNumber sets CarRentalRentalAgreementNumber field to given value.

### HasCarRentalRentalAgreementNumber

`func (o *AdditionalDataCarRental) HasCarRentalRentalAgreementNumber() bool`

HasCarRentalRentalAgreementNumber returns a boolean if a field has been set.

### GetCarRentalRentalClassId

`func (o *AdditionalDataCarRental) GetCarRentalRentalClassId() string`

GetCarRentalRentalClassId returns the CarRentalRentalClassId field if non-nil, zero value otherwise.

### GetCarRentalRentalClassIdOk

`func (o *AdditionalDataCarRental) GetCarRentalRentalClassIdOk() (*string, bool)`

GetCarRentalRentalClassIdOk returns a tuple with the CarRentalRentalClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarRentalRentalClassId

`func (o *AdditionalDataCarRental) SetCarRentalRentalClassId(v string)`

SetCarRentalRentalClassId sets CarRentalRentalClassId field to given value.

### HasCarRentalRentalClassId

`func (o *AdditionalDataCarRental) HasCarRentalRentalClassId() bool`

HasCarRentalRentalClassId returns a boolean if a field has been set.

### GetCarRentalRenterName

`func (o *AdditionalDataCarRental) GetCarRentalRenterName() string`

GetCarRentalRenterName returns the CarRentalRenterName field if non-nil, zero value otherwise.

### GetCarRentalRenterNameOk

`func (o *AdditionalDataCarRental) GetCarRentalRenterNameOk() (*string, bool)`

GetCarRentalRenterNameOk returns a tuple with the CarRentalRenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarRentalRenterName

`func (o *AdditionalDataCarRental) SetCarRentalRenterName(v string)`

SetCarRentalRenterName sets CarRentalRenterName field to given value.

### HasCarRentalRenterName

`func (o *AdditionalDataCarRental) HasCarRentalRenterName() bool`

HasCarRentalRenterName returns a boolean if a field has been set.

### GetCarRentalReturnCity

`func (o *AdditionalDataCarRental) GetCarRentalReturnCity() string`

GetCarRentalReturnCity returns the CarRentalReturnCity field if non-nil, zero value otherwise.

### GetCarRentalReturnCityOk

`func (o *AdditionalDataCarRental) GetCarRentalReturnCityOk() (*string, bool)`

GetCarRentalReturnCityOk returns a tuple with the CarRentalReturnCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarRentalReturnCity

`func (o *AdditionalDataCarRental) SetCarRentalReturnCity(v string)`

SetCarRentalReturnCity sets CarRentalReturnCity field to given value.

### HasCarRentalReturnCity

`func (o *AdditionalDataCarRental) HasCarRentalReturnCity() bool`

HasCarRentalReturnCity returns a boolean if a field has been set.

### GetCarRentalReturnCountry

`func (o *AdditionalDataCarRental) GetCarRentalReturnCountry() string`

GetCarRentalReturnCountry returns the CarRentalReturnCountry field if non-nil, zero value otherwise.

### GetCarRentalReturnCountryOk

`func (o *AdditionalDataCarRental) GetCarRentalReturnCountryOk() (*string, bool)`

GetCarRentalReturnCountryOk returns a tuple with the CarRentalReturnCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarRentalReturnCountry

`func (o *AdditionalDataCarRental) SetCarRentalReturnCountry(v string)`

SetCarRentalReturnCountry sets CarRentalReturnCountry field to given value.

### HasCarRentalReturnCountry

`func (o *AdditionalDataCarRental) HasCarRentalReturnCountry() bool`

HasCarRentalReturnCountry returns a boolean if a field has been set.

### GetCarRentalReturnDate

`func (o *AdditionalDataCarRental) GetCarRentalReturnDate() string`

GetCarRentalReturnDate returns the CarRentalReturnDate field if non-nil, zero value otherwise.

### GetCarRentalReturnDateOk

`func (o *AdditionalDataCarRental) GetCarRentalReturnDateOk() (*string, bool)`

GetCarRentalReturnDateOk returns a tuple with the CarRentalReturnDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarRentalReturnDate

`func (o *AdditionalDataCarRental) SetCarRentalReturnDate(v string)`

SetCarRentalReturnDate sets CarRentalReturnDate field to given value.

### HasCarRentalReturnDate

`func (o *AdditionalDataCarRental) HasCarRentalReturnDate() bool`

HasCarRentalReturnDate returns a boolean if a field has been set.

### GetCarRentalReturnLocationId

`func (o *AdditionalDataCarRental) GetCarRentalReturnLocationId() string`

GetCarRentalReturnLocationId returns the CarRentalReturnLocationId field if non-nil, zero value otherwise.

### GetCarRentalReturnLocationIdOk

`func (o *AdditionalDataCarRental) GetCarRentalReturnLocationIdOk() (*string, bool)`

GetCarRentalReturnLocationIdOk returns a tuple with the CarRentalReturnLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarRentalReturnLocationId

`func (o *AdditionalDataCarRental) SetCarRentalReturnLocationId(v string)`

SetCarRentalReturnLocationId sets CarRentalReturnLocationId field to given value.

### HasCarRentalReturnLocationId

`func (o *AdditionalDataCarRental) HasCarRentalReturnLocationId() bool`

HasCarRentalReturnLocationId returns a boolean if a field has been set.

### GetCarRentalReturnStateProvince

`func (o *AdditionalDataCarRental) GetCarRentalReturnStateProvince() string`

GetCarRentalReturnStateProvince returns the CarRentalReturnStateProvince field if non-nil, zero value otherwise.

### GetCarRentalReturnStateProvinceOk

`func (o *AdditionalDataCarRental) GetCarRentalReturnStateProvinceOk() (*string, bool)`

GetCarRentalReturnStateProvinceOk returns a tuple with the CarRentalReturnStateProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarRentalReturnStateProvince

`func (o *AdditionalDataCarRental) SetCarRentalReturnStateProvince(v string)`

SetCarRentalReturnStateProvince sets CarRentalReturnStateProvince field to given value.

### HasCarRentalReturnStateProvince

`func (o *AdditionalDataCarRental) HasCarRentalReturnStateProvince() bool`

HasCarRentalReturnStateProvince returns a boolean if a field has been set.

### GetCarRentalTaxExemptIndicator

`func (o *AdditionalDataCarRental) GetCarRentalTaxExemptIndicator() string`

GetCarRentalTaxExemptIndicator returns the CarRentalTaxExemptIndicator field if non-nil, zero value otherwise.

### GetCarRentalTaxExemptIndicatorOk

`func (o *AdditionalDataCarRental) GetCarRentalTaxExemptIndicatorOk() (*string, bool)`

GetCarRentalTaxExemptIndicatorOk returns a tuple with the CarRentalTaxExemptIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarRentalTaxExemptIndicator

`func (o *AdditionalDataCarRental) SetCarRentalTaxExemptIndicator(v string)`

SetCarRentalTaxExemptIndicator sets CarRentalTaxExemptIndicator field to given value.

### HasCarRentalTaxExemptIndicator

`func (o *AdditionalDataCarRental) HasCarRentalTaxExemptIndicator() bool`

HasCarRentalTaxExemptIndicator returns a boolean if a field has been set.

### GetTravelEntertainmentAuthDataDuration

`func (o *AdditionalDataCarRental) GetTravelEntertainmentAuthDataDuration() string`

GetTravelEntertainmentAuthDataDuration returns the TravelEntertainmentAuthDataDuration field if non-nil, zero value otherwise.

### GetTravelEntertainmentAuthDataDurationOk

`func (o *AdditionalDataCarRental) GetTravelEntertainmentAuthDataDurationOk() (*string, bool)`

GetTravelEntertainmentAuthDataDurationOk returns a tuple with the TravelEntertainmentAuthDataDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelEntertainmentAuthDataDuration

`func (o *AdditionalDataCarRental) SetTravelEntertainmentAuthDataDuration(v string)`

SetTravelEntertainmentAuthDataDuration sets TravelEntertainmentAuthDataDuration field to given value.

### HasTravelEntertainmentAuthDataDuration

`func (o *AdditionalDataCarRental) HasTravelEntertainmentAuthDataDuration() bool`

HasTravelEntertainmentAuthDataDuration returns a boolean if a field has been set.

### GetTravelEntertainmentAuthDataMarket

`func (o *AdditionalDataCarRental) GetTravelEntertainmentAuthDataMarket() string`

GetTravelEntertainmentAuthDataMarket returns the TravelEntertainmentAuthDataMarket field if non-nil, zero value otherwise.

### GetTravelEntertainmentAuthDataMarketOk

`func (o *AdditionalDataCarRental) GetTravelEntertainmentAuthDataMarketOk() (*string, bool)`

GetTravelEntertainmentAuthDataMarketOk returns a tuple with the TravelEntertainmentAuthDataMarket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelEntertainmentAuthDataMarket

`func (o *AdditionalDataCarRental) SetTravelEntertainmentAuthDataMarket(v string)`

SetTravelEntertainmentAuthDataMarket sets TravelEntertainmentAuthDataMarket field to given value.

### HasTravelEntertainmentAuthDataMarket

`func (o *AdditionalDataCarRental) HasTravelEntertainmentAuthDataMarket() bool`

HasTravelEntertainmentAuthDataMarket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


