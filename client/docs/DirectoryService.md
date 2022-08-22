# DirectoryService

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A locally unique, system-generated name. The name cannot be modified. | [optional] [default to null]
**BaseDn** | **string** | Base of the Distinguished Name (DN) of the directory service groups. | [optional] [default to null]
**BindPassword** | **string** | Masked password used to query the directory. | [optional] [default to null]
**BindUser** | **string** | Username used to query the directory. | [optional] [default to null]
**CaCertificate** | **string** | The certificate of the Certificate Authority (CA) that signed the certificates of the directory servers, which is used to validate the authenticity of the configured servers. | [optional] [default to null]
**CheckPeer** | **bool** | Whether or not server authenticity is enforced when a certificate is provided. | [optional] [default to null]
**Enabled** | **bool** | Whether or not the directory service is enabled. | [optional] [default to null]
**Services** | **[]string** | Services for which the directory service configuration is used. | [optional] [default to null]
**Uris** | **[]string** | List of URIs for the configured directory servers. | [optional] [default to null]
**Management** | [***DirectoryServiceManagement**](DirectoryService_management.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

