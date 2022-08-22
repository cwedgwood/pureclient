# RemoteProtectionGroupSnapshotTransfer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A user-specified name. The name must be locally unique and can be changed. | [optional] [default to null]
**Destroyed** | **bool** | Returns a value of &#x60;true&#x60; if the snapshot has been destroyed and is pending eradication. The destroyed snapshot can be recovered by setting &#x60;destroyed&#x3D;false&#x60;. Once the eradication pending period has elapsed, the snapshot is permanently eradicated and can no longer be recovered. | [optional] [default to null]
**Started** | **int64** | The timestamp of when the snapshot replication process started. Measured in milliseconds since the UNIX epoch. | [optional] [default to null]
**Progress** | **float32** | The percentage progress of the snapshot transfer from the source array to the target. Displayed in decimal format. | [optional] [default to null]
**Completed** | **int64** | The timestamp of when the snapshot replication process completed. Measured in milliseconds since the UNIX epoch. | [optional] [default to null]
**DataTransferred** | **int64** | The number of bytes transferred from the source to the target as part of the replication process. Measured in bytes. | [optional] [default to null]
**PhysicalBytesWritten** | **int64** | The amount of physical/logical data written to the target due to replication. Measured in bytes. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

