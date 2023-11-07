package checkout

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/adyen/adyen-go-api-library/v8/src/adyen"
	"github.com/adyen/adyen-go-api-library/v8/src/checkout"
	"github.com/adyen/adyen-go-api-library/v8/src/common"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var idempotencyKeys *sync.Map

type referenceAndIdempotencyKeyTransport struct {
	errChan chan error
}

func (rl *referenceAndIdempotencyKeyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	body, _ := ioutil.ReadAll(req.Body)
	data := map[string]interface{}{}
	json.Unmarshal(body, &data)
	reference := data["reference"]
	idempotencyKey := req.Header.Get("Idempotency-Key")
	idempotencyKeys.Store(reference.(string), idempotencyKey)
	resp := httptest.NewRecorder()
	resp.WriteHeader(http.StatusOK)
	return resp.Result(), nil
}

func Test_Checkout_Idempotency_Race(t *testing.T) {
	idempotencyKeys = &sync.Map{}
	client := adyen.NewClient(&common.Config{
		HTTPClient: &http.Client{
			Transport: &referenceAndIdempotencyKeyTransport{},
		},
	})
	service := client.Checkout()

	for r := 0; r < 10; r++ {
		t.Run(fmt.Sprintf("Routine %d", r), func(t *testing.T) {
			ir := r
			t.Parallel()
			for i := 0; i < 100; i++ {
				idempotencyKey := uuid.Must(uuid.NewRandom()).String()
				ref := fmt.Sprintf("%d/%d", ir, i)
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
				defer cancel()
				body := checkout.PaymentRequest{
					Reference: ref,
					PaymentMethod: checkout.IdealDetailsAsCheckoutPaymentMethod(&checkout.IdealDetails{
						Issuer: "1121",
					}),
				}
				req := service.PaymentsApi.PaymentsInput().IdempotencyKey(idempotencyKey).PaymentRequest(body)
				_, _, err := service.PaymentsApi.Payments(ctx, req)
				require.NoError(t, err)
				v, ok := idempotencyKeys.Load(ref)
				require.True(t, ok)
				require.Equal(t, v.(string), idempotencyKey)
			}
		})
	}
}
