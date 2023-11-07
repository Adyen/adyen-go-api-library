/*
 * Adyen API Client
 *
 * Contact: support@adyen.com
 */

package adyen

import (
	"fmt"
	"github.com/adyen/adyen-go-api-library/v8/src/dataprotection"
	"net/http"

	"github.com/adyen/adyen-go-api-library/v8/src/balancecontrol"
	"github.com/adyen/adyen-go-api-library/v8/src/balanceplatform"
	"github.com/adyen/adyen-go-api-library/v8/src/binlookup"
	"github.com/adyen/adyen-go-api-library/v8/src/checkout"
	"github.com/adyen/adyen-go-api-library/v8/src/common"
	"github.com/adyen/adyen-go-api-library/v8/src/disputes"
	"github.com/adyen/adyen-go-api-library/v8/src/legalentity"
	"github.com/adyen/adyen-go-api-library/v8/src/management"
	"github.com/adyen/adyen-go-api-library/v8/src/payments"
	"github.com/adyen/adyen-go-api-library/v8/src/payout"
	"github.com/adyen/adyen-go-api-library/v8/src/platformsaccount"
	"github.com/adyen/adyen-go-api-library/v8/src/platformsfund"
	"github.com/adyen/adyen-go-api-library/v8/src/platformshostedonboardingpage"
	"github.com/adyen/adyen-go-api-library/v8/src/platformsnotificationconfiguration"
	"github.com/adyen/adyen-go-api-library/v8/src/posterminalmanagement"
	"github.com/adyen/adyen-go-api-library/v8/src/recurring"
	"github.com/adyen/adyen-go-api-library/v8/src/storedvalue"
	"github.com/adyen/adyen-go-api-library/v8/src/transfers"
)

// Constants used for the client API
const (
	EndpointTest                      = "https://pal-test.adyen.com"
	EndpointLive                      = "https://pal-live.adyen.com"
	EndpointLiveSuffix                = "-pal-live.adyenpayments.com"
	MarketpayEndpointTest             = "https://cal-test.adyen.com/cal/services"
	MarketpayEndpointLive             = "https://cal-live.adyen.com/cal/services"
	CheckoutEndpointTest              = "https://checkout-test.adyen.com/checkout"
	CheckoutEndpointLiveSuffix        = "-checkout-live.adyenpayments.com/checkout"
	TerminalAPIEndpointTest           = "https://terminal-api-test.adyen.com"
	TerminalAPIEndpointLive           = "https://terminal-api-live.adyen.com"
	DisputesEndpointTest              = "https://ca-test.adyen.com/ca/services/DisputeService"
	DisputesEndpointLive              = "https://ca-live.adyen.com/ca/services/DisputeService"
	BalancePlatformEndpointTest       = "https://balanceplatform-api-test.adyen.com/bcl"
	BalancePlatformEndpointLive       = "https://balanceplatform-api-live.adyen.com/bcl"
	TransfersEndpointTest             = "https://balanceplatform-api-test.adyen.com/btl"
	TransfersEndpointLive             = "https://balanceplatform-api-live.adyen.com/btl"
	ManagementEndpointTest            = "https://management-test.adyen.com"
	ManagementEndpointLive            = "https://management-live.adyen.com"
	LegalEntityEntityTest             = "https://kyc-test.adyen.com/lem"
	LegalEntityEntityLive             = "https://kyc-live.adyen.com/lem"
	PosTerminalManagementEndpointTest = "https://postfmapi-test.adyen.com/postfmapi/terminal"
	PosTerminalManagementEndpointLive = "https://postfmapi-live.adyen.com/postfmapi/terminal"
	DataProtectionEndpointTest        = "https://ca-test.adyen.com/ca/services/DataProtectionService"
	DataProtectionEndpointLive        = "https://ca-live.adyen.com/ca/services/DataProtectionService"
)

// also update LibVersion in src/common/configuration.go when a version is updated and a major lib version is released
const (
	MarketpayAccountAPIVersion      = "v6"
	MarketpayFundAPIVersion         = "v6"
	MarketpayNotificationAPIVersion = "v6"
	MarketpayHopAPIVersion          = "v6"
	PaymentAPIVersion               = "v68"
	RecurringAPIVersion             = "v68"
	CheckoutAPIVersion              = "v71"
	BinLookupAPIVersion             = "v54"
	BalanceControlAPIVersion        = "v1"
	EndpointProtocol                = "https://"
	DisputesAPIVersion              = "v30"
	StoredValueAPIVersion           = "v46"
	BalancePlatformAPIVersion       = "v2"
	TransfersAPIVersion             = "v4"
	ManagementAPIVersion            = "v3"
	LegalEntityAPIVersion           = "v3"
	PosTerminalManagementAPIVersion = "v1"
	DataProtectionAPIVersion        = "v1"
)

