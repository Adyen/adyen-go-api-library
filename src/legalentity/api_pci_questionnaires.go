/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"context"
	_context "context"
	_nethttp "net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// PCIQuestionnairesApi PCIQuestionnairesApi service
type PCIQuestionnairesApi common.Service

type PCIQuestionnairesApiGeneratePciQuestionnaireConfig struct {
	ctx                           context.Context
	id                            string
	generatePciDescriptionRequest *GeneratePciDescriptionRequest
}

func (r PCIQuestionnairesApiGeneratePciQuestionnaireConfig) GeneratePciDescriptionRequest(generatePciDescriptionRequest GeneratePciDescriptionRequest) PCIQuestionnairesApiGeneratePciQuestionnaireConfig {
	r.generatePciDescriptionRequest = &generatePciDescriptionRequest
	return r
}

/*
GeneratePciQuestionnaire Generate PCI questionnaire

Generates the required PCI questionnaire based on the user's [salesChannel](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/businessLines__reqParam_salesChannels). If multiple questionnaires are required, this request creates a single consodilated document to be signed.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The legal entity ID of the individual who will sign the PCI questionnaire.
 @return PCIQuestionnairesApiGeneratePciQuestionnaireConfig
*/
func (a *PCIQuestionnairesApi) GeneratePciQuestionnaireConfig(ctx context.Context, id string) PCIQuestionnairesApiGeneratePciQuestionnaireConfig {
	return PCIQuestionnairesApiGeneratePciQuestionnaireConfig{
		ctx: ctx,
		id:  id,
	}
}

/*
Generate PCI questionnaire
Generates the required PCI questionnaire based on the user&#39;s [salesChannel](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/businessLines__reqParam_salesChannels). If multiple questionnaires are required, this request creates a single consodilated document to be signed.
 * @param id The legal entity ID of the individual who will sign the PCI questionnaire.
 * @param req GeneratePciDescriptionRequest - reference of GeneratePciDescriptionRequest).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return GeneratePciDescriptionResponse
*/

func (a *PCIQuestionnairesApi) GeneratePciQuestionnaire(r PCIQuestionnairesApiGeneratePciQuestionnaireConfig) (GeneratePciDescriptionResponse, *_nethttp.Response, error) {
	res := &GeneratePciDescriptionResponse{}
	path := "/legalEntities/{id}/pciQuestionnaires/generatePciTemplates"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodPost, r.generatePciDescriptionRequest, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type PCIQuestionnairesApiGetPciQuestionnaireConfig struct {
	ctx   context.Context
	id    string
	pciid string
}

/*
GetPciQuestionnaire Get PCI questionnaire

Returns the signed PCI questionnaire.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The legal entity ID of the individual who signed the PCI questionnaire.
 @param pciid The unique identifier of the signed PCI questionnaire.
 @return PCIQuestionnairesApiGetPciQuestionnaireConfig
*/
func (a *PCIQuestionnairesApi) GetPciQuestionnaireConfig(ctx context.Context, id string, pciid string) PCIQuestionnairesApiGetPciQuestionnaireConfig {
	return PCIQuestionnairesApiGetPciQuestionnaireConfig{
		ctx:   ctx,
		id:    id,
		pciid: pciid,
	}
}

/*
Get PCI questionnaire
Returns the signed PCI questionnaire.
 * @param id The legal entity ID of the individual who signed the PCI questionnaire.
 * @param pciid The unique identifier of the signed PCI questionnaire.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return GetPciQuestionnaireResponse
*/

func (a *PCIQuestionnairesApi) GetPciQuestionnaire(r PCIQuestionnairesApiGetPciQuestionnaireConfig) (GetPciQuestionnaireResponse, *_nethttp.Response, error) {
	res := &GetPciQuestionnaireResponse{}
	path := "/legalEntities/{id}/pciQuestionnaires/{pciid}"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	path = strings.Replace(path, "{"+"pciid"+"}", url.PathEscape(common.ParameterValueToString(r.pciid, "pciid")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodGet, nil, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type PCIQuestionnairesApiGetPciQuestionnaireDetailsConfig struct {
	ctx context.Context
	id  string
}

/*
GetPciQuestionnaireDetails Get PCI questionnaire details

Get a list of signed PCI questionnaires.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier of the legal entity to get PCI questionnaire information.
 @return PCIQuestionnairesApiGetPciQuestionnaireDetailsConfig
*/
func (a *PCIQuestionnairesApi) GetPciQuestionnaireDetailsConfig(ctx context.Context, id string) PCIQuestionnairesApiGetPciQuestionnaireDetailsConfig {
	return PCIQuestionnairesApiGetPciQuestionnaireDetailsConfig{
		ctx: ctx,
		id:  id,
	}
}

/*
Get PCI questionnaire details
Get a list of signed PCI questionnaires.
 * @param id The unique identifier of the legal entity to get PCI questionnaire information.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return GetPciQuestionnaireInfosResponse
*/

func (a *PCIQuestionnairesApi) GetPciQuestionnaireDetails(r PCIQuestionnairesApiGetPciQuestionnaireDetailsConfig) (GetPciQuestionnaireInfosResponse, *_nethttp.Response, error) {
	res := &GetPciQuestionnaireInfosResponse{}
	path := "/legalEntities/{id}/pciQuestionnaires"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodGet, nil, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type PCIQuestionnairesApiSignPciQuestionnaireConfig struct {
	ctx               context.Context
	id                string
	pciSigningRequest *PciSigningRequest
}

func (r PCIQuestionnairesApiSignPciQuestionnaireConfig) PciSigningRequest(pciSigningRequest PciSigningRequest) PCIQuestionnairesApiSignPciQuestionnaireConfig {
	r.pciSigningRequest = &pciSigningRequest
	return r
}

/*
SignPciQuestionnaire Sign PCI questionnaire

Signs the required PCI questionnaire.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The legal entity ID of the individual who signed the PCI questionnaire.
 @return PCIQuestionnairesApiSignPciQuestionnaireConfig
*/
func (a *PCIQuestionnairesApi) SignPciQuestionnaireConfig(ctx context.Context, id string) PCIQuestionnairesApiSignPciQuestionnaireConfig {
	return PCIQuestionnairesApiSignPciQuestionnaireConfig{
		ctx: ctx,
		id:  id,
	}
}

/*
Sign PCI questionnaire
Signs the required PCI questionnaire.
 * @param id The legal entity ID of the individual who signed the PCI questionnaire.
 * @param req PciSigningRequest - reference of PciSigningRequest).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return PciSigningResponse
*/

func (a *PCIQuestionnairesApi) SignPciQuestionnaire(r PCIQuestionnairesApiSignPciQuestionnaireConfig) (PciSigningResponse, *_nethttp.Response, error) {
	res := &PciSigningResponse{}
	path := "/legalEntities/{id}/pciQuestionnaires/signPciTemplates"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodPost, r.pciSigningRequest, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}