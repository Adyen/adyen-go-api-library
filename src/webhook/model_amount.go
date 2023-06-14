package webhook

type Amount struct {
	// The three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes).
	Currency string `json:"currency"`
	// The payable amount that can be charged for the transaction.  The transaction amount needs to be represented in minor units according to the [following table](https://docs.adyen.com/development-resources/currency-codes).
	Value int64 `json:"value"`
}
