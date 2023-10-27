package tests

import (
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type MockHandlerFunc func(status int, method, endpoint, fixture string)

func MockResponse(t *testing.T, mux *http.ServeMux) MockHandlerFunc {
	return func(status int, method, endpoint, fixture string) {
		mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
			assert.Contains(t, strings.Fields(method), r.Method)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(status)

			file, err := os.Open(filepath.Join("fixtures", fixture))
			if err != nil {
				file, err = os.Open(filepath.Join("..", "fixtures", fixture))
			}
			if err != nil {
				log.Fatal(err)
			}
			if _, err := io.Copy(w, file); err != nil {
				log.Fatal(err)
			}
		})
	}
}
