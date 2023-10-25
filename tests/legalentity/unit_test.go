package legalentity

import (
	"context"
	"github.com/adyen/adyen-go-api-library/v8/src/adyen"
	"github.com/adyen/adyen-go-api-library/v8/src/common"
	"github.com/adyen/adyen-go-api-library/v8/src/legalentity"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_LegalEntity(t *testing.T) {
	client := adyen.NewClient(&common.Config{
		ApiKey:      "YOUR_ADYEN_API_KEY",
		Environment: "TEST",
		Debug:       false,
	})
	service := client.LegalEntity()

	var (
		mockServer  *httptest.Server
		mockRequest *http.Request
	)

	mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mockRequest = r

		switch strings.TrimSpace(r.Method) {
		case "GET":
			switch strings.TrimSpace(r.URL.Path) {
			case "/legalEntities/123":
				legalEntityId := "123"
				model := legalentity.LegalEntity{Id: legalEntityId}
				mockResponse(http.StatusOK, w, model)
			case "/legalEntities/123/businessLines":
				model := legalentity.NewBusinessLinesWithDefaults()
				mockResponse(http.StatusOK, w, model)
			}
		case "PATCH":
			switch strings.TrimSpace(r.URL.Path) {
			case "/legalEntities/123":
				legalEntityId := "123"
				model := legalentity.LegalEntity{Id: legalEntityId}
				mockResponse(http.StatusOK, w, model)
			}
		case "POST":
			switch strings.TrimSpace(r.URL.Path) {
			case "/legalEntities":
				legalEntityId := "123"
				model := legalentity.LegalEntity{Id: legalEntityId}
				mockResponse(http.StatusOK, w, model)
			}
		}
	}))
	defer mockServer.Close()

	// base path is shared between all endpoints
	client.LegalEntity().BusinessLinesApi.BasePath = func() string {
		return mockServer.URL
	}

	t.Run("Test BusinessLinesApiService DeleteBusinessLine", func(t *testing.T) {
		req := service.BusinessLinesApi.DeleteBusinessLineInput("123")

		service.BusinessLinesApi.DeleteBusinessLine(context.Background(), req)

		assert.Equal(t, "DELETE", mockRequest.Method)
		assert.Equal(t, "/businessLines/123", mockRequest.URL.Path)
	})

	t.Run("Test BusinessLinesApiService GetBusinessLine", func(t *testing.T) {
		req := service.BusinessLinesApi.GetBusinessLineInput("123")

		service.BusinessLinesApi.GetBusinessLine(context.Background(), req)

		assert.Equal(t, "GET", mockRequest.Method)
		assert.Equal(t, "/businessLines/123", mockRequest.URL.Path)

	})

	t.Run("Test BusinessLinesApiService UpdateBusinessLine", func(t *testing.T) {
		req := service.BusinessLinesApi.UpdateBusinessLineInput("123")

		service.BusinessLinesApi.UpdateBusinessLine(context.Background(), req)

		assert.Equal(t, "PATCH", mockRequest.Method)
		assert.Equal(t, "/businessLines/123", mockRequest.URL.Path)
	})

	t.Run("Test BusinessLinesApiService CreateBusinessLine", func(t *testing.T) {
		req := service.BusinessLinesApi.CreateBusinessLineInput()

		service.BusinessLinesApi.CreateBusinessLine(context.Background(), req)

		assert.Equal(t, "POST", mockRequest.Method)
		assert.Equal(t, "/businessLines", mockRequest.URL.Path)
	})

	t.Run("Test DocumentsApiService UploadDocumentForVerificationChecks", func(t *testing.T) {
		req := service.DocumentsApi.UploadDocumentForVerificationChecksInput()

		service.DocumentsApi.UploadDocumentForVerificationChecks(context.Background(), req)

		assert.Equal(t, "POST", mockRequest.Method)
		assert.Equal(t, "/documents", mockRequest.URL.Path)

	})

	t.Run("Test HostedOnboardingApiService ListHostedOnboardingPageThemes", func(t *testing.T) {
		req := service.HostedOnboardingApi.ListHostedOnboardingPageThemesInput()

		service.HostedOnboardingApi.ListHostedOnboardingPageThemes(context.Background(), req)

		assert.Equal(t, "GET", mockRequest.Method)
		assert.Equal(t, "/themes", mockRequest.URL.Path)
	})

	t.Run("Test HostedOnboardingApiService GetLinkToAdyenhostedOnboardingPage", func(t *testing.T) {
		req := service.HostedOnboardingApi.GetLinkToAdyenhostedOnboardingPageInput("123")

		service.HostedOnboardingApi.GetLinkToAdyenhostedOnboardingPage(context.Background(), req)

		assert.Equal(t, "POST", mockRequest.Method)
		assert.Equal(t, "/legalEntities/123/onboardingLinks", mockRequest.URL.Path)
	})

	t.Run("Test LegalEntitiesApiService GetLegalEntity", func(t *testing.T) {
		req := service.LegalEntitiesApi.GetLegalEntityInput("123")

		service.LegalEntitiesApi.GetLegalEntity(context.Background(), req)

		assert.Equal(t, "GET", mockRequest.Method)
		assert.Equal(t, "/legalEntities/123", mockRequest.URL.Path)

	})

	t.Run("Test LegalEntitiesApiService GetAllBusinessLinesUnderLegalEntity", func(t *testing.T) {
		req := service.LegalEntitiesApi.GetAllBusinessLinesUnderLegalEntityInput("123")

		service.LegalEntitiesApi.GetAllBusinessLinesUnderLegalEntity(context.Background(), req)

		assert.Equal(t, "GET", mockRequest.Method)
		assert.Equal(t, "/legalEntities/123/businessLines", mockRequest.URL.Path)

	})

	t.Run("Test LegalEntitiesApiService UpdateLegalEntity", func(t *testing.T) {
		req := service.LegalEntitiesApi.UpdateLegalEntityInput("123").
			LegalEntityInfo(*legalentity.NewLegalEntityInfoWithDefaults())

		service.LegalEntitiesApi.UpdateLegalEntity(context.Background(), req)

		assert.Equal(t, "PATCH", mockRequest.Method)
		assert.Equal(t, "/legalEntities/123", mockRequest.URL.Path)

	})

	t.Run("Test TermsOfServiceApiService GetTermsOfServiceInformationForLegalEntity", func(t *testing.T) {
		req := service.TermsOfServiceApi.GetTermsOfServiceInformationForLegalEntityInput("123")

		service.TermsOfServiceApi.GetTermsOfServiceInformationForLegalEntity(context.Background(), req)

		assert.Equal(t, "GET", mockRequest.Method)
		assert.Equal(t, "/legalEntities/123/termsOfServiceAcceptanceInfos", mockRequest.URL.Path)
	})

	t.Run("Test TermsOfServiceApiService AcceptTermsOfService", func(t *testing.T) {
		legalEntityId := "123"
		termsOfServiceDocumentId := "456"
		req := service.TermsOfServiceApi.AcceptTermsOfServiceInput(legalEntityId, termsOfServiceDocumentId)

		service.TermsOfServiceApi.AcceptTermsOfService(context.Background(), req)

		assert.Equal(t, "PATCH", mockRequest.Method)
		assert.Equal(t, "/legalEntities/123/termsOfService/456", mockRequest.URL.Path)
	})
}
