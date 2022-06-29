# Model28RemoteprotectiongroupsnapshotsBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A locally unique, system-generated name. The name cannot be modified. | [optional] [default to null]
**Created** | **int64** | Creation time of the snapshot on the original source of the snapshot. Measured in milliseconds since the UNIX epoch. | [optional] [default to null]
**Destroyed** | **bool** |  | [optional] [default to null]
**IsLocal** | **bool** | Whether or not this remote protection group snapshot is replicated from the current array. | [optional] [default to null]
**Remote** | [***RemoteProtectionGroupSnapshotRemote**](RemoteProtectionGroupSnapshot_remote.md) |  | [optional] [default to null]
**Source** | [***RemoteProtectionGroupSnapshotSource**](RemoteProtectionGroupSnapshot_source.md) |  | [optional] [default to null]
**Suffix** | **string** | Specifies a name suffix for the snapshots created. The snapshot is created on the FlashArray specified by the &#x60;on&#x60; parameter. The &#x60;on&#x60; parameter cannot refer to an offload target. Snapshots with suffixes specified have names in the form &#x60;PGROUP.SUFFIX&#x60; instead of the default &#x60;PGROUP.NNN&#x60; form. The names of all snapshots created by a single command that specifies this option have the same suffix. | [optional] [default to null]
**TimeRemaining** | **int64** | Milliseconds remaining until eradication, if the snapshot has been destroyed. | [optional] [default to null]
**Id** | **string** | A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

