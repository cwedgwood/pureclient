# Api28offloadsGooglecloud

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | **string** | The access key ID of the Google Cloud account used to create a connection between the array and a Google Cloud offload target. The access key ID is 24 characters in length and is only accepted when creating the connection between the array and the Google Cloud offload target. The &#x60;access_key_id&#x60;, &#x60;secret_access_key&#x60;, and &#x60;bucket&#x60; parameters must be set together. | [optional] [default to null]
**Bucket** | **string** | The name of the Google Cloud Storage bucket to which the data will be offloaded. Grant basic read and write access permissions to the bucket and verify that the bucket is empty of all objects. The &#x60;access_key_id&#x60;, &#x60;secret_access_key&#x60;, and &#x60;bucket&#x60; parameters must be set together. | [optional] [default to null]
**SecretAccessKey** | **string** | The secret access key that goes with the access key ID of the Google Cloud account. The secret access key is 40 characters in length is only accepted when creating the connection between the array and the Google Cloud offload target. The &#x60;access_key_id&#x60;, &#x60;secret_access_key&#x60;, and &#x60;bucket&#x60; parameters must be set together. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

