# StoreDetailAndSubmitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** |  Pointer to [**map[string]interface{}**](.md) | This field contains additional data, which may be returned in a particular response. | [optional] 
**PspReference** | **string** | A new reference to uniquely identify this request. | 
**RefusalReason** | **string** | In case of refusal, an informational message for the reason. | [optional] 
**ResultCode** | **string** | The response:  * In case of success is payout-submit-received. * In case of an error, an informational message is returned. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


