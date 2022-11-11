# \TerminalActionsTerminalLevelApi

All URIs are relative to *https://management-test.adyen.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostTerminalsScheduleActions**](TerminalActionsTerminalLevelApi.md#PostTerminalsScheduleActions) | **Post** /terminals/scheduleActions | Create a terminal action



## PostTerminalsScheduleActions

> ScheduleTerminalActionsResponse PostTerminalsScheduleActions(ctx).ScheduleTerminalActionsRequest(scheduleTerminalActionsRequest).Execute()

Create a terminal action



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    scheduleTerminalActionsRequest := *openapiclient.NewScheduleTerminalActionsRequest() // ScheduleTerminalActionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerminalActionsTerminalLevelApi.PostTerminalsScheduleActions(context.Background()).ScheduleTerminalActionsRequest(scheduleTerminalActionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerminalActionsTerminalLevelApi.PostTerminalsScheduleActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTerminalsScheduleActions`: ScheduleTerminalActionsResponse
    fmt.Fprintf(os.Stdout, "Response from `TerminalActionsTerminalLevelApi.PostTerminalsScheduleActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTerminalsScheduleActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scheduleTerminalActionsRequest** | [**ScheduleTerminalActionsRequest**](ScheduleTerminalActionsRequest.md) |  | 

### Return type

[**ScheduleTerminalActionsResponse**](ScheduleTerminalActionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

