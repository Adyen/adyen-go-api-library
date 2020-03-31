# AdditionalData3DSecure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allow3DS2** | **string** | This parameter indicates that you are able to process 3D Secure 2 transactions natively on your payment page. Send this field when you are using &#x60;/payments&#x60; endpoint with any of our [native 3D Secure 2 solutions](https://docs.adyen.com/checkout/3d-secure/native-3ds2), such as Components or Drop-in. Possible values: * **true** - Ready to support native 3D Secure 2 authentication. Setting this to true does not mean always applying 3D Secure 2. Adyen still selects the version of 3D Secure based on configuration to optimize authorisation rates and improve the shopper&#39;s experience. * **false** – Not ready to support native 3D Secure 2 authentication. Adyen will not offer 3D Secure 2 to your shopper regardless of your configuration. &gt; This parameter only indicates your readiness to support 3D Secure 2 natively on Drop-in or Components. To specify that you want to perform 3D Secure on a transaction, use Dynamic 3D Secure or send the executeThreeD parameter. | [optional] 
**ExecuteThreeD** | **string** | This parameter indicates if you want to perform 3D Secure authentication on a transaction or not. Allowed values: * **true** – Perform 3D Secure authentication. * **false** – Don&#39;t perform 3D Secure authentication. &gt; Alternatively, you can also use Dynamic 3D Secure to configure rules for applying 3D Secure. | [optional] 
**MpiImplementationType** | **string** | In case of Secure+, this field must be set to **CUPSecurePlus**. | [optional] 
**ScaExemption** | **string** | Indicates the [exemption type](https://docs-admin.is.adyen.com/payments-fundamentals/psd2-sca-compliance-and-implementation-guide#specifypreferenceinyourapirequest) that you want to request for the transaction. Possible values: * **lowValue** * **secureCorporate** * **trustedBeneficiary** * **transactionRiskAnalysis** | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


