# RemoteProtectionGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource. | [optional] [default to null]
**Name** | **string** | A locally unique, system-generated name. The name cannot be modified. | [optional] [default to null]
**Destroyed** | **bool** | Returns a value of &#x60;true&#x60; if the remote protection group has been destroyed and is pending eradication. The &#x60;time_remaining&#x60; value displays the amount of time left until the destroyed remote protection group is permanently eradicated. Before the &#x60;time_remaining&#x60; period has elapsed, the destroyed remote protection group can be recovered by setting &#x60;destroyed&#x3D;false&#x60;. Once the &#x60;time_remaining&#x60; period has elapsed, the remote protection group is permanently eradicated and can no longer be recovered. | [optional] [default to null]
**IsLocal** | **bool** | If set to &#x60;true&#x60;, the location reference is to the local array. If set to &#x60;false&#x60;, the location reference is to a remote location, such as a remote array or offload target. | [optional] [default to null]
**Remote** | [***RemoteProtectionGroupRemote**](RemoteProtectionGroup_remote.md) |  | [optional] [default to null]
**Source** | [***RemoteProtectionGroupSource**](RemoteProtectionGroup_source.md) |  | [optional] [default to null]
**TargetRetention** | [***ProtectionGroupSourceRetention**](ProtectionGroup_source_retention.md) |  | [optional] [default to null]
**TimeRemaining** | **int64** | Milliseconds remaining until eradication, if remote protection group has been destroyed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

