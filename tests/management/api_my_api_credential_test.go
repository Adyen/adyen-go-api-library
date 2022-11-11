/*
 * Testing MyAPICredentialApiService
 *
 * Contact: support@adyen.com
 */

package management

import (
    "context"
    "encoding/json"
    "fmt"
    Management "github.com/adyen/adyen-go-api-library/v6/src/management"
    "github.com/joho/godotenv"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "net/http"
    "net/http/httptest"
    "os"
    "strings"
    "testing"
)

func Test_ManagementAPI_MyAPICredentialApiService(t *testing.T) {
    godotenv.Load("./../../.env")

    var (
        APIKey = "n/a"
        env    = Management.TestEnv
        server *httptest.Server
    )

    server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // mock here
        switch strings.TrimSpace(r.URL.Path) {
        case "/me":
            model := Management.MeApiCredential{Id: "me", Username: "myUsername"}
            mockResponse(http.StatusOK, w, model)
        case "/me/allowedOrigins":
            model := Management.AllowedOriginsResponse{Data: []Management.AllowedOrigin{{Domain: "adyen.com"}}}
            mockResponse(http.StatusOK, w, model)
        default:
            t.Errorf("Mock not found")
            http.NotFoundHandler().ServeHTTP(w, r)
        }

    }))

    configuration, err := Management.NewManagementAPIConfiguration(APIKey, env)
    configuration.Servers = Management.ServerConfigurations{
        {
            URL:         server.URL,
            Description: "Mock Server",
        },
    }
    require.Nil(t, err, "Error creating Config object")
    apiClient := Management.NewAPIClient(configuration)
    require.NotNil(t, apiClient)

    t.Run("Test MyAPICredentialApiService GetMe", func(t *testing.T) {

        t.Run("Create an API request that should pass", func(t *testing.T) {

            resp, httpRes, err := apiClient.MyAPICredentialApi.GetMe(context.Background()).Execute()
            if err != nil {
                t.Errorf("Error when calling `MyAPICredentialApi.GetMe``: %v\n", err)
                t.Errorf("Full HTTP response: %v\n", resp)
            }

            require.Nil(t, err)
            assert.Equal(t, 200, httpRes.StatusCode)
            require.NotNil(t, resp)
            assert.Equal(t, "me", resp.Id)

        })

    })

    t.Run("Test MyAPICredentialApiService GetMeAllowedOrigins", func(t *testing.T) {

        t.Run("Create an API request that should pass", func(t *testing.T) {

            resp, httpRes, err := apiClient.MyAPICredentialApi.GetMeAllowedOrigins(context.Background()).Execute()
            if err != nil {
                fmt.Fprintf(os.Stderr, "Error when calling `MyAPICredentialApi.GetMeAllowedOrigins``: %v\n", err)
                fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
            }

            require.Nil(t, err)
            assert.Equal(t, 200, httpRes.StatusCode)
            require.NotNil(t, resp)
            assert.Equal(t, 1, len(resp.Data))
        })
    })

    t.Run("Test MyAPICredentialApiService PostMeAllowedOrigins", func(t *testing.T) {

        t.Run("Create an API request that should pass", func(t *testing.T) {

            createAllowedOriginRequest := Management.NewCreateAllowedOriginRequest("https://adyen.com")

            resp, httpRes, err := apiClient.MyAPICredentialApi.PostMeAllowedOrigins(context.Background()).CreateAllowedOriginRequest(*createAllowedOriginRequest).Execute()

            if err != nil {
                fmt.Fprintf(os.Stderr, "Error when calling `PostMeAllowedOrigins.GetMeAllowedOrigins``: %v\n", err)
                fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
            }

            require.Nil(t, err)
            assert.Equal(t, 200, httpRes.StatusCode)
            require.NotNil(t, resp)
            assert.Equal(t, 1, len(resp.Data))
        })
    })

}

// mock Response given the model
func mockResponse(status int, w http.ResponseWriter, model interface{}) {

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(model)
}