// APIClient Manages access to Adyen API services.
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	client *common.Client
	// API Services
	checkout       *checkout.APIClient
	payments       *payments.APIClient
	payout         *payout.APIClient
	recurring      *recurring.GeneralApi
	binLookup      *binlookup.GeneralApi
	balancecontrol *balancecontrol.GeneralApi
	// Deprecated: Please migrate to the new Adyen For Platforms.
	platformsAccount *platformsaccount.PlatformsAccount
	// Deprecated: Please migrate to the new Adyen For Platforms.
	platformsFund *platformsfund.PlatformsFund
	// Deprecated: Please migrate to the new Adyen For Platforms.
	platformsHostedOnboardingPage *platformshostedonboardingpage.PlatformsHostedOnboardingPage
	// Deprecated: Please migrate to the new Adyen For Platforms.
	platformsNotificationConfiguration *platformsnotificationconfiguration.PlatformsNotificationConfiguration
	posTerminalManagement              *posterminalmanagement.GeneralApi
	disputes                           *disputes.GeneralApi
	storedValue                        *storedvalue.GeneralApi
	balancePlatform                    *balanceplatform.APIClient
	transfers                          *transfers.APIClient
	management                         *management.APIClient
	legalEntity                        *legalentity.APIClient
	dataProtection                     *dataprotection.GeneralApi
}

// NewClient creates a new API client. Requires Config object.
//
// @param config              A reference to common.Config object
//
// create a new API client based on provided api key & url prefix for LIVE environment
//
//	client := NewClient(&common.Config{
//	   ApiKey:                "apiKey",
//	   Environment:           common.LiveEnv,
//	   LiveEndpointURLPrefix: "liveEndpointURLPrefix",
//	})
//
//	ApiKey                Defines the api key that can be retrieved by back office
//	Environment           This defines the payment environment live or test
//	LiveEndpointURLPrefix Provide the unique live url prefix from the "API URLs and Response" menu in the Adyen Customer Area
//
// create a new API client based on provided api key for TEST environment
//
//	client := NewClient(&common.Config{
//	   ApiKey:                "apiKey",
//	   Environment:           common.TestEnv,
//	})
//
//	ApiKey                Defines the api key that can be retrieved by back office
//	Environment           This defines the payment environment live or test
//
// creates a new API client based on provided credentials &  url prefix for LIVE environment
//
//	client := NewClient(&common.Config{
//	    Username:              "username",
//	    Password:              "password",
//	    ApplicationName:       "applicationName",
//	    Environment:           common.LiveEnv,
//	    LiveEndpointURLPrefix: "liveEndpointURLPrefix",
//	})
//
//	Username              Your merchant account Username
//	Password              Your merchant accont Password
//	Environment           This defines the payment environment live or test
//	ApplicationName       Application name to be used in user agent
//	LiveEndpointUrlPrefix Provide the unique live url prefix from the "API URLs and Response" menu in the Adyen Customer Area
//
// creates a new API client based on provided credentials for TEST environment
//
//	client := NewClient(&common.Config{
//	    Username:              "username",
//	    Password:              "password",
//	    ApplicationName:       "applicationName",
//	    Environment:           common.TestEnv,
//	})
//
//	Username              Your merchant account Username
//	Password              Your merchant accont Password
//	Environment           This defines the payment environment live or test
//	ApplicationName       Application name to be used in user agent
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

	c.SetEnvironment(cfg.Environment, cfg.LiveEndpointURLPrefix)

	return c
}

// API Services

func (c *APIClient) Checkout() *checkout.APIClient {
	if c.checkout == nil {
		c.checkout = checkout.NewAPIClient(c.client)
	}
	return c.checkout
}

func (c *APIClient) Payments() *payments.APIClient {
	if c.payments == nil {
		c.payments = payments.NewAPIClient(c.client)
		c.payments.PaymentsApi.BasePath = func() string {
			return fmt.Sprintf("%s/pal/servlet/Payment/%s", c.client.Cfg.Endpoint, PaymentAPIVersion)
		}
	}
	return c.payments
}

func (c *APIClient) Payout() *payout.APIClient {
	if c.payout == nil {
		c.payout = payout.NewAPIClient(c.client)
		c.payout.InitializationApi.BasePath = func() string {
			return fmt.Sprintf("%s/pal/servlet/Payout/%s", c.client.Cfg.Endpoint, PaymentAPIVersion)
		}
	}
	return c.payout
}

