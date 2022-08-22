# Model28ProtectiongroupsnapshotsBody1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A user-specified name. The name must be locally unique and can be changed. | [optional] [default to null]
**Created** | **int64** | The snapshot creation time of the original snapshot source. Measured in milliseconds since the UNIX epoch. | [optional] [default to null]
**Destroyed** | **bool** | Returns a value of &#x60;true&#x60; if the protection group snapshot has been destroyed and is pending eradication. The &#x60;time_remaining&#x60; value displays the amount of time left until the destroyed snapshot is permanently eradicated. Before the &#x60;time_remaining&#x60; period has elapsed, the destroyed snapshot can be recovered by setting &#x60;destroyed&#x3D;false&#x60;. Once the &#x60;time_remaining&#x60; period has elapsed, the snapshot is permanently eradicated and can no longer be recovered. | [optional] [default to null]
**Pod** | [***ProtectionGroupSnapshotPod**](ProtectionGroupSnapshot_pod.md) |  | [optional] [default to null]
**Source** | [***ProtectionGroupSnapshotSource**](ProtectionGroupSnapshot_source.md) |  | [optional] [default to null]
**Space** | [***ProtectionGroupSpace**](ProtectionGroup_space.md) |  | [optional] [default to null]
**Suffix** | **string** |  | [optional] [default to null]
**TimeRemaining** | **int64** | The amount of time left until the destroyed snapshot is permanently eradicated. Measured in milliseconds. Before the &#x60;time_remaining&#x60; period has elapsed, the destroyed snapshot can be recovered by setting &#x60;destroyed&#x3D;false&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

