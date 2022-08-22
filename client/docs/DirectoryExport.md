# DirectoryExport

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destroyed** | **bool** | Returns a value of &#x60;true&#x60; if the managed directory of the export has been destroyed and is pending eradication. The export can be recovered by recovering the destroyed managed directory. | [optional] [default to null]
**Directory** | [***AllOfDirectoryExportDirectory**](AllOfDirectoryExportDirectory.md) | The managed directory of the export. | [optional] [default to null]
**Enabled** | **bool** | Returns a value of &#x60;true&#x60; if the export policy that manages this export is enabled. | [optional] [default to null]
**ExportName** | **string** | The export name for accessing this export. | [optional] [default to null]
**Path** | **string** | The path of the exported managed directory. | [optional] [default to null]
**Policy** | [***AllOfDirectoryExportPolicy**](AllOfDirectoryExportPolicy.md) | The export policy that manages this export. An export can be managed by at most one export policy. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

