//go:build integration
// +build integration

package legalentity

import (
	"context"
	"github.com/adyen/adyen-go-api-library/v8/src/adyen"
	"github.com/adyen/adyen-go-api-library/v8/src/common"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func Test_LegalEntity_Integration(t *testing.T) {
	godotenv.Load("./../../.env")

	var username, password, legalEntityId, proxy string

	username = os.Getenv("ADYEN_LEM_USERNAME")
	password = os.Getenv("ADYEN_LEM_PASSWORD")
	legalEntityId = os.Getenv("ADYEN_LEM_ID")
	proxy = os.Getenv("HTTP_PROXY")

	if username == "" || password == "" || legalEntityId == "" {
		t.Skip("Integration tests require credentials")
	}

	config := common.Config{
		Environment: common.TestEnv,
		Username:    os.Getenv("ADYEN_LEM_USERNAME"),
		Password:    os.Getenv("ADYEN_LEM_PASSWORD"),
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
	service := client.LegalEntity()

	t.Run("Get a legal entity", func(t *testing.T) {
		req := service.LegalEntitiesApi.GetLegalEntityInput(legalEntityId)

		res, httpRes, err := service.LegalEntitiesApi.GetLegalEntity(context.Background(), req)

		require.NoError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, legalEntityId, res.GetId())
	})
}
