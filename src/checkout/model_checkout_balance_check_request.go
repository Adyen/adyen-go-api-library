/*
 * Adyen Checkout API
 *
 * Adyen Checkout API provides a simple and flexible way to initiate and authorise online payments. You can use the same integration for payments made with cards (including 3D Secure), mobile wallets, and local payment methods (for example, iDEAL and Sofort).  This API reference provides information on available endpoints and how to interact with them. To learn more about the API, visit [Checkout documentation](https://docs.adyen.com/checkout).  ## Authentication Each request to the Checkout API must be signed with an API key. For this, obtain an API Key from your Customer Area, as described in [How to get the API key](https://docs.adyen.com/user-management/how-to-get-the-api-key). Then set this key to the `X-API-Key` header value, for example:  ``` curl -H \"Content-Type: application/json\" \\ -H \"X-API-Key: Your_Checkout_API_key\" \\ ... ``` Note that when going live, you need to generate a new API Key to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Checkout API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://checkout-test.adyen.com/v67/payments ```
 *
 * API version: 65
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package checkout

import (
	"time"
)

// CheckoutBalanceCheckRequest struct for CheckoutBalanceCheckRequest
type CheckoutBalanceCheckRequest struct {
	AccountInfo      *AccountInfo `json:"accountInfo,omitempty"`
	AdditionalAmount *Amount      `json:"additionalAmount,omitempty"`
	// This field contains additional data, which may be required for a particular payment request.  The `additionalData` object consists of entries, each of which includes the key and value.
	AdditionalData  *map[string]interface{} `json:"additionalData,omitempty"`
	Amount          *Amount                 `json:"amount,omitempty"`
	ApplicationInfo *ApplicationInfo        `json:"applicationInfo,omitempty"`
	BillingAddress  *Address                `json:"billingAddress,omitempty"`
	BrowserInfo     *BrowserInfo            `json:"browserInfo,omitempty"`
	// The delay between the authorisation and scheduled auto-capture, specified in hours.
	CaptureDelayHours int32 `json:"captureDelayHours,omitempty"`
	// The shopper's date of birth.  Format [ISO-8601](https://www.w3.org/TR/NOTE-datetime): YYYY-MM-DD
	DateOfBirth     *time.Time  `json:"dateOfBirth,omitempty"`
	DccQuote        *ForexQuote `json:"dccQuote,omitempty"`
	DeliveryAddress *Address    `json:"deliveryAddress,omitempty"`
	// The date and time the purchased goods should be delivered.  Format [ISO 8601](https://www.w3.org/TR/NOTE-datetime): YYYY-MM-DDThh:mm:ss.sssTZD  Example: 2017-07-17T13:42:40.428+01:00
	DeliveryDate *time.Time `json:"deliveryDate,omitempty"`
	// A string containing the shopper's device fingerprint. For more information, refer to [Device fingerprinting](https://docs.adyen.com/risk-management/device-fingerprinting).
	DeviceFingerprint string `json:"deviceFingerprint,omitempty"`
	// An integer value that is added to the normal fraud score. The value can be either positive or negative.
	FraudOffset  int32         `json:"fraudOffset,omitempty"`
	Installments *Installments `json:"installments,omitempty"`
	// The [merchant category code](https://en.wikipedia.org/wiki/Merchant_category_code) (MCC) is a four-digit number, which relates to a particular market segment. This code reflects the predominant activity that is conducted by the merchant.
	Mcc string `json:"mcc,omitempty"`
	// The merchant account identifier, with which you want to process the transaction.
	MerchantAccount string `json:"merchantAccount"`
	// This reference allows linking multiple transactions to each other for reporting purposes (i.e. order auth-rate). The reference should be unique per billing cycle. The same merchant order reference should never be reused after the first authorised attempt. If used, this field should be supplied for all incoming authorisations. > We strongly recommend you send the `merchantOrderReference` value to benefit from linking payment requests when authorisation retries take place. In addition, we recommend you provide `retry.orderAttemptNumber`, `retry.chainAttemptNumber`, and `retry.skipRetry` values in `PaymentRequest.additionalData`.
	MerchantOrderReference string                 `json:"merchantOrderReference,omitempty"`
	MerchantRiskIndicator  *MerchantRiskIndicator `json:"merchantRiskIndicator,omitempty"`
	// Metadata consists of entries, each of which includes a key and a value. Limitations: Maximum 20 key-value pairs per request. When exceeding, the \"177\" error occurs: \"Metadata size exceeds limit\".
	Metadata map[string]string `json:"metadata,omitempty"`
	// When you are doing multiple partial (gift card) payments, this is the `pspReference` of the first payment. We use this to link the multiple payments to each other. As your own reference for linking multiple payments, use the `merchantOrderReference`instead.
	OrderReference string `json:"orderReference,omitempty"`
	// The collection that contains the type of the payment method and its specific information.
	PaymentMethod map[string]interface{} `json:"paymentMethod"`
	Recurring     *Recurring             `json:"recurring,omitempty"`
	// Defines a recurring payment type. Allowed values: * `Subscription` – A transaction for a fixed or variable amount, which follows a fixed schedule. * `CardOnFile` – With a card-on-file (CoF) transaction, card details are stored to enable one-click or omnichannel journeys, or simply to streamline the checkout process. Any subscription not following a fixed schedule is also considered a card-on-file transaction. * `UnscheduledCardOnFile` – An unscheduled card-on-file (UCoF) transaction is a transaction that occurs on a non-fixed schedule and/or have variable amounts. For example, automatic top-ups when a cardholder's balance drops below a certain amount.
	RecurringProcessingModel string `json:"recurringProcessingModel,omitempty"`
	// The reference to uniquely identify a payment. This reference is used in all communication with you about the payment status. We recommend using a unique value per payment; however, it is not a requirement. If you need to provide multiple references for a transaction, separate them with hyphens (\"-\"). Maximum length: 80 characters.
	Reference string `json:"reference,omitempty"`
	// Some payment methods require defining a value for this field to specify how to process the transaction.  For the Bancontact payment method, it can be set to: * `maestro` (default), to be processed like a Maestro card, or * `bcmc`, to be processed like a Bancontact card.
	SelectedBrand string `json:"selectedBrand,omitempty"`
	// The `recurringDetailReference` you want to use for this payment. The value `LATEST` can be used to select the most recently stored recurring detail.
	SelectedRecurringDetailReference string `json:"selectedRecurringDetailReference,omitempty"`
	// A session ID used to identify a payment session.
	SessionId string `json:"sessionId,omitempty"`
	// The shopper's email address. We recommend that you provide this data, as it is used in velocity fraud checks. > For 3D Secure 2 transactions, schemes require the `shopperEmail` for both `deviceChannel` **browser** and **app**.
	ShopperEmail string `json:"shopperEmail,omitempty"`
	// The shopper's IP address. In general, we recommend that you provide this data, as it is used in a number of risk checks (for instance, number of payment attempts or location-based checks). > Required for 3D Secure 2 transactions. This field is also mandatory for some merchants depending on your business model. For more information, [contact Support](https://support.adyen.com/hc/en-us/requests/new).
	ShopperIP string `json:"shopperIP,omitempty"`
	// Specifies the sales channel, through which the shopper gives their card details, and whether the shopper is a returning customer. For the web service API, Adyen assumes Ecommerce shopper interaction by default.  This field has the following possible values: * `Ecommerce` - Online transactions where the cardholder is present (online). For better authorisation rates, we recommend sending the card security code (CSC) along with the request. * `ContAuth` - Card on file and/or subscription transactions, where the cardholder is known to the merchant (returning customer). If the shopper is present (online), you can supply also the CSC to improve authorisation (one-click payment). * `Moto` - Mail-order and telephone-order transactions where the shopper is in contact with the merchant via email or telephone. * `POS` - Point-of-sale transactions where the shopper is physically present to make a payment using a secure payment terminal.
	ShopperInteraction string `json:"shopperInteraction,omitempty"`
	// The combination of a language code and a country code to specify the language to be used in the payment.
	ShopperLocale string `json:"shopperLocale,omitempty"`
	ShopperName   *Name  `json:"shopperName,omitempty"`
	// Your reference to uniquely identify this shopper (for example, user ID or account ID). Minimum length: 3 characters. > This field is required for recurring payments.
	ShopperReference string `json:"shopperReference,omitempty"`
	// The text to be shown on the shopper's bank statement. To enable this field, contact our [Support Team](https://support.adyen.com/hc/en-us/requests/new).  We recommend sending a maximum of 25 characters, otherwise banks might truncate the string.
	ShopperStatement string `json:"shopperStatement,omitempty"`
	// The shopper's social security number.
	SocialSecurityNumber string `json:"socialSecurityNumber,omitempty"`
	// Information on how the payment should be split between accounts when using [Adyen for Platforms](https://docs.adyen.com/platforms/processing-payments#providing-split-information).
	Splits *[]Split `json:"splits,omitempty"`
	// The physical store, for which this payment is processed.
	Store string `json:"store,omitempty"`
	// The shopper's telephone number.
	TelephoneNumber     string               `json:"telephoneNumber,omitempty"`
	ThreeDS2RequestData *ThreeDS2RequestData `json:"threeDS2RequestData,omitempty"`
	// If set to true, you will only perform the [3D Secure 2 authentication](https://docs.adyen.com/checkout/3d-secure/other-3ds-flows/authentication-only), and not the payment authorisation.
	ThreeDSAuthenticationOnly bool `json:"threeDSAuthenticationOnly,omitempty"`
	// The reference value to aggregate sales totals in reporting. When not specified, the store field is used (if available).
	TotalsGroup string `json:"totalsGroup,omitempty"`
	// Set to true if the payment should be routed to a trusted MID.
	TrustedShopper bool `json:"trustedShopper,omitempty"`
}