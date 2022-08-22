# VolumeGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource. | [optional] [default to null]
**Name** | **string** | A user-specified name. The name must be locally unique and can be changed. | [optional] [default to null]
**Destroyed** | **bool** | Returns a value of &#x60;true&#x60; if the volume group has been destroyed and is pending eradication. The &#x60;time_remaining&#x60; value displays the amount of time left until the destroyed volume group is permanently eradicated. Before the &#x60;time_remaining&#x60; period has elapsed, the destroyed volume group can be recovered by setting &#x60;destroyed&#x3D;false&#x60;. Once the &#x60;time_remaining&#x60; period has elapsed, the volume group is permanently eradicated and can no longer be recovered. | [optional] [default to null]
**Qos** | [***VolumeGroupQos**](VolumeGroup_qos.md) |  | [optional] [default to null]
**Space** | [***OffloadSpace**](Offload_space.md) |  | [optional] [default to null]
**TimeRemaining** | **int64** | The amount of time left until the destroyed volume group is permanently eradicated. Measured in milliseconds. Before the &#x60;time_remaining&#x60; period has elapsed, the destroyed volume group can be recovered by setting &#x60;destroyed&#x3D;false&#x60;. | [optional] [default to null]
**VolumeCount** | **int64** | The number of volumes in the volume group. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

