package common

import (
	"fmt"
	"net/http"
	"time"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes an oauth2.TokenSource as authentication for the Request.
	ContextOAuth2 = contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the Request.
	ContextBasicAuth = contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the Request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextAPIKey takes an APIKey as authentication for the Request
	ContextAPIKey = contextKey("apikey")
)

// BasicAuth provides basic http authentication to a Request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a Request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

type Environment string

const (
	TestEnv Environment = "TEST"
	LiveEnv Environment = "LIVE"
)

const (
	LibName    = "adyen-go-api-library"
	LibVersion = "8.0.0"
)

// Config stores the configuration of the API client
type Config struct {
	Username                      string        `json:"username,omitempty"`
	Password                      string        `json:"password,omitempty"`
	MerchantAccount               string        `json:"merchantAccount,omitempty"`
	Environment                   Environment   `json:"environment,omitempty"`
	Endpoint                      string        `json:"endpoint,omitempty"`
	MarketPayEndpoint             string        `json:"marketPayEndpoint,omitempty"`
	ApiKey                        string        `json:"apiKey,omitempty"`
	ConnectionTimeoutMillis       time.Duration `json:"connectionTimeoutMillis,omitempty"`
	CertificatePath               string        `json:"certificatePath,omitempty"`
	DisputesEndpoint              string        `json:"disputesEndpoint,omitempty"`
	BalancePlatformEndpoint       string        `json:"balancePlatformEndpoint,omitempty"`
	ManagementEndpoint            string        `json:"managementEndpoint,omitempty"`
	LegalEntityEndpoint           string        `json:"legalEntityEndpoint,omitempty"`
	TransfersEndpoint             string        `json:"transfersEndpoint,omitempty"`
	PosTerminalManagementEndpoint string        `json:"posTerminalManagementEndpoint,omitempty"`
	DataProtectionEndpoint        string        `json:"dataProtectionEndpoint,omitempty"`

	//Checkout Specific
	CheckoutEndpoint string `json:"checkoutEndpoint,omitempty"`

	//Terminal API Specific
	TerminalApiCloudEndpoint string `json:"terminalApiCloudEndpoint,omitempty"`
	TerminalApiLocalEndpoint string `json:"terminalApiLocalEndpoint,omitempty"`
	TerminalCertificatePath  string `json:"terminalCertificatePath,omitempty"`

	LiveEndpointURLPrefix string            `json:"liveEndpointURLPrefix,omitempty"`
	DefaultHeader         map[string]string `json:"defaultHeader,omitempty"`
	Debug                 bool              `json:"debug,omitempty"`
	UserAgent             string            `json:"userAgent,omitempty"`
	HTTPClient            *http.Client
}

func (c *Config) GetCheckoutEndpoint() (string, error) {
	if c.CheckoutEndpoint == "" {
		message := "please provide your unique live url prefix on the SetEnvironment() call on the APIClient or provide checkoutEndpoint in your config object"
		return "", fmt.Errorf(message)
	}
	return c.CheckoutEndpoint, nil
}
