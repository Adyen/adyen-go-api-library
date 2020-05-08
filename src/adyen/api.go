/*
 * Adyen API Client
 *
 * Contact: support@adyen.com
 */

package adyen

import (
	"fmt"
	"net/http"

	binlookup "github.com/adyen/adyen-go-api-library/src/binlookup"
	checkout "github.com/adyen/adyen-go-api-library/src/checkout"
	checkoututility "github.com/adyen/adyen-go-api-library/src/checkoututility"
	common "github.com/adyen/adyen-go-api-library/src/common"
	notification "github.com/adyen/adyen-go-api-library/src/notification"
	payments "github.com/adyen/adyen-go-api-library/src/payments"
	payouts "github.com/adyen/adyen-go-api-library/src/payouts"
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
)

// APIClient manages communication with the Adyen Checkout API API v51
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	client *common.Client
	// API Services
	Checkout        *checkout.Checkout
	CheckoutUtility *checkoututility.CheckoutUtility
	Payments        *payments.Payments
	Payouts         *payouts.Payouts
	Recurring       *recurring.Recurring
	BinLookup       *binlookup.BinLookup
	Notification    *notification.NotificationService
}

// NewClient creates a new API client. Requires Config object.
//
// @param config              A reference to common.Config object
//
// create a new API client based on provided api key & url prefix for LIVE environment
//
//  client := NewClient(&common.Config{
//     ApiKey:                "apiKey",
//     Environment:           common.LiveEnv,
//     LiveEndpointURLPrefix: "liveEndpointURLPrefix",
//  })
//
//  ApiKey                Defines the api key that can be retrieved by back office
//  Environment           This defines the payment environment live or test
//  LiveEndpointURLPrefix Provide the unique live url prefix from the "API URLs and Response" menu in the Adyen Customer Area
//
//
// create a new API client based on provided api key for TEST environment
//
//  client := NewClient(&common.Config{
//     ApiKey:                "apiKey",
//     Environment:           common.TestEnv,
//  })
//
//  ApiKey                Defines the api key that can be retrieved by back office
//  Environment           This defines the payment environment live or test
//
//
// creates a new API client based on provided credentials &  url prefix for LIVE environment
//
//  client := NewClient(&common.Config{
//      Username:              "username",
//      Password:              "password",
//      ApplicationName:       "applicationName",
//      Environment:           common.LiveEnv,
//      LiveEndpointURLPrefix: "liveEndpointURLPrefix",
//  })
//
//  Username              Your merchant account Username
//  Password              Your merchant accont Password
//  Environment           This defines the payment environment live or test
//  ApplicationName       Application name to be used in user agent
//  LiveEndpointUrlPrefix Provide the unique live url prefix from the "API URLs and Response" menu in the Adyen Customer Area
//
//
// creates a new API client based on provided credentials for TEST environment
//
//  client := NewClient(&common.Config{
//      Username:              "username",
//      Password:              "password",
//      ApplicationName:       "applicationName",
//      Environment:           common.TestEnv,
//  })
//
//  Username              Your merchant account Username
//  Password              Your merchant accont Password
//  Environment           This defines the payment environment live or test
//  ApplicationName       Application name to be used in user agent
//
// optionally a custom http.Client can be passed via the Config allow for advanced features such as caching.
func NewClient(cfg *common.Config) *APIClient {
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
		cfg.UserAgent = fmt.Sprintf("%s/%s", common.LibName, common.LibVersion)
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
			return fmt.Sprintf("%s/%s", c.client.Cfg.CheckoutEndpoint, CheckoutAPIVersion)
		},
	}
	c.CheckoutUtility = &checkoututility.CheckoutUtility{
		Client: c.client,
		BasePath: func() string {
			return fmt.Sprintf("%s/%s", c.client.Cfg.CheckoutEndpoint, CheckoutUtilityAPIVersion)
		},
	}
	c.Payments = &payments.Payments{
		Client: c.client,
		BasePath: func() string {
			return fmt.Sprintf("%s/pal/servlet/Payment/%s", c.client.Cfg.Endpoint, APIVersion)
		},
	}

	c.Payouts = &payouts.Payouts{
		Client: c.client,
		BasePath: func() string {
			return fmt.Sprintf("%s/pal/servlet/Payout/%s", c.client.Cfg.Endpoint, APIVersion)
		},
	}

	c.Recurring = &recurring.Recurring{
		Client: c.client,
		BasePath: func() string {
			return fmt.Sprintf("%s/pal/servlet/Recurring/%s", c.client.Cfg.Endpoint, RecurringAPIVersion)
		},
	}

	c.BinLookup = &binlookup.BinLookup{
		Client: c.client,
		BasePath: func() string {
			return fmt.Sprintf("%s/%s/%s", c.client.Cfg.Endpoint, BinLookupPalSuffix, BinLookupAPIVersion)
		},
	}

	return c
}

/*
SetEnvironment This defines the payment environment for live or test

 * @param environment           This defines the payment environment live or test
 * @param liveEndpointUrlPrefix Provide the unique live url prefix from the "API URLs and Response" menu in the Adyen Customer Area
*/
func (c *APIClient) SetEnvironment(env common.Environment, liveEndpointURLPrefix string) {
	if env == common.LiveEnv {
		c.client.Cfg.Environment = env
		c.client.Cfg.MarketPayEndpoint = MarketpayEndpointLive
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
		c.client.Cfg.CheckoutEndpoint = CheckoutEndpointTest
		c.client.Cfg.TerminalApiCloudEndpoint = TerminalAPIEndpointTest
	}
}

// GetConfig Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetConfig() *common.Config {
	return c.client.Cfg
}
