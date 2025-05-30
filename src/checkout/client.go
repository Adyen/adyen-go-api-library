/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// APIClient manages communication with the Adyen Checkout API API v71
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	common common.Service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	DonationsApi *DonationsApi

	ModificationsApi *ModificationsApi

	OrdersApi *OrdersApi

	PaymentLinksApi *PaymentLinksApi

	PaymentsApi *PaymentsApi

	RecurringApi *RecurringApi

	UtilityApi *UtilityApi
}

// NewAPIClient creates a new API client.
func NewAPIClient(client *common.Client) *APIClient {
	c := &APIClient{}
	c.common.Client = client
	c.common.BasePath = func() string {
		return client.Cfg.CheckoutEndpoint
	}

	// API Services
	c.DonationsApi = (*DonationsApi)(&c.common)
	c.ModificationsApi = (*ModificationsApi)(&c.common)
	c.OrdersApi = (*OrdersApi)(&c.common)
	c.PaymentLinksApi = (*PaymentLinksApi)(&c.common)
	c.PaymentsApi = (*PaymentsApi)(&c.common)
	c.RecurringApi = (*RecurringApi)(&c.common)
	c.UtilityApi = (*UtilityApi)(&c.common)

	return c
}
