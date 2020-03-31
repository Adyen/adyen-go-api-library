# DetailsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | [**map[string]interface{}**](.md) | Use this collection to submit the details that were returned as a result of the &#x60;/payments&#x60; call. | 
**PaymentData** | **string** | The &#x60;paymentData&#x60; value that you received in the response to the &#x60;/payments&#x60; call. | [optional] 
**ThreeDSAuthenticationOnly** | **bool** | Change the &#x60;authenticationOnly&#x60; indicator originally set in the &#x60;/payments&#x60; request. Only needs to be set if you want to modify the value set previously. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


