/*
 * Adyen API Client
 *
 * Contact: support@adyen.com
 */

package main

import (
	"fmt"
	"net/http"

	checkout "github.com/adyen/adyen-go-api-library/src/checkout"
	checkoututility "github.com/adyen/adyen-go-api-library/src/checkoututility"
	common "github.com/adyen/adyen-go-api-library/src/common"
)

// APIClient manages communication with the Adyen Checkout API API v51
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	client *common.Client
	// API Services
	Checkout        *checkout.Checkout
	CheckoutUtility *checkoututility.CheckoutUtility
}

// NewConfig returns a new Config object
func NewConfig() *common.Config {
	cfg := &common.Config{
		BasePath:      "https://checkout-test.adyen.com/v51",
		DefaultHeader: make(map[string]string),
		UserAgent:     "OpenAPI-Generator/1.0.0/go",
		Debug:         false,
		Servers: []common.ServerConfiguration{
			{
				Url:         "https://checkout-test.adyen.com/v51",
				Description: "No description provided",
			},
		},
	}
	return cfg
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *common.Config) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.client = &common.Client{}
	c.client.Cfg = cfg
	c.client.Common.Client = c.client

	// API Services
	c.Checkout = (*checkout.Checkout)(&c.client.Common)

	return c
}

// ChangeBasePath changes base path to allow switching to mocks
func (c *APIClient) ChangeBasePath(path string) {
	c.client.Cfg.BasePath = path
}

// GetConfig Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetConfig() *common.Config {
	return c.client.Cfg
}

func main() {
	fmt.Println("Welcome to Adyen API Client")
}
