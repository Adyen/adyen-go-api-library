/*
 * Adyen Checkout API
 *
 * Adyen Checkout API provides a simple and flexible way to initiate and authorise online payments. You can use the same integration for payments made with cards (including 3D Secure), mobile wallets, and local payment methods (for example, iDEAL and Sofort).  This API reference provides information on available endpoints and how to interact with them. To learn more about the API, visit [online payments documentation](https://docs.adyen.com/online-payments).  ## Authentication Each request to Checkout API must be signed with an API key. For this, [get your API key](https://docs.adyen.com/development-resources/api-credentials#generate-api-key) from your Customer Area, and set this key to the `X-API-Key` header value, for example:  ``` curl -H \"Content-Type: application/json\" \\ -H \"X-API-Key: YOUR_API_KEY\" \\ ... ``` ## Versioning Checkout API supports [versioning](https://docs.adyen.com/development-resources/versioning) using a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://checkout-test.adyen.com/v70/payments ```  ## Going live  To access the live endpoints, you need an API key from your live Customer Area.  The live endpoint URLs contain a prefix which is unique to your company account, for example: ``` https://{PREFIX}-checkout-live.adyenpayments.com/checkout/v70/payments ```  Get your `{PREFIX}` from your live Customer Area under **Developers** > **API URLs** > **Prefix**.  When preparing to do live transactions with Checkout API, follow the [go-live checklist](https://docs.adyen.com/online-payments/go-live-checklist) to make sure you've got all the required configuration in place.  ## Release notes Have a look at the [release notes](https://docs.adyen.com/online-payments/release-notes?integration_type=api&version=70) to find out what changed in this version!
 *
 * API version: 70
 * Contact: developer-experience@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package checkout

// StoredPaymentMethodResource struct for StoredPaymentMethodResource
type StoredPaymentMethodResource struct {
	// The brand of the card.
	Brand string `json:"brand,omitempty"`
	// The month the card expires.
	ExpiryMonth string `json:"expiryMonth,omitempty"`
	// The last two digits of the year the card expires. For example, **22** for the year 2022.
	ExpiryYear string `json:"expiryYear,omitempty"`
	// The response code returned by an external system (for example after a provisioning operation).
	ExternalResponseCode string `json:"externalResponseCode,omitempty"`
	// The token reference of a linked token in an external system (for example a network token reference).
	ExternalTokenReference string `json:"externalTokenReference,omitempty"`
	// The unique payment method code.
	HolderName string `json:"holderName,omitempty"`
	// The IBAN of the bank account.
	Iban string `json:"iban,omitempty"`
	// A unique identifier of this stored payment method.
	Id string `json:"id,omitempty"`
	// The name of the issuer of token or card.
	IssuerName string `json:"issuerName,omitempty"`
	// The last four digits of the PAN.
	LastFour string `json:"lastFour,omitempty"`
	// The display name of the stored payment method.
	Name string `json:"name,omitempty"`
	// Returned in the response if you are not tokenizing with Adyen and are using the Merchant-initiated transactions (MIT) framework from Mastercard or Visa.  This contains either the Mastercard Trace ID or the Visa Transaction ID.
	NetworkTxReference string `json:"networkTxReference,omitempty"`
	// The name of the bank account holder.
	OwnerName string `json:"ownerName,omitempty"`
	// The shopper’s email address.
	ShopperEmail string `json:"shopperEmail,omitempty"`
	// Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. > Your reference must not include personally identifiable information (PII), for example name or email address.
	ShopperReference string `json:"shopperReference,omitempty"`
	// Defines a recurring payment type. Allowed values: * `Subscription` – A transaction for a fixed or variable amount, which follows a fixed schedule. * `CardOnFile` – With a card-on-file (CoF) transaction, card details are stored to enable one-click or omnichannel journeys, or simply to streamline the checkout process. Any subscription not following a fixed schedule is also considered a card-on-file transaction. * `UnscheduledCardOnFile` – An unscheduled card-on-file (UCoF) transaction is a transaction that occurs on a non-fixed schedule and/or have variable amounts. For example, automatic top-ups when a cardholder's balance drops below a certain amount.
	SupportedRecurringProcessingModels []string `json:"supportedRecurringProcessingModels,omitempty"`
	// The type of payment method.
	Type string `json:"type,omitempty"`
}