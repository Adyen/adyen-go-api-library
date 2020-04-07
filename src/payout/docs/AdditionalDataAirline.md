# AdditionalDataAirline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AirlinePassengerName** | **string** | Passenger name, initials, and a title. * Format: last name + first name or initials + title. * Example: *FLYER / MARY MS*. * minLength: 1 * maxLength: 49 | [optional] 
**AirlineComputerizedReservationSystem** | **string** | The [CRS](https://en.wikipedia.org/wiki/Computer_reservation_system) used to make the reservation and purchase the ticket. * Format: alphanumeric. * minLength: 4 * maxLength: 4 | [optional] 
**AirlineAirlineCode** | **string** | [IATA](https://www.iata.org/services/pages/codes.aspx) 3-digit accounting code (PAX); numeric. It identifies the carrier. * Format: IATA 3-digit accounting code (PAX) * Example: KLM &#x3D; 074 * minLength: 3 * maxLength: 3 | [optional] 
**AirlineTicketNumber** | **string** | The ticket&#39;s unique identifier. * minLength: 1 * maxLength: 150 | [optional] 
**AirlineFlightDate** | **string** | Flight departure date. Local time &#x60;(HH:mm)&#x60; is optional. * Date format: &#x60;yyyy-MM-dd&#x60; * Date and time format: &#x60;yyyy-MM-dd HH:mm&#x60; * minLength: 10 * maxLength: 16 | [optional] 
**AirlineCustomerReferenceNumber** | **string** | Reference number; alphanumeric. * minLength: 0 * maxLength: 20 | [optional] 
**AirlineTicketIssueAddress** | **string** | Address of the place/agency that issued the ticket. * minLength: 0 * maxLength: 16 | [optional] 
**AirlineAirlineDesignatorCode** | **string** | [IATA](https://www.iata.org/services/pages/codes.aspx) 2-letter accounting code (PAX); alphabetical. It identifies the carrier. * Format: [IATA](https://www.iata.org/services/pages/codes.aspx) 2-letter airline code * Example: KLM &#x3D; KL * minLength: 2 * maxLength: 2 | [optional] 
**AirlineTravelAgencyCode** | **string** | IATA number, also ARC number or ARC/IATA number. Unique identifier number for travel agencies. * minLength: 1 * maxLength: 8 | [optional] 
**AirlineTravelAgencyName** | **string** | The name of the travel agency. * minLength: 1 * maxLength: 25 | [optional] 
**AirlineAgencyPlanName** | **string** | 2-letter agency plan identifier; alphabetical. * minLength: 2 * maxLength: 2 | [optional] 
**AirlineAgencyInvoiceNumber** | **string** | Reference number for the invoice, issued by the agency. * minLength: 1 * maxLength: 6 | [optional] 
**AirlineBoardingFee** | **string** | Chargeable amount for boarding the plane. The transaction amount needs to be represented in minor units according to the [following table](https://docs.adyen.com/development-resources/currency-codes). * minLength: 1 * maxLength: 18 | [optional] 
**AirlineDocumentType** | **string** | Optional 2-digit code; alphanumeric. It identifies the type of product of the transaction. The description of the code may appear on credit card statements. * Format: 2-digit code * Example: Passenger ticket &#x3D; 01 * minLength: 2 * maxLength: 2 | [optional] 
**AirlineLegDepartAirport** | **string** | Alphabetical identifier of the departure airport. This field is required if the airline data includes leg details. * Format: [IATA](https://www.iata.org/services/pages/codes.aspx) 3-letter airport code. * Example: Amsterdam &#x3D; AMS * minLength: 3 * maxLength: 3 | [optional] 
**AirlineLegFlightNumber** | **string** | The flight identifier. * minLength: 1 * maxLength: 5 | [optional] 
**AirlineLegCarrierCode** | **string** | [IATA](https://www.iata.org/services/pages/codes.aspx) 2-letter accounting code (PAX); alphabetical. It identifies the carrier. This field is required/mandatory if the airline data includes leg details. * Format: IATA 2-letter airline code * Example: KLM &#x3D; KL * minLength: 2 * maxLength: 2 | [optional] 
**AirlineLegFareBaseCode** | **string** | [Fare basis code](https://en.wikipedia.org/wiki/Fare_basis_code); alphanumeric. * minLength: 1 * maxLength: 7 | [optional] 
**AirlineLegClassOfTravel** | **string** | 1-letter travel class identifier; alphabetical. There is no standard; however, the following codes are used rather consistently: * F: first class * J: business class * Y: economy class * W: premium economy  Limitations: * minLength: 1 * maxLength: 1 | [optional] 
**AirlineLegStopOverCode** | **string** | 1-letter code that indicates whether the passenger is entitled to make a stopover. Only two types of characters are allowed: * O: Stopover allowed * X: Stopover not allowed  Limitations: * minLength: 1 * maxLength: 1 | [optional] 
**AirlineLegDestinationCode** | **string** | Alphabetical identifier of the destination/arrival airport. This field is required/mandatory if the airline data includes leg details. * Format: [IATA](https://www.iata.org/services/pages/codes.aspx) 3-letter airport code. * Example: Amsterdam &#x3D; AMS * minLength: 3 * maxLength: 3 | [optional] 
**AirlineLegDateOfTravel** | **string** |   Date and time of travel. [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)-compliant. * Format: &#x60;yyyy-MM-dd HH:mm&#x60; * minLength: 16 * maxLength: 16 | [optional] 
**AirlineLegDepartTax** | **string** | [Departure tax](https://en.wikipedia.org/wiki/Departure_tax). Amount charged by a country to an individual upon their leaving. The transaction amount needs to be represented in minor units according to the [following table](https://docs.adyen.com/development-resources/currency-codes). * minLength: 1 * maxLength: 12 | [optional] 
**AirlinePassengerFirstName** | **string** | Passenger first name/given name. &gt; This field is required/mandatory if the airline data includes passenger details or leg details. | [optional] 
**AirlinePassengerLastName** | **string** | Passenger last name/family name. &gt; This field is required/mandatory if the airline data includes passenger details or leg details. | [optional] 
**AirlinePassengerTravellerType** | **string** | Passenger type code (PTC). IATA PTC values are 3-letter alphabetical. Example: ADT, SRC, CNN, INS.  However, several carriers use non-standard codes that can be up to 5 alphanumeric characters. * minLength: 3 * maxLength: 6 | [optional] 
**AirlinePassengerTelephoneNumber** | **string** | Telephone number of the passenger, including country code. This is an alphanumeric field that can include the &#39;+&#39; and &#39;-&#39; signs. * minLength: 3 * maxLength: 30 | [optional] 
**AirlinePassengerDateOfBirth** | **string** | Date of birth of the passenger.  Date format: &#x60;yyyy-MM-dd&#x60; * minLength: 10 * maxLength: 10 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


