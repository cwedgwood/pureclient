# ProtectionGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A user-specified name. The name must be locally unique and can be changed. | [optional] [default to null]
**Destroyed** | **bool** | Has this protection group been destroyed? To destroy a protection group, patch to &#x60;true&#x60;. To recover a destroyed protection group, patch to &#x60;false&#x60;. If not specified, defaults to &#x60;false&#x60;. | [optional] [default to null]
**HostCount** | **int64** | Number of hosts in this protection group. | [optional] [default to null]
**HostGroupCount** | **int64** | Number of host groups in this protection group. | [optional] [default to null]
**IsLocal** | **bool** | If set to &#x60;true&#x60;, the protection group belongs to the local array. If set to &#x60;false&#x60;, the protection group belongs to the remote array. | [optional] [default to null]
**Pod** | [***ProtectionGroupPod**](ProtectionGroup_pod.md) |  | [optional] [default to null]
**ReplicationSchedule** | [***Object**](.md) | The schedule settings for asynchronous replication. | [optional] [default to null]
**SnapshotSchedule** | [***ProtectionGroupSnapshotSchedule**](ProtectionGroup_snapshot_schedule.md) |  | [optional] [default to null]
**Source** | [***ProtectionGroupSource**](ProtectionGroup_source.md) |  | [optional] [default to null]
**SourceRetention** | [***ProtectionGroupSourceRetention**](ProtectionGroup_source_retention.md) |  | [optional] [default to null]
**Space** | [***ProtectionGroupSpace**](ProtectionGroup_space.md) |  | [optional] [default to null]
**TargetCount** | **int64** | The number of targets to where this protection group replicates. | [optional] [default to null]
**TargetRetention** | [***ProtectionGroupSourceRetention**](ProtectionGroup_source_retention.md) |  | [optional] [default to null]
**TimeRemaining** | **int64** | The amount of time left until the destroyed protection group is permanently eradicated. Measured in milliseconds. Before the &#x60;time_remaining&#x60; period has elapsed, the destroyed protection group can be recovered by setting &#x60;destroyed&#x3D;false&#x60;. | [optional] [default to null]
**VolumeCount** | **int64** | The number of volumes in the protection group. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

