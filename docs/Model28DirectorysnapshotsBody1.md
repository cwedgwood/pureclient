# Model28DirectorysnapshotsBody1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destroyed** | **bool** | If set to &#x60;true&#x60;, destroys a resource. Once set to &#x60;true&#x60;, the &#x60;time_remaining&#x60; value will display the amount of time left until the destroyed resource is permanently eradicated. Before the &#x60;time_remaining&#x60; period has elapsed, the destroyed resource can be recovered by setting &#x60;destroyed&#x3D;false&#x60;. Once the &#x60;time_remaining&#x60; period has elapsed, the resource is permanently eradicated and can no longer be recovered. | [optional] [default to null]
**KeepFor** | **int64** | The amount of time to keep the snapshots, in milliseconds. Can only be set on snapshots that are not managed by any snapshot policy. Set to &#x60;\&quot;\&quot;&#x60; to clear the keep_for value. | [optional] [default to null]
**Policy** | [***Api28directorysnapshotsPolicy**](api2.8directorysnapshots_policy.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

