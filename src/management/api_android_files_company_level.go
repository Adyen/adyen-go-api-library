/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// AndroidFilesCompanyLevelApi service
type AndroidFilesCompanyLevelApi common.Service

// All parameters accepted by AndroidFilesCompanyLevelApi.GetAndroidApp
type AndroidFilesCompanyLevelApiGetAndroidAppInput struct {
	companyId string
	id        string
}

/*
Prepare a request for GetAndroidApp
@param companyId The unique identifier of the company account.@param id The unique identifier of the app.
@return AndroidFilesCompanyLevelApiGetAndroidAppInput
*/
func (a *AndroidFilesCompanyLevelApi) GetAndroidAppInput(companyId string, id string) AndroidFilesCompanyLevelApiGetAndroidAppInput {
	return AndroidFilesCompanyLevelApiGetAndroidAppInput{
		companyId: companyId,
		id:        id,
	}
}

/*
GetAndroidApp Get Android app

Returns the details of the Android app identified in the path.
These apps have been uploaded to Adyen and can be installed or uninstalled on Android payment terminals through [terminal actions](https://docs.adyen.com/point-of-sale/automating-terminal-management/terminal-actions-api).

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Android files read
* Management API—Android files read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AndroidFilesCompanyLevelApiGetAndroidAppInput - Request parameters, see GetAndroidAppInput
@return AndroidApp, *http.Response, error
*/
func (a *AndroidFilesCompanyLevelApi) GetAndroidApp(ctx context.Context, r AndroidFilesCompanyLevelApiGetAndroidAppInput) (AndroidApp, *http.Response, error) {
	res := &AndroidApp{}
	path := "/companies/{companyId}/androidApps/{id}"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodGet,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError

	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by AndroidFilesCompanyLevelApi.ListAndroidApps
type AndroidFilesCompanyLevelApiListAndroidAppsInput struct {
	companyId   string
	pageNumber  *int32
	pageSize    *int32
	packageName *string
	versionCode *int32
}

// The number of the page to fetch.
func (r AndroidFilesCompanyLevelApiListAndroidAppsInput) PageNumber(pageNumber int32) AndroidFilesCompanyLevelApiListAndroidAppsInput {
	r.pageNumber = &pageNumber
	return r
}

// The number of items to have on a page, maximum 100. The default is 20 items on a page.
func (r AndroidFilesCompanyLevelApiListAndroidAppsInput) PageSize(pageSize int32) AndroidFilesCompanyLevelApiListAndroidAppsInput {
	r.pageSize = &pageSize
	return r
}

// The package name that uniquely identifies the Android app.
func (r AndroidFilesCompanyLevelApiListAndroidAppsInput) PackageName(packageName string) AndroidFilesCompanyLevelApiListAndroidAppsInput {
	r.packageName = &packageName
	return r
}

// The version number of the app.
func (r AndroidFilesCompanyLevelApiListAndroidAppsInput) VersionCode(versionCode int32) AndroidFilesCompanyLevelApiListAndroidAppsInput {
	r.versionCode = &versionCode
	return r
}

/*
Prepare a request for ListAndroidApps
@param companyId The unique identifier of the company account.
@return AndroidFilesCompanyLevelApiListAndroidAppsInput
*/
func (a *AndroidFilesCompanyLevelApi) ListAndroidAppsInput(companyId string) AndroidFilesCompanyLevelApiListAndroidAppsInput {
	return AndroidFilesCompanyLevelApiListAndroidAppsInput{
		companyId: companyId,
	}
}

/*
ListAndroidApps Get a list of Android apps

Returns a list of the Android apps that are available for the company identified in the path.
These apps have been uploaded to Adyen and can be installed or uninstalled on Android payment terminals through [terminal actions](https://docs.adyen.com/point-of-sale/automating-terminal-management/terminal-actions-api).

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Android files read
* Management API—Android files read and write
* Management API—Terminal actions read
* Management API—Terminal actions read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AndroidFilesCompanyLevelApiListAndroidAppsInput - Request parameters, see ListAndroidAppsInput
@return AndroidAppsResponse, *http.Response, error
*/
func (a *AndroidFilesCompanyLevelApi) ListAndroidApps(ctx context.Context, r AndroidFilesCompanyLevelApiListAndroidAppsInput) (AndroidAppsResponse, *http.Response, error) {
	res := &AndroidAppsResponse{}
	path := "/companies/{companyId}/androidApps"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.pageNumber != nil {
		common.ParameterAddToQuery(queryParams, "pageNumber", r.pageNumber, "")
	}
	if r.pageSize != nil {
		common.ParameterAddToQuery(queryParams, "pageSize", r.pageSize, "")
	}
	if r.packageName != nil {
		common.ParameterAddToQuery(queryParams, "packageName", r.packageName, "")
	}
	if r.versionCode != nil {
		common.ParameterAddToQuery(queryParams, "versionCode", r.versionCode, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodGet,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError

	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by AndroidFilesCompanyLevelApi.ListAndroidCertificates
type AndroidFilesCompanyLevelApiListAndroidCertificatesInput struct {
	companyId       string
	pageNumber      *int32
	pageSize        *int32
	certificateName *string
}

// The number of the page to fetch.
func (r AndroidFilesCompanyLevelApiListAndroidCertificatesInput) PageNumber(pageNumber int32) AndroidFilesCompanyLevelApiListAndroidCertificatesInput {
	r.pageNumber = &pageNumber
	return r
}

// The number of items to have on a page, maximum 100. The default is 20 items on a page.
func (r AndroidFilesCompanyLevelApiListAndroidCertificatesInput) PageSize(pageSize int32) AndroidFilesCompanyLevelApiListAndroidCertificatesInput {
	r.pageSize = &pageSize
	return r
}

// The name of the certificate.
func (r AndroidFilesCompanyLevelApiListAndroidCertificatesInput) CertificateName(certificateName string) AndroidFilesCompanyLevelApiListAndroidCertificatesInput {
	r.certificateName = &certificateName
	return r
}

/*
Prepare a request for ListAndroidCertificates
@param companyId The unique identifier of the company account.
@return AndroidFilesCompanyLevelApiListAndroidCertificatesInput
*/
func (a *AndroidFilesCompanyLevelApi) ListAndroidCertificatesInput(companyId string) AndroidFilesCompanyLevelApiListAndroidCertificatesInput {
	return AndroidFilesCompanyLevelApiListAndroidCertificatesInput{
		companyId: companyId,
	}
}

/*
ListAndroidCertificates Get a list of Android certificates

Returns a list of the Android certificates that are available for the company identified in the path.
Typically, these certificates enable running apps on Android payment terminals. The certifcates in the list have been uploaded to Adyen and can be installed or uninstalled on Android terminals through [terminal actions](https://docs.adyen.com/point-of-sale/automating-terminal-management/terminal-actions-api).

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Android files read
* Management API—Android files read and write
* Management API—Terminal actions read
* Management API—Terminal actions read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AndroidFilesCompanyLevelApiListAndroidCertificatesInput - Request parameters, see ListAndroidCertificatesInput
@return AndroidCertificatesResponse, *http.Response, error
*/
func (a *AndroidFilesCompanyLevelApi) ListAndroidCertificates(ctx context.Context, r AndroidFilesCompanyLevelApiListAndroidCertificatesInput) (AndroidCertificatesResponse, *http.Response, error) {
	res := &AndroidCertificatesResponse{}
	path := "/companies/{companyId}/androidCertificates"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.pageNumber != nil {
		common.ParameterAddToQuery(queryParams, "pageNumber", r.pageNumber, "")
	}
	if r.pageSize != nil {
		common.ParameterAddToQuery(queryParams, "pageSize", r.pageSize, "")
	}
	if r.certificateName != nil {
		common.ParameterAddToQuery(queryParams, "certificateName", r.certificateName, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodGet,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError

	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by AndroidFilesCompanyLevelApi.UploadAndroidApp
type AndroidFilesCompanyLevelApiUploadAndroidAppInput struct {
	companyId string
}

/*
Prepare a request for UploadAndroidApp
@param companyId The unique identifier of the company account.
@return AndroidFilesCompanyLevelApiUploadAndroidAppInput
*/
func (a *AndroidFilesCompanyLevelApi) UploadAndroidAppInput(companyId string) AndroidFilesCompanyLevelApiUploadAndroidAppInput {
	return AndroidFilesCompanyLevelApiUploadAndroidAppInput{
		companyId: companyId,
	}
}

/*
UploadAndroidApp Upload Android App

Uploads an Android APK file to Adyen.
The maximum APK file size is 200 MB.
To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Android files read and write

>By choosing to upload, install, or run any third-party applications on an Adyen payment terminal, you accept full responsibility and liability for any consequences of uploading, installing, or running any such applications.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AndroidFilesCompanyLevelApiUploadAndroidAppInput - Request parameters, see UploadAndroidAppInput
@return *http.Response, error
*/
func (a *AndroidFilesCompanyLevelApi) UploadAndroidApp(ctx context.Context, r AndroidFilesCompanyLevelApiUploadAndroidAppInput) (*http.Response, error) {
	var res interface{}
	path := "/companies/{companyId}/androidApps"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return httpRes, decodeError
		}
		return httpRes, serviceError
	}

	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return httpRes, decodeError
		}
		return httpRes, serviceError
	}

	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return httpRes, decodeError
		}
		return httpRes, serviceError
	}

	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return httpRes, decodeError
		}
		return httpRes, serviceError
	}

	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return httpRes, decodeError
		}
		return httpRes, serviceError
	}

	return httpRes, err
}