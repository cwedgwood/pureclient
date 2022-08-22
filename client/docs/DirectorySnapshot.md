# DirectorySnapshot

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource. | [optional] [default to null]
**Name** | **string** | A user-specified name. The name must be locally unique and can be changed. | [optional] [default to null]
**ClientName** | **string** | The customizable portion of the client visible snapshot name. A full snapshot name is constructed in the form of &#x60;DIR.CLIENT_NAME.SUFFIX&#x60; where &#x60;DIR&#x60; is the full managed directory name, &#x60;CLIENT_NAME&#x60; is the client name, and &#x60;SUFFIX&#x60; is the suffix. The client visible snapshot name is &#x60;CLIENT_NAME.SUFFIX&#x60;. | [optional] [default to null]
**Created** | **int64** | The snapshot creation time, measured in milliseconds since the UNIX epoch. | [optional] [default to null]
**Destroyed** | **bool** | Returns a value of &#x60;true&#x60; if the snapshot has been destroyed and is pending eradication. The &#x60;time_remaining&#x60; value displays the amount of time left until the destroyed directory snapshot is permanently eradicated. Before the &#x60;time_remaining&#x60; period has elapsed, the destroyed directory snapshot can be recovered by setting &#x60;destroyed&#x3D;false&#x60;. Once the &#x60;time_remaining&#x60; period has elapsed, the directory snapshot is permanently eradicated and can no longer be recovered. | [optional] [default to null]
**Policy** | [***DirectorySnapshotPolicy**](DirectorySnapshot_policy.md) |  | [optional] [default to null]
**Source** | [***DirectorySnapshotSource**](DirectorySnapshot_source.md) |  | [optional] [default to null]
**Space** | [***DirectorySpaceSpace**](DirectorySpace_space.md) |  | [optional] [default to null]
**Suffix** | **int64** | The suffix that is appended to the &#x60;source_name&#x60; value and the &#x60;client_name&#x60; value to generate the full directory snapshot name in the form of &#x60;DIR.CLIENT_NAME.SUFFIX&#x60; where &#x60;DIR&#x60; is the managed directory name, &#x60;CLIENT_NAME&#x60; is the client name, and &#x60;SUFFIX&#x60; is the suffix. If the suffix is a string, this field returns &#x60;null&#x60;. See the &#x60;name&#x60; value for the full snapshot name including the suffix. | [optional] [default to null]
**TimeRemaining** | **int64** | The amount of time left until the directory snapshot is permanently eradicated. Measured in milliseconds. Before the &#x60;time_remaining&#x60; period has elapsed, the snapshot can be recovered by setting &#x60;destroyed&#x3D;false&#x60; if it is destroyed, by setting &#x60;policy&#x3D;\&quot;\&quot;&#x60; if it is managed by a snapshot policy, or by setting &#x60;keep_for&#x3D;\&quot;\&quot;&#x60; if it is a manual snapshot. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