func (c *APIClient) Recurring() *recurring.GeneralApi {
	if c.recurring == nil {
		c.recurring = &recurring.GeneralApi{
			Client: c.client,
			BasePath: func() string {
				return fmt.Sprintf("%s/pal/servlet/Recurring/%s", c.client.Cfg.Endpoint, RecurringAPIVersion)
			},
		}
	}
	return c.recurring
}

func (c *APIClient) BinLookup() *binlookup.GeneralApi {
	if c.binLookup == nil {
		c.binLookup = &binlookup.GeneralApi{
			Client: c.client,
			BasePath: func() string {
				return fmt.Sprintf("%s/pal/servlet/BinLookup/%s", c.client.Cfg.Endpoint, BinLookupAPIVersion)
			},
		}
	}
	return c.binLookup
}

func (c *APIClient) BalanceControl() *balancecontrol.GeneralApi {
	if c.balancecontrol == nil {
		c.balancecontrol = &balancecontrol.GeneralApi{
			Client: c.client,
			BasePath: func() string {
				return fmt.Sprintf("%s/pal/servlet/BalanceControl/%s", c.client.Cfg.Endpoint, BalanceControlAPIVersion)
			},
		}
	}
	return c.balancecontrol
}

func (c *APIClient) StoredValue() *storedvalue.GeneralApi {
	if c.storedValue == nil {
		c.storedValue = &storedvalue.GeneralApi{
			Client: c.client,
			BasePath: func() string {
				return fmt.Sprintf("%s/pal/servlet/StoredValue/%s", c.client.Cfg.Endpoint, StoredValueAPIVersion)
			},
		}
	}
	return c.storedValue
}

func (c *APIClient) PosTerminalManagement() *posterminalmanagement.GeneralApi {
	if c.posTerminalManagement == nil {
		c.posTerminalManagement = &posterminalmanagement.GeneralApi{
			Client: c.client,
			BasePath: func() string {
				return fmt.Sprintf("%s/%s", c.client.Cfg.PosTerminalManagementEndpoint, PosTerminalManagementAPIVersion)
			},
		}
	}
	return c.posTerminalManagement
}

func (c *APIClient) Transfers() *transfers.APIClient {
	if c.transfers == nil {
		c.transfers = transfers.NewAPIClient(c.client)
		c.transfers.TransfersApi.BasePath = func() string {
			return fmt.Sprintf("%s/%s", c.client.Cfg.TransfersEndpoint, TransfersAPIVersion)
		}
	}
	return c.transfers
}

func (c *APIClient) BalancePlatform() *balanceplatform.APIClient {
	if c.balancePlatform == nil {
		c.balancePlatform = balanceplatform.NewAPIClient(c.client)
	}
	return c.balancePlatform
}

func (c *APIClient) Management() *management.APIClient {
	if c.management == nil {
		c.management = management.NewAPIClient(c.client)
	}
	return c.management
}

func (c *APIClient) LegalEntity() *legalentity.APIClient {
	if c.legalEntity == nil {
		c.legalEntity = legalentity.NewAPIClient(c.client)
	}
	return c.legalEntity
}

func (c *APIClient) Disputes() *disputes.GeneralApi {
	if c.disputes == nil {
		c.disputes = &disputes.GeneralApi{
			Client: c.client,
			BasePath: func() string {
				return fmt.Sprintf("%s/%s", c.client.Cfg.DisputesEndpoint, DisputesAPIVersion)
			},
		}
	}
	return c.disputes
}

// Deprecated: Please migrate to the new Adyen For Platforms.
func (c *APIClient) PlatformsAccount() *platformsaccount.PlatformsAccount {
	if c.platformsAccount == nil {
		c.platformsAccount = &platformsaccount.PlatformsAccount{
			Client: c.client,
			BasePath: func() string {
				return fmt.Sprintf("%s/Account/%s", c.client.Cfg.MarketPayEndpoint, MarketpayAccountAPIVersion)
			},
		}
	}
	return c.platformsAccount
}

// Deprecated: Please migrate to the new Adyen For Platforms.
func (c *APIClient) PlatformsFund() *platformsfund.PlatformsFund {
	if c.platformsFund == nil {
		c.platformsFund = &platformsfund.PlatformsFund{
			Client: c.client,
			BasePath: func() string {
				return fmt.Sprintf("%s/Fund/%s", c.client.Cfg.MarketPayEndpoint, MarketpayFundAPIVersion)
			},
		}
	}
	return c.platformsFund
}

