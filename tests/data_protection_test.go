package tests

import (
	"bytes"
	"context"
	"errors"
	"github.com/adyen/adyen-go-api-library/v8/src/adyen"
	"github.com/adyen/adyen-go-api-library/v8/src/common"
	"github.com/adyen/adyen-go-api-library/v8/src/dataprotection"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestDataProtection(t *testing.T) {
	client := adyen.NewClient(&common.Config{
		ApiKey:      "YOUR_ADYEN_API_KEY",
		Environment: "TEST",
		Debug:       "true" == os.Getenv("DEBUG"),
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/requestSubjectErasure", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		w.Header().Set("Content-Type", "application/json")
		body, _ := ioutil.ReadAll(r.Body)
		if bytes.Contains(body, []byte("invalidMerchant")) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			io.WriteString(w, `{
                "status": 422,
                "errorCode": "901",
                "message": "Invalid Merchant Account",
                "errorType": "validation"
            }`)
			return
		}
		io.WriteString(w, `{ "result": "SUCCESS" }`)
	})

	mockServer := httptest.NewServer(mux)
	defer mockServer.Close()
	client.DataProtection().BasePath = func() string { return mockServer.URL }
	service := client.DataProtection()

	t.Run("Configuration", func(t *testing.T) {
		testClient := adyen.NewClient(&common.Config{
			Environment: common.TestEnv,
		})
		assert.Equal(t, "https://ca-test.adyen.com/ca/services/DataProtectionService/v1", testClient.DataProtection().BasePath())
		liveClient := adyen.NewClient(&common.Config{
			Environment: common.LiveEnv,
		})
		assert.Equal(t, "https://ca-live.adyen.com/ca/services/DataProtectionService/v1", liveClient.DataProtection().BasePath())
	})

	t.Run("Submit a Subject Erasure Request", func(t *testing.T) {
		req := service.RequestSubjectErasureInput().SubjectErasureByPspReferenceRequest(dataprotection.SubjectErasureByPspReferenceRequest{
			ForceErasure:    common.PtrBool(true),
			MerchantAccount: common.PtrString("string"),
			PspReference:    common.PtrString("string"),
		})
		res, httpRes, err := service.RequestSubjectErasure(context.Background(), req)

		require.NotNil(t, res)
		assert.Equal(t, 200, httpRes.StatusCode)
		require.NoError(t, err)
		assert.Equal(t, "SUCCESS", res.GetResult())
	})

	t.Run("Invalid Merchant Account", func(t *testing.T) {
		req := service.RequestSubjectErasureInput().SubjectErasureByPspReferenceRequest(dataprotection.SubjectErasureByPspReferenceRequest{
			ForceErasure:    common.PtrBool(false),
			MerchantAccount: common.PtrString("invalidMerchant"),
			PspReference:    common.PtrString("string"),
		})
		res, httpRes, err := service.RequestSubjectErasure(context.Background(), req)

		assert.Nil(t, res.Result)
		assert.Equal(t, 422, httpRes.StatusCode)
		require.Error(t, err)
		var apiError common.APIError
		errors.As(err, &apiError)
		assert.Equal(t, float64(422), apiError.Status)
		assert.Equal(t, "901", apiError.Code)
		assert.Equal(t, "Invalid Merchant Account", apiError.Message)
		assert.Equal(t, "validation", apiError.Type)
	})
}
