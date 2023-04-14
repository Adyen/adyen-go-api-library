# AdditionalDataAirline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AirlineAgencyInvoiceNumber** | Pointer to **string** | Reference number for the invoice, issued by the agency. * minLength: 1 * maxLength: 6 | [optional] 
**AirlineAgencyPlanName** | Pointer to **string** | 2-letter agency plan identifier; alphabetical. * minLength: 2 * maxLength: 2 | [optional] 
**AirlineAirlineCode** | Pointer to **string** | [IATA](https://www.iata.org/services/pages/codes.aspx) 3-digit accounting code (PAX); numeric. It identifies the carrier. * Format: IATA 3-digit accounting code (PAX) * Example: KLM &#x3D; 074 * minLength: 3 * maxLength: 3 | [optional] 
**AirlineAirlineDesignatorCode** | Pointer to **string** | [IATA](https://www.iata.org/services/pages/codes.aspx) 2-letter accounting code (PAX); alphabetical. It identifies the carrier. * Format: [IATA](https://www.iata.org/services/pages/codes.aspx) 2-letter airline code * Example: KLM &#x3D; KL * minLength: 2 * maxLength: 2 | [optional] 
**AirlineBoardingFee** | Pointer to **string** | Chargeable amount for boarding the plane. The transaction amount needs to be represented in minor units according to the [following table](https://docs.adyen.com/development-resources/currency-codes). * minLength: 1 * maxLength: 18 | [optional] 
**AirlineComputerizedReservationSystem** | Pointer to **string** | The [CRS](https://en.wikipedia.org/wiki/Computer_reservation_system) used to make the reservation and purchase the ticket. * Format: alphanumeric. * minLength: 4 * maxLength: 4 | [optional] 
**AirlineCustomerReferenceNumber** | Pointer to **string** | Reference number; alphanumeric. * minLength: 0 * maxLength: 20 | [optional] 
**AirlineDocumentType** | Pointer to **string** | Optional 2-digit code; alphanumeric. It identifies the type of product of the transaction. The description of the code may appear on credit card statements. * Format: 2-digit code * Example: Passenger ticket &#x3D; 01 * minLength: 2 * maxLength: 2 | [optional] 
**AirlineFlightDate** | Pointer to **string** | Flight departure date. Local time &#x60;(HH:mm)&#x60; is optional. * Date format: &#x60;yyyy-MM-dd&#x60; * Date and time format: &#x60;yyyy-MM-dd HH:mm&#x60; * minLength: 10 * maxLength: 16 | [optional] 
**AirlineLegCarrierCode** | Pointer to **string** | [IATA](https://www.iata.org/services/pages/codes.aspx) 2-letter accounting code (PAX); alphabetical. It identifies the carrier. This field is required/mandatory if the airline data includes leg details. * Format: IATA 2-letter airline code * Example: KLM &#x3D; KL * minLength: 2 * maxLength: 2 | [optional] 
**AirlineLegClassOfTravel** | Pointer to **string** | 1-letter travel class identifier; alphabetical. There is no standard; however, the following codes are used rather consistently: * F: first class * J: business class * Y: economy class * W: premium economy  Limitations: * minLength: 1 * maxLength: 1 | [optional] 
**AirlineLegDateOfTravel** | Pointer to **string** |   Date and time of travel. [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)-compliant. * Format: &#x60;yyyy-MM-dd HH:mm&#x60; * minLength: 16 * maxLength: 16 | [optional] 
**AirlineLegDepartAirport** | Pointer to **string** | Alphabetical identifier of the departure airport. This field is required if the airline data includes leg details. * Format: [IATA](https://www.iata.org/services/pages/codes.aspx) 3-letter airport code. * Example: Amsterdam &#x3D; AMS * minLength: 3 * maxLength: 3 | [optional] 
**AirlineLegDepartTax** | Pointer to **string** | [Departure tax](https://en.wikipedia.org/wiki/Departure_tax). Amount charged by a country to an individual upon their leaving. The transaction amount needs to be represented in minor units according to the [following table](https://docs.adyen.com/development-resources/currency-codes). * minLength: 1 * maxLength: 12 | [optional] 
**AirlineLegDestinationCode** | Pointer to **string** | Alphabetical identifier of the destination/arrival airport. This field is required/mandatory if the airline data includes leg details. * Format: [IATA](https://www.iata.org/services/pages/codes.aspx) 3-letter airport code. * Example: Amsterdam &#x3D; AMS * minLength: 3 * maxLength: 3 | [optional] 
**AirlineLegFareBaseCode** | Pointer to **string** | [Fare basis code](https://en.wikipedia.org/wiki/Fare_basis_code); alphanumeric. * minLength: 1 * maxLength: 7 | [optional] 
**AirlineLegFlightNumber** | Pointer to **string** | The flight identifier. * minLength: 1 * maxLength: 5 | [optional] 
**AirlineLegStopOverCode** | Pointer to **string** | 1-letter code that indicates whether the passenger is entitled to make a stopover. Only two types of characters are allowed: * O: Stopover allowed * X: Stopover not allowed  Limitations: * minLength: 1 * maxLength: 1 | [optional] 
**AirlinePassengerDateOfBirth** | Pointer to **string** | Date of birth of the passenger.  Date format: &#x60;yyyy-MM-dd&#x60; * minLength: 10 * maxLength: 10 | [optional] 
**AirlinePassengerFirstName** | Pointer to **string** | Passenger first name/given name. &gt; This field is required/mandatory if the airline data includes passenger details or leg details. | [optional] 
**AirlinePassengerLastName** | Pointer to **string** | Passenger last name/family name. &gt; This field is required/mandatory if the airline data includes passenger details or leg details. | [optional] 
**AirlinePassengerTelephoneNumber** | Pointer to **string** | Telephone number of the passenger, including country code. This is an alphanumeric field that can include the &#39;+&#39; and &#39;-&#39; signs. * minLength: 3 * maxLength: 30 | [optional] 
**AirlinePassengerTravellerType** | Pointer to **string** | Passenger type code (PTC). IATA PTC values are 3-letter alphabetical. Example: ADT, SRC, CNN, INS.  However, several carriers use non-standard codes that can be up to 5 alphanumeric characters. * minLength: 3 * maxLength: 6 | [optional] 
**AirlinePassengerName** | **string** | Passenger name, initials, and a title. * Format: last name + first name or initials + title. * Example: *FLYER / MARY MS*. * minLength: 1 * maxLength: 49 | 
**AirlineTicketIssueAddress** | Pointer to **string** | Address of the place/agency that issued the ticket. * minLength: 0 * maxLength: 16 | [optional] 
**AirlineTicketNumber** | Pointer to **string** | The ticket&#39;s unique identifier. * minLength: 1 * maxLength: 150 | [optional] 
**AirlineTravelAgencyCode** | Pointer to **string** | IATA number, also ARC number or ARC/IATA number. Unique identifier number for travel agencies. * minLength: 1 * maxLength: 8 | [optional] 
**AirlineTravelAgencyName** | Pointer to **string** | The name of the travel agency. * minLength: 1 * maxLength: 25 | [optional] 

## Methods

### NewAdditionalDataAirline

`func NewAdditionalDataAirline(airlinePassengerName string, ) *AdditionalDataAirline`

NewAdditionalDataAirline instantiates a new AdditionalDataAirline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalDataAirlineWithDefaults

`func NewAdditionalDataAirlineWithDefaults() *AdditionalDataAirline`

NewAdditionalDataAirlineWithDefaults instantiates a new AdditionalDataAirline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAirlineAgencyInvoiceNumber

`func (o *AdditionalDataAirline) GetAirlineAgencyInvoiceNumber() string`

GetAirlineAgencyInvoiceNumber returns the AirlineAgencyInvoiceNumber field if non-nil, zero value otherwise.

### GetAirlineAgencyInvoiceNumberOk

`func (o *AdditionalDataAirline) GetAirlineAgencyInvoiceNumberOk() (*string, bool)`

GetAirlineAgencyInvoiceNumberOk returns a tuple with the AirlineAgencyInvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlineAgencyInvoiceNumber

`func (o *AdditionalDataAirline) SetAirlineAgencyInvoiceNumber(v string)`

SetAirlineAgencyInvoiceNumber sets AirlineAgencyInvoiceNumber field to given value.

### HasAirlineAgencyInvoiceNumber

`func (o *AdditionalDataAirline) HasAirlineAgencyInvoiceNumber() bool`

HasAirlineAgencyInvoiceNumber returns a boolean if a field has been set.

### GetAirlineAgencyPlanName

`func (o *AdditionalDataAirline) GetAirlineAgencyPlanName() string`

GetAirlineAgencyPlanName returns the AirlineAgencyPlanName field if non-nil, zero value otherwise.

### GetAirlineAgencyPlanNameOk

`func (o *AdditionalDataAirline) GetAirlineAgencyPlanNameOk() (*string, bool)`

GetAirlineAgencyPlanNameOk returns a tuple with the AirlineAgencyPlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlineAgencyPlanName

`func (o *AdditionalDataAirline) SetAirlineAgencyPlanName(v string)`

SetAirlineAgencyPlanName sets AirlineAgencyPlanName field to given value.

### HasAirlineAgencyPlanName

`func (o *AdditionalDataAirline) HasAirlineAgencyPlanName() bool`

HasAirlineAgencyPlanName returns a boolean if a field has been set.

### GetAirlineAirlineCode

`func (o *AdditionalDataAirline) GetAirlineAirlineCode() string`

GetAirlineAirlineCode returns the AirlineAirlineCode field if non-nil, zero value otherwise.

### GetAirlineAirlineCodeOk

`func (o *AdditionalDataAirline) GetAirlineAirlineCodeOk() (*string, bool)`

GetAirlineAirlineCodeOk returns a tuple with the AirlineAirlineCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlineAirlineCode

`func (o *AdditionalDataAirline) SetAirlineAirlineCode(v string)`

SetAirlineAirlineCode sets AirlineAirlineCode field to given value.

### HasAirlineAirlineCode

`func (o *AdditionalDataAirline) HasAirlineAirlineCode() bool`

HasAirlineAirlineCode returns a boolean if a field has been set.

### GetAirlineAirlineDesignatorCode

`func (o *AdditionalDataAirline) GetAirlineAirlineDesignatorCode() string`

GetAirlineAirlineDesignatorCode returns the AirlineAirlineDesignatorCode field if non-nil, zero value otherwise.

### GetAirlineAirlineDesignatorCodeOk

`func (o *AdditionalDataAirline) GetAirlineAirlineDesignatorCodeOk() (*string, bool)`

GetAirlineAirlineDesignatorCodeOk returns a tuple with the AirlineAirlineDesignatorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlineAirlineDesignatorCode

`func (o *AdditionalDataAirline) SetAirlineAirlineDesignatorCode(v string)`

SetAirlineAirlineDesignatorCode sets AirlineAirlineDesignatorCode field to given value.

### HasAirlineAirlineDesignatorCode

`func (o *AdditionalDataAirline) HasAirlineAirlineDesignatorCode() bool`

HasAirlineAirlineDesignatorCode returns a boolean if a field has been set.

### GetAirlineBoardingFee

`func (o *AdditionalDataAirline) GetAirlineBoardingFee() string`

GetAirlineBoardingFee returns the AirlineBoardingFee field if non-nil, zero value otherwise.

### GetAirlineBoardingFeeOk

`func (o *AdditionalDataAirline) GetAirlineBoardingFeeOk() (*string, bool)`

GetAirlineBoardingFeeOk returns a tuple with the AirlineBoardingFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlineBoardingFee

`func (o *AdditionalDataAirline) SetAirlineBoardingFee(v string)`

SetAirlineBoardingFee sets AirlineBoardingFee field to given value.

### HasAirlineBoardingFee

`func (o *AdditionalDataAirline) HasAirlineBoardingFee() bool`

HasAirlineBoardingFee returns a boolean if a field has been set.

### GetAirlineComputerizedReservationSystem

`func (o *AdditionalDataAirline) GetAirlineComputerizedReservationSystem() string`

GetAirlineComputerizedReservationSystem returns the AirlineComputerizedReservationSystem field if non-nil, zero value otherwise.

### GetAirlineComputerizedReservationSystemOk

`func (o *AdditionalDataAirline) GetAirlineComputerizedReservationSystemOk() (*string, bool)`

GetAirlineComputerizedReservationSystemOk returns a tuple with the AirlineComputerizedReservationSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlineComputerizedReservationSystem

`func (o *AdditionalDataAirline) SetAirlineComputerizedReservationSystem(v string)`

SetAirlineComputerizedReservationSystem sets AirlineComputerizedReservationSystem field to given value.

### HasAirlineComputerizedReservationSystem

`func (o *AdditionalDataAirline) HasAirlineComputerizedReservationSystem() bool`

HasAirlineComputerizedReservationSystem returns a boolean if a field has been set.

### GetAirlineCustomerReferenceNumber

`func (o *AdditionalDataAirline) GetAirlineCustomerReferenceNumber() string`

GetAirlineCustomerReferenceNumber returns the AirlineCustomerReferenceNumber field if non-nil, zero value otherwise.

### GetAirlineCustomerReferenceNumberOk

`func (o *AdditionalDataAirline) GetAirlineCustomerReferenceNumberOk() (*string, bool)`

GetAirlineCustomerReferenceNumberOk returns a tuple with the AirlineCustomerReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlineCustomerReferenceNumber

`func (o *AdditionalDataAirline) SetAirlineCustomerReferenceNumber(v string)`

SetAirlineCustomerReferenceNumber sets AirlineCustomerReferenceNumber field to given value.

### HasAirlineCustomerReferenceNumber

`func (o *AdditionalDataAirline) HasAirlineCustomerReferenceNumber() bool`

HasAirlineCustomerReferenceNumber returns a boolean if a field has been set.

### GetAirlineDocumentType

`func (o *AdditionalDataAirline) GetAirlineDocumentType() string`

GetAirlineDocumentType returns the AirlineDocumentType field if non-nil, zero value otherwise.

### GetAirlineDocumentTypeOk

`func (o *AdditionalDataAirline) GetAirlineDocumentTypeOk() (*string, bool)`

GetAirlineDocumentTypeOk returns a tuple with the AirlineDocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlineDocumentType

`func (o *AdditionalDataAirline) SetAirlineDocumentType(v string)`

SetAirlineDocumentType sets AirlineDocumentType field to given value.

### HasAirlineDocumentType

`func (o *AdditionalDataAirline) HasAirlineDocumentType() bool`

HasAirlineDocumentType returns a boolean if a field has been set.

### GetAirlineFlightDate

`func (o *AdditionalDataAirline) GetAirlineFlightDate() string`

GetAirlineFlightDate returns the AirlineFlightDate field if non-nil, zero value otherwise.

### GetAirlineFlightDateOk

`func (o *AdditionalDataAirline) GetAirlineFlightDateOk() (*string, bool)`

GetAirlineFlightDateOk returns a tuple with the AirlineFlightDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlineFlightDate

`func (o *AdditionalDataAirline) SetAirlineFlightDate(v string)`

SetAirlineFlightDate sets AirlineFlightDate field to given value.

### HasAirlineFlightDate

`func (o *AdditionalDataAirline) HasAirlineFlightDate() bool`

HasAirlineFlightDate returns a boolean if a field has been set.

### GetAirlineLegCarrierCode

`func (o *AdditionalDataAirline) GetAirlineLegCarrierCode() string`

GetAirlineLegCarrierCode returns the AirlineLegCarrierCode field if non-nil, zero value otherwise.

### GetAirlineLegCarrierCodeOk

`func (o *AdditionalDataAirline) GetAirlineLegCarrierCodeOk() (*string, bool)`

GetAirlineLegCarrierCodeOk returns a tuple with the AirlineLegCarrierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlineLegCarrierCode

`func (o *AdditionalDataAirline) SetAirlineLegCarrierCode(v string)`

SetAirlineLegCarrierCode sets AirlineLegCarrierCode field to given value.

### HasAirlineLegCarrierCode

`func (o *AdditionalDataAirline) HasAirlineLegCarrierCode() bool`

HasAirlineLegCarrierCode returns a boolean if a field has been set.

### GetAirlineLegClassOfTravel

`func (o *AdditionalDataAirline) GetAirlineLegClassOfTravel() string`

GetAirlineLegClassOfTravel returns the AirlineLegClassOfTravel field if non-nil, zero value otherwise.

### GetAirlineLegClassOfTravelOk

`func (o *AdditionalDataAirline) GetAirlineLegClassOfTravelOk() (*string, bool)`

GetAirlineLegClassOfTravelOk returns a tuple with the AirlineLegClassOfTravel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlineLegClassOfTravel

`func (o *AdditionalDataAirline) SetAirlineLegClassOfTravel(v string)`

SetAirlineLegClassOfTravel sets AirlineLegClassOfTravel field to given value.

### HasAirlineLegClassOfTravel

`func (o *AdditionalDataAirline) HasAirlineLegClassOfTravel() bool`

HasAirlineLegClassOfTravel returns a boolean if a field has been set.

### GetAirlineLegDateOfTravel

`func (o *AdditionalDataAirline) GetAirlineLegDateOfTravel() string`

GetAirlineLegDateOfTravel returns the AirlineLegDateOfTravel field if non-nil, zero value otherwise.

### GetAirlineLegDateOfTravelOk

`func (o *AdditionalDataAirline) GetAirlineLegDateOfTravelOk() (*string, bool)`

GetAirlineLegDateOfTravelOk returns a tuple with the AirlineLegDateOfTravel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlineLegDateOfTravel

`func (o *AdditionalDataAirline) SetAirlineLegDateOfTravel(v string)`

SetAirlineLegDateOfTravel sets AirlineLegDateOfTravel field to given value.

### HasAirlineLegDateOfTravel

`func (o *AdditionalDataAirline) HasAirlineLegDateOfTravel() bool`

HasAirlineLegDateOfTravel returns a boolean if a field has been set.

### GetAirlineLegDepartAirport

`func (o *AdditionalDataAirline) GetAirlineLegDepartAirport() string`

GetAirlineLegDepartAirport returns the AirlineLegDepartAirport field if non-nil, zero value otherwise.

### GetAirlineLegDepartAirportOk

`func (o *AdditionalDataAirline) GetAirlineLegDepartAirportOk() (*string, bool)`

GetAirlineLegDepartAirportOk returns a tuple with the AirlineLegDepartAirport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlineLegDepartAirport

`func (o *AdditionalDataAirline) SetAirlineLegDepartAirport(v string)`

SetAirlineLegDepartAirport sets AirlineLegDepartAirport field to given value.

### HasAirlineLegDepartAirport

`func (o *AdditionalDataAirline) HasAirlineLegDepartAirport() bool`

HasAirlineLegDepartAirport returns a boolean if a field has been set.

### GetAirlineLegDepartTax

`func (o *AdditionalDataAirline) GetAirlineLegDepartTax() string`

GetAirlineLegDepartTax returns the AirlineLegDepartTax field if non-nil, zero value otherwise.

### GetAirlineLegDepartTaxOk

`func (o *AdditionalDataAirline) GetAirlineLegDepartTaxOk() (*string, bool)`

GetAirlineLegDepartTaxOk returns a tuple with the AirlineLegDepartTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlineLegDepartTax

`func (o *AdditionalDataAirline) SetAirlineLegDepartTax(v string)`

SetAirlineLegDepartTax sets AirlineLegDepartTax field to given value.

### HasAirlineLegDepartTax

`func (o *AdditionalDataAirline) HasAirlineLegDepartTax() bool`

HasAirlineLegDepartTax returns a boolean if a field has been set.

### GetAirlineLegDestinationCode

`func (o *AdditionalDataAirline) GetAirlineLegDestinationCode() string`

GetAirlineLegDestinationCode returns the AirlineLegDestinationCode field if non-nil, zero value otherwise.

### GetAirlineLegDestinationCodeOk

`func (o *AdditionalDataAirline) GetAirlineLegDestinationCodeOk() (*string, bool)`

GetAirlineLegDestinationCodeOk returns a tuple with the AirlineLegDestinationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlineLegDestinationCode

`func (o *AdditionalDataAirline) SetAirlineLegDestinationCode(v string)`

SetAirlineLegDestinationCode sets AirlineLegDestinationCode field to given value.

### HasAirlineLegDestinationCode

`func (o *AdditionalDataAirline) HasAirlineLegDestinationCode() bool`

HasAirlineLegDestinationCode returns a boolean if a field has been set.

### GetAirlineLegFareBaseCode

`func (o *AdditionalDataAirline) GetAirlineLegFareBaseCode() string`

GetAirlineLegFareBaseCode returns the AirlineLegFareBaseCode field if non-nil, zero value otherwise.

### GetAirlineLegFareBaseCodeOk

`func (o *AdditionalDataAirline) GetAirlineLegFareBaseCodeOk() (*string, bool)`

GetAirlineLegFareBaseCodeOk returns a tuple with the AirlineLegFareBaseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlineLegFareBaseCode

`func (o *AdditionalDataAirline) SetAirlineLegFareBaseCode(v string)`

SetAirlineLegFareBaseCode sets AirlineLegFareBaseCode field to given value.

### HasAirlineLegFareBaseCode

`func (o *AdditionalDataAirline) HasAirlineLegFareBaseCode() bool`

HasAirlineLegFareBaseCode returns a boolean if a field has been set.

### GetAirlineLegFlightNumber

`func (o *AdditionalDataAirline) GetAirlineLegFlightNumber() string`

GetAirlineLegFlightNumber returns the AirlineLegFlightNumber field if non-nil, zero value otherwise.

### GetAirlineLegFlightNumberOk

`func (o *AdditionalDataAirline) GetAirlineLegFlightNumberOk() (*string, bool)`

GetAirlineLegFlightNumberOk returns a tuple with the AirlineLegFlightNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlineLegFlightNumber

`func (o *AdditionalDataAirline) SetAirlineLegFlightNumber(v string)`

SetAirlineLegFlightNumber sets AirlineLegFlightNumber field to given value.

### HasAirlineLegFlightNumber

`func (o *AdditionalDataAirline) HasAirlineLegFlightNumber() bool`

HasAirlineLegFlightNumber returns a boolean if a field has been set.

### GetAirlineLegStopOverCode

`func (o *AdditionalDataAirline) GetAirlineLegStopOverCode() string`

GetAirlineLegStopOverCode returns the AirlineLegStopOverCode field if non-nil, zero value otherwise.

### GetAirlineLegStopOverCodeOk

`func (o *AdditionalDataAirline) GetAirlineLegStopOverCodeOk() (*string, bool)`

GetAirlineLegStopOverCodeOk returns a tuple with the AirlineLegStopOverCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlineLegStopOverCode

`func (o *AdditionalDataAirline) SetAirlineLegStopOverCode(v string)`

SetAirlineLegStopOverCode sets AirlineLegStopOverCode field to given value.

### HasAirlineLegStopOverCode

`func (o *AdditionalDataAirline) HasAirlineLegStopOverCode() bool`

HasAirlineLegStopOverCode returns a boolean if a field has been set.

### GetAirlinePassengerDateOfBirth

`func (o *AdditionalDataAirline) GetAirlinePassengerDateOfBirth() string`

GetAirlinePassengerDateOfBirth returns the AirlinePassengerDateOfBirth field if non-nil, zero value otherwise.

### GetAirlinePassengerDateOfBirthOk

`func (o *AdditionalDataAirline) GetAirlinePassengerDateOfBirthOk() (*string, bool)`

GetAirlinePassengerDateOfBirthOk returns a tuple with the AirlinePassengerDateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlinePassengerDateOfBirth

`func (o *AdditionalDataAirline) SetAirlinePassengerDateOfBirth(v string)`

SetAirlinePassengerDateOfBirth sets AirlinePassengerDateOfBirth field to given value.

### HasAirlinePassengerDateOfBirth

`func (o *AdditionalDataAirline) HasAirlinePassengerDateOfBirth() bool`

HasAirlinePassengerDateOfBirth returns a boolean if a field has been set.

### GetAirlinePassengerFirstName

`func (o *AdditionalDataAirline) GetAirlinePassengerFirstName() string`

GetAirlinePassengerFirstName returns the AirlinePassengerFirstName field if non-nil, zero value otherwise.

### GetAirlinePassengerFirstNameOk

`func (o *AdditionalDataAirline) GetAirlinePassengerFirstNameOk() (*string, bool)`

GetAirlinePassengerFirstNameOk returns a tuple with the AirlinePassengerFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlinePassengerFirstName

`func (o *AdditionalDataAirline) SetAirlinePassengerFirstName(v string)`

SetAirlinePassengerFirstName sets AirlinePassengerFirstName field to given value.

### HasAirlinePassengerFirstName

`func (o *AdditionalDataAirline) HasAirlinePassengerFirstName() bool`

HasAirlinePassengerFirstName returns a boolean if a field has been set.

### GetAirlinePassengerLastName

`func (o *AdditionalDataAirline) GetAirlinePassengerLastName() string`

GetAirlinePassengerLastName returns the AirlinePassengerLastName field if non-nil, zero value otherwise.

### GetAirlinePassengerLastNameOk

`func (o *AdditionalDataAirline) GetAirlinePassengerLastNameOk() (*string, bool)`

GetAirlinePassengerLastNameOk returns a tuple with the AirlinePassengerLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlinePassengerLastName

`func (o *AdditionalDataAirline) SetAirlinePassengerLastName(v string)`

SetAirlinePassengerLastName sets AirlinePassengerLastName field to given value.

### HasAirlinePassengerLastName

`func (o *AdditionalDataAirline) HasAirlinePassengerLastName() bool`

HasAirlinePassengerLastName returns a boolean if a field has been set.

### GetAirlinePassengerTelephoneNumber

`func (o *AdditionalDataAirline) GetAirlinePassengerTelephoneNumber() string`

GetAirlinePassengerTelephoneNumber returns the AirlinePassengerTelephoneNumber field if non-nil, zero value otherwise.

### GetAirlinePassengerTelephoneNumberOk

`func (o *AdditionalDataAirline) GetAirlinePassengerTelephoneNumberOk() (*string, bool)`

GetAirlinePassengerTelephoneNumberOk returns a tuple with the AirlinePassengerTelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlinePassengerTelephoneNumber

`func (o *AdditionalDataAirline) SetAirlinePassengerTelephoneNumber(v string)`

SetAirlinePassengerTelephoneNumber sets AirlinePassengerTelephoneNumber field to given value.

### HasAirlinePassengerTelephoneNumber

`func (o *AdditionalDataAirline) HasAirlinePassengerTelephoneNumber() bool`

HasAirlinePassengerTelephoneNumber returns a boolean if a field has been set.

### GetAirlinePassengerTravellerType

`func (o *AdditionalDataAirline) GetAirlinePassengerTravellerType() string`

GetAirlinePassengerTravellerType returns the AirlinePassengerTravellerType field if non-nil, zero value otherwise.

### GetAirlinePassengerTravellerTypeOk

`func (o *AdditionalDataAirline) GetAirlinePassengerTravellerTypeOk() (*string, bool)`

GetAirlinePassengerTravellerTypeOk returns a tuple with the AirlinePassengerTravellerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlinePassengerTravellerType

`func (o *AdditionalDataAirline) SetAirlinePassengerTravellerType(v string)`

SetAirlinePassengerTravellerType sets AirlinePassengerTravellerType field to given value.

### HasAirlinePassengerTravellerType

`func (o *AdditionalDataAirline) HasAirlinePassengerTravellerType() bool`

HasAirlinePassengerTravellerType returns a boolean if a field has been set.

### GetAirlinePassengerName

`func (o *AdditionalDataAirline) GetAirlinePassengerName() string`

GetAirlinePassengerName returns the AirlinePassengerName field if non-nil, zero value otherwise.

### GetAirlinePassengerNameOk

`func (o *AdditionalDataAirline) GetAirlinePassengerNameOk() (*string, bool)`

GetAirlinePassengerNameOk returns a tuple with the AirlinePassengerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlinePassengerName

`func (o *AdditionalDataAirline) SetAirlinePassengerName(v string)`

SetAirlinePassengerName sets AirlinePassengerName field to given value.


### GetAirlineTicketIssueAddress

`func (o *AdditionalDataAirline) GetAirlineTicketIssueAddress() string`

GetAirlineTicketIssueAddress returns the AirlineTicketIssueAddress field if non-nil, zero value otherwise.

### GetAirlineTicketIssueAddressOk

`func (o *AdditionalDataAirline) GetAirlineTicketIssueAddressOk() (*string, bool)`

GetAirlineTicketIssueAddressOk returns a tuple with the AirlineTicketIssueAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlineTicketIssueAddress

`func (o *AdditionalDataAirline) SetAirlineTicketIssueAddress(v string)`

SetAirlineTicketIssueAddress sets AirlineTicketIssueAddress field to given value.

### HasAirlineTicketIssueAddress

`func (o *AdditionalDataAirline) HasAirlineTicketIssueAddress() bool`

HasAirlineTicketIssueAddress returns a boolean if a field has been set.

### GetAirlineTicketNumber

`func (o *AdditionalDataAirline) GetAirlineTicketNumber() string`

GetAirlineTicketNumber returns the AirlineTicketNumber field if non-nil, zero value otherwise.

### GetAirlineTicketNumberOk

`func (o *AdditionalDataAirline) GetAirlineTicketNumberOk() (*string, bool)`

GetAirlineTicketNumberOk returns a tuple with the AirlineTicketNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlineTicketNumber

`func (o *AdditionalDataAirline) SetAirlineTicketNumber(v string)`

SetAirlineTicketNumber sets AirlineTicketNumber field to given value.

### HasAirlineTicketNumber

`func (o *AdditionalDataAirline) HasAirlineTicketNumber() bool`

HasAirlineTicketNumber returns a boolean if a field has been set.

### GetAirlineTravelAgencyCode

`func (o *AdditionalDataAirline) GetAirlineTravelAgencyCode() string`

GetAirlineTravelAgencyCode returns the AirlineTravelAgencyCode field if non-nil, zero value otherwise.

### GetAirlineTravelAgencyCodeOk

`func (o *AdditionalDataAirline) GetAirlineTravelAgencyCodeOk() (*string, bool)`

GetAirlineTravelAgencyCodeOk returns a tuple with the AirlineTravelAgencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlineTravelAgencyCode

`func (o *AdditionalDataAirline) SetAirlineTravelAgencyCode(v string)`

SetAirlineTravelAgencyCode sets AirlineTravelAgencyCode field to given value.

### HasAirlineTravelAgencyCode

`func (o *AdditionalDataAirline) HasAirlineTravelAgencyCode() bool`

HasAirlineTravelAgencyCode returns a boolean if a field has been set.

### GetAirlineTravelAgencyName

`func (o *AdditionalDataAirline) GetAirlineTravelAgencyName() string`

GetAirlineTravelAgencyName returns the AirlineTravelAgencyName field if non-nil, zero value otherwise.

### GetAirlineTravelAgencyNameOk

`func (o *AdditionalDataAirline) GetAirlineTravelAgencyNameOk() (*string, bool)`

GetAirlineTravelAgencyNameOk returns a tuple with the AirlineTravelAgencyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlineTravelAgencyName

`func (o *AdditionalDataAirline) SetAirlineTravelAgencyName(v string)`

SetAirlineTravelAgencyName sets AirlineTravelAgencyName field to given value.

### HasAirlineTravelAgencyName

`func (o *AdditionalDataAirline) HasAirlineTravelAgencyName() bool`

HasAirlineTravelAgencyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


