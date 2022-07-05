/*
 * Adyen API
 *
 * Contact: support@adyen.com
 */

package common

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	_ioutil "io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/oauth2"
)

var (
	jsonCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
	xmlCheck  = regexp.MustCompile(`(?i:(?:application|text)/xml)`)
)

// Service type is the struct implemented by individual API services
type Service struct {
	Client   *Client
	BasePath func() string
}

// Client manages communication with the Adyen API
// In most cases there should be only one, shared, APIClient.
type Client struct {
	Cfg *Config
}

// CreateHTTPRequest is used as base to create HTTP request for all methods (GET, POST, PATCH...)
func CreateHTTPRequest(c *Client, httpMethod string, req interface{}, res interface{}, path string, ctxs []context.Context) (*http.Response, error) {
	// create path and map variables
	headerParams := make(map[string]string)
	queryParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// set Content-Type header
	contentType := SelectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"application/json"}

	// set Accept header
	headerAccept := SelectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
	}

	var ctx context.Context
	if len(ctxs) > 0 {
		ctx = ctxs[0]
	}

	r, err := c.PrepareRequest(ctx, path, httpMethod, req, headerParams, queryParams)
	if err != nil {
		return nil, err
	}

	httpResponse, err := c.CallAPI(r)

	if err != nil || httpResponse == nil {
		return httpResponse, err
	}

	body, err := _ioutil.ReadAll(httpResponse.Body)
	httpResponse.Body.Close()
	if err != nil {
		return httpResponse, err
	}

	if httpResponse.StatusCode >= 300 {
		newErr := NewAPIError(body, httpResponse.Status)
		return httpResponse, newErr
	}

	err = c.Decode(&res, body, httpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := NewAPIError(body, err.Error())
		return httpResponse, newErr
	}

	return httpResponse, nil
}

// MakeHTTPPostRequest is a generic method used to make HTTP POST requests
func (c *Client) MakeHTTPPostRequest(req interface{}, res interface{}, path string, ctxs ...context.Context) (*http.Response, error) {
	return CreateHTTPRequest(c, http.MethodPost, req, res, path, ctxs)
}

// MakeHTTPGetRequest is a generic method used to make HTTP GET requests
func (c *Client) MakeHTTPGetRequest(res interface{}, path string, ctxs ...context.Context) (*http.Response, error) {
	var req interface{}
	return CreateHTTPRequest(c, http.MethodGet, req, res, path, ctxs)
}

// MakeHTTPPatchRequest is a generic method used to make HTTP PATCH requests
func (c *Client) MakeHTTPPatchRequest(req interface{}, res interface{}, path string, ctxs ...context.Context) (*http.Response, error) {
	return CreateHTTPRequest(c, http.MethodPatch, req, res, path, ctxs)
}

// CallAPI do the Request.
func (c *Client) CallAPI(request *http.Request) (*http.Response, error) {
	if c.Cfg.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	resp, err := c.Cfg.HTTPClient.Do(request)
	if err != nil {
		return resp, err
	}

	if c.Cfg.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return resp, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	return resp, err
}

// PrepareRequest build the Request
func (c *Client) PrepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new Request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	// Add the user agent to the Request.
	localVarRequest.Header.Add("User-Agent", c.Cfg.UserAgent)
	localVarRequest.Header.Add("Cache-Control", "no-cache")

	// Add authentication headers
	if c.Cfg.ApiKey != "" {
		localVarRequest.Header.Add("x-API-key", c.Cfg.ApiKey)
	} else if c.Cfg.Username != "" && c.Cfg.Password != "" {
		localVarRequest.SetBasicAuth(c.Cfg.Username, c.Cfg.Password)
	}

	if ctx != nil {
		// add context to the Request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

		// OAuth2 authentication
		if tok, ok := ctx.Value(ContextOAuth2).(oauth2.TokenSource); ok {
			// We were able to grab an oauth2 token from the context
			var latestToken *oauth2.Token
			if latestToken, err = tok.Token(); err != nil {
				return nil, err
			}

			latestToken.SetAuthHeader(localVarRequest)
		}

		// Basic HTTP Authentication
		if auth, ok := ctx.Value(ContextBasicAuth).(BasicAuth); ok {
			localVarRequest.SetBasicAuth(auth.UserName, auth.Password)
		}

		// AccessToken Authentication
		if auth, ok := ctx.Value(ContextAccessToken).(string); ok {
			localVarRequest.Header.Add("Authorization", "Bearer "+auth)
		}

	}

	for header, value := range c.Cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}

	if key, ok := IdempotencyKey(ctx); ok {
		localVarRequest.Header.Add("Idempotency-Key", key)
	}

	return localVarRequest, nil
}

func (c *Client) Decode(v interface{}, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}
	if xmlCheck.MatchString(contentType) {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	if jsonCheck.MatchString(contentType) {
		if err = json.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	return errors.New("undefined Response type")
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// SelectHeaderContentType select a content type from the available list.
func SelectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// SelectHeaderAccept join all accept types and return
func SelectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insenstive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("Expected %s to be of type %s but received %s.", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	} else if t, ok := obj.(time.Time); ok {
		return t.Format(time.RFC3339)
	}

	return fmt.Sprintf("%v", obj)
}

// parameterToJson is helper for converting interface{} parameters to json strings
func parameterToJson(obj interface{}) (string, error) {
	jsonBuf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBuf), err
}

// Add a file to the multipart Request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// Set Request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		err = xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("Invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for Request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a Request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		} else {
			expires = now.Add(lifetime)
		}
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// APIError Provides access to the body, error and model on returned errors.
type APIError struct {
	RawBody []byte
	Err     string
	Status  float64
	Message string
	Code    string
	Type    string
}

// NewAPIError returns a new error instance
func NewAPIError(body []byte, errMsg string) APIError {
	errBody := map[string]interface{}{}
	apiError := APIError{
		Err:     errMsg,
		RawBody: body,
	}
	if len(body) != 0 {
		err := json.Unmarshal(body, &errBody)
		if err != nil {
			log.Printf("\nError while parsing error body for %s: %s\n", errMsg, err.Error())
		}
		if val, ok := errBody["message"]; ok {
			apiError.Message = val.(string)
		}
		if val, ok := errBody["errorCode"]; ok {
			apiError.Code = fmt.Sprintf("%v", val)
		}
		if val, ok := errBody["errorType"]; ok {
			apiError.Type = val.(string)
		}
		if val, ok := errBody["status"]; ok {
			apiError.Status = val.(float64)
		}
	}
	return apiError
}

// Error returns non-empty string if there was an error.
func (e APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("%s: %s (%s: %s)", e.Err, e.Message, e.Type, e.Code)
	}
	return e.Err
}

var ctxKeyIdempotencyKey = 1

/*
WithIdempotencyKey returns a context with an Idempotency-Key added to the provided context.
Pass this context as the first context to a call to Adyen, and the idempotency
key will be added to the header
*/
func WithIdempotencyKey(ctx context.Context, idempotencyKey string) context.Context {
	return context.WithValue(ctx, ctxKeyIdempotencyKey, idempotencyKey)
}

func IdempotencyKey(ctx context.Context) (string, bool) {
	if ctx == nil {
		return "", false
	}
	ikey := ctx.Value(ctxKeyIdempotencyKey)
	key, ok := ikey.(string)
	return key, ok
}
