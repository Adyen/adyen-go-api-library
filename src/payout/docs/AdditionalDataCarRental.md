# AdditionalDataCarRental

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarRentalRentalAgreementNumber** | **string** | The rental agreement number associated with this car rental. * Format: Alphanumeric * maxLength: 14 | [optional] 
**CarRentalRenterName** | **string** | The name of the person renting the car. * Format: Alphanumeric * maxLength: 26 | [optional] 
**CarRentalReturnCity** | **string** | The city where the car must be returned. * Format: Alphanumeric * maxLength: 18 | [optional] 
**CarRentalReturnStateProvince** | **string** | The state or province where the car must be returned. * Format: Alphanumeric * maxLength: 3 | [optional] 
**CarRentalReturnCountry** | **string** | The country where the car must be returned. * Format: Alphanumeric * maxLength: 2 | [optional] 
**CarRentalReturnLocationId** | **string** | Agency code, phone number, or address abbreviation * Format: Alphanumeric * maxLength: 10 | [optional] 
**CarRentalReturnDate** | **string** | The last date to return the car by. * Date format: &#x60;yyyyMMdd&#x60; | [optional] 
**CarRentalCheckOutDate** | **string** | Pick-up date. * Date format: &#x60;yyyyMMdd&#x60; | [optional] 
**CarRentalCustomerServiceTollFreeNumber** | **string** | The customer service phone number of the car rental company. * Format: Alphanumeric * maxLength: 17 | [optional] 
**CarRentalRate** | **string** | Daily rental rate. * Format: Alphanumeric * maxLength: 12 | [optional] 
**CarRentalRateIndicator** | **string** | Specifies whether the given rate is applied daily or weekly. * D - Daily rate. * W - Weekly rate. | [optional] 
**CarRentalLocationCity** | **string** | The location from which the car is rented. * Format: Alphanumeric * maxLength: 18 | [optional] 
**CarRentalLocationStateProvince** | **string** | Pick-up date. * Date format: &#x60;yyyyMMdd&#x60; | [optional] 
**CarRentalLocationCountry** | **string** | The customer service phone number of the car rental company. * Format: Alphanumeric * maxLength: 17 | [optional] 
**CarRentalRentalClassId** | **string** | Daily rental rate. * Format: Alphanumeric * maxLength: 12 | [optional] 
**CarRentalDaysRented** | **string** | Specifies whether the given rate is applied daily or weekly. * D - Daily rate. * W - Weekly rate. | [optional] 
**CarRentalTaxExemptIndicator** | **string** | Indicates whether the goods or services were tax-exempt, or tax was not collected.  Values: * 0 - Tax was not collected * 1 - Goods or services were tax exempt | [optional] 
**TravelEntertainmentAuthDataMarket** | **string** | Indicates what market-specific dataset will be submitted or is being submitted. Value should be \&quot;A\&quot; for Car rental. This should be included in the auth message. * Format: Alphanumeric * maxLength: 1 | [optional] 
**TravelEntertainmentAuthDataDuration** | **string** | Number of nights.  This should be included in the auth message. * Format: Numeric * maxLength: 2 | [optional] 
**CarRentalFuelCharges** | **string** | Any fuel charges associated with the rental. * Format: Numeric * maxLength: 12 | [optional] 
**CarRentalInsuranceCharges** | **string** | Any insurance charges associated with the rental. * Format: Numeric * maxLength: 12 | [optional] 
**CarRentalNoShowIndicator** | **string** | Indicates if the customer was a \&quot;no-show\&quot; (neither keeps nor cancels their booking). * 0 - Not applicable. * 1 - Customer was a no show. | [optional] 
**CarRentalOneWayDropOffCharges** | **string** | Charge associated with not returning a vehicle to the original rental location. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


