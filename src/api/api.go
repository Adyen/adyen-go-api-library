/*
 * Adyen API Client
 *
 * Contact: support@adyen.com
 */

package api

import (
	"fmt"
	"net/http"
	"time"

	checkout "github.com/adyen/adyen-go-api-library/src/checkout"
	checkoututility "github.com/adyen/adyen-go-api-library/src/checkoututility"
	common "github.com/adyen/adyen-go-api-library/src/common"
	payment "github.com/adyen/adyen-go-api-library/src/payment"
	recurring "github.com/adyen/adyen-go-api-library/src/recurring"
)

// Constants used for the client API
const (
	EndpointTest                    = "https://pal-test.adyen.com"
	EndpointLive                    = "https://pal-live.adyen.com"
	EndpointLiveSuffix              = "-pal-live.adyenpayments.com"
	HppTest                         = "https://test.adyen.com/hpp"
	HppLive                         = "https://live.adyen.com/hpp"
	MarketpayEndpointTest           = "https://cal-test.adyen.com/cal/services"
	MarketpayEndpointLive           = "https://cal-live.adyen.com/cal/services"
	MarketpayAccountAPIVersion      = "v5"
	MarketpayFundAPIVersion         = "v5"
	MarketpayNotificationAPIVersion = "v5"
	APIVersion                      = "v52"
	RecurringAPIVersion             = "v49"
	CheckoutEndpointTest            = "https://checkout-test.adyen.com/checkout"
	CheckoutEndpointLiveSuffix      = "-checkout-live.adyenpayments.com/checkout"
	CheckoutAPIVersion              = "v52"
	BinLookupPalSuffix              = "/pal/servlet/BinLookup/"
	BinLookupAPIVersion             = "v50"
	CheckoutUtilityAPIVersion       = "v1"
	TerminalAPIEndpointTest         = "https://terminal-api-test.adyen.com"
	TerminalAPIEndpointLive         = "https://terminal-api-live.adyen.com"
	EndpointProtocol                = "https://"
	LibName                         = "adyen-go-api-library"
	LibVersion                      = "0.0.1"
)

// APIClient manages communication with the Adyen Checkout API API v51
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	client *common.Client
	// API Services
	Checkout        *checkout.Checkout
	CheckoutUtility *checkoututility.CheckoutUtility
	Payment         *payment.Payment
	Recurring       *recurring.Recurring
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *common.Config) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}
	if cfg.ConnectionTimeoutMillis != 0 {
		cfg.HTTPClient.Timeout = cfg.ConnectionTimeoutMillis
	}
	if cfg.DefaultHeader == nil {
		cfg.DefaultHeader = make(map[string]string)
	}
	if cfg.UserAgent == "" {
		cfg.UserAgent = fmt.Sprintf("%s %s/%s", cfg.ApplicationName, LibName, LibVersion)
	}

	c := &APIClient{}
	c.client = &common.Client{}
	c.client.Cfg = cfg

	if cfg.Environment != "" {
		c.SetEnvironment(cfg.Environment, cfg.LiveEndpointURLPrefix)
	}

	// API Services
	c.Checkout = &checkout.Checkout{
		Client: c.client,
		BasePath: func() string {
			return getURL(c.client.Cfg.CheckoutEndpoint, CheckoutAPIVersion)
		},
	}
	c.CheckoutUtility = &checkoututility.CheckoutUtility{
		Client: c.client,
		BasePath: func() string {
			return getURL(c.client.Cfg.CheckoutEndpoint, CheckoutUtilityAPIVersion)
		},
	}
	c.Payment = &payment.Payment{
		Client: c.client,
		BasePath: func() string {
			return fmt.Sprintf("%s/pal/servlet/Payment/%s", c.client.Cfg.Endpoint, APIVersion)
		},
	}
	c.Recurring = &recurring.Recurring{
		Client: c.client,
		BasePath: func() string {
			return fmt.Sprintf("%s/pal/servlet/Recurring/%s", c.client.Cfg.Endpoint, RecurringAPIVersion)
		},
	}

	return c
}

// NewAPIClientWithCredentials creates a new API client based on provided credentials and environment
/*
 * @param username              Your merchant account Username
 * @param password              Your merchant accont Password
 * @param environment           This defines the payment environment live or test
 * @param applicationName       Application name to be used in user agent
 */
func NewAPIClientWithCredentials(username string, password string, environment common.Environment, applicationName string) *APIClient {
	return NewAPIClientWithLiveCredentials(username, password, environment, applicationName, "")
}

