//go:build integration
// +build integration

package legalentity

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/adyen/adyen-go-api-library/v21/src/adyen"
	"github.com/adyen/adyen-go-api-library/v21/src/common"
	"github.com/adyen/adyen-go-api-library/v21/src/legalentity"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_LegalEntity_Integration(t *testing.T) {
	godotenv.Load("./../../.env")

	var username, password, legalEntityId, proxy string

	username = os.Getenv("ADYEN_LEM_USERNAME")
	password = os.Getenv("ADYEN_LEM_PASSWORD")
	legalEntityId = os.Getenv("LEM_LEGAL_ENTITY_ID")
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

	t.Run("Get a legal entity that should fail", func(t *testing.T) {
		req := service.LegalEntitiesApi.GetLegalEntityInput("INVALID_ID")

		_, httpRes, err := service.LegalEntitiesApi.GetLegalEntity(context.Background(), req)

		require.NotNil(t, err)
		assert.Equal(t, 422, httpRes.StatusCode)

		// cast to API Error
		e := err.(common.APIError)

		require.NotNil(t, e.RawBody)
		assert.Equal(t, float64(422), e.Status)
		assert.Equal(t, "30_112", e.Code)

	})

	t.Run("Create a legal entity", func(t *testing.T) {

		legalEntityInfoRequiredType := legalentity.LegalEntityInfoRequiredType{
			Type: "individual",
			Individual: legalentity.NewIndividual(legalentity.Name{
				FirstName: "John",
				LastName:  "Smith",
			},
				legalentity.Address{
					Country: "NL",
				}),
		}

		req := service.LegalEntitiesApi.CreateLegalEntityInput().
			LegalEntityInfoRequiredType(legalEntityInfoRequiredType)

		res, httpRes, err := service.LegalEntitiesApi.CreateLegalEntity(context.Background(), req)

		require.NoError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.NotNil(t, res.GetId())

	})

	t.Run("Create a legal entity that fails", func(t *testing.T) {

		// log 4xx error
		service.LegalEntitiesApi.Client.Cfg.Log4XXError = true

		legalEntityInfoRequiredType := legalentity.LegalEntityInfoRequiredType{
			Type: "individual",
			Individual: legalentity.NewIndividual(legalentity.Name{
				FirstName: "John",
				LastName:  "Smith",
			},
				legalentity.Address{
					Country: "The Netherlands", // invalid: must pass country code
				}),
		}

		req := service.LegalEntitiesApi.CreateLegalEntityInput().
			LegalEntityInfoRequiredType(legalEntityInfoRequiredType)

		_, httpRes, err := service.LegalEntitiesApi.CreateLegalEntity(context.Background(), req)

		assert.Equal(t, 422, httpRes.StatusCode)

		// cast to API Error
		e := err.(common.APIError)

		require.NotNil(t, e.RawBody)
		assert.Equal(t, float64(422), e.Status)
		assert.Equal(t, "30_102", e.Code)

	})
}
