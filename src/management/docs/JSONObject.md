# JSONObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paths** | Pointer to [**[]JSONPath**](JSONPath.md) |  | [optional] 
**RootPath** | Pointer to [**JSONPath**](JSONPath.md) |  | [optional] 

## Methods

### NewJSONObject

`func NewJSONObject() *JSONObject`

NewJSONObject instantiates a new JSONObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSONObjectWithDefaults

`func NewJSONObjectWithDefaults() *JSONObject`

NewJSONObjectWithDefaults instantiates a new JSONObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaths

`func (o *JSONObject) GetPaths() []JSONPath`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *JSONObject) GetPathsOk() (*[]JSONPath, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *JSONObject) SetPaths(v []JSONPath)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *JSONObject) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetRootPath

`func (o *JSONObject) GetRootPath() JSONPath`

GetRootPath returns the RootPath field if non-nil, zero value otherwise.

### GetRootPathOk

`func (o *JSONObject) GetRootPathOk() (*JSONPath, bool)`

GetRootPathOk returns a tuple with the RootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPath

`func (o *JSONObject) SetRootPath(v JSONPath)`

SetRootPath sets RootPath field to given value.

### HasRootPath

`func (o *JSONObject) HasRootPath() bool`

HasRootPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