// NewAPIClientWithLiveCredentials creates a new API client based on provided credentials, environment & live environment url
/*
 * @param username              Your merchant account Username
 * @param password              Your merchant accont Password
 * @param environment           This defines the payment environment live or test
 * @param applicationName       Application name to be used in user agent
 * @param liveEndpointUrlPrefix Provide the unique live url prefix from the "API URLs and Response" menu in the Adyen Customer Area
 */
func NewAPIClientWithLiveCredentials(username string, password string, environment common.Environment, applicationName string, liveEndpointURLPrefix string) *APIClient {
	return NewAPIClient(&common.Config{
		Username:              username,
		Password:              password,
		ApplicationName:       applicationName,
		Environment:           environment,
		LiveEndpointURLPrefix: liveEndpointURLPrefix,
	})
}

// NewAPIClientWithAPIKey creates a new API client based on provided api key and environment
/*
 * @param apiKey                Defines the api key that can be retrieved by back office
 * @param environment           This defines the payment environment live or test
 */
func NewAPIClientWithAPIKey(apiKey string, environment common.Environment) *APIClient {
	return NewAPIClientWithLiveAPIKey(apiKey, environment, "")
}

// NewAPIClientWithLiveAPIKey creates a new API client based on provided api key, environment & live environment url
/*
 * @param apiKey                Defines the api key that can be retrieved by back office
 * @param environment           This defines the payment environment live or test
 * @param liveEndpointUrlPrefix Provide the unique live url prefix from the "API URLs and Response" menu in the Adyen Customer Area
 */
func NewAPIClientWithLiveAPIKey(apiKey string, environment common.Environment, liveEndpointURLPrefix string) *APIClient {
	return NewAPIClient(&common.Config{
		ApiKey:                apiKey,
		Environment:           environment,
		LiveEndpointURLPrefix: liveEndpointURLPrefix,
	})
}

/*
SetEnvironment This defines the payment environment for live or test

 * @param environment           This defines the payment environment live or test
 * @param liveEndpointUrlPrefix Provide the unique live url prefix from the "API URLs and Response" menu in the Adyen Customer Area
*/
func (c *APIClient) SetEnvironment(env common.Environment, liveEndpointURLPrefix string) {
	if env == common.Live {
		c.client.Cfg.Environment = env
		c.client.Cfg.MarketPayEndpoint = MarketpayEndpointLive
		c.client.Cfg.HppEndpoint = HppLive
		if liveEndpointURLPrefix != "" {
			c.client.Cfg.Endpoint = EndpointProtocol + liveEndpointURLPrefix + EndpointLiveSuffix
			c.client.Cfg.CheckoutEndpoint = EndpointProtocol + liveEndpointURLPrefix + CheckoutEndpointLiveSuffix
		} else {
			c.client.Cfg.Endpoint = EndpointLive
			c.client.Cfg.CheckoutEndpoint = ""
		}
		c.client.Cfg.TerminalApiCloudEndpoint = TerminalAPIEndpointLive
	} else {
		c.client.Cfg.Environment = env
		c.client.Cfg.Endpoint = EndpointTest
		c.client.Cfg.MarketPayEndpoint = MarketpayEndpointTest
		c.client.Cfg.HppEndpoint = HppTest
		c.client.Cfg.CheckoutEndpoint = CheckoutEndpointTest
		c.client.Cfg.TerminalApiCloudEndpoint = TerminalAPIEndpointTest
	}
}

// GetConfig Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetConfig() *common.Config {
	return c.client.Cfg
}

func (c *APIClient) getHTTPClient() *http.Client {
	if c.client.Cfg.HTTPClient == nil {
		c.client.Cfg.HTTPClient = http.DefaultClient
	}

	return c.client.Cfg.HTTPClient
}

func (c *APIClient) setHTTPClient(httpClient *http.Client) {
	c.client.Cfg.HTTPClient = httpClient
}

func (c *APIClient) setApplicationName(applicationName string) {
	c.client.Cfg.ApplicationName = applicationName
}

func (c *APIClient) setTimeouts(connectionTimeoutMillis, readTimeoutMillis time.Duration) {
	c.client.Cfg.ConnectionTimeoutMillis = connectionTimeoutMillis
	c.client.Cfg.ReadTimeoutMillis = readTimeoutMillis
}

func getURL(endpoint, version string) string {
	return fmt.Sprintf("%s/%s", endpoint, version)
}
