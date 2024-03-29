/*
 * Adyen for Platforms: Account API
 *
 * The Account API provides endpoints for managing account-related entities on your platform. These related entities include account holders, accounts, bank accounts, shareholders, and KYC-related documents. The management operations include actions such as creation, retrieval, updating, and deletion of them.  For more information, refer to our [documentation](https://docs.adyen.com/platforms). ## Authentication To connect to the Account API, you must use basic authentication credentials of your web service user. If you don't have one, please contact the [Adyen Support Team](https://support.adyen.com/hc/en-us/requests/new). Then use its credentials to authenticate your request, for example:  ``` curl -U \"ws@MarketPlace.YourMarketPlace\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning The Account API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://cal-test.adyen.com/cal/services/Account/v6/createAccountHolder ```
 *
 * API version: 6
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package platformsaccount

// ViasAddress struct for ViasAddress
type ViasAddress struct {
	// The name of the city. >Required if either `houseNumberOrName`, `street`, `postalCode`, or `stateOrProvince` are provided.
	City string `json:"city,omitempty"`
	// The two-character country code of the address. The permitted country codes are defined in ISO-3166-1 alpha-2 (e.g. 'NL'). > If you don't know the country or are not collecting the country from the shopper, provide `country` as `ZZ`.
	Country string `json:"country"`
	// The number or name of the house.
	HouseNumberOrName string `json:"houseNumberOrName,omitempty"`
	// The postal code. >A maximum of five (5) digits for an address in the USA, or a maximum of ten (10) characters for an address in all other countries. >Required if either `houseNumberOrName`, `street`, `city`, or `stateOrProvince` are provided.
	PostalCode string `json:"postalCode,omitempty"`
	// The abbreviation of the state or province. >Two (2) characters for an address in the USA or Canada, or a maximum of three (3) characters for an address in all other countries. >Required for an address in the USA or Canada if either `houseNumberOrName`, `street`, `city`, or `postalCode` are provided.
	StateOrProvince string `json:"stateOrProvince,omitempty"`
	// The name of the street. >The house number should not be included in this field; it should be separately provided via `houseNumberOrName`. >Required if either `houseNumberOrName`, `city`, `postalCode`, or `stateOrProvince` are provided.
	Street string `json:"street,omitempty"`
}
