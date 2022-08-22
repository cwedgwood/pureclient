# ProtectionGroupSnapshot

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A user-specified name. The name must be locally unique and can be changed. | [optional] [default to null]
**Created** | **int64** | The snapshot creation time of the original snapshot source. Measured in milliseconds since the UNIX epoch. | [optional] [default to null]
**Destroyed** | **bool** | Returns a value of &#x60;true&#x60; if the protection group snapshot has been destroyed and is pending eradication. The &#x60;time_remaining&#x60; value displays the amount of time left until the destroyed snapshot is permanently eradicated. Before the &#x60;time_remaining&#x60; period has elapsed, the destroyed snapshot can be recovered by setting &#x60;destroyed&#x3D;false&#x60;. Once the &#x60;time_remaining&#x60; period has elapsed, the snapshot is permanently eradicated and can no longer be recovered. | [optional] [default to null]
**Pod** | [***ProtectionGroupSnapshotPod**](ProtectionGroupSnapshot_pod.md) |  | [optional] [default to null]
**Source** | [***ProtectionGroupSnapshotSource**](ProtectionGroupSnapshot_source.md) |  | [optional] [default to null]
**Space** | [***ProtectionGroupSpace**](ProtectionGroup_space.md) |  | [optional] [default to null]
**Suffix** | **string** | The name suffix appended to the protection group name to make up the full protection group snapshot name in the form &#x60;PGROUP.SUFFIX&#x60;. If &#x60;suffix&#x60; is not specified, the protection group name is in the form &#x60;PGROUP.NNN&#x60;, where &#x60;NNN&#x60; is a unique monotonically increasing number. If multiple protection group snapshots are created at a time, the suffix name is appended to those snapshots. | [optional] [default to null]
**TimeRemaining** | **int64** | The amount of time left until the destroyed snapshot is permanently eradicated. Measured in milliseconds. Before the &#x60;time_remaining&#x60; period has elapsed, the destroyed snapshot can be recovered by setting &#x60;destroyed&#x3D;false&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

