# AdditionalDataAirline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AirlineAgencyInvoiceNumber** | Pointer to **string** | The reference number for the invoice, issued by the agency. * Encoding: ASCII * minLength: 1 character * maxLength: 6 characters | [optional] 
**AirlineAgencyPlanName** | Pointer to **string** | The two-letter agency plan identifier. * Encoding: ASCII * minLength: 2 characters * maxLength: 2 characters | [optional] 
**AirlineAirlineCode** | Pointer to **string** | The [IATA](https://www.iata.org/services/pages/codes.aspx) 3-digit accounting code (PAX) that identifies the carrier. * Format: IATA 3-digit accounting code (PAX) * Example: KLM &#x3D; 074 * minLength: 3 characters * maxLength: 3 characters * Must not be all spaces * Must not be all zeros | [optional] 
**AirlineAirlineDesignatorCode** | Pointer to **string** | The [IATA](https://www.iata.org/services/pages/codes.aspx) 2-letter accounting code (PAX) that identifies the carrier. * Encoding: ASCII * Example: KLM &#x3D; KL * minLength: 2 characters * maxLength: 2 characters * Must not be all spaces * Must not be all zeros | [optional] 
**AirlineBoardingFee** | Pointer to **string** | The amount charged for boarding the plane, in [minor units](https://docs.adyen.com/development-resources/currency-codes). * Encoding: Numeric * minLength: 1 character * maxLength: 18 characters | [optional] 
**AirlineComputerizedReservationSystem** | Pointer to **string** | The [CRS](https://en.wikipedia.org/wiki/Computer_reservation_system) used to make the reservation and purchase the ticket. * Encoding: ASCII * minLength: 4 characters * maxLength: 4 characters | [optional] 
**AirlineCustomerReferenceNumber** | Pointer to **string** | The alphanumeric customer reference number. * Encoding: ASCII * maxLength: 20 characters * If you send more than 20 characters, the customer reference number is truncated * Must not be all spaces | [optional] 
**AirlineDocumentType** | Pointer to **string** | A code that identifies the type of item bought. The description of the code can appear on credit card statements. * Encoding: ASCII * Example: Passenger ticket &#x3D; 01 * minLength: 2 characters * maxLength: 2 characters | [optional] 
**AirlineFlightDate** | Pointer to **string** | The flight departure date. Local time &#x60;(HH:mm)&#x60; is optional. * Date format: &#x60;yyyy-MM-dd&#x60; * Date and time format: &#x60;yyyy-MM-dd HH:mm&#x60; * minLength: 10 characters * maxLength: 16 characters | [optional] 
**AirlineLegCarrierCode** | Pointer to **string** | The [IATA](https://www.iata.org/services/pages/codes.aspx) 2-letter accounting code (PAX) that identifies the carrier. This field is required if the airline data includes leg details. * Example: KLM &#x3D; KL * minLength: 2 characters * maxLength: 2 characters * Must not be all spaces * Must not be all zeros | [optional] 
**AirlineLegClassOfTravel** | Pointer to **string** | A one-letter travel class identifier.  The following are common:  * F: first class * J: business class * Y: economy class * W: premium economy  * Encoding: ASCII * minLength: 1 character * maxLength: 1 character * Must not be all spaces * Must not be all zeros | [optional] 
**AirlineLegDateOfTravel** | Pointer to **string** |   Date and time of travel in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format &#x60;yyyy-MM-dd HH:mm&#x60;. * Encoding: ASCII * minLength: 16 characters * maxLength: 16 characters | [optional] 
**AirlineLegDepartAirport** | Pointer to **string** | The [IATA](https://www.iata.org/services/pages/codes.aspx) three-letter airport code of the departure airport. This field is required if the airline data includes leg details.  * Encoding: ASCII * Example: Amsterdam &#x3D; AMS * minLength: 3 characters * maxLength: 3 characters * Must not be all spaces * Must not be all zeros | [optional] 
**AirlineLegDepartTax** | Pointer to **string** | The amount of [departure tax](https://en.wikipedia.org/wiki/Departure_tax) charged, in [minor units](https://docs.adyen.com/development-resources/currency-codes). * Encoding: Numeric * minLength: 1 * maxLength: 12 * Must not be all zeros | [optional] 
**AirlineLegDestinationCode** | Pointer to **string** | The [IATA](https://www.iata.org/services/pages/codes.aspx) 3-letter airport code of the destination airport. This field is required if the airline data includes leg details. * Example: Amsterdam &#x3D; AMS * Encoding: ASCII * minLength: 3 characters * maxLength: 3 characters * Must not be all spaces * Must not be all zeros | [optional] 
**AirlineLegFareBaseCode** | Pointer to **string** | The [fare basis code](https://en.wikipedia.org/wiki/Fare_basis_code), alphanumeric. * minLength: 1 character * maxLength: 6 characters * Must not be all spaces * Must not be all zeros | [optional] 
**AirlineLegFlightNumber** | Pointer to **string** | The flight identifier. * minLength: 1 character * maxLength: 5 characters * Must not be all spaces * Must not be all zeros | [optional] 
**AirlineLegStopOverCode** | Pointer to **string** | A one-letter code that indicates whether the passenger is entitled to make a stopover. Can be a space, O if the passenger is entitled to make a stopover, or X if they are not. * Encoding: ASCII * minLength: 1 character * maxLength: 1 character | [optional] 
**AirlinePassengerDateOfBirth** | Pointer to **string** | The passenger&#39;s date of birth.  Date format: &#x60;yyyy-MM-dd&#x60; * minLength: 10 * maxLength: 10 | [optional] 
**AirlinePassengerFirstName** | Pointer to **string** | The passenger&#39;s first name. &gt; This field is required if the airline data includes passenger details or leg details. * Encoding: ASCII | [optional] 
**AirlinePassengerLastName** | Pointer to **string** | The passenger&#39;s last name. &gt; This field is required if the airline data includes passenger details or leg details. * Encoding: ASCII | [optional] 
**AirlinePassengerTelephoneNumber** | Pointer to **string** | The passenger&#39;s telephone number, including country code. This is an alphanumeric field that can include the &#39;+&#39; and &#39;-&#39; signs. * Encoding: ASCII * minLength: 3 characters * maxLength: 30 characters | [optional] 
**AirlinePassengerTravellerType** | Pointer to **string** | The IATA passenger type code (PTC). * Encoding: ASCII * minLength: 3 characters * maxLength: 6 characters | [optional] 
**AirlinePassengerName** | **string** | The passenger&#39;s name, initials, and title. * Format: last name + first name or initials + title * Example: *FLYER / MARY MS* * minLength: 1 character * maxLength: 20 characters * If you send more than 20 characters, the name is truncated * Must not be all spaces * Must not be all zeros | 
**AirlineTicketIssueAddress** | Pointer to **string** | The address of the organization that issued the ticket. * minLength: 0 characters * maxLength: 16 characters | [optional] 
**AirlineTicketNumber** | Pointer to **string** | The ticket&#39;s unique identifier. * minLength: 1 character * maxLength: 15 characters * Must not be all spaces * Must not be all zeros | [optional] 
**AirlineTravelAgencyCode** | Pointer to **string** | The unique identifier from IATA or ARC for the travel agency that issues the ticket. * Encoding: ASCII * minLength: 1 character * maxLength: 8 characters * Must not be all spaces * Must not be all zeros | [optional] 
**AirlineTravelAgencyName** | Pointer to **string** | The name of the travel agency.  * Encoding: ASCII * minLength: 1 character * maxLength: 25 characters * Must not be all spaces * Must not be all zeros | [optional] 

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


