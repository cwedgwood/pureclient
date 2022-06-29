# RemoteProtectionGroupSnapshot

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A locally unique, system-generated name. The name cannot be modified. | [optional] [default to null]
**Created** | **int64** | Creation time of the snapshot on the original source of the snapshot. Measured in milliseconds since the UNIX epoch. | [optional] [default to null]
**Destroyed** | **bool** | Destroyed and pending eradication? If not specified, defaults to false. | [optional] [default to null]
**IsLocal** | **bool** | Whether or not this remote protection group snapshot is replicated from the current array. | [optional] [default to null]
**Remote** | [***RemoteProtectionGroupSnapshotRemote**](RemoteProtectionGroupSnapshot_remote.md) |  | [optional] [default to null]
**Source** | [***RemoteProtectionGroupSnapshotSource**](RemoteProtectionGroupSnapshot_source.md) |  | [optional] [default to null]
**Suffix** | **string** | The suffix that is appended to the &#x60;source_name&#x60; value to generate the full remote protection group snapshot name in the form &#x60;PGROUP.SUFFIX&#x60;. If the suffix is not specified, the system constructs the snapshot name in the form &#x60;PGROUP.NNN&#x60;, where &#x60;PGROUP&#x60; is the protection group name, and &#x60;NNN&#x60; is a monotonically increasing number. | [optional] [default to null]
**TimeRemaining** | **int64** | Milliseconds remaining until eradication, if the snapshot has been destroyed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