// Deprecated: Please migrate to the new Adyen For Platforms.
func (c *APIClient) PlatformsHostedOnboardingPage() *platformshostedonboardingpage.PlatformsHostedOnboardingPage {
	if c.platformsHostedOnboardingPage == nil {
		c.platformsHostedOnboardingPage = &platformshostedonboardingpage.PlatformsHostedOnboardingPage{
			Client: c.client,
			BasePath: func() string {
				return fmt.Sprintf("%s/Hop/%s", c.client.Cfg.MarketPayEndpoint, MarketpayHopAPIVersion)
			},
		}
	}
	return c.platformsHostedOnboardingPage
}

// Deprecated: Please migrate to the new Adyen For Platforms.
func (c *APIClient) PlatformsNotificationConfiguration() *platformsnotificationconfiguration.PlatformsNotificationConfiguration {
	if c.platformsNotificationConfiguration == nil {
		c.platformsNotificationConfiguration = &platformsnotificationconfiguration.PlatformsNotificationConfiguration{
			Client: c.client,
			BasePath: func() string {
				return fmt.Sprintf("%s/Notification/%s", c.client.Cfg.MarketPayEndpoint, MarketpayNotificationAPIVersion)
			},
		}
	}
	return c.platformsNotificationConfiguration
}

func (c *APIClient) DataProtection() *dataprotection.GeneralApi {
	if c.dataProtection == nil {
		c.dataProtection = &dataprotection.GeneralApi{
			Client: c.client,
			BasePath: func() string {
				return fmt.Sprintf("%s/%s", c.client.Cfg.DataProtectionEndpoint, DataProtectionAPIVersion)
			},
		}
	}
	return c.dataProtection
}

/*
SetEnvironment This defines the payment environment for live or test

  - @param environment           This defines the payment environment live or test
  - @param liveEndpointUrlPrefix Provide the unique live url prefix from the "API URLs and Response" menu in the Adyen Customer Area
*/
func (c *APIClient) SetEnvironment(env common.Environment, liveEndpointURLPrefix string) {
	if env == common.LiveEnv {
		c.client.Cfg.Environment = env
		c.client.Cfg.MarketPayEndpoint = MarketpayEndpointLive
		c.client.Cfg.DisputesEndpoint = DisputesEndpointLive
		if liveEndpointURLPrefix != "" {
			c.client.Cfg.Endpoint = EndpointProtocol + liveEndpointURLPrefix + EndpointLiveSuffix
			c.client.Cfg.CheckoutEndpoint = EndpointProtocol + liveEndpointURLPrefix + CheckoutEndpointLiveSuffix
		} else {
			c.client.Cfg.Endpoint = EndpointLive
			c.client.Cfg.CheckoutEndpoint = ""
		}
		c.client.Cfg.TerminalApiCloudEndpoint = TerminalAPIEndpointLive
		c.client.Cfg.BalancePlatformEndpoint = BalancePlatformEndpointLive
		c.client.Cfg.TransfersEndpoint = TransfersEndpointLive
		c.client.Cfg.ManagementEndpoint = ManagementEndpointLive
		c.client.Cfg.LegalEntityEndpoint = LegalEntityEntityLive
		c.client.Cfg.PosTerminalManagementEndpoint = PosTerminalManagementEndpointLive
		c.client.Cfg.DataProtectionEndpoint = DataProtectionEndpointLive
	} else {
		c.client.Cfg.Environment = env
		c.client.Cfg.Endpoint = EndpointTest
		c.client.Cfg.MarketPayEndpoint = MarketpayEndpointTest
		c.client.Cfg.CheckoutEndpoint = CheckoutEndpointTest
		c.client.Cfg.TerminalApiCloudEndpoint = TerminalAPIEndpointTest
		c.client.Cfg.DisputesEndpoint = DisputesEndpointTest
		c.client.Cfg.BalancePlatformEndpoint = BalancePlatformEndpointTest
		c.client.Cfg.TransfersEndpoint = TransfersEndpointTest
		c.client.Cfg.ManagementEndpoint = ManagementEndpointTest
		c.client.Cfg.LegalEntityEndpoint = LegalEntityEntityTest
		c.client.Cfg.PosTerminalManagementEndpoint = PosTerminalManagementEndpointTest
		c.client.Cfg.DataProtectionEndpoint = DataProtectionEndpointTest
	}

	c.client.Cfg.CheckoutEndpoint += "/" + CheckoutAPIVersion
	c.client.Cfg.BalancePlatformEndpoint += "/" + BalancePlatformAPIVersion
	c.client.Cfg.ManagementEndpoint += "/" + ManagementAPIVersion
	c.client.Cfg.LegalEntityEndpoint += "/" + LegalEntityAPIVersion
}

// GetConfig Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetConfig() *common.Config {
	return c.client.Cfg
}
