/*
 * Adyen Payout API
 *
 * A set of API endpoints that allow you to store payout details, confirm, or decline a payout.  For more information, refer to [Online payouts](https://docs.adyen.com/checkout/online-payouts).
 *
 * API version: 52
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package payouts

// BrowserInfo struct for BrowserInfo
type BrowserInfo struct {
	// The accept header value of the shopper's browser.
	AcceptHeader string `json:"acceptHeader"`
	// The color depth of the shopper's browser in bits per pixel. This should be obtained by using the browser's `screen.colorDepth` property. Accepted values: 1, 4, 8, 15, 16, 24, 32 or 48 bit color depth.
	ColorDepth int32 `json:"colorDepth"`
	// Boolean value indicating if the shopper's browser is able to execute Java.
	JavaEnabled bool `json:"javaEnabled"`
	// Boolean value indicating if the shopper's browser is able to execute JavaScript. A default 'true' value is assumed if the field is not present.
	JavaScriptEnabled bool `json:"javaScriptEnabled,omitempty"`
	// The `navigator.language` value of the shopper's browser (as defined in IETF BCP 47).
	Language string `json:"language"`
	// The total height of the shopper's device screen in pixels.
	ScreenHeight int32 `json:"screenHeight"`
	// The total width of the shopper's device screen in pixels.
	ScreenWidth int32 `json:"screenWidth"`
	// Time difference between UTC time and the shopper's browser local time, in minutes.
	TimeZoneOffset int32 `json:"timeZoneOffset"`
	// The user agent value of the shopper's browser.
	UserAgent string `json:"userAgent"`
}
