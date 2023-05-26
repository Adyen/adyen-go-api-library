# AssignTerminalsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | **map[string]string** | Array that returns a list of the terminals, and for each terminal the result of assigning it to an account or store.  The results can be:    - &#x60;Done&#x60;: The terminal has been assigned.   - &#x60;AssignmentScheduled&#x60;: The terminal will be assigned asynschronously.   - &#x60;RemoveConfigScheduled&#x60;: The terminal was previously assigned and boarded. Wait for the terminal to synchronize with the Adyen platform. For more information, refer to [Reassigning boarded terminals](https://docs.adyen.com/point-of-sale/managing-terminals/assign-terminals#reassign-boarded-terminals).   - &#x60;Error&#x60;: There was an error when assigning the terminal.  | 

## Methods

### NewAssignTerminalsResponse

`func NewAssignTerminalsResponse(results map[string]string, ) *AssignTerminalsResponse`

NewAssignTerminalsResponse instantiates a new AssignTerminalsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignTerminalsResponseWithDefaults

`func NewAssignTerminalsResponseWithDefaults() *AssignTerminalsResponse`

NewAssignTerminalsResponseWithDefaults instantiates a new AssignTerminalsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *AssignTerminalsResponse) GetResults() map[string]string`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AssignTerminalsResponse) GetResultsOk() (*map[string]string, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AssignTerminalsResponse) SetResults(v map[string]string)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


