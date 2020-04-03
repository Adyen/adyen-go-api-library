# ThreeDSecureData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationResponse** | **string** | In 3D Secure 1, the authentication response if the shopper was redirected.  In 3D Secure 2, this is the &#x60;transStatus&#x60; from the challenge result. If the transaction was frictionless, set this value to **Y**. | [optional] 
**Cavv** | **string** | The cardholder authentication value (base64 encoded, 20 bytes in a decoded form). | [optional] 
**CavvAlgorithm** | **string** | The CAVV algorithm used. Include this only for 3D Secure 1. | [optional] 
**DirectoryResponse** | **string** | In 3D Secure 1, this is the enrollment response from the 3D directory server.  In 3D Secure 2, this is the &#x60;transStatus&#x60; from the &#x60;ARes&#x60;. The possible values are **A** or **Y** for a frictionless flow, or **C** for a challenge flow. | [optional] 
**DsTransID** | **string** | Supported for 3D Secure 2. The unique transaction identifier assigned by the Directory Server (DS) to identify a single transaction. | [optional] 
**Eci** | **string** | The electronic commerce indicator. | [optional] 
**ThreeDSVersion** | **string** | The version of the 3D Secure protocol. | [optional] 
**Xid** | **string** | Supported for 3D Secure 1. The transaction identifier (Base64-encoded, 20 bytes in a decoded form). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


