# Api28offloadsAzure

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerName** | **string** | The name of the container in the Azure Blob storage account to where the data will be offloaded. The name must be a valid DNS name. If not specified, defaults to &#x60;offload&#x60;. | [optional] [default to null]
**AccountName** | **string** | The name of the existing Azure Blob storage account. | [optional] [default to null]
**SecretAccessKey** | **string** | The secret access key that goes with the account name (&#x60;account_name&#x60;) of the Azure Blob storage account. The secret access key is only accepted when creating the connection between the array and the Azure Blob storage account. The &#x60;account_name&#x60; and &#x60;container_name&#x60;, and &#x60;secret_access_key&#x60; parameters must be set together. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

