# RemoteVolumeSnapshot

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource. | [optional] [default to null]
**Name** | **string** | A user-specified name. The name must be locally unique and can be changed. | [optional] [default to null]
**Created** | **int64** | The snapshot creation time. Measured in milliseconds since the UNIX epoch. | [optional] [default to null]
**Destroyed** | **bool** | Returns a value of &#x60;true&#x60; if the snapshot has been destroyed and is pending eradication. The &#x60;time_remaining&#x60; value displays the amount of time left until the destroyed volume snapshot is permanently eradicated. Before the &#x60;time_remaining&#x60; period has elapsed, the destroyed volume snapshot can be recovered by setting &#x60;destroyed&#x3D;false&#x60;. Once the &#x60;time_remaining&#x60; period has elapsed, the volume snapshot is permanently eradicated and can no longer be recovered. | [optional] [default to null]
**Pod** | [***RemoteVolumeSnapshotPod**](RemoteVolumeSnapshot_pod.md) |  | [optional] [default to null]
**Provisioned** | **int64** | The provisioned space of the snapshot. Measured in bytes. | [optional] [default to null]
**Source** | [***RemoteVolumeSnapshotSource**](RemoteVolumeSnapshot_source.md) |  | [optional] [default to null]
**Suffix** | **string** | The suffix that is appended to the &#x60;source_name&#x60; value to generate the full volume snapshot name in the form &#x60;VOL.SUFFIX&#x60;. If the suffix is not specified, the system constructs the snapshot name in the form &#x60;VOL.NNN&#x60;, where &#x60;VOL&#x60; is the volume name, and &#x60;NNN&#x60; is a monotonically increasing number. | [optional] [default to null]
**TimeRemaining** | **int64** | The amount of time left until the destroyed snapshot is permanently eradicated. Measured in milliseconds. Before the &#x60;time_remaining&#x60; period has elapsed, the destroyed snapshot can be recovered by setting &#x60;destroyed&#x3D;false&#x60;. | [optional] [default to null]
**Remote** | [***RemoteVolumeSnapshotRemote**](RemoteVolumeSnapshot_remote.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

