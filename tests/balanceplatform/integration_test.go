//go:build integration
// +build integration

package balanceplatform

import (
    "context"
    "github.com/adyen/adyen-go-api-library/v21/src/adyen"
    "github.com/adyen/adyen-go-api-library/v21/src/common"
    "github.com/joho/godotenv"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "net/http"
    "net/url"
    "os"
    "testing"
)

func Test_BalancePlatform_Integration(t *testing.T) {
	godotenv.Load("./../../.env")

	var apiKey, balancePlatformId, proxy string

    apiKey = os.Getenv("ADYEN_PLATFORM_API_KEY")
    balancePlatformId = os.Getenv("BALANCE_PLATFORM_ID")
	proxy = os.Getenv("HTTP_PROXY")

	if apiKey == "" || balancePlatformId == "" {
		t.Skip("Integration tests require credentials")
	}

	config := common.Config{
		Environment: common.TestEnv,
		ApiKey:    apiKey,
		Debug:       "true" == os.Getenv("DEBUG"),
	}

	if proxy != "" {
		proxyURL, _ := url.Parse(proxy)
		config.HTTPClient = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			},
		}
	}
	client := adyen.NewClient(&config)
	service := client.BalancePlatform()

	t.Run("Get a Balance Platform", func(t *testing.T) {
		req := service.PlatformApi.GetBalancePlatformInput(balancePlatformId)
		res, httpRes, err := service.PlatformApi.GetBalancePlatform(context.Background(), req)

		require.NoError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, balancePlatformId, res.GetId())
	})

    t.Run("Get a Balance Platform that fails with 422", func(t *testing.T) {
        req := service.PlatformApi.GetBalancePlatformInput("na")
        _, httpRes, err := service.PlatformApi.GetBalancePlatform(context.Background(), req)

        require.Error(t, err)
        assert.Equal(t, 422, httpRes.StatusCode)

        // cast to RestServiceError
        restServiceError := err.(common.RestServiceError)
        assert.Equal(t, "https://docs.adyen.com/errors/not-found", restServiceError.Type)
        assert.Equal(t, "30_002", restServiceError.ErrorCode)
    })

    t.Run("Get a Balance Platform that fails with 400", func(t *testing.T) {

        // log 4xx error
        service.PlatformApi.Client.Cfg.Log4XXError = true

        req := service.PlatformApi.GetBalancePlatformInput("with invalid char//")
        _, httpRes, err := service.PlatformApi.GetBalancePlatform(context.Background(), req)

        require.Error(t, err)
        assert.Equal(t, 400, httpRes.StatusCode)
    })
}
