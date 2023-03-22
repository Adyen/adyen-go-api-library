# RiskData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientData** | Pointer to **string** | Contains client-side data, like the device fingerprint, cookies, and specific browser settings. | [optional] 
**CustomFields** | Pointer to **map[string]string** | Any custom fields used as part of the input to configured risk rules. | [optional] 
**FraudOffset** | Pointer to **int32** | An integer value that is added to the normal fraud score. The value can be either positive or negative. | [optional] 
**ProfileReference** | Pointer to **string** | The risk profile to assign to this payment. When left empty, the merchant-level account&#39;s default risk profile will be applied. | [optional] 

## Methods

### NewRiskData

`func NewRiskData() *RiskData`

NewRiskData instantiates a new RiskData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskDataWithDefaults

`func NewRiskDataWithDefaults() *RiskData`

NewRiskDataWithDefaults instantiates a new RiskData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientData

`func (o *RiskData) GetClientData() string`

GetClientData returns the ClientData field if non-nil, zero value otherwise.

### GetClientDataOk

`func (o *RiskData) GetClientDataOk() (*string, bool)`

GetClientDataOk returns a tuple with the ClientData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientData

`func (o *RiskData) SetClientData(v string)`

SetClientData sets ClientData field to given value.

### HasClientData

`func (o *RiskData) HasClientData() bool`

HasClientData returns a boolean if a field has been set.

### GetCustomFields

`func (o *RiskData) GetCustomFields() map[string]string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *RiskData) GetCustomFieldsOk() (*map[string]string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *RiskData) SetCustomFields(v map[string]string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *RiskData) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetFraudOffset

`func (o *RiskData) GetFraudOffset() int32`

GetFraudOffset returns the FraudOffset field if non-nil, zero value otherwise.

### GetFraudOffsetOk

`func (o *RiskData) GetFraudOffsetOk() (*int32, bool)`

GetFraudOffsetOk returns a tuple with the FraudOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudOffset

`func (o *RiskData) SetFraudOffset(v int32)`

SetFraudOffset sets FraudOffset field to given value.

### HasFraudOffset

`func (o *RiskData) HasFraudOffset() bool`

HasFraudOffset returns a boolean if a field has been set.

### GetProfileReference

`func (o *RiskData) GetProfileReference() string`

GetProfileReference returns the ProfileReference field if non-nil, zero value otherwise.

### GetProfileReferenceOk

`func (o *RiskData) GetProfileReferenceOk() (*string, bool)`

GetProfileReferenceOk returns a tuple with the ProfileReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileReference

`func (o *RiskData) SetProfileReference(v string)`

SetProfileReference sets ProfileReference field to given value.

### HasProfileReference

`func (o *RiskData) HasProfileReference() bool`

HasProfileReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


