# ModifyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** |  Pointer to [**map[string]interface{}**](.md) | This field contains additional data, which may be returned in a particular response. | [optional] 
**PspReference** | **string** | Adyen&#39;s 16-character string reference associated with the transaction. This value is globally unique; quote it when communicating with us about this response. | 
**Response** | **string** | The response: * In case of success, it is either &#x60;payout-confirm-received&#x60; or &#x60;payout-decline-received&#x60;. * In case of an error, an informational message is returned. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


